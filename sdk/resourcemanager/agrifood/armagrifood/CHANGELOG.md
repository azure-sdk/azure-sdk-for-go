# Release History

## 0.8.0 (2022-09-13)
### Breaking Changes

- Const `PrivateEndpointConnectionProvisioningStateDeleting` has been removed
- Const `ProvisioningStateCreating` has been removed
- Const `ProvisioningStateDeleting` has been removed
- Const `PublicNetworkAccessHybrid` has been removed
- Const `PrivateEndpointConnectionProvisioningStateSucceeded` has been removed
- Const `ProvisioningStateUpdating` has been removed
- Const `PrivateEndpointServiceConnectionStatusApproved` has been removed
- Const `PrivateEndpointServiceConnectionStatusPending` has been removed
- Const `PublicNetworkAccessEnabled` has been removed
- Const `PrivateEndpointServiceConnectionStatusRejected` has been removed
- Const `PrivateEndpointConnectionProvisioningStateCreating` has been removed
- Const `PrivateEndpointConnectionProvisioningStateFailed` has been removed
- Type alias `PrivateEndpointConnectionProvisioningState` has been removed
- Type alias `PrivateEndpointServiceConnectionStatus` has been removed
- Type alias `PublicNetworkAccess` has been removed
- Function `*PrivateLinkResourcesClient.Get` has been removed
- Function `*PrivateLinkResourcesClient.NewListByResourcePager` has been removed
- Function `NewPrivateEndpointConnectionsClient` has been removed
- Function `PossiblePrivateEndpointConnectionProvisioningStateValues` has been removed
- Function `*PrivateEndpointConnectionsClient.BeginDelete` has been removed
- Function `*PrivateEndpointConnectionsClient.Get` has been removed
- Function `PossiblePrivateEndpointServiceConnectionStatusValues` has been removed
- Function `NewPrivateLinkResourcesClient` has been removed
- Function `PossiblePublicNetworkAccessValues` has been removed
- Function `*FarmBeatsModelsClient.GetOperationResult` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByResourcePager` has been removed
- Function `*FarmBeatsModelsClient.BeginUpdate` has been removed
- Function `*PrivateEndpointConnectionsClient.CreateOrUpdate` has been removed
- Struct `ArmAsyncOperation` has been removed
- Struct `FarmBeatsModelsClientBeginUpdateOptions` has been removed
- Struct `FarmBeatsModelsClientGetOperationResultOptions` has been removed
- Struct `FarmBeatsModelsClientGetOperationResultResponse` has been removed
- Struct `FarmBeatsUpdateProperties` has been removed
- Struct `Identity` has been removed
- Struct `PrivateEndpoint` has been removed
- Struct `PrivateEndpointConnection` has been removed
- Struct `PrivateEndpointConnectionListResult` has been removed
- Struct `PrivateEndpointConnectionProperties` has been removed
- Struct `PrivateEndpointConnectionsClient` has been removed
- Struct `PrivateEndpointConnectionsClientBeginDeleteOptions` has been removed
- Struct `PrivateEndpointConnectionsClientCreateOrUpdateOptions` has been removed
- Struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse` has been removed
- Struct `PrivateEndpointConnectionsClientDeleteResponse` has been removed
- Struct `PrivateEndpointConnectionsClientGetOptions` has been removed
- Struct `PrivateEndpointConnectionsClientGetResponse` has been removed
- Struct `PrivateEndpointConnectionsClientListByResourceOptions` has been removed
- Struct `PrivateEndpointConnectionsClientListByResourceResponse` has been removed
- Struct `PrivateLinkResource` has been removed
- Struct `PrivateLinkResourceListResult` has been removed
- Struct `PrivateLinkResourceProperties` has been removed
- Struct `PrivateLinkResourcesClient` has been removed
- Struct `PrivateLinkResourcesClientGetOptions` has been removed
- Struct `PrivateLinkResourcesClientGetResponse` has been removed
- Struct `PrivateLinkResourcesClientListByResourceOptions` has been removed
- Struct `PrivateLinkResourcesClientListByResourceResponse` has been removed
- Struct `PrivateLinkServiceConnectionState` has been removed
- Struct `SensorIntegration` has been removed
- Field `SystemData` of struct `Resource` has been removed
- Field `Identity` of struct `FarmBeatsUpdateRequestModel` has been removed
- Field `Properties` of struct `FarmBeatsUpdateRequestModel` has been removed
- Field `SystemData` of struct `TrackedResource` has been removed
- Field `SystemData` of struct `ProxyResource` has been removed
- Field `SensorIntegration` of struct `FarmBeatsProperties` has been removed
- Field `PrivateEndpointConnections` of struct `FarmBeatsProperties` has been removed
- Field `PublicNetworkAccess` of struct `FarmBeatsProperties` has been removed
- Field `Identity` of struct `FarmBeats` has been removed

### Features Added

- New function `*FarmBeatsModelsClient.Update(context.Context, string, string, FarmBeatsUpdateRequestModel, *FarmBeatsModelsClientUpdateOptions) (FarmBeatsModelsClientUpdateResponse, error)`
- New struct `FarmBeatsModelsClientUpdateOptions`


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
