//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/scvmm/armscvmm"
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

// AllocationMethod - Allocation method.
type AllocationMethod string

const (
	AllocationMethodDynamic AllocationMethod = "Dynamic"
	AllocationMethodStatic  AllocationMethod = "Static"
)

// PossibleAllocationMethodValues returns the possible values for the AllocationMethod const type.
func PossibleAllocationMethodValues() []AllocationMethod {
	return []AllocationMethod{
		AllocationMethodDynamic,
		AllocationMethodStatic,
	}
}

// CreateDiffDisk - Gets or sets a value indicating diff disk.
type CreateDiffDisk string

const (
	CreateDiffDiskFalse CreateDiffDisk = "false"
	CreateDiffDiskTrue  CreateDiffDisk = "true"
)

// PossibleCreateDiffDiskValues returns the possible values for the CreateDiffDisk const type.
func PossibleCreateDiffDiskValues() []CreateDiffDisk {
	return []CreateDiffDisk{
		CreateDiffDiskFalse,
		CreateDiffDiskTrue,
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

type DeleteFromHost string

const (
	DeleteFromHostFalse DeleteFromHost = "false"
	DeleteFromHostTrue  DeleteFromHost = "true"
)

// PossibleDeleteFromHostValues returns the possible values for the DeleteFromHost const type.
func PossibleDeleteFromHostValues() []DeleteFromHost {
	return []DeleteFromHost{
		DeleteFromHostFalse,
		DeleteFromHostTrue,
	}
}

// DynamicMemoryEnabled - Gets or sets a value indicating whether to enable dynamic memory or not.
type DynamicMemoryEnabled string

const (
	DynamicMemoryEnabledFalse DynamicMemoryEnabled = "false"
	DynamicMemoryEnabledTrue  DynamicMemoryEnabled = "true"
)

// PossibleDynamicMemoryEnabledValues returns the possible values for the DynamicMemoryEnabled const type.
func PossibleDynamicMemoryEnabledValues() []DynamicMemoryEnabled {
	return []DynamicMemoryEnabled{
		DynamicMemoryEnabledFalse,
		DynamicMemoryEnabledTrue,
	}
}

type Force string

const (
	ForceFalse Force = "false"
	ForceTrue  Force = "true"
)

// PossibleForceValues returns the possible values for the Force const type.
func PossibleForceValues() []Force {
	return []Force{
		ForceFalse,
		ForceTrue,
	}
}

// IdentityType - The type of managed service identity.
type IdentityType string

const (
	IdentityTypeNone           IdentityType = "None"
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeNone,
		IdentityTypeSystemAssigned,
	}
}

// InventoryType - The inventory type.
type InventoryType string

const (
	InventoryTypeCloud                  InventoryType = "Cloud"
	InventoryTypeVirtualMachine         InventoryType = "VirtualMachine"
	InventoryTypeVirtualMachineTemplate InventoryType = "VirtualMachineTemplate"
	InventoryTypeVirtualNetwork         InventoryType = "VirtualNetwork"
)

// PossibleInventoryTypeValues returns the possible values for the InventoryType const type.
func PossibleInventoryTypeValues() []InventoryType {
	return []InventoryType{
		InventoryTypeCloud,
		InventoryTypeVirtualMachine,
		InventoryTypeVirtualMachineTemplate,
		InventoryTypeVirtualNetwork,
	}
}

// IsCustomizable - Gets a value indicating whether the vm template is customizable or not.
type IsCustomizable string

const (
	IsCustomizableFalse IsCustomizable = "false"
	IsCustomizableTrue  IsCustomizable = "true"
)

// PossibleIsCustomizableValues returns the possible values for the IsCustomizable const type.
func PossibleIsCustomizableValues() []IsCustomizable {
	return []IsCustomizable{
		IsCustomizableFalse,
		IsCustomizableTrue,
	}
}

// IsHighlyAvailable - Gets highly available property.
type IsHighlyAvailable string

const (
	IsHighlyAvailableFalse IsHighlyAvailable = "false"
	IsHighlyAvailableTrue  IsHighlyAvailable = "true"
)

// PossibleIsHighlyAvailableValues returns the possible values for the IsHighlyAvailable const type.
func PossibleIsHighlyAvailableValues() []IsHighlyAvailable {
	return []IsHighlyAvailable{
		IsHighlyAvailableFalse,
		IsHighlyAvailableTrue,
	}
}

// LimitCPUForMigration - Gets or sets a value indicating whether to enable processor compatibility mode for live migration
// of VMs.
type LimitCPUForMigration string

const (
	LimitCPUForMigrationFalse LimitCPUForMigration = "false"
	LimitCPUForMigrationTrue  LimitCPUForMigration = "true"
)

// PossibleLimitCPUForMigrationValues returns the possible values for the LimitCPUForMigration const type.
func PossibleLimitCPUForMigrationValues() []LimitCPUForMigration {
	return []LimitCPUForMigration{
		LimitCPUForMigrationFalse,
		LimitCPUForMigrationTrue,
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

// OsType - Defines the different types of VM guest operating systems.
type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeOther   OsType = "Other"
	OsTypeWindows OsType = "Windows"
)

// PossibleOsTypeValues returns the possible values for the OsType const type.
func PossibleOsTypeValues() []OsType {
	return []OsType{
		OsTypeLinux,
		OsTypeOther,
		OsTypeWindows,
	}
}

// ProvisioningAction - Defines the different types of operations for guest agent.
type ProvisioningAction string

const (
	ProvisioningActionInstall   ProvisioningAction = "install"
	ProvisioningActionRepair    ProvisioningAction = "repair"
	ProvisioningActionUninstall ProvisioningAction = "uninstall"
)

// PossibleProvisioningActionValues returns the possible values for the ProvisioningAction const type.
func PossibleProvisioningActionValues() []ProvisioningAction {
	return []ProvisioningAction{
		ProvisioningActionInstall,
		ProvisioningActionRepair,
		ProvisioningActionUninstall,
	}
}

// ProvisioningState - The provisioning state of a resource.
type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreated      ProvisioningState = "Created"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateCreated,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SkipShutdown - Gets or sets a value indicating whether to request non-graceful VM shutdown. True value for this flag indicates
// non-graceful shutdown whereas false indicates otherwise. Defaults to false.
type SkipShutdown string

const (
	SkipShutdownFalse SkipShutdown = "false"
	SkipShutdownTrue  SkipShutdown = "true"
)

// PossibleSkipShutdownValues returns the possible values for the SkipShutdown const type.
func PossibleSkipShutdownValues() []SkipShutdown {
	return []SkipShutdown{
		SkipShutdownFalse,
		SkipShutdownTrue,
	}
}
