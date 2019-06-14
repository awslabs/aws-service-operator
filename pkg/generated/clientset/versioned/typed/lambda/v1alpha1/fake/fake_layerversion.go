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
	v1alpha1 "awsoperator.io/pkg/apis/lambda/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLayerVersions implements LayerVersionInterface
type FakeLayerVersions struct {
	Fake *FakeLambdaV1alpha1
	ns   string
}

var layerversionsResource = schema.GroupVersionResource{Group: "lambda.awsoperator.io", Version: "v1alpha1", Resource: "layerversions"}

var layerversionsKind = schema.GroupVersionKind{Group: "lambda.awsoperator.io", Version: "v1alpha1", Kind: "LayerVersion"}

// Get takes name of the layerVersion, and returns the corresponding layerVersion object, and an error if there is any.
func (c *FakeLayerVersions) Get(name string, options v1.GetOptions) (result *v1alpha1.LayerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(layerversionsResource, c.ns, name), &v1alpha1.LayerVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersion), err
}

// List takes label and field selectors, and returns the list of LayerVersions that match those selectors.
func (c *FakeLayerVersions) List(opts v1.ListOptions) (result *v1alpha1.LayerVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(layerversionsResource, layerversionsKind, c.ns, opts), &v1alpha1.LayerVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LayerVersionList{ListMeta: obj.(*v1alpha1.LayerVersionList).ListMeta}
	for _, item := range obj.(*v1alpha1.LayerVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested layerVersions.
func (c *FakeLayerVersions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(layerversionsResource, c.ns, opts))

}

// Create takes the representation of a layerVersion and creates it.  Returns the server's representation of the layerVersion, and an error, if there is any.
func (c *FakeLayerVersions) Create(layerVersion *v1alpha1.LayerVersion) (result *v1alpha1.LayerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(layerversionsResource, c.ns, layerVersion), &v1alpha1.LayerVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersion), err
}

// Update takes the representation of a layerVersion and updates it. Returns the server's representation of the layerVersion, and an error, if there is any.
func (c *FakeLayerVersions) Update(layerVersion *v1alpha1.LayerVersion) (result *v1alpha1.LayerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(layerversionsResource, c.ns, layerVersion), &v1alpha1.LayerVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLayerVersions) UpdateStatus(layerVersion *v1alpha1.LayerVersion) (*v1alpha1.LayerVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(layerversionsResource, "status", c.ns, layerVersion), &v1alpha1.LayerVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersion), err
}

// Delete takes name of the layerVersion and deletes it. Returns an error if one occurs.
func (c *FakeLayerVersions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(layerversionsResource, c.ns, name), &v1alpha1.LayerVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLayerVersions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(layerversionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LayerVersionList{})
	return err
}

// Patch applies the patch and returns the patched layerVersion.
func (c *FakeLayerVersions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LayerVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(layerversionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.LayerVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LayerVersion), err
}
