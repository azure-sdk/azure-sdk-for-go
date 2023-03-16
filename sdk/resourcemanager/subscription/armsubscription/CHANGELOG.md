# Release History

## 2.0.0-beta.1 (2023-03-16)
### Breaking Changes

- Type alias `SpendingLimit` has been removed
- Type alias `SubscriptionState` has been removed
- Function `NewSubscriptionsClient` has been removed
- Function `*SubscriptionsClient.Get` has been removed
- Function `*SubscriptionsClient.NewListLocationsPager` has been removed
- Function `*SubscriptionsClient.NewListPager` has been removed
- Function `NewTenantsClient` has been removed
- Function `*TenantsClient.NewListPager` has been removed
- Struct `ListResult` has been removed
- Struct `Location` has been removed
- Struct `LocationListResult` has been removed
- Struct `Policies` has been removed
- Struct `Subscription` has been removed
- Struct `SubscriptionsClient` has been removed
- Struct `TenantIDDescription` has been removed
- Struct `TenantListResult` has been removed
- Struct `TenantsClient` has been removed

### Features Added

- New type alias `Provisioning` with values `ProvisioningAccepted`, `ProvisioningPending`, `ProvisioningSucceeded`
- New field `ProvisioningState` in struct `AcceptOwnershipStatusResponse`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/subscription/armsubscription` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).