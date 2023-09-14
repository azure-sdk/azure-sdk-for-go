# Release History

## 1.6.0-beta.1 (2023-09-14)
### Features Added

- New enum type `AllowedContentLevel` with values `AllowedContentLevelHigh`, `AllowedContentLevelLow`, `AllowedContentLevelMedium`
- New enum type `ByPassSelection` with values `ByPassSelectionAzureServices`, `ByPassSelectionNone`
- New enum type `RaiContentFilterType` with values `RaiContentFilterTypeMultiLevel`, `RaiContentFilterTypeSwitch`
- New enum type `RaiPolicyContentSource` with values `RaiPolicyContentSourceCompletion`, `RaiPolicyContentSourcePrompt`
- New enum type `RaiPolicyMode` with values `RaiPolicyModeBlocking`, `RaiPolicyModeDefault`, `RaiPolicyModeDeferred`
- New enum type `RaiPolicyType` with values `RaiPolicyTypeSystemManaged`, `RaiPolicyTypeUserManaged`
- New function `*ClientFactory.NewRaiBlocklistClient() *RaiBlocklistClient`
- New function `*ClientFactory.NewRaiBlocklistItemClient() *RaiBlocklistItemClient`
- New function `*ClientFactory.NewRaiBlocklistItemsClient() *RaiBlocklistItemsClient`
- New function `*ClientFactory.NewRaiBlocklistsClient() *RaiBlocklistsClient`
- New function `*ClientFactory.NewRaiContentFiltersClient() *RaiContentFiltersClient`
- New function `*ClientFactory.NewRaiPoliciesClient() *RaiPoliciesClient`
- New function `NewRaiBlocklistClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RaiBlocklistClient, error)`
- New function `*RaiBlocklistClient.CreateOrUpdate(context.Context, string, string, string, RaiBlocklist, *RaiBlocklistClientCreateOrUpdateOptions) (RaiBlocklistClientCreateOrUpdateResponse, error)`
- New function `*RaiBlocklistClient.Get(context.Context, string, string, string, *RaiBlocklistClientGetOptions) (RaiBlocklistClientGetResponse, error)`
- New function `NewRaiBlocklistItemClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RaiBlocklistItemClient, error)`
- New function `*RaiBlocklistItemClient.CreateOrUpdate(context.Context, string, string, string, string, RaiBlocklistItem, *RaiBlocklistItemClientCreateOrUpdateOptions) (RaiBlocklistItemClientCreateOrUpdateResponse, error)`
- New function `*RaiBlocklistItemClient.BeginDelete(context.Context, string, string, string, string, *RaiBlocklistItemClientBeginDeleteOptions) (*runtime.Poller[RaiBlocklistItemClientDeleteResponse], error)`
- New function `*RaiBlocklistItemClient.Get(context.Context, string, string, string, string, *RaiBlocklistItemClientGetOptions) (RaiBlocklistItemClientGetResponse, error)`
- New function `NewRaiBlocklistItemsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RaiBlocklistItemsClient, error)`
- New function `*RaiBlocklistItemsClient.NewListPager(string, string, string, *RaiBlocklistItemsClientListOptions) *runtime.Pager[RaiBlocklistItemsClientListResponse]`
- New function `NewRaiBlocklistsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RaiBlocklistsClient, error)`
- New function `*RaiBlocklistsClient.BeginDelete(context.Context, string, string, string, *RaiBlocklistsClientBeginDeleteOptions) (*runtime.Poller[RaiBlocklistsClientDeleteResponse], error)`
- New function `*RaiBlocklistsClient.NewListPager(string, string, *RaiBlocklistsClientListOptions) *runtime.Pager[RaiBlocklistsClientListResponse]`
- New function `NewRaiContentFiltersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RaiContentFiltersClient, error)`
- New function `*RaiContentFiltersClient.List(context.Context, string, *RaiContentFiltersClientListOptions) (RaiContentFiltersClientListResponse, error)`
- New function `NewRaiPoliciesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*RaiPoliciesClient, error)`
- New function `*RaiPoliciesClient.CreateOrUpdate(context.Context, string, string, string, RaiPolicy, *RaiPoliciesClientCreateOrUpdateOptions) (RaiPoliciesClientCreateOrUpdateResponse, error)`
- New function `*RaiPoliciesClient.BeginDelete(context.Context, string, string, string, *RaiPoliciesClientBeginDeleteOptions) (*runtime.Poller[RaiPoliciesClientDeleteResponse], error)`
- New function `*RaiPoliciesClient.Get(context.Context, string, string, string, *RaiPoliciesClientGetOptions) (RaiPoliciesClientGetResponse, error)`
- New function `*RaiPoliciesClient.NewListPager(string, string, *RaiPoliciesClientListOptions) *runtime.Pager[RaiPoliciesClientListResponse]`
- New struct `RaiBlockListItemsResult`
- New struct `RaiBlockListResult`
- New struct `RaiBlocklist`
- New struct `RaiBlocklistConfig`
- New struct `RaiBlocklistItem`
- New struct `RaiBlocklistItemProperties`
- New struct `RaiBlocklistProperties`
- New struct `RaiContentFilter`
- New struct `RaiContentFilterListResult`
- New struct `RaiPolicy`
- New struct `RaiPolicyContentFilter`
- New struct `RaiPolicyListResult`
- New struct `RaiPolicyProperties`
- New field `Bypass` in struct `NetworkRuleSet`


## 1.5.0 (2023-07-28)
### Features Added

- New value `DeploymentProvisioningStateCanceled`, `DeploymentProvisioningStateDisabled` added to enum type `DeploymentProvisioningState`
- New value `HostingModelProvisionedWeb` added to enum type `HostingModel`
- New enum type `AbusePenaltyAction` with values `AbusePenaltyActionBlock`, `AbusePenaltyActionThrottle`
- New enum type `DeploymentModelVersionUpgradeOption` with values `DeploymentModelVersionUpgradeOptionNoAutoUpgrade`, `DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired`, `DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable`
- New function `*ClientFactory.NewModelsClient() *ModelsClient`
- New function `*ClientFactory.NewUsagesClient() *UsagesClient`
- New function `NewModelsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ModelsClient, error)`
- New function `*ModelsClient.NewListPager(string, *ModelsClientListOptions) *runtime.Pager[ModelsClientListResponse]`
- New function `NewUsagesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UsagesClient, error)`
- New function `*UsagesClient.NewListPager(string, *UsagesClientListOptions) *runtime.Pager[UsagesClientListResponse]`
- New struct `AbusePenalty`
- New struct `CapacityConfig`
- New struct `Model`
- New struct `ModelListResult`
- New struct `ModelSKU`
- New field `IsDefaultVersion`, `SKUs`, `Source` in struct `AccountModel`
- New field `AbusePenalty` in struct `AccountProperties`
- New field `ProvisioningIssues` in struct `CommitmentPlanProperties`
- New field `SKU` in struct `Deployment`
- New field `Source` in struct `DeploymentModel`
- New field `RateLimits`, `VersionUpgradeOption` in struct `DeploymentProperties`
- New field `NextLink` in struct `UsageListResult`


## 1.4.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 1.4.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.3.0 (2023-02-24)
### Features Added

- New type alias `CommitmentPlanProvisioningState` with values `CommitmentPlanProvisioningStateAccepted`, `CommitmentPlanProvisioningStateCanceled`, `CommitmentPlanProvisioningStateCreating`, `CommitmentPlanProvisioningStateDeleting`, `CommitmentPlanProvisioningStateFailed`, `CommitmentPlanProvisioningStateMoving`, `CommitmentPlanProvisioningStateSucceeded`
- New type alias `ModelLifecycleStatus` with values `ModelLifecycleStatusGenerallyAvailable`, `ModelLifecycleStatusPreview`
- New type alias `RoutingMethods` with values `RoutingMethodsPerformance`, `RoutingMethodsPriority`, `RoutingMethodsWeighted`
- New function `*CommitmentPlansClient.BeginCreateOrUpdateAssociation(context.Context, string, string, string, CommitmentPlanAccountAssociation, *CommitmentPlansClientBeginCreateOrUpdateAssociationOptions) (*runtime.Poller[CommitmentPlansClientCreateOrUpdateAssociationResponse], error)`
- New function `*CommitmentPlansClient.BeginCreateOrUpdatePlan(context.Context, string, string, CommitmentPlan, *CommitmentPlansClientBeginCreateOrUpdatePlanOptions) (*runtime.Poller[CommitmentPlansClientCreateOrUpdatePlanResponse], error)`
- New function `*CommitmentPlansClient.BeginDeleteAssociation(context.Context, string, string, string, *CommitmentPlansClientBeginDeleteAssociationOptions) (*runtime.Poller[CommitmentPlansClientDeleteAssociationResponse], error)`
- New function `*CommitmentPlansClient.BeginDeletePlan(context.Context, string, string, *CommitmentPlansClientBeginDeletePlanOptions) (*runtime.Poller[CommitmentPlansClientDeletePlanResponse], error)`
- New function `*CommitmentPlansClient.GetAssociation(context.Context, string, string, string, *CommitmentPlansClientGetAssociationOptions) (CommitmentPlansClientGetAssociationResponse, error)`
- New function `*CommitmentPlansClient.GetPlan(context.Context, string, string, *CommitmentPlansClientGetPlanOptions) (CommitmentPlansClientGetPlanResponse, error)`
- New function `*CommitmentPlansClient.NewListAssociationsPager(string, string, *CommitmentPlansClientListAssociationsOptions) *runtime.Pager[CommitmentPlansClientListAssociationsResponse]`
- New function `*CommitmentPlansClient.NewListPlansByResourceGroupPager(string, *CommitmentPlansClientListPlansByResourceGroupOptions) *runtime.Pager[CommitmentPlansClientListPlansByResourceGroupResponse]`
- New function `*CommitmentPlansClient.NewListPlansBySubscriptionPager(*CommitmentPlansClientListPlansBySubscriptionOptions) *runtime.Pager[CommitmentPlansClientListPlansBySubscriptionResponse]`
- New function `*CommitmentPlansClient.BeginUpdatePlan(context.Context, string, string, PatchResourceTagsAndSKU, *CommitmentPlansClientBeginUpdatePlanOptions) (*runtime.Poller[CommitmentPlansClientUpdatePlanResponse], error)`
- New struct `CommitmentPlanAccountAssociation`
- New struct `CommitmentPlanAccountAssociationListResult`
- New struct `CommitmentPlanAccountAssociationProperties`
- New struct `CommitmentPlanAssociation`
- New struct `CommitmentPlansClientCreateOrUpdateAssociationResponse`
- New struct `CommitmentPlansClientCreateOrUpdatePlanResponse`
- New struct `CommitmentPlansClientDeleteAssociationResponse`
- New struct `CommitmentPlansClientDeletePlanResponse`
- New struct `CommitmentPlansClientListAssociationsResponse`
- New struct `CommitmentPlansClientListPlansByResourceGroupResponse`
- New struct `CommitmentPlansClientListPlansBySubscriptionResponse`
- New struct `CommitmentPlansClientUpdatePlanResponse`
- New struct `MultiRegionSettings`
- New struct `PatchResourceTags`
- New struct `PatchResourceTagsAndSKU`
- New struct `RegionSetting`
- New field `FinetuneCapabilities` in struct `AccountModel`
- New field `LifecycleStatus` in struct `AccountModel`
- New field `CommitmentPlanAssociations` in struct `AccountProperties`
- New field `Locations` in struct `AccountProperties`
- New field `Kind` in struct `CommitmentPlan`
- New field `Location` in struct `CommitmentPlan`
- New field `SKU` in struct `CommitmentPlan`
- New field `Tags` in struct `CommitmentPlan`
- New field `CommitmentPlanGUID` in struct `CommitmentPlanProperties`
- New field `ProvisioningState` in struct `CommitmentPlanProperties`


## 1.2.0 (2022-10-20)
### Features Added

- New field `CallRateLimit` in struct `DeploymentProperties`
- New field `Capabilities` in struct `DeploymentProperties`
- New field `RaiPolicyName` in struct `DeploymentProperties`
- New field `CallRateLimit` in struct `AccountModel`
- New field `CallRateLimit` in struct `DeploymentModel`


## 1.1.0 (2022-06-09)
### Features Added

- New const `DeploymentScaleTypeStandard`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/cognitiveservices/armcognitiveservices` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).