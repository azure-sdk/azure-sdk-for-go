# Release History

## 2.2.0 (2025-04-11)
### Features Added

- New value `MarketplaceSubscriptionStatusUnsubscribed` added to enum type `MarketplaceSubscriptionStatus`
- New value `MonitoringTypeDISCOVERY` added to enum type `MonitoringType`
- New enum type `Action` with values `ActionInstall`, `ActionUninstall`
- New enum type `ManagedServiceIdentityType` with values `ManagedServiceIdentityTypeNone`, `ManagedServiceIdentityTypeSystemAssigned`, `ManagedServiceIdentityTypeSystemAssignedUserAssigned`, `ManagedServiceIdentityTypeUserAssigned`
- New enum type `MarketplaceSaasAutoRenew` with values `MarketplaceSaasAutoRenewOff`, `MarketplaceSaasAutoRenewOn`
- New enum type `Status` with values `StatusActive`, `StatusDeleting`, `StatusFailed`, `StatusInProgress`
- New enum type `SubscriptionListOperation` with values `SubscriptionListOperationActive`, `SubscriptionListOperationAddBegin`, `SubscriptionListOperationAddComplete`, `SubscriptionListOperationDeleteBegin`, `SubscriptionListOperationDeleteComplete`
- New function `*ClientFactory.NewCreationSupportedClient() *CreationSupportedClient`
- New function `*ClientFactory.NewMonitoredSubscriptionsClient() *MonitoredSubscriptionsClient`
- New function `NewCreationSupportedClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CreationSupportedClient, error)`
- New function `*CreationSupportedClient.Get(context.Context, string, *CreationSupportedClientGetOptions) (CreationSupportedClientGetResponse, error)`
- New function `*CreationSupportedClient.List(context.Context, string, *CreationSupportedClientListOptions) (CreationSupportedClientListResponse, error)`
- New function `NewMonitoredSubscriptionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MonitoredSubscriptionsClient, error)`
- New function `*MonitoredSubscriptionsClient.BeginCreateOrUpdate(context.Context, string, string, MonitoredSubscriptionProperties, *MonitoredSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientCreateOrUpdateResponse], error)`
- New function `*MonitoredSubscriptionsClient.BeginDelete(context.Context, string, string, *MonitoredSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[MonitoredSubscriptionsClientDeleteResponse], error)`
- New function `*MonitoredSubscriptionsClient.Get(context.Context, string, string, *MonitoredSubscriptionsClientGetOptions) (MonitoredSubscriptionsClientGetResponse, error)`
- New function `*MonitoredSubscriptionsClient.NewListPager(string, string, *MonitoredSubscriptionsClientListOptions) *runtime.Pager[MonitoredSubscriptionsClientListResponse]`
- New function `*MonitoredSubscriptionsClient.BeginUpdate(context.Context, string, string, MonitoredSubscriptionProperties, *MonitoredSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientUpdateResponse], error)`
- New function `*MonitorsClient.GetAllConnectedResourcesCount(context.Context, MarketplaceSubscriptionIDRequest, *MonitorsClientGetAllConnectedResourcesCountOptions) (MonitorsClientGetAllConnectedResourcesCountResponse, error)`
- New function `*MonitorsClient.UpdateAgentStatus(context.Context, string, string, AgentStatusRequest, *MonitorsClientUpdateAgentStatusOptions) (MonitorsClientUpdateAgentStatusResponse, error)`
- New function `*MonitorsClient.BeginUpgradePlan(context.Context, string, string, UpgradePlanRequest, *MonitorsClientBeginUpgradePlanOptions) (*runtime.Poller[MonitorsClientUpgradePlanResponse], error)`
- New struct `AgentStatus`
- New struct `AgentStatusRequest`
- New struct `ConnectedResourcesCountResponse`
- New struct `CreateResourceSupportedProperties`
- New struct `CreateResourceSupportedResponse`
- New struct `LogStatusRequest`
- New struct `ManagedServiceIdentity`
- New struct `MarketplaceSubscriptionIDRequest`
- New struct `MetricStatusRequest`
- New struct `MonitorUpdateProperties`
- New struct `MonitoredSubscription`
- New struct `MonitoredSubscriptionProperties`
- New struct `MonitoredSubscriptionPropertiesList`
- New struct `SubscriptionList`
- New struct `UpgradePlanRequest`
- New field `CompanyName` in struct `AccountInfo`
- New field `MarketplaceSaaSResourceName` in struct `MarketplaceSaaSResourceDetailsResponse`
- New field `MarketplaceSaasAutoRenew` in struct `MonitorProperties`
- New field `Identity`, `Properties` in struct `MonitorResourceUpdate`
- New field `Request` in struct `MonitorsClientGetMetricStatusOptions`
- New field `Request` in struct `MonitorsClientListMonitoredResourcesOptions`


## 2.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.0.0 (2023-08-25)
### Breaking Changes

- Function `*MonitorsClient.GetAccountCredentials` has been removed
- Function `*TagRulesClient.Update` has been removed
- Struct `AccountInfoSecure` has been removed
- Struct `TagRuleUpdate` has been removed
- Field `DynatraceEnvironmentProperties`, `MarketplaceSubscriptionStatus`, `MonitoringStatus`, `PlanData`, `UserInfo` of struct `MonitorResourceUpdate` has been removed

### Features Added

- New function `*MonitorsClient.GetMarketplaceSaaSResourceDetails(context.Context, MarketplaceSaaSResourceDetailsRequest, *MonitorsClientGetMarketplaceSaaSResourceDetailsOptions) (MonitorsClientGetMarketplaceSaaSResourceDetailsResponse, error)`
- New function `*MonitorsClient.GetMetricStatus(context.Context, string, string, *MonitorsClientGetMetricStatusOptions) (MonitorsClientGetMetricStatusResponse, error)`
- New struct `MarketplaceSaaSResourceDetailsRequest`
- New struct `MarketplaceSaaSResourceDetailsResponse`
- New struct `MetricsStatusResponse`
- New field `SendingMetrics` in struct `MetricRules`


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-09-20)
### Other Changes

- Release stable version.

## 0.1.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/dynatrace/armdynatrace` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.1.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).