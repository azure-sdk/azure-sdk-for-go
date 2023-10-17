# Release History

## 1.1.0-beta.1 (2023-10-17)
### Features Added

- New function `NewBillingInfoClient(string, azcore.TokenCredential, *arm.ClientOptions) (*BillingInfoClient, error)`
- New function `*BillingInfoClient.Get(context.Context, string, string, *BillingInfoClientGetOptions) (BillingInfoClientGetResponse, error)`
- New function `*ClientFactory.NewBillingInfoClient() *BillingInfoClient`
- New function `*ClientFactory.NewConnectedPartnerResourcesClient() *ConnectedPartnerResourcesClient`
- New function `NewConnectedPartnerResourcesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ConnectedPartnerResourcesClient, error)`
- New function `*ConnectedPartnerResourcesClient.NewListPager(string, string, *ConnectedPartnerResourcesClientListOptions) *runtime.Pager[ConnectedPartnerResourcesClientListResponse]`
- New struct `BillingInfoResponse`
- New struct `ConnectedPartnerResourceProperties`
- New struct `ConnectedPartnerResourcesListFormat`
- New struct `ConnectedPartnerResourcesListResponse`
- New struct `MarketplaceSaaSInfo`
- New struct `PartnerBillingEntity`


## 1.0.0 (2023-05-26)
### Features Added

- New anonymous field `NewRelicMonitorResource` in struct `MonitorsClientSwitchBillingResponse`
- New field `RetryAfter` in struct `MonitorsClientSwitchBillingResponse`


## 0.1.0 (2023-03-24)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/newrelic/armnewrelicobservability` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).