# Release History

## 0.9.0 (2023-06-23)
### Breaking Changes

- Function `*ExtensionsClient.Create` has been removed
- Function `*ExtensionsClient.Update` has been removed
- Function `*FarmBeatsModelsClient.GetOperationResult` has been removed

### Features Added

- New function `*ClientFactory.NewOperationResultsClient() *OperationResultsClient`
- New function `*ClientFactory.NewSolutionsClient() *SolutionsClient`
- New function `*ClientFactory.NewSolutionsDiscoverabilityClient() *SolutionsDiscoverabilityClient`
- New function `*ExtensionsClient.CreateOrUpdate2(context.Context, string, string, string, *ExtensionsClientCreateOrUpdate2Options) (ExtensionsClientCreateOrUpdate2Response, error)`
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
- New struct `ExtensionInstallationRequest`
- New struct `FarmBeatsSolution`
- New struct `FarmBeatsSolutionListResponse`
- New struct `FarmBeatsSolutionProperties`
- New struct `Insight`
- New struct `InsightAttachment`
- New struct `MarketplaceOfferDetails`
- New struct `Measure`
- New struct `ResourceParameter`
- New struct `Solution`
- New struct `SolutionEvaluatedOutput`
- New struct `SolutionInstallationRequest`
- New struct `SolutionListResponse`
- New struct `SolutionProperties`
- New field `Error` in struct `ArmAsyncOperation`
- New field `AdditionalAPIProperties` in struct `ExtensionProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`


## 0.8.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 0.8.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
