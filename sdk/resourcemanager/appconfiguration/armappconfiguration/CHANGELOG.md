# Release History

## 2.0.0 (2022-08-04)
### Breaking Changes

- Function `*KeyValuesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *KeyValuesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, KeyValue, *KeyValuesClientCreateOrUpdateOptions)`
- Function `*KeyValuesClient.NewListByConfigurationStorePager` has been removed
- Struct `KeyValuesClientListByConfigurationStoreOptions` has been removed
- Struct `KeyValuesClientListByConfigurationStoreResponse` has been removed
- Field `KeyValueParameters` of struct `KeyValuesClientCreateOrUpdateOptions` has been removed


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).