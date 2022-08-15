# Release History

## 1.1.0 (2022-08-15)
### Features Added

- New const `PrivateEndpointConnectionProvisioningStateCancelled`
- New const `EndpointAccessDefaultActionAllow`
- New const `EndpointAccessDefaultActionDeny`
- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New function `PossibleEndpointAccessDefaultActionValues() []EndpointAccessDefaultAction`
- New function `*PrivateEndpointConnectionClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionClientDeleteResponse], error)`
- New struct `EndpointAccessProfile`
- New struct `IPRule`
- New struct `NetworkProfile`
- New struct `PrivateEndpointConnectionClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionClientDeleteResponse`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`
- New field `NodeManagementEndpoint` in struct `AccountProperties`
- New field `NetworkProfile` in struct `AccountProperties`
- New field `NetworkProfile` in struct `AccountCreateProperties`
- New field `NetworkProfile` in struct `AccountUpdateProperties`
- New field `PublicNetworkAccess` in struct `AccountUpdateProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).