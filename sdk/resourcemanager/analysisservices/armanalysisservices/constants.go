//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/analysisservices/armanalysisservices"
	moduleVersion = "v2.0.0"
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

type ConnectionMode string

const (
	ConnectionModeAll      ConnectionMode = "All"
	ConnectionModeReadOnly ConnectionMode = "ReadOnly"
)

// PossibleConnectionModeValues returns the possible values for the ConnectionMode const type.
func PossibleConnectionModeValues() []ConnectionMode {
	return []ConnectionMode{
		ConnectionModeAll,
		ConnectionModeReadOnly,
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

// GatewayListStatusLiveStatus - Live message of list gateway. Status: 0 - Live
type GatewayListStatusLiveStatus float32

const (
	GatewayListStatusLiveStatusZero GatewayListStatusLiveStatus = 0
)

// PossibleGatewayListStatusLiveStatusValues returns the possible values for the GatewayListStatusLiveStatus const type.
func PossibleGatewayListStatusLiveStatusValues() []GatewayListStatusLiveStatus {
	return []GatewayListStatusLiveStatus{
		GatewayListStatusLiveStatusZero,
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

type ProvisioningState string

const (
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStatePaused       ProvisioningState = "Paused"
	ProvisioningStatePausing      ProvisioningState = "Pausing"
	ProvisioningStatePreparing    ProvisioningState = "Preparing"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateResuming     ProvisioningState = "Resuming"
	ProvisioningStateScaling      ProvisioningState = "Scaling"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateSuspended    ProvisioningState = "Suspended"
	ProvisioningStateSuspending   ProvisioningState = "Suspending"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStatePaused,
		ProvisioningStatePausing,
		ProvisioningStatePreparing,
		ProvisioningStateProvisioning,
		ProvisioningStateResuming,
		ProvisioningStateScaling,
		ProvisioningStateSucceeded,
		ProvisioningStateSuspended,
		ProvisioningStateSuspending,
		ProvisioningStateUpdating,
	}
}

// SKUTier - This field is required to be implemented by the Resource Provider if the service has more than one tier, but
// is not required on a PUT.
type SKUTier string

const (
	SKUTierBasic       SKUTier = "Basic"
	SKUTierDevelopment SKUTier = "Development"
	SKUTierFree        SKUTier = "Free"
	SKUTierPremium     SKUTier = "Premium"
	SKUTierStandard    SKUTier = "Standard"
)

// PossibleSKUTierValues returns the possible values for the SKUTier const type.
func PossibleSKUTierValues() []SKUTier {
	return []SKUTier{
		SKUTierBasic,
		SKUTierDevelopment,
		SKUTierFree,
		SKUTierPremium,
		SKUTierStandard,
	}
}

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

// Versions - The available API versions.
type Versions string

const (
	VersionsV20170801 Versions = "2017-08-01"
)

// PossibleVersionsValues returns the possible values for the Versions const type.
func PossibleVersionsValues() []Versions {
	return []Versions{
		VersionsV20170801,
	}
}
