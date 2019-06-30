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
	v1alpha1 "awsoperator.io/pkg/apis/greengrass/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FunctionDefinitionVersionLister helps list FunctionDefinitionVersions.
type FunctionDefinitionVersionLister interface {
	// List lists all FunctionDefinitionVersions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.FunctionDefinitionVersion, err error)
	// FunctionDefinitionVersions returns an object that can list and get FunctionDefinitionVersions.
	FunctionDefinitionVersions(namespace string) FunctionDefinitionVersionNamespaceLister
	FunctionDefinitionVersionListerExpansion
}

// functionDefinitionVersionLister implements the FunctionDefinitionVersionLister interface.
type functionDefinitionVersionLister struct {
	indexer cache.Indexer
}

// NewFunctionDefinitionVersionLister returns a new FunctionDefinitionVersionLister.
func NewFunctionDefinitionVersionLister(indexer cache.Indexer) FunctionDefinitionVersionLister {
	return &functionDefinitionVersionLister{indexer: indexer}
}

// List lists all FunctionDefinitionVersions in the indexer.
func (s *functionDefinitionVersionLister) List(selector labels.Selector) (ret []*v1alpha1.FunctionDefinitionVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FunctionDefinitionVersion))
	})
	return ret, err
}

// FunctionDefinitionVersions returns an object that can list and get FunctionDefinitionVersions.
func (s *functionDefinitionVersionLister) FunctionDefinitionVersions(namespace string) FunctionDefinitionVersionNamespaceLister {
	return functionDefinitionVersionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// FunctionDefinitionVersionNamespaceLister helps list and get FunctionDefinitionVersions.
type FunctionDefinitionVersionNamespaceLister interface {
	// List lists all FunctionDefinitionVersions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.FunctionDefinitionVersion, err error)
	// Get retrieves the FunctionDefinitionVersion from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.FunctionDefinitionVersion, error)
	FunctionDefinitionVersionNamespaceListerExpansion
}

// functionDefinitionVersionNamespaceLister implements the FunctionDefinitionVersionNamespaceLister
// interface.
type functionDefinitionVersionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all FunctionDefinitionVersions in the indexer for a given namespace.
func (s functionDefinitionVersionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.FunctionDefinitionVersion, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.FunctionDefinitionVersion))
	})
	return ret, err
}

// Get retrieves the FunctionDefinitionVersion from the indexer for a given namespace and name.
func (s functionDefinitionVersionNamespaceLister) Get(name string) (*v1alpha1.FunctionDefinitionVersion, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("functiondefinitionversion"), name)
	}
	return obj.(*v1alpha1.FunctionDefinitionVersion), nil
}
