# Release History

## 0.11.0 (2024-07-16)
### Breaking Changes

- Function `*AlertsClient.ChangeState` parameter(s) have been changed from `(context.Context, string, AlertState, *AlertsClientChangeStateOptions)` to `(context.Context, string, string, AlertState, *AlertsClientChangeStateOptions)`
- Function `*AlertsClient.GetByID` parameter(s) have been changed from `(context.Context, string, *AlertsClientGetByIDOptions)` to `(context.Context, string, string, *AlertsClientGetByIDOptions)`
- Function `*AlertsClient.GetHistory` parameter(s) have been changed from `(context.Context, string, *AlertsClientGetHistoryOptions)` to `(context.Context, string, string, *AlertsClientGetHistoryOptions)`
- Function `*AlertsClient.GetSummary` parameter(s) have been changed from `(context.Context, AlertsSummaryGroupByFields, *AlertsClientGetSummaryOptions)` to `(context.Context, string, AlertsSummaryGroupByFields, *AlertsClientGetSummaryOptions)`
- Function `*AlertsClient.NewGetAllPager` parameter(s) have been changed from `(*AlertsClientGetAllOptions)` to `(string, *AlertsClientGetAllOptions)`
- Function `*PrometheusRuleGroupsClient.Update` parameter(s) have been changed from `(context.Context, string, string, PrometheusRuleGroupResourcePatch, *PrometheusRuleGroupsClientUpdateOptions)` to `(context.Context, string, string, PrometheusRuleGroupResourcePatchParameters, *PrometheusRuleGroupsClientUpdateOptions)`
- Function `*ClientFactory.NewTenantActivityLogAlertsClient` has been removed
- Function `*PrometheusRuleGroupsClient.NewListBySubscriptionPager` has been removed
- Function `NewTenantActivityLogAlertsClient` has been removed
- Function `*TenantActivityLogAlertsClient.CreateOrUpdate` has been removed
- Function `*TenantActivityLogAlertsClient.Delete` has been removed
- Function `*TenantActivityLogAlertsClient.Get` has been removed
- Function `*TenantActivityLogAlertsClient.NewListByManagementGroupPager` has been removed
- Function `*TenantActivityLogAlertsClient.NewListByTenantPager` has been removed
- Function `*TenantActivityLogAlertsClient.Update` has been removed
- Struct `ActionGroup` has been removed
- Struct `ActionList` has been removed
- Struct `AlertRuleAllOfCondition` has been removed
- Struct `AlertRuleAnyOfOrLeafCondition` has been removed
- Struct `AlertRuleLeafCondition` has been removed
- Struct `AlertRuleProperties` has been removed
- Struct `PrometheusRuleGroupResourcePatch` has been removed
- Struct `PrometheusRuleGroupResourcePatchProperties` has been removed
- Struct `TenantActivityLogAlertResource` has been removed
- Struct `TenantAlertRuleList` has been removed
- Struct `TenantAlertRulePatchObject` has been removed
- Struct `TenantAlertRulePatchProperties` has been removed

### Features Added

- New value `ActionTypeCorrelateAlerts` added to enum type `ActionType`
- New value `MonitorServiceResourceHealth` added to enum type `MonitorService`
- New enum type `NotificationsForCorrelatedAlerts` with values `NotificationsForCorrelatedAlertsNotifyAlways`, `NotificationsForCorrelatedAlertsSuppressAlways`
- New enum type `Status` with values `StatusFailed`, `StatusSucceeded`
- New enum type `Type` with values `TypePrometheusInstantQuery`, `TypePrometheusRangeQuery`
- New enum type `UpdateType` with values `UpdateTypeTimeBased`
- New function `*AlertEnrichmentItem.GetAlertEnrichmentItem() *AlertEnrichmentItem`
- New function `*AlertsClient.GetEnrichments(context.Context, string, string, *AlertsClientGetEnrichmentsOptions) (AlertsClientGetEnrichmentsResponse, error)`
- New function `*AlertsClient.NewListEnrichmentsPager(string, string, *AlertsClientListEnrichmentsOptions) *runtime.Pager[AlertsClientListEnrichmentsResponse]`
- New function `NewCROSSClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CROSSClient, error)`
- New function `*CROSSClient.NewVERBREAKINGCHANGETESTPrometheusRuleGroupsListBySubscriptionPager(*CROSSClientVERBREAKINGCHANGETESTPrometheusRuleGroupsListBySubscriptionOptions) *runtime.Pager[CROSSClientVERBREAKINGCHANGETESTPrometheusRuleGroupsListBySubscriptionResponse]`
- New function `*ClientFactory.NewCROSSClient() *CROSSClient`
- New function `*CorrelateAlerts.GetAction() *Action`
- New function `*PrometheusEnrichmentItem.GetAlertEnrichmentItem() *AlertEnrichmentItem`
- New function `*PrometheusEnrichmentItem.GetPrometheusEnrichmentItem() *PrometheusEnrichmentItem`
- New function `*PrometheusInstantQuery.GetAlertEnrichmentItem() *AlertEnrichmentItem`
- New function `*PrometheusInstantQuery.GetPrometheusEnrichmentItem() *PrometheusEnrichmentItem`
- New function `*PrometheusRangeQuery.GetAlertEnrichmentItem() *AlertEnrichmentItem`
- New function `*PrometheusRangeQuery.GetPrometheusEnrichmentItem() *PrometheusEnrichmentItem`
- New struct `AlertEnrichmentProperties`
- New struct `AlertEnrichmentResponse`
- New struct `AlertEnrichmentsList`
- New struct `CorrelateAlerts`
- New struct `CorrelateBy`
- New struct `CorrelationDetails`
- New struct `CorrelationUpdates`
- New struct `PrometheusInstantQuery`
- New struct `PrometheusRangeQuery`
- New struct `PrometheusRuleGroupResourcePatchParameters`
- New struct `PrometheusRuleGroupResourcePatchParametersProperties`
- New field `Category` in struct `AlertRuleRecommendationProperties`
- New field `CorrelationDetails`, `HasEnrichments`, `IsStatefulAlert` in struct `Essentials`


## 0.10.0 (2024-03-01)
### Breaking Changes

- Type of `AlertsClientChangeStateOptions.Comment` has been changed from `*string` to `*Comments`

### Features Added

- New function `NewAlertRuleRecommendationsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*AlertRuleRecommendationsClient, error)`
- New function `*AlertRuleRecommendationsClient.NewListByResourcePager(string, *AlertRuleRecommendationsClientListByResourceOptions) *runtime.Pager[AlertRuleRecommendationsClientListByResourceResponse]`
- New function `*AlertRuleRecommendationsClient.NewListByTargetTypePager(string, *AlertRuleRecommendationsClientListByTargetTypeOptions) *runtime.Pager[AlertRuleRecommendationsClientListByTargetTypeResponse]`
- New function `*ClientFactory.NewAlertRuleRecommendationsClient() *AlertRuleRecommendationsClient`
- New function `*ClientFactory.NewPrometheusRuleGroupsClient() *PrometheusRuleGroupsClient`
- New function `*ClientFactory.NewTenantActivityLogAlertsClient() *TenantActivityLogAlertsClient`
- New function `NewPrometheusRuleGroupsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*PrometheusRuleGroupsClient, error)`
- New function `*PrometheusRuleGroupsClient.CreateOrUpdate(context.Context, string, string, PrometheusRuleGroupResource, *PrometheusRuleGroupsClientCreateOrUpdateOptions) (PrometheusRuleGroupsClientCreateOrUpdateResponse, error)`
- New function `*PrometheusRuleGroupsClient.Delete(context.Context, string, string, *PrometheusRuleGroupsClientDeleteOptions) (PrometheusRuleGroupsClientDeleteResponse, error)`
- New function `*PrometheusRuleGroupsClient.Get(context.Context, string, string, *PrometheusRuleGroupsClientGetOptions) (PrometheusRuleGroupsClientGetResponse, error)`
- New function `*PrometheusRuleGroupsClient.NewListByResourceGroupPager(string, *PrometheusRuleGroupsClientListByResourceGroupOptions) *runtime.Pager[PrometheusRuleGroupsClientListByResourceGroupResponse]`
- New function `*PrometheusRuleGroupsClient.NewListBySubscriptionPager(*PrometheusRuleGroupsClientListBySubscriptionOptions) *runtime.Pager[PrometheusRuleGroupsClientListBySubscriptionResponse]`
- New function `*PrometheusRuleGroupsClient.Update(context.Context, string, string, PrometheusRuleGroupResourcePatch, *PrometheusRuleGroupsClientUpdateOptions) (PrometheusRuleGroupsClientUpdateResponse, error)`
- New function `NewTenantActivityLogAlertsClient(azcore.TokenCredential, *arm.ClientOptions) (*TenantActivityLogAlertsClient, error)`
- New function `*TenantActivityLogAlertsClient.CreateOrUpdate(context.Context, string, string, TenantActivityLogAlertResource, *TenantActivityLogAlertsClientCreateOrUpdateOptions) (TenantActivityLogAlertsClientCreateOrUpdateResponse, error)`
- New function `*TenantActivityLogAlertsClient.Delete(context.Context, string, string, *TenantActivityLogAlertsClientDeleteOptions) (TenantActivityLogAlertsClientDeleteResponse, error)`
- New function `*TenantActivityLogAlertsClient.Get(context.Context, string, string, *TenantActivityLogAlertsClientGetOptions) (TenantActivityLogAlertsClientGetResponse, error)`
- New function `*TenantActivityLogAlertsClient.NewListByManagementGroupPager(string, *TenantActivityLogAlertsClientListByManagementGroupOptions) *runtime.Pager[TenantActivityLogAlertsClientListByManagementGroupResponse]`
- New function `*TenantActivityLogAlertsClient.NewListByTenantPager(*TenantActivityLogAlertsClientListByTenantOptions) *runtime.Pager[TenantActivityLogAlertsClientListByTenantResponse]`
- New function `*TenantActivityLogAlertsClient.Update(context.Context, string, string, TenantAlertRulePatchObject, *TenantActivityLogAlertsClientUpdateOptions) (TenantActivityLogAlertsClientUpdateResponse, error)`
- New struct `ActionGroup`
- New struct `ActionList`
- New struct `AlertRuleAllOfCondition`
- New struct `AlertRuleAnyOfOrLeafCondition`
- New struct `AlertRuleLeafCondition`
- New struct `AlertRuleProperties`
- New struct `AlertRuleRecommendationProperties`
- New struct `AlertRuleRecommendationResource`
- New struct `AlertRuleRecommendationsListResponse`
- New struct `Comments`
- New struct `PrometheusRule`
- New struct `PrometheusRuleGroupAction`
- New struct `PrometheusRuleGroupProperties`
- New struct `PrometheusRuleGroupResource`
- New struct `PrometheusRuleGroupResourceCollection`
- New struct `PrometheusRuleGroupResourcePatch`
- New struct `PrometheusRuleGroupResourcePatchProperties`
- New struct `PrometheusRuleResolveConfiguration`
- New struct `RuleArmTemplate`
- New struct `TenantActivityLogAlertResource`
- New struct `TenantAlertRuleList`
- New struct `TenantAlertRulePatchObject`
- New struct `TenantAlertRulePatchProperties`


## 0.9.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 0.8.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.

## 0.8.0 (2023-03-27)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.7.0 (2022-08-19)
### Features Added

- New field `Origin` in struct `Operation`
- New field `Comment` in struct `AlertsClientChangeStateOptions`


## 0.6.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/alertsmanagement/armalertsmanagement` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.6.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).