// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armhybridconnectivity

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/hybridconnectivity/armhybridconnectivity"
	moduleVersion = "v2.0.0"
)

// ActionType - Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	// ActionTypeInternal - Actions are for internal-only APIs.
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
	}
}

// CloudNativeType - Cloud Native Type enum.
type CloudNativeType string

const (
	// CloudNativeTypeEc2 - ec2 enum.
	CloudNativeTypeEc2 CloudNativeType = "ec2"
)

// PossibleCloudNativeTypeValues returns the possible values for the CloudNativeType const type.
func PossibleCloudNativeTypeValues() []CloudNativeType {
	return []CloudNativeType{
		CloudNativeTypeEc2,
	}
}

// CreatedByType - The kind of entity that created the resource.
type CreatedByType string

const (
	// CreatedByTypeApplication - The entity was created by an application.
	CreatedByTypeApplication CreatedByType = "Application"
	// CreatedByTypeKey - The entity was created by a key.
	CreatedByTypeKey CreatedByType = "Key"
	// CreatedByTypeManagedIdentity - The entity was created by a managed identity.
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	// CreatedByTypeUser - The entity was created by a user.
	CreatedByTypeUser CreatedByType = "User"
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

// HostType - Enum of host cloud the public cloud connector is referencing.
type HostType string

const (
	// HostTypeAWS - AWS state
	HostTypeAWS HostType = "AWS"
)

// PossibleHostTypeValues returns the possible values for the HostType const type.
func PossibleHostTypeValues() []HostType {
	return []HostType{
		HostTypeAWS,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	// OriginSystem - Indicates the operation is initiated by a system.
	OriginSystem Origin = "system"
	// OriginUser - Indicates the operation is initiated by a user.
	OriginUser Origin = "user"
	// OriginUserSystem - Indicates the operation is initiated by a user or system.
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

// ProvisioningState - The resource provisioning state.
type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCanceled,
		ProvisioningStateCreating,
		ProvisioningStateFailed,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ResourceProvisioningState - The provisioning state of a resource type.
type ResourceProvisioningState string

const (
	// ResourceProvisioningStateCanceled - Resource creation was canceled.
	ResourceProvisioningStateCanceled ResourceProvisioningState = "Canceled"
	// ResourceProvisioningStateFailed - Resource creation failed.
	ResourceProvisioningStateFailed ResourceProvisioningState = "Failed"
	// ResourceProvisioningStateSucceeded - Resource has been created.
	ResourceProvisioningStateSucceeded ResourceProvisioningState = "Succeeded"
)

// PossibleResourceProvisioningStateValues returns the possible values for the ResourceProvisioningState const type.
func PossibleResourceProvisioningStateValues() []ResourceProvisioningState {
	return []ResourceProvisioningState{
		ResourceProvisioningStateCanceled,
		ResourceProvisioningStateFailed,
		ResourceProvisioningStateSucceeded,
	}
}

// ServiceName - Name of the service.
type ServiceName string

const (
	ServiceNameSSH ServiceName = "SSH"
	ServiceNameWAC ServiceName = "WAC"
)

// PossibleServiceNameValues returns the possible values for the ServiceName const type.
func PossibleServiceNameValues() []ServiceName {
	return []ServiceName{
		ServiceNameSSH,
		ServiceNameWAC,
	}
}

// SolutionConfigurationStatus - Solution Configuration Status.
type SolutionConfigurationStatus string

const (
	// SolutionConfigurationStatusCompleted - Canceled status
	SolutionConfigurationStatusCompleted SolutionConfigurationStatus = "Completed"
	// SolutionConfigurationStatusFailed - Failed status
	SolutionConfigurationStatusFailed SolutionConfigurationStatus = "Failed"
	// SolutionConfigurationStatusInProgress - InProgress status
	SolutionConfigurationStatusInProgress SolutionConfigurationStatus = "InProgress"
	// SolutionConfigurationStatusNew - New status
	SolutionConfigurationStatusNew SolutionConfigurationStatus = "New"
)

// PossibleSolutionConfigurationStatusValues returns the possible values for the SolutionConfigurationStatus const type.
func PossibleSolutionConfigurationStatusValues() []SolutionConfigurationStatus {
	return []SolutionConfigurationStatus{
		SolutionConfigurationStatusCompleted,
		SolutionConfigurationStatusFailed,
		SolutionConfigurationStatusInProgress,
		SolutionConfigurationStatusNew,
	}
}

// Type - The type of endpoint.
type Type string

const (
	TypeCustom  Type = "custom"
	TypeDefault Type = "default"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeCustom,
		TypeDefault,
	}
}
