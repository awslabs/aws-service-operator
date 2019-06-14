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

// Cluster is a specification for a Cluster resource
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ClusterSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ClusterStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ClusterOutputs `json:"outputs"`
}

// ClusterSpec is the spec for the Cluster resource
type ClusterSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// ResourcesVpcConfig http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-resourcesvpcconfig
	ResourcesVpcConfig ClusterSpecResourcesVpcConfig `json:"resourcesVpcConfig,omitempty"`

	// RoleRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-rolearn
	RoleRef string `json:"roleRef,omitempty" cloudformation:"RoleArn,Parameter"`

	// Version http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-eks-cluster.html#cfn-eks-cluster-version
	Version string `json:"version,omitempty" cloudformation:"Version,Parameter"`
}

// ClusterSpecResourcesVpcConfig is the definition for a Cluster resource
type ClusterSpecResourcesVpcConfig struct {
	// SecurityGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-securitygroupids
	SecurityGroupRef []string `json:"securityGroupRef,omitempty" cloudformation:"SecurityGroupIds,Parameter"`

	// SubnetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-eks-cluster-resourcesvpcconfig.html#cfn-eks-cluster-resourcesvpcconfig-subnetids
	SubnetRef []string `json:"subnetRef,omitempty" cloudformation:"SubnetIds,Parameter"`
}

// ClusterStatus is the status for a  resource
type ClusterStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ClusterOutputs is the status for a  resource
type ClusterOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterList is a list of Cluster resources
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Cluster `json:"items"`
}
