# Release History

## 1.1.0 (2023-02-13)
### Features Added

- New value `ArcSettingAggregateStateAccepted`, `ArcSettingAggregateStateDisableInProgress`, `ArcSettingAggregateStateProvisioning` added to type alias `ArcSettingAggregateState`
- New value `ExtensionAggregateStateAccepted`, `ExtensionAggregateStateProvisioning`, `ExtensionAggregateStateUpgradeFailedRollbackSucceeded` added to type alias `ExtensionAggregateState`
- New value `NodeArcStateAccepted`, `NodeArcStateDisableInProgress`, `NodeArcStateInProgress`, `NodeArcStatePartiallyConnected`, `NodeArcStatePartiallySucceeded`, `NodeArcStateProvisioning` added to type alias `NodeArcState`
- New value `NodeExtensionStateAccepted`, `NodeExtensionStateInProgress`, `NodeExtensionStatePartiallyConnected`, `NodeExtensionStatePartiallySucceeded`, `NodeExtensionStateProvisioning` added to type alias `NodeExtensionState`
- New value `ProvisioningStateConnected`, `ProvisioningStateCreating`, `ProvisioningStateDeleted`, `ProvisioningStateDeleting`, `ProvisioningStateDisableInProgress`, `ProvisioningStateDisconnected`, `ProvisioningStateError`, `ProvisioningStateInProgress`, `ProvisioningStateMoving`, `ProvisioningStateNotSpecified`, `ProvisioningStatePartiallyConnected`, `ProvisioningStatePartiallySucceeded`, `ProvisioningStateUpdating` added to type alias `ProvisioningState`
- New value `StatusFailed`, `StatusInProgress`, `StatusNotSpecified`, `StatusSucceeded` added to type alias `Status`
- New type alias `AvailabilityType` with values `AvailabilityTypeLocal`, `AvailabilityTypeNotify`, `AvailabilityTypeOnline`
- New type alias `ClusterNodeType` with values `ClusterNodeTypeFirstParty`, `ClusterNodeTypeThirdParty`
- New type alias `HealthState` with values `HealthStateError`, `HealthStateFailure`, `HealthStateInProgress`, `HealthStateSuccess`, `HealthStateUnknown`, `HealthStateWarning`
- New type alias `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New type alias `RebootRequirement` with values `RebootRequirementFalse`, `RebootRequirementTrue`, `RebootRequirementUnknown`
- New type alias `Severity` with values `SeverityCritical`, `SeverityHidden`, `SeverityInformational`, `SeverityWarning`
- New type alias `SoftwareAssuranceIntent` with values `SoftwareAssuranceIntentDisable`, `SoftwareAssuranceIntentEnable`
- New type alias `SoftwareAssuranceStatus` with values `SoftwareAssuranceStatusDisabled`, `SoftwareAssuranceStatusEnabled`
- New type alias `State` with values `StateDownloadFailed`, `StateDownloading`, `StateHasPrerequisite`, `StateHealthCheckFailed`, `StateHealthChecking`, `StateInstallationFailed`, `StateInstalled`, `StateInstalling`, `StateInvalid`, `StateNotApplicableBecauseAnotherUpdateIsInProgress`, `StateObsolete`, `StatePreparationFailed`, `StatePreparing`, `StateReady`, `StateReadyToInstall`, `StateRecalled`, `StateScanFailed`, `StateScanInProgress`
- New type alias `StatusLevelTypes` with values `StatusLevelTypesError`, `StatusLevelTypesInfo`, `StatusLevelTypesWarning`
- New type alias `UpdateRunPropertiesState` with values `UpdateRunPropertiesStateFailed`, `UpdateRunPropertiesStateInProgress`, `UpdateRunPropertiesStateSucceeded`, `UpdateRunPropertiesStateUnknown`
- New type alias `UpdateSummariesPropertiesState` with values `UpdateSummariesPropertiesStateAppliedSuccessfully`, `UpdateSummariesPropertiesStateNeedsAttention`, `UpdateSummariesPropertiesStatePreparationFailed`, `UpdateSummariesPropertiesStatePreparationInProgress`, `UpdateSummariesPropertiesStateUnknown`, `UpdateSummariesPropertiesStateUpdateAvailable`, `UpdateSummariesPropertiesStateUpdateFailed`, `UpdateSummariesPropertiesStateUpdateInProgress`
- New function `*ClustersClient.BeginExtendSoftwareAssuranceBenefit(context.Context, string, string, SoftwareAssuranceChangeRequest, *ClustersClientBeginExtendSoftwareAssuranceBenefitOptions) (*runtime.Poller[ClustersClientExtendSoftwareAssuranceBenefitResponse], error)`
- New function `*ExtensionsClient.BeginUpgrade(context.Context, string, string, string, string, ExtensionUpgradeParameters, *ExtensionsClientBeginUpgradeOptions) (*runtime.Poller[ExtensionsClientUpgradeResponse], error)`
- New function `NewOffersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*OffersClient, error)`
- New function `*OffersClient.Get(context.Context, string, string, string, string, *OffersClientGetOptions) (OffersClientGetResponse, error)`
- New function `*OffersClient.NewListByClusterPager(string, string, *OffersClientListByClusterOptions) *runtime.Pager[OffersClientListByClusterResponse]`
- New function `*OffersClient.NewListByPublisherPager(string, string, string, *OffersClientListByPublisherOptions) *runtime.Pager[OffersClientListByPublisherResponse]`
- New function `NewPublishersClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PublishersClient, error)`
- New function `*PublishersClient.Get(context.Context, string, string, string, *PublishersClientGetOptions) (PublishersClientGetResponse, error)`
- New function `*PublishersClient.NewListByClusterPager(string, string, *PublishersClientListByClusterOptions) *runtime.Pager[PublishersClientListByClusterResponse]`
- New function `NewSKUsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SKUsClient, error)`
- New function `*SKUsClient.Get(context.Context, string, string, string, string, string, *SKUsClientGetOptions) (SKUsClientGetResponse, error)`
- New function `*SKUsClient.NewListByOfferPager(string, string, string, string, *SKUsClientListByOfferOptions) *runtime.Pager[SKUsClientListByOfferResponse]`
- New function `NewUpdateRunsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UpdateRunsClient, error)`
- New function `*UpdateRunsClient.BeginDelete(context.Context, string, string, string, string, *UpdateRunsClientBeginDeleteOptions) (*runtime.Poller[UpdateRunsClientDeleteResponse], error)`
- New function `*UpdateRunsClient.Get(context.Context, string, string, string, string, *UpdateRunsClientGetOptions) (UpdateRunsClientGetResponse, error)`
- New function `*UpdateRunsClient.NewListPager(string, string, string, *UpdateRunsClientListOptions) *runtime.Pager[UpdateRunsClientListResponse]`
- New function `*UpdateRunsClient.Put(context.Context, string, string, string, string, UpdateRun, *UpdateRunsClientPutOptions) (UpdateRunsClientPutResponse, error)`
- New function `NewUpdateSummariesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UpdateSummariesClient, error)`
- New function `*UpdateSummariesClient.BeginDelete(context.Context, string, string, *UpdateSummariesClientBeginDeleteOptions) (*runtime.Poller[UpdateSummariesClientDeleteResponse], error)`
- New function `*UpdateSummariesClient.Get(context.Context, string, string, *UpdateSummariesClientGetOptions) (UpdateSummariesClientGetResponse, error)`
- New function `*UpdateSummariesClient.NewListPager(string, string, *UpdateSummariesClientListOptions) *runtime.Pager[UpdateSummariesClientListResponse]`
- New function `*UpdateSummariesClient.Put(context.Context, string, string, UpdateSummaries, *UpdateSummariesClientPutOptions) (UpdateSummariesClientPutResponse, error)`
- New function `NewUpdatesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*UpdatesClient, error)`
- New function `*UpdatesClient.BeginDelete(context.Context, string, string, string, *UpdatesClientBeginDeleteOptions) (*runtime.Poller[UpdatesClientDeleteResponse], error)`
- New function `*UpdatesClient.Get(context.Context, string, string, string, *UpdatesClientGetOptions) (UpdatesClientGetResponse, error)`
- New function `*UpdatesClient.NewListPager(string, string, *UpdatesClientListOptions) *runtime.Pager[UpdatesClientListResponse]`
- New function `*UpdatesClient.BeginPost(context.Context, string, string, string, *UpdatesClientBeginPostOptions) (*runtime.Poller[UpdatesClientPostResponse], error)`
- New function `*UpdatesClient.Put(context.Context, string, string, string, Update, *UpdatesClientPutOptions) (UpdatesClientPutResponse, error)`
- New struct `ClustersClientExtendSoftwareAssuranceBenefitResponse`
- New struct `ExtensionInstanceView`
- New struct `ExtensionInstanceViewStatus`
- New struct `ExtensionUpgradeParameters`
- New struct `ExtensionsClientUpgradeResponse`
- New struct `ManagedServiceIdentity`
- New struct `Offer`
- New struct `OfferList`
- New struct `OfferProperties`
- New struct `OffersClient`
- New struct `OffersClientListByClusterResponse`
- New struct `OffersClientListByPublisherResponse`
- New struct `PackageVersionInfo`
- New struct `PrecheckResult`
- New struct `PrecheckResultTags`
- New struct `Publisher`
- New struct `PublisherList`
- New struct `PublisherProperties`
- New struct `PublishersClient`
- New struct `PublishersClientListByClusterResponse`
- New struct `SKU`
- New struct `SKUList`
- New struct `SKUMappings`
- New struct `SKUProperties`
- New struct `SKUsClient`
- New struct `SKUsClientListByOfferResponse`
- New struct `SoftwareAssuranceChangeRequest`
- New struct `SoftwareAssuranceChangeRequestProperties`
- New struct `SoftwareAssuranceProperties`
- New struct `Step`
- New struct `Update`
- New struct `UpdateList`
- New struct `UpdatePrerequisite`
- New struct `UpdateProperties`
- New struct `UpdateRun`
- New struct `UpdateRunList`
- New struct `UpdateRunProperties`
- New struct `UpdateRunsClient`
- New struct `UpdateRunsClientDeleteResponse`
- New struct `UpdateRunsClientListResponse`
- New struct `UpdateStateProperties`
- New struct `UpdateSummaries`
- New struct `UpdateSummariesClient`
- New struct `UpdateSummariesClientDeleteResponse`
- New struct `UpdateSummariesClientListResponse`
- New struct `UpdateSummariesList`
- New struct `UpdateSummariesProperties`
- New struct `UpdatesClient`
- New struct `UpdatesClientDeleteResponse`
- New struct `UpdatesClientListResponse`
- New struct `UpdatesClientPostResponse`
- New struct `UserAssignedIdentity`
- New field `Identity` in struct `Cluster`
- New field `EhcResourceID` in struct `ClusterNode`
- New field `LastLicensingTimestamp` in struct `ClusterNode`
- New field `NodeType` in struct `ClusterNode`
- New field `OSDisplayVersion` in struct `ClusterNode`
- New field `Identity` in struct `ClusterPatch`
- New field `ResourceProviderObjectID` in struct `ClusterProperties`
- New field `SoftwareAssuranceProperties` in struct `ClusterProperties`
- New field `SupportedCapabilities` in struct `ClusterReportedProperties`
- New field `EnableAutomaticUpgrade` in struct `ExtensionParameters`
- New field `InstanceView` in struct `PerNodeExtensionState`
- New field `TypeHandlerVersion` in struct `PerNodeExtensionState`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.0.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/azurestackhci/armazurestackhci` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).