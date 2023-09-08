//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragetasks

const (
	moduleName    = "armstoragetasks"
	moduleVersion = "v1.0.0"
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

// ManagedServiceIdentityType - Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeNone                       ManagedServiceIdentityType = "None"
	ManagedServiceIdentityTypeSystemAssigned             ManagedServiceIdentityType = "SystemAssigned"
	ManagedServiceIdentityTypeSystemAssignedUserAssigned ManagedServiceIdentityType = "SystemAssigned,UserAssigned"
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

// MatchedBlockName - Represents the condition block name that matched blob properties.
type MatchedBlockName string

const (
	MatchedBlockNameElse MatchedBlockName = "Else"
	MatchedBlockNameIf   MatchedBlockName = "If"
	MatchedBlockNameNone MatchedBlockName = "None"
)

// PossibleMatchedBlockNameValues returns the possible values for the MatchedBlockName const type.
func PossibleMatchedBlockNameValues() []MatchedBlockName {
	return []MatchedBlockName{
		MatchedBlockNameElse,
		MatchedBlockNameIf,
		MatchedBlockNameNone,
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

// ProvisioningState - Represents the provisioning state of the storage task.
type ProvisioningState string

const (
	ProvisioningStateCanceled                       ProvisioningState = "Canceled"
	ProvisioningStateCreating                       ProvisioningState = "Creating"
	ProvisioningStateDeleting                       ProvisioningState = "Deleting"
	ProvisioningStateFailed                         ProvisioningState = "Failed"
	ProvisioningStateSucceeded                      ProvisioningState = "Succeeded"
	ProvisioningStateValidateSubscriptionQuotaBegin ProvisioningState = "ValidateSubscriptionQuotaBegin"
	ProvisioningStateValidateSubscriptionQuotaEnd   ProvisioningState = "ValidateSubscriptionQuotaEnd"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateValidateSubscriptionQuotaBegin,
		ProvisioningStateValidateSubscriptionQuotaEnd,
	}
}

// RunResult - Represents the overall result of the execution for the run instance
type RunResult string

const (
	RunResultFailed    RunResult = "Failed"
	RunResultSucceeded RunResult = "Succeeded"
)

// PossibleRunResultValues returns the possible values for the RunResult const type.
func PossibleRunResultValues() []RunResult {
	return []RunResult{
		RunResultFailed,
		RunResultSucceeded,
	}
}

// RunStatusEnum - Represents the status of the execution.
type RunStatusEnum string

const (
	RunStatusEnumFinished   RunStatusEnum = "Finished"
	RunStatusEnumInProgress RunStatusEnum = "InProgress"
)

// PossibleRunStatusEnumValues returns the possible values for the RunStatusEnum const type.
func PossibleRunStatusEnumValues() []RunStatusEnum {
	return []RunStatusEnum{
		RunStatusEnumFinished,
		RunStatusEnumInProgress,
	}
}

// StorageTaskOperationName - The operation to be performed on the object.
type StorageTaskOperationName string

const (
	StorageTaskOperationNameDeleteBlob                StorageTaskOperationName = "DeleteBlob"
	StorageTaskOperationNameSetBlobExpiry             StorageTaskOperationName = "SetBlobExpiry"
	StorageTaskOperationNameSetBlobImmutabilityPolicy StorageTaskOperationName = "SetBlobImmutabilityPolicy"
	StorageTaskOperationNameSetBlobLegalHold          StorageTaskOperationName = "SetBlobLegalHold"
	StorageTaskOperationNameSetBlobTags               StorageTaskOperationName = "SetBlobTags"
	StorageTaskOperationNameSetBlobTier               StorageTaskOperationName = "SetBlobTier"
	StorageTaskOperationNameUndeleteBlob              StorageTaskOperationName = "UndeleteBlob"
)

// PossibleStorageTaskOperationNameValues returns the possible values for the StorageTaskOperationName const type.
func PossibleStorageTaskOperationNameValues() []StorageTaskOperationName {
	return []StorageTaskOperationName{
		StorageTaskOperationNameDeleteBlob,
		StorageTaskOperationNameSetBlobExpiry,
		StorageTaskOperationNameSetBlobImmutabilityPolicy,
		StorageTaskOperationNameSetBlobLegalHold,
		StorageTaskOperationNameSetBlobTags,
		StorageTaskOperationNameSetBlobTier,
		StorageTaskOperationNameUndeleteBlob,
	}
}
