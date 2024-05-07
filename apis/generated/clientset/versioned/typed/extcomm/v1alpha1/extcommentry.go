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

	v1alpha1 "github.com/kuidio/kuid/apis/backend/extcomm/v1alpha1"
	scheme "github.com/kuidio/kuid/apis/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EXTCOMMEntriesGetter has a method to return a EXTCOMMEntryInterface.
// A group's client should implement this interface.
type EXTCOMMEntriesGetter interface {
	EXTCOMMEntries(namespace string) EXTCOMMEntryInterface
}

// EXTCOMMEntryInterface has methods to work with EXTCOMMEntry resources.
type EXTCOMMEntryInterface interface {
	Create(ctx context.Context, eXTCOMMEntry *v1alpha1.EXTCOMMEntry, opts v1.CreateOptions) (*v1alpha1.EXTCOMMEntry, error)
	Update(ctx context.Context, eXTCOMMEntry *v1alpha1.EXTCOMMEntry, opts v1.UpdateOptions) (*v1alpha1.EXTCOMMEntry, error)
	UpdateStatus(ctx context.Context, eXTCOMMEntry *v1alpha1.EXTCOMMEntry, opts v1.UpdateOptions) (*v1alpha1.EXTCOMMEntry, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.EXTCOMMEntry, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.EXTCOMMEntryList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EXTCOMMEntry, err error)
	EXTCOMMEntryExpansion
}

// eXTCOMMEntries implements EXTCOMMEntryInterface
type eXTCOMMEntries struct {
	client rest.Interface
	ns     string
}

// newEXTCOMMEntries returns a EXTCOMMEntries
func newEXTCOMMEntries(c *ExtcommV1alpha1Client, namespace string) *eXTCOMMEntries {
	return &eXTCOMMEntries{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eXTCOMMEntry, and returns the corresponding eXTCOMMEntry object, and an error if there is any.
func (c *eXTCOMMEntries) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EXTCOMMEntry, err error) {
	result = &v1alpha1.EXTCOMMEntry{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extcommentries").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EXTCOMMEntries that match those selectors.
func (c *eXTCOMMEntries) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EXTCOMMEntryList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EXTCOMMEntryList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("extcommentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eXTCOMMEntries.
func (c *eXTCOMMEntries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("extcommentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a eXTCOMMEntry and creates it.  Returns the server's representation of the eXTCOMMEntry, and an error, if there is any.
func (c *eXTCOMMEntries) Create(ctx context.Context, eXTCOMMEntry *v1alpha1.EXTCOMMEntry, opts v1.CreateOptions) (result *v1alpha1.EXTCOMMEntry, err error) {
	result = &v1alpha1.EXTCOMMEntry{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("extcommentries").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eXTCOMMEntry).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a eXTCOMMEntry and updates it. Returns the server's representation of the eXTCOMMEntry, and an error, if there is any.
func (c *eXTCOMMEntries) Update(ctx context.Context, eXTCOMMEntry *v1alpha1.EXTCOMMEntry, opts v1.UpdateOptions) (result *v1alpha1.EXTCOMMEntry, err error) {
	result = &v1alpha1.EXTCOMMEntry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extcommentries").
		Name(eXTCOMMEntry.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eXTCOMMEntry).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *eXTCOMMEntries) UpdateStatus(ctx context.Context, eXTCOMMEntry *v1alpha1.EXTCOMMEntry, opts v1.UpdateOptions) (result *v1alpha1.EXTCOMMEntry, err error) {
	result = &v1alpha1.EXTCOMMEntry{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("extcommentries").
		Name(eXTCOMMEntry.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(eXTCOMMEntry).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the eXTCOMMEntry and deletes it. Returns an error if one occurs.
func (c *eXTCOMMEntries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extcommentries").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eXTCOMMEntries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("extcommentries").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched eXTCOMMEntry.
func (c *eXTCOMMEntries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EXTCOMMEntry, err error) {
	result = &v1alpha1.EXTCOMMEntry{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("extcommentries").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}