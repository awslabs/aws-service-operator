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
	v1alpha1 "awsoperator.io/pkg/apis/directoryservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMicrosoftADs implements MicrosoftADInterface
type FakeMicrosoftADs struct {
	Fake *FakeDirectoryserviceV1alpha1
	ns   string
}

var microsoftadsResource = schema.GroupVersionResource{Group: "directoryservice.awsoperator.io", Version: "v1alpha1", Resource: "microsoftads"}

var microsoftadsKind = schema.GroupVersionKind{Group: "directoryservice.awsoperator.io", Version: "v1alpha1", Kind: "MicrosoftAD"}

// Get takes name of the microsoftAD, and returns the corresponding microsoftAD object, and an error if there is any.
func (c *FakeMicrosoftADs) Get(name string, options v1.GetOptions) (result *v1alpha1.MicrosoftAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(microsoftadsResource, c.ns, name), &v1alpha1.MicrosoftAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MicrosoftAD), err
}

// List takes label and field selectors, and returns the list of MicrosoftADs that match those selectors.
func (c *FakeMicrosoftADs) List(opts v1.ListOptions) (result *v1alpha1.MicrosoftADList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(microsoftadsResource, microsoftadsKind, c.ns, opts), &v1alpha1.MicrosoftADList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MicrosoftADList{ListMeta: obj.(*v1alpha1.MicrosoftADList).ListMeta}
	for _, item := range obj.(*v1alpha1.MicrosoftADList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested microsoftADs.
func (c *FakeMicrosoftADs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(microsoftadsResource, c.ns, opts))

}

// Create takes the representation of a microsoftAD and creates it.  Returns the server's representation of the microsoftAD, and an error, if there is any.
func (c *FakeMicrosoftADs) Create(microsoftAD *v1alpha1.MicrosoftAD) (result *v1alpha1.MicrosoftAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(microsoftadsResource, c.ns, microsoftAD), &v1alpha1.MicrosoftAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MicrosoftAD), err
}

// Update takes the representation of a microsoftAD and updates it. Returns the server's representation of the microsoftAD, and an error, if there is any.
func (c *FakeMicrosoftADs) Update(microsoftAD *v1alpha1.MicrosoftAD) (result *v1alpha1.MicrosoftAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(microsoftadsResource, c.ns, microsoftAD), &v1alpha1.MicrosoftAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MicrosoftAD), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMicrosoftADs) UpdateStatus(microsoftAD *v1alpha1.MicrosoftAD) (*v1alpha1.MicrosoftAD, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(microsoftadsResource, "status", c.ns, microsoftAD), &v1alpha1.MicrosoftAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MicrosoftAD), err
}

// Delete takes name of the microsoftAD and deletes it. Returns an error if one occurs.
func (c *FakeMicrosoftADs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(microsoftadsResource, c.ns, name), &v1alpha1.MicrosoftAD{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMicrosoftADs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(microsoftadsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.MicrosoftADList{})
	return err
}

// Patch applies the patch and returns the patched microsoftAD.
func (c *FakeMicrosoftADs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.MicrosoftAD, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(microsoftadsResource, c.ns, name, pt, data, subresources...), &v1alpha1.MicrosoftAD{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MicrosoftAD), err
}
