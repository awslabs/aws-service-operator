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
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// InstanceFleetConfigs returns a InstanceFleetConfigInformer.
	InstanceFleetConfigs() InstanceFleetConfigInformer
	// InstanceGroupConfigs returns a InstanceGroupConfigInformer.
	InstanceGroupConfigs() InstanceGroupConfigInformer
	// SecurityConfigurations returns a SecurityConfigurationInformer.
	SecurityConfigurations() SecurityConfigurationInformer
	// Steps returns a StepInformer.
	Steps() StepInformer
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

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceFleetConfigs returns a InstanceFleetConfigInformer.
func (v *version) InstanceFleetConfigs() InstanceFleetConfigInformer {
	return &instanceFleetConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InstanceGroupConfigs returns a InstanceGroupConfigInformer.
func (v *version) InstanceGroupConfigs() InstanceGroupConfigInformer {
	return &instanceGroupConfigInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityConfigurations returns a SecurityConfigurationInformer.
func (v *version) SecurityConfigurations() SecurityConfigurationInformer {
	return &securityConfigurationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Steps returns a StepInformer.
func (v *version) Steps() StepInformer {
	return &stepInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
