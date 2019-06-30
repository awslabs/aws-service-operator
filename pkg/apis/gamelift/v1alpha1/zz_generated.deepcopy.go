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
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alias) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasList) DeepCopyInto(out *AliasList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasList.
func (in *AliasList) DeepCopy() *AliasList {
	if in == nil {
		return nil
	}
	out := new(AliasList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AliasList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasOutputs) DeepCopyInto(out *AliasOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasOutputs.
func (in *AliasOutputs) DeepCopy() *AliasOutputs {
	if in == nil {
		return nil
	}
	out := new(AliasOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpec) DeepCopyInto(out *AliasSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.RoutingStrategy = in.RoutingStrategy
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpec.
func (in *AliasSpec) DeepCopy() *AliasSpec {
	if in == nil {
		return nil
	}
	out := new(AliasSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasSpecRoutingStrategy) DeepCopyInto(out *AliasSpecRoutingStrategy) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasSpecRoutingStrategy.
func (in *AliasSpecRoutingStrategy) DeepCopy() *AliasSpecRoutingStrategy {
	if in == nil {
		return nil
	}
	out := new(AliasSpecRoutingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasStatus) DeepCopyInto(out *AliasStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasStatus.
func (in *AliasStatus) DeepCopy() *AliasStatus {
	if in == nil {
		return nil
	}
	out := new(AliasStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Build) DeepCopyInto(out *Build) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Build.
func (in *Build) DeepCopy() *Build {
	if in == nil {
		return nil
	}
	out := new(Build)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Build) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildList) DeepCopyInto(out *BuildList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Build, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildList.
func (in *BuildList) DeepCopy() *BuildList {
	if in == nil {
		return nil
	}
	out := new(BuildList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BuildList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildOutputs) DeepCopyInto(out *BuildOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildOutputs.
func (in *BuildOutputs) DeepCopy() *BuildOutputs {
	if in == nil {
		return nil
	}
	out := new(BuildOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildSpec) DeepCopyInto(out *BuildSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.StorageLocation = in.StorageLocation
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildSpec.
func (in *BuildSpec) DeepCopy() *BuildSpec {
	if in == nil {
		return nil
	}
	out := new(BuildSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildSpecStorageLocation) DeepCopyInto(out *BuildSpecStorageLocation) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildSpecStorageLocation.
func (in *BuildSpecStorageLocation) DeepCopy() *BuildSpecStorageLocation {
	if in == nil {
		return nil
	}
	out := new(BuildSpecStorageLocation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BuildStatus) DeepCopyInto(out *BuildStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BuildStatus.
func (in *BuildStatus) DeepCopy() *BuildStatus {
	if in == nil {
		return nil
	}
	out := new(BuildStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Fleet) DeepCopyInto(out *Fleet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Fleet.
func (in *Fleet) DeepCopy() *Fleet {
	if in == nil {
		return nil
	}
	out := new(Fleet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Fleet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetList) DeepCopyInto(out *FleetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Fleet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetList.
func (in *FleetList) DeepCopy() *FleetList {
	if in == nil {
		return nil
	}
	out := new(FleetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FleetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetOutputs) DeepCopyInto(out *FleetOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetOutputs.
func (in *FleetOutputs) DeepCopy() *FleetOutputs {
	if in == nil {
		return nil
	}
	out := new(FleetOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpec) DeepCopyInto(out *FleetSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.EC2InboundPermissions != nil {
		in, out := &in.EC2InboundPermissions, &out.EC2InboundPermissions
		*out = make([]FleetSpecEC2InboundPermission, len(*in))
		copy(*out, *in)
	}
	if in.LogPaths != nil {
		in, out := &in.LogPaths, &out.LogPaths
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpec.
func (in *FleetSpec) DeepCopy() *FleetSpec {
	if in == nil {
		return nil
	}
	out := new(FleetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetSpecEC2InboundPermission) DeepCopyInto(out *FleetSpecEC2InboundPermission) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetSpecEC2InboundPermission.
func (in *FleetSpecEC2InboundPermission) DeepCopy() *FleetSpecEC2InboundPermission {
	if in == nil {
		return nil
	}
	out := new(FleetSpecEC2InboundPermission)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FleetStatus) DeepCopyInto(out *FleetStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FleetStatus.
func (in *FleetStatus) DeepCopy() *FleetStatus {
	if in == nil {
		return nil
	}
	out := new(FleetStatus)
	in.DeepCopyInto(out)
	return out
}
