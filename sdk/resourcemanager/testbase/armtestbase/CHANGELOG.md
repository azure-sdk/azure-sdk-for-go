# Release History

## 0.7.0 (2023-11-07)
### Breaking Changes

- Function `*AccountsClient.NewListBySubscriptionPager` has been removed

### Features Added

- New function `NewBillingHubServiceClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BillingHubServiceClient, error)`
- New function `*BillingHubServiceClient.GetFreeHourBalance(context.Context, string, string, *BillingHubServiceClientGetFreeHourBalanceOptions) (BillingHubServiceClientGetFreeHourBalanceResponse, error)`
- New function `*BillingHubServiceClient.GetUsage(context.Context, string, string, *BillingHubServiceClientGetUsageOptions) (BillingHubServiceClientGetUsageResponse, error)`
- New function `*ClientFactory.NewBillingHubServiceClient() *BillingHubServiceClient`
- New function `*PackagesClient.RunTest(context.Context, string, string, string, *PackagesClientRunTestOptions) (PackagesClientRunTestResponse, error)`
- New function `*TestResultsClient.GetConsoleLogDownloadURL(context.Context, string, string, string, string, TestResultConsoleLogDownloadURLParameters, *TestResultsClientGetConsoleLogDownloadURLOptions) (TestResultsClientGetConsoleLogDownloadURLResponse, error)`
- New struct `BillingHubExecutionUsageDetail`
- New struct `BillingHubFreeHourIncrementEntry`
- New struct `BillingHubGetFreeHourBalanceResponse`
- New struct `BillingHubGetUsageRequest`
- New struct `BillingHubGetUsageResponse`
- New struct `BillingHubPackageUsage`
- New struct `BillingHubUsage`
- New struct `BillingHubUsageGroup`
- New struct `BillingHubUsageGroupedByUpdateType`
- New struct `PackageRunTestParameters`
- New struct `TestResultConsoleLogDownloadURLParameters`
- New field `StderrLogFileName`, `StdoutLogFileName` in struct `ScriptExecutionResult`
- New field `BaselineOSs` in struct `TargetOSInfo`
- New field `ValidationResultID` in struct `Test`
- New field `InteropMediaType`, `InteropMediaVersion` in struct `TestResultProperties`
- New field `PackageTags` in struct `TestSummaryProperties`


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/testbase/armtestbase` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).