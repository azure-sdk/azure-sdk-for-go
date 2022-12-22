# Release History

## 1.0.0 (2022-12-22)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateEndpointConnectionsClient.NewCreateOrUpdatePager(string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionsClientCreateOrUpdateOptions) *runtime.Pager[PrivateEndpointConnectionsClientCreateOrUpdateResponse]`
- New function `*PrivateEndpointConnectionsClient.NewDeletePager(string, string, string, *PrivateEndpointConnectionsClientDeleteOptions) *runtime.Pager[PrivateEndpointConnectionsClientDeleteResponse]`
- New function `*PrivateEndpointConnectionsClient.NewGetPager(string, string, string, *PrivateEndpointConnectionsClientGetOptions) *runtime.Pager[PrivateEndpointConnectionsClientGetResponse]`
- New function `*PrivateEndpointConnectionsClient.NewListByServerGroupPager(string, string, *PrivateEndpointConnectionsClientListByServerGroupOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByServerGroupResponse]`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.NewGetPager(string, string, string, *PrivateLinkResourcesClientGetOptions) *runtime.Pager[PrivateLinkResourcesClientGetResponse]`
- New function `*PrivateLinkResourcesClient.NewListByServerGroupPager(string, string, *PrivateLinkResourcesClientListByServerGroupOptions) *runtime.Pager[PrivateLinkResourcesClientListByServerGroupResponse]`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResult`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionsClient`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New struct `PrivateEndpointConnectionsClientDeleteResponse`
- New struct `PrivateEndpointConnectionsClientGetResponse`
- New struct `PrivateEndpointConnectionsClientListByServerGroupResponse`
- New struct `PrivateEndpointProperty`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkResourcesClient`
- New struct `PrivateLinkResourcesClientGetResponse`
- New struct `PrivateLinkResourcesClientListByServerGroupResponse`
- New struct `PrivateLinkServiceConnectionStateProperty`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).