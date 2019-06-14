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

	v1alpha1 "awsoperator.io/pkg/apis/guardduty/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FiltersGetter has a method to return a FilterInterface.
// A group's client should implement this interface.
type FiltersGetter interface {
	Filters(namespace string) FilterInterface
}

// FilterInterface has methods to work with Filter resources.
type FilterInterface interface {
	Create(*v1alpha1.Filter) (*v1alpha1.Filter, error)
	Update(*v1alpha1.Filter) (*v1alpha1.Filter, error)
	UpdateStatus(*v1alpha1.Filter) (*v1alpha1.Filter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Filter, error)
	List(opts v1.ListOptions) (*v1alpha1.FilterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Filter, err error)
	FilterExpansion
}

// filters implements FilterInterface
type filters struct {
	client rest.Interface
	ns     string
}

// newFilters returns a Filters
func newFilters(c *GuarddutyV1alpha1Client, namespace string) *filters {
	return &filters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the filter, and returns the corresponding filter object, and an error if there is any.
func (c *filters) Get(name string, options v1.GetOptions) (result *v1alpha1.Filter, err error) {
	result = &v1alpha1.Filter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("filters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Filters that match those selectors.
func (c *filters) List(opts v1.ListOptions) (result *v1alpha1.FilterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.FilterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("filters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested filters.
func (c *filters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("filters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a filter and creates it.  Returns the server's representation of the filter, and an error, if there is any.
func (c *filters) Create(filter *v1alpha1.Filter) (result *v1alpha1.Filter, err error) {
	result = &v1alpha1.Filter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("filters").
		Body(filter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a filter and updates it. Returns the server's representation of the filter, and an error, if there is any.
func (c *filters) Update(filter *v1alpha1.Filter) (result *v1alpha1.Filter, err error) {
	result = &v1alpha1.Filter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("filters").
		Name(filter.Name).
		Body(filter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *filters) UpdateStatus(filter *v1alpha1.Filter) (result *v1alpha1.Filter, err error) {
	result = &v1alpha1.Filter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("filters").
		Name(filter.Name).
		SubResource("status").
		Body(filter).
		Do().
		Into(result)
	return
}

// Delete takes name of the filter and deletes it. Returns an error if one occurs.
func (c *filters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("filters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *filters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("filters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched filter.
func (c *filters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Filter, err error) {
	result = &v1alpha1.Filter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("filters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
