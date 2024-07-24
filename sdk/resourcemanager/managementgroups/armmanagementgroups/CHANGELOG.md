# Release History

## 2.0.0 (2024-07-24)
### Breaking Changes

- Function `*APIClient.CheckNameAvailability` parameter(s) have been changed from `(context.Context, CheckNameAvailabilityRequest, *APIClientCheckNameAvailabilityOptions)` to `(context.Context, PathsItdwrvProvidersMicrosoftManagementChecknameavailabilityPostRequestbodyContentApplicationJSONSchema, *APIClientCheckNameAvailabilityOptions)`
- Function `*Client.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, CreateManagementGroupRequest, *ClientBeginCreateOrUpdateOptions)` to `(context.Context, string, ManagementGroup, *ClientBeginCreateOrUpdateOptions)`
- Type of `ErrorResponse.Error` has been changed from `*ErrorDetails` to `*ErrorDetail`
- Type of `ManagementGroupListResult.Value` has been changed from `[]*ManagementGroupInfo` to `[]*ManagementGroup`
- Type of `Operation.Display` has been changed from `*OperationDisplayProperties` to `*OperationDisplay`
- Function `*Client.NewListPager` has been removed
- Function `*ClientFactory.NewEntitiesClient` has been removed
- Function `*ClientFactory.NewHierarchySettingsClient` has been removed
- Function `*ClientFactory.NewManagementGroupSubscriptionsClient` has been removed
- Function `NewEntitiesClient` has been removed
- Function `*EntitiesClient.NewListPager` has been removed
- Function `NewHierarchySettingsClient` has been removed
- Function `*HierarchySettingsClient.CreateOrUpdate` has been removed
- Function `*HierarchySettingsClient.Delete` has been removed
- Function `*HierarchySettingsClient.Get` has been removed
- Function `*HierarchySettingsClient.List` has been removed
- Function `*HierarchySettingsClient.Update` has been removed
- Function `NewManagementGroupSubscriptionsClient` has been removed
- Function `*ManagementGroupSubscriptionsClient.Create` has been removed
- Function `*ManagementGroupSubscriptionsClient.Delete` has been removed
- Function `*ManagementGroupSubscriptionsClient.GetSubscription` has been removed
- Function `*ManagementGroupSubscriptionsClient.NewGetSubscriptionsUnderManagementGroupPager` has been removed
- Struct `AzureAsyncOperationResults` has been removed
- Struct `CreateManagementGroupChildInfo` has been removed
- Struct `CreateManagementGroupDetails` has been removed
- Struct `CreateManagementGroupProperties` has been removed
- Struct `CreateManagementGroupRequest` has been removed
- Struct `CreateParentGroupInfo` has been removed
- Struct `EntityHierarchyItem` has been removed
- Struct `EntityHierarchyItemProperties` has been removed
- Struct `ErrorDetails` has been removed
- Struct `HierarchySettingsInfo` has been removed
- Struct `HierarchySettingsList` has been removed
- Struct `ListSubscriptionUnderManagementGroup` has been removed
- Struct `ManagementGroupInfo` has been removed
- Struct `ManagementGroupInfoProperties` has been removed
- Struct `OperationDisplayProperties` has been removed
- Struct `OperationResults` has been removed
- Field `AzureAsyncOperationResults` of struct `ClientDeleteResponse` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New function `*Client.NewListSettingsPager(string, *ClientListSettingsOptions) *runtime.Pager[ClientListSettingsResponse]`
- New function `*ClientFactory.NewEntitiesOperationsClient() *EntitiesOperationsClient`
- New function `*ClientFactory.NewHierarchySettingsOperationGroupClient() *HierarchySettingsOperationGroupClient`
- New function `*ClientFactory.NewSubscriptionUnderManagementGroupsClient() *SubscriptionUnderManagementGroupsClient`
- New function `NewEntitiesOperationsClient(azcore.TokenCredential, *arm.ClientOptions) (*EntitiesOperationsClient, error)`
- New function `*EntitiesOperationsClient.NewListPager(*EntitiesOperationsClientListOptions) *runtime.Pager[EntitiesOperationsClientListResponse]`
- New function `NewHierarchySettingsOperationGroupClient(azcore.TokenCredential, *arm.ClientOptions) (*HierarchySettingsOperationGroupClient, error)`
- New function `*HierarchySettingsOperationGroupClient.CreateOrUpdate(context.Context, string, HierarchySettings, *HierarchySettingsOperationGroupClientCreateOrUpdateOptions) (HierarchySettingsOperationGroupClientCreateOrUpdateResponse, error)`
- New function `*HierarchySettingsOperationGroupClient.Delete(context.Context, string, *HierarchySettingsOperationGroupClientDeleteOptions) (HierarchySettingsOperationGroupClientDeleteResponse, error)`
- New function `*HierarchySettingsOperationGroupClient.Get(context.Context, string, *HierarchySettingsOperationGroupClientGetOptions) (HierarchySettingsOperationGroupClientGetResponse, error)`
- New function `*HierarchySettingsOperationGroupClient.NewListPager(string, *HierarchySettingsOperationGroupClientListOptions) *runtime.Pager[HierarchySettingsOperationGroupClientListResponse]`
- New function `*HierarchySettingsOperationGroupClient.Update(context.Context, string, CreateOrUpdateSettingsRequest, *HierarchySettingsOperationGroupClientUpdateOptions) (HierarchySettingsOperationGroupClientUpdateResponse, error)`
- New function `NewSubscriptionUnderManagementGroupsClient(azcore.TokenCredential, *arm.ClientOptions) (*SubscriptionUnderManagementGroupsClient, error)`
- New function `*SubscriptionUnderManagementGroupsClient.Create(context.Context, string, string, *SubscriptionUnderManagementGroupsClientCreateOptions) (SubscriptionUnderManagementGroupsClientCreateResponse, error)`
- New function `*SubscriptionUnderManagementGroupsClient.Delete(context.Context, string, string, *SubscriptionUnderManagementGroupsClientDeleteOptions) (SubscriptionUnderManagementGroupsClientDeleteResponse, error)`
- New function `*SubscriptionUnderManagementGroupsClient.GetSubscription(context.Context, string, string, *SubscriptionUnderManagementGroupsClientGetSubscriptionOptions) (SubscriptionUnderManagementGroupsClientGetSubscriptionResponse, error)`
- New function `*SubscriptionUnderManagementGroupsClient.NewGetSubscriptionsUnderManagementGroupPager(string, *SubscriptionUnderManagementGroupsClientGetSubscriptionsUnderManagementGroupOptions) *runtime.Pager[SubscriptionUnderManagementGroupsClientGetSubscriptionsUnderManagementGroupResponse]`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `HierarchySettingsListResult`
- New struct `ManagementGroupOperationAcceptance`
- New struct `OperationDisplay`
- New struct `PathsItdwrvProvidersMicrosoftManagementChecknameavailabilityPostRequestbodyContentApplicationJSONSchema`
- New struct `ProxyResource`
- New struct `Resource`
- New struct `SubscriptionUnderManagementGroupListResult`
- New struct `SystemData`
- New anonymous field `ManagementGroupOperationAcceptance` in struct `ClientDeleteResponse`
- New field `SystemData` in struct `HierarchySettings`
- New field `SystemData` in struct `ManagementGroup`
- New field `ActionType`, `IsDataAction`, `Origin` in struct `Operation`
- New field `SystemData` in struct `SubscriptionUnderManagementGroup`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/managementgroups/armmanagementgroups` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).