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

// FakeWebACLAssociations implements WebACLAssociationInterface
type FakeWebACLAssociations struct {
	Fake *FakeWafregionalV1alpha1
	ns   string
}

var webaclassociationsResource = schema.GroupVersionResource{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Resource: "webaclassociations"}

var webaclassociationsKind = schema.GroupVersionKind{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Kind: "WebACLAssociation"}

// Get takes name of the webACLAssociation, and returns the corresponding webACLAssociation object, and an error if there is any.
func (c *FakeWebACLAssociations) Get(name string, options v1.GetOptions) (result *v1alpha1.WebACLAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(webaclassociationsResource, c.ns, name), &v1alpha1.WebACLAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebACLAssociation), err
}

// List takes label and field selectors, and returns the list of WebACLAssociations that match those selectors.
func (c *FakeWebACLAssociations) List(opts v1.ListOptions) (result *v1alpha1.WebACLAssociationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(webaclassociationsResource, webaclassociationsKind, c.ns, opts), &v1alpha1.WebACLAssociationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.WebACLAssociationList{ListMeta: obj.(*v1alpha1.WebACLAssociationList).ListMeta}
	for _, item := range obj.(*v1alpha1.WebACLAssociationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested webACLAssociations.
func (c *FakeWebACLAssociations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(webaclassociationsResource, c.ns, opts))

}

// Create takes the representation of a webACLAssociation and creates it.  Returns the server's representation of the webACLAssociation, and an error, if there is any.
func (c *FakeWebACLAssociations) Create(webACLAssociation *v1alpha1.WebACLAssociation) (result *v1alpha1.WebACLAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(webaclassociationsResource, c.ns, webACLAssociation), &v1alpha1.WebACLAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebACLAssociation), err
}

// Update takes the representation of a webACLAssociation and updates it. Returns the server's representation of the webACLAssociation, and an error, if there is any.
func (c *FakeWebACLAssociations) Update(webACLAssociation *v1alpha1.WebACLAssociation) (result *v1alpha1.WebACLAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(webaclassociationsResource, c.ns, webACLAssociation), &v1alpha1.WebACLAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebACLAssociation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeWebACLAssociations) UpdateStatus(webACLAssociation *v1alpha1.WebACLAssociation) (*v1alpha1.WebACLAssociation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(webaclassociationsResource, "status", c.ns, webACLAssociation), &v1alpha1.WebACLAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebACLAssociation), err
}

// Delete takes name of the webACLAssociation and deletes it. Returns an error if one occurs.
func (c *FakeWebACLAssociations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(webaclassociationsResource, c.ns, name), &v1alpha1.WebACLAssociation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeWebACLAssociations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(webaclassociationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.WebACLAssociationList{})
	return err
}

// Patch applies the patch and returns the patched webACLAssociation.
func (c *FakeWebACLAssociations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.WebACLAssociation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(webaclassociationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.WebACLAssociation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.WebACLAssociation), err
}
