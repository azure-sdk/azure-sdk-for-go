//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmysqlflexibleservers

// AzureADAdministratorsClientCreateOrUpdateResponse contains the response from method AzureADAdministratorsClient.BeginCreateOrUpdate.
type AzureADAdministratorsClientCreateOrUpdateResponse struct {
	AzureADAdministrator
}

// AzureADAdministratorsClientDeleteResponse contains the response from method AzureADAdministratorsClient.BeginDelete.
type AzureADAdministratorsClientDeleteResponse struct {
	// placeholder for future response values
}

// AzureADAdministratorsClientGetResponse contains the response from method AzureADAdministratorsClient.Get.
type AzureADAdministratorsClientGetResponse struct {
	AzureADAdministrator
}

// AzureADAdministratorsClientListByServerResponse contains the response from method AzureADAdministratorsClient.NewListByServerPager.
type AzureADAdministratorsClientListByServerResponse struct {
	AdministratorListResult
}

// BackupsClientGetResponse contains the response from method BackupsClient.Get.
type BackupsClientGetResponse struct {
	ServerBackup
}

// BackupsClientListByServerResponse contains the response from method BackupsClient.NewListByServerPager.
type BackupsClientListByServerResponse struct {
	ServerBackupListResult
}

// BackupsClientPutResponse contains the response from method BackupsClient.Put.
type BackupsClientPutResponse struct {
	ServerBackup
}

// CheckNameAvailabilityClientExecuteResponse contains the response from method CheckNameAvailabilityClient.Execute.
type CheckNameAvailabilityClientExecuteResponse struct {
	NameAvailability
}

// CheckNameAvailabilityWithoutLocationClientExecuteResponse contains the response from method CheckNameAvailabilityWithoutLocationClient.Execute.
type CheckNameAvailabilityWithoutLocationClientExecuteResponse struct {
	NameAvailability
}

// CheckVirtualNetworkSubnetUsageClientExecuteResponse contains the response from method CheckVirtualNetworkSubnetUsageClient.Execute.
type CheckVirtualNetworkSubnetUsageClientExecuteResponse struct {
	VirtualNetworkSubnetUsageResult
}

// ConfigurationsClientBatchUpdateResponse contains the response from method ConfigurationsClient.BeginBatchUpdate.
type ConfigurationsClientBatchUpdateResponse struct {
	ConfigurationListResult
}

// ConfigurationsClientCreateOrUpdateResponse contains the response from method ConfigurationsClient.BeginCreateOrUpdate.
type ConfigurationsClientCreateOrUpdateResponse struct {
	Configuration
}

// ConfigurationsClientGetResponse contains the response from method ConfigurationsClient.Get.
type ConfigurationsClientGetResponse struct {
	Configuration
}

// ConfigurationsClientListByServerResponse contains the response from method ConfigurationsClient.NewListByServerPager.
type ConfigurationsClientListByServerResponse struct {
	ConfigurationListResult
}

// ConfigurationsClientUpdateResponse contains the response from method ConfigurationsClient.BeginUpdate.
type ConfigurationsClientUpdateResponse struct {
	Configuration
}

// DatabasesClientCreateOrUpdateResponse contains the response from method DatabasesClient.BeginCreateOrUpdate.
type DatabasesClientCreateOrUpdateResponse struct {
	Database
}

// DatabasesClientDeleteResponse contains the response from method DatabasesClient.BeginDelete.
type DatabasesClientDeleteResponse struct {
	// placeholder for future response values
}

// DatabasesClientGetResponse contains the response from method DatabasesClient.Get.
type DatabasesClientGetResponse struct {
	Database
}

// DatabasesClientListByServerResponse contains the response from method DatabasesClient.NewListByServerPager.
type DatabasesClientListByServerResponse struct {
	DatabaseListResult
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

// FirewallRulesClientListByServerResponse contains the response from method FirewallRulesClient.NewListByServerPager.
type FirewallRulesClientListByServerResponse struct {
	FirewallRuleListResult
}

// GetPrivateDNSZoneSuffixClientExecuteResponse contains the response from method GetPrivateDNSZoneSuffixClient.Execute.
type GetPrivateDNSZoneSuffixClientExecuteResponse struct {
	GetPrivateDNSZoneSuffixResponse
}

// LocationBasedCapabilitiesClientListResponse contains the response from method LocationBasedCapabilitiesClient.NewListPager.
type LocationBasedCapabilitiesClientListResponse struct {
	CapabilitiesListResult
}

// LogFilesClientListByServerResponse contains the response from method LogFilesClient.NewListByServerPager.
type LogFilesClientListByServerResponse struct {
	LogFileListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// ReplicasClientListByServerResponse contains the response from method ReplicasClient.NewListByServerPager.
type ReplicasClientListByServerResponse struct {
	ServerListResult
}

// ServersClientCreateResponse contains the response from method ServersClient.BeginCreate.
type ServersClientCreateResponse struct {
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.BeginDelete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientFailoverResponse contains the response from method ServersClient.BeginFailover.
type ServersClientFailoverResponse struct {
	// placeholder for future response values
}

// ServersClientGetResponse contains the response from method ServersClient.Get.
type ServersClientGetResponse struct {
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.NewListByResourceGroupPager.
type ServersClientListByResourceGroupResponse struct {
	ServerListResult
}

// ServersClientListResponse contains the response from method ServersClient.NewListPager.
type ServersClientListResponse struct {
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
	Server
}
