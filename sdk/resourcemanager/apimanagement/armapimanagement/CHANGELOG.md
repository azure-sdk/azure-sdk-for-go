# Release History

## 2.0.0 (2023-01-22)
### Breaking Changes

- Function `*ContentItemClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, *ContentItemClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, string, ContentItemContract, *ContentItemClientCreateOrUpdateOptions)`
- Function `*ContentTypeClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *ContentTypeClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, ContentTypeContract, *ContentTypeClientCreateOrUpdateOptions)`
- Type of `ContentItemContract.Properties` has been changed from `map[string]interface{}` to `map[string]any`
- Type of `ContentTypeContractProperties.Schema` has been changed from `interface{}` to `any`
- Type of `GlobalSchemaContractProperties.Document` has been changed from `interface{}` to `any`
- Type of `GlobalSchemaContractProperties.Value` has been changed from `interface{}` to `any`
- Type of `Operation.Properties` has been changed from `interface{}` to `any`
- Type of `ParameterExampleContract.Value` has been changed from `interface{}` to `any`
- Type of `SchemaDocumentProperties.Components` has been changed from `interface{}` to `any`
- Type of `SchemaDocumentProperties.Definitions` has been changed from `interface{}` to `any`

### Features Added

- New field `Metrics` in struct `DiagnosticContractProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).