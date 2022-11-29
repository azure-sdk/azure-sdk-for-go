# Release History

## 2.0.0-beta.1 (2022-11-29)
### Breaking Changes

- Struct `LinkerList` has been removed
- Struct `SourceConfigurationResult` has been removed
- Field `SourceConfigurationResult` of struct `LinkerClientListConfigurationsResponse` has been removed
- Field `LinkerList` of struct `LinkerClientListResponse` has been removed

### Features Added

- New const `AccessKeyPermissionsListen`
- New const `AccessKeyPermissionsManage`
- New const `AccessKeyPermissionsRead`
- New const `AccessKeyPermissionsSend`
- New const `AccessKeyPermissionsWrite`
- New const `ActionTypeEnable`
- New const `ActionTypeOptOut`
- New const `AllowTypeFalse`
- New const `AllowTypeTrue`
- New const `AuthTypeAccessKey`
- New const `AuthTypeUserAccount`
- New const `ClientTypeKafkaSpringBoot`
- New const `DeleteOrUpdateBehaviorDefault`
- New const `DeleteOrUpdateBehaviorForcedCleanup`
- New const `DryrunActionNameCreateOrUpdate`
- New const `DryrunPrerequisiteResultTypeBasicError`
- New const `DryrunPrerequisiteResultTypePermissionsMissing`
- New const `DryrunPreviewOperationTypeConfigAuth`
- New const `DryrunPreviewOperationTypeConfigConnection`
- New const `DryrunPreviewOperationTypeConfigNetwork`
- New const `TargetServiceTypeSelfHostedServer`
- New type alias `AccessKeyPermissions`
- New type alias `AllowType`
- New type alias `DeleteOrUpdateBehavior`
- New type alias `DryrunActionName`
- New type alias `DryrunPrerequisiteResultType`
- New type alias `DryrunPreviewOperationType`
- New function `*AccessKeyInfoBase.GetAuthInfoBase() *AuthInfoBase`
- New function `*BasicErrorDryrunPrerequisiteResult.GetDryrunPrerequisiteResult() *DryrunPrerequisiteResult`
- New function `NewConfigurationNamesClient(azcore.TokenCredential, *arm.ClientOptions) (*ConfigurationNamesClient, error)`
- New function `*ConfigurationNamesClient.NewListPager(*ConfigurationNamesClientListOptions) *runtime.Pager[ConfigurationNamesClientListResponse]`
- New function `NewConnectorClient(azcore.TokenCredential, *arm.ClientOptions) (*ConnectorClient, error)`
- New function `*ConnectorClient.BeginCreateDryrun(context.Context, string, string, string, string, DryrunResource, *ConnectorClientBeginCreateDryrunOptions) (*runtime.Poller[ConnectorClientCreateDryrunResponse], error)`
- New function `*ConnectorClient.BeginCreateOrUpdate(context.Context, string, string, string, string, LinkerResource, *ConnectorClientBeginCreateOrUpdateOptions) (*runtime.Poller[ConnectorClientCreateOrUpdateResponse], error)`
- New function `*ConnectorClient.BeginDelete(context.Context, string, string, string, string, *ConnectorClientBeginDeleteOptions) (*runtime.Poller[ConnectorClientDeleteResponse], error)`
- New function `*ConnectorClient.DeleteDryrun(context.Context, string, string, string, string, *ConnectorClientDeleteDryrunOptions) (ConnectorClientDeleteDryrunResponse, error)`
- New function `*ConnectorClient.GenerateConfigurations(context.Context, string, string, string, string, *ConnectorClientGenerateConfigurationsOptions) (ConnectorClientGenerateConfigurationsResponse, error)`
- New function `*ConnectorClient.Get(context.Context, string, string, string, string, *ConnectorClientGetOptions) (ConnectorClientGetResponse, error)`
- New function `*ConnectorClient.GetDryrun(context.Context, string, string, string, string, *ConnectorClientGetDryrunOptions) (ConnectorClientGetDryrunResponse, error)`
- New function `*ConnectorClient.NewListDryrunPager(string, string, string, *ConnectorClientListDryrunOptions) *runtime.Pager[ConnectorClientListDryrunResponse]`
- New function `*ConnectorClient.NewListPager(string, string, string, *ConnectorClientListOptions) *runtime.Pager[ConnectorClientListResponse]`
- New function `*ConnectorClient.BeginUpdate(context.Context, string, string, string, string, LinkerPatch, *ConnectorClientBeginUpdateOptions) (*runtime.Poller[ConnectorClientUpdateResponse], error)`
- New function `*ConnectorClient.BeginUpdateDryrun(context.Context, string, string, string, string, DryrunPatch, *ConnectorClientBeginUpdateDryrunOptions) (*runtime.Poller[ConnectorClientUpdateDryrunResponse], error)`
- New function `*ConnectorClient.BeginValidate(context.Context, string, string, string, string, *ConnectorClientBeginValidateOptions) (*runtime.Poller[ConnectorClientValidateResponse], error)`
- New function `*CreateOrUpdateDryrunParameters.GetDryrunParameters() *DryrunParameters`
- New function `*DryrunParameters.GetDryrunParameters() *DryrunParameters`
- New function `*DryrunPrerequisiteResult.GetDryrunPrerequisiteResult() *DryrunPrerequisiteResult`
- New function `NewLinkersClient(azcore.TokenCredential, *arm.ClientOptions) (*LinkersClient, error)`
- New function `*LinkersClient.BeginCreateDryrun(context.Context, string, string, DryrunResource, *LinkersClientBeginCreateDryrunOptions) (*runtime.Poller[LinkersClientCreateDryrunResponse], error)`
- New function `*LinkersClient.DeleteDryrun(context.Context, string, string, *LinkersClientDeleteDryrunOptions) (LinkersClientDeleteDryrunResponse, error)`
- New function `*LinkersClient.GenerateConfigurations(context.Context, string, string, *LinkersClientGenerateConfigurationsOptions) (LinkersClientGenerateConfigurationsResponse, error)`
- New function `*LinkersClient.GetDryrun(context.Context, string, string, *LinkersClientGetDryrunOptions) (LinkersClientGetDryrunResponse, error)`
- New function `*LinkersClient.NewListDryrunPager(string, *LinkersClientListDryrunOptions) *runtime.Pager[LinkersClientListDryrunResponse]`
- New function `*LinkersClient.BeginUpdateDryrun(context.Context, string, string, DryrunPatch, *LinkersClientBeginUpdateDryrunOptions) (*runtime.Poller[LinkersClientUpdateDryrunResponse], error)`
- New function `*PermissionsMissingDryrunPrerequisiteResult.GetDryrunPrerequisiteResult() *DryrunPrerequisiteResult`
- New function `PossibleAccessKeyPermissionsValues() []AccessKeyPermissions`
- New function `PossibleAllowTypeValues() []AllowType`
- New function `PossibleDeleteOrUpdateBehaviorValues() []DeleteOrUpdateBehavior`
- New function `PossibleDryrunActionNameValues() []DryrunActionName`
- New function `PossibleDryrunPrerequisiteResultTypeValues() []DryrunPrerequisiteResultType`
- New function `PossibleDryrunPreviewOperationTypeValues() []DryrunPreviewOperationType`
- New function `*SelfHostedServer.GetTargetServiceBase() *TargetServiceBase`
- New function `*UserAccountAuthInfo.GetAuthInfoBase() *AuthInfoBase`
- New struct `AccessKeyInfoBase`
- New struct `BasicErrorDryrunPrerequisiteResult`
- New struct `ConfigurationInfo`
- New struct `ConfigurationName`
- New struct `ConfigurationNameItem`
- New struct `ConfigurationNameResult`
- New struct `ConfigurationNames`
- New struct `ConfigurationNamesClient`
- New struct `ConfigurationNamesClientListOptions`
- New struct `ConfigurationNamesClientListResponse`
- New struct `ConfigurationResult`
- New struct `ConnectorClient`
- New struct `ConnectorClientBeginCreateDryrunOptions`
- New struct `ConnectorClientBeginCreateOrUpdateOptions`
- New struct `ConnectorClientBeginDeleteOptions`
- New struct `ConnectorClientBeginUpdateDryrunOptions`
- New struct `ConnectorClientBeginUpdateOptions`
- New struct `ConnectorClientBeginValidateOptions`
- New struct `ConnectorClientCreateDryrunResponse`
- New struct `ConnectorClientCreateOrUpdateResponse`
- New struct `ConnectorClientDeleteDryrunOptions`
- New struct `ConnectorClientDeleteDryrunResponse`
- New struct `ConnectorClientDeleteResponse`
- New struct `ConnectorClientGenerateConfigurationsOptions`
- New struct `ConnectorClientGenerateConfigurationsResponse`
- New struct `ConnectorClientGetDryrunOptions`
- New struct `ConnectorClientGetDryrunResponse`
- New struct `ConnectorClientGetOptions`
- New struct `ConnectorClientGetResponse`
- New struct `ConnectorClientListDryrunOptions`
- New struct `ConnectorClientListDryrunResponse`
- New struct `ConnectorClientListOptions`
- New struct `ConnectorClientListResponse`
- New struct `ConnectorClientUpdateDryrunResponse`
- New struct `ConnectorClientUpdateResponse`
- New struct `ConnectorClientValidateResponse`
- New struct `CreateOrUpdateDryrunParameters`
- New struct `DatabaseAADAuthInfo`
- New struct `DryrunList`
- New struct `DryrunOperationPreview`
- New struct `DryrunParameters`
- New struct `DryrunPatch`
- New struct `DryrunPrerequisiteResult`
- New struct `DryrunProperties`
- New struct `DryrunResource`
- New struct `FirewallRules`
- New struct `LinkersClient`
- New struct `LinkersClientBeginCreateDryrunOptions`
- New struct `LinkersClientBeginUpdateDryrunOptions`
- New struct `LinkersClientCreateDryrunResponse`
- New struct `LinkersClientDeleteDryrunOptions`
- New struct `LinkersClientDeleteDryrunResponse`
- New struct `LinkersClientGenerateConfigurationsOptions`
- New struct `LinkersClientGenerateConfigurationsResponse`
- New struct `LinkersClientGetDryrunOptions`
- New struct `LinkersClientGetDryrunResponse`
- New struct `LinkersClientListDryrunOptions`
- New struct `LinkersClientListDryrunResponse`
- New struct `LinkersClientUpdateDryrunResponse`
- New struct `PermissionsMissingDryrunPrerequisiteResult`
- New struct `PublicNetworkSolution`
- New struct `ResourceList`
- New struct `SelfHostedServer`
- New struct `UserAccountAuthInfo`
- New anonymous field `ConfigurationResult` in struct `LinkerClientListConfigurationsResponse`
- New anonymous field `ResourceList` in struct `LinkerClientListResponse`
- New field `ConfigurationInfo` in struct `LinkerProperties`
- New field `PublicNetworkSolution` in struct `LinkerProperties`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `KeyVaultSecretName` in struct `SecretStore`
- New field `DeleteOrUpdateBehavior` in struct `ServicePrincipalCertificateAuthInfo`
- New field `Roles` in struct `ServicePrincipalCertificateAuthInfo`
- New field `DeleteOrUpdateBehavior` in struct `ServicePrincipalSecretAuthInfo`
- New field `Roles` in struct `ServicePrincipalSecretAuthInfo`
- New field `UserName` in struct `ServicePrincipalSecretAuthInfo`
- New field `DeleteOrUpdateBehavior` in struct `SystemAssignedIdentityAuthInfo`
- New field `Roles` in struct `SystemAssignedIdentityAuthInfo`
- New field `UserName` in struct `SystemAssignedIdentityAuthInfo`
- New field `DeleteOrUpdateBehavior` in struct `UserAssignedIdentityAuthInfo`
- New field `Roles` in struct `UserAssignedIdentityAuthInfo`
- New field `UserName` in struct `UserAssignedIdentityAuthInfo`
- New field `DeleteOrUpdateBehavior` in struct `VNetSolution`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/servicelinker/armservicelinker` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).