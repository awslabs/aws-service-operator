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
	v1alpha1 "awsoperator.io/pkg/apis/ec2/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VPCEndpointLister helps list VPCEndpoints.
type VPCEndpointLister interface {
	// List lists all VPCEndpoints in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.VPCEndpoint, err error)
	// VPCEndpoints returns an object that can list and get VPCEndpoints.
	VPCEndpoints(namespace string) VPCEndpointNamespaceLister
	VPCEndpointListerExpansion
}

// vPCEndpointLister implements the VPCEndpointLister interface.
type vPCEndpointLister struct {
	indexer cache.Indexer
}

// NewVPCEndpointLister returns a new VPCEndpointLister.
func NewVPCEndpointLister(indexer cache.Indexer) VPCEndpointLister {
	return &vPCEndpointLister{indexer: indexer}
}

// List lists all VPCEndpoints in the indexer.
func (s *vPCEndpointLister) List(selector labels.Selector) (ret []*v1alpha1.VPCEndpoint, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VPCEndpoint))
	})
	return ret, err
}

// VPCEndpoints returns an object that can list and get VPCEndpoints.
func (s *vPCEndpointLister) VPCEndpoints(namespace string) VPCEndpointNamespaceLister {
	return vPCEndpointNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VPCEndpointNamespaceLister helps list and get VPCEndpoints.
type VPCEndpointNamespaceLister interface {
	// List lists all VPCEndpoints in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.VPCEndpoint, err error)
	// Get retrieves the VPCEndpoint from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.VPCEndpoint, error)
	VPCEndpointNamespaceListerExpansion
}

// vPCEndpointNamespaceLister implements the VPCEndpointNamespaceLister
// interface.
type vPCEndpointNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VPCEndpoints in the indexer for a given namespace.
func (s vPCEndpointNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.VPCEndpoint, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VPCEndpoint))
	})
	return ret, err
}

// Get retrieves the VPCEndpoint from the indexer for a given namespace and name.
func (s vPCEndpointNamespaceLister) Get(name string) (*v1alpha1.VPCEndpoint, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcendpoint"), name)
	}
	return obj.(*v1alpha1.VPCEndpoint), nil
}
