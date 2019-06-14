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

	v1alpha1 "awsoperator.io/pkg/apis/kms/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KeysGetter has a method to return a KeyInterface.
// A group's client should implement this interface.
type KeysGetter interface {
	Keys(namespace string) KeyInterface
}

// KeyInterface has methods to work with Key resources.
type KeyInterface interface {
	Create(*v1alpha1.Key) (*v1alpha1.Key, error)
	Update(*v1alpha1.Key) (*v1alpha1.Key, error)
	UpdateStatus(*v1alpha1.Key) (*v1alpha1.Key, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Key, error)
	List(opts v1.ListOptions) (*v1alpha1.KeyList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Key, err error)
	KeyExpansion
}

// keys implements KeyInterface
type keys struct {
	client rest.Interface
	ns     string
}

// newKeys returns a Keys
func newKeys(c *KmsV1alpha1Client, namespace string) *keys {
	return &keys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the key, and returns the corresponding key object, and an error if there is any.
func (c *keys) Get(name string, options v1.GetOptions) (result *v1alpha1.Key, err error) {
	result = &v1alpha1.Key{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Keys that match those selectors.
func (c *keys) List(opts v1.ListOptions) (result *v1alpha1.KeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.KeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("keys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested keys.
func (c *keys) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("keys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a key and creates it.  Returns the server's representation of the key, and an error, if there is any.
func (c *keys) Create(key *v1alpha1.Key) (result *v1alpha1.Key, err error) {
	result = &v1alpha1.Key{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("keys").
		Body(key).
		Do().
		Into(result)
	return
}

// Update takes the representation of a key and updates it. Returns the server's representation of the key, and an error, if there is any.
func (c *keys) Update(key *v1alpha1.Key) (result *v1alpha1.Key, err error) {
	result = &v1alpha1.Key{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keys").
		Name(key.Name).
		Body(key).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *keys) UpdateStatus(key *v1alpha1.Key) (result *v1alpha1.Key, err error) {
	result = &v1alpha1.Key{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("keys").
		Name(key.Name).
		SubResource("status").
		Body(key).
		Do().
		Into(result)
	return
}

// Delete takes name of the key and deletes it. Returns an error if one occurs.
func (c *keys) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keys").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *keys) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("keys").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched key.
func (c *keys) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Key, err error) {
	result = &v1alpha1.Key{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("keys").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
