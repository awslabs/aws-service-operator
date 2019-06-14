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
	v1alpha1 "awsoperator.io/pkg/apis/apigatewayv2/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApiMappingLister helps list ApiMappings.
type ApiMappingLister interface {
	// List lists all ApiMappings in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ApiMapping, err error)
	// ApiMappings returns an object that can list and get ApiMappings.
	ApiMappings(namespace string) ApiMappingNamespaceLister
	ApiMappingListerExpansion
}

// apiMappingLister implements the ApiMappingLister interface.
type apiMappingLister struct {
	indexer cache.Indexer
}

// NewApiMappingLister returns a new ApiMappingLister.
func NewApiMappingLister(indexer cache.Indexer) ApiMappingLister {
	return &apiMappingLister{indexer: indexer}
}

// List lists all ApiMappings in the indexer.
func (s *apiMappingLister) List(selector labels.Selector) (ret []*v1alpha1.ApiMapping, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiMapping))
	})
	return ret, err
}

// ApiMappings returns an object that can list and get ApiMappings.
func (s *apiMappingLister) ApiMappings(namespace string) ApiMappingNamespaceLister {
	return apiMappingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApiMappingNamespaceLister helps list and get ApiMappings.
type ApiMappingNamespaceLister interface {
	// List lists all ApiMappings in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ApiMapping, err error)
	// Get retrieves the ApiMapping from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ApiMapping, error)
	ApiMappingNamespaceListerExpansion
}

// apiMappingNamespaceLister implements the ApiMappingNamespaceLister
// interface.
type apiMappingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApiMappings in the indexer for a given namespace.
func (s apiMappingNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ApiMapping, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ApiMapping))
	})
	return ret, err
}

// Get retrieves the ApiMapping from the indexer for a given namespace and name.
func (s apiMappingNamespaceLister) Get(name string) (*v1alpha1.ApiMapping, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("apimapping"), name)
	}
	return obj.(*v1alpha1.ApiMapping), nil
}
