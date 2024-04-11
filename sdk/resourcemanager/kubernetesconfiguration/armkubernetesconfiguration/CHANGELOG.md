# Release History

## 3.0.0-beta.1 (2024-04-11)
### Breaking Changes

- Function `*SourceControlConfigurationsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *SourceControlConfigurationsClientBeginDeleteOptions)` to `(context.Context, string, string, *SourceControlConfigurationsClientBeginDeleteOptions)`
- Function `*SourceControlConfigurationsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, SourceControlConfiguration, *SourceControlConfigurationsClientCreateOrUpdateOptions)` to `(context.Context, string, string, SourceControlConfiguration, *SourceControlConfigurationsClientCreateOrUpdateOptions)`
- Function `*SourceControlConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *SourceControlConfigurationsClientGetOptions)` to `(context.Context, string, string, *SourceControlConfigurationsClientGetOptions)`
- Type of `Extension.Identity` has been changed from `*Identity` to `*ManagedServiceIdentity`
- `OperatorScopeTypeNamespace` from enum `OperatorScopeType` has been removed
- `ScopeTypeNamespace` from enum `ScopeType` has been removed
- Function `*ClientFactory.NewFluxConfigOperationStatusClient` has been removed
- Function `*ExtensionsClient.NewListPager` has been removed
- Function `NewFluxConfigOperationStatusClient` has been removed
- Function `*FluxConfigOperationStatusClient.Get` has been removed
- Function `*FluxConfigurationsClient.BeginCreateOrUpdate` has been removed
- Function `*FluxConfigurationsClient.NewListPager` has been removed
- Function `*OperationStatusClient.Get` has been removed
- Function `*OperationStatusClient.NewListPager` has been removed
- Function `*SourceControlConfigurationsClient.NewListPager` has been removed
- Operation `*ExtensionsClient.BeginUpdate` has been changed to non-LRO, use `*ExtensionsClient.Update` instead.
- Operation `*FluxConfigurationsClient.BeginUpdate` has been changed to non-LRO, use `*FluxConfigurationsClient.Update` instead.
- Struct `ExtensionsList` has been removed
- Struct `FluxConfigurationsList` has been removed
- Struct `OperationStatusList` has been removed
- Struct `OperationStatusResult` has been removed
- Struct `ResourceProviderOperation` has been removed
- Struct `ResourceProviderOperationDisplay` has been removed
- Struct `ResourceProviderOperationList` has been removed
- Struct `SourceControlConfigurationList` has been removed
- Field `ReconciliationWaitDuration` of struct `FluxConfigurationProperties` has been removed
- Field `ResourceProviderOperationList` of struct `OperationsClientListResponse` has been removed

### Features Added

- New value `OperatorScopeTypeNameSpace` added to enum type `OperatorScopeType`
- New value `ProvisioningStateTypeCanceled` added to enum type `ProvisioningStateType`
- New value `ScopeTypeNameSpace` added to enum type `ScopeType`
- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New enum type `ResourceProvisioningState` with values `ResourceProvisioningStateCanceled`, `ResourceProvisioningStateFailed`, `ResourceProvisioningStateSucceeded`
- New enum type `Versions` with values `VersionsV20240601Preview`
- New function `*ExtensionsClient.NewListByResourceGroupPager(string, string, string, string, *ExtensionsClientListByResourceGroupOptions) *runtime.Pager[ExtensionsClientListByResourceGroupResponse]`
- New function `*ExtensionsClient.OperationStatus(context.Context, string, string, string, string, string, string, any, *ExtensionsClientOperationStatusOptions) (ExtensionsClientOperationStatusResponse, error)`
- New function `*FluxConfigurationsClient.BeginCreate(context.Context, string, string, string, string, string, FluxConfiguration, *FluxConfigurationsClientBeginCreateOptions) (*runtime.Poller[FluxConfigurationsClientCreateResponse], error)`
- New function `*FluxConfigurationsClient.NewListByResourceGroupPager(string, string, string, string, *FluxConfigurationsClientListByResourceGroupOptions) *runtime.Pager[FluxConfigurationsClientListByResourceGroupResponse]`
- New function `*FluxConfigurationsClient.OperationStatus(context.Context, string, string, string, string, string, string, any, *FluxConfigurationsClientOperationStatusOptions) (FluxConfigurationsClientOperationStatusResponse, error)`
- New function `*OperationStatusClient.NewListByResourceGroupPager(string, string, string, string, *OperationStatusClientListByResourceGroupOptions) *runtime.Pager[OperationStatusClientListByResourceGroupResponse]`
- New function `*SourceControlConfigurationsClient.NewListByResourceGroupPager(string, *SourceControlConfigurationsClientListByResourceGroupOptions) *runtime.Pager[SourceControlConfigurationsClientListByResourceGroupResponse]`
- New struct `ExtensionListResult`
- New struct `ExtensionUpdate`
- New struct `ExtensionUpdateProperties`
- New struct `FluxConfigurationListResult`
- New struct `FluxConfigurationUpdate`
- New struct `FluxConfigurationUpdateProperties`
- New struct `ManagedServiceIdentity`
- New struct `Operation`
- New struct `OperationDisplay`
- New struct `OperationListResult`
- New struct `OperationModel`
- New struct `OperationModelListResult`
- New struct `OperationsParameter`
- New struct `Paths1B0Hq6PSubscriptionsSubscriptionidResourcegroupsResourcegroupnameProviderClusterrpClusterresourcenameClusternameProvidersMicrosoftKubernetesconfigurationExtensionsExtensionnameOperationsOperationidGetResponses200ContentApplicationJSONSchema`
- New struct `PathsT3WamfSubscriptionsSubscriptionidResourcegroupsResourcegroupnameProviderClusterrpClusterresourcenameClusternameProvidersMicrosoftKubernetesconfigurationFluxconfigurationsFluxconfigurationnameOperationsOperationidGetResponses200ContentApplicationJSONSchema`
- New struct `SourceControlConfigurationListResult`
- New struct `UserAssignedIdentity`
- New field `AutoUpgradeChannel` in struct `ExtensionProperties`
- New field `ReconciliationWait` in struct `FluxConfigurationProperties`
- New anonymous field `OperationListResult` in struct `OperationsClientListResponse`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`


## 2.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.1.0 (2023-09-22)
### Features Added

- New struct `PostBuildDefinition`
- New struct `SubstituteFromDefinition`
- New field `ReconciliationWaitDuration`, `WaitForReconciliation` in struct `FluxConfigurationProperties`
- New field `PostBuild`, `Wait` in struct `KustomizationDefinition`
- New field `PostBuild`, `Wait` in struct `KustomizationPatchDefinition`


## 2.0.0 (2023-05-26)
### Breaking Changes

- Field `InstalledVersion` of struct `ExtensionProperties` has been removed

### Features Added

- New value `SourceKindTypeAzureBlob` added to enum type `SourceKindType`
- New struct `AzureBlobDefinition`
- New struct `AzureBlobPatchDefinition`
- New struct `ManagedIdentityDefinition`
- New struct `ManagedIdentityPatchDefinition`
- New struct `Plan`
- New struct `ServicePrincipalDefinition`
- New struct `ServicePrincipalPatchDefinition`
- New field `Plan` in struct `Extension`
- New field `CurrentVersion`, `IsSystemExtension` in struct `ExtensionProperties`
- New field `AzureBlob` in struct `FluxConfigurationPatchProperties`
- New field `AzureBlob` in struct `FluxConfigurationProperties`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).