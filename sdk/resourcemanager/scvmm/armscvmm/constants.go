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

// AllocationMethod - Network address allocation method.
type AllocationMethod string

const (
	// AllocationMethodDynamic - Dynamically allocated address.
	AllocationMethodDynamic AllocationMethod = "Dynamic"
	// AllocationMethodStatic - Statically allocated address.
	AllocationMethodStatic AllocationMethod = "Static"
)

// PossibleAllocationMethodValues returns the possible values for the AllocationMethod const type.
func PossibleAllocationMethodValues() []AllocationMethod {
	return []AllocationMethod{
		AllocationMethodDynamic,
		AllocationMethodStatic,
	}
}

// CreateDiffDisk - Create diff disk.
type CreateDiffDisk string

const (
	// CreateDiffDiskFalse - Disable create diff disk.
	CreateDiffDiskFalse CreateDiffDisk = "false"
	// CreateDiffDiskTrue - Enable create diff disk.
	CreateDiffDiskTrue CreateDiffDisk = "true"
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
	// DeleteFromHostFalse - Disable delete from host.
	DeleteFromHostFalse DeleteFromHost = "false"
	// DeleteFromHostTrue - Enable delete from host.
	DeleteFromHostTrue DeleteFromHost = "true"
)

// PossibleDeleteFromHostValues returns the possible values for the DeleteFromHost const type.
func PossibleDeleteFromHostValues() []DeleteFromHost {
	return []DeleteFromHost{
		DeleteFromHostFalse,
		DeleteFromHostTrue,
	}
}

// DynamicMemoryEnabled - Dynamic memory enabled.
type DynamicMemoryEnabled string

const (
	// DynamicMemoryEnabledFalse - Disable dynamic memory.
	DynamicMemoryEnabledFalse DynamicMemoryEnabled = "false"
	// DynamicMemoryEnabledTrue - Enable dynamic memory.
	DynamicMemoryEnabledTrue DynamicMemoryEnabled = "true"
)

// PossibleDynamicMemoryEnabledValues returns the possible values for the DynamicMemoryEnabled const type.
func PossibleDynamicMemoryEnabledValues() []DynamicMemoryEnabled {
	return []DynamicMemoryEnabled{
		DynamicMemoryEnabledFalse,
		DynamicMemoryEnabledTrue,
	}
}

type ForceDelete string

const (
	// ForceDeleteFalse - Disable force delete.
	ForceDeleteFalse ForceDelete = "false"
	// ForceDeleteTrue - Enable force delete.
	ForceDeleteTrue ForceDelete = "true"
)

// PossibleForceDeleteValues returns the possible values for the ForceDelete const type.
func PossibleForceDeleteValues() []ForceDelete {
	return []ForceDelete{
		ForceDeleteFalse,
		ForceDeleteTrue,
	}
}

// InventoryType - The inventory type
type InventoryType string

const (
	// InventoryTypeCloud - Cloud inventory type
	InventoryTypeCloud InventoryType = "Cloud"
	// InventoryTypeVirtualMachine - VirtualMachine inventory type
	InventoryTypeVirtualMachine InventoryType = "VirtualMachine"
	// InventoryTypeVirtualMachineTemplate - VirtualMachineTemplate inventory type
	InventoryTypeVirtualMachineTemplate InventoryType = "VirtualMachineTemplate"
	// InventoryTypeVirtualNetwork - VirtualNetwork inventory type
	InventoryTypeVirtualNetwork InventoryType = "VirtualNetwork"
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

// IsCustomizable - Customizable.
type IsCustomizable string

const (
	// IsCustomizableFalse - Disable customizable.
	IsCustomizableFalse IsCustomizable = "false"
	// IsCustomizableTrue - Enable customizable.
	IsCustomizableTrue IsCustomizable = "true"
)

// PossibleIsCustomizableValues returns the possible values for the IsCustomizable const type.
func PossibleIsCustomizableValues() []IsCustomizable {
	return []IsCustomizable{
		IsCustomizableFalse,
		IsCustomizableTrue,
	}
}

// IsHighlyAvailable - Highly available.
type IsHighlyAvailable string

const (
	// IsHighlyAvailableFalse - Disable highly available.
	IsHighlyAvailableFalse IsHighlyAvailable = "false"
	// IsHighlyAvailableTrue - Enable highly available.
	IsHighlyAvailableTrue IsHighlyAvailable = "true"
)

// PossibleIsHighlyAvailableValues returns the possible values for the IsHighlyAvailable const type.
func PossibleIsHighlyAvailableValues() []IsHighlyAvailable {
	return []IsHighlyAvailable{
		IsHighlyAvailableFalse,
		IsHighlyAvailableTrue,
	}
}

// LimitCPUForMigration - Limit CPU for migration.
type LimitCPUForMigration string

const (
	// LimitCPUForMigrationFalse - Disable limit CPU for migration.
	LimitCPUForMigrationFalse LimitCPUForMigration = "false"
	// LimitCPUForMigrationTrue - Enable limit CPU for migration.
	LimitCPUForMigrationTrue LimitCPUForMigration = "true"
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

// OsType - Virtual machine operating system type.
type OsType string

const (
	// OsTypeLinux - Linux operating system.
	OsTypeLinux OsType = "Linux"
	// OsTypeOther - Other operating system.
	OsTypeOther OsType = "Other"
	// OsTypeWindows - Windows operating system.
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

// ProvisioningAction - Guest agent provisioning action.
type ProvisioningAction string

const (
	// ProvisioningActionInstall - Install guest agent.
	ProvisioningActionInstall ProvisioningAction = "install"
	// ProvisioningActionRepair - Repair guest agent.
	ProvisioningActionRepair ProvisioningAction = "repair"
	// ProvisioningActionUninstall - Uninstall guest agent.
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

// ResourceProvisioningState - The provisioning state of the resource.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateAccepted - The resource has been accepted.
	ResourceProvisioningStateAccepted ResourceProvisioningState = "Accepted"
	// ResourceProvisioningStateCanceled - Resource creation was canceled.
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateCreated - The resource was created.
	ResourceProvisioningStateCreated ResourceProvisioningState = "Created"
	// ResourceProvisioningStateDeleting - The resource is being deleted.
	ResourceProvisioningStateDeleting ResourceProvisioningState = "Deleting"
	// ResourceProvisioningStateFailed - Resource creation failed.
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateProvisioning - The resource is provisioning.
	ResourceProvisioningStateProvisioning ResourceProvisioningState = "Provisioning"
	// ResourceProvisioningStateSucceeded - Resource has been created.
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
	// ResourceProvisioningStateUpdating - The resource is updating.
	ResourceProvisioningStateUpdating ResourceProvisioningState = "Updating"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateAccepted,
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateCreated,
		ResourceProvisioningStateDeleting,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateProvisioning,
		ResourceProvisioningStateSucceeded,
		ResourceProvisioningStateUpdating,
	}
}

// SkipShutdown - Gets or sets a value indicating whether to request non-graceful VM shutdown. True value for this flag indicates
// non-graceful shutdown whereas false indicates otherwise. Defaults to false.
type SkipShutdown string

const (
	// SkipShutdownFalse - Disable skip shutdown.
	SkipShutdownFalse SkipShutdown = "false"
	// SkipShutdownTrue - Enable skip shutdown.
	SkipShutdownTrue SkipShutdown = "true"
)

// PossibleSkipShutdownValues returns the possible values for the SkipShutdown const type.
func PossibleSkipShutdownValues() []SkipShutdown {
	return []SkipShutdown{
		SkipShutdownFalse,
		SkipShutdownTrue,
	}
}
