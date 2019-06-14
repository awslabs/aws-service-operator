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
	v1alpha1 "awsoperator.io/pkg/apis/fsx/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFileSystems implements FileSystemInterface
type FakeFileSystems struct {
	Fake *FakeFsxV1alpha1
	ns   string
}

var filesystemsResource = schema.GroupVersionResource{Group: "fsx.awsoperator.io", Version: "v1alpha1", Resource: "filesystems"}

var filesystemsKind = schema.GroupVersionKind{Group: "fsx.awsoperator.io", Version: "v1alpha1", Kind: "FileSystem"}

// Get takes name of the fileSystem, and returns the corresponding fileSystem object, and an error if there is any.
func (c *FakeFileSystems) Get(name string, options v1.GetOptions) (result *v1alpha1.FileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(filesystemsResource, c.ns, name), &v1alpha1.FileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FileSystem), err
}

// List takes label and field selectors, and returns the list of FileSystems that match those selectors.
func (c *FakeFileSystems) List(opts v1.ListOptions) (result *v1alpha1.FileSystemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(filesystemsResource, filesystemsKind, c.ns, opts), &v1alpha1.FileSystemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.FileSystemList{ListMeta: obj.(*v1alpha1.FileSystemList).ListMeta}
	for _, item := range obj.(*v1alpha1.FileSystemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested fileSystems.
func (c *FakeFileSystems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(filesystemsResource, c.ns, opts))

}

// Create takes the representation of a fileSystem and creates it.  Returns the server's representation of the fileSystem, and an error, if there is any.
func (c *FakeFileSystems) Create(fileSystem *v1alpha1.FileSystem) (result *v1alpha1.FileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(filesystemsResource, c.ns, fileSystem), &v1alpha1.FileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FileSystem), err
}

// Update takes the representation of a fileSystem and updates it. Returns the server's representation of the fileSystem, and an error, if there is any.
func (c *FakeFileSystems) Update(fileSystem *v1alpha1.FileSystem) (result *v1alpha1.FileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(filesystemsResource, c.ns, fileSystem), &v1alpha1.FileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FileSystem), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFileSystems) UpdateStatus(fileSystem *v1alpha1.FileSystem) (*v1alpha1.FileSystem, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(filesystemsResource, "status", c.ns, fileSystem), &v1alpha1.FileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FileSystem), err
}

// Delete takes name of the fileSystem and deletes it. Returns an error if one occurs.
func (c *FakeFileSystems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(filesystemsResource, c.ns, name), &v1alpha1.FileSystem{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFileSystems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(filesystemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.FileSystemList{})
	return err
}

// Patch applies the patch and returns the patched fileSystem.
func (c *FakeFileSystems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.FileSystem, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(filesystemsResource, c.ns, name, pt, data, subresources...), &v1alpha1.FileSystem{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.FileSystem), err
}
