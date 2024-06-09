# Release History

## 2.0.0 (2024-06-09)
### Breaking Changes

- Type of `DeploymentProperties.Parameters` has been changed from `any` to `map[string]*DeploymentParameter`
- Type of `DeploymentWhatIfProperties.Parameters` has been changed from `any` to `map[string]*DeploymentParameter`
- Operation `*TagsClient.CreateOrUpdateAtScope` has been changed to LRO, use `*TagsClient.BeginCreateOrUpdateAtScope` instead.
- Operation `*TagsClient.DeleteAtScope` has been changed to LRO, use `*TagsClient.BeginDeleteAtScope` instead.
- Operation `*TagsClient.UpdateAtScope` has been changed to LRO, use `*TagsClient.BeginUpdateAtScope` instead.

### Features Added

- New enum type `ExportTemplateOutputFormat` with values `ExportTemplateOutputFormatBicep`, `ExportTemplateOutputFormatJSON`
- New struct `DeploymentParameter`
- New struct `KeyVaultParameterReference`
- New struct `KeyVaultReference`
- New field `OutputFormat` in struct `ExportTemplateRequest`
- New field `Output` in struct `ResourceGroupExportResult`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.2.0-beta.3 (2023-10-09)

### Other Changes

- Updated to latest `azcore` beta.

## 1.2.0-beta.2 (2023-07-19)

### Bug Fixes

- Fixed a potential panic in faked paged and long-running operations.

## 1.2.0-beta.1 (2023-06-12)

### Features Added

- Support for test fakes and OpenTelemetry trace spans.

## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).