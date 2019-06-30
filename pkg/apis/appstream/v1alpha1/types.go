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

// DirectoryConfig is a specification for a DirectoryConfig resource
type DirectoryConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DirectoryConfigSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DirectoryConfigStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DirectoryConfigOutputs `json:"outputs"`
}

// DirectoryConfigSpec is the spec for the DirectoryConfig resource
type DirectoryConfigSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DirectoryName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-directoryname
	DirectoryName string `json:"directoryName,omitempty" cloudformation:"DirectoryName,Parameter"`

	// OrganizationalUnitDistinguishedNames http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-organizationalunitdistinguishednames
	OrganizationalUnitDistinguishedNames []string `json:"organizationalUnitDistinguishedNames,omitempty" cloudformation:"OrganizationalUnitDistinguishedNames,Parameter"`

	// ServiceAccountCredentials http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html#cfn-appstream-directoryconfig-serviceaccountcredentials
	ServiceAccountCredentials DirectoryConfigSpecServiceAccountCredentials `json:"serviceAccountCredentials,omitempty"`
}

// DirectoryConfigSpecServiceAccountCredentials is the definition for a DirectoryConfig resource
type DirectoryConfigSpecServiceAccountCredentials struct {
	// AccountName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountname
	AccountName string `json:"accountName,omitempty" cloudformation:"AccountName,Parameter"`

	// AccountPassword http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountpassword
	AccountPassword string `json:"accountPassword,omitempty" cloudformation:"AccountPassword,Parameter"`
}

// DirectoryConfigStatus is the status for a  resource
type DirectoryConfigStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DirectoryConfigOutputs is the status for a  resource
type DirectoryConfigOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DirectoryConfigList is a list of DirectoryConfig resources
type DirectoryConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DirectoryConfig `json:"items"`
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

	// ComputeCapacity http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-computecapacity
	ComputeCapacity FleetSpecComputeCapacity `json:"computeCapacity,omitempty"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DisconnectTimeoutInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-disconnecttimeoutinseconds
	DisconnectTimeoutInSeconds int32 `json:"disconnectTimeoutInSeconds,omitempty" cloudformation:"DisconnectTimeoutInSeconds,Parameter"`

	// DisplayName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-displayname
	DisplayName string `json:"displayName,omitempty" cloudformation:"DisplayName,Parameter"`

	// DomainJoinInfo http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-domainjoininfo
	DomainJoinInfo FleetSpecDomainJoinInfo `json:"domainJoinInfo,omitempty"`

	// EnableDefaultInternetAccess http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-enabledefaultinternetaccess
	EnableDefaultInternetAccess bool `json:"enableDefaultInternetAccess,omitempty" cloudformation:"EnableDefaultInternetAccess,Parameter"`

	// FleetType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-fleettype
	FleetType string `json:"fleetType,omitempty" cloudformation:"FleetType,Parameter"`

	// ImageRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-imagearn
	ImageRef string `json:"imageRef,omitempty" cloudformation:"ImageArn,Parameter"`

	// ImageName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-imagename
	ImageName string `json:"imageName,omitempty" cloudformation:"ImageName,Parameter"`

	// InstanceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-instancetype
	InstanceType string `json:"instanceType,omitempty" cloudformation:"InstanceType,Parameter"`

	// MaxUserDurationInSeconds http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-maxuserdurationinseconds
	MaxUserDurationInSeconds int32 `json:"maxUserDurationInSeconds,omitempty" cloudformation:"MaxUserDurationInSeconds,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-tags
	Tags []FleetSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// VpcConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-fleet.html#cfn-appstream-fleet-vpcconfig
	VpcConfig FleetSpecVpcConfig `json:"vpcConfig,omitempty"`
}

// FleetSpecComputeCapacity is the definition for a Fleet resource
type FleetSpecComputeCapacity struct {
	// DesiredInstances http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-computecapacity.html#cfn-appstream-fleet-computecapacity-desiredinstances
	DesiredInstances int32 `json:"desiredInstances,omitempty" cloudformation:"DesiredInstances,Parameter"`
}

// FleetSpecDomainJoinInfo is the definition for a Fleet resource
type FleetSpecDomainJoinInfo struct {
	// DirectoryName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-domainjoininfo.html#cfn-appstream-fleet-domainjoininfo-directoryname
	DirectoryName string `json:"directoryName,omitempty" cloudformation:"DirectoryName,Parameter"`

	// OrganizationalUnitDistinguishedName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-domainjoininfo.html#cfn-appstream-fleet-domainjoininfo-organizationalunitdistinguishedname
	OrganizationalUnitDistinguishedName string `json:"organizationalUnitDistinguishedName,omitempty" cloudformation:"OrganizationalUnitDistinguishedName,Parameter"`
}

// FleetSpecTag is the definition for a Fleet resource
type FleetSpecTag struct {
}

// FleetSpecVpcConfig is the definition for a Fleet resource
type FleetSpecVpcConfig struct {
	// SecurityGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-vpcconfig.html#cfn-appstream-fleet-vpcconfig-securitygroupids
	SecurityGroupRef []string `json:"securityGroupRef,omitempty" cloudformation:"SecurityGroupIds,Parameter"`

	// SubnetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-fleet-vpcconfig.html#cfn-appstream-fleet-vpcconfig-subnetids
	SubnetRef []string `json:"subnetRef,omitempty" cloudformation:"SubnetIds,Parameter"`
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImageBuilder is a specification for a ImageBuilder resource
type ImageBuilder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ImageBuilderSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ImageBuilderStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ImageBuilderOutputs `json:"outputs"`
}

// ImageBuilderSpec is the spec for the ImageBuilder resource
type ImageBuilderSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AppstreamAgentVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-appstreamagentversion
	AppstreamAgentVersion string `json:"appstreamAgentVersion,omitempty" cloudformation:"AppstreamAgentVersion,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DisplayName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-displayname
	DisplayName string `json:"displayName,omitempty" cloudformation:"DisplayName,Parameter"`

	// DomainJoinInfo http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-domainjoininfo
	DomainJoinInfo ImageBuilderSpecDomainJoinInfo `json:"domainJoinInfo,omitempty"`

	// EnableDefaultInternetAccess http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-enabledefaultinternetaccess
	EnableDefaultInternetAccess bool `json:"enableDefaultInternetAccess,omitempty" cloudformation:"EnableDefaultInternetAccess,Parameter"`

	// ImageRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagearn
	ImageRef string `json:"imageRef,omitempty" cloudformation:"ImageArn,Parameter"`

	// ImageName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-imagename
	ImageName string `json:"imageName,omitempty" cloudformation:"ImageName,Parameter"`

	// InstanceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-instancetype
	InstanceType string `json:"instanceType,omitempty" cloudformation:"InstanceType,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-tags
	Tags []ImageBuilderSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// VpcConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-imagebuilder.html#cfn-appstream-imagebuilder-vpcconfig
	VpcConfig ImageBuilderSpecVpcConfig `json:"vpcConfig,omitempty"`
}

// ImageBuilderSpecDomainJoinInfo is the definition for a ImageBuilder resource
type ImageBuilderSpecDomainJoinInfo struct {
	// DirectoryName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-domainjoininfo.html#cfn-appstream-imagebuilder-domainjoininfo-directoryname
	DirectoryName string `json:"directoryName,omitempty" cloudformation:"DirectoryName,Parameter"`

	// OrganizationalUnitDistinguishedName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-domainjoininfo.html#cfn-appstream-imagebuilder-domainjoininfo-organizationalunitdistinguishedname
	OrganizationalUnitDistinguishedName string `json:"organizationalUnitDistinguishedName,omitempty" cloudformation:"OrganizationalUnitDistinguishedName,Parameter"`
}

// ImageBuilderSpecTag is the definition for a ImageBuilder resource
type ImageBuilderSpecTag struct {
}

// ImageBuilderSpecVpcConfig is the definition for a ImageBuilder resource
type ImageBuilderSpecVpcConfig struct {
	// SecurityGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-vpcconfig.html#cfn-appstream-imagebuilder-vpcconfig-securitygroupids
	SecurityGroupRef []string `json:"securityGroupRef,omitempty" cloudformation:"SecurityGroupIds,Parameter"`

	// SubnetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-imagebuilder-vpcconfig.html#cfn-appstream-imagebuilder-vpcconfig-subnetids
	SubnetRef []string `json:"subnetRef,omitempty" cloudformation:"SubnetIds,Parameter"`
}

// ImageBuilderStatus is the status for a  resource
type ImageBuilderStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ImageBuilderOutputs is the status for a  resource
type ImageBuilderOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ImageBuilderList is a list of ImageBuilder resources
type ImageBuilderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ImageBuilder `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Stack is a specification for a Stack resource
type Stack struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec StackSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status StackStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs StackOutputs `json:"outputs"`
}

// StackSpec is the spec for the Stack resource
type StackSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ApplicationSettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-applicationsettings
	ApplicationSettings StackSpecApplicationSettings `json:"applicationSettings,omitempty"`

	// AttributesToDelete http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-attributestodelete
	AttributesToDelete []string `json:"attributesToDelete,omitempty" cloudformation:"AttributesToDelete,Parameter"`

	// DeleteStorageConnectors http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-deletestorageconnectors
	DeleteStorageConnectors bool `json:"deleteStorageConnectors,omitempty" cloudformation:"DeleteStorageConnectors,Parameter"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// DisplayName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-displayname
	DisplayName string `json:"displayName,omitempty" cloudformation:"DisplayName,Parameter"`

	// FeedbackURL http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-feedbackurl
	FeedbackURL string `json:"feedbackURL,omitempty" cloudformation:"FeedbackURL,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// RedirectURL http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-redirecturl
	RedirectURL string `json:"redirectURL,omitempty" cloudformation:"RedirectURL,Parameter"`

	// StorageConnectors http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-storageconnectors
	StorageConnectors []StackSpecStorageConnector `json:"storageConnectors,omitempty" cloudformation:"StorageConnectors,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-tags
	Tags []StackSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// UserSettings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stack.html#cfn-appstream-stack-usersettings
	UserSettings []StackSpecUserSetting `json:"userSettings,omitempty" cloudformation:"UserSettings,Parameter"`
}

// StackSpecApplicationSettings is the definition for a Stack resource
type StackSpecApplicationSettings struct {
	// Enabled http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html#cfn-appstream-stack-applicationsettings-enabled
	Enabled bool `json:"enabled,omitempty" cloudformation:"Enabled,Parameter"`

	// SettingsGroup http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html#cfn-appstream-stack-applicationsettings-settingsgroup
	SettingsGroup string `json:"settingsGroup,omitempty" cloudformation:"SettingsGroup,Parameter"`
}

// StackSpecStorageConnector is the definition for a Stack resource
type StackSpecStorageConnector struct {
}

// StackSpecTag is the definition for a Stack resource
type StackSpecTag struct {
}

// StackSpecUserSetting is the definition for a Stack resource
type StackSpecUserSetting struct {
}

// StackStatus is the status for a  resource
type StackStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// StackOutputs is the status for a  resource
type StackOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackList is a list of Stack resources
type StackList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Stack `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackFleetAssociation is a specification for a StackFleetAssociation resource
type StackFleetAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec StackFleetAssociationSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status StackFleetAssociationStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs StackFleetAssociationOutputs `json:"outputs"`
}

// StackFleetAssociationSpec is the spec for the StackFleetAssociation resource
type StackFleetAssociationSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// FleetName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-fleetname
	FleetName string `json:"fleetName,omitempty" cloudformation:"FleetName,Parameter"`

	// StackName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackfleetassociation.html#cfn-appstream-stackfleetassociation-stackname
	StackName string `json:"stackName,omitempty" cloudformation:"StackName,Parameter"`
}

// StackFleetAssociationStatus is the status for a  resource
type StackFleetAssociationStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// StackFleetAssociationOutputs is the status for a  resource
type StackFleetAssociationOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackFleetAssociationList is a list of StackFleetAssociation resources
type StackFleetAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []StackFleetAssociation `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackUserAssociation is a specification for a StackUserAssociation resource
type StackUserAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec StackUserAssociationSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status StackUserAssociationStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs StackUserAssociationOutputs `json:"outputs"`
}

// StackUserAssociationSpec is the spec for the StackUserAssociation resource
type StackUserAssociationSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AuthenticationType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-authenticationtype
	AuthenticationType string `json:"authenticationType,omitempty" cloudformation:"AuthenticationType,Parameter"`

	// SendEmailNotification http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-sendemailnotification
	SendEmailNotification bool `json:"sendEmailNotification,omitempty" cloudformation:"SendEmailNotification,Parameter"`

	// StackName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-stackname
	StackName string `json:"stackName,omitempty" cloudformation:"StackName,Parameter"`

	// UserName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-stackuserassociation.html#cfn-appstream-stackuserassociation-username
	UserName string `json:"userName,omitempty" cloudformation:"UserName,Parameter"`
}

// StackUserAssociationStatus is the status for a  resource
type StackUserAssociationStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// StackUserAssociationOutputs is the status for a  resource
type StackUserAssociationOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// StackUserAssociationList is a list of StackUserAssociation resources
type StackUserAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []StackUserAssociation `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// User is a specification for a User resource
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec UserSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status UserStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs UserOutputs `json:"outputs"`
}

// UserSpec is the spec for the User resource
type UserSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AuthenticationType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-authenticationtype
	AuthenticationType string `json:"authenticationType,omitempty" cloudformation:"AuthenticationType,Parameter"`

	// FirstName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-firstname
	FirstName string `json:"firstName,omitempty" cloudformation:"FirstName,Parameter"`

	// LastName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-lastname
	LastName string `json:"lastName,omitempty" cloudformation:"LastName,Parameter"`

	// MessageAction http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-messageaction
	MessageAction string `json:"messageAction,omitempty" cloudformation:"MessageAction,Parameter"`

	// UserName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-user.html#cfn-appstream-user-username
	UserName string `json:"userName,omitempty" cloudformation:"UserName,Parameter"`
}

// UserStatus is the status for a  resource
type UserStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// UserOutputs is the status for a  resource
type UserOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserList is a list of User resources
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []User `json:"items"`
}
