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

// Domain is a specification for a Domain resource
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DomainSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DomainStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DomainOutputs `json:"outputs"`
}

// DomainSpec is the spec for the Domain resource
type DomainSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-simpledb.html#cfn-sdb-domain-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`
}

// DomainStatus is the status for a  resource
type DomainStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DomainOutputs is the status for a  resource
type DomainOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DomainList is a list of Domain resources
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Domain `json:"items"`
}
