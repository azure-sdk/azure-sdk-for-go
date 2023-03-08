# Release History

## 2.0.0 (2023-03-08)
### Breaking Changes

- Type of `ExportExecutionListResult.Value` has been changed from `[]*ExportExecution` to `[]*ExportRun`
- Type of `ForecastDataset.Aggregation` has been changed from `map[string]*QueryAggregation` to `map[string]*ForecastAggregation`
- Type of `ForecastDataset.Configuration` has been changed from `*QueryDatasetConfiguration` to `*ForecastDatasetConfiguration`
- Type of `ForecastDataset.Filter` has been changed from `*QueryFilter` to `*ForecastFilter`
- Type of `ForecastDefinition.TimePeriod` has been changed from `*QueryTimePeriod` to `*ForecastTimePeriod`
- Type of `ForecastDefinition.Timeframe` has been changed from `*ForecastTimeframeType` to `*ForecastTimeframe`
- Type of `OperationListResult.Value` has been changed from `[]*Operation` to `[]*OperationForCostManagement`
- Type of `ReportConfigGrouping.Type` has been changed from `*ReportConfigColumnType` to `*QueryColumnType`
- Const `QueryColumnTypeTag` from type alias `QueryColumnType` has been removed
- Type alias `ForecastTimeframeType` has been removed
- Type alias `ReportConfigColumnType` has been removed
- Operation `*GenerateDetailedCostReportOperationResultsClient.Get` has been changed to LRO, use `*GenerateDetailedCostReportOperationResultsClient.BeginGet` instead.
- Struct `ExportExecution` has been removed
- Struct `ExportExecutionProperties` has been removed
- Field `QueryResult` of struct `ForecastClientExternalCloudProviderUsageResponse` has been removed
- Field `QueryResult` of struct `ForecastClientUsageResponse` has been removed
- Field `ID` of struct `Operation` has been removed
- Field `ETag` of struct `ProxyResource` has been removed
- Field `ETag` of struct `Resource` has been removed
- Field `Location` of struct `Resource` has been removed
- Field `SKU` of struct `Resource` has been removed
- Field `Tags` of struct `Resource` has been removed

### Features Added

- New value `QueryColumnTypeTagKey` added to type alias `QueryColumnType`
- New type alias `ActionType` with values `ActionTypeInternal`
- New type alias `BenefitKind` with values `BenefitKindIncludedQuantity`, `BenefitKindReservation`, `BenefitKindSavingsPlan`
- New type alias `CheckNameAvailabilityReason` with values `CheckNameAvailabilityReasonAlreadyExists`, `CheckNameAvailabilityReasonInvalid`
- New type alias `CostDetailsDataFormat` with values `CostDetailsDataFormatCSVCostDetailsDataFormat`
- New type alias `CostDetailsMetricType` with values `CostDetailsMetricTypeActualCostCostDetailsMetricType`, `CostDetailsMetricTypeAmortizedCostCostDetailsMetricType`
- New type alias `CostDetailsStatusType` with values `CostDetailsStatusTypeCompletedCostDetailsStatusType`, `CostDetailsStatusTypeFailedCostDetailsStatusType`, `CostDetailsStatusTypeNoDataFoundCostDetailsStatusType`
- New type alias `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New type alias `DaysOfWeek` with values `DaysOfWeekFriday`, `DaysOfWeekMonday`, `DaysOfWeekSaturday`, `DaysOfWeekSunday`, `DaysOfWeekThursday`, `DaysOfWeekTuesday`, `DaysOfWeekWednesday`
- New type alias `FileFormat` with values `FileFormatCSV`
- New type alias `ForecastOperatorType` with values `ForecastOperatorTypeIn`
- New type alias `ForecastTimeframe` with values `ForecastTimeframeCustom`
- New type alias `FunctionName` with values `FunctionNameCost`, `FunctionNameCostUSD`, `FunctionNamePreTaxCost`, `FunctionNamePreTaxCostUSD`
- New type alias `Grain` with values `GrainDaily`, `GrainHourly`, `GrainMonthly`
- New type alias `GrainParameter` with values `GrainParameterDaily`, `GrainParameterHourly`, `GrainParameterMonthly`
- New type alias `LookBackPeriod` with values `LookBackPeriodLast30Days`, `LookBackPeriodLast60Days`, `LookBackPeriodLast7Days`
- New type alias `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New type alias `ScheduleFrequency` with values `ScheduleFrequencyDaily`, `ScheduleFrequencyMonthly`, `ScheduleFrequencyWeekly`
- New type alias `ScheduledActionKind` with values `ScheduledActionKindEmail`, `ScheduledActionKindInsightAlert`
- New type alias `ScheduledActionStatus` with values `ScheduledActionStatusDisabled`, `ScheduledActionStatusEnabled`, `ScheduledActionStatusExpired`
- New type alias `Scope` with values `ScopeShared`, `ScopeSingle`
- New type alias `Term` with values `TermP1Y`, `TermP3Y`
- New type alias `WeeksOfMonth` with values `WeeksOfMonthFirst`, `WeeksOfMonthFourth`, `WeeksOfMonthLast`, `WeeksOfMonthSecond`, `WeeksOfMonthThird`
- New function `*BenefitRecommendationProperties.GetBenefitRecommendationProperties() *BenefitRecommendationProperties`
- New function `NewBenefitRecommendationsClient(azcore.TokenCredential, *arm.ClientOptions) (*BenefitRecommendationsClient, error)`
- New function `*BenefitRecommendationsClient.NewListPager(string, *BenefitRecommendationsClientListOptions) *runtime.Pager[BenefitRecommendationsClientListResponse]`
- New function `NewBenefitUtilizationSummariesClient(azcore.TokenCredential, *arm.ClientOptions) (*BenefitUtilizationSummariesClient, error)`
- New function `*BenefitUtilizationSummariesClient.NewListByBillingAccountIDPager(string, *BenefitUtilizationSummariesClientListByBillingAccountIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListByBillingAccountIDResponse]`
- New function `*BenefitUtilizationSummariesClient.NewListByBillingProfileIDPager(string, string, *BenefitUtilizationSummariesClientListByBillingProfileIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListByBillingProfileIDResponse]`
- New function `*BenefitUtilizationSummariesClient.NewListBySavingsPlanIDPager(string, string, *BenefitUtilizationSummariesClientListBySavingsPlanIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListBySavingsPlanIDResponse]`
- New function `*BenefitUtilizationSummariesClient.NewListBySavingsPlanOrderPager(string, *BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions) *runtime.Pager[BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse]`
- New function `*BenefitUtilizationSummary.GetBenefitUtilizationSummary() *BenefitUtilizationSummary`
- New function `NewGenerateCostDetailsReportClient(azcore.TokenCredential, *arm.ClientOptions) (*GenerateCostDetailsReportClient, error)`
- New function `*GenerateCostDetailsReportClient.BeginCreateOperation(context.Context, string, GenerateCostDetailsReportRequestDefinition, *GenerateCostDetailsReportClientBeginCreateOperationOptions) (*runtime.Poller[GenerateCostDetailsReportClientCreateOperationResponse], error)`
- New function `*GenerateCostDetailsReportClient.BeginGetOperationResults(context.Context, string, string, *GenerateCostDetailsReportClientBeginGetOperationResultsOptions) (*runtime.Poller[GenerateCostDetailsReportClientGetOperationResultsResponse], error)`
- New function `*IncludedQuantityUtilizationSummary.GetBenefitUtilizationSummary() *BenefitUtilizationSummary`
- New function `PossibleGrainValues() []Grain`
- New function `NewPriceSheetClient(azcore.TokenCredential, *arm.ClientOptions) (*PriceSheetClient, error)`
- New function `*PriceSheetClient.BeginDownload(context.Context, string, string, string, *PriceSheetClientBeginDownloadOptions) (*runtime.Poller[PriceSheetClientDownloadResponse], error)`
- New function `*PriceSheetClient.BeginDownloadByBillingProfile(context.Context, string, string, *PriceSheetClientBeginDownloadByBillingProfileOptions) (*runtime.Poller[PriceSheetClientDownloadByBillingProfileResponse], error)`
- New function `*SavingsPlanUtilizationSummary.GetBenefitUtilizationSummary() *BenefitUtilizationSummary`
- New function `NewScheduledActionsClient(azcore.TokenCredential, *arm.ClientOptions) (*ScheduledActionsClient, error)`
- New function `*ScheduledActionsClient.CheckNameAvailability(context.Context, CheckNameAvailabilityRequest, *ScheduledActionsClientCheckNameAvailabilityOptions) (ScheduledActionsClientCheckNameAvailabilityResponse, error)`
- New function `*ScheduledActionsClient.CheckNameAvailabilityByScope(context.Context, string, CheckNameAvailabilityRequest, *ScheduledActionsClientCheckNameAvailabilityByScopeOptions) (ScheduledActionsClientCheckNameAvailabilityByScopeResponse, error)`
- New function `*ScheduledActionsClient.CreateOrUpdate(context.Context, string, ScheduledAction, *ScheduledActionsClientCreateOrUpdateOptions) (ScheduledActionsClientCreateOrUpdateResponse, error)`
- New function `*ScheduledActionsClient.CreateOrUpdateByScope(context.Context, string, string, ScheduledAction, *ScheduledActionsClientCreateOrUpdateByScopeOptions) (ScheduledActionsClientCreateOrUpdateByScopeResponse, error)`
- New function `*ScheduledActionsClient.Delete(context.Context, string, *ScheduledActionsClientDeleteOptions) (ScheduledActionsClientDeleteResponse, error)`
- New function `*ScheduledActionsClient.DeleteByScope(context.Context, string, string, *ScheduledActionsClientDeleteByScopeOptions) (ScheduledActionsClientDeleteByScopeResponse, error)`
- New function `*ScheduledActionsClient.Get(context.Context, string, *ScheduledActionsClientGetOptions) (ScheduledActionsClientGetResponse, error)`
- New function `*ScheduledActionsClient.GetByScope(context.Context, string, string, *ScheduledActionsClientGetByScopeOptions) (ScheduledActionsClientGetByScopeResponse, error)`
- New function `*ScheduledActionsClient.NewListByScopePager(string, *ScheduledActionsClientListByScopeOptions) *runtime.Pager[ScheduledActionsClientListByScopeResponse]`
- New function `*ScheduledActionsClient.NewListPager(*ScheduledActionsClientListOptions) *runtime.Pager[ScheduledActionsClientListResponse]`
- New function `*ScheduledActionsClient.Run(context.Context, string, *ScheduledActionsClientRunOptions) (ScheduledActionsClientRunResponse, error)`
- New function `*ScheduledActionsClient.RunByScope(context.Context, string, string, *ScheduledActionsClientRunByScopeOptions) (ScheduledActionsClientRunByScopeResponse, error)`
- New function `*SharedScopeBenefitRecommendationProperties.GetBenefitRecommendationProperties() *BenefitRecommendationProperties`
- New function `*SingleScopeBenefitRecommendationProperties.GetBenefitRecommendationProperties() *BenefitRecommendationProperties`
- New struct `AllSavingsBenefitDetails`
- New struct `AllSavingsList`
- New struct `BenefitRecommendationModel`
- New struct `BenefitRecommendationsClient`
- New struct `BenefitRecommendationsListResult`
- New struct `BenefitResource`
- New struct `BenefitUtilizationSummariesClient`
- New struct `BenefitUtilizationSummariesListResult`
- New struct `BenefitUtilizationSummaryProperties`
- New struct `BlobInfo`
- New struct `CheckNameAvailabilityRequest`
- New struct `CheckNameAvailabilityResponse`
- New struct `CostDetailsOperationResults`
- New struct `CostDetailsTimePeriod`
- New struct `ErrorDetailsWithNestedDetails`
- New struct `ErrorResponseWithNestedDetails`
- New struct `ExportRun`
- New struct `ExportRunProperties`
- New struct `FileDestination`
- New struct `ForecastAggregation`
- New struct `ForecastColumn`
- New struct `ForecastComparisonExpression`
- New struct `ForecastDatasetConfiguration`
- New struct `ForecastFilter`
- New struct `ForecastProperties`
- New struct `ForecastResult`
- New struct `ForecastTimePeriod`
- New struct `GenerateCostDetailsReportClient`
- New struct `GenerateCostDetailsReportErrorResponse`
- New struct `GenerateCostDetailsReportRequestDefinition`
- New struct `IncludedQuantityUtilizationSummary`
- New struct `IncludedQuantityUtilizationSummaryProperties`
- New struct `NotificationProperties`
- New struct `OperationForCostManagement`
- New struct `PriceSheetClient`
- New struct `ProxyResourceForCostManagement`
- New struct `RecommendationUsageDetails`
- New struct `ReportManifest`
- New struct `RequestContext`
- New struct `ResourceForCostManagement`
- New struct `SavingsPlanUtilizationSummary`
- New struct `SavingsPlanUtilizationSummaryProperties`
- New struct `ScheduleProperties`
- New struct `ScheduledAction`
- New struct `ScheduledActionListResult`
- New struct `ScheduledActionProperties`
- New struct `ScheduledActionProxyResource`
- New struct `ScheduledActionsClient`
- New struct `SharedScopeBenefitRecommendationProperties`
- New struct `SingleScopeBenefitRecommendationProperties`
- New struct `SystemData`
- New field `ExpiryTime` in struct `DownloadURL`
- New anonymous field `ForecastResult` in struct `ForecastClientExternalCloudProviderUsageResponse`
- New anonymous field `ForecastResult` in struct `ForecastClientUsageResponse`
- New field `ActionType` in struct `Operation`
- New field `IsDataAction` in struct `Operation`
- New field `Origin` in struct `Operation`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).