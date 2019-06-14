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

	v1alpha1 "awsoperator.io/pkg/apis/iotanalytics/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DatasetsGetter has a method to return a DatasetInterface.
// A group's client should implement this interface.
type DatasetsGetter interface {
	Datasets(namespace string) DatasetInterface
}

// DatasetInterface has methods to work with Dataset resources.
type DatasetInterface interface {
	Create(*v1alpha1.Dataset) (*v1alpha1.Dataset, error)
	Update(*v1alpha1.Dataset) (*v1alpha1.Dataset, error)
	UpdateStatus(*v1alpha1.Dataset) (*v1alpha1.Dataset, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Dataset, error)
	List(opts v1.ListOptions) (*v1alpha1.DatasetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Dataset, err error)
	DatasetExpansion
}

// datasets implements DatasetInterface
type datasets struct {
	client rest.Interface
	ns     string
}

// newDatasets returns a Datasets
func newDatasets(c *IotanalyticsV1alpha1Client, namespace string) *datasets {
	return &datasets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the dataset, and returns the corresponding dataset object, and an error if there is any.
func (c *datasets) Get(name string, options v1.GetOptions) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Datasets that match those selectors.
func (c *datasets) List(opts v1.ListOptions) (result *v1alpha1.DatasetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DatasetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested datasets.
func (c *datasets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a dataset and creates it.  Returns the server's representation of the dataset, and an error, if there is any.
func (c *datasets) Create(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("datasets").
		Body(dataset).
		Do().
		Into(result)
	return
}

// Update takes the representation of a dataset and updates it. Returns the server's representation of the dataset, and an error, if there is any.
func (c *datasets) Update(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasets").
		Name(dataset.Name).
		Body(dataset).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *datasets) UpdateStatus(dataset *v1alpha1.Dataset) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("datasets").
		Name(dataset.Name).
		SubResource("status").
		Body(dataset).
		Do().
		Into(result)
	return
}

// Delete takes name of the dataset and deletes it. Returns an error if one occurs.
func (c *datasets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *datasets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("datasets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched dataset.
func (c *datasets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Dataset, err error) {
	result = &v1alpha1.Dataset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("datasets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
