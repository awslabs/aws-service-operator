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
	v1alpha1 "awsoperator.io/pkg/apis/dms/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeReplicationSubnetGroups implements ReplicationSubnetGroupInterface
type FakeReplicationSubnetGroups struct {
	Fake *FakeDmsV1alpha1
	ns   string
}

var replicationsubnetgroupsResource = schema.GroupVersionResource{Group: "dms.awsoperator.io", Version: "v1alpha1", Resource: "replicationsubnetgroups"}

var replicationsubnetgroupsKind = schema.GroupVersionKind{Group: "dms.awsoperator.io", Version: "v1alpha1", Kind: "ReplicationSubnetGroup"}

// Get takes name of the replicationSubnetGroup, and returns the corresponding replicationSubnetGroup object, and an error if there is any.
func (c *FakeReplicationSubnetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.ReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(replicationsubnetgroupsResource, c.ns, name), &v1alpha1.ReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationSubnetGroup), err
}

// List takes label and field selectors, and returns the list of ReplicationSubnetGroups that match those selectors.
func (c *FakeReplicationSubnetGroups) List(opts v1.ListOptions) (result *v1alpha1.ReplicationSubnetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(replicationsubnetgroupsResource, replicationsubnetgroupsKind, c.ns, opts), &v1alpha1.ReplicationSubnetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ReplicationSubnetGroupList{ListMeta: obj.(*v1alpha1.ReplicationSubnetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.ReplicationSubnetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicationSubnetGroups.
func (c *FakeReplicationSubnetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(replicationsubnetgroupsResource, c.ns, opts))

}

// Create takes the representation of a replicationSubnetGroup and creates it.  Returns the server's representation of the replicationSubnetGroup, and an error, if there is any.
func (c *FakeReplicationSubnetGroups) Create(replicationSubnetGroup *v1alpha1.ReplicationSubnetGroup) (result *v1alpha1.ReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(replicationsubnetgroupsResource, c.ns, replicationSubnetGroup), &v1alpha1.ReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationSubnetGroup), err
}

// Update takes the representation of a replicationSubnetGroup and updates it. Returns the server's representation of the replicationSubnetGroup, and an error, if there is any.
func (c *FakeReplicationSubnetGroups) Update(replicationSubnetGroup *v1alpha1.ReplicationSubnetGroup) (result *v1alpha1.ReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(replicationsubnetgroupsResource, c.ns, replicationSubnetGroup), &v1alpha1.ReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationSubnetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeReplicationSubnetGroups) UpdateStatus(replicationSubnetGroup *v1alpha1.ReplicationSubnetGroup) (*v1alpha1.ReplicationSubnetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(replicationsubnetgroupsResource, "status", c.ns, replicationSubnetGroup), &v1alpha1.ReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationSubnetGroup), err
}

// Delete takes name of the replicationSubnetGroup and deletes it. Returns an error if one occurs.
func (c *FakeReplicationSubnetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(replicationsubnetgroupsResource, c.ns, name), &v1alpha1.ReplicationSubnetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeReplicationSubnetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(replicationsubnetgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ReplicationSubnetGroupList{})
	return err
}

// Patch applies the patch and returns the patched replicationSubnetGroup.
func (c *FakeReplicationSubnetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ReplicationSubnetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(replicationsubnetgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ReplicationSubnetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ReplicationSubnetGroup), err
}
