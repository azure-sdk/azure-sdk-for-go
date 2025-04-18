// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

// AdvancedThreatProtectionSettingsClientGetResponse contains the response from method AdvancedThreatProtectionSettingsClient.Get.
type AdvancedThreatProtectionSettingsClientGetResponse struct {
	// A server's Advanced Threat Protection.
	AdvancedThreatProtection
}

// AdvancedThreatProtectionSettingsClientListResponse contains the response from method AdvancedThreatProtectionSettingsClient.NewListPager.
type AdvancedThreatProtectionSettingsClientListResponse struct {
	// A list of the server's Advanced Threat Protection configurations.
	AdvancedThreatProtectionListResult
}

// AdvancedThreatProtectionSettingsClientUpdatePutResponse contains the response from method AdvancedThreatProtectionSettingsClient.BeginUpdatePut.
type AdvancedThreatProtectionSettingsClientUpdatePutResponse struct {
	// A server's Advanced Threat Protection.
	AdvancedThreatProtection
}

// AdvancedThreatProtectionSettingsClientUpdateResponse contains the response from method AdvancedThreatProtectionSettingsClient.BeginUpdate.
type AdvancedThreatProtectionSettingsClientUpdateResponse struct {
	// A server's Advanced Threat Protection.
	AdvancedThreatProtection
}

// AzureADAdministratorsClientCreateOrUpdateResponse contains the response from method AzureADAdministratorsClient.BeginCreateOrUpdate.
type AzureADAdministratorsClientCreateOrUpdateResponse struct {
	// Represents a Administrator.
	AzureADAdministrator
}

// AzureADAdministratorsClientDeleteResponse contains the response from method AzureADAdministratorsClient.BeginDelete.
type AzureADAdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// AzureADAdministratorsClientGetResponse contains the response from method AzureADAdministratorsClient.Get.
type AzureADAdministratorsClientGetResponse struct {
	// Represents a Administrator.
	AzureADAdministrator
}

// AzureADAdministratorsClientListByServerResponse contains the response from method AzureADAdministratorsClient.NewListByServerPager.
type AzureADAdministratorsClientListByServerResponse struct {
	// A List of azure ad administrators.
	AdministratorListResult
}

// BackupAndExportClientCreateResponse contains the response from method BackupAndExportClient.BeginCreate.
type BackupAndExportClientCreateResponse struct {
	// Represents BackupAndExport API Response
	BackupAndExportResponse
}

// BackupAndExportClientValidateBackupResponse contains the response from method BackupAndExportClient.ValidateBackup.
type BackupAndExportClientValidateBackupResponse struct {
	// Represents ValidateBackup API Response
	ValidateBackupResponse
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

// BackupsClientPutResponse contains the response from method BackupsClient.Put.
type BackupsClientPutResponse struct {
	// Server backup properties
	ServerBackup
}

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
}

// CheckNameAvailabilityWithoutLocationClientExecuteResponse contains the response from method CheckNameAvailabilityWithoutLocationClient.Execute.
type CheckNameAvailabilityWithoutLocationClientExecuteResponse struct {
	// Represents a resource name availability.
	NameAvailability
}

// CheckVirtualNetworkSubnetUsageClientExecuteResponse contains the response from method CheckVirtualNetworkSubnetUsageClient.Execute.
type CheckVirtualNetworkSubnetUsageClientExecuteResponse struct {
	// Virtual network subnet usage data.
	VirtualNetworkSubnetUsageResult
}

// ConfigurationsClientBatchUpdateResponse contains the response from method ConfigurationsClient.BeginBatchUpdate.
type ConfigurationsClientBatchUpdateResponse struct {
	// A list of server configurations.
	ConfigurationListResult
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

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.BeginUpdate.
type ConfigurationsClientUpdateResponse struct {
	// Represents a Configuration.
	Configuration
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

// GetPrivateDNSZoneSuffixClientExecuteResponse contains the response from method GetPrivateDNSZoneSuffixClient.Execute.
type GetPrivateDNSZoneSuffixClientExecuteResponse struct {
	// The response of get private dns zone suffix.
	GetPrivateDNSZoneSuffixResponse
}

// LocationBasedCapabilitiesClientListResponse contains the response from method LocationBasedCapabilitiesClient.NewListPager.
type LocationBasedCapabilitiesClientListResponse struct {
	// location capability
	CapabilitiesListResult
}

// LocationBasedCapabilitySetClientGetResponse contains the response from method LocationBasedCapabilitySetClient.Get.
type LocationBasedCapabilitySetClientGetResponse struct {
	// Represents a location capability set.
	Capability
}

// LocationBasedCapabilitySetClientListResponse contains the response from method LocationBasedCapabilitySetClient.NewListPager.
type LocationBasedCapabilitySetClientListResponse struct {
	// location capability set
	CapabilitySetsList
}

// LogFilesClientListByServerResponse contains the response from method LogFilesClient.NewListByServerPager.
type LogFilesClientListByServerResponse struct {
	// A List of logFiles.
	LogFileListResult
}

// LongRunningBackupClientCreateResponse contains the response from method LongRunningBackupClient.BeginCreate.
type LongRunningBackupClientCreateResponse struct {
	// Server backup properties
	ServerBackupV2
}

// LongRunningBackupsClientGetResponse contains the response from method LongRunningBackupsClient.Get.
type LongRunningBackupsClientGetResponse struct {
	// Server backup properties
	ServerBackupV2
}

// LongRunningBackupsClientListResponse contains the response from method LongRunningBackupsClient.NewListPager.
type LongRunningBackupsClientListResponse struct {
	// A list of server backups.
	ServerBackupV2ListResult
}

// MaintenancesClientListResponse contains the response from method MaintenancesClient.NewListPager.
type MaintenancesClientListResponse struct {
	// A list of maintenances.
	MaintenanceListResult
}

// MaintenancesClientReadResponse contains the response from method MaintenancesClient.Read.
type MaintenancesClientReadResponse struct {
	// Represents a maintenance.
	Maintenance
}

// MaintenancesClientUpdateResponse contains the response from method MaintenancesClient.BeginUpdate.
type MaintenancesClientUpdateResponse struct {
	// Represents a maintenance.
	Maintenance
}

// OperationProgressClientGetResponse contains the response from method OperationProgressClient.Get.
type OperationProgressClientGetResponse struct {
	// Represents Operation Results API Response
	OperationProgressResult
}

// OperationResultsClientGetResponse contains the response from method OperationResultsClient.Get.
type OperationResultsClientGetResponse struct {
	// Represents Operation Results API Response
	OperationStatusExtendedResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of resource provider operations.
	OperationListResult
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.NewListByServerPager.
type ReplicasClientListByServerResponse struct {
	// A list of servers.
	ServerListResult
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

// ServersClientDetachVNetResponse contains the response from method ServersClient.BeginDetachVNet.
type ServersClientDetachVNetResponse struct {
	// Represents a server.
	Server
}

// ServersClientFailoverResponse contains the response from method ServersClient.BeginFailover.
type ServersClientFailoverResponse struct {
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

// ServersClientResetGtidResponse contains the response from method ServersClient.BeginResetGtid.
type ServersClientResetGtidResponse struct {
	// placeholder for future response values
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

// ServersClientValidateEstimateHighAvailabilityResponse contains the response from method ServersClient.ValidateEstimateHighAvailability.
type ServersClientValidateEstimateHighAvailabilityResponse struct {
	// High availability validation properties of a server
	HighAvailabilityValidationEstimation
}

// ServersMigrationClientCutoverMigrationResponse contains the response from method ServersMigrationClient.BeginCutoverMigration.
type ServersMigrationClientCutoverMigrationResponse struct {
	// Represents a server.
	Server
}
