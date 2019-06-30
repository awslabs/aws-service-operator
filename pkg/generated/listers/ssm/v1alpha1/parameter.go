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
	v1alpha1 "awsoperator.io/pkg/apis/ssm/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ParameterLister helps list Parameters.
type ParameterLister interface {
	// List lists all Parameters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Parameter, err error)
	// Parameters returns an object that can list and get Parameters.
	Parameters(namespace string) ParameterNamespaceLister
	ParameterListerExpansion
}

// parameterLister implements the ParameterLister interface.
type parameterLister struct {
	indexer cache.Indexer
}

// NewParameterLister returns a new ParameterLister.
func NewParameterLister(indexer cache.Indexer) ParameterLister {
	return &parameterLister{indexer: indexer}
}

// List lists all Parameters in the indexer.
func (s *parameterLister) List(selector labels.Selector) (ret []*v1alpha1.Parameter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Parameter))
	})
	return ret, err
}

// Parameters returns an object that can list and get Parameters.
func (s *parameterLister) Parameters(namespace string) ParameterNamespaceLister {
	return parameterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ParameterNamespaceLister helps list and get Parameters.
type ParameterNamespaceLister interface {
	// List lists all Parameters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Parameter, err error)
	// Get retrieves the Parameter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Parameter, error)
	ParameterNamespaceListerExpansion
}

// parameterNamespaceLister implements the ParameterNamespaceLister
// interface.
type parameterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Parameters in the indexer for a given namespace.
func (s parameterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Parameter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Parameter))
	})
	return ret, err
}

// Get retrieves the Parameter from the indexer for a given namespace and name.
func (s parameterNamespaceLister) Get(name string) (*v1alpha1.Parameter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("parameter"), name)
	}
	return obj.(*v1alpha1.Parameter), nil
}
