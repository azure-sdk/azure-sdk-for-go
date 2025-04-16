// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlabservices

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
	moduleVersion = "v1.2.0"
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

// ConnectionType - A connection type for access labs and VMs (Public, Private or None).
type ConnectionType string

const (
	ConnectionTypeNone    ConnectionType = "None"
	ConnectionTypePrivate ConnectionType = "Private"
	ConnectionTypePublic  ConnectionType = "Public"
)

// PossibleConnectionTypeValues returns the possible values for the ConnectionType const type.
func PossibleConnectionTypeValues() []ConnectionType {
	return []ConnectionType{
		ConnectionTypeNone,
		ConnectionTypePrivate,
		ConnectionTypePublic,
	}
}

// CreateOption - Indicates what lab virtual machines are created from.
type CreateOption string

const (
	// CreateOptionImage - An image is used to create all lab user virtual machines. When this option is set, no template VM will
	// be created.
	CreateOptionImage CreateOption = "Image"
	// CreateOptionTemplateVM - A template VM will be used to create all lab user virtual machines.
	CreateOptionTemplateVM CreateOption = "TemplateVM"
)

// PossibleCreateOptionValues returns the possible values for the CreateOption const type.
func PossibleCreateOptionValues() []CreateOption {
	return []CreateOption{
		CreateOptionImage,
		CreateOptionTemplateVM,
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

// EnableState - Property enabled state.
type EnableState string

const (
	EnableStateDisabled EnableState = "Disabled"
	EnableStateEnabled  EnableState = "Enabled"
)

// PossibleEnableStateValues returns the possible values for the EnableState const type.
func PossibleEnableStateValues() []EnableState {
	return []EnableState{
		EnableStateDisabled,
		EnableStateEnabled,
	}
}

// InvitationState - The lab user invitation state.
type InvitationState string

const (
	// InvitationStateFailed - There was an error while sending the invitation.
	InvitationStateFailed InvitationState = "Failed"
	// InvitationStateNotSent - The invitation has not been sent.
	InvitationStateNotSent InvitationState = "NotSent"
	// InvitationStateSending - Currently sending the invitation.
	InvitationStateSending InvitationState = "Sending"
	// InvitationStateSent - The invitation has been successfully sent.
	InvitationStateSent InvitationState = "Sent"
)

// PossibleInvitationStateValues returns the possible values for the InvitationState const type.
func PossibleInvitationStateValues() []InvitationState {
	return []InvitationState{
		InvitationStateFailed,
		InvitationStateNotSent,
		InvitationStateSending,
		InvitationStateSent,
	}
}

// LabServicesSKUTier - The tier of the SKU.
type LabServicesSKUTier string

const (
	LabServicesSKUTierPremium  LabServicesSKUTier = "Premium"
	LabServicesSKUTierStandard LabServicesSKUTier = "Standard"
)

// PossibleLabServicesSKUTierValues returns the possible values for the LabServicesSKUTier const type.
func PossibleLabServicesSKUTierValues() []LabServicesSKUTier {
	return []LabServicesSKUTier{
		LabServicesSKUTierPremium,
		LabServicesSKUTierStandard,
	}
}

// LabState - The state of a virtual machine.
type LabState string

const (
	// LabStateDraft - The lab is currently in draft (has not been published).
	LabStateDraft LabState = "Draft"
	// LabStatePublished - The lab has been published.
	LabStatePublished LabState = "Published"
	// LabStatePublishing - The lab is publishing.
	LabStatePublishing LabState = "Publishing"
	// LabStateScaling - The lab is scaling.
	LabStateScaling LabState = "Scaling"
	// LabStateSyncing - The lab is syncing users.
	LabStateSyncing LabState = "Syncing"
)

// PossibleLabStateValues returns the possible values for the LabState const type.
func PossibleLabStateValues() []LabState {
	return []LabState{
		LabStateDraft,
		LabStatePublished,
		LabStatePublishing,
		LabStateScaling,
		LabStateSyncing,
	}
}

// OperationStatus - The operation status
type OperationStatus string

const (
	// OperationStatusCanceled - Not supported yet
	OperationStatusCanceled OperationStatus = "Canceled"
	// OperationStatusFailed - The operation failed
	OperationStatusFailed OperationStatus = "Failed"
	// OperationStatusInProgress - The operation is running
	OperationStatusInProgress OperationStatus = "InProgress"
	// OperationStatusNotStarted - The operation has been accepted but hasn't started.
	OperationStatusNotStarted OperationStatus = "NotStarted"
	// OperationStatusSucceeded - The operation Succeeded
	OperationStatusSucceeded OperationStatus = "Succeeded"
)

// PossibleOperationStatusValues returns the possible values for the OperationStatus const type.
func PossibleOperationStatusValues() []OperationStatus {
	return []OperationStatus{
		OperationStatusCanceled,
		OperationStatusFailed,
		OperationStatusInProgress,
		OperationStatusNotStarted,
		OperationStatusSucceeded,
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

// OsState - The operating system state.
type OsState string

const (
	// OsStateGeneralized - Image does not contain any machine and user specific information.
	OsStateGeneralized OsState = "Generalized"
	// OsStateSpecialized - Image contains machine and user specific information.
	OsStateSpecialized OsState = "Specialized"
)

// PossibleOsStateValues returns the possible values for the OsState const type.
func PossibleOsStateValues() []OsState {
	return []OsState{
		OsStateGeneralized,
		OsStateSpecialized,
	}
}

// OsType - The operating system type.
type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeLinux,
		OsTypeWindows,
	}
}

// ProvisioningState - Resource provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateCreating - Resource is in the process of being created.
	ProvisioningStateCreating ProvisioningState = "Creating"
	// ProvisioningStateDeleting - Resource is in the process of being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Previous operation on the resource has failed leaving resource in unhealthy state.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateLocked - The resource is locked and changes are currently blocked. This could be due to maintenance or
	// a scheduled operation. The state will go back to succeeded once the locking operation has finished.
	ProvisioningStateLocked ProvisioningState = "Locked"
	// ProvisioningStateSucceeded - Resource is in healthy state after creation or update operation.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - New property values are being applied to the resource.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateLocked,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// RecurrenceFrequency - Schedule recurrence frequencies.
type RecurrenceFrequency string

const (
	// RecurrenceFrequencyDaily - Schedule will run every days.
	RecurrenceFrequencyDaily RecurrenceFrequency = "Daily"
	// RecurrenceFrequencyWeekly - Schedule will run every week on days specified in weekDays.
	RecurrenceFrequencyWeekly RecurrenceFrequency = "Weekly"
)

// PossibleRecurrenceFrequencyValues returns the possible values for the RecurrenceFrequency const type.
func PossibleRecurrenceFrequencyValues() []RecurrenceFrequency {
	return []RecurrenceFrequency{
		RecurrenceFrequencyDaily,
		RecurrenceFrequencyWeekly,
	}
}

// RegistrationState - The user lab registration state.
type RegistrationState string

const (
	// RegistrationStateNotRegistered - User has registered with the lab.
	RegistrationStateNotRegistered RegistrationState = "NotRegistered"
	// RegistrationStateRegistered - User has not yet registered with the lab.
	RegistrationStateRegistered RegistrationState = "Registered"
)

// PossibleRegistrationStateValues returns the possible values for the RegistrationState const type.
func PossibleRegistrationStateValues() []RegistrationState {
	return []RegistrationState{
		RegistrationStateNotRegistered,
		RegistrationStateRegistered,
	}
}

// RestrictionReasonCode - The reason for the restriction.
type RestrictionReasonCode string

const (
	RestrictionReasonCodeNotAvailableForSubscription RestrictionReasonCode = "NotAvailableForSubscription"
	RestrictionReasonCodeQuotaID                     RestrictionReasonCode = "QuotaId"
)

// PossibleRestrictionReasonCodeValues returns the possible values for the RestrictionReasonCode const type.
func PossibleRestrictionReasonCodeValues() []RestrictionReasonCode {
	return []RestrictionReasonCode{
		RestrictionReasonCodeNotAvailableForSubscription,
		RestrictionReasonCodeQuotaID,
	}
}

// RestrictionType - The type of restriction.
type RestrictionType string

const (
	RestrictionTypeLocation RestrictionType = "Location"
)

// PossibleRestrictionTypeValues returns the possible values for the RestrictionType const type.
func PossibleRestrictionTypeValues() []RestrictionType {
	return []RestrictionType{
		RestrictionTypeLocation,
	}
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierBasic    SKUTier = "Basic"
	SKUTierFree     SKUTier = "Free"
	SKUTierPremium  SKUTier = "Premium"
	SKUTierStandard SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierFree,
		SKUTierPremium,
		SKUTierStandard,
	}
}

// ScaleType - The localized name of the resource.
type ScaleType string

const (
	// ScaleTypeAutomatic - The user is permitted to scale this SKU in and out.
	ScaleTypeAutomatic ScaleType = "Automatic"
	// ScaleTypeManual - The user must manually scale this SKU in and out.
	ScaleTypeManual ScaleType = "Manual"
	// ScaleTypeNone - The capacity is not adjustable in any way.
	ScaleTypeNone ScaleType = "None"
)

// PossibleScaleTypeValues returns the possible values for the ScaleType const type.
func PossibleScaleTypeValues() []ScaleType {
	return []ScaleType{
		ScaleTypeAutomatic,
		ScaleTypeManual,
		ScaleTypeNone,
	}
}

// ShutdownOnIdleMode - Defines whether to shut down VM on idle and the criteria for idle detection.
type ShutdownOnIdleMode string

const (
	// ShutdownOnIdleModeLowUsage - The VM will be considered as idle when user is absent and the resource (CPU and disk) consumption
	// is low.
	ShutdownOnIdleModeLowUsage ShutdownOnIdleMode = "LowUsage"
	// ShutdownOnIdleModeNone - The VM won't be shut down when it is idle.
	ShutdownOnIdleModeNone ShutdownOnIdleMode = "None"
	// ShutdownOnIdleModeUserAbsence - The VM will be considered as idle when there is no keyboard or mouse input.
	ShutdownOnIdleModeUserAbsence ShutdownOnIdleMode = "UserAbsence"
)

// PossibleShutdownOnIdleModeValues returns the possible values for the ShutdownOnIdleMode const type.
func PossibleShutdownOnIdleModeValues() []ShutdownOnIdleMode {
	return []ShutdownOnIdleMode{
		ShutdownOnIdleModeLowUsage,
		ShutdownOnIdleModeNone,
		ShutdownOnIdleModeUserAbsence,
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

// VirtualMachineState - The state of a virtual machine.
type VirtualMachineState string

const (
	// VirtualMachineStateRedeploying - The VM is being redeployed.
	VirtualMachineStateRedeploying VirtualMachineState = "Redeploying"
	// VirtualMachineStateReimaging - The VM is being reimaged.
	VirtualMachineStateReimaging VirtualMachineState = "Reimaging"
	// VirtualMachineStateResettingPassword - The VM password is being reset.
	VirtualMachineStateResettingPassword VirtualMachineState = "ResettingPassword"
	// VirtualMachineStateRunning - The VM is running.
	VirtualMachineStateRunning VirtualMachineState = "Running"
	// VirtualMachineStateStarting - The VM is starting.
	VirtualMachineStateStarting VirtualMachineState = "Starting"
	// VirtualMachineStateStopped - The VM is currently stopped.
	VirtualMachineStateStopped VirtualMachineState = "Stopped"
	// VirtualMachineStateStopping - The VM is stopping.
	VirtualMachineStateStopping VirtualMachineState = "Stopping"
)

// PossibleVirtualMachineStateValues returns the possible values for the VirtualMachineState const type.
func PossibleVirtualMachineStateValues() []VirtualMachineState {
	return []VirtualMachineState{
		VirtualMachineStateRedeploying,
		VirtualMachineStateReimaging,
		VirtualMachineStateResettingPassword,
		VirtualMachineStateRunning,
		VirtualMachineStateStarting,
		VirtualMachineStateStopped,
		VirtualMachineStateStopping,
	}
}

// VirtualMachineType - The type of the lab virtual machine.
type VirtualMachineType string

const (
	// VirtualMachineTypeTemplate - A template VM
	VirtualMachineTypeTemplate VirtualMachineType = "Template"
	// VirtualMachineTypeUser - A user VM
	VirtualMachineTypeUser VirtualMachineType = "User"
)

// PossibleVirtualMachineTypeValues returns the possible values for the VirtualMachineType const type.
func PossibleVirtualMachineTypeValues() []VirtualMachineType {
	return []VirtualMachineType{
		VirtualMachineTypeTemplate,
		VirtualMachineTypeUser,
	}
}

// WeekDay - Days of the week.
type WeekDay string

const (
	// WeekDayFriday - Schedule will run on Friday
	WeekDayFriday WeekDay = "Friday"
	// WeekDayMonday - Schedule will run on Monday
	WeekDayMonday WeekDay = "Monday"
	// WeekDaySaturday - Schedule will run on Saturday
	WeekDaySaturday WeekDay = "Saturday"
	// WeekDaySunday - Schedule will run on Sunday
	WeekDaySunday WeekDay = "Sunday"
	// WeekDayThursday - Schedule will run on Thursday
	WeekDayThursday WeekDay = "Thursday"
	// WeekDayTuesday - Schedule will run on Tuesday
	WeekDayTuesday WeekDay = "Tuesday"
	// WeekDayWednesday - Schedule will run on Wednesday
	WeekDayWednesday WeekDay = "Wednesday"
)

// PossibleWeekDayValues returns the possible values for the WeekDay const type.
func PossibleWeekDayValues() []WeekDay {
	return []WeekDay{
		WeekDayFriday,
		WeekDayMonday,
		WeekDaySaturday,
		WeekDaySunday,
		WeekDayThursday,
		WeekDayTuesday,
		WeekDayWednesday,
	}
}
