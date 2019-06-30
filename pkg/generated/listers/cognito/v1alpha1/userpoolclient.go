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
	v1alpha1 "awsoperator.io/pkg/apis/cognito/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserPoolClientLister helps list UserPoolClients.
type UserPoolClientLister interface {
	// List lists all UserPoolClients in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.UserPoolClient, err error)
	// UserPoolClients returns an object that can list and get UserPoolClients.
	UserPoolClients(namespace string) UserPoolClientNamespaceLister
	UserPoolClientListerExpansion
}

// userPoolClientLister implements the UserPoolClientLister interface.
type userPoolClientLister struct {
	indexer cache.Indexer
}

// NewUserPoolClientLister returns a new UserPoolClientLister.
func NewUserPoolClientLister(indexer cache.Indexer) UserPoolClientLister {
	return &userPoolClientLister{indexer: indexer}
}

// List lists all UserPoolClients in the indexer.
func (s *userPoolClientLister) List(selector labels.Selector) (ret []*v1alpha1.UserPoolClient, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserPoolClient))
	})
	return ret, err
}

// UserPoolClients returns an object that can list and get UserPoolClients.
func (s *userPoolClientLister) UserPoolClients(namespace string) UserPoolClientNamespaceLister {
	return userPoolClientNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserPoolClientNamespaceLister helps list and get UserPoolClients.
type UserPoolClientNamespaceLister interface {
	// List lists all UserPoolClients in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.UserPoolClient, err error)
	// Get retrieves the UserPoolClient from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.UserPoolClient, error)
	UserPoolClientNamespaceListerExpansion
}

// userPoolClientNamespaceLister implements the UserPoolClientNamespaceLister
// interface.
type userPoolClientNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserPoolClients in the indexer for a given namespace.
func (s userPoolClientNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.UserPoolClient, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.UserPoolClient))
	})
	return ret, err
}

// Get retrieves the UserPoolClient from the indexer for a given namespace and name.
func (s userPoolClientNamespaceLister) Get(name string) (*v1alpha1.UserPoolClient, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("userpoolclient"), name)
	}
	return obj.(*v1alpha1.UserPoolClient), nil
}
