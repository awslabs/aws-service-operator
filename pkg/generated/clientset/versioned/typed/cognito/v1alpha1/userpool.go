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

	v1alpha1 "awsoperator.io/pkg/apis/cognito/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UserPoolsGetter has a method to return a UserPoolInterface.
// A group's client should implement this interface.
type UserPoolsGetter interface {
	UserPools(namespace string) UserPoolInterface
}

// UserPoolInterface has methods to work with UserPool resources.
type UserPoolInterface interface {
	Create(*v1alpha1.UserPool) (*v1alpha1.UserPool, error)
	Update(*v1alpha1.UserPool) (*v1alpha1.UserPool, error)
	UpdateStatus(*v1alpha1.UserPool) (*v1alpha1.UserPool, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.UserPool, error)
	List(opts v1.ListOptions) (*v1alpha1.UserPoolList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UserPool, err error)
	UserPoolExpansion
}

// userPools implements UserPoolInterface
type userPools struct {
	client rest.Interface
	ns     string
}

// newUserPools returns a UserPools
func newUserPools(c *CognitoV1alpha1Client, namespace string) *userPools {
	return &userPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the userPool, and returns the corresponding userPool object, and an error if there is any.
func (c *userPools) Get(name string, options v1.GetOptions) (result *v1alpha1.UserPool, err error) {
	result = &v1alpha1.UserPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userpools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserPools that match those selectors.
func (c *userPools) List(opts v1.ListOptions) (result *v1alpha1.UserPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.UserPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userPools.
func (c *userPools) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("userpools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a userPool and creates it.  Returns the server's representation of the userPool, and an error, if there is any.
func (c *userPools) Create(userPool *v1alpha1.UserPool) (result *v1alpha1.UserPool, err error) {
	result = &v1alpha1.UserPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("userpools").
		Body(userPool).
		Do().
		Into(result)
	return
}

// Update takes the representation of a userPool and updates it. Returns the server's representation of the userPool, and an error, if there is any.
func (c *userPools) Update(userPool *v1alpha1.UserPool) (result *v1alpha1.UserPool, err error) {
	result = &v1alpha1.UserPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userpools").
		Name(userPool.Name).
		Body(userPool).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *userPools) UpdateStatus(userPool *v1alpha1.UserPool) (result *v1alpha1.UserPool, err error) {
	result = &v1alpha1.UserPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userpools").
		Name(userPool.Name).
		SubResource("status").
		Body(userPool).
		Do().
		Into(result)
	return
}

// Delete takes name of the userPool and deletes it. Returns an error if one occurs.
func (c *userPools) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userpools").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userPools) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userpools").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched userPool.
func (c *userPools) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UserPool, err error) {
	result = &v1alpha1.UserPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("userpools").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
