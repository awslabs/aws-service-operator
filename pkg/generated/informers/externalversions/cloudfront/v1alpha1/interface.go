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
	// CloudFrontOriginAccessIdentities returns a CloudFrontOriginAccessIdentityInformer.
	CloudFrontOriginAccessIdentities() CloudFrontOriginAccessIdentityInformer
	// Distributions returns a DistributionInformer.
	Distributions() DistributionInformer
	// StreamingDistributions returns a StreamingDistributionInformer.
	StreamingDistributions() StreamingDistributionInformer
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

// CloudFrontOriginAccessIdentities returns a CloudFrontOriginAccessIdentityInformer.
func (v *version) CloudFrontOriginAccessIdentities() CloudFrontOriginAccessIdentityInformer {
	return &cloudFrontOriginAccessIdentityInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Distributions returns a DistributionInformer.
func (v *version) Distributions() DistributionInformer {
	return &distributionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StreamingDistributions returns a StreamingDistributionInformer.
func (v *version) StreamingDistributions() StreamingDistributionInformer {
	return &streamingDistributionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
