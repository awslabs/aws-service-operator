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
	// CapacityReservations returns a CapacityReservationInformer.
	CapacityReservations() CapacityReservationInformer
	// CustomerGateways returns a CustomerGatewayInformer.
	CustomerGateways() CustomerGatewayInformer
	// DHCPOptionses returns a DHCPOptionsInformer.
	DHCPOptionses() DHCPOptionsInformer
	// EC2Fleets returns a EC2FleetInformer.
	EC2Fleets() EC2FleetInformer
	// EIPs returns a EIPInformer.
	EIPs() EIPInformer
	// EIPAssociations returns a EIPAssociationInformer.
	EIPAssociations() EIPAssociationInformer
	// EgressOnlyInternetGateways returns a EgressOnlyInternetGatewayInformer.
	EgressOnlyInternetGateways() EgressOnlyInternetGatewayInformer
	// FlowLogs returns a FlowLogInformer.
	FlowLogs() FlowLogInformer
	// Hosts returns a HostInformer.
	Hosts() HostInformer
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// InternetGateways returns a InternetGatewayInformer.
	InternetGateways() InternetGatewayInformer
	// LaunchTemplates returns a LaunchTemplateInformer.
	LaunchTemplates() LaunchTemplateInformer
	// NatGateways returns a NatGatewayInformer.
	NatGateways() NatGatewayInformer
	// NetworkAcls returns a NetworkAclInformer.
	NetworkAcls() NetworkAclInformer
	// NetworkAclEntries returns a NetworkAclEntryInformer.
	NetworkAclEntries() NetworkAclEntryInformer
	// NetworkInterfaces returns a NetworkInterfaceInformer.
	NetworkInterfaces() NetworkInterfaceInformer
	// NetworkInterfaceAttachments returns a NetworkInterfaceAttachmentInformer.
	NetworkInterfaceAttachments() NetworkInterfaceAttachmentInformer
	// NetworkInterfacePermissions returns a NetworkInterfacePermissionInformer.
	NetworkInterfacePermissions() NetworkInterfacePermissionInformer
	// PlacementGroups returns a PlacementGroupInformer.
	PlacementGroups() PlacementGroupInformer
	// Routes returns a RouteInformer.
	Routes() RouteInformer
	// RouteTables returns a RouteTableInformer.
	RouteTables() RouteTableInformer
	// SecurityGroups returns a SecurityGroupInformer.
	SecurityGroups() SecurityGroupInformer
	// SecurityGroupEgresses returns a SecurityGroupEgressInformer.
	SecurityGroupEgresses() SecurityGroupEgressInformer
	// SecurityGroupIngresses returns a SecurityGroupIngressInformer.
	SecurityGroupIngresses() SecurityGroupIngressInformer
	// SpotFleets returns a SpotFleetInformer.
	SpotFleets() SpotFleetInformer
	// Subnets returns a SubnetInformer.
	Subnets() SubnetInformer
	// SubnetCidrBlocks returns a SubnetCidrBlockInformer.
	SubnetCidrBlocks() SubnetCidrBlockInformer
	// SubnetNetworkAclAssociations returns a SubnetNetworkAclAssociationInformer.
	SubnetNetworkAclAssociations() SubnetNetworkAclAssociationInformer
	// SubnetRouteTableAssociations returns a SubnetRouteTableAssociationInformer.
	SubnetRouteTableAssociations() SubnetRouteTableAssociationInformer
	// TransitGateways returns a TransitGatewayInformer.
	TransitGateways() TransitGatewayInformer
	// TransitGatewayAttachments returns a TransitGatewayAttachmentInformer.
	TransitGatewayAttachments() TransitGatewayAttachmentInformer
	// TransitGatewayRoutes returns a TransitGatewayRouteInformer.
	TransitGatewayRoutes() TransitGatewayRouteInformer
	// TransitGatewayRouteTables returns a TransitGatewayRouteTableInformer.
	TransitGatewayRouteTables() TransitGatewayRouteTableInformer
	// TransitGatewayRouteTableAssociations returns a TransitGatewayRouteTableAssociationInformer.
	TransitGatewayRouteTableAssociations() TransitGatewayRouteTableAssociationInformer
	// TransitGatewayRouteTablePropagations returns a TransitGatewayRouteTablePropagationInformer.
	TransitGatewayRouteTablePropagations() TransitGatewayRouteTablePropagationInformer
	// VPCs returns a VPCInformer.
	VPCs() VPCInformer
	// VPCCidrBlocks returns a VPCCidrBlockInformer.
	VPCCidrBlocks() VPCCidrBlockInformer
	// VPCDHCPOptionsAssociations returns a VPCDHCPOptionsAssociationInformer.
	VPCDHCPOptionsAssociations() VPCDHCPOptionsAssociationInformer
	// VPCEndpoints returns a VPCEndpointInformer.
	VPCEndpoints() VPCEndpointInformer
	// VPCEndpointConnectionNotifications returns a VPCEndpointConnectionNotificationInformer.
	VPCEndpointConnectionNotifications() VPCEndpointConnectionNotificationInformer
	// VPCEndpointServices returns a VPCEndpointServiceInformer.
	VPCEndpointServices() VPCEndpointServiceInformer
	// VPCEndpointServicePermissionses returns a VPCEndpointServicePermissionsInformer.
	VPCEndpointServicePermissionses() VPCEndpointServicePermissionsInformer
	// VPCGatewayAttachments returns a VPCGatewayAttachmentInformer.
	VPCGatewayAttachments() VPCGatewayAttachmentInformer
	// VPCPeeringConnections returns a VPCPeeringConnectionInformer.
	VPCPeeringConnections() VPCPeeringConnectionInformer
	// VPNConnections returns a VPNConnectionInformer.
	VPNConnections() VPNConnectionInformer
	// VPNConnectionRoutes returns a VPNConnectionRouteInformer.
	VPNConnectionRoutes() VPNConnectionRouteInformer
	// VPNGateways returns a VPNGatewayInformer.
	VPNGateways() VPNGatewayInformer
	// VPNGatewayRoutePropagations returns a VPNGatewayRoutePropagationInformer.
	VPNGatewayRoutePropagations() VPNGatewayRoutePropagationInformer
	// Volumes returns a VolumeInformer.
	Volumes() VolumeInformer
	// VolumeAttachments returns a VolumeAttachmentInformer.
	VolumeAttachments() VolumeAttachmentInformer
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

// CapacityReservations returns a CapacityReservationInformer.
func (v *version) CapacityReservations() CapacityReservationInformer {
	return &capacityReservationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CustomerGateways returns a CustomerGatewayInformer.
func (v *version) CustomerGateways() CustomerGatewayInformer {
	return &customerGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// DHCPOptionses returns a DHCPOptionsInformer.
func (v *version) DHCPOptionses() DHCPOptionsInformer {
	return &dHCPOptionsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EC2Fleets returns a EC2FleetInformer.
func (v *version) EC2Fleets() EC2FleetInformer {
	return &eC2FleetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EIPs returns a EIPInformer.
func (v *version) EIPs() EIPInformer {
	return &eIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EIPAssociations returns a EIPAssociationInformer.
func (v *version) EIPAssociations() EIPAssociationInformer {
	return &eIPAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// EgressOnlyInternetGateways returns a EgressOnlyInternetGatewayInformer.
func (v *version) EgressOnlyInternetGateways() EgressOnlyInternetGatewayInformer {
	return &egressOnlyInternetGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// FlowLogs returns a FlowLogInformer.
func (v *version) FlowLogs() FlowLogInformer {
	return &flowLogInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Hosts returns a HostInformer.
func (v *version) Hosts() HostInformer {
	return &hostInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// InternetGateways returns a InternetGatewayInformer.
func (v *version) InternetGateways() InternetGatewayInformer {
	return &internetGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// LaunchTemplates returns a LaunchTemplateInformer.
func (v *version) LaunchTemplates() LaunchTemplateInformer {
	return &launchTemplateInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NatGateways returns a NatGatewayInformer.
func (v *version) NatGateways() NatGatewayInformer {
	return &natGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkAcls returns a NetworkAclInformer.
func (v *version) NetworkAcls() NetworkAclInformer {
	return &networkAclInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkAclEntries returns a NetworkAclEntryInformer.
func (v *version) NetworkAclEntries() NetworkAclEntryInformer {
	return &networkAclEntryInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkInterfaces returns a NetworkInterfaceInformer.
func (v *version) NetworkInterfaces() NetworkInterfaceInformer {
	return &networkInterfaceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkInterfaceAttachments returns a NetworkInterfaceAttachmentInformer.
func (v *version) NetworkInterfaceAttachments() NetworkInterfaceAttachmentInformer {
	return &networkInterfaceAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// NetworkInterfacePermissions returns a NetworkInterfacePermissionInformer.
func (v *version) NetworkInterfacePermissions() NetworkInterfacePermissionInformer {
	return &networkInterfacePermissionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PlacementGroups returns a PlacementGroupInformer.
func (v *version) PlacementGroups() PlacementGroupInformer {
	return &placementGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Routes returns a RouteInformer.
func (v *version) Routes() RouteInformer {
	return &routeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// RouteTables returns a RouteTableInformer.
func (v *version) RouteTables() RouteTableInformer {
	return &routeTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityGroups returns a SecurityGroupInformer.
func (v *version) SecurityGroups() SecurityGroupInformer {
	return &securityGroupInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityGroupEgresses returns a SecurityGroupEgressInformer.
func (v *version) SecurityGroupEgresses() SecurityGroupEgressInformer {
	return &securityGroupEgressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SecurityGroupIngresses returns a SecurityGroupIngressInformer.
func (v *version) SecurityGroupIngresses() SecurityGroupIngressInformer {
	return &securityGroupIngressInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SpotFleets returns a SpotFleetInformer.
func (v *version) SpotFleets() SpotFleetInformer {
	return &spotFleetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Subnets returns a SubnetInformer.
func (v *version) Subnets() SubnetInformer {
	return &subnetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubnetCidrBlocks returns a SubnetCidrBlockInformer.
func (v *version) SubnetCidrBlocks() SubnetCidrBlockInformer {
	return &subnetCidrBlockInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubnetNetworkAclAssociations returns a SubnetNetworkAclAssociationInformer.
func (v *version) SubnetNetworkAclAssociations() SubnetNetworkAclAssociationInformer {
	return &subnetNetworkAclAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// SubnetRouteTableAssociations returns a SubnetRouteTableAssociationInformer.
func (v *version) SubnetRouteTableAssociations() SubnetRouteTableAssociationInformer {
	return &subnetRouteTableAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitGateways returns a TransitGatewayInformer.
func (v *version) TransitGateways() TransitGatewayInformer {
	return &transitGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitGatewayAttachments returns a TransitGatewayAttachmentInformer.
func (v *version) TransitGatewayAttachments() TransitGatewayAttachmentInformer {
	return &transitGatewayAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitGatewayRoutes returns a TransitGatewayRouteInformer.
func (v *version) TransitGatewayRoutes() TransitGatewayRouteInformer {
	return &transitGatewayRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitGatewayRouteTables returns a TransitGatewayRouteTableInformer.
func (v *version) TransitGatewayRouteTables() TransitGatewayRouteTableInformer {
	return &transitGatewayRouteTableInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitGatewayRouteTableAssociations returns a TransitGatewayRouteTableAssociationInformer.
func (v *version) TransitGatewayRouteTableAssociations() TransitGatewayRouteTableAssociationInformer {
	return &transitGatewayRouteTableAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// TransitGatewayRouteTablePropagations returns a TransitGatewayRouteTablePropagationInformer.
func (v *version) TransitGatewayRouteTablePropagations() TransitGatewayRouteTablePropagationInformer {
	return &transitGatewayRouteTablePropagationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCs returns a VPCInformer.
func (v *version) VPCs() VPCInformer {
	return &vPCInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCCidrBlocks returns a VPCCidrBlockInformer.
func (v *version) VPCCidrBlocks() VPCCidrBlockInformer {
	return &vPCCidrBlockInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCDHCPOptionsAssociations returns a VPCDHCPOptionsAssociationInformer.
func (v *version) VPCDHCPOptionsAssociations() VPCDHCPOptionsAssociationInformer {
	return &vPCDHCPOptionsAssociationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCEndpoints returns a VPCEndpointInformer.
func (v *version) VPCEndpoints() VPCEndpointInformer {
	return &vPCEndpointInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCEndpointConnectionNotifications returns a VPCEndpointConnectionNotificationInformer.
func (v *version) VPCEndpointConnectionNotifications() VPCEndpointConnectionNotificationInformer {
	return &vPCEndpointConnectionNotificationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCEndpointServices returns a VPCEndpointServiceInformer.
func (v *version) VPCEndpointServices() VPCEndpointServiceInformer {
	return &vPCEndpointServiceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCEndpointServicePermissionses returns a VPCEndpointServicePermissionsInformer.
func (v *version) VPCEndpointServicePermissionses() VPCEndpointServicePermissionsInformer {
	return &vPCEndpointServicePermissionsInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCGatewayAttachments returns a VPCGatewayAttachmentInformer.
func (v *version) VPCGatewayAttachments() VPCGatewayAttachmentInformer {
	return &vPCGatewayAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPCPeeringConnections returns a VPCPeeringConnectionInformer.
func (v *version) VPCPeeringConnections() VPCPeeringConnectionInformer {
	return &vPCPeeringConnectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPNConnections returns a VPNConnectionInformer.
func (v *version) VPNConnections() VPNConnectionInformer {
	return &vPNConnectionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPNConnectionRoutes returns a VPNConnectionRouteInformer.
func (v *version) VPNConnectionRoutes() VPNConnectionRouteInformer {
	return &vPNConnectionRouteInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPNGateways returns a VPNGatewayInformer.
func (v *version) VPNGateways() VPNGatewayInformer {
	return &vPNGatewayInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VPNGatewayRoutePropagations returns a VPNGatewayRoutePropagationInformer.
func (v *version) VPNGatewayRoutePropagations() VPNGatewayRoutePropagationInformer {
	return &vPNGatewayRoutePropagationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Volumes returns a VolumeInformer.
func (v *version) Volumes() VolumeInformer {
	return &volumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// VolumeAttachments returns a VolumeAttachmentInformer.
func (v *version) VolumeAttachments() VolumeAttachmentInformer {
	return &volumeAttachmentInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
