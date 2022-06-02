# Release History

## 1.1.0-beta.1 (2022-06-02)
### Features Added

- New const `PolicyFragmentContentFormatRawxml`
- New const `PolicyFragmentContentFormatXML`
- New const `PortalSettingsCspModeEnabled`
- New const `PortalSettingsCspModeDisabled`
- New const `PortalSettingsCspModeReportOnly`
- New function `PossiblePolicyFragmentContentFormatValues() []PolicyFragmentContentFormat`
- New function `PossiblePortalSettingsCspModeValues() []PortalSettingsCspMode`
- New struct `PolicyFragmentClientBeginCreateOrUpdateOptions`
- New struct `PolicyFragmentClientCreateOrUpdateResponse`
- New struct `PolicyFragmentClientDeleteOptions`
- New struct `PolicyFragmentClientDeleteResponse`
- New struct `PolicyFragmentClientGetEntityTagOptions`
- New struct `PolicyFragmentClientGetEntityTagResponse`
- New struct `PolicyFragmentClientGetOptions`
- New struct `PolicyFragmentClientGetResponse`
- New struct `PolicyFragmentClientListByServiceOptions`
- New struct `PolicyFragmentClientListByServiceResponse`
- New struct `PolicyFragmentClientListReferencesOptions`
- New struct `PolicyFragmentClientListReferencesResponse`
- New struct `PolicyFragmentCollection`
- New struct `PolicyFragmentContract`
- New struct `PolicyFragmentContractProperties`
- New struct `PortalConfigClientCreateOrUpdateOptions`
- New struct `PortalConfigClientCreateOrUpdateResponse`
- New struct `PortalConfigClientGetEntityTagOptions`
- New struct `PortalConfigClientGetEntityTagResponse`
- New struct `PortalConfigClientGetOptions`
- New struct `PortalConfigClientGetResponse`
- New struct `PortalConfigClientListByServiceOptions`
- New struct `PortalConfigClientListByServiceResponse`
- New struct `PortalConfigClientUpdateOptions`
- New struct `PortalConfigClientUpdateResponse`
- New struct `PortalConfigCollection`
- New struct `PortalConfigContract`
- New struct `PortalConfigCorsProperties`
- New struct `PortalConfigCspProperties`
- New struct `PortalConfigDelegationProperties`
- New struct `PortalConfigProperties`
- New struct `PortalConfigPropertiesSignin`
- New struct `PortalConfigPropertiesSignup`
- New struct `PortalConfigTermsOfServiceProperties`
- New struct `ResourceCollection`
- New struct `ResourceCollectionValueItem`
- New field `Parameters` in struct `ContentTypeClientCreateOrUpdateOptions`
- New field `OAuth2AuthenticationSettings` in struct `AuthenticationSettingsContract`
- New field `UseInAPIDocumentation` in struct `AuthorizationServerUpdateContractProperties`
- New field `UseInTestConsole` in struct `AuthorizationServerUpdateContractProperties`
- New field `ClientLibrary` in struct `IdentityProviderCreateContractProperties`
- New field `ClientLibrary` in struct `IdentityProviderUpdateProperties`
- New field `ClientLibrary` in struct `IdentityProviderContractProperties`
- New field `ClientLibrary` in struct `IdentityProviderBaseParameters`
- New field `UseInTestConsole` in struct `AuthorizationServerContractProperties`
- New field `UseInAPIDocumentation` in struct `AuthorizationServerContractProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).