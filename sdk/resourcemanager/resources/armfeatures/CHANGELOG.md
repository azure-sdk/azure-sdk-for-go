# Release History

## 2.0.0 (2025-03-26)
### Breaking Changes

- `SubscriptionFeatureRegistrationStateRegistering`, `SubscriptionFeatureRegistrationStateUnregistering` from enum `SubscriptionFeatureRegistrationState` has been removed
- Function `NewClient` has been removed
- Function `*Client.Get` has been removed
- Function `*Client.NewListAllPager` has been removed
- Function `*Client.NewListPager` has been removed
- Function `*Client.Register` has been removed
- Function `*Client.Unregister` has been removed
- Function `*ClientFactory.NewClient` has been removed
- Function `*ClientFactory.NewFeatureClient` has been removed
- Function `NewFeatureClient` has been removed
- Function `*FeatureClient.NewListOperationsPager` has been removed
- Struct `FeatureOperationsListResult` has been removed
- Struct `FeatureProperties` has been removed
- Struct `FeatureResult` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module

## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armfeatures` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).