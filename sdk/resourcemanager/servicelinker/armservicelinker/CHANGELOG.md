# Release History

## 2.0.0-beta.2 (2024-03-20)
### Breaking Changes

- `ActionTypeEnable`, `ActionTypeOptOut` from enum `ActionType` has been removed
- `AuthTypeAccessKey`, `AuthTypeEasyAuthMicrosoftEntraID`, `AuthTypeUserAccount` from enum `AuthType` has been removed
- `ClientTypeDapr`, `ClientTypeJmsSpringBoot`, `ClientTypeKafkaSpringBoot` from enum `ClientType` has been removed
- Enum `AccessKeyPermissions` has been removed
- Enum `AllowType` has been removed
- Enum `AuthMode` has been removed
- Enum `AzureResourceType` has been removed
- Enum `DaprBindingComponentDirection` has been removed
- Enum `DaprMetadataRequired` has been removed
- Enum `DeleteOrUpdateBehavior` has been removed
- Enum `DryrunActionName` has been removed
- Enum `DryrunPrerequisiteResultType` has been removed
- Enum `DryrunPreviewOperationType` has been removed
- Enum `LinkerConfigurationType` has been removed
- Enum `SecretSourceType` has been removed
- Enum `SecretType` has been removed
- Enum `TargetServiceType` has been removed
- Enum `ValidationResultStatus` has been removed
- Function `*AccessKeyInfoBase.GetAuthInfoBase` has been removed
- Function `*AzureKeyVaultProperties.GetAzureResourcePropertiesBase` has been removed
- Function `*AzureResource.GetTargetServiceBase` has been removed
- Function `*AzureResourcePropertiesBase.GetAzureResourcePropertiesBase` has been removed
- Function `*BasicErrorDryrunPrerequisiteResult.GetDryrunPrerequisiteResult` has been removed
- Function `*ClientFactory.NewConfigurationNamesClient` has been removed
- Function `*ClientFactory.NewConnectorClient` has been removed
- Function `*ClientFactory.NewLinkersClient` has been removed
- Function `NewConfigurationNamesClient` has been removed
- Function `*ConfigurationNamesClient.NewListPager` has been removed
- Function `*ConfluentBootstrapServer.GetTargetServiceBase` has been removed
- Function `*ConfluentSchemaRegistry.GetTargetServiceBase` has been removed
- Function `NewConnectorClient` has been removed
- Function `*ConnectorClient.BeginCreateDryrun` has been removed
- Function `*ConnectorClient.BeginCreateOrUpdate` has been removed
- Function `*ConnectorClient.BeginDelete` has been removed
- Function `*ConnectorClient.DeleteDryrun` has been removed
- Function `*ConnectorClient.GenerateConfigurations` has been removed
- Function `*ConnectorClient.Get` has been removed
- Function `*ConnectorClient.GetDryrun` has been removed
- Function `*ConnectorClient.NewListDryrunPager` has been removed
- Function `*ConnectorClient.NewListPager` has been removed
- Function `*ConnectorClient.BeginUpdate` has been removed
- Function `*ConnectorClient.BeginUpdateDryrun` has been removed
- Function `*ConnectorClient.BeginValidate` has been removed
- Function `*CreateOrUpdateDryrunParameters.GetDryrunParameters` has been removed
- Function `*DryrunParameters.GetDryrunParameters` has been removed
- Function `*DryrunPrerequisiteResult.GetDryrunPrerequisiteResult` has been removed
- Function `*EasyAuthMicrosoftEntraIDAuthInfo.GetAuthInfoBase` has been removed
- Function `*KeyVaultSecretReferenceSecretInfo.GetSecretInfoBase` has been removed
- Function `*KeyVaultSecretURISecretInfo.GetSecretInfoBase` has been removed
- Function `NewLinkersClient` has been removed
- Function `*LinkersClient.BeginCreateDryrun` has been removed
- Function `*LinkersClient.DeleteDryrun` has been removed
- Function `*LinkersClient.GenerateConfigurations` has been removed
- Function `*LinkersClient.GetDryrun` has been removed
- Function `*LinkersClient.NewListDaprConfigurationsPager` has been removed
- Function `*LinkersClient.NewListDryrunPager` has been removed
- Function `*LinkersClient.BeginUpdateDryrun` has been removed
- Function `*PermissionsMissingDryrunPrerequisiteResult.GetDryrunPrerequisiteResult` has been removed
- Function `*SecretInfoBase.GetSecretInfoBase` has been removed
- Function `*SelfHostedServer.GetTargetServiceBase` has been removed
- Function `*TargetServiceBase.GetTargetServiceBase` has been removed
- Function `*UserAccountAuthInfo.GetAuthInfoBase` has been removed
- Function `*ValueSecretInfo.GetSecretInfoBase` has been removed
- Struct `AccessKeyInfoBase` has been removed
- Struct `AzureKeyVaultProperties` has been removed
- Struct `AzureResource` has been removed
- Struct `BasicErrorDryrunPrerequisiteResult` has been removed
- Struct `ConfigurationInfo` has been removed
- Struct `ConfigurationName` has been removed
- Struct `ConfigurationNameItem` has been removed
- Struct `ConfigurationNameResult` has been removed
- Struct `ConfigurationNames` has been removed
- Struct `ConfigurationResult` has been removed
- Struct `ConfigurationStore` has been removed
- Struct `ConfluentBootstrapServer` has been removed
- Struct `ConfluentSchemaRegistry` has been removed
- Struct `CreateOrUpdateDryrunParameters` has been removed
- Struct `DaprConfigurationList` has been removed
- Struct `DaprConfigurationProperties` has been removed
- Struct `DaprConfigurationResource` has been removed
- Struct `DaprMetadata` has been removed
- Struct `DaprProperties` has been removed
- Struct `DatabaseAADAuthInfo` has been removed
- Struct `DryrunList` has been removed
- Struct `DryrunOperationPreview` has been removed
- Struct `DryrunPatch` has been removed
- Struct `DryrunProperties` has been removed
- Struct `DryrunResource` has been removed
- Struct `EasyAuthMicrosoftEntraIDAuthInfo` has been removed
- Struct `FirewallRules` has been removed
- Struct `KeyVaultSecretReferenceSecretInfo` has been removed
- Struct `KeyVaultSecretURISecretInfo` has been removed
- Struct `PermissionsMissingDryrunPrerequisiteResult` has been removed
- Struct `PublicNetworkSolution` has been removed
- Struct `ResourceList` has been removed
- Struct `SelfHostedServer` has been removed
- Struct `UserAccountAuthInfo` has been removed
- Struct `ValidateOperationResult` has been removed
- Struct `ValidationResultItem` has been removed
- Struct `ValueSecretInfo` has been removed
- Field `ConfigurationResult` of struct `LinkerClientListConfigurationsResponse` has been removed
- Field `ResourceList` of struct `LinkerClientListResponse` has been removed
- Field `ValidateOperationResult` of struct `LinkerClientValidateResponse` has been removed
- Field `ConfigurationInfo`, `PublicNetworkSolution`, `Scope`, `TargetService` of struct `LinkerProperties` has been removed
- Field `SystemData` of struct `ProxyResource` has been removed
- Field `SystemData` of struct `Resource` has been removed
- Field `AuthMode`, `SecretInfo` of struct `SecretAuthInfo` has been removed
- Field `KeyVaultSecretName` of struct `SecretStore` has been removed
- Field `AuthMode`, `DeleteOrUpdateBehavior`, `Roles` of struct `ServicePrincipalCertificateAuthInfo` has been removed
- Field `AuthMode`, `DeleteOrUpdateBehavior`, `Roles`, `UserName` of struct `ServicePrincipalSecretAuthInfo` has been removed
- Field `ConfigType`, `Description`, `KeyVaultReferenceIdentity` of struct `SourceConfiguration` has been removed
- Field `AuthMode`, `DeleteOrUpdateBehavior`, `Roles`, `UserName` of struct `SystemAssignedIdentityAuthInfo` has been removed
- Field `AuthMode`, `DeleteOrUpdateBehavior`, `Roles`, `UserName` of struct `UserAssignedIdentityAuthInfo` has been removed
- Field `DeleteOrUpdateBehavior` of struct `VNetSolution` has been removed
- Field `IsConnectionAvailable`, `LinkerName`, `SourceID`, `ValidationDetail` of struct `ValidateResult` has been removed

### Features Added

- New enum type `LinkerStatus` with values `LinkerStatusHealthy`, `LinkerStatusNotHealthy`
- New struct `LinkerList`
- New struct `SourceConfigurationResult`
- New anonymous field `SourceConfigurationResult` in struct `LinkerClientListConfigurationsResponse`
- New anonymous field `LinkerList` in struct `LinkerClientListResponse`
- New anonymous field `ValidateResult` in struct `LinkerClientValidateResponse`
- New field `TargetID` in struct `LinkerProperties`
- New field `Secret` in struct `SecretAuthInfo`
- New field `LinkerStatus`, `Name`, `Reason` in struct `ValidateResult`


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

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).