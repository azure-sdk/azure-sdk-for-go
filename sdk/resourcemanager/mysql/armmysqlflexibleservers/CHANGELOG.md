# Release History

## 2.0.0-beta.1 (2022-10-20)
### Breaking Changes

- Struct `CloudError` has been removed

### Features Added

- New const `AdministratorNameActiveDirectory`
- New const `AdministratorTypeActiveDirectory`
- New const `ResetAllToDefaultFalse`
- New const `ResetAllToDefaultTrue`
- New type alias `AdministratorName`
- New type alias `AdministratorType`
- New type alias `ResetAllToDefault`
- New function `*LogFilesClient.NewListByServerPager(string, string, *LogFilesClientListByServerOptions) *runtime.Pager[LogFilesClientListByServerResponse]`
- New function `NewCheckNameAvailabilityWithoutLocationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityWithoutLocationClient, error)`
- New function `*AzureADAdministratorsClient.Get(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientGetOptions) (AzureADAdministratorsClientGetResponse, error)`
- New function `PossibleResetAllToDefaultValues() []ResetAllToDefault`
- New function `NewLogFilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LogFilesClient, error)`
- New function `*BackupsClient.Put(context.Context, string, string, string, *BackupsClientPutOptions) (BackupsClientPutResponse, error)`
- New function `PossibleAdministratorNameValues() []AdministratorName`
- New function `*AzureADAdministratorsClient.BeginCreateOrUpdate(context.Context, string, string, AdministratorName, AzureADAdministrator, *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureADAdministratorsClientCreateOrUpdateResponse], error)`
- New function `*CheckNameAvailabilityWithoutLocationClient.Execute(context.Context, NameAvailabilityRequest, *CheckNameAvailabilityWithoutLocationClientExecuteOptions) (CheckNameAvailabilityWithoutLocationClientExecuteResponse, error)`
- New function `NewAzureADAdministratorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureADAdministratorsClient, error)`
- New function `*AzureADAdministratorsClient.BeginDelete(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[AzureADAdministratorsClientDeleteResponse], error)`
- New function `*AzureADAdministratorsClient.NewListByServerPager(string, string, *AzureADAdministratorsClientListByServerOptions) *runtime.Pager[AzureADAdministratorsClientListByServerResponse]`
- New function `PossibleAdministratorTypeValues() []AdministratorType`
- New struct `AdministratorListResult`
- New struct `AdministratorProperties`
- New struct `AzureADAdministrator`
- New struct `AzureADAdministratorsClient`
- New struct `AzureADAdministratorsClientBeginCreateOrUpdateOptions`
- New struct `AzureADAdministratorsClientBeginDeleteOptions`
- New struct `AzureADAdministratorsClientCreateOrUpdateResponse`
- New struct `AzureADAdministratorsClientDeleteResponse`
- New struct `AzureADAdministratorsClientGetOptions`
- New struct `AzureADAdministratorsClientGetResponse`
- New struct `AzureADAdministratorsClientListByServerOptions`
- New struct `AzureADAdministratorsClientListByServerResponse`
- New struct `BackupsClientPutOptions`
- New struct `BackupsClientPutResponse`
- New struct `CheckNameAvailabilityWithoutLocationClient`
- New struct `CheckNameAvailabilityWithoutLocationClientExecuteOptions`
- New struct `CheckNameAvailabilityWithoutLocationClientExecuteResponse`
- New struct `LogFile`
- New struct `LogFileListResult`
- New struct `LogFileProperties`
- New struct `LogFilesClient`
- New struct `LogFilesClientListByServerOptions`
- New struct `LogFilesClientListByServerResponse`
- New field `Version` in struct `ServerPropertiesForUpdate`
- New field `ResetAllToDefault` in struct `ConfigurationListForBatchUpdate`
- New field `Tags` in struct `ConfigurationsClientListByServerOptions`
- New field `Keyword` in struct `ConfigurationsClientListByServerOptions`
- New field `Page` in struct `ConfigurationsClientListByServerOptions`
- New field `PageSize` in struct `ConfigurationsClientListByServerOptions`
- New field `SubscriptionID` in struct `VirtualNetworkSubnetUsageResult`
- New field `Location` in struct `VirtualNetworkSubnetUsageResult`
- New field `AutoIoScaling` in struct `Storage`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).