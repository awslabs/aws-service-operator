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

	v1alpha1 "awsoperator.io/pkg/apis/route53/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RecordSetsGetter has a method to return a RecordSetInterface.
// A group's client should implement this interface.
type RecordSetsGetter interface {
	RecordSets(namespace string) RecordSetInterface
}

// RecordSetInterface has methods to work with RecordSet resources.
type RecordSetInterface interface {
	Create(*v1alpha1.RecordSet) (*v1alpha1.RecordSet, error)
	Update(*v1alpha1.RecordSet) (*v1alpha1.RecordSet, error)
	UpdateStatus(*v1alpha1.RecordSet) (*v1alpha1.RecordSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RecordSet, error)
	List(opts v1.ListOptions) (*v1alpha1.RecordSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RecordSet, err error)
	RecordSetExpansion
}

// recordSets implements RecordSetInterface
type recordSets struct {
	client rest.Interface
	ns     string
}

// newRecordSets returns a RecordSets
func newRecordSets(c *Route53V1alpha1Client, namespace string) *recordSets {
	return &recordSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the recordSet, and returns the corresponding recordSet object, and an error if there is any.
func (c *recordSets) Get(name string, options v1.GetOptions) (result *v1alpha1.RecordSet, err error) {
	result = &v1alpha1.RecordSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("recordsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RecordSets that match those selectors.
func (c *recordSets) List(opts v1.ListOptions) (result *v1alpha1.RecordSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RecordSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("recordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested recordSets.
func (c *recordSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("recordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a recordSet and creates it.  Returns the server's representation of the recordSet, and an error, if there is any.
func (c *recordSets) Create(recordSet *v1alpha1.RecordSet) (result *v1alpha1.RecordSet, err error) {
	result = &v1alpha1.RecordSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("recordsets").
		Body(recordSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a recordSet and updates it. Returns the server's representation of the recordSet, and an error, if there is any.
func (c *recordSets) Update(recordSet *v1alpha1.RecordSet) (result *v1alpha1.RecordSet, err error) {
	result = &v1alpha1.RecordSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("recordsets").
		Name(recordSet.Name).
		Body(recordSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *recordSets) UpdateStatus(recordSet *v1alpha1.RecordSet) (result *v1alpha1.RecordSet, err error) {
	result = &v1alpha1.RecordSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("recordsets").
		Name(recordSet.Name).
		SubResource("status").
		Body(recordSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the recordSet and deletes it. Returns an error if one occurs.
func (c *recordSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("recordsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *recordSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("recordsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched recordSet.
func (c *recordSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RecordSet, err error) {
	result = &v1alpha1.RecordSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("recordsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
