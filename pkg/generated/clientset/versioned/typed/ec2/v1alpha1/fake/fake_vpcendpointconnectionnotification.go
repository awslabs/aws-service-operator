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

package fake

import (
	v1alpha1 "awsoperator.io/pkg/apis/ec2/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeVPCEndpointConnectionNotifications implements VPCEndpointConnectionNotificationInterface
type FakeVPCEndpointConnectionNotifications struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var vpcendpointconnectionnotificationsResource = schema.GroupVersionResource{Group: "ec2.awsoperator.io", Version: "v1alpha1", Resource: "vpcendpointconnectionnotifications"}

var vpcendpointconnectionnotificationsKind = schema.GroupVersionKind{Group: "ec2.awsoperator.io", Version: "v1alpha1", Kind: "VPCEndpointConnectionNotification"}

// Get takes name of the vPCEndpointConnectionNotification, and returns the corresponding vPCEndpointConnectionNotification object, and an error if there is any.
func (c *FakeVPCEndpointConnectionNotifications) Get(name string, options v1.GetOptions) (result *v1alpha1.VPCEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(vpcendpointconnectionnotificationsResource, c.ns, name), &v1alpha1.VPCEndpointConnectionNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VPCEndpointConnectionNotification), err
}

// List takes label and field selectors, and returns the list of VPCEndpointConnectionNotifications that match those selectors.
func (c *FakeVPCEndpointConnectionNotifications) List(opts v1.ListOptions) (result *v1alpha1.VPCEndpointConnectionNotificationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(vpcendpointconnectionnotificationsResource, vpcendpointconnectionnotificationsKind, c.ns, opts), &v1alpha1.VPCEndpointConnectionNotificationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.VPCEndpointConnectionNotificationList{ListMeta: obj.(*v1alpha1.VPCEndpointConnectionNotificationList).ListMeta}
	for _, item := range obj.(*v1alpha1.VPCEndpointConnectionNotificationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested vPCEndpointConnectionNotifications.
func (c *FakeVPCEndpointConnectionNotifications) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(vpcendpointconnectionnotificationsResource, c.ns, opts))

}

// Create takes the representation of a vPCEndpointConnectionNotification and creates it.  Returns the server's representation of the vPCEndpointConnectionNotification, and an error, if there is any.
func (c *FakeVPCEndpointConnectionNotifications) Create(vPCEndpointConnectionNotification *v1alpha1.VPCEndpointConnectionNotification) (result *v1alpha1.VPCEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(vpcendpointconnectionnotificationsResource, c.ns, vPCEndpointConnectionNotification), &v1alpha1.VPCEndpointConnectionNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VPCEndpointConnectionNotification), err
}

// Update takes the representation of a vPCEndpointConnectionNotification and updates it. Returns the server's representation of the vPCEndpointConnectionNotification, and an error, if there is any.
func (c *FakeVPCEndpointConnectionNotifications) Update(vPCEndpointConnectionNotification *v1alpha1.VPCEndpointConnectionNotification) (result *v1alpha1.VPCEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(vpcendpointconnectionnotificationsResource, c.ns, vPCEndpointConnectionNotification), &v1alpha1.VPCEndpointConnectionNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VPCEndpointConnectionNotification), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeVPCEndpointConnectionNotifications) UpdateStatus(vPCEndpointConnectionNotification *v1alpha1.VPCEndpointConnectionNotification) (*v1alpha1.VPCEndpointConnectionNotification, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(vpcendpointconnectionnotificationsResource, "status", c.ns, vPCEndpointConnectionNotification), &v1alpha1.VPCEndpointConnectionNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VPCEndpointConnectionNotification), err
}

// Delete takes name of the vPCEndpointConnectionNotification and deletes it. Returns an error if one occurs.
func (c *FakeVPCEndpointConnectionNotifications) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(vpcendpointconnectionnotificationsResource, c.ns, name), &v1alpha1.VPCEndpointConnectionNotification{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeVPCEndpointConnectionNotifications) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(vpcendpointconnectionnotificationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.VPCEndpointConnectionNotificationList{})
	return err
}

// Patch applies the patch and returns the patched vPCEndpointConnectionNotification.
func (c *FakeVPCEndpointConnectionNotifications) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.VPCEndpointConnectionNotification, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(vpcendpointconnectionnotificationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.VPCEndpointConnectionNotification{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.VPCEndpointConnectionNotification), err
}
