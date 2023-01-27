//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpostgresqlhsc

// ClustersClientCheckNameAvailabilityResponse contains the response from method ClustersClient.CheckNameAvailability.
type ClustersClientCheckNameAvailabilityResponse struct {
	NameAvailability
}

// ClustersClientCreateOrUpdateResponse contains the response from method ClustersClient.BeginCreateOrUpdate.
type ClustersClientCreateOrUpdateResponse struct {
	ClusterResponse
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	ClusterResponse
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	ClusterListResult
}

// ClustersClientListResponse contains the response from method ClustersClient.NewListPager.
type ClustersClientListResponse struct {
	ClusterListResult
}

// ClustersClientPromoteReadReplicaResponse contains the response from method ClustersClient.BeginPromoteReadReplica.
type ClustersClientPromoteReadReplicaResponse struct {
	// placeholder for future response values
}

// ClustersClientRestartResponse contains the response from method ClustersClient.BeginRestart.
type ClustersClientRestartResponse struct {
	// placeholder for future response values
}

// ClustersClientStartResponse contains the response from method ClustersClient.BeginStart.
type ClustersClientStartResponse struct {
	// placeholder for future response values
}

// ClustersClientStopResponse contains the response from method ClustersClient.BeginStop.
type ClustersClientStopResponse struct {
	// placeholder for future response values
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.BeginUpdate.
type ClustersClientUpdateResponse struct {
	ClusterResponse
}

// ConfigurationsClientCreateOrUpdateCoordinatorResponse contains the response from method ConfigurationsClient.BeginCreateOrUpdateCoordinator.
type ConfigurationsClientCreateOrUpdateCoordinatorResponse struct {
	ServerConfiguration
}

// ConfigurationsClientCreateOrUpdateNodeResponse contains the response from method ConfigurationsClient.BeginCreateOrUpdateNode.
type ConfigurationsClientCreateOrUpdateNodeResponse struct {
	ServerConfiguration
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	Configuration
}

// ConfigurationsClientListByClusterResponse contains the response from method ConfigurationsClient.NewListByClusterPager.
type ConfigurationsClientListByClusterResponse struct {
	ClusterConfigurationListResult
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.NewListByServerPager.
type ConfigurationsClientListByServerResponse struct {
	ServerConfigurationListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.BeginCreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.BeginDelete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	FirewallRule
}

// FirewallRulesClientListByClusterResponse contains the response from method FirewallRulesClient.NewListByClusterPager.
type FirewallRulesClientListByClusterResponse struct {
	FirewallRuleListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByClusterResponse contains the response from method PrivateEndpointConnectionsClient.NewListByClusterPager.
type PrivateEndpointConnectionsClientListByClusterResponse struct {
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByClusterResponse contains the response from method PrivateLinkResourcesClient.NewListByClusterPager.
type PrivateLinkResourcesClientListByClusterResponse struct {
	PrivateLinkResourceListResult
}

// RolesClientCreateResponse contains the response from method RolesClient.BeginCreate.
type RolesClientCreateResponse struct {
	Role
}

// RolesClientDeleteResponse contains the response from method RolesClient.BeginDelete.
type RolesClientDeleteResponse struct {
	// placeholder for future response values
}

// RolesClientListByClusterResponse contains the response from method RolesClient.NewListByClusterPager.
type RolesClientListByClusterResponse struct {
	RoleListResult
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	ClusterServer
}

// ServersClientListByClusterResponse contains the response from method ServersClient.NewListByClusterPager.
type ServersClientListByClusterResponse struct {
	ClusterServerListResult
}
