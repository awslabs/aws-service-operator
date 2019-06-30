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
	v1alpha1 "awsoperator.io/pkg/apis/apigateway/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AccountLister helps list Accounts.
type AccountLister interface {
	// List lists all Accounts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Account, err error)
	// Accounts returns an object that can list and get Accounts.
	Accounts(namespace string) AccountNamespaceLister
	AccountListerExpansion
}

// accountLister implements the AccountLister interface.
type accountLister struct {
	indexer cache.Indexer
}

// NewAccountLister returns a new AccountLister.
func NewAccountLister(indexer cache.Indexer) AccountLister {
	return &accountLister{indexer: indexer}
}

// List lists all Accounts in the indexer.
func (s *accountLister) List(selector labels.Selector) (ret []*v1alpha1.Account, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Account))
	})
	return ret, err
}

// Accounts returns an object that can list and get Accounts.
func (s *accountLister) Accounts(namespace string) AccountNamespaceLister {
	return accountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AccountNamespaceLister helps list and get Accounts.
type AccountNamespaceLister interface {
	// List lists all Accounts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Account, err error)
	// Get retrieves the Account from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Account, error)
	AccountNamespaceListerExpansion
}

// accountNamespaceLister implements the AccountNamespaceLister
// interface.
type accountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Accounts in the indexer for a given namespace.
func (s accountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Account, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Account))
	})
	return ret, err
}

// Get retrieves the Account from the indexer for a given namespace and name.
func (s accountNamespaceLister) Get(name string) (*v1alpha1.Account, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("account"), name)
	}
	return obj.(*v1alpha1.Account), nil
}
