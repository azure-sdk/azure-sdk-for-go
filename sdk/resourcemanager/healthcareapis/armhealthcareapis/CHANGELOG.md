# Release History

## 1.1.0-beta.2 (2022-10-25)
### Features Added

- New const `FhirServiceVersionSTU3`
- New const `FhirServiceVersionR4`
- New const `AnalyticsConnectorMappingTypeFhirToParquet`
- New const `AnalyticsConnectorDataSourceTypeFhirservice`
- New const `AnalyticsConnectorDataDestinationTypeDatalake`
- New type alias `AnalyticsConnectorDataSourceType`
- New type alias `FhirServiceVersion`
- New type alias `AnalyticsConnectorMappingType`
- New type alias `AnalyticsConnectorDataDestinationType`
- New function `*AnalyticsConnectorClient.BeginUpdate(context.Context, string, string, string, AnalyticsConnectorPatchResource, *AnalyticsConnectorClientBeginUpdateOptions) (*runtime.Poller[AnalyticsConnectorClientUpdateResponse], error)`
- New function `*AnalyticsConnectorDataDestination.GetAnalyticsConnectorDataDestination() *AnalyticsConnectorDataDestination`
- New function `PossibleAnalyticsConnectorMappingTypeValues() []AnalyticsConnectorMappingType`
- New function `PossibleAnalyticsConnectorDataSourceTypeValues() []AnalyticsConnectorDataSourceType`
- New function `*AnalyticsConnectorDataSource.GetAnalyticsConnectorDataSource() *AnalyticsConnectorDataSource`
- New function `*AnalyticsConnectorsClient.BeginDelete(context.Context, string, string, string, *AnalyticsConnectorsClientBeginDeleteOptions) (*runtime.Poller[AnalyticsConnectorsClientDeleteResponse], error)`
- New function `*AnalyticsConnectorFhirServiceDataSource.GetAnalyticsConnectorDataSource() *AnalyticsConnectorDataSource`
- New function `NewAnalyticsConnectorClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AnalyticsConnectorClient, error)`
- New function `*AnalyticsConnectorFhirToParquetMapping.GetAnalyticsConnectorMapping() *AnalyticsConnectorMapping`
- New function `*AnalyticsConnectorDataLakeDataDestination.GetAnalyticsConnectorDataDestination() *AnalyticsConnectorDataDestination`
- New function `PossibleAnalyticsConnectorDataDestinationTypeValues() []AnalyticsConnectorDataDestinationType`
- New function `*AnalyticsConnectorsClient.BeginCreateOrUpdate(context.Context, string, string, string, AnalyticsConnector, *AnalyticsConnectorsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AnalyticsConnectorsClientCreateOrUpdateResponse], error)`
- New function `*AnalyticsConnectorMapping.GetAnalyticsConnectorMapping() *AnalyticsConnectorMapping`
- New function `*AnalyticsConnectorsClient.Get(context.Context, string, string, string, *AnalyticsConnectorsClientGetOptions) (AnalyticsConnectorsClientGetResponse, error)`
- New function `*AnalyticsConnectorsClient.NewListByWorkspacePager(string, string, *AnalyticsConnectorsClientListByWorkspaceOptions) *runtime.Pager[AnalyticsConnectorsClientListByWorkspaceResponse]`
- New function `PossibleFhirServiceVersionValues() []FhirServiceVersion`
- New function `NewAnalyticsConnectorsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AnalyticsConnectorsClient, error)`
- New struct `AnalyticsConnector`
- New struct `AnalyticsConnectorClient`
- New struct `AnalyticsConnectorClientBeginUpdateOptions`
- New struct `AnalyticsConnectorClientUpdateResponse`
- New struct `AnalyticsConnectorCollection`
- New struct `AnalyticsConnectorDataDestination`
- New struct `AnalyticsConnectorDataLakeDataDestination`
- New struct `AnalyticsConnectorDataSource`
- New struct `AnalyticsConnectorFhirServiceDataSource`
- New struct `AnalyticsConnectorFhirToParquetMapping`
- New struct `AnalyticsConnectorMapping`
- New struct `AnalyticsConnectorPatchResource`
- New struct `AnalyticsConnectorProperties`
- New struct `AnalyticsConnectorsClient`
- New struct `AnalyticsConnectorsClientBeginCreateOrUpdateOptions`
- New struct `AnalyticsConnectorsClientBeginDeleteOptions`
- New struct `AnalyticsConnectorsClientCreateOrUpdateResponse`
- New struct `AnalyticsConnectorsClientDeleteResponse`
- New struct `AnalyticsConnectorsClientGetOptions`
- New struct `AnalyticsConnectorsClientGetResponse`
- New struct `AnalyticsConnectorsClientListByWorkspaceOptions`
- New struct `AnalyticsConnectorsClientListByWorkspaceResponse`
- New struct `CorsConfiguration`
- New field `CorsConfiguration` in struct `DicomServiceProperties`
- New field `IsInternal` in struct `MetricSpecification`
- New field `MetricFilterPattern` in struct `MetricSpecification`
- New field `ResourceIDDimensionNameOverride` in struct `MetricSpecification`
- New field `SourceMdmAccount` in struct `MetricSpecification`
- New field `EnableRegionalMdmAccount` in struct `MetricSpecification`


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