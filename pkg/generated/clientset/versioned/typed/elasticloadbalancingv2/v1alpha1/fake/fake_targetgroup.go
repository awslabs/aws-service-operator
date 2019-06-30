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
	v1alpha1 "awsoperator.io/pkg/apis/elasticloadbalancingv2/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTargetGroups implements TargetGroupInterface
type FakeTargetGroups struct {
	Fake *FakeElasticloadbalancingv2V1alpha1
	ns   string
}

var targetgroupsResource = schema.GroupVersionResource{Group: "elasticloadbalancingv2.awsoperator.io", Version: "v1alpha1", Resource: "targetgroups"}

var targetgroupsKind = schema.GroupVersionKind{Group: "elasticloadbalancingv2.awsoperator.io", Version: "v1alpha1", Kind: "TargetGroup"}

// Get takes name of the targetGroup, and returns the corresponding targetGroup object, and an error if there is any.
func (c *FakeTargetGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.TargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(targetgroupsResource, c.ns, name), &v1alpha1.TargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetGroup), err
}

// List takes label and field selectors, and returns the list of TargetGroups that match those selectors.
func (c *FakeTargetGroups) List(opts v1.ListOptions) (result *v1alpha1.TargetGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(targetgroupsResource, targetgroupsKind, c.ns, opts), &v1alpha1.TargetGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.TargetGroupList{ListMeta: obj.(*v1alpha1.TargetGroupList).ListMeta}
	for _, item := range obj.(*v1alpha1.TargetGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested targetGroups.
func (c *FakeTargetGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(targetgroupsResource, c.ns, opts))

}

// Create takes the representation of a targetGroup and creates it.  Returns the server's representation of the targetGroup, and an error, if there is any.
func (c *FakeTargetGroups) Create(targetGroup *v1alpha1.TargetGroup) (result *v1alpha1.TargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(targetgroupsResource, c.ns, targetGroup), &v1alpha1.TargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetGroup), err
}

// Update takes the representation of a targetGroup and updates it. Returns the server's representation of the targetGroup, and an error, if there is any.
func (c *FakeTargetGroups) Update(targetGroup *v1alpha1.TargetGroup) (result *v1alpha1.TargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(targetgroupsResource, c.ns, targetGroup), &v1alpha1.TargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTargetGroups) UpdateStatus(targetGroup *v1alpha1.TargetGroup) (*v1alpha1.TargetGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(targetgroupsResource, "status", c.ns, targetGroup), &v1alpha1.TargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetGroup), err
}

// Delete takes name of the targetGroup and deletes it. Returns an error if one occurs.
func (c *FakeTargetGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(targetgroupsResource, c.ns, name), &v1alpha1.TargetGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTargetGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(targetgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.TargetGroupList{})
	return err
}

// Patch applies the patch and returns the patched targetGroup.
func (c *FakeTargetGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.TargetGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(targetgroupsResource, c.ns, name, pt, data, subresources...), &v1alpha1.TargetGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.TargetGroup), err
}
