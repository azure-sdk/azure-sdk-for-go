//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrustedsigning

import "time"

// AccountSKU - SKU of the trusted signing account.
type AccountSKU struct {
	// REQUIRED; Name of the SKU.
	Name *SKUName
}

// Certificate - Properties of the certificate.
type Certificate struct {
	// Certificate created date.
	CreatedDate *string

	// Certificate expiry date.
	ExpiryDate *string

	// Revocations history of a certificate.
	Revocation *Revocation

	// Serial number of the certificate.
	SerialNumber *string

	// Status of the certificate.
	Status *CertificateStatus

	// Subject name of the certificate.
	SubjectName *string

	// Thumbprint of the certificate.
	Thumbprint *string
}

// CertificateProfile - Certificate profile resource.
type CertificateProfile struct {
	// The resource-specific properties for this resource.
	Properties *CertificateProfileProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CertificateProfileListResult - The response of a CertificateProfile list operation.
type CertificateProfileListResult struct {
	// REQUIRED; The CertificateProfile items on this page
	Value []*CertificateProfile

	// The link to the next page of items
	NextLink *string
}

// CertificateProfileProperties - Properties of the certificate profile.
type CertificateProfileProperties struct {
	// REQUIRED; Profile type of the certificate.
	ProfileType *ProfileType

	// Identity validation id used for the certificate subject name.
	IdentityValidationID *string

	// Whether to include L in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCity *bool

	// Whether to include C in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeCountry *bool

	// Whether to include PC in the certificate subject name.
	IncludePostalCode *bool

	// Whether to include S in the certificate subject name. Applicable only for private trust, private trust ci profile types
	IncludeState *bool

	// Whether to include STREET in the certificate subject name.
	IncludeStreetAddress *bool

	// READ-ONLY; List of renewed certificates.
	Certificates []*Certificate

	// READ-ONLY; Used as L in the certificate subject name.
	City *string

	// READ-ONLY; Used as CN in the certificate subject name.
	CommonName *string

	// READ-ONLY; Used as C in the certificate subject name.
	Country *string

	// READ-ONLY; Enhanced key usage of the certificate.
	EnhancedKeyUsage *string

	// READ-ONLY; Used as O in the certificate subject name.
	Organization *string

	// READ-ONLY; Used as OU in the private trust certificate subject name.
	OrganizationUnit *string

	// READ-ONLY; Used as PC in the certificate subject name.
	PostalCode *string

	// READ-ONLY; Status of the current operation on certificate profile.
	ProvisioningState *ProvisioningState

	// READ-ONLY; Used as S in the certificate subject name.
	State *string

	// READ-ONLY; Status of the certificate profile.
	Status *CertificateProfileStatus

	// READ-ONLY; Used as STREET in the certificate subject name.
	StreetAddress *string
}

// CheckNameAvailability - The parameters used to check the availability of the trusted signing account name.
type CheckNameAvailability struct {
	// REQUIRED; Trusted signing account name.
	Name *string
}

// CheckNameAvailabilityResult - The CheckNameAvailability operation response.
type CheckNameAvailabilityResult struct {
	// READ-ONLY; An error message explaining the Reason value in more detail.
	Message *string

	// READ-ONLY; A boolean value that indicates whether the name is available for you to use. If true, the name is available.
	// If false, the name has already been taken or is invalid and cannot be used.
	NameAvailable *bool

	// READ-ONLY; The reason that a trusted signing account name could not be used. The Reason element is only returned if nameAvailable
	// is false.
	Reason *NameUnavailabilityReason
}

// CodeSigningAccount - Trusted signing account resource.
type CodeSigningAccount struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *CodeSigningAccountProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// CodeSigningAccountListResult - The response of a CodeSigningAccount list operation.
type CodeSigningAccountListResult struct {
	// REQUIRED; The CodeSigningAccount items on this page
	Value []*CodeSigningAccount

	// The link to the next page of items
	NextLink *string
}

// CodeSigningAccountPatch - Parameters for creating or updating a trusted signing account.
type CodeSigningAccountPatch struct {
	// Properties of the trusted signing account.
	Properties *CodeSigningAccountPatchProperties

	// Resource tags.
	Tags map[string]*string
}

// CodeSigningAccountPatchProperties - Properties of the trusted signing account.
type CodeSigningAccountPatchProperties struct {
	// SKU of the trusted signing account.
	SKU *AccountSKU
}

// CodeSigningAccountProperties - Properties of the trusted signing account.
type CodeSigningAccountProperties struct {
	// SKU of the trusted signing account.
	SKU *AccountSKU

	// READ-ONLY; The URI of the trusted signing account which is used during signing files.
	AccountURI *string

	// READ-ONLY; Status of the current operation on trusted signing account.
	ProvisioningState *ProvisioningState
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// Revocation details of the certificate.
type Revocation struct {
	// The timestamp when the revocation is effective.
	EffectiveAt *time.Time

	// Reason for the revocation failure.
	FailureReason *string

	// Reason for revocation.
	Reason *string

	// Remarks for the revocation.
	Remarks *string

	// The timestamp when the revocation is requested.
	RequestedAt *time.Time

	// READ-ONLY; Status of the revocation.
	Status *RevocationStatus
}

// RevokeCertificate - Defines the certificate revocation properties.
type RevokeCertificate struct {
	// REQUIRED; The timestamp when the revocation is effective.
	EffectiveAt *time.Time

	// REQUIRED; Reason for the revocation.
	Reason *string

	// REQUIRED; Serial number of the certificate.
	SerialNumber *string

	// REQUIRED; Thumbprint of the certificate.
	Thumbprint *string

	// Remarks for the revocation.
	Remarks *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}
