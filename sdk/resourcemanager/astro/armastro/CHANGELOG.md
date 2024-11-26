# Release History

## 0.2.0 (2024-11-26)
### Features Added

- New enum type `RenewalMode` with values `RenewalModeAuto`, `RenewalModeManual`
- New function `*OrganizationsClient.GetResources(context.Context, string, string, GetResourcesRequest, *OrganizationsClientGetResourcesOptions) (OrganizationsClientGetResourcesResponse, error)`
- New function `*OrganizationsClient.GetRoles(context.Context, string, string, GetRolesRequest, *OrganizationsClientGetRolesOptions) (OrganizationsClientGetRolesResponse, error)`
- New function `*OrganizationsClient.GetUsers(context.Context, string, string, GetUsersRequest, *OrganizationsClientGetUsersOptions) (OrganizationsClientGetUsersResponse, error)`
- New function `*OrganizationsClient.ManageRoles(context.Context, string, string, ManageRolesModel, *OrganizationsClientManageRolesOptions) (OrganizationsClientManageRolesResponse, error)`
- New function `*OrganizationsClient.RemoveUser(context.Context, string, string, RemoveUserRequest, *OrganizationsClientRemoveUserOptions) (OrganizationsClientRemoveUserResponse, error)`
- New struct `GetResourcesRequest`
- New struct `GetResourcesSuccessResponse`
- New struct `GetRolesRequest`
- New struct `GetRolesSuccessResponse`
- New struct `GetUsersRequest`
- New struct `GetUsersSuccessResponse`
- New struct `LiftrBaseMarketplaceDetailsUpdate`
- New struct `LiftrBaseOfferDetailsUpdate`
- New struct `ManageRolesModel`
- New struct `PageInfo`
- New struct `PartnerResource`
- New struct `RemoveUserRequest`
- New struct `Role`
- New struct `User`
- New field `EndDate`, `RenewalMode` in struct `LiftrBaseOfferDetails`
- New field `Marketplace` in struct `OrganizationResourceUpdateProperties`


## 0.1.0 (2024-02-23)
### Other Changes

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/astro/armastro` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).