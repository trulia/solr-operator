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
	"reflect"
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	solrv1beta1 "github.com/bloomberg/solr-operator/api/v1beta1"
	"github.com/bloomberg/solr-operator/controllers/util"
)

// SolrShardReconciler reconciles a SolrShard object
type SolrShardReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=solr.bloomberg.com,resources=solrshards,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=solr.bloomberg.com,resources=solrshards/status,verbs=get;update;patch

func (r *SolrShardReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("solrshard", req.NamespacedName)

	shard := &solrv1beta1.SolrShard{}
	shardFinalizer := "shard.finalizers.bloomberg.com"

	err := r.Get(context.TODO(), req.NamespacedName, shard)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the req.
		return reconcile.Result{}, err
	}

	oldStatus := shard.Status.DeepCopy()
	requeueOrNot := reconcile.Result{Requeue: true, RequeueAfter: time.Second * 5}

	shardCreationStatus := reconcileSolrShard(r, shard, shard.Spec.SolrCloud, shard.Name, shard.Spec.CreateNodeSet, shard.Spec.Collection, shard.Namespace)

	if err != nil {
		r.Log.Error(err, "Error while creating SolrCloud alias")
	}

	if shard.ObjectMeta.DeletionTimestamp.IsZero() {
		// The object is not being deleted, so if it does not have our finalizer,
		// then lets add the finalizer and update the object
		if !util.ContainsString(shard.ObjectMeta.Finalizers, shardFinalizer) {
			shard.ObjectMeta.Finalizers = append(shard.ObjectMeta.Finalizers, shardFinalizer)
			if err := r.Update(context.Background(), shard); err != nil {
				return reconcile.Result{}, err
			}
		}
	} else {
		// The object is being deleted, get associated SolrCloud
		solrCloud := &solrv1beta1.SolrCloud{}
		err = r.Get(context.TODO(), types.NamespacedName{Namespace: shard.Namespace, Name: shard.Spec.SolrCloud}, solrCloud)

		if util.ContainsString(shard.ObjectMeta.Finalizers, shardFinalizer) && solrCloud != nil && shardCreationStatus {
			r.Log.Info("Deleting Solr collection shard", "cloud", shard.Spec.SolrCloud, "namespace", shard.Namespace, "Collection Name", shard.Spec.Collection)
			// our finalizer is present, along with the associated SolrCloud and shard lets delete shard
			if util.CheckIfShardExists(shard.Spec.SolrCloud, shard.Spec.Collection, shard.Name, shard.Namespace) {
				delete, err := util.DeleteShard(shard.Spec.SolrCloud, shard.Spec.Collection, shard.Name, shard.Namespace)
				if err != nil {
					r.Log.Error(err, "Failed to delete Solr shard")
					return reconcile.Result{}, err
				}
				r.Log.Info("Deleted Solr collection", "cloud", shard.Spec.SolrCloud, "namespace", shard.Namespace, "Alias", shard.Name, "Deleted", delete)
			}
		}

		// remove our finalizer from the list and update it.
		shard.ObjectMeta.Finalizers = util.RemoveString(shard.ObjectMeta.Finalizers, shardFinalizer)
		if err := r.Update(context.Background(), shard); err != nil {
			return reconcile.Result{}, err
		}
	}

	if shard.Status.CreatedTime == nil {
		now := metav1.Now()
		shard.Status.CreatedTime = &now
		shard.Status.Created = shardCreationStatus
	}

	if !reflect.DeepEqual(oldStatus, shard.Status) {
		r.Log.Info("Updating status for collection shard", "shard", shard, "namespace", shard.Namespace, "name", shard.Name)
		err = r.Status().Update(context.TODO(), shard)
	}

	if shardCreationStatus {
		requeueOrNot = reconcile.Result{}
	}

	return requeueOrNot, nil
}

func reconcileSolrShard(r *SolrShardReconciler, shard *solrv1beta1.SolrShard, solrCloudName string, shardName string, createNodeSet []string, collection string, namespace string) (shardCreationStatus bool) {
	success := util.CheckIfShardExists(solrCloudName, collection, shardName, namespace)

	// If not created, or if alias status differs from spec requirements create alias
	if !success {
		r.Log.Info("Creating shard", "shard", shard, "collection", collection, "createNodeSet", createNodeSet)
		result, err := util.CreateShard(solrCloudName, collection, shard.Name, createNodeSet, namespace)
		if err == nil {
			shard.Status.Created = result
			shard.Status.NodeSet = createNodeSet
		}
	} else {
		shard.Status.Created = true
	}

	return shard.Status.Created
}

func (r *SolrShardReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&solrv1beta1.SolrShard{}).
		Complete(r)
}
