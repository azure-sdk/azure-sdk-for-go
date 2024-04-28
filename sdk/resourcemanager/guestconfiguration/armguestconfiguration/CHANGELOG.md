# Release History

## 1.3.0 (2024-04-28)
### Features Added

- New function `*AssignmentsVMSSClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *AssignmentsVMSSClientCreateOrUpdateOptions) (AssignmentsVMSSClientCreateOrUpdateResponse, error)`
- New function `*ClientFactory.NewConnectedVMwarevSphereAssignmentsClient() *ConnectedVMwarevSphereAssignmentsClient`
- New function `*ClientFactory.NewConnectedVMwarevSphereAssignmentsReportsClient() *ConnectedVMwarevSphereAssignmentsReportsClient`
- New function `NewConnectedVMwarevSphereAssignmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedVMwarevSphereAssignmentsClient, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *ConnectedVMwarevSphereAssignmentsClientCreateOrUpdateOptions) (ConnectedVMwarevSphereAssignmentsClientCreateOrUpdateResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.Delete(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsClientDeleteOptions) (ConnectedVMwarevSphereAssignmentsClientDeleteResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.Get(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsClientGetOptions) (ConnectedVMwarevSphereAssignmentsClientGetResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.NewListPager(string, string, *ConnectedVMwarevSphereAssignmentsClientListOptions) *runtime.Pager[ConnectedVMwarevSphereAssignmentsClientListResponse]`
- New function `NewConnectedVMwarevSphereAssignmentsReportsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedVMwarevSphereAssignmentsReportsClient, error)`
- New function `*ConnectedVMwarevSphereAssignmentsReportsClient.Get(context.Context, string, string, string, string, *ConnectedVMwarevSphereAssignmentsReportsClientGetOptions) (ConnectedVMwarevSphereAssignmentsReportsClientGetResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsReportsClient.List(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsReportsClientListOptions) (ConnectedVMwarevSphereAssignmentsReportsClientListResponse, error)`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).