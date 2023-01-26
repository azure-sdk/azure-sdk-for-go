# Release History

## 2.0.0 (2023-01-26)
### Breaking Changes

- Type alias `AcceptOwnership` has been removed
- Type alias `CreatedByType` has been removed
- Type alias `ProvisioningState` has been removed
- Type alias `Workload` has been removed
- Function `NewAliasClient` has been removed
- Function `*AliasClient.BeginCreate` has been removed
- Function `*AliasClient.Delete` has been removed
- Function `*AliasClient.Get` has been removed
- Function `*AliasClient.List` has been removed
- Function `NewBillingAccountClient` has been removed
- Function `*BillingAccountClient.GetPolicy` has been removed
- Function `*Client.BeginAcceptOwnership` has been removed
- Function `*Client.AcceptOwnershipStatus` has been removed
- Function `*Client.Rename` has been removed
- Function `NewOperationsClient` has been removed
- Function `*OperationsClient.NewListPager` has been removed
- Function `NewPolicyClient` has been removed
- Function `*PolicyClient.AddUpdatePolicyForTenant` has been removed
- Function `*PolicyClient.GetPolicyForTenant` has been removed
- Function `*PolicyClient.NewListPolicyForTenantPager` has been removed
- Function `timeRFC3339.MarshalText` has been removed
- Function `*timeRFC3339.Parse` has been removed
- Function `*timeRFC3339.UnmarshalText` has been removed
- Operation `*Client.Cancel` has been changed to LRO, use `*Client.BeginCancel` instead.
- Operation `*Client.Enable` has been changed to LRO, use `*Client.BeginEnable` instead.
- Struct `AcceptOwnershipRequest` has been removed
- Struct `AcceptOwnershipRequestProperties` has been removed
- Struct `AcceptOwnershipStatusResponse` has been removed
- Struct `AliasClient` has been removed
- Struct `AliasClientCreateResponse` has been removed
- Struct `AliasListResult` has been removed
- Struct `AliasResponse` has been removed
- Struct `AliasResponseProperties` has been removed
- Struct `BillingAccountClient` has been removed
- Struct `BillingAccountPoliciesResponse` has been removed
- Struct `BillingAccountPoliciesResponseProperties` has been removed
- Struct `CanceledSubscriptionID` has been removed
- Struct `ClientAcceptOwnershipResponse` has been removed
- Struct `EnabledSubscriptionID` has been removed
- Struct `GetTenantPolicyListResponse` has been removed
- Struct `GetTenantPolicyResponse` has been removed
- Struct `Name` has been removed
- Struct `Operation` has been removed
- Struct `OperationDisplay` has been removed
- Struct `OperationListResult` has been removed
- Struct `OperationsClient` has been removed
- Struct `OperationsClientListResponse` has been removed
- Struct `PolicyClient` has been removed
- Struct `PolicyClientListPolicyForTenantResponse` has been removed
- Struct `PutAliasRequest` has been removed
- Struct `PutAliasRequestAdditionalProperties` has been removed
- Struct `PutAliasRequestProperties` has been removed
- Struct `PutTenantPolicyRequestProperties` has been removed
- Struct `RenamedSubscriptionID` has been removed
- Struct `ServiceTenantResponse` has been removed
- Struct `SystemData` has been removed
- Struct `TenantPolicy` has been removed

### Features Added

- New struct `ClientBeginOperationResultsOptions`
- New struct `ClientOperationResultsResponse`
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