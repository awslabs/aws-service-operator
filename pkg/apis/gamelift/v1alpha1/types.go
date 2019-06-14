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

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RoutingStrategy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-alias.html#cfn-gamelift-alias-routingstrategy
	RoutingStrategy AliasSpecRoutingStrategy `json:"routingStrategy,omitempty"`
}

// AliasSpecRoutingStrategy is the definition for a Alias resource
type AliasSpecRoutingStrategy struct {
	// FleetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-fleetid
	FleetRef string `json:"fleetRef,omitempty" cloudformation:"FleetId,Parameter"`

	// Message http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-message
	Message string `json:"message,omitempty" cloudformation:"Message,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-alias-routingstrategy.html#cfn-gamelift-alias-routingstrategy-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
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

// Build is a specification for a Build resource
type Build struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec BuildSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status BuildStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs BuildOutputs `json:"outputs"`
}

// BuildSpec is the spec for the Build resource
type BuildSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// StorageLocation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-storagelocation
	StorageLocation BuildSpecStorageLocation `json:"storageLocation,omitempty"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-build.html#cfn-gamelift-build-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// BuildSpecStorageLocation is the definition for a Build resource
type BuildSpecStorageLocation struct {
	// Bucket http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-bucket
	Bucket string `json:"bucket,omitempty" cloudformation:"Bucket,Parameter"`

	// Key http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-key
	Key string `json:"key,omitempty" cloudformation:"Key,Parameter"`

	// RoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-gamelift-build-storagelocation.html#cfn-gamelift-build-storage-rolearn
	RoleRef string `json:"roleRef,omitempty" cloudformation:"RoleArn,Parameter"`
}

// BuildStatus is the status for a  resource
type BuildStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// BuildOutputs is the status for a  resource
type BuildOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BuildList is a list of Build resources
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Build `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Fleet is a specification for a Fleet resource
type Fleet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec FleetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status FleetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs FleetOutputs `json:"outputs"`
}

// FleetSpec is the spec for the Fleet resource
type FleetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// BuildRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-buildid
	BuildRef string `json:"buildRef,omitempty" cloudformation:"BuildId,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DesiredEC2Instances http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-desiredec2instances
	DesiredEC2Instances int32 `json:"desiredEC2Instances,omitempty" cloudformation:"DesiredEC2Instances,Parameter"`

	// EC2InboundPermissions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2inboundpermissions
	EC2InboundPermissions []FleetSpecEC2InboundPermission `json:"eC2InboundPermissions,omitempty" cloudformation:"EC2InboundPermissions,Parameter"`

	// EC2InstanceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-ec2instancetype
	EC2InstanceType string `json:"eC2InstanceType,omitempty" cloudformation:"EC2InstanceType,Parameter"`

	// LogPaths http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-logpaths
	LogPaths []string `json:"logPaths,omitempty" cloudformation:"LogPaths,Parameter"`

	// MaxSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-maxsize
	MaxSize int32 `json:"maxSize,omitempty" cloudformation:"MaxSize,Parameter"`

	// MinSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-minsize
	MinSize int32 `json:"minSize,omitempty" cloudformation:"MinSize,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// ServerLaunchParameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchparameters
	ServerLaunchParameters string `json:"serverLaunchParameters,omitempty" cloudformation:"ServerLaunchParameters,Parameter"`

	// ServerLaunchPath http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-gamelift-fleet.html#cfn-gamelift-fleet-serverlaunchpath
	ServerLaunchPath string `json:"serverLaunchPath,omitempty" cloudformation:"ServerLaunchPath,Parameter"`
}

// FleetSpecEC2InboundPermission is the definition for a Fleet resource
type FleetSpecEC2InboundPermission struct {
}

// FleetStatus is the status for a  resource
type FleetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// FleetOutputs is the status for a  resource
type FleetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FleetList is a list of Fleet resources
type FleetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Fleet `json:"items"`
}
