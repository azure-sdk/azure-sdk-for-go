# Release History

## 3.0.0-beta.1 (2025-05-12)
### Breaking Changes

- Function `NewClientFactory` parameter(s) have been changed from `(azcore.TokenCredential, *arm.ClientOptions)` to `(string, azcore.TokenCredential, *arm.ClientOptions)`
- Type of `PurchaseRequest.SKU` has been changed from `*SKU` to `*ResourceSKU`
- Type of `ReservationOrderAliasRequest.SKU` has been changed from `*SKU` to `*ResourceSKU`
- Type of `ReservationOrderAliasResponse.SKU` has been changed from `*SKU` to `*ResourceSKU`
- Type of `SavingsPlanModel.SKU` has been changed from `*SKU` to `*ResourceSKU`
- Type of `SavingsPlanOrderAliasModel.SKU` has been changed from `*SKU` to `*ResourceSKU`
- Type of `SavingsPlanOrderModel.SKU` has been changed from `*SKU` to `*ResourceSKU`
- Function `*ClientFactory.NewRPClient` has been removed
- Function `*ClientFactory.NewReservationOrderAliasClient` has been removed
- Function `*ClientFactory.NewSavingsPlanClient` has been removed
- Function `*ClientFactory.NewSavingsPlanOrderAliasClient` has been removed
- Function `*ClientFactory.NewSavingsPlanOrderClient` has been removed
- Function `NewRPClient` has been removed
- Function `*RPClient.ValidatePurchase` has been removed
- Function `NewReservationOrderAliasClient` has been removed
- Function `*ReservationOrderAliasClient.BeginCreate` has been removed
- Function `*ReservationOrderAliasClient.Get` has been removed
- Function `NewSavingsPlanClient` has been removed
- Function `*SavingsPlanClient.Get` has been removed
- Function `*SavingsPlanClient.NewListAllPager` has been removed
- Function `*SavingsPlanClient.NewListPager` has been removed
- Function `*SavingsPlanClient.Update` has been removed
- Function `*SavingsPlanClient.ValidateUpdate` has been removed
- Function `NewSavingsPlanOrderAliasClient` has been removed
- Function `*SavingsPlanOrderAliasClient.BeginCreate` has been removed
- Function `*SavingsPlanOrderAliasClient.Get` has been removed
- Function `NewSavingsPlanOrderClient` has been removed
- Function `*SavingsPlanOrderClient.Elevate` has been removed
- Function `*SavingsPlanOrderClient.Get` has been removed
- Function `*SavingsPlanOrderClient.NewListPager` has been removed
- Struct `SKU` has been removed

### Features Added

- New value `CommitmentGrainFullTerm`, `CommitmentGrainUnknown` added to enum type `CommitmentGrain`
- New enum type `ApplyDiscountOn` with values `ApplyDiscountOnConsume`, `ApplyDiscountOnPurchase`, `ApplyDiscountOnRenew`
- New enum type `DiscountAppliedScopeType` with values `DiscountAppliedScopeTypeBillingAccount`, `DiscountAppliedScopeTypeBillingProfile`, `DiscountAppliedScopeTypeCustomer`
- New enum type `DiscountCombinationRule` with values `DiscountCombinationRuleBestOf`, `DiscountCombinationRuleStackable`
- New enum type `DiscountEntityType` with values `DiscountEntityTypeAffiliate`, `DiscountEntityTypePrimary`
- New enum type `DiscountProvisioningState` with values `DiscountProvisioningStateCanceled`, `DiscountProvisioningStateFailed`, `DiscountProvisioningStatePending`, `DiscountProvisioningStateSucceeded`, `DiscountProvisioningStateUnknown`
- New enum type `DiscountStatus` with values `DiscountStatusActive`, `DiscountStatusCanceled`, `DiscountStatusExpired`, `DiscountStatusFailed`, `DiscountStatusPending`
- New enum type `DiscountType` with values `DiscountTypeCustomPrice`, `DiscountTypeCustomPriceMultiCurrency`, `DiscountTypeProduct`, `DiscountTypeProductFamily`, `DiscountTypeSKU`
- New enum type `PricingPolicy` with values `PricingPolicyLocked`, `PricingPolicyProtected`
- New function `*ClientFactory.NewDiscountsClient() *DiscountsClient`
- New function `*ClientFactory.NewDiscountsOperationGroupClient() *DiscountsOperationGroupClient`
- New function `*ClientFactory.NewReservationOrderAliasResponsesClient() *ReservationOrderAliasResponsesClient`
- New function `*ClientFactory.NewSavingsPlanModelsClient() *SavingsPlanModelsClient`
- New function `*ClientFactory.NewSavingsPlanOperationGroupClient() *SavingsPlanOperationGroupClient`
- New function `*ClientFactory.NewSavingsPlanOrderAliasModelsClient() *SavingsPlanOrderAliasModelsClient`
- New function `*ClientFactory.NewSavingsPlanOrderModelsClient() *SavingsPlanOrderModelsClient`
- New function `*DiscountProperties.GetDiscountProperties() *DiscountProperties`
- New function `*DiscountTypeProduct.GetDiscountTypeProperties() *DiscountTypeProperties`
- New function `*DiscountTypeProductFamily.GetDiscountTypeProperties() *DiscountTypeProperties`
- New function `*DiscountTypeProductSKU.GetDiscountTypeProperties() *DiscountTypeProperties`
- New function `*DiscountTypeProperties.GetDiscountTypeProperties() *DiscountTypeProperties`
- New function `NewDiscountsClient(string, azcore.TokenCredential, *arm.ClientOptions) (*DiscountsClient, error)`
- New function `*DiscountsClient.BeginCancel(context.Context, string, string, *DiscountsClientBeginCancelOptions) (*runtime.Poller[DiscountsClientCancelResponse], error)`
- New function `*DiscountsClient.BeginCreate(context.Context, string, string, Discount, *DiscountsClientBeginCreateOptions) (*runtime.Poller[DiscountsClientCreateResponse], error)`
- New function `*DiscountsClient.BeginDelete(context.Context, string, string, *DiscountsClientBeginDeleteOptions) (*runtime.Poller[DiscountsClientDeleteResponse], error)`
- New function `*DiscountsClient.Get(context.Context, string, string, *DiscountsClientGetOptions) (DiscountsClientGetResponse, error)`
- New function `*DiscountsClient.NewResourceGroupListPager(string, *DiscountsClientResourceGroupListOptions) *runtime.Pager[DiscountsClientResourceGroupListResponse]`
- New function `*DiscountsClient.NewSubscriptionListPager(*DiscountsClientSubscriptionListOptions) *runtime.Pager[DiscountsClientSubscriptionListResponse]`
- New function `*DiscountsClient.BeginUpdate(context.Context, string, string, DiscountPatchRequest, *DiscountsClientBeginUpdateOptions) (*runtime.Poller[DiscountsClientUpdateResponse], error)`
- New function `NewDiscountsOperationGroupClient(azcore.TokenCredential, *arm.ClientOptions) (*DiscountsOperationGroupClient, error)`
- New function `*DiscountsOperationGroupClient.NewScopeListPager(string, *DiscountsOperationGroupClientScopeListOptions) *runtime.Pager[DiscountsOperationGroupClientScopeListResponse]`
- New function `*EntityTypeAffiliateDiscount.GetDiscountProperties() *DiscountProperties`
- New function `*EntityTypePrimaryDiscount.GetDiscountProperties() *DiscountProperties`
- New function `NewReservationOrderAliasResponsesClient(azcore.TokenCredential, *arm.ClientOptions) (*ReservationOrderAliasResponsesClient, error)`
- New function `*ReservationOrderAliasResponsesClient.BeginCreate(context.Context, string, ReservationOrderAliasRequest, *ReservationOrderAliasResponsesClientBeginCreateOptions) (*runtime.Poller[ReservationOrderAliasResponsesClientCreateResponse], error)`
- New function `*ReservationOrderAliasResponsesClient.Get(context.Context, string, *ReservationOrderAliasResponsesClientGetOptions) (ReservationOrderAliasResponsesClientGetResponse, error)`
- New function `NewSavingsPlanModelsClient(azcore.TokenCredential, *arm.ClientOptions) (*SavingsPlanModelsClient, error)`
- New function `*SavingsPlanModelsClient.Get(context.Context, string, string, *SavingsPlanModelsClientGetOptions) (SavingsPlanModelsClientGetResponse, error)`
- New function `*SavingsPlanModelsClient.NewListPager(string, *SavingsPlanModelsClientListOptions) *runtime.Pager[SavingsPlanModelsClientListResponse]`
- New function `*SavingsPlanModelsClient.BeginUpdate(context.Context, string, string, SavingsPlanUpdateRequest, *SavingsPlanModelsClientBeginUpdateOptions) (*runtime.Poller[SavingsPlanModelsClientUpdateResponse], error)`
- New function `*SavingsPlanModelsClient.ValidateUpdate(context.Context, string, string, SavingsPlanUpdateValidateRequest, *SavingsPlanModelsClientValidateUpdateOptions) (SavingsPlanModelsClientValidateUpdateResponse, error)`
- New function `NewSavingsPlanOperationGroupClient(azcore.TokenCredential, *arm.ClientOptions) (*SavingsPlanOperationGroupClient, error)`
- New function `*SavingsPlanOperationGroupClient.NewListAllPager(*SavingsPlanOperationGroupClientListAllOptions) *runtime.Pager[SavingsPlanOperationGroupClientListAllResponse]`
- New function `*SavingsPlanOperationGroupClient.ValidatePurchase(context.Context, SavingsPlanPurchaseValidateRequest, *SavingsPlanOperationGroupClientValidatePurchaseOptions) (SavingsPlanOperationGroupClientValidatePurchaseResponse, error)`
- New function `NewSavingsPlanOrderAliasModelsClient(azcore.TokenCredential, *arm.ClientOptions) (*SavingsPlanOrderAliasModelsClient, error)`
- New function `*SavingsPlanOrderAliasModelsClient.BeginCreate(context.Context, string, SavingsPlanOrderAliasModel, *SavingsPlanOrderAliasModelsClientBeginCreateOptions) (*runtime.Poller[SavingsPlanOrderAliasModelsClientCreateResponse], error)`
- New function `*SavingsPlanOrderAliasModelsClient.Get(context.Context, string, *SavingsPlanOrderAliasModelsClientGetOptions) (SavingsPlanOrderAliasModelsClientGetResponse, error)`
- New function `NewSavingsPlanOrderModelsClient(azcore.TokenCredential, *arm.ClientOptions) (*SavingsPlanOrderModelsClient, error)`
- New function `*SavingsPlanOrderModelsClient.Elevate(context.Context, string, *SavingsPlanOrderModelsClientElevateOptions) (SavingsPlanOrderModelsClientElevateResponse, error)`
- New function `*SavingsPlanOrderModelsClient.Get(context.Context, string, *SavingsPlanOrderModelsClientGetOptions) (SavingsPlanOrderModelsClientGetResponse, error)`
- New function `*SavingsPlanOrderModelsClient.NewListPager(*SavingsPlanOrderModelsClientListOptions) *runtime.Pager[SavingsPlanOrderModelsClientListResponse]`
- New struct `ConditionsItem`
- New struct `Discount`
- New struct `DiscountList`
- New struct `DiscountPatchRequest`
- New struct `DiscountPatchRequestProperties`
- New struct `DiscountTypeProduct`
- New struct `DiscountTypeProductFamily`
- New struct `DiscountTypeProductSKU`
- New struct `EntityTypeAffiliateDiscount`
- New struct `EntityTypePrimaryDiscount`
- New struct `PriceGuaranteeProperties`
- New struct `ResourceSKU`
- New field `Renew` in struct `SavingsPlanOrderAliasProperties`


## 2.1.0 (2023-11-24)
### Features Added

- Support for test fakes and OpenTelemetry trace spans.


## 2.0.0 (2023-04-03)
### Breaking Changes

- Function `NewSavingsPlanClient` parameter(s) have been changed from `(*string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`
- Function `NewSavingsPlanOrderClient` parameter(s) have been changed from `(*string, azcore.TokenCredential, *arm.ClientOptions)` to `(azcore.TokenCredential, *arm.ClientOptions)`

### Features Added

- New struct `ClientFactory` which is a client factory used to create any client in this module
- New field `Expand` in struct `SavingsPlanClientGetOptions`
- New field `Expand` in struct `SavingsPlanOrderClientGetOptions`


## 1.0.0 (2022-12-23)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/billingbenefits/armbillingbenefits` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).