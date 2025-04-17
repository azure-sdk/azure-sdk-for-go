# Release History

## 1.0.0 (2025-04-17)
### Breaking Changes

- Type of `Identity.UserAssignedIdentities` has been changed from `map[string]*ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties` to `map[string]*UserAssignedIdentitiesProperties`
- `ModuleProvisioningStateCancelled` from enum `ModuleProvisioningState` has been removed
- Function `*ClientFactory.NewDscCompilationJobClient` has been removed
- Function `*ClientFactory.NewDscCompilationJobStreamClient` has been removed
- Function `NewDscCompilationJobClient` has been removed
- Function `*DscCompilationJobClient.BeginCreate` has been removed
- Function `*DscCompilationJobClient.Get` has been removed
- Function `*DscCompilationJobClient.GetStream` has been removed
- Function `*DscCompilationJobClient.NewListByAutomationAccountPager` has been removed
- Function `NewDscCompilationJobStreamClient` has been removed
- Function `*DscCompilationJobStreamClient.ListByJob` has been removed
- Struct `ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties` has been removed
- Struct `DscCompilationJob` has been removed
- Struct `DscCompilationJobCreateParameters` has been removed
- Struct `DscCompilationJobCreateProperties` has been removed
- Struct `DscCompilationJobListResult` has been removed
- Struct `DscCompilationJobProperties` has been removed
- Field `Value` of struct `DscConfigurationClientGetContentResponse` has been removed

### Features Added

- New value `ModuleProvisioningStateCanceled` added to enum type `ModuleProvisioningState`
- New value `RunbookTypeEnumPowerShell72`, `RunbookTypeEnumPython` added to enum type `RunbookTypeEnum`
- New enum type `PackageProvisioningState` with values `PackageProvisioningStateActivitiesStored`, `PackageProvisioningStateCanceled`, `PackageProvisioningStateConnectionTypeImported`, `PackageProvisioningStateContentDownloaded`, `PackageProvisioningStateContentRetrieved`, `PackageProvisioningStateContentStored`, `PackageProvisioningStateContentValidated`, `PackageProvisioningStateCreated`, `PackageProvisioningStateCreating`, `PackageProvisioningStateFailed`, `PackageProvisioningStateModuleDataStored`, `PackageProvisioningStateModuleImportRunbookComplete`, `PackageProvisioningStateRunningImportModuleRunbook`, `PackageProvisioningStateStartingImportModuleRunbook`, `PackageProvisioningStateSucceeded`, `PackageProvisioningStateUpdating`
- New function `*AccountClient.NewListDeletedRunbooksPager(string, string, *AccountClientListDeletedRunbooksOptions) *runtime.Pager[AccountClientListDeletedRunbooksResponse]`
- New function `*ClientFactory.NewPackageClient() *PackageClient`
- New function `*ClientFactory.NewPython3PackageClient() *Python3PackageClient`
- New function `*ClientFactory.NewRuntimeEnvironmentsClient() *RuntimeEnvironmentsClient`
- New function `*HybridRunbookWorkersClient.Patch(context.Context, string, string, string, string, *HybridRunbookWorkersClientPatchOptions) (HybridRunbookWorkersClientPatchResponse, error)`
- New function `NewPackageClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PackageClient, error)`
- New function `*PackageClient.CreateOrUpdate(context.Context, string, string, string, string, PackageCreateOrUpdateParameters, *PackageClientCreateOrUpdateOptions) (PackageClientCreateOrUpdateResponse, error)`
- New function `*PackageClient.Delete(context.Context, string, string, string, string, *PackageClientDeleteOptions) (PackageClientDeleteResponse, error)`
- New function `*PackageClient.Get(context.Context, string, string, string, string, *PackageClientGetOptions) (PackageClientGetResponse, error)`
- New function `*PackageClient.NewListByRuntimeEnvironmentPager(string, string, string, *PackageClientListByRuntimeEnvironmentOptions) *runtime.Pager[PackageClientListByRuntimeEnvironmentResponse]`
- New function `*PackageClient.Update(context.Context, string, string, string, string, PackageUpdateParameters, *PackageClientUpdateOptions) (PackageClientUpdateResponse, error)`
- New function `NewPython3PackageClient(string, azcore.TokenCredential, *arm.ClientOptions) (*Python3PackageClient, error)`
- New function `*Python3PackageClient.CreateOrUpdate(context.Context, string, string, string, PythonPackageCreateParameters, *Python3PackageClientCreateOrUpdateOptions) (Python3PackageClientCreateOrUpdateResponse, error)`
- New function `*Python3PackageClient.Delete(context.Context, string, string, string, *Python3PackageClientDeleteOptions) (Python3PackageClientDeleteResponse, error)`
- New function `*Python3PackageClient.Get(context.Context, string, string, string, *Python3PackageClientGetOptions) (Python3PackageClientGetResponse, error)`
- New function `*Python3PackageClient.NewListByAutomationAccountPager(string, string, *Python3PackageClientListByAutomationAccountOptions) *runtime.Pager[Python3PackageClientListByAutomationAccountResponse]`
- New function `*Python3PackageClient.Update(context.Context, string, string, string, PythonPackageUpdateParameters, *Python3PackageClientUpdateOptions) (Python3PackageClientUpdateResponse, error)`
- New function `NewRuntimeEnvironmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RuntimeEnvironmentsClient, error)`
- New function `*RuntimeEnvironmentsClient.Create(context.Context, string, string, string, RuntimeEnvironment, *RuntimeEnvironmentsClientCreateOptions) (RuntimeEnvironmentsClientCreateResponse, error)`
- New function `*RuntimeEnvironmentsClient.Delete(context.Context, string, string, string, *RuntimeEnvironmentsClientDeleteOptions) (RuntimeEnvironmentsClientDeleteResponse, error)`
- New function `*RuntimeEnvironmentsClient.Get(context.Context, string, string, string, *RuntimeEnvironmentsClientGetOptions) (RuntimeEnvironmentsClientGetResponse, error)`
- New function `*RuntimeEnvironmentsClient.NewListByAutomationAccountPager(string, string, *RuntimeEnvironmentsClientListByAutomationAccountOptions) *runtime.Pager[RuntimeEnvironmentsClientListByAutomationAccountResponse]`
- New function `*RuntimeEnvironmentsClient.Update(context.Context, string, string, string, RuntimeEnvironmentUpdateParameters, *RuntimeEnvironmentsClientUpdateOptions) (RuntimeEnvironmentsClientUpdateResponse, error)`
- New struct `DeletedRunbook`
- New struct `DeletedRunbookListResult`
- New struct `DeletedRunbookProperties`
- New struct `Dimension`
- New struct `JobRuntimeEnvironment`
- New struct `LogSpecification`
- New struct `MetricSpecification`
- New struct `OperationPropertiesFormat`
- New struct `OperationPropertiesFormatServiceSpecification`
- New struct `Package`
- New struct `PackageCreateOrUpdateParameters`
- New struct `PackageCreateOrUpdateProperties`
- New struct `PackageErrorInfo`
- New struct `PackageListResult`
- New struct `PackageProperties`
- New struct `PackageUpdateParameters`
- New struct `PackageUpdateProperties`
- New struct `RuntimeEnvironment`
- New struct `RuntimeEnvironmentListResult`
- New struct `RuntimeEnvironmentProperties`
- New struct `RuntimeEnvironmentUpdateParameters`
- New struct `RuntimeEnvironmentUpdateProperties`
- New struct `RuntimeProperties`
- New struct `UserAssignedIdentitiesProperties`
- New field `SystemData` in struct `Certificate`
- New field `SystemData` in struct `Connection`
- New field `SystemData` in struct `Credential`
- New field `SystemData` in struct `DscConfiguration`
- New field `SystemData` in struct `DscNode`
- New field `SystemData` in struct `DscNodeConfiguration`
- New field `Location`, `Tags` in struct `HybridRunbookWorker`
- New field `Location`, `Tags` in struct `HybridRunbookWorkerGroup`
- New field `SystemData` in struct `Job`
- New field `SystemData` in struct `JobCollectionItem`
- New field `JobRuntimeEnvironment`, `StartedBy` in struct `JobCollectionItemProperties`
- New field `JobRuntimeEnvironment` in struct `JobProperties`
- New field `SystemData` in struct `Module`
- New field `Origin`, `Properties` in struct `Operation`
- New field `Description` in struct `OperationDisplay`
- New field `SystemData` in struct `PrivateEndpointConnection`
- New field `SystemData` in struct `PrivateLinkResource`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `Runbook`
- New field `RuntimeEnvironment` in struct `RunbookCreateOrUpdateDraftProperties`
- New field `RuntimeEnvironment` in struct `RunbookCreateOrUpdateProperties`
- New field `RuntimeEnvironment` in struct `RunbookProperties`
- New field `SystemData` in struct `Schedule`
- New field `SystemData` in struct `SourceControl`
- New field `RuntimeEnvironment` in struct `TestJobCreateParameters`
- New field `SystemData` in struct `TrackedResource`
- New field `SystemData` in struct `Variable`
- New field `SystemData` in struct `Watcher`
- New field `SystemData` in struct `Webhook`


## 0.9.0 (2023-11-30)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.8.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.8.0 (2023-04-03)
### Breaking Changes

- Function `*DscConfigurationClient.UpdateWithJSON` parameter(s) have been changed from `(context.Context, string, string, string, DscConfigurationUpdateParameters, *DscConfigurationClientUpdateWithJSONOptions)` to `(context.Context, string, string, string, *DscConfigurationClientUpdateWithJSONOptions)`
- Function `*DscConfigurationClient.UpdateWithText` parameter(s) have been changed from `(context.Context, string, string, string, string, *DscConfigurationClientUpdateWithTextOptions)` to `(context.Context, string, string, string, *DscConfigurationClientUpdateWithTextOptions)`

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New field `Parameters` in struct `DscConfigurationClientUpdateWithJSONOptions`
- New field `Parameters` in struct `DscConfigurationClientUpdateWithTextOptions`


## 0.7.0 (2022-07-12)
### Breaking Changes

- Function `*DscConfigurationClient.UpdateWithJSON` parameter(s) have been changed from `(context.Context, string, string, string, *DscConfigurationClientUpdateWithJSONOptions)` to `(context.Context, string, string, string, DscConfigurationUpdateParameters, *DscConfigurationClientUpdateWithJSONOptions)`
- Function `*DscConfigurationClient.UpdateWithText` parameter(s) have been changed from `(context.Context, string, string, string, *DscConfigurationClientUpdateWithTextOptions)` to `(context.Context, string, string, string, string, *DscConfigurationClientUpdateWithTextOptions)`
- Struct `HybridRunbookWorkerGroupUpdateParameters` has been removed
- Struct `HybridRunbookWorkerLegacy` has been removed
- Field `Parameters` of struct `DscConfigurationClientUpdateWithJSONOptions` has been removed
- Field `Credential` of struct `HybridRunbookWorkerGroup` has been removed
- Field `GroupType` of struct `HybridRunbookWorkerGroup` has been removed
- Field `HybridRunbookWorkers` of struct `HybridRunbookWorkerGroup` has been removed
- Field `Parameters` of struct `DscConfigurationClientUpdateWithTextOptions` has been removed
- Field `Credential` of struct `HybridRunbookWorkerGroupCreateOrUpdateParameters` has been removed

### Features Added

- New const `RunbookTypeEnumPython3`
- New const `RunbookTypeEnumPython2`
- New function `NewDeletedAutomationAccountsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DeletedAutomationAccountsClient, error)`
- New function `*DeletedAutomationAccountsClient.ListBySubscription(context.Context, *DeletedAutomationAccountsClientListBySubscriptionOptions) (DeletedAutomationAccountsClientListBySubscriptionResponse, error)`
- New struct `DeletedAutomationAccount`
- New struct `DeletedAutomationAccountListResult`
- New struct `DeletedAutomationAccountProperties`
- New struct `DeletedAutomationAccountsClient`
- New struct `DeletedAutomationAccountsClientListBySubscriptionOptions`
- New struct `DeletedAutomationAccountsClientListBySubscriptionResponse`
- New struct `HybridRunbookWorkerGroupCreateOrUpdateProperties`
- New struct `HybridRunbookWorkerGroupProperties`
- New field `Name` in struct `HybridRunbookWorkerGroupCreateOrUpdateParameters`
- New field `Properties` in struct `HybridRunbookWorkerGroupCreateOrUpdateParameters`
- New field `Properties` in struct `HybridRunbookWorkerGroup`


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/automation/armautomation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).