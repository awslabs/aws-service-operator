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

// TransitGatewayRouteTableAssociationLister helps list TransitGatewayRouteTableAssociations.
type TransitGatewayRouteTableAssociationLister interface {
	// List lists all TransitGatewayRouteTableAssociations in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayRouteTableAssociation, err error)
	// TransitGatewayRouteTableAssociations returns an object that can list and get TransitGatewayRouteTableAssociations.
	TransitGatewayRouteTableAssociations(namespace string) TransitGatewayRouteTableAssociationNamespaceLister
	TransitGatewayRouteTableAssociationListerExpansion
}

// transitGatewayRouteTableAssociationLister implements the TransitGatewayRouteTableAssociationLister interface.
type transitGatewayRouteTableAssociationLister struct {
	indexer cache.Indexer
}

// NewTransitGatewayRouteTableAssociationLister returns a new TransitGatewayRouteTableAssociationLister.
func NewTransitGatewayRouteTableAssociationLister(indexer cache.Indexer) TransitGatewayRouteTableAssociationLister {
	return &transitGatewayRouteTableAssociationLister{indexer: indexer}
}

// List lists all TransitGatewayRouteTableAssociations in the indexer.
func (s *transitGatewayRouteTableAssociationLister) List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayRouteTableAssociation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransitGatewayRouteTableAssociation))
	})
	return ret, err
}

// TransitGatewayRouteTableAssociations returns an object that can list and get TransitGatewayRouteTableAssociations.
func (s *transitGatewayRouteTableAssociationLister) TransitGatewayRouteTableAssociations(namespace string) TransitGatewayRouteTableAssociationNamespaceLister {
	return transitGatewayRouteTableAssociationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TransitGatewayRouteTableAssociationNamespaceLister helps list and get TransitGatewayRouteTableAssociations.
type TransitGatewayRouteTableAssociationNamespaceLister interface {
	// List lists all TransitGatewayRouteTableAssociations in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayRouteTableAssociation, err error)
	// Get retrieves the TransitGatewayRouteTableAssociation from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TransitGatewayRouteTableAssociation, error)
	TransitGatewayRouteTableAssociationNamespaceListerExpansion
}

// transitGatewayRouteTableAssociationNamespaceLister implements the TransitGatewayRouteTableAssociationNamespaceLister
// interface.
type transitGatewayRouteTableAssociationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TransitGatewayRouteTableAssociations in the indexer for a given namespace.
func (s transitGatewayRouteTableAssociationNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TransitGatewayRouteTableAssociation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TransitGatewayRouteTableAssociation))
	})
	return ret, err
}

// Get retrieves the TransitGatewayRouteTableAssociation from the indexer for a given namespace and name.
func (s transitGatewayRouteTableAssociationNamespaceLister) Get(name string) (*v1alpha1.TransitGatewayRouteTableAssociation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("transitgatewayroutetableassociation"), name)
	}
	return obj.(*v1alpha1.TransitGatewayRouteTableAssociation), nil
}
