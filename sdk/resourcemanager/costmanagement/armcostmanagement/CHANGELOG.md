# Release History

## 2.0.0 (2022-09-26)
### Breaking Changes

- Type of `ForecastDefinition.Timeframe` has been changed from `*ForecastTimeframeType` to `*ForecastTimeframe`
- Type of `ForecastDefinition.TimePeriod` has been changed from `*QueryTimePeriod` to `*ForecastTimePeriod`
- Type of `ForecastDataset.Aggregation` has been changed from `map[string]*QueryAggregation` to `map[string]*ForecastAggregation`
- Type of `ForecastDataset.Configuration` has been changed from `*QueryDatasetConfiguration` to `*ForecastDatasetConfiguration`
- Type of `ForecastDataset.Filter` has been changed from `*QueryFilter` to `*ForecastFilter`
- Const `ForecastTimeframeTypeTheLastBillingMonth` has been removed
- Const `ForecastTimeframeTypeCustom` has been removed
- Const `ForecastTimeframeTypeWeekToDate` has been removed
- Const `ForecastTimeframeTypeTheLastMonth` has been removed
- Const `ForecastTimeframeTypeMonthToDate` has been removed
- Const `ForecastTimeframeTypeBillingMonthToDate` has been removed
- Type alias `ForecastTimeframeType` has been removed
- Function `PossibleForecastTimeframeTypeValues` has been removed
- Field `QueryResult` of struct `ForecastClientExternalCloudProviderUsageResponse` has been removed
- Field `QueryResult` of struct `ForecastClientUsageResponse` has been removed

### Features Added

- New const `FunctionNameCost`
- New const `ForecastOperatorTypeIn`
- New const `FunctionNamePreTaxCost`
- New const `CostDetailsStatusTypeNoDataFoundCostDetailsStatusType`
- New const `ForecastTimeframeCustom`
- New const `FunctionNamePreTaxCostUSD`
- New const `CostDetailsStatusTypeFailedCostDetailsStatusType`
- New const `CostDetailsMetricTypeActualCostCostDetailsMetricType`
- New const `CostDetailsMetricTypeAmortizedCostCostDetailsMetricType`
- New const `CostDetailsDataFormatCSVCostDetailsDataFormat`
- New const `CostDetailsStatusTypeCompletedCostDetailsStatusType`
- New const `FunctionNameCostUSD`
- New type alias `ForecastOperatorType`
- New type alias `CostDetailsStatusType`
- New type alias `ForecastTimeframe`
- New type alias `FunctionName`
- New type alias `CostDetailsDataFormat`
- New type alias `CostDetailsMetricType`
- New function `*GenerateCostDetailsReportClient.BeginGetOperationResults(context.Context, string, string, *GenerateCostDetailsReportClientBeginGetOperationResultsOptions) (*runtime.Poller[GenerateCostDetailsReportClientGetOperationResultsResponse], error)`
- New function `PossibleForecastOperatorTypeValues() []ForecastOperatorType`
- New function `PossibleCostDetailsStatusTypeValues() []CostDetailsStatusType`
- New function `NewGenerateCostDetailsReportClient(azcore.TokenCredential, *arm.ClientOptions) (*GenerateCostDetailsReportClient, error)`
- New function `PossibleCostDetailsDataFormatValues() []CostDetailsDataFormat`
- New function `PossibleForecastTimeframeValues() []ForecastTimeframe`
- New function `PossibleFunctionNameValues() []FunctionName`
- New function `PossibleCostDetailsMetricTypeValues() []CostDetailsMetricType`
- New function `*GenerateCostDetailsReportClient.BeginCreateOperation(context.Context, string, GenerateCostDetailsReportRequestDefinition, *GenerateCostDetailsReportClientBeginCreateOperationOptions) (*runtime.Poller[GenerateCostDetailsReportClientCreateOperationResponse], error)`
- New struct `BlobInfo`
- New struct `CostDetailsOperationResults`
- New struct `CostDetailsTimePeriod`
- New struct `ForecastAggregation`
- New struct `ForecastColumn`
- New struct `ForecastComparisonExpression`
- New struct `ForecastDatasetConfiguration`
- New struct `ForecastFilter`
- New struct `ForecastProperties`
- New struct `ForecastResult`
- New struct `ForecastTimePeriod`
- New struct `GenerateCostDetailsReportClient`
- New struct `GenerateCostDetailsReportClientBeginCreateOperationOptions`
- New struct `GenerateCostDetailsReportClientBeginGetOperationResultsOptions`
- New struct `GenerateCostDetailsReportClientCreateOperationResponse`
- New struct `GenerateCostDetailsReportClientGetOperationResultsResponse`
- New struct `GenerateCostDetailsReportErrorResponse`
- New struct `GenerateCostDetailsReportRequestDefinition`
- New struct `ReportManifest`
- New struct `RequestContext`
- New anonymous field `ForecastResult` in struct `ForecastClientUsageResponse`
- New anonymous field `ForecastResult` in struct `ForecastClientExternalCloudProviderUsageResponse`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).