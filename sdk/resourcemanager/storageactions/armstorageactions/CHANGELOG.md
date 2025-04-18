# Release History

## 1.0.0 (2025-04-18)
### Breaking Changes

- Type of `StorageTaskAssignmentClientListOptions.Maxpagesize` has been changed from `*string` to `*int32`
- Type of `StorageTaskOperation.OnFailure` has been changed from `*string` to `*OnFailure`
- Type of `StorageTaskOperation.OnSuccess` has been changed from `*string` to `*OnSuccess`
- Type of `StorageTaskUpdateParameters.Properties` has been changed from `*StorageTaskProperties` to `*StorageTaskUpdateProperties`
- Type of `StorageTasksReportClientListOptions.Maxpagesize` has been changed from `*string` to `*int32`

### Features Added

- New value `ProvisioningStateAccepted` added to enum type `ProvisioningState`
- New enum type `OnFailure` with values `OnFailureBreak`
- New enum type `OnSuccess` with values `OnSuccessContinue`
- New struct `StorageTaskUpdateProperties`


## 0.1.0 (2024-03-08)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storageactions/armstorageactions` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).