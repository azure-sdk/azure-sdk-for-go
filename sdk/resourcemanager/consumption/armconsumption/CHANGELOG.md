# Release History

## 2.0.0 (2022-11-04)
### Breaking Changes

- Function `*PriceSheetClient.Get` has been removed
- Function `*PriceSheetClient.GetByBillingPeriod` has been removed
- Field `ETag` of struct `LegacyChargeSummary` has been removed
- Field `ETag` of struct `ChargeSummary` has been removed
- Field `MarketplaceCharges` of struct `LegacyChargeSummaryProperties` has been removed
- Field `Tags` of struct `CreditSummary` has been removed
- Field `Etag` of struct `CreditSummary` has been removed
- Field `ETag` of struct `ModernChargeSummary` has been removed

### Features Added

- New const `EventTypeCreditExpired`
- New function `*PriceSheetClient.NewGetPager(*PriceSheetClientGetOptions) *runtime.Pager[PriceSheetClientGetResponse]`
- New function `*PriceSheetClient.NewGetByBillingPeriodPager(string, *PriceSheetClientGetByBillingPeriodOptions) *runtime.Pager[PriceSheetClientGetByBillingPeriodResponse]`
- New field `AzureMarketplaceCharges` in struct `LegacyChargeSummaryProperties`
- New field `Etag` in struct `LegacyChargeSummary`
- New field `Tags` in struct `LegacyChargeSummary`
- New field `ResourceType` in struct `ModernReservationRecommendationProperties`
- New field `SubscriptionID` in struct `ModernReservationRecommendationProperties`
- New field `Etag` in struct `ChargeSummary`
- New field `Tags` in struct `ChargeSummary`
- New field `Etag` in struct `ModernChargeSummary`
- New field `Tags` in struct `ModernChargeSummary`
- New field `NextLink` in struct `PriceSheetResult`
- New field `ETag` in struct `CreditSummary`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).