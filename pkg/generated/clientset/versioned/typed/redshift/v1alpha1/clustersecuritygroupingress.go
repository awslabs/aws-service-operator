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

	v1alpha1 "awsoperator.io/pkg/apis/redshift/v1alpha1"
	scheme "awsoperator.io/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ClusterSecurityGroupIngressesGetter has a method to return a ClusterSecurityGroupIngressInterface.
// A group's client should implement this interface.
type ClusterSecurityGroupIngressesGetter interface {
	ClusterSecurityGroupIngresses(namespace string) ClusterSecurityGroupIngressInterface
}

// ClusterSecurityGroupIngressInterface has methods to work with ClusterSecurityGroupIngress resources.
type ClusterSecurityGroupIngressInterface interface {
	Create(*v1alpha1.ClusterSecurityGroupIngress) (*v1alpha1.ClusterSecurityGroupIngress, error)
	Update(*v1alpha1.ClusterSecurityGroupIngress) (*v1alpha1.ClusterSecurityGroupIngress, error)
	UpdateStatus(*v1alpha1.ClusterSecurityGroupIngress) (*v1alpha1.ClusterSecurityGroupIngress, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterSecurityGroupIngress, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterSecurityGroupIngressList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterSecurityGroupIngress, err error)
	ClusterSecurityGroupIngressExpansion
}

// clusterSecurityGroupIngresses implements ClusterSecurityGroupIngressInterface
type clusterSecurityGroupIngresses struct {
	client rest.Interface
	ns     string
}

// newClusterSecurityGroupIngresses returns a ClusterSecurityGroupIngresses
func newClusterSecurityGroupIngresses(c *RedshiftV1alpha1Client, namespace string) *clusterSecurityGroupIngresses {
	return &clusterSecurityGroupIngresses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterSecurityGroupIngress, and returns the corresponding clusterSecurityGroupIngress object, and an error if there is any.
func (c *clusterSecurityGroupIngresses) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterSecurityGroupIngress, err error) {
	result = &v1alpha1.ClusterSecurityGroupIngress{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterSecurityGroupIngresses that match those selectors.
func (c *clusterSecurityGroupIngresses) List(opts v1.ListOptions) (result *v1alpha1.ClusterSecurityGroupIngressList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ClusterSecurityGroupIngressList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterSecurityGroupIngresses.
func (c *clusterSecurityGroupIngresses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a clusterSecurityGroupIngress and creates it.  Returns the server's representation of the clusterSecurityGroupIngress, and an error, if there is any.
func (c *clusterSecurityGroupIngresses) Create(clusterSecurityGroupIngress *v1alpha1.ClusterSecurityGroupIngress) (result *v1alpha1.ClusterSecurityGroupIngress, err error) {
	result = &v1alpha1.ClusterSecurityGroupIngress{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		Body(clusterSecurityGroupIngress).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterSecurityGroupIngress and updates it. Returns the server's representation of the clusterSecurityGroupIngress, and an error, if there is any.
func (c *clusterSecurityGroupIngresses) Update(clusterSecurityGroupIngress *v1alpha1.ClusterSecurityGroupIngress) (result *v1alpha1.ClusterSecurityGroupIngress, err error) {
	result = &v1alpha1.ClusterSecurityGroupIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		Name(clusterSecurityGroupIngress.Name).
		Body(clusterSecurityGroupIngress).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *clusterSecurityGroupIngresses) UpdateStatus(clusterSecurityGroupIngress *v1alpha1.ClusterSecurityGroupIngress) (result *v1alpha1.ClusterSecurityGroupIngress, err error) {
	result = &v1alpha1.ClusterSecurityGroupIngress{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		Name(clusterSecurityGroupIngress.Name).
		SubResource("status").
		Body(clusterSecurityGroupIngress).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterSecurityGroupIngress and deletes it. Returns an error if one occurs.
func (c *clusterSecurityGroupIngresses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterSecurityGroupIngresses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterSecurityGroupIngress.
func (c *clusterSecurityGroupIngresses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterSecurityGroupIngress, err error) {
	result = &v1alpha1.ClusterSecurityGroupIngress{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clustersecuritygroupingresses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
