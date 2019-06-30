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
	v1alpha1 "awsoperator.io/pkg/apis/config/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigRuleLister helps list ConfigRules.
type ConfigRuleLister interface {
	// List lists all ConfigRules in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigRule, err error)
	// ConfigRules returns an object that can list and get ConfigRules.
	ConfigRules(namespace string) ConfigRuleNamespaceLister
	ConfigRuleListerExpansion
}

// configRuleLister implements the ConfigRuleLister interface.
type configRuleLister struct {
	indexer cache.Indexer
}

// NewConfigRuleLister returns a new ConfigRuleLister.
func NewConfigRuleLister(indexer cache.Indexer) ConfigRuleLister {
	return &configRuleLister{indexer: indexer}
}

// List lists all ConfigRules in the indexer.
func (s *configRuleLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigRule))
	})
	return ret, err
}

// ConfigRules returns an object that can list and get ConfigRules.
func (s *configRuleLister) ConfigRules(namespace string) ConfigRuleNamespaceLister {
	return configRuleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ConfigRuleNamespaceLister helps list and get ConfigRules.
type ConfigRuleNamespaceLister interface {
	// List lists all ConfigRules in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ConfigRule, err error)
	// Get retrieves the ConfigRule from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ConfigRule, error)
	ConfigRuleNamespaceListerExpansion
}

// configRuleNamespaceLister implements the ConfigRuleNamespaceLister
// interface.
type configRuleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ConfigRules in the indexer for a given namespace.
func (s configRuleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ConfigRule, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ConfigRule))
	})
	return ret, err
}

// Get retrieves the ConfigRule from the indexer for a given namespace and name.
func (s configRuleNamespaceLister) Get(name string) (*v1alpha1.ConfigRule, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("configrule"), name)
	}
	return obj.(*v1alpha1.ConfigRule), nil
}
