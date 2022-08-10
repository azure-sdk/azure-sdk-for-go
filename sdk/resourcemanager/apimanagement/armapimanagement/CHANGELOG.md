# Release History

## 1.1.0-beta.1 (2022-08-10)
### Features Added

- New const `PolicyFragmentContentFormatXML`
- New const `PortalSettingsCspModeDisabled`
- New const `PortalSettingsCspModeEnabled`
- New const `PolicyFragmentContentFormatRawxml`
- New const `PortalSettingsCspModeReportOnly`
- New function `PossiblePolicyFragmentContentFormatValues() []PolicyFragmentContentFormat`
- New function `*PolicyFragmentClient.ListReferences(context.Context, string, string, string, *PolicyFragmentClientListReferencesOptions) (PolicyFragmentClientListReferencesResponse, error)`
- New function `*PolicyFragmentClient.BeginCreateOrUpdate(context.Context, string, string, string, PolicyFragmentContract, *PolicyFragmentClientBeginCreateOrUpdateOptions) (*runtime.Poller[PolicyFragmentClientCreateOrUpdateResponse], error)`
- New function `*PortalConfigClient.CreateOrUpdate(context.Context, string, string, string, string, PortalConfigContract, *PortalConfigClientCreateOrUpdateOptions) (PortalConfigClientCreateOrUpdateResponse, error)`
- New function `*PolicyFragmentClient.ListByService(context.Context, string, string, *PolicyFragmentClientListByServiceOptions) (PolicyFragmentClientListByServiceResponse, error)`
- New function `*PolicyFragmentClient.Get(context.Context, string, string, string, *PolicyFragmentClientGetOptions) (PolicyFragmentClientGetResponse, error)`
- New function `NewPortalConfigClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PortalConfigClient, error)`
- New function `PossiblePortalSettingsCspModeValues() []PortalSettingsCspMode`
- New function `*PolicyFragmentClient.GetEntityTag(context.Context, string, string, string, *PolicyFragmentClientGetEntityTagOptions) (PolicyFragmentClientGetEntityTagResponse, error)`
- New function `*PortalConfigClient.Get(context.Context, string, string, string, *PortalConfigClientGetOptions) (PortalConfigClientGetResponse, error)`
- New function `*PortalConfigClient.Update(context.Context, string, string, string, string, PortalConfigContract, *PortalConfigClientUpdateOptions) (PortalConfigClientUpdateResponse, error)`
- New function `NewPolicyFragmentClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PolicyFragmentClient, error)`
- New function `*PortalConfigClient.ListByService(context.Context, string, string, *PortalConfigClientListByServiceOptions) (PortalConfigClientListByServiceResponse, error)`
- New function `*PortalConfigClient.GetEntityTag(context.Context, string, string, string, *PortalConfigClientGetEntityTagOptions) (PortalConfigClientGetEntityTagResponse, error)`
- New function `*PolicyFragmentClient.Delete(context.Context, string, string, string, string, *PolicyFragmentClientDeleteOptions) (PolicyFragmentClientDeleteResponse, error)`
- New struct `PolicyFragmentClient`
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
- New struct `PortalConfigClient`
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
- New field `ClientLibrary` in struct `IdentityProviderBaseParameters`
- New field `ClientLibrary` in struct `IdentityProviderCreateContractProperties`
- New field `Parameters` in struct `ContentTypeClientCreateOrUpdateOptions`
- New field `ClientLibrary` in struct `IdentityProviderContractProperties`
- New field `ClientLibrary` in struct `IdentityProviderUpdateProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/apimanagement/armapimanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).