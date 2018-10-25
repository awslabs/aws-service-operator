/*
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
	time "time"

	service_operator_aws_v1alpha1 "github.com/awslabs/aws-service-operator/pkg/apis/service-operator.aws/v1alpha1"
	versioned "github.com/awslabs/aws-service-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/awslabs/aws-service-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/awslabs/aws-service-operator/pkg/client/listers/service-operator.aws/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SNSTopicInformer provides access to a shared informer and lister for
// SNSTopics.
type SNSTopicInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.SNSTopicLister
}

type sNSTopicInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSNSTopicInformer constructs a new informer for SNSTopic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSNSTopicInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSNSTopicInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSNSTopicInformer constructs a new informer for SNSTopic type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSNSTopicInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServiceoperatorV1alpha1().SNSTopics(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ServiceoperatorV1alpha1().SNSTopics(namespace).Watch(options)
			},
		},
		&service_operator_aws_v1alpha1.SNSTopic{},
		resyncPeriod,
		indexers,
	)
}

func (f *sNSTopicInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSNSTopicInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sNSTopicInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&service_operator_aws_v1alpha1.SNSTopic{}, f.defaultInformer)
}

func (f *sNSTopicInformer) Lister() v1alpha1.SNSTopicLister {
	return v1alpha1.NewSNSTopicLister(f.Informer().GetIndexer())
}
