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

	v1alpha1 "awsoperator.io/pkg/apis/robomaker/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SimulationApplicationVersionsGetter has a method to return a SimulationApplicationVersionInterface.
// A group's client should implement this interface.
type SimulationApplicationVersionsGetter interface {
	SimulationApplicationVersions(namespace string) SimulationApplicationVersionInterface
}

// SimulationApplicationVersionInterface has methods to work with SimulationApplicationVersion resources.
type SimulationApplicationVersionInterface interface {
	Create(*v1alpha1.SimulationApplicationVersion) (*v1alpha1.SimulationApplicationVersion, error)
	Update(*v1alpha1.SimulationApplicationVersion) (*v1alpha1.SimulationApplicationVersion, error)
	UpdateStatus(*v1alpha1.SimulationApplicationVersion) (*v1alpha1.SimulationApplicationVersion, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SimulationApplicationVersion, error)
	List(opts v1.ListOptions) (*v1alpha1.SimulationApplicationVersionList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SimulationApplicationVersion, err error)
	SimulationApplicationVersionExpansion
}

// simulationApplicationVersions implements SimulationApplicationVersionInterface
type simulationApplicationVersions struct {
	client rest.Interface
	ns     string
}

// newSimulationApplicationVersions returns a SimulationApplicationVersions
func newSimulationApplicationVersions(c *RobomakerV1alpha1Client, namespace string) *simulationApplicationVersions {
	return &simulationApplicationVersions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the simulationApplicationVersion, and returns the corresponding simulationApplicationVersion object, and an error if there is any.
func (c *simulationApplicationVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.SimulationApplicationVersion, err error) {
	result = &v1alpha1.SimulationApplicationVersion{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SimulationApplicationVersions that match those selectors.
func (c *simulationApplicationVersions) List(opts v1.ListOptions) (result *v1alpha1.SimulationApplicationVersionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SimulationApplicationVersionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested simulationApplicationVersions.
func (c *simulationApplicationVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a simulationApplicationVersion and creates it.  Returns the server's representation of the simulationApplicationVersion, and an error, if there is any.
func (c *simulationApplicationVersions) Create(simulationApplicationVersion *v1alpha1.SimulationApplicationVersion) (result *v1alpha1.SimulationApplicationVersion, err error) {
	result = &v1alpha1.SimulationApplicationVersion{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		Body(simulationApplicationVersion).
		Do().
		Into(result)
	return
}

// Update takes the representation of a simulationApplicationVersion and updates it. Returns the server's representation of the simulationApplicationVersion, and an error, if there is any.
func (c *simulationApplicationVersions) Update(simulationApplicationVersion *v1alpha1.SimulationApplicationVersion) (result *v1alpha1.SimulationApplicationVersion, err error) {
	result = &v1alpha1.SimulationApplicationVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		Name(simulationApplicationVersion.Name).
		Body(simulationApplicationVersion).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *simulationApplicationVersions) UpdateStatus(simulationApplicationVersion *v1alpha1.SimulationApplicationVersion) (result *v1alpha1.SimulationApplicationVersion, err error) {
	result = &v1alpha1.SimulationApplicationVersion{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		Name(simulationApplicationVersion.Name).
		SubResource("status").
		Body(simulationApplicationVersion).
		Do().
		Into(result)
	return
}

// Delete takes name of the simulationApplicationVersion and deletes it. Returns an error if one occurs.
func (c *simulationApplicationVersions) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *simulationApplicationVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched simulationApplicationVersion.
func (c *simulationApplicationVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SimulationApplicationVersion, err error) {
	result = &v1alpha1.SimulationApplicationVersion{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("simulationapplicationversions").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
