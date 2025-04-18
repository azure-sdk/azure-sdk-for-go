// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armextensions

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/kubernetesconfiguration/armextensions"
	moduleVersion = "v0.1.0"
)

// AKSIdentityType - The identity type.
type AKSIdentityType string

const (
	AKSIdentityTypeSystemAssigned AKSIdentityType = "SystemAssigned"
	AKSIdentityTypeUserAssigned   AKSIdentityType = "UserAssigned"
)

// PossibleAKSIdentityTypeValues returns the possible values for the AKSIdentityType const type.
func PossibleAKSIdentityTypeValues() []AKSIdentityType {
	return []AKSIdentityType{
		AKSIdentityTypeSystemAssigned,
		AKSIdentityTypeUserAssigned,
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

// LevelType - Level of the status.
type LevelType string

const (
	LevelTypeError       LevelType = "Error"
	LevelTypeInformation LevelType = "Information"
	LevelTypeWarning     LevelType = "Warning"
)

// PossibleLevelTypeValues returns the possible values for the LevelType const type.
func PossibleLevelTypeValues() []LevelType {
	return []LevelType{
		LevelTypeError,
		LevelTypeInformation,
		LevelTypeWarning,
	}
}

// ProvisioningState - The provisioning state of the resource.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}
