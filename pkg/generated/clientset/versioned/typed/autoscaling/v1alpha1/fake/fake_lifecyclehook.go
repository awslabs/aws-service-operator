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
	v1alpha1 "awsoperator.io/pkg/apis/autoscaling/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLifecycleHooks implements LifecycleHookInterface
type FakeLifecycleHooks struct {
	Fake *FakeAutoscalingV1alpha1
	ns   string
}

var lifecyclehooksResource = schema.GroupVersionResource{Group: "autoscaling.awsoperator.io", Version: "v1alpha1", Resource: "lifecyclehooks"}

var lifecyclehooksKind = schema.GroupVersionKind{Group: "autoscaling.awsoperator.io", Version: "v1alpha1", Kind: "LifecycleHook"}

// Get takes name of the lifecycleHook, and returns the corresponding lifecycleHook object, and an error if there is any.
func (c *FakeLifecycleHooks) Get(name string, options v1.GetOptions) (result *v1alpha1.LifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(lifecyclehooksResource, c.ns, name), &v1alpha1.LifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecycleHook), err
}

// List takes label and field selectors, and returns the list of LifecycleHooks that match those selectors.
func (c *FakeLifecycleHooks) List(opts v1.ListOptions) (result *v1alpha1.LifecycleHookList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(lifecyclehooksResource, lifecyclehooksKind, c.ns, opts), &v1alpha1.LifecycleHookList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LifecycleHookList{ListMeta: obj.(*v1alpha1.LifecycleHookList).ListMeta}
	for _, item := range obj.(*v1alpha1.LifecycleHookList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested lifecycleHooks.
func (c *FakeLifecycleHooks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(lifecyclehooksResource, c.ns, opts))

}

// Create takes the representation of a lifecycleHook and creates it.  Returns the server's representation of the lifecycleHook, and an error, if there is any.
func (c *FakeLifecycleHooks) Create(lifecycleHook *v1alpha1.LifecycleHook) (result *v1alpha1.LifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(lifecyclehooksResource, c.ns, lifecycleHook), &v1alpha1.LifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecycleHook), err
}

// Update takes the representation of a lifecycleHook and updates it. Returns the server's representation of the lifecycleHook, and an error, if there is any.
func (c *FakeLifecycleHooks) Update(lifecycleHook *v1alpha1.LifecycleHook) (result *v1alpha1.LifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(lifecyclehooksResource, c.ns, lifecycleHook), &v1alpha1.LifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecycleHook), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLifecycleHooks) UpdateStatus(lifecycleHook *v1alpha1.LifecycleHook) (*v1alpha1.LifecycleHook, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(lifecyclehooksResource, "status", c.ns, lifecycleHook), &v1alpha1.LifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecycleHook), err
}

// Delete takes name of the lifecycleHook and deletes it. Returns an error if one occurs.
func (c *FakeLifecycleHooks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(lifecyclehooksResource, c.ns, name), &v1alpha1.LifecycleHook{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLifecycleHooks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(lifecyclehooksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.LifecycleHookList{})
	return err
}

// Patch applies the patch and returns the patched lifecycleHook.
func (c *FakeLifecycleHooks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.LifecycleHook, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(lifecyclehooksResource, c.ns, name, pt, data, subresources...), &v1alpha1.LifecycleHook{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.LifecycleHook), err
}
