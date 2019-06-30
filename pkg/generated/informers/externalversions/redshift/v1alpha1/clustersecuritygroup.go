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

	redshiftv1alpha1 "awsoperator.io/pkg/apis/redshift/v1alpha1"
	versioned "awsoperator.io/pkg/generated/clientset/versioned"
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "awsoperator.io/pkg/generated/listers/redshift/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterSecurityGroupInformer provides access to a shared informer and lister for
// ClusterSecurityGroups.
type ClusterSecurityGroupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterSecurityGroupLister
}

type clusterSecurityGroupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewClusterSecurityGroupInformer constructs a new informer for ClusterSecurityGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterSecurityGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterSecurityGroupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredClusterSecurityGroupInformer constructs a new informer for ClusterSecurityGroup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterSecurityGroupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RedshiftV1alpha1().ClusterSecurityGroups(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.RedshiftV1alpha1().ClusterSecurityGroups(namespace).Watch(options)
			},
		},
		&redshiftv1alpha1.ClusterSecurityGroup{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterSecurityGroupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterSecurityGroupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterSecurityGroupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&redshiftv1alpha1.ClusterSecurityGroup{}, f.defaultInformer)
}

func (f *clusterSecurityGroupInformer) Lister() v1alpha1.ClusterSecurityGroupLister {
	return v1alpha1.NewClusterSecurityGroupLister(f.Informer().GetIndexer())
}
