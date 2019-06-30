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
	v1alpha1 "awsoperator.io/pkg/apis/emr/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StepLister helps list Steps.
type StepLister interface {
	// List lists all Steps in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Step, err error)
	// Steps returns an object that can list and get Steps.
	Steps(namespace string) StepNamespaceLister
	StepListerExpansion
}

// stepLister implements the StepLister interface.
type stepLister struct {
	indexer cache.Indexer
}

// NewStepLister returns a new StepLister.
func NewStepLister(indexer cache.Indexer) StepLister {
	return &stepLister{indexer: indexer}
}

// List lists all Steps in the indexer.
func (s *stepLister) List(selector labels.Selector) (ret []*v1alpha1.Step, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Step))
	})
	return ret, err
}

// Steps returns an object that can list and get Steps.
func (s *stepLister) Steps(namespace string) StepNamespaceLister {
	return stepNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// StepNamespaceLister helps list and get Steps.
type StepNamespaceLister interface {
	// List lists all Steps in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Step, err error)
	// Get retrieves the Step from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Step, error)
	StepNamespaceListerExpansion
}

// stepNamespaceLister implements the StepNamespaceLister
// interface.
type stepNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Steps in the indexer for a given namespace.
func (s stepNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Step, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Step))
	})
	return ret, err
}

// Get retrieves the Step from the indexer for a given namespace and name.
func (s stepNamespaceLister) Get(name string) (*v1alpha1.Step, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("step"), name)
	}
	return obj.(*v1alpha1.Step), nil
}
