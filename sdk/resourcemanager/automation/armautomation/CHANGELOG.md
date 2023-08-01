# Release History

## 1.0.0 (2023-08-01)
### Breaking Changes

- Type of `Identity.UserAssignedIdentities` has been changed from `map[string]*ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties` to `map[string]*UserAssignedIdentitiesProperties`
- Enum `AgentRegistrationKeyName` has been removed
- Enum `CountType` has been removed
- Enum `LinuxUpdateClasses` has been removed
- Enum `OperatingSystemType` has been removed
- Enum `TagOperators` has been removed
- Enum `WindowsUpdateClasses` has been removed
- Function `NewAgentRegistrationInformationClient` has been removed
- Function `*AgentRegistrationInformationClient.Get` has been removed
- Function `*AgentRegistrationInformationClient.RegenerateKey` has been removed
- Function `*ClientFactory.NewAgentRegistrationInformationClient` has been removed
- Function `*ClientFactory.NewDeletedAutomationAccountsClient` has been removed
- Function `*ClientFactory.NewDscCompilationJobClient` has been removed
- Function `*ClientFactory.NewDscCompilationJobStreamClient` has been removed
- Function `*ClientFactory.NewDscNodeClient` has been removed
- Function `*ClientFactory.NewNodeCountInformationClient` has been removed
- Function `*ClientFactory.NewNodeReportsClient` has been removed
- Function `*ClientFactory.NewPrivateEndpointConnectionsClient` has been removed
- Function `*ClientFactory.NewPrivateLinkResourcesClient` has been removed
- Function `*ClientFactory.NewSoftwareUpdateConfigurationsClient` has been removed
- Function `*ClientFactory.NewWatcherClient` has been removed
- Function `*ClientFactory.NewWebhookClient` has been removed
- Function `NewDeletedAutomationAccountsClient` has been removed
- Function `*DeletedAutomationAccountsClient.ListBySubscription` has been removed
- Function `NewDscCompilationJobClient` has been removed
- Function `*DscCompilationJobClient.BeginCreate` has been removed
- Function `*DscCompilationJobClient.Get` has been removed
- Function `*DscCompilationJobClient.GetStream` has been removed
- Function `*DscCompilationJobClient.NewListByAutomationAccountPager` has been removed
- Function `NewDscCompilationJobStreamClient` has been removed
- Function `*DscCompilationJobStreamClient.ListByJob` has been removed
- Function `NewDscNodeClient` has been removed
- Function `*DscNodeClient.Delete` has been removed
- Function `*DscNodeClient.Get` has been removed
- Function `*DscNodeClient.NewListByAutomationAccountPager` has been removed
- Function `*DscNodeClient.Update` has been removed
- Function `NewPrivateEndpointConnectionsClient` has been removed
- Function `*PrivateEndpointConnectionsClient.BeginCreateOrUpdate` has been removed
- Function `*PrivateEndpointConnectionsClient.BeginDelete` has been removed
- Function `*PrivateEndpointConnectionsClient.Get` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByAutomationAccountPager` has been removed
- Function `NewPrivateLinkResourcesClient` has been removed
- Function `*PrivateLinkResourcesClient.NewAutomationPager` has been removed
- Function `NewSoftwareUpdateConfigurationsClient` has been removed
- Function `*SoftwareUpdateConfigurationsClient.Create` has been removed
- Function `*SoftwareUpdateConfigurationsClient.Delete` has been removed
- Function `*SoftwareUpdateConfigurationsClient.GetByName` has been removed
- Function `*SoftwareUpdateConfigurationsClient.List` has been removed
- Function `NewWatcherClient` has been removed
- Function `*WatcherClient.CreateOrUpdate` has been removed
- Function `*WatcherClient.Delete` has been removed
- Function `*WatcherClient.Get` has been removed
- Function `*WatcherClient.NewListByAutomationAccountPager` has been removed
- Function `*WatcherClient.Start` has been removed
- Function `*WatcherClient.Stop` has been removed
- Function `*WatcherClient.Update` has been removed
- Function `NewWebhookClient` has been removed
- Function `*WebhookClient.CreateOrUpdate` has been removed
- Function `*WebhookClient.Delete` has been removed
- Function `*WebhookClient.GenerateURI` has been removed
- Function `*WebhookClient.Get` has been removed
- Function `*WebhookClient.NewListByAutomationAccountPager` has been removed
- Function `*WebhookClient.Update` has been removed
- Function `NewNodeCountInformationClient` has been removed
- Function `*NodeCountInformationClient.Get` has been removed
- Function `NewNodeReportsClient` has been removed
- Function `*NodeReportsClient.Get` has been removed
- Function `*NodeReportsClient.GetContent` has been removed
- Function `*NodeReportsClient.NewListByNodePager` has been removed
- Struct `AgentRegistration` has been removed
- Struct `AgentRegistrationKeys` has been removed
- Struct `AgentRegistrationRegenerateKeyParameter` has been removed
- Struct `AzureQueryProperties` has been removed
- Struct `ComponentsSgqdofSchemasIdentityPropertiesUserassignedidentitiesAdditionalproperties` has been removed
- Struct `DeletedAutomationAccount` has been removed
- Struct `DeletedAutomationAccountListResult` has been removed
- Struct `DeletedAutomationAccountProperties` has been removed
- Struct `DscCompilationJob` has been removed
- Struct `DscCompilationJobCreateParameters` has been removed
- Struct `DscCompilationJobCreateProperties` has been removed
- Struct `DscCompilationJobListResult` has been removed
- Struct `DscCompilationJobProperties` has been removed
- Struct `DscMetaConfiguration` has been removed
- Struct `DscNode` has been removed
- Struct `DscNodeConfigurationAssociationProperty` has been removed
- Struct `DscNodeListResult` has been removed
- Struct `DscNodeProperties` has been removed
- Struct `DscNodeReport` has been removed
- Struct `DscNodeReportListResult` has been removed
- Struct `DscNodeUpdateParameters` has been removed
- Struct `DscNodeUpdateParametersProperties` has been removed
- Struct `DscReportError` has been removed
- Struct `DscReportResource` has been removed
- Struct `DscReportResourceNavigation` has been removed
- Struct `LinuxProperties` has been removed
- Struct `NodeCount` has been removed
- Struct `NodeCountProperties` has been removed
- Struct `NodeCounts` has been removed
- Struct `NonAzureQueryProperties` has been removed
- Struct `PrivateEndpointConnectionListResult` has been removed
- Struct `PrivateLinkResource` has been removed
- Struct `PrivateLinkResourceListResult` has been removed
- Struct `PrivateLinkResourceProperties` has been removed
- Struct `SUCScheduleProperties` has been removed
- Struct `SoftwareUpdateConfiguration` has been removed
- Struct `SoftwareUpdateConfigurationCollectionItem` has been removed
- Struct `SoftwareUpdateConfigurationCollectionItemProperties` has been removed
- Struct `SoftwareUpdateConfigurationListResult` has been removed
- Struct `SoftwareUpdateConfigurationProperties` has been removed
- Struct `SoftwareUpdateConfigurationTasks` has been removed
- Struct `TagSettingsProperties` has been removed
- Struct `TargetProperties` has been removed
- Struct `TaskProperties` has been removed
- Struct `UpdateConfiguration` has been removed
- Struct `Watcher` has been removed
- Struct `WatcherListResult` has been removed
- Struct `WatcherProperties` has been removed
- Struct `WatcherUpdateParameters` has been removed
- Struct `WatcherUpdateProperties` has been removed
- Struct `Webhook` has been removed
- Struct `WebhookCreateOrUpdateParameters` has been removed
- Struct `WebhookCreateOrUpdateProperties` has been removed
- Struct `WebhookListResult` has been removed
- Struct `WebhookProperties` has been removed
- Struct `WebhookUpdateParameters` has been removed
- Struct `WebhookUpdateProperties` has been removed
- Struct `WindowsProperties` has been removed
- Field `Value` of struct `DscConfigurationClientGetContentResponse` has been removed

### Features Added

- New function `*AccountClient.NewListDeletedRunbooksPager(string, string, *AccountClientListDeletedRunbooksOptions) *runtime.Pager[AccountClientListDeletedRunbooksResponse]`
- New function `*ClientFactory.NewPython3PackageClient() *Python3PackageClient`
- New function `NewPython3PackageClient(string, azcore.TokenCredential, *arm.ClientOptions) (*Python3PackageClient, error)`
- New function `*Python3PackageClient.CreateOrUpdate(context.Context, string, string, string, PythonPackageCreateParameters, *Python3PackageClientCreateOrUpdateOptions) (Python3PackageClientCreateOrUpdateResponse, error)`
- New function `*Python3PackageClient.Delete(context.Context, string, string, string, *Python3PackageClientDeleteOptions) (Python3PackageClientDeleteResponse, error)`
- New function `*Python3PackageClient.Get(context.Context, string, string, string, *Python3PackageClientGetOptions) (Python3PackageClientGetResponse, error)`
- New function `*Python3PackageClient.NewListByAutomationAccountPager(string, string, *Python3PackageClientListByAutomationAccountOptions) *runtime.Pager[Python3PackageClientListByAutomationAccountResponse]`
- New function `*Python3PackageClient.Update(context.Context, string, string, string, PythonPackageUpdateParameters, *Python3PackageClientUpdateOptions) (Python3PackageClientUpdateResponse, error)`
- New struct `DeletedRunbook`
- New struct `DeletedRunbookListResult`
- New struct `DeletedRunbookProperties`
- New struct `Dimension`
- New struct `LogSpecification`
- New struct `MetricSpecification`
- New struct `OperationPropertiesFormat`
- New struct `OperationPropertiesFormatServiceSpecification`
- New struct `UserAssignedIdentitiesProperties`
- New field `Origin`, `Properties` in struct `Operation`
- New field `Description` in struct `OperationDisplay`


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