# Release History

## 2.0.0 (2025-03-07)
### Breaking Changes

- Function `NewClientFactory` parameter(s) have been changed from `(string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`
- Function `*ClientFactory.NewEndpointsClient` parameter(s) have been changed from `(<none>)` to `(string)`
- Function `*ClientFactory.NewHeatMapClient` parameter(s) have been changed from `(<none>)` to `(string)`
- Function `*ClientFactory.NewProfilesClient` parameter(s) have been changed from `(<none>)` to `(string)`
- Function `*ClientFactory.NewUserMetricsKeysClient` parameter(s) have been changed from `(<none>)` to `(string)`


## 1.3.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.2.0 (2023-06-23)
### Features Added

- New value `EndpointMonitorStatusUnmonitored` added to enum type `EndpointMonitorStatus`
- New enum type `AlwaysServe` with values `AlwaysServeDisabled`, `AlwaysServeEnabled`
- New function `*ProfilesClient.CheckTrafficManagerNameAvailabilityV2(context.Context, CheckTrafficManagerRelativeDNSNameAvailabilityParameters, *ProfilesClientCheckTrafficManagerNameAvailabilityV2Options) (ProfilesClientCheckTrafficManagerNameAvailabilityV2Response, error)`
- New field `AlwaysServe` in struct `EndpointProperties`


## 1.1.0 (2023-04-07)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).