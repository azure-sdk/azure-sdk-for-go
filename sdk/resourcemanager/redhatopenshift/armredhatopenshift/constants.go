//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredhatopenshift

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/redhatopenshift/armredhatopenshift"
	moduleVersion = "v1.6.1"
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

// EncryptionAtHost - EncryptionAtHost represents encryption at host state
type EncryptionAtHost string

const (
	EncryptionAtHostDisabled EncryptionAtHost = "Disabled"
	EncryptionAtHostEnabled  EncryptionAtHost = "Enabled"
)

// PossibleEncryptionAtHostValues returns the possible values for the EncryptionAtHost const type.
func PossibleEncryptionAtHostValues() []EncryptionAtHost {
	return []EncryptionAtHost{
		EncryptionAtHostDisabled,
		EncryptionAtHostEnabled,
	}
}

// FipsValidatedModules - FipsValidatedModules determines if FIPS is used.
type FipsValidatedModules string

const (
	FipsValidatedModulesDisabled FipsValidatedModules = "Disabled"
	FipsValidatedModulesEnabled  FipsValidatedModules = "Enabled"
)

// PossibleFipsValidatedModulesValues returns the possible values for the FipsValidatedModules const type.
func PossibleFipsValidatedModulesValues() []FipsValidatedModules {
	return []FipsValidatedModules{
		FipsValidatedModulesDisabled,
		FipsValidatedModulesEnabled,
	}
}

// OutboundType - The outbound routing strategy used to provide your cluster egress to the internet.
type OutboundType string

const (
	OutboundTypeLoadbalancer       OutboundType = "Loadbalancer"
	OutboundTypeUserDefinedRouting OutboundType = "UserDefinedRouting"
)

// PossibleOutboundTypeValues returns the possible values for the OutboundType const type.
func PossibleOutboundTypeValues() []OutboundType {
	return []OutboundType{
		OutboundTypeLoadbalancer,
		OutboundTypeUserDefinedRouting,
	}
}

// PreconfiguredNSG - PreconfiguredNSG represents whether customers want to use their own NSG attached to the subnets
type PreconfiguredNSG string

const (
	PreconfiguredNSGDisabled PreconfiguredNSG = "Disabled"
	PreconfiguredNSGEnabled  PreconfiguredNSG = "Enabled"
)

// PossiblePreconfiguredNSGValues returns the possible values for the PreconfiguredNSG const type.
func PossiblePreconfiguredNSGValues() []PreconfiguredNSG {
	return []PreconfiguredNSG{
		PreconfiguredNSGDisabled,
		PreconfiguredNSGEnabled,
	}
}

// ProvisioningState - ProvisioningState represents a provisioning state.
type ProvisioningState string

const (
	ProvisioningStateAdminUpdating ProvisioningState = "AdminUpdating"
	ProvisioningStateCanceled      ProvisioningState = "Canceled"
	ProvisioningStateCreating      ProvisioningState = "Creating"
	ProvisioningStateDeleting      ProvisioningState = "Deleting"
	ProvisioningStateFailed        ProvisioningState = "Failed"
	ProvisioningStateSucceeded     ProvisioningState = "Succeeded"
	ProvisioningStateUpdating      ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAdminUpdating,
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// Visibility - Visibility represents visibility.
type Visibility string

const (
	VisibilityPrivate Visibility = "Private"
	VisibilityPublic  Visibility = "Public"
)

// PossibleVisibilityValues returns the possible values for the Visibility const type.
func PossibleVisibilityValues() []Visibility {
	return []Visibility{
		VisibilityPrivate,
		VisibilityPublic,
	}
}
