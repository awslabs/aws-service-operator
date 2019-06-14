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

// ResourceShare is a specification for a ResourceShare resource
type ResourceShare struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ResourceShareSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ResourceShareStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ResourceShareOutputs `json:"outputs"`
}

// ResourceShareSpec is the spec for the ResourceShare resource
type ResourceShareSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AllowExternalPrincipals http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-allowexternalprincipals
	AllowExternalPrincipals bool `json:"allowExternalPrincipals,omitempty" cloudformation:"AllowExternalPrincipals,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Principals http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-principals
	Principals []string `json:"principals,omitempty" cloudformation:"Principals,Parameter"`

	// ResourceRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-resourcearns
	ResourceRef []string `json:"resourceRef,omitempty" cloudformation:"ResourceArns,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ram-resourceshare.html#cfn-ram-resourceshare-tags
	Tags []ResourceShareSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// ResourceShareSpecTag is the definition for a ResourceShare resource
type ResourceShareSpecTag struct {
}

// ResourceShareStatus is the status for a  resource
type ResourceShareStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ResourceShareOutputs is the status for a  resource
type ResourceShareOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceShareList is a list of ResourceShare resources
type ResourceShareList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ResourceShare `json:"items"`
}
