# Release History

## 2.0.0-beta.1 (2023-04-07)
### Breaking Changes

- Type of `Identity.Type` has been changed from `*string` to `*ManagedServiceIdentityType`
- Function `*LocationBasedCapabilitiesClient.NewListPager` has been removed
- Struct `CapabilitiesListResult` has been removed
- Field `SupportedHAMode` of struct `CapabilityProperties` has been removed
- Field `Zone` of struct `CapabilityProperties` has been removed
- Field `SupportedServerVersions` of struct `ServerEditionCapability` has been removed
- Field `SupportedSKUs` of struct `ServerVersionCapability` has been removed

### Features Added

- New enum type `AdministratorName` with values `AdministratorNameActiveDirectory`
- New enum type `AdministratorType` with values `AdministratorTypeActiveDirectory`
- New enum type `BackupFormat` with values `BackupFormatCollatedFormat`, `BackupFormatNone`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeUserAssigned`
- New enum type `OperationStatus` with values `OperationStatusCancelInProgress`, `OperationStatusCanceled`, `OperationStatusFailed`, `OperationStatusInProgress`, `OperationStatusPending`, `OperationStatusSucceeded`
- New enum type `ResetAllToDefault` with values `ResetAllToDefaultFalse`, `ResetAllToDefaultTrue`
- New function `NewAzureADAdministratorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AzureADAdministratorsClient, error)`
- New function `*AzureADAdministratorsClient.BeginCreateOrUpdate(context.Context, string, string, AdministratorName, AzureADAdministrator, *AzureADAdministratorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AzureADAdministratorsClientCreateOrUpdateResponse], error)`
- New function `*AzureADAdministratorsClient.BeginDelete(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientBeginDeleteOptions) (*runtime.Poller[AzureADAdministratorsClientDeleteResponse], error)`
- New function `*AzureADAdministratorsClient.Get(context.Context, string, string, AdministratorName, *AzureADAdministratorsClientGetOptions) (AzureADAdministratorsClientGetResponse, error)`
- New function `*AzureADAdministratorsClient.NewListByServerPager(string, string, *AzureADAdministratorsClientListByServerOptions) *runtime.Pager[AzureADAdministratorsClientListByServerResponse]`
- New function `NewBackupAndExportClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BackupAndExportClient, error)`
- New function `*BackupAndExportClient.BeginCreate(context.Context, string, string, BackupAndExportRequest, *BackupAndExportClientBeginCreateOptions) (*runtime.Poller[BackupAndExportClientCreateResponse], error)`
- New function `*BackupAndExportClient.ValidateBackup(context.Context, string, string, *BackupAndExportClientValidateBackupOptions) (BackupAndExportClientValidateBackupResponse, error)`
- New function `*BackupStoreDetails.GetBackupStoreDetails() *BackupStoreDetails`
- New function `*BackupsClient.Put(context.Context, string, string, string, *BackupsClientPutOptions) (BackupsClientPutResponse, error)`
- New function `NewCheckNameAvailabilityWithoutLocationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityWithoutLocationClient, error)`
- New function `*CheckNameAvailabilityWithoutLocationClient.Execute(context.Context, NameAvailabilityRequest, *CheckNameAvailabilityWithoutLocationClientExecuteOptions) (CheckNameAvailabilityWithoutLocationClientExecuteResponse, error)`
- New function `*ClientFactory.NewAzureADAdministratorsClient() *AzureADAdministratorsClient`
- New function `*ClientFactory.NewBackupAndExportClient() *BackupAndExportClient`
- New function `*ClientFactory.NewCheckNameAvailabilityWithoutLocationClient() *CheckNameAvailabilityWithoutLocationClient`
- New function `*ClientFactory.NewLogFilesClient() *LogFilesClient`
- New function `*FullBackupStoreDetails.GetBackupStoreDetails() *BackupStoreDetails`
- New function `*LocationBasedCapabilitiesClient.Get(context.Context, string, *LocationBasedCapabilitiesClientGetOptions) (LocationBasedCapabilitiesClientGetResponse, error)`
- New function `NewLogFilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LogFilesClient, error)`
- New function `*LogFilesClient.NewListByServerPager(string, string, *LogFilesClientListByServerOptions) *runtime.Pager[LogFilesClientListByServerResponse]`
- New function `*ServersClient.BeginResetGtid(context.Context, string, string, ServerGtidSetParameter, *ServersClientBeginResetGtidOptions) (*runtime.Poller[ServersClientResetGtidResponse], error)`
- New struct `AdministratorListResult`
- New struct `AdministratorProperties`
- New struct `AzureADAdministrator`
- New struct `BackupAndExportRequest`
- New struct `BackupAndExportResponse`
- New struct `BackupAndExportResponseProperties`
- New struct `BackupRequestBase`
- New struct `BackupSettings`
- New struct `FullBackupStoreDetails`
- New struct `LogFile`
- New struct `LogFileListResult`
- New struct `LogFileProperties`
- New struct `ServerGtidSetParameter`
- New struct `ValidateBackupResponse`
- New struct `ValidateBackupResponseProperties`
- New field `SupportedServerVersions` in struct `CapabilityProperties`
- New field `ResetAllToDefault` in struct `ConfigurationListForBatchUpdate`
- New field `Keyword` in struct `ConfigurationsClientListByServerOptions`
- New field `Page` in struct `ConfigurationsClientListByServerOptions`
- New field `PageSize` in struct `ConfigurationsClientListByServerOptions`
- New field `Tags` in struct `ConfigurationsClientListByServerOptions`
- New field `SupportedHAMode` in struct `SKUCapability`
- New field `SupportedZones` in struct `SKUCapability`
- New field `DefaultSKU` in struct `ServerEditionCapability`
- New field `DefaultStorageSize` in struct `ServerEditionCapability`
- New field `SupportedSKUs` in struct `ServerEditionCapability`
- New field `Version` in struct `ServerPropertiesForUpdate`
- New field `AutoIoScaling` in struct `Storage`
- New field `Location` in struct `VirtualNetworkSubnetUsageResult`
- New field `SubscriptionID` in struct `VirtualNetworkSubnetUsageResult`


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mysql/armmysqlflexibleservers` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).