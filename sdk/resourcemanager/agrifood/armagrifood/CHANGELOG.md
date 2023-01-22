# Release History

## 0.8.0 (2023-01-22)
### Breaking Changes

- Type of `ErrorAdditionalInfo.Info` has been changed from `interface{}` to `any`
- Function `*ExtensionsClient.Create` has been removed
- Function `*ExtensionsClient.Update` has been removed

### Features Added

- New function `*ExtensionsClient.CreateOrUpdate(context.Context, string, string, string, *ExtensionsClientCreateOrUpdateOptions) (ExtensionsClientCreateOrUpdateResponse, error)`
- New function `NewSolutionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SolutionsClient, error)`
- New function `*SolutionsClient.CreateOrUpdate(context.Context, string, string, string, *SolutionsClientCreateOrUpdateOptions) (SolutionsClientCreateOrUpdateResponse, error)`
- New function `*SolutionsClient.Delete(context.Context, string, string, string, *SolutionsClientDeleteOptions) (SolutionsClientDeleteResponse, error)`
- New function `*SolutionsClient.Get(context.Context, string, string, string, *SolutionsClientGetOptions) (SolutionsClientGetResponse, error)`
- New function `*SolutionsClient.NewListPager(string, string, *SolutionsClientListOptions) *runtime.Pager[SolutionsClientListResponse]`
- New function `NewSolutionsDiscoverabilityClient(azcore.TokenCredential, *arm.ClientOptions) (*SolutionsDiscoverabilityClient, error)`
- New function `*SolutionsDiscoverabilityClient.Get(context.Context, string, *SolutionsDiscoverabilityClientGetOptions) (SolutionsDiscoverabilityClientGetResponse, error)`
- New function `*SolutionsDiscoverabilityClient.NewListPager(*SolutionsDiscoverabilityClientListOptions) *runtime.Pager[SolutionsDiscoverabilityClientListResponse]`
- New struct `APIProperties`
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
- New struct `SolutionsClient`
- New struct `SolutionsClientListResponse`
- New struct `SolutionsDiscoverabilityClient`
- New struct `SolutionsDiscoverabilityClientListResponse`
- New field `AdditionalAPIProperties` in struct `ExtensionProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
