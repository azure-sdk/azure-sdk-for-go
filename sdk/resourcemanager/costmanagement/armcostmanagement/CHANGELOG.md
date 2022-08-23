# Release History

## 1.1.0 (2022-08-23)
### Features Added

- New const `CostDetailsMetricTypeActualCostCostDetailsMetricType`
- New const `CostDetailsDataFormatCSVCostDetailsDataFormat`
- New const `CostDetailsStatusTypeCompletedCostDetailsStatusType`
- New const `CostDetailsStatusTypeFailedCostDetailsStatusType`
- New const `CostDetailsMetricTypeAmortizedCostCostDetailsMetricType`
- New const `CostDetailsStatusTypeNoDataFoundCostDetailsStatusType`
- New type alias `CostDetailsStatusType`
- New type alias `CostDetailsMetricType`
- New type alias `CostDetailsDataFormat`
- New function `NewGenerateCostDetailsReportClient(azcore.TokenCredential, *arm.ClientOptions) (*GenerateCostDetailsReportClient, error)`
- New function `*GenerateCostDetailsReportClient.BeginCreateOperation(context.Context, string, GenerateCostDetailsReportRequestDefinition, *GenerateCostDetailsReportClientBeginCreateOperationOptions) (*runtime.Poller[GenerateCostDetailsReportClientCreateOperationResponse], error)`
- New function `PossibleCostDetailsMetricTypeValues() []CostDetailsMetricType`
- New function `PossibleCostDetailsStatusTypeValues() []CostDetailsStatusType`
- New function `*GenerateCostDetailsReportClient.BeginGetOperationResults(context.Context, string, string, *GenerateCostDetailsReportClientBeginGetOperationResultsOptions) (*runtime.Poller[GenerateCostDetailsReportClientGetOperationResultsResponse], error)`
- New function `PossibleCostDetailsDataFormatValues() []CostDetailsDataFormat`
- New struct `BlobInfo`
- New struct `CostDetailsOperationResults`
- New struct `CostDetailsTimePeriod`
- New struct `GenerateCostDetailsReportClient`
- New struct `GenerateCostDetailsReportClientBeginCreateOperationOptions`
- New struct `GenerateCostDetailsReportClientBeginGetOperationResultsOptions`
- New struct `GenerateCostDetailsReportClientCreateOperationResponse`
- New struct `GenerateCostDetailsReportClientGetOperationResultsResponse`
- New struct `GenerateCostDetailsReportErrorResponse`
- New struct `GenerateCostDetailsReportRequestDefinition`
- New struct `ReportManifest`
- New struct `RequestContext`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/costmanagement/armcostmanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).