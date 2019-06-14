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

// FakeSubnetNetworkAclAssociations implements SubnetNetworkAclAssociationInterface
type FakeSubnetNetworkAclAssociations struct {
	Fake *FakeEc2V1alpha1
	ns   string
}

var subnetnetworkaclassociationsResource = schema.GroupVersionResource{Group: "ec2.awsoperator.io", Version: "v1alpha1", Resource: "subnetnetworkaclassociations"}

var subnetnetworkaclassociationsKind = schema.GroupVersionKind{Group: "ec2.awsoperator.io", Version: "v1alpha1", Kind: "SubnetNetworkAclAssociation"}

// Get takes name of the subnetNetworkAclAssociation, and returns the corresponding subnetNetworkAclAssociation object, and an error if there is any.
func (c *FakeSubnetNetworkAclAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.SubnetNetworkAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subnetnetworkaclassociationsResource, c.ns, name), &v1alpha1.SubnetNetworkAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkAclAssociation), err
}

// List takes label and field selectors, and returns the list of SubnetNetworkAclAssociations that match those selectors.
func (c *FakeSubnetNetworkAclAssociations) List(opts v1.ListOptions) (result *v1alpha1.SubnetNetworkAclAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subnetnetworkaclassociationsResource, subnetnetworkaclassociationsKind, c.ns, opts), &v1alpha1.SubnetNetworkAclAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubnetNetworkAclAssociationList{ListMeta: obj.(*v1alpha1.SubnetNetworkAclAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubnetNetworkAclAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subnetNetworkAclAssociations.
func (c *FakeSubnetNetworkAclAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subnetnetworkaclassociationsResource, c.ns, opts))

}

// Create takes the representation of a subnetNetworkAclAssociation and creates it.  Returns the server's representation of the subnetNetworkAclAssociation, and an error, if there is any.
func (c *FakeSubnetNetworkAclAssociations) Create(subnetNetworkAclAssociation *v1alpha1.SubnetNetworkAclAssociation) (result *v1alpha1.SubnetNetworkAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subnetnetworkaclassociationsResource, c.ns, subnetNetworkAclAssociation), &v1alpha1.SubnetNetworkAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkAclAssociation), err
}

// Update takes the representation of a subnetNetworkAclAssociation and updates it. Returns the server's representation of the subnetNetworkAclAssociation, and an error, if there is any.
func (c *FakeSubnetNetworkAclAssociations) Update(subnetNetworkAclAssociation *v1alpha1.SubnetNetworkAclAssociation) (result *v1alpha1.SubnetNetworkAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subnetnetworkaclassociationsResource, c.ns, subnetNetworkAclAssociation), &v1alpha1.SubnetNetworkAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkAclAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubnetNetworkAclAssociations) UpdateStatus(subnetNetworkAclAssociation *v1alpha1.SubnetNetworkAclAssociation) (*v1alpha1.SubnetNetworkAclAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subnetnetworkaclassociationsResource, "status", c.ns, subnetNetworkAclAssociation), &v1alpha1.SubnetNetworkAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkAclAssociation), err
}

// Delete takes name of the subnetNetworkAclAssociation and deletes it. Returns an error if one occurs.
func (c *FakeSubnetNetworkAclAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(subnetnetworkaclassociationsResource, c.ns, name), &v1alpha1.SubnetNetworkAclAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnetNetworkAclAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subnetnetworkaclassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubnetNetworkAclAssociationList{})
	return err
}

// Patch applies the patch and returns the patched subnetNetworkAclAssociation.
func (c *FakeSubnetNetworkAclAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SubnetNetworkAclAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subnetnetworkaclassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubnetNetworkAclAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetNetworkAclAssociation), err
}
