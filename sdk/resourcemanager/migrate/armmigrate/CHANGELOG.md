# Release History

## 2.0.0 (2022-08-03)
### Breaking Changes

- Function `*ImportCollectorsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *ImportCollectorsClientCreateOptions)` to `(context.Context, string, string, string, ImportCollector, *ImportCollectorsClientCreateOptions)`
- Function `*VMwareCollectorsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *VMwareCollectorsClientCreateOptions)` to `(context.Context, string, string, string, VMwareCollector, *VMwareCollectorsClientCreateOptions)`
- Function `*AssessmentsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, string, *AssessmentsClientCreateOptions)` to `(context.Context, string, string, string, string, Assessment, *AssessmentsClientCreateOptions)`
- Function `*GroupsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *GroupsClientCreateOptions)` to `(context.Context, string, string, string, Group, *GroupsClientCreateOptions)`
- Function `*ProjectsClient.Update` parameter(s) have been changed from `(context.Context, string, string, *ProjectsClientUpdateOptions)` to `(context.Context, string, string, Project, *ProjectsClientUpdateOptions)`
- Function `*HyperVCollectorsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *HyperVCollectorsClientCreateOptions)` to `(context.Context, string, string, string, HyperVCollector, *HyperVCollectorsClientCreateOptions)`
- Function `*ProjectsClient.Create` parameter(s) have been changed from `(context.Context, string, string, *ProjectsClientCreateOptions)` to `(context.Context, string, string, Project, *ProjectsClientCreateOptions)`
- Function `*PrivateEndpointConnectionClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *PrivateEndpointConnectionClientUpdateOptions)` to `(context.Context, string, string, string, PrivateEndpointConnection, *PrivateEndpointConnectionClientUpdateOptions)`
- Function `*ServerCollectorsClient.Create` parameter(s) have been changed from `(context.Context, string, string, string, *ServerCollectorsClientCreateOptions)` to `(context.Context, string, string, string, ServerCollector, *ServerCollectorsClientCreateOptions)`
- Field `Project` of struct `ProjectsClientCreateOptions` has been removed
- Field `CollectorBody` of struct `VMwareCollectorsClientCreateOptions` has been removed
- Field `Project` of struct `ProjectsClientUpdateOptions` has been removed
- Field `CollectorBody` of struct `ImportCollectorsClientCreateOptions` has been removed
- Field `Assessment` of struct `AssessmentsClientCreateOptions` has been removed
- Field `Group` of struct `GroupsClientCreateOptions` has been removed
- Field `PrivateEndpointConnectionBody` of struct `PrivateEndpointConnectionClientUpdateOptions` has been removed
- Field `CollectorBody` of struct `ServerCollectorsClientCreateOptions` has been removed
- Field `CollectorBody` of struct `HyperVCollectorsClientCreateOptions` has been removed


## 1.0.0 (2022-06-10)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/migrate/armmigrate` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).