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

	v1alpha1 "awsoperator.io/pkg/apis/inspector/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ResourceGroupsGetter has a method to return a ResourceGroupInterface.
// A group's client should implement this interface.
type ResourceGroupsGetter interface {
	ResourceGroups(namespace string) ResourceGroupInterface
}

// ResourceGroupInterface has methods to work with ResourceGroup resources.
type ResourceGroupInterface interface {
	Create(*v1alpha1.ResourceGroup) (*v1alpha1.ResourceGroup, error)
	Update(*v1alpha1.ResourceGroup) (*v1alpha1.ResourceGroup, error)
	UpdateStatus(*v1alpha1.ResourceGroup) (*v1alpha1.ResourceGroup, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ResourceGroup, error)
	List(opts v1.ListOptions) (*v1alpha1.ResourceGroupList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceGroup, err error)
	ResourceGroupExpansion
}

// resourceGroups implements ResourceGroupInterface
type resourceGroups struct {
	client rest.Interface
	ns     string
}

// newResourceGroups returns a ResourceGroups
func newResourceGroups(c *InspectorV1alpha1Client, namespace string) *resourceGroups {
	return &resourceGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the resourceGroup, and returns the corresponding resourceGroup object, and an error if there is any.
func (c *resourceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceGroup, err error) {
	result = &v1alpha1.ResourceGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ResourceGroups that match those selectors.
func (c *resourceGroups) List(opts v1.ListOptions) (result *v1alpha1.ResourceGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ResourceGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("resourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested resourceGroups.
func (c *resourceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("resourcegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a resourceGroup and creates it.  Returns the server's representation of the resourceGroup, and an error, if there is any.
func (c *resourceGroups) Create(resourceGroup *v1alpha1.ResourceGroup) (result *v1alpha1.ResourceGroup, err error) {
	result = &v1alpha1.ResourceGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("resourcegroups").
		Body(resourceGroup).
		Do().
		Into(result)
	return
}

// Update takes the representation of a resourceGroup and updates it. Returns the server's representation of the resourceGroup, and an error, if there is any.
func (c *resourceGroups) Update(resourceGroup *v1alpha1.ResourceGroup) (result *v1alpha1.ResourceGroup, err error) {
	result = &v1alpha1.ResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcegroups").
		Name(resourceGroup.Name).
		Body(resourceGroup).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *resourceGroups) UpdateStatus(resourceGroup *v1alpha1.ResourceGroup) (result *v1alpha1.ResourceGroup, err error) {
	result = &v1alpha1.ResourceGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("resourcegroups").
		Name(resourceGroup.Name).
		SubResource("status").
		Body(resourceGroup).
		Do().
		Into(result)
	return
}

// Delete takes name of the resourceGroup and deletes it. Returns an error if one occurs.
func (c *resourceGroups) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcegroups").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *resourceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("resourcegroups").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched resourceGroup.
func (c *resourceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceGroup, err error) {
	result = &v1alpha1.ResourceGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("resourcegroups").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
