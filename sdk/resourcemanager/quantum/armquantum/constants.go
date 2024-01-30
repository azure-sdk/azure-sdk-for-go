//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/quantum/armquantum"
	moduleVersion = "v0.8.0"
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

type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSecondary KeyType = "Secondary"
)

// PossibleKeyTypeValues returns the possible values for the KeyType const type.
func PossibleKeyTypeValues() []KeyType {
	return []KeyType{
		KeyTypePrimary,
		KeyTypeSecondary,
	}
}

// ProvisioningStatus - Provisioning status field
type ProvisioningStatus string

const (
	ProvisioningStatusFailed               ProvisioningStatus = "Failed"
	ProvisioningStatusProviderDeleting     ProvisioningStatus = "ProviderDeleting"
	ProvisioningStatusProviderLaunching    ProvisioningStatus = "ProviderLaunching"
	ProvisioningStatusProviderProvisioning ProvisioningStatus = "ProviderProvisioning"
	ProvisioningStatusProviderUpdating     ProvisioningStatus = "ProviderUpdating"
	ProvisioningStatusSucceeded            ProvisioningStatus = "Succeeded"
)

// PossibleProvisioningStatusValues returns the possible values for the ProvisioningStatus const type.
func PossibleProvisioningStatusValues() []ProvisioningStatus {
	return []ProvisioningStatus{
		ProvisioningStatusFailed,
		ProvisioningStatusProviderDeleting,
		ProvisioningStatusProviderLaunching,
		ProvisioningStatusProviderProvisioning,
		ProvisioningStatusProviderUpdating,
		ProvisioningStatusSucceeded,
	}
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
	}
}

// Status - Provisioning status field
type Status string

const (
	StatusDeleted   Status = "Deleted"
	StatusDeleting  Status = "Deleting"
	StatusFailed    Status = "Failed"
	StatusLaunching Status = "Launching"
	StatusSucceeded Status = "Succeeded"
	StatusUpdating  Status = "Updating"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusDeleted,
		StatusDeleting,
		StatusFailed,
		StatusLaunching,
		StatusSucceeded,
		StatusUpdating,
	}
}

// UsableStatus - Whether the current workspace is ready to accept Jobs.
type UsableStatus string

const (
	UsableStatusNo      UsableStatus = "No"
	UsableStatusPartial UsableStatus = "Partial"
	UsableStatusYes     UsableStatus = "Yes"
)

// PossibleUsableStatusValues returns the possible values for the UsableStatus const type.
func PossibleUsableStatusValues() []UsableStatus {
	return []UsableStatus{
		UsableStatusNo,
		UsableStatusPartial,
		UsableStatusYes,
	}
}
