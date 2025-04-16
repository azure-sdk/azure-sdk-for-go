// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbidedicated

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/powerbidedicated/armpowerbidedicated"
	moduleVersion = "v1.2.0"
)

// CapacityProvisioningState - The current deployment state of PowerBI Dedicated resource. The provisioningState is to indicate
// states for resource provisioning.
type CapacityProvisioningState string

const (
	CapacityProvisioningStateDeleting     CapacityProvisioningState = "Deleting"
	CapacityProvisioningStateFailed       CapacityProvisioningState = "Failed"
	CapacityProvisioningStatePaused       CapacityProvisioningState = "Paused"
	CapacityProvisioningStatePausing      CapacityProvisioningState = "Pausing"
	CapacityProvisioningStatePreparing    CapacityProvisioningState = "Preparing"
	CapacityProvisioningStateProvisioning CapacityProvisioningState = "Provisioning"
	CapacityProvisioningStateResuming     CapacityProvisioningState = "Resuming"
	CapacityProvisioningStateScaling      CapacityProvisioningState = "Scaling"
	CapacityProvisioningStateSucceeded    CapacityProvisioningState = "Succeeded"
	CapacityProvisioningStateSuspended    CapacityProvisioningState = "Suspended"
	CapacityProvisioningStateSuspending   CapacityProvisioningState = "Suspending"
	CapacityProvisioningStateUpdating     CapacityProvisioningState = "Updating"
)

// PossibleCapacityProvisioningStateValues returns the possible values for the CapacityProvisioningState const type.
func PossibleCapacityProvisioningStateValues() []CapacityProvisioningState {
	return []CapacityProvisioningState{
		CapacityProvisioningStateDeleting,
		CapacityProvisioningStateFailed,
		CapacityProvisioningStatePaused,
		CapacityProvisioningStatePausing,
		CapacityProvisioningStatePreparing,
		CapacityProvisioningStateProvisioning,
		CapacityProvisioningStateResuming,
		CapacityProvisioningStateScaling,
		CapacityProvisioningStateSucceeded,
		CapacityProvisioningStateSuspended,
		CapacityProvisioningStateSuspending,
		CapacityProvisioningStateUpdating,
	}
}

// CapacitySKUTier - The name of the Azure pricing tier to which the SKU applies.
type CapacitySKUTier string

const (
	CapacitySKUTierAutoPremiumHost CapacitySKUTier = "AutoPremiumHost"
	CapacitySKUTierPBIEAzure       CapacitySKUTier = "PBIE_Azure"
	CapacitySKUTierPremium         CapacitySKUTier = "Premium"
)

// PossibleCapacitySKUTierValues returns the possible values for the CapacitySKUTier const type.
func PossibleCapacitySKUTierValues() []CapacitySKUTier {
	return []CapacitySKUTier{
		CapacitySKUTierAutoPremiumHost,
		CapacitySKUTierPBIEAzure,
		CapacitySKUTierPremium,
	}
}

// IdentityType - The type of identity that created/modified the resource.
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "Application"
	IdentityTypeKey             IdentityType = "Key"
	IdentityTypeManagedIdentity IdentityType = "ManagedIdentity"
	IdentityTypeUser            IdentityType = "User"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// Mode - Specifies the generation of the Power BI Embedded capacity. If no value is specified, the default value 'Gen2' is
// used. Learn More
// [https://docs.microsoft.com/power-bi/developer/embedded/power-bi-embedded-generation-2]
type Mode string

const (
	ModeGen1 Mode = "Gen1"
	ModeGen2 Mode = "Gen2"
)

// PossibleModeValues returns the possible values for the Mode const type.
func PossibleModeValues() []Mode {
	return []Mode{
		ModeGen1,
		ModeGen2,
	}
}

// State - The current state of PowerBI Dedicated resource. The state is to indicate more states outside of resource provisioning.
type State string

const (
	StateDeleting     State = "Deleting"
	StateFailed       State = "Failed"
	StatePaused       State = "Paused"
	StatePausing      State = "Pausing"
	StatePreparing    State = "Preparing"
	StateProvisioning State = "Provisioning"
	StateResuming     State = "Resuming"
	StateScaling      State = "Scaling"
	StateSucceeded    State = "Succeeded"
	StateSuspended    State = "Suspended"
	StateSuspending   State = "Suspending"
	StateUpdating     State = "Updating"
)

// PossibleStateValues returns the possible values for the State const type.
func PossibleStateValues() []State {
	return []State{
		StateDeleting,
		StateFailed,
		StatePaused,
		StatePausing,
		StatePreparing,
		StateProvisioning,
		StateResuming,
		StateScaling,
		StateSucceeded,
		StateSuspended,
		StateSuspending,
		StateUpdating,
	}
}

// VCoreProvisioningState - The current deployment state of an auto scale v-core resource. The provisioningState is to indicate
// states for resource provisioning.
type VCoreProvisioningState string

const (
	VCoreProvisioningStateSucceeded VCoreProvisioningState = "Succeeded"
)

// PossibleVCoreProvisioningStateValues returns the possible values for the VCoreProvisioningState const type.
func PossibleVCoreProvisioningStateValues() []VCoreProvisioningState {
	return []VCoreProvisioningState{
		VCoreProvisioningStateSucceeded,
	}
}

// VCoreSKUTier - The name of the Azure pricing tier to which the SKU applies.
type VCoreSKUTier string

const (
	VCoreSKUTierAutoScale VCoreSKUTier = "AutoScale"
)

// PossibleVCoreSKUTierValues returns the possible values for the VCoreSKUTier const type.
func PossibleVCoreSKUTierValues() []VCoreSKUTier {
	return []VCoreSKUTier{
		VCoreSKUTierAutoScale,
	}
}
