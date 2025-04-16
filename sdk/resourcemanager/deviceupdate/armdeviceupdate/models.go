// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdeviceupdate

import "time"

// Account - Device Update account details.
type Account struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The type of identity used for the resource.
	Identity *ManagedServiceIdentity

	// Device Update account properties.
	Properties *AccountProperties

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

// AccountList - List of Accounts.
type AccountList struct {
	// The link used to get the next page of Accounts list.
	NextLink *string

	// List of Accounts.
	Value []*Account
}

// AccountProperties - Device Update account properties.
type AccountProperties struct {
	// CMK encryption at rest properties
	Encryption *Encryption

	// List of private endpoint connections associated with the account.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// Whether or not public network access is allowed for the account.
	PublicNetworkAccess *PublicNetworkAccess

	// Device Update Sku
	SKU *SKU

	// READ-ONLY; API host name.
	HostName *string

	// READ-ONLY; Device Update account primary and failover location details
	Locations []*Location

	// READ-ONLY; Provisioning state.
	ProvisioningState *ProvisioningState
}

// AccountUpdate - Request payload used to update and existing Accounts.
type AccountUpdate struct {
	// The type of identity used for the resource.
	Identity *ManagedServiceIdentity

	// The geo-location where the resource lives
	Location *string

	// List of key value pairs that describe the resource. This will overwrite the existing tags.
	Tags map[string]*string
}

// CheckNameAvailabilityRequest - The check availability request body.
type CheckNameAvailabilityRequest struct {
	// The name of the resource for which availability needs to be checked.
	Name *string

	// The resource type.
	Type *string
}

// CheckNameAvailabilityResponse - The check availability result.
type CheckNameAvailabilityResponse struct {
	// Detailed reason why the given name is available.
	Message *string

	// Indicates if the resource name is available.
	NameAvailable *bool

	// The reason why the given name is not available.
	Reason *CheckNameAvailabilityReason
}

// ConnectionDetails - Private endpoint connection proxy object properties.
type ConnectionDetails struct {
	// READ-ONLY; Group ID.
	GroupID *string

	// READ-ONLY; Connection details ID.
	ID *string

	// READ-ONLY; Link ID.
	LinkIdentifier *string

	// READ-ONLY; Member name.
	MemberName *string

	// READ-ONLY; Private IP address.
	PrivateIPAddress *string
}

// DiagnosticStorageProperties - Customer-initiated diagnostic log collection storage properties
type DiagnosticStorageProperties struct {
	// REQUIRED; Authentication Type
	AuthenticationType *AuthenticationType

	// REQUIRED; ResourceId of the diagnostic storage account
	ResourceID *string

	// ConnectionString of the diagnostic storage account
	ConnectionString *string
}

// Encryption - The CMK encryption settings on the Device Update account.
type Encryption struct {
	// The URI of the key vault
	KeyVaultKeyURI *string

	// The full resourceId of the user assigned identity to be used for key vault access. Identity has to be also assigned to
	// the Account
	UserAssignedIdentity *string
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

// GroupConnectivityInformation - Group connectivity details.
type GroupConnectivityInformation struct {
	// List of customer visible FQDNs.
	CustomerVisibleFqdns []*string

	// PrivateLinkService ARM region.
	PrivateLinkServiceArmRegion *string

	// Redirect map ID.
	RedirectMapID *string

	// READ-ONLY; Group ID.
	GroupID *string

	// READ-ONLY; Internal FQDN.
	InternalFqdn *string

	// READ-ONLY; Member name.
	MemberName *string
}

// GroupInformation - The group information for creating a private endpoint on an Account
type GroupInformation struct {
	// REQUIRED; The properties for a group information object
	Properties *GroupInformationProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// GroupInformationProperties - The properties for a group information object
type GroupInformationProperties struct {
	// The private link resource Private link DNS zone name.
	RequiredZoneNames []*string

	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The provisioning state of private link group ID.
	ProvisioningState *GroupIDProvisioningState

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string
}

// Instance - Device Update instance details.
type Instance struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; Device Update instance properties.
	Properties *InstanceProperties

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

// InstanceList - List of Instances.
type InstanceList struct {
	// The link used to get the next page of Instances list.
	NextLink *string

	// List of Instances.
	Value []*Instance
}

// InstanceProperties - Device Update instance properties.
type InstanceProperties struct {
	// Customer-initiated diagnostic log collection storage properties
	DiagnosticStorageProperties *DiagnosticStorageProperties

	// Enables or Disables the diagnostic logs collection
	EnableDiagnostics *bool

	// List of IoT Hubs associated with the account.
	IotHubs []*IotHubSettings

	// READ-ONLY; Parent Device Update Account name which Instance belongs to.
	AccountName *string

	// READ-ONLY; Provisioning state.
	ProvisioningState *ProvisioningState
}

// IotHubSettings - Device Update account integration with IoT Hub settings.
type IotHubSettings struct {
	// REQUIRED; IoTHub resource ID
	ResourceID *string
}

type Location struct {
	Name *string

	// Whether the location is primary or failover
	Role *Role
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; Type of managed service identity (where both SystemAssigned and UserAssigned types are allowed).
	Type *ManagedServiceIdentityType

	// The set of user assigned identities associated with the resource. The userAssignedIdentities dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}.
	// The dictionary values can be empty objects ({}) in
	// requests.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
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

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - The Private Endpoint Connection resource.
type PrivateEndpointConnection struct {
	// REQUIRED; Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - List of private endpoint connection associated with the specified storage account
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// Array of group IDs.
	GroupIDs []*string

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateEndpointConnectionProxy - Private endpoint connection proxy details.
type PrivateEndpointConnectionProxy struct {
	// Private endpoint connection proxy object property bag.
	Properties *PrivateEndpointConnectionProxyProperties

	// Remote private endpoint details.
	RemotePrivateEndpoint *RemotePrivateEndpoint

	// Operation status.
	Status *string

	// READ-ONLY; ETag from NRP.
	ETag *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionProxyListResult - The available private endpoint connection proxies for an Account (not to be
// used by anyone, here because of ARM requirements)
type PrivateEndpointConnectionProxyListResult struct {
	// The URI that can be used to request the next list of private endpoint connection proxies.
	NextLink *string

	// The list of available private endpoint connection proxies for an Account
	Value []*PrivateEndpointConnectionProxy
}

// PrivateEndpointConnectionProxyProperties - Private endpoint connection proxy object property bag.
type PrivateEndpointConnectionProxyProperties struct {
	// READ-ONLY; The provisioning state of the private endpoint connection proxy resource.
	ProvisioningState *PrivateEndpointConnectionProxyProvisioningState
}

// PrivateEndpointConnectionProxyPropertiesAutoGenerated - Private endpoint connection proxy object properties.
type PrivateEndpointConnectionProxyPropertiesAutoGenerated struct {
	// Remote private endpoint details.
	RemotePrivateEndpoint *RemotePrivateEndpoint

	// Operation status.
	Status *string

	// READ-ONLY; ETag from NRP.
	ETag *string
}

// PrivateEndpointUpdate - Private endpoint update details.
type PrivateEndpointUpdate struct {
	// Remote endpoint resource ID.
	ID *string

	// Original resource ID needed by Microsoft.Network.
	ImmutableResourceID *string

	// Original subscription ID needed by Microsoft.Network.
	ImmutableSubscriptionID *string

	// ARM location of the remote private endpoint.
	Location *string

	// Virtual network traffic tag.
	VnetTrafficTag *string
}

// PrivateLinkResourceListResult - The available private link resources for an Account
type PrivateLinkResourceListResult struct {
	// The URI that can be used to request the next list of private link resources.
	NextLink *string

	// The list of available private link resources for an Account
	Value []*GroupInformation
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

// PrivateLinkServiceConnection - Private link service connection details.
type PrivateLinkServiceConnection struct {
	// List of group IDs.
	GroupIDs []*string

	// Private link service connection name.
	Name *string

	// Request message.
	RequestMessage *string
}

// PrivateLinkServiceConnectionState - A collection of information about the state of the connection between service consumer
// and provider.
type PrivateLinkServiceConnectionState struct {
	// A message indicating if changes on the service provider require any updates on the consumer.
	ActionsRequired *string

	// The reason for approval/rejection of the connection.
	Description *string

	// Indicates whether the connection has been Approved/Rejected/Removed by the owner of the service.
	Status *PrivateEndpointServiceConnectionStatus
}

// PrivateLinkServiceProxy - Private link service proxy details.
type PrivateLinkServiceProxy struct {
	// Group connectivity information.
	GroupConnectivityInformation []*GroupConnectivityInformation

	// NRP resource ID.
	ID *string

	// Remote private endpoint connection details.
	RemotePrivateEndpointConnection *PrivateLinkServiceProxyRemotePrivateEndpointConnection

	// Remote private link service connection state
	RemotePrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState
}

// PrivateLinkServiceProxyRemotePrivateEndpointConnection - Remote private endpoint connection details.
type PrivateLinkServiceProxyRemotePrivateEndpointConnection struct {
	// READ-ONLY; Remote private endpoint connection ID.
	ID *string
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RemotePrivateEndpoint - Remote private endpoint details.
type RemotePrivateEndpoint struct {
	// List of connection details.
	ConnectionDetails []*ConnectionDetails

	// Remote endpoint resource ID.
	ID *string

	// Original resource ID needed by Microsoft.Network.
	ImmutableResourceID *string

	// Original subscription ID needed by Microsoft.Network.
	ImmutableSubscriptionID *string

	// ARM location of the remote private endpoint.
	Location *string

	// List of private link service connections that need manual approval.
	ManualPrivateLinkServiceConnections []*PrivateLinkServiceConnection

	// List of automatically approved private link service connections.
	PrivateLinkServiceConnections []*PrivateLinkServiceConnection

	// List of private link service proxies.
	PrivateLinkServiceProxies []*PrivateLinkServiceProxy

	// Virtual network traffic tag.
	VnetTrafficTag *string
}

// RemotePrivateEndpointConnection - Remote private endpoint connection details.
type RemotePrivateEndpointConnection struct {
	// READ-ONLY; Remote private endpoint connection ID.
	ID *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
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

// TagUpdate - Request payload used to update an existing resource's tags.
type TagUpdate struct {
	// List of key value pairs that describe the resource. This will overwrite the existing tags.
	Tags map[string]*string
}

// TrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has 'tags'
// and a 'location'
type TrackedResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

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

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
