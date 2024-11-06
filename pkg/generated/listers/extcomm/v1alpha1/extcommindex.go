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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kuidio/kuid/apis/backend/extcomm/v1alpha1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// EXTCOMMIndexLister helps list EXTCOMMIndexes.
// All objects returned here must be treated as read-only.
type EXTCOMMIndexLister interface {
	// List lists all EXTCOMMIndexes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EXTCOMMIndex, err error)
	// EXTCOMMIndexes returns an object that can list and get EXTCOMMIndexes.
	EXTCOMMIndexes(namespace string) EXTCOMMIndexNamespaceLister
	EXTCOMMIndexListerExpansion
}

// eXTCOMMIndexLister implements the EXTCOMMIndexLister interface.
type eXTCOMMIndexLister struct {
	listers.ResourceIndexer[*v1alpha1.EXTCOMMIndex]
}

// NewEXTCOMMIndexLister returns a new EXTCOMMIndexLister.
func NewEXTCOMMIndexLister(indexer cache.Indexer) EXTCOMMIndexLister {
	return &eXTCOMMIndexLister{listers.New[*v1alpha1.EXTCOMMIndex](indexer, v1alpha1.Resource("extcommindex"))}
}

// EXTCOMMIndexes returns an object that can list and get EXTCOMMIndexes.
func (s *eXTCOMMIndexLister) EXTCOMMIndexes(namespace string) EXTCOMMIndexNamespaceLister {
	return eXTCOMMIndexNamespaceLister{listers.NewNamespaced[*v1alpha1.EXTCOMMIndex](s.ResourceIndexer, namespace)}
}

// EXTCOMMIndexNamespaceLister helps list and get EXTCOMMIndexes.
// All objects returned here must be treated as read-only.
type EXTCOMMIndexNamespaceLister interface {
	// List lists all EXTCOMMIndexes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.EXTCOMMIndex, err error)
	// Get retrieves the EXTCOMMIndex from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.EXTCOMMIndex, error)
	EXTCOMMIndexNamespaceListerExpansion
}

// eXTCOMMIndexNamespaceLister implements the EXTCOMMIndexNamespaceLister
// interface.
type eXTCOMMIndexNamespaceLister struct {
	listers.ResourceIndexer[*v1alpha1.EXTCOMMIndex]
}