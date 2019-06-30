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

	v1alpha1 "awsoperator.io/pkg/apis/logs/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// LogStreamsGetter has a method to return a LogStreamInterface.
// A group's client should implement this interface.
type LogStreamsGetter interface {
	LogStreams(namespace string) LogStreamInterface
}

// LogStreamInterface has methods to work with LogStream resources.
type LogStreamInterface interface {
	Create(*v1alpha1.LogStream) (*v1alpha1.LogStream, error)
	Update(*v1alpha1.LogStream) (*v1alpha1.LogStream, error)
	UpdateStatus(*v1alpha1.LogStream) (*v1alpha1.LogStream, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.LogStream, error)
	List(opts v1.ListOptions) (*v1alpha1.LogStreamList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogStream, err error)
	LogStreamExpansion
}

// logStreams implements LogStreamInterface
type logStreams struct {
	client rest.Interface
	ns     string
}

// newLogStreams returns a LogStreams
func newLogStreams(c *LogsV1alpha1Client, namespace string) *logStreams {
	return &logStreams{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the logStream, and returns the corresponding logStream object, and an error if there is any.
func (c *logStreams) Get(name string, options v1.GetOptions) (result *v1alpha1.LogStream, err error) {
	result = &v1alpha1.LogStream{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("logstreams").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LogStreams that match those selectors.
func (c *logStreams) List(opts v1.ListOptions) (result *v1alpha1.LogStreamList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.LogStreamList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("logstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested logStreams.
func (c *logStreams) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("logstreams").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a logStream and creates it.  Returns the server's representation of the logStream, and an error, if there is any.
func (c *logStreams) Create(logStream *v1alpha1.LogStream) (result *v1alpha1.LogStream, err error) {
	result = &v1alpha1.LogStream{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("logstreams").
		Body(logStream).
		Do().
		Into(result)
	return
}

// Update takes the representation of a logStream and updates it. Returns the server's representation of the logStream, and an error, if there is any.
func (c *logStreams) Update(logStream *v1alpha1.LogStream) (result *v1alpha1.LogStream, err error) {
	result = &v1alpha1.LogStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("logstreams").
		Name(logStream.Name).
		Body(logStream).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *logStreams) UpdateStatus(logStream *v1alpha1.LogStream) (result *v1alpha1.LogStream, err error) {
	result = &v1alpha1.LogStream{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("logstreams").
		Name(logStream.Name).
		SubResource("status").
		Body(logStream).
		Do().
		Into(result)
	return
}

// Delete takes name of the logStream and deletes it. Returns an error if one occurs.
func (c *logStreams) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("logstreams").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *logStreams) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("logstreams").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched logStream.
func (c *logStreams) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LogStream, err error) {
	result = &v1alpha1.LogStream{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("logstreams").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
