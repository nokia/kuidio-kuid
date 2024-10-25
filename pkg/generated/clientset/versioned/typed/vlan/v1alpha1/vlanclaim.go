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

	v1alpha1 "github.com/kuidio/kuid/apis/backend/vlan/v1alpha1"
	scheme "github.com/kuidio/kuid/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// VLANClaimsGetter has a method to return a VLANClaimInterface.
// A group's client should implement this interface.
type VLANClaimsGetter interface {
	VLANClaims(namespace string) VLANClaimInterface
}

// VLANClaimInterface has methods to work with VLANClaim resources.
type VLANClaimInterface interface {
	Create(ctx context.Context, vLANClaim *v1alpha1.VLANClaim, opts v1.CreateOptions) (*v1alpha1.VLANClaim, error)
	Update(ctx context.Context, vLANClaim *v1alpha1.VLANClaim, opts v1.UpdateOptions) (*v1alpha1.VLANClaim, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, vLANClaim *v1alpha1.VLANClaim, opts v1.UpdateOptions) (*v1alpha1.VLANClaim, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.VLANClaim, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.VLANClaimList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.VLANClaim, err error)
	VLANClaimExpansion
}

// vLANClaims implements VLANClaimInterface
type vLANClaims struct {
	*gentype.ClientWithList[*v1alpha1.VLANClaim, *v1alpha1.VLANClaimList]
}

// newVLANClaims returns a VLANClaims
func newVLANClaims(c *VlanV1alpha1Client, namespace string) *vLANClaims {
	return &vLANClaims{
		gentype.NewClientWithList[*v1alpha1.VLANClaim, *v1alpha1.VLANClaimList](
			"vlanclaims",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha1.VLANClaim { return &v1alpha1.VLANClaim{} },
			func() *v1alpha1.VLANClaimList { return &v1alpha1.VLANClaimList{} }),
	}
}
