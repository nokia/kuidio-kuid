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

	commonv1alpha1 "github.com/kuidio/kuid/apis/common/v1alpha1"
	conditionv1alpha1 "github.com/kuidio/kuid/apis/condition/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VLANClaimSpec defines the desired state of VLANClaim
type VLANClaimSpec struct {
	// VLANIndex defines the vlan index for the VLAN Claim
	Index string `json:"index" yaml:"index" protobuf:"bytes,1,opt,name=index"`
	// VLANID defines the vlan for the VLAN claim
	// +optional
	ID *uint32 `json:"id,omitempty" yaml:"id,omitempty" protobuf:"bytes,2,opt,name=id"`
	// Range defines the vlan range for the VLAN claim
	// The following notation is used: start-end <start-vlanID>-<end-vlanID>
	// the vlans in the range must be consecutive
	// +optional
	Range *string `json:"range,omitempty" yaml:"range,omitempty" protobuf:"bytes,3,opt,name=range"`
	// ClaimLabels define the user defined labels and selector labels used
	// in resource claim
	commonv1alpha1.ClaimLabels `json:",inline" yaml:",inline" protobuf:"bytes,4,opt,name=claimLabels"`
	// Owner defines the ownerReference of the VLANClaim
	// Allow for different namesapces, hence it is part of the spec
	Owner *commonv1alpha1.OwnerReference `json:"owner,omitempty" yaml:"owner,omitempty" protobuf:"bytes,5,opt,name=owner"`
}

// VLANClaimStatus defines the observed state of VLANClaim
type VLANClaimStatus struct {
	// ConditionedStatus provides the status of the Claim using conditions
	// - a ready condition indicates the overall status of the resource
	conditionv1alpha1.ConditionedStatus `json:",inline" yaml:",inline" protobuf:"bytes,1,opt,name=conditionedStatus"`
	// VLANID defines the vlan for the VLAN claim
	// +optional
	ID *uint32 `json:"id,omitempty" yaml:"id,omitempty" protobuf:"bytes,2,opt,name=id"`
	// VLANRange defines the vlan range for the VLAN claim
	// +optional
	Range *string `json:"range,omitempty" yaml:"range,omitempty" protobuf:"bytes,3,opt,name=range"`
	// ExpiryTime defines when the claim expires
	// +kubebuilder:validation:Optional
	// +optional
	ExpiryTime *string `json:"expiryTime,omitempty" yaml:"expiryTime,omitempty" protobuf:"bytes,4,opt,name=expiryTime"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VLANClaim is the Schema for the VLANClaim API
//
// +k8s:openapi-gen=true
type VLANClaim struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   VLANClaimSpec   `json:"spec,omitempty" yaml:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status VLANClaimStatus `json:"status,omitempty" yaml:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// VLANClaimList contains a list of VLANClaims
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type VLANClaimList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []VLANClaim `json:"items" yaml:"items" protobuf:"bytes,2,rep,name=items"`
}

var (
	VLANClaimKind     = reflect.TypeOf(VLANClaim{}).Name()
	VLANClaimListKind = reflect.TypeOf(VLANClaimList{}).Name()
)
