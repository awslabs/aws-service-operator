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
	"time"

	v1alpha1 "awsoperator.io/pkg/apis/servicecatalog/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PortfoliosGetter has a method to return a PortfolioInterface.
// A group's client should implement this interface.
type PortfoliosGetter interface {
	Portfolios(namespace string) PortfolioInterface
}

// PortfolioInterface has methods to work with Portfolio resources.
type PortfolioInterface interface {
	Create(*v1alpha1.Portfolio) (*v1alpha1.Portfolio, error)
	Update(*v1alpha1.Portfolio) (*v1alpha1.Portfolio, error)
	UpdateStatus(*v1alpha1.Portfolio) (*v1alpha1.Portfolio, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Portfolio, error)
	List(opts v1.ListOptions) (*v1alpha1.PortfolioList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Portfolio, err error)
	PortfolioExpansion
}

// portfolios implements PortfolioInterface
type portfolios struct {
	client rest.Interface
	ns     string
}

// newPortfolios returns a Portfolios
func newPortfolios(c *ServicecatalogV1alpha1Client, namespace string) *portfolios {
	return &portfolios{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the portfolio, and returns the corresponding portfolio object, and an error if there is any.
func (c *portfolios) Get(name string, options v1.GetOptions) (result *v1alpha1.Portfolio, err error) {
	result = &v1alpha1.Portfolio{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("portfolios").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Portfolios that match those selectors.
func (c *portfolios) List(opts v1.ListOptions) (result *v1alpha1.PortfolioList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PortfolioList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("portfolios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested portfolios.
func (c *portfolios) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("portfolios").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a portfolio and creates it.  Returns the server's representation of the portfolio, and an error, if there is any.
func (c *portfolios) Create(portfolio *v1alpha1.Portfolio) (result *v1alpha1.Portfolio, err error) {
	result = &v1alpha1.Portfolio{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("portfolios").
		Body(portfolio).
		Do().
		Into(result)
	return
}

// Update takes the representation of a portfolio and updates it. Returns the server's representation of the portfolio, and an error, if there is any.
func (c *portfolios) Update(portfolio *v1alpha1.Portfolio) (result *v1alpha1.Portfolio, err error) {
	result = &v1alpha1.Portfolio{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("portfolios").
		Name(portfolio.Name).
		Body(portfolio).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *portfolios) UpdateStatus(portfolio *v1alpha1.Portfolio) (result *v1alpha1.Portfolio, err error) {
	result = &v1alpha1.Portfolio{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("portfolios").
		Name(portfolio.Name).
		SubResource("status").
		Body(portfolio).
		Do().
		Into(result)
	return
}

// Delete takes name of the portfolio and deletes it. Returns an error if one occurs.
func (c *portfolios) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("portfolios").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *portfolios) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("portfolios").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched portfolio.
func (c *portfolios) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Portfolio, err error) {
	result = &v1alpha1.Portfolio{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("portfolios").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
