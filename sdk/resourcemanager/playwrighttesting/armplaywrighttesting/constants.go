//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armplaywrighttesting

const (
	moduleName    = "armplaywrighttesting"
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

// EnablementStatus - The enablement status of a feature.
type EnablementStatus string

const (
	// EnablementStatusDisabled - The feature is Disabled.
	EnablementStatusDisabled EnablementStatus = "Disabled"
	// EnablementStatusEnabled - The feature is Enabled.
	EnablementStatusEnabled EnablementStatus = "Enabled"
)

// PossibleEnablementStatusValues returns the possible values for the EnablementStatus const type.
func PossibleEnablementStatusValues() []EnablementStatus {
	return []EnablementStatus{
		EnablementStatusDisabled,
		EnablementStatusEnabled,
	}
}

// FreeTrialState - The free-trial state.
type FreeTrialState string

const (
	// FreeTrialStateActive - The free-trial is Active.
	FreeTrialStateActive FreeTrialState = "Active"
	// FreeTrialStateExpired - The free-trial is Expired.
	FreeTrialStateExpired FreeTrialState = "Expired"
)

// PossibleFreeTrialStateValues returns the possible values for the FreeTrialState const type.
func PossibleFreeTrialStateValues() []FreeTrialState {
	return []FreeTrialState{
		FreeTrialStateActive,
		FreeTrialStateExpired,
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
	// ProvisioningStateAccepted - Change accepted for processing
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Deletion in progress
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

type QuotaNames string

const (
	// QuotaNamesScalableExecution - The quota details for scalable execution feature. When enabled, Playwright client workers
	// can connect to cloud-hosted browsers. This can increase the number of parallel workers for a test run, significantly minimizing
	// test completion durations.
	QuotaNamesScalableExecution QuotaNames = "ScalableExecution"
)

// PossibleQuotaNamesValues returns the possible values for the QuotaNames const type.
func PossibleQuotaNamesValues() []QuotaNames {
	return []QuotaNames{
		QuotaNamesScalableExecution,
	}
}
