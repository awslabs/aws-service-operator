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
// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "awsoperator.io/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// CacheClusters returns a CacheClusterInformer.
	CacheClusters() CacheClusterInformer
	// ParameterGroups returns a ParameterGroupInformer.
	ParameterGroups() ParameterGroupInformer
	// ReplicationGroups returns a ReplicationGroupInformer.
	ReplicationGroups() ReplicationGroupInformer
	// SecurityGroups returns a SecurityGroupInformer.
	SecurityGroups() SecurityGroupInformer
	// SecurityGroupIngresses returns a SecurityGroupIngressInformer.
	SecurityGroupIngresses() SecurityGroupIngressInformer
	// SubnetGroups returns a SubnetGroupInformer.
	SubnetGroups() SubnetGroupInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// CacheClusters returns a CacheClusterInformer.
func (v *version) CacheClusters() CacheClusterInformer {
	return &cacheClusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ParameterGroups returns a ParameterGroupInformer.
func (v *version) ParameterGroups() ParameterGroupInformer {
	return &parameterGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ReplicationGroups returns a ReplicationGroupInformer.
func (v *version) ReplicationGroups() ReplicationGroupInformer {
	return &replicationGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityGroups returns a SecurityGroupInformer.
func (v *version) SecurityGroups() SecurityGroupInformer {
	return &securityGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityGroupIngresses returns a SecurityGroupIngressInformer.
func (v *version) SecurityGroupIngresses() SecurityGroupIngressInformer {
	return &securityGroupIngressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubnetGroups returns a SubnetGroupInformer.
func (v *version) SubnetGroups() SubnetGroupInformer {
	return &subnetGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
