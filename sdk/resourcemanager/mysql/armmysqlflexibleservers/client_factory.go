//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
)

// ClientFactory is a client factory used to create any client in this module.
// Don't use this type directly, use NewClientFactory instead.
type ClientFactory struct {
	subscriptionID string
	credential     azcore.TokenCredential
	options        *arm.ClientOptions
}

// NewClientFactory creates a new instance of ClientFactory with the specified values.
// The parameter values will be propagated to any client created from this factory.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClientFactory(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClientFactory, error) {
	_, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	return &ClientFactory{
		subscriptionID: subscriptionID, credential: credential,
		options: options.Clone(),
	}, nil
}

// NewAdvancedThreatProtectionSettingsClient creates a new instance of AdvancedThreatProtectionSettingsClient.
func (c *ClientFactory) NewAdvancedThreatProtectionSettingsClient() *AdvancedThreatProtectionSettingsClient {
	subClient, _ := NewAdvancedThreatProtectionSettingsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewAzureADAdministratorsClient creates a new instance of AzureADAdministratorsClient.
func (c *ClientFactory) NewAzureADAdministratorsClient() *AzureADAdministratorsClient {
	subClient, _ := NewAzureADAdministratorsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBackupAndExportClient creates a new instance of BackupAndExportClient.
func (c *ClientFactory) NewBackupAndExportClient() *BackupAndExportClient {
	subClient, _ := NewBackupAndExportClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewBackupsClient creates a new instance of BackupsClient.
func (c *ClientFactory) NewBackupsClient() *BackupsClient {
	subClient, _ := NewBackupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCheckNameAvailabilityClient creates a new instance of CheckNameAvailabilityClient.
func (c *ClientFactory) NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient {
	subClient, _ := NewCheckNameAvailabilityClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCheckNameAvailabilityWithoutLocationClient creates a new instance of CheckNameAvailabilityWithoutLocationClient.
func (c *ClientFactory) NewCheckNameAvailabilityWithoutLocationClient() *CheckNameAvailabilityWithoutLocationClient {
	subClient, _ := NewCheckNameAvailabilityWithoutLocationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewCheckVirtualNetworkSubnetUsageClient creates a new instance of CheckVirtualNetworkSubnetUsageClient.
func (c *ClientFactory) NewCheckVirtualNetworkSubnetUsageClient() *CheckVirtualNetworkSubnetUsageClient {
	subClient, _ := NewCheckVirtualNetworkSubnetUsageClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewConfigurationsClient creates a new instance of ConfigurationsClient.
func (c *ClientFactory) NewConfigurationsClient() *ConfigurationsClient {
	subClient, _ := NewConfigurationsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewDatabasesClient creates a new instance of DatabasesClient.
func (c *ClientFactory) NewDatabasesClient() *DatabasesClient {
	subClient, _ := NewDatabasesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewFirewallRulesClient creates a new instance of FirewallRulesClient.
func (c *ClientFactory) NewFirewallRulesClient() *FirewallRulesClient {
	subClient, _ := NewFirewallRulesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewGetPrivateDNSZoneSuffixClient creates a new instance of GetPrivateDNSZoneSuffixClient.
func (c *ClientFactory) NewGetPrivateDNSZoneSuffixClient() *GetPrivateDNSZoneSuffixClient {
	subClient, _ := NewGetPrivateDNSZoneSuffixClient(c.credential, c.options)
	return subClient
}

// NewLocationBasedCapabilitiesClient creates a new instance of LocationBasedCapabilitiesClient.
func (c *ClientFactory) NewLocationBasedCapabilitiesClient() *LocationBasedCapabilitiesClient {
	subClient, _ := NewLocationBasedCapabilitiesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLocationBasedCapabilitySetClient creates a new instance of LocationBasedCapabilitySetClient.
func (c *ClientFactory) NewLocationBasedCapabilitySetClient() *LocationBasedCapabilitySetClient {
	subClient, _ := NewLocationBasedCapabilitySetClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLogFilesClient creates a new instance of LogFilesClient.
func (c *ClientFactory) NewLogFilesClient() *LogFilesClient {
	subClient, _ := NewLogFilesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLongRunningBackupClient creates a new instance of LongRunningBackupClient.
func (c *ClientFactory) NewLongRunningBackupClient() *LongRunningBackupClient {
	subClient, _ := NewLongRunningBackupClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewLongRunningBackupsClient creates a new instance of LongRunningBackupsClient.
func (c *ClientFactory) NewLongRunningBackupsClient() *LongRunningBackupsClient {
	subClient, _ := NewLongRunningBackupsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewMaintenancesClient creates a new instance of MaintenancesClient.
func (c *ClientFactory) NewMaintenancesClient() *MaintenancesClient {
	subClient, _ := NewMaintenancesClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationProgressClient creates a new instance of OperationProgressClient.
func (c *ClientFactory) NewOperationProgressClient() *OperationProgressClient {
	subClient, _ := NewOperationProgressClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationResultsClient creates a new instance of OperationResultsClient.
func (c *ClientFactory) NewOperationResultsClient() *OperationResultsClient {
	subClient, _ := NewOperationResultsClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewOperationsClient creates a new instance of OperationsClient.
func (c *ClientFactory) NewOperationsClient() *OperationsClient {
	subClient, _ := NewOperationsClient(c.credential, c.options)
	return subClient
}

// NewReplicasClient creates a new instance of ReplicasClient.
func (c *ClientFactory) NewReplicasClient() *ReplicasClient {
	subClient, _ := NewReplicasClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServersClient creates a new instance of ServersClient.
func (c *ClientFactory) NewServersClient() *ServersClient {
	subClient, _ := NewServersClient(c.subscriptionID, c.credential, c.options)
	return subClient
}

// NewServersMigrationClient creates a new instance of ServersMigrationClient.
func (c *ClientFactory) NewServersMigrationClient() *ServersMigrationClient {
	subClient, _ := NewServersMigrationClient(c.subscriptionID, c.credential, c.options)
	return subClient
}
