# Release History

## 2.0.0 (2022-07-04)
### Breaking Changes

- Function `*PrivateStoreCollectionClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, *PrivateStoreCollectionClientCreateOrUpdateOptions)` to `(context.Context, string, string, Collection, *PrivateStoreCollectionClientCreateOrUpdateOptions)`
- Function `*PrivateStoreClient.UpdateAdminRequestApproval` parameter(s) have been changed from `(context.Context, string, string, *PrivateStoreClientUpdateAdminRequestApprovalOptions)` to `(context.Context, string, string, AdminRequestApprovalsResource, *PrivateStoreClientUpdateAdminRequestApprovalOptions)`
- Function `*PrivateStoreCollectionOfferClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, *PrivateStoreCollectionOfferClientCreateOrUpdateOptions)` to `(context.Context, string, string, string, Offer, *PrivateStoreCollectionOfferClientCreateOrUpdateOptions)`
- Function `*PrivateStoreClient.CreateApprovalRequest` parameter(s) have been changed from `(context.Context, string, string, *PrivateStoreClientCreateApprovalRequestOptions)` to `(context.Context, string, string, RequestApprovalResource, *PrivateStoreClientCreateApprovalRequestOptions)`
- Function `*PrivateStoreClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, *PrivateStoreClientCreateOrUpdateOptions)` to `(context.Context, string, PrivateStore, *PrivateStoreClientCreateOrUpdateOptions)`
- Field `Payload` of struct `PrivateStoreClientCreateOrUpdateOptions` has been removed
- Field `Payload` of struct `PrivateStoreCollectionOfferClientCreateOrUpdateOptions` has been removed
- Field `Payload` of struct `PrivateStoreClientCreateApprovalRequestOptions` has been removed
- Field `Payload` of struct `PrivateStoreCollectionClientCreateOrUpdateOptions` has been removed
- Field `Payload` of struct `PrivateStoreClientUpdateAdminRequestApprovalOptions` has been removed

### Features Added

- New function `*PrivateStoreClient.QueryUserOffers(context.Context, string, *PrivateStoreClientQueryUserOffersOptions) (PrivateStoreClientQueryUserOffersResponse, error)`
- New function `*PrivateStoreCollectionClient.ApproveAllItems(context.Context, string, string, *PrivateStoreCollectionClientApproveAllItemsOptions) (PrivateStoreCollectionClientApproveAllItemsResponse, error)`
- New function `*PrivateStoreCollectionOfferClient.UpsertOfferWithMultiContext(context.Context, string, string, string, *PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions) (PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse, error)`
- New function `*PrivateStoreCollectionClient.DisableApproveAllItems(context.Context, string, string, *PrivateStoreCollectionClientDisableApproveAllItemsOptions) (PrivateStoreCollectionClientDisableApproveAllItemsResponse, error)`
- New function `*PrivateStoreClient.AnyExistingOffersInTheCollections(context.Context, string, *PrivateStoreClientAnyExistingOffersInTheCollectionsOptions) (PrivateStoreClientAnyExistingOffersInTheCollectionsResponse, error)`
- New struct `AnyExistingOffersInTheCollectionsResponse`
- New struct `ContextAndPlansDetails`
- New struct `MultiContextAndPlansPayload`
- New struct `MultiContextAndPlansProperties`
- New struct `PrivateStoreClientAnyExistingOffersInTheCollectionsOptions`
- New struct `PrivateStoreClientAnyExistingOffersInTheCollectionsResponse`
- New struct `PrivateStoreClientQueryUserOffersOptions`
- New struct `PrivateStoreClientQueryUserOffersResponse`
- New struct `PrivateStoreCollectionClientApproveAllItemsOptions`
- New struct `PrivateStoreCollectionClientApproveAllItemsResponse`
- New struct `PrivateStoreCollectionClientDisableApproveAllItemsOptions`
- New struct `PrivateStoreCollectionClientDisableApproveAllItemsResponse`
- New struct `PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions`
- New struct `PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse`
- New struct `QueryUserOffersDetails`
- New struct `QueryUserOffersProperties`
- New field `SubscriptionIDs` in struct `QueryApprovedPlans`
- New field `ID` in struct `SingleOperation`
- New field `ApproveAllItemsModifiedAt` in struct `CollectionProperties`
- New field `ApproveAllItems` in struct `CollectionProperties`
- New field `Icon` in struct `AdminRequestApprovalProperties`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).