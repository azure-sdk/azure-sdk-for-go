# Release History

## 2.5.0-beta.2 (2025-01-28)
### Breaking Changes

- Type of `ContainerGroupProfile.Properties` has been changed from `*ContainerGroupProfilePropertiesProperties` to `*ContainerGroupProfileProperties`
- Function `*ClientFactory.NewContainerGroupProfileClient` has been removed
- Function `*ClientFactory.NewContainerGroupProfilesClient` has been removed
- Function `NewContainerGroupProfileClient` has been removed
- Function `*ContainerGroupProfileClient.GetByRevisionNumber` has been removed
- Function `*ContainerGroupProfileClient.NewListAllRevisionsPager` has been removed
- Function `NewContainerGroupProfilesClient` has been removed
- Function `*ContainerGroupProfilesClient.CreateOrUpdate` has been removed
- Function `*ContainerGroupProfilesClient.Delete` has been removed
- Function `*ContainerGroupProfilesClient.Get` has been removed
- Function `*ContainerGroupProfilesClient.NewListByResourceGroupPager` has been removed
- Function `*ContainerGroupProfilesClient.NewListPager` has been removed
- Function `*ContainerGroupProfilesClient.Patch` has been removed
- Struct `ConfigMap` has been removed
- Struct `ContainerGroupProfilePropertiesProperties` has been removed
- Struct `ContainerGroupProfileReferenceDefinition` has been removed
- Struct `StandbyPoolProfileDefinition` has been removed
- Field `Properties` of struct `ContainerGroupProfileProperties` has been removed
- Field `ContainerGroupProfile`, `IsCreatedFromStandbyPool`, `StandbyPoolProfile` of struct `ContainerGroupPropertiesProperties` has been removed
- Field `ConfigMap` of struct `ContainerProperties` has been removed

### Features Added

- New value `ContainerGroupSKUNotSpecified` added to enum type `ContainerGroupSKU`
- New enum type `AzureFileShareAccessTier` with values `AzureFileShareAccessTierCool`, `AzureFileShareAccessTierHot`, `AzureFileShareAccessTierPremium`, `AzureFileShareAccessTierTransactionOptimized`
- New enum type `AzureFileShareAccessType` with values `AzureFileShareAccessTypeExclusive`, `AzureFileShareAccessTypeShared`
- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `IdentityAccessLevel` with values `IdentityAccessLevelAll`, `IdentityAccessLevelSystem`, `IdentityAccessLevelUser`
- New enum type `NGroupProvisioningState` with values `NGroupProvisioningStateCanceled`, `NGroupProvisioningStateCreating`, `NGroupProvisioningStateDeleting`, `NGroupProvisioningStateFailed`, `NGroupProvisioningStateMigrating`, `NGroupProvisioningStateSucceeded`, `NGroupProvisioningStateUpdating`
- New enum type `NGroupUpdateMode` with values `NGroupUpdateModeManual`, `NGroupUpdateModeRolling`
- New function `NewCGProfileClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CGProfileClient, error)`
- New function `*CGProfileClient.CreateOrUpdate(context.Context, string, string, ContainerGroupProfile, *CGProfileClientCreateOrUpdateOptions) (CGProfileClientCreateOrUpdateResponse, error)`
- New function `*CGProfileClient.BeginDelete(context.Context, string, string, *CGProfileClientBeginDeleteOptions) (*runtime.Poller[CGProfileClientDeleteResponse], error)`
- New function `*CGProfileClient.Get(context.Context, string, string, *CGProfileClientGetOptions) (CGProfileClientGetResponse, error)`
- New function `*CGProfileClient.GetByRevisionNumber(context.Context, string, string, string, *CGProfileClientGetByRevisionNumberOptions) (CGProfileClientGetByRevisionNumberResponse, error)`
- New function `*CGProfileClient.NewListAllRevisionsPager(string, string, *CGProfileClientListAllRevisionsOptions) *runtime.Pager[CGProfileClientListAllRevisionsResponse]`
- New function `*CGProfileClient.Update(context.Context, string, string, ContainerGroupProfilePatch, *CGProfileClientUpdateOptions) (CGProfileClientUpdateResponse, error)`
- New function `NewCGProfilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CGProfilesClient, error)`
- New function `*CGProfilesClient.NewListByResourceGroupPager(string, *CGProfilesClientListByResourceGroupOptions) *runtime.Pager[CGProfilesClientListByResourceGroupResponse]`
- New function `*CGProfilesClient.NewListBySubscriptionPager(*CGProfilesClientListBySubscriptionOptions) *runtime.Pager[CGProfilesClientListBySubscriptionResponse]`
- New function `*ClientFactory.NewCGProfileClient() *CGProfileClient`
- New function `*ClientFactory.NewCGProfilesClient() *CGProfilesClient`
- New function `*ClientFactory.NewNGroupsClient() *NGroupsClient`
- New function `NewNGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*NGroupsClient, error)`
- New function `*NGroupsClient.BeginCreateOrUpdate(context.Context, string, string, NGroup, *NGroupsClientBeginCreateOrUpdateOptions) (*runtime.Poller[NGroupsClientCreateOrUpdateResponse], error)`
- New function `*NGroupsClient.BeginDelete(context.Context, string, string, *NGroupsClientBeginDeleteOptions) (*runtime.Poller[NGroupsClientDeleteResponse], error)`
- New function `*NGroupsClient.Get(context.Context, string, string, *NGroupsClientGetOptions) (NGroupsClientGetResponse, error)`
- New function `*NGroupsClient.NewListByResourceGroupPager(string, *NGroupsClientListByResourceGroupOptions) *runtime.Pager[NGroupsClientListByResourceGroupResponse]`
- New function `*NGroupsClient.NewListPager(*NGroupsClientListOptions) *runtime.Pager[NGroupsClientListResponse]`
- New function `*NGroupsClient.BeginRestart(context.Context, string, string, *NGroupsClientBeginRestartOptions) (*runtime.Poller[NGroupsClientRestartResponse], error)`
- New function `*NGroupsClient.BeginStart(context.Context, string, string, *NGroupsClientBeginStartOptions) (*runtime.Poller[NGroupsClientStartResponse], error)`
- New function `*NGroupsClient.Stop(context.Context, string, string, *NGroupsClientStopOptions) (NGroupsClientStopResponse, error)`
- New function `*NGroupsClient.BeginUpdate(context.Context, string, string, NGroup, *NGroupsClientBeginUpdateOptions) (*runtime.Poller[NGroupsClientUpdateResponse], error)`
- New struct `APIEntityReference`
- New struct `ApplicationGateway`
- New struct `ApplicationGatewayBackendAddressPool`
- New struct `ContainerGroupProfileStub`
- New struct `ElasticProfile`
- New struct `ElasticProfileContainerGroupNamingPolicy`
- New struct `ElasticProfileContainerGroupNamingPolicyGUIDNamingPolicy`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `FileShare`
- New struct `FileShareProperties`
- New struct `IdentityACLs`
- New struct `IdentityAccessControl`
- New struct `LoadBalancer`
- New struct `LoadBalancerBackendAddressPool`
- New struct `NGroup`
- New struct `NGroupCGPropertyContainer`
- New struct `NGroupCGPropertyContainerProperties`
- New struct `NGroupCGPropertyVolume`
- New struct `NGroupContainerGroupProperties`
- New struct `NGroupIdentity`
- New struct `NGroupProperties`
- New struct `NGroupSKUs`
- New struct `NGroupsListResult`
- New struct `NGroupsSKUsList`
- New struct `NetworkProfile`
- New struct `PlacementProfile`
- New struct `SecretReference`
- New struct `StorageProfile`
- New struct `SystemData`
- New struct `UpdateProfile`
- New struct `UpdateProfileRollingUpdateProfile`
- New field `StorageAccountKeyReference` in struct `AzureFileVolume`
- New field `SystemData` in struct `ContainerGroupProfile`
- New field `ConfidentialComputeProperties`, `Containers`, `Diagnostics`, `EncryptionProperties`, `Extensions`, `IPAddress`, `ImageRegistryCredentials`, `InitContainers`, `OSType`, `Priority`, `RegisteredRevisions`, `RestartPolicy`, `Revision`, `SKU`, `SecurityContext`, `ShutdownGracePeriod`, `TimeToLive`, `UseKrypton`, `Volumes` in struct `ContainerGroupProfileProperties`
- New field `IdentityACLs`, `SecretReferences` in struct `ContainerGroupPropertiesProperties`
- New field `SecureValueReference` in struct `EnvironmentVariable`
- New field `PasswordReference` in struct `ImageRegistryCredential`
- New field `SecretReference` in struct `Volume`


## 2.5.0-beta.1 (2024-10-23)
### Features Added

- New function `*ClientFactory.NewContainerGroupProfileClient() *ContainerGroupProfileClient`
- New function `*ClientFactory.NewContainerGroupProfilesClient() *ContainerGroupProfilesClient`
- New function `NewContainerGroupProfileClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ContainerGroupProfileClient, error)`
- New function `*ContainerGroupProfileClient.GetByRevisionNumber(context.Context, string, string, string, *ContainerGroupProfileClientGetByRevisionNumberOptions) (ContainerGroupProfileClientGetByRevisionNumberResponse, error)`
- New function `*ContainerGroupProfileClient.NewListAllRevisionsPager(string, string, *ContainerGroupProfileClientListAllRevisionsOptions) *runtime.Pager[ContainerGroupProfileClientListAllRevisionsResponse]`
- New function `NewContainerGroupProfilesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ContainerGroupProfilesClient, error)`
- New function `*ContainerGroupProfilesClient.CreateOrUpdate(context.Context, string, string, ContainerGroupProfile, *ContainerGroupProfilesClientCreateOrUpdateOptions) (ContainerGroupProfilesClientCreateOrUpdateResponse, error)`
- New function `*ContainerGroupProfilesClient.Delete(context.Context, string, string, *ContainerGroupProfilesClientDeleteOptions) (ContainerGroupProfilesClientDeleteResponse, error)`
- New function `*ContainerGroupProfilesClient.Get(context.Context, string, string, *ContainerGroupProfilesClientGetOptions) (ContainerGroupProfilesClientGetResponse, error)`
- New function `*ContainerGroupProfilesClient.NewListByResourceGroupPager(string, *ContainerGroupProfilesClientListByResourceGroupOptions) *runtime.Pager[ContainerGroupProfilesClientListByResourceGroupResponse]`
- New function `*ContainerGroupProfilesClient.NewListPager(*ContainerGroupProfilesClientListOptions) *runtime.Pager[ContainerGroupProfilesClientListResponse]`
- New function `*ContainerGroupProfilesClient.Patch(context.Context, string, string, ContainerGroupProfilePatch, *ContainerGroupProfilesClientPatchOptions) (ContainerGroupProfilesClientPatchResponse, error)`
- New struct `ConfigMap`
- New struct `ContainerGroupProfile`
- New struct `ContainerGroupProfileListResult`
- New struct `ContainerGroupProfilePatch`
- New struct `ContainerGroupProfileProperties`
- New struct `ContainerGroupProfilePropertiesProperties`
- New struct `ContainerGroupProfileReferenceDefinition`
- New struct `StandbyPoolProfileDefinition`
- New field `ContainerGroupProfile`, `IsCreatedFromStandbyPool`, `StandbyPoolProfile` in struct `ContainerGroupPropertiesProperties`
- New field `ConfigMap` in struct `ContainerProperties`


## 2.4.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.3.0 (2023-04-28)
### Features Added

- New value `ContainerGroupSKUConfidential` added to enum type `ContainerGroupSKU`
- New enum type `ContainerGroupPriority` with values `ContainerGroupPriorityRegular`, `ContainerGroupPrioritySpot`
- New struct `ConfidentialComputeProperties`
- New struct `SecurityContextCapabilitiesDefinition`
- New struct `SecurityContextDefinition`
- New field `ConfidentialComputeProperties`, `Priority` in struct `ContainerGroupPropertiesProperties`
- New field `SecurityContext` in struct `ContainerProperties`
- New field `SecurityContext` in struct `InitContainerPropertiesDefinition`


## 2.2.1 (2023-04-18)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 2.2.0 (2023-04-07)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 2.1.0 (2022-11-08)
### Features Added

- New struct `DeploymentExtensionSpec`
- New struct `DeploymentExtensionSpecProperties`
- New field `Extensions` in struct `ContainerGroupPropertiesProperties`
- New field `Identity` in struct `EncryptionProperties`


## 2.0.0 (2022-08-26)
### Breaking Changes

- Type of `ContainerGroup.Properties` has been changed from `*ContainerGroupProperties` to `*ContainerGroupPropertiesProperties`
- Type of `ContainerGroupIdentity.UserAssignedIdentities` has been changed from `map[string]*Components10Wh5UdSchemasContainergroupidentityPropertiesUserassignedidentitiesAdditionalproperties` to `map[string]*UserAssignedIdentities`
- Const `AutoGeneratedDomainNameLabelScopeResourceGroupReuse` has been removed
- Const `AutoGeneratedDomainNameLabelScopeSubscriptionReuse` has been removed
- Const `AutoGeneratedDomainNameLabelScopeNoreuse` has been removed
- Const `AutoGeneratedDomainNameLabelScopeTenantReuse` has been removed
- Const `AutoGeneratedDomainNameLabelScopeUnsecure` has been removed
- Type alias `AutoGeneratedDomainNameLabelScope` has been removed
- Function `PossibleAutoGeneratedDomainNameLabelScopeValues` has been removed
- Struct `Components10Wh5UdSchemasContainergroupidentityPropertiesUserassignedidentitiesAdditionalproperties` has been removed
- Field `DNSNameLabelReusePolicy` of struct `IPAddress` has been removed
- Field `Containers` of struct `ContainerGroupProperties` has been removed
- Field `SubnetIDs` of struct `ContainerGroupProperties` has been removed
- Field `RestartPolicy` of struct `ContainerGroupProperties` has been removed
- Field `Volumes` of struct `ContainerGroupProperties` has been removed
- Field `ProvisioningState` of struct `ContainerGroupProperties` has been removed
- Field `Diagnostics` of struct `ContainerGroupProperties` has been removed
- Field `SKU` of struct `ContainerGroupProperties` has been removed
- Field `InstanceView` of struct `ContainerGroupProperties` has been removed
- Field `OSType` of struct `ContainerGroupProperties` has been removed
- Field `EncryptionProperties` of struct `ContainerGroupProperties` has been removed
- Field `InitContainers` of struct `ContainerGroupProperties` has been removed
- Field `DNSConfig` of struct `ContainerGroupProperties` has been removed
- Field `IPAddress` of struct `ContainerGroupProperties` has been removed
- Field `ImageRegistryCredentials` of struct `ContainerGroupProperties` has been removed

### Features Added

- New const `DNSNameLabelReusePolicyNoreuse`
- New const `DNSNameLabelReusePolicyUnsecure`
- New const `DNSNameLabelReusePolicySubscriptionReuse`
- New const `DNSNameLabelReusePolicyTenantReuse`
- New const `DNSNameLabelReusePolicyResourceGroupReuse`
- New type alias `DNSNameLabelReusePolicy`
- New function `*SubnetServiceAssociationLinkClient.BeginDelete(context.Context, string, string, string, *SubnetServiceAssociationLinkClientBeginDeleteOptions) (*runtime.Poller[SubnetServiceAssociationLinkClientDeleteResponse], error)`
- New function `NewSubnetServiceAssociationLinkClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SubnetServiceAssociationLinkClient, error)`
- New function `PossibleDNSNameLabelReusePolicyValues() []DNSNameLabelReusePolicy`
- New struct `ContainerGroupPropertiesProperties`
- New struct `SubnetServiceAssociationLinkClient`
- New struct `SubnetServiceAssociationLinkClientBeginDeleteOptions`
- New struct `SubnetServiceAssociationLinkClientDeleteResponse`
- New struct `UserAssignedIdentities`
- New field `Properties` in struct `ContainerGroupProperties`
- New field `Identity` in struct `ContainerGroupProperties`
- New field `ID` in struct `Usage`
- New field `AutoGeneratedDomainNameLabelScope` in struct `IPAddress`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerinstance/armcontainerinstance` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).