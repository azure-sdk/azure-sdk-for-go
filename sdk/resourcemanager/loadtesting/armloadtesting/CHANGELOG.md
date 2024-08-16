# Release History

## 2.0.0 (2024-08-16)
### Breaking Changes

- Function `*LoadTestsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, LoadTestResourcePatchRequestBody, *LoadTestsClientBeginUpdateOptions)` to `(context.Context, string, string, LoadTestResourceUpdate, *LoadTestsClientBeginUpdateOptions)`
- Struct `LoadTestResourcePageList` has been removed
- Struct `LoadTestResourcePatchRequestBody` has been removed
- Struct `LoadTestResourcePatchRequestBodyProperties` has been removed
- Struct `OutboundEnvironmentEndpointCollection` has been removed
- Struct `QuotaResourceList` has been removed
- Field `LoadTestResourcePageList` of struct `LoadTestsClientListByResourceGroupResponse` has been removed
- Field `LoadTestResourcePageList` of struct `LoadTestsClientListBySubscriptionResponse` has been removed
- Field `OutboundEnvironmentEndpointCollection` of struct `LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse` has been removed
- Field `ID`, `Name`, `SystemData`, `Type` of struct `QuotaBucketRequest` has been removed
- Field `QuotaResourceList` of struct `QuotasClientListResponse` has been removed

### Features Added

- New struct `LoadTestResourceListResult`
- New struct `LoadTestResourceUpdate`
- New struct `LoadTestResourceUpdateProperties`
- New struct `PagedOutboundEnvironmentEndpoint`
- New struct `QuotaResourceListResult`
- New anonymous field `LoadTestResourceListResult` in struct `LoadTestsClientListByResourceGroupResponse`
- New anonymous field `LoadTestResourceListResult` in struct `LoadTestsClientListBySubscriptionResponse`
- New anonymous field `PagedOutboundEnvironmentEndpoint` in struct `LoadTestsClientListOutboundNetworkDependenciesEndpointsResponse`
- New anonymous field `QuotaResourceListResult` in struct `QuotasClientListResponse`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.2 (2023-02-24)

Release stable version.

## 0.1.0 (2022-11-21)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/loadtesting/armloadtesting` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).