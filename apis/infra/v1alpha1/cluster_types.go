/*
Copyright 2024 Nokia.

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

package v1alpha1

import (
	"reflect"

	condv1alpha1 "github.com/kform-dev/choreo/apis/condition/v1alpha1"
	commonv1alpha1 "github.com/kuidio/kuid/apis/common/v1alpha1"
	idv1alpha1 "github.com/kuidio/kuid/apis/id/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	// PartitionClusterID defines the cluster partition
	idv1alpha1.PartitionClusterID `json:",inline" yaml:",inline" protobuf:"bytes,1,opt,name=nodeID"`
	// Provider defines the provider implementing this resource.
	Provider string `json:"provider" yaml:"provider" protobuf:"bytes,2,opt,name=provider"`
	// Location defines the location information where this resource is located
	// in lon/lat coordinates
	Location *Location `json:"location,omitempty" yaml:"location,omitempty" protobuf:"bytes,3,opt,name=location"`
	// UserDefinedLabels define metadata to the resource.
	// defined in the spec to distingiush metadata labels from user defined labels
	commonv1alpha1.UserDefinedLabels `json:",inline" yaml:",inline" protobuf:"bytes,4,opt,name=userDefinedLabels"`
	// ParametersRef points to a provider specific configuration of the resource
	// +optional
	//ParametersRef *ObjectReference `json:"parametersRef,omitempty" yaml:"parametersRef,omitempty" protobuf:"bytes,5,opt,name=parametersRef"`
}

// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
	// ConditionedStatus provides the status of the IPClain using conditions
	// - a ready condition indicates the overall status of the resource
	condv1alpha1.ConditionedStatus `json:",inline" yaml:",inline" protobuf:"bytes,1,opt,name=conditionedStatus"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories={kuid}
// A Cluster represents a kubernetes cluster and is typically used as a nodeGroup identifier.
type Cluster struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   ClusterSpec   `json:"spec,omitempty" yaml:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status ClusterStatus `json:"status,omitempty" yaml:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ClusterList contains a list of Clusters
type ClusterList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []Cluster `json:"items" yaml:"items" protobuf:"bytes,2,rep,name=items"`
}

var (
	ClusterKind     = reflect.TypeOf(Cluster{}).Name()
	ClusterKindList = reflect.TypeOf(ClusterList{}).Name()
)
