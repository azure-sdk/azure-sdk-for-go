# Release History

## 2.0.0 (2022-09-21)
### Breaking Changes

- Function `*ApplicationDefinitionsClient.BeginCreateOrUpdate` has been removed
- Function `*ApplicationsClient.Update` has been removed
- Function `*ApplicationDefinitionsClient.BeginDelete` has been removed
- Struct `ApplicationDefinitionsClientBeginCreateOrUpdateOptions` has been removed
- Struct `ApplicationDefinitionsClientBeginDeleteOptions` has been removed
- Struct `ApplicationPropertiesPatchable` has been removed
- Struct `ApplicationsClientUpdateOptions` has been removed
- Field `Application` of struct `ApplicationsClientUpdateResponse` has been removed
- Field `ProvisioningState` of struct `ApplicationDefinitionProperties` has been removed

### Features Added

- New const `StatusRemove`
- New const `SubstatusExpired`
- New const `SubstatusFailed`
- New const `SubstatusApproved`
- New const `SubstatusTimeout`
- New const `StatusNotSpecified`
- New const `SubstatusDenied`
- New const `StatusElevate`
- New const `SubstatusNotSpecified`
- New type alias `Status`
- New type alias `Substatus`
- New function `*ApplicationsClient.GetByID(context.Context, string, *ApplicationsClientGetByIDOptions) (ApplicationsClientGetByIDResponse, error)`
- New function `*ApplicationDefinitionsClient.CreateOrUpdateByID(context.Context, string, string, ApplicationDefinition, *ApplicationDefinitionsClientCreateOrUpdateByIDOptions) (ApplicationDefinitionsClientCreateOrUpdateByIDResponse, error)`
- New function `*ApplicationDefinitionsClient.UpdateByID(context.Context, string, string, ApplicationDefinitionPatchable, *ApplicationDefinitionsClientUpdateByIDOptions) (ApplicationDefinitionsClientUpdateByIDResponse, error)`
- New function `*ApplicationsClient.BeginDeleteByID(context.Context, string, *ApplicationsClientBeginDeleteByIDOptions) (*runtime.Poller[ApplicationsClientDeleteByIDResponse], error)`
- New function `*ApplicationDefinitionsClient.GetByID(context.Context, string, string, *ApplicationDefinitionsClientGetByIDOptions) (ApplicationDefinitionsClientGetByIDResponse, error)`
- New function `*ApplicationsClient.BeginUpdateAccess(context.Context, string, string, UpdateAccessDefinition, *ApplicationsClientBeginUpdateAccessOptions) (*runtime.Poller[ApplicationsClientUpdateAccessResponse], error)`
- New function `*ApplicationDefinitionsClient.Delete(context.Context, string, string, *ApplicationDefinitionsClientDeleteOptions) (ApplicationDefinitionsClientDeleteResponse, error)`
- New function `*ApplicationsClient.ListTokens(context.Context, string, string, ListTokenRequest, *ApplicationsClientListTokensOptions) (ApplicationsClientListTokensResponse, error)`
- New function `*ApplicationDefinitionsClient.CreateOrUpdate(context.Context, string, string, ApplicationDefinition, *ApplicationDefinitionsClientCreateOrUpdateOptions) (ApplicationDefinitionsClientCreateOrUpdateResponse, error)`
- New function `*ApplicationDefinitionsClient.DeleteByID(context.Context, string, string, *ApplicationDefinitionsClientDeleteByIDOptions) (ApplicationDefinitionsClientDeleteByIDResponse, error)`
- New function `*ApplicationsClient.BeginCreateOrUpdateByID(context.Context, string, Application, *ApplicationsClientBeginCreateOrUpdateByIDOptions) (*runtime.Poller[ApplicationsClientCreateOrUpdateByIDResponse], error)`
- New function `PossibleSubstatusValues() []Substatus`
- New function `*ApplicationsClient.BeginUpdateByID(context.Context, string, *ApplicationsClientBeginUpdateByIDOptions) (*runtime.Poller[ApplicationsClientUpdateByIDResponse], error)`
- New function `PossibleStatusValues() []Status`
- New function `*ApplicationsClient.BeginUpdate(context.Context, string, string, *ApplicationsClientBeginUpdateOptions) (*runtime.Poller[ApplicationsClientUpdateResponse], error)`
- New struct `AllowedUpgradePlansResult`
- New struct `ApplicationDefinitionsClientCreateOrUpdateByIDOptions`
- New struct `ApplicationDefinitionsClientCreateOrUpdateByIDResponse`
- New struct `ApplicationDefinitionsClientCreateOrUpdateOptions`
- New struct `ApplicationDefinitionsClientDeleteByIDOptions`
- New struct `ApplicationDefinitionsClientDeleteByIDResponse`
- New struct `ApplicationDefinitionsClientDeleteOptions`
- New struct `ApplicationDefinitionsClientGetByIDOptions`
- New struct `ApplicationDefinitionsClientGetByIDResponse`
- New struct `ApplicationDefinitionsClientUpdateByIDOptions`
- New struct `ApplicationDefinitionsClientUpdateByIDResponse`
- New struct `ApplicationsClientBeginCreateOrUpdateByIDOptions`
- New struct `ApplicationsClientBeginDeleteByIDOptions`
- New struct `ApplicationsClientBeginUpdateAccessOptions`
- New struct `ApplicationsClientBeginUpdateByIDOptions`
- New struct `ApplicationsClientBeginUpdateOptions`
- New struct `ApplicationsClientCreateOrUpdateByIDResponse`
- New struct `ApplicationsClientDeleteByIDResponse`
- New struct `ApplicationsClientGetByIDOptions`
- New struct `ApplicationsClientGetByIDResponse`
- New struct `ApplicationsClientListTokensOptions`
- New struct `ApplicationsClientListTokensResponse`
- New struct `ApplicationsClientUpdateAccessResponse`
- New struct `ApplicationsClientUpdateByIDResponse`
- New struct `JitRequestMetadata`
- New struct `ListTokenRequest`
- New struct `ManagedIdentityToken`
- New struct `ManagedIdentityTokenResult`
- New struct `UpdateAccessDefinition`
- New anonymous field `AllowedUpgradePlansResult` in struct `ApplicationsClientListAllowedUpgradePlansResponse`
- New anonymous field `ApplicationPatchable` in struct `ApplicationsClientUpdateResponse`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/solutions/armmanagedapplications` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).