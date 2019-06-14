/*
Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License"). You may
not use this file except in compliance with the License. A copy of the
License is located at

     http://aws.amazon.com/apache2.0/

or in the "license" file accompanying this file. This file is distributed
on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
express or implied. See the License for the specific language governing
permissions and limitations under the License.
*/
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	greengrassv1alpha1 "awsoperator.io/pkg/apis/greengrass/v1alpha1"
	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "awsoperator.io/pkg/generated/listers/greengrass/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// FunctionDefinitionVersionInformer provides access to a shared informer and lister for
// FunctionDefinitionVersions.
type FunctionDefinitionVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.FunctionDefinitionVersionLister
}

type functionDefinitionVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFunctionDefinitionVersionInformer constructs a new informer for FunctionDefinitionVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFunctionDefinitionVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFunctionDefinitionVersionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFunctionDefinitionVersionInformer constructs a new informer for FunctionDefinitionVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFunctionDefinitionVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GreengrassV1alpha1().FunctionDefinitionVersions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.GreengrassV1alpha1().FunctionDefinitionVersions(namespace).Watch(options)
			},
		},
		&greengrassv1alpha1.FunctionDefinitionVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *functionDefinitionVersionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFunctionDefinitionVersionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *functionDefinitionVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&greengrassv1alpha1.FunctionDefinitionVersion{}, f.defaultInformer)
}

func (f *functionDefinitionVersionInformer) Lister() v1alpha1.FunctionDefinitionVersionLister {
	return v1alpha1.NewFunctionDefinitionVersionLister(f.Informer().GetIndexer())
}
