# Release History

## 2.0.0 (2023-03-13)
### Breaking Changes

- Struct `ErrorResponseError` has been removed
- Field `OccurredTime` of struct `AvailabilityStatusProperties` has been removed
- Field `UnavailabilitySummary` of struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed
- Field `UnavailableOccurredTime` of struct `AvailabilityStatusPropertiesRecentlyResolved` has been removed
- Field `Error` of struct `ErrorResponse` has been removed

### Features Added

- New type alias `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New type alias `EventLevelValues` with values `EventLevelValuesCritical`, `EventLevelValuesError`, `EventLevelValuesInformational`, `EventLevelValuesWarning`
- New type alias `EventSourceValues` with values `EventSourceValuesResourceHealth`, `EventSourceValuesServiceHealth`
- New type alias `EventStatusValues` with values `EventStatusValuesActive`, `EventStatusValuesResolved`
- New type alias `EventTypeValues` with values `EventTypeValuesEmergingIssues`, `EventTypeValuesHealthAdvisory`, `EventTypeValuesPlannedMaintenance`, `EventTypeValuesRCA`, `EventTypeValuesSecurityAdvisory`, `EventTypeValuesServiceIssue`
- New type alias `LevelValues` with values `LevelValuesCritical`, `LevelValuesWarning`
- New type alias `LinkTypeValues` with values `LinkTypeValuesButton`, `LinkTypeValuesHyperlink`
- New function `NewEventClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EventClient, error)`
- New function `*EventClient.GetBySubscriptionIDAndTrackingID(context.Context, string, *EventClientGetBySubscriptionIDAndTrackingIDOptions) (EventClientGetBySubscriptionIDAndTrackingIDResponse, error)`
- New function `*EventClient.GetByTenantIDAndTrackingID(context.Context, string, *EventClientGetByTenantIDAndTrackingIDOptions) (EventClientGetByTenantIDAndTrackingIDResponse, error)`
- New function `NewEventsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*EventsClient, error)`
- New function `*EventsClient.NewListBySingleResourcePager(string, *EventsClientListBySingleResourceOptions) *runtime.Pager[EventsClientListBySingleResourceResponse]`
- New function `*EventsClient.NewListBySubscriptionIDPager(*EventsClientListBySubscriptionIDOptions) *runtime.Pager[EventsClientListBySubscriptionIDResponse]`
- New function `*EventsClient.NewListByTenantIDPager(*EventsClientListByTenantIDOptions) *runtime.Pager[EventsClientListByTenantIDResponse]`
- New function `NewImpactedResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ImpactedResourcesClient, error)`
- New function `*ImpactedResourcesClient.Get(context.Context, string, string, *ImpactedResourcesClientGetOptions) (ImpactedResourcesClientGetResponse, error)`
- New function `*ImpactedResourcesClient.NewListBySubscriptionIDAndEventIDPager(string, *ImpactedResourcesClientListBySubscriptionIDAndEventIDOptions) *runtime.Pager[ImpactedResourcesClientListBySubscriptionIDAndEventIDResponse]`
- New struct `Event`
- New struct `EventClient`
- New struct `EventImpactedResource`
- New struct `EventImpactedResourceListResult`
- New struct `EventImpactedResourceProperties`
- New struct `EventProperties`
- New struct `EventPropertiesAdditionalInformation`
- New struct `EventPropertiesArticle`
- New struct `EventPropertiesRecommendedActions`
- New struct `EventPropertiesRecommendedActionsItem`
- New struct `Events`
- New struct `EventsClient`
- New struct `Faq`
- New struct `Impact`
- New struct `ImpactedResourcesClient`
- New struct `ImpactedServiceRegion`
- New struct `KeyValueItem`
- New struct `Link`
- New struct `LinkDisplayText`
- New struct `ProxyResource`
- New struct `SystemData`
- New struct `Update`
- New field `OccuredTime` in struct `AvailabilityStatusProperties`
- New field `UnavailableOccuredTime` in struct `AvailabilityStatusPropertiesRecentlyResolved`
- New field `UnavailableSummary` in struct `AvailabilityStatusPropertiesRecentlyResolved`
- New field `Code` in struct `ErrorResponse`
- New field `Details` in struct `ErrorResponse`
- New field `Message` in struct `ErrorResponse`
- New field `SystemData` in struct `ImpactedResourceStatus`
- New field `SystemData` in struct `Resource`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcehealth/armresourcehealth` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).