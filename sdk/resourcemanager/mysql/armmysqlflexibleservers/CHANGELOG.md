# Release History

## 2.0.0-beta.1 (2023-03-22)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New enum type `AdministratorName` with values `AdministratorNameActiveDirectory`
- New enum type `AdministratorType` with values `AdministratorTypeActiveDirectory`
- New enum type `ResetAllToDefault` with values `ResetAllToDefaultFalse`, `ResetAllToDefaultTrue`
- New function `NewAzureADAdministratorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureADAdministratorsClient, error)`
- New function `*AzureADAdministratorsClient.BeginCreateOrUpdate(context.Context, string, string, AdministratorName, AzureADAdministrator, *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureADAdministratorsClientCreateOrUpdateResponse], error)`
- New function `*AzureADAdministratorsClient.BeginDelete(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[AzureADAdministratorsClientDeleteResponse], error)`
- New function `*AzureADAdministratorsClient.Get(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientGetOptions) (AzureADAdministratorsClientGetResponse, error)`
- New function `*AzureADAdministratorsClient.NewListByServerPager(string, string, *AzureADAdministratorsClientListByServerOptions) *runtime.Pager[AzureADAdministratorsClientListByServerResponse]`
- New function `*BackupsClient.Put(context.Context, string, string, string, *BackupsClientPutOptions) (BackupsClientPutResponse, error)`
- New function `NewCheckNameAvailabilityWithoutLocationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityWithoutLocationClient, error)`
- New function `*CheckNameAvailabilityWithoutLocationClient.Execute(context.Context, NameAvailabilityRequest, *CheckNameAvailabilityWithoutLocationClientExecuteOptions) (CheckNameAvailabilityWithoutLocationClientExecuteResponse, error)`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAzureADAdministratorsClient() *AzureADAdministratorsClient`
- New function `*ClientFactory.NewBackupsClient() *BackupsClient`
- New function `*ClientFactory.NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient`
- New function `*ClientFactory.NewCheckNameAvailabilityWithoutLocationClient() *CheckNameAvailabilityWithoutLocationClient`
- New function `*ClientFactory.NewCheckVirtualNetworkSubnetUsageClient() *CheckVirtualNetworkSubnetUsageClient`
- New function `*ClientFactory.NewConfigurationsClient() *ConfigurationsClient`
- New function `*ClientFactory.NewDatabasesClient() *DatabasesClient`
- New function `*ClientFactory.NewFirewallRulesClient() *FirewallRulesClient`
- New function `*ClientFactory.NewGetPrivateDNSZoneSuffixClient() *GetPrivateDNSZoneSuffixClient`
- New function `*ClientFactory.NewLocationBasedCapabilitiesClient() *LocationBasedCapabilitiesClient`
- New function `*ClientFactory.NewLogFilesClient() *LogFilesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewReplicasClient() *ReplicasClient`
- New function `*ClientFactory.NewServersClient() *ServersClient`
- New function `NewLogFilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LogFilesClient, error)`
- New function `*LogFilesClient.NewListByServerPager(string, string, *LogFilesClientListByServerOptions) *runtime.Pager[LogFilesClientListByServerResponse]`
- New struct `AdministratorListResult`
- New struct `AdministratorProperties`
- New struct `AzureADAdministrator`
- New struct `ClientFactory`
- New struct `LogFile`
- New struct `LogFileListResult`
- New struct `LogFileProperties`
- New field `ResetAllToDefault` in struct `ConfigurationListForBatchUpdate`
- New field `Keyword` in struct `ConfigurationsClientListByServerOptions`
- New field `Page` in struct `ConfigurationsClientListByServerOptions`
- New field `PageSize` in struct `ConfigurationsClientListByServerOptions`
- New field `Tags` in struct `ConfigurationsClientListByServerOptions`
- New field `Version` in struct `ServerPropertiesForUpdate`
- New field `AutoIoScaling` in struct `Storage`
- New field `Location` in struct `VirtualNetworkSubnetUsageResult`
- New field `SubscriptionID` in struct `VirtualNetworkSubnetUsageResult`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).