//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdelegatednetwork

// ControllerDetails - controller details
type ControllerDetails struct {
	// controller arm resource id
	ID *string
}

// ControllerResource - Represents an instance of a resource.
type ControllerResource struct {
	// Location of the resource.
	Location *string

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of resource.
	Type *string
}

// ControllerResourceUpdateParameters - Parameters for updating a resource.
type ControllerResourceUpdateParameters struct {
	// The resource tags.
	Tags map[string]*string
}

// DelegatedController - Represents an instance of a DNC controller.
type DelegatedController struct {
	// Location of the resource.
	Location *string

	// Properties of the provision operation request.
	Properties *DelegatedControllerProperties

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of resource.
	Type *string
}

// DelegatedControllerProperties - Properties of Delegated controller resource.
type DelegatedControllerProperties struct {
	// The purpose of the dnc controller resource.
	Purpose *ControllerPurpose

	// READ-ONLY; dnc application id should be used by customer to authenticate with dnc gateway.
	DncAppID *string

	// READ-ONLY; dnc endpoint url that customers can use to connect to
	DncEndpoint *string

	// READ-ONLY; tenant id of dnc application id
	DncTenantID *string

	// READ-ONLY; The current state of dnc controller resource.
	ProvisioningState *ControllerState

	// READ-ONLY; Resource guid.
	ResourceGUID *string
}

// DelegatedControllers - An array of Delegated controller resources.
type DelegatedControllers struct {
	// REQUIRED; An array of Delegated controller resources.
	Value []*DelegatedController

	// READ-ONLY; The URL to get the next set of controllers.
	NextLink *string
}

// DelegatedSubnet - Represents an instance of a orchestrator.
type DelegatedSubnet struct {
	// Location of the resource.
	Location *string

	// Properties of the provision operation request.
	Properties *DelegatedSubnetProperties

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of resource.
	Type *string
}

// DelegatedSubnetProperties - Properties of delegated subnet
type DelegatedSubnetProperties struct {
	// Defines prefix size of CIDR blocks allocated to nodes in VnetBlock Mode. Delegated subnet's prefix size should be smaller
	// than this by a minimum of 3.
	AllocationBlockPrefixSize *int32

	// Properties of the controller.
	ControllerDetails *ControllerDetails

	// subnet details
	SubnetDetails *SubnetDetails

	// READ-ONLY; The current state of dnc delegated subnet resource.
	ProvisioningState *DelegatedSubnetState

	// READ-ONLY; Resource guid.
	ResourceGUID *string
}

// DelegatedSubnetResource - Represents an instance of a resource.
type DelegatedSubnetResource struct {
	// Location of the resource.
	Location *string

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of resource.
	Type *string
}

// DelegatedSubnets - An array of DelegatedSubnet resources.
type DelegatedSubnets struct {
	// REQUIRED; An array of DelegatedSubnet resources.
	Value []*DelegatedSubnet

	// READ-ONLY; The URL to get the next set of DelegatedSubnet resources.
	NextLink *string
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetail - The error detail.
type ErrorDetail struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorDetail

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.).
type ErrorResponse struct {
	// The error object.
	Error *ErrorDetail
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation
}

// Orchestrator - Represents an instance of a orchestrator.
type Orchestrator struct {
	// REQUIRED; The kind of workbook. Choices are user and shared.
	Kind *OrchestratorKind

	// The identity of the orchestrator
	Identity *OrchestratorIdentity

	// Location of the resource.
	Location *string

	// Properties of the provision operation request.
	Properties *OrchestratorResourceProperties

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of resource.
	Type *string
}

type OrchestratorIdentity struct {
	// The type of identity used for orchestrator cluster. Type 'SystemAssigned' will use an implicitly created identity orchestrator
	// clusters
	Type *ResourceIdentityType

	// READ-ONLY; The principal id of the system assigned identity which is used by orchestrator.
	PrincipalID *string

	// READ-ONLY; The tenant id of the system assigned identity which is used by orchestrator.
	TenantID *string
}

// OrchestratorResource - Represents an instance of a resource.
type OrchestratorResource struct {
	// REQUIRED; The kind of workbook. Choices are user and shared.
	Kind *OrchestratorKind

	// The identity of the orchestrator
	Identity *OrchestratorIdentity

	// Location of the resource.
	Location *string

	// The resource tags.
	Tags map[string]*string

	// READ-ONLY; An identifier that represents the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of resource.
	Type *string
}

// OrchestratorResourceProperties - Properties of orchestrator
type OrchestratorResourceProperties struct {
	// REQUIRED; Properties of the controller.
	ControllerDetails *ControllerDetails

	// K8s APIServer url. Either one of apiServerEndpoint or privateLinkResourceId can be specified
	APIServerEndpoint *string

	// RootCA certificate of kubernetes cluster base64 encoded
	ClusterRootCA *string

	// AAD ID used with apiserver
	OrchestratorAppID *string

	// TenantID of server App ID
	OrchestratorTenantID *string

	// private link arm resource id. Either one of apiServerEndpoint or privateLinkResourceId can be specified
	PrivateLinkResourceID *string

	// READ-ONLY; The current state of orchestratorInstance resource.
	ProvisioningState *OrchestratorInstanceState

	// READ-ONLY; Resource guid.
	ResourceGUID *string
}

// OrchestratorResourceUpdateParameters - Parameters for updating a resource.
type OrchestratorResourceUpdateParameters struct {
	// The resource tags.
	Tags map[string]*string
}

// Orchestrators - An array of OrchestratorInstance resources.
type Orchestrators struct {
	// REQUIRED; An array of OrchestratorInstance resources.
	Value []*Orchestrator

	// READ-ONLY; The URL to get the next set of orchestrators.
	NextLink *string
}

// ResourceUpdateParameters - Parameters for updating a resource.
type ResourceUpdateParameters struct {
	// The resource tags.
	Tags map[string]*string
}

// SubnetDetails - Properties of orchestrator
type SubnetDetails struct {
	// subnet arm resource id
	ID *string
}
