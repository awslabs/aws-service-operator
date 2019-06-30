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

// EventBusPolicy is a specification for a EventBusPolicy resource
type EventBusPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec EventBusPolicySpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status EventBusPolicyStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs EventBusPolicyOutputs `json:"outputs"`
}

// EventBusPolicySpec is the spec for the EventBusPolicy resource
type EventBusPolicySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Action http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-action
	Action string `json:"action,omitempty" cloudformation:"Action,Parameter"`

	// Condition http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-condition
	Condition EventBusPolicySpecCondition `json:"condition,omitempty"`

	// Principal http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-principal
	Principal string `json:"principal,omitempty" cloudformation:"Principal,Parameter"`

	// StatementRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-eventbuspolicy.html#cfn-events-eventbuspolicy-statementid
	StatementRef string `json:"statementRef,omitempty" cloudformation:"StatementId,Parameter"`
}

// EventBusPolicySpecCondition is the definition for a EventBusPolicy resource
type EventBusPolicySpecCondition struct {
	// Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-key
	Key string `json:"key,omitempty" cloudformation:"Key,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`

	// Value http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-eventbuspolicy-condition.html#cfn-events-eventbuspolicy-condition-value
	Value string `json:"value,omitempty" cloudformation:"Value,Parameter"`
}

// EventBusPolicyStatus is the status for a  resource
type EventBusPolicyStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// EventBusPolicyOutputs is the status for a  resource
type EventBusPolicyOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EventBusPolicyList is a list of EventBusPolicy resources
type EventBusPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EventBusPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Rule is a specification for a Rule resource
type Rule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec RuleSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status RuleStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs RuleOutputs `json:"outputs"`
}

// RuleSpec is the spec for the Rule resource
type RuleSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// EventPattern http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventpattern
	EventPattern RuleSpecEventPattern `json:"eventPattern,omitempty" cloudformation:"EventPattern,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-rolearn
	RoleRef string `json:"roleRef,omitempty" cloudformation:"RoleArn,Parameter"`

	// ScheduleExpression http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-scheduleexpression
	ScheduleExpression string `json:"scheduleExpression,omitempty" cloudformation:"ScheduleExpression,Parameter"`

	// State http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-state
	State string `json:"state,omitempty" cloudformation:"State,Parameter"`

	// Targets http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-targets
	Targets []RuleSpecTarget `json:"targets,omitempty" cloudformation:"Targets,Parameter"`
}

// RuleSpecEventPattern is the definition for a Rule resource
type RuleSpecEventPattern struct {
}

// RuleSpecTarget is the definition for a Rule resource
type RuleSpecTarget struct {
}

// RuleSpecTargetEcsParameters is the definition for a Rule resource
type RuleSpecTargetEcsParameters struct {
	// TaskCount http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-ecsparameters.html#cfn-events-rule-ecsparameters-taskcount
	TaskCount int32 `json:"taskCount,omitempty" cloudformation:"TaskCount,Parameter"`

	// TaskDefinitionRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-ecsparameters.html#cfn-events-rule-ecsparameters-taskdefinitionarn
	TaskDefinitionRef string `json:"taskDefinitionRef,omitempty" cloudformation:"TaskDefinitionArn,Parameter"`
}

// RuleSpecTargetInputTransformer is the definition for a Rule resource
type RuleSpecTargetInputTransformer struct {
	// InputPathsMap http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-inputtransformer.html#cfn-events-rule-inputtransformer-inputpathsmap
	InputPathsMap map[string]string `json:"inputPathsMap,omitempty" cloudformation:"InputPathsMap,Parameter"`

	// InputTemplate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-inputtransformer.html#cfn-events-rule-inputtransformer-inputtemplate
	InputTemplate string `json:"inputTemplate,omitempty" cloudformation:"InputTemplate,Parameter"`
}

// RuleSpecTargetKinesisParameters is the definition for a Rule resource
type RuleSpecTargetKinesisParameters struct {
	// PartitionKeyPath http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-kinesisparameters.html#cfn-events-rule-kinesisparameters-partitionkeypath
	PartitionKeyPath string `json:"partitionKeyPath,omitempty" cloudformation:"PartitionKeyPath,Parameter"`
}

// RuleSpecTargetRunCommandParameters is the definition for a Rule resource
type RuleSpecTargetRunCommandParameters struct {
	// RunCommandTargets http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-runcommandparameters.html#cfn-events-rule-runcommandparameters-runcommandtargets
	RunCommandTargets []RuleSpecTargetRunCommandParametersRunCommandTarget `json:"runCommandTargets,omitempty" cloudformation:"RunCommandTargets,Parameter"`
}

// RuleSpecTargetRunCommandParametersRunCommandTarget is the definition for a Rule resource
type RuleSpecTargetRunCommandParametersRunCommandTarget struct {
}

// RuleSpecTargetSqsParameters is the definition for a Rule resource
type RuleSpecTargetSqsParameters struct {
	// MessageGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-events-rule-sqsparameters.html#cfn-events-rule-sqsparameters-messagegroupid
	MessageGroupRef string `json:"messageGroupRef,omitempty" cloudformation:"MessageGroupId,Parameter"`
}

// RuleStatus is the status for a  resource
type RuleStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// RuleOutputs is the status for a  resource
type RuleOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RuleList is a list of Rule resources
type RuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Rule `json:"items"`
}
