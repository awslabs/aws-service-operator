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
func (in *ResolverEndpoint) DeepCopyInto(out *ResolverEndpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpoint.
func (in *ResolverEndpoint) DeepCopy() *ResolverEndpoint {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverEndpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverEndpointList) DeepCopyInto(out *ResolverEndpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResolverEndpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpointList.
func (in *ResolverEndpointList) DeepCopy() *ResolverEndpointList {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverEndpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverEndpointOutputs) DeepCopyInto(out *ResolverEndpointOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpointOutputs.
func (in *ResolverEndpointOutputs) DeepCopy() *ResolverEndpointOutputs {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpointOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverEndpointSpec) DeepCopyInto(out *ResolverEndpointSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.IpAddresses != nil {
		in, out := &in.IpAddresses, &out.IpAddresses
		*out = make([]ResolverEndpointSpecIpAddress, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroupRef != nil {
		in, out := &in.SecurityGroupRef, &out.SecurityGroupRef
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]ResolverEndpointSpecTag, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpointSpec.
func (in *ResolverEndpointSpec) DeepCopy() *ResolverEndpointSpec {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverEndpointSpecIpAddress) DeepCopyInto(out *ResolverEndpointSpecIpAddress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpointSpecIpAddress.
func (in *ResolverEndpointSpecIpAddress) DeepCopy() *ResolverEndpointSpecIpAddress {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpointSpecIpAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverEndpointSpecTag) DeepCopyInto(out *ResolverEndpointSpecTag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpointSpecTag.
func (in *ResolverEndpointSpecTag) DeepCopy() *ResolverEndpointSpecTag {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpointSpecTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverEndpointStatus) DeepCopyInto(out *ResolverEndpointStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverEndpointStatus.
func (in *ResolverEndpointStatus) DeepCopy() *ResolverEndpointStatus {
	if in == nil {
		return nil
	}
	out := new(ResolverEndpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRule) DeepCopyInto(out *ResolverRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRule.
func (in *ResolverRule) DeepCopy() *ResolverRule {
	if in == nil {
		return nil
	}
	out := new(ResolverRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociation) DeepCopyInto(out *ResolverRuleAssociation) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociation.
func (in *ResolverRuleAssociation) DeepCopy() *ResolverRuleAssociation {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverRuleAssociation) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationList) DeepCopyInto(out *ResolverRuleAssociationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResolverRuleAssociation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationList.
func (in *ResolverRuleAssociationList) DeepCopy() *ResolverRuleAssociationList {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverRuleAssociationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationOutputs) DeepCopyInto(out *ResolverRuleAssociationOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationOutputs.
func (in *ResolverRuleAssociationOutputs) DeepCopy() *ResolverRuleAssociationOutputs {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationSpec) DeepCopyInto(out *ResolverRuleAssociationSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationSpec.
func (in *ResolverRuleAssociationSpec) DeepCopy() *ResolverRuleAssociationSpec {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleAssociationStatus) DeepCopyInto(out *ResolverRuleAssociationStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleAssociationStatus.
func (in *ResolverRuleAssociationStatus) DeepCopy() *ResolverRuleAssociationStatus {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleAssociationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleList) DeepCopyInto(out *ResolverRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResolverRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleList.
func (in *ResolverRuleList) DeepCopy() *ResolverRuleList {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResolverRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleOutputs) DeepCopyInto(out *ResolverRuleOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleOutputs.
func (in *ResolverRuleOutputs) DeepCopy() *ResolverRuleOutputs {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleSpec) DeepCopyInto(out *ResolverRuleSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]ResolverRuleSpecTag, len(*in))
		copy(*out, *in)
	}
	if in.TargetIps != nil {
		in, out := &in.TargetIps, &out.TargetIps
		*out = make([]ResolverRuleSpecTargetIp, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleSpec.
func (in *ResolverRuleSpec) DeepCopy() *ResolverRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleSpecTag) DeepCopyInto(out *ResolverRuleSpecTag) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleSpecTag.
func (in *ResolverRuleSpecTag) DeepCopy() *ResolverRuleSpecTag {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleSpecTag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleSpecTargetIp) DeepCopyInto(out *ResolverRuleSpecTargetIp) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleSpecTargetIp.
func (in *ResolverRuleSpecTargetIp) DeepCopy() *ResolverRuleSpecTargetIp {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleSpecTargetIp)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResolverRuleStatus) DeepCopyInto(out *ResolverRuleStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResolverRuleStatus.
func (in *ResolverRuleStatus) DeepCopy() *ResolverRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ResolverRuleStatus)
	in.DeepCopyInto(out)
	return out
}
