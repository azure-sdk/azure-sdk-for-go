//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscom

const (
	moduleName    = "armscom"
	moduleVersion = "v0.1.0"
)

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

// HybridLicenseType - License type for various SCOM infra components
type HybridLicenseType string

const (
	HybridLicenseTypeAzureHybridBenefit HybridLicenseType = "AzureHybridBenefit"
	HybridLicenseTypeNone               HybridLicenseType = "None"
)

// PossibleHybridLicenseTypeValues returns the possible values for the HybridLicenseType const type.
func PossibleHybridLicenseTypeValues() []HybridLicenseType {
	return []HybridLicenseType{
		HybridLicenseTypeAzureHybridBenefit,
		HybridLicenseTypeNone,
	}
}

// ManagedIdentityType - The identity type
type ManagedIdentityType string

const (
	ManagedIdentityTypeNone                       ManagedIdentityType = "None"
	ManagedIdentityTypeSystemAssigned             ManagedIdentityType = "SystemAssigned"
	ManagedIdentityTypeSystemAssignedUserAssigned ManagedIdentityType = "SystemAssigned,UserAssigned"
	ManagedIdentityTypeUserAssigned               ManagedIdentityType = "UserAssigned"
)

// PossibleManagedIdentityTypeValues returns the possible values for the ManagedIdentityType const type.
func PossibleManagedIdentityTypeValues() []ManagedIdentityType {
	return []ManagedIdentityType{
		ManagedIdentityTypeNone,
		ManagedIdentityTypeSystemAssigned,
		ManagedIdentityTypeSystemAssignedUserAssigned,
		ManagedIdentityTypeUserAssigned,
	}
}

// OperationOrigin - The intended executor of the operation.
type OperationOrigin string

const (
	OperationOriginNotSpecified OperationOrigin = "NotSpecified"
	OperationOriginSystem       OperationOrigin = "system"
	OperationOriginUser         OperationOrigin = "user"
)

// PossibleOperationOriginValues returns the possible values for the OperationOrigin const type.
func PossibleOperationOriginValues() []OperationOrigin {
	return []OperationOrigin{
		OperationOriginNotSpecified,
		OperationOriginSystem,
		OperationOriginUser,
	}
}
