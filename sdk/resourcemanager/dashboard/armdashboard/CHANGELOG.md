# Release History

## 1.0.0 (2022-08-02)
### Breaking Changes

- Type of `SystemData.LastModifiedByType` has been changed from `*LastModifiedByType` to `*CreatedByType`
- Type of `OperationListResult.Value` has been changed from `[]*OperationResult` to `[]*Operation`
- Type of `ManagedGrafana.Identity` has been changed from `*ManagedIdentity` to `*ManagedServiceIdentity`
- Type of `ManagedGrafanaUpdateParameters.Identity` has been changed from `*ManagedIdentity` to `*ManagedServiceIdentity`
- Const `IdentityTypeSystemAssigned` has been removed
- Const `LastModifiedByTypeManagedIdentity` has been removed
- Const `LastModifiedByTypeUser` has been removed
- Const `LastModifiedByTypeKey` has been removed
- Const `IdentityTypeNone` has been removed
- Const `LastModifiedByTypeApplication` has been removed
- Function `PossibleIdentityTypeValues` has been removed
- Function `PossibleLastModifiedByTypeValues` has been removed
- Struct `ManagedIdentity` has been removed
- Struct `OperationResult` has been removed

### Features Added

- New const `PrivateEndpointServiceConnectionStatusPending`
- New const `PrivateEndpointConnectionProvisioningStateFailed`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `ManagedServiceIdentityTypeSystemAssigned`
- New const `PublicNetworkAccessDisabled`
- New const `APIKeyEnabled`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `PrivateEndpointServiceConnectionStatusApproved`
- New const `ManagedServiceIdentityTypeSystemAssignedUserAssigned`
- New const `APIKeyDisabled`
- New const `PrivateEndpointServiceConnectionStatusRejected`
- New const `DeterministicOutboundIPDisabled`
- New const `DeterministicOutboundIPEnabled`
- New const `ManagedServiceIdentityTypeUserAssigned`
- New const `PrivateEndpointConnectionProvisioningStateSucceeded`
- New const `ManagedServiceIdentityTypeNone`
- New const `PublicNetworkAccessEnabled`
- New function `PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus`
- New function `PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState`
- New function `*PrivateLinkResourcesClient.NewListPager(string, string, *PrivateLinkResourcesClientListOptions) *runtime.Pager[PrivateLinkResourcesClientListResponse]`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `PossiblePublicNetworkAccessValues() []PublicNetworkAccess`
- New function `*PrivateEndpointConnectionsClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionsClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionsClientDeleteResponse], error)`
- New function `*PrivateLinkResourcesClient.Get(context.Context, string, string, string, *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error)`
- New function `*PrivateEndpointConnectionsClient.BeginApprove(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientBeginApproveOptions) (*runtime.Poller[PrivateEndpointConnectionsClientApproveResponse], error)`
- New function `PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType`
- New function `*PrivateEndpointConnectionsClient.NewListPager(string, string, *PrivateEndpointConnectionsClientListOptions) *runtime.Pager[PrivateEndpointConnectionsClientListResponse]`
- New function `PossibleDeterministicOutboundIPValues() []DeterministicOutboundIP`
- New function `PossibleAPIKeyValues() []APIKey`
- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateEndpointConnectionsClient.Get(context.Context, string, string, string, *PrivateEndpointConnectionsClientGetOptions) (PrivateEndpointConnectionsClientGetResponse, error)`
- New struct `AzureMonitorWorkspaceIntegration`
- New struct `GrafanaIntegrations`
- New struct `ManagedGrafanaPropertiesUpdateParameters`
- New struct `ManagedServiceIdentity`
- New struct `Operation`
- New struct `PrivateEndpoint`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResult`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionsClient`
- New struct `PrivateEndpointConnectionsClientApproveResponse`
- New struct `PrivateEndpointConnectionsClientBeginApproveOptions`
- New struct `PrivateEndpointConnectionsClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionsClientDeleteResponse`
- New struct `PrivateEndpointConnectionsClientGetOptions`
- New struct `PrivateEndpointConnectionsClientGetResponse`
- New struct `PrivateEndpointConnectionsClientListOptions`
- New struct `PrivateEndpointConnectionsClientListResponse`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkResourcesClient`
- New struct `PrivateLinkResourcesClientGetOptions`
- New struct `PrivateLinkResourcesClientGetResponse`
- New struct `PrivateLinkResourcesClientListOptions`
- New struct `PrivateLinkResourcesClientListResponse`
- New struct `PrivateLinkServiceConnectionState`
- New struct `Resource`
- New field `Properties` in struct `ManagedGrafanaUpdateParameters`
- New field `PublicNetworkAccess` in struct `ManagedGrafanaProperties`
- New field `OutboundIPs` in struct `ManagedGrafanaProperties`
- New field `APIKey` in struct `ManagedGrafanaProperties`
- New field `GrafanaIntegrations` in struct `ManagedGrafanaProperties`
- New field `PrivateEndpointConnections` in struct `ManagedGrafanaProperties`
- New field `DeterministicOutboundIP` in struct `ManagedGrafanaProperties`


## 0.3.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dashboard/armdashboard` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.3.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).