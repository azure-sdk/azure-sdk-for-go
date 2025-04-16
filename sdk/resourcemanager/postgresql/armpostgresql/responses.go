// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
}

// ConfigurationsClientCreateOrUpdateResponse contains the response from method ConfigurationsClient.BeginCreateOrUpdate.
type ConfigurationsClientCreateOrUpdateResponse struct {
	// Represents a Configuration.
	Configuration
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	// Represents a Configuration.
	Configuration
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.NewListByServerPager.
type ConfigurationsClientListByServerResponse struct {
	// A list of server configurations.
	ConfigurationListResult
}

// DatabasesClientCreateOrUpdateResponse contains the response from method DatabasesClient.BeginCreateOrUpdate.
type DatabasesClientCreateOrUpdateResponse struct {
	// Represents a Database.
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.BeginDelete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	// Represents a Database.
	Database
}

// DatabasesClientListByServerResponse contains the response from method DatabasesClient.NewListByServerPager.
type DatabasesClientListByServerResponse struct {
	// A List of databases.
	DatabaseListResult
}

// FirewallRulesClientCreateOrUpdateResponse contains the response from method FirewallRulesClient.BeginCreateOrUpdate.
type FirewallRulesClientCreateOrUpdateResponse struct {
	// Represents a server firewall rule.
	FirewallRule
}

// FirewallRulesClientDeleteResponse contains the response from method FirewallRulesClient.BeginDelete.
type FirewallRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// FirewallRulesClientGetResponse contains the response from method FirewallRulesClient.Get.
type FirewallRulesClientGetResponse struct {
	// Represents a server firewall rule.
	FirewallRule
}

// FirewallRulesClientListByServerResponse contains the response from method FirewallRulesClient.NewListByServerPager.
type FirewallRulesClientListByServerResponse struct {
	// A list of firewall rules.
	FirewallRuleListResult
}

// LocationBasedPerformanceTierClientListResponse contains the response from method LocationBasedPerformanceTierClient.NewListPager.
type LocationBasedPerformanceTierClientListResponse struct {
	// A list of performance tiers.
	PerformanceTierListResult
}

// LogFilesClientListByServerResponse contains the response from method LogFilesClient.NewListByServerPager.
type LogFilesClientListByServerResponse struct {
	// A list of log files.
	LogFileListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of resource provider operations.
	OperationListResult
}

// PrivateEndpointConnectionsClientCreateOrUpdateResponse contains the response from method PrivateEndpointConnectionsClient.BeginCreateOrUpdate.
type PrivateEndpointConnectionsClientCreateOrUpdateResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.BeginDelete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByServerResponse contains the response from method PrivateEndpointConnectionsClient.NewListByServerPager.
type PrivateEndpointConnectionsClientListByServerResponse struct {
	// A list of private endpoint connections.
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateTagsResponse contains the response from method PrivateEndpointConnectionsClient.BeginUpdateTags.
type PrivateEndpointConnectionsClientUpdateTagsResponse struct {
	// A private endpoint connection
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByServerResponse contains the response from method PrivateLinkResourcesClient.NewListByServerPager.
type PrivateLinkResourcesClientListByServerResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// RecoverableServersClientGetResponse contains the response from method RecoverableServersClient.Get.
type RecoverableServersClientGetResponse struct {
	// A recoverable server resource.
	RecoverableServerResource
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.NewListByServerPager.
type ReplicasClientListByServerResponse struct {
	// A list of servers.
	ServerListResult
}

// ServerAdministratorsClientCreateOrUpdateResponse contains the response from method ServerAdministratorsClient.BeginCreateOrUpdate.
type ServerAdministratorsClientCreateOrUpdateResponse struct {
	// Represents a and external administrator to be created.
	ServerAdministratorResource
}

// ServerAdministratorsClientDeleteResponse contains the response from method ServerAdministratorsClient.BeginDelete.
type ServerAdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerAdministratorsClientGetResponse contains the response from method ServerAdministratorsClient.Get.
type ServerAdministratorsClientGetResponse struct {
	// Represents a and external administrator to be created.
	ServerAdministratorResource
}

// ServerAdministratorsClientListResponse contains the response from method ServerAdministratorsClient.NewListPager.
type ServerAdministratorsClientListResponse struct {
	// The response to a list Active Directory Administrators request.
	ServerAdministratorResourceListResult
}

// ServerBasedPerformanceTierClientListResponse contains the response from method ServerBasedPerformanceTierClient.NewListPager.
type ServerBasedPerformanceTierClientListResponse struct {
	// A list of performance tiers.
	PerformanceTierListResult
}

// ServerKeysClientCreateOrUpdateResponse contains the response from method ServerKeysClient.BeginCreateOrUpdate.
type ServerKeysClientCreateOrUpdateResponse struct {
	// A PostgreSQL Server key.
	ServerKey
}

// ServerKeysClientDeleteResponse contains the response from method ServerKeysClient.BeginDelete.
type ServerKeysClientDeleteResponse struct {
	// placeholder for future response values
}

// ServerKeysClientGetResponse contains the response from method ServerKeysClient.Get.
type ServerKeysClientGetResponse struct {
	// A PostgreSQL Server key.
	ServerKey
}

// ServerKeysClientListResponse contains the response from method ServerKeysClient.NewListPager.
type ServerKeysClientListResponse struct {
	// A list of PostgreSQL Server keys.
	ServerKeyListResult
}

// ServerParametersClientListUpdateConfigurationsResponse contains the response from method ServerParametersClient.BeginListUpdateConfigurations.
type ServerParametersClientListUpdateConfigurationsResponse struct {
	// A list of server configurations.
	ConfigurationListResult
}

// ServerSecurityAlertPoliciesClientCreateOrUpdateResponse contains the response from method ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate.
type ServerSecurityAlertPoliciesClientCreateOrUpdateResponse struct {
	// A server security alert policy.
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesClientGetResponse contains the response from method ServerSecurityAlertPoliciesClient.Get.
type ServerSecurityAlertPoliciesClientGetResponse struct {
	// A server security alert policy.
	ServerSecurityAlertPolicy
}

// ServerSecurityAlertPoliciesClientListByServerResponse contains the response from method ServerSecurityAlertPoliciesClient.NewListByServerPager.
type ServerSecurityAlertPoliciesClientListByServerResponse struct {
	// A list of the server's security alert policies.
	ServerSecurityAlertPolicyListResult
}

// ServersClientCreateResponse contains the response from method ServersClient.BeginCreate.
type ServersClientCreateResponse struct {
	// Represents a server.
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.BeginDelete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	// Represents a server.
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.NewListByResourceGroupPager.
type ServersClientListByResourceGroupResponse struct {
	// A list of servers.
	ServerListResult
}

// ServersClientListResponse contains the response from method ServersClient.NewListPager.
type ServersClientListResponse struct {
	// A list of servers.
	ServerListResult
}

// ServersClientRestartResponse contains the response from method ServersClient.BeginRestart.
type ServersClientRestartResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.BeginUpdate.
type ServersClientUpdateResponse struct {
	// Represents a server.
	Server
}

// VirtualNetworkRulesClientCreateOrUpdateResponse contains the response from method VirtualNetworkRulesClient.BeginCreateOrUpdate.
type VirtualNetworkRulesClientCreateOrUpdateResponse struct {
	// A virtual network rule.
	VirtualNetworkRule
}

// VirtualNetworkRulesClientDeleteResponse contains the response from method VirtualNetworkRulesClient.BeginDelete.
type VirtualNetworkRulesClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualNetworkRulesClientGetResponse contains the response from method VirtualNetworkRulesClient.Get.
type VirtualNetworkRulesClientGetResponse struct {
	// A virtual network rule.
	VirtualNetworkRule
}

// VirtualNetworkRulesClientListByServerResponse contains the response from method VirtualNetworkRulesClient.NewListByServerPager.
type VirtualNetworkRulesClientListByServerResponse struct {
	// A list of virtual network rules.
	VirtualNetworkRuleListResult
}
