//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armazuresphere

const (
	moduleName    = "armazuresphere"
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

// AllowCrashDumpCollection - Allow crash dumps values.
type AllowCrashDumpCollection string

const (
	// AllowCrashDumpCollectionDisabled - Crash dump collection disabled
	AllowCrashDumpCollectionDisabled AllowCrashDumpCollection = "Disabled"
	// AllowCrashDumpCollectionEnabled - Crash dump collection enabled
	AllowCrashDumpCollectionEnabled AllowCrashDumpCollection = "Enabled"
)

// PossibleAllowCrashDumpCollectionValues returns the possible values for the AllowCrashDumpCollection const type.
func PossibleAllowCrashDumpCollectionValues() []AllowCrashDumpCollection {
	return []AllowCrashDumpCollection{
		AllowCrashDumpCollectionDisabled,
		AllowCrashDumpCollectionEnabled,
	}
}

// CapabilityType - Capability image type
type CapabilityType string

const (
	// CapabilityTypeApplicationDevelopment - Application development capability
	CapabilityTypeApplicationDevelopment CapabilityType = "ApplicationDevelopment"
	// CapabilityTypeFieldServicing - Field servicing capability
	CapabilityTypeFieldServicing CapabilityType = "FieldServicing"
)

// PossibleCapabilityTypeValues returns the possible values for the CapabilityType const type.
func PossibleCapabilityTypeValues() []CapabilityType {
	return []CapabilityType{
		CapabilityTypeApplicationDevelopment,
		CapabilityTypeFieldServicing,
	}
}

// CertificateStatus - Certificate status values.
type CertificateStatus string

const (
	// CertificateStatusActive - Certificate is active
	CertificateStatusActive CertificateStatus = "Active"
	// CertificateStatusExpired - Certificate has expired
	CertificateStatusExpired CertificateStatus = "Expired"
	// CertificateStatusInactive - Certificate is inactive
	CertificateStatusInactive CertificateStatus = "Inactive"
	// CertificateStatusRevoked - Certificate has been revoked
	CertificateStatusRevoked CertificateStatus = "Revoked"
)

// PossibleCertificateStatusValues returns the possible values for the CertificateStatus const type.
func PossibleCertificateStatusValues() []CertificateStatus {
	return []CertificateStatus{
		CertificateStatusActive,
		CertificateStatusExpired,
		CertificateStatusInactive,
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

// ImageType - Image type values.
type ImageType string

const (
	// ImageTypeApplications - Applications image type
	ImageTypeApplications ImageType = "Applications"
	// ImageTypeBaseSystemUpdateManifest - Base system update manifest image type
	ImageTypeBaseSystemUpdateManifest ImageType = "BaseSystemUpdateManifest"
	// ImageTypeBootManifest - Boot manifest image type
	ImageTypeBootManifest ImageType = "BootManifest"
	// ImageTypeCustomerBoardConfig - Customer board config image type
	ImageTypeCustomerBoardConfig ImageType = "CustomerBoardConfig"
	// ImageTypeCustomerUpdateManifest - Customer update manifest image type
	ImageTypeCustomerUpdateManifest ImageType = "CustomerUpdateManifest"
	// ImageTypeFirmwareUpdateManifest - Firmware update manifest image type
	ImageTypeFirmwareUpdateManifest ImageType = "FirmwareUpdateManifest"
	// ImageTypeFwConfig - FW config image type
	ImageTypeFwConfig ImageType = "FwConfig"
	// ImageTypeInvalidImageType - Invalid image.
	ImageTypeInvalidImageType ImageType = "InvalidImageType"
	// ImageTypeManifestSet - manifest set image type
	ImageTypeManifestSet ImageType = "ManifestSet"
	// ImageTypeNormalWorldDtb - Normal world dtb image type
	ImageTypeNormalWorldDtb ImageType = "NormalWorldDtb"
	// ImageTypeNormalWorldKernel - Normal world kernel image type
	ImageTypeNormalWorldKernel ImageType = "NormalWorldKernel"
	// ImageTypeNormalWorldLoader - Normal world loader image type
	ImageTypeNormalWorldLoader ImageType = "NormalWorldLoader"
	// ImageTypeNwfs - Nwfs image type
	ImageTypeNwfs ImageType = "Nwfs"
	// ImageTypeOneBl - One Bl image type
	ImageTypeOneBl ImageType = "OneBl"
	// ImageTypeOther - Other image type
	ImageTypeOther ImageType = "Other"
	// ImageTypePlutonRuntime - Pluton image type
	ImageTypePlutonRuntime ImageType = "PlutonRuntime"
	// ImageTypePolicy - Policy image type
	ImageTypePolicy ImageType = "Policy"
	// ImageTypeRecoveryManifest - Recovery manifest image type
	ImageTypeRecoveryManifest ImageType = "RecoveryManifest"
	// ImageTypeRootFs - Root FS image type
	ImageTypeRootFs ImageType = "RootFs"
	// ImageTypeSecurityMonitor - Security monitor image type
	ImageTypeSecurityMonitor ImageType = "SecurityMonitor"
	// ImageTypeServices - Services image type
	ImageTypeServices ImageType = "Services"
	// ImageTypeTrustedKeystore - Trusted key store image type
	ImageTypeTrustedKeystore ImageType = "TrustedKeystore"
	// ImageTypeUpdateCertStore - Update certificate store image type
	ImageTypeUpdateCertStore ImageType = "UpdateCertStore"
	// ImageTypeWifiFirmware - Wifi firmware image type
	ImageTypeWifiFirmware ImageType = "WifiFirmware"
)

// PossibleImageTypeValues returns the possible values for the ImageType const type.
func PossibleImageTypeValues() []ImageType {
	return []ImageType{
		ImageTypeApplications,
		ImageTypeBaseSystemUpdateManifest,
		ImageTypeBootManifest,
		ImageTypeCustomerBoardConfig,
		ImageTypeCustomerUpdateManifest,
		ImageTypeFirmwareUpdateManifest,
		ImageTypeFwConfig,
		ImageTypeInvalidImageType,
		ImageTypeManifestSet,
		ImageTypeNormalWorldDtb,
		ImageTypeNormalWorldKernel,
		ImageTypeNormalWorldLoader,
		ImageTypeNwfs,
		ImageTypeOneBl,
		ImageTypeOther,
		ImageTypePlutonRuntime,
		ImageTypePolicy,
		ImageTypeRecoveryManifest,
		ImageTypeRootFs,
		ImageTypeSecurityMonitor,
		ImageTypeServices,
		ImageTypeTrustedKeystore,
		ImageTypeUpdateCertStore,
		ImageTypeWifiFirmware,
	}
}

// OSFeedType - OS feed type values.
type OSFeedType string

const (
	// OSFeedTypeRetail - Retail OS feed type.
	OSFeedTypeRetail OSFeedType = "Retail"
	// OSFeedTypeRetailEval - Retail evaluation OS feed type.
	OSFeedTypeRetailEval OSFeedType = "RetailEval"
)

// PossibleOSFeedTypeValues returns the possible values for the OSFeedType const type.
func PossibleOSFeedTypeValues() []OSFeedType {
	return []OSFeedType{
		OSFeedTypeRetail,
		OSFeedTypeRetailEval,
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

// ProvisioningState - Provisioning state of the resource.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - The resource create request has been accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - The resource is being deleted
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - The resource is being provisioned
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The resource is being updated
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RegionalDataBoundary - Regional data boundary values.
type RegionalDataBoundary string

const (
	// RegionalDataBoundaryEU - EU data boundary
	RegionalDataBoundaryEU RegionalDataBoundary = "EU"
	// RegionalDataBoundaryNone - No data boundary
	RegionalDataBoundaryNone RegionalDataBoundary = "None"
)

// PossibleRegionalDataBoundaryValues returns the possible values for the RegionalDataBoundary const type.
func PossibleRegionalDataBoundaryValues() []RegionalDataBoundary {
	return []RegionalDataBoundary{
		RegionalDataBoundaryEU,
		RegionalDataBoundaryNone,
	}
}

// UpdatePolicy - Update policy values.
type UpdatePolicy string

const (
	// UpdatePolicyNo3RdPartyAppUpdates - No update for 3rd party app policy.
	UpdatePolicyNo3RdPartyAppUpdates UpdatePolicy = "No3rdPartyAppUpdates"
	// UpdatePolicyUpdateAll - Update all policy.
	UpdatePolicyUpdateAll UpdatePolicy = "UpdateAll"
)

// PossibleUpdatePolicyValues returns the possible values for the UpdatePolicy const type.
func PossibleUpdatePolicyValues() []UpdatePolicy {
	return []UpdatePolicy{
		UpdatePolicyNo3RdPartyAppUpdates,
		UpdatePolicyUpdateAll,
	}
}
