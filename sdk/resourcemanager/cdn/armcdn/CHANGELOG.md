# Release History

## 2.0.0-beta.1 (2022-11-28)
### Breaking Changes

- Type of `EndpointProperties.CustomDomains` has been changed from `[]*CustomDomain` to `[]*DeepCreatedCustomDomain`
- Type of `MetricsResponse.Granularity` has been changed from `*MetricsResponseGranularity` to `*MetricsGranularity`
- Type of `MetricsResponseSeriesItem.Unit` has been changed from `*MetricsResponseSeriesItemUnit` to `*MetricsSeriesUnit`
- Type of `WafMetricsResponse.Granularity` has been changed from `*WafMetricsResponseGranularity` to `*WafMetricsGranularity`
- Type of `WafMetricsResponseSeriesItem.Unit` has been changed from `*WafMetricsResponseSeriesItemUnit` to `*WafMetricsSeriesUnit`
- Const `MetricsResponseGranularityP1D` has been removed
- Const `MetricsResponseGranularityPT1H` has been removed
- Const `MetricsResponseGranularityPT5M` has been removed
- Const `MetricsResponseSeriesItemUnitBitsPerSecond` has been removed
- Const `MetricsResponseSeriesItemUnitBytes` has been removed
- Const `MetricsResponseSeriesItemUnitCount` has been removed
- Const `MetricsResponseSeriesItemUnitMilliSeconds` has been removed
- Const `WafMetricsResponseGranularityP1D` has been removed
- Const `WafMetricsResponseGranularityPT1H` has been removed
- Const `WafMetricsResponseGranularityPT5M` has been removed
- Const `WafMetricsResponseSeriesItemUnitCount` has been removed
- Type alias `MetricsResponseGranularity` has been removed
- Type alias `MetricsResponseSeriesItemUnit` has been removed
- Type alias `WafMetricsResponseGranularity` has been removed
- Type alias `WafMetricsResponseSeriesItemUnit` has been removed
- Function `*CustomDomainsClient.DisableCustomHTTPS` has been removed
- Function `*CustomDomainsClient.EnableCustomHTTPS` has been removed
- Function `PossibleMetricsResponseGranularityValues` has been removed
- Function `PossibleMetricsResponseSeriesItemUnitValues` has been removed
- Function `PossibleWafMetricsResponseGranularityValues` has been removed
- Function `PossibleWafMetricsResponseSeriesItemUnitValues` has been removed
- Function `NewValidateClient` has been removed
- Function `*ValidateClient.Secret` has been removed
- Struct `CustomDomainsClientDisableCustomHTTPSOptions` has been removed
- Struct `CustomDomainsClientEnableCustomHTTPSOptions` has been removed
- Struct `ValidateClient` has been removed
- Struct `ValidateClientSecretOptions` has been removed
- Struct `ValidateClientSecretResponse` has been removed

### Features Added

- New const `CanMigrateDefaultSKUPremiumAzureFrontDoor`
- New const `CanMigrateDefaultSKUStandardAzureFrontDoor`
- New const `ManagedServiceIdentityTypeNone`
- New const `ManagedServiceIdentityTypeSystemAssigned`
- New const `ManagedServiceIdentityTypeSystemAssignedUserAssigned`
- New const `ManagedServiceIdentityTypeUserAssigned`
- New const `MetricsGranularityP1D`
- New const `MetricsGranularityPT1H`
- New const `MetricsGranularityPT5M`
- New const `MetricsSeriesUnitBitsPerSecond`
- New const `MetricsSeriesUnitBytes`
- New const `MetricsSeriesUnitCount`
- New const `MetricsSeriesUnitMilliSeconds`
- New const `ProfileResourceStateAbortingMigration`
- New const `ProfileResourceStateCommittingMigration`
- New const `ProfileResourceStateMigrated`
- New const `ProfileResourceStateMigrating`
- New const `ProfileResourceStatePendingMigrationCommit`
- New const `WafMetricsGranularityP1D`
- New const `WafMetricsGranularityPT1H`
- New const `WafMetricsGranularityPT5M`
- New const `WafMetricsSeriesUnitCount`
- New type alias `CanMigrateDefaultSKU`
- New type alias `ManagedServiceIdentityType`
- New type alias `MetricsGranularity`
- New type alias `MetricsSeriesUnit`
- New type alias `WafMetricsGranularity`
- New type alias `WafMetricsSeriesUnit`
- New function `*AFDProfilesClient.CheckEndpointNameAvailability(context.Context, string, string, CheckEndpointNameAvailabilityInput, *AFDProfilesClientCheckEndpointNameAvailabilityOptions) (AFDProfilesClientCheckEndpointNameAvailabilityResponse, error)`
- New function `*AFDProfilesClient.BeginUpgrade(context.Context, string, string, ProfileUpgradeParameters, *AFDProfilesClientBeginUpgradeOptions) (*runtime.Poller[AFDProfilesClientUpgradeResponse], error)`
- New function `*AFDProfilesClient.ValidateSecret(context.Context, string, string, ValidateSecretInput, *AFDProfilesClientValidateSecretOptions) (AFDProfilesClientValidateSecretResponse, error)`
- New function `*CustomDomainsClient.BeginDisableCustomHTTPS(context.Context, string, string, string, string, *CustomDomainsClientBeginDisableCustomHTTPSOptions) (*runtime.Poller[CustomDomainsClientDisableCustomHTTPSResponse], error)`
- New function `*CustomDomainsClient.BeginEnableCustomHTTPS(context.Context, string, string, string, string, *CustomDomainsClientBeginEnableCustomHTTPSOptions) (*runtime.Poller[CustomDomainsClientEnableCustomHTTPSResponse], error)`
- New function `PossibleCanMigrateDefaultSKUValues() []CanMigrateDefaultSKU`
- New function `PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType`
- New function `PossibleMetricsGranularityValues() []MetricsGranularity`
- New function `PossibleMetricsSeriesUnitValues() []MetricsSeriesUnit`
- New function `PossibleWafMetricsGranularityValues() []WafMetricsGranularity`
- New function `PossibleWafMetricsSeriesUnitValues() []WafMetricsSeriesUnit`
- New function `*ProfilesClient.CanMigrate(context.Context, string, CanMigrateParameters, *ProfilesClientCanMigrateOptions) (ProfilesClientCanMigrateResponse, error)`
- New function `*ProfilesClient.BeginMigrate(context.Context, string, MigrationParameters, *ProfilesClientBeginMigrateOptions) (*runtime.Poller[ProfilesClientMigrateResponse], error)`
- New function `*ProfilesClient.BeginMigrationCommit(context.Context, string, string, *ProfilesClientBeginMigrationCommitOptions) (*runtime.Poller[ProfilesClientMigrationCommitResponse], error)`
- New struct `AFDProfilesClientBeginUpgradeOptions`
- New struct `AFDProfilesClientCheckEndpointNameAvailabilityOptions`
- New struct `AFDProfilesClientCheckEndpointNameAvailabilityResponse`
- New struct `AFDProfilesClientUpgradeResponse`
- New struct `AFDProfilesClientValidateSecretOptions`
- New struct `AFDProfilesClientValidateSecretResponse`
- New struct `CanMigrateParameters`
- New struct `CanMigrateResult`
- New struct `CustomDomainsClientBeginDisableCustomHTTPSOptions`
- New struct `CustomDomainsClientBeginEnableCustomHTTPSOptions`
- New struct `DeepCreatedCustomDomain`
- New struct `DeepCreatedCustomDomainProperties`
- New struct `ManagedServiceIdentity`
- New struct `MigrateResult`
- New struct `MigrationErrorType`
- New struct `MigrationParameters`
- New struct `MigrationWebApplicationFirewallMapping`
- New struct `ProfileChangeSKUWafMapping`
- New struct `ProfileUpgradeParameters`
- New struct `ProfilesClientBeginMigrateOptions`
- New struct `ProfilesClientBeginMigrationCommitOptions`
- New struct `ProfilesClientCanMigrateOptions`
- New struct `ProfilesClientCanMigrateResponse`
- New struct `ProfilesClientMigrateResponse`
- New struct `ProfilesClientMigrationCommitResponse`
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