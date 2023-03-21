# Release History

## 1.1.0 (2023-03-21)
### Features Added

- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewDicomServicesClient() *DicomServicesClient`
- New function `*ClientFactory.NewFhirDestinationsClient() *FhirDestinationsClient`
- New function `*ClientFactory.NewFhirServicesClient() *FhirServicesClient`
- New function `*ClientFactory.NewIotConnectorFhirDestinationClient() *IotConnectorFhirDestinationClient`
- New function `*ClientFactory.NewIotConnectorsClient() *IotConnectorsClient`
- New function `*ClientFactory.NewOperationResultsClient() *OperationResultsClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewServicesClient() *ServicesClient`
- New function `*ClientFactory.NewWorkspacePrivateEndpointConnectionsClient() *WorkspacePrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewWorkspacePrivateLinkResourcesClient() *WorkspacePrivateLinkResourcesClient`
- New function `*ClientFactory.NewWorkspacesClient() *WorkspacesClient`
- New struct `ClientFactory`
- New struct `CorsConfiguration`
- New struct `FhirServiceImportConfiguration`
- New struct `ImplementationGuidesConfiguration`
- New struct `ServiceImportConfigurationInfo`
- New field `CorsConfiguration` in struct `DicomServiceProperties`
- New field `EventState` in struct `DicomServiceProperties`
- New field `ImplementationGuidesConfiguration` in struct `FhirServiceProperties`
- New field `ImportConfiguration` in struct `FhirServiceProperties`
- New field `EnableRegionalMdmAccount` in struct `MetricSpecification`
- New field `IsInternal` in struct `MetricSpecification`
- New field `MetricFilterPattern` in struct `MetricSpecification`
- New field `ResourceIDDimensionNameOverride` in struct `MetricSpecification`
- New field `SourceMdmAccount` in struct `MetricSpecification`
- New field `CrossTenantCmkApplicationID` in struct `ServiceCosmosDbConfigurationInfo`
- New field `ImportConfiguration` in struct `ServicesProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthcareapis/armhealthcareapis` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).