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

// Repository is a specification for a Repository resource
type Repository struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec RepositorySpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status RepositoryStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs RepositoryOutputs `json:"outputs"`
}

// RepositorySpec is the spec for the Repository resource
type RepositorySpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// LifecyclePolicy http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-lifecyclepolicy
	LifecyclePolicy RepositorySpecLifecyclePolicy `json:"lifecyclePolicy,omitempty"`

	// RepositoryName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositoryname
	RepositoryName string `json:"repositoryName,omitempty" cloudformation:"RepositoryName,Parameter"`

	// RepositoryPolicyText http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ecr-repository.html#cfn-ecr-repository-repositorypolicytext
	RepositoryPolicyText RepositorySpecRepositoryPolicyText `json:"repositoryPolicyText,omitempty" cloudformation:"RepositoryPolicyText,Parameter"`
}

// RepositorySpecLifecyclePolicy is the definition for a Repository resource
type RepositorySpecLifecyclePolicy struct {
	// LifecyclePolicyText http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-lifecyclepolicytext
	LifecyclePolicyText string `json:"lifecyclePolicyText,omitempty" cloudformation:"LifecyclePolicyText,Parameter"`

	// RegistryRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecr-repository-lifecyclepolicy.html#cfn-ecr-repository-lifecyclepolicy-registryid
	RegistryRef string `json:"registryRef,omitempty" cloudformation:"RegistryId,Parameter"`
}

// RepositorySpecRepositoryPolicyText is the definition for a Repository resource
type RepositorySpecRepositoryPolicyText struct {
}

// RepositoryStatus is the status for a  resource
type RepositoryStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// RepositoryOutputs is the status for a  resource
type RepositoryOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RepositoryList is a list of Repository resources
type RepositoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Repository `json:"items"`
}
