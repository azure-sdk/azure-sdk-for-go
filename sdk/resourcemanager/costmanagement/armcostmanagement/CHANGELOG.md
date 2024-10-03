# Release History

## 3.0.0 (2024-10-03)
### Breaking Changes

- Function `*PriceSheetClient.BeginDownload` has been removed
- Field `DownloadURL` of struct `PriceSheetClientDownloadByBillingProfileResponse` has been removed

### Features Added

- New value `OperationStatusTypeComplete` added to enum type `OperationStatusType`
- New enum type `BenefitUtilizationSummaryReportSchema` with values `BenefitUtilizationSummaryReportSchemaAvgUtilizationPercentage`, `BenefitUtilizationSummaryReportSchemaBenefitID`, `BenefitUtilizationSummaryReportSchemaBenefitOrderID`, `BenefitUtilizationSummaryReportSchemaBenefitType`, `BenefitUtilizationSummaryReportSchemaKind`, `BenefitUtilizationSummaryReportSchemaMaxUtilizationPercentage`, `BenefitUtilizationSummaryReportSchemaMinUtilizationPercentage`, `BenefitUtilizationSummaryReportSchemaUsageDate`, `BenefitUtilizationSummaryReportSchemaUtilizedPercentage`
- New enum type `BudgetNotificationOperatorType` with values `BudgetNotificationOperatorTypeEqualTo`, `BudgetNotificationOperatorTypeGreaterThan`, `BudgetNotificationOperatorTypeGreaterThanOrEqualTo`, `BudgetNotificationOperatorTypeLessThan`
- New enum type `BudgetOperatorType` with values `BudgetOperatorTypeIn`
- New enum type `CategoryType` with values `CategoryTypeCost`, `CategoryTypeReservationUtilization`
- New enum type `CostAllocationPolicyType` with values `CostAllocationPolicyTypeFixedProportion`
- New enum type `CostAllocationResourceType` with values `CostAllocationResourceTypeDimension`, `CostAllocationResourceTypeTag`
- New enum type `CultureCode` with values `CultureCodeCsCz`, `CultureCodeDaDk`, `CultureCodeDeDe`, `CultureCodeEnGb`, `CultureCodeEnUs`, `CultureCodeEsEs`, `CultureCodeFrFr`, `CultureCodeHuHu`, `CultureCodeItIt`, `CultureCodeJaJp`, `CultureCodeKoKr`, `CultureCodeNbNo`, `CultureCodeNlNl`, `CultureCodePlPl`, `CultureCodePtBr`, `CultureCodePtPt`, `CultureCodeRuRu`, `CultureCodeSvSe`, `CultureCodeTrTr`, `CultureCodeZhCn`, `CultureCodeZhTw`
- New enum type `Frequency` with values `FrequencyDaily`, `FrequencyMonthly`, `FrequencyWeekly`
- New enum type `Reason` with values `ReasonAlreadyExists`, `ReasonInvalid`, `ReasonValid`
- New enum type `RuleStatus` with values `RuleStatusActive`, `RuleStatusNotActive`, `RuleStatusProcessing`
- New enum type `SettingType` with values `SettingTypeTaginheritance`
- New enum type `SettingsKind` with values `SettingsKindTaginheritance`
- New enum type `SystemAssignedServiceIdentityType` with values `SystemAssignedServiceIdentityTypeNone`, `SystemAssignedServiceIdentityTypeSystemAssigned`
- New enum type `ThresholdType` with values `ThresholdTypeActual`, `ThresholdTypeForecasted`
- New enum type `TimeGrainType` with values `TimeGrainTypeAnnually`, `TimeGrainTypeBillingAnnual`, `TimeGrainTypeBillingMonth`, `TimeGrainTypeBillingQuarter`, `TimeGrainTypeLast30Days`, `TimeGrainTypeLast7Days`, `TimeGrainTypeMonthly`, `TimeGrainTypeQuarterly`
- New function `NewBudgetsClient(azcore.TokenCredential, *arm.ClientOptions) (*BudgetsClient, error)`
- New function `*BudgetsClient.CreateOrUpdate(context.Context, string, string, Budget, *BudgetsClientCreateOrUpdateOptions) (BudgetsClientCreateOrUpdateResponse, error)`
- New function `*BudgetsClient.Delete(context.Context, string, string, *BudgetsClientDeleteOptions) (BudgetsClientDeleteResponse, error)`
- New function `*BudgetsClient.Get(context.Context, string, string, *BudgetsClientGetOptions) (BudgetsClientGetResponse, error)`
- New function `*BudgetsClient.NewListPager(string, *BudgetsClientListOptions) *runtime.Pager[BudgetsClientListResponse]`
- New function `*ClientFactory.NewBudgetsClient() *BudgetsClient`
- New function `*ClientFactory.NewCostAllocationRulesClient() *CostAllocationRulesClient`
- New function `*ClientFactory.NewGenerateBenefitUtilizationSummariesReportClient() *GenerateBenefitUtilizationSummariesReportClient`
- New function `*ClientFactory.NewSettingsClient() *SettingsClient`
- New function `NewCostAllocationRulesClient(azcore.TokenCredential, *arm.ClientOptions) (*CostAllocationRulesClient, error)`
- New function `*CostAllocationRulesClient.CheckNameAvailability(context.Context, string, CostAllocationRuleCheckNameAvailabilityRequest, *CostAllocationRulesClientCheckNameAvailabilityOptions) (CostAllocationRulesClientCheckNameAvailabilityResponse, error)`
- New function `*CostAllocationRulesClient.CreateOrUpdate(context.Context, string, string, CostAllocationRuleDefinition, *CostAllocationRulesClientCreateOrUpdateOptions) (CostAllocationRulesClientCreateOrUpdateResponse, error)`
- New function `*CostAllocationRulesClient.Delete(context.Context, string, string, *CostAllocationRulesClientDeleteOptions) (CostAllocationRulesClientDeleteResponse, error)`
- New function `*CostAllocationRulesClient.Get(context.Context, string, string, *CostAllocationRulesClientGetOptions) (CostAllocationRulesClientGetResponse, error)`
- New function `*CostAllocationRulesClient.NewListPager(string, *CostAllocationRulesClientListOptions) *runtime.Pager[CostAllocationRulesClientListResponse]`
- New function `NewGenerateBenefitUtilizationSummariesReportClient(azcore.TokenCredential, *arm.ClientOptions) (*GenerateBenefitUtilizationSummariesReportClient, error)`
- New function `*GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByBillingAccount(context.Context, string, BenefitUtilizationSummariesRequest, *GenerateBenefitUtilizationSummariesReportClientBeginGenerateByBillingAccountOptions) (*runtime.Poller[GenerateBenefitUtilizationSummariesReportClientGenerateByBillingAccountResponse], error)`
- New function `*GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByBillingProfile(context.Context, string, string, BenefitUtilizationSummariesRequest, *GenerateBenefitUtilizationSummariesReportClientBeginGenerateByBillingProfileOptions) (*runtime.Poller[GenerateBenefitUtilizationSummariesReportClientGenerateByBillingProfileResponse], error)`
- New function `*GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByReservationID(context.Context, string, string, BenefitUtilizationSummariesRequest, *GenerateBenefitUtilizationSummariesReportClientBeginGenerateByReservationIDOptions) (*runtime.Poller[GenerateBenefitUtilizationSummariesReportClientGenerateByReservationIDResponse], error)`
- New function `*GenerateBenefitUtilizationSummariesReportClient.BeginGenerateByReservationOrderID(context.Context, string, BenefitUtilizationSummariesRequest, *GenerateBenefitUtilizationSummariesReportClientBeginGenerateByReservationOrderIDOptions) (*runtime.Poller[GenerateBenefitUtilizationSummariesReportClientGenerateByReservationOrderIDResponse], error)`
- New function `*GenerateBenefitUtilizationSummariesReportClient.BeginGenerateBySavingsPlanID(context.Context, string, string, BenefitUtilizationSummariesRequest, *GenerateBenefitUtilizationSummariesReportClientBeginGenerateBySavingsPlanIDOptions) (*runtime.Poller[GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanIDResponse], error)`
- New function `*GenerateBenefitUtilizationSummariesReportClient.BeginGenerateBySavingsPlanOrderID(context.Context, string, BenefitUtilizationSummariesRequest, *GenerateBenefitUtilizationSummariesReportClientBeginGenerateBySavingsPlanOrderIDOptions) (*runtime.Poller[GenerateBenefitUtilizationSummariesReportClientGenerateBySavingsPlanOrderIDResponse], error)`
- New function `*PriceSheetClient.BeginDownloadByBillingAccount(context.Context, string, string, *PriceSheetClientBeginDownloadByBillingAccountOptions) (*runtime.Poller[PriceSheetClientDownloadByBillingAccountResponse], error)`
- New function `*PriceSheetClient.BeginDownloadByInvoice(context.Context, string, string, string, *PriceSheetClientBeginDownloadByInvoiceOptions) (*runtime.Poller[PriceSheetClientDownloadByInvoiceResponse], error)`
- New function `*Setting.GetSetting() *Setting`
- New function `NewSettingsClient(azcore.TokenCredential, *arm.ClientOptions) (*SettingsClient, error)`
- New function `*SettingsClient.CreateOrUpdateByScope(context.Context, string, SettingType, SettingClassification, *SettingsClientCreateOrUpdateByScopeOptions) (SettingsClientCreateOrUpdateByScopeResponse, error)`
- New function `*SettingsClient.DeleteByScope(context.Context, string, SettingType, *SettingsClientDeleteByScopeOptions) (SettingsClientDeleteByScopeResponse, error)`
- New function `*SettingsClient.GetByScope(context.Context, string, SettingType, *SettingsClientGetByScopeOptions) (SettingsClientGetByScopeResponse, error)`
- New function `*SettingsClient.List(context.Context, string, *SettingsClientListOptions) (SettingsClientListResponse, error)`
- New function `*TagInheritanceSetting.GetSetting() *Setting`
- New struct `AsyncOperationStatusProperties`
- New struct `BenefitUtilizationSummariesOperationStatus`
- New struct `BenefitUtilizationSummariesRequest`
- New struct `Budget`
- New struct `BudgetComparisonExpression`
- New struct `BudgetFilter`
- New struct `BudgetFilterProperties`
- New struct `BudgetProperties`
- New struct `BudgetTimePeriod`
- New struct `BudgetsListResult`
- New struct `CostAllocationProportion`
- New struct `CostAllocationRuleCheckNameAvailabilityRequest`
- New struct `CostAllocationRuleCheckNameAvailabilityResponse`
- New struct `CostAllocationRuleDefinition`
- New struct `CostAllocationRuleDetails`
- New struct `CostAllocationRuleList`
- New struct `CostAllocationRuleProperties`
- New struct `CurrentSpend`
- New struct `EAPriceSheetProperties`
- New struct `EAPricesheetDownloadProperties`
- New struct `ForecastSpend`
- New struct `MCAPriceSheetProperties`
- New struct `Notification`
- New struct `OperationStatusAutoGenerated`
- New struct `PricesheetDownloadProperties`
- New struct `SettingsListResult`
- New struct `SourceCostAllocationResource`
- New struct `SystemAssignedServiceIdentity`
- New struct `TagInheritanceProperties`
- New struct `TagInheritanceSetting`
- New struct `TargetCostAllocationResource`
- New field `Identity`, `Location` in struct `Export`
- New anonymous field `PricesheetDownloadProperties` in struct `PriceSheetClientDownloadByBillingProfileResponse`


## 2.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.0.0 (2023-05-26)
### Breaking Changes

- Type of `ExportExecutionListResult.Value` has been changed from `[]*ExportExecution` to `[]*ExportRun`
- Type of `ForecastDataset.Aggregation` has been changed from `map[string]*QueryAggregation` to `map[string]*ForecastAggregation`
- Type of `ForecastDataset.Configuration` has been changed from `*QueryDatasetConfiguration` to `*ForecastDatasetConfiguration`
- Type of `ForecastDataset.Filter` has been changed from `*QueryFilter` to `*ForecastFilter`
- Type of `ForecastDefinition.TimePeriod` has been changed from `*QueryTimePeriod` to `*ForecastTimePeriod`
- Type of `ForecastDefinition.Timeframe` has been changed from `*ForecastTimeframeType` to `*ForecastTimeframe`
- Type of `OperationListResult.Value` has been changed from `[]*Operation` to `[]*OperationForCostManagement`
- Type of `ReportConfigGrouping.Type` has been changed from `*ReportConfigColumnType` to `*QueryColumnType`
- `QueryColumnTypeTag` from enum `QueryColumnType` has been removed
- Enum `ForecastTimeframeType` has been removed
- Enum `ReportConfigColumnType` has been removed
- Operation `*GenerateDetailedCostReportOperationResultsClient.Get` has been changed to LRO, use `*GenerateDetailedCostReportOperationResultsClient.BeginGet` instead.
- Field `QueryResult` of struct `ForecastClientExternalCloudProviderUsageResponse` has been removed
- Field `QueryResult` of struct `ForecastClientUsageResponse` has been removed

### Features Added

- New value `QueryColumnTypeTagKey` added to enum type `QueryColumnType`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `BenefitKind` with values `BenefitKindIncludedQuantity`, `BenefitKindReservation`, `BenefitKindSavingsPlan`
- New enum type `CheckNameAvailabilityReason` with values `CheckNameAvailabilityReasonAlreadyExists`, `CheckNameAvailabilityReasonInvalid`
- New enum type `CostDetailsDataFormat` with values `CostDetailsDataFormatCSVCostDetailsDataFormat`
- New enum type `CostDetailsMetricType` with values `CostDetailsMetricTypeActualCostCostDetailsMetricType`, `CostDetailsMetricTypeAmortizedCostCostDetailsMetricType`
- New enum type `CostDetailsStatusType` with values `CostDetailsStatusTypeCompletedCostDetailsStatusType`, `CostDetailsStatusTypeFailedCostDetailsStatusType`, `CostDetailsStatusTypeNoDataFoundCostDetailsStatusType`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `DaysOfWeek` with values `DaysOfWeekFriday`, `DaysOfWeekMonday`, `DaysOfWeekSaturday`, `DaysOfWeekSunday`, `DaysOfWeekThursday`, `DaysOfWeekTuesday`, `DaysOfWeekWednesday`
- New enum type `FileFormat` with values `FileFormatCSV`
- New enum type `ForecastOperatorType` with values `ForecastOperatorTypeIn`
- New enum type `ForecastTimeframe` with values `ForecastTimeframeCustom`
- New enum type `FunctionName` with values `FunctionNameCost`, `FunctionNameCostUSD`, `FunctionNamePreTaxCost`, `FunctionNamePreTaxCostUSD`
- New enum type `Grain` with values `GrainDaily`, `GrainHourly`, `GrainMonthly`
- New enum type `GrainParameter` with values `GrainParameterDaily`, `GrainParameterHourly`, `GrainParameterMonthly`
- New enum type `LookBackPeriod` with values `LookBackPeriodLast30Days`, `LookBackPeriodLast60Days`, `LookBackPeriodLast7Days`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `ScheduleFrequency` with values `ScheduleFrequencyDaily`, `ScheduleFrequencyMonthly`, `ScheduleFrequencyWeekly`
- New enum type `ScheduledActionKind` with values `ScheduledActionKindEmail`, `ScheduledActionKindInsightAlert`
- New enum type `ScheduledActionStatus` with values `ScheduledActionStatusDisabled`, `ScheduledActionStatusEnabled`, `ScheduledActionStatusExpired`
- New enum type `Scope` with values `ScopeShared`, `ScopeSingle`
- New enum type `Term` with values `TermP1Y`, `TermP3Y`
- New enum type `WeeksOfMonth` with values `WeeksOfMonthFirst`, `WeeksOfMonthFourth`, `WeeksOfMonthLast`, `WeeksOfMonthSecond`, `WeeksOfMonthThird`
- New function `*BenefitRecommendationProperties.GetBenefitRecommendationProperties() *BenefitRecommendationProperties`
- New function `NewBenefitRecommendationsClient(azcore.TokenCredential, *arm.ClientOptions) (*BenefitRecommendationsClient, error)`
- New function `*BenefitRecommendationsClient.NewListPager(string, *BenefitRecommendationsClientListOptions) *runtime.Pager[BenefitRecommendationsClientListResponse]`
- New function `NewBenefitUtilizationSummariesClient(azcore.TokenCredential, *arm.ClientOptions) (*BenefitUtilizationSummariesClient, error)`
- New function `*BenefitUtilizationSummariesClient.NewListByBillingAccountIDPager(string, *BenefitUtilizationSummariesClientListByBillingAccountIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListByBillingAccountIDResponse]`
- New function `*BenefitUtilizationSummariesClient.NewListByBillingProfileIDPager(string, string, *BenefitUtilizationSummariesClientListByBillingProfileIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListByBillingProfileIDResponse]`
- New function `*BenefitUtilizationSummariesClient.NewListBySavingsPlanIDPager(string, string, *BenefitUtilizationSummariesClientListBySavingsPlanIDOptions) *runtime.Pager[BenefitUtilizationSummariesClientListBySavingsPlanIDResponse]`
- New function `*BenefitUtilizationSummariesClient.NewListBySavingsPlanOrderPager(string, *BenefitUtilizationSummariesClientListBySavingsPlanOrderOptions) *runtime.Pager[BenefitUtilizationSummariesClientListBySavingsPlanOrderResponse]`
- New function `*BenefitUtilizationSummary.GetBenefitUtilizationSummary() *BenefitUtilizationSummary`
- New function `*ClientFactory.NewBenefitRecommendationsClient() *BenefitRecommendationsClient`
- New function `*ClientFactory.NewBenefitUtilizationSummariesClient() *BenefitUtilizationSummariesClient`
- New function `*ClientFactory.NewGenerateCostDetailsReportClient() *GenerateCostDetailsReportClient`
- New function `*ClientFactory.NewPriceSheetClient() *PriceSheetClient`
- New function `*ClientFactory.NewScheduledActionsClient() *ScheduledActionsClient`
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
- New struct `BenefitRecommendationsListResult`
- New struct `BenefitUtilizationSummariesListResult`
- New struct `BlobInfo`
- New struct `CheckNameAvailabilityRequest`
- New struct `CheckNameAvailabilityResponse`
- New struct `CostDetailsOperationResults`
- New struct `CostDetailsTimePeriod`
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
- New struct `GenerateCostDetailsReportRequestDefinition`
- New struct `IncludedQuantityUtilizationSummary`
- New struct `IncludedQuantityUtilizationSummaryProperties`
- New struct `NotificationProperties`
- New struct `OperationForCostManagement`
- New struct `RecommendationUsageDetails`
- New struct `ReportManifest`
- New struct `RequestContext`
- New struct `SavingsPlanUtilizationSummary`
- New struct `SavingsPlanUtilizationSummaryProperties`
- New struct `ScheduleProperties`
- New struct `ScheduledAction`
- New struct `ScheduledActionListResult`
- New struct `ScheduledActionProperties`
- New struct `SharedScopeBenefitRecommendationProperties`
- New struct `SingleScopeBenefitRecommendationProperties`
- New struct `SystemData`
- New field `ExpiryTime` in struct `DownloadURL`
- New anonymous field `ForecastResult` in struct `ForecastClientExternalCloudProviderUsageResponse`
- New anonymous field `ForecastResult` in struct `ForecastClientUsageResponse`
- New field `EndTime`, `StartTime` in struct `GenerateDetailedCostReportOperationStatuses`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.



## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).