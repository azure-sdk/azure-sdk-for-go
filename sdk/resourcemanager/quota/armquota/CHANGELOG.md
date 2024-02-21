# Release History

## 1.1.0-beta.1 (2024-02-21)
### Features Added

- New enum type `CreatedByType` with values `CreatedByTypeApplication`, `CreatedByTypeKey`, `CreatedByTypeManagedIdentity`, `CreatedByTypeUser`
- New enum type `EnvironmentType` with values `EnvironmentTypeNonProduction`, `EnvironmentTypeProduction`
- New enum type `GroupingIDType` with values `GroupingIDTypeBillingID`, `GroupingIDTypeServiceTreeID`
- New enum type `RequestState` with values `RequestStateAccepted`, `RequestStateCanceled`, `RequestStateCreated`, `RequestStateFailed`, `RequestStateInProgress`, `RequestStateInvalid`, `RequestStateSucceeded`
- New function `*ClientFactory.NewGroupQuotaEnforcementClient() *GroupQuotaEnforcementClient`
- New function `*ClientFactory.NewGroupQuotaLimitsClient() *GroupQuotaLimitsClient`
- New function `*ClientFactory.NewGroupQuotaLimitsRequestsClient() *GroupQuotaLimitsRequestsClient`
- New function `*ClientFactory.NewGroupQuotaSubscriptionQuotaAllocationClient() *GroupQuotaSubscriptionQuotaAllocationClient`
- New function `*ClientFactory.NewGroupQuotaSubscriptionQuotaAllocationRequestsClient() *GroupQuotaSubscriptionQuotaAllocationRequestsClient`
- New function `*ClientFactory.NewGroupQuotaSubscriptionsClient() *GroupQuotaSubscriptionsClient`
- New function `*ClientFactory.NewGroupQuotasClient() *GroupQuotasClient`
- New function `*ClientFactory.NewSubscriptionRequestsClient() *SubscriptionRequestsClient`
- New function `NewGroupQuotaEnforcementClient(azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotaEnforcementClient, error)`
- New function `*GroupQuotaEnforcementClient.BeginCreateOrUpdate(context.Context, string, string, string, string, *GroupQuotaEnforcementClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotaEnforcementClientCreateOrUpdateResponse], error)`
- New function `*GroupQuotaEnforcementClient.Get(context.Context, string, string, string, string, *GroupQuotaEnforcementClientGetOptions) (GroupQuotaEnforcementClientGetResponse, error)`
- New function `*GroupQuotaEnforcementClient.NewListPager(string, string, string, *GroupQuotaEnforcementClientListOptions) *runtime.Pager[GroupQuotaEnforcementClientListResponse]`
- New function `*GroupQuotaEnforcementClient.BeginUpdate(context.Context, string, string, string, string, *GroupQuotaEnforcementClientBeginUpdateOptions) (*runtime.Poller[GroupQuotaEnforcementClientUpdateResponse], error)`
- New function `NewGroupQuotaLimitsClient(azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotaLimitsClient, error)`
- New function `*GroupQuotaLimitsClient.BeginCreateOrUpdate(context.Context, string, string, string, *GroupQuotaLimitsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotaLimitsClientCreateOrUpdateResponse], error)`
- New function `*GroupQuotaLimitsClient.Get(context.Context, string, string, string, string, *GroupQuotaLimitsClientGetOptions) (GroupQuotaLimitsClientGetResponse, error)`
- New function `*GroupQuotaLimitsClient.NewListPager(string, string, string, *GroupQuotaLimitsClientListOptions) *runtime.Pager[GroupQuotaLimitsClientListResponse]`
- New function `*GroupQuotaLimitsClient.BeginUpdate(context.Context, string, string, string, *GroupQuotaLimitsClientBeginUpdateOptions) (*runtime.Poller[GroupQuotaLimitsClientUpdateResponse], error)`
- New function `NewGroupQuotaLimitsRequestsClient(azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotaLimitsRequestsClient, error)`
- New function `*GroupQuotaLimitsRequestsClient.Get(context.Context, string, string, string, string, *GroupQuotaLimitsRequestsClientGetOptions) (GroupQuotaLimitsRequestsClientGetResponse, error)`
- New function `*GroupQuotaLimitsRequestsClient.NewListPager(string, string, string, *GroupQuotaLimitsRequestsClientListOptions) *runtime.Pager[GroupQuotaLimitsRequestsClientListResponse]`
- New function `NewGroupQuotaSubscriptionQuotaAllocationClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotaSubscriptionQuotaAllocationClient, error)`
- New function `*GroupQuotaSubscriptionQuotaAllocationClient.BeginCreateOrUpdate(context.Context, string, string, string, SubscriptionQuotaAllocations, *GroupQuotaSubscriptionQuotaAllocationClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotaSubscriptionQuotaAllocationClientCreateOrUpdateResponse], error)`
- New function `*GroupQuotaSubscriptionQuotaAllocationClient.Get(context.Context, string, string, string, string, *GroupQuotaSubscriptionQuotaAllocationClientGetOptions) (GroupQuotaSubscriptionQuotaAllocationClientGetResponse, error)`
- New function `*GroupQuotaSubscriptionQuotaAllocationClient.NewListPager(string, string, string, *GroupQuotaSubscriptionQuotaAllocationClientListOptions) *runtime.Pager[GroupQuotaSubscriptionQuotaAllocationClientListResponse]`
- New function `*GroupQuotaSubscriptionQuotaAllocationClient.BeginUpdate(context.Context, string, string, string, SubscriptionQuotaAllocations, *GroupQuotaSubscriptionQuotaAllocationClientBeginUpdateOptions) (*runtime.Poller[GroupQuotaSubscriptionQuotaAllocationClientUpdateResponse], error)`
- New function `NewGroupQuotaSubscriptionQuotaAllocationRequestsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotaSubscriptionQuotaAllocationRequestsClient, error)`
- New function `*GroupQuotaSubscriptionQuotaAllocationRequestsClient.Get(context.Context, string, string, string, *GroupQuotaSubscriptionQuotaAllocationRequestsClientGetOptions) (GroupQuotaSubscriptionQuotaAllocationRequestsClientGetResponse, error)`
- New function `*GroupQuotaSubscriptionQuotaAllocationRequestsClient.NewListPager(string, string, string, *GroupQuotaSubscriptionQuotaAllocationRequestsClientListOptions) *runtime.Pager[GroupQuotaSubscriptionQuotaAllocationRequestsClientListResponse]`
- New function `NewGroupQuotaSubscriptionsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotaSubscriptionsClient, error)`
- New function `*GroupQuotaSubscriptionsClient.BeginCreateOrUpdate(context.Context, string, string, *GroupQuotaSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotaSubscriptionsClientCreateOrUpdateResponse], error)`
- New function `*GroupQuotaSubscriptionsClient.Delete(context.Context, string, string, *GroupQuotaSubscriptionsClientDeleteOptions) (GroupQuotaSubscriptionsClientDeleteResponse, error)`
- New function `*GroupQuotaSubscriptionsClient.Get(context.Context, string, string, *GroupQuotaSubscriptionsClientGetOptions) (GroupQuotaSubscriptionsClientGetResponse, error)`
- New function `*GroupQuotaSubscriptionsClient.NewListPager(string, string, *GroupQuotaSubscriptionsClientListOptions) *runtime.Pager[GroupQuotaSubscriptionsClientListResponse]`
- New function `*GroupQuotaSubscriptionsClient.BeginUpdate(context.Context, string, string, *GroupQuotaSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[GroupQuotaSubscriptionsClientUpdateResponse], error)`
- New function `NewGroupQuotasClient(azcore.TokenCredential, *arm.ClientOptions) (*GroupQuotasClient, error)`
- New function `*GroupQuotasClient.BeginCreateOrUpdate(context.Context, string, string, *GroupQuotasClientBeginCreateOrUpdateOptions) (*runtime.Poller[GroupQuotasClientCreateOrUpdateResponse], error)`
- New function `*GroupQuotasClient.Delete(context.Context, string, string, *GroupQuotasClientDeleteOptions) (GroupQuotasClientDeleteResponse, error)`
- New function `*GroupQuotasClient.Get(context.Context, string, string, *GroupQuotasClientGetOptions) (GroupQuotasClientGetResponse, error)`
- New function `*GroupQuotasClient.NewListPager(string, *GroupQuotasClientListOptions) *runtime.Pager[GroupQuotasClientListResponse]`
- New function `*GroupQuotasClient.BeginUpdate(context.Context, string, string, *GroupQuotasClientBeginUpdateOptions) (*runtime.Poller[GroupQuotasClientUpdateResponse], error)`
- New function `NewSubscriptionRequestsClient(azcore.TokenCredential, *arm.ClientOptions) (*SubscriptionRequestsClient, error)`
- New function `*SubscriptionRequestsClient.Get(context.Context, string, string, string, *SubscriptionRequestsClientGetOptions) (SubscriptionRequestsClientGetResponse, error)`
- New function `*SubscriptionRequestsClient.NewListPager(string, string, *SubscriptionRequestsClientListOptions) *runtime.Pager[SubscriptionRequestsClientListResponse]`
- New struct `AdditionalAttributes`
- New struct `AllocatedToSubscription`
- New struct `AllocationRequestBase`
- New struct `AllocationRequestBaseProperties`
- New struct `AllocationRequestBasePropertiesName`
- New struct `AllocationRequestStatus`
- New struct `AllocationRequestStatusList`
- New struct `AllocationRequestStatusProperties`
- New struct `BillingAccountID`
- New struct `ErrorAdditionalInfo`
- New struct `ErrorDetail`
- New struct `ErrorResponse`
- New struct `GroupQuotaDetails`
- New struct `GroupQuotaDetailsName`
- New struct `GroupQuotaLimit`
- New struct `GroupQuotaLimitList`
- New struct `GroupQuotaList`
- New struct `GroupQuotaSubscriptionID`
- New struct `GroupQuotaSubscriptionIDList`
- New struct `GroupQuotaSubscriptionIDProperties`
- New struct `GroupQuotaSubscriptionRequestStatus`
- New struct `GroupQuotaSubscriptionRequestStatusList`
- New struct `GroupQuotaSubscriptionRequestStatusProperties`
- New struct `GroupQuotasEnforcementListResponse`
- New struct `GroupQuotasEnforcementResponse`
- New struct `GroupQuotasEnforcementResponseProperties`
- New struct `GroupQuotasEntity`
- New struct `GroupQuotasEntityBase`
- New struct `GroupingID`
- New struct `LROResponse`
- New struct `LROResponseProperties`
- New struct `ProxyResource`
- New struct `Resource`
- New struct `ResourceBaseRequest`
- New struct `SubmittedResourceRequestStatus`
- New struct `SubmittedResourceRequestStatusList`
- New struct `SubscriptionGroupQuotaAssignment`
- New struct `SubscriptionQuotaAllocationRequestList`
- New struct `SubscriptionQuotaAllocations`
- New struct `SubscriptionQuotaAllocationsList`
- New struct `SubscriptionQuotaDetails`
- New struct `SubscriptionQuotaDetailsName`
- New struct `SystemData`


## 1.0.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.

### Other Changes

- Release stable version.


## 0.6.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 0.6.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 0.5.0 (2022-05-17)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quota/armquota` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 0.5.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).
