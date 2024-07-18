# Release History

## 2.0.0-beta.1 (2024-07-18)
### Breaking Changes

- `ManagedServiceIdentityTypeSystemAssignedUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Enum `ActionType` has been removed
- Enum `Origin` has been removed
- Function `*ClientFactory.NewLoadTestsClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewQuotasClient` has been removed
- Function `NewLoadTestsClient` has been removed
- Function `*LoadTestsClient.BeginCreateOrUpdate` has been removed
- Function `*LoadTestsClient.BeginDelete` has been removed
- Function `*LoadTestsClient.Get` has been removed
- Function `*LoadTestsClient.NewListByResourceGroupPager` has been removed
- Function `*LoadTestsClient.NewListBySubscriptionPager` has been removed
- Function `*LoadTestsClient.NewListOutboundNetworkDependenciesEndpointsPager` has been removed
- Function `*LoadTestsClient.BeginUpdate` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewQuotasClient` has been removed
- Function `*QuotasClient.CheckAvailability` has been removed
- Function `*QuotasClient.Get` has been removed
- Function `*QuotasClient.NewListPager` has been removed
- Struct `LoadTestResourcePageList` has been removed
- Struct `LoadTestResourcePatchRequestBody` has been removed
- Struct `LoadTestResourcePatchRequestBodyProperties` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed
- Struct `QuotaResourceList` has been removed
- Field `ID`, `Name`, `SystemData`, `Type` of struct `QuotaBucketRequest` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAndUserAssigned` added to enum type `ManagedServiceIdentityType`
- New function `*ClientFactory.NewLoadTestMgmtClient() *LoadTestMgmtClient`
- New function `NewLoadTestMgmtClient(string, azcore.TokenCredential, *arm.ClientOptions) (*LoadTestMgmtClient, error)`
- New function `*LoadTestMgmtClient.CheckAvailabilityQuota(context.Context, string, string, QuotaBucketRequest, *LoadTestMgmtClientCheckAvailabilityQuotaOptions) (LoadTestMgmtClientCheckAvailabilityQuotaResponse, error)`
- New function `*LoadTestMgmtClient.BeginCreateOrUpdateLoadtest(context.Context, string, string, LoadTestResource, *LoadTestMgmtClientBeginCreateOrUpdateLoadtestOptions) (*runtime.Poller[LoadTestMgmtClientCreateOrUpdateLoadtestResponse], error)`
- New function `*LoadTestMgmtClient.BeginDeleteLoadtest(context.Context, string, string, *LoadTestMgmtClientBeginDeleteLoadtestOptions) (*runtime.Poller[LoadTestMgmtClientDeleteLoadtestResponse], error)`
- New function `*LoadTestMgmtClient.GetLoadtest(context.Context, string, string, *LoadTestMgmtClientGetLoadtestOptions) (LoadTestMgmtClientGetLoadtestResponse, error)`
- New function `*LoadTestMgmtClient.GetQuota(context.Context, string, string, *LoadTestMgmtClientGetQuotaOptions) (LoadTestMgmtClientGetQuotaResponse, error)`
- New function `*LoadTestMgmtClient.NewListByResourceGroupPager(string, *LoadTestMgmtClientListByResourceGroupOptions) *runtime.Pager[LoadTestMgmtClientListByResourceGroupResponse]`
- New function `*LoadTestMgmtClient.NewListBySubscriptionPager(*LoadTestMgmtClientListBySubscriptionOptions) *runtime.Pager[LoadTestMgmtClientListBySubscriptionResponse]`
- New function `*LoadTestMgmtClient.NewListQuotaPager(string, *LoadTestMgmtClientListQuotaOptions) *runtime.Pager[LoadTestMgmtClientListQuotaResponse]`
- New function `*LoadTestMgmtClient.NewOutboundNetworkDependenciesEndpointsPager(string, string, *LoadTestMgmtClientOutboundNetworkDependenciesEndpointsOptions) *runtime.Pager[LoadTestMgmtClientOutboundNetworkDependenciesEndpointsResponse]`
- New function `*LoadTestMgmtClient.BeginUpdateLoadtest(context.Context, string, string, LoadTestResource, *LoadTestMgmtClientBeginUpdateLoadtestOptions) (*runtime.Poller[LoadTestMgmtClientUpdateLoadtestResponse], error)`
- New struct `LoadTestResourceListResult`
- New struct `QuotaResourceListResult`


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