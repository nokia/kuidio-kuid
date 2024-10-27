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

package ipam

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BuildIPClaim returns a reource from a client Object a Spec/Status
func BuildIPIndex(meta metav1.ObjectMeta, spec *IPIndexSpec, status *IPIndexStatus) *IPIndex {
	aspec := IPIndexSpec{}
	if spec != nil {
		aspec = *spec
	}
	astatus := IPIndexStatus{}
	if status != nil {
		astatus = *status
	}
	return &IPIndex{
		TypeMeta: metav1.TypeMeta{
			APIVersion: SchemeGroupVersion.Identifier(),
			Kind:       IPIndexKind,
		},
		ObjectMeta: meta,
		Spec:       aspec,
		Status:     astatus,
	}
}
