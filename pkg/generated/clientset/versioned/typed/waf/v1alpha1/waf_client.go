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
// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "awsoperator.io/pkg/apis/waf/v1alpha1"
	"awsoperator.io/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type WafV1alpha1Interface interface {
	RESTClient() rest.Interface
	ByteMatchSetsGetter
	IPSetsGetter
	RulesGetter
	SizeConstraintSetsGetter
	SqlInjectionMatchSetsGetter
	WebACLsGetter
	XssMatchSetsGetter
}

// WafV1alpha1Client is used to interact with features provided by the waf.awsoperator.io group.
type WafV1alpha1Client struct {
	restClient rest.Interface
}

func (c *WafV1alpha1Client) ByteMatchSets(namespace string) ByteMatchSetInterface {
	return newByteMatchSets(c, namespace)
}

func (c *WafV1alpha1Client) IPSets(namespace string) IPSetInterface {
	return newIPSets(c, namespace)
}

func (c *WafV1alpha1Client) Rules(namespace string) RuleInterface {
	return newRules(c, namespace)
}

func (c *WafV1alpha1Client) SizeConstraintSets(namespace string) SizeConstraintSetInterface {
	return newSizeConstraintSets(c, namespace)
}

func (c *WafV1alpha1Client) SqlInjectionMatchSets(namespace string) SqlInjectionMatchSetInterface {
	return newSqlInjectionMatchSets(c, namespace)
}

func (c *WafV1alpha1Client) WebACLs(namespace string) WebACLInterface {
	return newWebACLs(c, namespace)
}

func (c *WafV1alpha1Client) XssMatchSets(namespace string) XssMatchSetInterface {
	return newXssMatchSets(c, namespace)
}

// NewForConfig creates a new WafV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*WafV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &WafV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new WafV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *WafV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new WafV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *WafV1alpha1Client {
	return &WafV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *WafV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
