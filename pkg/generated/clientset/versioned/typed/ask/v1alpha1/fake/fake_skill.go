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
	v1alpha1 "awsoperator.io/pkg/apis/ask/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSkills implements SkillInterface
type FakeSkills struct {
	Fake *FakeAskV1alpha1
	ns   string
}

var skillsResource = schema.GroupVersionResource{Group: "ask.awsoperator.io", Version: "v1alpha1", Resource: "skills"}

var skillsKind = schema.GroupVersionKind{Group: "ask.awsoperator.io", Version: "v1alpha1", Kind: "Skill"}

// Get takes name of the skill, and returns the corresponding skill object, and an error if there is any.
func (c *FakeSkills) Get(name string, options v1.GetOptions) (result *v1alpha1.Skill, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(skillsResource, c.ns, name), &v1alpha1.Skill{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Skill), err
}

// List takes label and field selectors, and returns the list of Skills that match those selectors.
func (c *FakeSkills) List(opts v1.ListOptions) (result *v1alpha1.SkillList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(skillsResource, skillsKind, c.ns, opts), &v1alpha1.SkillList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SkillList{ListMeta: obj.(*v1alpha1.SkillList).ListMeta}
	for _, item := range obj.(*v1alpha1.SkillList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested skills.
func (c *FakeSkills) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(skillsResource, c.ns, opts))

}

// Create takes the representation of a skill and creates it.  Returns the server's representation of the skill, and an error, if there is any.
func (c *FakeSkills) Create(skill *v1alpha1.Skill) (result *v1alpha1.Skill, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(skillsResource, c.ns, skill), &v1alpha1.Skill{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Skill), err
}

// Update takes the representation of a skill and updates it. Returns the server's representation of the skill, and an error, if there is any.
func (c *FakeSkills) Update(skill *v1alpha1.Skill) (result *v1alpha1.Skill, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(skillsResource, c.ns, skill), &v1alpha1.Skill{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Skill), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSkills) UpdateStatus(skill *v1alpha1.Skill) (*v1alpha1.Skill, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(skillsResource, "status", c.ns, skill), &v1alpha1.Skill{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Skill), err
}

// Delete takes name of the skill and deletes it. Returns an error if one occurs.
func (c *FakeSkills) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(skillsResource, c.ns, name), &v1alpha1.Skill{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSkills) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(skillsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SkillList{})
	return err
}

// Patch applies the patch and returns the patched skill.
func (c *FakeSkills) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Skill, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(skillsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Skill{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Skill), err
}
