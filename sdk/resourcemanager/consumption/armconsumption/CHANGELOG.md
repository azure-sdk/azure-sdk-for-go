# Release History

## 2.0.0 (2023-08-29)
### Breaking Changes

- Type of `ModernReservationRecommendation.Properties` has been changed from `*ModernReservationRecommendationProperties` to `ModernReservationRecommendationPropertiesClassification`
- Field `Etag`, `Tags` of struct `CreditSummary` has been removed
- Field `MarketplaceCharges` of struct `LegacyChargeSummaryProperties` has been removed

### Features Added

- New value `EventTypeCreditExpired` added to enum type `EventType`
- New enum type `OperationStatusType` with values `OperationStatusTypeCompleted`, `OperationStatusTypeFailed`, `OperationStatusTypeRunning`
- New enum type `OrgType` with values `OrgTypeContributorOrgType`, `OrgTypePrimaryOrgType`
- New function `*ModernReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*ModernSharedScopeReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*ModernSingleScopeReservationRecommendationProperties.GetModernReservationRecommendationProperties() *ModernReservationRecommendationProperties`
- New function `*PriceSheetClient.BeginDownloadByBillingAccountPeriod(context.Context, string, string, *PriceSheetClientBeginDownloadByBillingAccountPeriodOptions) (*runtime.Poller[PriceSheetClientDownloadByBillingAccountPeriodResponse], error)`
- New struct `ModernSharedScopeReservationRecommendationProperties`
- New struct `ModernSingleScopeReservationRecommendationProperties`
- New struct `OperationStatus`
- New struct `PricesheetDownloadProperties`
- New struct `SavingsPlan`
- New field `OverageRefund` in struct `BalanceProperties`
- New field `ETag` in struct `CreditSummary`
- New field `IsEstimatedBalance` in struct `CreditSummaryProperties`
- New field `IsEstimatedBalance` in struct `EventProperties`
- New field `AzureMarketplaceCharges` in struct `LegacyChargeSummaryProperties`
- New field `IsEstimatedBalance`, `OrgType`, `UsedAmount` in struct `LotProperties`
- New field `SubscriptionID` in struct `ModernChargeSummaryProperties`
- New field `SavingsPlan` in struct `PriceSheetProperties`
- New field `Filter` in struct `ReservationRecommendationDetailsClientGetOptions`
- New field `PreviewMarkupPercentage`, `UseMarkupIfPartner` in struct `ReservationTransactionsClientListOptions`


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/consumption/armconsumption` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).