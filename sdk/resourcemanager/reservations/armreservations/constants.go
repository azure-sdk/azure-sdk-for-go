//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armreservations

const (
	moduleName    = "armreservations"
	moduleVersion = "v3.0.2"
)

// AppliedScopeType - Type of the Applied Scope.
type AppliedScopeType string

const (
	AppliedScopeTypeManagementGroup AppliedScopeType = "ManagementGroup"
	AppliedScopeTypeShared          AppliedScopeType = "Shared"
	AppliedScopeTypeSingle          AppliedScopeType = "Single"
)

// PossibleAppliedScopeTypeValues returns the possible values for the AppliedScopeType const type.
func PossibleAppliedScopeTypeValues() []AppliedScopeType {
	return []AppliedScopeType{
		AppliedScopeTypeManagementGroup,
		AppliedScopeTypeShared,
		AppliedScopeTypeSingle,
	}
}

// BillingPlan - Represents the billing plan in ISO 8601 format. Required only for monthly billing plans.
type BillingPlan string

const (
	BillingPlanP1M BillingPlan = "P1M"
)

// PossibleBillingPlanValues returns the possible values for the BillingPlan const type.
func PossibleBillingPlanValues() []BillingPlan {
	return []BillingPlan{
		BillingPlanP1M,
	}
}

// CalculateExchangeOperationResultStatus - Status of the operation.
type CalculateExchangeOperationResultStatus string

const (
	CalculateExchangeOperationResultStatusCancelled CalculateExchangeOperationResultStatus = "Cancelled"
	CalculateExchangeOperationResultStatusFailed    CalculateExchangeOperationResultStatus = "Failed"
	CalculateExchangeOperationResultStatusPending   CalculateExchangeOperationResultStatus = "Pending"
	CalculateExchangeOperationResultStatusSucceeded CalculateExchangeOperationResultStatus = "Succeeded"
)

// PossibleCalculateExchangeOperationResultStatusValues returns the possible values for the CalculateExchangeOperationResultStatus const type.
func PossibleCalculateExchangeOperationResultStatusValues() []CalculateExchangeOperationResultStatus {
	return []CalculateExchangeOperationResultStatus{
		CalculateExchangeOperationResultStatusCancelled,
		CalculateExchangeOperationResultStatusFailed,
		CalculateExchangeOperationResultStatusPending,
		CalculateExchangeOperationResultStatusSucceeded,
	}
}

// CommitmentGrain - Commitment grain.
type CommitmentGrain string

const (
	CommitmentGrainHourly CommitmentGrain = "Hourly"
)

// PossibleCommitmentGrainValues returns the possible values for the CommitmentGrain const type.
func PossibleCommitmentGrainValues() []CommitmentGrain {
	return []CommitmentGrain{
		CommitmentGrainHourly,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// DisplayProvisioningState - Represent the current display state of the reservation.
type DisplayProvisioningState string

const (
	DisplayProvisioningStateCancelled  DisplayProvisioningState = "Cancelled"
	DisplayProvisioningStateExpired    DisplayProvisioningState = "Expired"
	DisplayProvisioningStateExpiring   DisplayProvisioningState = "Expiring"
	DisplayProvisioningStateFailed     DisplayProvisioningState = "Failed"
	DisplayProvisioningStateNoBenefit  DisplayProvisioningState = "NoBenefit"
	DisplayProvisioningStatePending    DisplayProvisioningState = "Pending"
	DisplayProvisioningStateProcessing DisplayProvisioningState = "Processing"
	DisplayProvisioningStateSucceeded  DisplayProvisioningState = "Succeeded"
	DisplayProvisioningStateWarning    DisplayProvisioningState = "Warning"
)

// PossibleDisplayProvisioningStateValues returns the possible values for the DisplayProvisioningState const type.
func PossibleDisplayProvisioningStateValues() []DisplayProvisioningState {
	return []DisplayProvisioningState{
		DisplayProvisioningStateCancelled,
		DisplayProvisioningStateExpired,
		DisplayProvisioningStateExpiring,
		DisplayProvisioningStateFailed,
		DisplayProvisioningStateNoBenefit,
		DisplayProvisioningStatePending,
		DisplayProvisioningStateProcessing,
		DisplayProvisioningStateSucceeded,
		DisplayProvisioningStateWarning,
	}
}

// ErrorResponseCode - Error code describing the reason that service is not able to process the incoming request
type ErrorResponseCode string

const (
	ErrorResponseCodeActivateQuoteFailed                           ErrorResponseCode = "ActivateQuoteFailed"
	ErrorResponseCodeAppliedScopesNotAssociatedWithCommerceAccount ErrorResponseCode = "AppliedScopesNotAssociatedWithCommerceAccount"
	ErrorResponseCodeAppliedScopesSameAsExisting                   ErrorResponseCode = "AppliedScopesSameAsExisting"
	ErrorResponseCodeAuthorizationFailed                           ErrorResponseCode = "AuthorizationFailed"
	ErrorResponseCodeBadRequest                                    ErrorResponseCode = "BadRequest"
	ErrorResponseCodeBillingCustomerInputError                     ErrorResponseCode = "BillingCustomerInputError"
	ErrorResponseCodeBillingError                                  ErrorResponseCode = "BillingError"
	ErrorResponseCodeBillingPaymentInstrumentHardError             ErrorResponseCode = "BillingPaymentInstrumentHardError"
	ErrorResponseCodeBillingPaymentInstrumentSoftError             ErrorResponseCode = "BillingPaymentInstrumentSoftError"
	ErrorResponseCodeBillingScopeIDCannotBeChanged                 ErrorResponseCode = "BillingScopeIdCannotBeChanged"
	ErrorResponseCodeBillingTransientError                         ErrorResponseCode = "BillingTransientError"
	ErrorResponseCodeCalculatePriceFailed                          ErrorResponseCode = "CalculatePriceFailed"
	ErrorResponseCodeCapacityUpdateScopesFailed                    ErrorResponseCode = "CapacityUpdateScopesFailed"
	ErrorResponseCodeClientCertificateThumbprintNotSet             ErrorResponseCode = "ClientCertificateThumbprintNotSet"
	ErrorResponseCodeCreateQuoteFailed                             ErrorResponseCode = "CreateQuoteFailed"
	ErrorResponseCodeForbidden                                     ErrorResponseCode = "Forbidden"
	ErrorResponseCodeFulfillmentConfigurationError                 ErrorResponseCode = "FulfillmentConfigurationError"
	ErrorResponseCodeFulfillmentError                              ErrorResponseCode = "FulfillmentError"
	ErrorResponseCodeFulfillmentOutOfStockError                    ErrorResponseCode = "FulfillmentOutOfStockError"
	ErrorResponseCodeFulfillmentTransientError                     ErrorResponseCode = "FulfillmentTransientError"
	ErrorResponseCodeHTTPMethodNotSupported                        ErrorResponseCode = "HttpMethodNotSupported"
	ErrorResponseCodeInternalServerError                           ErrorResponseCode = "InternalServerError"
	ErrorResponseCodeInvalidAccessToken                            ErrorResponseCode = "InvalidAccessToken"
	ErrorResponseCodeInvalidFulfillmentRequestParameters           ErrorResponseCode = "InvalidFulfillmentRequestParameters"
	ErrorResponseCodeInvalidHealthCheckType                        ErrorResponseCode = "InvalidHealthCheckType"
	ErrorResponseCodeInvalidLocationID                             ErrorResponseCode = "InvalidLocationId"
	ErrorResponseCodeInvalidRefundQuantity                         ErrorResponseCode = "InvalidRefundQuantity"
	ErrorResponseCodeInvalidRequestContent                         ErrorResponseCode = "InvalidRequestContent"
	ErrorResponseCodeInvalidRequestURI                             ErrorResponseCode = "InvalidRequestUri"
	ErrorResponseCodeInvalidReservationID                          ErrorResponseCode = "InvalidReservationId"
	ErrorResponseCodeInvalidReservationOrderID                     ErrorResponseCode = "InvalidReservationOrderId"
	ErrorResponseCodeInvalidSingleAppliedScopesCount               ErrorResponseCode = "InvalidSingleAppliedScopesCount"
	ErrorResponseCodeInvalidSubscriptionID                         ErrorResponseCode = "InvalidSubscriptionId"
	ErrorResponseCodeInvalidTenantID                               ErrorResponseCode = "InvalidTenantId"
	ErrorResponseCodeMissingAppliedScopesForSingle                 ErrorResponseCode = "MissingAppliedScopesForSingle"
	ErrorResponseCodeMissingTenantID                               ErrorResponseCode = "MissingTenantId"
	ErrorResponseCodeNoValidReservationsToReRate                   ErrorResponseCode = "NoValidReservationsToReRate"
	ErrorResponseCodeNonsupportedAccountID                         ErrorResponseCode = "NonsupportedAccountId"
	ErrorResponseCodeNotSpecified                                  ErrorResponseCode = "NotSpecified"
	ErrorResponseCodeNotSupportedCountry                           ErrorResponseCode = "NotSupportedCountry"
	ErrorResponseCodeOperationCannotBePerformedInCurrentState      ErrorResponseCode = "OperationCannotBePerformedInCurrentState"
	ErrorResponseCodeOperationFailed                               ErrorResponseCode = "OperationFailed"
	ErrorResponseCodePatchValuesSameAsExisting                     ErrorResponseCode = "PatchValuesSameAsExisting"
	ErrorResponseCodePaymentInstrumentNotFound                     ErrorResponseCode = "PaymentInstrumentNotFound"
	ErrorResponseCodePurchaseError                                 ErrorResponseCode = "PurchaseError"
	ErrorResponseCodeReRateOnlyAllowedForEA                        ErrorResponseCode = "ReRateOnlyAllowedForEA"
	ErrorResponseCodeRefundLimitExceeded                           ErrorResponseCode = "RefundLimitExceeded"
	ErrorResponseCodeReservationIDNotInReservationOrder            ErrorResponseCode = "ReservationIdNotInReservationOrder"
	ErrorResponseCodeReservationOrderCreationFailed                ErrorResponseCode = "ReservationOrderCreationFailed"
	ErrorResponseCodeReservationOrderIDAlreadyExists               ErrorResponseCode = "ReservationOrderIdAlreadyExists"
	ErrorResponseCodeReservationOrderNotEnabled                    ErrorResponseCode = "ReservationOrderNotEnabled"
	ErrorResponseCodeReservationOrderNotFound                      ErrorResponseCode = "ReservationOrderNotFound"
	ErrorResponseCodeRiskCheckFailed                               ErrorResponseCode = "RiskCheckFailed"
	ErrorResponseCodeRoleAssignmentCreationFailed                  ErrorResponseCode = "RoleAssignmentCreationFailed"
	ErrorResponseCodeSelfServiceRefundNotSupported                 ErrorResponseCode = "SelfServiceRefundNotSupported"
	ErrorResponseCodeServerTimeout                                 ErrorResponseCode = "ServerTimeout"
	ErrorResponseCodeUnauthenticatedRequestsThrottled              ErrorResponseCode = "UnauthenticatedRequestsThrottled"
	ErrorResponseCodeUnsupportedReservationTerm                    ErrorResponseCode = "UnsupportedReservationTerm"
)

// PossibleErrorResponseCodeValues returns the possible values for the ErrorResponseCode const type.
func PossibleErrorResponseCodeValues() []ErrorResponseCode {
	return []ErrorResponseCode{
		ErrorResponseCodeActivateQuoteFailed,
		ErrorResponseCodeAppliedScopesNotAssociatedWithCommerceAccount,
		ErrorResponseCodeAppliedScopesSameAsExisting,
		ErrorResponseCodeAuthorizationFailed,
		ErrorResponseCodeBadRequest,
		ErrorResponseCodeBillingCustomerInputError,
		ErrorResponseCodeBillingError,
		ErrorResponseCodeBillingPaymentInstrumentHardError,
		ErrorResponseCodeBillingPaymentInstrumentSoftError,
		ErrorResponseCodeBillingScopeIDCannotBeChanged,
		ErrorResponseCodeBillingTransientError,
		ErrorResponseCodeCalculatePriceFailed,
		ErrorResponseCodeCapacityUpdateScopesFailed,
		ErrorResponseCodeClientCertificateThumbprintNotSet,
		ErrorResponseCodeCreateQuoteFailed,
		ErrorResponseCodeForbidden,
		ErrorResponseCodeFulfillmentConfigurationError,
		ErrorResponseCodeFulfillmentError,
		ErrorResponseCodeFulfillmentOutOfStockError,
		ErrorResponseCodeFulfillmentTransientError,
		ErrorResponseCodeHTTPMethodNotSupported,
		ErrorResponseCodeInternalServerError,
		ErrorResponseCodeInvalidAccessToken,
		ErrorResponseCodeInvalidFulfillmentRequestParameters,
		ErrorResponseCodeInvalidHealthCheckType,
		ErrorResponseCodeInvalidLocationID,
		ErrorResponseCodeInvalidRefundQuantity,
		ErrorResponseCodeInvalidRequestContent,
		ErrorResponseCodeInvalidRequestURI,
		ErrorResponseCodeInvalidReservationID,
		ErrorResponseCodeInvalidReservationOrderID,
		ErrorResponseCodeInvalidSingleAppliedScopesCount,
		ErrorResponseCodeInvalidSubscriptionID,
		ErrorResponseCodeInvalidTenantID,
		ErrorResponseCodeMissingAppliedScopesForSingle,
		ErrorResponseCodeMissingTenantID,
		ErrorResponseCodeNoValidReservationsToReRate,
		ErrorResponseCodeNonsupportedAccountID,
		ErrorResponseCodeNotSpecified,
		ErrorResponseCodeNotSupportedCountry,
		ErrorResponseCodeOperationCannotBePerformedInCurrentState,
		ErrorResponseCodeOperationFailed,
		ErrorResponseCodePatchValuesSameAsExisting,
		ErrorResponseCodePaymentInstrumentNotFound,
		ErrorResponseCodePurchaseError,
		ErrorResponseCodeReRateOnlyAllowedForEA,
		ErrorResponseCodeRefundLimitExceeded,
		ErrorResponseCodeReservationIDNotInReservationOrder,
		ErrorResponseCodeReservationOrderCreationFailed,
		ErrorResponseCodeReservationOrderIDAlreadyExists,
		ErrorResponseCodeReservationOrderNotEnabled,
		ErrorResponseCodeReservationOrderNotFound,
		ErrorResponseCodeRiskCheckFailed,
		ErrorResponseCodeRoleAssignmentCreationFailed,
		ErrorResponseCodeSelfServiceRefundNotSupported,
		ErrorResponseCodeServerTimeout,
		ErrorResponseCodeUnauthenticatedRequestsThrottled,
		ErrorResponseCodeUnsupportedReservationTerm,
	}
}

// ExchangeOperationResultStatus - Status of the operation.
type ExchangeOperationResultStatus string

const (
	ExchangeOperationResultStatusCancelled        ExchangeOperationResultStatus = "Cancelled"
	ExchangeOperationResultStatusFailed           ExchangeOperationResultStatus = "Failed"
	ExchangeOperationResultStatusPendingPurchases ExchangeOperationResultStatus = "PendingPurchases"
	ExchangeOperationResultStatusPendingRefunds   ExchangeOperationResultStatus = "PendingRefunds"
	ExchangeOperationResultStatusSucceeded        ExchangeOperationResultStatus = "Succeeded"
)

// PossibleExchangeOperationResultStatusValues returns the possible values for the ExchangeOperationResultStatus const type.
func PossibleExchangeOperationResultStatusValues() []ExchangeOperationResultStatus {
	return []ExchangeOperationResultStatus{
		ExchangeOperationResultStatusCancelled,
		ExchangeOperationResultStatusFailed,
		ExchangeOperationResultStatusPendingPurchases,
		ExchangeOperationResultStatusPendingRefunds,
		ExchangeOperationResultStatusSucceeded,
	}
}

// InstanceFlexibility - Turning this on will apply the reservation discount to other VMs in the same VM size group. Only
// specify for VirtualMachines reserved resource type.
type InstanceFlexibility string

const (
	InstanceFlexibilityOff InstanceFlexibility = "Off"
	InstanceFlexibilityOn  InstanceFlexibility = "On"
)

// PossibleInstanceFlexibilityValues returns the possible values for the InstanceFlexibility const type.
func PossibleInstanceFlexibilityValues() []InstanceFlexibility {
	return []InstanceFlexibility{
		InstanceFlexibilityOff,
		InstanceFlexibilityOn,
	}
}

// Location - Location in which the Resources needs to be reserved. It cannot be changed after the resource has been created.
type Location string

const (
	LocationAustraliaeast      Location = "australiaeast"
	LocationAustraliasoutheast Location = "australiasoutheast"
	LocationBrazilsouth        Location = "brazilsouth"
	LocationCanadacentral      Location = "canadacentral"
	LocationCanadaeast         Location = "canadaeast"
	LocationCentralindia       Location = "centralindia"
	LocationCentralus          Location = "centralus"
	LocationEastasia           Location = "eastasia"
	LocationEastus             Location = "eastus"
	LocationEastus2            Location = "eastus2"
	LocationJapaneast          Location = "japaneast"
	LocationJapanwest          Location = "japanwest"
	LocationNorthcentralus     Location = "northcentralus"
	LocationNortheurope        Location = "northeurope"
	LocationSouthcentralus     Location = "southcentralus"
	LocationSoutheastasia      Location = "southeastasia"
	LocationSouthindia         Location = "southindia"
	LocationUksouth            Location = "uksouth"
	LocationUkwest             Location = "ukwest"
	LocationWestcentralus      Location = "westcentralus"
	LocationWesteurope         Location = "westeurope"
	LocationWestindia          Location = "westindia"
	LocationWestus             Location = "westus"
	LocationWestus2            Location = "westus2"
)

// PossibleLocationValues returns the possible values for the Location const type.
func PossibleLocationValues() []Location {
	return []Location{
		LocationAustraliaeast,
		LocationAustraliasoutheast,
		LocationBrazilsouth,
		LocationCanadacentral,
		LocationCanadaeast,
		LocationCentralindia,
		LocationCentralus,
		LocationEastasia,
		LocationEastus,
		LocationEastus2,
		LocationJapaneast,
		LocationJapanwest,
		LocationNorthcentralus,
		LocationNortheurope,
		LocationSouthcentralus,
		LocationSoutheastasia,
		LocationSouthindia,
		LocationUksouth,
		LocationUkwest,
		LocationWestcentralus,
		LocationWesteurope,
		LocationWestindia,
		LocationWestus,
		LocationWestus2,
	}
}

// OperationStatus - Status of the individual operation.
type OperationStatus string

const (
	OperationStatusCancelled OperationStatus = "Cancelled"
	OperationStatusFailed    OperationStatus = "Failed"
	OperationStatusPending   OperationStatus = "Pending"
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// PossibleOperationStatusValues returns the possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{
		OperationStatusCancelled,
		OperationStatusFailed,
		OperationStatusPending,
		OperationStatusSucceeded,
	}
}

// PaymentStatus - Describes whether the payment is completed, failed, cancelled or scheduled in the future.
type PaymentStatus string

const (
	PaymentStatusCancelled PaymentStatus = "Cancelled"
	PaymentStatusFailed    PaymentStatus = "Failed"
	PaymentStatusScheduled PaymentStatus = "Scheduled"
	PaymentStatusSucceeded PaymentStatus = "Succeeded"
)

// PossiblePaymentStatusValues returns the possible values for the PaymentStatus const type.
func PossiblePaymentStatusValues() []PaymentStatus {
	return []PaymentStatus{
		PaymentStatusCancelled,
		PaymentStatusFailed,
		PaymentStatusScheduled,
		PaymentStatusSucceeded,
	}
}

// ProvisioningState - Represent the current state of the Reservation.
type ProvisioningState string

const (
	ProvisioningStateBillingFailed         ProvisioningState = "BillingFailed"
	ProvisioningStateCancelled             ProvisioningState = "Cancelled"
	ProvisioningStateConfirmedBilling      ProvisioningState = "ConfirmedBilling"
	ProvisioningStateConfirmedResourceHold ProvisioningState = "ConfirmedResourceHold"
	ProvisioningStateCreated               ProvisioningState = "Created"
	ProvisioningStateCreating              ProvisioningState = "Creating"
	ProvisioningStateExpired               ProvisioningState = "Expired"
	ProvisioningStateFailed                ProvisioningState = "Failed"
	ProvisioningStateMerged                ProvisioningState = "Merged"
	ProvisioningStatePendingBilling        ProvisioningState = "PendingBilling"
	ProvisioningStatePendingResourceHold   ProvisioningState = "PendingResourceHold"
	ProvisioningStateSplit                 ProvisioningState = "Split"
	ProvisioningStateSucceeded             ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateBillingFailed,
		ProvisioningStateCancelled,
		ProvisioningStateConfirmedBilling,
		ProvisioningStateConfirmedResourceHold,
		ProvisioningStateCreated,
		ProvisioningStateCreating,
		ProvisioningStateExpired,
		ProvisioningStateFailed,
		ProvisioningStateMerged,
		ProvisioningStatePendingBilling,
		ProvisioningStatePendingResourceHold,
		ProvisioningStateSplit,
		ProvisioningStateSucceeded,
	}
}

// QuotaRequestState - The quota request status.
type QuotaRequestState string

const (
	QuotaRequestStateAccepted   QuotaRequestState = "Accepted"
	QuotaRequestStateFailed     QuotaRequestState = "Failed"
	QuotaRequestStateInProgress QuotaRequestState = "InProgress"
	QuotaRequestStateInvalid    QuotaRequestState = "Invalid"
	QuotaRequestStateSucceeded  QuotaRequestState = "Succeeded"
)

// PossibleQuotaRequestStateValues returns the possible values for the QuotaRequestState const type.
func PossibleQuotaRequestStateValues() []QuotaRequestState {
	return []QuotaRequestState{
		QuotaRequestStateAccepted,
		QuotaRequestStateFailed,
		QuotaRequestStateInProgress,
		QuotaRequestStateInvalid,
		QuotaRequestStateSucceeded,
	}
}

// ReservationBillingPlan - Represent the billing plans.
type ReservationBillingPlan string

const (
	ReservationBillingPlanMonthly ReservationBillingPlan = "Monthly"
	ReservationBillingPlanUpfront ReservationBillingPlan = "Upfront"
)

// PossibleReservationBillingPlanValues returns the possible values for the ReservationBillingPlan const type.
func PossibleReservationBillingPlanValues() []ReservationBillingPlan {
	return []ReservationBillingPlan{
		ReservationBillingPlanMonthly,
		ReservationBillingPlanUpfront,
	}
}

type ReservationStatusCode string

const (
	ReservationStatusCodeActive                 ReservationStatusCode = "Active"
	ReservationStatusCodeExpired                ReservationStatusCode = "Expired"
	ReservationStatusCodeMerged                 ReservationStatusCode = "Merged"
	ReservationStatusCodeNone                   ReservationStatusCode = "None"
	ReservationStatusCodePaymentInstrumentError ReservationStatusCode = "PaymentInstrumentError"
	ReservationStatusCodePending                ReservationStatusCode = "Pending"
	ReservationStatusCodeProcessing             ReservationStatusCode = "Processing"
	ReservationStatusCodePurchaseError          ReservationStatusCode = "PurchaseError"
	ReservationStatusCodeSplit                  ReservationStatusCode = "Split"
	ReservationStatusCodeSucceeded              ReservationStatusCode = "Succeeded"
)

// PossibleReservationStatusCodeValues returns the possible values for the ReservationStatusCode const type.
func PossibleReservationStatusCodeValues() []ReservationStatusCode {
	return []ReservationStatusCode{
		ReservationStatusCodeActive,
		ReservationStatusCodeExpired,
		ReservationStatusCodeMerged,
		ReservationStatusCodeNone,
		ReservationStatusCodePaymentInstrumentError,
		ReservationStatusCodePending,
		ReservationStatusCodeProcessing,
		ReservationStatusCodePurchaseError,
		ReservationStatusCodeSplit,
		ReservationStatusCodeSucceeded,
	}
}

// ReservationTerm - Represent the term of reservation.
type ReservationTerm string

const (
	ReservationTermP1Y ReservationTerm = "P1Y"
	ReservationTermP3Y ReservationTerm = "P3Y"
	ReservationTermP5Y ReservationTerm = "P5Y"
)

// PossibleReservationTermValues returns the possible values for the ReservationTerm const type.
func PossibleReservationTermValues() []ReservationTerm {
	return []ReservationTerm{
		ReservationTermP1Y,
		ReservationTermP3Y,
		ReservationTermP5Y,
	}
}

// ReservedResourceType - The type of the resource that is being reserved.
type ReservedResourceType string

const (
	ReservedResourceTypeAVS                    ReservedResourceType = "AVS"
	ReservedResourceTypeAppService             ReservedResourceType = "AppService"
	ReservedResourceTypeAzureDataExplorer      ReservedResourceType = "AzureDataExplorer"
	ReservedResourceTypeAzureFiles             ReservedResourceType = "AzureFiles"
	ReservedResourceTypeBlockBlob              ReservedResourceType = "BlockBlob"
	ReservedResourceTypeCosmosDb               ReservedResourceType = "CosmosDb"
	ReservedResourceTypeDataFactory            ReservedResourceType = "DataFactory"
	ReservedResourceTypeDatabricks             ReservedResourceType = "Databricks"
	ReservedResourceTypeDedicatedHost          ReservedResourceType = "DedicatedHost"
	ReservedResourceTypeManagedDisk            ReservedResourceType = "ManagedDisk"
	ReservedResourceTypeMariaDb                ReservedResourceType = "MariaDb"
	ReservedResourceTypeMySQL                  ReservedResourceType = "MySql"
	ReservedResourceTypeNetAppStorage          ReservedResourceType = "NetAppStorage"
	ReservedResourceTypePostgreSQL             ReservedResourceType = "PostgreSql"
	ReservedResourceTypeRedHat                 ReservedResourceType = "RedHat"
	ReservedResourceTypeRedHatOsa              ReservedResourceType = "RedHatOsa"
	ReservedResourceTypeRedisCache             ReservedResourceType = "RedisCache"
	ReservedResourceTypeSQLAzureHybridBenefit  ReservedResourceType = "SqlAzureHybridBenefit"
	ReservedResourceTypeSQLDataWarehouse       ReservedResourceType = "SqlDataWarehouse"
	ReservedResourceTypeSQLDatabases           ReservedResourceType = "SqlDatabases"
	ReservedResourceTypeSQLEdge                ReservedResourceType = "SqlEdge"
	ReservedResourceTypeSapHana                ReservedResourceType = "SapHana"
	ReservedResourceTypeSuseLinux              ReservedResourceType = "SuseLinux"
	ReservedResourceTypeVMwareCloudSimple      ReservedResourceType = "VMwareCloudSimple"
	ReservedResourceTypeVirtualMachineSoftware ReservedResourceType = "VirtualMachineSoftware"
	ReservedResourceTypeVirtualMachines        ReservedResourceType = "VirtualMachines"
)

// PossibleReservedResourceTypeValues returns the possible values for the ReservedResourceType const type.
func PossibleReservedResourceTypeValues() []ReservedResourceType {
	return []ReservedResourceType{
		ReservedResourceTypeAVS,
		ReservedResourceTypeAppService,
		ReservedResourceTypeAzureDataExplorer,
		ReservedResourceTypeAzureFiles,
		ReservedResourceTypeBlockBlob,
		ReservedResourceTypeCosmosDb,
		ReservedResourceTypeDataFactory,
		ReservedResourceTypeDatabricks,
		ReservedResourceTypeDedicatedHost,
		ReservedResourceTypeManagedDisk,
		ReservedResourceTypeMariaDb,
		ReservedResourceTypeMySQL,
		ReservedResourceTypeNetAppStorage,
		ReservedResourceTypePostgreSQL,
		ReservedResourceTypeRedHat,
		ReservedResourceTypeRedHatOsa,
		ReservedResourceTypeRedisCache,
		ReservedResourceTypeSQLAzureHybridBenefit,
		ReservedResourceTypeSQLDataWarehouse,
		ReservedResourceTypeSQLDatabases,
		ReservedResourceTypeSQLEdge,
		ReservedResourceTypeSapHana,
		ReservedResourceTypeSuseLinux,
		ReservedResourceTypeVMwareCloudSimple,
		ReservedResourceTypeVirtualMachineSoftware,
		ReservedResourceTypeVirtualMachines,
	}
}

// ResourceType - The resource types.
type ResourceType string

const (
	ResourceTypeDedicated       ResourceType = "dedicated"
	ResourceTypeLowPriority     ResourceType = "lowPriority"
	ResourceTypeServiceSpecific ResourceType = "serviceSpecific"
	ResourceTypeShared          ResourceType = "shared"
	ResourceTypeStandard        ResourceType = "standard"
)

// PossibleResourceTypeValues returns the possible values for the ResourceType const type.
func PossibleResourceTypeValues() []ResourceType {
	return []ResourceType{
		ResourceTypeDedicated,
		ResourceTypeLowPriority,
		ResourceTypeServiceSpecific,
		ResourceTypeShared,
		ResourceTypeStandard,
	}
}

// SavingsPlanTerm - Represent savings plan term in ISO 8601 format.
type SavingsPlanTerm string

const (
	SavingsPlanTermP1Y SavingsPlanTerm = "P1Y"
	SavingsPlanTermP3Y SavingsPlanTerm = "P3Y"
)

// PossibleSavingsPlanTermValues returns the possible values for the SavingsPlanTerm const type.
func PossibleSavingsPlanTermValues() []SavingsPlanTerm {
	return []SavingsPlanTerm{
		SavingsPlanTermP1Y,
		SavingsPlanTermP3Y,
	}
}

// UserFriendlyAppliedScopeType - The applied scope type
type UserFriendlyAppliedScopeType string

const (
	UserFriendlyAppliedScopeTypeManagementGroup UserFriendlyAppliedScopeType = "ManagementGroup"
	UserFriendlyAppliedScopeTypeNone            UserFriendlyAppliedScopeType = "None"
	UserFriendlyAppliedScopeTypeResourceGroup   UserFriendlyAppliedScopeType = "ResourceGroup"
	UserFriendlyAppliedScopeTypeShared          UserFriendlyAppliedScopeType = "Shared"
	UserFriendlyAppliedScopeTypeSingle          UserFriendlyAppliedScopeType = "Single"
)

// PossibleUserFriendlyAppliedScopeTypeValues returns the possible values for the UserFriendlyAppliedScopeType const type.
func PossibleUserFriendlyAppliedScopeTypeValues() []UserFriendlyAppliedScopeType {
	return []UserFriendlyAppliedScopeType{
		UserFriendlyAppliedScopeTypeManagementGroup,
		UserFriendlyAppliedScopeTypeNone,
		UserFriendlyAppliedScopeTypeResourceGroup,
		UserFriendlyAppliedScopeTypeShared,
		UserFriendlyAppliedScopeTypeSingle,
	}
}

// UserFriendlyRenewState - The renew state of the reservation
type UserFriendlyRenewState string

const (
	UserFriendlyRenewStateNotApplicable UserFriendlyRenewState = "NotApplicable"
	UserFriendlyRenewStateNotRenewed    UserFriendlyRenewState = "NotRenewed"
	UserFriendlyRenewStateOff           UserFriendlyRenewState = "Off"
	UserFriendlyRenewStateOn            UserFriendlyRenewState = "On"
	UserFriendlyRenewStateRenewed       UserFriendlyRenewState = "Renewed"
)

// PossibleUserFriendlyRenewStateValues returns the possible values for the UserFriendlyRenewState const type.
func PossibleUserFriendlyRenewStateValues() []UserFriendlyRenewState {
	return []UserFriendlyRenewState{
		UserFriendlyRenewStateNotApplicable,
		UserFriendlyRenewStateNotRenewed,
		UserFriendlyRenewStateOff,
		UserFriendlyRenewStateOn,
		UserFriendlyRenewStateRenewed,
	}
}
