# Release History

## 1.3.0-beta.2 (2025-04-22)
### Breaking Changes

- Function `*CustomLocationsClient.NewListOperationsPager` has been removed
- Operation `*ResourceSyncRulesClient.BeginUpdate` has been changed to non-LRO, use `*ResourceSyncRulesClient.Update` instead.
- Struct `CustomLocationOperation` has been removed
- Struct `CustomLocationOperationValueDisplay` has been removed
- Struct `CustomLocationOperationsList` has been removed
- Struct `EnabledResourceTypesListResult` has been removed
- Field `EnabledResourceTypesListResult` of struct `CustomLocationsClientListEnabledResourceTypesResponse` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `NewOperationsClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationsClient, error)`
- New function `*OperationsClient.NewListPager(*OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse]`
- New struct `EnabledResourceTypeListResult`
- New struct `Operation`
- New struct `OperationDisplay`
- New struct `OperationListResult`
- New anonymous field `EnabledResourceTypeListResult` in struct `CustomLocationsClientListEnabledResourceTypesResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.3.0-beta.1 (2023-11-30)
### Features Added

- New function `*ClientFactory.NewResourceSyncRulesClient() *ResourceSyncRulesClient`
- New function `*CustomLocationsClient.FindTargetResourceGroup(context.Context, string, string, CustomLocationFindTargetResourceGroupProperties, *CustomLocationsClientFindTargetResourceGroupOptions) (CustomLocationsClientFindTargetResourceGroupResponse, error)`
- New function `NewResourceSyncRulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ResourceSyncRulesClient, error)`
- New function `*ResourceSyncRulesClient.BeginCreateOrUpdate(context.Context, string, string, string, ResourceSyncRule, *ResourceSyncRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientCreateOrUpdateResponse], error)`
- New function `*ResourceSyncRulesClient.Delete(context.Context, string, string, string, *ResourceSyncRulesClientDeleteOptions) (ResourceSyncRulesClientDeleteResponse, error)`
- New function `*ResourceSyncRulesClient.Get(context.Context, string, string, string, *ResourceSyncRulesClientGetOptions) (ResourceSyncRulesClientGetResponse, error)`
- New function `*ResourceSyncRulesClient.NewListByCustomLocationIDPager(string, string, *ResourceSyncRulesClientListByCustomLocationIDOptions) *runtime.Pager[ResourceSyncRulesClientListByCustomLocationIDResponse]`
- New function `*ResourceSyncRulesClient.BeginUpdate(context.Context, string, string, string, PatchableResourceSyncRule, *ResourceSyncRulesClientBeginUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientUpdateResponse], error)`
- New struct `CustomLocationFindTargetResourceGroupProperties`
- New struct `CustomLocationFindTargetResourceGroupResult`
- New struct `MatchExpressionsProperties`
- New struct `PatchableResourceSyncRule`
- New struct `ResourceSyncRule`
- New struct `ResourceSyncRuleListResult`
- New struct `ResourceSyncRuleProperties`
- New struct `ResourceSyncRulePropertiesSelector`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.2.0-beta.1 (2023-05-11)
### Features Added

- New function `*ClientFactory.NewResourceSyncRulesClient() *ResourceSyncRulesClient`
- New function `*CustomLocationsClient.FindTargetResourceGroup(context.Context, string, string, CustomLocationFindTargetResourceGroupProperties, *CustomLocationsClientFindTargetResourceGroupOptions) (CustomLocationsClientFindTargetResourceGroupResponse, error)`
- New function `NewResourceSyncRulesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ResourceSyncRulesClient, error)`
- New function `*ResourceSyncRulesClient.BeginCreateOrUpdate(context.Context, string, string, string, ResourceSyncRule, *ResourceSyncRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientCreateOrUpdateResponse], error)`
- New function `*ResourceSyncRulesClient.Delete(context.Context, string, string, string, *ResourceSyncRulesClientDeleteOptions) (ResourceSyncRulesClientDeleteResponse, error)`
- New function `*ResourceSyncRulesClient.Get(context.Context, string, string, string, *ResourceSyncRulesClientGetOptions) (ResourceSyncRulesClientGetResponse, error)`
- New function `*ResourceSyncRulesClient.NewListByCustomLocationIDPager(string, string, *ResourceSyncRulesClientListByCustomLocationIDOptions) *runtime.Pager[ResourceSyncRulesClientListByCustomLocationIDResponse]`
- New function `*ResourceSyncRulesClient.BeginUpdate(context.Context, string, string, string, PatchableResourceSyncRule, *ResourceSyncRulesClientBeginUpdateOptions) (*runtime.Poller[ResourceSyncRulesClientUpdateResponse], error)`
- New struct `CustomLocationFindTargetResourceGroupProperties`
- New struct `CustomLocationFindTargetResourceGroupResult`
- New struct `MatchExpressionsProperties`
- New struct `PatchableResourceSyncRule`
- New struct `ResourceSyncRule`
- New struct `ResourceSyncRuleListResult`
- New struct `ResourceSyncRuleProperties`
- New struct `ResourceSyncRulePropertiesSelector`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-04-07)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module

## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/extendedlocation/armextendedlocation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).