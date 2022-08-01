# Release History

## 2.0.0 (2022-08-01)
### Breaking Changes

- Type of `ErrorResponse.Error` has been changed from `*ErrorResponseError` to `*ErrorDetail`
- Struct `ErrorResponseError` has been removed
- Field `UnavailabilitySummary` of struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed
- Field `UnavailableOccurredTime` of struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed
- Field `OccurredTime` of struct `ImpactedResourceStatusProperties` has been removed
- Field `OccurredTime` of struct `AvailabilityStatusProperties` has been removed

### Features Added

- New const `EventSourceValuesResourceHealth`
- New const `ActionTypeInternal`
- New const `EventTypeValuesHealthAdvisory`
- New const `EventTypeValuesEmergingIssues`
- New const `EventSourceValuesServiceHealth`
- New const `OriginUser`
- New const `OriginUserSystem`
- New const `SeverityValuesInformation`
- New const `EventTypeValuesSecurityAdvisory`
- New const `LinkTypeValuesButton`
- New const `StageValuesResolve`
- New const `LinkTypeValuesHyperlink`
- New const `StageValuesActive`
- New const `EventLevelValuesInformational`
- New const `EventLevelValuesWarning`
- New const `ScenarioAlerts`
- New const `IssueNameParameterDefault`
- New const `EventLevelValuesError`
- New const `LevelValuesCritical`
- New const `SeverityValuesWarning`
- New const `OriginSystem`
- New const `EventTypeValuesServiceIssue`
- New const `LevelValuesWarning`
- New const `StageValuesArchived`
- New const `EventStatusValuesActive`
- New const `EventTypeValuesPlannedMaintenance`
- New const `EventLevelValuesCritical`
- New const `EventStatusValuesResolved`
- New const `EventTypeValuesRCA`
- New const `SeverityValuesError`
- New function `*EventsClient.NewListBySingleResourcePager(string, *EventsClientListBySingleResourceOptions) *runtime.Pager[EventsClientListBySingleResourceResponse]`
- New function `NewMetadataClient(azcore.TokenCredential, *arm.ClientOptions) (*MetadataClient, error)`
- New function `PossibleSeverityValuesValues() []SeverityValues`
- New function `*ChildAvailabilityStatusesClient.NewListPager(string, *ChildAvailabilityStatusesClientListOptions) *runtime.Pager[ChildAvailabilityStatusesClientListResponse]`
- New function `*MetadataClient.GetEntity(context.Context, string, *MetadataClientGetEntityOptions) (MetadataClientGetEntityResponse, error)`
- New function `PossibleLevelValuesValues() []LevelValues`
- New function `*ChildResourcesClient.NewListPager(string, *ChildResourcesClientListOptions) *runtime.Pager[ChildResourcesClientListResponse]`
- New function `NewEventsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EventsClient, error)`
- New function `PossibleEventTypeValuesValues() []EventTypeValues`
- New function `PossibleScenarioValues() []Scenario`
- New function `PossibleIssueNameParameterValues() []IssueNameParameter`
- New function `PossibleStageValuesValues() []StageValues`
- New function `PossibleEventLevelValuesValues() []EventLevelValues`
- New function `PossibleEventStatusValuesValues() []EventStatusValues`
- New function `PossibleEventSourceValuesValues() []EventSourceValues`
- New function `*EmergingIssuesClient.Get(context.Context, IssueNameParameter, *EmergingIssuesClientGetOptions) (EmergingIssuesClientGetResponse, error)`
- New function `*EmergingIssuesClient.NewListPager(*EmergingIssuesClientListOptions) *runtime.Pager[EmergingIssuesClientListResponse]`
- New function `NewChildResourcesClient(azcore.TokenCredential, *arm.ClientOptions) (*ChildResourcesClient, error)`
- New function `*ChildAvailabilityStatusesClient.GetByResource(context.Context, string, *ChildAvailabilityStatusesClientGetByResourceOptions) (ChildAvailabilityStatusesClientGetByResourceResponse, error)`
- New function `PossibleActionTypeValues() []ActionType`
- New function `NewEmergingIssuesClient(azcore.TokenCredential, *arm.ClientOptions) (*EmergingIssuesClient, error)`
- New function `*ImpactedResourcesClient.NewListBySubscriptionIDPager(*ImpactedResourcesClientListBySubscriptionIDOptions) *runtime.Pager[ImpactedResourcesClientListBySubscriptionIDResponse]`
- New function `*EventsClient.NewListBySubscriptionIDPager(*EventsClientListBySubscriptionIDOptions) *runtime.Pager[EventsClientListBySubscriptionIDResponse]`
- New function `NewImpactedResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ImpactedResourcesClient, error)`
- New function `PossibleOriginValues() []Origin`
- New function `NewChildAvailabilityStatusesClient(azcore.TokenCredential, *arm.ClientOptions) (*ChildAvailabilityStatusesClient, error)`
- New function `*MetadataClient.NewListPager(*MetadataClientListOptions) *runtime.Pager[MetadataClientListResponse]`
- New function `PossibleLinkTypeValuesValues() []LinkTypeValues`
- New struct `ChildAvailabilityStatusesClient`
- New struct `ChildAvailabilityStatusesClientGetByResourceOptions`
- New struct `ChildAvailabilityStatusesClientGetByResourceResponse`
- New struct `ChildAvailabilityStatusesClientListOptions`
- New struct `ChildAvailabilityStatusesClientListResponse`
- New struct `ChildResourcesClient`
- New struct `ChildResourcesClientListOptions`
- New struct `ChildResourcesClientListResponse`
- New struct `EmergingIssue`
- New struct `EmergingIssueImpact`
- New struct `EmergingIssueListResult`
- New struct `EmergingIssuesClient`
- New struct `EmergingIssuesClientGetOptions`
- New struct `EmergingIssuesClientGetResponse`
- New struct `EmergingIssuesClientListOptions`
- New struct `EmergingIssuesClientListResponse`
- New struct `EmergingIssuesGetResult`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `Event`
- New struct `EventProperties`
- New struct `EventPropertiesAdditionalInformation`
- New struct `EventPropertiesArticle`
- New struct `EventPropertiesRecommendedActions`
- New struct `EventPropertiesRecommendedActionsItem`
- New struct `Events`
- New struct `EventsClient`
- New struct `EventsClientListBySingleResourceOptions`
- New struct `EventsClientListBySingleResourceResponse`
- New struct `EventsClientListBySubscriptionIDOptions`
- New struct `EventsClientListBySubscriptionIDResponse`
- New struct `Faq`
- New struct `Impact`
- New struct `ImpactedResourceListResult`
- New struct `ImpactedResourcesClient`
- New struct `ImpactedResourcesClientListBySubscriptionIDOptions`
- New struct `ImpactedResourcesClientListBySubscriptionIDResponse`
- New struct `ImpactedServiceRegion`
- New struct `Link`
- New struct `LinkDisplayText`
- New struct `MetadataClient`
- New struct `MetadataClientGetEntityOptions`
- New struct `MetadataClientGetEntityResponse`
- New struct `MetadataClientListOptions`
- New struct `MetadataClientListResponse`
- New struct `MetadataEntity`
- New struct `MetadataEntityListResult`
- New struct `MetadataEntityProperties`
- New struct `MetadataSupportedValueDetail`
- New struct `StatusActiveEvent`
- New struct `Update`
- New field `OccuredTime` in struct `AvailabilityStatusProperties`
- New field `OccuredTime` in struct `ImpactedResourceStatusProperties`
- New field `ActionType` in struct `Operation`
- New field `IsDataAction` in struct `Operation`
- New field `Origin` in struct `Operation`
- New field `UnavailableOccuredTime` in struct `AvailabilityStatusPropertiesRecentlyResolved`
- New field `UnavailableSummary` in struct `AvailabilityStatusPropertiesRecentlyResolved`
- New field `NextLink` in struct `OperationListResult`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).