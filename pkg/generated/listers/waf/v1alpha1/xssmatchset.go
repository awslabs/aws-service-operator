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
	v1alpha1 "awsoperator.io/pkg/apis/waf/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// XssMatchSetLister helps list XssMatchSets.
type XssMatchSetLister interface {
	// List lists all XssMatchSets in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.XssMatchSet, err error)
	// XssMatchSets returns an object that can list and get XssMatchSets.
	XssMatchSets(namespace string) XssMatchSetNamespaceLister
	XssMatchSetListerExpansion
}

// xssMatchSetLister implements the XssMatchSetLister interface.
type xssMatchSetLister struct {
	indexer cache.Indexer
}

// NewXssMatchSetLister returns a new XssMatchSetLister.
func NewXssMatchSetLister(indexer cache.Indexer) XssMatchSetLister {
	return &xssMatchSetLister{indexer: indexer}
}

// List lists all XssMatchSets in the indexer.
func (s *xssMatchSetLister) List(selector labels.Selector) (ret []*v1alpha1.XssMatchSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XssMatchSet))
	})
	return ret, err
}

// XssMatchSets returns an object that can list and get XssMatchSets.
func (s *xssMatchSetLister) XssMatchSets(namespace string) XssMatchSetNamespaceLister {
	return xssMatchSetNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// XssMatchSetNamespaceLister helps list and get XssMatchSets.
type XssMatchSetNamespaceLister interface {
	// List lists all XssMatchSets in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.XssMatchSet, err error)
	// Get retrieves the XssMatchSet from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.XssMatchSet, error)
	XssMatchSetNamespaceListerExpansion
}

// xssMatchSetNamespaceLister implements the XssMatchSetNamespaceLister
// interface.
type xssMatchSetNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all XssMatchSets in the indexer for a given namespace.
func (s xssMatchSetNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.XssMatchSet, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.XssMatchSet))
	})
	return ret, err
}

// Get retrieves the XssMatchSet from the indexer for a given namespace and name.
func (s xssMatchSetNamespaceLister) Get(name string) (*v1alpha1.XssMatchSet, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("xssmatchset"), name)
	}
	return obj.(*v1alpha1.XssMatchSet), nil
}
