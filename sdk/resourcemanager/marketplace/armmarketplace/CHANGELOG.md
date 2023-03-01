# Release History

## 1.1.0 (2023-03-01)
### Features Added

- New type alias `RuleType` with values `RuleTypePrivateProducts`, `RuleTypeTermsAndCondition`
- New function `*PrivateStoreClient.AnyExistingOffersInTheCollections(context.Context, string, *PrivateStoreClientAnyExistingOffersInTheCollectionsOptions) (PrivateStoreClientAnyExistingOffersInTheCollectionsResponse, error)`
- New function `*PrivateStoreClient.QueryUserOffers(context.Context, string, *PrivateStoreClientQueryUserOffersOptions) (PrivateStoreClientQueryUserOffersResponse, error)`
- New function `*PrivateStoreCollectionClient.ApproveAllItems(context.Context, string, string, *PrivateStoreCollectionClientApproveAllItemsOptions) (PrivateStoreCollectionClientApproveAllItemsResponse, error)`
- New function `*PrivateStoreCollectionClient.DisableApproveAllItems(context.Context, string, string, *PrivateStoreCollectionClientDisableApproveAllItemsOptions) (PrivateStoreCollectionClientDisableApproveAllItemsResponse, error)`
- New function `*PrivateStoreCollectionOfferClient.NewListByContextsPager(string, string, *PrivateStoreCollectionOfferClientListByContextsOptions) *runtime.Pager[PrivateStoreCollectionOfferClientListByContextsResponse]`
- New function `*PrivateStoreCollectionOfferClient.UpsertOfferWithMultiContext(context.Context, string, string, string, *PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextOptions) (PrivateStoreCollectionOfferClientUpsertOfferWithMultiContextResponse, error)`
- New function `NewRPServiceClient(azcore.TokenCredential, *arm.ClientOptions) (*RPServiceClient, error)`
- New function `*RPServiceClient.QueryRules(context.Context, string, string, *RPServiceClientQueryRulesOptions) (RPServiceClientQueryRulesResponse, error)`
- New function `*RPServiceClient.QueryUserRules(context.Context, string, *RPServiceClientQueryUserRulesOptions) (RPServiceClientQueryUserRulesResponse, error)`
- New function `*RPServiceClient.SetCollectionRules(context.Context, string, string, *RPServiceClientSetCollectionRulesOptions) (RPServiceClientSetCollectionRulesResponse, error)`
- New struct `AnyExistingOffersInTheCollectionsResponse`
- New struct `CollectionOffersByAllContextsPayload`
- New struct `CollectionOffersByAllContextsProperties`
- New struct `CollectionOffersByContext`
- New struct `CollectionOffersByContextList`
- New struct `CollectionOffersByContextOffers`
- New struct `ContextAndPlansDetails`
- New struct `MultiContextAndPlansPayload`
- New struct `MultiContextAndPlansProperties`
- New struct `QueryUserOffersDetails`
- New struct `QueryUserOffersProperties`
- New struct `QueryUserRulesDetails`
- New struct `QueryUserRulesProperties`
- New struct `RPServiceClient`
- New struct `Rule`
- New struct `RuleListResponse`
- New struct `SetRulesRequest`
- New field `Icon` in struct `AdminRequestApprovalProperties`
- New field `AppliedRules` in struct `CollectionProperties`
- New field `ApproveAllItems` in struct `CollectionProperties`
- New field `ApproveAllItemsModifiedAt` in struct `CollectionProperties`
- New field `SubscriptionIDs` in struct `QueryApprovedPlans`
- New field `ID` in struct `SingleOperation`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/marketplace/armmarketplace` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).