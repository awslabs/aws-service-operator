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

// NetworkInterfacePermissionLister helps list NetworkInterfacePermissions.
type NetworkInterfacePermissionLister interface {
	// List lists all NetworkInterfacePermissions in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfacePermission, err error)
	// NetworkInterfacePermissions returns an object that can list and get NetworkInterfacePermissions.
	NetworkInterfacePermissions(namespace string) NetworkInterfacePermissionNamespaceLister
	NetworkInterfacePermissionListerExpansion
}

// networkInterfacePermissionLister implements the NetworkInterfacePermissionLister interface.
type networkInterfacePermissionLister struct {
	indexer cache.Indexer
}

// NewNetworkInterfacePermissionLister returns a new NetworkInterfacePermissionLister.
func NewNetworkInterfacePermissionLister(indexer cache.Indexer) NetworkInterfacePermissionLister {
	return &networkInterfacePermissionLister{indexer: indexer}
}

// List lists all NetworkInterfacePermissions in the indexer.
func (s *networkInterfacePermissionLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfacePermission, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkInterfacePermission))
	})
	return ret, err
}

// NetworkInterfacePermissions returns an object that can list and get NetworkInterfacePermissions.
func (s *networkInterfacePermissionLister) NetworkInterfacePermissions(namespace string) NetworkInterfacePermissionNamespaceLister {
	return networkInterfacePermissionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkInterfacePermissionNamespaceLister helps list and get NetworkInterfacePermissions.
type NetworkInterfacePermissionNamespaceLister interface {
	// List lists all NetworkInterfacePermissions in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfacePermission, err error)
	// Get retrieves the NetworkInterfacePermission from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.NetworkInterfacePermission, error)
	NetworkInterfacePermissionNamespaceListerExpansion
}

// networkInterfacePermissionNamespaceLister implements the NetworkInterfacePermissionNamespaceLister
// interface.
type networkInterfacePermissionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkInterfacePermissions in the indexer for a given namespace.
func (s networkInterfacePermissionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NetworkInterfacePermission, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NetworkInterfacePermission))
	})
	return ret, err
}

// Get retrieves the NetworkInterfacePermission from the indexer for a given namespace and name.
func (s networkInterfacePermissionNamespaceLister) Get(name string) (*v1alpha1.NetworkInterfacePermission, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("networkinterfacepermission"), name)
	}
	return obj.(*v1alpha1.NetworkInterfacePermission), nil
}
