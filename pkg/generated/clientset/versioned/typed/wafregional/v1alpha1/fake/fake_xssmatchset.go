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
	v1alpha1 "awsoperator.io/pkg/apis/wafregional/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeXssMatchSets implements XssMatchSetInterface
type FakeXssMatchSets struct {
	Fake *FakeWafregionalV1alpha1
	ns   string
}

var xssmatchsetsResource = schema.GroupVersionResource{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Resource: "xssmatchsets"}

var xssmatchsetsKind = schema.GroupVersionKind{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Kind: "XssMatchSet"}

// Get takes name of the xssMatchSet, and returns the corresponding xssMatchSet object, and an error if there is any.
func (c *FakeXssMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.XssMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(xssmatchsetsResource, c.ns, name), &v1alpha1.XssMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XssMatchSet), err
}

// List takes label and field selectors, and returns the list of XssMatchSets that match those selectors.
func (c *FakeXssMatchSets) List(opts v1.ListOptions) (result *v1alpha1.XssMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(xssmatchsetsResource, xssmatchsetsKind, c.ns, opts), &v1alpha1.XssMatchSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.XssMatchSetList{ListMeta: obj.(*v1alpha1.XssMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.XssMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested xssMatchSets.
func (c *FakeXssMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(xssmatchsetsResource, c.ns, opts))

}

// Create takes the representation of a xssMatchSet and creates it.  Returns the server's representation of the xssMatchSet, and an error, if there is any.
func (c *FakeXssMatchSets) Create(xssMatchSet *v1alpha1.XssMatchSet) (result *v1alpha1.XssMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(xssmatchsetsResource, c.ns, xssMatchSet), &v1alpha1.XssMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XssMatchSet), err
}

// Update takes the representation of a xssMatchSet and updates it. Returns the server's representation of the xssMatchSet, and an error, if there is any.
func (c *FakeXssMatchSets) Update(xssMatchSet *v1alpha1.XssMatchSet) (result *v1alpha1.XssMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(xssmatchsetsResource, c.ns, xssMatchSet), &v1alpha1.XssMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XssMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeXssMatchSets) UpdateStatus(xssMatchSet *v1alpha1.XssMatchSet) (*v1alpha1.XssMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(xssmatchsetsResource, "status", c.ns, xssMatchSet), &v1alpha1.XssMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XssMatchSet), err
}

// Delete takes name of the xssMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeXssMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(xssmatchsetsResource, c.ns, name), &v1alpha1.XssMatchSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeXssMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(xssmatchsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.XssMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched xssMatchSet.
func (c *FakeXssMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.XssMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(xssmatchsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.XssMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.XssMatchSet), err
}
