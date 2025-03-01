// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupPersistenceStatus) DeepCopyInto(out *BackupPersistenceStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupPersistenceStatus.
func (in *BackupPersistenceStatus) DeepCopy() *BackupPersistenceStatus {
	if in == nil {
		return nil
	}
	out := new(BackupPersistenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CollectionBackupStatus) DeepCopyInto(out *CollectionBackupStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CollectionBackupStatus.
func (in *CollectionBackupStatus) DeepCopy() *CollectionBackupStatus {
	if in == nil {
		return nil
	}
	out := new(CollectionBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerImage) DeepCopyInto(out *ContainerImage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerImage.
func (in *ContainerImage) DeepCopy() *ContainerImage {
	if in == nil {
		return nil
	}
	out := new(ContainerImage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdPodPolicy) DeepCopyInto(out *EtcdPodPolicy) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdPodPolicy.
func (in *EtcdPodPolicy) DeepCopy() *EtcdPodPolicy {
	if in == nil {
		return nil
	}
	out := new(EtcdPodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdSpec) DeepCopyInto(out *EtcdSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	in.EtcdPod.DeepCopyInto(&out.EtcdPod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdSpec.
func (in *EtcdSpec) DeepCopy() *EtcdSpec {
	if in == nil {
		return nil
	}
	out := new(EtcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FullZetcdSpec) DeepCopyInto(out *FullZetcdSpec) {
	*out = *in
	if in.EtcdSpec != nil {
		in, out := &in.EtcdSpec, &out.EtcdSpec
		*out = new(EtcdSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ZetcdSpec != nil {
		in, out := &in.ZetcdSpec, &out.ZetcdSpec
		*out = new(ZetcdSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FullZetcdSpec.
func (in *FullZetcdSpec) DeepCopy() *FullZetcdSpec {
	if in == nil {
		return nil
	}
	out := new(FullZetcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistenceSource) DeepCopyInto(out *PersistenceSource) {
	*out = *in
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = new(S3PersistenceSource)
		(*in).DeepCopyInto(*out)
	}
	if in.Volume != nil {
		in, out := &in.Volume, &out.Volume
		*out = new(VolumePersistenceSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistenceSource.
func (in *PersistenceSource) DeepCopy() *PersistenceSource {
	if in == nil {
		return nil
	}
	out := new(PersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProvidedZookeeper) DeepCopyInto(out *ProvidedZookeeper) {
	*out = *in
	if in.Zookeeper != nil {
		in, out := &in.Zookeeper, &out.Zookeeper
		*out = new(ZookeeperSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Zetcd != nil {
		in, out := &in.Zetcd, &out.Zetcd
		*out = new(FullZetcdSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProvidedZookeeper.
func (in *ProvidedZookeeper) DeepCopy() *ProvidedZookeeper {
	if in == nil {
		return nil
	}
	out := new(ProvidedZookeeper)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3PersistenceSource) DeepCopyInto(out *S3PersistenceSource) {
	*out = *in
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(int32)
		**out = **in
	}
	out.Secrets = in.Secrets
	out.AWSCliImage = in.AWSCliImage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3PersistenceSource.
func (in *S3PersistenceSource) DeepCopy() *S3PersistenceSource {
	if in == nil {
		return nil
	}
	out := new(S3PersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Secrets) DeepCopyInto(out *S3Secrets) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Secrets.
func (in *S3Secrets) DeepCopy() *S3Secrets {
	if in == nil {
		return nil
	}
	out := new(S3Secrets)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackup) DeepCopyInto(out *SolrBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackup.
func (in *SolrBackup) DeepCopy() *SolrBackup {
	if in == nil {
		return nil
	}
	out := new(SolrBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupList) DeepCopyInto(out *SolrBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupList.
func (in *SolrBackupList) DeepCopy() *SolrBackupList {
	if in == nil {
		return nil
	}
	out := new(SolrBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupSpec) DeepCopyInto(out *SolrBackupSpec) {
	*out = *in
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Persistence.DeepCopyInto(&out.Persistence)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupSpec.
func (in *SolrBackupSpec) DeepCopy() *SolrBackupSpec {
	if in == nil {
		return nil
	}
	out := new(SolrBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrBackupStatus) DeepCopyInto(out *SolrBackupStatus) {
	*out = *in
	if in.CollectionBackupStatuses != nil {
		in, out := &in.CollectionBackupStatuses, &out.CollectionBackupStatuses
		*out = make([]CollectionBackupStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.PersistenceStatus.DeepCopyInto(&out.PersistenceStatus)
	if in.FinishTime != nil {
		in, out := &in.FinishTime, &out.FinishTime
		*out = (*in).DeepCopy()
	}
	if in.Successful != nil {
		in, out := &in.Successful, &out.Successful
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrBackupStatus.
func (in *SolrBackupStatus) DeepCopy() *SolrBackupStatus {
	if in == nil {
		return nil
	}
	out := new(SolrBackupStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloud) DeepCopyInto(out *SolrCloud) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloud.
func (in *SolrCloud) DeepCopy() *SolrCloud {
	if in == nil {
		return nil
	}
	out := new(SolrCloud)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCloud) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudList) DeepCopyInto(out *SolrCloudList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrCloud, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudList.
func (in *SolrCloudList) DeepCopy() *SolrCloudList {
	if in == nil {
		return nil
	}
	out := new(SolrCloudList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCloudList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudReference) DeepCopyInto(out *SolrCloudReference) {
	*out = *in
	if in.ZookeeperConnectionInfo != nil {
		in, out := &in.ZookeeperConnectionInfo, &out.ZookeeperConnectionInfo
		*out = new(ZookeeperConnectionInfo)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudReference.
func (in *SolrCloudReference) DeepCopy() *SolrCloudReference {
	if in == nil {
		return nil
	}
	out := new(SolrCloudReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudSpec) DeepCopyInto(out *SolrCloudSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.ZookeeperRef != nil {
		in, out := &in.ZookeeperRef, &out.ZookeeperRef
		*out = new(ZookeeperRef)
		(*in).DeepCopyInto(*out)
	}
	if in.SolrImage != nil {
		in, out := &in.SolrImage, &out.SolrImage
		*out = new(ContainerImage)
		**out = **in
	}
	in.SolrPod.DeepCopyInto(&out.SolrPod)
	if in.DataPvcSpec != nil {
		in, out := &in.DataPvcSpec, &out.DataPvcSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.BackupRestoreVolume != nil {
		in, out := &in.BackupRestoreVolume, &out.BackupRestoreVolume
		*out = new(v1.VolumeSource)
		(*in).DeepCopyInto(*out)
	}
	if in.BusyBoxImage != nil {
		in, out := &in.BusyBoxImage, &out.BusyBoxImage
		*out = new(ContainerImage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudSpec.
func (in *SolrCloudSpec) DeepCopy() *SolrCloudSpec {
	if in == nil {
		return nil
	}
	out := new(SolrCloudSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCloudStatus) DeepCopyInto(out *SolrCloudStatus) {
	*out = *in
	if in.SolrNodes != nil {
		in, out := &in.SolrNodes, &out.SolrNodes
		*out = make([]SolrNodeStatus, len(*in))
		copy(*out, *in)
	}
	if in.ExternalCommonAddress != nil {
		in, out := &in.ExternalCommonAddress, &out.ExternalCommonAddress
		*out = new(string)
		**out = **in
	}
	in.ZookeeperConnectionInfo.DeepCopyInto(&out.ZookeeperConnectionInfo)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCloudStatus.
func (in *SolrCloudStatus) DeepCopy() *SolrCloudStatus {
	if in == nil {
		return nil
	}
	out := new(SolrCloudStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollection) DeepCopyInto(out *SolrCollection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollection.
func (in *SolrCollection) DeepCopy() *SolrCollection {
	if in == nil {
		return nil
	}
	out := new(SolrCollection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCollection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionAlias) DeepCopyInto(out *SolrCollectionAlias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionAlias.
func (in *SolrCollectionAlias) DeepCopy() *SolrCollectionAlias {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionAlias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCollectionAlias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionAliasList) DeepCopyInto(out *SolrCollectionAliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrCollectionAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionAliasList.
func (in *SolrCollectionAliasList) DeepCopy() *SolrCollectionAliasList {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionAliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCollectionAliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionAliasSpec) DeepCopyInto(out *SolrCollectionAliasSpec) {
	*out = *in
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionAliasSpec.
func (in *SolrCollectionAliasSpec) DeepCopy() *SolrCollectionAliasSpec {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionAliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionAliasStatus) DeepCopyInto(out *SolrCollectionAliasStatus) {
	*out = *in
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = (*in).DeepCopy()
	}
	if in.Collections != nil {
		in, out := &in.Collections, &out.Collections
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionAliasStatus.
func (in *SolrCollectionAliasStatus) DeepCopy() *SolrCollectionAliasStatus {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionAliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionList) DeepCopyInto(out *SolrCollectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrCollection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionList.
func (in *SolrCollectionList) DeepCopy() *SolrCollectionList {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrCollectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionSpec) DeepCopyInto(out *SolrCollectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionSpec.
func (in *SolrCollectionSpec) DeepCopy() *SolrCollectionSpec {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrCollectionStatus) DeepCopyInto(out *SolrCollectionStatus) {
	*out = *in
	if in.CreatedTime != nil {
		in, out := &in.CreatedTime, &out.CreatedTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrCollectionStatus.
func (in *SolrCollectionStatus) DeepCopy() *SolrCollectionStatus {
	if in == nil {
		return nil
	}
	out := new(SolrCollectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrNodeStatus) DeepCopyInto(out *SolrNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrNodeStatus.
func (in *SolrNodeStatus) DeepCopy() *SolrNodeStatus {
	if in == nil {
		return nil
	}
	out := new(SolrNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPodPolicy) DeepCopyInto(out *SolrPodPolicy) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPodPolicy.
func (in *SolrPodPolicy) DeepCopy() *SolrPodPolicy {
	if in == nil {
		return nil
	}
	out := new(SolrPodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporter) DeepCopyInto(out *SolrPrometheusExporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporter.
func (in *SolrPrometheusExporter) DeepCopy() *SolrPrometheusExporter {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrPrometheusExporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporterList) DeepCopyInto(out *SolrPrometheusExporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SolrPrometheusExporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporterList.
func (in *SolrPrometheusExporterList) DeepCopy() *SolrPrometheusExporterList {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SolrPrometheusExporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporterSpec) DeepCopyInto(out *SolrPrometheusExporterSpec) {
	*out = *in
	in.SolrReference.DeepCopyInto(&out.SolrReference)
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	in.PodPolicy.DeepCopyInto(&out.PodPolicy)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporterSpec.
func (in *SolrPrometheusExporterSpec) DeepCopy() *SolrPrometheusExporterSpec {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrPrometheusExporterStatus) DeepCopyInto(out *SolrPrometheusExporterStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrPrometheusExporterStatus.
func (in *SolrPrometheusExporterStatus) DeepCopy() *SolrPrometheusExporterStatus {
	if in == nil {
		return nil
	}
	out := new(SolrPrometheusExporterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SolrReference) DeepCopyInto(out *SolrReference) {
	*out = *in
	if in.Cloud != nil {
		in, out := &in.Cloud, &out.Cloud
		*out = new(SolrCloudReference)
		(*in).DeepCopyInto(*out)
	}
	if in.Standalone != nil {
		in, out := &in.Standalone, &out.Standalone
		*out = new(StandaloneSolrReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SolrReference.
func (in *SolrReference) DeepCopy() *SolrReference {
	if in == nil {
		return nil
	}
	out := new(SolrReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StandaloneSolrReference) DeepCopyInto(out *StandaloneSolrReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StandaloneSolrReference.
func (in *StandaloneSolrReference) DeepCopy() *StandaloneSolrReference {
	if in == nil {
		return nil
	}
	out := new(StandaloneSolrReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumePersistenceSource) DeepCopyInto(out *VolumePersistenceSource) {
	*out = *in
	in.VolumeSource.DeepCopyInto(&out.VolumeSource)
	out.BusyBoxImage = in.BusyBoxImage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumePersistenceSource.
func (in *VolumePersistenceSource) DeepCopy() *VolumePersistenceSource {
	if in == nil {
		return nil
	}
	out := new(VolumePersistenceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZetcdPodPolicy) DeepCopyInto(out *ZetcdPodPolicy) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZetcdPodPolicy.
func (in *ZetcdPodPolicy) DeepCopy() *ZetcdPodPolicy {
	if in == nil {
		return nil
	}
	out := new(ZetcdPodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZetcdSpec) DeepCopyInto(out *ZetcdSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	in.ZetcdPod.DeepCopyInto(&out.ZetcdPod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZetcdSpec.
func (in *ZetcdSpec) DeepCopy() *ZetcdSpec {
	if in == nil {
		return nil
	}
	out := new(ZetcdSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperConnectionInfo) DeepCopyInto(out *ZookeeperConnectionInfo) {
	*out = *in
	if in.ExternalConnectionString != nil {
		in, out := &in.ExternalConnectionString, &out.ExternalConnectionString
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperConnectionInfo.
func (in *ZookeeperConnectionInfo) DeepCopy() *ZookeeperConnectionInfo {
	if in == nil {
		return nil
	}
	out := new(ZookeeperConnectionInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperPodPolicy) DeepCopyInto(out *ZookeeperPodPolicy) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperPodPolicy.
func (in *ZookeeperPodPolicy) DeepCopy() *ZookeeperPodPolicy {
	if in == nil {
		return nil
	}
	out := new(ZookeeperPodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperRef) DeepCopyInto(out *ZookeeperRef) {
	*out = *in
	if in.ConnectionInfo != nil {
		in, out := &in.ConnectionInfo, &out.ConnectionInfo
		*out = new(ZookeeperConnectionInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.ProvidedZookeeper != nil {
		in, out := &in.ProvidedZookeeper, &out.ProvidedZookeeper
		*out = new(ProvidedZookeeper)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperRef.
func (in *ZookeeperRef) DeepCopy() *ZookeeperRef {
	if in == nil {
		return nil
	}
	out := new(ZookeeperRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ZookeeperSpec) DeepCopyInto(out *ZookeeperSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ContainerImage)
		**out = **in
	}
	if in.PersistentVolumeClaimSpec != nil {
		in, out := &in.PersistentVolumeClaimSpec, &out.PersistentVolumeClaimSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
	in.ZookeeperPod.DeepCopyInto(&out.ZookeeperPod)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ZookeeperSpec.
func (in *ZookeeperSpec) DeepCopy() *ZookeeperSpec {
	if in == nil {
		return nil
	}
	out := new(ZookeeperSpec)
	in.DeepCopyInto(out)
	return out
}
