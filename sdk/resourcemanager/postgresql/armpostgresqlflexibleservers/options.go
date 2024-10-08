//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

// AdministratorsClientBeginCreateOptions contains the optional parameters for the AdministratorsClient.BeginCreate method.
type AdministratorsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AdministratorsClientBeginDeleteOptions contains the optional parameters for the AdministratorsClient.BeginDelete method.
type AdministratorsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AdministratorsClientGetOptions contains the optional parameters for the AdministratorsClient.Get method.
type AdministratorsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AdministratorsClientListByServerOptions contains the optional parameters for the AdministratorsClient.NewListByServerPager
// method.
type AdministratorsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// BackupsClientBeginCreateOptions contains the optional parameters for the BackupsClient.BeginCreate method.
type BackupsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupsClientBeginDeleteOptions contains the optional parameters for the BackupsClient.BeginDelete method.
type BackupsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// BackupsClientGetOptions contains the optional parameters for the BackupsClient.Get method.
type BackupsClientGetOptions struct {
	// placeholder for future optional parameters
}

// BackupsClientListByServerOptions contains the optional parameters for the BackupsClient.NewListByServerPager method.
type BackupsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// CheckNameAvailabilityClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityClient.Execute
// method.
type CheckNameAvailabilityClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// CheckNameAvailabilityWithLocationClientExecuteOptions contains the optional parameters for the CheckNameAvailabilityWithLocationClient.Execute
// method.
type CheckNameAvailabilityWithLocationClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientBeginPutOptions contains the optional parameters for the ConfigurationsClient.BeginPut method.
type ConfigurationsClientBeginPutOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientBeginUpdateOptions contains the optional parameters for the ConfigurationsClient.BeginUpdate method.
type ConfigurationsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ConfigurationsClientGetOptions contains the optional parameters for the ConfigurationsClient.Get method.
type ConfigurationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ConfigurationsClientListByServerOptions contains the optional parameters for the ConfigurationsClient.NewListByServerPager
// method.
type ConfigurationsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// DatabasesClientBeginCreateOptions contains the optional parameters for the DatabasesClient.BeginCreate method.
type DatabasesClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabasesClientBeginDeleteOptions contains the optional parameters for the DatabasesClient.BeginDelete method.
type DatabasesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// DatabasesClientGetOptions contains the optional parameters for the DatabasesClient.Get method.
type DatabasesClientGetOptions struct {
	// placeholder for future optional parameters
}

// DatabasesClientListByServerOptions contains the optional parameters for the DatabasesClient.NewListByServerPager method.
type DatabasesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.BeginCreateOrUpdate
// method.
type FirewallRulesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FirewallRulesClientBeginDeleteOptions contains the optional parameters for the FirewallRulesClient.BeginDelete method.
type FirewallRulesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
type FirewallRulesClientGetOptions struct {
	// placeholder for future optional parameters
}

// FirewallRulesClientListByServerOptions contains the optional parameters for the FirewallRulesClient.NewListByServerPager
// method.
type FirewallRulesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// FlexibleServerClientBeginStartLtrBackupOptions contains the optional parameters for the FlexibleServerClient.BeginStartLtrBackup
// method.
type FlexibleServerClientBeginStartLtrBackupOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FlexibleServerClientTriggerLtrPreBackupOptions contains the optional parameters for the FlexibleServerClient.TriggerLtrPreBackup
// method.
type FlexibleServerClientTriggerLtrPreBackupOptions struct {
	// placeholder for future optional parameters
}

// GetPrivateDNSZoneSuffixClientExecuteOptions contains the optional parameters for the GetPrivateDNSZoneSuffixClient.Execute
// method.
type GetPrivateDNSZoneSuffixClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// LocationBasedCapabilitiesClientExecuteOptions contains the optional parameters for the LocationBasedCapabilitiesClient.NewExecutePager
// method.
type LocationBasedCapabilitiesClientExecuteOptions struct {
	// placeholder for future optional parameters
}

// LogFilesClientListByServerOptions contains the optional parameters for the LogFilesClient.NewListByServerPager method.
type LogFilesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// LtrBackupOperationsClientGetOptions contains the optional parameters for the LtrBackupOperationsClient.Get method.
type LtrBackupOperationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// LtrBackupOperationsClientListByServerOptions contains the optional parameters for the LtrBackupOperationsClient.NewListByServerPager
// method.
type LtrBackupOperationsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// MigrationsClientCreateOptions contains the optional parameters for the MigrationsClient.Create method.
type MigrationsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// MigrationsClientDeleteOptions contains the optional parameters for the MigrationsClient.Delete method.
type MigrationsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// MigrationsClientGetOptions contains the optional parameters for the MigrationsClient.Get method.
type MigrationsClientGetOptions struct {
	// placeholder for future optional parameters
}

// MigrationsClientListByTargetServerOptions contains the optional parameters for the MigrationsClient.NewListByTargetServerPager
// method.
type MigrationsClientListByTargetServerOptions struct {
	// Migration list filter. Retrieves either active migrations or all migrations.
	MigrationListFilter *MigrationListFilter
}

// MigrationsClientUpdateOptions contains the optional parameters for the MigrationsClient.Update method.
type MigrationsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.List method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PostgreSQLServerManagementClientCheckMigrationNameAvailabilityOptions contains the optional parameters for the PostgreSQLServerManagementClient.CheckMigrationNameAvailability
// method.
type PostgreSQLServerManagementClientCheckMigrationNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionClientBeginDeleteOptions contains the optional parameters for the PrivateEndpointConnectionClient.BeginDelete
// method.
type PrivateEndpointConnectionClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionClientBeginUpdateOptions contains the optional parameters for the PrivateEndpointConnectionClient.BeginUpdate
// method.
type PrivateEndpointConnectionClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByServerOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByServerPager
// method.
type PrivateEndpointConnectionsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByServerOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByServerPager
// method.
type PrivateLinkResourcesClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// ReplicasClientListByServerOptions contains the optional parameters for the ReplicasClient.NewListByServerPager method.
type ReplicasClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// ServerCapabilitiesClientListOptions contains the optional parameters for the ServerCapabilitiesClient.NewListPager method.
type ServerCapabilitiesClientListOptions struct {
	// placeholder for future optional parameters
}

// ServerThreatProtectionSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerThreatProtectionSettingsClient.BeginCreateOrUpdate
// method.
type ServerThreatProtectionSettingsClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServerThreatProtectionSettingsClientGetOptions contains the optional parameters for the ServerThreatProtectionSettingsClient.Get
// method.
type ServerThreatProtectionSettingsClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServerThreatProtectionSettingsClientListByServerOptions contains the optional parameters for the ServerThreatProtectionSettingsClient.NewListByServerPager
// method.
type ServerThreatProtectionSettingsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// ServersClientBeginCreateOptions contains the optional parameters for the ServersClient.BeginCreate method.
type ServersClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginDeleteOptions contains the optional parameters for the ServersClient.BeginDelete method.
type ServersClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginRestartOptions contains the optional parameters for the ServersClient.BeginRestart method.
type ServersClientBeginRestartOptions struct {
	// The parameters for restarting a server.
	Parameters *RestartParameter

	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginStartOptions contains the optional parameters for the ServersClient.BeginStart method.
type ServersClientBeginStartOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginStopOptions contains the optional parameters for the ServersClient.BeginStop method.
type ServersClientBeginStopOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientBeginUpdateOptions contains the optional parameters for the ServersClient.BeginUpdate method.
type ServersClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServersClientGetOptions contains the optional parameters for the ServersClient.Get method.
type ServersClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServersClientListByResourceGroupOptions contains the optional parameters for the ServersClient.NewListByResourceGroupPager
// method.
type ServersClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServersClientListOptions contains the optional parameters for the ServersClient.NewListPager method.
type ServersClientListOptions struct {
	// placeholder for future optional parameters
}

// VirtualEndpointsClientBeginCreateOptions contains the optional parameters for the VirtualEndpointsClient.BeginCreate method.
type VirtualEndpointsClientBeginCreateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualEndpointsClientBeginDeleteOptions contains the optional parameters for the VirtualEndpointsClient.BeginDelete method.
type VirtualEndpointsClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualEndpointsClientBeginUpdateOptions contains the optional parameters for the VirtualEndpointsClient.BeginUpdate method.
type VirtualEndpointsClientBeginUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// VirtualEndpointsClientGetOptions contains the optional parameters for the VirtualEndpointsClient.Get method.
type VirtualEndpointsClientGetOptions struct {
	// placeholder for future optional parameters
}

// VirtualEndpointsClientListByServerOptions contains the optional parameters for the VirtualEndpointsClient.NewListByServerPager
// method.
type VirtualEndpointsClientListByServerOptions struct {
	// placeholder for future optional parameters
}

// VirtualNetworkSubnetUsageClientExecuteOptions contains the optional parameters for the VirtualNetworkSubnetUsageClient.Execute
// method.
type VirtualNetworkSubnetUsageClientExecuteOptions struct {
	// placeholder for future optional parameters
}
