# Release History

## 1.2.0 (2023-08-31)
### Features Added

- New value `ConnectionStateActiveNoBilling`, `ConnectionStateTypeChangeInProgress`, `ConnectionStateTypeChangeRequested` added to enum type `ConnectionState`
- New function `*ClientFactory.NewResourceMoveClient() *ResourceMoveClient`
- New function `*ClientFactory.NewRpUnbilledPrefixesClient() *RpUnbilledPrefixesClient`
- New function `*RegisteredPrefixesClient.Validate(context.Context, string, string, string, *RegisteredPrefixesClientValidateOptions) (RegisteredPrefixesClientValidateResponse, error)`
- New function `NewResourceMoveClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ResourceMoveClient, error)`
- New function `*ResourceMoveClient.MoveResources(context.Context, string, ResourceMoveRequest, *ResourceMoveClientMoveResourcesOptions) (ResourceMoveClientMoveResourcesResponse, error)`
- New function `*ResourceMoveClient.ValidateMoveResources(context.Context, string, ResourceMoveRequest, *ResourceMoveClientValidateMoveResourcesOptions) (ResourceMoveClientValidateMoveResourcesResponse, error)`
- New function `NewRpUnbilledPrefixesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RpUnbilledPrefixesClient, error)`
- New function `*RpUnbilledPrefixesClient.NewListPager(string, string, *RpUnbilledPrefixesClientListOptions) *runtime.Pager[RpUnbilledPrefixesClientListResponse]`
- New struct `ResourceMoveRequest`
- New struct `RpUnbilledPrefix`
- New struct `RpUnbilledPrefixListResult`
- New field `DirectPeeringType` in struct `LegacyPeeringsClientListOptions`


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/peering/armpeering` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).