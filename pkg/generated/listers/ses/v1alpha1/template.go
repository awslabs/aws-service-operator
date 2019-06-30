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
	v1alpha1 "awsoperator.io/pkg/apis/ses/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TemplateLister helps list Templates.
type TemplateLister interface {
	// List lists all Templates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Template, err error)
	// Templates returns an object that can list and get Templates.
	Templates(namespace string) TemplateNamespaceLister
	TemplateListerExpansion
}

// templateLister implements the TemplateLister interface.
type templateLister struct {
	indexer cache.Indexer
}

// NewTemplateLister returns a new TemplateLister.
func NewTemplateLister(indexer cache.Indexer) TemplateLister {
	return &templateLister{indexer: indexer}
}

// List lists all Templates in the indexer.
func (s *templateLister) List(selector labels.Selector) (ret []*v1alpha1.Template, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Template))
	})
	return ret, err
}

// Templates returns an object that can list and get Templates.
func (s *templateLister) Templates(namespace string) TemplateNamespaceLister {
	return templateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TemplateNamespaceLister helps list and get Templates.
type TemplateNamespaceLister interface {
	// List lists all Templates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Template, err error)
	// Get retrieves the Template from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Template, error)
	TemplateNamespaceListerExpansion
}

// templateNamespaceLister implements the TemplateNamespaceLister
// interface.
type templateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Templates in the indexer for a given namespace.
func (s templateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Template, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Template))
	})
	return ret, err
}

// Get retrieves the Template from the indexer for a given namespace and name.
func (s templateNamespaceLister) Get(name string) (*v1alpha1.Template, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("template"), name)
	}
	return obj.(*v1alpha1.Template), nil
}
