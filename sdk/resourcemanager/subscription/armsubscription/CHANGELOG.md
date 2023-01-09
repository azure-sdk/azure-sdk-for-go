# Release History

## 1.1.0 (2023-01-09)
### Features Added

- New type alias `Provisioning` with values `ProvisioningAccepted`, `ProvisioningPending`, `ProvisioningSucceeded`
- New field `ProvisioningState` in struct `AcceptOwnershipStatusResponse`
- New field `Tags` in struct `Subscription`
- New field `TenantID` in struct `Subscription`
- New field `Country` in struct `TenantIDDescription`
- New field `CountryCode` in struct `TenantIDDescription`
- New field `DefaultDomain` in struct `TenantIDDescription`
- New field `DisplayName` in struct `TenantIDDescription`
- New field `Domains` in struct `TenantIDDescription`
- New field `TenantCategory` in struct `TenantIDDescription`
- New field `TenantType` in struct `TenantIDDescription`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).