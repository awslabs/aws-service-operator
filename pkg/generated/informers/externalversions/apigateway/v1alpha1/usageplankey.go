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

	apigatewayv1alpha1 "awsoperator.io/pkg/apis/apigateway/v1alpha1"
	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "awsoperator.io/pkg/generated/listers/apigateway/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UsagePlanKeyInformer provides access to a shared informer and lister for
// UsagePlanKeys.
type UsagePlanKeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.UsagePlanKeyLister
}

type usagePlanKeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewUsagePlanKeyInformer constructs a new informer for UsagePlanKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUsagePlanKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUsagePlanKeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredUsagePlanKeyInformer constructs a new informer for UsagePlanKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUsagePlanKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApigatewayV1alpha1().UsagePlanKeys(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ApigatewayV1alpha1().UsagePlanKeys(namespace).Watch(options)
			},
		},
		&apigatewayv1alpha1.UsagePlanKey{},
		resyncPeriod,
		indexers,
	)
}

func (f *usagePlanKeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUsagePlanKeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *usagePlanKeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apigatewayv1alpha1.UsagePlanKey{}, f.defaultInformer)
}

func (f *usagePlanKeyInformer) Lister() v1alpha1.UsagePlanKeyLister {
	return v1alpha1.NewUsagePlanKeyLister(f.Informer().GetIndexer())
}
