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

// FakeSqlInjectionMatchSets implements SqlInjectionMatchSetInterface
type FakeSqlInjectionMatchSets struct {
	Fake *FakeWafregionalV1alpha1
	ns   string
}

var sqlinjectionmatchsetsResource = schema.GroupVersionResource{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Resource: "sqlinjectionmatchsets"}

var sqlinjectionmatchsetsKind = schema.GroupVersionKind{Group: "wafregional.awsoperator.io", Version: "v1alpha1", Kind: "SqlInjectionMatchSet"}

// Get takes name of the sqlInjectionMatchSet, and returns the corresponding sqlInjectionMatchSet object, and an error if there is any.
func (c *FakeSqlInjectionMatchSets) Get(name string, options v1.GetOptions) (result *v1alpha1.SqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(sqlinjectionmatchsetsResource, c.ns, name), &v1alpha1.SqlInjectionMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlInjectionMatchSet), err
}

// List takes label and field selectors, and returns the list of SqlInjectionMatchSets that match those selectors.
func (c *FakeSqlInjectionMatchSets) List(opts v1.ListOptions) (result *v1alpha1.SqlInjectionMatchSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(sqlinjectionmatchsetsResource, sqlinjectionmatchsetsKind, c.ns, opts), &v1alpha1.SqlInjectionMatchSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SqlInjectionMatchSetList{ListMeta: obj.(*v1alpha1.SqlInjectionMatchSetList).ListMeta}
	for _, item := range obj.(*v1alpha1.SqlInjectionMatchSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested sqlInjectionMatchSets.
func (c *FakeSqlInjectionMatchSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(sqlinjectionmatchsetsResource, c.ns, opts))

}

// Create takes the representation of a sqlInjectionMatchSet and creates it.  Returns the server's representation of the sqlInjectionMatchSet, and an error, if there is any.
func (c *FakeSqlInjectionMatchSets) Create(sqlInjectionMatchSet *v1alpha1.SqlInjectionMatchSet) (result *v1alpha1.SqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(sqlinjectionmatchsetsResource, c.ns, sqlInjectionMatchSet), &v1alpha1.SqlInjectionMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlInjectionMatchSet), err
}

// Update takes the representation of a sqlInjectionMatchSet and updates it. Returns the server's representation of the sqlInjectionMatchSet, and an error, if there is any.
func (c *FakeSqlInjectionMatchSets) Update(sqlInjectionMatchSet *v1alpha1.SqlInjectionMatchSet) (result *v1alpha1.SqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(sqlinjectionmatchsetsResource, c.ns, sqlInjectionMatchSet), &v1alpha1.SqlInjectionMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlInjectionMatchSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSqlInjectionMatchSets) UpdateStatus(sqlInjectionMatchSet *v1alpha1.SqlInjectionMatchSet) (*v1alpha1.SqlInjectionMatchSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(sqlinjectionmatchsetsResource, "status", c.ns, sqlInjectionMatchSet), &v1alpha1.SqlInjectionMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlInjectionMatchSet), err
}

// Delete takes name of the sqlInjectionMatchSet and deletes it. Returns an error if one occurs.
func (c *FakeSqlInjectionMatchSets) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(sqlinjectionmatchsetsResource, c.ns, name), &v1alpha1.SqlInjectionMatchSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSqlInjectionMatchSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(sqlinjectionmatchsetsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.SqlInjectionMatchSetList{})
	return err
}

// Patch applies the patch and returns the patched sqlInjectionMatchSet.
func (c *FakeSqlInjectionMatchSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SqlInjectionMatchSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(sqlinjectionmatchsetsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SqlInjectionMatchSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SqlInjectionMatchSet), err
}
