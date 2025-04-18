// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbiprivatelinks

import "time"

// AsyncOperationDetail
type AsyncOperationDetail struct {
	// The operation end time.
	EndTime *string

	// The error.
	Error *ErrorDetail

	// The operation id.
	ID *string

	// The operation name.
	Name *string

	// The operation start time.
	StartTime *string

	// The operation status.
	Status *string
}

// ConnectionState information.
type ConnectionState struct {
	// Actions required (if any).
	ActionsRequired *string

	// Description of the connection state.
	Description *string

	// Status of the connection.
	Status *PersistedConnectionStatus
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

// GroupConnectivityInformation
type GroupConnectivityInformation struct {
	// Specifies the customer visible FQDNs of the group connectivity information.
	CustomerVisibleFqdns []*string

	// Specifies the group id of the group connectivity information.
	GroupID *string

	// Specifies the internal FQDN of the group connectivity information.
	InternalFqdn *string

	// Specifies the member name of the group connectivity information.
	MemberName *string

	// Specifies the ARM region of the group connectivity information.
	PrivateLinkServiceArmRegion *string
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

// PrivateEndpoint
type PrivateEndpoint struct {
	// Specifies the id of private endpoint.
	ID *string
}

// PrivateEndpointConnection
type PrivateEndpointConnection struct {
	// Specifies the properties of the private endpoint connection.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Specifies the id of the resource.
	ID *string

	// READ-ONLY; Specifies the name of the resource.
	Name *string

	// READ-ONLY; The system meta data relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Specifies the type of the resource.
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connections.
type PrivateEndpointConnectionListResult struct {
	// Specifies the name of the private endpoint connection.
	Value []*PrivateEndpointConnection

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string
}

// PrivateEndpointConnectionProperties
type PrivateEndpointConnectionProperties struct {
	// Specifies the private endpoint.
	PrivateEndpoint *PrivateEndpoint

	// Specifies the connection state.
	PrivateLinkServiceConnectionState *ConnectionState

	// Provisioning state of the Private Endpoint Connection.
	ProvisioningState *ResourceProvisioningState
}

// PrivateLinkConnectionDetail
type PrivateLinkConnectionDetail struct {
	// Specifies the group id of the connection detail.
	GroupID *string

	// Specifies the type of the connection detail.
	ID *string

	// Specifies the link id of the connection detail.
	LinkIdentifier *string

	// Specifies the member name of the connection detail.
	MemberName *string

	// Specifies the private ip address of the connection detail.
	PrivateIPAddress *string
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Fully qualified identifier of the resource.
	ID *string

	// Name of the resource.
	Name *string

	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// Type of the resource.
	Type *string
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// PrivateLinkResourcesListResult - Specifies list of the private link resource.
type PrivateLinkResourcesListResult struct {
	// A collection of private endpoint connection resources.
	Value []*PrivateLinkResource

	// READ-ONLY; URL to get the next set of operation list results (if there are any).
	NextLink *string
}

// PrivateLinkServiceConnection
type PrivateLinkServiceConnection struct {
	// Specifies the group ids of the private link service connection.
	GroupIDs []*string

	// Specifies the name of the private link service connection.
	Name *string

	// Specifies the request message of the private link service connection.
	RequestMessage *string
}

// PrivateLinkServiceProxy
type PrivateLinkServiceProxy struct {
	// Specifies the group connectivity information of the private link service proxy.
	GroupConnectivityInformation []*GroupConnectivityInformation

	// Specifies the id of the private link service proxy.
	ID *string

	// Specifies the private endpoint connection of the private link service proxy.
	RemotePrivateEndpointConnection *RemotePrivateEndpointConnection

	// Specifies the connection state of the private link service proxy.
	RemotePrivateLinkServiceConnectionState *ConnectionState
}

// RemotePrivateEndpointConnection
type RemotePrivateEndpointConnection struct {
	// Specifies the id of private endpoint connection.
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

// TenantProperties
type TenantProperties struct {
	// Specifies the private endpoint connections of the resource.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// Specifies the tenant id of the resource.
	TenantID *string
}

// TenantResource
type TenantResource struct {
	// Specifies the location of the resource.
	Location *string

	// Specifies the properties of the resource.
	Properties *TenantProperties

	// Specifies the tags of the resource.
	Tags map[string]*string

	// READ-ONLY; Specifies the resource identifier of the resource.
	ID *string

	// READ-ONLY; Specifies the name of the resource.
	Name *string

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData

	// READ-ONLY; Specifies the type of the resource.
	Type *string
}
