# Release History

## 2.0.0-beta.1 (2023-02-03)
### Breaking Changes

- Struct `CloudError` has been removed
- Field `AdditionalInfo` of struct `ErrorResponse` has been removed
- Field `Code` of struct `ErrorResponse` has been removed
- Field `Details` of struct `ErrorResponse` has been removed
- Field `Message` of struct `ErrorResponse` has been removed
- Field `Target` of struct `ErrorResponse` has been removed

### Features Added

- New type alias `AdministratorName` with values `AdministratorNameActiveDirectory`
- New type alias `AdministratorType` with values `AdministratorTypeActiveDirectory`
- New type alias `PrivateEndpointConnectionProvisioningState` with values `PrivateEndpointConnectionProvisioningStateCreating`, `PrivateEndpointConnectionProvisioningStateDeleting`, `PrivateEndpointConnectionProvisioningStateFailed`, `PrivateEndpointConnectionProvisioningStateSucceeded`
- New type alias `PrivateEndpointServiceConnectionStatus` with values `PrivateEndpointServiceConnectionStatusApproved`, `PrivateEndpointServiceConnectionStatusPending`, `PrivateEndpointServiceConnectionStatusRejected`
- New type alias `ResetAllToDefault` with values `ResetAllToDefaultFalse`, `ResetAllToDefaultTrue`
- New function `NewAzureADAdministratorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureADAdministratorsClient, error)`
- New function `*AzureADAdministratorsClient.BeginCreateOrUpdate(context.Context, string, string, AdministratorName, AzureADAdministrator, *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureADAdministratorsClientCreateOrUpdateResponse], error)`
- New function `*AzureADAdministratorsClient.BeginDelete(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[AzureADAdministratorsClientDeleteResponse], error)`
- New function `*AzureADAdministratorsClient.Get(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientGetOptions) (AzureADAdministratorsClientGetResponse, error)`
- New function `*AzureADAdministratorsClient.NewListByServerPager(string, string, *AzureADAdministratorsClientListByServerOptions) *runtime.Pager[AzureADAdministratorsClientListByServerResponse]`
- New function `*BackupsClient.Put(context.Context, string, string, string, *BackupsClientPutOptions) (BackupsClientPutResponse, error)`
- New function `NewCheckNameAvailabilityWithoutLocationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityWithoutLocationClient, error)`
- New function `*CheckNameAvailabilityWithoutLocationClient.Execute(context.Context, NameAvailabilityRequest, *CheckNameAvailabilityWithoutLocationClientExecuteOptions) (CheckNameAvailabilityWithoutLocationClientExecuteResponse, error)`
- New function `NewLogFilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LogFilesClient, error)`
- New function `*LogFilesClient.NewListByServerPager(string, string, *LogFilesClientListByServerOptions) *runtime.Pager[LogFilesClientListByServerResponse]`
- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateEndpointConnectionsClient.BeginCreateOrUpdate(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PrivateEndpointConnectionsClientCreateOrUpdateResponse], error)`
- New function `*PrivateEndpointConnectionsClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error)`
- New function `*PrivateEndpointConnectionsClient.Get(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error)`
- New function `*PrivateEndpointConnectionsClient.ListByServer(context.Context, string, string, *PrivateEndpointConnectionsClientListByServerOptions) (PrivateEndpointConnectionsClientListByServerResponse, error)`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.Get(context.Context, string, string, string, *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error)`
- New function `*PrivateLinkResourcesClient.ListByServer(context.Context, string, string, *PrivateLinkResourcesClientListByServerOptions) (PrivateLinkResourcesClientListByServerResponse, error)`
- New struct `AdministratorListResult`
- New struct `AdministratorProperties`
- New struct `AzureADAdministrator`
- New struct `AzureADAdministratorsClient`
- New struct `AzureADAdministratorsClientCreateOrUpdateResponse`
- New struct `AzureADAdministratorsClientDeleteResponse`
- New struct `AzureADAdministratorsClientListByServerResponse`
- New struct `CheckNameAvailabilityWithoutLocationClient`
- New struct `ErrorDetail`
- New struct `LogFile`
- New struct `LogFileListResult`
- New struct `LogFileProperties`
- New struct `LogFilesClient`
- New struct `LogFilesClientListByServerResponse`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResult`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionsClient`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New struct `PrivateEndpointConnectionsClientDeleteResponse`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkResourcesClient`
- New struct `PrivateLinkServiceConnectionState`
- New struct `TagsObject`
- New field `ResetAllToDefault` in struct `ConfigurationListForBatchUpdate`
- New field `Keyword` in struct `ConfigurationsClientListByServerOptions`
- New field `Page` in struct `ConfigurationsClientListByServerOptions`
- New field `PageSize` in struct `ConfigurationsClientListByServerOptions`
- New field `Tags` in struct `ConfigurationsClientListByServerOptions`
- New field `Error` in struct `ErrorResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `Version` in struct `ServerPropertiesForUpdate`
- New field `AutoIoScaling` in struct `Storage`
- New field `SystemData` in struct `TrackedResource`
- New field `Location` in struct `VirtualNetworkSubnetUsageResult`
- New field `SubscriptionID` in struct `VirtualNetworkSubnetUsageResult`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).