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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	infrav1alpha1 "github.com/kuidio/kuid/apis/infra/v1alpha1"
	versioned "github.com/kuidio/kuid/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/kuidio/kuid/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kuidio/kuid/pkg/generated/listers/infra/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ModuleInformer provides access to a shared informer and lister for
// Modules.
type ModuleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ModuleLister
}

type moduleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewModuleInformer constructs a new informer for Module type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewModuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredModuleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredModuleInformer constructs a new informer for Module type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredModuleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InfraV1alpha1().Modules(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.InfraV1alpha1().Modules(namespace).Watch(context.TODO(), options)
			},
		},
		&infrav1alpha1.Module{},
		resyncPeriod,
		indexers,
	)
}

func (f *moduleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredModuleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *moduleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&infrav1alpha1.Module{}, f.defaultInformer)
}

func (f *moduleInformer) Lister() v1alpha1.ModuleLister {
	return v1alpha1.NewModuleLister(f.Informer().GetIndexer())
}
