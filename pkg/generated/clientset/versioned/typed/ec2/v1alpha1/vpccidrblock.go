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

	v1alpha1 "awsoperator.io/pkg/apis/ec2/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VPCCidrBlocksGetter has a method to return a VPCCidrBlockInterface.
// A group's client should implement this interface.
type VPCCidrBlocksGetter interface {
	VPCCidrBlocks(namespace string) VPCCidrBlockInterface
}

// VPCCidrBlockInterface has methods to work with VPCCidrBlock resources.
type VPCCidrBlockInterface interface {
	Create(*v1alpha1.VPCCidrBlock) (*v1alpha1.VPCCidrBlock, error)
	Update(*v1alpha1.VPCCidrBlock) (*v1alpha1.VPCCidrBlock, error)
	UpdateStatus(*v1alpha1.VPCCidrBlock) (*v1alpha1.VPCCidrBlock, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.VPCCidrBlock, error)
	List(opts v1.ListOptions) (*v1alpha1.VPCCidrBlockList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VPCCidrBlock, err error)
	VPCCidrBlockExpansion
}

// vPCCidrBlocks implements VPCCidrBlockInterface
type vPCCidrBlocks struct {
	client rest.Interface
	ns     string
}

// newVPCCidrBlocks returns a VPCCidrBlocks
func newVPCCidrBlocks(c *Ec2V1alpha1Client, namespace string) *vPCCidrBlocks {
	return &vPCCidrBlocks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the vPCCidrBlock, and returns the corresponding vPCCidrBlock object, and an error if there is any.
func (c *vPCCidrBlocks) Get(name string, options v1.GetOptions) (result *v1alpha1.VPCCidrBlock, err error) {
	result = &v1alpha1.VPCCidrBlock{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of VPCCidrBlocks that match those selectors.
func (c *vPCCidrBlocks) List(opts v1.ListOptions) (result *v1alpha1.VPCCidrBlockList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VPCCidrBlockList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested vPCCidrBlocks.
func (c *vPCCidrBlocks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a vPCCidrBlock and creates it.  Returns the server's representation of the vPCCidrBlock, and an error, if there is any.
func (c *vPCCidrBlocks) Create(vPCCidrBlock *v1alpha1.VPCCidrBlock) (result *v1alpha1.VPCCidrBlock, err error) {
	result = &v1alpha1.VPCCidrBlock{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		Body(vPCCidrBlock).
		Do().
		Into(result)
	return
}

// Update takes the representation of a vPCCidrBlock and updates it. Returns the server's representation of the vPCCidrBlock, and an error, if there is any.
func (c *vPCCidrBlocks) Update(vPCCidrBlock *v1alpha1.VPCCidrBlock) (result *v1alpha1.VPCCidrBlock, err error) {
	result = &v1alpha1.VPCCidrBlock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		Name(vPCCidrBlock.Name).
		Body(vPCCidrBlock).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *vPCCidrBlocks) UpdateStatus(vPCCidrBlock *v1alpha1.VPCCidrBlock) (result *v1alpha1.VPCCidrBlock, err error) {
	result = &v1alpha1.VPCCidrBlock{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		Name(vPCCidrBlock.Name).
		SubResource("status").
		Body(vPCCidrBlock).
		Do().
		Into(result)
	return
}

// Delete takes name of the vPCCidrBlock and deletes it. Returns an error if one occurs.
func (c *vPCCidrBlocks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *vPCCidrBlocks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("vpccidrblocks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched vPCCidrBlock.
func (c *vPCCidrBlocks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VPCCidrBlock, err error) {
	result = &v1alpha1.VPCCidrBlock{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("vpccidrblocks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
