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

// ByteMatchSet is a specification for a ByteMatchSet resource
type ByteMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec ByteMatchSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status ByteMatchSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs ByteMatchSetOutputs `json:"outputs"`
}

// ByteMatchSetSpec is the spec for the ByteMatchSet resource
type ByteMatchSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// ByteMatchTuples http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html#cfn-waf-bytematchset-bytematchtuples
	ByteMatchTuples []ByteMatchSetSpecByteMatchTuple `json:"byteMatchTuples,omitempty" cloudformation:"ByteMatchTuples,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-bytematchset.html#cfn-waf-bytematchset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`
}

// ByteMatchSetSpecByteMatchTuple is the definition for a ByteMatchSet resource
type ByteMatchSetSpecByteMatchTuple struct {
}

// ByteMatchSetSpecByteMatchTupleFieldToMatch is the definition for a ByteMatchSet resource
type ByteMatchSetSpecByteMatchTupleFieldToMatch struct {
	// Data http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch-data
	Data string `json:"data,omitempty" cloudformation:"Data,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-bytematchset-bytematchtuples-fieldtomatch-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// ByteMatchSetStatus is the status for a  resource
type ByteMatchSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// ByteMatchSetOutputs is the status for a  resource
type ByteMatchSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ByteMatchSetList is a list of ByteMatchSet resources
type ByteMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ByteMatchSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IPSet is a specification for a IPSet resource
type IPSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec IPSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status IPSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs IPSetOutputs `json:"outputs"`
}

// IPSetSpec is the spec for the IPSet resource
type IPSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// IPSetDescriptors http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-ipsetdescriptors
	IPSetDescriptors []IPSetSpecIPSetDescriptor `json:"iPSetDescriptors,omitempty" cloudformation:"IPSetDescriptors,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-ipset.html#cfn-waf-ipset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`
}

// IPSetSpecIPSetDescriptor is the definition for a IPSet resource
type IPSetSpecIPSetDescriptor struct {
}

// IPSetStatus is the status for a  resource
type IPSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// IPSetOutputs is the status for a  resource
type IPSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IPSetList is a list of IPSet resources
type IPSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []IPSet `json:"items"`
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

	// MetricName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-metricname
	MetricName string `json:"metricName,omitempty" cloudformation:"MetricName,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Predicates http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-rule.html#cfn-waf-rule-predicates
	Predicates []RuleSpecPredicate `json:"predicates,omitempty" cloudformation:"Predicates,Parameter"`
}

// RuleSpecPredicate is the definition for a Rule resource
type RuleSpecPredicate struct {
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

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SizeConstraintSet is a specification for a SizeConstraintSet resource
type SizeConstraintSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec SizeConstraintSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status SizeConstraintSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs SizeConstraintSetOutputs `json:"outputs"`
}

// SizeConstraintSetSpec is the spec for the SizeConstraintSet resource
type SizeConstraintSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// SizeConstraints http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sizeconstraintset.html#cfn-waf-sizeconstraintset-sizeconstraints
	SizeConstraints []SizeConstraintSetSpecSizeConstraint `json:"sizeConstraints,omitempty" cloudformation:"SizeConstraints,Parameter"`
}

// SizeConstraintSetSpecSizeConstraint is the definition for a SizeConstraintSet resource
type SizeConstraintSetSpecSizeConstraint struct {
}

// SizeConstraintSetSpecSizeConstraintFieldToMatch is the definition for a SizeConstraintSet resource
type SizeConstraintSetSpecSizeConstraintFieldToMatch struct {
	// Data http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data
	Data string `json:"data,omitempty" cloudformation:"Data,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-sizeconstraintset-sizeconstraint-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// SizeConstraintSetStatus is the status for a  resource
type SizeConstraintSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// SizeConstraintSetOutputs is the status for a  resource
type SizeConstraintSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SizeConstraintSetList is a list of SizeConstraintSet resources
type SizeConstraintSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SizeConstraintSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SqlInjectionMatchSet is a specification for a SqlInjectionMatchSet resource
type SqlInjectionMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec SqlInjectionMatchSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status SqlInjectionMatchSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs SqlInjectionMatchSetOutputs `json:"outputs"`
}

// SqlInjectionMatchSetSpec is the spec for the SqlInjectionMatchSet resource
type SqlInjectionMatchSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// SqlInjectionMatchTuples http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-sqlinjectionmatchset.html#cfn-waf-sqlinjectionmatchset-sqlinjectionmatchtuples
	SqlInjectionMatchTuples []SqlInjectionMatchSetSpecSqlInjectionMatchTuple `json:"sqlInjectionMatchTuples,omitempty" cloudformation:"SqlInjectionMatchTuples,Parameter"`
}

// SqlInjectionMatchSetSpecSqlInjectionMatchTuple is the definition for a SqlInjectionMatchSet resource
type SqlInjectionMatchSetSpecSqlInjectionMatchTuple struct {
}

// SqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch is the definition for a SqlInjectionMatchSet resource
type SqlInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch struct {
	// Data http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-data
	Data string `json:"data,omitempty" cloudformation:"Data,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-bytematchset-bytematchtuples-fieldtomatch.html#cfn-waf-sizeconstraintset-sizeconstraint-fieldtomatch-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// SqlInjectionMatchSetStatus is the status for a  resource
type SqlInjectionMatchSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// SqlInjectionMatchSetOutputs is the status for a  resource
type SqlInjectionMatchSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SqlInjectionMatchSetList is a list of SqlInjectionMatchSet resources
type SqlInjectionMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SqlInjectionMatchSet `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebACL is a specification for a WebACL resource
type WebACL struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec WebACLSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status WebACLStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs WebACLOutputs `json:"outputs"`
}

// WebACLSpec is the spec for the WebACL resource
type WebACLSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DefaultAction http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-defaultaction
	DefaultAction WebACLSpecDefaultAction `json:"defaultAction,omitempty"`

	// MetricName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-metricname
	MetricName string `json:"metricName,omitempty" cloudformation:"MetricName,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Rules http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-rules
	Rules []WebACLSpecRule `json:"rules,omitempty" cloudformation:"Rules,Parameter"`
}

// WebACLSpecDefaultAction is the definition for a WebACL resource
type WebACLSpecDefaultAction struct {
	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html#cfn-waf-webacl-action-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// WebACLSpecRule is the definition for a WebACL resource
type WebACLSpecRule struct {
}

// WebACLSpecRuleAction is the definition for a WebACL resource
type WebACLSpecRuleAction struct {
	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html#cfn-waf-webacl-action-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// WebACLStatus is the status for a  resource
type WebACLStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// WebACLOutputs is the status for a  resource
type WebACLOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WebACLList is a list of WebACL resources
type WebACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []WebACL `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// XssMatchSet is a specification for a XssMatchSet resource
type XssMatchSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec XssMatchSetSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status XssMatchSetStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs XssMatchSetOutputs `json:"outputs"`
}

// XssMatchSetSpec is the spec for the XssMatchSet resource
type XssMatchSetSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// XssMatchTuples http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-xssmatchset.html#cfn-waf-xssmatchset-xssmatchtuples
	XssMatchTuples []XssMatchSetSpecXssMatchTuple `json:"xssMatchTuples,omitempty" cloudformation:"XssMatchTuples,Parameter"`
}

// XssMatchSetSpecXssMatchTuple is the definition for a XssMatchSet resource
type XssMatchSetSpecXssMatchTuple struct {
}

// XssMatchSetSpecXssMatchTupleFieldToMatch is the definition for a XssMatchSet resource
type XssMatchSetSpecXssMatchTupleFieldToMatch struct {
	// Data http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-data
	Data string `json:"data,omitempty" cloudformation:"Data,Parameter"`

	// Type http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-xssmatchset-xssmatchtuple-fieldtomatch.html#cfn-waf-xssmatchset-xssmatchtuple-fieldtomatch-type
	Type string `json:"type,omitempty" cloudformation:"Type,Parameter"`
}

// XssMatchSetStatus is the status for a  resource
type XssMatchSetStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// XssMatchSetOutputs is the status for a  resource
type XssMatchSetOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// XssMatchSetList is a list of XssMatchSet resources
type XssMatchSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []XssMatchSet `json:"items"`
}
