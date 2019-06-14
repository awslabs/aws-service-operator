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

// DBCluster is a specification for a DBCluster resource
type DBCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DBClusterSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DBClusterStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DBClusterOutputs `json:"outputs"`
}

// DBClusterSpec is the spec for the DBCluster resource
type DBClusterSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AvailabilityZones http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-availabilityzones
	AvailabilityZones []string `json:"availabilityZones,omitempty" cloudformation:"AvailabilityZones,Parameter"`

	// BackupRetentionPeriod http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-backupretentionperiod
	BackupRetentionPeriod int32 `json:"backupRetentionPeriod,omitempty" cloudformation:"BackupRetentionPeriod,Parameter"`

	// DBClusterentifierRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbclusteridentifier
	DBClusterentifierRef string `json:"dBClusterentifierRef,omitempty" cloudformation:"DBClusterIdentifier,Parameter"`

	// DBClusterParameterGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbclusterparametergroupname
	DBClusterParameterGroupName string `json:"dBClusterParameterGroupName,omitempty" cloudformation:"DBClusterParameterGroupName,Parameter"`

	// DBSubnetGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbsubnetgroupname
	DBSubnetGroupName string `json:"dBSubnetGroupName,omitempty" cloudformation:"DBSubnetGroupName,Parameter"`

	// EngineVersion http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-engineversion
	EngineVersion string `json:"engineVersion,omitempty" cloudformation:"EngineVersion,Parameter"`

	// KmsKeyRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-kmskeyid
	KmsKeyRef string `json:"kmsKeyRef,omitempty" cloudformation:"KmsKeyId,Parameter"`

	// MasterUserPassword http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masteruserpassword
	MasterUserPassword string `json:"masterUserPassword,omitempty" cloudformation:"MasterUserPassword,Parameter"`

	// MasterUsername http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masterusername
	MasterUsername string `json:"masterUsername,omitempty" cloudformation:"MasterUsername,Parameter"`

	// Port http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-port
	Port int32 `json:"port,omitempty" cloudformation:"Port,Parameter"`

	// PreferredBackupWindow http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-preferredbackupwindow
	PreferredBackupWindow string `json:"preferredBackupWindow,omitempty" cloudformation:"PreferredBackupWindow,Parameter"`

	// PreferredMaintenanceWindow http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" cloudformation:"PreferredMaintenanceWindow,Parameter"`

	// SnapshotentifierRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-snapshotidentifier
	SnapshotentifierRef string `json:"snapshotentifierRef,omitempty" cloudformation:"SnapshotIdentifier,Parameter"`

	// StorageEncrypted http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-storageencrypted
	StorageEncrypted bool `json:"storageEncrypted,omitempty" cloudformation:"StorageEncrypted,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-tags
	Tags []DBClusterSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`

	// VpcSecurityGroupRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-vpcsecuritygroupids
	VpcSecurityGroupRef []string `json:"vpcSecurityGroupRef,omitempty" cloudformation:"VpcSecurityGroupIds,Parameter"`
}

// DBClusterSpecTag is the definition for a DBCluster resource
type DBClusterSpecTag struct {
}

// DBClusterStatus is the status for a  resource
type DBClusterStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DBClusterOutputs is the status for a  resource
type DBClusterOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBClusterList is a list of DBCluster resources
type DBClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DBCluster `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBClusterParameterGroup is a specification for a DBClusterParameterGroup resource
type DBClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DBClusterParameterGroupSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DBClusterParameterGroupStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DBClusterParameterGroupOutputs `json:"outputs"`
}

// DBClusterParameterGroupSpec is the spec for the DBClusterParameterGroup resource
type DBClusterParameterGroupSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// Description http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html#cfn-docdb-dbclusterparametergroup-description
	Description string `json:"description,omitempty" cloudformation:"Description,Parameter"`

	// Family http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html#cfn-docdb-dbclusterparametergroup-family
	Family string `json:"family,omitempty" cloudformation:"Family,Parameter"`

	// Name http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html#cfn-docdb-dbclusterparametergroup-name
	Name string `json:"name,omitempty" cloudformation:"Name,Parameter"`

	// Parameters http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html#cfn-docdb-dbclusterparametergroup-parameters
	Parameters DBClusterParameterGroupSpecParameters `json:"parameters,omitempty" cloudformation:"Parameters,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbclusterparametergroup.html#cfn-docdb-dbclusterparametergroup-tags
	Tags []DBClusterParameterGroupSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// DBClusterParameterGroupSpecParameters is the definition for a DBClusterParameterGroup resource
type DBClusterParameterGroupSpecParameters struct {
}

// DBClusterParameterGroupSpecTag is the definition for a DBClusterParameterGroup resource
type DBClusterParameterGroupSpecTag struct {
}

// DBClusterParameterGroupStatus is the status for a  resource
type DBClusterParameterGroupStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DBClusterParameterGroupOutputs is the status for a  resource
type DBClusterParameterGroupOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBClusterParameterGroupList is a list of DBClusterParameterGroup resources
type DBClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DBClusterParameterGroup `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBInstance is a specification for a DBInstance resource
type DBInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DBInstanceSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DBInstanceStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DBInstanceOutputs `json:"outputs"`
}

// DBInstanceSpec is the spec for the DBInstance resource
type DBInstanceSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// AutoMinorVersionUpgrade http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-autominorversionupgrade
	AutoMinorVersionUpgrade bool `json:"autoMinorVersionUpgrade,omitempty" cloudformation:"AutoMinorVersionUpgrade,Parameter"`

	// AvailabilityZone http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-availabilityzone
	AvailabilityZone string `json:"availabilityZone,omitempty" cloudformation:"AvailabilityZone,Parameter"`

	// DBClusterentifierRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-dbclusteridentifier
	DBClusterentifierRef string `json:"dBClusterentifierRef,omitempty" cloudformation:"DBClusterIdentifier,Parameter"`

	// DBInstanceClass http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-dbinstanceclass
	DBInstanceClass string `json:"dBInstanceClass,omitempty" cloudformation:"DBInstanceClass,Parameter"`

	// DBInstanceentifierRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-dbinstanceidentifier
	DBInstanceentifierRef string `json:"dBInstanceentifierRef,omitempty" cloudformation:"DBInstanceIdentifier,Parameter"`

	// PreferredMaintenanceWindow http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"preferredMaintenanceWindow,omitempty" cloudformation:"PreferredMaintenanceWindow,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbinstance.html#cfn-docdb-dbinstance-tags
	Tags []DBInstanceSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// DBInstanceSpecTag is the definition for a DBInstance resource
type DBInstanceSpecTag struct {
}

// DBInstanceStatus is the status for a  resource
type DBInstanceStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DBInstanceOutputs is the status for a  resource
type DBInstanceOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBInstanceList is a list of DBInstance resources
type DBInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DBInstance `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBSubnetGroup is a specification for a DBSubnetGroup resource
type DBSubnetGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" cloudformation:"Name,Parameter"`

	// Spec stores all the CRD information
	Spec DBSubnetGroupSpec `json:"spec"`

	// +optional
	// Status stores the CloudFormation Status
	Status DBSubnetGroupStatus `json:"status"`

	// +optional
	// Outputs store any Cloudformation outputs
	Outputs DBSubnetGroupOutputs `json:"outputs"`
}

// DBSubnetGroupSpec is the spec for the DBSubnetGroup resource
type DBSubnetGroupSpec struct {
	metav1alpha1.CloudFormationMeta `json:",inline"`

	// DBSubnetGroupDescription http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbsubnetgroup.html#cfn-docdb-dbsubnetgroup-dbsubnetgroupdescription
	DBSubnetGroupDescription string `json:"dBSubnetGroupDescription,omitempty" cloudformation:"DBSubnetGroupDescription,Parameter"`

	// DBSubnetGroupName http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbsubnetgroup.html#cfn-docdb-dbsubnetgroup-dbsubnetgroupname
	DBSubnetGroupName string `json:"dBSubnetGroupName,omitempty" cloudformation:"DBSubnetGroupName,Parameter"`

	// SubnetRef http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbsubnetgroup.html#cfn-docdb-dbsubnetgroup-subnetids
	SubnetRef []string `json:"subnetRef,omitempty" cloudformation:"SubnetIds,Parameter"`

	// Tags http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbsubnetgroup.html#cfn-docdb-dbsubnetgroup-tags
	Tags []DBSubnetGroupSpecTag `json:"tags,omitempty" cloudformation:"Tags,Parameter"`
}

// DBSubnetGroupSpecTag is the definition for a DBSubnetGroup resource
type DBSubnetGroupSpecTag struct {
}

// DBSubnetGroupStatus is the status for a  resource
type DBSubnetGroupStatus struct {
	metav1alpha1.StatusMeta `json:",inline" `
}

// DBSubnetGroupOutputs is the status for a  resource
type DBSubnetGroupOutputs struct {
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DBSubnetGroupList is a list of DBSubnetGroup resources
type DBSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []DBSubnetGroup `json:"items"`
}
