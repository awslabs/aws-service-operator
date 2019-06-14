// +build !ignore_autogenerated

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
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQuery) DeepCopyInto(out *NamedQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQuery.
func (in *NamedQuery) DeepCopy() *NamedQuery {
	if in == nil {
		return nil
	}
	out := new(NamedQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamedQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryList) DeepCopyInto(out *NamedQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamedQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryList.
func (in *NamedQueryList) DeepCopy() *NamedQueryList {
	if in == nil {
		return nil
	}
	out := new(NamedQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamedQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryOutputs) DeepCopyInto(out *NamedQueryOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryOutputs.
func (in *NamedQueryOutputs) DeepCopy() *NamedQueryOutputs {
	if in == nil {
		return nil
	}
	out := new(NamedQueryOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQuerySpec) DeepCopyInto(out *NamedQuerySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQuerySpec.
func (in *NamedQuerySpec) DeepCopy() *NamedQuerySpec {
	if in == nil {
		return nil
	}
	out := new(NamedQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryStatus) DeepCopyInto(out *NamedQueryStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryStatus.
func (in *NamedQueryStatus) DeepCopy() *NamedQueryStatus {
	if in == nil {
		return nil
	}
	out := new(NamedQueryStatus)
	in.DeepCopyInto(out)
	return out
}
