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
	v1alpha1 "awsoperator.io/pkg/apis/cloudtrail/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTrails implements TrailInterface
type FakeTrails struct {
	Fake *FakeCloudtrailV1alpha1
	ns   string
}

var trailsResource = schema.GroupVersionResource{Group: "cloudtrail.awsoperator.io", Version: "v1alpha1", Resource: "trails"}

var trailsKind = schema.GroupVersionKind{Group: "cloudtrail.awsoperator.io", Version: "v1alpha1", Kind: "Trail"}

// Get takes name of the trail, and returns the corresponding trail object, and an error if there is any.
func (c *FakeTrails) Get(name string, options v1.GetOptions) (result *v1alpha1.Trail, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(trailsResource, c.ns, name), &v1alpha1.Trail{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Trail), err
}

// List takes label and field selectors, and returns the list of Trails that match those selectors.
func (c *FakeTrails) List(opts v1.ListOptions) (result *v1alpha1.TrailList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(trailsResource, trailsKind, c.ns, opts), &v1alpha1.TrailList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TrailList{ListMeta: obj.(*v1alpha1.TrailList).ListMeta}
	for _, item := range obj.(*v1alpha1.TrailList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested trails.
func (c *FakeTrails) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(trailsResource, c.ns, opts))

}

// Create takes the representation of a trail and creates it.  Returns the server's representation of the trail, and an error, if there is any.
func (c *FakeTrails) Create(trail *v1alpha1.Trail) (result *v1alpha1.Trail, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(trailsResource, c.ns, trail), &v1alpha1.Trail{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Trail), err
}

// Update takes the representation of a trail and updates it. Returns the server's representation of the trail, and an error, if there is any.
func (c *FakeTrails) Update(trail *v1alpha1.Trail) (result *v1alpha1.Trail, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(trailsResource, c.ns, trail), &v1alpha1.Trail{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Trail), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTrails) UpdateStatus(trail *v1alpha1.Trail) (*v1alpha1.Trail, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(trailsResource, "status", c.ns, trail), &v1alpha1.Trail{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Trail), err
}

// Delete takes name of the trail and deletes it. Returns an error if one occurs.
func (c *FakeTrails) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(trailsResource, c.ns, name), &v1alpha1.Trail{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTrails) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(trailsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TrailList{})
	return err
}

// Patch applies the patch and returns the patched trail.
func (c *FakeTrails) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Trail, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(trailsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Trail{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Trail), err
}
