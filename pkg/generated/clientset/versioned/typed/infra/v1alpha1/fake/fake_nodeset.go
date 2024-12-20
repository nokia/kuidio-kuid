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
	"context"

	v1alpha1 "github.com/kuidio/kuid/apis/infra/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeSets implements NodeSetInterface
type FakeNodeSets struct {
	Fake *FakeInfraV1alpha1
	ns   string
}

var nodesetsResource = v1alpha1.SchemeGroupVersion.WithResource("nodesets")

var nodesetsKind = v1alpha1.SchemeGroupVersion.WithKind("NodeSet")

// Get takes name of the nodeSet, and returns the corresponding nodeSet object, and an error if there is any.
func (c *FakeNodeSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.NodeSet, err error) {
	emptyResult := &v1alpha1.NodeSet{}
	obj, err := c.Fake.
		Invokes(testing.NewGetActionWithOptions(nodesetsResource, c.ns, name, options), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NodeSet), err
}

// List takes label and field selectors, and returns the list of NodeSets that match those selectors.
func (c *FakeNodeSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.NodeSetList, err error) {
	emptyResult := &v1alpha1.NodeSetList{}
	obj, err := c.Fake.
		Invokes(testing.NewListActionWithOptions(nodesetsResource, nodesetsKind, c.ns, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeSetList{ListMeta: obj.(*v1alpha1.NodeSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeSets.
func (c *FakeNodeSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchActionWithOptions(nodesetsResource, c.ns, opts))

}

// Create takes the representation of a nodeSet and creates it.  Returns the server's representation of the nodeSet, and an error, if there is any.
func (c *FakeNodeSets) Create(ctx context.Context, nodeSet *v1alpha1.NodeSet, opts v1.CreateOptions) (result *v1alpha1.NodeSet, err error) {
	emptyResult := &v1alpha1.NodeSet{}
	obj, err := c.Fake.
		Invokes(testing.NewCreateActionWithOptions(nodesetsResource, c.ns, nodeSet, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NodeSet), err
}

// Update takes the representation of a nodeSet and updates it. Returns the server's representation of the nodeSet, and an error, if there is any.
func (c *FakeNodeSets) Update(ctx context.Context, nodeSet *v1alpha1.NodeSet, opts v1.UpdateOptions) (result *v1alpha1.NodeSet, err error) {
	emptyResult := &v1alpha1.NodeSet{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateActionWithOptions(nodesetsResource, c.ns, nodeSet, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NodeSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeSets) UpdateStatus(ctx context.Context, nodeSet *v1alpha1.NodeSet, opts v1.UpdateOptions) (result *v1alpha1.NodeSet, err error) {
	emptyResult := &v1alpha1.NodeSet{}
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceActionWithOptions(nodesetsResource, "status", c.ns, nodeSet, opts), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NodeSet), err
}

// Delete takes name of the nodeSet and deletes it. Returns an error if one occurs.
func (c *FakeNodeSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(nodesetsResource, c.ns, name, opts), &v1alpha1.NodeSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionActionWithOptions(nodesetsResource, c.ns, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeSetList{})
	return err
}

// Patch applies the patch and returns the patched nodeSet.
func (c *FakeNodeSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.NodeSet, err error) {
	emptyResult := &v1alpha1.NodeSet{}
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceActionWithOptions(nodesetsResource, c.ns, name, pt, data, opts, subresources...), emptyResult)

	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1alpha1.NodeSet), err
}
