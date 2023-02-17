# Release History

## 1.1.0-beta.2 (2023-02-17)
### Features Added

- New type alias `AnalyticsConnectorDataDestinationType` with values `AnalyticsConnectorDataDestinationTypeDatalake`
- New type alias `AnalyticsConnectorDataSourceType` with values `AnalyticsConnectorDataSourceTypeDicomservice`, `AnalyticsConnectorDataSourceTypeFhirservice`
- New type alias `AnalyticsConnectorMappingType` with values `AnalyticsConnectorMappingTypeDicomToParquet`, `AnalyticsConnectorMappingTypeFhirToParquet`
- New type alias `FhirServiceVersion` with values `FhirServiceVersionR4`, `FhirServiceVersionSTU3`
- New function `*AnalyticsConnectorDataDestination.GetAnalyticsConnectorDataDestination() *AnalyticsConnectorDataDestination`
- New function `*AnalyticsConnectorDataLakeDataDestination.GetAnalyticsConnectorDataDestination() *AnalyticsConnectorDataDestination`
- New function `*AnalyticsConnectorDataSource.GetAnalyticsConnectorDataSource() *AnalyticsConnectorDataSource`
- New function `*AnalyticsConnectorDicomServiceDataSource.GetAnalyticsConnectorDataSource() *AnalyticsConnectorDataSource`
- New function `*AnalyticsConnectorDicomToParquetMapping.GetAnalyticsConnectorMapping() *AnalyticsConnectorMapping`
- New function `*AnalyticsConnectorFhirServiceDataSource.GetAnalyticsConnectorDataSource() *AnalyticsConnectorDataSource`
- New function `*AnalyticsConnectorFhirToParquetMapping.GetAnalyticsConnectorMapping() *AnalyticsConnectorMapping`
- New function `*AnalyticsConnectorMapping.GetAnalyticsConnectorMapping() *AnalyticsConnectorMapping`
- New function `NewAnalyticsConnectorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AnalyticsConnectorsClient, error)`
- New function `*AnalyticsConnectorsClient.BeginCreateOrUpdate(context.Context, string, string, string, AnalyticsConnector, *AnalyticsConnectorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AnalyticsConnectorsClientCreateOrUpdateResponse], error)`
- New function `*AnalyticsConnectorsClient.BeginDelete(context.Context, string, string, string, *AnalyticsConnectorsClientBeginDeleteOptions) (*runtime.Poller[AnalyticsConnectorsClientDeleteResponse], error)`
- New function `*AnalyticsConnectorsClient.Get(context.Context, string, string, string, *AnalyticsConnectorsClientGetOptions) (AnalyticsConnectorsClientGetResponse, error)`
- New function `*AnalyticsConnectorsClient.NewListByWorkspacePager(string, string, *AnalyticsConnectorsClientListByWorkspaceOptions) *runtime.Pager[AnalyticsConnectorsClientListByWorkspaceResponse]`
- New function `*AnalyticsConnectorsClient.BeginUpdate(context.Context, string, string, string, AnalyticsConnectorPatchResource, *AnalyticsConnectorsClientBeginUpdateOptions) (*runtime.Poller[AnalyticsConnectorsClientUpdateResponse], error)`
- New struct `AnalyticsConnector`
- New struct `AnalyticsConnectorCollection`
- New struct `AnalyticsConnectorDataLakeDataDestination`
- New struct `AnalyticsConnectorDicomServiceDataSource`
- New struct `AnalyticsConnectorDicomToParquetMapping`
- New struct `AnalyticsConnectorFhirServiceDataSource`
- New struct `AnalyticsConnectorFhirToParquetMapping`
- New struct `AnalyticsConnectorPatchResource`
- New struct `AnalyticsConnectorProperties`
- New struct `AnalyticsConnectorsClient`
- New struct `CorsConfiguration`
- New struct `ImplementationGuidesConfiguration`
- New field `CorsConfiguration` in struct `DicomServiceProperties`
- New field `EventState` in struct `DicomServiceProperties`
- New field `ImplementationGuidesConfiguration` in struct `FhirServiceProperties`
- New field `EnableRegionalMdmAccount` in struct `MetricSpecification`
- New field `IsInternal` in struct `MetricSpecification`
- New field `MetricFilterPattern` in struct `MetricSpecification`
- New field `ResourceIDDimensionNameOverride` in struct `MetricSpecification`
- New field `SourceMdmAccount` in struct `MetricSpecification`


## 1.1.0-beta.1 (2022-05-19)
### Features Added

- New struct `FhirServiceImportConfiguration`
- New struct `ServiceImportConfigurationInfo`
- New field `ImportConfiguration` in struct `ServicesProperties`
- New field `ImportConfiguration` in struct `FhirServiceProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).