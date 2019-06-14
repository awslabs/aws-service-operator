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
	v1alpha1 "awsoperator.io/pkg/apis/neptune/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DBParameterGroupLister helps list DBParameterGroups.
type DBParameterGroupLister interface {
	// List lists all DBParameterGroups in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DBParameterGroup, err error)
	// DBParameterGroups returns an object that can list and get DBParameterGroups.
	DBParameterGroups(namespace string) DBParameterGroupNamespaceLister
	DBParameterGroupListerExpansion
}

// dBParameterGroupLister implements the DBParameterGroupLister interface.
type dBParameterGroupLister struct {
	indexer cache.Indexer
}

// NewDBParameterGroupLister returns a new DBParameterGroupLister.
func NewDBParameterGroupLister(indexer cache.Indexer) DBParameterGroupLister {
	return &dBParameterGroupLister{indexer: indexer}
}

// List lists all DBParameterGroups in the indexer.
func (s *dBParameterGroupLister) List(selector labels.Selector) (ret []*v1alpha1.DBParameterGroup, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DBParameterGroup))
	})
	return ret, err
}

// DBParameterGroups returns an object that can list and get DBParameterGroups.
func (s *dBParameterGroupLister) DBParameterGroups(namespace string) DBParameterGroupNamespaceLister {
	return dBParameterGroupNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DBParameterGroupNamespaceLister helps list and get DBParameterGroups.
type DBParameterGroupNamespaceLister interface {
	// List lists all DBParameterGroups in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DBParameterGroup, err error)
	// Get retrieves the DBParameterGroup from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DBParameterGroup, error)
	DBParameterGroupNamespaceListerExpansion
}

// dBParameterGroupNamespaceLister implements the DBParameterGroupNamespaceLister
// interface.
type dBParameterGroupNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DBParameterGroups in the indexer for a given namespace.
func (s dBParameterGroupNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DBParameterGroup, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DBParameterGroup))
	})
	return ret, err
}

// Get retrieves the DBParameterGroup from the indexer for a given namespace and name.
func (s dBParameterGroupNamespaceLister) Get(name string) (*v1alpha1.DBParameterGroup, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbparametergroup"), name)
	}
	return obj.(*v1alpha1.DBParameterGroup), nil
}
