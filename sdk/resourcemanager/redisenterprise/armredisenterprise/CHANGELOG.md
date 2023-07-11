# Release History

## 2.0.0-beta.1 (2023-07-11)
### Breaking Changes

- Operation `*PrivateEndpointConnectionsClient.Delete` has been changed to LRO, use `*PrivateEndpointConnectionsClient.BeginDelete` instead.

### Features Added

- New value `ResourceStateScaling`, `ResourceStateScalingFailed` added to enum type `ResourceState`
- New enum type `CmkIdentityType` with values `CmkIdentityTypeSystemAssignedIdentity`, `CmkIdentityTypeUserAssignedIdentity`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New function `*Client.CheckNameAvailability(context.Context, CheckNameAvailabilityParameters, *ClientCheckNameAvailabilityOptions) (ClientCheckNameAvailabilityResponse, error)`
- New function `*ClientFactory.NewSKUsClient() *SKUsClient`
- New function `*DatabasesClient.BeginFlush(context.Context, string, string, string, FlushParameters, *DatabasesClientBeginFlushOptions) (*runtime.Poller[DatabasesClientFlushResponse], error)`
- New function `NewSKUsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SKUsClient, error)`
- New function `*SKUsClient.NewListPager(string, *SKUsClientListOptions) *runtime.Pager[SKUsClientListResponse]`
- New function `timeRFC3339.MarshalText() ([]byte, error)`
- New function `*timeRFC3339.Parse(string) error`
- New function `*timeRFC3339.UnmarshalText([]byte) error`
- New struct `Capability`
- New struct `CheckNameAvailabilityParameters`
- New struct `ClusterPropertiesEncryption`
- New struct `ClusterPropertiesEncryptionCustomerManagedKeyEncryption`
- New struct `ClusterPropertiesEncryptionCustomerManagedKeyEncryptionKeyIdentity`
- New struct `FlushParameters`
- New struct `LocationInfo`
- New struct `ManagedServiceIdentity`
- New struct `RegionSKUDetail`
- New struct `RegionSKUDetails`
- New struct `SKUDetail`
- New struct `SystemData`
- New struct `UserAssignedIdentity`
- New field `Identity`, `SystemData` in struct `Cluster`
- New field `Encryption` in struct `ClusterProperties`
- New field `Identity` in struct `ClusterUpdate`
- New field `SystemData` in struct `Database`
- New field `SystemData` in struct `PrivateEndpointConnection`
- New field `SystemData` in struct `PrivateLinkResource`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-04-07)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redisenterprise/armredisenterprise` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).