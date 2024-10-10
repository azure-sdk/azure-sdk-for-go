# Release History

## 3.0.0 (2024-10-10)
### Breaking Changes

- Type of `KustomizationPatchDefinition.PostBuild` has been changed from `*PostBuildDefinition` to `*PostBuildPatchDefinition`
- Enum `AKSIdentityType` has been removed
- Enum `ComplianceStateType` has been removed
- Enum `LevelType` has been removed
- Enum `MessageLevelType` has been removed
- Enum `OperatorScopeType` has been removed
- Enum `OperatorType` has been removed
- Enum `ProvisioningStateType` has been removed
- Function `*ClientFactory.NewExtensionsClient` has been removed
- Function `*ClientFactory.NewOperationStatusClient` has been removed
- Function `*ClientFactory.NewOperationsClient` has been removed
- Function `*ClientFactory.NewSourceControlConfigurationsClient` has been removed
- Function `NewExtensionsClient` has been removed
- Function `*ExtensionsClient.BeginCreate` has been removed
- Function `*ExtensionsClient.BeginDelete` has been removed
- Function `*ExtensionsClient.Get` has been removed
- Function `*ExtensionsClient.NewListPager` has been removed
- Function `*ExtensionsClient.BeginUpdate` has been removed
- Function `NewOperationStatusClient` has been removed
- Function `*OperationStatusClient.Get` has been removed
- Function `*OperationStatusClient.NewListPager` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewSourceControlConfigurationsClient` has been removed
- Function `*SourceControlConfigurationsClient.CreateOrUpdate` has been removed
- Function `*SourceControlConfigurationsClient.BeginDelete` has been removed
- Function `*SourceControlConfigurationsClient.Get` has been removed
- Function `*SourceControlConfigurationsClient.NewListPager` has been removed
- Struct `ComplianceStatus` has been removed
- Struct `Extension` has been removed
- Struct `ExtensionProperties` has been removed
- Struct `ExtensionPropertiesAksAssignedIdentity` has been removed
- Struct `ExtensionStatus` has been removed
- Struct `ExtensionsList` has been removed
- Struct `HelmOperatorProperties` has been removed
- Struct `Identity` has been removed
- Struct `OperationStatusList` has been removed
- Struct `PatchExtension` has been removed
- Struct `PatchExtensionProperties` has been removed
- Struct `Plan` has been removed
- Struct `ResourceProviderOperation` has been removed
- Struct `ResourceProviderOperationDisplay` has been removed
- Struct `ResourceProviderOperationList` has been removed
- Struct `Scope` has been removed
- Struct `ScopeCluster` has been removed
- Struct `ScopeNamespace` has been removed
- Struct `SourceControlConfiguration` has been removed
- Struct `SourceControlConfigurationList` has been removed
- Struct `SourceControlConfigurationProperties` has been removed

### Features Added

- New value `SourceKindTypeOCIRepository` added to enum type `SourceKindType`
- New enum type `OperationType` with values `OperationTypeCopy`, `OperationTypeExtract`
- New struct `LayerSelectorDefinition`
- New struct `LayerSelectorPatchDefinition`
- New struct `MatchOidcIdentityDefinition`
- New struct `MatchOidcIdentityPatchDefinition`
- New struct `OCIRepositoryDefinition`
- New struct `OCIRepositoryPatchDefinition`
- New struct `OCIRepositoryRefDefinition`
- New struct `OCIRepositoryRefPatchDefinition`
- New struct `PostBuildPatchDefinition`
- New struct `SubstituteFromPatchDefinition`
- New struct `TLSConfigDefinition`
- New struct `TLSConfigPatchDefinition`
- New struct `VerifyDefinition`
- New struct `VerifyPatchDefinition`
- New field `OciRepository` in struct `FluxConfigurationPatchProperties`
- New field `OciRepository` in struct `FluxConfigurationProperties`
- New field `Provider` in struct `GitRepositoryDefinition`
- New field `Provider` in struct `GitRepositoryPatchDefinition`


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