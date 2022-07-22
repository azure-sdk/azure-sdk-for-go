# Release History

## 2.0.0 (2022-07-22)
### Breaking Changes

- Function `*ApplicationClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *ApplicationClientCreateOptions)` to `(context.Context, string, string, string, Application, *ApplicationClientCreateOptions)`
- Function `*ApplicationPackageClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, string, *ApplicationPackageClientCreateOptions)` to `(context.Context, string, string, string, string, ApplicationPackage, *ApplicationPackageClientCreateOptions)`
- Field `Parameters` of struct `ApplicationClientCreateOptions` has been removed
- Field `Parameters` of struct `ApplicationPackageClientCreateOptions` has been removed

### Features Added

- New const `PrivateEndpointConnectionProvisioningStateDeleting`
- New const `PrivateEndpointConnectionProvisioningStateCreating`
- New const `EndpointAccessDefaultActionDeny`
- New const `EndpointAccessDefaultActionAllow`
- New const `PrivateEndpointConnectionProvisioningStateCancelled`
- New function `PossibleEndpointAccessDefaultActionValues() []EndpointAccessDefaultAction`
- New function `*PrivateEndpointConnectionClient.BeginDelete(context.Context, string, string, string, *PrivateEndpointConnectionClientBeginDeleteOptions) (*runtime.Poller[PrivateEndpointConnectionClientDeleteResponse], error)`
- New struct `EndpointAccessProfile`
- New struct `IPRule`
- New struct `NetworkProfile`
- New struct `PrivateEndpointConnectionClientBeginDeleteOptions`
- New struct `PrivateEndpointConnectionClientDeleteResponse`
- New field `NetworkProfile` in struct `AccountCreateProperties`
- New field `GroupIDs` in struct `PrivateEndpointConnectionProperties`
- New field `NodeManagementEndpoint` in struct `AccountProperties`
- New field `NetworkProfile` in struct `AccountProperties`
- New field `NetworkProfile` in struct `AccountUpdateProperties`
- New field `PublicNetworkAccess` in struct `AccountUpdateProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/batch/armbatch` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).