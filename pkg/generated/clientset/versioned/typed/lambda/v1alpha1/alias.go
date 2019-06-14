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

	v1alpha1 "awsoperator.io/pkg/apis/lambda/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AliasesGetter has a method to return a AliasInterface.
// A group's client should implement this interface.
type AliasesGetter interface {
	Aliases(namespace string) AliasInterface
}

// AliasInterface has methods to work with Alias resources.
type AliasInterface interface {
	Create(*v1alpha1.Alias) (*v1alpha1.Alias, error)
	Update(*v1alpha1.Alias) (*v1alpha1.Alias, error)
	UpdateStatus(*v1alpha1.Alias) (*v1alpha1.Alias, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Alias, error)
	List(opts v1.ListOptions) (*v1alpha1.AliasList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alias, err error)
	AliasExpansion
}

// aliases implements AliasInterface
type aliases struct {
	client rest.Interface
	ns     string
}

// newAliases returns a Aliases
func newAliases(c *LambdaV1alpha1Client, namespace string) *aliases {
	return &aliases{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the alias, and returns the corresponding alias object, and an error if there is any.
func (c *aliases) Get(name string, options v1.GetOptions) (result *v1alpha1.Alias, err error) {
	result = &v1alpha1.Alias{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("aliases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Aliases that match those selectors.
func (c *aliases) List(opts v1.ListOptions) (result *v1alpha1.AliasList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AliasList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("aliases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aliases.
func (c *aliases) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("aliases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a alias and creates it.  Returns the server's representation of the alias, and an error, if there is any.
func (c *aliases) Create(alias *v1alpha1.Alias) (result *v1alpha1.Alias, err error) {
	result = &v1alpha1.Alias{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("aliases").
		Body(alias).
		Do().
		Into(result)
	return
}

// Update takes the representation of a alias and updates it. Returns the server's representation of the alias, and an error, if there is any.
func (c *aliases) Update(alias *v1alpha1.Alias) (result *v1alpha1.Alias, err error) {
	result = &v1alpha1.Alias{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("aliases").
		Name(alias.Name).
		Body(alias).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *aliases) UpdateStatus(alias *v1alpha1.Alias) (result *v1alpha1.Alias, err error) {
	result = &v1alpha1.Alias{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("aliases").
		Name(alias.Name).
		SubResource("status").
		Body(alias).
		Do().
		Into(result)
	return
}

// Delete takes name of the alias and deletes it. Returns an error if one occurs.
func (c *aliases) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("aliases").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aliases) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("aliases").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched alias.
func (c *aliases) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Alias, err error) {
	result = &v1alpha1.Alias{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("aliases").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
