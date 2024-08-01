# Release History

## 0.8.0 (2024-08-01)
### Breaking Changes

- Function `*TenantConfigurationsClient.Create` parameter(s) have been changed from `(context.Context, ConfigurationName, Configuration, *TenantConfigurationsClientCreateOptions)` to `(context.Context, string, Configuration, *TenantConfigurationsClientCreateOptions)`
- Function `*TenantConfigurationsClient.Delete` parameter(s) have been changed from `(context.Context, ConfigurationName, *TenantConfigurationsClientDeleteOptions)` to `(context.Context, string, *TenantConfigurationsClientDeleteOptions)`
- Function `*TenantConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, ConfigurationName, *TenantConfigurationsClientGetOptions)` to `(context.Context, string, *TenantConfigurationsClientGetOptions)`
- Type of `Dashboard.Properties` has been changed from `*DashboardProperties` to `*DashboardPropertiesWithProvisioningState`
- Type of `DashboardPartMetadata.Type` has been changed from `*string` to `*DashboardPartMetadataType`
- Type of `ErrorResponse.Error` has been changed from `*ErrorDefinition` to `*ErrorDetail`
- Type of `MarkdownPartMetadata.Type` has been changed from `*string` to `*DashboardPartMetadataType`
- Enum `ConfigurationName` has been removed
- Struct `ConfigurationList` has been removed
- Struct `ErrorDefinition` has been removed
- Struct `ResourceProviderOperation` has been removed
- Struct `ResourceProviderOperationDisplay` has been removed
- Struct `ResourceProviderOperationList` has been removed
- Field `AdditionalProperties` of struct `DashboardPartMetadata` has been removed
- Field `AdditionalProperties` of struct `MarkdownPartMetadata` has been removed
- Field `ResourceProviderOperationList` of struct `OperationsClientListResponse` has been removed
- Field `ConfigurationList` of struct `TenantConfigurationsClientListResponse` has been removed

### Features Added

- Type of `DashboardLens.Metadata` has been changed from `map[string]any` to `any`
- Type of `DashboardPartsPosition.Metadata` has been changed from `map[string]any` to `any`
- Type of `DashboardProperties.Metadata` has been changed from `map[string]any` to `any`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `DashboardPartMetadataType` with values `DashboardPartMetadataTypeMarkdown`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `ResourceProvisioningState` with values `ResourceProvisioningStateCanceled`, `ResourceProvisioningStateFailed`, `ResourceProvisioningStateSucceeded`
- New struct `ConfigurationListResult`
- New struct `DashboardPropertiesWithProvisioningState`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `Operation`
- New struct `OperationDisplay`
- New struct `OperationListResult`
- New struct `SystemData`
- New struct `TrackedResource`
- New field `SystemData` in struct `Configuration`
- New field `ProvisioningState` in struct `ConfigurationProperties`
- New field `SystemData` in struct `Dashboard`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New anonymous field `ConfigurationListResult` in struct `TenantConfigurationsClientListResponse`


## 0.7.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/portal/armportal` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).