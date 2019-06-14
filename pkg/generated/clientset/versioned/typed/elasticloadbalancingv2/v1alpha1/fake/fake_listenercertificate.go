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

// FakeListenerCertificates implements ListenerCertificateInterface
type FakeListenerCertificates struct {
	Fake *FakeElasticloadbalancingv2V1alpha1
	ns   string
}

var listenercertificatesResource = schema.GroupVersionResource{Group: "elasticloadbalancingv2.awsoperator.io", Version: "v1alpha1", Resource: "listenercertificates"}

var listenercertificatesKind = schema.GroupVersionKind{Group: "elasticloadbalancingv2.awsoperator.io", Version: "v1alpha1", Kind: "ListenerCertificate"}

// Get takes name of the listenerCertificate, and returns the corresponding listenerCertificate object, and an error if there is any.
func (c *FakeListenerCertificates) Get(name string, options v1.GetOptions) (result *v1alpha1.ListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(listenercertificatesResource, c.ns, name), &v1alpha1.ListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerCertificate), err
}

// List takes label and field selectors, and returns the list of ListenerCertificates that match those selectors.
func (c *FakeListenerCertificates) List(opts v1.ListOptions) (result *v1alpha1.ListenerCertificateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(listenercertificatesResource, listenercertificatesKind, c.ns, opts), &v1alpha1.ListenerCertificateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ListenerCertificateList{ListMeta: obj.(*v1alpha1.ListenerCertificateList).ListMeta}
	for _, item := range obj.(*v1alpha1.ListenerCertificateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested listenerCertificates.
func (c *FakeListenerCertificates) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(listenercertificatesResource, c.ns, opts))

}

// Create takes the representation of a listenerCertificate and creates it.  Returns the server's representation of the listenerCertificate, and an error, if there is any.
func (c *FakeListenerCertificates) Create(listenerCertificate *v1alpha1.ListenerCertificate) (result *v1alpha1.ListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(listenercertificatesResource, c.ns, listenerCertificate), &v1alpha1.ListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerCertificate), err
}

// Update takes the representation of a listenerCertificate and updates it. Returns the server's representation of the listenerCertificate, and an error, if there is any.
func (c *FakeListenerCertificates) Update(listenerCertificate *v1alpha1.ListenerCertificate) (result *v1alpha1.ListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(listenercertificatesResource, c.ns, listenerCertificate), &v1alpha1.ListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerCertificate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeListenerCertificates) UpdateStatus(listenerCertificate *v1alpha1.ListenerCertificate) (*v1alpha1.ListenerCertificate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(listenercertificatesResource, "status", c.ns, listenerCertificate), &v1alpha1.ListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerCertificate), err
}

// Delete takes name of the listenerCertificate and deletes it. Returns an error if one occurs.
func (c *FakeListenerCertificates) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(listenercertificatesResource, c.ns, name), &v1alpha1.ListenerCertificate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeListenerCertificates) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(listenercertificatesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ListenerCertificateList{})
	return err
}

// Patch applies the patch and returns the patched listenerCertificate.
func (c *FakeListenerCertificates) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ListenerCertificate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(listenercertificatesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ListenerCertificate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ListenerCertificate), err
}
