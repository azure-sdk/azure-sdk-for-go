# Release History

## 1.1.0 (2022-08-17)
### Features Added

- New const `ColumnTypeDimension`
- New const `CostDetailsStatusTypeNoDataFoundCostDetailsStatusType`
- New const `ColumnTypeTag`
- New const `CostDetailsMetricTypeActualCostCostDetailsMetricType`
- New const `CostDetailsDataFormatCSVCostDetailsDataFormat`
- New const `ColumnTypeMetric`
- New const `ColumnTypeTagKey`
- New const `ColumnTypeTagValue`
- New const `CostDetailsStatusTypeFailedCostDetailsStatusType`
- New const `CostDetailsStatusTypeCompletedCostDetailsStatusType`
- New const `ColumnTypeNone`
- New const `CostDetailsMetricTypeAmortizedCostCostDetailsMetricType`
- New function `NewGenerateCostDetailsReportClient(azcore.TokenCredential, *arm.ClientOptions) (*GenerateCostDetailsReportClient, error)`
- New function `PossibleColumnTypeValues() []ColumnType`
- New function `PossibleCostDetailsDataFormatValues() []CostDetailsDataFormat`
- New function `*GenerateCostDetailsReportClient.BeginGetOperationResults(context.Context, string, string, *GenerateCostDetailsReportClientBeginGetOperationResultsOptions) (*runtime.Poller[GenerateCostDetailsReportClientGetOperationResultsResponse], error)`
- New function `*GenerateCostDetailsReportClient.BeginCreateOperation(context.Context, string, GenerateCostDetailsReportRequestDefinition, *GenerateCostDetailsReportClientBeginCreateOperationOptions) (*runtime.Poller[GenerateCostDetailsReportClientCreateOperationResponse], error)`
- New function `PossibleCostDetailsMetricTypeValues() []CostDetailsMetricType`
- New function `PossibleCostDetailsStatusTypeValues() []CostDetailsStatusType`
- New struct `BlobInfo`
- New struct `CostDetailsOperationResults`
- New struct `CostDetailsTimePeriod`
- New struct `ExportDatasetGrouping`
- New struct `GenerateCostDetailsReportClient`
- New struct `GenerateCostDetailsReportClientBeginCreateOperationOptions`
- New struct `GenerateCostDetailsReportClientBeginGetOperationResultsOptions`
- New struct `GenerateCostDetailsReportClientCreateOperationResponse`
- New struct `GenerateCostDetailsReportClientGetOperationResultsResponse`
- New struct `GenerateCostDetailsReportErrorResponse`
- New struct `GenerateCostDetailsReportRequestDefinition`
- New struct `ReportManifest`
- New struct `RequestContext`
- New field `Grouping` in struct `ExportDataset`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).