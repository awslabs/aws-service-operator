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
	v1alpha1 "awsoperator.io/pkg/apis/neptune/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDBParameterGroups implements DBParameterGroupInterface
type FakeDBParameterGroups struct {
	Fake *FakeNeptuneV1alpha1
	ns   string
}

var dbparametergroupsResource = schema.GroupVersionResource{Group: "neptune.awsoperator.io", Version: "v1alpha1", Resource: "dbparametergroups"}

var dbparametergroupsKind = schema.GroupVersionKind{Group: "neptune.awsoperator.io", Version: "v1alpha1", Kind: "DBParameterGroup"}

// Get takes name of the dBParameterGroup, and returns the corresponding dBParameterGroup object, and an error if there is any.
func (c *FakeDBParameterGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.DBParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dbparametergroupsResource, c.ns, name), &v1alpha1.DBParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DBParameterGroup), err
}

// List takes label and field selectors, and returns the list of DBParameterGroups that match those selectors.
func (c *FakeDBParameterGroups) List(opts v1.ListOptions) (result *v1alpha1.DBParameterGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dbparametergroupsResource, dbparametergroupsKind, c.ns, opts), &v1alpha1.DBParameterGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DBParameterGroupList{ListMeta: obj.(*v1alpha1.DBParameterGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.DBParameterGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dBParameterGroups.
func (c *FakeDBParameterGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dbparametergroupsResource, c.ns, opts))

}

// Create takes the representation of a dBParameterGroup and creates it.  Returns the server's representation of the dBParameterGroup, and an error, if there is any.
func (c *FakeDBParameterGroups) Create(dBParameterGroup *v1alpha1.DBParameterGroup) (result *v1alpha1.DBParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dbparametergroupsResource, c.ns, dBParameterGroup), &v1alpha1.DBParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DBParameterGroup), err
}

// Update takes the representation of a dBParameterGroup and updates it. Returns the server's representation of the dBParameterGroup, and an error, if there is any.
func (c *FakeDBParameterGroups) Update(dBParameterGroup *v1alpha1.DBParameterGroup) (result *v1alpha1.DBParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dbparametergroupsResource, c.ns, dBParameterGroup), &v1alpha1.DBParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DBParameterGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDBParameterGroups) UpdateStatus(dBParameterGroup *v1alpha1.DBParameterGroup) (*v1alpha1.DBParameterGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dbparametergroupsResource, "status", c.ns, dBParameterGroup), &v1alpha1.DBParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DBParameterGroup), err
}

// Delete takes name of the dBParameterGroup and deletes it. Returns an error if one occurs.
func (c *FakeDBParameterGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(dbparametergroupsResource, c.ns, name), &v1alpha1.DBParameterGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDBParameterGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dbparametergroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DBParameterGroupList{})
	return err
}

// Patch applies the patch and returns the patched dBParameterGroup.
func (c *FakeDBParameterGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.DBParameterGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dbparametergroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DBParameterGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DBParameterGroup), err
}
