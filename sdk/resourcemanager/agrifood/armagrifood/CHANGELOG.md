# Release History

## 0.8.0 (2022-10-03)
### Breaking Changes

- Function `*ExtensionsClient.Create` has been removed
- Function `*ExtensionsClient.Update` has been removed
- Struct `ExtensionsClientCreateOptions` has been removed
- Struct `ExtensionsClientCreateResponse` has been removed
- Struct `ExtensionsClientUpdateOptions` has been removed
- Struct `ExtensionsClientUpdateResponse` has been removed

### Features Added

- New function `*ExtensionsClient.CreateOrUpdate(context.Context, string, string, string, *ExtensionsClientCreateOrUpdateOptions) (ExtensionsClientCreateOrUpdateResponse, error)`
- New struct `APIProperties`
- New struct `ExtensionInstallationRequest`
- New struct `ExtensionsClientCreateOrUpdateOptions`
- New struct `ExtensionsClientCreateOrUpdateResponse`
- New field `AdditionalAPIProperties` in struct `ExtensionProperties`


## 0.7.0 (2022-08-23)
### Breaking Changes

- Operation `FarmBeatsModelsClient.Update` has been changed to LRO, use `FarmBeatsModels.BeginUpdate` instead.

### Features Added

- Sensor, Private endpoint & managedIdentity support added.

## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/agrifood/armagrifood` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
