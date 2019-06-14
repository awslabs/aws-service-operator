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

// AutoScalingGroup is a specification for a AutoScalingGroup resource
type AutoScalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec AutoScalingGroupSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status AutoScalingGroupStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs AutoScalingGroupOutputs `json:"outputs"`
}

// AutoScalingGroupSpec is the spec for the AutoScalingGroup resource
type AutoScalingGroupSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AutoScalingGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-autoscaling-autoscalinggroup-autoscalinggroupname
	AutoScalingGroupName string `json:"autoScalingGroupName,omitempty" cloudformation:"AutoScalingGroupName,Parameter"`

	// AvailabilityZones http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-availabilityzones
	AvailabilityZones []string `json:"availabilityZones,omitempty" cloudformation:"AvailabilityZones,Parameter"`

	// Cooldown http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-cooldown
	Cooldown string `json:"cooldown,omitempty" cloudformation:"Cooldown,Parameter"`

	// DesiredCapacity http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	DesiredCapacity string `json:"desiredCapacity,omitempty" cloudformation:"DesiredCapacity,Parameter"`

	// HealthCheckGracePeriod http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthcheckgraceperiod
	HealthCheckGracePeriod int32 `json:"healthCheckGracePeriod,omitempty" cloudformation:"HealthCheckGracePeriod,Parameter"`

	// HealthCheckType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthchecktype
	HealthCheckType string `json:"healthCheckType,omitempty" cloudformation:"HealthCheckType,Parameter"`

	// InstanceRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-instanceid
	InstanceRef string `json:"instanceRef,omitempty" cloudformation:"InstanceId,Parameter"`

	// LaunchConfigurationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-launchconfigurationname
	LaunchConfigurationName string `json:"launchConfigurationName,omitempty" cloudformation:"LaunchConfigurationName,Parameter"`

	// LaunchTemplate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-launchtemplate
	LaunchTemplate AutoScalingGroupSpecLaunchTemplate `json:"launchTemplate,omitempty"`

	// LifecycleHookSpecificationList http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-autoscaling-autoscalinggroup-lifecyclehookspecificationlist
	LifecycleHookSpecificationList []AutoScalingGroupSpecLifecycleHookSpecificationList `json:"lifecycleHookSpecificationList,omitempty" cloudformation:"LifecycleHookSpecificationList,Parameter"`

	// LoadBalancerNames http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-loadbalancernames
	LoadBalancerNames []string `json:"loadBalancerNames,omitempty" cloudformation:"LoadBalancerNames,Parameter"`

	// MaxSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-maxsize
	MaxSize string `json:"maxSize,omitempty" cloudformation:"MaxSize,Parameter"`

	// MetricsCollection http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-metricscollection
	MetricsCollection []AutoScalingGroupSpecMetricsCollection `json:"metricsCollection,omitempty" cloudformation:"MetricsCollection,Parameter"`

	// MinSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-minsize
	MinSize string `json:"minSize,omitempty" cloudformation:"MinSize,Parameter"`

	// MixedInstancesPolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-mixedinstancespolicy
	MixedInstancesPolicy AutoScalingGroupSpecMixedInstancesPolicy `json:"mixedInstancesPolicy,omitempty"`

	// NotificationConfigurations http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	NotificationConfigurations []AutoScalingGroupSpecNotificationConfiguration `json:"notificationConfigurations,omitempty" cloudformation:"NotificationConfigurations,Parameter"`

	// PlacementGroup http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-placementgroup
	PlacementGroup string `json:"placementGroup,omitempty" cloudformation:"PlacementGroup,Parameter"`

	// ServiceLinkedRoleARN http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-autoscaling-autoscalinggroup-servicelinkedrolearn
	ServiceLinkedRoleARN string `json:"serviceLinkedRoleARN,omitempty" cloudformation:"ServiceLinkedRoleARN,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-tags
	Tags []AutoScalingGroupSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// TargetGroupARNs http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-targetgrouparns
	TargetGroupARNs []string `json:"targetGroupARNs,omitempty" cloudformation:"TargetGroupARNs,Parameter"`

	// TerminationPolicies http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-termpolicy
	TerminationPolicies []string `json:"terminationPolicies,omitempty" cloudformation:"TerminationPolicies,Parameter"`

	// VPCZoneentifierRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-vpczoneidentifier
	VPCZoneentifierRef []string `json:"vPCZoneentifierRef,omitempty" cloudformation:"VPCZoneIdentifier,Parameter"`
}

// AutoScalingGroupSpecLaunchTemplate is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecLaunchTemplate struct {
	// LaunchTemplateRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html#cfn-autoscaling-autoscalinggroup-launchtemplatespecification-launchtemplateid
	LaunchTemplateRef string `json:"launchTemplateRef,omitempty" cloudformation:"LaunchTemplateId,Parameter"`

	// LaunchTemplateName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html#cfn-autoscaling-autoscalinggroup-launchtemplatespecification-launchtemplatename
	LaunchTemplateName string `json:"launchTemplateName,omitempty" cloudformation:"LaunchTemplateName,Parameter"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html#cfn-autoscaling-autoscalinggroup-launchtemplatespecification-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// AutoScalingGroupSpecLifecycleHookSpecificationList is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecLifecycleHookSpecificationList struct {
}

// AutoScalingGroupSpecMetricsCollection is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecMetricsCollection struct {
}

// AutoScalingGroupSpecMixedInstancesPolicy is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecMixedInstancesPolicy struct {
	// InstancesDistribution http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-group-mixedinstancespolicy.html#cfn-as-mixedinstancespolicy-instancesdistribution
	InstancesDistribution AutoScalingGroupSpecMixedInstancesPolicyInstancesDistribution `json:"instancesDistribution,omitempty"`

	// LaunchTemplate http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-group-mixedinstancespolicy.html#cfn-as-mixedinstancespolicy-launchtemplate
	LaunchTemplate AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplate `json:"launchTemplate,omitempty"`
}

// AutoScalingGroupSpecMixedInstancesPolicyInstancesDistribution is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecMixedInstancesPolicyInstancesDistribution struct {
	// OnDemandAllocationStrategy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-ondemandallocationstrategy
	OnDemandAllocationStrategy string `json:"onDemandAllocationStrategy,omitempty" cloudformation:"OnDemandAllocationStrategy,Parameter"`

	// OnDemandBaseCapacity http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-ondemandbasecapacity
	OnDemandBaseCapacity int32 `json:"onDemandBaseCapacity,omitempty" cloudformation:"OnDemandBaseCapacity,Parameter"`

	// OnDemandPercentageAboveBaseCapacity http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-ondemandpercentageabovebasecapacity
	OnDemandPercentageAboveBaseCapacity int32 `json:"onDemandPercentageAboveBaseCapacity,omitempty" cloudformation:"OnDemandPercentageAboveBaseCapacity,Parameter"`

	// SpotAllocationStrategy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-spotallocationstrategy
	SpotAllocationStrategy string `json:"spotAllocationStrategy,omitempty" cloudformation:"SpotAllocationStrategy,Parameter"`

	// SpotInstancePools http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-spotinstancepools
	SpotInstancePools int32 `json:"spotInstancePools,omitempty" cloudformation:"SpotInstancePools,Parameter"`

	// SpotMaxPrice http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-instancesdistribution.html#cfn-autoscaling-autoscalinggroup-instancesdistribution-spotmaxprice
	SpotMaxPrice string `json:"spotMaxPrice,omitempty" cloudformation:"SpotMaxPrice,Parameter"`
}

// AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplate is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplate struct {
	// LaunchTemplateSpecification http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-launchtemplate.html#cfn-as-group-launchtemplate
	LaunchTemplateSpecification AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification `json:"launchTemplateSpecification,omitempty"`

	// Overrides http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-as-mixedinstancespolicy-launchtemplate.html#cfn-as-mixedinstancespolicy-overrides
	Overrides []AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride `json:"overrides,omitempty" cloudformation:"Overrides,Parameter"`
}

// AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplateLaunchTemplateSpecification struct {
	// LaunchTemplateRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html#cfn-autoscaling-autoscalinggroup-launchtemplatespecification-launchtemplateid
	LaunchTemplateRef string `json:"launchTemplateRef,omitempty" cloudformation:"LaunchTemplateId,Parameter"`

	// LaunchTemplateName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html#cfn-autoscaling-autoscalinggroup-launchtemplatespecification-launchtemplatename
	LaunchTemplateName string `json:"launchTemplateName,omitempty" cloudformation:"LaunchTemplateName,Parameter"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-autoscalinggroup-launchtemplatespecification.html#cfn-autoscaling-autoscalinggroup-launchtemplatespecification-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecMixedInstancesPolicyLaunchTemplateOverride struct {
}

// AutoScalingGroupSpecNotificationConfiguration is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecNotificationConfiguration struct {
}

// AutoScalingGroupSpecTag is the definition for a AutoScalingGroup resource
type AutoScalingGroupSpecTag struct {
}

// AutoScalingGroupStatus is the status for a  resource
type AutoScalingGroupStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// AutoScalingGroupOutputs is the status for a  resource
type AutoScalingGroupOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AutoScalingGroupList is a list of AutoScalingGroup resources
type AutoScalingGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []AutoScalingGroup `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LaunchConfiguration is a specification for a LaunchConfiguration resource
type LaunchConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec LaunchConfigurationSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status LaunchConfigurationStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs LaunchConfigurationOutputs `json:"outputs"`
}

// LaunchConfigurationSpec is the spec for the LaunchConfiguration resource
type LaunchConfigurationSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AssociatePublicIpAddress http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cf-as-launchconfig-associatepubip
	AssociatePublicIpAddress bool `json:"associatePublicIpAddress,omitempty" cloudformation:"AssociatePublicIpAddress,Parameter"`

	// BlockDeviceMappings http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-blockdevicemappings
	BlockDeviceMappings []LaunchConfigurationSpecBlockDeviceMapping `json:"blockDeviceMappings,omitempty" cloudformation:"BlockDeviceMappings,Parameter"`

	// ClassicLinkVPCRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcid
	ClassicLinkVPCRef string `json:"classicLinkVPCRef,omitempty" cloudformation:"ClassicLinkVPCId,Parameter"`

	// ClassicLinkVPCSecurityGroups http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-classiclinkvpcsecuritygroups
	ClassicLinkVPCSecurityGroups []string `json:"classicLinkVPCSecurityGroups,omitempty" cloudformation:"ClassicLinkVPCSecurityGroups,Parameter"`

	// EbsOptimized http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ebsoptimized
	EbsOptimized bool `json:"ebsOptimized,omitempty" cloudformation:"EbsOptimized,Parameter"`

	// IamInstanceProfile http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-iaminstanceprofile
	IamInstanceProfile string `json:"iamInstanceProfile,omitempty" cloudformation:"IamInstanceProfile,Parameter"`

	// ImageRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-imageid
	ImageRef string `json:"imageRef,omitempty" cloudformation:"ImageId,Parameter"`

	// InstanceRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instanceid
	InstanceRef string `json:"instanceRef,omitempty" cloudformation:"InstanceId,Parameter"`

	// InstanceMonitoring http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancemonitoring
	InstanceMonitoring bool `json:"instanceMonitoring,omitempty" cloudformation:"InstanceMonitoring,Parameter"`

	// InstanceType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-instancetype
	InstanceType string `json:"instanceType,omitempty" cloudformation:"InstanceType,Parameter"`

	// KernelRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-kernelid
	KernelRef string `json:"kernelRef,omitempty" cloudformation:"KernelId,Parameter"`

	// KeyName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-keyname
	KeyName string `json:"keyName,omitempty" cloudformation:"KeyName,Parameter"`

	// LaunchConfigurationName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-autoscaling-launchconfig-launchconfigurationname
	LaunchConfigurationName string `json:"launchConfigurationName,omitempty" cloudformation:"LaunchConfigurationName,Parameter"`

	// PlacementTenancy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-placementtenancy
	PlacementTenancy string `json:"placementTenancy,omitempty" cloudformation:"PlacementTenancy,Parameter"`

	// RamDiskRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-ramdiskid
	RamDiskRef string `json:"ramDiskRef,omitempty" cloudformation:"RamDiskId,Parameter"`

	// SecurityGroupsRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-securitygroups
	SecurityGroupsRef []string `json:"securityGroupsRef,omitempty" cloudformation:"SecurityGroups,Parameter"`

	// SpotPrice http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-spotprice
	SpotPrice string `json:"spotPrice,omitempty" cloudformation:"SpotPrice,Parameter"`

	// UserData http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig.html#cfn-as-launchconfig-userdata
	UserData string `json:"userData,omitempty" cloudformation:"UserData,Parameter"`
}

// LaunchConfigurationSpecBlockDeviceMapping is the definition for a LaunchConfiguration resource
type LaunchConfigurationSpecBlockDeviceMapping struct {
}

// LaunchConfigurationSpecBlockDeviceMappingEbs is the definition for a LaunchConfiguration resource
type LaunchConfigurationSpecBlockDeviceMappingEbs struct {
	// DeleteOnTermination http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-deleteonterm
	DeleteOnTermination bool `json:"deleteOnTermination,omitempty" cloudformation:"DeleteOnTermination,Parameter"`

	// Encrypted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-encrypted
	Encrypted bool `json:"encrypted,omitempty" cloudformation:"Encrypted,Parameter"`

	// Iops http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-iops
	Iops int32 `json:"iops,omitempty" cloudformation:"Iops,Parameter"`

	// SnapshotRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-snapshotid
	SnapshotRef string `json:"snapshotRef,omitempty" cloudformation:"SnapshotId,Parameter"`

	// VolumeSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-volumesize
	VolumeSize int32 `json:"volumeSize,omitempty" cloudformation:"VolumeSize,Parameter"`

	// VolumeType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-launchconfig-blockdev-template.html#cfn-as-launchconfig-blockdev-template-volumetype
	VolumeType string `json:"volumeType,omitempty" cloudformation:"VolumeType,Parameter"`
}

// LaunchConfigurationStatus is the status for a  resource
type LaunchConfigurationStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// LaunchConfigurationOutputs is the status for a  resource
type LaunchConfigurationOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LaunchConfigurationList is a list of LaunchConfiguration resources
type LaunchConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LaunchConfiguration `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LifecycleHook is a specification for a LifecycleHook resource
type LifecycleHook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec LifecycleHookSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status LifecycleHookStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs LifecycleHookOutputs `json:"outputs"`
}

// LifecycleHookSpec is the spec for the LifecycleHook resource
type LifecycleHookSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AutoScalingGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-autoscalinggroupname
	AutoScalingGroupName string `json:"autoScalingGroupName,omitempty" cloudformation:"AutoScalingGroupName,Parameter"`

	// DefaultResult http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-defaultresult
	DefaultResult string `json:"defaultResult,omitempty" cloudformation:"DefaultResult,Parameter"`

	// HeartbeatTimeout http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-heartbeattimeout
	HeartbeatTimeout int32 `json:"heartbeatTimeout,omitempty" cloudformation:"HeartbeatTimeout,Parameter"`

	// LifecycleHookName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-autoscaling-lifecyclehook-lifecyclehookname
	LifecycleHookName string `json:"lifecycleHookName,omitempty" cloudformation:"LifecycleHookName,Parameter"`

	// LifecycleTransition http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-lifecycletransition
	LifecycleTransition string `json:"lifecycleTransition,omitempty" cloudformation:"LifecycleTransition,Parameter"`

	// NotificationMetadata http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-notificationmetadata
	NotificationMetadata string `json:"notificationMetadata,omitempty" cloudformation:"NotificationMetadata,Parameter"`

	// NotificationTargetARN http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-notificationtargetarn
	NotificationTargetARN string `json:"notificationTargetARN,omitempty" cloudformation:"NotificationTargetARN,Parameter"`

	// RoleARN http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-lifecyclehook.html#cfn-as-lifecyclehook-rolearn
	RoleARN string `json:"roleARN,omitempty" cloudformation:"RoleARN,Parameter"`
}

// LifecycleHookStatus is the status for a  resource
type LifecycleHookStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// LifecycleHookOutputs is the status for a  resource
type LifecycleHookOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// LifecycleHookList is a list of LifecycleHook resources
type LifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []LifecycleHook `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScalingPolicy is a specification for a ScalingPolicy resource
type ScalingPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ScalingPolicySpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ScalingPolicyStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ScalingPolicyOutputs `json:"outputs"`
}

// ScalingPolicySpec is the spec for the ScalingPolicy resource
type ScalingPolicySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AdjustmentType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-adjustmenttype
	AdjustmentType string `json:"adjustmentType,omitempty" cloudformation:"AdjustmentType,Parameter"`

	// AutoScalingGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-autoscalinggroupname
	AutoScalingGroupName string `json:"autoScalingGroupName,omitempty" cloudformation:"AutoScalingGroupName,Parameter"`

	// Cooldown http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-cooldown
	Cooldown string `json:"cooldown,omitempty" cloudformation:"Cooldown,Parameter"`

	// EstimatedInstanceWarmup http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-estimatedinstancewarmup
	EstimatedInstanceWarmup int32 `json:"estimatedInstanceWarmup,omitempty" cloudformation:"EstimatedInstanceWarmup,Parameter"`

	// MetricAggregationType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-metricaggregationtype
	MetricAggregationType string `json:"metricAggregationType,omitempty" cloudformation:"MetricAggregationType,Parameter"`

	// MinAdjustmentMagnitude http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-minadjustmentmagnitude
	MinAdjustmentMagnitude int32 `json:"minAdjustmentMagnitude,omitempty" cloudformation:"MinAdjustmentMagnitude,Parameter"`

	// PolicyType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-policytype
	PolicyType string `json:"policyType,omitempty" cloudformation:"PolicyType,Parameter"`

	// ScalingAdjustment http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-scalingadjustment
	ScalingAdjustment int32 `json:"scalingAdjustment,omitempty" cloudformation:"ScalingAdjustment,Parameter"`

	// StepAdjustments http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-as-scalingpolicy-stepadjustments
	StepAdjustments []ScalingPolicySpecStepAdjustment `json:"stepAdjustments,omitempty" cloudformation:"StepAdjustments,Parameter"`

	// TargetTrackingConfiguration http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-policy.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration
	TargetTrackingConfiguration ScalingPolicySpecTargetTrackingConfiguration `json:"targetTrackingConfiguration,omitempty"`
}

// ScalingPolicySpecStepAdjustment is the definition for a ScalingPolicy resource
type ScalingPolicySpecStepAdjustment struct {
}

// ScalingPolicySpecTargetTrackingConfiguration is the definition for a ScalingPolicy resource
type ScalingPolicySpecTargetTrackingConfiguration struct {
	// CustomizedMetricSpecification http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-customizedmetricspecification
	CustomizedMetricSpecification ScalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification `json:"customizedMetricSpecification,omitempty"`

	// DisableScaleIn http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-disablescalein
	DisableScaleIn bool `json:"disableScaleIn,omitempty" cloudformation:"DisableScaleIn,Parameter"`

	// PredefinedMetricSpecification http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-predefinedmetricspecification
	PredefinedMetricSpecification ScalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification `json:"predefinedMetricSpecification,omitempty"`

	// TargetValue http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-targettrackingconfiguration.html#cfn-autoscaling-scalingpolicy-targettrackingconfiguration-targetvalue
	TargetValue float64 `json:"targetValue,omitempty" cloudformation:"TargetValue,Parameter"`
}

// ScalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification is the definition for a ScalingPolicy resource
type ScalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecification struct {
	// Dimensions http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-dimensions
	Dimensions []ScalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationDimension `json:"dimensions,omitempty" cloudformation:"Dimensions,Parameter"`

	// MetricName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-metricname
	MetricName string `json:"metricName,omitempty" cloudformation:"MetricName,Parameter"`

	// Namespace http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-namespace
	Namespace string `json:"namespace,omitempty" cloudformation:"Namespace,Parameter"`

	// Statistic http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-statistic
	Statistic string `json:"statistic,omitempty" cloudformation:"Statistic,Parameter"`

	// Unit http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-customizedmetricspecification.html#cfn-autoscaling-scalingpolicy-customizedmetricspecification-unit
	Unit string `json:"unit,omitempty" cloudformation:"Unit,Parameter"`
}

// ScalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationDimension is the definition for a ScalingPolicy resource
type ScalingPolicySpecTargetTrackingConfigurationCustomizedMetricSpecificationDimension struct {
}

// ScalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification is the definition for a ScalingPolicy resource
type ScalingPolicySpecTargetTrackingConfigurationPredefinedMetricSpecification struct {
	// PredefinedMetricType http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-autoscaling-scalingpolicy-predefinedmetricspecification-predefinedmetrictype
	PredefinedMetricType string `json:"predefinedMetricType,omitempty" cloudformation:"PredefinedMetricType,Parameter"`

	// ResourceLabel http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-autoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-autoscaling-scalingpolicy-predefinedmetricspecification-resourcelabel
	ResourceLabel string `json:"resourceLabel,omitempty" cloudformation:"ResourceLabel,Parameter"`
}

// ScalingPolicyStatus is the status for a  resource
type ScalingPolicyStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ScalingPolicyOutputs is the status for a  resource
type ScalingPolicyOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScalingPolicyList is a list of ScalingPolicy resources
type ScalingPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ScalingPolicy `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScheduledAction is a specification for a ScheduledAction resource
type ScheduledAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ScheduledActionSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ScheduledActionStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ScheduledActionOutputs `json:"outputs"`
}

// ScheduledActionSpec is the spec for the ScheduledAction resource
type ScheduledActionSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AutoScalingGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-asgname
	AutoScalingGroupName string `json:"autoScalingGroupName,omitempty" cloudformation:"AutoScalingGroupName,Parameter"`

	// DesiredCapacity http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-desiredcapacity
	DesiredCapacity int32 `json:"desiredCapacity,omitempty" cloudformation:"DesiredCapacity,Parameter"`

	// EndTime http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-endtime
	EndTime string `json:"endTime,omitempty" cloudformation:"EndTime,Parameter"`

	// MaxSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-maxsize
	MaxSize int32 `json:"maxSize,omitempty" cloudformation:"MaxSize,Parameter"`

	// MinSize http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-minsize
	MinSize int32 `json:"minSize,omitempty" cloudformation:"MinSize,Parameter"`

	// Recurrence http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-recurrence
	Recurrence string `json:"recurrence,omitempty" cloudformation:"Recurrence,Parameter"`

	// StartTime http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-as-scheduledaction.html#cfn-as-scheduledaction-starttime
	StartTime string `json:"startTime,omitempty" cloudformation:"StartTime,Parameter"`
}

// ScheduledActionStatus is the status for a  resource
type ScheduledActionStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ScheduledActionOutputs is the status for a  resource
type ScheduledActionOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ScheduledActionList is a list of ScheduledAction resources
type ScheduledActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ScheduledAction `json:"items"`
}
