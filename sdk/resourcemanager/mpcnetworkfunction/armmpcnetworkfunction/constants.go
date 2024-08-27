//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmpcnetworkfunction

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mpcnetworkfunction/armmpcnetworkfunction"
	moduleVersion = "v0.1.0"
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

// NetworkFunctionAdministrativeState - Network Function Administrative State enumerations
type NetworkFunctionAdministrativeState string

const (
	// NetworkFunctionAdministrativeStateCommissioned - Resource has been commissioned
	NetworkFunctionAdministrativeStateCommissioned NetworkFunctionAdministrativeState = "Commissioned"
	// NetworkFunctionAdministrativeStateDecommissioned - Resource has been decommissioned
	NetworkFunctionAdministrativeStateDecommissioned NetworkFunctionAdministrativeState = "Decommissioned"
)

// PossibleNetworkFunctionAdministrativeStateValues returns the possible values for the NetworkFunctionAdministrativeState const type.
func PossibleNetworkFunctionAdministrativeStateValues() []NetworkFunctionAdministrativeState {
	return []NetworkFunctionAdministrativeState{
		NetworkFunctionAdministrativeStateCommissioned,
		NetworkFunctionAdministrativeStateDecommissioned,
	}
}

// NetworkFunctionOperationalStatus - Network Function Operational Status enumerations
type NetworkFunctionOperationalStatus string

const (
	// NetworkFunctionOperationalStatusActive - Resource is fully deployed and operational
	NetworkFunctionOperationalStatusActive NetworkFunctionOperationalStatus = "Active"
	// NetworkFunctionOperationalStatusInactive - Resource is inactive
	NetworkFunctionOperationalStatusInactive NetworkFunctionOperationalStatus = "Inactive"
	// NetworkFunctionOperationalStatusInstantiatedNotProvisioned - Resource has been deployed, awaiting provisioning
	NetworkFunctionOperationalStatusInstantiatedNotProvisioned NetworkFunctionOperationalStatus = "InstantiatedNotProvisioned"
	// NetworkFunctionOperationalStatusInstantiatedProvisioned - Resource has been deployed and provisioned, awaiting testing
	NetworkFunctionOperationalStatusInstantiatedProvisioned NetworkFunctionOperationalStatus = "InstantiatedProvisioned"
)

// PossibleNetworkFunctionOperationalStatusValues returns the possible values for the NetworkFunctionOperationalStatus const type.
func PossibleNetworkFunctionOperationalStatusValues() []NetworkFunctionOperationalStatus {
	return []NetworkFunctionOperationalStatus{
		NetworkFunctionOperationalStatusActive,
		NetworkFunctionOperationalStatusInactive,
		NetworkFunctionOperationalStatusInstantiatedNotProvisioned,
		NetworkFunctionOperationalStatusInstantiatedProvisioned,
	}
}

// NetworkFunctionType - Type of Network Function
type NetworkFunctionType string

const (
	// NetworkFunctionTypeAMF - Access and Mobility Function
	NetworkFunctionTypeAMF NetworkFunctionType = "AMF"
	// NetworkFunctionTypeEMS - Element Management System
	NetworkFunctionTypeEMS NetworkFunctionType = "EMS"
	// NetworkFunctionTypeEPDG - Evolved Packet Data Gateway
	NetworkFunctionTypeEPDG NetworkFunctionType = "ePDG"
	// NetworkFunctionTypeMME - Mobility Management Entity
	NetworkFunctionTypeMME NetworkFunctionType = "MME"
	// NetworkFunctionTypeN3IWF - Non-3GPP Interworking Function
	NetworkFunctionTypeN3IWF NetworkFunctionType = "N3IWF"
	// NetworkFunctionTypeNRF - Network Repository Function
	NetworkFunctionTypeNRF NetworkFunctionType = "NRF"
	// NetworkFunctionTypeNSSF - Network Slice Selection Function
	NetworkFunctionTypeNSSF NetworkFunctionType = "NSSF"
	// NetworkFunctionTypeOperationsPolicyManager - Operations and Policy Manager
	NetworkFunctionTypeOperationsPolicyManager NetworkFunctionType = "OperationsPolicyManager"
	// NetworkFunctionTypeRemotePaaS - Remote Platform As A Service Components
	NetworkFunctionTypeRemotePaaS NetworkFunctionType = "RemotePaaS"
	// NetworkFunctionTypeSMF - Session Management Function
	NetworkFunctionTypeSMF NetworkFunctionType = "SMF"
	// NetworkFunctionTypeSaegw - System Architecture Evolution Gateway. This combines the Serving Gateway (SGW) and the Packet
	// Data Network Gateway (PGW) functionality
	NetworkFunctionTypeSaegw NetworkFunctionType = "Saegw"
	// NetworkFunctionTypeSaegwControlPlane - System Architecture Evolution Gateway Control Plane, control and user plane separation
	// (CUPS) architecture
	NetworkFunctionTypeSaegwControlPlane NetworkFunctionType = "SaegwControlPlane"
	// NetworkFunctionTypeSaegwUserPlane - System Architecture Evolution Gateway User Plane, control and user plane separation
	// (CUPS) architecture
	NetworkFunctionTypeSaegwUserPlane NetworkFunctionType = "SaegwUserPlane"
	// NetworkFunctionTypeUPF - User Plane Function
	NetworkFunctionTypeUPF NetworkFunctionType = "UPF"
)

// PossibleNetworkFunctionTypeValues returns the possible values for the NetworkFunctionType const type.
func PossibleNetworkFunctionTypeValues() []NetworkFunctionType {
	return []NetworkFunctionType{
		NetworkFunctionTypeAMF,
		NetworkFunctionTypeEMS,
		NetworkFunctionTypeEPDG,
		NetworkFunctionTypeMME,
		NetworkFunctionTypeN3IWF,
		NetworkFunctionTypeNRF,
		NetworkFunctionTypeNSSF,
		NetworkFunctionTypeOperationsPolicyManager,
		NetworkFunctionTypeRemotePaaS,
		NetworkFunctionTypeSMF,
		NetworkFunctionTypeSaegw,
		NetworkFunctionTypeSaegwControlPlane,
		NetworkFunctionTypeSaegwUserPlane,
		NetworkFunctionTypeUPF,
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

// ProvisioningState - Provisioning state of the resource
type ProvisioningState string

const (
	// ProvisioningStateAccepted - Resource has been accepted
	ProvisioningStateAccepted ProvisioningState = "Accepted"
	// ProvisioningStateCanceled - Resource creation was canceled.
	ProvisioningStateCanceled ProvisioningState = "Canceled"
	// ProvisioningStateDeleting - Resource is getting deleted
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Resource creation failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - Resource is getting provisioned
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateSucceeded - Resource has been created.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - Resource is updating
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateAccepted,
		ProvisioningStateCanceled,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// SKUDefinitions - SKU Definitions
type SKUDefinitions string

const (
	// SKUDefinitionsAzureLab - Azure Lab SKU
	SKUDefinitionsAzureLab SKUDefinitions = "AzureLab"
	// SKUDefinitionsAzureProduction - Azure Production SKU
	SKUDefinitionsAzureProduction SKUDefinitions = "AzureProduction"
	// SKUDefinitionsNexusLab - Nexus Lab SKU
	SKUDefinitionsNexusLab SKUDefinitions = "NexusLab"
	// SKUDefinitionsNexusProduction - Nexus Production SKU
	SKUDefinitionsNexusProduction SKUDefinitions = "NexusProduction"
)

// PossibleSKUDefinitionsValues returns the possible values for the SKUDefinitions const type.
func PossibleSKUDefinitionsValues() []SKUDefinitions {
	return []SKUDefinitions{
		SKUDefinitionsAzureLab,
		SKUDefinitionsAzureProduction,
		SKUDefinitionsNexusLab,
		SKUDefinitionsNexusProduction,
	}
}