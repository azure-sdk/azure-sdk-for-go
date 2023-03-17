# Release History

## 1.1.0-beta.2 (2023-03-17)
### Features Added

- New function `NewReplicasClient(string, azcore.TokenCredential, *arm.ClientOptions) (*ReplicasClient, error)`
- New function `*ReplicasClient.BeginCreateOrUpdate(context.Context, string, string, string, Replica, *ReplicasClientBeginCreateOrUpdateOptions) (*runtime.Poller[ReplicasClientCreateOrUpdateResponse], error)`
- New function `*ReplicasClient.Delete(context.Context, string, string, string, *ReplicasClientDeleteOptions) (ReplicasClientDeleteResponse, error)`
- New function `*ReplicasClient.Get(context.Context, string, string, string, *ReplicasClientGetOptions) (ReplicasClientGetResponse, error)`
- New function `*ReplicasClient.NewListPager(string, string, *ReplicasClientListOptions) *runtime.Pager[ReplicasClientListResponse]`
- New function `*ReplicasClient.BeginRestart(context.Context, string, string, string, *ReplicasClientBeginRestartOptions) (*runtime.Poller[ReplicasClientRestartResponse], error)`
- New function `*ReplicasClient.BeginUpdate(context.Context, string, string, string, Replica, *ReplicasClientBeginUpdateOptions) (*runtime.Poller[ReplicasClientUpdateResponse], error)`
- New struct `Replica`
- New struct `ReplicaList`
- New struct `ReplicaProperties`
- New struct `ReplicasClient`
- New field `SystemData` in struct `PrivateLinkResource`
- New field `SystemData` in struct `ProxyResource`
- New field `SystemData` in struct `Resource`
- New field `SystemData` in struct `TrackedResource`


## 1.1.0-beta.1 (2022-11-03)
### Features Added

- New const `EventListenerFilterDiscriminatorEventName`
- New const `EventListenerEndpointDiscriminatorEventHub`
- New type alias `EventListenerEndpointDiscriminator`
- New type alias `EventListenerFilterDiscriminator`
- New function `PossibleEventListenerFilterDiscriminatorValues() []EventListenerFilterDiscriminator`
- New function `*EventHubEndpoint.GetEventListenerEndpoint() *EventListenerEndpoint`
- New function `*CustomCertificatesClient.BeginCreateOrUpdate(context.Context, string, string, string, CustomCertificate, *CustomCertificatesClientBeginCreateOrUpdateOptions) (*runtime.Poller[CustomCertificatesClientCreateOrUpdateResponse], error)`
- New function `*CustomCertificatesClient.NewListPager(string, string, *CustomCertificatesClientListOptions) *runtime.Pager[CustomCertificatesClientListResponse]`
- New function `*CustomCertificatesClient.Get(context.Context, string, string, string, *CustomCertificatesClientGetOptions) (CustomCertificatesClientGetResponse, error)`
- New function `*CustomDomainsClient.BeginDelete(context.Context, string, string, string, *CustomDomainsClientBeginDeleteOptions) (*runtime.Poller[CustomDomainsClientDeleteResponse], error)`
- New function `*CustomDomainsClient.Get(context.Context, string, string, string, *CustomDomainsClientGetOptions) (CustomDomainsClientGetResponse, error)`
- New function `NewCustomCertificatesClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CustomCertificatesClient, error)`
- New function `PossibleEventListenerEndpointDiscriminatorValues() []EventListenerEndpointDiscriminator`
- New function `*CustomCertificatesClient.Delete(context.Context, string, string, string, *CustomCertificatesClientDeleteOptions) (CustomCertificatesClientDeleteResponse, error)`
- New function `*EventListenerEndpoint.GetEventListenerEndpoint() *EventListenerEndpoint`
- New function `*CustomDomainsClient.NewListPager(string, string, *CustomDomainsClientListOptions) *runtime.Pager[CustomDomainsClientListResponse]`
- New function `*CustomDomainsClient.BeginCreateOrUpdate(context.Context, string, string, string, CustomDomain, *CustomDomainsClientBeginCreateOrUpdateOptions) (*runtime.Poller[CustomDomainsClientCreateOrUpdateResponse], error)`
- New function `*EventListenerFilter.GetEventListenerFilter() *EventListenerFilter`
- New function `NewCustomDomainsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*CustomDomainsClient, error)`
- New function `*EventNameFilter.GetEventListenerFilter() *EventListenerFilter`
- New struct `CustomCertificate`
- New struct `CustomCertificateList`
- New struct `CustomCertificateProperties`
- New struct `CustomCertificatesClient`
- New struct `CustomCertificatesClientBeginCreateOrUpdateOptions`
- New struct `CustomCertificatesClientCreateOrUpdateResponse`
- New struct `CustomCertificatesClientDeleteOptions`
- New struct `CustomCertificatesClientDeleteResponse`
- New struct `CustomCertificatesClientGetOptions`
- New struct `CustomCertificatesClientGetResponse`
- New struct `CustomCertificatesClientListOptions`
- New struct `CustomCertificatesClientListResponse`
- New struct `CustomDomain`
- New struct `CustomDomainList`
- New struct `CustomDomainProperties`
- New struct `CustomDomainsClient`
- New struct `CustomDomainsClientBeginCreateOrUpdateOptions`
- New struct `CustomDomainsClientBeginDeleteOptions`
- New struct `CustomDomainsClientCreateOrUpdateResponse`
- New struct `CustomDomainsClientDeleteResponse`
- New struct `CustomDomainsClientGetOptions`
- New struct `CustomDomainsClientGetResponse`
- New struct `CustomDomainsClientListOptions`
- New struct `CustomDomainsClientListResponse`
- New struct `EventHubEndpoint`
- New struct `EventListener`
- New struct `EventListenerEndpoint`
- New struct `EventListenerFilter`
- New struct `EventNameFilter`
- New struct `ResourceReference`
- New field `EventListeners` in struct `HubProperties`


## 1.0.0 (2022-05-16)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/webpubsub/armwebpubsub` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).