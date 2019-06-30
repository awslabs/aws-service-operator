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
	v1alpha1 "awsoperator.io/pkg/apis/rds/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DBSecurityGroupIngressLister helps list DBSecurityGroupIngresses.
type DBSecurityGroupIngressLister interface {
	// List lists all DBSecurityGroupIngresses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DBSecurityGroupIngress, err error)
	// DBSecurityGroupIngresses returns an object that can list and get DBSecurityGroupIngresses.
	DBSecurityGroupIngresses(namespace string) DBSecurityGroupIngressNamespaceLister
	DBSecurityGroupIngressListerExpansion
}

// dBSecurityGroupIngressLister implements the DBSecurityGroupIngressLister interface.
type dBSecurityGroupIngressLister struct {
	indexer cache.Indexer
}

// NewDBSecurityGroupIngressLister returns a new DBSecurityGroupIngressLister.
func NewDBSecurityGroupIngressLister(indexer cache.Indexer) DBSecurityGroupIngressLister {
	return &dBSecurityGroupIngressLister{indexer: indexer}
}

// List lists all DBSecurityGroupIngresses in the indexer.
func (s *dBSecurityGroupIngressLister) List(selector labels.Selector) (ret []*v1alpha1.DBSecurityGroupIngress, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DBSecurityGroupIngress))
	})
	return ret, err
}

// DBSecurityGroupIngresses returns an object that can list and get DBSecurityGroupIngresses.
func (s *dBSecurityGroupIngressLister) DBSecurityGroupIngresses(namespace string) DBSecurityGroupIngressNamespaceLister {
	return dBSecurityGroupIngressNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DBSecurityGroupIngressNamespaceLister helps list and get DBSecurityGroupIngresses.
type DBSecurityGroupIngressNamespaceLister interface {
	// List lists all DBSecurityGroupIngresses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DBSecurityGroupIngress, err error)
	// Get retrieves the DBSecurityGroupIngress from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DBSecurityGroupIngress, error)
	DBSecurityGroupIngressNamespaceListerExpansion
}

// dBSecurityGroupIngressNamespaceLister implements the DBSecurityGroupIngressNamespaceLister
// interface.
type dBSecurityGroupIngressNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DBSecurityGroupIngresses in the indexer for a given namespace.
func (s dBSecurityGroupIngressNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DBSecurityGroupIngress, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DBSecurityGroupIngress))
	})
	return ret, err
}

// Get retrieves the DBSecurityGroupIngress from the indexer for a given namespace and name.
func (s dBSecurityGroupIngressNamespaceLister) Get(name string) (*v1alpha1.DBSecurityGroupIngress, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dbsecuritygroupingress"), name)
	}
	return obj.(*v1alpha1.DBSecurityGroupIngress), nil
}
