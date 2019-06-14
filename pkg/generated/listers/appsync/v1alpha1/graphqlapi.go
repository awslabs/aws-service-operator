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
	v1alpha1 "awsoperator.io/pkg/apis/appsync/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// GraphQLApiLister helps list GraphQLApis.
type GraphQLApiLister interface {
	// List lists all GraphQLApis in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.GraphQLApi, err error)
	// GraphQLApis returns an object that can list and get GraphQLApis.
	GraphQLApis(namespace string) GraphQLApiNamespaceLister
	GraphQLApiListerExpansion
}

// graphQLApiLister implements the GraphQLApiLister interface.
type graphQLApiLister struct {
	indexer cache.Indexer
}

// NewGraphQLApiLister returns a new GraphQLApiLister.
func NewGraphQLApiLister(indexer cache.Indexer) GraphQLApiLister {
	return &graphQLApiLister{indexer: indexer}
}

// List lists all GraphQLApis in the indexer.
func (s *graphQLApiLister) List(selector labels.Selector) (ret []*v1alpha1.GraphQLApi, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GraphQLApi))
	})
	return ret, err
}

// GraphQLApis returns an object that can list and get GraphQLApis.
func (s *graphQLApiLister) GraphQLApis(namespace string) GraphQLApiNamespaceLister {
	return graphQLApiNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GraphQLApiNamespaceLister helps list and get GraphQLApis.
type GraphQLApiNamespaceLister interface {
	// List lists all GraphQLApis in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.GraphQLApi, err error)
	// Get retrieves the GraphQLApi from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.GraphQLApi, error)
	GraphQLApiNamespaceListerExpansion
}

// graphQLApiNamespaceLister implements the GraphQLApiNamespaceLister
// interface.
type graphQLApiNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GraphQLApis in the indexer for a given namespace.
func (s graphQLApiNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.GraphQLApi, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GraphQLApi))
	})
	return ret, err
}

// Get retrieves the GraphQLApi from the indexer for a given namespace and name.
func (s graphQLApiNamespaceLister) Get(name string) (*v1alpha1.GraphQLApi, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("graphqlapi"), name)
	}
	return obj.(*v1alpha1.GraphQLApi), nil
}
