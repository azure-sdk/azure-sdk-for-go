# Release History

## 2.0.0 (2025-05-08)
### Breaking Changes

- Type of `Error.Error` has been changed from `*ErrorError` to `*ErrorDetail`
- Type of `SystemData.CreatedByType` has been changed from `*IdentityType` to `*CreatedByType`
- Type of `SystemData.LastModifiedByType` has been changed from `*IdentityType` to `*CreatedByType`
- Enum `IdentityType` has been removed
- Operation `*BotsClient.Update` has been changed to LRO, use `*BotsClient.BeginUpdate` instead.
- Struct `ErrorError` has been removed
- Struct `ValidationResult` has been removed

### Features Added

- New value `SKUNameC1`, `SKUNamePES` added to enum type `SKUName`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New function `*BotsClient.ListSecrets(context.Context, string, string, *BotsClientListSecretsOptions) (BotsClientListSecretsResponse, error)`
- New function `*BotsClient.RegenerateAPIJwtSecret(context.Context, string, string, *BotsClientRegenerateAPIJwtSecretOptions) (BotsClientRegenerateAPIJwtSecretResponse, error)`
- New struct `ErrorDetail`
- New struct `Key`
- New struct `KeyVaultProperties`
- New struct `KeysResponse`
- New field `KeyVaultProperties` in struct `Properties`
- New field `Properties` in struct `UpdateParameters`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/healthbot/armhealthbot` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).