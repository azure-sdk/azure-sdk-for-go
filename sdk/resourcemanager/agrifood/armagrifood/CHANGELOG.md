# Release History

## 0.7.0 (2022-09-01)
### Breaking Changes

- Function `*FarmBeatsModelsClient.Update` has been removed
- Struct `FarmBeatsModelsClientUpdateOptions` has been removed

### Features Added

- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `ProvisioningStateCreating`
- New const `PrivateEndpointConnectionProvisioningStateFailed`
- New const `PrivateEndpointServiceConnectionStatusRejected`
- New const `PublicNetworkAccessHybrid`
- New const `ProvisioningStateUpdating`
- New const `PublicNetworkAccessEnabled`
- New const `ProvisioningStateDeleting`
- New const `PrivateEndpointServiceConnectionStatusApproved`
- New const `PrivateEndpointConnectionProvisioningStateSucceeded`
- New const `PrivateEndpointServiceConnectionStatusPending`
- New type alias `PublicNetworkAccess`
- New type alias `PrivateEndpointConnectionProvisioningState`
- New type alias `PrivateEndpointServiceConnectionStatus`
- New function `*PrivateEndpointConnectionsClient.Get(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error)`
- New function `*PrivateEndpointConnectionsClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error)`
- New function `*FarmBeatsModelsClient.BeginUpdate(context.Context, string, string, FarmBeatsUpdateRequestModel, *FarmBeatsModelsClientBeginUpdateOptions) (*runtime.Poller[FarmBeatsModelsClientUpdateResponse], error)`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.Get(context.Context, string, string, string, *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error)`
- New function `PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus`
- New function `PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState`
- New function `*FarmBeatsModelsClient.GetOperationResult(context.Context, string, string, string, *FarmBeatsModelsClientGetOperationResultOptions) (FarmBeatsModelsClientGetOperationResultResponse, error)`
- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateLinkResourcesClient.NewListByResourcePager(string, string, *PrivateLinkResourcesClientListByResourceOptions) *runtime.Pager[PrivateLinkResourcesClientListByResourceResponse]`
- New function `PossiblePublicNetworkAccessValues() []PublicNetworkAccess`
- New function `*PrivateEndpointConnectionsClient.NewListByResourcePager(string, string, *PrivateEndpointConnectionsClientListByResourceOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByResourceResponse]`
- New function `*PrivateEndpointConnectionsClient.CreateOrUpdate(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientCreateOrUpdateOptions) (PrivateEndpointConnectionsClientCreateOrUpdateResponse, error)`
- New struct `ArmAsyncOperation`
- New struct `FarmBeatsModelsClientBeginUpdateOptions`
- New struct `FarmBeatsModelsClientGetOperationResultOptions`
- New struct `FarmBeatsModelsClientGetOperationResultResponse`
- New struct `FarmBeatsUpdateProperties`
- New struct `Identity`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResult`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionsClient`
- New struct `PrivateEndpointConnectionsClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateOptions`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New struct `PrivateEndpointConnectionsClientDeleteResponse`
- New struct `PrivateEndpointConnectionsClientGetOptions`
- New struct `PrivateEndpointConnectionsClientGetResponse`
- New struct `PrivateEndpointConnectionsClientListByResourceOptions`
- New struct `PrivateEndpointConnectionsClientListByResourceResponse`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkResourcesClient`
- New struct `PrivateLinkResourcesClientGetOptions`
- New struct `PrivateLinkResourcesClientGetResponse`
- New struct `PrivateLinkResourcesClientListByResourceOptions`
- New struct `PrivateLinkResourcesClientListByResourceResponse`
- New struct `PrivateLinkServiceConnectionState`
- New struct `SensorIntegration`
- New field `SystemData` in struct `TrackedResource`
- New field `Identity` in struct `FarmBeats`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `ProxyResource`
- New field `Identity` in struct `FarmBeatsUpdateRequestModel`
- New field `Properties` in struct `FarmBeatsUpdateRequestModel`
- New field `PrivateEndpointConnections` in struct `FarmBeatsProperties`
- New field `PublicNetworkAccess` in struct `FarmBeatsProperties`
- New field `SensorIntegration` in struct `FarmBeatsProperties`


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).