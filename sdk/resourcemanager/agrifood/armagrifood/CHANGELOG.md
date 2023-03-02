# Release History

## 0.8.0 (2023-03-02)
### Breaking Changes

- Const `PublicNetworkAccessHybrid` from type alias `PublicNetworkAccess` has been removed
- Function `*ExtensionsClient.Create` has been removed
- Function `*ExtensionsClient.NewListByFarmBeatsPager` has been removed
- Function `*ExtensionsClient.Update` has been removed
- Function `NewFarmBeatsExtensionsClient` has been removed
- Function `*FarmBeatsExtensionsClient.Get` has been removed
- Function `*FarmBeatsExtensionsClient.NewListPager` has been removed
- Function `NewFarmBeatsModelsClient` has been removed
- Function `*FarmBeatsModelsClient.CreateOrUpdate` has been removed
- Function `*FarmBeatsModelsClient.Delete` has been removed
- Function `*FarmBeatsModelsClient.Get` has been removed
- Function `*FarmBeatsModelsClient.GetOperationResult` has been removed
- Function `*FarmBeatsModelsClient.NewListByResourceGroupPager` has been removed
- Function `*FarmBeatsModelsClient.NewListBySubscriptionPager` has been removed
- Function `*FarmBeatsModelsClient.BeginUpdate` has been removed
- Function `NewLocationsClient` has been removed
- Function `*LocationsClient.CheckNameAvailability` has been removed
- Operation `*PrivateEndpointConnectionsClient.NewListByResourcePager` does not support pagination anymore, use `*PrivateEndpointConnectionsClient.ListByResource` instead.
- Operation `*PrivateLinkResourcesClient.NewListByResourcePager` does not support pagination anymore, use `*PrivateLinkResourcesClient.ListByResource` instead.
- Struct `FarmBeats` has been removed
- Struct `FarmBeatsExtension` has been removed
- Struct `FarmBeatsExtensionListResponse` has been removed
- Struct `FarmBeatsExtensionProperties` has been removed
- Struct `FarmBeatsExtensionsClient` has been removed
- Struct `FarmBeatsListResponse` has been removed
- Struct `FarmBeatsModelsClient` has been removed
- Struct `FarmBeatsProperties` has been removed
- Struct `FarmBeatsUpdateProperties` has been removed
- Struct `FarmBeatsUpdateRequestModel` has been removed
- Struct `LocationsClient` has been removed
- Field `PrivateEndpointConnection` of struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse` has been removed
- Field `PrivateEndpointConnection` of struct `PrivateEndpointConnectionsClientGetResponse` has been removed
- Field `PrivateLinkResource` of struct `PrivateLinkResourcesClientGetResponse` has been removed

### Features Added

- New value `ProvisioningStateRunning` added to type alias `ProvisioningState`
- New value `PublicNetworkAccessDisabled` added to type alias `PublicNetworkAccess`
- New function `NewCheckNameAvailabilityClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CheckNameAvailabilityClient, error)`
- New function `*CheckNameAvailabilityClient.CheckNameAvailability(context.Context, CheckNameAvailabilityRequest, *CheckNameAvailabilityClientCheckNameAvailabilityOptions) (CheckNameAvailabilityClientCheckNameAvailabilityResponse, error)`
- New function `NewDataManagerForAgricultureExtensionsClient(azcore.TokenCredential, *arm.ClientOptions) (*DataManagerForAgricultureExtensionsClient, error)`
- New function `*DataManagerForAgricultureExtensionsClient.Get(context.Context, string, *DataManagerForAgricultureExtensionsClientGetOptions) (DataManagerForAgricultureExtensionsClientGetResponse, error)`
- New function `*DataManagerForAgricultureExtensionsClient.NewListPager(*DataManagerForAgricultureExtensionsClientListOptions) *runtime.Pager[DataManagerForAgricultureExtensionsClientListResponse]`
- New function `NewDataManagerForAgricultureResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DataManagerForAgricultureResourcesClient, error)`
- New function `*DataManagerForAgricultureResourcesClient.CreateOrUpdate(context.Context, string, string, DataManagerForAgriculture, *DataManagerForAgricultureResourcesClientCreateOrUpdateOptions) (DataManagerForAgricultureResourcesClientCreateOrUpdateResponse, error)`
- New function `*DataManagerForAgricultureResourcesClient.Delete(context.Context, string, string, *DataManagerForAgricultureResourcesClientDeleteOptions) (DataManagerForAgricultureResourcesClientDeleteResponse, error)`
- New function `*DataManagerForAgricultureResourcesClient.Get(context.Context, string, string, *DataManagerForAgricultureResourcesClientGetOptions) (DataManagerForAgricultureResourcesClientGetResponse, error)`
- New function `*DataManagerForAgricultureResourcesClient.NewListByResourceGroupPager(string, *DataManagerForAgricultureResourcesClientListByResourceGroupOptions) *runtime.Pager[DataManagerForAgricultureResourcesClientListByResourceGroupResponse]`
- New function `*DataManagerForAgricultureResourcesClient.NewListBySubscriptionPager(*DataManagerForAgricultureResourcesClientListBySubscriptionOptions) *runtime.Pager[DataManagerForAgricultureResourcesClientListBySubscriptionResponse]`
- New function `*DataManagerForAgricultureResourcesClient.BeginUpdate(context.Context, string, string, DataManagerForAgricultureUpdateRequestModel, *DataManagerForAgricultureResourcesClientBeginUpdateOptions) (*runtime.Poller[DataManagerForAgricultureResourcesClientUpdateResponse], error)`
- New function `*ExtensionsClient.CreateOrUpdate(context.Context, string, string, string, *ExtensionsClientCreateOrUpdateOptions) (ExtensionsClientCreateOrUpdateResponse, error)`
- New function `*ExtensionsClient.NewListByDataManagerForAgriculturePager(string, string, *ExtensionsClientListByDataManagerForAgricultureOptions) *runtime.Pager[ExtensionsClientListByDataManagerForAgricultureResponse]`
- New function `NewOperationResultsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OperationResultsClient, error)`
- New function `*OperationResultsClient.Get(context.Context, string, string, *OperationResultsClientGetOptions) (OperationResultsClientGetResponse, error)`
- New function `NewSolutionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SolutionsClient, error)`
- New function `*SolutionsClient.CreateOrUpdate(context.Context, string, string, string, *SolutionsClientCreateOrUpdateOptions) (SolutionsClientCreateOrUpdateResponse, error)`
- New function `*SolutionsClient.Delete(context.Context, string, string, string, *SolutionsClientDeleteOptions) (SolutionsClientDeleteResponse, error)`
- New function `*SolutionsClient.Get(context.Context, string, string, string, *SolutionsClientGetOptions) (SolutionsClientGetResponse, error)`
- New function `*SolutionsClient.NewListPager(string, string, *SolutionsClientListOptions) *runtime.Pager[SolutionsClientListResponse]`
- New function `NewSolutionsDiscoverabilityClient(azcore.TokenCredential, *arm.ClientOptions) (*SolutionsDiscoverabilityClient, error)`
- New function `*SolutionsDiscoverabilityClient.Get(context.Context, string, *SolutionsDiscoverabilityClientGetOptions) (SolutionsDiscoverabilityClientGetResponse, error)`
- New function `*SolutionsDiscoverabilityClient.NewListPager(*SolutionsDiscoverabilityClientListOptions) *runtime.Pager[SolutionsDiscoverabilityClientListResponse]`
- New struct `APIProperties`
- New struct `ArmAsyncOperationError`
- New struct `CheckNameAvailabilityClient`
- New struct `DataManagerForAgriculture`
- New struct `DataManagerForAgricultureExtension`
- New struct `DataManagerForAgricultureExtensionListResponse`
- New struct `DataManagerForAgricultureExtensionProperties`
- New struct `DataManagerForAgricultureExtensionsClient`
- New struct `DataManagerForAgricultureListResponse`
- New struct `DataManagerForAgricultureProperties`
- New struct `DataManagerForAgricultureResourcesClient`
- New struct `DataManagerForAgricultureSolution`
- New struct `DataManagerForAgricultureSolutionListResponse`
- New struct `DataManagerForAgricultureSolutionProperties`
- New struct `DataManagerForAgricultureUpdateProperties`
- New struct `DataManagerForAgricultureUpdateRequestModel`
- New struct `ExtensionInstallationRequest`
- New struct `Insight`
- New struct `InsightAttachment`
- New struct `MarketplaceOfferDetails`
- New struct `Measure`
- New struct `OperationResultsClient`
- New struct `ResourceParameter`
- New struct `Solution`
- New struct `SolutionEvaluatedOutput`
- New struct `SolutionInstallationRequest`
- New struct `SolutionListResponse`
- New struct `SolutionProperties`
- New struct `SolutionsClient`
- New struct `SolutionsDiscoverabilityClient`
- New field `Error` in struct `ArmAsyncOperation`
- New field `APIDefaultInputParameters` in struct `DetailedInformation`
- New field `APIDocsLink` in struct `DetailedInformation`
- New field `APIType` in struct `DetailedInformation`
- New field `SkipToken` in struct `ExtensionListResponse`
- New field `AdditionalAPIProperties` in struct `ExtensionProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`
- New field `Value` in struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New field `Value` in struct `PrivateEndpointConnectionsClientGetResponse`
- New field `Value` in struct `PrivateLinkResourcesClientGetResponse`


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
