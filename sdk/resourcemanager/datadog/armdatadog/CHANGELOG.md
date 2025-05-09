# Release History

## 2.0.0 (2025-05-09)
### Breaking Changes

- Function `*MonitoredSubscriptionsClient.BeginCreateorUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginCreateorUpdateOptions)` to `(context.Context, string, string, string, MonitoredSubscriptionProperties, *MonitoredSubscriptionsClientBeginCreateorUpdateOptions)`
- Function `*MonitoredSubscriptionsClient.BeginUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginUpdateOptions)` to `(context.Context, string, string, string, MonitoredSubscriptionProperties, *MonitoredSubscriptionsClientBeginUpdateOptions)`
- Function `*TagRulesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *TagRulesClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, MonitoringTagRules, *TagRulesClientCreateOrUpdateOptions)`
- Function `*ClientFactory.NewCreationSupportedClient` has been removed
- Function `*ClientFactory.NewMarketplaceAgreementsClient` has been removed
- Function `*ClientFactory.NewMonitorsClient` has been removed
- Function `*ClientFactory.NewSingleSignOnConfigurationsClient` has been removed
- Function `NewCreationSupportedClient` has been removed
- Function `*CreationSupportedClient.Get` has been removed
- Function `*CreationSupportedClient.NewListPager` has been removed
- Function `NewMarketplaceAgreementsClient` has been removed
- Function `*MarketplaceAgreementsClient.CreateOrUpdate` has been removed
- Function `*MarketplaceAgreementsClient.NewListPager` has been removed
- Function `NewMonitorsClient` has been removed
- Function `*MonitorsClient.BeginCreate` has been removed
- Function `*MonitorsClient.BeginDelete` has been removed
- Function `*MonitorsClient.Get` has been removed
- Function `*MonitorsClient.GetDefaultKey` has been removed
- Function `*MonitorsClient.NewListAPIKeysPager` has been removed
- Function `*MonitorsClient.NewListByResourceGroupPager` has been removed
- Function `*MonitorsClient.NewListHostsPager` has been removed
- Function `*MonitorsClient.NewListLinkedResourcesPager` has been removed
- Function `*MonitorsClient.NewListMonitoredResourcesPager` has been removed
- Function `*MonitorsClient.NewListPager` has been removed
- Function `*MonitorsClient.RefreshSetPasswordLink` has been removed
- Function `*MonitorsClient.SetDefaultKey` has been removed
- Function `*MonitorsClient.BeginUpdate` has been removed
- Function `NewSingleSignOnConfigurationsClient` has been removed
- Function `*SingleSignOnConfigurationsClient.BeginCreateOrUpdate` has been removed
- Function `*SingleSignOnConfigurationsClient.Get` has been removed
- Function `*SingleSignOnConfigurationsClient.NewListPager` has been removed
- Field `Body` of struct `MonitoredSubscriptionsClientBeginCreateorUpdateOptions` has been removed
- Field `Body` of struct `MonitoredSubscriptionsClientBeginUpdateOptions` has been removed
- Field `MonitoredSubscriptionProperties` of struct `MonitoredSubscriptionsClientCreateorUpdateResponse` has been removed
- Field `Body` of struct `TagRulesClientCreateOrUpdateOptions` has been removed

### Features Added

- New function `*ClientFactory.NewCreationSupportedOperationGroupClient() *CreationSupportedOperationGroupClient`
- New function `*ClientFactory.NewMarketplaceAgreementsOperationGroupClient() *MarketplaceAgreementsOperationGroupClient`
- New function `*ClientFactory.NewMonitorResourcesClient() *MonitorResourcesClient`
- New function `*ClientFactory.NewSingleSignOnResourcesClient() *SingleSignOnResourcesClient`
- New function `NewCreationSupportedOperationGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CreationSupportedOperationGroupClient, error)`
- New function `*CreationSupportedOperationGroupClient.Get(context.Context, string, *CreationSupportedOperationGroupClientGetOptions) (CreationSupportedOperationGroupClientGetResponse, error)`
- New function `*CreationSupportedOperationGroupClient.NewListPager(string, *CreationSupportedOperationGroupClientListOptions) *runtime.Pager[CreationSupportedOperationGroupClientListResponse]`
- New function `NewMarketplaceAgreementsOperationGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MarketplaceAgreementsOperationGroupClient, error)`
- New function `*MarketplaceAgreementsOperationGroupClient.CreateOrUpdate(context.Context, AgreementResource, *MarketplaceAgreementsOperationGroupClientCreateOrUpdateOptions) (MarketplaceAgreementsOperationGroupClientCreateOrUpdateResponse, error)`
- New function `*MarketplaceAgreementsOperationGroupClient.NewListPager(*MarketplaceAgreementsOperationGroupClientListOptions) *runtime.Pager[MarketplaceAgreementsOperationGroupClientListResponse]`
- New function `NewMonitorResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MonitorResourcesClient, error)`
- New function `*MonitorResourcesClient.BillingInfoGet(context.Context, string, string, *MonitorResourcesClientBillingInfoGetOptions) (MonitorResourcesClientBillingInfoGetResponse, error)`
- New function `*MonitorResourcesClient.BeginCreate(context.Context, string, string, MonitorResource, *MonitorResourcesClientBeginCreateOptions) (*runtime.Poller[MonitorResourcesClientCreateResponse], error)`
- New function `*MonitorResourcesClient.BeginDelete(context.Context, string, string, *MonitorResourcesClientBeginDeleteOptions) (*runtime.Poller[MonitorResourcesClientDeleteResponse], error)`
- New function `*MonitorResourcesClient.Get(context.Context, string, string, *MonitorResourcesClientGetOptions) (MonitorResourcesClientGetResponse, error)`
- New function `*MonitorResourcesClient.GetDefaultKey(context.Context, string, string, *MonitorResourcesClientGetDefaultKeyOptions) (MonitorResourcesClientGetDefaultKeyResponse, error)`
- New function `*MonitorResourcesClient.NewListAPIKeysPager(string, string, *MonitorResourcesClientListAPIKeysOptions) *runtime.Pager[MonitorResourcesClientListAPIKeysResponse]`
- New function `*MonitorResourcesClient.NewListByResourceGroupPager(string, *MonitorResourcesClientListByResourceGroupOptions) *runtime.Pager[MonitorResourcesClientListByResourceGroupResponse]`
- New function `*MonitorResourcesClient.NewListHostsPager(string, string, *MonitorResourcesClientListHostsOptions) *runtime.Pager[MonitorResourcesClientListHostsResponse]`
- New function `*MonitorResourcesClient.NewListLinkedResourcesPager(string, string, *MonitorResourcesClientListLinkedResourcesOptions) *runtime.Pager[MonitorResourcesClientListLinkedResourcesResponse]`
- New function `*MonitorResourcesClient.NewListMonitoredResourcesPager(string, string, *MonitorResourcesClientListMonitoredResourcesOptions) *runtime.Pager[MonitorResourcesClientListMonitoredResourcesResponse]`
- New function `*MonitorResourcesClient.NewListPager(*MonitorResourcesClientListOptions) *runtime.Pager[MonitorResourcesClientListResponse]`
- New function `*MonitorResourcesClient.RefreshSetPasswordLink(context.Context, string, string, *MonitorResourcesClientRefreshSetPasswordLinkOptions) (MonitorResourcesClientRefreshSetPasswordLinkResponse, error)`
- New function `*MonitorResourcesClient.SetDefaultKey(context.Context, string, string, *MonitorResourcesClientSetDefaultKeyOptions) (MonitorResourcesClientSetDefaultKeyResponse, error)`
- New function `*MonitorResourcesClient.BeginUpdate(context.Context, string, string, MonitorResourceUpdateParameters, *MonitorResourcesClientBeginUpdateOptions) (*runtime.Poller[MonitorResourcesClientUpdateResponse], error)`
- New function `NewSingleSignOnResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*SingleSignOnResourcesClient, error)`
- New function `*SingleSignOnResourcesClient.BeginCreateOrUpdate(context.Context, string, string, string, SingleSignOnResource, *SingleSignOnResourcesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SingleSignOnResourcesClientCreateOrUpdateResponse], error)`
- New function `*SingleSignOnResourcesClient.Get(context.Context, string, string, string, *SingleSignOnResourcesClientGetOptions) (SingleSignOnResourcesClientGetResponse, error)`
- New function `*SingleSignOnResourcesClient.NewListPager(string, string, *SingleSignOnResourcesClientListOptions) *runtime.Pager[SingleSignOnResourcesClientListResponse]`
- New struct `BillingInfoResponse`
- New struct `MarketplaceSaaSInfo`
- New struct `PartnerBillingEntity`
- New struct `SetDefaultKeyParameterBody`
- New field `NextLink` in struct `CreateResourceSupportedResponseList`
- New field `Location` in struct `LinkedResource`
- New field `SystemData` in struct `MonitoredSubscriptionProperties`
- New field `NextLink` in struct `MonitoredSubscriptionPropertiesList`
- New field `CustomMetrics` in struct `MonitoringTagRulesProperties`


## 1.3.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.2.0 (2023-10-27)
### Features Added

- New enum type `Operation` with values `OperationActive`, `OperationAddBegin`, `OperationAddComplete`, `OperationDeleteBegin`, `OperationDeleteComplete`
- New enum type `Status` with values `StatusActive`, `StatusDeleting`, `StatusFailed`, `StatusInProgress`
- New function `*ClientFactory.NewCreationSupportedClient() *CreationSupportedClient`
- New function `*ClientFactory.NewMonitoredSubscriptionsClient() *MonitoredSubscriptionsClient`
- New function `NewCreationSupportedClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CreationSupportedClient, error)`
- New function `*CreationSupportedClient.Get(context.Context, string, *CreationSupportedClientGetOptions) (CreationSupportedClientGetResponse, error)`
- New function `*CreationSupportedClient.NewListPager(string, *CreationSupportedClientListOptions) *runtime.Pager[CreationSupportedClientListResponse]`
- New function `NewMonitoredSubscriptionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*MonitoredSubscriptionsClient, error)`
- New function `*MonitoredSubscriptionsClient.BeginCreateorUpdate(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginCreateorUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientCreateorUpdateResponse], error)`
- New function `*MonitoredSubscriptionsClient.BeginDelete(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[MonitoredSubscriptionsClientDeleteResponse], error)`
- New function `*MonitoredSubscriptionsClient.Get(context.Context, string, string, string, *MonitoredSubscriptionsClientGetOptions) (MonitoredSubscriptionsClientGetResponse, error)`
- New function `*MonitoredSubscriptionsClient.NewListPager(string, string, *MonitoredSubscriptionsClientListOptions) *runtime.Pager[MonitoredSubscriptionsClientListResponse]`
- New function `*MonitoredSubscriptionsClient.BeginUpdate(context.Context, string, string, string, *MonitoredSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientUpdateResponse], error)`
- New struct `CreateResourceSupportedProperties`
- New struct `CreateResourceSupportedResponse`
- New struct `CreateResourceSupportedResponseList`
- New struct `MonitoredSubscription`
- New struct `MonitoredSubscriptionProperties`
- New struct `MonitoredSubscriptionPropertiesList`
- New struct `SubscriptionList`
- New field `Cspm` in struct `MonitorUpdateProperties`
- New field `Automuting` in struct `MonitoringTagRulesProperties`
- New field `Cspm` in struct `OrganizationProperties`


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-28)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/datadog/armdatadog` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).