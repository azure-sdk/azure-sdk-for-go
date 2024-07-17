# Release History

## 0.2.0 (2024-07-17)
### Breaking Changes

- Type of `AgentProfile.ResourcePredictions` has been changed from `any` to `*ResourcePredictions`
- Type of `PoolUpdateProperties.AgentProfile` has been changed from `AgentProfileUpdateClassification` to `AgentProfileClassification`
- Type of `Quota.Name` has been changed from `*string` to `*QuotaName`
- Type of `Stateful.ResourcePredictions` has been changed from `any` to `*ResourcePredictions`
- Type of `StatelessAgentProfile.ResourcePredictions` has been changed from `any` to `*ResourcePredictions`
- `ManagedServiceIdentityTypeSystemAssignedUserAssigned` from enum `ManagedServiceIdentityType` has been removed
- Function `*AgentProfileUpdate.GetAgentProfileUpdate` has been removed
- Function `*AutomaticResourcePredictionsProfileUpdate.GetResourcePredictionsProfileUpdate` has been removed
- Function `*ManualResourcePredictionsProfileUpdate.GetResourcePredictionsProfileUpdate` has been removed
- Function `*ResourcePredictionsProfileUpdate.GetResourcePredictionsProfileUpdate` has been removed
- Function `*StatefulUpdate.GetAgentProfileUpdate` has been removed
- Function `*StatelessAgentProfileUpdate.GetAgentProfileUpdate` has been removed
- Function `*SubscriptionUsagesClient.NewListByLocationPager` has been removed
- Struct `AutomaticResourcePredictionsProfileUpdate` has been removed
- Struct `ManualResourcePredictionsProfileUpdate` has been removed
- Struct `QuotaListResult` has been removed
- Struct `QuotaProperties` has been removed
- Struct `StatefulUpdate` has been removed
- Struct `StatelessAgentProfileUpdate` has been removed
- Field `Properties`, `SystemData`, `Type` of struct `Quota` has been removed

### Features Added

- New value `ManagedServiceIdentityTypeSystemAndUserAssigned` added to enum type `ManagedServiceIdentityType`
- New function `*SubscriptionUsagesClient.NewUsagesPager(string, *SubscriptionUsagesClientUsagesOptions) *runtime.Pager[SubscriptionUsagesClientUsagesResponse]`
- New struct `PagedQuota`
- New struct `ResourcePredictions`
- New field `CurrentValue`, `Limit`, `Unit` in struct `Quota`


## 0.1.0 (2024-05-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).