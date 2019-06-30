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

	v1alpha1 "awsoperator.io/pkg/apis/opsworks/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VolumesGetter has a method to return a VolumeInterface.
// A group's client should implement this interface.
type VolumesGetter interface {
	Volumes(namespace string) VolumeInterface
}

// VolumeInterface has methods to work with Volume resources.
type VolumeInterface interface {
	Create(*v1alpha1.Volume) (*v1alpha1.Volume, error)
	Update(*v1alpha1.Volume) (*v1alpha1.Volume, error)
	UpdateStatus(*v1alpha1.Volume) (*v1alpha1.Volume, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Volume, error)
	List(opts v1.ListOptions) (*v1alpha1.VolumeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Volume, err error)
	VolumeExpansion
}

// volumes implements VolumeInterface
type volumes struct {
	client rest.Interface
	ns     string
}

// newVolumes returns a Volumes
func newVolumes(c *OpsworksV1alpha1Client, namespace string) *volumes {
	return &volumes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the volume, and returns the corresponding volume object, and an error if there is any.
func (c *volumes) Get(name string, options v1.GetOptions) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Volumes that match those selectors.
func (c *volumes) List(opts v1.ListOptions) (result *v1alpha1.VolumeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VolumeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested volumes.
func (c *volumes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a volume and creates it.  Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Create(volume *v1alpha1.Volume) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("volumes").
		Body(volume).
		Do().
		Into(result)
	return
}

// Update takes the representation of a volume and updates it. Returns the server's representation of the volume, and an error, if there is any.
func (c *volumes) Update(volume *v1alpha1.Volume) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumes").
		Name(volume.Name).
		Body(volume).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *volumes) UpdateStatus(volume *v1alpha1.Volume) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("volumes").
		Name(volume.Name).
		SubResource("status").
		Body(volume).
		Do().
		Into(result)
	return
}

// Delete takes name of the volume and deletes it. Returns an error if one occurs.
func (c *volumes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *volumes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("volumes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched volume.
func (c *volumes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Volume, err error) {
	result = &v1alpha1.Volume{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("volumes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
