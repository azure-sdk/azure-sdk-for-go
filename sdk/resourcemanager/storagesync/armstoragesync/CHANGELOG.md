# Release History

## 2.0.0 (2022-08-01)
### Breaking Changes

- Function `*ServerEndpointsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, *ServerEndpointsClientBeginUpdateOptions)` to `(context.Context, string, string, string, string, ServerEndpointUpdateParameters, *ServerEndpointsClientBeginUpdateOptions)`
- Function `*ServicesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, *ServicesClientBeginUpdateOptions)` to `(context.Context, string, string, ServiceUpdateParameters, *ServicesClientBeginUpdateOptions)`
- Field `Parameters` of struct `ServicesClientBeginUpdateOptions` has been removed
- Field `Parameters` of struct `ServerEndpointsClientBeginUpdateOptions` has been removed

### Features Added

- New const `ManagedServiceIdentityTypeUserAssigned`
- New const `ManagedServiceIdentityTypeNone`
- New const `ManagedServiceIdentityTypeSystemAssigned`
- New const `ManagedServiceIdentityTypeSystemAssignedUserAssigned`
- New const `CloudTieringLowDiskModeStateEnabled`
- New const `CloudTieringLowDiskModeStateDisabled`
- New function `PossibleCloudTieringLowDiskModeStateValues() []CloudTieringLowDiskModeState`
- New function `PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType`
- New function `*RegisteredServersClient.BeginUpdate(context.Context, string, string, string, RegisteredServerUpdateParameters, *RegisteredServersClientBeginUpdateOptions) (*runtime.Poller[RegisteredServersClientUpdateResponse], error)`
- New struct `CloudTieringLowDiskMode`
- New struct `ManagedServiceIdentity`
- New struct `RegisteredServerUpdateParameters`
- New struct `RegisteredServerUpdateParametersProperties`
- New struct `RegisteredServersClientBeginUpdateOptions`
- New struct `RegisteredServersClientUpdateResponse`
- New struct `UserAssignedIdentity`
- New field `IdentityID` in struct `ServiceCreateParametersProperties`
- New field `IdentityType` in struct `ServiceCreateParametersProperties`
- New field `IdentityType` in struct `RegisteredServerCreateParametersProperties`
- New field `IdentityID` in struct `RegisteredServerCreateParametersProperties`
- New field `IdentityType` in struct `RegisteredServerProperties`
- New field `IdentityID` in struct `RegisteredServerProperties`
- New field `LowDiskMode` in struct `ServerEndpointCloudTieringStatus`
- New field `Identity` in struct `ServiceProperties`
- New field `IdentityType` in struct `ServiceProperties`
- New field `IdentityID` in struct `ServiceProperties`
- New field `IdentityID` in struct `ServiceUpdateProperties`
- New field `IdentityType` in struct `ServiceUpdateProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).