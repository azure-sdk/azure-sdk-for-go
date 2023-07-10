# Release History

## 2.0.0 (2023-07-10)
### Breaking Changes

- Enum `ConfigurationSource` has been removed
- Enum `CreateMode` has been removed
- Enum `DataEncryptionType` has been removed
- Enum `EnableStatusEnum` has been removed
- Enum `HighAvailabilityMode` has been removed
- Enum `HighAvailabilityState` has been removed
- Enum `IsConfigPendingRestart` has been removed
- Enum `IsDynamicConfig` has been removed
- Enum `IsReadOnly` has been removed
- Enum `ReplicationRole` has been removed
- Enum `SKUTier` has been removed
- Enum `ServerState` has been removed
- Enum `ServerVersion` has been removed
- Function `NewBackupsClient` has been removed
- Function `*BackupsClient.Get` has been removed
- Function `*BackupsClient.NewListByServerPager` has been removed
- Function `NewCheckNameAvailabilityClient` has been removed
- Function `*CheckNameAvailabilityClient.Execute` has been removed
- Function `NewCheckVirtualNetworkSubnetUsageClient` has been removed
- Function `*CheckVirtualNetworkSubnetUsageClient.Execute` has been removed
- Function `*ClientFactory.NewBackupsClient` has been removed
- Function `*ClientFactory.NewCheckNameAvailabilityClient` has been removed
- Function `*ClientFactory.NewCheckVirtualNetworkSubnetUsageClient` has been removed
- Function `*ClientFactory.NewConfigurationsClient` has been removed
- Function `*ClientFactory.NewDatabasesClient` has been removed
- Function `*ClientFactory.NewFirewallRulesClient` has been removed
- Function `*ClientFactory.NewGetPrivateDNSZoneSuffixClient` has been removed
- Function `*ClientFactory.NewLocationBasedCapabilitiesClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewReplicasClient` has been removed
- Function `*ClientFactory.NewServersClient` has been removed
- Function `NewConfigurationsClient` has been removed
- Function `*ConfigurationsClient.Get` has been removed
- Function `*ConfigurationsClient.NewListByServerPager` has been removed
- Function `*ConfigurationsClient.BeginUpdate` has been removed
- Function `*ConfigurationsClient.BeginBatchUpdate` has been removed
- Function `NewDatabasesClient` has been removed
- Function `*DatabasesClient.BeginCreateOrUpdate` has been removed
- Function `*DatabasesClient.BeginDelete` has been removed
- Function `*DatabasesClient.Get` has been removed
- Function `*DatabasesClient.NewListByServerPager` has been removed
- Function `NewFirewallRulesClient` has been removed
- Function `*FirewallRulesClient.BeginCreateOrUpdate` has been removed
- Function `*FirewallRulesClient.BeginDelete` has been removed
- Function `*FirewallRulesClient.Get` has been removed
- Function `*FirewallRulesClient.NewListByServerPager` has been removed
- Function `NewGetPrivateDNSZoneSuffixClient` has been removed
- Function `*GetPrivateDNSZoneSuffixClient.Execute` has been removed
- Function `NewLocationBasedCapabilitiesClient` has been removed
- Function `*LocationBasedCapabilitiesClient.NewListPager` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewReplicasClient` has been removed
- Function `*ReplicasClient.NewListByServerPager` has been removed
- Function `NewServersClient` has been removed
- Function `*ServersClient.BeginCreate` has been removed
- Function `*ServersClient.BeginDelete` has been removed
- Function `*ServersClient.BeginFailover` has been removed
- Function `*ServersClient.Get` has been removed
- Function `*ServersClient.NewListByResourceGroupPager` has been removed
- Function `*ServersClient.NewListPager` has been removed
- Function `*ServersClient.BeginRestart` has been removed
- Function `*ServersClient.BeginStart` has been removed
- Function `*ServersClient.BeginStop` has been removed
- Function `*ServersClient.BeginUpdate` has been removed
- Struct `Backup` has been removed
- Struct `CapabilitiesListResult` has been removed
- Struct `CapabilityProperties` has been removed
- Struct `Configuration` has been removed
- Struct `ConfigurationForBatchUpdate` has been removed
- Struct `ConfigurationForBatchUpdateProperties` has been removed
- Struct `ConfigurationListForBatchUpdate` has been removed
- Struct `ConfigurationListResult` has been removed
- Struct `ConfigurationProperties` has been removed
- Struct `DataEncryption` has been removed
- Struct `Database` has been removed
- Struct `DatabaseListResult` has been removed
- Struct `DatabaseProperties` has been removed
- Struct `DelegatedSubnetUsage` has been removed
- Struct `FirewallRule` has been removed
- Struct `FirewallRuleListResult` has been removed
- Struct `FirewallRuleProperties` has been removed
- Struct `GetPrivateDNSZoneSuffixResponse` has been removed
- Struct `HighAvailability` has been removed
- Struct `Identity` has been removed
- Struct `MaintenanceWindow` has been removed
- Struct `NameAvailability` has been removed
- Struct `NameAvailabilityRequest` has been removed
- Struct `Network` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed
- Struct `ProxyResource` has been removed
- Struct `SKU` has been removed
- Struct `SKUCapability` has been removed
- Struct `Server` has been removed
- Struct `ServerBackup` has been removed
- Struct `ServerBackupListResult` has been removed
- Struct `ServerBackupProperties` has been removed
- Struct `ServerEditionCapability` has been removed
- Struct `ServerForUpdate` has been removed
- Struct `ServerListResult` has been removed
- Struct `ServerProperties` has been removed
- Struct `ServerPropertiesForUpdate` has been removed
- Struct `ServerRestartParameter` has been removed
- Struct `ServerVersionCapability` has been removed
- Struct `Storage` has been removed
- Struct `StorageEditionCapability` has been removed
- Struct `TrackedResource` has been removed
- Struct `UserAssignedIdentity` has been removed
- Struct `VirtualNetworkSubnetUsageParameter` has been removed
- Struct `VirtualNetworkSubnetUsageResult` has been removed
- Field `AdditionalInfo`, `Code`, `Details`, `Message`, `Target` of struct `ErrorResponse` has been removed

### Features Added

- New enum type `PrivateEndpointConnectionProvisioningState` with values `PrivateEndpointConnectionProvisioningStateCreating`, `PrivateEndpointConnectionProvisioningStateDeleting`, `PrivateEndpointConnectionProvisioningStateFailed`, `PrivateEndpointConnectionProvisioningStateSucceeded`
- New enum type `PrivateEndpointServiceConnectionStatus` with values `PrivateEndpointServiceConnectionStatusApproved`, `PrivateEndpointServiceConnectionStatusPending`, `PrivateEndpointServiceConnectionStatusRejected`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateEndpointConnectionsClient.BeginCreateOrUpdate(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientCreateOrUpdateResponse], error)`
- New function `*PrivateEndpointConnectionsClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error)`
- New function `*PrivateEndpointConnectionsClient.Get(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error)`
- New function `*PrivateEndpointConnectionsClient.ListByServer(context.Context, string, string, *PrivateEndpointConnectionsClientListByServerOptions) (PrivateEndpointConnectionsClientListByServerResponse, error)`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.Get(context.Context, string, string, string, *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error)`
- New function `*PrivateLinkResourcesClient.ListByServer(context.Context, string, string, *PrivateLinkResourcesClientListByServerOptions) (PrivateLinkResourcesClientListByServerResponse, error)`
- New struct `ErrorDetail`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResult`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkServiceConnectionState`
- New struct `TagsObject`
- New field `Error` in struct `ErrorResponse`
- New field `SystemData` in struct `Resource`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).