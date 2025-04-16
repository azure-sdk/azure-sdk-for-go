// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetwork

// ConnectivityCollection - The collection of Connectivity related groups and policies within the Managed Network
type ConnectivityCollection struct {
	// READ-ONLY; The collection of connectivity related Managed Network Groups within the Managed Network
	Groups []*Group

	// READ-ONLY; The collection of Managed Network Peering Policies within the Managed Network
	Peerings []*PeeringPolicy
}

// ErrorResponse - The error response that indicates why an operation has failed.
type ErrorResponse struct {
	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error message.
	Message *string
}

// Group - The Managed Network Group resource
type Group struct {
	// Responsibility role under which this Managed Network Group will be created
	Kind *Kind

	// The geo-location where the resource lives
	Location *string

	// Gets or sets the properties of a network group
	Properties *GroupProperties

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// GroupListResult - Result of the request to list Managed Network Groups. It contains a list of groups and a URL link to
// get the next set of results.
type GroupListResult struct {
	// Gets the URL to get the next set of results.
	NextLink *string

	// Gets a page of ManagedNetworkGroup
	Value []*Group
}

// GroupProperties - Properties of a Managed Network Group
type GroupProperties struct {
	// The collection of management groups covered by the Managed Network
	ManagementGroups []*ResourceID

	// The collection of subnets covered by the Managed Network
	Subnets []*ResourceID

	// The collection of subscriptions covered by the Managed Network
	Subscriptions []*ResourceID

	// The collection of virtual nets covered by the Managed Network
	VirtualNetworks []*ResourceID

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// HubAndSpokePeeringPolicyProperties - Properties of a Hub and Spoke Peering Policy
type HubAndSpokePeeringPolicyProperties struct {
	// REQUIRED; Gets or sets the connectivity type of a network structure policy
	Type *Type

	// Gets or sets the hub virtual network ID
	Hub *ResourceID

	// Gets or sets the mesh group IDs
	Mesh []*ResourceID

	// Gets or sets the spokes group IDs
	Spokes []*ResourceID

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// ListResult - Result of the request to list Managed Network. It contains a list of Managed Networks and a URL link to get
// the next set of results.
type ListResult struct {
	// Gets the URL to get the next page of results.
	NextLink *string

	// Gets a page of ManagedNetworks
	Value []*ManagedNetwork
}

// ManagedNetwork - The Managed Network resource
type ManagedNetwork struct {
	// The geo-location where the resource lives
	Location *string

	// The MNC properties
	Properties *Properties

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// MeshPeeringPolicyProperties - Properties of a Mesh Peering Policy
type MeshPeeringPolicyProperties struct {
	// REQUIRED; Gets or sets the connectivity type of a network structure policy
	Type *Type

	// Gets or sets the hub virtual network ID
	Hub *ResourceID

	// Gets or sets the mesh group IDs
	Mesh []*ResourceID

	// Gets or sets the spokes group IDs
	Spokes []*ResourceID

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// Operation - REST API operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.ManagedNetwork
	Provider *string

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string
}

// OperationListResult - Result of the request to list Managed Network operations. It contains a list of operations and a
// URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Resource Provider operations supported by the Managed Network resource provider.
	Value []*Operation
}

// PeeringPolicy - The Managed Network Peering Policy resource
type PeeringPolicy struct {
	// The geo-location where the resource lives
	Location *string

	// Gets or sets the properties of a Managed Network Policy
	Properties *PeeringPolicyProperties

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// PeeringPolicyListResult - Result of the request to list Managed Network Peering Policies. It contains a list of policies
// and a URL link to get the next set of results.
type PeeringPolicyListResult struct {
	// Gets the URL to get the next page of results.
	NextLink *string

	// Gets a page of Peering Policies
	Value []*PeeringPolicy
}

// PeeringPolicyProperties - Properties of a Managed Network Peering Policy
type PeeringPolicyProperties struct {
	// REQUIRED; Gets or sets the connectivity type of a network structure policy
	Type *Type

	// Gets or sets the hub virtual network ID
	Hub *ResourceID

	// Gets or sets the mesh group IDs
	Mesh []*ResourceID

	// Gets or sets the spokes group IDs
	Spokes []*ResourceID

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// Properties of Managed Network
type Properties struct {
	// The collection of management groups, subscriptions, virtual networks, and subnets by the Managed Network. This is a read-only
	// property that is reflective of all ScopeAssignments for this Managed
	// Network
	Scope *Scope

	// READ-ONLY; The collection of groups and policies concerned with connectivity
	Connectivity *ConnectivityCollection

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// ProxyResource - The resource model definition for a ARM proxy resource. It will have everything other than required location
// and tags
type ProxyResource struct {
	// The geo-location where the resource lives
	Location *string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// Resource - The general resource model definition
type Resource struct {
	// The geo-location where the resource lives
	Location *string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// ResourceID - Generic pointer to a resource
type ResourceID struct {
	// Resource Id
	ID *string
}

// ResourceProperties - Base for resource properties.
type ResourceProperties struct {
	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// Scope of a Managed Network
type Scope struct {
	// The collection of management groups covered by the Managed Network
	ManagementGroups []*ResourceID

	// The collection of subnets covered by the Managed Network
	Subnets []*ResourceID

	// The collection of subscriptions covered by the Managed Network
	Subscriptions []*ResourceID

	// The collection of virtual nets covered by the Managed Network
	VirtualNetworks []*ResourceID
}

// ScopeAssignment - The Managed Network resource
type ScopeAssignment struct {
	// The geo-location where the resource lives
	Location *string

	// The Scope Assignment properties
	Properties *ScopeAssignmentProperties

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// ScopeAssignmentListResult - Result of the request to list ScopeAssignment. It contains a list of groups and a URL link
// to get the next set of results.
type ScopeAssignmentListResult struct {
	// Gets the URL to get the next set of results.
	NextLink *string

	// Gets a page of ScopeAssignment
	Value []*ScopeAssignment
}

// ScopeAssignmentProperties - Properties of Managed Network
type ScopeAssignmentProperties struct {
	// The managed network ID with scope will be assigned to.
	AssignedManagedNetwork *string

	// READ-ONLY; A unique read-only string that changes whenever the resource is updated.
	Etag *string

	// READ-ONLY; Provisioning state of the ManagedNetwork resource.
	ProvisioningState *ProvisioningState
}

// TrackedResource - The resource model definition for a ARM tracked top level resource
type TrackedResource struct {
	// The geo-location where the resource lives
	Location *string

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type *string
}

// Update Tags of Managed Network
type Update struct {
	// Resource tags
	Tags map[string]*string
}
