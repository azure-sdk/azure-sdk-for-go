# Release History

## 0.8.0 (2022-09-05)
### Breaking Changes

- Function `*DscConfigurationClient.UpdateWithJSON` parameter(s) have been changed from `(context.Context, string, string, string, DscConfigurationUpdateParameters, *DscConfigurationClientUpdateWithJSONOptions)` to `(context.Context, string, string, string, *DscConfigurationClientUpdateWithJSONOptions)`
- Function `*DscConfigurationClient.UpdateWithText` parameter(s) have been changed from `(context.Context, string, string, string, string, *DscConfigurationClientUpdateWithTextOptions)` to `(context.Context, string, string, string, *DscConfigurationClientUpdateWithTextOptions)`
- Struct `RunbookCreateOrUpdateDraftProperties` has been removed

### Features Added

- New const `RuntimeLanguagePython`
- New const `RuntimeLanguagePowerShell`
- New const `RuntimeLanguagePowerShellWorkflow`
- New const `RuntimeLanguageGraphPowerShellWorkflow`
- New const `RuntimeLanguageGraphPowerShell`
- New const `RunbookTypeEnumPython`
- New type alias `RuntimeLanguage`
- New function `*RuntimeClient.Delete(context.Context, string, string, string, *RuntimeClientDeleteOptions) (RuntimeClientDeleteResponse, error)`
- New function `NewRuntimeClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RuntimeClient, error)`
- New function `*RuntimeClient.NewListByAutomationAccountPager(string, string, *RuntimeClientListByAutomationAccountOptions) *runtime.Pager[RuntimeClientListByAutomationAccountResponse]`
- New function `*RuntimeClient.Update(context.Context, string, string, string, RuntimeUpdateProperties, *RuntimeClientUpdateOptions) (RuntimeClientUpdateResponse, error)`
- New function `*RuntimeClient.Get(context.Context, string, string, string, *RuntimeClientGetOptions) (RuntimeClientGetResponse, error)`
- New function `*RuntimeClient.Create(context.Context, string, string, string, Runtime, *RuntimeClientCreateOptions) (RuntimeClientCreateResponse, error)`
- New function `PossibleRuntimeLanguageValues() []RuntimeLanguage`
- New struct `Runtime`
- New struct `RuntimeClient`
- New struct `RuntimeClientCreateOptions`
- New struct `RuntimeClientCreateResponse`
- New struct `RuntimeClientDeleteOptions`
- New struct `RuntimeClientDeleteResponse`
- New struct `RuntimeClientGetOptions`
- New struct `RuntimeClientGetResponse`
- New struct `RuntimeClientListByAutomationAccountOptions`
- New struct `RuntimeClientListByAutomationAccountResponse`
- New struct `RuntimeClientUpdateOptions`
- New struct `RuntimeClientUpdateResponse`
- New struct `RuntimeListResult`
- New struct `RuntimeProperties`
- New struct `RuntimeUpdateParameters`
- New struct `RuntimeUpdateProperties`
- New field `Runtime` in struct `RunbookCreateOrUpdateProperties`
- New field `Runtime` in struct `RunbookProperties`
- New field `Parameters` in struct `DscConfigurationClientUpdateWithJSONOptions`
- New field `Runtime` in struct `RunbookUpdateProperties`
- New field `RunbookType` in struct `RunbookUpdateProperties`
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