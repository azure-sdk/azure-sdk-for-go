# Release History

## 2.0.0-beta.1 (2023-03-13)
### Breaking Changes

- Type of `EndpointProperties.CustomDomains` has been changed from `[]*CustomDomain` to `[]*DeepCreatedCustomDomain`
- Type of `MetricsResponse.Granularity` has been changed from `*MetricsResponseGranularity` to `*MetricsGranularity`
- Type of `MetricsResponseSeriesItem.Unit` has been changed from `*MetricsResponseSeriesItemUnit` to `*MetricsSeriesUnit`
- Type of `WafMetricsResponse.Granularity` has been changed from `*WafMetricsResponseGranularity` to `*WafMetricsGranularity`
- Type of `WafMetricsResponseSeriesItem.Unit` has been changed from `*WafMetricsResponseSeriesItemUnit` to `*WafMetricsSeriesUnit`
- Type alias `MetricsResponseGranularity` has been removed
- Type alias `MetricsResponseSeriesItemUnit` has been removed
- Type alias `WafMetricsResponseGranularity` has been removed
- Type alias `WafMetricsResponseSeriesItemUnit` has been removed
- Function `NewValidateClient` has been removed
- Function `*ValidateClient.Secret` has been removed
- Operation `*CustomDomainsClient.DisableCustomHTTPS` has been changed to LRO, use `*CustomDomainsClient.BeginDisableCustomHTTPS` instead.
- Operation `*CustomDomainsClient.EnableCustomHTTPS` has been changed to LRO, use `*CustomDomainsClient.BeginEnableCustomHTTPS` instead.
- Struct `ValidateClient` has been removed

### Features Added

- New value `ProfileResourceStateAbortingMigration`, `ProfileResourceStateCommittingMigration`, `ProfileResourceStateMigrated`, `ProfileResourceStateMigrating`, `ProfileResourceStatePendingMigrationCommit` added to type alias `ProfileResourceState`
- New type alias `CanMigrateDefaultSKU` with values `CanMigrateDefaultSKUPremiumAzureFrontDoor`, `CanMigrateDefaultSKUStandardAzureFrontDoor`
- New type alias `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New type alias `MetricsGranularity` with values `MetricsGranularityP1D`, `MetricsGranularityPT1H`, `MetricsGranularityPT5M`
- New type alias `MetricsSeriesUnit` with values `MetricsSeriesUnitBitsPerSecond`, `MetricsSeriesUnitBytes`, `MetricsSeriesUnitCount`, `MetricsSeriesUnitMilliSeconds`
- New type alias `WafMetricsGranularity` with values `WafMetricsGranularityP1D`, `WafMetricsGranularityPT1H`, `WafMetricsGranularityPT5M`
- New type alias `WafMetricsSeriesUnit` with values `WafMetricsSeriesUnitCount`
- New function `*AFDProfilesClient.CheckEndpointNameAvailability(context.Context, string, string, CheckEndpointNameAvailabilityInput, *AFDProfilesClientCheckEndpointNameAvailabilityOptions) (AFDProfilesClientCheckEndpointNameAvailabilityResponse, error)`
- New function `*AFDProfilesClient.BeginUpgrade(context.Context, string, string, ProfileUpgradeParameters, *AFDProfilesClientBeginUpgradeOptions) (*runtime.Poller[AFDProfilesClientUpgradeResponse], error)`
- New function `*AFDProfilesClient.ValidateSecret(context.Context, string, string, ValidateSecretInput, *AFDProfilesClientValidateSecretOptions) (AFDProfilesClientValidateSecretResponse, error)`
- New function `*ProfilesClient.CanMigrate(context.Context, string, CanMigrateParameters, *ProfilesClientCanMigrateOptions) (ProfilesClientCanMigrateResponse, error)`
- New function `*ProfilesClient.BeginMigrate(context.Context, string, MigrationParameters, *ProfilesClientBeginMigrateOptions) (*runtime.Poller[ProfilesClientMigrateResponse], error)`
- New function `*ProfilesClient.BeginMigrationCommit(context.Context, string, string, *ProfilesClientBeginMigrationCommitOptions) (*runtime.Poller[ProfilesClientMigrationCommitResponse], error)`
- New struct `CanMigrateParameters`
- New struct `CanMigrateResult`
- New struct `DeepCreatedCustomDomain`
- New struct `DeepCreatedCustomDomainProperties`
- New struct `ManagedServiceIdentity`
- New struct `MigrateResult`
- New struct `MigrationErrorType`
- New struct `MigrationParameters`
- New struct `MigrationWebApplicationFirewallMapping`
- New struct `ProfileChangeSKUWafMapping`
- New struct `ProfileUpgradeParameters`
- New struct `UserAssignedIdentity`
- New field `ExtendedProperties` in struct `AFDDomainProperties`
- New field `Identity` in struct `Profile`
- New field `ExtendedProperties` in struct `ProfileProperties`
- New field `Identity` in struct `ProfileUpdateParameters`
- New field `ExtendedProperties` in struct `WebApplicationFirewallPolicyProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).