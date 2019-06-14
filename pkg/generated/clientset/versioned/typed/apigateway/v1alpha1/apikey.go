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

	v1alpha1 "awsoperator.io/pkg/apis/apigateway/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApiKeysGetter has a method to return a ApiKeyInterface.
// A group's client should implement this interface.
type ApiKeysGetter interface {
	ApiKeys(namespace string) ApiKeyInterface
}

// ApiKeyInterface has methods to work with ApiKey resources.
type ApiKeyInterface interface {
	Create(*v1alpha1.ApiKey) (*v1alpha1.ApiKey, error)
	Update(*v1alpha1.ApiKey) (*v1alpha1.ApiKey, error)
	UpdateStatus(*v1alpha1.ApiKey) (*v1alpha1.ApiKey, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApiKey, error)
	List(opts v1.ListOptions) (*v1alpha1.ApiKeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiKey, err error)
	ApiKeyExpansion
}

// apiKeys implements ApiKeyInterface
type apiKeys struct {
	client rest.Interface
	ns     string
}

// newApiKeys returns a ApiKeys
func newApiKeys(c *ApigatewayV1alpha1Client, namespace string) *apiKeys {
	return &apiKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the apiKey, and returns the corresponding apiKey object, and an error if there is any.
func (c *apiKeys) Get(name string, options v1.GetOptions) (result *v1alpha1.ApiKey, err error) {
	result = &v1alpha1.ApiKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apikeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApiKeys that match those selectors.
func (c *apiKeys) List(opts v1.ListOptions) (result *v1alpha1.ApiKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApiKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apiKeys.
func (c *apiKeys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apikeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a apiKey and creates it.  Returns the server's representation of the apiKey, and an error, if there is any.
func (c *apiKeys) Create(apiKey *v1alpha1.ApiKey) (result *v1alpha1.ApiKey, err error) {
	result = &v1alpha1.ApiKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apikeys").
		Body(apiKey).
		Do().
		Into(result)
	return
}

// Update takes the representation of a apiKey and updates it. Returns the server's representation of the apiKey, and an error, if there is any.
func (c *apiKeys) Update(apiKey *v1alpha1.ApiKey) (result *v1alpha1.ApiKey, err error) {
	result = &v1alpha1.ApiKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apikeys").
		Name(apiKey.Name).
		Body(apiKey).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *apiKeys) UpdateStatus(apiKey *v1alpha1.ApiKey) (result *v1alpha1.ApiKey, err error) {
	result = &v1alpha1.ApiKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apikeys").
		Name(apiKey.Name).
		SubResource("status").
		Body(apiKey).
		Do().
		Into(result)
	return
}

// Delete takes name of the apiKey and deletes it. Returns an error if one occurs.
func (c *apiKeys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apikeys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apiKeys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apikeys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched apiKey.
func (c *apiKeys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApiKey, err error) {
	result = &v1alpha1.ApiKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apikeys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
