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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "awsoperator.io/pkg/generated/clientset/versioned/typed/batch/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeBatchV1alpha1 struct {
	*testing.Fake
}

func (c *FakeBatchV1alpha1) ComputeEnvironments(namespace string) v1alpha1.ComputeEnvironmentInterface {
	return &FakeComputeEnvironments{c, namespace}
}

func (c *FakeBatchV1alpha1) JobDefinitions(namespace string) v1alpha1.JobDefinitionInterface {
	return &FakeJobDefinitions{c, namespace}
}

func (c *FakeBatchV1alpha1) JobQueues(namespace string) v1alpha1.JobQueueInterface {
	return &FakeJobQueues{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeBatchV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
