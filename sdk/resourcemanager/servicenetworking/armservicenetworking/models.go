//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicenetworking

import "time"

// Association Subresource of Traffic Controller
type Association struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *AssociationProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AssociationListResult - The response of a Association list operation.
type AssociationListResult struct {
	// REQUIRED; The Association items on this page
	Value []*Association

	// The link to the next page of items
	NextLink *string
}

// AssociationProperties - Association Properties.
type AssociationProperties struct {
	// REQUIRED; Association Type
	AssociationType *AssociationType

	// Association Subnet
	Subnet *AssociationSubnet

	// READ-ONLY; Provisioning State of Traffic Controller Association Resource
	ProvisioningState *ProvisioningState
}

// AssociationSubnet - Association Subnet.
type AssociationSubnet struct {
	// REQUIRED; Association ID.
	ID *string
}

// AssociationSubnetUpdate - Association Subnet.
type AssociationSubnetUpdate struct {
	// Association ID.
	ID *string
}

// AssociationUpdate - The type used for update operations of the Association.
type AssociationUpdate struct {
	// The updatable properties of the Association.
	Properties *AssociationUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// AssociationUpdateProperties - The updatable properties of the Association.
type AssociationUpdateProperties struct {
	// Association Type
	AssociationType *AssociationType

	// Association Subnet
	Subnet *AssociationSubnetUpdate
}

// Frontend Subresource of Traffic Controller.
type Frontend struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *FrontendProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// FrontendListResult - The response of a Frontend list operation.
type FrontendListResult struct {
	// REQUIRED; The Frontend items on this page
	Value []*Frontend

	// The link to the next page of items
	NextLink *string
}

// FrontendProperties - Frontend Properties.
type FrontendProperties struct {
	// READ-ONLY; The Fully Qualified Domain Name of the DNS record associated to a Traffic Controller frontend.
	Fqdn *string

	// READ-ONLY; Provisioning State of Traffic Controller Frontend Resource
	ProvisioningState *ProvisioningState
}

// FrontendUpdate - The type used for update operations of the Frontend.
type FrontendUpdate struct {
	// Resource tags.
	Tags map[string]*string
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

// ResourceID - Resource ID definition used by parent to reference child resources.
type ResourceID struct {
	// REQUIRED; Resource ID of child resource.
	ID *string
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time

	// The identity that created the resource.
	CreatedBy *string

	// The type of identity that created the resource.
	CreatedByType *CreatedByType

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time

	// The identity that last modified the resource.
	LastModifiedBy *string

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType
}

// TrafficController - Concrete tracked resource types can be created by aliasing this type using a specific property type.
type TrafficController struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The resource-specific properties for this resource.
	Properties *TrafficControllerProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// TrafficControllerListResult - The response of a TrafficController list operation.
type TrafficControllerListResult struct {
	// REQUIRED; The TrafficController items on this page
	Value []*TrafficController

	// The link to the next page of items
	NextLink *string
}

// TrafficControllerProperties - Traffic Controller Properties.
type TrafficControllerProperties struct {
	// READ-ONLY; Associations References List
	Associations []*ResourceID

	// READ-ONLY; Configuration Endpoints.
	ConfigurationEndpoints []*string

	// READ-ONLY; Frontends References List
	Frontends []*ResourceID

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState
}

// TrafficControllerUpdate - The type used for update operations of the TrafficController.
type TrafficControllerUpdate struct {
	// Resource tags.
	Tags map[string]*string
}
