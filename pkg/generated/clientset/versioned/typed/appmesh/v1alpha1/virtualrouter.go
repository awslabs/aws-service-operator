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

	v1alpha1 "awsoperator.io/pkg/apis/appmesh/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualRoutersGetter has a method to return a VirtualRouterInterface.
// A group's client should implement this interface.
type VirtualRoutersGetter interface {
	VirtualRouters(namespace string) VirtualRouterInterface
}

// VirtualRouterInterface has methods to work with VirtualRouter resources.
type VirtualRouterInterface interface {
	Create(*v1alpha1.VirtualRouter) (*v1alpha1.VirtualRouter, error)
	Update(*v1alpha1.VirtualRouter) (*v1alpha1.VirtualRouter, error)
	UpdateStatus(*v1alpha1.VirtualRouter) (*v1alpha1.VirtualRouter, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VirtualRouter, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualRouterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualRouter, err error)
	VirtualRouterExpansion
}

// virtualRouters implements VirtualRouterInterface
type virtualRouters struct {
	client rest.Interface
	ns     string
}

// newVirtualRouters returns a VirtualRouters
func newVirtualRouters(c *AppmeshV1alpha1Client, namespace string) *virtualRouters {
	return &virtualRouters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualRouter, and returns the corresponding virtualRouter object, and an error if there is any.
func (c *virtualRouters) Get(name string, options v1.GetOptions) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VirtualRouters that match those selectors.
func (c *virtualRouters) List(opts v1.ListOptions) (result *v1alpha1.VirtualRouterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualRouterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualRouters.
func (c *virtualRouters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualRouter and creates it.  Returns the server's representation of the virtualRouter, and an error, if there is any.
func (c *virtualRouters) Create(virtualRouter *v1alpha1.VirtualRouter) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualrouters").
		Body(virtualRouter).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualRouter and updates it. Returns the server's representation of the virtualRouter, and an error, if there is any.
func (c *virtualRouters) Update(virtualRouter *v1alpha1.VirtualRouter) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(virtualRouter.Name).
		Body(virtualRouter).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualRouters) UpdateStatus(virtualRouter *v1alpha1.VirtualRouter) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(virtualRouter.Name).
		SubResource("status").
		Body(virtualRouter).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualRouter and deletes it. Returns an error if one occurs.
func (c *virtualRouters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualrouters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualRouters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualrouters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualRouter.
func (c *virtualRouters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VirtualRouter, err error) {
	result = &v1alpha1.VirtualRouter{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualrouters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
