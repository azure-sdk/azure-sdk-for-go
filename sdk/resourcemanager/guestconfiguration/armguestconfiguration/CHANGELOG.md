# Release History

## 2.0.0 (2025-05-13)
### Breaking Changes

- Type of `AssignmentProperties.VmssVMList` has been changed from `[]*VMSSVMInfo` to `[]*VmssvmInfo`
- Type of `AssignmentReportResource.Properties` has been changed from `any` to `map[string]any`
- Type of `ErrorResponse.Error` has been changed from `*ErrorResponseError` to `*ErrorDetail`
- Function `*AssignmentsClient.NewSubscriptionListPager` has been removed
- Operation `*AssignmentReportsVMSSClient.NewListPager` does not support pagination anymore, use `*AssignmentReportsVMSSClient.List` instead.
- Operation `*AssignmentsClient.NewListPager` does not support pagination anymore, use `*AssignmentsClient.List` instead.
- Operation `*AssignmentsClient.NewRGListPager` does not support pagination anymore, use `*AssignmentsClient.RGList` instead.
- Operation `*AssignmentsVMSSClient.NewListPager` does not support pagination anymore, use `*AssignmentsVMSSClient.List` instead.
- Operation `*HCRPAssignmentsClient.NewListPager` does not support pagination anymore, use `*HCRPAssignmentsClient.List` instead.
- Struct `ErrorResponseError` has been removed
- Struct `OperationList` has been removed
- Struct `OperationProperties` has been removed
- Struct `VMSSVMInfo` has been removed
- Field `Location` of struct `Assignment` has been removed
- Field `Properties` of struct `Operation` has been removed
- Field `OperationList` of struct `OperationsClientListResponse` has been removed
- Field `Location` of struct `ProxyResource` has been removed
- Field `Location` of struct `Resource` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New function `*AssignmentsVMSSClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *AssignmentsVMSSClientCreateOrUpdateOptions) (AssignmentsVMSSClientCreateOrUpdateResponse, error)`
- New function `*ClientFactory.NewConnectedVMwarevSphereAssignmentsClient() *ConnectedVMwarevSphereAssignmentsClient`
- New function `*ClientFactory.NewConnectedVMwarevSphereAssignmentsReportsClient() *ConnectedVMwarevSphereAssignmentsReportsClient`
- New function `NewConnectedVMwarevSphereAssignmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedVMwarevSphereAssignmentsClient, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *ConnectedVMwarevSphereAssignmentsClientCreateOrUpdateOptions) (ConnectedVMwarevSphereAssignmentsClientCreateOrUpdateResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.Delete(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsClientDeleteOptions) (ConnectedVMwarevSphereAssignmentsClientDeleteResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.Get(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsClientGetOptions) (ConnectedVMwarevSphereAssignmentsClientGetResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.List(context.Context, string, string, *ConnectedVMwarevSphereAssignmentsClientListOptions) (ConnectedVMwarevSphereAssignmentsClientListResponse, error)`
- New function `NewConnectedVMwarevSphereAssignmentsReportsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedVMwarevSphereAssignmentsReportsClient, error)`
- New function `*ConnectedVMwarevSphereAssignmentsReportsClient.Get(context.Context, string, string, string, string, *ConnectedVMwarevSphereAssignmentsReportsClientGetOptions) (ConnectedVMwarevSphereAssignmentsReportsClientGetResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsReportsClient.List(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsReportsClientListOptions) (ConnectedVMwarevSphereAssignmentsReportsClientListResponse, error)`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `OperationListResult`
- New struct `VmssvmInfo`
- New field `ContentManagedIdentity` in struct `Navigation`
- New field `ActionType`, `IsDataAction`, `Origin` in struct `Operation`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`


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