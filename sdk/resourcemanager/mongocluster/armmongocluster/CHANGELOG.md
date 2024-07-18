# Release History

## 0.2.0 (2024-07-18)
### Breaking Changes

- Function `*MongoClustersClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, Update, *MongoClustersClientBeginUpdateOptions)` to `(context.Context, string, string, MongoCluster, *MongoClustersClientBeginUpdateOptions)`
- Function `*FirewallRulesClient.NewListByMongoClusterPager` has been removed
- Function `*PrivateEndpointConnectionsClient.NewListByMongoClusterPager` has been removed
- Function `*PrivateLinksClient.NewListByMongoClusterPager` has been removed
- Struct `Update` has been removed
- Struct `UpdateProperties` has been removed

### Features Added

- New function `*FirewallRulesClient.NewListByParentPager(string, string, *FirewallRulesClientListByParentOptions) *runtime.Pager[FirewallRulesClientListByParentResponse]`
- New function `*PrivateEndpointConnectionsClient.NewListConnectionsPager(string, string, *PrivateEndpointConnectionsClientListConnectionsOptions) *runtime.Pager[PrivateEndpointConnectionsClientListConnectionsResponse]`
- New function `*PrivateLinksClient.NewListPager(string, string, *PrivateLinksClientListOptions) *runtime.Pager[PrivateLinksClientListResponse]`


## 0.1.0 (2024-07-05)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mongocluster/armmongocluster` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).