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

	v1alpha1 "awsoperator.io/pkg/apis/greengrass/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DeviceDefinitionVersionsGetter has a method to return a DeviceDefinitionVersionInterface.
// A group's client should implement this interface.
type DeviceDefinitionVersionsGetter interface {
	DeviceDefinitionVersions(namespace string) DeviceDefinitionVersionInterface
}

// DeviceDefinitionVersionInterface has methods to work with DeviceDefinitionVersion resources.
type DeviceDefinitionVersionInterface interface {
	Create(*v1alpha1.DeviceDefinitionVersion) (*v1alpha1.DeviceDefinitionVersion, error)
	Update(*v1alpha1.DeviceDefinitionVersion) (*v1alpha1.DeviceDefinitionVersion, error)
	UpdateStatus(*v1alpha1.DeviceDefinitionVersion) (*v1alpha1.DeviceDefinitionVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.DeviceDefinitionVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.DeviceDefinitionVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceDefinitionVersion, err error)
	DeviceDefinitionVersionExpansion
}

// deviceDefinitionVersions implements DeviceDefinitionVersionInterface
type deviceDefinitionVersions struct {
	client rest.Interface
	ns     string
}

// newDeviceDefinitionVersions returns a DeviceDefinitionVersions
func newDeviceDefinitionVersions(c *GreengrassV1alpha1Client, namespace string) *deviceDefinitionVersions {
	return &deviceDefinitionVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the deviceDefinitionVersion, and returns the corresponding deviceDefinitionVersion object, and an error if there is any.
func (c *deviceDefinitionVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.DeviceDefinitionVersion, err error) {
	result = &v1alpha1.DeviceDefinitionVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DeviceDefinitionVersions that match those selectors.
func (c *deviceDefinitionVersions) List(opts v1.ListOptions) (result *v1alpha1.DeviceDefinitionVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DeviceDefinitionVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested deviceDefinitionVersions.
func (c *deviceDefinitionVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a deviceDefinitionVersion and creates it.  Returns the server's representation of the deviceDefinitionVersion, and an error, if there is any.
func (c *deviceDefinitionVersions) Create(deviceDefinitionVersion *v1alpha1.DeviceDefinitionVersion) (result *v1alpha1.DeviceDefinitionVersion, err error) {
	result = &v1alpha1.DeviceDefinitionVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		Body(deviceDefinitionVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a deviceDefinitionVersion and updates it. Returns the server's representation of the deviceDefinitionVersion, and an error, if there is any.
func (c *deviceDefinitionVersions) Update(deviceDefinitionVersion *v1alpha1.DeviceDefinitionVersion) (result *v1alpha1.DeviceDefinitionVersion, err error) {
	result = &v1alpha1.DeviceDefinitionVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		Name(deviceDefinitionVersion.Name).
		Body(deviceDefinitionVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *deviceDefinitionVersions) UpdateStatus(deviceDefinitionVersion *v1alpha1.DeviceDefinitionVersion) (result *v1alpha1.DeviceDefinitionVersion, err error) {
	result = &v1alpha1.DeviceDefinitionVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		Name(deviceDefinitionVersion.Name).
		SubResource("status").
		Body(deviceDefinitionVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the deviceDefinitionVersion and deletes it. Returns an error if one occurs.
func (c *deviceDefinitionVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *deviceDefinitionVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched deviceDefinitionVersion.
func (c *deviceDefinitionVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DeviceDefinitionVersion, err error) {
	result = &v1alpha1.DeviceDefinitionVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("devicedefinitionversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
