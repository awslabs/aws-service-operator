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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "awsoperator.io/pkg/apis/glue/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PartitionLister helps list Partitions.
type PartitionLister interface {
	// List lists all Partitions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Partition, err error)
	// Partitions returns an object that can list and get Partitions.
	Partitions(namespace string) PartitionNamespaceLister
	PartitionListerExpansion
}

// partitionLister implements the PartitionLister interface.
type partitionLister struct {
	indexer cache.Indexer
}

// NewPartitionLister returns a new PartitionLister.
func NewPartitionLister(indexer cache.Indexer) PartitionLister {
	return &partitionLister{indexer: indexer}
}

// List lists all Partitions in the indexer.
func (s *partitionLister) List(selector labels.Selector) (ret []*v1alpha1.Partition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Partition))
	})
	return ret, err
}

// Partitions returns an object that can list and get Partitions.
func (s *partitionLister) Partitions(namespace string) PartitionNamespaceLister {
	return partitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PartitionNamespaceLister helps list and get Partitions.
type PartitionNamespaceLister interface {
	// List lists all Partitions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Partition, err error)
	// Get retrieves the Partition from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Partition, error)
	PartitionNamespaceListerExpansion
}

// partitionNamespaceLister implements the PartitionNamespaceLister
// interface.
type partitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Partitions in the indexer for a given namespace.
func (s partitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Partition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Partition))
	})
	return ret, err
}

// Get retrieves the Partition from the indexer for a given namespace and name.
func (s partitionNamespaceLister) Get(name string) (*v1alpha1.Partition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("partition"), name)
	}
	return obj.(*v1alpha1.Partition), nil
}
