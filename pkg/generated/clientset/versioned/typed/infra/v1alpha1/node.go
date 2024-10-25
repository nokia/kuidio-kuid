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

	v1alpha1 "github.com/kuidio/kuid/apis/infra/v1alpha1"
	scheme "github.com/kuidio/kuid/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// NodesGetter has a method to return a NodeInterface.
// A group's client should implement this interface.
type NodesGetter interface {
	Nodes(namespace string) NodeInterface
}

// NodeInterface has methods to work with Node resources.
type NodeInterface interface {
	Create(ctx context.Context, node *v1alpha1.Node, opts v1.CreateOptions) (*v1alpha1.Node, error)
	Update(ctx context.Context, node *v1alpha1.Node, opts v1.UpdateOptions) (*v1alpha1.Node, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, node *v1alpha1.Node, opts v1.UpdateOptions) (*v1alpha1.Node, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.Node, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.NodeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Node, err error)
	NodeExpansion
}

// nodes implements NodeInterface
type nodes struct {
	*gentype.ClientWithList[*v1alpha1.Node, *v1alpha1.NodeList]
}

// newNodes returns a Nodes
func newNodes(c *InfraV1alpha1Client, namespace string) *nodes {
	return &nodes{
		gentype.NewClientWithList[*v1alpha1.Node, *v1alpha1.NodeList](
			"nodes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.Node { return &v1alpha1.Node{} },
			func() *v1alpha1.NodeList { return &v1alpha1.NodeList{} }),
	}
}
