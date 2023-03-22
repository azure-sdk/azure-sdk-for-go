# Release History

## 2.0.0-beta.1 (2023-03-22)
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

- New value `DeliveryRuleActionCustomErrorPageURL`, `DeliveryRuleActionOverrideResponseStatusCode` added to enum type `DeliveryRuleAction`
- New value `MatchVariableResponseStatusCode` added to enum type `MatchVariable`
- New value `ProfileResourceStateAbortingMigration`, `ProfileResourceStateCommittingMigration`, `ProfileResourceStateMigrated`, `ProfileResourceStateMigrating`, `ProfileResourceStatePendingMigrationCommit` added to enum type `ProfileResourceState`
- New enum type `CanMigrateDefaultSKU` with values `CanMigrateDefaultSKUPremiumAzureFrontDoor`, `CanMigrateDefaultSKUStandardAzureFrontDoor`
- New enum type `DeliveryRuleCustomErrorPageActionParameters` with values `DeliveryRuleCustomErrorPageActionParametersDeliveryRuleCustomErrorPageActionParameters`
- New enum type `DeliveryRuleOverrideResponseStatusCodeActionParameters` with values `DeliveryRuleOverrideResponseStatusCodeActionParametersDeliveryRuleOverrideResponseStatusCodeActionParameters`
- New enum type `DeliveryRuleResponseStatusCodeConditionParameters` with values `DeliveryRuleResponseStatusCodeConditionParametersDeliveryRuleResponseStatusCodeConditionParameters`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `MetricsGranularity` with values `MetricsGranularityP1D`, `MetricsGranularityPT1H`, `MetricsGranularityPT5M`
- New enum type `MetricsSeriesUnit` with values `MetricsSeriesUnitBitsPerSecond`, `MetricsSeriesUnitBytes`, `MetricsSeriesUnitCount`, `MetricsSeriesUnitMilliSeconds`
- New enum type `RuleIsNegativeCachingEnabled` with values `RuleIsNegativeCachingEnabledDisabled`, `RuleIsNegativeCachingEnabledEnabled`
- New enum type `WafMetricsGranularity` with values `WafMetricsGranularityP1D`, `WafMetricsGranularityPT1H`, `WafMetricsGranularityPT5M`
- New enum type `WafMetricsSeriesUnit` with values `WafMetricsSeriesUnitCount`
- New function `*AFDProfilesClient.CheckEndpointNameAvailability(context.Context, string, string, CheckEndpointNameAvailabilityInput, *AFDProfilesClientCheckEndpointNameAvailabilityOptions) (AFDProfilesClientCheckEndpointNameAvailabilityResponse, error)`
- New function `*AFDProfilesClient.BeginUpgrade(context.Context, string, string, ProfileUpgradeParameters, *AFDProfilesClientBeginUpgradeOptions) (*runtime.Poller[AFDProfilesClientUpgradeResponse], error)`
- New function `*AFDProfilesClient.ValidateSecret(context.Context, string, string, ValidateSecretInput, *AFDProfilesClientValidateSecretOptions) (AFDProfilesClientValidateSecretResponse, error)`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAFDCustomDomainsClient() *AFDCustomDomainsClient`
- New function `*ClientFactory.NewAFDEndpointsClient() *AFDEndpointsClient`
- New function `*ClientFactory.NewAFDOriginGroupsClient() *AFDOriginGroupsClient`
- New function `*ClientFactory.NewAFDOriginsClient() *AFDOriginsClient`
- New function `*ClientFactory.NewAFDProfilesClient() *AFDProfilesClient`
- New function `*ClientFactory.NewCustomDomainsClient() *CustomDomainsClient`
- New function `*ClientFactory.NewEdgeNodesClient() *EdgeNodesClient`
- New function `*ClientFactory.NewEndpointsClient() *EndpointsClient`
- New function `*ClientFactory.NewLogAnalyticsClient() *LogAnalyticsClient`
- New function `*ClientFactory.NewManagedRuleSetsClient() *ManagedRuleSetsClient`
- New function `*ClientFactory.NewManagementClient() *ManagementClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewOriginGroupsClient() *OriginGroupsClient`
- New function `*ClientFactory.NewOriginsClient() *OriginsClient`
- New function `*ClientFactory.NewPoliciesClient() *PoliciesClient`
- New function `*ClientFactory.NewProfilesClient() *ProfilesClient`
- New function `*ClientFactory.NewResourceUsageClient() *ResourceUsageClient`
- New function `*ClientFactory.NewRoutesClient() *RoutesClient`
- New function `*ClientFactory.NewRuleSetsClient() *RuleSetsClient`
- New function `*ClientFactory.NewRulesClient() *RulesClient`
- New function `*ClientFactory.NewSecretsClient() *SecretsClient`
- New function `*ClientFactory.NewSecurityPoliciesClient() *SecurityPoliciesClient`
- New function `*DeliveryRuleCustomErrorPageAction.GetDeliveryRuleActionAutoGenerated() *DeliveryRuleActionAutoGenerated`
- New function `*DeliveryRuleOverrideResponseStatusCodeAction.GetDeliveryRuleActionAutoGenerated() *DeliveryRuleActionAutoGenerated`
- New function `*DeliveryRuleResponseStatusCodeCondition.GetDeliveryRuleCondition() *DeliveryRuleCondition`
- New function `*ProfilesClient.BeginCanMigrate(context.Context, string, CanMigrateParameters, *ProfilesClientBeginCanMigrateOptions) (*runtime.Poller[ProfilesClientCanMigrateResponse], error)`
- New function `*ProfilesClient.BeginMigrate(context.Context, string, MigrationParameters, *ProfilesClientBeginMigrateOptions) (*runtime.Poller[ProfilesClientMigrateResponse], error)`
- New function `*ProfilesClient.BeginMigrationCommit(context.Context, string, string, *ProfilesClientBeginMigrationCommitOptions) (*runtime.Poller[ProfilesClientMigrationCommitResponse], error)`
- New struct `CanMigrateParameters`
- New struct `CanMigrateResult`
- New struct `ClientFactory`
- New struct `CustomErrorPageActionParameters`
- New struct `DeepCreatedCustomDomain`
- New struct `DeepCreatedCustomDomainProperties`
- New struct `DeliveryRuleCustomErrorPageAction`
- New struct `DeliveryRuleOverrideResponseStatusCodeAction`
- New struct `DeliveryRuleResponseStatusCodeCondition`
- New struct `ManagedServiceIdentity`
- New struct `MigrateResult`
- New struct `MigrationErrorType`
- New struct `MigrationParameters`
- New struct `MigrationWebApplicationFirewallMapping`
- New struct `OverrideResponseStatusCodeActionParameters`
- New struct `ProfileChangeSKUWafMapping`
- New struct `ProfileUpgradeParameters`
- New struct `ResponseStatusCodeMatchConditionParameters`
- New struct `UserAssignedIdentity`
- New field `ExtendedProperties` in struct `AFDDomainProperties`
- New field `IsNegativeCachingEnabled` in struct `CacheConfiguration`
- New field `CapacityConsciousThreshold` in struct `LoadBalancingSettingsParameters`
- New field `Identity` in struct `Profile`
- New field `ExtendedProperties` in struct `ProfileProperties`
- New field `Identity` in struct `ProfileUpdateParameters`
- New field `ExtendedProperties` in struct `WebApplicationFirewallPolicyProperties`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cdn/armcdn` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).