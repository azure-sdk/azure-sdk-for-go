# Release History

## 0.2.0 (2024-07-16)
### Breaking Changes

- Function `*PoolsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, PoolUpdate, *PoolsClientBeginUpdateOptions)` to `(context.Context, string, string, Pool, *PoolsClientBeginUpdateOptions)`
- Type of `ImageVersion.Properties` has been changed from `*ImageVersionProperties` to `*VersionProperties`
- Type of `Quota.Name` has been changed from `*string` to `*QuotaName`
- Function `*AgentProfileUpdate.GetAgentProfileUpdate` has been removed
- Function `*AutomaticResourcePredictionsProfileUpdate.GetResourcePredictionsProfileUpdate` has been removed
- Function `*ManualResourcePredictionsProfileUpdate.GetResourcePredictionsProfileUpdate` has been removed
- Function `*ResourcePredictionsProfileUpdate.GetResourcePredictionsProfileUpdate` has been removed
- Function `*StatefulUpdate.GetAgentProfileUpdate` has been removed
- Function `*StatelessAgentProfileUpdate.GetAgentProfileUpdate` has been removed
- Function `*SubscriptionUsagesClient.NewListByLocationPager` has been removed
- Struct `AutomaticResourcePredictionsProfileUpdate` has been removed
- Struct `ImageVersionProperties` has been removed
- Struct `ManualResourcePredictionsProfileUpdate` has been removed
- Struct `PoolUpdate` has been removed
- Struct `PoolUpdateProperties` has been removed
- Struct `QuotaListResult` has been removed
- Struct `QuotaProperties` has been removed
- Struct `StatefulUpdate` has been removed
- Struct `StatelessAgentProfileUpdate` has been removed
- Field `Properties`, `SystemData`, `Type` of struct `Quota` has been removed

### Features Added

- New function `*SubscriptionUsagesClient.NewUsagesPager(string, *SubscriptionUsagesClientUsagesOptions) *runtime.Pager[SubscriptionUsagesClientUsagesResponse]`
- New struct `PagedQuota`
- New struct `VersionProperties`
- New field `CurrentValue`, `Limit`, `Unit` in struct `Quota`


## 0.1.0 (2024-05-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).