# Release History

## 1.1.0 (2023-03-03)
### Features Added

- New function `*AssignmentsVMSSClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *AssignmentsVMSSClientCreateOrUpdateOptions) (AssignmentsVMSSClientCreateOrUpdateResponse, error)`
- New function `NewConnectedVMwarevSphereAssignmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedVMwarevSphereAssignmentsClient, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *ConnectedVMwarevSphereAssignmentsClientCreateOrUpdateOptions) (ConnectedVMwarevSphereAssignmentsClientCreateOrUpdateResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.Delete(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsClientDeleteOptions) (ConnectedVMwarevSphereAssignmentsClientDeleteResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.Get(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsClientGetOptions) (ConnectedVMwarevSphereAssignmentsClientGetResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsClient.NewListPager(string, string, *ConnectedVMwarevSphereAssignmentsClientListOptions) *runtime.Pager[ConnectedVMwarevSphereAssignmentsClientListResponse]`
- New function `NewConnectedVMwarevSphereAssignmentsReportsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedVMwarevSphereAssignmentsReportsClient, error)`
- New function `*ConnectedVMwarevSphereAssignmentsReportsClient.Get(context.Context, string, string, string, string, *ConnectedVMwarevSphereAssignmentsReportsClientGetOptions) (ConnectedVMwarevSphereAssignmentsReportsClientGetResponse, error)`
- New function `*ConnectedVMwarevSphereAssignmentsReportsClient.List(context.Context, string, string, string, *ConnectedVMwarevSphereAssignmentsReportsClientListOptions) (ConnectedVMwarevSphereAssignmentsReportsClientListResponse, error)`
- New struct `ConnectedVMwarevSphereAssignmentsClient`
- New struct `ConnectedVMwarevSphereAssignmentsReportsClient`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).