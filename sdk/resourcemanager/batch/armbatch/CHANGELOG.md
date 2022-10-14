# Release History

## 2.0.0 (2022-10-14)
### Breaking Changes

- Struct `CloudError` has been removed
- Struct `CloudErrorBody` has been removed

### Features Added

- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `NodeCommunicationModeDefault`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `NodeCommunicationModeClassic`
- New const `PrivateEndpointConnectionProvisioningStateCancelled`
- New const `NodeCommunicationModeSimplified`
- New const `EndpointAccessDefaultActionAllow`
- New const `EndpointAccessDefaultActionDeny`
- New type alias `NodeCommunicationMode`
- New type alias `EndpointAccessDefaultAction`
- New function `PossibleEndpointAccessDefaultActionValues() []EndpointAccessDefaultAction`
- New function `*PrivateEndpointConnectionClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionClientDeleteResponse], error)`
- New function `PossibleNodeCommunicationModeValues() []NodeCommunicationMode`
- New struct `EndpointAccessProfile`
- New struct `IPRule`
- New struct `NetworkProfile`
- New struct `PrivateEndpointConnectionClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionClientDeleteResponse`
- New field `NetworkProfile` in struct `AccountProperties`
- New field `NodeManagementEndpoint` in struct `AccountProperties`
- New field `NetworkProfile` in struct `AccountCreateProperties`
- New field `NetworkProfile` in struct `AccountUpdateProperties`
- New field `PublicNetworkAccess` in struct `AccountUpdateProperties`
- New field `TargetNodeCommunicationMode` in struct `PoolProperties`
- New field `CurrentNodeCommunicationMode` in struct `PoolProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).