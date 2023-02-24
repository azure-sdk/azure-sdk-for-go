# Release History

## 2.0.0 (2023-02-24)
### Breaking Changes

- Type of `ModernReservationRecommendation.Properties` has been changed from `*ModernReservationRecommendationProperties` to `ModernReservationRecommendationPropertiesClassification`
- Field `Etag` of struct `CreditSummary` has been removed
- Field `Tags` of struct `CreditSummary` has been removed
- Field `MarketplaceCharges` of struct `LegacyChargeSummaryProperties` has been removed

### Features Added

- New value `EventTypeCreditExpired` added to type alias `EventType`
- New function `*ModernReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*ModernSharedScopeReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*ModernSingleScopeReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `NewOperationsResultClient(azcore.TokenCredential, *arm.ClientOptions) (*OperationsResultClient, error)`
- New function `*OperationsResultClient.NewListPager(*OperationsResultClientListOptions) *runtime.Pager[OperationsResultClientListResponse]`
- New function `NewPriceSheetsClient(azcore.TokenCredential, *arm.ClientOptions) (*PriceSheetsClient, error)`
- New function `*PriceSheetsClient.Get(context.Context, string, *PriceSheetsClientGetOptions) (PriceSheetsClientGetResponse, error)`
- New function `*PriceSheetsClient.GetByBillingPeriod(context.Context, string, string, *PriceSheetsClientGetByBillingPeriodOptions) (PriceSheetsClientGetByBillingPeriodResponse, error)`
- New struct `MeterDetailsV2`
- New struct `ModernSharedScopeReservationRecommendationProperties`
- New struct `ModernSingleScopeReservationRecommendationProperties`
- New struct `OperationListResultV2`
- New struct `OperationV2`
- New struct `OperationV2Display`
- New struct `OperationsResultClient`
- New struct `PriceSheetModelV2`
- New struct `PriceSheetPropertiesV2`
- New struct `PriceSheetResultV2`
- New struct `PriceSheetsClient`
- New field `ETag` in struct `CreditSummary`
- New field `AzureMarketplaceCharges` in struct `LegacyChargeSummaryProperties`
- New field `SubscriptionID` in struct `ModernChargeSummaryProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).