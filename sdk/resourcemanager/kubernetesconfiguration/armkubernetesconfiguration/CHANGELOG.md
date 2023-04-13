# Release History

## 2.0.0 (2023-04-13)
### Breaking Changes

- Function `*ExtensionsClient.BeginCreate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, Extension, *ExtensionsClientBeginCreateOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, Extension, *ExtensionsClientBeginCreateOptions)`
- Function `*ExtensionsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *ExtensionsClientBeginDeleteOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, *ExtensionsClientBeginDeleteOptions)`
- Function `*ExtensionsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, PatchExtension, *ExtensionsClientBeginUpdateOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, PatchExtension, *ExtensionsClientBeginUpdateOptions)`
- Function `*ExtensionsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *ExtensionsClientGetOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, *ExtensionsClientGetOptions)`
- Function `*ExtensionsClient.NewListPager` parameter(s) have been changed from `(string, string, string, string, *ExtensionsClientListOptions)` to `(string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, *ExtensionsClientListOptions)`
- Function `*FluxConfigOperationStatusClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, *FluxConfigOperationStatusClientGetOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, string, *FluxConfigOperationStatusClientGetOptions)`
- Function `*FluxConfigurationsClient.BeginCreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, FluxConfiguration, *FluxConfigurationsClientBeginCreateOrUpdateOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, FluxConfiguration, *FluxConfigurationsClientBeginCreateOrUpdateOptions)`
- Function `*FluxConfigurationsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *FluxConfigurationsClientBeginDeleteOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, *FluxConfigurationsClientBeginDeleteOptions)`
- Function `*FluxConfigurationsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, FluxConfigurationPatch, *FluxConfigurationsClientBeginUpdateOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, FluxConfigurationPatch, *FluxConfigurationsClientBeginUpdateOptions)`
- Function `*FluxConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *FluxConfigurationsClientGetOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, *FluxConfigurationsClientGetOptions)`
- Function `*FluxConfigurationsClient.NewListPager` parameter(s) have been changed from `(string, string, string, string, *FluxConfigurationsClientListOptions)` to `(string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, *FluxConfigurationsClientListOptions)`
- Function `*OperationStatusClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, *OperationStatusClientGetOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, string, *OperationStatusClientGetOptions)`
- Function `*OperationStatusClient.NewListPager` parameter(s) have been changed from `(string, string, string, string, *OperationStatusClientListOptions)` to `(string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, *OperationStatusClientListOptions)`
- Function `*SourceControlConfigurationsClient.BeginDelete` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *SourceControlConfigurationsClientBeginDeleteOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, *SourceControlConfigurationsClientBeginDeleteOptions)`
- Function `*SourceControlConfigurationsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, SourceControlConfiguration, *SourceControlConfigurationsClientCreateOrUpdateOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, SourceControlConfiguration, *SourceControlConfigurationsClientCreateOrUpdateOptions)`
- Function `*SourceControlConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *SourceControlConfigurationsClientGetOptions)` to `(context.Context, string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, string, *SourceControlConfigurationsClientGetOptions)`
- Function `*SourceControlConfigurationsClient.NewListPager` parameter(s) have been changed from `(string, string, string, string, *SourceControlConfigurationsClientListOptions)` to `(string, KubernetesClusterResourceProviderName, KubernetesClusterResourceName, string, *SourceControlConfigurationsClientListOptions)`
- Field `InstalledVersion` of struct `ExtensionProperties` has been removed

### Features Added

- New value `SourceKindTypeAzureBlob` added to enum type `SourceKindType`
- New enum type `KubernetesClusterResourceName` with values `KubernetesClusterResourceNameConnectedClusters`, `KubernetesClusterResourceNameManagedClusters`, `KubernetesClusterResourceNameProvisionedClusters`
- New enum type `KubernetesClusterResourceProviderName` with values `KubernetesClusterResourceProviderNameMicrosoftContainerService`, `KubernetesClusterResourceProviderNameMicrosoftHybridContainerService`, `KubernetesClusterResourceProviderNameMicrosoftKubernetes`
- New struct `AzureBlobDefinition`
- New struct `AzureBlobPatchDefinition`
- New struct `ManagedIdentityDefinition`
- New struct `ManagedIdentityPatchDefinition`
- New struct `Plan`
- New struct `ServicePrincipalDefinition`
- New struct `ServicePrincipalPatchDefinition`
- New field `Plan` in struct `Extension`
- New field `CurrentVersion` in struct `ExtensionProperties`
- New field `IsSystemExtension` in struct `ExtensionProperties`
- New field `AzureBlob` in struct `FluxConfigurationPatchProperties`
- New field `AzureBlob` in struct `FluxConfigurationProperties`


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armkubernetesconfiguration` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).