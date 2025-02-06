//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
	moduleVersion = "v0.9.0"
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

// CheckNameAvailabilityReason - The reason why the given name is not available.
type CheckNameAvailabilityReason string

const (
	CheckNameAvailabilityReasonAlreadyExists CheckNameAvailabilityReason = "AlreadyExists"
	CheckNameAvailabilityReasonInvalid       CheckNameAvailabilityReason = "Invalid"
)

// PossibleCheckNameAvailabilityReasonValues returns the possible values for the CheckNameAvailabilityReason const type.
func PossibleCheckNameAvailabilityReasonValues() []CheckNameAvailabilityReason {
	return []CheckNameAvailabilityReason{
		CheckNameAvailabilityReasonAlreadyExists,
		CheckNameAvailabilityReasonInvalid,
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

// KeyType - The API key type.
type KeyType string

const (
	// KeyTypePrimary - The primary API key.
	KeyTypePrimary KeyType = "Primary"
	// KeyTypeSecondary - The secondary API key.
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimary,
		KeyTypeSecondary,
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

// ProviderStatus - Provisioning status field
type ProviderStatus string

const (
	// ProviderStatusDeleted - The provider is deleted.
	ProviderStatusDeleted ProviderStatus = "Deleted"
	// ProviderStatusDeleting - The provider is deleting.
	ProviderStatusDeleting ProviderStatus = "Deleting"
	// ProviderStatusFailed - The provider is failed.
	ProviderStatusFailed ProviderStatus = "Failed"
	// ProviderStatusLaunching - The provider is starting provisioning.
	ProviderStatusLaunching ProviderStatus = "Launching"
	// ProviderStatusSucceeded - The provider is successfully provisioned.
	ProviderStatusSucceeded ProviderStatus = "Succeeded"
	// ProviderStatusUpdating - The provider is updating.
	ProviderStatusUpdating ProviderStatus = "Updating"
)

// PossibleProviderStatusValues returns the possible values for the ProviderStatus const type.
func PossibleProviderStatusValues() []ProviderStatus {
	return []ProviderStatus{
		ProviderStatusDeleted,
		ProviderStatusDeleting,
		ProviderStatusFailed,
		ProviderStatusLaunching,
		ProviderStatusSucceeded,
		ProviderStatusUpdating,
	}
}

// UsableStatus - Whether the current workspace is ready to accept Jobs.
type UsableStatus string

const (
	// UsableStatusNo - The workspace is not usable and cannot accept jobs.
	UsableStatusNo UsableStatus = "No"
	// UsableStatusPartial - The workspace is partially usable.
	UsableStatusPartial UsableStatus = "Partial"
	// UsableStatusYes - The workspace is usable and can accept jobs.
	UsableStatusYes UsableStatus = "Yes"
)

// PossibleUsableStatusValues returns the possible values for the UsableStatus const type.
func PossibleUsableStatusValues() []UsableStatus {
	return []UsableStatus{
		UsableStatusNo,
		UsableStatusPartial,
		UsableStatusYes,
	}
}

// WorkspaceProvisioningStatus - The Workspace provisioning status.
type WorkspaceProvisioningStatus string

const (
	// WorkspaceProvisioningStatusCanceled - Resource creation was canceled.
	WorkspaceProvisioningStatusCanceled WorkspaceProvisioningStatus = "Canceled"
	// WorkspaceProvisioningStatusFailed - The Workspace provisioning failed.
	WorkspaceProvisioningStatusFailed WorkspaceProvisioningStatus = "Failed"
	// WorkspaceProvisioningStatusProviderDeleting - The Workspace is currently deleting a provider.
	WorkspaceProvisioningStatusProviderDeleting WorkspaceProvisioningStatus = "ProviderDeleting"
	// WorkspaceProvisioningStatusProviderLaunching - The Workspace is currently starting to provision a provider.
	WorkspaceProvisioningStatusProviderLaunching WorkspaceProvisioningStatus = "ProviderLaunching"
	// WorkspaceProvisioningStatusProviderProvisioning - The Workspace is currently provisioning a provider.
	WorkspaceProvisioningStatusProviderProvisioning WorkspaceProvisioningStatus = "ProviderProvisioning"
	// WorkspaceProvisioningStatusProviderUpdating - The Workspace is currently updating a provider.
	WorkspaceProvisioningStatusProviderUpdating WorkspaceProvisioningStatus = "ProviderUpdating"
	// WorkspaceProvisioningStatusSucceeded - The Workspace provisioning is succeeded.
	WorkspaceProvisioningStatusSucceeded WorkspaceProvisioningStatus = "Succeeded"
)

// PossibleWorkspaceProvisioningStatusValues returns the possible values for the WorkspaceProvisioningStatus const type.
func PossibleWorkspaceProvisioningStatusValues() []WorkspaceProvisioningStatus {
	return []WorkspaceProvisioningStatus{
		WorkspaceProvisioningStatusCanceled,
		WorkspaceProvisioningStatusFailed,
		WorkspaceProvisioningStatusProviderDeleting,
		WorkspaceProvisioningStatusProviderLaunching,
		WorkspaceProvisioningStatusProviderProvisioning,
		WorkspaceProvisioningStatusProviderUpdating,
		WorkspaceProvisioningStatusSucceeded,
	}
}
