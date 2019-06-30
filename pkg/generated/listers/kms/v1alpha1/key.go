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
	v1alpha1 "awsoperator.io/pkg/apis/kms/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// KeyLister helps list Keys.
type KeyLister interface {
	// List lists all Keys in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Key, err error)
	// Keys returns an object that can list and get Keys.
	Keys(namespace string) KeyNamespaceLister
	KeyListerExpansion
}

// keyLister implements the KeyLister interface.
type keyLister struct {
	indexer cache.Indexer
}

// NewKeyLister returns a new KeyLister.
func NewKeyLister(indexer cache.Indexer) KeyLister {
	return &keyLister{indexer: indexer}
}

// List lists all Keys in the indexer.
func (s *keyLister) List(selector labels.Selector) (ret []*v1alpha1.Key, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Key))
	})
	return ret, err
}

// Keys returns an object that can list and get Keys.
func (s *keyLister) Keys(namespace string) KeyNamespaceLister {
	return keyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// KeyNamespaceLister helps list and get Keys.
type KeyNamespaceLister interface {
	// List lists all Keys in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Key, err error)
	// Get retrieves the Key from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Key, error)
	KeyNamespaceListerExpansion
}

// keyNamespaceLister implements the KeyNamespaceLister
// interface.
type keyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Keys in the indexer for a given namespace.
func (s keyNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Key, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Key))
	})
	return ret, err
}

// Get retrieves the Key from the indexer for a given namespace and name.
func (s keyNamespaceLister) Get(name string) (*v1alpha1.Key, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("key"), name)
	}
	return obj.(*v1alpha1.Key), nil
}
