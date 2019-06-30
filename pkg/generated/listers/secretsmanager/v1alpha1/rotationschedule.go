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
	v1alpha1 "awsoperator.io/pkg/apis/secretsmanager/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RotationScheduleLister helps list RotationSchedules.
type RotationScheduleLister interface {
	// List lists all RotationSchedules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.RotationSchedule, err error)
	// RotationSchedules returns an object that can list and get RotationSchedules.
	RotationSchedules(namespace string) RotationScheduleNamespaceLister
	RotationScheduleListerExpansion
}

// rotationScheduleLister implements the RotationScheduleLister interface.
type rotationScheduleLister struct {
	indexer cache.Indexer
}

// NewRotationScheduleLister returns a new RotationScheduleLister.
func NewRotationScheduleLister(indexer cache.Indexer) RotationScheduleLister {
	return &rotationScheduleLister{indexer: indexer}
}

// List lists all RotationSchedules in the indexer.
func (s *rotationScheduleLister) List(selector labels.Selector) (ret []*v1alpha1.RotationSchedule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RotationSchedule))
	})
	return ret, err
}

// RotationSchedules returns an object that can list and get RotationSchedules.
func (s *rotationScheduleLister) RotationSchedules(namespace string) RotationScheduleNamespaceLister {
	return rotationScheduleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RotationScheduleNamespaceLister helps list and get RotationSchedules.
type RotationScheduleNamespaceLister interface {
	// List lists all RotationSchedules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.RotationSchedule, err error)
	// Get retrieves the RotationSchedule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.RotationSchedule, error)
	RotationScheduleNamespaceListerExpansion
}

// rotationScheduleNamespaceLister implements the RotationScheduleNamespaceLister
// interface.
type rotationScheduleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all RotationSchedules in the indexer for a given namespace.
func (s rotationScheduleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.RotationSchedule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.RotationSchedule))
	})
	return ret, err
}

// Get retrieves the RotationSchedule from the indexer for a given namespace and name.
func (s rotationScheduleNamespaceLister) Get(name string) (*v1alpha1.RotationSchedule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("rotationschedule"), name)
	}
	return obj.(*v1alpha1.RotationSchedule), nil
}
