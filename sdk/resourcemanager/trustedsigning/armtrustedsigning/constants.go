//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrustedsigning

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trustedsigning/armtrustedsigning"
	moduleVersion = "v0.1.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CertificateProfileStatus - Status of the certificate profiles.
type CertificateProfileStatus string

const (
	// CertificateProfileStatusActive - The certificate profile is active.
	CertificateProfileStatusActive CertificateProfileStatus = "Active"
	// CertificateProfileStatusDisabled - The certificate profile is disabled.
	CertificateProfileStatusDisabled CertificateProfileStatus = "Disabled"
	// CertificateProfileStatusSuspended - The certificate profile is suspended.
	CertificateProfileStatusSuspended CertificateProfileStatus = "Suspended"
)

// PossibleCertificateProfileStatusValues returns the possible values for the CertificateProfileStatus const type.
func PossibleCertificateProfileStatusValues() []CertificateProfileStatus {
	return []CertificateProfileStatus{
		CertificateProfileStatusActive,
		CertificateProfileStatusDisabled,
		CertificateProfileStatusSuspended,
	}
}

// CertificateStatus - Status of the certificate
type CertificateStatus string

const (
	// CertificateStatusActive - The certificate is active.
	CertificateStatusActive CertificateStatus = "Active"
	// CertificateStatusExpired - The certificate is expired.
	CertificateStatusExpired CertificateStatus = "Expired"
	// CertificateStatusRevoked - The certificate is revoked.
	CertificateStatusRevoked CertificateStatus = "Revoked"
)

// PossibleCertificateStatusValues returns the possible values for the CertificateStatus const type.
func PossibleCertificateStatusValues() []CertificateStatus {
	return []CertificateStatus{
		CertificateStatusActive,
		CertificateStatusExpired,
		CertificateStatusRevoked,
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

// NameUnavailabilityReason - The reason that a trusted signing account name could not be used. The Reason element is only
// returned if nameAvailable is false.
type NameUnavailabilityReason string

const (
	// NameUnavailabilityReasonAccountNameInvalid - Account name is invalid
	NameUnavailabilityReasonAccountNameInvalid NameUnavailabilityReason = "AccountNameInvalid"
	// NameUnavailabilityReasonAlreadyExists - Account name already exists
	NameUnavailabilityReasonAlreadyExists NameUnavailabilityReason = "AlreadyExists"
)

// PossibleNameUnavailabilityReasonValues returns the possible values for the NameUnavailabilityReason const type.
func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return []NameUnavailabilityReason{
		NameUnavailabilityReasonAccountNameInvalid,
		NameUnavailabilityReasonAlreadyExists,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ProfileType - Type of the certificate
type ProfileType string

const (
	// ProfileTypePrivateTrust - Used for signing files which are distributed internally within organization or group boundary.
	ProfileTypePrivateTrust ProfileType = "PrivateTrust"
	// ProfileTypePrivateTrustCIPolicy - Used for signing CI policy files.
	ProfileTypePrivateTrustCIPolicy ProfileType = "PrivateTrustCIPolicy"
	// ProfileTypePublicTrust - Used for signing files which are distributed publicly.
	ProfileTypePublicTrust ProfileType = "PublicTrust"
	// ProfileTypePublicTrustTest - Used for signing files for testing purpose.
	ProfileTypePublicTrustTest ProfileType = "PublicTrustTest"
	// ProfileTypeVBSEnclave - Used for signing files which are run in secure vbs enclave.
	ProfileTypeVBSEnclave ProfileType = "VBSEnclave"
)

// PossibleProfileTypeValues returns the possible values for the ProfileType const type.
func PossibleProfileTypeValues() []ProfileType {
	return []ProfileType{
		ProfileTypePrivateTrust,
		ProfileTypePrivateTrustCIPolicy,
		ProfileTypePublicTrust,
		ProfileTypePublicTrustTest,
		ProfileTypeVBSEnclave,
	}
}

// ProvisioningState - The status of the current operation.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Resource creation started.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Deletion in progress.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Updating in progress.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RevocationStatus - Revocation status of the certificate.
type RevocationStatus string

const (
	// RevocationStatusFailed - Certificate revocation failed.
	RevocationStatusFailed RevocationStatus = "Failed"
	// RevocationStatusInProgress - Certificate revocation is in progress.
	RevocationStatusInProgress RevocationStatus = "InProgress"
	// RevocationStatusSucceeded - Certificate revocation succeeded.
	RevocationStatusSucceeded RevocationStatus = "Succeeded"
)

// PossibleRevocationStatusValues returns the possible values for the RevocationStatus const type.
func PossibleRevocationStatusValues() []RevocationStatus {
	return []RevocationStatus{
		RevocationStatusFailed,
		RevocationStatusInProgress,
		RevocationStatusSucceeded,
	}
}

// SKUName - Name of the sku.
type SKUName string

const (
	// SKUNameBasic - Basic sku.
	SKUNameBasic SKUName = "Basic"
	// SKUNamePremium - Premium sku.
	SKUNamePremium SKUName = "Premium"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameBasic,
		SKUNamePremium,
	}
}
