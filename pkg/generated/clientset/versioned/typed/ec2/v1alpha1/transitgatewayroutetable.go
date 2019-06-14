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

// TransitGatewayRouteTablesGetter has a method to return a TransitGatewayRouteTableInterface.
// A group's client should implement this interface.
type TransitGatewayRouteTablesGetter interface {
	TransitGatewayRouteTables(namespace string) TransitGatewayRouteTableInterface
}

// TransitGatewayRouteTableInterface has methods to work with TransitGatewayRouteTable resources.
type TransitGatewayRouteTableInterface interface {
	Create(*v1alpha1.TransitGatewayRouteTable) (*v1alpha1.TransitGatewayRouteTable, error)
	Update(*v1alpha1.TransitGatewayRouteTable) (*v1alpha1.TransitGatewayRouteTable, error)
	UpdateStatus(*v1alpha1.TransitGatewayRouteTable) (*v1alpha1.TransitGatewayRouteTable, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.TransitGatewayRouteTable, error)
	List(opts v1.ListOptions) (*v1alpha1.TransitGatewayRouteTableList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TransitGatewayRouteTable, err error)
	TransitGatewayRouteTableExpansion
}

// transitGatewayRouteTables implements TransitGatewayRouteTableInterface
type transitGatewayRouteTables struct {
	client rest.Interface
	ns     string
}

// newTransitGatewayRouteTables returns a TransitGatewayRouteTables
func newTransitGatewayRouteTables(c *Ec2V1alpha1Client, namespace string) *transitGatewayRouteTables {
	return &transitGatewayRouteTables{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the transitGatewayRouteTable, and returns the corresponding transitGatewayRouteTable object, and an error if there is any.
func (c *transitGatewayRouteTables) Get(name string, options v1.GetOptions) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TransitGatewayRouteTables that match those selectors.
func (c *transitGatewayRouteTables) List(opts v1.ListOptions) (result *v1alpha1.TransitGatewayRouteTableList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.TransitGatewayRouteTableList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested transitGatewayRouteTables.
func (c *transitGatewayRouteTables) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a transitGatewayRouteTable and creates it.  Returns the server's representation of the transitGatewayRouteTable, and an error, if there is any.
func (c *transitGatewayRouteTables) Create(transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Body(transitGatewayRouteTable).
		Do().
		Into(result)
	return
}

// Update takes the representation of a transitGatewayRouteTable and updates it. Returns the server's representation of the transitGatewayRouteTable, and an error, if there is any.
func (c *transitGatewayRouteTables) Update(transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(transitGatewayRouteTable.Name).
		Body(transitGatewayRouteTable).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *transitGatewayRouteTables) UpdateStatus(transitGatewayRouteTable *v1alpha1.TransitGatewayRouteTable) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(transitGatewayRouteTable.Name).
		SubResource("status").
		Body(transitGatewayRouteTable).
		Do().
		Into(result)
	return
}

// Delete takes name of the transitGatewayRouteTable and deletes it. Returns an error if one occurs.
func (c *transitGatewayRouteTables) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *transitGatewayRouteTables) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched transitGatewayRouteTable.
func (c *transitGatewayRouteTables) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TransitGatewayRouteTable, err error) {
	result = &v1alpha1.TransitGatewayRouteTable{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("transitgatewayroutetables").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
