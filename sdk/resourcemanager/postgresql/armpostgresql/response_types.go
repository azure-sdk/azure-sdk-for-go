//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

// AdministratorsClientCreateResponse contains the response from method AdministratorsClient.BeginCreate.
type AdministratorsClientCreateResponse struct {
	// Represents an Active Directory administrator.
	ActiveDirectoryAdministrator
}

// AdministratorsClientDeleteResponse contains the response from method AdministratorsClient.BeginDelete.
type AdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// AdministratorsClientGetResponse contains the response from method AdministratorsClient.Get.
type AdministratorsClientGetResponse struct {
	// Represents an Active Directory administrator.
	ActiveDirectoryAdministrator
}

// AdministratorsClientListByServerResponse contains the response from method AdministratorsClient.NewListByServerPager.
type AdministratorsClientListByServerResponse struct {
	// A list of active directory administrators.
	AdministratorListResult
}

// BackupsClientGetResponse contains the response from method BackupsClient.Get.
type BackupsClientGetResponse struct {
	// Server backup properties
	ServerBackup
}

// BackupsClientListByServerResponse contains the response from method BackupsClient.NewListByServerPager.
type BackupsClientListByServerResponse struct {
	// A list of server backups.
	ServerBackupListResult
}

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
}

// CheckNameAvailabilityWithLocationClientExecuteResponse contains the response from method CheckNameAvailabilityWithLocationClient.Execute.
type CheckNameAvailabilityWithLocationClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
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

// ConfigurationsClientPutResponse contains the response from method ConfigurationsClient.BeginPut.
type ConfigurationsClientPutResponse struct {
	// Represents a Configuration.
	Configuration
}

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.BeginUpdate.
type ConfigurationsClientUpdateResponse struct {
	// Represents a Configuration.
	Configuration
}

// DatabasesClientCreateResponse contains the response from method DatabasesClient.BeginCreate.
type DatabasesClientCreateResponse struct {
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

// FlexibleServerClientStartLtrBackupResponse contains the response from method FlexibleServerClient.BeginStartLtrBackup.
type FlexibleServerClientStartLtrBackupResponse struct {
	// Response for the LTR backup API call
	LtrBackupResponse
}

// FlexibleServerClientTriggerLtrPreBackupResponse contains the response from method FlexibleServerClient.TriggerLtrPreBackup.
type FlexibleServerClientTriggerLtrPreBackupResponse struct {
	// Response for the LTR pre-backup API call
	LtrPreBackupResponse

	// XMSRequestID contains the information returned from the x-ms-request-id header response.
	XMSRequestID *string
}

// GetPrivateDNSZoneSuffixClientExecuteResponse contains the response from method GetPrivateDNSZoneSuffixClient.Execute.
type GetPrivateDNSZoneSuffixClientExecuteResponse struct {
	// Represents a resource name availability.
	Value *string
}

// LocationBasedCapabilitiesClientExecuteResponse contains the response from method LocationBasedCapabilitiesClient.NewExecutePager.
type LocationBasedCapabilitiesClientExecuteResponse struct {
	// Capability for the PostgreSQL server
	CapabilitiesListResult
}

// LogFilesClientListByServerResponse contains the response from method LogFilesClient.NewListByServerPager.
type LogFilesClientListByServerResponse struct {
	// A List of logFiles.
	LogFileListResult
}

// LtrBackupOperationsClientGetResponse contains the response from method LtrBackupOperationsClient.Get.
type LtrBackupOperationsClientGetResponse struct {
	// Response for the LTR backup Operation API call
	LtrServerBackupOperation
}

// LtrBackupOperationsClientListByServerResponse contains the response from method LtrBackupOperationsClient.NewListByServerPager.
type LtrBackupOperationsClientListByServerResponse struct {
	// A list of long term retention backup operations for server.
	LtrServerBackupOperationList
}

// ManagementClientCheckMigrationNameAvailabilityResponse contains the response from method ManagementClient.CheckMigrationNameAvailability.
type ManagementClientCheckMigrationNameAvailabilityResponse struct {
	// Represents a migration name's availability.
	MigrationNameAvailabilityResource
}

// MigrationsClientCreateResponse contains the response from method MigrationsClient.Create.
type MigrationsClientCreateResponse struct {
	// Represents a migration resource.
	MigrationResource
}

// MigrationsClientDeleteResponse contains the response from method MigrationsClient.Delete.
type MigrationsClientDeleteResponse struct {
	// placeholder for future response values
}

// MigrationsClientGetResponse contains the response from method MigrationsClient.Get.
type MigrationsClientGetResponse struct {
	// Represents a migration resource.
	MigrationResource
}

// MigrationsClientListByTargetServerResponse contains the response from method MigrationsClient.NewListByTargetServerPager.
type MigrationsClientListByTargetServerResponse struct {
	// A list of migration resources.
	MigrationResourceListResult
}

// MigrationsClientUpdateResponse contains the response from method MigrationsClient.Update.
type MigrationsClientUpdateResponse struct {
	// Represents a migration resource.
	MigrationResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of resource provider operations.
	OperationListResult
}

// PrivateEndpointConnectionClientDeleteResponse contains the response from method PrivateEndpointConnectionClient.BeginDelete.
type PrivateEndpointConnectionClientDeleteResponse struct {
	// placeholder for future response values
}

// PrivateEndpointConnectionClientUpdateResponse contains the response from method PrivateEndpointConnectionClient.BeginUpdate.
type PrivateEndpointConnectionClientUpdateResponse struct {
	// The private endpoint connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// The private endpoint connection resource.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByServerResponse contains the response from method PrivateEndpointConnectionsClient.NewListByServerPager.
type PrivateEndpointConnectionsClientListByServerResponse struct {
	// A list of private endpoint connections.
	PrivateEndpointConnectionListResult
}

// PrivateLinkResourcesClientGetResponse contains the response from method PrivateLinkResourcesClient.Get.
type PrivateLinkResourcesClientGetResponse struct {
	// A private link resource.
	PrivateLinkResource
}

// PrivateLinkResourcesClientListByServerResponse contains the response from method PrivateLinkResourcesClient.NewListByServerPager.
type PrivateLinkResourcesClientListByServerResponse struct {
	// A list of private link resources
	PrivateLinkResourceListResult
}

// QuotaUsagesClientListResponse contains the response from method QuotaUsagesClient.NewListPager.
type QuotaUsagesClientListResponse struct {
	// Capability for the PostgreSQL server
	QuotaUsagesListResult
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.NewListByServerPager.
type ReplicasClientListByServerResponse struct {
	// A list of servers.
	ServerListResult
}

// ServerCapabilitiesClientListResponse contains the response from method ServerCapabilitiesClient.NewListPager.
type ServerCapabilitiesClientListResponse struct {
	// Capability for the PostgreSQL server
	CapabilitiesListResult
}

// ServerThreatProtectionSettingsClientCreateOrUpdateResponse contains the response from method ServerThreatProtectionSettingsClient.BeginCreateOrUpdate.
type ServerThreatProtectionSettingsClientCreateOrUpdateResponse struct {
	// Server's Advanced Threat Protection settings.
	ServerThreatProtectionSettingsModel
}

// ServerThreatProtectionSettingsClientGetResponse contains the response from method ServerThreatProtectionSettingsClient.Get.
type ServerThreatProtectionSettingsClientGetResponse struct {
	// Server's Advanced Threat Protection settings.
	ServerThreatProtectionSettingsModel
}

// ServerThreatProtectionSettingsClientListByServerResponse contains the response from method ServerThreatProtectionSettingsClient.NewListByServerPager.
type ServerThreatProtectionSettingsClientListByServerResponse struct {
	// A list of the server's Advanced Threat Protection settings.
	ServerThreatProtectionListResult
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

// ServersClientStartResponse contains the response from method ServersClient.BeginStart.
type ServersClientStartResponse struct {
	// placeholder for future response values
}

// ServersClientStopResponse contains the response from method ServersClient.BeginStop.
type ServersClientStopResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.BeginUpdate.
type ServersClientUpdateResponse struct {
	// Represents a server.
	Server
}

// VirtualEndpointsClientCreateResponse contains the response from method VirtualEndpointsClient.BeginCreate.
type VirtualEndpointsClientCreateResponse struct {
	// Represents a virtual endpoint for a server.
	VirtualEndpointResource
}

// VirtualEndpointsClientDeleteResponse contains the response from method VirtualEndpointsClient.BeginDelete.
type VirtualEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// VirtualEndpointsClientGetResponse contains the response from method VirtualEndpointsClient.Get.
type VirtualEndpointsClientGetResponse struct {
	// Represents a virtual endpoint for a server.
	VirtualEndpointResource
}

// VirtualEndpointsClientListByServerResponse contains the response from method VirtualEndpointsClient.NewListByServerPager.
type VirtualEndpointsClientListByServerResponse struct {
	// A list of virtual endpoints.
	VirtualEndpointsListResult
}

// VirtualEndpointsClientUpdateResponse contains the response from method VirtualEndpointsClient.BeginUpdate.
type VirtualEndpointsClientUpdateResponse struct {
	// Represents a virtual endpoint for a server.
	VirtualEndpointResource
}

// VirtualNetworkSubnetUsageClientExecuteResponse contains the response from method VirtualNetworkSubnetUsageClient.Execute.
type VirtualNetworkSubnetUsageClientExecuteResponse struct {
	// Virtual network subnet usage data.
	VirtualNetworkSubnetUsageResult
}
