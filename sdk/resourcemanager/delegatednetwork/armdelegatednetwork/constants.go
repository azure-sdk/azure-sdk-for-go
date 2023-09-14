//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdelegatednetwork

const (
	moduleName    = "armdelegatednetwork"
	moduleVersion = "v1.2.0-beta.1"
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

// ControllerPurpose - The purpose of the dnc controller resource.
type ControllerPurpose string

const (
	ControllerPurposeProd ControllerPurpose = "prod"
	ControllerPurposeTest ControllerPurpose = "test"
)

// PossibleControllerPurposeValues returns the possible values for the ControllerPurpose const type.
func PossibleControllerPurposeValues() []ControllerPurpose {
	return []ControllerPurpose{
		ControllerPurposeProd,
		ControllerPurposeTest,
	}
}

// ControllerState - The current state of dnc controller resource.
type ControllerState string

const (
	ControllerStateDeleting     ControllerState = "Deleting"
	ControllerStateFailed       ControllerState = "Failed"
	ControllerStateProvisioning ControllerState = "Provisioning"
	ControllerStateSucceeded    ControllerState = "Succeeded"
)

// PossibleControllerStateValues returns the possible values for the ControllerState const type.
func PossibleControllerStateValues() []ControllerState {
	return []ControllerState{
		ControllerStateDeleting,
		ControllerStateFailed,
		ControllerStateProvisioning,
		ControllerStateSucceeded,
	}
}

// DelegatedSubnetState - The current state of dnc delegated subnet resource.
type DelegatedSubnetState string

const (
	DelegatedSubnetStateDeleting     DelegatedSubnetState = "Deleting"
	DelegatedSubnetStateFailed       DelegatedSubnetState = "Failed"
	DelegatedSubnetStateProvisioning DelegatedSubnetState = "Provisioning"
	DelegatedSubnetStateSucceeded    DelegatedSubnetState = "Succeeded"
)

// PossibleDelegatedSubnetStateValues returns the possible values for the DelegatedSubnetState const type.
func PossibleDelegatedSubnetStateValues() []DelegatedSubnetState {
	return []DelegatedSubnetState{
		DelegatedSubnetStateDeleting,
		DelegatedSubnetStateFailed,
		DelegatedSubnetStateProvisioning,
		DelegatedSubnetStateSucceeded,
	}
}

// OrchestratorInstanceState - The current state of orchestratorInstance resource.
type OrchestratorInstanceState string

const (
	OrchestratorInstanceStateDeleting     OrchestratorInstanceState = "Deleting"
	OrchestratorInstanceStateFailed       OrchestratorInstanceState = "Failed"
	OrchestratorInstanceStateProvisioning OrchestratorInstanceState = "Provisioning"
	OrchestratorInstanceStateSucceeded    OrchestratorInstanceState = "Succeeded"
)

// PossibleOrchestratorInstanceStateValues returns the possible values for the OrchestratorInstanceState const type.
func PossibleOrchestratorInstanceStateValues() []OrchestratorInstanceState {
	return []OrchestratorInstanceState{
		OrchestratorInstanceStateDeleting,
		OrchestratorInstanceStateFailed,
		OrchestratorInstanceStateProvisioning,
		OrchestratorInstanceStateSucceeded,
	}
}

// OrchestratorKind - The kind of workbook. Choices are user and shared.
type OrchestratorKind string

const (
	OrchestratorKindKubernetes OrchestratorKind = "Kubernetes"
)

// PossibleOrchestratorKindValues returns the possible values for the OrchestratorKind const type.
func PossibleOrchestratorKindValues() []OrchestratorKind {
	return []OrchestratorKind{
		OrchestratorKindKubernetes,
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

// ResourceIdentityType - The type of identity used for orchestrator cluster. Type 'SystemAssigned' will use an implicitly
// created identity orchestrator clusters
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
