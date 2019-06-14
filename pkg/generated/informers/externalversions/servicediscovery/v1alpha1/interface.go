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
	// HttpNamespaces returns a HttpNamespaceInformer.
	HttpNamespaces() HttpNamespaceInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// PrivateDnsNamespaces returns a PrivateDnsNamespaceInformer.
	PrivateDnsNamespaces() PrivateDnsNamespaceInformer
	// PublicDnsNamespaces returns a PublicDnsNamespaceInformer.
	PublicDnsNamespaces() PublicDnsNamespaceInformer
	// Services returns a ServiceInformer.
	Services() ServiceInformer
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

// HttpNamespaces returns a HttpNamespaceInformer.
func (v *version) HttpNamespaces() HttpNamespaceInformer {
	return &httpNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PrivateDnsNamespaces returns a PrivateDnsNamespaceInformer.
func (v *version) PrivateDnsNamespaces() PrivateDnsNamespaceInformer {
	return &privateDnsNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PublicDnsNamespaces returns a PublicDnsNamespaceInformer.
func (v *version) PublicDnsNamespaces() PublicDnsNamespaceInformer {
	return &publicDnsNamespaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Services returns a ServiceInformer.
func (v *version) Services() ServiceInformer {
	return &serviceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
