# Release History

## 2.0.0 (2022-07-08)
### Breaking Changes

- Function `*ServicesClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, *ServicesClientBeginUpdateOptions)` to `(context.Context, string, string, ServiceUpdateParameters, *ServicesClientBeginUpdateOptions)`
- Function `*ServerEndpointsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, *ServerEndpointsClientBeginUpdateOptions)` to `(context.Context, string, string, string, string, ServerEndpointUpdateParameters, *ServerEndpointsClientBeginUpdateOptions)`
- Field `Parameters` of struct `ServerEndpointsClientBeginUpdateOptions` has been removed
- Field `Parameters` of struct `ServicesClientBeginUpdateOptions` has been removed

### Features Added

- New const `CloudTieringLowDiskModeStateEnabled`
- New const `CloudTieringLowDiskModeStateDisabled`
- New function `PossibleCloudTieringLowDiskModeStateValues() []CloudTieringLowDiskModeState`
- New struct `CloudTieringLowDiskMode`
- New field `LowDiskMode` in struct `ServerEndpointCloudTieringStatus`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storagesync/armstoragesync` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).