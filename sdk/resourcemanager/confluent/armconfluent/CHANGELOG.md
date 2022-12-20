# Release History

## 1.1.0-beta.1 (2022-12-20)
### Features Added

- New function `*ValidationsClient.ValidateOrganizationV2(context.Context, string, string, OrganizationResource, *ValidationsClientValidateOrganizationV2Options) (ValidationsClientValidateOrganizationV2Response, error)`
- New struct `LinkOrganization`
- New struct `ValidationResponse`
- New field `PrivateOfferID` in struct `OfferDetail`
- New field `PrivateOfferIDs` in struct `OfferDetail`
- New field `TermID` in struct `OfferDetail`
- New field `LinkOrganization` in struct `OrganizationResourceProperties`
- New field `AADEmail` in struct `UserDetail`
- New field `UserPrincipalName` in struct `UserDetail`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/confluent/armconfluent` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).