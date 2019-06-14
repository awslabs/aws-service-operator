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

	elasticbeanstalkv1alpha1 "awsoperator.io/pkg/apis/elasticbeanstalk/v1alpha1"
	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "awsoperator.io/pkg/generated/listers/elasticbeanstalk/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApplicationVersionInformer provides access to a shared informer and lister for
// ApplicationVersions.
type ApplicationVersionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ApplicationVersionLister
}

type applicationVersionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApplicationVersionInformer constructs a new informer for ApplicationVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApplicationVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApplicationVersionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApplicationVersionInformer constructs a new informer for ApplicationVersion type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApplicationVersionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ElasticbeanstalkV1alpha1().ApplicationVersions(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ElasticbeanstalkV1alpha1().ApplicationVersions(namespace).Watch(options)
			},
		},
		&elasticbeanstalkv1alpha1.ApplicationVersion{},
		resyncPeriod,
		indexers,
	)
}

func (f *applicationVersionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApplicationVersionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *applicationVersionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&elasticbeanstalkv1alpha1.ApplicationVersion{}, f.defaultInformer)
}

func (f *applicationVersionInformer) Lister() v1alpha1.ApplicationVersionLister {
	return v1alpha1.NewApplicationVersionLister(f.Informer().GetIndexer())
}
