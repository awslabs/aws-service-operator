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
func (in *ScalingPlan) DeepCopyInto(out *ScalingPlan) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlan.
func (in *ScalingPlan) DeepCopy() *ScalingPlan {
	if in == nil {
		return nil
	}
	out := new(ScalingPlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPlan) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanList) DeepCopyInto(out *ScalingPlanList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPlan, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanList.
func (in *ScalingPlanList) DeepCopy() *ScalingPlanList {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPlanList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanOutputs) DeepCopyInto(out *ScalingPlanOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanOutputs.
func (in *ScalingPlanOutputs) DeepCopy() *ScalingPlanOutputs {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpec) DeepCopyInto(out *ScalingPlanSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	in.ApplicationSource.DeepCopyInto(&out.ApplicationSource)
	if in.ScalingInstructions != nil {
		in, out := &in.ScalingInstructions, &out.ScalingInstructions
		*out = make([]ScalingPlanSpecScalingInstruction, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpec.
func (in *ScalingPlanSpec) DeepCopy() *ScalingPlanSpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecApplicationSource) DeepCopyInto(out *ScalingPlanSpecApplicationSource) {
	*out = *in
	if in.TagFilters != nil {
		in, out := &in.TagFilters, &out.TagFilters
		*out = make([]ScalingPlanSpecApplicationSourceTagFilter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecApplicationSource.
func (in *ScalingPlanSpecApplicationSource) DeepCopy() *ScalingPlanSpecApplicationSource {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecApplicationSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecApplicationSourceTagFilter) DeepCopyInto(out *ScalingPlanSpecApplicationSourceTagFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecApplicationSourceTagFilter.
func (in *ScalingPlanSpecApplicationSourceTagFilter) DeepCopy() *ScalingPlanSpecApplicationSourceTagFilter {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecApplicationSourceTagFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstruction) DeepCopyInto(out *ScalingPlanSpecScalingInstruction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstruction.
func (in *ScalingPlanSpecScalingInstruction) DeepCopy() *ScalingPlanSpecScalingInstruction {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstruction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension) DeepCopyInto(out *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension.
func (in *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension) DeepCopy() *ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionCustomizedLoadMetricSpecificationDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionPredefinedLoadMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfiguration.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfiguration {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationCustomizedScalingMetricSpecificationDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification) DeepCopyInto(out *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification.
func (in *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification) DeepCopy() *ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanSpecScalingInstructionTargetTrackingConfigurationPredefinedScalingMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPlanStatus) DeepCopyInto(out *ScalingPlanStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPlanStatus.
func (in *ScalingPlanStatus) DeepCopy() *ScalingPlanStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPlanStatus)
	in.DeepCopyInto(out)
	return out
}
