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
	v1alpha1 "awsoperator.io/pkg/apis/wafregional/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// WebACLLister helps list WebACLs.
type WebACLLister interface {
	// List lists all WebACLs in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.WebACL, err error)
	// WebACLs returns an object that can list and get WebACLs.
	WebACLs(namespace string) WebACLNamespaceLister
	WebACLListerExpansion
}

// webACLLister implements the WebACLLister interface.
type webACLLister struct {
	indexer cache.Indexer
}

// NewWebACLLister returns a new WebACLLister.
func NewWebACLLister(indexer cache.Indexer) WebACLLister {
	return &webACLLister{indexer: indexer}
}

// List lists all WebACLs in the indexer.
func (s *webACLLister) List(selector labels.Selector) (ret []*v1alpha1.WebACL, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebACL))
	})
	return ret, err
}

// WebACLs returns an object that can list and get WebACLs.
func (s *webACLLister) WebACLs(namespace string) WebACLNamespaceLister {
	return webACLNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// WebACLNamespaceLister helps list and get WebACLs.
type WebACLNamespaceLister interface {
	// List lists all WebACLs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.WebACL, err error)
	// Get retrieves the WebACL from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.WebACL, error)
	WebACLNamespaceListerExpansion
}

// webACLNamespaceLister implements the WebACLNamespaceLister
// interface.
type webACLNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all WebACLs in the indexer for a given namespace.
func (s webACLNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.WebACL, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.WebACL))
	})
	return ret, err
}

// Get retrieves the WebACL from the indexer for a given namespace and name.
func (s webACLNamespaceLister) Get(name string) (*v1alpha1.WebACL, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("webacl"), name)
	}
	return obj.(*v1alpha1.WebACL), nil
}
