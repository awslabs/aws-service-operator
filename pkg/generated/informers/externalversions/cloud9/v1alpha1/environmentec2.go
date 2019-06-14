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

	cloud9v1alpha1 "awsoperator.io/pkg/apis/cloud9/v1alpha1"
	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "awsoperator.io/pkg/generated/listers/cloud9/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// EnvironmentEC2Informer provides access to a shared informer and lister for
// EnvironmentEC2s.
type EnvironmentEC2Informer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.EnvironmentEC2Lister
}

type environmentEC2Informer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEnvironmentEC2Informer constructs a new informer for EnvironmentEC2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEnvironmentEC2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEnvironmentEC2Informer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEnvironmentEC2Informer constructs a new informer for EnvironmentEC2 type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEnvironmentEC2Informer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Cloud9V1alpha1().EnvironmentEC2s(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Cloud9V1alpha1().EnvironmentEC2s(namespace).Watch(options)
			},
		},
		&cloud9v1alpha1.EnvironmentEC2{},
		resyncPeriod,
		indexers,
	)
}

func (f *environmentEC2Informer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEnvironmentEC2Informer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *environmentEC2Informer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cloud9v1alpha1.EnvironmentEC2{}, f.defaultInformer)
}

func (f *environmentEC2Informer) Lister() v1alpha1.EnvironmentEC2Lister {
	return v1alpha1.NewEnvironmentEC2Lister(f.Informer().GetIndexer())
}
