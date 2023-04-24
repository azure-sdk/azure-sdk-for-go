//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsearch

// AdminKeyResult - Response containing the primary and secondary admin API keys for a given Azure Cognitive Search service.
type AdminKeyResult struct {
	// READ-ONLY; The primary admin API key of the search service.
	PrimaryKey *string

	// READ-ONLY; The secondary admin API key of the search service.
	SecondaryKey *string
}

// AdminKeysClientGetOptions contains the optional parameters for the AdminKeysClient.Get method.
type AdminKeysClientGetOptions struct {
	// placeholder for future optional parameters
}

// AdminKeysClientRegenerateOptions contains the optional parameters for the AdminKeysClient.Regenerate method.
type AdminKeysClientRegenerateOptions struct {
	// placeholder for future optional parameters
}

// AsyncOperationResult - The details of a long running asynchronous shared private link resource operation
type AsyncOperationResult struct {
	// The current status of the long running asynchronous shared private link resource operation.
	Status *SharedPrivateLinkResourceAsyncOperationResult
}

// CheckNameAvailabilityInput - Input of check name availability API.
type CheckNameAvailabilityInput struct {
	// REQUIRED; The search service name to validate. Search service names must only contain lowercase letters, digits or dashes,
	// cannot use dash as the first two or last one characters, cannot contain consecutive
	// dashes, and must be between 2 and 60 characters in length.
	Name *string

	// CONSTANT; The type of the resource whose name is to be validated. This value must always be 'searchServices'.
	// Field has constant value "searchServices", any specified value is ignored.
	Type *string
}

// CheckNameAvailabilityOutput - Output of check name availability API.
type CheckNameAvailabilityOutput struct {
	// READ-ONLY; A value indicating whether the name is available.
	IsNameAvailable *bool

	// READ-ONLY; A message that explains why the name is invalid and provides resource naming requirements. Available only if
	// 'Invalid' is returned in the 'reason' property.
	Message *string

	// READ-ONLY; The reason why the name is not available. 'Invalid' indicates the name provided does not match the naming requirements
	// (incorrect length, unsupported characters, etc.). 'AlreadyExists' indicates that
	// the name is already in use and is therefore unavailable.
	Reason *UnavailableNameReason
}

// DataPlaneAADOrAPIKeyAuthOption - Indicates that either the API key or an access token from Azure Active Directory can be
// used for authentication.
type DataPlaneAADOrAPIKeyAuthOption struct {
	// Describes what response the data plane API of a Search service would send for requests that failed authentication.
	AADAuthFailureMode *AADAuthFailureMode
}

// DataPlaneAuthOptions - Defines the options for how the data plane API of a Search service authenticates requests. This
// cannot be set if 'disableLocalAuth' is set to true.
type DataPlaneAuthOptions struct {
	// Indicates that either the API key or an access token from Azure Active Directory can be used for authentication.
	AADOrAPIKey *DataPlaneAADOrAPIKeyAuthOption

	// Indicates that only the API key needs to be used for authentication.
	APIKeyOnly any
}

// EncryptionWithCmk - Describes a policy that determines how resources within the search service are to be encrypted with
// Customer Managed Keys.
type EncryptionWithCmk struct {
	// Describes how a search service should enforce having one or more non customer encrypted resources.
	Enforcement *SearchEncryptionWithCmk

	// READ-ONLY; Describes whether the search service is compliant or not with respect to having non customer encrypted resources.
	// If a service has more than one non customer encrypted resource and 'Enforcement' is
	// 'enabled' then the service will be marked as 'nonCompliant'.
	EncryptionComplianceStatus *SearchEncryptionComplianceStatus
}

// IPRule - The IP restriction rule of the Azure Cognitive Search service.
type IPRule struct {
	// Value corresponding to a single IPv4 address (eg., 123.1.2.3) or an IP range in CIDR format (eg., 123.1.2.3/24) to be allowed.
	Value *string
}

// Identity for the resource.
type Identity struct {
	// REQUIRED; The identity type.
	Type *IdentityType

	// READ-ONLY; The principal ID of the system-assigned identity of the search service.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system-assigned identity of the search service.
	TenantID *string
}

// ListQueryKeysResult - Response containing the query API keys for a given Azure Cognitive Search service.
type ListQueryKeysResult struct {
	// READ-ONLY; Request URL that can be used to query next page of query keys. Returned when the total number of requested query
	// keys exceed maximum page size.
	NextLink *string

	// READ-ONLY; The query keys for the Azure Cognitive Search service.
	Value []*QueryKey
}

// NetworkRuleSet - Network specific rules that determine how the Azure Cognitive Search service may be reached.
type NetworkRuleSet struct {
	// A list of IP restriction rules that defines the inbound network(s) with allowing access to the search service endpoint.
	// At the meantime, all other public IP networks are blocked by the firewall. These
	// restriction rules are applied only when the 'publicNetworkAccess' of the search service is 'enabled'; otherwise, traffic
	// over public interface is not allowed even with any public IP rules, and private
	// endpoint connections would be the exclusive access method.
	IPRules []*IPRule
}

// Operation - Describes a REST API operation.
type Operation struct {
	// READ-ONLY; The object that describes the operation.
	Display *OperationDisplay

	// READ-ONLY; The name of the operation. This name is of the form {provider}/{resource}/{operation}.
	Name *string
}

// OperationDisplay - The object that describes the operation.
type OperationDisplay struct {
	// READ-ONLY; The friendly name of the operation.
	Description *string

	// READ-ONLY; The operation type: read, write, delete, listKeys/action, etc.
	Operation *string

	// READ-ONLY; The friendly name of the resource provider.
	Provider *string

	// READ-ONLY; The resource type on which the operation is performed.
	Resource *string
}

// OperationListResult - The result of the request to list REST API operations. It contains a list of operations and a URL
// to get the next set of results.
type OperationListResult struct {
	// READ-ONLY; The URL to get the next set of operation list results, if any.
	NextLink *string

	// READ-ONLY; The list of operations supported by the resource provider.
	Value []*Operation
}

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnection - Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
type PrivateEndpointConnection struct {
	// Describes the properties of an existing Private Endpoint connection to the Azure Cognitive Search service.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - Response containing a list of Private Endpoint connections.
type PrivateEndpointConnectionListResult struct {
	// READ-ONLY; Request URL that can be used to query next page of private endpoint connections. Returned when the total number
	// of requested private endpoint connections exceed maximum page size.
	NextLink *string

	// READ-ONLY; The list of Private Endpoint connections.
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Describes the properties of an existing Private Endpoint connection to the Azure
// Cognitive Search service.
type PrivateEndpointConnectionProperties struct {
	// The group id from the provider of resource the private link service connection is for.
	GroupID *string

	// The private endpoint resource from Microsoft.Network provider.
	PrivateEndpoint *PrivateEndpointConnectionPropertiesPrivateEndpoint

	// Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
	PrivateLinkServiceConnectionState *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState

	// The provisioning state of the private link service connection. Can be Updating, Deleting, Failed, Succeeded, or Incomplete
	ProvisioningState *PrivateLinkServiceConnectionProvisioningState
}

// PrivateEndpointConnectionPropertiesPrivateEndpoint - The private endpoint resource from Microsoft.Network provider.
type PrivateEndpointConnectionPropertiesPrivateEndpoint struct {
	// The resource id of the private endpoint resource from Microsoft.Network provider.
	ID *string
}

// PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState - Describes the current state of an existing Private
// Link Service connection to the Azure Private Endpoint.
type PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState struct {
	// A description of any extra actions that may be required.
	ActionsRequired *string

	// The description for the private link service connection state.
	Description *string

	// Status of the the private link service connection. Can be Pending, Approved, Rejected, or Disconnected.
	Status *PrivateLinkServiceConnectionStatus
}

// PrivateEndpointConnectionsClientDeleteOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Delete
// method.
type PrivateEndpointConnectionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Get
// method.
type PrivateEndpointConnectionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientListByServiceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByServicePager
// method.
type PrivateEndpointConnectionsClientListByServiceOptions struct {
	// placeholder for future optional parameters
}

// PrivateEndpointConnectionsClientUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.Update
// method.
type PrivateEndpointConnectionsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - Describes a supported private link resource for the Azure Cognitive Search service.
type PrivateLinkResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Describes the properties of a supported private link resource for the Azure Cognitive Search service.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceProperties - Describes the properties of a supported private link resource for the Azure Cognitive Search
// service. For a given API version, this represents the 'supported' groupIds when creating a shared private
// link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; The group ID of the private link resource.
	GroupID *string

	// READ-ONLY; The list of required members of the private link resource.
	RequiredMembers []*string

	// READ-ONLY; The list of required DNS zone names of the private link resource.
	RequiredZoneNames []*string

	// READ-ONLY; The list of resources that are onboarded to private link service, that are supported by Azure Cognitive Search.
	ShareablePrivateLinkResourceTypes []*ShareablePrivateLinkResourceType
}

// PrivateLinkResourcesClientListSupportedOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListSupportedPager
// method.
type PrivateLinkResourcesClientListSupportedOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesResult - Response containing a list of supported Private Link Resources.
type PrivateLinkResourcesResult struct {
	// READ-ONLY; The list of supported Private Link Resources.
	Value []*PrivateLinkResource
}

// QueryKey - Describes an API key for a given Azure Cognitive Search service that has permissions for query operations only.
type QueryKey struct {
	// READ-ONLY; The value of the query API key.
	Key *string

	// READ-ONLY; The name of the query API key; may be empty.
	Name *string
}

// QueryKeysClientCreateOptions contains the optional parameters for the QueryKeysClient.Create method.
type QueryKeysClientCreateOptions struct {
	// placeholder for future optional parameters
}

// QueryKeysClientDeleteOptions contains the optional parameters for the QueryKeysClient.Delete method.
type QueryKeysClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// QueryKeysClientListBySearchServiceOptions contains the optional parameters for the QueryKeysClient.NewListBySearchServicePager
// method.
type QueryKeysClientListBySearchServiceOptions struct {
	// placeholder for future optional parameters
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SKU - Defines the SKU of an Azure Cognitive Search Service, which determines price tier and capacity limits.
type SKU struct {
	// The SKU of the search service. Valid values include: 'free': Shared service. 'basic': Dedicated service with up to 3 replicas.
	// 'standard': Dedicated service with up to 12 partitions and 12 replicas.
	// 'standard2': Similar to standard, but with more capacity per search unit. 'standard3': The largest Standard offering with
	// up to 12 partitions and 12 replicas (or up to 3 partitions with more indexes
	// if you also set the hostingMode property to 'highDensity'). 'storageoptimizedl1': Supports 1TB per partition, up to 12
	// partitions. 'storageoptimizedl2': Supports 2TB per partition, up to 12
	// partitions.'
	Name *SKUName
}

// SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get method.
type SearchManagementRequestOptions struct {
	// A client-generated GUID value that identifies this request. If specified, this will be included in response information
	// as a way to track the request.
	ClientRequestID *string
}

// Service - Describes an Azure Cognitive Search service and its current state.
type Service struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The identity of the resource.
	Identity *Identity

	// Properties of the search service.
	Properties *ServiceProperties

	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating
	// a new Search Service.
	SKU *SKU

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServiceListResult - Response containing a list of Azure Cognitive Search services.
type ServiceListResult struct {
	// READ-ONLY; Request URL that can be used to query next page of search services. Returned when the total number of requested
	// search services exceed maximum page size.
	NextLink *string

	// READ-ONLY; The list of Search services.
	Value []*Service
}

// ServiceProperties - Properties of the search service.
type ServiceProperties struct {
	// Defines the options for how the data plane API of a search service authenticates requests. This cannot be set if 'disableLocalAuth'
	// is set to true.
	AuthOptions *DataPlaneAuthOptions

	// When set to true, calls to the search service will not be permitted to utilize API keys for authentication. This cannot
	// be set to true if 'dataPlaneAuthOptions' are defined.
	DisableLocalAuth *bool

	// Specifies any policy regarding encryption of resources (such as indexes) using customer manager keys within a search service.
	EncryptionWithCmk *EncryptionWithCmk

	// Applicable only for the standard3 SKU. You can set this property to enable up to 3 high density partitions that allow up
	// to 1000 indexes, which is much higher than the maximum indexes allowed for any
	// other SKU. For the standard3 SKU, the value is either 'default' or 'highDensity'. For all other SKUs, this value must be
	// 'default'.
	HostingMode *HostingMode

	// Network specific rules that determine how the Azure Cognitive Search service may be reached.
	NetworkRuleSet *NetworkRuleSet

	// The number of partitions in the search service; if specified, it can be 1, 2, 3, 4, 6, or 12. Values greater than 1 are
	// only valid for standard SKUs. For 'standard3' services with hostingMode set to
	// 'highDensity', the allowed values are between 1 and 3.
	PartitionCount *int32

	// This value can be set to 'enabled' to avoid breaking changes on existing customer resources and templates. If set to 'disabled',
	// traffic over public interface is not allowed, and private endpoint
	// connections would be the exclusive access method.
	PublicNetworkAccess *PublicNetworkAccess

	// The number of replicas in the search service. If specified, it must be a value between 1 and 12 inclusive for standard
	// SKUs or between 1 and 3 inclusive for basic SKU.
	ReplicaCount *int32

	// READ-ONLY; The list of private endpoint connections to the Azure Cognitive Search service.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The state of the last provisioning operation performed on the search service. Provisioning is an intermediate
	// state that occurs while service capacity is being established. After capacity is set up,
	// provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended
	// polling interval is from 30 seconds to one minute) by using the Get
	// Search Service operation to see when an operation is completed. If you are using the free service, this value tends to
	// come back as 'succeeded' directly in the call to Create search service. This is
	// because the free service uses capacity that is already set up.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The list of shared private link resources managed by the Azure Cognitive Search service.
	SharedPrivateLinkResources []*SharedPrivateLinkResource

	// READ-ONLY; The status of the search service. Possible values include: 'running': The search service is running and no provisioning
	// operations are underway. 'provisioning': The search service is being provisioned
	// or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This
	// can occur when the underlying search units are not healthy. The search service
	// is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service
	// is disabled. In this state, the service will reject all API requests.
	// 'error': The search service is in an error state. If your service is in the degraded, disabled, or error states, it means
	// the Azure Cognitive Search team is actively investigating the underlying
	// issue. Dedicated services in these states are still chargeable based on the number of search units provisioned.
	Status *SearchServiceStatus

	// READ-ONLY; The details of the search service status.
	StatusDetails *string
}

// ServiceUpdate - The parameters used to update an Azure Cognitive Search service.
type ServiceUpdate struct {
	// The identity of the resource.
	Identity *Identity

	// The geographic location of the resource. This must be one of the supported and registered Azure Geo Regions (for example,
	// West US, East US, Southeast Asia, and so forth). This property is required
	// when creating a new resource.
	Location *string

	// Properties of the search service.
	Properties *ServiceProperties

	// The SKU of the Search Service, which determines price tier and capacity limits. This property is required when creating
	// a new Search Service.
	SKU *SKU

	// Tags to help categorize the resource in the Azure portal.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServicesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServicesClient.BeginCreateOrUpdate method.
type ServicesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// ServicesClientCheckNameAvailabilityOptions contains the optional parameters for the ServicesClient.CheckNameAvailability
// method.
type ServicesClientCheckNameAvailabilityOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientDeleteOptions contains the optional parameters for the ServicesClient.Delete method.
type ServicesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientGetOptions contains the optional parameters for the ServicesClient.Get method.
type ServicesClientGetOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListByResourceGroupOptions contains the optional parameters for the ServicesClient.NewListByResourceGroupPager
// method.
type ServicesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientListBySubscriptionOptions contains the optional parameters for the ServicesClient.NewListBySubscriptionPager
// method.
type ServicesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// ServicesClientUpdateOptions contains the optional parameters for the ServicesClient.Update method.
type ServicesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// ShareablePrivateLinkResourceProperties - Describes the properties of a resource type that has been onboarded to private
// link service, supported by Azure Cognitive Search.
type ShareablePrivateLinkResourceProperties struct {
	// READ-ONLY; The description of the resource type that has been onboarded to private link service, supported by Azure Cognitive
	// Search.
	Description *string

	// READ-ONLY; The resource provider group id for the resource that has been onboarded to private link service, supported by
	// Azure Cognitive Search.
	GroupID *string

	// READ-ONLY; The resource provider type for the resource that has been onboarded to private link service, supported by Azure
	// Cognitive Search.
	Type *string
}

// ShareablePrivateLinkResourceType - Describes an resource type that has been onboarded to private link service, supported
// by Azure Cognitive Search.
type ShareablePrivateLinkResourceType struct {
	// READ-ONLY; The name of the resource type that has been onboarded to private link service, supported by Azure Cognitive
	// Search.
	Name *string

	// READ-ONLY; Describes the properties of a resource type that has been onboarded to private link service, supported by Azure
	// Cognitive Search.
	Properties *ShareablePrivateLinkResourceProperties
}

// SharedPrivateLinkResource - Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
type SharedPrivateLinkResource struct {
	// Describes the properties of a Shared Private Link Resource managed by the Azure Cognitive Search service.
	Properties *SharedPrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// SharedPrivateLinkResourceListResult - Response containing a list of Shared Private Link Resources.
type SharedPrivateLinkResourceListResult struct {
	// The URL to get the next set of shared private link resources, if there are any.
	NextLink *string

	// READ-ONLY; The list of Shared Private Link Resources.
	Value []*SharedPrivateLinkResource
}

// SharedPrivateLinkResourceProperties - Describes the properties of an existing Shared Private Link Resource managed by the
// Azure Cognitive Search service.
type SharedPrivateLinkResourceProperties struct {
	// The group id from the provider of resource the shared private link resource is for.
	GroupID *string

	// The resource id of the resource the shared private link resource is for.
	PrivateLinkResourceID *string

	// The provisioning state of the shared private link resource. Can be Updating, Deleting, Failed, Succeeded or Incomplete.
	ProvisioningState *SharedPrivateLinkResourceProvisioningState

	// The request message for requesting approval of the shared private link resource.
	RequestMessage *string

	// Optional. Can be used to specify the Azure Resource Manager location of the resource to which a shared private link is
	// to be created. This is only required for those resources whose DNS configuration
	// are regional (such as Azure Kubernetes Service).
	ResourceRegion *string

	// Status of the shared private link resource. Can be Pending, Approved, Rejected or Disconnected.
	Status *SharedPrivateLinkResourceStatus
}

// SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginCreateOrUpdate
// method.
type SharedPrivateLinkResourcesClientBeginCreateOrUpdateOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SharedPrivateLinkResourcesClientBeginDeleteOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.BeginDelete
// method.
type SharedPrivateLinkResourcesClientBeginDeleteOptions struct {
	// Resumes the LRO from the provided token.
	ResumeToken string
}

// SharedPrivateLinkResourcesClientGetOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.Get
// method.
type SharedPrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// SharedPrivateLinkResourcesClientListByServiceOptions contains the optional parameters for the SharedPrivateLinkResourcesClient.NewListByServicePager
// method.
type SharedPrivateLinkResourcesClientListByServiceOptions struct {
	// placeholder for future optional parameters
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

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}
