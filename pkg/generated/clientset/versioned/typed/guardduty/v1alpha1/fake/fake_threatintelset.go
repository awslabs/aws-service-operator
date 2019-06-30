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
	v1alpha1 "awsoperator.io/pkg/apis/guardduty/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeThreatIntelSets implements ThreatIntelSetInterface
type FakeThreatIntelSets struct {
	Fake *FakeGuarddutyV1alpha1
	ns   string
}

var threatintelsetsResource = schema.GroupVersionResource{Group: "guardduty.awsoperator.io", Version: "v1alpha1", Resource: "threatintelsets"}

var threatintelsetsKind = schema.GroupVersionKind{Group: "guardduty.awsoperator.io", Version: "v1alpha1", Kind: "ThreatIntelSet"}

// Get takes name of the threatIntelSet, and returns the corresponding threatIntelSet object, and an error if there is any.
func (c *FakeThreatIntelSets) Get(name string, options v1.GetOptions) (result *v1alpha1.ThreatIntelSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(threatintelsetsResource, c.ns, name), &v1alpha1.ThreatIntelSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThreatIntelSet), err
}

// List takes label and field selectors, and returns the list of ThreatIntelSets that match those selectors.
func (c *FakeThreatIntelSets) List(opts v1.ListOptions) (result *v1alpha1.ThreatIntelSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(threatintelsetsResource, threatintelsetsKind, c.ns, opts), &v1alpha1.ThreatIntelSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ThreatIntelSetList{ListMeta: obj.(*v1alpha1.ThreatIntelSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.ThreatIntelSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested threatIntelSets.
func (c *FakeThreatIntelSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(threatintelsetsResource, c.ns, opts))

}

// Create takes the representation of a threatIntelSet and creates it.  Returns the server's representation of the threatIntelSet, and an error, if there is any.
func (c *FakeThreatIntelSets) Create(threatIntelSet *v1alpha1.ThreatIntelSet) (result *v1alpha1.ThreatIntelSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(threatintelsetsResource, c.ns, threatIntelSet), &v1alpha1.ThreatIntelSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThreatIntelSet), err
}

// Update takes the representation of a threatIntelSet and updates it. Returns the server's representation of the threatIntelSet, and an error, if there is any.
func (c *FakeThreatIntelSets) Update(threatIntelSet *v1alpha1.ThreatIntelSet) (result *v1alpha1.ThreatIntelSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(threatintelsetsResource, c.ns, threatIntelSet), &v1alpha1.ThreatIntelSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThreatIntelSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeThreatIntelSets) UpdateStatus(threatIntelSet *v1alpha1.ThreatIntelSet) (*v1alpha1.ThreatIntelSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(threatintelsetsResource, "status", c.ns, threatIntelSet), &v1alpha1.ThreatIntelSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThreatIntelSet), err
}

// Delete takes name of the threatIntelSet and deletes it. Returns an error if one occurs.
func (c *FakeThreatIntelSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(threatintelsetsResource, c.ns, name), &v1alpha1.ThreatIntelSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeThreatIntelSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(threatintelsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ThreatIntelSetList{})
	return err
}

// Patch applies the patch and returns the patched threatIntelSet.
func (c *FakeThreatIntelSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ThreatIntelSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(threatintelsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ThreatIntelSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ThreatIntelSet), err
}
