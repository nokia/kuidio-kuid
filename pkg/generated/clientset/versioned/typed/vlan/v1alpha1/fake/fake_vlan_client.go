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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kuidio/kuid/pkg/generated/clientset/versioned/typed/vlan/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeVlanV1alpha1 struct {
	*testing.Fake
}

func (c *FakeVlanV1alpha1) VLANClaims(namespace string) v1alpha1.VLANClaimInterface {
	return &FakeVLANClaims{c, namespace}
}

func (c *FakeVlanV1alpha1) VLANEntries(namespace string) v1alpha1.VLANEntryInterface {
	return &FakeVLANEntries{c, namespace}
}

func (c *FakeVlanV1alpha1) VLANIndexes(namespace string) v1alpha1.VLANIndexInterface {
	return &FakeVLANIndexes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeVlanV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
