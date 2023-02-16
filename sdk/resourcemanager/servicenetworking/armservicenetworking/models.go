//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicenetworking

import "time"

// Association Subresource of Traffic Controller
type Association struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *AssociationProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// AssociationListResult - The response of a Traffic Controller Association list operation.
type AssociationListResult struct {
	// REQUIRED; The Association items on this page
	Value []*Association `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// AssociationProperties - Association Properties.
type AssociationProperties struct {
	// CONSTANT; Association Type
	// Field has constant value "Subnet", any specified value is ignored.
	AssociationType *string `json:"associationType,omitempty"`

	// Association Subnet
	Subnet *AssociationSubnet `json:"subnet,omitempty"`

	// READ-ONLY; Provisioning State
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// AssociationSubnet - Association Subnet.
type AssociationSubnet struct {
	// REQUIRED; Association ID.
	ID *string `json:"id,omitempty"`
}

// AssociationUpdate - The type used for update operations of the Association.
type AssociationUpdate struct {
	// The updatable properties of the Association.
	Properties *AssociationUpdateProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// AssociationUpdateProperties - The updatable properties of the Association.
type AssociationUpdateProperties struct {
	// Association Type
	AssociationType *string `json:"associationType,omitempty"`

	// Association Subnet
	Subnet *AssociationSubnet `json:"subnet,omitempty"`
}

// AssociationsInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the AssociationsInterfaceClient.BeginCreateOrUpdate
// method.
type AssociationsInterfaceClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AssociationsInterfaceClientBeginDeleteOptions contains the optional parameters for the AssociationsInterfaceClient.BeginDelete
// method.
type AssociationsInterfaceClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// AssociationsInterfaceClientGetOptions contains the optional parameters for the AssociationsInterfaceClient.Get method.
type AssociationsInterfaceClientGetOptions struct {
	// placeholder for future optional parameters
}

// AssociationsInterfaceClientListByTrafficControllerOptions contains the optional parameters for the AssociationsInterfaceClient.NewListByTrafficControllerPager
// method.
type AssociationsInterfaceClientListByTrafficControllerOptions struct {
	// placeholder for future optional parameters
}

// AssociationsInterfaceClientUpdateOptions contains the optional parameters for the AssociationsInterfaceClient.Update method.
type AssociationsInterfaceClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Frontend Subresource of Traffic Controller.
type Frontend struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *FrontendProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// FrontendListResult - The response of a Traffic Controller Frontend list operation.
type FrontendListResult struct {
	// REQUIRED; The Frontend items on this page
	Value []*Frontend `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// FrontendProperties - Frontend Properties.
type FrontendProperties struct {
	// Frontend IP Address Version (Optional).
	IPVersion *FrontendIPVersion `json:"ipVersion,omitempty"`

	// Frontend Mode (Optional).
	Mode *string `json:"mode,omitempty"`

	// Frontend Public IP Address (Optional).
	PublicIPAddress *FrontendPropertiesIPAddress `json:"publicIPAddress,omitempty"`

	// READ-ONLY; test doc
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// FrontendPropertiesIPAddress - Frontend IP Address.
type FrontendPropertiesIPAddress struct {
	// REQUIRED; IP Address.
	ID *string `json:"id,omitempty"`
}

// FrontendUpdate - The type used for update operations of the Frontend.
type FrontendUpdate struct {
	// The updatable properties of the Frontend.
	Properties *FrontendUpdateProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}

// FrontendUpdateProperties - The updatable properties of the Frontend.
type FrontendUpdateProperties struct {
	// Frontend IP Address Type (Optional).
	IPVersion *FrontendIPVersion `json:"ipVersion,omitempty"`

	// Frontend Mode (Optional).
	Mode *string `json:"mode,omitempty"`

	// Frontend Public IP Address (Optional).
	PublicIPAddress *FrontendPropertiesIPAddress `json:"publicIPAddress,omitempty"`
}

// FrontendsInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the FrontendsInterfaceClient.BeginCreateOrUpdate
// method.
type FrontendsInterfaceClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FrontendsInterfaceClientBeginDeleteOptions contains the optional parameters for the FrontendsInterfaceClient.BeginDelete
// method.
type FrontendsInterfaceClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// FrontendsInterfaceClientGetOptions contains the optional parameters for the FrontendsInterfaceClient.Get method.
type FrontendsInterfaceClientGetOptions struct {
	// placeholder for future optional parameters
}

// FrontendsInterfaceClientListByTrafficControllerOptions contains the optional parameters for the FrontendsInterfaceClient.NewListByTrafficControllerPager
// method.
type FrontendsInterfaceClientListByTrafficControllerOptions struct {
	// placeholder for future optional parameters
}

// FrontendsInterfaceClientUpdateOptions contains the optional parameters for the FrontendsInterfaceClient.Update method.
type FrontendsInterfaceClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Localized display information for this particular operation.
	Display *OperationDisplay `json:"display,omitempty"`

	// READ-ONLY; Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType `json:"actionType,omitempty" azure:"ro"`

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for ARM/control-plane
	// operations.
	IsDataAction *bool `json:"isDataAction,omitempty" azure:"ro"`

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin `json:"origin,omitempty" azure:"ro"`
}

// OperationDisplay - Localized display information for this particular operation.
type OperationDisplay struct {
	// READ-ONLY; The short, localized friendly description of the operation; suitable for tool tips and detailed views.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The concise, localized friendly name for the operation; suitable for dropdowns. E.g. "Create or Update Virtual
	// Machine", "Restart Virtual Machine".
	Operation *string `json:"operation,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly form of the resource provider name, e.g. "Microsoft Monitoring Insights" or "Microsoft
	// Compute".
	Provider *string `json:"provider,omitempty" azure:"ro"`

	// READ-ONLY; The localized friendly name of the resource type related to this operation. E.g. "Virtual Machines" or "Job
	// Schedule Collections".
	Resource *string `json:"resource,omitempty" azure:"ro"`
}

// OperationListResult - A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to
// get the next set of results.
type OperationListResult struct {
	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`

	// READ-ONLY; List of operations supported by the resource provider
	Value []*Operation `json:"value,omitempty" azure:"ro"`
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// ResourceID - Resource ID definition used by parent to reference child resources.
type ResourceID struct {
	// REQUIRED; Resource ID of child resource.
	ID *string `json:"id,omitempty"`
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// TrafficController - Concrete tracked resource types can be created by aliasing this type using a specific property type.
type TrafficController struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// The resource-specific properties for this resource.
	Properties *TrafficControllerProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; The name of the resource
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty" azure:"ro"`
}

// TrafficControllerInterfaceClientBeginCreateOrUpdateOptions contains the optional parameters for the TrafficControllerInterfaceClient.BeginCreateOrUpdate
// method.
type TrafficControllerInterfaceClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TrafficControllerInterfaceClientBeginDeleteOptions contains the optional parameters for the TrafficControllerInterfaceClient.BeginDelete
// method.
type TrafficControllerInterfaceClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// TrafficControllerInterfaceClientGetOptions contains the optional parameters for the TrafficControllerInterfaceClient.Get
// method.
type TrafficControllerInterfaceClientGetOptions struct {
	// placeholder for future optional parameters
}

// TrafficControllerInterfaceClientListByResourceGroupOptions contains the optional parameters for the TrafficControllerInterfaceClient.NewListByResourceGroupPager
// method.
type TrafficControllerInterfaceClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// TrafficControllerInterfaceClientListBySubscriptionOptions contains the optional parameters for the TrafficControllerInterfaceClient.NewListBySubscriptionPager
// method.
type TrafficControllerInterfaceClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// TrafficControllerInterfaceClientUpdateOptions contains the optional parameters for the TrafficControllerInterfaceClient.Update
// method.
type TrafficControllerInterfaceClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// TrafficControllerListResult - The response of a TrafficController list operation.
type TrafficControllerListResult struct {
	// REQUIRED; The TrafficController items on this page
	Value []*TrafficController `json:"value,omitempty"`

	// The link to the next page of items
	NextLink *string `json:"nextLink,omitempty"`
}

// TrafficControllerProperties - Traffic Controller Properties.
type TrafficControllerProperties struct {
	// READ-ONLY; Associations References List
	Associations []*ResourceID `json:"associations,omitempty" azure:"ro"`

	// READ-ONLY; Configuration Endpoints.
	ConfigurationEndpoints []*string `json:"configurationEndpoints,omitempty" azure:"ro"`

	// READ-ONLY; Frontends References List
	Frontends []*ResourceID `json:"frontends,omitempty" azure:"ro"`

	// READ-ONLY; The status of the last operation.
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty" azure:"ro"`
}

// TrafficControllerUpdate - The type used for update operations of the TrafficController.
type TrafficControllerUpdate struct {
	// The updatable properties of the TrafficController.
	Properties any `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`
}
