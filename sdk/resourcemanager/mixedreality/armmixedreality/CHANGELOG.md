# Release History

## 0.8.0 (2025-01-15)
### Breaking Changes

- Function `*ClientFactory.NewObjectAnchorsAccountsClient` has been removed
- Function `*ClientFactory.NewSpatialAnchorsAccountsClient` has been removed
- Function `NewObjectAnchorsAccountsClient` has been removed
- Function `*ObjectAnchorsAccountsClient.Create` has been removed
- Function `*ObjectAnchorsAccountsClient.Delete` has been removed
- Function `*ObjectAnchorsAccountsClient.Get` has been removed
- Function `*ObjectAnchorsAccountsClient.NewListByResourceGroupPager` has been removed
- Function `*ObjectAnchorsAccountsClient.NewListBySubscriptionPager` has been removed
- Function `*ObjectAnchorsAccountsClient.ListKeys` has been removed
- Function `*ObjectAnchorsAccountsClient.RegenerateKeys` has been removed
- Function `*ObjectAnchorsAccountsClient.Update` has been removed
- Function `NewSpatialAnchorsAccountsClient` has been removed
- Function `*SpatialAnchorsAccountsClient.Create` has been removed
- Function `*SpatialAnchorsAccountsClient.Delete` has been removed
- Function `*SpatialAnchorsAccountsClient.Get` has been removed
- Function `*SpatialAnchorsAccountsClient.NewListByResourceGroupPager` has been removed
- Function `*SpatialAnchorsAccountsClient.NewListBySubscriptionPager` has been removed
- Function `*SpatialAnchorsAccountsClient.ListKeys` has been removed
- Function `*SpatialAnchorsAccountsClient.RegenerateKeys` has been removed
- Function `*SpatialAnchorsAccountsClient.Update` has been removed
- Struct `ObjectAnchorsAccount` has been removed
- Struct `ObjectAnchorsAccountIdentity` has been removed
- Struct `ObjectAnchorsAccountPage` has been removed
- Struct `SpatialAnchorsAccount` has been removed
- Struct `SpatialAnchorsAccountPage` has been removed


## 0.7.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mixedreality/armmixedreality` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).