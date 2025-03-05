# Release History

## 2.0.0-beta.1 (2025-03-05)
### Breaking Changes

- Type of `OperationListResult.Value` has been changed from `[]*SingleOperation` to `[]*Operation`
- Enum `Accessibility` has been removed
- Enum `AdminAction` has been removed
- Enum `Availability` has been removed
- Enum `IdentityType` has been removed
- Enum `Operation` has been removed
- Enum `Status` has been removed
- Enum `SubscriptionState` has been removed
- Function `*ClientFactory.NewPrivateStoreClient` has been removed
- Function `*ClientFactory.NewPrivateStoreCollectionClient` has been removed
- Function `*ClientFactory.NewPrivateStoreCollectionOfferClient` has been removed
- Function `NewPrivateStoreClient` has been removed
- Function `*PrivateStoreClient.AcknowledgeOfferNotification` has been removed
- Function `*PrivateStoreClient.AdminRequestApprovalsList` has been removed
- Function `*PrivateStoreClient.BillingAccounts` has been removed
- Function `*PrivateStoreClient.BulkCollectionsAction` has been removed
- Function `*PrivateStoreClient.CollectionsToSubscriptionsMapping` has been removed
- Function `*PrivateStoreClient.CreateApprovalRequest` has been removed
- Function `*PrivateStoreClient.CreateOrUpdate` has been removed
- Function `*PrivateStoreClient.Delete` has been removed
- Function `*PrivateStoreClient.FetchAllSubscriptionsInTenant` has been removed
- Function `*PrivateStoreClient.Get` has been removed
- Function `*PrivateStoreClient.GetAdminRequestApproval` has been removed
- Function `*PrivateStoreClient.GetApprovalRequestsList` has been removed
- Function `*PrivateStoreClient.GetRequestApproval` has been removed
- Function `*PrivateStoreClient.ListNewPlansNotifications` has been removed
- Function `*PrivateStoreClient.NewListPager` has been removed
- Function `*PrivateStoreClient.ListStopSellOffersPlansNotifications` has been removed
- Function `*PrivateStoreClient.ListSubscriptionsContext` has been removed
- Function `*PrivateStoreClient.QueryApprovedPlans` has been removed
- Function `*PrivateStoreClient.QueryNotificationsState` has been removed
- Function `*PrivateStoreClient.QueryOffers` has been removed
- Function `*PrivateStoreClient.QueryRequestApproval` has been removed
- Function `*PrivateStoreClient.UpdateAdminRequestApproval` has been removed
- Function `*PrivateStoreClient.WithdrawPlan` has been removed
- Function `NewPrivateStoreCollectionClient` has been removed
- Function `*PrivateStoreCollectionClient.CreateOrUpdate` has been removed
- Function `*PrivateStoreCollectionClient.Delete` has been removed
- Function `*PrivateStoreCollectionClient.Get` has been removed
- Function `*PrivateStoreCollectionClient.List` has been removed
- Function `*PrivateStoreCollectionClient.Post` has been removed
- Function `*PrivateStoreCollectionClient.TransferOffers` has been removed
- Function `NewPrivateStoreCollectionOfferClient` has been removed
- Function `*PrivateStoreCollectionOfferClient.CreateOrUpdate` has been removed
- Function `*PrivateStoreCollectionOfferClient.Delete` has been removed
- Function `*PrivateStoreCollectionOfferClient.Get` has been removed
- Function `*PrivateStoreCollectionOfferClient.NewListPager` has been removed
- Function `*PrivateStoreCollectionOfferClient.Post` has been removed
- Operation `*OperationsClient.List` has supported pagination, use `*OperationsClient.NewListPager` instead.
- Struct `AcknowledgeOfferNotificationDetails` has been removed
- Struct `AcknowledgeOfferNotificationProperties` has been removed
- Struct `AdminRequestApprovalProperties` has been removed
- Struct `AdminRequestApprovalsList` has been removed
- Struct `AdminRequestApprovalsResource` has been removed
- Struct `BillingAccountsResponse` has been removed
- Struct `BulkCollectionsDetails` has been removed
- Struct `BulkCollectionsPayload` has been removed
- Struct `BulkCollectionsResponse` has been removed
- Struct `Collection` has been removed
- Struct `CollectionProperties` has been removed
- Struct `CollectionsDetails` has been removed
- Struct `CollectionsList` has been removed
- Struct `CollectionsSubscriptionsMappingDetails` has been removed
- Struct `CollectionsToSubscriptionsMappingPayload` has been removed
- Struct `CollectionsToSubscriptionsMappingProperties` has been removed
- Struct `CollectionsToSubscriptionsMappingResponse` has been removed
- Struct `ErrorResponse` has been removed
- Struct `ErrorResponseError` has been removed
- Struct `NewNotifications` has been removed
- Struct `NewPlansNotificationsList` has been removed
- Struct `NotificationsSettingsProperties` has been removed
- Struct `Offer` has been removed
- Struct `OfferListResponse` has been removed
- Struct `OfferProperties` has been removed
- Struct `Plan` has been removed
- Struct `PlanDetails` has been removed
- Struct `PlanNotificationDetails` has been removed
- Struct `PlanRequesterDetails` has been removed
- Struct `PrivateStore` has been removed
- Struct `PrivateStoreList` has been removed
- Struct `PrivateStoreNotificationsState` has been removed
- Struct `PrivateStoreProperties` has been removed
- Struct `QueryApprovedPlans` has been removed
- Struct `QueryApprovedPlansDetails` has been removed
- Struct `QueryApprovedPlansPayload` has been removed
- Struct `QueryApprovedPlansResponse` has been removed
- Struct `QueryOffers` has been removed
- Struct `QueryRequestApproval` has been removed
- Struct `QueryRequestApprovalProperties` has been removed
- Struct `Recipient` has been removed
- Struct `RequestApprovalProperties` has been removed
- Struct `RequestApprovalResource` has been removed
- Struct `RequestApprovalsDetails` has been removed
- Struct `RequestApprovalsList` has been removed
- Struct `RequestDetails` has been removed
- Struct `Resource` has been removed
- Struct `SingleOperation` has been removed
- Struct `SingleOperationDisplay` has been removed
- Struct `StopSellNotifications` has been removed
- Struct `StopSellOffersPlansNotificationsList` has been removed
- Struct `StopSellOffersPlansNotificationsListProperties` has been removed
- Struct `StopSellSubscriptions` has been removed
- Struct `Subscription` has been removed
- Struct `SubscriptionsContextList` has been removed
- Struct `SubscriptionsResponse` has been removed
- Struct `SystemData` has been removed
- Struct `TransferOffersDetails` has been removed
- Struct `TransferOffersProperties` has been removed
- Struct `TransferOffersResponse` has been removed
- Struct `UserRequestDetails` has been removed
- Struct `WithdrawDetails` has been removed
- Struct `WithdrawProperties` has been removed

### Features Added

- New enum type `ActionType` with values `ActionTypeInternal`
- New enum type `Origin` with values `OriginSystem`, `OriginUser`, `OriginUserSystem`
- New function `*ClientFactory.NewRatingAndReviewsOperationsClient() *RatingAndReviewsOperationsClient`
- New function `NewRatingAndReviewsOperationsClient(azcore.TokenCredential, *arm.ClientOptions) (*RatingAndReviewsOperationsClient, error)`
- New function `*RatingAndReviewsOperationsClient.GetUserHasReview(context.Context, string, *RatingAndReviewsOperationsClientGetUserHasReviewOptions) (RatingAndReviewsOperationsClientGetUserHasReviewResponse, error)`
- New struct `Operation`
- New struct `OperationDisplay`
- New struct `UserHasReview`


## 1.2.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 1.1.1 (2023-04-14)
### Bug Fixes

- Fix serialization bug of empty value of `any` type.


## 1.1.0 (2023-03-31)
### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).