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
// Code generated by controller-gen. DO NOT EDIT.

package controllers

import (
	"fmt"

	awsclientset "awsoperator.io/pkg/generated/clientset/versioned"
	cloudformationInformers "awsoperator.io/pkg/generated/informers/externalversions/cloudformation/v1alpha1"
	"github.com/aws/aws-sdk-go/service/cloudformation/cloudformationiface"
	"k8s.io/client-go/kubernetes"
)

// Controller allows dynamica controllers to be triggered
type Controller interface {
	Run(int, <-chan struct{}) error
}

// Get will return an a controller which responds to Run()
func Get(
	name string,
	kubeclientset kubernetes.Interface,
	awsclientset awsclientset.Interface,
	stackInformer cloudformationInformers.StackInformer,
	clients map[string]cloudformationiface.CloudFormationAPI) (Controller, error) {
	// switch name {
	// case "cloudformation":
	// 	return cloudformation.New(kubeclientset, awsclientset, stackInformer, clients)
	// }
	return nil, fmt.Errorf("No controller found for '%s'", name)
}
