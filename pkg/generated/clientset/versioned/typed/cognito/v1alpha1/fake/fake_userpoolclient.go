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
	v1alpha1 "awsoperator.io/pkg/apis/cognito/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeUserPoolClients implements UserPoolClientInterface
type FakeUserPoolClients struct {
	Fake *FakeCognitoV1alpha1
	ns   string
}

var userpoolclientsResource = schema.GroupVersionResource{Group: "cognito.awsoperator.io", Version: "v1alpha1", Resource: "userpoolclients"}

var userpoolclientsKind = schema.GroupVersionKind{Group: "cognito.awsoperator.io", Version: "v1alpha1", Kind: "UserPoolClient"}

// Get takes name of the userPoolClient, and returns the corresponding userPoolClient object, and an error if there is any.
func (c *FakeUserPoolClients) Get(name string, options v1.GetOptions) (result *v1alpha1.UserPoolClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(userpoolclientsResource, c.ns, name), &v1alpha1.UserPoolClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolClient), err
}

// List takes label and field selectors, and returns the list of UserPoolClients that match those selectors.
func (c *FakeUserPoolClients) List(opts v1.ListOptions) (result *v1alpha1.UserPoolClientList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(userpoolclientsResource, userpoolclientsKind, c.ns, opts), &v1alpha1.UserPoolClientList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.UserPoolClientList{ListMeta: obj.(*v1alpha1.UserPoolClientList).ListMeta}
	for _, item := range obj.(*v1alpha1.UserPoolClientList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userPoolClients.
func (c *FakeUserPoolClients) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(userpoolclientsResource, c.ns, opts))

}

// Create takes the representation of a userPoolClient and creates it.  Returns the server's representation of the userPoolClient, and an error, if there is any.
func (c *FakeUserPoolClients) Create(userPoolClient *v1alpha1.UserPoolClient) (result *v1alpha1.UserPoolClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(userpoolclientsResource, c.ns, userPoolClient), &v1alpha1.UserPoolClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolClient), err
}

// Update takes the representation of a userPoolClient and updates it. Returns the server's representation of the userPoolClient, and an error, if there is any.
func (c *FakeUserPoolClients) Update(userPoolClient *v1alpha1.UserPoolClient) (result *v1alpha1.UserPoolClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(userpoolclientsResource, c.ns, userPoolClient), &v1alpha1.UserPoolClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolClient), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserPoolClients) UpdateStatus(userPoolClient *v1alpha1.UserPoolClient) (*v1alpha1.UserPoolClient, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(userpoolclientsResource, "status", c.ns, userPoolClient), &v1alpha1.UserPoolClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolClient), err
}

// Delete takes name of the userPoolClient and deletes it. Returns an error if one occurs.
func (c *FakeUserPoolClients) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(userpoolclientsResource, c.ns, name), &v1alpha1.UserPoolClient{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserPoolClients) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(userpoolclientsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.UserPoolClientList{})
	return err
}

// Patch applies the patch and returns the patched userPoolClient.
func (c *FakeUserPoolClients) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.UserPoolClient, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(userpoolclientsResource, c.ns, name, pt, data, subresources...), &v1alpha1.UserPoolClient{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.UserPoolClient), err
}
