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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SolrShardSpec defines the desired state of SolrShard
type SolrShardSpec struct {
	// A reference to the SolrCloud to create the shard on
	SolrCloud string `json:"solrCloud"`

	// A reference to the collection to create the shard for
	Collection string `json:"collection"`

	// A referenec to the node set to create the shard on
	CreateNodeSet []string `json:"createNodeSet"`
}

// SolrShardStatus defines the observed state of SolrShard
type SolrShardStatus struct {
	// Status of the shard
	// +optional
	Created bool `json:"created,omitempty"`

	// +optional
	CreatedTime *metav1.Time `json:"createdTime,omitempty"`

	// Node set that the shard is for
	// +optional
	NodeSet []string `json:"nodeSet,omitempty"`
}

// +kubebuilder:object:root=true

// SolrShard is the Schema for the solrshards API
type SolrShard struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SolrShardSpec   `json:"spec,omitempty"`
	Status SolrShardStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SolrShardList contains a list of SolrShard
type SolrShardList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SolrShard `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SolrShard{}, &SolrShardList{})
}
