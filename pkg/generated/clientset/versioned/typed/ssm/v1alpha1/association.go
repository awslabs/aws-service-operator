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

	v1alpha1 "awsoperator.io/pkg/apis/ssm/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AssociationsGetter has a method to return a AssociationInterface.
// A group's client should implement this interface.
type AssociationsGetter interface {
	Associations(namespace string) AssociationInterface
}

// AssociationInterface has methods to work with Association resources.
type AssociationInterface interface {
	Create(*v1alpha1.Association) (*v1alpha1.Association, error)
	Update(*v1alpha1.Association) (*v1alpha1.Association, error)
	UpdateStatus(*v1alpha1.Association) (*v1alpha1.Association, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Association, error)
	List(opts v1.ListOptions) (*v1alpha1.AssociationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Association, err error)
	AssociationExpansion
}

// associations implements AssociationInterface
type associations struct {
	client rest.Interface
	ns     string
}

// newAssociations returns a Associations
func newAssociations(c *SsmV1alpha1Client, namespace string) *associations {
	return &associations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the association, and returns the corresponding association object, and an error if there is any.
func (c *associations) Get(name string, options v1.GetOptions) (result *v1alpha1.Association, err error) {
	result = &v1alpha1.Association{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("associations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Associations that match those selectors.
func (c *associations) List(opts v1.ListOptions) (result *v1alpha1.AssociationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AssociationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("associations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested associations.
func (c *associations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("associations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a association and creates it.  Returns the server's representation of the association, and an error, if there is any.
func (c *associations) Create(association *v1alpha1.Association) (result *v1alpha1.Association, err error) {
	result = &v1alpha1.Association{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("associations").
		Body(association).
		Do().
		Into(result)
	return
}

// Update takes the representation of a association and updates it. Returns the server's representation of the association, and an error, if there is any.
func (c *associations) Update(association *v1alpha1.Association) (result *v1alpha1.Association, err error) {
	result = &v1alpha1.Association{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("associations").
		Name(association.Name).
		Body(association).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *associations) UpdateStatus(association *v1alpha1.Association) (result *v1alpha1.Association, err error) {
	result = &v1alpha1.Association{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("associations").
		Name(association.Name).
		SubResource("status").
		Body(association).
		Do().
		Into(result)
	return
}

// Delete takes name of the association and deletes it. Returns an error if one occurs.
func (c *associations) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("associations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *associations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("associations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched association.
func (c *associations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Association, err error) {
	result = &v1alpha1.Association{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("associations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
