//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevopsinfrastructure

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/devopsinfrastructure/armdevopsinfrastructure"
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

// AzureDevOpsPermissionType - Determines who has admin permissions to the Azure DevOps pool.
type AzureDevOpsPermissionType string

const (
	// AzureDevOpsPermissionTypeCreatorOnly - Only the pool creator will be an admin of the pool.
	AzureDevOpsPermissionTypeCreatorOnly AzureDevOpsPermissionType = "CreatorOnly"
	// AzureDevOpsPermissionTypeInherit - Pool will inherit permissions from the project or organization.
	AzureDevOpsPermissionTypeInherit AzureDevOpsPermissionType = "Inherit"
	// AzureDevOpsPermissionTypeSpecificAccounts - Only the specified accounts will be admins of the pool.
	AzureDevOpsPermissionTypeSpecificAccounts AzureDevOpsPermissionType = "SpecificAccounts"
)

// PossibleAzureDevOpsPermissionTypeValues returns the possible values for the AzureDevOpsPermissionType const type.
func PossibleAzureDevOpsPermissionTypeValues() []AzureDevOpsPermissionType {
	return []AzureDevOpsPermissionType{
		AzureDevOpsPermissionTypeCreatorOnly,
		AzureDevOpsPermissionTypeInherit,
		AzureDevOpsPermissionTypeSpecificAccounts,
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

// LogonType - Determines how the service should be run.
type LogonType string

const (
	// LogonTypeInteractive - Run in interactive mode.
	LogonTypeInteractive LogonType = "Interactive"
	// LogonTypeService - Run as a service.
	LogonTypeService LogonType = "Service"
)

// PossibleLogonTypeValues returns the possible values for the LogonType const type.
func PossibleLogonTypeValues() []LogonType {
	return []LogonType{
		LogonTypeInteractive,
		LogonTypeService,
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

// OsDiskStorageAccountType - The storage account type of the OS disk.
type OsDiskStorageAccountType string

const (
	// OsDiskStorageAccountTypePremium - Premium OS disk type.
	OsDiskStorageAccountTypePremium OsDiskStorageAccountType = "Premium"
	// OsDiskStorageAccountTypeStandard - Standard OS disk type.
	OsDiskStorageAccountTypeStandard OsDiskStorageAccountType = "Standard"
	// OsDiskStorageAccountTypeStandardSSD - Standard SSD OS disk type.
	OsDiskStorageAccountTypeStandardSSD OsDiskStorageAccountType = "StandardSSD"
)

// PossibleOsDiskStorageAccountTypeValues returns the possible values for the OsDiskStorageAccountType const type.
func PossibleOsDiskStorageAccountTypeValues() []OsDiskStorageAccountType {
	return []OsDiskStorageAccountType{
		OsDiskStorageAccountTypePremium,
		OsDiskStorageAccountTypeStandard,
		OsDiskStorageAccountTypeStandardSSD,
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
