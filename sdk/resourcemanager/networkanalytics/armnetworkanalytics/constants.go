//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkanalytics

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkanalytics/armnetworkanalytics"
	moduleVersion = "v1.0.1"
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

// ControlState - The data type state
type ControlState string

const (
	// ControlStateDisabled - Field to disable a setting.
	ControlStateDisabled ControlState = "Disabled"
	// ControlStateEnabled - Field to enable a setting.
	ControlStateEnabled ControlState = "Enabled"
)

// PossibleControlStateValues returns the possible values for the ControlState const type.
func PossibleControlStateValues() []ControlState {
	return []ControlState{
		ControlStateDisabled,
		ControlStateEnabled,
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

// DataProductUserRole - The data type state
type DataProductUserRole string

const (
	// DataProductUserRoleReader - Field to specify user of type Reader.
	DataProductUserRoleReader DataProductUserRole = "Reader"
	// DataProductUserRoleSensitiveReader - Field to specify user of type SensitiveReader.
	// This user has privileged access to read sensitive data of a data product.
	DataProductUserRoleSensitiveReader DataProductUserRole = "SensitiveReader"
)

// PossibleDataProductUserRoleValues returns the possible values for the DataProductUserRole const type.
func PossibleDataProductUserRoleValues() []DataProductUserRole {
	return []DataProductUserRole{
		DataProductUserRoleReader,
		DataProductUserRoleSensitiveReader,
	}
}

// DataTypeState - The data type state
type DataTypeState string

const (
	// DataTypeStateRunning - Field to specify running state.
	DataTypeStateRunning DataTypeState = "Running"
	// DataTypeStateStopped - Field to specify stopped state.
	DataTypeStateStopped DataTypeState = "Stopped"
)

// PossibleDataTypeStateValues returns the possible values for the DataTypeState const type.
func PossibleDataTypeStateValues() []DataTypeState {
	return []DataTypeState{
		DataTypeStateRunning,
		DataTypeStateStopped,
	}
}

// DefaultAction - Specifies the default action of allow or deny when no other rules match.
type DefaultAction string

const (
	// DefaultActionAllow - Represents allow action.
	DefaultActionAllow DefaultAction = "Allow"
	// DefaultActionDeny - Represents deny action.
	DefaultActionDeny DefaultAction = "Deny"
)

// PossibleDefaultActionValues returns the possible values for the DefaultAction const type.
func PossibleDefaultActionValues() []DefaultAction {
	return []DefaultAction{
		DefaultActionAllow,
		DefaultActionDeny,
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

// ProvisioningState - The status of the current operation.
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Represents an accepted operation.
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Represents a canceled operation.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Represents an operation under deletion.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Represents a failed operation.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - Represents a pending operation.
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Represents a succeeded operation.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Represents a pending operation.
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
