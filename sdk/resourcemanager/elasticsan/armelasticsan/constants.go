//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armelasticsan

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/elasticsan/armelasticsan"
	moduleVersion = "v1.0.0"
)

// Action - The action of virtual network rule.
type Action string

const (
	ActionAllow Action = "Allow"
)

// PossibleActionValues returns the possible values for the Action const type.
func PossibleActionValues() []Action {
	return []Action{
		ActionAllow,
	}
}

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

// EncryptionType - The type of key used to encrypt the data of the disk.
type EncryptionType string

const (
	// EncryptionTypeEncryptionAtRestWithCustomerManagedKey - Volume is encrypted at rest with Customer managed key that can be
	// changed and revoked by a customer.
	EncryptionTypeEncryptionAtRestWithCustomerManagedKey EncryptionType = "EncryptionAtRestWithCustomerManagedKey"
	// EncryptionTypeEncryptionAtRestWithPlatformKey - Volume is encrypted at rest with Platform managed key. It is the default
	// encryption type.
	EncryptionTypeEncryptionAtRestWithPlatformKey EncryptionType = "EncryptionAtRestWithPlatformKey"
)

// PossibleEncryptionTypeValues returns the possible values for the EncryptionType const type.
func PossibleEncryptionTypeValues() []EncryptionType {
	return []EncryptionType{
		EncryptionTypeEncryptionAtRestWithCustomerManagedKey,
		EncryptionTypeEncryptionAtRestWithPlatformKey,
	}
}

// IdentityType - The identity type.
type IdentityType string

const (
	IdentityTypeNone           IdentityType = "None"
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
		IdentityTypeUserAssigned,
	}
}

// OperationalStatus - Operational status of the resource.
type OperationalStatus string

const (
	OperationalStatusHealthy            OperationalStatus = "Healthy"
	OperationalStatusInvalid            OperationalStatus = "Invalid"
	OperationalStatusRunning            OperationalStatus = "Running"
	OperationalStatusStopped            OperationalStatus = "Stopped"
	OperationalStatusStoppedDeallocated OperationalStatus = "Stopped (deallocated)"
	OperationalStatusUnhealthy          OperationalStatus = "Unhealthy"
	OperationalStatusUnknown            OperationalStatus = "Unknown"
	OperationalStatusUpdating           OperationalStatus = "Updating"
)

// PossibleOperationalStatusValues returns the possible values for the OperationalStatus const type.
func PossibleOperationalStatusValues() []OperationalStatus {
	return []OperationalStatus{
		OperationalStatusHealthy,
		OperationalStatusInvalid,
		OperationalStatusRunning,
		OperationalStatusStopped,
		OperationalStatusStoppedDeallocated,
		OperationalStatusUnhealthy,
		OperationalStatusUnknown,
		OperationalStatusUpdating,
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

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusFailed   PrivateEndpointServiceConnectionStatus = "Failed"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusFailed,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ProvisioningStates - Provisioning state of the iSCSI Target.
type ProvisioningStates string

const (
	ProvisioningStatesCanceled  ProvisioningStates = "Canceled"
	ProvisioningStatesCreating  ProvisioningStates = "Creating"
	ProvisioningStatesDeleting  ProvisioningStates = "Deleting"
	ProvisioningStatesFailed    ProvisioningStates = "Failed"
	ProvisioningStatesInvalid   ProvisioningStates = "Invalid"
	ProvisioningStatesPending   ProvisioningStates = "Pending"
	ProvisioningStatesSucceeded ProvisioningStates = "Succeeded"
	ProvisioningStatesUpdating  ProvisioningStates = "Updating"
)

// PossibleProvisioningStatesValues returns the possible values for the ProvisioningStates const type.
func PossibleProvisioningStatesValues() []ProvisioningStates {
	return []ProvisioningStates{
		ProvisioningStatesCanceled,
		ProvisioningStatesCreating,
		ProvisioningStatesDeleting,
		ProvisioningStatesFailed,
		ProvisioningStatesInvalid,
		ProvisioningStatesPending,
		ProvisioningStatesSucceeded,
		ProvisioningStatesUpdating,
	}
}

// PublicNetworkAccess - Allow or disallow public network access to ElasticSan. Value is optional but if passed in, must be
// 'Enabled' or 'Disabled'.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// SKUName - The sku name.
type SKUName string

const (
	// SKUNamePremiumLRS - Premium locally redundant storage
	SKUNamePremiumLRS SKUName = "Premium_LRS"
	// SKUNamePremiumZRS - Premium zone redundant storage
	SKUNamePremiumZRS SKUName = "Premium_ZRS"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNamePremiumLRS,
		SKUNamePremiumZRS,
	}
}

// SKUTier - The sku tier.
type SKUTier string

const (
	// SKUTierPremium - Premium Tier
	SKUTierPremium SKUTier = "Premium"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierPremium,
	}
}

// StorageTargetType - Storage Target type.
type StorageTargetType string

const (
	StorageTargetTypeIscsi StorageTargetType = "Iscsi"
	StorageTargetTypeNone  StorageTargetType = "None"
)

// PossibleStorageTargetTypeValues returns the possible values for the StorageTargetType const type.
func PossibleStorageTargetTypeValues() []StorageTargetType {
	return []StorageTargetType{
		StorageTargetTypeIscsi,
		StorageTargetTypeNone,
	}
}

// VolumeCreateOption - This enumerates the possible sources of a volume creation.
type VolumeCreateOption string

const (
	VolumeCreateOptionDisk             VolumeCreateOption = "Disk"
	VolumeCreateOptionDiskRestorePoint VolumeCreateOption = "DiskRestorePoint"
	VolumeCreateOptionDiskSnapshot     VolumeCreateOption = "DiskSnapshot"
	VolumeCreateOptionNone             VolumeCreateOption = "None"
	VolumeCreateOptionVolumeSnapshot   VolumeCreateOption = "VolumeSnapshot"
)

// PossibleVolumeCreateOptionValues returns the possible values for the VolumeCreateOption const type.
func PossibleVolumeCreateOptionValues() []VolumeCreateOption {
	return []VolumeCreateOption{
		VolumeCreateOptionDisk,
		VolumeCreateOptionDiskRestorePoint,
		VolumeCreateOptionDiskSnapshot,
		VolumeCreateOptionNone,
		VolumeCreateOptionVolumeSnapshot,
	}
}

type XMSDeleteSnapshots string

const (
	XMSDeleteSnapshotsFalse XMSDeleteSnapshots = "false"
	XMSDeleteSnapshotsTrue  XMSDeleteSnapshots = "true"
)

// PossibleXMSDeleteSnapshotsValues returns the possible values for the XMSDeleteSnapshots const type.
func PossibleXMSDeleteSnapshotsValues() []XMSDeleteSnapshots {
	return []XMSDeleteSnapshots{
		XMSDeleteSnapshotsFalse,
		XMSDeleteSnapshotsTrue,
	}
}

type XMSForceDelete string

const (
	XMSForceDeleteFalse XMSForceDelete = "false"
	XMSForceDeleteTrue  XMSForceDelete = "true"
)

// PossibleXMSForceDeleteValues returns the possible values for the XMSForceDelete const type.
func PossibleXMSForceDeleteValues() []XMSForceDelete {
	return []XMSForceDelete{
		XMSForceDeleteFalse,
		XMSForceDeleteTrue,
	}
}
