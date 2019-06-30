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
	v1alpha1 "awsoperator.io/pkg/apis/dax/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeParameterGroups implements ParameterGroupInterface
type FakeParameterGroups struct {
	Fake *FakeDaxV1alpha1
	ns   string
}

var parametergroupsResource = schema.GroupVersionResource{Group: "dax.awsoperator.io", Version: "v1alpha1", Resource: "parametergroups"}

var parametergroupsKind = schema.GroupVersionKind{Group: "dax.awsoperator.io", Version: "v1alpha1", Kind: "ParameterGroup"}

// Get takes name of the parameterGroup, and returns the corresponding parameterGroup object, and an error if there is any.
func (c *FakeParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(parametergroupsResource, c.ns, name), &v1alpha1.ParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ParameterGroup), err
}

// List takes label and field selectors, and returns the list of ParameterGroups that match those selectors.
func (c *FakeParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.ParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(parametergroupsResource, parametergroupsKind, c.ns, opts), &v1alpha1.ParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ParameterGroupList{ListMeta: obj.(*v1alpha1.ParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested parameterGroups.
func (c *FakeParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(parametergroupsResource, c.ns, opts))

}

// Create takes the representation of a parameterGroup and creates it.  Returns the server's representation of the parameterGroup, and an error, if there is any.
func (c *FakeParameterGroups) Create(parameterGroup *v1alpha1.ParameterGroup) (result *v1alpha1.ParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(parametergroupsResource, c.ns, parameterGroup), &v1alpha1.ParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ParameterGroup), err
}

// Update takes the representation of a parameterGroup and updates it. Returns the server's representation of the parameterGroup, and an error, if there is any.
func (c *FakeParameterGroups) Update(parameterGroup *v1alpha1.ParameterGroup) (result *v1alpha1.ParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(parametergroupsResource, c.ns, parameterGroup), &v1alpha1.ParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeParameterGroups) UpdateStatus(parameterGroup *v1alpha1.ParameterGroup) (*v1alpha1.ParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(parametergroupsResource, "status", c.ns, parameterGroup), &v1alpha1.ParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ParameterGroup), err
}

// Delete takes name of the parameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(parametergroupsResource, c.ns, name), &v1alpha1.ParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(parametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched parameterGroup.
func (c *FakeParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(parametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ParameterGroup), err
}
