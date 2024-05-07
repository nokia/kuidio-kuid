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
	v1alpha1 "github.com/kuidio/kuid/apis/backend/genid/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GENIDEntryLister helps list GENIDEntries.
// All objects returned here must be treated as read-only.
type GENIDEntryLister interface {
	// List lists all GENIDEntries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GENIDEntry, err error)
	// GENIDEntries returns an object that can list and get GENIDEntries.
	GENIDEntries(namespace string) GENIDEntryNamespaceLister
	GENIDEntryListerExpansion
}

// gENIDEntryLister implements the GENIDEntryLister interface.
type gENIDEntryLister struct {
	indexer cache.Indexer
}

// NewGENIDEntryLister returns a new GENIDEntryLister.
func NewGENIDEntryLister(indexer cache.Indexer) GENIDEntryLister {
	return &gENIDEntryLister{indexer: indexer}
}

// List lists all GENIDEntries in the indexer.
func (s *gENIDEntryLister) List(selector labels.Selector) (ret []*v1alpha1.GENIDEntry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GENIDEntry))
	})
	return ret, err
}

// GENIDEntries returns an object that can list and get GENIDEntries.
func (s *gENIDEntryLister) GENIDEntries(namespace string) GENIDEntryNamespaceLister {
	return gENIDEntryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GENIDEntryNamespaceLister helps list and get GENIDEntries.
// All objects returned here must be treated as read-only.
type GENIDEntryNamespaceLister interface {
	// List lists all GENIDEntries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.GENIDEntry, err error)
	// Get retrieves the GENIDEntry from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.GENIDEntry, error)
	GENIDEntryNamespaceListerExpansion
}

// gENIDEntryNamespaceLister implements the GENIDEntryNamespaceLister
// interface.
type gENIDEntryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GENIDEntries in the indexer for a given namespace.
func (s gENIDEntryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GENIDEntry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GENIDEntry))
	})
	return ret, err
}

// Get retrieves the GENIDEntry from the indexer for a given namespace and name.
func (s gENIDEntryNamespaceLister) Get(name string) (*v1alpha1.GENIDEntry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("genidentry"), name)
	}
	return obj.(*v1alpha1.GENIDEntry), nil
}