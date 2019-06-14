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

	v1alpha1 "awsoperator.io/pkg/apis/sagemaker/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EndpointsGetter has a method to return a EndpointInterface.
// A group's client should implement this interface.
type EndpointsGetter interface {
	Endpoints(namespace string) EndpointInterface
}

// EndpointInterface has methods to work with Endpoint resources.
type EndpointInterface interface {
	Create(*v1alpha1.Endpoint) (*v1alpha1.Endpoint, error)
	Update(*v1alpha1.Endpoint) (*v1alpha1.Endpoint, error)
	UpdateStatus(*v1alpha1.Endpoint) (*v1alpha1.Endpoint, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Endpoint, error)
	List(opts v1.ListOptions) (*v1alpha1.EndpointList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Endpoint, err error)
	EndpointExpansion
}

// endpoints implements EndpointInterface
type endpoints struct {
	client rest.Interface
	ns     string
}

// newEndpoints returns a Endpoints
func newEndpoints(c *SagemakerV1alpha1Client, namespace string) *endpoints {
	return &endpoints{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the endpoint, and returns the corresponding endpoint object, and an error if there is any.
func (c *endpoints) Get(name string, options v1.GetOptions) (result *v1alpha1.Endpoint, err error) {
	result = &v1alpha1.Endpoint{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Endpoints that match those selectors.
func (c *endpoints) List(opts v1.ListOptions) (result *v1alpha1.EndpointList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.EndpointList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested endpoints.
func (c *endpoints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a endpoint and creates it.  Returns the server's representation of the endpoint, and an error, if there is any.
func (c *endpoints) Create(endpoint *v1alpha1.Endpoint) (result *v1alpha1.Endpoint, err error) {
	result = &v1alpha1.Endpoint{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("endpoints").
		Body(endpoint).
		Do().
		Into(result)
	return
}

// Update takes the representation of a endpoint and updates it. Returns the server's representation of the endpoint, and an error, if there is any.
func (c *endpoints) Update(endpoint *v1alpha1.Endpoint) (result *v1alpha1.Endpoint, err error) {
	result = &v1alpha1.Endpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpoints").
		Name(endpoint.Name).
		Body(endpoint).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *endpoints) UpdateStatus(endpoint *v1alpha1.Endpoint) (result *v1alpha1.Endpoint, err error) {
	result = &v1alpha1.Endpoint{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("endpoints").
		Name(endpoint.Name).
		SubResource("status").
		Body(endpoint).
		Do().
		Into(result)
	return
}

// Delete takes name of the endpoint and deletes it. Returns an error if one occurs.
func (c *endpoints) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *endpoints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("endpoints").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched endpoint.
func (c *endpoints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Endpoint, err error) {
	result = &v1alpha1.Endpoint{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("endpoints").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
