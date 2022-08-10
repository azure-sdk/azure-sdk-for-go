# Release History

## 2.0.0 (2022-08-10)
### Breaking Changes

- Function `*TablesClient.Update` has been removed
- Struct `TablesClientUpdateOptions` has been removed
- Field `Etag` of struct `Workspace` has been removed

### Features Added

- New const `IdentityTypeKey`
- New const `IdentityTypeManagedIdentity`
- New const `ColumnTypeEnumString`
- New const `SourceEnumMicrosoft`
- New const `CreatedByTypeManagedIdentity`
- New const `ColumnTypeEnumDynamic`
- New const `TableSubTypeEnumClassic`
- New const `TablePlanEnumBasic`
- New const `ColumnTypeEnumDateTime`
- New const `SourceEnumCustomer`
- New const `ColumnDataTypeHintEnumURI`
- New const `TablePlanEnumAnalytics`
- New const `ColumnTypeEnumBoolean`
- New const `ColumnDataTypeHintEnumIP`
- New const `TableTypeEnumRestoredLogs`
- New const `IdentityTypeApplication`
- New const `TableSubTypeEnumAny`
- New const `TableTypeEnumCustomLog`
- New const `CreatedByTypeApplication`
- New const `ColumnTypeEnumReal`
- New const `ColumnDataTypeHintEnumGUID`
- New const `ColumnTypeEnumInt`
- New const `TableTypeEnumMicrosoft`
- New const `ColumnTypeEnumLong`
- New const `ColumnTypeEnumGUID`
- New const `ColumnDataTypeHintEnumArmPath`
- New const `IdentityTypeUser`
- New const `ProvisioningStateEnumInProgress`
- New const `DataSourceTypeIngestion`
- New const `TableTypeEnumSearchResults`
- New const `CreatedByTypeKey`
- New const `ProvisioningStateEnumUpdating`
- New const `CreatedByTypeUser`
- New const `TableSubTypeEnumDataCollectionRuleBased`
- New const `ProvisioningStateEnumSucceeded`
- New function `*QueryPacksClient.NewListPager(*QueryPacksClientListOptions) *runtime.Pager[QueryPacksClientListResponse]`
- New function `*TablesClient.BeginDelete(context.Context, string, string, string, *TablesClientBeginDeleteOptions) (*runtime.Poller[TablesClientDeleteResponse], error)`
- New function `PossibleTableSubTypeEnumValues() []TableSubTypeEnum`
- New function `*QueriesClient.Get(context.Context, string, string, string, *QueriesClientGetOptions) (QueriesClientGetResponse, error)`
- New function `PossibleColumnDataTypeHintEnumValues() []ColumnDataTypeHintEnum`
- New function `*TablesClient.BeginUpdate(context.Context, string, string, string, Table, *TablesClientBeginUpdateOptions) (*runtime.Poller[TablesClientUpdateResponse], error)`
- New function `*TablesClient.BeginCreateOrUpdate(context.Context, string, string, string, Table, *TablesClientBeginCreateOrUpdateOptions) (*runtime.Poller[TablesClientCreateOrUpdateResponse], error)`
- New function `*QueriesClient.Delete(context.Context, string, string, string, *QueriesClientDeleteOptions) (QueriesClientDeleteResponse, error)`
- New function `*QueriesClient.NewListPager(string, string, *QueriesClientListOptions) *runtime.Pager[QueriesClientListResponse]`
- New function `*QueriesClient.NewSearchPager(string, string, LogAnalyticsQueryPackQuerySearchProperties, *QueriesClientSearchOptions) *runtime.Pager[QueriesClientSearchResponse]`
- New function `*QueryPacksClient.Get(context.Context, string, string, *QueryPacksClientGetOptions) (QueryPacksClientGetResponse, error)`
- New function `*QueryPacksClient.Delete(context.Context, string, string, *QueryPacksClientDeleteOptions) (QueryPacksClientDeleteResponse, error)`
- New function `PossibleTableTypeEnumValues() []TableTypeEnum`
- New function `*QueryPacksClient.NewListByResourceGroupPager(string, *QueryPacksClientListByResourceGroupOptions) *runtime.Pager[QueryPacksClientListByResourceGroupResponse]`
- New function `PossibleSourceEnumValues() []SourceEnum`
- New function `*QueryPacksClient.CreateOrUpdate(context.Context, string, string, LogAnalyticsQueryPack, *QueryPacksClientCreateOrUpdateOptions) (QueryPacksClientCreateOrUpdateResponse, error)`
- New function `*QueryPacksClient.UpdateTags(context.Context, string, string, TagsResource, *QueryPacksClientUpdateTagsOptions) (QueryPacksClientUpdateTagsResponse, error)`
- New function `PossibleColumnTypeEnumValues() []ColumnTypeEnum`
- New function `*QueriesClient.Update(context.Context, string, string, string, LogAnalyticsQueryPackQuery, *QueriesClientUpdateOptions) (QueriesClientUpdateResponse, error)`
- New function `*QueriesClient.Put(context.Context, string, string, string, LogAnalyticsQueryPackQuery, *QueriesClientPutOptions) (QueriesClientPutResponse, error)`
- New function `*TablesClient.Migrate(context.Context, string, string, string, *TablesClientMigrateOptions) (TablesClientMigrateResponse, error)`
- New function `PossibleCreatedByTypeValues() []CreatedByType`
- New function `NewQueriesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*QueriesClient, error)`
- New function `NewQueryPacksClient(string, azcore.TokenCredential, *arm.ClientOptions) (*QueryPacksClient, error)`
- New function `PossibleTablePlanEnumValues() []TablePlanEnum`
- New function `*TablesClient.CancelSearch(context.Context, string, string, string, *TablesClientCancelSearchOptions) (TablesClientCancelSearchResponse, error)`
- New function `PossibleProvisioningStateEnumValues() []ProvisioningStateEnum`
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
- New struct `QueriesClient`
- New struct `QueriesClientDeleteOptions`
- New struct `QueriesClientDeleteResponse`
- New struct `QueriesClientGetOptions`
- New struct `QueriesClientGetResponse`
- New struct `QueriesClientListOptions`
- New struct `QueriesClientListResponse`
- New struct `QueriesClientPutOptions`
- New struct `QueriesClientPutResponse`
- New struct `QueriesClientSearchOptions`
- New struct `QueriesClientSearchResponse`
- New struct `QueriesClientUpdateOptions`
- New struct `QueriesClientUpdateResponse`
- New struct `QueryPacksClient`
- New struct `QueryPacksClientCreateOrUpdateOptions`
- New struct `QueryPacksClientCreateOrUpdateResponse`
- New struct `QueryPacksClientDeleteOptions`
- New struct `QueryPacksClientDeleteResponse`
- New struct `QueryPacksClientGetOptions`
- New struct `QueryPacksClientGetResponse`
- New struct `QueryPacksClientListByResourceGroupOptions`
- New struct `QueryPacksClientListByResourceGroupResponse`
- New struct `QueryPacksClientListOptions`
- New struct `QueryPacksClientListResponse`
- New struct `QueryPacksClientUpdateTagsOptions`
- New struct `QueryPacksClientUpdateTagsResponse`
- New struct `QueryPacksResource`
- New struct `RestoredLogs`
- New struct `ResultStatistics`
- New struct `Schema`
- New struct `SearchResults`
- New struct `SystemData`
- New struct `SystemDataAutoGenerated`
- New struct `TablesClientBeginCreateOrUpdateOptions`
- New struct `TablesClientBeginDeleteOptions`
- New struct `TablesClientBeginUpdateOptions`
- New struct `TablesClientCancelSearchOptions`
- New struct `TablesClientCancelSearchResponse`
- New struct `TablesClientCreateOrUpdateResponse`
- New struct `TablesClientDeleteResponse`
- New struct `TablesClientMigrateOptions`
- New struct `TablesClientMigrateResponse`
- New struct `TagsResource`
- New field `DefaultDataCollectionRuleResourceID` in struct `WorkspaceProperties`
- New field `SystemData` in struct `Table`
- New field `TotalRetentionInDaysAsDefault` in struct `TableProperties`
- New field `TotalRetentionInDays` in struct `TableProperties`
- New field `ArchiveRetentionInDays` in struct `TableProperties`
- New field `SearchResults` in struct `TableProperties`
- New field `Plan` in struct `TableProperties`
- New field `ResultStatistics` in struct `TableProperties`
- New field `LastPlanModifiedDate` in struct `TableProperties`
- New field `ProvisioningState` in struct `TableProperties`
- New field `RetentionInDaysAsDefault` in struct `TableProperties`
- New field `RestoredLogs` in struct `TableProperties`
- New field `Schema` in struct `TableProperties`
- New field `ETag` in struct `Workspace`
- New field `Identity` in struct `Workspace`
- New field `SystemData` in struct `Workspace`
- New field `Identity` in struct `WorkspacePatch`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).