# Release History

## 0.10.0 (2024-10-21)
### Breaking Changes

- Type of `Identity.UserAssignedIdentities` has been changed from `map[string]*ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties` to `map[string]*UserAssignedIdentitiesProperties`
- `ModuleProvisioningStateCancelled` from enum `ModuleProvisioningState` has been removed
- Struct `ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties` has been removed
- Field `Value` of struct `DscConfigurationClientGetContentResponse` has been removed

### Features Added

- New value `ModuleProvisioningStateCanceled` added to enum type `ModuleProvisioningState`
- New value `RunbookTypeEnumPowerShell72` added to enum type `RunbookTypeEnum`
- New function `*ClientFactory.NewPowerShell72ModuleClient() *PowerShell72ModuleClient`
- New function `*ClientFactory.NewPython3PackageClient() *Python3PackageClient`
- New function `NewPowerShell72ModuleClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PowerShell72ModuleClient, error)`
- New function `*PowerShell72ModuleClient.CreateOrUpdate(context.Context, string, string, string, ModuleCreateOrUpdateParameters, *PowerShell72ModuleClientCreateOrUpdateOptions) (PowerShell72ModuleClientCreateOrUpdateResponse, error)`
- New function `*PowerShell72ModuleClient.Delete(context.Context, string, string, string, *PowerShell72ModuleClientDeleteOptions) (PowerShell72ModuleClientDeleteResponse, error)`
- New function `*PowerShell72ModuleClient.Get(context.Context, string, string, string, *PowerShell72ModuleClientGetOptions) (PowerShell72ModuleClientGetResponse, error)`
- New function `*PowerShell72ModuleClient.NewListByAutomationAccountPager(string, string, *PowerShell72ModuleClientListByAutomationAccountOptions) *runtime.Pager[PowerShell72ModuleClientListByAutomationAccountResponse]`
- New function `*PowerShell72ModuleClient.Update(context.Context, string, string, string, ModuleUpdateParameters, *PowerShell72ModuleClientUpdateOptions) (PowerShell72ModuleClientUpdateResponse, error)`
- New function `NewPython3PackageClient(string, azcore.TokenCredential, *arm.ClientOptions) (*Python3PackageClient, error)`
- New function `*Python3PackageClient.CreateOrUpdate(context.Context, string, string, string, PythonPackageCreateParameters, *Python3PackageClientCreateOrUpdateOptions) (Python3PackageClientCreateOrUpdateResponse, error)`
- New function `*Python3PackageClient.Delete(context.Context, string, string, string, *Python3PackageClientDeleteOptions) (Python3PackageClientDeleteResponse, error)`
- New function `*Python3PackageClient.Get(context.Context, string, string, string, *Python3PackageClientGetOptions) (Python3PackageClientGetResponse, error)`
- New function `*Python3PackageClient.NewListByAutomationAccountPager(string, string, *Python3PackageClientListByAutomationAccountOptions) *runtime.Pager[Python3PackageClientListByAutomationAccountResponse]`
- New function `*Python3PackageClient.Update(context.Context, string, string, string, PythonPackageUpdateParameters, *Python3PackageClientUpdateOptions) (Python3PackageClientUpdateResponse, error)`
- New struct `Dimension`
- New struct `LogSpecification`
- New struct `MetricSpecification`
- New struct `OperationPropertiesFormat`
- New struct `OperationPropertiesFormatServiceSpecification`
- New struct `UserAssignedIdentitiesProperties`
- New field `Origin`, `Properties` in struct `Operation`
- New field `Description` in struct `OperationDisplay`


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