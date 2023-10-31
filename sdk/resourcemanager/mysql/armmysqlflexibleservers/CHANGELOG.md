# Release History

## 2.0.0-beta.2 (2023-10-31)
### Breaking Changes

- Type of `BackupAndExportResponse.Error` has been changed from `*ErrorResponse` to `*ErrorDetail`
- Type of `Server.Identity` has been changed from `*Identity` to `*MySQLServerIdentity`
- Type of `Server.SKU` has been changed from `*SKU` to `*MySQLServerSKU`
- Type of `ServerForUpdate.Identity` has been changed from `*Identity` to `*MySQLServerIdentity`
- Type of `ServerForUpdate.SKU` has been changed from `*SKU` to `*MySQLServerSKU`
- `BackupFormatNone` from enum `BackupFormat` has been removed
- Enum `SKUTier` has been removed
- Struct `Identity` has been removed
- Struct `SKU` has been removed
- Field `AdditionalInfo`, `Code`, `Details`, `Message`, `Target` of struct `ErrorResponse` has been removed

### Features Added

- New value `BackupFormatRaw` added to enum type `BackupFormat`
- New enum type `AdvancedThreatProtectionName` with values `AdvancedThreatProtectionNameDefault`
- New enum type `AdvancedThreatProtectionProvisioningState` with values `AdvancedThreatProtectionProvisioningStateCanceled`, `AdvancedThreatProtectionProvisioningStateFailed`, `AdvancedThreatProtectionProvisioningStateSucceeded`, `AdvancedThreatProtectionProvisioningStateUpdating`
- New enum type `AdvancedThreatProtectionState` with values `AdvancedThreatProtectionStateDisabled`, `AdvancedThreatProtectionStateEnabled`
- New enum type `ImportSourceStorageType` with values `ImportSourceStorageTypeAzureBlob`
- New enum type `PrivateEndpointConnectionProvisioningState` with values `PrivateEndpointConnectionProvisioningStateCreating`, `PrivateEndpointConnectionProvisioningStateDeleting`, `PrivateEndpointConnectionProvisioningStateFailed`, `PrivateEndpointConnectionProvisioningStateSucceeded`
- New enum type `PrivateEndpointServiceConnectionStatus` with values `PrivateEndpointServiceConnectionStatusApproved`, `PrivateEndpointServiceConnectionStatusPending`, `PrivateEndpointServiceConnectionStatusRejected`
- New enum type `ServerSKUTier` with values `ServerSKUTierBurstable`, `ServerSKUTierGeneralPurpose`, `ServerSKUTierMemoryOptimized`
- New function `NewAdvancedThreatProtectionSettingsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AdvancedThreatProtectionSettingsClient, error)`
- New function `*AdvancedThreatProtectionSettingsClient.Get(context.Context, string, string, AdvancedThreatProtectionName, *AdvancedThreatProtectionSettingsClientGetOptions) (AdvancedThreatProtectionSettingsClientGetResponse, error)`
- New function `*AdvancedThreatProtectionSettingsClient.NewListPager(string, string, *AdvancedThreatProtectionSettingsClientListOptions) *runtime.Pager[AdvancedThreatProtectionSettingsClientListResponse]`
- New function `*AdvancedThreatProtectionSettingsClient.BeginUpdate(context.Context, string, string, AdvancedThreatProtectionName, AdvancedThreatProtectionForUpdate, *AdvancedThreatProtectionSettingsClientBeginUpdateOptions) (*runtime.Poller[AdvancedThreatProtectionSettingsClientUpdateResponse], error)`
- New function `*ClientFactory.NewAdvancedThreatProtectionSettingsClient() *AdvancedThreatProtectionSettingsClient`
- New function `*ClientFactory.NewLocationBasedCapabilitySetClient() *LocationBasedCapabilitySetClient`
- New function `*ClientFactory.NewOperationResultsClient() *OperationResultsClient`
- New function `*ClientFactory.NewServersMigrationClient() *ServersMigrationClient`
- New function `NewLocationBasedCapabilitySetClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LocationBasedCapabilitySetClient, error)`
- New function `*LocationBasedCapabilitySetClient.Get(context.Context, string, string, *LocationBasedCapabilitySetClientGetOptions) (LocationBasedCapabilitySetClientGetResponse, error)`
- New function `*LocationBasedCapabilitySetClient.NewListPager(string, *LocationBasedCapabilitySetClientListOptions) *runtime.Pager[LocationBasedCapabilitySetClientListResponse]`
- New function `NewOperationResultsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OperationResultsClient, error)`
- New function `*OperationResultsClient.Get(context.Context, string, string, *OperationResultsClientGetOptions) (OperationResultsClientGetResponse, error)`
- New function `NewServersMigrationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ServersMigrationClient, error)`
- New function `*ServersMigrationClient.BeginCutoverMigration(context.Context, string, string, *ServersMigrationClientBeginCutoverMigrationOptions) (*runtime.Poller[ServersMigrationClientCutoverMigrationResponse], error)`
- New struct `AdvancedThreatProtection`
- New struct `AdvancedThreatProtectionForUpdate`
- New struct `AdvancedThreatProtectionListResult`
- New struct `AdvancedThreatProtectionProperties`
- New struct `AdvancedThreatProtectionUpdateProperties`
- New struct `Capability`
- New struct `CapabilityPropertiesV2`
- New struct `CapabilitySetsList`
- New struct `ErrorDetail`
- New struct `ImportSourceProperties`
- New struct `MySQLServerIdentity`
- New struct `MySQLServerSKU`
- New struct `OperationStatusExtendedResult`
- New struct `OperationStatusResult`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateLinkServiceConnectionState`
- New struct `SKUCapabilityV2`
- New struct `ServerEditionCapabilityV2`
- New struct `ServerVersionCapabilityV2`
- New field `SystemData` in struct `BackupAndExportResponse`
- New field `Error` in struct `ErrorResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `ImportSourceProperties`, `PrivateEndpointConnections` in struct `ServerProperties`
- New field `Network` in struct `ServerPropertiesForUpdate`
- New field `SystemData` in struct `TrackedResource`


## 2.0.0-beta.1 (2023-05-26)
### Breaking Changes

- Type of `Identity.Type` has been changed from `*string` to `*ManagedServiceIdentityType`

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
- New function `*ConfigurationsClient.BeginCreateOrUpdate(context.Context, string, string, string, Configuration, *ConfigurationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ConfigurationsClientCreateOrUpdateResponse], error)`
- New function `*FullBackupStoreDetails.GetBackupStoreDetails() *BackupStoreDetails`
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
- New field `ResetAllToDefault` in struct `ConfigurationListForBatchUpdate`
- New field `CurrentValue`, `DocumentationLink` in struct `ConfigurationProperties`
- New field `Keyword`, `Page`, `PageSize`, `Tags` in struct `ConfigurationsClientListByServerOptions`
- New field `Version` in struct `ServerPropertiesForUpdate`
- New field `AutoIoScaling`, `LogOnDisk` in struct `Storage`
- New field `Location`, `SubscriptionID` in struct `VirtualNetworkSubnetUsageResult`


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