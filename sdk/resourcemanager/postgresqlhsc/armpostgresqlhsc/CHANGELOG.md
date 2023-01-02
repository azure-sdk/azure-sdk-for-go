# Release History

## 1.0.0 (2023-01-02)
### Breaking Changes

- Function `*ConfigurationsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, ServerGroupConfiguration, *ConfigurationsClientBeginUpdateOptions)` to `(context.Context, string, string, string, ServerConfigurationRequest, *ConfigurationsClientBeginUpdateOptions)`
- Type of `ServerProperties.ServerEdition` has been changed from `*ServerEdition` to `*string`
- Type alias `CitusVersion` has been removed
- Type alias `CreateMode` has been removed
- Type alias `PostgreSQLVersion` has been removed
- Type alias `ResourceProviderType` has been removed
- Type alias `ServerEdition` has been removed
- Type alias `ServerHaState` has been removed
- Type alias `ServerState` has been removed
- Function `*ConfigurationsClient.Get` has been removed
- Function `*ConfigurationsClient.NewListByServerGroupPager` has been removed
- Function `*FirewallRulesClient.NewListByServerGroupPager` has been removed
- Function `*RolesClient.NewListByServerGroupPager` has been removed
- Function `NewServerGroupsClient` has been removed
- Function `*ServerGroupsClient.CheckNameAvailability` has been removed
- Function `*ServerGroupsClient.BeginCreateOrUpdate` has been removed
- Function `*ServerGroupsClient.BeginDelete` has been removed
- Function `*ServerGroupsClient.Get` has been removed
- Function `*ServerGroupsClient.NewListByResourceGroupPager` has been removed
- Function `*ServerGroupsClient.NewListPager` has been removed
- Function `*ServerGroupsClient.BeginRestart` has been removed
- Function `*ServerGroupsClient.BeginStart` has been removed
- Function `*ServerGroupsClient.BeginStop` has been removed
- Function `*ServerGroupsClient.BeginUpdate` has been removed
- Function `*ServersClient.NewListByServerGroupPager` has been removed
- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed
- Struct `ConfigurationsClientListByServerGroupResponse` has been removed
- Struct `FirewallRulesClientListByServerGroupResponse` has been removed
- Struct `RolesClientListByServerGroupResponse` has been removed
- Struct `ServerGroup` has been removed
- Struct `ServerGroupConfiguration` has been removed
- Struct `ServerGroupConfigurationListResult` has been removed
- Struct `ServerGroupConfigurationProperties` has been removed
- Struct `ServerGroupForUpdate` has been removed
- Struct `ServerGroupListResult` has been removed
- Struct `ServerGroupProperties` has been removed
- Struct `ServerGroupPropertiesDelegatedSubnetArguments` has been removed
- Struct `ServerGroupPropertiesForUpdate` has been removed
- Struct `ServerGroupPropertiesPrivateDNSZoneArguments` has been removed
- Struct `ServerGroupServer` has been removed
- Struct `ServerGroupServerListResult` has been removed
- Struct `ServerGroupServerProperties` has been removed
- Struct `ServerGroupsClient` has been removed
- Struct `ServerGroupsClientCreateOrUpdateResponse` has been removed
- Struct `ServerGroupsClientDeleteResponse` has been removed
- Struct `ServerGroupsClientListByResourceGroupResponse` has been removed
- Struct `ServerGroupsClientListResponse` has been removed
- Struct `ServerGroupsClientRestartResponse` has been removed
- Struct `ServerGroupsClientStartResponse` has been removed
- Struct `ServerGroupsClientStopResponse` has been removed
- Struct `ServerGroupsClientUpdateResponse` has been removed
- Struct `ServerRoleGroup` has been removed
- Struct `ServersClientListByServerGroupResponse` has been removed
- Field `ServerGroupConfiguration` of struct `ConfigurationsClientUpdateResponse` has been removed
- Field `ServerGroupServer` of struct `ServersClientGetResponse` has been removed

### Features Added

- New function `NewClusterClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ClusterClient, error)`
- New function `*ClusterClient.BeginPromoteReadReplica(context.Context, string, string, *ClusterClientBeginPromoteReadReplicaOptions) (*runtime.Poller[ClusterClientPromoteReadReplicaResponse], error)`
- New function `*ClusterClient.BeginRestart(context.Context, string, string, *ClusterClientBeginRestartOptions) (*runtime.Poller[ClusterClientRestartResponse], error)`
- New function `*ClusterClient.BeginStart(context.Context, string, string, *ClusterClientBeginStartOptions) (*runtime.Poller[ClusterClientStartResponse], error)`
- New function `*ClusterClient.BeginStop(context.Context, string, string, *ClusterClientBeginStopOptions) (*runtime.Poller[ClusterClientStopResponse], error)`
- New function `NewClustersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ClustersClient, error)`
- New function `*ClustersClient.CheckNameAvailability(context.Context, NameAvailabilityRequest, *ClustersClientCheckNameAvailabilityOptions) (ClustersClientCheckNameAvailabilityResponse, error)`
- New function `*ClustersClient.BeginCreateOrUpdate(context.Context, string, string, Cluster, *ClustersClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClustersClientCreateOrUpdateResponse], error)`
- New function `*ClustersClient.BeginDelete(context.Context, string, string, *ClustersClientBeginDeleteOptions) (*runtime.Poller[ClustersClientDeleteResponse], error)`
- New function `*ClustersClient.Get(context.Context, string, string, *ClustersClientGetOptions) (ClustersClientGetResponse, error)`
- New function `*ClustersClient.NewListByResourceGroupPager(string, *ClustersClientListByResourceGroupOptions) *runtime.Pager[ClustersClientListByResourceGroupResponse]`
- New function `*ClustersClient.NewListPager(*ClustersClientListOptions) *runtime.Pager[ClustersClientListResponse]`
- New function `*ClustersClient.BeginUpdate(context.Context, string, string, ClusterForUpdate, *ClustersClientBeginUpdateOptions) (*runtime.Poller[ClustersClientUpdateResponse], error)`
- New function `NewConfigurationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConfigurationClient, error)`
- New function `*ConfigurationClient.GetCoordinator(context.Context, string, string, string, *ConfigurationClientGetCoordinatorOptions) (ConfigurationClientGetCoordinatorResponse, error)`
- New function `*ConfigurationClient.GetNode(context.Context, string, string, string, *ConfigurationClientGetNodeOptions) (ConfigurationClientGetNodeResponse, error)`
- New function `*ConfigurationsClient.NewListByClusterPager(string, string, *ConfigurationsClientListByClusterOptions) *runtime.Pager[ConfigurationsClientListByClusterResponse]`
- New function `*ConfigurationsClient.BeginUpdateNode(context.Context, string, string, string, ServerConfigurationRequest, *ConfigurationsClientBeginUpdateNodeOptions) (*runtime.Poller[ConfigurationsClientUpdateNodeResponse], error)`
- New function `*FirewallRulesClient.NewListByClusterPager(string, string, *FirewallRulesClientListByClusterOptions) *runtime.Pager[FirewallRulesClientListByClusterResponse]`
- New function `NewPrivateEndpointConnectionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error)`
- New function `*PrivateEndpointConnectionsClient.NewCreateOrUpdatePager(string, string, string, PrivateEndpointConnectionRequest, *PrivateEndpointConnectionsClientCreateOrUpdateOptions) *runtime.Pager[PrivateEndpointConnectionsClientCreateOrUpdateResponse]`
- New function `*PrivateEndpointConnectionsClient.NewDeletePager(string, string, string, *PrivateEndpointConnectionsClientDeleteOptions) *runtime.Pager[PrivateEndpointConnectionsClientDeleteResponse]`
- New function `*PrivateEndpointConnectionsClient.NewGetPager(string, string, string, *PrivateEndpointConnectionsClientGetOptions) *runtime.Pager[PrivateEndpointConnectionsClientGetResponse]`
- New function `*PrivateEndpointConnectionsClient.NewListByClusterPager(string, string, *PrivateEndpointConnectionsClientListByClusterOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByClusterResponse]`
- New function `NewPrivateLinkResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrivateLinkResourcesClient, error)`
- New function `*PrivateLinkResourcesClient.NewGetPager(string, string, string, *PrivateLinkResourcesClientGetOptions) *runtime.Pager[PrivateLinkResourcesClientGetResponse]`
- New function `*PrivateLinkResourcesClient.NewListByClusterPager(string, string, *PrivateLinkResourcesClientListByClusterOptions) *runtime.Pager[PrivateLinkResourcesClientListByClusterResponse]`
- New function `*RolesClient.NewListByClusterPager(string, string, *RolesClientListByClusterOptions) *runtime.Pager[RolesClientListByClusterResponse]`
- New function `*ServersClient.NewListByClusterPager(string, string, *ServersClientListByClusterOptions) *runtime.Pager[ServersClientListByClusterResponse]`
- New struct `Cluster`
- New struct `ClusterClient`
- New struct `ClusterClientPromoteReadReplicaResponse`
- New struct `ClusterClientRestartResponse`
- New struct `ClusterClientStartResponse`
- New struct `ClusterClientStopResponse`
- New struct `ClusterConfiguration`
- New struct `ClusterConfigurationListResult`
- New struct `ClusterConfigurationProperties`
- New struct `ClusterForUpdate`
- New struct `ClusterListResult`
- New struct `ClusterProperties`
- New struct `ClusterPropertiesForUpdate`
- New struct `ClusterResponse`
- New struct `ClusterResponseProperties`
- New struct `ClusterServer`
- New struct `ClusterServerListResult`
- New struct `ClusterServerProperties`
- New struct `ClustersClient`
- New struct `ClustersClientCreateOrUpdateResponse`
- New struct `ClustersClientDeleteResponse`
- New struct `ClustersClientListByResourceGroupResponse`
- New struct `ClustersClientListResponse`
- New struct `ClustersClientUpdateResponse`
- New struct `ConfigurationClient`
- New struct `ConfigurationsClientListByClusterResponse`
- New struct `ConfigurationsClientUpdateNodeResponse`
- New struct `FirewallRulesClientListByClusterResponse`
- New struct `PrivateEndpointConnection`
- New struct `PrivateEndpointConnectionListResult`
- New struct `PrivateEndpointConnectionProperties`
- New struct `PrivateEndpointConnectionProperty`
- New struct `PrivateEndpointConnectionRequest`
- New struct `PrivateEndpointConnectionRequestProperties`
- New struct `PrivateEndpointConnectionSimpleProperties`
- New struct `PrivateEndpointConnectionsClient`
- New struct `PrivateEndpointConnectionsClientCreateOrUpdateResponse`
- New struct `PrivateEndpointConnectionsClientDeleteResponse`
- New struct `PrivateEndpointConnectionsClientGetResponse`
- New struct `PrivateEndpointConnectionsClientListByClusterResponse`
- New struct `PrivateEndpointProperty`
- New struct `PrivateLinkResource`
- New struct `PrivateLinkResourceListResult`
- New struct `PrivateLinkResourceProperties`
- New struct `PrivateLinkResourcesClient`
- New struct `PrivateLinkResourcesClientGetResponse`
- New struct `PrivateLinkResourcesClientListByClusterResponse`
- New struct `PrivateLinkServiceConnectionStateProperty`
- New struct `RolesClientListByClusterResponse`
- New struct `ServerConfigurationRequest`
- New struct `ServersClientListByClusterResponse`
- New anonymous field `ServerConfiguration` in struct `ConfigurationsClientUpdateResponse`
- New field `RequiresRestart` in struct `ServerConfigurationProperties`
- New anonymous field `ClusterServer` in struct `ServersClientGetResponse`


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/postgresqlhsc/armpostgresqlhsc` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).