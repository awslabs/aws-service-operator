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
	v1alpha1 "awsoperator.io/pkg/apis/elasticsearch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDomains implements DomainInterface
type FakeDomains struct {
	Fake *FakeElasticsearchV1alpha1
	ns   string
}

var domainsResource = schema.GroupVersionResource{Group: "elasticsearch.awsoperator.io", Version: "v1alpha1", Resource: "domains"}

var domainsKind = schema.GroupVersionKind{Group: "elasticsearch.awsoperator.io", Version: "v1alpha1", Kind: "Domain"}

// Get takes name of the domain, and returns the corresponding domain object, and an error if there is any.
func (c *FakeDomains) Get(name string, options v1.GetOptions) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(domainsResource, c.ns, name), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// List takes label and field selectors, and returns the list of Domains that match those selectors.
func (c *FakeDomains) List(opts v1.ListOptions) (result *v1alpha1.DomainList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(domainsResource, domainsKind, c.ns, opts), &v1alpha1.DomainList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DomainList{ListMeta: obj.(*v1alpha1.DomainList).ListMeta}
	for _, item := range obj.(*v1alpha1.DomainList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested domains.
func (c *FakeDomains) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(domainsResource, c.ns, opts))

}

// Create takes the representation of a domain and creates it.  Returns the server's representation of the domain, and an error, if there is any.
func (c *FakeDomains) Create(domain *v1alpha1.Domain) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(domainsResource, c.ns, domain), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// Update takes the representation of a domain and updates it. Returns the server's representation of the domain, and an error, if there is any.
func (c *FakeDomains) Update(domain *v1alpha1.Domain) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(domainsResource, c.ns, domain), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDomains) UpdateStatus(domain *v1alpha1.Domain) (*v1alpha1.Domain, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(domainsResource, "status", c.ns, domain), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}

// Delete takes name of the domain and deletes it. Returns an error if one occurs.
func (c *FakeDomains) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(domainsResource, c.ns, name), &v1alpha1.Domain{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDomains) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(domainsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.DomainList{})
	return err
}

// Patch applies the patch and returns the patched domain.
func (c *FakeDomains) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Domain, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(domainsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Domain{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Domain), err
}
