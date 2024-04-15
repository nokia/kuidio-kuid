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

	ipambev1alpha1 "github.com/kuidio/kuid/apis/backend/ipam/v1alpha1"
	resourcev1alpha1 "github.com/kuidio/kuid/apis/common/v1alpha1"
	conditionv1alpha1 "github.com/kuidio/kuid/apis/condition/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IPPrefixSpec defines the desired state of IPPrefix
type IPPrefixSpec struct {
	// Kind defines the kind of prefix for the IP Prefix
	// - network kind is used for physical, virtual nics on a device
	// - loopback kind is used for loopback interfaces
	// - pool kind is used for pools for dhcp/radius/bng/upf/etc
	// - aggregate kind is used for claiming an aggregate prefix
	// +kubebuilder:validation:Enum=`network`;`loopback`;`pool`;`aggregate`
	// +kubebuilder:default=network
	Kind ipambev1alpha1.PrefixKind `json:"kind" yaml:"kind" protobuf:"bytes,1,opt,name=kind"`
	// NetworkInstance defines the networkInstance context for the IP prefix
	// Name and optionally Namespace is used here
	NetworkInstance string `json:"networkInstance" yaml:"networkInstance" protobuf:"bytes,2,opt,name=networkInstance"`
	// Prefix defines the ip cidr in prefix or address notation.
	// +kubebuilder:validation:Pattern=`(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])/(([0-9])|([1-2][0-9])|(3[0-2]))|((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(/(([0-9])|([0-9]{2})|(1[0-1][0-9])|(12[0-8])))`
	Prefix string `json:"prefix" yaml:"prefix" protobuf:"bytes,3,opt,name=prefix"`
	// Gateway defines if the prefix/address is a gateway
	// +kubebuilder:validation:Optional
	// +optional
	Gateway *bool `json:"gateway,omitempty" yaml:"gateway,omitempty" protobuf:"varint,4,opt,name=gateway"`
	// +optional
	// UserDefinedLabels define metadata to the resource.
	// defined in the spec to distingiush metadata labels from user defined labels
	resourcev1alpha1.UserDefinedLabels `json:",inline" yaml:",inline" protobuf:"bytes,5,opt,name=userDefinedLabels"`
}

// IPPrefixStatus defines the observed state of IPPrefix
type IPPrefixStatus struct {
	// ConditionedStatus provides the status of the IPPrefix using conditions
	// - a ready condition indicates the overall status of the resource
	conditionv1alpha1.ConditionedStatus `json:",inline" yaml:",inline" protobuf:"bytes,1,opt,name=conditionedStatus"`
	// Prefix defines the prefix, claimed through the IPAM backend
	// +kubebuilder:validation:Optional
	Prefix *string `json:"prefix,omitempty" yaml:"prefix,omitempty" protobuf:"bytes,2,opt,name=prefix"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="NETWORK-INSTANCE",type="string",JSONPath=".spec.networkInstance"
// +kubebuilder:printcolumn:name="KIND",type="string",JSONPath=".spec.kind"
// +kubebuilder:printcolumn:name="SUBNET",type="string",JSONPath=".spec.subnetName"
// +kubebuilder:printcolumn:name="PREFIX-REQ",type="string",JSONPath=".spec.prefix"
// +kubebuilder:printcolumn:name="PREFIX-ALLOC",type="string",JSONPath=".status.prefix"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:categories={kuid,ipam}

// IPPrefix is the Schema for the ipprefixes API
type IPPrefix struct {
	metav1.TypeMeta   `json:",inline" yaml:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   IPPrefixSpec   `json:"spec,omitempty" yaml:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status IPPrefixStatus `json:"status,omitempty" yaml:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

//+kubebuilder:object:root=true

// IPPrefixList contains a list of IPPrefix
type IPPrefixList struct {
	metav1.TypeMeta `json:",inline" yaml:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" yaml:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []IPPrefix `json:"items" yaml:"items" protobuf:"bytes,2,rep,name=items"`
}

var (
	IPPrefixKind = reflect.TypeOf(IPPrefix{}).Name()
)
