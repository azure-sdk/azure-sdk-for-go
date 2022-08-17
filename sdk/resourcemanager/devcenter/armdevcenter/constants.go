//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armdevcenter

const (
	moduleName    = "armdevcenter"
	moduleVersion = "v0.1.1"
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

// DomainJoinType - Active Directory join type
type DomainJoinType string

const (
	DomainJoinTypeAzureADJoin       DomainJoinType = "AzureADJoin"
	DomainJoinTypeHybridAzureADJoin DomainJoinType = "HybridAzureADJoin"
)

// PossibleDomainJoinTypeValues returns the possible values for the DomainJoinType const type.
func PossibleDomainJoinTypeValues() []DomainJoinType {
	return []DomainJoinType{
		DomainJoinTypeAzureADJoin,
		DomainJoinTypeHybridAzureADJoin,
	}
}

// EnableStatus - Enable or disable status. Indicates whether the property applied to is either enabled or disabled.
type EnableStatus string

const (
	EnableStatusDisabled EnableStatus = "Disabled"
	EnableStatusEnabled  EnableStatus = "Enabled"
)

// PossibleEnableStatusValues returns the possible values for the EnableStatus const type.
func PossibleEnableStatusValues() []EnableStatus {
	return []EnableStatus{
		EnableStatusDisabled,
		EnableStatusEnabled,
	}
}

// HealthCheckStatus - Health check status values
type HealthCheckStatus string

const (
	HealthCheckStatusFailed  HealthCheckStatus = "Failed"
	HealthCheckStatusPassed  HealthCheckStatus = "Passed"
	HealthCheckStatusPending HealthCheckStatus = "Pending"
	HealthCheckStatusRunning HealthCheckStatus = "Running"
	HealthCheckStatusUnknown HealthCheckStatus = "Unknown"
	HealthCheckStatusWarning HealthCheckStatus = "Warning"
)

// PossibleHealthCheckStatusValues returns the possible values for the HealthCheckStatus const type.
func PossibleHealthCheckStatusValues() []HealthCheckStatus {
	return []HealthCheckStatus{
		HealthCheckStatusFailed,
		HealthCheckStatusPassed,
		HealthCheckStatusPending,
		HealthCheckStatusRunning,
		HealthCheckStatusUnknown,
		HealthCheckStatusWarning,
	}
}

// ImageValidationStatus - Image validation status
type ImageValidationStatus string

const (
	ImageValidationStatusFailed    ImageValidationStatus = "Failed"
	ImageValidationStatusPending   ImageValidationStatus = "Pending"
	ImageValidationStatusSucceeded ImageValidationStatus = "Succeeded"
	ImageValidationStatusTimedOut  ImageValidationStatus = "TimedOut"
	ImageValidationStatusUnknown   ImageValidationStatus = "Unknown"
)

// PossibleImageValidationStatusValues returns the possible values for the ImageValidationStatus const type.
func PossibleImageValidationStatusValues() []ImageValidationStatus {
	return []ImageValidationStatus{
		ImageValidationStatusFailed,
		ImageValidationStatusPending,
		ImageValidationStatusSucceeded,
		ImageValidationStatusTimedOut,
		ImageValidationStatusUnknown,
	}
}

// LicenseType - License Types
type LicenseType string

const (
	LicenseTypeWindowsClient LicenseType = "Windows_Client"
)

// PossibleLicenseTypeValues returns the possible values for the LicenseType const type.
func PossibleLicenseTypeValues() []LicenseType {
	return []LicenseType{
		LicenseTypeWindowsClient,
	}
}

type LocalAdminStatus string

const (
	LocalAdminStatusDisabled LocalAdminStatus = "Disabled"
	LocalAdminStatusEnabled  LocalAdminStatus = "Enabled"
)

// PossibleLocalAdminStatusValues returns the possible values for the LocalAdminStatus const type.
func PossibleLocalAdminStatusValues() []LocalAdminStatus {
	return []LocalAdminStatus{
		LocalAdminStatusDisabled,
		LocalAdminStatusEnabled,
	}
}

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned, UserAssigned"
	ManagedServiceIdentityTypeUserAssigned               ManagedServiceIdentityType = "UserAssigned"
)

// PossibleManagedServiceIdentityTypeValues returns the possible values for the ManagedServiceIdentityType const type.
func PossibleManagedServiceIdentityTypeValues() []ManagedServiceIdentityType {
	return []ManagedServiceIdentityType{
		ManagedServiceIdentityTypeNone,
		ManagedServiceIdentityTypeSystemAssigned,
		ManagedServiceIdentityTypeSystemAssignedUserAssigned,
		ManagedServiceIdentityTypeUserAssigned,
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

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierFree     SKUTier = "Free"
	SKUTierBasic    SKUTier = "Basic"
	SKUTierStandard SKUTier = "Standard"
	SKUTierPremium  SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierFree,
		SKUTierBasic,
		SKUTierStandard,
		SKUTierPremium,
	}
}

// ScheduledFrequency - The frequency of task execution.
type ScheduledFrequency string

const (
	ScheduledFrequencyDaily ScheduledFrequency = "Daily"
)

// PossibleScheduledFrequencyValues returns the possible values for the ScheduledFrequency const type.
func PossibleScheduledFrequencyValues() []ScheduledFrequency {
	return []ScheduledFrequency{
		ScheduledFrequencyDaily,
	}
}

// ScheduledType - The supported types for a scheduled task.
type ScheduledType string

const (
	ScheduledTypeStopDevBox ScheduledType = "StopDevBox"
)

// PossibleScheduledTypeValues returns the possible values for the ScheduledType const type.
func PossibleScheduledTypeValues() []ScheduledType {
	return []ScheduledType{
		ScheduledTypeStopDevBox,
	}
}

// UsageUnit - The unit details.
type UsageUnit string

const (
	UsageUnitCount UsageUnit = "Count"
)

// PossibleUsageUnitValues returns the possible values for the UsageUnit const type.
func PossibleUsageUnitValues() []UsageUnit {
	return []UsageUnit{
		UsageUnitCount,
	}
}
