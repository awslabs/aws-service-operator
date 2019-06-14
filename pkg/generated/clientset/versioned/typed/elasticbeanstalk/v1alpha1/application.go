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

	v1alpha1 "awsoperator.io/pkg/apis/elasticbeanstalk/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationsGetter has a method to return a ApplicationInterface.
// A group's client should implement this interface.
type ApplicationsGetter interface {
	Applications(namespace string) ApplicationInterface
}

// ApplicationInterface has methods to work with Application resources.
type ApplicationInterface interface {
	Create(*v1alpha1.Application) (*v1alpha1.Application, error)
	Update(*v1alpha1.Application) (*v1alpha1.Application, error)
	UpdateStatus(*v1alpha1.Application) (*v1alpha1.Application, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Application, error)
	List(opts v1.ListOptions) (*v1alpha1.ApplicationList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Application, err error)
	ApplicationExpansion
}

// applications implements ApplicationInterface
type applications struct {
	client rest.Interface
	ns     string
}

// newApplications returns a Applications
func newApplications(c *ElasticbeanstalkV1alpha1Client, namespace string) *applications {
	return &applications{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the application, and returns the corresponding application object, and an error if there is any.
func (c *applications) Get(name string, options v1.GetOptions) (result *v1alpha1.Application, err error) {
	result = &v1alpha1.Application{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Applications that match those selectors.
func (c *applications) List(opts v1.ListOptions) (result *v1alpha1.ApplicationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ApplicationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applications.
func (c *applications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a application and creates it.  Returns the server's representation of the application, and an error, if there is any.
func (c *applications) Create(application *v1alpha1.Application) (result *v1alpha1.Application, err error) {
	result = &v1alpha1.Application{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applications").
		Body(application).
		Do().
		Into(result)
	return
}

// Update takes the representation of a application and updates it. Returns the server's representation of the application, and an error, if there is any.
func (c *applications) Update(application *v1alpha1.Application) (result *v1alpha1.Application, err error) {
	result = &v1alpha1.Application{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applications").
		Name(application.Name).
		Body(application).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *applications) UpdateStatus(application *v1alpha1.Application) (result *v1alpha1.Application, err error) {
	result = &v1alpha1.Application{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applications").
		Name(application.Name).
		SubResource("status").
		Body(application).
		Do().
		Into(result)
	return
}

// Delete takes name of the application and deletes it. Returns an error if one occurs.
func (c *applications) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applications").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applications").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched application.
func (c *applications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Application, err error) {
	result = &v1alpha1.Application{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applications").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
