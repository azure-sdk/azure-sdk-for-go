# Release History

## 0.4.0 (2023-11-30)
### Breaking Changes

- Function `*ClientFactory.NewReportClient` has been removed
- Function `*ClientFactory.NewReportsClient` has been removed
- Function `*ClientFactory.NewSnapshotClient` has been removed
- Function `*ClientFactory.NewSnapshotsClient` has been removed
- Function `NewReportClient` has been removed
- Function `*ReportClient.BeginCreateOrUpdate` has been removed
- Function `*ReportClient.BeginDelete` has been removed
- Function `*ReportClient.Get` has been removed
- Function `*ReportClient.BeginUpdate` has been removed
- Function `NewReportsClient` has been removed
- Function `*ReportsClient.NewListPager` has been removed
- Function `NewSnapshotClient` has been removed
- Function `*SnapshotClient.BeginDownload` has been removed
- Function `*SnapshotClient.Get` has been removed
- Function `NewSnapshotsClient` has been removed
- Function `*SnapshotsClient.NewListPager` has been removed
- Struct `ReportResourceList` has been removed
- Struct `ReportResourcePatch` has been removed
- Struct `SnapshotResourceList` has been removed
- Field `ReportSystemData` of struct `SnapshotProperties` has been removed

### Features Added

- New function `*ClientFactory.NewReportResourcesClient() *ReportResourcesClient`
- New function `*ClientFactory.NewSnapshotResourcesClient() *SnapshotResourcesClient`
- New function `NewReportResourcesClient(azcore.TokenCredential, *arm.ClientOptions) (*ReportResourcesClient, error)`
- New function `*ReportResourcesClient.BeginCreateOrUpdate(context.Context, string, ReportResource, *ReportResourcesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ReportResourcesClientCreateOrUpdateResponse], error)`
- New function `*ReportResourcesClient.BeginDelete(context.Context, string, *ReportResourcesClientBeginDeleteOptions) (*runtime.Poller[ReportResourcesClientDeleteResponse], error)`
- New function `*ReportResourcesClient.Get(context.Context, string, *ReportResourcesClientGetOptions) (ReportResourcesClientGetResponse, error)`
- New function `*ReportResourcesClient.NewListByTenantPager(*ReportResourcesClientListByTenantOptions) *runtime.Pager[ReportResourcesClientListByTenantResponse]`
- New function `*ReportResourcesClient.BeginUpdate(context.Context, string, ReportResourceUpdate, *ReportResourcesClientBeginUpdateOptions) (*runtime.Poller[ReportResourcesClientUpdateResponse], error)`
- New function `NewSnapshotResourcesClient(azcore.TokenCredential, *arm.ClientOptions) (*SnapshotResourcesClient, error)`
- New function `*SnapshotResourcesClient.BeginDownload(context.Context, string, string, SnapshotDownloadRequest, *SnapshotResourcesClientBeginDownloadOptions) (*runtime.Poller[SnapshotResourcesClientDownloadResponse], error)`
- New function `*SnapshotResourcesClient.Get(context.Context, string, string, *SnapshotResourcesClientGetOptions) (SnapshotResourcesClientGetResponse, error)`
- New function `*SnapshotResourcesClient.NewListByReportResourcePager(string, *SnapshotResourcesClientListByReportResourceOptions) *runtime.Pager[SnapshotResourcesClientListByReportResourceResponse]`
- New struct `ReportResourceListResult`
- New struct `ReportResourceUpdate`
- New struct `ReportResourceUpdateProperties`
- New struct `SnapshotResourceListResult`


## 0.3.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.2.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.1.0 (2022-10-26)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appcomplianceautomation/armappcomplianceautomation` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).