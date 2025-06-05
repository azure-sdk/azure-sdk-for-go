# Release History

## 1.0.0 (2025-05-09)
### Breaking Changes

- Type of `StorageTaskUpdateParameters.Properties` has been changed from `*StorageTaskProperties` to `*StorageTaskUpdateProperties`
- Function `*ClientFactory.NewStorageTaskAssignmentClient` has been removed
- Function `*ClientFactory.NewStorageTasksReportClient` has been removed
- Function `NewStorageTaskAssignmentClient` has been removed
- Function `*StorageTaskAssignmentClient.NewListPager` has been removed
- Function `*StorageTasksClient.PreviewActions` has been removed
- Function `NewStorageTasksReportClient` has been removed
- Function `*StorageTasksReportClient.NewListPager` has been removed
- Field `StorageTask` of struct `StorageTasksClientCreateResponse` has been removed

### Features Added

- New function `*ClientFactory.NewStorageTasksOperationGroupClient() *StorageTasksOperationGroupClient`
- New function `*StorageTasksClient.NewListPager(string, string, *StorageTasksClientListOptions) *runtime.Pager[StorageTasksClientListResponse]`
- New function `*StorageTasksClient.NewStorageTaskAssignmentListPager(string, string, *StorageTasksClientStorageTaskAssignmentListOptions) *runtime.Pager[StorageTasksClientStorageTaskAssignmentListResponse]`
- New function `NewStorageTasksOperationGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) (*StorageTasksOperationGroupClient, error)`
- New function `*StorageTasksOperationGroupClient.PreviewActions(context.Context, string, StorageTaskPreviewAction, *StorageTasksOperationGroupClientPreviewActionsOptions) (StorageTasksOperationGroupClientPreviewActionsResponse, error)`
- New struct `StorageTaskUpdateProperties`


## 0.2.0 (2025-04-21)
### Breaking Changes

- Type of `StorageTaskAssignmentClientListOptions.Maxpagesize` has been changed from `*string` to `*int32`
- Type of `StorageTaskOperation.OnFailure` has been changed from `*string` to `*OnFailure`
- Type of `StorageTaskOperation.OnSuccess` has been changed from `*string` to `*OnSuccess`
- Type of `StorageTasksReportClientListOptions.Maxpagesize` has been changed from `*string` to `*int32`

### Features Added

- New value `ProvisioningStateAccepted` added to enum type `ProvisioningState`
- New enum type `OnFailure` with values `OnFailureBreak`
- New enum type `OnSuccess` with values `OnSuccessContinue`


## 0.1.0 (2024-03-08)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).