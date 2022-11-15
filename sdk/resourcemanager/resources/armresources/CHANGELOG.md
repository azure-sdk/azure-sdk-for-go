# Release History

## 2.0.0 (2022-11-15)
### Breaking Changes

- Function `*TagsClient.DeleteAtScope` has been removed
- Function `*TagsClient.UpdateAtScope` has been removed
- Function `*TagsClient.CreateOrUpdateAtScope` has been removed
- Struct `CloudError` has been removed
- Struct `TagsClientCreateOrUpdateAtScopeOptions` has been removed
- Struct `TagsClientDeleteAtScopeOptions` has been removed
- Struct `TagsClientUpdateAtScopeOptions` has been removed

### Features Added

- New function `*TagsClient.BeginDeleteAtScope(context.Context, string, *TagsClientBeginDeleteAtScopeOptions) (*runtime.Poller[TagsClientDeleteAtScopeResponse], error)`
- New function `*TagsClient.BeginUpdateAtScope(context.Context, string, TagsPatchResource, *TagsClientBeginUpdateAtScopeOptions) (*runtime.Poller[TagsClientUpdateAtScopeResponse], error)`
- New function `*TagsClient.BeginCreateOrUpdateAtScope(context.Context, string, TagsResource, *TagsClientBeginCreateOrUpdateAtScopeOptions) (*runtime.Poller[TagsClientCreateOrUpdateAtScopeResponse], error)`
- New struct `TagsClientBeginCreateOrUpdateAtScopeOptions`
- New struct `TagsClientBeginDeleteAtScopeOptions`
- New struct `TagsClientBeginUpdateAtScopeOptions`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).