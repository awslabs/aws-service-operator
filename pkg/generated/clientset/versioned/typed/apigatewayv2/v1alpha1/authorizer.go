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

	v1alpha1 "awsoperator.io/pkg/apis/apigatewayv2/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AuthorizersGetter has a method to return a AuthorizerInterface.
// A group's client should implement this interface.
type AuthorizersGetter interface {
	Authorizers(namespace string) AuthorizerInterface
}

// AuthorizerInterface has methods to work with Authorizer resources.
type AuthorizerInterface interface {
	Create(*v1alpha1.Authorizer) (*v1alpha1.Authorizer, error)
	Update(*v1alpha1.Authorizer) (*v1alpha1.Authorizer, error)
	UpdateStatus(*v1alpha1.Authorizer) (*v1alpha1.Authorizer, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Authorizer, error)
	List(opts v1.ListOptions) (*v1alpha1.AuthorizerList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Authorizer, err error)
	AuthorizerExpansion
}

// authorizers implements AuthorizerInterface
type authorizers struct {
	client rest.Interface
	ns     string
}

// newAuthorizers returns a Authorizers
func newAuthorizers(c *Apigatewayv2V1alpha1Client, namespace string) *authorizers {
	return &authorizers{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the authorizer, and returns the corresponding authorizer object, and an error if there is any.
func (c *authorizers) Get(name string, options v1.GetOptions) (result *v1alpha1.Authorizer, err error) {
	result = &v1alpha1.Authorizer{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authorizers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Authorizers that match those selectors.
func (c *authorizers) List(opts v1.ListOptions) (result *v1alpha1.AuthorizerList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AuthorizerList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("authorizers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested authorizers.
func (c *authorizers) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("authorizers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a authorizer and creates it.  Returns the server's representation of the authorizer, and an error, if there is any.
func (c *authorizers) Create(authorizer *v1alpha1.Authorizer) (result *v1alpha1.Authorizer, err error) {
	result = &v1alpha1.Authorizer{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("authorizers").
		Body(authorizer).
		Do().
		Into(result)
	return
}

// Update takes the representation of a authorizer and updates it. Returns the server's representation of the authorizer, and an error, if there is any.
func (c *authorizers) Update(authorizer *v1alpha1.Authorizer) (result *v1alpha1.Authorizer, err error) {
	result = &v1alpha1.Authorizer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authorizers").
		Name(authorizer.Name).
		Body(authorizer).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *authorizers) UpdateStatus(authorizer *v1alpha1.Authorizer) (result *v1alpha1.Authorizer, err error) {
	result = &v1alpha1.Authorizer{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("authorizers").
		Name(authorizer.Name).
		SubResource("status").
		Body(authorizer).
		Do().
		Into(result)
	return
}

// Delete takes name of the authorizer and deletes it. Returns an error if one occurs.
func (c *authorizers) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authorizers").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *authorizers) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("authorizers").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched authorizer.
func (c *authorizers) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Authorizer, err error) {
	result = &v1alpha1.Authorizer{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("authorizers").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
