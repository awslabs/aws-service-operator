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

// DHCPOptionsLister helps list DHCPOptionses.
type DHCPOptionsLister interface {
	// List lists all DHCPOptionses in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.DHCPOptions, err error)
	// DHCPOptionses returns an object that can list and get DHCPOptionses.
	DHCPOptionses(namespace string) DHCPOptionsNamespaceLister
	DHCPOptionsListerExpansion
}

// dHCPOptionsLister implements the DHCPOptionsLister interface.
type dHCPOptionsLister struct {
	indexer cache.Indexer
}

// NewDHCPOptionsLister returns a new DHCPOptionsLister.
func NewDHCPOptionsLister(indexer cache.Indexer) DHCPOptionsLister {
	return &dHCPOptionsLister{indexer: indexer}
}

// List lists all DHCPOptionses in the indexer.
func (s *dHCPOptionsLister) List(selector labels.Selector) (ret []*v1alpha1.DHCPOptions, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DHCPOptions))
	})
	return ret, err
}

// DHCPOptionses returns an object that can list and get DHCPOptionses.
func (s *dHCPOptionsLister) DHCPOptionses(namespace string) DHCPOptionsNamespaceLister {
	return dHCPOptionsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DHCPOptionsNamespaceLister helps list and get DHCPOptionses.
type DHCPOptionsNamespaceLister interface {
	// List lists all DHCPOptionses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.DHCPOptions, err error)
	// Get retrieves the DHCPOptions from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.DHCPOptions, error)
	DHCPOptionsNamespaceListerExpansion
}

// dHCPOptionsNamespaceLister implements the DHCPOptionsNamespaceLister
// interface.
type dHCPOptionsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all DHCPOptionses in the indexer for a given namespace.
func (s dHCPOptionsNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.DHCPOptions, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.DHCPOptions))
	})
	return ret, err
}

// Get retrieves the DHCPOptions from the indexer for a given namespace and name.
func (s dHCPOptionsNamespaceLister) Get(name string) (*v1alpha1.DHCPOptions, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("dhcpoptions"), name)
	}
	return obj.(*v1alpha1.DHCPOptions), nil
}
