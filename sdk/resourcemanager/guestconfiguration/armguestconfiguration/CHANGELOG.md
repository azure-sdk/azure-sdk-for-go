# Release History

## 1.1.0 (2023-02-23)
### Features Added

- New function `NewVMWareSpehereAssignmentsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VMWareSpehereAssignmentsClient, error)`
- New function `*VMWareSpehereAssignmentsClient.CreateOrUpdate(context.Context, string, string, string, Assignment, *VMWareSpehereAssignmentsClientCreateOrUpdateOptions) (VMWareSpehereAssignmentsClientCreateOrUpdateResponse, error)`
- New function `*VMWareSpehereAssignmentsClient.Delete(context.Context, string, string, string, *VMWareSpehereAssignmentsClientDeleteOptions) (VMWareSpehereAssignmentsClientDeleteResponse, error)`
- New function `*VMWareSpehereAssignmentsClient.Get(context.Context, string, string, string, *VMWareSpehereAssignmentsClientGetOptions) (VMWareSpehereAssignmentsClientGetResponse, error)`
- New function `*VMWareSpehereAssignmentsClient.NewListPager(string, string, *VMWareSpehereAssignmentsClientListOptions) *runtime.Pager[VMWareSpehereAssignmentsClientListResponse]`
- New function `NewVMWareSpehereAssignmentsReportsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*VMWareSpehereAssignmentsReportsClient, error)`
- New function `*VMWareSpehereAssignmentsReportsClient.Get(context.Context, string, string, string, string, *VMWareSpehereAssignmentsReportsClientGetOptions) (VMWareSpehereAssignmentsReportsClientGetResponse, error)`
- New function `*VMWareSpehereAssignmentsReportsClient.List(context.Context, string, string, string, *VMWareSpehereAssignmentsReportsClientListOptions) (VMWareSpehereAssignmentsReportsClientListResponse, error)`
- New struct `VMWareSpehereAssignmentsClient`
- New struct `VMWareSpehereAssignmentsReportsClient`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/guestconfiguration/armguestconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).