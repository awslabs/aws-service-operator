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

// CrawlerLister helps list Crawlers.
type CrawlerLister interface {
	// List lists all Crawlers in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Crawler, err error)
	// Crawlers returns an object that can list and get Crawlers.
	Crawlers(namespace string) CrawlerNamespaceLister
	CrawlerListerExpansion
}

// crawlerLister implements the CrawlerLister interface.
type crawlerLister struct {
	indexer cache.Indexer
}

// NewCrawlerLister returns a new CrawlerLister.
func NewCrawlerLister(indexer cache.Indexer) CrawlerLister {
	return &crawlerLister{indexer: indexer}
}

// List lists all Crawlers in the indexer.
func (s *crawlerLister) List(selector labels.Selector) (ret []*v1alpha1.Crawler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Crawler))
	})
	return ret, err
}

// Crawlers returns an object that can list and get Crawlers.
func (s *crawlerLister) Crawlers(namespace string) CrawlerNamespaceLister {
	return crawlerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CrawlerNamespaceLister helps list and get Crawlers.
type CrawlerNamespaceLister interface {
	// List lists all Crawlers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Crawler, err error)
	// Get retrieves the Crawler from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Crawler, error)
	CrawlerNamespaceListerExpansion
}

// crawlerNamespaceLister implements the CrawlerNamespaceLister
// interface.
type crawlerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Crawlers in the indexer for a given namespace.
func (s crawlerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Crawler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Crawler))
	})
	return ret, err
}

// Get retrieves the Crawler from the indexer for a given namespace and name.
func (s crawlerNamespaceLister) Get(name string) (*v1alpha1.Crawler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("crawler"), name)
	}
	return obj.(*v1alpha1.Crawler), nil
}
