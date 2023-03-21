# Release History

## 0.5.0 (2023-03-21)
### Breaking Changes

- Type of `ProjectEnvironmentTypeProperties.Status` has been changed from `*EnableStatus` to `*EnvironmentTypeEnableStatus`
- Type of `ProjectEnvironmentTypeUpdateProperties.Status` has been changed from `*EnableStatus` to `*EnvironmentTypeEnableStatus`
- Type of `ScheduleProperties.State` has been changed from `*EnableStatus` to `*ScheduleEnableStatus`
- Type of `ScheduleUpdateProperties.State` has been changed from `*EnableStatus` to `*ScheduleEnableStatus`
- Type alias `EnableStatus` has been removed
- Type alias `UsageUnit` has been removed
- Function `NewUsagesClient` has been removed
- Function `*UsagesClient.NewListByLocationPager` has been removed
- Struct `ListUsagesResult` has been removed
- Struct `Usage` has been removed
- Struct `UsageName` has been removed
- Struct `UsagesClient` has been removed
- Field `Offer` of struct `ImageReference` has been removed
- Field `Publisher` of struct `ImageReference` has been removed
- Field `SKU` of struct `ImageReference` has been removed

### Features Added

- New enum type `EnvironmentTypeEnableStatus` with values `EnvironmentTypeEnableStatusDisabled`, `EnvironmentTypeEnableStatusEnabled`
- New enum type `HealthStatus` with values `HealthStatusHealthy`, `HealthStatusPending`, `HealthStatusUnhealthy`, `HealthStatusUnknown`, `HealthStatusWarning`
- New enum type `ScheduleEnableStatus` with values `ScheduleEnableStatusDisabled`, `ScheduleEnableStatusEnabled`
- New enum type `StopOnDisconnectEnableStatus` with values `StopOnDisconnectEnableStatusDisabled`, `StopOnDisconnectEnableStatusEnabled`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAttachedNetworksClient() *AttachedNetworksClient`
- New function `*ClientFactory.NewCatalogsClient() *CatalogsClient`
- New function `*ClientFactory.NewCheckNameAvailabilityClient() *CheckNameAvailabilityClient`
- New function `*ClientFactory.NewDevBoxDefinitionsClient() *DevBoxDefinitionsClient`
- New function `*ClientFactory.NewDevCentersClient() *DevCentersClient`
- New function `*ClientFactory.NewEnvironmentTypesClient() *EnvironmentTypesClient`
- New function `*ClientFactory.NewGalleriesClient() *GalleriesClient`
- New function `*ClientFactory.NewImageVersionsClient() *ImageVersionsClient`
- New function `*ClientFactory.NewImagesClient() *ImagesClient`
- New function `*ClientFactory.NewNetworkConnectionsClient() *NetworkConnectionsClient`
- New function `*ClientFactory.NewOperationStatusesClient() *OperationStatusesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPoolsClient() *PoolsClient`
- New function `*ClientFactory.NewProjectAllowedEnvironmentTypesClient() *ProjectAllowedEnvironmentTypesClient`
- New function `*ClientFactory.NewProjectEnvironmentTypesClient() *ProjectEnvironmentTypesClient`
- New function `*ClientFactory.NewProjectsClient() *ProjectsClient`
- New function `*ClientFactory.NewSKUsClient() *SKUsClient`
- New function `*ClientFactory.NewSchedulesClient() *SchedulesClient`
- New function `*PoolsClient.BeginRunHealthChecks(context.Context, string, string, string, *PoolsClientBeginRunHealthChecksOptions) (*runtime.Poller[PoolsClientRunHealthChecksResponse], error)`
- New function `*NetworkConnectionsClient.NewListOutboundNetworkDependenciesEndpointsPager(string, string, *NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsOptions) *runtime.Pager[NetworkConnectionsClientListOutboundNetworkDependenciesEndpointsResponse]`
- New struct `ClientFactory`
- New struct `EndpointDependency`
- New struct `EndpointDetail`
- New struct `HealthStatusDetail`
- New struct `OutboundEnvironmentEndpoint`
- New struct `OutboundEnvironmentEndpointCollection`
- New struct `StopOnDisconnectConfiguration`
- New field `HibernateSupport` in struct `ImageProperties`
- New field `HealthStatus` in struct `PoolProperties`
- New field `HealthStatusDetails` in struct `PoolProperties`
- New field `StopOnDisconnect` in struct `PoolProperties`
- New field `StopOnDisconnect` in struct `PoolUpdateProperties`
- New field `MaxDevBoxesPerUser` in struct `ProjectProperties`
- New field `MaxDevBoxesPerUser` in struct `ProjectUpdateProperties`


## 0.4.0 (2022-11-24)
### Breaking Changes

- Type of `AllowedEnvironmentTypeProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `AttachedNetworkConnectionProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `CatalogProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `DevBoxDefinitionProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `EnvironmentTypeProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `GalleryProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ImageProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ImageVersionProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `NetworkProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `PoolProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ProjectEnvironmentTypeProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ProjectProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `Properties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Type of `ScheduleProperties.ProvisioningState` has been changed from `*string` to `*ProvisioningState`
- Operation `*NetworkConnectionsClient.RunHealthChecks` has been changed to LRO, use `*NetworkConnectionsClient.BeginRunHealthChecks` instead.

### Features Added

- New type alias `CheckNameAvailabilityReason`
- New type alias `HibernateSupport`
- New type alias `ProvisioningState`
- New function `NewCheckNameAvailabilityClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityClient, error)`
- New function `*CheckNameAvailabilityClient.Execute(context.Context, CheckNameAvailabilityRequest, *CheckNameAvailabilityClientExecuteOptions) (CheckNameAvailabilityClientExecuteResponse, error)`
- New struct `CheckNameAvailabilityClient`
- New struct `CheckNameAvailabilityRequest`
- New struct `CheckNameAvailabilityResponse`
- New field `HibernateSupport` in struct `DevBoxDefinitionProperties`
- New field `HibernateSupport` in struct `DevBoxDefinitionUpdateProperties`
- New field `DevCenterURI` in struct `ProjectProperties`
- New field `DevCenterURI` in struct `Properties`


## 0.3.0 (2022-10-27)
### Breaking Changes

- Type of `OperationStatus.Error` has been changed from `*OperationStatusError` to `*ErrorDetail`
- Struct `OperationStatusError` has been removed

### Features Added

- New const `CatalogSyncStateFailed`
- New const `CatalogSyncStateSucceeded`
- New const `CatalogSyncStateInProgress`
- New const `CatalogSyncStateCanceled`
- New type alias `CatalogSyncState`
- New function `PossibleCatalogSyncStateValues() []CatalogSyncState`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `OperationStatusResult`
- New anonymous field `Schedule` in struct `SchedulesClientUpdateResponse`
- New field `Operations` in struct `OperationStatus`
- New field `ResourceID` in struct `OperationStatus`
- New field `SyncState` in struct `CatalogProperties`


## 0.2.0 (2022-09-29)
### Features Added

- New function `NewProjectAllowedEnvironmentTypesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ProjectAllowedEnvironmentTypesClient, error)`
- New function `*ProjectAllowedEnvironmentTypesClient.NewListPager(string, string, *ProjectAllowedEnvironmentTypesClientListOptions) *runtime.Pager[ProjectAllowedEnvironmentTypesClientListResponse]`
- New function `*ProjectAllowedEnvironmentTypesClient.Get(context.Context, string, string, string, *ProjectAllowedEnvironmentTypesClientGetOptions) (ProjectAllowedEnvironmentTypesClientGetResponse, error)`
- New struct `AllowedEnvironmentType`
- New struct `AllowedEnvironmentTypeListResult`
- New struct `AllowedEnvironmentTypeProperties`
- New struct `ProjectAllowedEnvironmentTypesClient`
- New struct `ProjectAllowedEnvironmentTypesClientGetOptions`
- New struct `ProjectAllowedEnvironmentTypesClientGetResponse`
- New struct `ProjectAllowedEnvironmentTypesClientListOptions`
- New struct `ProjectAllowedEnvironmentTypesClientListResponse`


## 0.1.0 (2022-07-25)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devcenter/armdevcenter` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).