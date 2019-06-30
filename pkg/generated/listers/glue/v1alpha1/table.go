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

// TableLister helps list Tables.
type TableLister interface {
	// List lists all Tables in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Table, err error)
	// Tables returns an object that can list and get Tables.
	Tables(namespace string) TableNamespaceLister
	TableListerExpansion
}

// tableLister implements the TableLister interface.
type tableLister struct {
	indexer cache.Indexer
}

// NewTableLister returns a new TableLister.
func NewTableLister(indexer cache.Indexer) TableLister {
	return &tableLister{indexer: indexer}
}

// List lists all Tables in the indexer.
func (s *tableLister) List(selector labels.Selector) (ret []*v1alpha1.Table, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Table))
	})
	return ret, err
}

// Tables returns an object that can list and get Tables.
func (s *tableLister) Tables(namespace string) TableNamespaceLister {
	return tableNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TableNamespaceLister helps list and get Tables.
type TableNamespaceLister interface {
	// List lists all Tables in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Table, err error)
	// Get retrieves the Table from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Table, error)
	TableNamespaceListerExpansion
}

// tableNamespaceLister implements the TableNamespaceLister
// interface.
type tableNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Tables in the indexer for a given namespace.
func (s tableNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Table, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Table))
	})
	return ret, err
}

// Get retrieves the Table from the indexer for a given namespace and name.
func (s tableNamespaceLister) Get(name string) (*v1alpha1.Table, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("table"), name)
	}
	return obj.(*v1alpha1.Table), nil
}
