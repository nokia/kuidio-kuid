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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kuidio/kuid/apis/backend/ipam/v1alpha1"
	scheme "github.com/kuidio/kuid/apis/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IPIndexesGetter has a method to return a IPIndexInterface.
// A group's client should implement this interface.
type IPIndexesGetter interface {
	IPIndexes(namespace string) IPIndexInterface
}

// IPIndexInterface has methods to work with IPIndex resources.
type IPIndexInterface interface {
	Create(ctx context.Context, iPIndex *v1alpha1.IPIndex, opts v1.CreateOptions) (*v1alpha1.IPIndex, error)
	Update(ctx context.Context, iPIndex *v1alpha1.IPIndex, opts v1.UpdateOptions) (*v1alpha1.IPIndex, error)
	UpdateStatus(ctx context.Context, iPIndex *v1alpha1.IPIndex, opts v1.UpdateOptions) (*v1alpha1.IPIndex, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.IPIndex, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.IPIndexList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPIndex, err error)
	IPIndexExpansion
}

// iPIndexes implements IPIndexInterface
type iPIndexes struct {
	client rest.Interface
	ns     string
}

// newIPIndexes returns a IPIndexes
func newIPIndexes(c *IpamV1alpha1Client, namespace string) *iPIndexes {
	return &iPIndexes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iPIndex, and returns the corresponding iPIndex object, and an error if there is any.
func (c *iPIndexes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IPIndex, err error) {
	result = &v1alpha1.IPIndex{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipindexes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPIndexes that match those selectors.
func (c *iPIndexes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPIndexList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.IPIndexList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipindexes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPIndexes.
func (c *iPIndexes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ipindexes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iPIndex and creates it.  Returns the server's representation of the iPIndex, and an error, if there is any.
func (c *iPIndexes) Create(ctx context.Context, iPIndex *v1alpha1.IPIndex, opts v1.CreateOptions) (result *v1alpha1.IPIndex, err error) {
	result = &v1alpha1.IPIndex{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ipindexes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPIndex).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iPIndex and updates it. Returns the server's representation of the iPIndex, and an error, if there is any.
func (c *iPIndexes) Update(ctx context.Context, iPIndex *v1alpha1.IPIndex, opts v1.UpdateOptions) (result *v1alpha1.IPIndex, err error) {
	result = &v1alpha1.IPIndex{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipindexes").
		Name(iPIndex.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPIndex).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *iPIndexes) UpdateStatus(ctx context.Context, iPIndex *v1alpha1.IPIndex, opts v1.UpdateOptions) (result *v1alpha1.IPIndex, err error) {
	result = &v1alpha1.IPIndex{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipindexes").
		Name(iPIndex.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPIndex).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iPIndex and deletes it. Returns an error if one occurs.
func (c *iPIndexes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipindexes").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPIndexes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipindexes").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iPIndex.
func (c *iPIndexes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPIndex, err error) {
	result = &v1alpha1.IPIndex{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ipindexes").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
