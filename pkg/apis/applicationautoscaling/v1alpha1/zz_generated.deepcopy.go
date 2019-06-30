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
func (in *ScalableTarget) DeepCopyInto(out *ScalableTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTarget.
func (in *ScalableTarget) DeepCopy() *ScalableTarget {
	if in == nil {
		return nil
	}
	out := new(ScalableTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalableTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetList) DeepCopyInto(out *ScalableTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalableTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetList.
func (in *ScalableTargetList) DeepCopy() *ScalableTargetList {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalableTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetOutputs) DeepCopyInto(out *ScalableTargetOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetOutputs.
func (in *ScalableTargetOutputs) DeepCopy() *ScalableTargetOutputs {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetSpec) DeepCopyInto(out *ScalableTargetSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.ScheduledActions != nil {
		in, out := &in.ScheduledActions, &out.ScheduledActions
		*out = make([]ScalableTargetSpecScheduledAction, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetSpec.
func (in *ScalableTargetSpec) DeepCopy() *ScalableTargetSpec {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetSpecScheduledAction) DeepCopyInto(out *ScalableTargetSpecScheduledAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetSpecScheduledAction.
func (in *ScalableTargetSpecScheduledAction) DeepCopy() *ScalableTargetSpecScheduledAction {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetSpecScheduledAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetSpecScheduledActionEndTime) DeepCopyInto(out *ScalableTargetSpecScheduledActionEndTime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetSpecScheduledActionEndTime.
func (in *ScalableTargetSpecScheduledActionEndTime) DeepCopy() *ScalableTargetSpecScheduledActionEndTime {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetSpecScheduledActionEndTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetSpecScheduledActionScalableTargetAction) DeepCopyInto(out *ScalableTargetSpecScheduledActionScalableTargetAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetSpecScheduledActionScalableTargetAction.
func (in *ScalableTargetSpecScheduledActionScalableTargetAction) DeepCopy() *ScalableTargetSpecScheduledActionScalableTargetAction {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetSpecScheduledActionScalableTargetAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetSpecScheduledActionStartTime) DeepCopyInto(out *ScalableTargetSpecScheduledActionStartTime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetSpecScheduledActionStartTime.
func (in *ScalableTargetSpecScheduledActionStartTime) DeepCopy() *ScalableTargetSpecScheduledActionStartTime {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetSpecScheduledActionStartTime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalableTargetStatus) DeepCopyInto(out *ScalableTargetStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalableTargetStatus.
func (in *ScalableTargetStatus) DeepCopy() *ScalableTargetStatus {
	if in == nil {
		return nil
	}
	out := new(ScalableTargetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicy) DeepCopyInto(out *ScalingPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicy.
func (in *ScalingPolicy) DeepCopy() *ScalingPolicy {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyList) DeepCopyInto(out *ScalingPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ScalingPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyList.
func (in *ScalingPolicyList) DeepCopy() *ScalingPolicyList {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ScalingPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyOutputs) DeepCopyInto(out *ScalingPolicyOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyOutputs.
func (in *ScalingPolicyOutputs) DeepCopy() *ScalingPolicyOutputs {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpec) DeepCopyInto(out *ScalingPolicySpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	in.StepScalingPolicyConfiguration.DeepCopyInto(&out.StepScalingPolicyConfiguration)
	in.TargetTrackingScalingPolicyConfiguration.DeepCopyInto(&out.TargetTrackingScalingPolicyConfiguration)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpec.
func (in *ScalingPolicySpec) DeepCopy() *ScalingPolicySpec {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpecStepScalingPolicyConfiguration) DeepCopyInto(out *ScalingPolicySpecStepScalingPolicyConfiguration) {
	*out = *in
	if in.StepAdjustments != nil {
		in, out := &in.StepAdjustments, &out.StepAdjustments
		*out = make([]ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpecStepScalingPolicyConfiguration.
func (in *ScalingPolicySpecStepScalingPolicyConfiguration) DeepCopy() *ScalingPolicySpecStepScalingPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpecStepScalingPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment) DeepCopyInto(out *ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment.
func (in *ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment) DeepCopy() *ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpecStepScalingPolicyConfigurationStepAdjustment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfiguration) DeepCopyInto(out *ScalingPolicySpecTargetTrackingScalingPolicyConfiguration) {
	*out = *in
	in.CustomizedMetricSpecification.DeepCopyInto(&out.CustomizedMetricSpecification)
	out.PredefinedMetricSpecification = in.PredefinedMetricSpecification
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpecTargetTrackingScalingPolicyConfiguration.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfiguration) DeepCopy() *ScalingPolicySpecTargetTrackingScalingPolicyConfiguration {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpecTargetTrackingScalingPolicyConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) DeepCopyInto(out *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) {
	*out = *in
	if in.Dimensions != nil {
		in, out := &in.Dimensions, &out.Dimensions
		*out = make([]ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification) DeepCopy() *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension) DeepCopyInto(out *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension) DeepCopy() *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpecTargetTrackingScalingPolicyConfigurationCustomizedMetricSpecificationDimension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification) DeepCopyInto(out *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification.
func (in *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification) DeepCopy() *ScalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicySpecTargetTrackingScalingPolicyConfigurationPredefinedMetricSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalingPolicyStatus) DeepCopyInto(out *ScalingPolicyStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalingPolicyStatus.
func (in *ScalingPolicyStatus) DeepCopy() *ScalingPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(ScalingPolicyStatus)
	in.DeepCopyInto(out)
	return out
}
