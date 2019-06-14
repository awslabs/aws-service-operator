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
	v1alpha1 "awsoperator.io/pkg/apis/servicecatalog/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PortfolioShareLister helps list PortfolioShares.
type PortfolioShareLister interface {
	// List lists all PortfolioShares in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.PortfolioShare, err error)
	// PortfolioShares returns an object that can list and get PortfolioShares.
	PortfolioShares(namespace string) PortfolioShareNamespaceLister
	PortfolioShareListerExpansion
}

// portfolioShareLister implements the PortfolioShareLister interface.
type portfolioShareLister struct {
	indexer cache.Indexer
}

// NewPortfolioShareLister returns a new PortfolioShareLister.
func NewPortfolioShareLister(indexer cache.Indexer) PortfolioShareLister {
	return &portfolioShareLister{indexer: indexer}
}

// List lists all PortfolioShares in the indexer.
func (s *portfolioShareLister) List(selector labels.Selector) (ret []*v1alpha1.PortfolioShare, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PortfolioShare))
	})
	return ret, err
}

// PortfolioShares returns an object that can list and get PortfolioShares.
func (s *portfolioShareLister) PortfolioShares(namespace string) PortfolioShareNamespaceLister {
	return portfolioShareNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PortfolioShareNamespaceLister helps list and get PortfolioShares.
type PortfolioShareNamespaceLister interface {
	// List lists all PortfolioShares in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.PortfolioShare, err error)
	// Get retrieves the PortfolioShare from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.PortfolioShare, error)
	PortfolioShareNamespaceListerExpansion
}

// portfolioShareNamespaceLister implements the PortfolioShareNamespaceLister
// interface.
type portfolioShareNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all PortfolioShares in the indexer for a given namespace.
func (s portfolioShareNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.PortfolioShare, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.PortfolioShare))
	})
	return ret, err
}

// Get retrieves the PortfolioShare from the indexer for a given namespace and name.
func (s portfolioShareNamespaceLister) Get(name string) (*v1alpha1.PortfolioShare, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("portfolioshare"), name)
	}
	return obj.(*v1alpha1.PortfolioShare), nil
}
