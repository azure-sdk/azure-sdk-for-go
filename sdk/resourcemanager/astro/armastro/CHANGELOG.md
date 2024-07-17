# Release History

## 1.0.0 (2024-07-17)
### Breaking Changes

- Type of `OrganizationResource.Identity` has been changed from `*ManagedServiceIdentity` to `*ManagedServiceIdentityV4`
- Type of `OrganizationResource.Properties` has been changed from `*LiftrBaseDataOrganizationProperties` to `*OrganizationProperties`
- Type of `OrganizationResourceUpdate.Identity` has been changed from `*ManagedServiceIdentity` to `*ManagedServiceIdentityV4`
- Type of `OrganizationResourceUpdateProperties.PartnerOrganizationProperties` has been changed from `*LiftrBaseDataPartnerOrganizationPropertiesUpdate` to `*PartnerOrganizationProperties`
- Type of `OrganizationResourceUpdateProperties.User` has been changed from `*LiftrBaseUserDetailsUpdate` to `*UserDetails`
- `ManagedServiceIdentityTypeSystemAssignedUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Struct `LiftrBaseDataOrganizationProperties` has been removed
- Struct `LiftrBaseDataPartnerOrganizationProperties` has been removed
- Struct `LiftrBaseDataPartnerOrganizationPropertiesUpdate` has been removed
- Struct `LiftrBaseMarketplaceDetails` has been removed
- Struct `LiftrBaseOfferDetails` has been removed
- Struct `LiftrBaseSingleSignOnProperties` has been removed
- Struct `LiftrBaseUserDetails` has been removed
- Struct `LiftrBaseUserDetailsUpdate` has been removed
- Struct `ManagedServiceIdentity` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAndUserAssigned` added to enum type `ManagedServiceIdentityType`
- New struct `ManagedServiceIdentityV4`
- New struct `MarketplaceDetails`
- New struct `OfferDetails`
- New struct `OrganizationProperties`
- New struct `PartnerOrganizationProperties`
- New struct `SingleSignOnProperties`
- New struct `UserDetails`


## 0.1.0 (2024-02-23)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/astro/armastro` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).