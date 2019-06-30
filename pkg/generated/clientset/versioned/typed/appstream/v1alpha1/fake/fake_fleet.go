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
	v1alpha1 "awsoperator.io/pkg/apis/appstream/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFleets implements FleetInterface
type FakeFleets struct {
	Fake *FakeAppstreamV1alpha1
	ns   string
}

var fleetsResource = schema.GroupVersionResource{Group: "appstream.awsoperator.io", Version: "v1alpha1", Resource: "fleets"}

var fleetsKind = schema.GroupVersionKind{Group: "appstream.awsoperator.io", Version: "v1alpha1", Kind: "Fleet"}

// Get takes name of the fleet, and returns the corresponding fleet object, and an error if there is any.
func (c *FakeFleets) Get(name string, options v1.GetOptions) (result *v1alpha1.Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(fleetsResource, c.ns, name), &v1alpha1.Fleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fleet), err
}

// List takes label and field selectors, and returns the list of Fleets that match those selectors.
func (c *FakeFleets) List(opts v1.ListOptions) (result *v1alpha1.FleetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(fleetsResource, fleetsKind, c.ns, opts), &v1alpha1.FleetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FleetList{ListMeta: obj.(*v1alpha1.FleetList).ListMeta}
	for _, item := range obj.(*v1alpha1.FleetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fleets.
func (c *FakeFleets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(fleetsResource, c.ns, opts))

}

// Create takes the representation of a fleet and creates it.  Returns the server's representation of the fleet, and an error, if there is any.
func (c *FakeFleets) Create(fleet *v1alpha1.Fleet) (result *v1alpha1.Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(fleetsResource, c.ns, fleet), &v1alpha1.Fleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fleet), err
}

// Update takes the representation of a fleet and updates it. Returns the server's representation of the fleet, and an error, if there is any.
func (c *FakeFleets) Update(fleet *v1alpha1.Fleet) (result *v1alpha1.Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(fleetsResource, c.ns, fleet), &v1alpha1.Fleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fleet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFleets) UpdateStatus(fleet *v1alpha1.Fleet) (*v1alpha1.Fleet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(fleetsResource, "status", c.ns, fleet), &v1alpha1.Fleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fleet), err
}

// Delete takes name of the fleet and deletes it. Returns an error if one occurs.
func (c *FakeFleets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(fleetsResource, c.ns, name), &v1alpha1.Fleet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFleets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(fleetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FleetList{})
	return err
}

// Patch applies the patch and returns the patched fleet.
func (c *FakeFleets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Fleet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(fleetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Fleet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Fleet), err
}
