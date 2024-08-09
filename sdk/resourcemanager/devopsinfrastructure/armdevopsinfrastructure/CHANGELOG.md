# Release History

## 0.2.0 (2024-08-09)
### Breaking Changes

- Type of `Quota.Name` has been changed from `*string` to `*QuotaName`
- Function `*SubscriptionUsagesClient.NewListByLocationPager` has been removed
- Struct `QuotaListResult` has been removed
- Struct `QuotaProperties` has been removed
- Field `Properties`, `SystemData`, `Type` of struct `Quota` has been removed

### Features Added

- New function `*SubscriptionUsagesClient.NewUsagesPager(string, *SubscriptionUsagesClientUsagesOptions) *runtime.Pager[SubscriptionUsagesClientUsagesResponse]`
- New struct `PagedQuota`
- New field `CurrentValue`, `Limit`, `Unit` in struct `Quota`


## 0.1.0 (2024-05-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).