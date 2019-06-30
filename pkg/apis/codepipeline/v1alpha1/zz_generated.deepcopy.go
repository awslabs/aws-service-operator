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
func (in *CustomActionType) DeepCopyInto(out *CustomActionType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionType.
func (in *CustomActionType) DeepCopy() *CustomActionType {
	if in == nil {
		return nil
	}
	out := new(CustomActionType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomActionType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeList) DeepCopyInto(out *CustomActionTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CustomActionType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeList.
func (in *CustomActionTypeList) DeepCopy() *CustomActionTypeList {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CustomActionTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeOutputs) DeepCopyInto(out *CustomActionTypeOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeOutputs.
func (in *CustomActionTypeOutputs) DeepCopy() *CustomActionTypeOutputs {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeSpec) DeepCopyInto(out *CustomActionTypeSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	if in.ConfigurationProperties != nil {
		in, out := &in.ConfigurationProperties, &out.ConfigurationProperties
		*out = make([]CustomActionTypeSpecConfigurationProperty, len(*in))
		copy(*out, *in)
	}
	out.InputArtifactDetails = in.InputArtifactDetails
	out.OutputArtifactDetails = in.OutputArtifactDetails
	out.Settings = in.Settings
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeSpec.
func (in *CustomActionTypeSpec) DeepCopy() *CustomActionTypeSpec {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeSpecConfigurationProperty) DeepCopyInto(out *CustomActionTypeSpecConfigurationProperty) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeSpecConfigurationProperty.
func (in *CustomActionTypeSpecConfigurationProperty) DeepCopy() *CustomActionTypeSpecConfigurationProperty {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeSpecConfigurationProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeSpecInputArtifactDetails) DeepCopyInto(out *CustomActionTypeSpecInputArtifactDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeSpecInputArtifactDetails.
func (in *CustomActionTypeSpecInputArtifactDetails) DeepCopy() *CustomActionTypeSpecInputArtifactDetails {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeSpecInputArtifactDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeSpecOutputArtifactDetails) DeepCopyInto(out *CustomActionTypeSpecOutputArtifactDetails) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeSpecOutputArtifactDetails.
func (in *CustomActionTypeSpecOutputArtifactDetails) DeepCopy() *CustomActionTypeSpecOutputArtifactDetails {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeSpecOutputArtifactDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeSpecSettings) DeepCopyInto(out *CustomActionTypeSpecSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeSpecSettings.
func (in *CustomActionTypeSpecSettings) DeepCopy() *CustomActionTypeSpecSettings {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeSpecSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomActionTypeStatus) DeepCopyInto(out *CustomActionTypeStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomActionTypeStatus.
func (in *CustomActionTypeStatus) DeepCopy() *CustomActionTypeStatus {
	if in == nil {
		return nil
	}
	out := new(CustomActionTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pipeline) DeepCopyInto(out *Pipeline) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pipeline.
func (in *Pipeline) DeepCopy() *Pipeline {
	if in == nil {
		return nil
	}
	out := new(Pipeline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pipeline) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineList) DeepCopyInto(out *PipelineList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Pipeline, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineList.
func (in *PipelineList) DeepCopy() *PipelineList {
	if in == nil {
		return nil
	}
	out := new(PipelineList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineOutputs) DeepCopyInto(out *PipelineOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineOutputs.
func (in *PipelineOutputs) DeepCopy() *PipelineOutputs {
	if in == nil {
		return nil
	}
	out := new(PipelineOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpec) DeepCopyInto(out *PipelineSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.ArtifactStore = in.ArtifactStore
	if in.ArtifactStores != nil {
		in, out := &in.ArtifactStores, &out.ArtifactStores
		*out = make([]PipelineSpecArtifactStore, len(*in))
		copy(*out, *in)
	}
	if in.DisableInboundStageTransitions != nil {
		in, out := &in.DisableInboundStageTransitions, &out.DisableInboundStageTransitions
		*out = make([]PipelineSpecDisableInboundStageTransition, len(*in))
		copy(*out, *in)
	}
	if in.Stages != nil {
		in, out := &in.Stages, &out.Stages
		*out = make([]PipelineSpecStage, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpec.
func (in *PipelineSpec) DeepCopy() *PipelineSpec {
	if in == nil {
		return nil
	}
	out := new(PipelineSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecArtifactStore) DeepCopyInto(out *PipelineSpecArtifactStore) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecArtifactStore.
func (in *PipelineSpecArtifactStore) DeepCopy() *PipelineSpecArtifactStore {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecArtifactStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecArtifactStoreArtifactStore) DeepCopyInto(out *PipelineSpecArtifactStoreArtifactStore) {
	*out = *in
	out.EncryptionKey = in.EncryptionKey
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecArtifactStoreArtifactStore.
func (in *PipelineSpecArtifactStoreArtifactStore) DeepCopy() *PipelineSpecArtifactStoreArtifactStore {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecArtifactStoreArtifactStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecArtifactStoreArtifactStoreEncryptionKey) DeepCopyInto(out *PipelineSpecArtifactStoreArtifactStoreEncryptionKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecArtifactStoreArtifactStoreEncryptionKey.
func (in *PipelineSpecArtifactStoreArtifactStoreEncryptionKey) DeepCopy() *PipelineSpecArtifactStoreArtifactStoreEncryptionKey {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecArtifactStoreArtifactStoreEncryptionKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecArtifactStoreEncryptionKey) DeepCopyInto(out *PipelineSpecArtifactStoreEncryptionKey) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecArtifactStoreEncryptionKey.
func (in *PipelineSpecArtifactStoreEncryptionKey) DeepCopy() *PipelineSpecArtifactStoreEncryptionKey {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecArtifactStoreEncryptionKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecDisableInboundStageTransition) DeepCopyInto(out *PipelineSpecDisableInboundStageTransition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecDisableInboundStageTransition.
func (in *PipelineSpecDisableInboundStageTransition) DeepCopy() *PipelineSpecDisableInboundStageTransition {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecDisableInboundStageTransition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStage) DeepCopyInto(out *PipelineSpecStage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStage.
func (in *PipelineSpecStage) DeepCopy() *PipelineSpecStage {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStageAction) DeepCopyInto(out *PipelineSpecStageAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStageAction.
func (in *PipelineSpecStageAction) DeepCopy() *PipelineSpecStageAction {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStageAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStageActionActionTypeRef) DeepCopyInto(out *PipelineSpecStageActionActionTypeRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStageActionActionTypeRef.
func (in *PipelineSpecStageActionActionTypeRef) DeepCopy() *PipelineSpecStageActionActionTypeRef {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStageActionActionTypeRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStageActionConfiguration) DeepCopyInto(out *PipelineSpecStageActionConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStageActionConfiguration.
func (in *PipelineSpecStageActionConfiguration) DeepCopy() *PipelineSpecStageActionConfiguration {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStageActionConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStageActionInputArtifact) DeepCopyInto(out *PipelineSpecStageActionInputArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStageActionInputArtifact.
func (in *PipelineSpecStageActionInputArtifact) DeepCopy() *PipelineSpecStageActionInputArtifact {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStageActionInputArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStageActionOutputArtifact) DeepCopyInto(out *PipelineSpecStageActionOutputArtifact) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStageActionOutputArtifact.
func (in *PipelineSpecStageActionOutputArtifact) DeepCopy() *PipelineSpecStageActionOutputArtifact {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStageActionOutputArtifact)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineSpecStageBlocker) DeepCopyInto(out *PipelineSpecStageBlocker) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineSpecStageBlocker.
func (in *PipelineSpecStageBlocker) DeepCopy() *PipelineSpecStageBlocker {
	if in == nil {
		return nil
	}
	out := new(PipelineSpecStageBlocker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineStatus) DeepCopyInto(out *PipelineStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineStatus.
func (in *PipelineStatus) DeepCopy() *PipelineStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Webhook) DeepCopyInto(out *Webhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	out.Outputs = in.Outputs
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Webhook.
func (in *Webhook) DeepCopy() *Webhook {
	if in == nil {
		return nil
	}
	out := new(Webhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Webhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookList) DeepCopyInto(out *WebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Webhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookList.
func (in *WebhookList) DeepCopy() *WebhookList {
	if in == nil {
		return nil
	}
	out := new(WebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookOutputs) DeepCopyInto(out *WebhookOutputs) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookOutputs.
func (in *WebhookOutputs) DeepCopy() *WebhookOutputs {
	if in == nil {
		return nil
	}
	out := new(WebhookOutputs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSpec) DeepCopyInto(out *WebhookSpec) {
	*out = *in
	in.CloudFormationMeta.DeepCopyInto(&out.CloudFormationMeta)
	out.AuthenticationConfiguration = in.AuthenticationConfiguration
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]WebhookSpecFilter, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSpec.
func (in *WebhookSpec) DeepCopy() *WebhookSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSpecAuthenticationConfiguration) DeepCopyInto(out *WebhookSpecAuthenticationConfiguration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSpecAuthenticationConfiguration.
func (in *WebhookSpecAuthenticationConfiguration) DeepCopy() *WebhookSpecAuthenticationConfiguration {
	if in == nil {
		return nil
	}
	out := new(WebhookSpecAuthenticationConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookSpecFilter) DeepCopyInto(out *WebhookSpecFilter) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookSpecFilter.
func (in *WebhookSpecFilter) DeepCopy() *WebhookSpecFilter {
	if in == nil {
		return nil
	}
	out := new(WebhookSpecFilter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookStatus) DeepCopyInto(out *WebhookStatus) {
	*out = *in
	in.StatusMeta.DeepCopyInto(&out.StatusMeta)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookStatus.
func (in *WebhookStatus) DeepCopy() *WebhookStatus {
	if in == nil {
		return nil
	}
	out := new(WebhookStatus)
	in.DeepCopyInto(out)
	return out
}
