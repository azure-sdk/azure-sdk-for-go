# Release History

## 2.0.0 (2023-03-23)
### Breaking Changes

- Type of `ModernReservationRecommendation.Properties` has been changed from `*ModernReservationRecommendationProperties` to `ModernReservationRecommendationPropertiesClassification`
- Field `Etag` of struct `CreditSummary` has been removed
- Field `Tags` of struct `CreditSummary` has been removed
- Field `MarketplaceCharges` of struct `LegacyChargeSummaryProperties` has been removed

### Features Added

- New value `EventTypeCreditExpired` added to enum type `EventType`
- New function `NewClientFactory(string, azcore.TokenCredential, *arm.ClientOptions) (*ClientFactory, error)`
- New function `*ClientFactory.NewAggregatedCostClient() *AggregatedCostClient`
- New function `*ClientFactory.NewBalancesClient() *BalancesClient`
- New function `*ClientFactory.NewBudgetsClient() *BudgetsClient`
- New function `*ClientFactory.NewChargesClient() *ChargesClient`
- New function `*ClientFactory.NewCreditsClient() *CreditsClient`
- New function `*ClientFactory.NewEventsClient() *EventsClient`
- New function `*ClientFactory.NewLotsClient() *LotsClient`
- New function `*ClientFactory.NewMarketplacesClient() *MarketplacesClient`
- New function `*ClientFactory.NewOperationsClient() *OperationsClient`
- New function `*ClientFactory.NewPriceSheetClient() *PriceSheetClient`
- New function `*ClientFactory.NewReservationRecommendationDetailsClient() *ReservationRecommendationDetailsClient`
- New function `*ClientFactory.NewReservationRecommendationsClient() *ReservationRecommendationsClient`
- New function `*ClientFactory.NewReservationTransactionsClient() *ReservationTransactionsClient`
- New function `*ClientFactory.NewReservationsDetailsClient() *ReservationsDetailsClient`
- New function `*ClientFactory.NewReservationsSummariesClient() *ReservationsSummariesClient`
- New function `*ClientFactory.NewTagsClient() *TagsClient`
- New function `*ClientFactory.NewUsageDetailsClient() *UsageDetailsClient`
- New function `*ModernReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*ModernSharedScopeReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*ModernSingleScopeReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New struct `ClientFactory`
- New struct `ModernSharedScopeReservationRecommendationProperties`
- New struct `ModernSingleScopeReservationRecommendationProperties`
- New field `ETag` in struct `CreditSummary`
- New field `IsEstimatedBalance` in struct `CreditSummaryProperties`
- New field `IsEstimatedBalance` in struct `EventProperties`
- New field `AzureMarketplaceCharges` in struct `LegacyChargeSummaryProperties`
- New field `IsEstimatedBalance` in struct `LotProperties`
- New field `SubscriptionID` in struct `ModernChargeSummaryProperties`
- New field `PreviewMarkupPercentage` in struct `ReservationTransactionsClientListOptions`
- New field `UseMarkupIfPartner` in struct `ReservationTransactionsClientListOptions`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).