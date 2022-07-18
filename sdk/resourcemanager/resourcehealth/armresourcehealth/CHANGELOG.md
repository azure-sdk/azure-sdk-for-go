# Release History

## 2.0.0 (2022-07-18)
### Breaking Changes

- Const `ReasonTypeValuesUnplanned` has been removed
- Const `AvailabilityStateValuesAvailable` has been removed
- Const `AvailabilityStateValuesUnavailable` has been removed
- Const `AvailabilityStateValuesDegraded` has been removed
- Const `AvailabilityStateValuesUnknown` has been removed
- Const `ReasonChronicityTypesPersistent` has been removed
- Const `ReasonChronicityTypesTransient` has been removed
- Const `ReasonTypeValuesPlanned` has been removed
- Const `ReasonTypeValuesUserInitiated` has been removed
- Function `*AvailabilityStatusesClient.GetByResource` has been removed
- Function `PossibleAvailabilityStateValuesValues` has been removed
- Function `PossibleReasonTypeValuesValues` has been removed
- Function `*AvailabilityStatusesClient.NewListByResourceGroupPager` has been removed
- Function `*AvailabilityStatusesClient.NewListPager` has been removed
- Function `NewAvailabilityStatusesClient` has been removed
- Function `PossibleReasonChronicityTypesValues` has been removed
- Function `*AvailabilityStatusesClient.NewListBySubscriptionIDPager` has been removed
- Struct `AvailabilityStatus` has been removed
- Struct `AvailabilityStatusListResult` has been removed
- Struct `AvailabilityStatusProperties` has been removed
- Struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed
- Struct `AvailabilityStatusesClient` has been removed
- Struct `AvailabilityStatusesClientGetByResourceOptions` has been removed
- Struct `AvailabilityStatusesClientGetByResourceResponse` has been removed
- Struct `AvailabilityStatusesClientListByResourceGroupOptions` has been removed
- Struct `AvailabilityStatusesClientListByResourceGroupResponse` has been removed
- Struct `AvailabilityStatusesClientListBySubscriptionIDOptions` has been removed
- Struct `AvailabilityStatusesClientListBySubscriptionIDResponse` has been removed
- Struct `AvailabilityStatusesClientListOptions` has been removed
- Struct `AvailabilityStatusesClientListResponse` has been removed
- Struct `ErrorResponseError` has been removed
- Struct `ImpactedRegion` has been removed
- Struct `ImpactedResourceStatus` has been removed
- Struct `ImpactedResourceStatusProperties` has been removed
- Struct `RecommendedAction` has been removed
- Struct `ServiceImpactingEvent` has been removed
- Struct `ServiceImpactingEventIncidentProperties` has been removed
- Struct `ServiceImpactingEventStatus` has been removed
- Struct `StatusBanner` has been removed
- Field `Error` of struct `ErrorResponse` has been removed

### Features Added

- New const `EventLevelValuesCritical`
- New const `LinkTypeValuesHyperlink`
- New const `LevelValuesCritical`
- New const `LinkTypeValuesButton`
- New const `EventLevelValuesWarning`
- New const `EventTypeValuesSecurityAdvisory`
- New const `LevelValuesWarning`
- New const `EventTypeValuesServiceIssue`
- New const `EventTypeValuesPlannedMaintenance`
- New const `EventTypeValuesRCA`
- New const `EventLevelValuesError`
- New const `EventLevelValuesInformational`
- New const `EventSourceValuesResourceHealth`
- New const `EventStatusValuesResolved`
- New const `EventStatusValuesActive`
- New const `EventTypeValuesHealthAdvisory`
- New const `EventTypeValuesEmergingIssues`
- New const `EventSourceValuesServiceHealth`
- New function `PossibleEventTypeValuesValues() []EventTypeValues`
- New function `PossibleLinkTypeValuesValues() []LinkTypeValues`
- New function `PossibleEventStatusValuesValues() []EventStatusValues`
- New function `PossibleEventSourceValuesValues() []EventSourceValues`
- New function `PossibleLevelValuesValues() []LevelValues`
- New function `*EventsClient.NewListBySingleResourcePager(string, *EventsClientListBySingleResourceOptions) *runtime.Pager[EventsClientListBySingleResourceResponse]`
- New function `PossibleEventLevelValuesValues() []EventLevelValues`
- New function `NewEventsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EventsClient, error)`
- New function `*EventsClient.NewListBySubscriptionIDPager(*EventsClientListBySubscriptionIDOptions) *runtime.Pager[EventsClientListBySubscriptionIDResponse]`
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
- New struct `ImpactedServiceRegion`
- New struct `Link`
- New struct `LinkDisplayText`
- New struct `Update`
- New field `IsDataAction` in struct `Operation`
- New field `Message` in struct `ErrorResponse`
- New field `Code` in struct `ErrorResponse`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).