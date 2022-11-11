# Release History

## 2.0.0 (2022-11-11)
### Breaking Changes

- Const `QuotaRequestStateFailed` has been removed
- Const `ResourceTypeLowPriority` has been removed
- Const `ResourceTypeStandard` has been removed
- Const `QuotaRequestStateInvalid` has been removed
- Const `ResourceTypeShared` has been removed
- Const `QuotaRequestStateAccepted` has been removed
- Const `QuotaRequestStateSucceeded` has been removed
- Const `ResourceTypeServiceSpecific` has been removed
- Const `ResourceTypeDedicated` has been removed
- Const `QuotaRequestStateInProgress` has been removed
- Type alias `QuotaRequestState` has been removed
- Type alias `ResourceType` has been removed
- Function `NewQuotaRequestStatusClient` has been removed
- Function `*QuotaClient.BeginCreateOrUpdate` has been removed
- Function `PossibleResourceTypeValues` has been removed
- Function `NewQuotaClient` has been removed
- Function `*QuotaClient.BeginUpdate` has been removed
- Function `PossibleQuotaRequestStateValues` has been removed
- Function `*QuotaRequestStatusClient.Get` has been removed
- Function `*QuotaClient.NewListPager` has been removed
- Function `*QuotaRequestStatusClient.NewListPager` has been removed
- Function `*QuotaClient.Get` has been removed
- Struct `CreateGenericQuotaRequestParameters` has been removed
- Struct `CurrentQuotaLimit` has been removed
- Struct `CurrentQuotaLimitBase` has been removed
- Struct `ExceptionResponse` has been removed
- Struct `QuotaClient` has been removed
- Struct `QuotaClientBeginCreateOrUpdateOptions` has been removed
- Struct `QuotaClientBeginUpdateOptions` has been removed
- Struct `QuotaClientCreateOrUpdateResponse` has been removed
- Struct `QuotaClientGetOptions` has been removed
- Struct `QuotaClientGetResponse` has been removed
- Struct `QuotaClientListOptions` has been removed
- Struct `QuotaClientListResponse` has been removed
- Struct `QuotaClientUpdateResponse` has been removed
- Struct `QuotaLimits` has been removed
- Struct `QuotaLimitsResponse` has been removed
- Struct `QuotaProperties` has been removed
- Struct `QuotaRequestDetails` has been removed
- Struct `QuotaRequestDetailsList` has been removed
- Struct `QuotaRequestOneResourceProperties` has been removed
- Struct `QuotaRequestOneResourceSubmitResponse` has been removed
- Struct `QuotaRequestProperties` has been removed
- Struct `QuotaRequestStatusClient` has been removed
- Struct `QuotaRequestStatusClientGetOptions` has been removed
- Struct `QuotaRequestStatusClientGetResponse` has been removed
- Struct `QuotaRequestStatusClientListOptions` has been removed
- Struct `QuotaRequestStatusClientListResponse` has been removed
- Struct `QuotaRequestStatusDetails` has been removed
- Struct `QuotaRequestSubmitResponse` has been removed
- Struct `QuotaRequestSubmitResponse201` has been removed
- Struct `ResourceName` has been removed
- Struct `ServiceError` has been removed
- Struct `ServiceErrorDetail` has been removed
- Struct `SubRequest` has been removed

### Features Added

- New struct `AppliedScopeProperties`
- New struct `ReservationSwapProperties`
- New field `SwapProperties` in struct `Properties`
- New field `AppliedScopeProperties` in struct `Properties`


## 1.1.0 (2022-09-16)
### Features Added

- New const `ErrorResponseCodeRefundLimitExceeded`
- New const `ErrorResponseCodeSelfServiceRefundNotSupported`
- New function `*ReservationClient.Archive(context.Context, string, string, *ReservationClientArchiveOptions) (ReservationClientArchiveResponse, error)`
- New function `NewCalculateRefundClient(azcore.TokenCredential, *arm.ClientOptions) (*CalculateRefundClient, error)`
- New function `*CalculateRefundClient.Post(context.Context, string, CalculateRefundRequest, *CalculateRefundClientPostOptions) (CalculateRefundClientPostResponse, error)`
- New function `*ReservationClient.Unarchive(context.Context, string, string, *ReservationClientUnarchiveOptions) (ReservationClientUnarchiveResponse, error)`
- New function `NewReturnClient(azcore.TokenCredential, *arm.ClientOptions) (*ReturnClient, error)`
- New function `*ReturnClient.Post(context.Context, string, RefundRequest, *ReturnClientPostOptions) (ReturnClientPostResponse, error)`
- New struct `CalculateRefundClient`
- New struct `CalculateRefundClientPostOptions`
- New struct `CalculateRefundClientPostResponse`
- New struct `CalculateRefundRequest`
- New struct `CalculateRefundRequestProperties`
- New struct `CalculateRefundResponse`
- New struct `RefundBillingInformation`
- New struct `RefundPolicyError`
- New struct `RefundPolicyResult`
- New struct `RefundPolicyResultProperty`
- New struct `RefundRequest`
- New struct `RefundRequestProperties`
- New struct `RefundResponse`
- New struct `RefundResponseProperties`
- New struct `ReservationClientArchiveOptions`
- New struct `ReservationClientArchiveResponse`
- New struct `ReservationClientUnarchiveOptions`
- New struct `ReservationClientUnarchiveResponse`
- New struct `ReturnClient`
- New struct `ReturnClientPostOptions`
- New struct `ReturnClientPostResponse`


## 1.0.0 (2022-05-18)

The package of `github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/reservations/armreservations` is using our [next generation design principles](https://azure.github.io/azure-sdk/general_introduction.html) since version 1.0.0, which contains breaking changes.

To migrate the existing applications to the latest version, please refer to [Migration Guide](https://aka.ms/azsdk/go/mgmt/migration).

To learn more, please refer to our documentation [Quick Start](https://aka.ms/azsdk/go/mgmt).