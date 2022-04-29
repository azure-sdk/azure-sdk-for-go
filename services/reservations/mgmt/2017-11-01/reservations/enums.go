package reservations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// AppliedScopeType enumerates the values for applied scope type.
type AppliedScopeType string

const (
	// Shared ...
	Shared AppliedScopeType = "Shared"
	// Single ...
	Single AppliedScopeType = "Single"
)

// PossibleAppliedScopeTypeValues returns an array of possible values for the AppliedScopeType const type.
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return []AppliedScopeType{Shared, Single}
}

// AppliedScopeType1 enumerates the values for applied scope type 1.
type AppliedScopeType1 string

const (
	// AppliedScopeType1Shared ...
	AppliedScopeType1Shared AppliedScopeType1 = "Shared"
	// AppliedScopeType1Single ...
	AppliedScopeType1Single AppliedScopeType1 = "Single"
)

// PossibleAppliedScopeType1Values returns an array of possible values for the AppliedScopeType1 const type.
func PossibleAppliedScopeType1Values() []AppliedScopeType1 {
	return []AppliedScopeType1{AppliedScopeType1Shared, AppliedScopeType1Single}
}

// Code enumerates the values for code.
type Code string

const (
	// ActivateQuoteFailed ...
	ActivateQuoteFailed Code = "ActivateQuoteFailed"
	// AppliedScopesNotAssociatedWithCommerceAccount ...
	AppliedScopesNotAssociatedWithCommerceAccount Code = "AppliedScopesNotAssociatedWithCommerceAccount"
	// AppliedScopesSameAsExisting ...
	AppliedScopesSameAsExisting Code = "AppliedScopesSameAsExisting"
	// AuthorizationFailed ...
	AuthorizationFailed Code = "AuthorizationFailed"
	// BadRequest ...
	BadRequest Code = "BadRequest"
	// BillingCustomerInputError ...
	BillingCustomerInputError Code = "BillingCustomerInputError"
	// BillingError ...
	BillingError Code = "BillingError"
	// BillingPaymentInstrumentHardError ...
	BillingPaymentInstrumentHardError Code = "BillingPaymentInstrumentHardError"
	// BillingPaymentInstrumentSoftError ...
	BillingPaymentInstrumentSoftError Code = "BillingPaymentInstrumentSoftError"
	// BillingScopeIDCannotBeChanged ...
	BillingScopeIDCannotBeChanged Code = "BillingScopeIdCannotBeChanged"
	// BillingTransientError ...
	BillingTransientError Code = "BillingTransientError"
	// CalculatePriceFailed ...
	CalculatePriceFailed Code = "CalculatePriceFailed"
	// CapacityUpdateScopesFailed ...
	CapacityUpdateScopesFailed Code = "CapacityUpdateScopesFailed"
	// ClientCertificateThumbprintNotSet ...
	ClientCertificateThumbprintNotSet Code = "ClientCertificateThumbprintNotSet"
	// CreateQuoteFailed ...
	CreateQuoteFailed Code = "CreateQuoteFailed"
	// Forbidden ...
	Forbidden Code = "Forbidden"
	// FulfillmentConfigurationError ...
	FulfillmentConfigurationError Code = "FulfillmentConfigurationError"
	// FulfillmentError ...
	FulfillmentError Code = "FulfillmentError"
	// FulfillmentOutOfStockError ...
	FulfillmentOutOfStockError Code = "FulfillmentOutOfStockError"
	// FulfillmentTransientError ...
	FulfillmentTransientError Code = "FulfillmentTransientError"
	// HTTPMethodNotSupported ...
	HTTPMethodNotSupported Code = "HttpMethodNotSupported"
	// InternalServerError ...
	InternalServerError Code = "InternalServerError"
	// InvalidAccessToken ...
	InvalidAccessToken Code = "InvalidAccessToken"
	// InvalidFulfillmentRequestParameters ...
	InvalidFulfillmentRequestParameters Code = "InvalidFulfillmentRequestParameters"
	// InvalidHealthCheckType ...
	InvalidHealthCheckType Code = "InvalidHealthCheckType"
	// InvalidLocationID ...
	InvalidLocationID Code = "InvalidLocationId"
	// InvalidRefundQuantity ...
	InvalidRefundQuantity Code = "InvalidRefundQuantity"
	// InvalidRequestContent ...
	InvalidRequestContent Code = "InvalidRequestContent"
	// InvalidRequestURI ...
	InvalidRequestURI Code = "InvalidRequestUri"
	// InvalidReservationID ...
	InvalidReservationID Code = "InvalidReservationId"
	// InvalidReservationOrderID ...
	InvalidReservationOrderID Code = "InvalidReservationOrderId"
	// InvalidSingleAppliedScopesCount ...
	InvalidSingleAppliedScopesCount Code = "InvalidSingleAppliedScopesCount"
	// InvalidSubscriptionID ...
	InvalidSubscriptionID Code = "InvalidSubscriptionId"
	// InvalidTenantID ...
	InvalidTenantID Code = "InvalidTenantId"
	// MissingAppliedScopesForSingle ...
	MissingAppliedScopesForSingle Code = "MissingAppliedScopesForSingle"
	// MissingTenantID ...
	MissingTenantID Code = "MissingTenantId"
	// NonsupportedAccountID ...
	NonsupportedAccountID Code = "NonsupportedAccountId"
	// NotSpecified ...
	NotSpecified Code = "NotSpecified"
	// NotSupportedCountry ...
	NotSupportedCountry Code = "NotSupportedCountry"
	// NoValidReservationsToReRate ...
	NoValidReservationsToReRate Code = "NoValidReservationsToReRate"
	// OperationCannotBePerformedInCurrentState ...
	OperationCannotBePerformedInCurrentState Code = "OperationCannotBePerformedInCurrentState"
	// OperationFailed ...
	OperationFailed Code = "OperationFailed"
	// PaymentInstrumentNotFound ...
	PaymentInstrumentNotFound Code = "PaymentInstrumentNotFound"
	// PurchaseError ...
	PurchaseError Code = "PurchaseError"
	// ReRateOnlyAllowedForEA ...
	ReRateOnlyAllowedForEA Code = "ReRateOnlyAllowedForEA"
	// ReservationIDNotInReservationOrder ...
	ReservationIDNotInReservationOrder Code = "ReservationIdNotInReservationOrder"
	// ReservationOrderCreationFailed ...
	ReservationOrderCreationFailed Code = "ReservationOrderCreationFailed"
	// ReservationOrderIDAlreadyExists ...
	ReservationOrderIDAlreadyExists Code = "ReservationOrderIdAlreadyExists"
	// ReservationOrderNotEnabled ...
	ReservationOrderNotEnabled Code = "ReservationOrderNotEnabled"
	// ReservationOrderNotFound ...
	ReservationOrderNotFound Code = "ReservationOrderNotFound"
	// RiskCheckFailed ...
	RiskCheckFailed Code = "RiskCheckFailed"
	// RoleAssignmentCreationFailed ...
	RoleAssignmentCreationFailed Code = "RoleAssignmentCreationFailed"
	// ServerTimeout ...
	ServerTimeout Code = "ServerTimeout"
	// UnauthenticatedRequestsThrottled ...
	UnauthenticatedRequestsThrottled Code = "UnauthenticatedRequestsThrottled"
	// UnsupportedReservationTerm ...
	UnsupportedReservationTerm Code = "UnsupportedReservationTerm"
)

// PossibleCodeValues returns an array of possible values for the Code const type.
func PossibleCodeValues() []Code {
	return []Code{ActivateQuoteFailed, AppliedScopesNotAssociatedWithCommerceAccount, AppliedScopesSameAsExisting, AuthorizationFailed, BadRequest, BillingCustomerInputError, BillingError, BillingPaymentInstrumentHardError, BillingPaymentInstrumentSoftError, BillingScopeIDCannotBeChanged, BillingTransientError, CalculatePriceFailed, CapacityUpdateScopesFailed, ClientCertificateThumbprintNotSet, CreateQuoteFailed, Forbidden, FulfillmentConfigurationError, FulfillmentError, FulfillmentOutOfStockError, FulfillmentTransientError, HTTPMethodNotSupported, InternalServerError, InvalidAccessToken, InvalidFulfillmentRequestParameters, InvalidHealthCheckType, InvalidLocationID, InvalidRefundQuantity, InvalidRequestContent, InvalidRequestURI, InvalidReservationID, InvalidReservationOrderID, InvalidSingleAppliedScopesCount, InvalidSubscriptionID, InvalidTenantID, MissingAppliedScopesForSingle, MissingTenantID, NonsupportedAccountID, NotSpecified, NotSupportedCountry, NoValidReservationsToReRate, OperationCannotBePerformedInCurrentState, OperationFailed, PaymentInstrumentNotFound, PurchaseError, ReRateOnlyAllowedForEA, ReservationIDNotInReservationOrder, ReservationOrderCreationFailed, ReservationOrderIDAlreadyExists, ReservationOrderNotEnabled, ReservationOrderNotFound, RiskCheckFailed, RoleAssignmentCreationFailed, ServerTimeout, UnauthenticatedRequestsThrottled, UnsupportedReservationTerm}
}

// Kind enumerates the values for kind.
type Kind string

const (
	// MicrosoftCompute ...
	MicrosoftCompute Kind = "Microsoft.Compute"
)

// PossibleKindValues returns an array of possible values for the Kind const type.
func PossibleKindValues() []Kind {
	return []Kind{MicrosoftCompute}
}

// Location enumerates the values for location.
type Location string

const (
	// Australiaeast ...
	Australiaeast Location = "australiaeast"
	// Australiasoutheast ...
	Australiasoutheast Location = "australiasoutheast"
	// Brazilsouth ...
	Brazilsouth Location = "brazilsouth"
	// Canadacentral ...
	Canadacentral Location = "canadacentral"
	// Canadaeast ...
	Canadaeast Location = "canadaeast"
	// Centralindia ...
	Centralindia Location = "centralindia"
	// Centralus ...
	Centralus Location = "centralus"
	// Eastasia ...
	Eastasia Location = "eastasia"
	// Eastus ...
	Eastus Location = "eastus"
	// Eastus2 ...
	Eastus2 Location = "eastus2"
	// Japaneast ...
	Japaneast Location = "japaneast"
	// Japanwest ...
	Japanwest Location = "japanwest"
	// Northcentralus ...
	Northcentralus Location = "northcentralus"
	// Northeurope ...
	Northeurope Location = "northeurope"
	// Southcentralus ...
	Southcentralus Location = "southcentralus"
	// Southeastasia ...
	Southeastasia Location = "southeastasia"
	// Southindia ...
	Southindia Location = "southindia"
	// Uksouth ...
	Uksouth Location = "uksouth"
	// Ukwest ...
	Ukwest Location = "ukwest"
	// Westcentralus ...
	Westcentralus Location = "westcentralus"
	// Westeurope ...
	Westeurope Location = "westeurope"
	// Westindia ...
	Westindia Location = "westindia"
	// Westus ...
	Westus Location = "westus"
	// Westus2 ...
	Westus2 Location = "westus2"
)

// PossibleLocationValues returns an array of possible values for the Location const type.
func PossibleLocationValues() []Location {
	return []Location{Australiaeast, Australiasoutheast, Brazilsouth, Canadacentral, Canadaeast, Centralindia, Centralus, Eastasia, Eastus, Eastus2, Japaneast, Japanwest, Northcentralus, Northeurope, Southcentralus, Southeastasia, Southindia, Uksouth, Ukwest, Westcentralus, Westeurope, Westindia, Westus, Westus2}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// BillingFailed ...
	BillingFailed ProvisioningState = "BillingFailed"
	// Cancelled ...
	Cancelled ProvisioningState = "Cancelled"
	// ConfirmedBilling ...
	ConfirmedBilling ProvisioningState = "ConfirmedBilling"
	// ConfirmedResourceHold ...
	ConfirmedResourceHold ProvisioningState = "ConfirmedResourceHold"
	// Created ...
	Created ProvisioningState = "Created"
	// Creating ...
	Creating ProvisioningState = "Creating"
	// Expired ...
	Expired ProvisioningState = "Expired"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Merged ...
	Merged ProvisioningState = "Merged"
	// PendingBilling ...
	PendingBilling ProvisioningState = "PendingBilling"
	// PendingResourceHold ...
	PendingResourceHold ProvisioningState = "PendingResourceHold"
	// Processing ...
	Processing ProvisioningState = "Processing"
	// Split ...
	Split ProvisioningState = "Split"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{BillingFailed, Cancelled, ConfirmedBilling, ConfirmedResourceHold, Created, Creating, Expired, Failed, Merged, PendingBilling, PendingResourceHold, Processing, Split, Succeeded}
}

// ProvisioningState1 enumerates the values for provisioning state 1.
type ProvisioningState1 string

const (
	// ProvisioningState1BillingFailed ...
	ProvisioningState1BillingFailed ProvisioningState1 = "BillingFailed"
	// ProvisioningState1Cancelled ...
	ProvisioningState1Cancelled ProvisioningState1 = "Cancelled"
	// ProvisioningState1ConfirmedBilling ...
	ProvisioningState1ConfirmedBilling ProvisioningState1 = "ConfirmedBilling"
	// ProvisioningState1ConfirmedResourceHold ...
	ProvisioningState1ConfirmedResourceHold ProvisioningState1 = "ConfirmedResourceHold"
	// ProvisioningState1Created ...
	ProvisioningState1Created ProvisioningState1 = "Created"
	// ProvisioningState1Creating ...
	ProvisioningState1Creating ProvisioningState1 = "Creating"
	// ProvisioningState1Expired ...
	ProvisioningState1Expired ProvisioningState1 = "Expired"
	// ProvisioningState1Failed ...
	ProvisioningState1Failed ProvisioningState1 = "Failed"
	// ProvisioningState1Merged ...
	ProvisioningState1Merged ProvisioningState1 = "Merged"
	// ProvisioningState1PendingBilling ...
	ProvisioningState1PendingBilling ProvisioningState1 = "PendingBilling"
	// ProvisioningState1PendingResourceHold ...
	ProvisioningState1PendingResourceHold ProvisioningState1 = "PendingResourceHold"
	// ProvisioningState1Processing ...
	ProvisioningState1Processing ProvisioningState1 = "Processing"
	// ProvisioningState1Split ...
	ProvisioningState1Split ProvisioningState1 = "Split"
	// ProvisioningState1Succeeded ...
	ProvisioningState1Succeeded ProvisioningState1 = "Succeeded"
)

// PossibleProvisioningState1Values returns an array of possible values for the ProvisioningState1 const type.
func PossibleProvisioningState1Values() []ProvisioningState1 {
	return []ProvisioningState1{ProvisioningState1BillingFailed, ProvisioningState1Cancelled, ProvisioningState1ConfirmedBilling, ProvisioningState1ConfirmedResourceHold, ProvisioningState1Created, ProvisioningState1Creating, ProvisioningState1Expired, ProvisioningState1Failed, ProvisioningState1Merged, ProvisioningState1PendingBilling, ProvisioningState1PendingResourceHold, ProvisioningState1Processing, ProvisioningState1Split, ProvisioningState1Succeeded}
}

// StatusCode enumerates the values for status code.
type StatusCode string

const (
	// StatusCodeActive ...
	StatusCodeActive StatusCode = "Active"
	// StatusCodeExpired ...
	StatusCodeExpired StatusCode = "Expired"
	// StatusCodeMerged ...
	StatusCodeMerged StatusCode = "Merged"
	// StatusCodeNone ...
	StatusCodeNone StatusCode = "None"
	// StatusCodePaymentInstrumentError ...
	StatusCodePaymentInstrumentError StatusCode = "PaymentInstrumentError"
	// StatusCodePending ...
	StatusCodePending StatusCode = "Pending"
	// StatusCodePurchaseError ...
	StatusCodePurchaseError StatusCode = "PurchaseError"
	// StatusCodeSplit ...
	StatusCodeSplit StatusCode = "Split"
	// StatusCodeSucceeded ...
	StatusCodeSucceeded StatusCode = "Succeeded"
)

// PossibleStatusCodeValues returns an array of possible values for the StatusCode const type.
func PossibleStatusCodeValues() []StatusCode {
	return []StatusCode{StatusCodeActive, StatusCodeExpired, StatusCodeMerged, StatusCodeNone, StatusCodePaymentInstrumentError, StatusCodePending, StatusCodePurchaseError, StatusCodeSplit, StatusCodeSucceeded}
}

// Term enumerates the values for term.
type Term string

const (
	// P1Y ...
	P1Y Term = "P1Y"
	// P3Y ...
	P3Y Term = "P3Y"
)

// PossibleTermValues returns an array of possible values for the Term const type.
func PossibleTermValues() []Term {
	return []Term{P1Y, P3Y}
}
