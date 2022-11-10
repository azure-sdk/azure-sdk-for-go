# Release History

## 2.0.0 (2022-11-10)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `NodeCommunicationModeSimplified`
- New const `NodeCommunicationModeClassic`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `EndpointAccessDefaultActionAllow`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `PrivateEndpointConnectionProvisioningStateCancelled`
- New const `NodeCommunicationModeDefault`
- New const `EndpointAccessDefaultActionDeny`
- New type alias `NodeCommunicationMode`
- New type alias `EndpointAccessDefaultAction`
- New function `*PrivateEndpointConnectionClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionClientDeleteResponse], error)`
- New function `PossibleEndpointAccessDefaultActionValues() []EndpointAccessDefaultAction`
- New function `PossibleNodeCommunicationModeValues() []NodeCommunicationMode`
- New struct `EndpointAccessProfile`
- New struct `IPRule`
- New struct `NetworkProfile`
- New struct `PrivateEndpointConnectionClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionClientDeleteResponse`
- New field `TargetNodeCommunicationMode` in struct `PoolProperties`
- New field `CurrentNodeCommunicationMode` in struct `PoolProperties`
- New field `NetworkProfile` in struct `AccountUpdateProperties`
- New field `PublicNetworkAccess` in struct `AccountUpdateProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`
- New field `NetworkProfile` in struct `AccountCreateProperties`
- New field `NodeManagementEndpoint` in struct `AccountProperties`
- New field `NetworkProfile` in struct `AccountProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).