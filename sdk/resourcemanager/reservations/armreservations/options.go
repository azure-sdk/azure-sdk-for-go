// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

// AzureReservationAPIClientGetAppliedReservationListOptions contains the optional parameters for the AzureReservationAPIClient.GetAppliedReservationList
// method.
type AzureReservationAPIClientGetAppliedReservationListOptions struct {
	// placeholder for future optional parameters
}

// AzureReservationAPIClientGetCatalogOptions contains the optional parameters for the AzureReservationAPIClient.NewGetCatalogPager
// method.
type AzureReservationAPIClientGetCatalogOptions struct {
	// May be used to filter by Catalog properties. The filter supports 'eq', 'or', and 'and'.
	Filter *string

	// Filters the skus based on the location specified in this parameter. This can be an Azure region or global
	Location *string

	// Offer id used to get the third party products
	OfferID *string

	// Plan id used to get the third party products
	PlanID *string

	// Publisher id used to get the third party products
	PublisherID *string

	// The type of the resource for which the skus should be provided.
	ReservedResourceType *string

	// The number of reservations to skip from the list before returning results
	Skip *float32

	// To number of reservations to return
	Take *float32
}

// CalculateExchangeClientBeginPostOptions contains the optional parameters for the CalculateExchangeClient.BeginPost method.
type CalculateExchangeClientBeginPostOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// CalculateRefundClientPostOptions contains the optional parameters for the CalculateRefundClient.Post method.
type CalculateRefundClientPostOptions struct {
	// placeholder for future optional parameters
}

// ExchangeClientBeginPostOptions contains the optional parameters for the ExchangeClient.BeginPost method.
type ExchangeClientBeginPostOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// OperationClientListOptions contains the optional parameters for the OperationClient.NewListPager method.
type OperationClientListOptions struct {
	// placeholder for future optional parameters
}

// QuotaClientBeginCreateOrUpdateOptions contains the optional parameters for the QuotaClient.BeginCreateOrUpdate method.
type QuotaClientBeginCreateOrUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// QuotaClientBeginUpdateOptions contains the optional parameters for the QuotaClient.BeginUpdate method.
type QuotaClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// QuotaClientGetOptions contains the optional parameters for the QuotaClient.Get method.
type QuotaClientGetOptions struct {
	// placeholder for future optional parameters
}

// QuotaClientListOptions contains the optional parameters for the QuotaClient.NewListPager method.
type QuotaClientListOptions struct {
	// placeholder for future optional parameters
}

// QuotaRequestStatusClientGetOptions contains the optional parameters for the QuotaRequestStatusClient.Get method.
type QuotaRequestStatusClientGetOptions struct {
	// placeholder for future optional parameters
}

// QuotaRequestStatusClientListOptions contains the optional parameters for the QuotaRequestStatusClient.NewListPager method.
type QuotaRequestStatusClientListOptions struct {
	// FIELD SUPPORTED OPERATORS
	// requestSubmitTime ge, le, eq, gt, lt
	Filter *string

	// Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a nextLink element,
	// the value of the nextLink element includes a skiptoken parameter that
	// specifies a starting point to use for subsequent calls.
	Skiptoken *string

	// Number of records to return.
	Top *int32
}

// ReservationClientArchiveOptions contains the optional parameters for the ReservationClient.Archive method.
type ReservationClientArchiveOptions struct {
	// placeholder for future optional parameters
}

// ReservationClientBeginAvailableScopesOptions contains the optional parameters for the ReservationClient.BeginAvailableScopes
// method.
type ReservationClientBeginAvailableScopesOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ReservationClientBeginMergeOptions contains the optional parameters for the ReservationClient.BeginMerge method.
type ReservationClientBeginMergeOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ReservationClientBeginSplitOptions contains the optional parameters for the ReservationClient.BeginSplit method.
type ReservationClientBeginSplitOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ReservationClientBeginUpdateOptions contains the optional parameters for the ReservationClient.BeginUpdate method.
type ReservationClientBeginUpdateOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ReservationClientGetOptions contains the optional parameters for the ReservationClient.Get method.
type ReservationClientGetOptions struct {
	// Supported value of this query is renewProperties
	Expand *string
}

// ReservationClientListAllOptions contains the optional parameters for the ReservationClient.NewListAllPager method.
type ReservationClientListAllOptions struct {
	// May be used to filter by reservation properties. The filter supports 'eq', 'or', and 'and'. It does not currently support
	// 'ne', 'gt', 'le', 'ge', or 'not'. Reservation properties include sku/name,
	// properties/{appliedScopeType, archived, displayName, displayProvisioningState, effectiveDateTime, expiryDate, expiryDateTime,
	// provisioningState, quantity, renew, reservedResourceType, term,
	// userFriendlyAppliedScopeType, userFriendlyRenewState}
	Filter *string

	// May be used to sort order by reservation properties.
	Orderby *string

	// To indicate whether to refresh the roll up counts of the reservations group by provisioning states
	RefreshSummary *string

	// The selected provisioning state
	SelectedState *string

	// The number of reservations to skip from the list before returning results
	Skiptoken *float32

	// To number of reservations to return
	Take *float32
}

// ReservationClientListOptions contains the optional parameters for the ReservationClient.NewListPager method.
type ReservationClientListOptions struct {
	// placeholder for future optional parameters
}

// ReservationClientListRevisionsOptions contains the optional parameters for the ReservationClient.NewListRevisionsPager
// method.
type ReservationClientListRevisionsOptions struct {
	// placeholder for future optional parameters
}

// ReservationClientUnarchiveOptions contains the optional parameters for the ReservationClient.Unarchive method.
type ReservationClientUnarchiveOptions struct {
	// placeholder for future optional parameters
}

// ReservationOrderClientBeginPurchaseOptions contains the optional parameters for the ReservationOrderClient.BeginPurchase
// method.
type ReservationOrderClientBeginPurchaseOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}

// ReservationOrderClientCalculateOptions contains the optional parameters for the ReservationOrderClient.Calculate method.
type ReservationOrderClientCalculateOptions struct {
	// placeholder for future optional parameters
}

// ReservationOrderClientChangeDirectoryOptions contains the optional parameters for the ReservationOrderClient.ChangeDirectory
// method.
type ReservationOrderClientChangeDirectoryOptions struct {
	// placeholder for future optional parameters
}

// ReservationOrderClientGetOptions contains the optional parameters for the ReservationOrderClient.Get method.
type ReservationOrderClientGetOptions struct {
	// May be used to expand the planInformation.
	Expand *string
}

// ReservationOrderClientListOptions contains the optional parameters for the ReservationOrderClient.NewListPager method.
type ReservationOrderClientListOptions struct {
	// placeholder for future optional parameters
}

// ReturnClientBeginPostOptions contains the optional parameters for the ReturnClient.BeginPost method.
type ReturnClientBeginPostOptions struct {
	// Resumes the long-running operation from the provided token.
	ResumeToken string
}
