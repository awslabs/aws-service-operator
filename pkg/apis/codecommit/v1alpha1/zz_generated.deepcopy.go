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
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Repository) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryList) DeepCopyInto(out *RepositoryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Repository, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryList.
func (in *RepositoryList) DeepCopy() *RepositoryList {
	if in == nil {
		return nil
	}
	out := new(RepositoryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RepositoryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryOutputs) DeepCopyInto(out *RepositoryOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryOutputs.
func (in *RepositoryOutputs) DeepCopy() *RepositoryOutputs {
	if in == nil {
		return nil
	}
	out := new(RepositoryOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpec) DeepCopyInto(out *RepositorySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.Code = in.Code
	if in.Triggers != nil {
		in, out := &in.Triggers, &out.Triggers
		*out = make([]RepositorySpecTrigger, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpec.
func (in *RepositorySpec) DeepCopy() *RepositorySpec {
	if in == nil {
		return nil
	}
	out := new(RepositorySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpecCode) DeepCopyInto(out *RepositorySpecCode) {
	*out = *in
	out.S3 = in.S3
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpecCode.
func (in *RepositorySpecCode) DeepCopy() *RepositorySpecCode {
	if in == nil {
		return nil
	}
	out := new(RepositorySpecCode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpecCodeS3) DeepCopyInto(out *RepositorySpecCodeS3) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpecCodeS3.
func (in *RepositorySpecCodeS3) DeepCopy() *RepositorySpecCodeS3 {
	if in == nil {
		return nil
	}
	out := new(RepositorySpecCodeS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositorySpecTrigger) DeepCopyInto(out *RepositorySpecTrigger) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositorySpecTrigger.
func (in *RepositorySpecTrigger) DeepCopy() *RepositorySpecTrigger {
	if in == nil {
		return nil
	}
	out := new(RepositorySpecTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RepositoryStatus) DeepCopyInto(out *RepositoryStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoryStatus.
func (in *RepositoryStatus) DeepCopy() *RepositoryStatus {
	if in == nil {
		return nil
	}
	out := new(RepositoryStatus)
	in.DeepCopyInto(out)
	return out
}
