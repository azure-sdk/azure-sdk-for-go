# Release History

## 2.0.0 (2022-07-08)
### Breaking Changes

- Function `*ApplicationsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *ApplicationsClientUpdateOptions)` to `(context.Context, string, string, ApplicationPatchable, *ApplicationsClientUpdateOptions)`
- Function `*ApplicationsClient.UpdateByID` parameter(s) have been changed from `(context.Context, string, *ApplicationsClientUpdateByIDOptions)` to `(context.Context, string, Application, *ApplicationsClientUpdateByIDOptions)`
- Field `Parameters` of struct `ApplicationsClientUpdateOptions` has been removed
- Field `Parameters` of struct `ApplicationsClientUpdateByIDOptions` has been removed


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armmanagedapplications` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).