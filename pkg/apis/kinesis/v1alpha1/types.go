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

// Stream is a specification for a Stream resource
type Stream struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec StreamSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status StreamStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs StreamOutputs `json:"outputs"`
}

// StreamSpec is the spec for the Stream resource
type StreamSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RetentionPeriodHours http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-retentionperiodhours
	RetentionPeriodHours int32 `json:"retentionPeriodHours,omitempty" cloudformation:"RetentionPeriodHours,Parameter"`

	// ShardCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-shardcount
	ShardCount int32 `json:"shardCount,omitempty" cloudformation:"ShardCount,Parameter"`

	// StreamEncryption http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-streamencryption
	StreamEncryption StreamSpecStreamEncryption `json:"streamEncryption,omitempty"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-stream.html#cfn-kinesis-stream-tags
	Tags []StreamSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// StreamSpecStreamEncryption is the definition for a Stream resource
type StreamSpecStreamEncryption struct {
	// EncryptionType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-encryptiontype
	EncryptionType string `json:"encryptionType,omitempty" cloudformation:"EncryptionType,Parameter"`

	// KeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesis-stream-streamencryption.html#cfn-kinesis-stream-streamencryption-keyid
	KeyRef string `json:"keyRef,omitempty" cloudformation:"KeyId,Parameter"`
}

// StreamSpecTag is the definition for a Stream resource
type StreamSpecTag struct {
}

// StreamStatus is the status for a  resource
type StreamStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// StreamOutputs is the status for a  resource
type StreamOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StreamList is a list of Stream resources
type StreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Stream `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StreamConsumer is a specification for a StreamConsumer resource
type StreamConsumer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec StreamConsumerSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status StreamConsumerStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs StreamConsumerOutputs `json:"outputs"`
}

// StreamConsumerSpec is the spec for the StreamConsumer resource
type StreamConsumerSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ConsumerName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-consumername
	ConsumerName string `json:"consumerName,omitempty" cloudformation:"ConsumerName,Parameter"`

	// StreamARN http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesis-streamconsumer.html#cfn-kinesis-streamconsumer-streamarn
	StreamARN string `json:"streamARN,omitempty" cloudformation:"StreamARN,Parameter"`
}

// StreamConsumerStatus is the status for a  resource
type StreamConsumerStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// StreamConsumerOutputs is the status for a  resource
type StreamConsumerOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StreamConsumerList is a list of StreamConsumer resources
type StreamConsumerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []StreamConsumer `json:"items"`
}
