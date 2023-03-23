# Release History

## 2.0.0 (2023-03-23)
### Breaking Changes

- Function `*KeyValuesClient.NewListByConfigurationStorePager` has been removed

### Features Added

- New enum type `ReplicaProvisioningState` with values `ReplicaProvisioningStateCanceled`, `ReplicaProvisioningStateCreating`, `ReplicaProvisioningStateDeleting`, `ReplicaProvisioningStateFailed`, `ReplicaProvisioningStateSucceeded`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewConfigurationStoresClient() *ConfigurationStoresClient`
- New function `*ClientFactory.NewKeyValuesClient() *KeyValuesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPrivateEndpointConnectionsClient() *PrivateEndpointConnectionsClient`
- New function `*ClientFactory.NewPrivateLinkResourcesClient() *PrivateLinkResourcesClient`
- New function `*ClientFactory.NewReplicasClient() *ReplicasClient`
- New function `NewReplicasClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ReplicasClient, error)`
- New function `*ReplicasClient.BeginCreate(context.Context, string, string, string, Replica, *ReplicasClientBeginCreateOptions) (*runtime.Poller[ReplicasClientCreateResponse], error)`
- New function `*ReplicasClient.BeginDelete(context.Context, string, string, string, *ReplicasClientBeginDeleteOptions) (*runtime.Poller[ReplicasClientDeleteResponse], error)`
- New function `*ReplicasClient.Get(context.Context, string, string, string, *ReplicasClientGetOptions) (ReplicasClientGetResponse, error)`
- New function `*ReplicasClient.NewListByConfigurationStorePager(string, string, *ReplicasClientListByConfigurationStoreOptions) *runtime.Pager[ReplicasClientListByConfigurationStoreResponse]`
- New struct `ClientFactory`
- New struct `Replica`
- New struct `ReplicaListResult`
- New struct `ReplicaProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).