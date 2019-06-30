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

	v1alpha1 "awsoperator.io/pkg/apis/kinesisanalyticsv2/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationReferenceDataSourcesGetter has a method to return a ApplicationReferenceDataSourceInterface.
// A group's client should implement this interface.
type ApplicationReferenceDataSourcesGetter interface {
	ApplicationReferenceDataSources(namespace string) ApplicationReferenceDataSourceInterface
}

// ApplicationReferenceDataSourceInterface has methods to work with ApplicationReferenceDataSource resources.
type ApplicationReferenceDataSourceInterface interface {
	Create(*v1alpha1.ApplicationReferenceDataSource) (*v1alpha1.ApplicationReferenceDataSource, error)
	Update(*v1alpha1.ApplicationReferenceDataSource) (*v1alpha1.ApplicationReferenceDataSource, error)
	UpdateStatus(*v1alpha1.ApplicationReferenceDataSource) (*v1alpha1.ApplicationReferenceDataSource, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ApplicationReferenceDataSource, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationReferenceDataSourceList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationReferenceDataSource, err error)
	ApplicationReferenceDataSourceExpansion
}

// applicationReferenceDataSources implements ApplicationReferenceDataSourceInterface
type applicationReferenceDataSources struct {
	client rest.Interface
	ns     string
}

// newApplicationReferenceDataSources returns a ApplicationReferenceDataSources
func newApplicationReferenceDataSources(c *Kinesisanalyticsv2V1alpha1Client, namespace string) *applicationReferenceDataSources {
	return &applicationReferenceDataSources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationReferenceDataSource, and returns the corresponding applicationReferenceDataSource object, and an error if there is any.
func (c *applicationReferenceDataSources) Get(name string, options v1.GetOptions) (result *v1alpha1.ApplicationReferenceDataSource, err error) {
	result = &v1alpha1.ApplicationReferenceDataSource{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationReferenceDataSources that match those selectors.
func (c *applicationReferenceDataSources) List(opts v1.ListOptions) (result *v1alpha1.ApplicationReferenceDataSourceList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationReferenceDataSourceList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationReferenceDataSources.
func (c *applicationReferenceDataSources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a applicationReferenceDataSource and creates it.  Returns the server's representation of the applicationReferenceDataSource, and an error, if there is any.
func (c *applicationReferenceDataSources) Create(applicationReferenceDataSource *v1alpha1.ApplicationReferenceDataSource) (result *v1alpha1.ApplicationReferenceDataSource, err error) {
	result = &v1alpha1.ApplicationReferenceDataSource{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		Body(applicationReferenceDataSource).
		Do().
		Into(result)
	return
}

// Update takes the representation of a applicationReferenceDataSource and updates it. Returns the server's representation of the applicationReferenceDataSource, and an error, if there is any.
func (c *applicationReferenceDataSources) Update(applicationReferenceDataSource *v1alpha1.ApplicationReferenceDataSource) (result *v1alpha1.ApplicationReferenceDataSource, err error) {
	result = &v1alpha1.ApplicationReferenceDataSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		Name(applicationReferenceDataSource.Name).
		Body(applicationReferenceDataSource).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applicationReferenceDataSources) UpdateStatus(applicationReferenceDataSource *v1alpha1.ApplicationReferenceDataSource) (result *v1alpha1.ApplicationReferenceDataSource, err error) {
	result = &v1alpha1.ApplicationReferenceDataSource{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		Name(applicationReferenceDataSource.Name).
		SubResource("status").
		Body(applicationReferenceDataSource).
		Do().
		Into(result)
	return
}

// Delete takes name of the applicationReferenceDataSource and deletes it. Returns an error if one occurs.
func (c *applicationReferenceDataSources) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationReferenceDataSources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched applicationReferenceDataSource.
func (c *applicationReferenceDataSources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ApplicationReferenceDataSource, err error) {
	result = &v1alpha1.ApplicationReferenceDataSource{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationreferencedatasources").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
