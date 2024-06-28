# Release History

## 0.2.0 (2024-06-25)
### Breaking Changes

- Function `*SubscriptionUsagesClient.NewListByLocationPager` parameter(s) have been changed from `(string, *SubscriptionUsagesClientListByLocationOptions)` to `(*SubscriptionUsagesClientListByLocationOptions)`
- Type of `Quota.Name` has been changed from `*string` to `*QuotaName`
- Struct `QuotaListResult` has been removed
- Struct `QuotaProperties` has been removed
- Field `ID`, `Properties`, `SystemData`, `Type` of struct `Quota` has been removed
- Field `QuotaListResult` of struct `SubscriptionUsagesClientListByLocationResponse` has been removed

### Features Added

- New struct `ResourceListResult`
- New field `CurrentValue`, `Limit`, `Unit` in struct `Quota`
- New anonymous field `ResourceListResult` in struct `SubscriptionUsagesClientListByLocationResponse`


## 0.1.0 (2024-05-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).