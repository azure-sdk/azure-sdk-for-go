//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstandbypool

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/standbypool/armstandbypool"
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

// ProvisioningState - Provisioning state
type ProvisioningState string

const (
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Resource is being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
	}
}

// RefillPolicy - Refill policy of standby pool
type RefillPolicy string

const (
	// RefillPolicyAlways - A refill policy that standby pool is automatically refilled to maintain maxReadyCapacity.
	RefillPolicyAlways RefillPolicy = "always"
)

// PossibleRefillPolicyValues returns the possible values for the RefillPolicy const type.
func PossibleRefillPolicyValues() []RefillPolicy {
	return []RefillPolicy{
		RefillPolicyAlways,
	}
}

// VirtualMachineState - State of standby virtual machines
type VirtualMachineState string

const (
	// VirtualMachineStateDeallocated - The virtual machine has released the lease on the underlying hardware and is powered off.
	VirtualMachineStateDeallocated VirtualMachineState = "Deallocated"
	// VirtualMachineStateRunning - The virtual machine is up and running.
	VirtualMachineStateRunning VirtualMachineState = "Running"
)

// PossibleVirtualMachineStateValues returns the possible values for the VirtualMachineState const type.
func PossibleVirtualMachineStateValues() []VirtualMachineState {
	return []VirtualMachineState{
		VirtualMachineStateDeallocated,
		VirtualMachineStateRunning,
	}
}
