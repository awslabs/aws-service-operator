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
	v1alpha1 "awsoperator.io/pkg/apis/appmesh/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VirtualRouterLister helps list VirtualRouters.
type VirtualRouterLister interface {
	// List lists all VirtualRouters in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualRouter, err error)
	// VirtualRouters returns an object that can list and get VirtualRouters.
	VirtualRouters(namespace string) VirtualRouterNamespaceLister
	VirtualRouterListerExpansion
}

// virtualRouterLister implements the VirtualRouterLister interface.
type virtualRouterLister struct {
	indexer cache.Indexer
}

// NewVirtualRouterLister returns a new VirtualRouterLister.
func NewVirtualRouterLister(indexer cache.Indexer) VirtualRouterLister {
	return &virtualRouterLister{indexer: indexer}
}

// List lists all VirtualRouters in the indexer.
func (s *virtualRouterLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualRouter, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualRouter))
	})
	return ret, err
}

// VirtualRouters returns an object that can list and get VirtualRouters.
func (s *virtualRouterLister) VirtualRouters(namespace string) VirtualRouterNamespaceLister {
	return virtualRouterNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VirtualRouterNamespaceLister helps list and get VirtualRouters.
type VirtualRouterNamespaceLister interface {
	// List lists all VirtualRouters in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VirtualRouter, err error)
	// Get retrieves the VirtualRouter from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VirtualRouter, error)
	VirtualRouterNamespaceListerExpansion
}

// virtualRouterNamespaceLister implements the VirtualRouterNamespaceLister
// interface.
type virtualRouterNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VirtualRouters in the indexer for a given namespace.
func (s virtualRouterNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VirtualRouter, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VirtualRouter))
	})
	return ret, err
}

// Get retrieves the VirtualRouter from the indexer for a given namespace and name.
func (s virtualRouterNamespaceLister) Get(name string) (*v1alpha1.VirtualRouter, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("virtualrouter"), name)
	}
	return obj.(*v1alpha1.VirtualRouter), nil
}
