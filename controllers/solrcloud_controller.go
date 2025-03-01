/*
Copyright 2019 Bloomberg Finance LP.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	solr "github.com/bloomberg/solr-operator/api/v1beta1"
	solrv1beta1 "github.com/bloomberg/solr-operator/api/v1beta1"
	"github.com/bloomberg/solr-operator/controllers/util"
	etcd "github.com/coreos/etcd-operator/pkg/apis/etcd/v1beta2"
	"github.com/go-logr/logr"
	zk "github.com/pravega/zookeeper-operator/pkg/apis/zookeeper/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	extv1 "k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"strings"
)

// SolrCloudReconciler reconciles a SolrCloud object
type SolrCloudReconciler struct {
	client.Client
	scheme *runtime.Scheme
	Log    logr.Logger
}

var useZkCRD bool
var useEtcdCRD bool
var IngressBaseUrl string

func UseZkCRD(useCRD bool) {
	useZkCRD = useCRD
}

func UseEtcdCRD(useCRD bool) {
	useEtcdCRD = useCRD
}

func SetIngressBaseUrl(ingressBaseUrl string) {
	IngressBaseUrl = ingressBaseUrl
}

// +kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch
// +kubebuilder:rbac:groups="",resources=pods/status,verbs=get
// +kubebuilder:rbac:groups="",resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=services/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=apps,resources=statefulsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apps,resources=statefulsets/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=extensions,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=extensions,resources=ingresses/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups="",resources=configmaps/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=solr.bloomberg.com,resources=solrclouds,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=solr.bloomberg.com,resources=solrclouds/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=zookeeper.pravega.io,resources=zookeeperclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=zookeeper.pravega.io,resources=zookeeperclusters/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=etcd.database.coreos.com,resources=etcdclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=etcd.database.coreos.com,resources=etcdclusters/status,verbs=get;update;patch

func (r *SolrCloudReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("solrcloud", req.NamespacedName)

	// Fetch the SolrCloud instance
	instance := &solr.SolrCloud{}
	err := r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the req.
		return reconcile.Result{}, err
	}

	changed := instance.WithDefaults()
	if changed {
		r.Log.Info("Setting default settings for solr-cloud", "namespace", instance.Namespace, "name", instance.Name)
		if err := r.Update(context.TODO(), instance); err != nil {
			return reconcile.Result{}, err
		}
		return reconcile.Result{Requeue: true}, nil
	}

	newStatus := solr.SolrCloudStatus{}

	busyBoxImage := *instance.Spec.BusyBoxImage

	if err := reconcileZk(r, req, instance, busyBoxImage, &newStatus); err != nil {
		return reconcile.Result{}, err
	}

	// Generate Service
	service := util.GenerateService(instance)
	if err := controllerutil.SetControllerReference(instance, service, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the Service already exists
	foundService := &corev1.Service{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		r.Log.Info("Creating Service", "namespace", service.Namespace, "name", service.Name)
		err = r.Create(context.TODO(), service)
	} else if err == nil {
		if util.CopyServiceFields(service, foundService) {
			// Update the found Service and write the result back if there are any changes
			r.Log.Info("Updating Service", "namespace", service.Namespace, "name", service.Name)
			err = r.Update(context.TODO(), foundService)
		}
		newStatus.InternalCommonAddress = "http://" + foundService.Name + "." + foundService.Namespace
	} else {
		return reconcile.Result{}, err
	}

	solrNodeNames := instance.GetAllSolrNodeNames()

	hostNameIpMap := make(map[string]string)
	// Generate a service for every Node
	for _, nodeName := range solrNodeNames {
		err, ip := reconcileNodeService(r, instance, nodeName)
		if err != nil {
			return reconcile.Result{}, err
		}
		if IngressBaseUrl != "" && ip != "" {
			hostNameIpMap[instance.NodeIngressUrl(nodeName, IngressBaseUrl)] = ip
		}
	}

	// Generate HeadlessService
	headless := util.GenerateHeadlessService(instance)
	if err := controllerutil.SetControllerReference(instance, headless, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the HeadlessService already exists
	foundHeadless := &corev1.Service{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: headless.Name, Namespace: headless.Namespace}, foundHeadless)
	if err != nil && errors.IsNotFound(err) {
		r.Log.Info("Creating HeadlessService", "namespace", headless.Namespace, "name", headless.Name)
		err = r.Create(context.TODO(), headless)
	} else if err == nil && util.CopyServiceFields(headless, foundHeadless) {
		// Update the found HeadlessService and write the result back if there are any changes
		r.Log.Info("Updating HeadlessService", "namespace", headless.Namespace, "name", headless.Name)
		err = r.Update(context.TODO(), foundHeadless)
	}
	if err != nil {
		return reconcile.Result{}, err
	}

	// Generate ConfigMap
	configMap := util.GenerateConfigMap(instance)
	if err := controllerutil.SetControllerReference(instance, configMap, r.scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the ConfigMap already exists
	foundConfigMap := &corev1.ConfigMap{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: configMap.Name, Namespace: configMap.Namespace}, foundConfigMap)
	if err != nil && errors.IsNotFound(err) {
		r.Log.Info("Creating ConfigMap", "namespace", configMap.Namespace, "name", configMap.Name)
		err = r.Create(context.TODO(), configMap)
	} else if err == nil && util.CopyConfigMapFields(configMap, foundConfigMap) {
		// Update the found ConfigMap and write the result back if there are any changes
		r.Log.Info("Updating ConfigMap", "namespace", configMap.Namespace, "name", configMap.Name)
		err = r.Update(context.TODO(), foundConfigMap)
	}
	if err != nil {
		return reconcile.Result{}, err
	}

	// Only create stateful set if zkConnectionString can be found (must contain host and port)
	if strings.Contains(newStatus.ZkConnectionString(), ":") {
		// Generate StatefulSet
		statefulSet := util.GenerateStatefulSet(instance, IngressBaseUrl, hostNameIpMap)
		if err := controllerutil.SetControllerReference(instance, statefulSet, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		// Check if the StatefulSet already exists
		foundStatefulSet := &appsv1.StatefulSet{}
		err = r.Get(context.TODO(), types.NamespacedName{Name: statefulSet.Name, Namespace: statefulSet.Namespace}, foundStatefulSet)
		if err != nil && errors.IsNotFound(err) {
			r.Log.Info("Creating StatefulSet", "namespace", statefulSet.Namespace, "name", statefulSet.Name)
			err = r.Create(context.TODO(), statefulSet)
		} else if err == nil {
			if util.CopyStatefulSetFields(statefulSet, foundStatefulSet) {
				// Update the found StatefulSet and write the result back if there are any changes
				r.Log.Info("Updating StatefulSet", "namespace", statefulSet.Namespace, "name", statefulSet.Name)
				err = r.Update(context.TODO(), foundStatefulSet)
			}
			newStatus.Replicas = foundStatefulSet.Status.Replicas
			newStatus.ReadyReplicas = foundStatefulSet.Status.ReadyReplicas
		}
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	err = reconcileCloudStatus(r, instance, &newStatus)
	if err != nil {
		return reconcile.Result{}, err
	}

	if IngressBaseUrl != "" {
		// Generate Ingress
		ingress := util.GenerateCommonIngress(instance, solrNodeNames, IngressBaseUrl)
		if err := controllerutil.SetControllerReference(instance, ingress, r.scheme); err != nil {
			return reconcile.Result{}, err
		}

		// Check if the Ingress already exists
		foundIngress := &extv1.Ingress{}
		err = r.Get(context.TODO(), types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, foundIngress)
		if err != nil && errors.IsNotFound(err) {
			r.Log.Info("Creating Common Ingress", "namespace", ingress.Namespace, "name", ingress.Name)
			err = r.Create(context.TODO(), ingress)
		} else if err == nil && util.CopyIngressFields(ingress, foundIngress) {
			// Update the found Ingress and write the result back if there are any changes
			r.Log.Info("Updating Common Ingress", "namespace", ingress.Namespace, "name", ingress.Name)
			err = r.Update(context.TODO(), foundIngress)
		}
		if err != nil {
			return reconcile.Result{}, err
		} else {
			address := "http://" + instance.CommonIngressUrl(IngressBaseUrl)
			newStatus.ExternalCommonAddress = &address
		}
	}

	if !reflect.DeepEqual(instance.Status, newStatus) {
		instance.Status = newStatus
		r.Log.Info("Updating SolrCloud Status: ", "namespace", instance.Namespace, "name", instance.Name)
		err = r.Status().Update(context.Background(), instance)
		if err != nil {
			return reconcile.Result{}, err
		}
	}

	return reconcile.Result{}, nil
}

func reconcileCloudStatus(r *SolrCloudReconciler, solrCloud *solr.SolrCloud, newStatus *solr.SolrCloudStatus) (err error) {
	foundPods := &corev1.PodList{}
	selectorLabels := solrCloud.SharedLabels()
	selectorLabels["technology"] = solr.SolrTechnologyLabel

	labelSelector := labels.SelectorFromSet(selectorLabels)
	listOps := &client.ListOptions{
		Namespace:     solrCloud.Namespace,
		LabelSelector: labelSelector,
	}

	err = r.List(context.TODO(), foundPods, listOps)
	if err != nil {
		return err
	}

	otherVersions := []string{}
	nodeStatuses := []solr.SolrNodeStatus{}
	backupRestoreReadyPods := 0
	for _, p := range foundPods.Items {
		nodeStatus := solr.SolrNodeStatus{}
		nodeStatus.NodeName = p.Name
		nodeStatus.InternalAddress = fmt.Sprintf("http://%s.%s.svc.cluster.local", p.Name, p.Namespace)
		if IngressBaseUrl != "" {
			nodeStatus.ExternalAddress = "http://" + solrCloud.NodeIngressUrl(nodeStatus.NodeName, IngressBaseUrl)
		}
		ready := false
		if len(p.Status.ContainerStatuses) > 0 {
			ready = true
			for _, c := range p.Status.ContainerStatuses {
				ready = ready && c.Ready
			}

			// The first container should always be running solr
			nodeStatus.Version = solr.ImageVersion(p.Spec.Containers[0].Image)
			if nodeStatus.Version != solrCloud.Spec.SolrImage.Tag {
				otherVersions = append(otherVersions, nodeStatus.Version)
			}
		}
		nodeStatus.Ready = ready

		nodeStatuses = append(nodeStatuses, nodeStatus)

		// Get Volumes for backup/restore
		if solrCloud.Spec.BackupRestoreVolume != nil {
			for _, volume := range p.Spec.Volumes {
				if volume.Name == util.BackupRestoreVolume {
					backupRestoreReadyPods += 1
				}
			}
		}
	}
	newStatus.SolrNodes = nodeStatuses

	if backupRestoreReadyPods == int(*solrCloud.Spec.Replicas) && backupRestoreReadyPods > 0 {
		newStatus.BackupRestoreReady = true
	}

	// If there are multiple versions of solr running, use the first otherVersion as the current running solr version of the cloud
	if len(otherVersions) > 0 {
		newStatus.TargetVersion = solrCloud.Spec.SolrImage.Tag
		newStatus.Version = otherVersions[0]
	} else {
		newStatus.TargetVersion = ""
		newStatus.Version = solrCloud.Spec.SolrImage.Tag
	}

	return nil
}

func reconcileNodeService(r *SolrCloudReconciler, instance *solr.SolrCloud, nodeName string) (err error, ip string) {
	// Generate Node Service
	service := util.GenerateNodeService(instance, nodeName)
	if err := controllerutil.SetControllerReference(instance, service, r.scheme); err != nil {
		return err, ip
	}

	// Check if the Ingress already exists
	foundService := &corev1.Service{}
	err = r.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundService)
	if err != nil && errors.IsNotFound(err) {
		r.Log.Info("Creating Node Service", "namespace", service.Namespace, "name", service.Name)
		err = r.Create(context.TODO(), service)
	} else if err == nil {
		if util.CopyServiceFields(service, foundService) {
			// Update the found Ingress and write the result back if there are any changes
			r.Log.Info("Updating Node Service", "namespace", service.Namespace, "name", service.Name)
			err = r.Update(context.TODO(), foundService)
		}
		ip = foundService.Spec.ClusterIP
	}
	if err != nil {
		return err, ip
	}

	return nil, ip
}

func reconcileZk(r *SolrCloudReconciler, request reconcile.Request, instance *solr.SolrCloud, busyBoxImage solr.ContainerImage, newStatus *solr.SolrCloudStatus) error {
	zkRef := instance.Spec.ZookeeperRef

	if zkRef.ConnectionInfo != nil {
		newStatus.ZookeeperConnectionInfo = *zkRef.ConnectionInfo
	} else {
		pzk := zkRef.ProvidedZookeeper
		if pzk == nil {
			return errors.NewBadRequest("No Zookeeper reference information provided.")
		}
		if pzk.Zetcd != nil {
			if !useEtcdCRD {
				return errors.NewBadRequest("Cannot create an Etcd Cluster, as the Solr Operator is not configured to use the Etcd CRD")
			}
			// Generate EtcdCluster
			etcdCluster := util.GenerateEtcdCluster(instance, *pzk.Zetcd.EtcdSpec, busyBoxImage)
			if err := controllerutil.SetControllerReference(instance, etcdCluster, r.scheme); err != nil {
				return err
			}

			// Check if the EtcdCluster already exists
			foundEtcdCluster := &etcd.EtcdCluster{}
			err := r.Get(context.TODO(), types.NamespacedName{Name: etcdCluster.Name, Namespace: etcdCluster.Namespace}, foundEtcdCluster)
			if err != nil && errors.IsNotFound(err) {
				r.Log.Info("Creating EtcdCluster", "namespace", etcdCluster.Namespace, "name", etcdCluster.Name)
				err = r.Create(context.TODO(), etcdCluster)
			} else if err == nil && util.CopyEtcdClusterFields(etcdCluster, foundEtcdCluster) {
				// Update the found EtcdCluster and write the result back if there are any changes
				r.Log.Info("Updating EtcdCluster", "namespace", etcdCluster.Namespace, "name", etcdCluster.Name)
				err = r.Update(context.TODO(), foundEtcdCluster)
			}
			if err != nil {
				return err
			}

			// Generate Zetcd Deployment
			deployment := util.GenerateZetcdDeployment(instance, *pzk.Zetcd.ZetcdSpec)
			if err := controllerutil.SetControllerReference(instance, deployment, r.scheme); err != nil {
				return err
			}

			// Check if the Zetcd Deployment already exists
			foundDeployment := &appsv1.Deployment{}
			err = r.Get(context.TODO(), types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, foundDeployment)
			if err != nil && errors.IsNotFound(err) {
				r.Log.Info("Creating Zetcd Deployment", "namespace", deployment.Namespace, "name", deployment.Name)
				err = r.Create(context.TODO(), foundDeployment)
			} else if err == nil && util.CopyDeploymentFields(deployment, foundDeployment) {
				// Update the found Zetcd Deployment and write the result back if there are any changes
				r.Log.Info("Updating Zetcd Deployment", "namespace", deployment.Namespace, "name", deployment.Name)
				err = r.Update(context.TODO(), foundDeployment)
			}
			if err != nil {
				return err
			}

			// Generate Zetcd Service
			service := util.GenerateZetcdService(instance, *pzk.Zetcd.ZetcdSpec)
			if err := controllerutil.SetControllerReference(instance, service, r.scheme); err != nil {
				return err
			}

			// Check if the Zetcd Service already exists
			foundService := &corev1.Service{}
			err = r.Get(context.TODO(), types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, foundService)
			if err != nil && errors.IsNotFound(err) {
				r.Log.Info("Creating Zetcd Service", "namespace", service.Namespace, "name", service.Name)
				err = r.Create(context.TODO(), service)
				newStatus.ZookeeperConnectionInfo = solr.ZookeeperConnectionInfo{
					InternalConnectionString: service.Name + "." + service.Namespace + ":2181",
					ChRoot:                   "/",
				}
			} else if err == nil {
				if util.CopyServiceFields(service, foundService) {
					// Update the found Zetcd Service and write the result back if there are any changes
					r.Log.Info("Updating Zetcd Service", "namespace", service.Namespace, "name", service.Name)
					err = r.Update(context.TODO(), foundService)
				}
				newStatus.ZookeeperConnectionInfo = solr.ZookeeperConnectionInfo{
					InternalConnectionString: service.Name + "." + service.Namespace + ":2181",
					ChRoot:                   "/",
				}
			} else {
				return err
			}

		} else if pzk.Zookeeper != nil {
			// Generate ZookeeperCluster
			if !useZkCRD {
				return errors.NewBadRequest("Cannot create a Zookeeper Cluster, as the Solr Operator is not configured to use the Zookeeper CRD")
			}
			zkCluster := util.GenerateZookeeperCluster(instance, *pzk.Zookeeper)
			if err := controllerutil.SetControllerReference(instance, zkCluster, r.scheme); err != nil {
				return err
			}

			// Check if the ZookeeperCluster already exists
			foundZkCluster := &zk.ZookeeperCluster{}
			err := r.Get(context.TODO(), types.NamespacedName{Name: zkCluster.Name, Namespace: zkCluster.Namespace}, foundZkCluster)
			if err != nil && errors.IsNotFound(err) {
				r.Log.Info("Creating Zookeeer Cluster", "namespace", zkCluster.Namespace, "name", zkCluster.Name)
				err = r.Create(context.TODO(), zkCluster)
			} else if err == nil {
				if util.CopyZookeeperClusterFields(zkCluster, foundZkCluster) {
					// Update the found ZookeeperCluster and write the result back if there are any changes
					r.Log.Info("Updating Zookeeer Cluster", "namespace", zkCluster.Namespace, "name", zkCluster.Name)
					err = r.Update(context.TODO(), foundZkCluster)
				}
				external := &foundZkCluster.Status.ExternalClientEndpoint
				if "" == *external {
					external = nil
				}
				newStatus.ZookeeperConnectionInfo = solr.ZookeeperConnectionInfo{
					InternalConnectionString: foundZkCluster.Status.InternalClientEndpoint,
					ExternalConnectionString: external,
					ChRoot:                   "/",
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func (r *SolrCloudReconciler) SetupWithManager(mgr ctrl.Manager) error {
	ctrlBuilder := ctrl.NewControllerManagedBy(mgr).
		For(&solrv1beta1.SolrCloud{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&appsv1.StatefulSet{}).
		Owns(&corev1.Service{}).
		Owns(&extv1.Ingress{}).
		Owns(&appsv1.Deployment{})
	if useZkCRD {
		ctrlBuilder = ctrlBuilder.Owns(&zk.ZookeeperCluster{})
	} else if useEtcdCRD {
		ctrlBuilder = ctrlBuilder.Owns(&etcd.EtcdCluster{})
	}
	r.scheme = mgr.GetScheme()
	return ctrlBuilder.Complete(r)
}
