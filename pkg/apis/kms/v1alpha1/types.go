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

// Code generated by make generate. DO NOT EDIT.

// Package v1alpha1 is the v1alpha1 version of the awsoperator.io api.
package v1alpha1

import (
	metav1alpha1 "awsoperator.io/pkg/apis/meta/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Alias is a specification for a Alias resource
type Alias struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec AliasSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status AliasStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs AliasOutputs `json:"outputs"`
}

// AliasSpec is the spec for the Alias resource
type AliasSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AliasName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName string `json:"aliasName,omitempty" cloudformation:"AliasName,Parameter"`

	// TargetKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyRef string `json:"targetKeyRef,omitempty" cloudformation:"TargetKeyId,Parameter"`
}

// AliasStatus is the status for a  resource
type AliasStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// AliasOutputs is the status for a  resource
type AliasOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AliasList is a list of Alias resources
type AliasList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Alias `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Key is a specification for a Key resource
type Key struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec KeySpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status KeyStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs KeyOutputs `json:"outputs"`
}

// KeySpec is the spec for the Key resource
type KeySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// EnableKeyRotation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
	EnableKeyRotation bool `json:"enableKeyRotation,omitempty" cloudformation:"EnableKeyRotation,Parameter"`

	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// KeyPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
	KeyPolicy KeySpecKeyPolicy `json:"keyPolicy,omitempty" cloudformation:"KeyPolicy,Parameter"`

	// KeyUsage http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
	KeyUsage string `json:"keyUsage,omitempty" cloudformation:"KeyUsage,Parameter"`

	// PendingWindowInDays http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-pendingwindowindays
	PendingWindowInDays int32 `json:"pendingWindowInDays,omitempty" cloudformation:"PendingWindowInDays,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-tags
	Tags []KeySpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// KeySpecKeyPolicy is the definition for a Key resource
type KeySpecKeyPolicy struct {
}

// KeySpecTag is the definition for a Key resource
type KeySpecTag struct {
}

// KeyStatus is the status for a  resource
type KeyStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// KeyOutputs is the status for a  resource
type KeyOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KeyList is a list of Key resources
type KeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Key `json:"items"`
}
