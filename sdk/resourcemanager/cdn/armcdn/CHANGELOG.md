# Release History

## 2.0.0-beta.1 (2023-04-25)
### Breaking Changes

- Type of `EndpointProperties.CustomDomains` has been changed from `[]*CustomDomain` to `[]*DeepCreatedCustomDomain`
- Type of `MetricsResponse.Granularity` has been changed from `*MetricsResponseGranularity` to `*MetricsGranularity`
- Type of `MetricsResponseSeriesItem.Unit` has been changed from `*MetricsResponseSeriesItemUnit` to `*MetricsSeriesUnit`
- Type of `WafMetricsResponse.Granularity` has been changed from `*WafMetricsResponseGranularity` to `*WafMetricsGranularity`
- Type of `WafMetricsResponseSeriesItem.Unit` has been changed from `*WafMetricsResponseSeriesItemUnit` to `*WafMetricsSeriesUnit`
- Enum `MetricsResponseGranularity` has been removed
- Enum `MetricsResponseSeriesItemUnit` has been removed
- Enum `WafMetricsResponseGranularity` has been removed
- Enum `WafMetricsResponseSeriesItemUnit` has been removed
- Function `*ClientFactory.NewValidateClient` has been removed
- Function `NewValidateClient` has been removed
- Function `*ValidateClient.Secret` has been removed
- Operation `*CustomDomainsClient.DisableCustomHTTPS` has been changed to LRO, use `*CustomDomainsClient.BeginDisableCustomHTTPS` instead.
- Operation `*CustomDomainsClient.EnableCustomHTTPS` has been changed to LRO, use `*CustomDomainsClient.BeginEnableCustomHTTPS` instead.

### Features Added

- New value `ProbeProtocolTCP`, `ProbeProtocolUDP` added to enum type `ProbeProtocol`
- New value `ProfileResourceStateAbortingMigration`, `ProfileResourceStateCommittingMigration`, `ProfileResourceStateMigrated`, `ProfileResourceStateMigrating`, `ProfileResourceStatePendingMigrationCommit` added to enum type `ProfileResourceState`
- New enum type `BackendProtocol` with values `BackendProtocolTCP`, `BackendProtocolUDP`
- New enum type `CanMigrateDefaultSKU` with values `CanMigrateDefaultSKUPremiumAzureFrontDoor`, `CanMigrateDefaultSKUStandardAzureFrontDoor`
- New enum type `ClientIPPreservationState` with values `ClientIPPreservationStateDisabled`, `ClientIPPreservationStateEnabled`
- New enum type `FrontendProtocol` with values `FrontendProtocolTCP`, `FrontendProtocolUDP`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `MetricsGranularity` with values `MetricsGranularityP1D`, `MetricsGranularityPT1H`, `MetricsGranularityPT5M`
- New enum type `MetricsSeriesUnit` with values `MetricsSeriesUnitBitsPerSecond`, `MetricsSeriesUnitBytes`, `MetricsSeriesUnitCount`, `MetricsSeriesUnitMilliSeconds`
- New enum type `RouteType` with values `RouteTypeL4Route`, `RouteTypeL7Route`
- New enum type `RoutingMethod` with values `RoutingMethodPortBased`, `RoutingMethodSNIBased`
- New enum type `SessionAffinityType` with values `SessionAffinityTypeNone`, `SessionAffinityTypeSourceIP`
- New enum type `WafMetricsGranularity` with values `WafMetricsGranularityP1D`, `WafMetricsGranularityPT1H`, `WafMetricsGranularityPT5M`
- New enum type `WafMetricsSeriesUnit` with values `WafMetricsSeriesUnitCount`
- New function `*AFDProfilesClient.CheckEndpointNameAvailability(context.Context, string, string, CheckEndpointNameAvailabilityInput, *AFDProfilesClientCheckEndpointNameAvailabilityOptions) (AFDProfilesClientCheckEndpointNameAvailabilityResponse, error)`
- New function `*AFDProfilesClient.BeginUpgrade(context.Context, string, string, ProfileUpgradeParameters, *AFDProfilesClientBeginUpgradeOptions) (*runtime.Poller[AFDProfilesClientUpgradeResponse], error)`
- New function `*AFDProfilesClient.ValidateSecret(context.Context, string, string, ValidateSecretInput, *AFDProfilesClientValidateSecretOptions) (AFDProfilesClientValidateSecretResponse, error)`
- New function `*ClientFactory.NewL4RoutesClient() *L4RoutesClient`
- New function `NewL4RoutesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*L4RoutesClient, error)`
- New function `*L4RoutesClient.BeginCreate(context.Context, string, string, string, string, L4Route, *L4RoutesClientBeginCreateOptions) (*runtime.Poller[L4RoutesClientCreateResponse], error)`
- New function `*L4RoutesClient.BeginDelete(context.Context, string, string, string, string, *L4RoutesClientBeginDeleteOptions) (*runtime.Poller[L4RoutesClientDeleteResponse], error)`
- New function `*L4RoutesClient.Get(context.Context, string, string, string, string, *L4RoutesClientGetOptions) (L4RoutesClientGetResponse, error)`
- New function `*L4RoutesClient.NewListByEndpointPager(string, string, string, *L4RoutesClientListByEndpointOptions) *runtime.Pager[L4RoutesClientListByEndpointResponse]`
- New function `*L4RoutesClient.BeginUpdate(context.Context, string, string, string, string, L4RouteUpdatePropertiesParameters, *L4RoutesClientBeginUpdateOptions) (*runtime.Poller[L4RoutesClientUpdateResponse], error)`
- New function `*ProfilesClient.CanMigrate(context.Context, string, CanMigrateParameters, *ProfilesClientCanMigrateOptions) (ProfilesClientCanMigrateResponse, error)`
- New function `*ProfilesClient.BeginMigrate(context.Context, string, MigrationParameters, *ProfilesClientBeginMigrateOptions) (*runtime.Poller[ProfilesClientMigrateResponse], error)`
- New function `*ProfilesClient.BeginMigrationCommit(context.Context, string, string, *ProfilesClientBeginMigrationCommitOptions) (*runtime.Poller[ProfilesClientMigrationCommitResponse], error)`
- New struct `CanMigrateParameters`
- New struct `CanMigrateResult`
- New struct `DeepCreatedCustomDomain`
- New struct `DeepCreatedCustomDomainProperties`
- New struct `L4Route`
- New struct `L4RouteBaseSettings`
- New struct `L4RouteListResult`
- New struct `L4RouteProperties`
- New struct `L4RouteUpdatePropertiesParameters`
- New struct `ManagedServiceIdentity`
- New struct `MigrateResult`
- New struct `MigrationErrorType`
- New struct `MigrationParameters`
- New struct `MigrationWebApplicationFirewallMapping`
- New struct `ProfileChangeSKUWafMapping`
- New struct `ProfileUpgradeParameters`
- New struct `UserAssignedIdentity`
- New field `ExtendedProperties` in struct `AFDDomainProperties`
- New field `DedicatedVIP` in struct `AFDEndpointProperties`
- New field `DedicatedVIP` in struct `AFDEndpointPropertiesUpdateParameters`
- New field `RouteType`, `SessionAffinityType` in struct `AFDOriginGroupProperties`
- New field `RouteType`, `SessionAffinityType` in struct `AFDOriginGroupUpdatePropertiesParameters`
- New field `ClientIPPreservationState`, `OriginPorts` in struct `AFDOriginProperties`
- New field `ClientIPPreservationState`, `OriginPorts` in struct `AFDOriginUpdatePropertiesParameters`
- New field `CertificateAuthority`, `ExpirationDate`, `SecretSource`, `Subject`, `SubjectAlternativeNames`, `Thumbprint` in struct `AzureFirstPartyManagedCertificateParameters`
- New field `ProbePort` in struct `HealthProbeParameters`
- New field `Identity` in struct `Profile`
- New field `ExtendedProperties` in struct `ProfileProperties`
- New field `Identity` in struct `ProfileUpdateParameters`
- New field `ExtendedProperties` in struct `WebApplicationFirewallPolicyProperties`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 1.1.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).