// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armterraform

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/terraform/armterraform"
	moduleVersion = "v0.1.1"
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

// TargetProvider - The target Azure Terraform Provider
type TargetProvider string

const (
	// TargetProviderAzapi - https://registry.terraform.io/providers/Azure/azapi/latest
	TargetProviderAzapi TargetProvider = "azapi"
	// TargetProviderAzurerm - https://registry.terraform.io/providers/hashicorp/azurerm/latest
	TargetProviderAzurerm TargetProvider = "azurerm"
)

// PossibleTargetProviderValues returns the possible values for the TargetProvider const type.
func PossibleTargetProviderValues() []TargetProvider {
	return []TargetProvider{
		TargetProviderAzapi,
		TargetProviderAzurerm,
	}
}

// Type - The parameter type
type Type string

const (
	TypeExportQuery         Type = "ExportQuery"
	TypeExportResource      Type = "ExportResource"
	TypeExportResourceGroup Type = "ExportResourceGroup"
)

// PossibleTypeValues returns the possible values for the Type const type.
func PossibleTypeValues() []Type {
	return []Type{
		TypeExportQuery,
		TypeExportResource,
		TypeExportResourceGroup,
	}
}
