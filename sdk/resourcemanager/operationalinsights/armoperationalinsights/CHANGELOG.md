# Release History

## 2.0.0 (2024-05-14)
### Breaking Changes

- Type of `Cluster.Identity` has been changed from `*Identity` to `*ManagedServiceIdentity`
- Type of `ClusterPatch.Identity` has been changed from `*Identity` to `*ManagedServiceIdentity`
- Operation `*TablesClient.Update` has been changed to LRO, use `*TablesClient.BeginUpdate` instead.

### Features Added

- New value `CapacityFiftyThousand`, `CapacityFourHundred`, `CapacityOneHundred`, `CapacityTenThousand`, `CapacityThreeHundred`, `CapacityTwentyFiveThousand`, `CapacityTwoHundred` added to enum type `Capacity`
- New value `CapacityReservationLevelFiftyThousand`, `CapacityReservationLevelTenThousand`, `CapacityReservationLevelTwentyFiveThousand` added to enum type `CapacityReservationLevel`
- New value `DataSourceTypeIngestion` added to enum type `DataSourceType`
- New value `IdentityTypeApplication`, `IdentityTypeKey`, `IdentityTypeManagedIdentity`, `IdentityTypeUser` added to enum type `IdentityType`
- New enum type `ColumnDataTypeHintEnum` with values `ColumnDataTypeHintEnumArmPath`, `ColumnDataTypeHintEnumGUID`, `ColumnDataTypeHintEnumIP`, `ColumnDataTypeHintEnumURI`
- New enum type `ColumnTypeEnum` with values `ColumnTypeEnumBoolean`, `ColumnTypeEnumDateTime`, `ColumnTypeEnumDynamic`, `ColumnTypeEnumGUID`, `ColumnTypeEnumInt`, `ColumnTypeEnumLong`, `ColumnTypeEnumReal`, `ColumnTypeEnumString`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `ProvisioningStateEnum` with values `ProvisioningStateEnumDeleting`, `ProvisioningStateEnumInProgress`, `ProvisioningStateEnumSucceeded`, `ProvisioningStateEnumUpdating`
- New enum type `SourceEnum` with values `SourceEnumCustomer`, `SourceEnumMicrosoft`
- New enum type `TablePlanEnum` with values `TablePlanEnumAnalytics`, `TablePlanEnumBasic`
- New enum type `TableSubTypeEnum` with values `TableSubTypeEnumAny`, `TableSubTypeEnumClassic`, `TableSubTypeEnumDataCollectionRuleBased`
- New enum type `TableTypeEnum` with values `TableTypeEnumCustomLog`, `TableTypeEnumMicrosoft`, `TableTypeEnumRestoredLogs`, `TableTypeEnumSearchResults`
- New function `*ClientFactory.NewQueriesClient() *QueriesClient`
- New function `*ClientFactory.NewQueryPacksClient() *QueryPacksClient`
- New function `NewQueriesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*QueriesClient, error)`
- New function `*QueriesClient.Delete(context.Context, string, string, string, *QueriesClientDeleteOptions) (QueriesClientDeleteResponse, error)`
- New function `*QueriesClient.Get(context.Context, string, string, string, *QueriesClientGetOptions) (QueriesClientGetResponse, error)`
- New function `*QueriesClient.NewListPager(string, string, *QueriesClientListOptions) *runtime.Pager[QueriesClientListResponse]`
- New function `*QueriesClient.Put(context.Context, string, string, string, LogAnalyticsQueryPackQuery, *QueriesClientPutOptions) (QueriesClientPutResponse, error)`
- New function `*QueriesClient.NewSearchPager(string, string, LogAnalyticsQueryPackQuerySearchProperties, *QueriesClientSearchOptions) *runtime.Pager[QueriesClientSearchResponse]`
- New function `*QueriesClient.Update(context.Context, string, string, string, LogAnalyticsQueryPackQuery, *QueriesClientUpdateOptions) (QueriesClientUpdateResponse, error)`
- New function `NewQueryPacksClient(string, azcore.TokenCredential, *arm.ClientOptions) (*QueryPacksClient, error)`
- New function `*QueryPacksClient.CreateOrUpdate(context.Context, string, string, LogAnalyticsQueryPack, *QueryPacksClientCreateOrUpdateOptions) (QueryPacksClientCreateOrUpdateResponse, error)`
- New function `*QueryPacksClient.CreateOrUpdateWithoutName(context.Context, string, LogAnalyticsQueryPack, *QueryPacksClientCreateOrUpdateWithoutNameOptions) (QueryPacksClientCreateOrUpdateWithoutNameResponse, error)`
- New function `*QueryPacksClient.Delete(context.Context, string, string, *QueryPacksClientDeleteOptions) (QueryPacksClientDeleteResponse, error)`
- New function `*QueryPacksClient.Get(context.Context, string, string, *QueryPacksClientGetOptions) (QueryPacksClientGetResponse, error)`
- New function `*QueryPacksClient.NewListByResourceGroupPager(string, *QueryPacksClientListByResourceGroupOptions) *runtime.Pager[QueryPacksClientListByResourceGroupResponse]`
- New function `*QueryPacksClient.NewListPager(*QueryPacksClientListOptions) *runtime.Pager[QueryPacksClientListResponse]`
- New function `*QueryPacksClient.UpdateTags(context.Context, string, string, TagsResource, *QueryPacksClientUpdateTagsOptions) (QueryPacksClientUpdateTagsResponse, error)`
- New function `*TablesClient.CancelSearch(context.Context, string, string, string, *TablesClientCancelSearchOptions) (TablesClientCancelSearchResponse, error)`
- New function `*TablesClient.BeginCreateOrUpdate(context.Context, string, string, string, Table, *TablesClientBeginCreateOrUpdateOptions) (*runtime.Poller[TablesClientCreateOrUpdateResponse], error)`
- New function `*TablesClient.BeginDelete(context.Context, string, string, string, *TablesClientBeginDeleteOptions) (*runtime.Poller[TablesClientDeleteResponse], error)`
- New function `*TablesClient.Migrate(context.Context, string, string, string, *TablesClientMigrateOptions) (TablesClientMigrateResponse, error)`
- New struct `AzureResourceProperties`
- New struct `Column`
- New struct `LogAnalyticsQueryPack`
- New struct `LogAnalyticsQueryPackListResult`
- New struct `LogAnalyticsQueryPackProperties`
- New struct `LogAnalyticsQueryPackQuery`
- New struct `LogAnalyticsQueryPackQueryListResult`
- New struct `LogAnalyticsQueryPackQueryProperties`
- New struct `LogAnalyticsQueryPackQueryPropertiesRelated`
- New struct `LogAnalyticsQueryPackQuerySearchProperties`
- New struct `LogAnalyticsQueryPackQuerySearchPropertiesRelated`
- New struct `ManagedServiceIdentity`
- New struct `QueryPacksResource`
- New struct `RestoredLogs`
- New struct `ResultStatistics`
- New struct `Schema`
- New struct `SearchResults`
- New struct `SystemData`
- New struct `SystemDataAutoGenerated`
- New struct `TagsResource`
- New struct `UserAssignedIdentity`
- New field `SystemData` in struct `Table`
- New field `ArchiveRetentionInDays`, `LastPlanModifiedDate`, `Plan`, `ProvisioningState`, `RestoredLogs`, `ResultStatistics`, `RetentionInDaysAsDefault`, `Schema`, `SearchResults`, `TotalRetentionInDays`, `TotalRetentionInDaysAsDefault` in struct `TableProperties`
- New field `Identity`, `SystemData` in struct `Workspace`
- New field `UnifiedSentinelBillingOnly` in struct `WorkspaceFeatures`
- New field `Identity` in struct `WorkspacePatch`
- New field `DefaultDataCollectionRuleResourceID` in struct `WorkspaceProperties`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-04-06)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module

## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).