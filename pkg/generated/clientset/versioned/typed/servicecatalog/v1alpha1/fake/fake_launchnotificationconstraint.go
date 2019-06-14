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
	v1alpha1 "awsoperator.io/pkg/apis/servicecatalog/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLaunchNotificationConstraints implements LaunchNotificationConstraintInterface
type FakeLaunchNotificationConstraints struct {
	Fake *FakeServicecatalogV1alpha1
	ns   string
}

var launchnotificationconstraintsResource = schema.GroupVersionResource{Group: "servicecatalog.awsoperator.io", Version: "v1alpha1", Resource: "launchnotificationconstraints"}

var launchnotificationconstraintsKind = schema.GroupVersionKind{Group: "servicecatalog.awsoperator.io", Version: "v1alpha1", Kind: "LaunchNotificationConstraint"}

// Get takes name of the launchNotificationConstraint, and returns the corresponding launchNotificationConstraint object, and an error if there is any.
func (c *FakeLaunchNotificationConstraints) Get(name string, options v1.GetOptions) (result *v1alpha1.LaunchNotificationConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(launchnotificationconstraintsResource, c.ns, name), &v1alpha1.LaunchNotificationConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchNotificationConstraint), err
}

// List takes label and field selectors, and returns the list of LaunchNotificationConstraints that match those selectors.
func (c *FakeLaunchNotificationConstraints) List(opts v1.ListOptions) (result *v1alpha1.LaunchNotificationConstraintList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(launchnotificationconstraintsResource, launchnotificationconstraintsKind, c.ns, opts), &v1alpha1.LaunchNotificationConstraintList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LaunchNotificationConstraintList{ListMeta: obj.(*v1alpha1.LaunchNotificationConstraintList).ListMeta}
	for _, item := range obj.(*v1alpha1.LaunchNotificationConstraintList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested launchNotificationConstraints.
func (c *FakeLaunchNotificationConstraints) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(launchnotificationconstraintsResource, c.ns, opts))

}

// Create takes the representation of a launchNotificationConstraint and creates it.  Returns the server's representation of the launchNotificationConstraint, and an error, if there is any.
func (c *FakeLaunchNotificationConstraints) Create(launchNotificationConstraint *v1alpha1.LaunchNotificationConstraint) (result *v1alpha1.LaunchNotificationConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(launchnotificationconstraintsResource, c.ns, launchNotificationConstraint), &v1alpha1.LaunchNotificationConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchNotificationConstraint), err
}

// Update takes the representation of a launchNotificationConstraint and updates it. Returns the server's representation of the launchNotificationConstraint, and an error, if there is any.
func (c *FakeLaunchNotificationConstraints) Update(launchNotificationConstraint *v1alpha1.LaunchNotificationConstraint) (result *v1alpha1.LaunchNotificationConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(launchnotificationconstraintsResource, c.ns, launchNotificationConstraint), &v1alpha1.LaunchNotificationConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchNotificationConstraint), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLaunchNotificationConstraints) UpdateStatus(launchNotificationConstraint *v1alpha1.LaunchNotificationConstraint) (*v1alpha1.LaunchNotificationConstraint, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(launchnotificationconstraintsResource, "status", c.ns, launchNotificationConstraint), &v1alpha1.LaunchNotificationConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchNotificationConstraint), err
}

// Delete takes name of the launchNotificationConstraint and deletes it. Returns an error if one occurs.
func (c *FakeLaunchNotificationConstraints) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(launchnotificationconstraintsResource, c.ns, name), &v1alpha1.LaunchNotificationConstraint{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLaunchNotificationConstraints) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(launchnotificationconstraintsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LaunchNotificationConstraintList{})
	return err
}

// Patch applies the patch and returns the patched launchNotificationConstraint.
func (c *FakeLaunchNotificationConstraints) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LaunchNotificationConstraint, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(launchnotificationconstraintsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LaunchNotificationConstraint{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LaunchNotificationConstraint), err
}
