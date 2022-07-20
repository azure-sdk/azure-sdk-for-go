# Release History

## 2.0.0 (2022-07-20)
### Breaking Changes

- Function `*TemplateSpecVersionsClient.Update` parameter(s) have been changed from `(context.Context, string, string, string, *TemplateSpecVersionsClientUpdateOptions)` to `(context.Context, string, string, string, TemplateSpecVersionUpdateModel, *TemplateSpecVersionsClientUpdateOptions)`
- Function `*Client.Update` parameter(s) have been changed from `(context.Context, string, string, *ClientUpdateOptions)` to `(context.Context, string, string, TemplateSpecUpdateModel, *ClientUpdateOptions)`
- Field `TemplateSpec` of struct `ClientUpdateOptions` has been removed
- Field `TemplateSpecVersionUpdateModel` of struct `TemplateSpecVersionsClientUpdateOptions` has been removed


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armtemplatespecs` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).