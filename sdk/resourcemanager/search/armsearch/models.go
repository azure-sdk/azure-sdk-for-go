//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

// AdminKeyResult - Response containing the primary and secondary admin API keys for a given search service.
type AdminKeyResult struct {
	// READ-ONLY; The primary admin API key of the search service.
	PrimaryKey *string

	// READ-ONLY; The secondary admin API key of the search service.
	SecondaryKey *string
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

// DataPlaneAADOrAPIKeyAuthOption - Indicates that either the API key or an access token from a Microsoft Entra ID tenant
// can be used for authentication.
type DataPlaneAADOrAPIKeyAuthOption struct {
	// Describes what response the data plane API of a search service would send for requests that failed authentication.
	AADAuthFailureMode *AADAuthFailureMode
}

// DataPlaneAuthOptions - Defines the options for how the search service authenticates a data plane request. This cannot be
// set if 'disableLocalAuth' is set to true.
type DataPlaneAuthOptions struct {
	// Indicates that either the API key or an access token from a Microsoft Entra ID tenant can be used for authentication.
	AADOrAPIKey *DataPlaneAADOrAPIKeyAuthOption

	// Indicates that only the API key can be used for authentication.
	APIKeyOnly any
}

// EncryptionWithCmk - Describes a policy that determines how resources within the search service are to be encrypted with
// customer=managed keys.
type EncryptionWithCmk struct {
	// Describes how a search service should enforce having one or more non-customer-encrypted resources.
	Enforcement *SearchEncryptionWithCmk

	// READ-ONLY; Describes whether the search service is compliant or not with respect to having non-customer-encrypted resources.
	// If a service has more than one non-customer-encrypted resource and 'Enforcement' is
	// 'enabled' then the service will be marked as 'nonCompliant'.
	EncryptionComplianceStatus *SearchEncryptionComplianceStatus
}

// IPRule - The IP restriction rule of the search service.
type IPRule struct {
	// Value corresponding to a single IPv4 address (for example, 123.1.2.3) or an IP range in CIDR format (for example, 123.1.2.3/24)
	// to be allowed.
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

// ListQueryKeysResult - Response containing the query API keys for a given search service.
type ListQueryKeysResult struct {
	// READ-ONLY; Request URL that can be used to query next page of query keys. Returned when the total number of requested query
	// keys exceed maximum page size.
	NextLink *string

	// READ-ONLY; The query keys for the search service.
	Value []*QueryKey
}

// NetworkRuleSet - Network-specific rules that determine how the search service can be reached.
type NetworkRuleSet struct {
	// A list of IP restriction rules used for an IP firewall. Any IPs that do not match the rules are blocked by the firewall.
	// These rules are only applied when the 'publicNetworkAccess' of the search
	// service is 'enabled'.
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

// PrivateEndpointConnection - Describes an existing private endpoint connection to the search service.
type PrivateEndpointConnection struct {
	// Describes the properties of an existing private endpoint connection to the search service.
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

// PrivateEndpointConnectionProperties - Describes the properties of an existing Private Endpoint connection to the search
// service.
type PrivateEndpointConnectionProperties struct {
	// The group id from the provider of resource the private link service connection is for.
	GroupID *string

	// The private endpoint resource from Microsoft.Network provider.
	PrivateEndpoint *PrivateEndpointConnectionPropertiesPrivateEndpoint

	// Describes the current state of an existing Private Link Service connection to the Azure Private Endpoint.
	PrivateLinkServiceConnectionState *PrivateEndpointConnectionPropertiesPrivateLinkServiceConnectionState

	// The provisioning state of the private link service connection. Valid values are Updating, Deleting, Failed, Succeeded,
	// or Incomplete
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

	// Status of the the private link service connection. Valid values are Pending, Approved, Rejected, or Disconnected.
	Status *PrivateLinkServiceConnectionStatus
}

// PrivateLinkResource - Describes a supported private link resource for the search service.
type PrivateLinkResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Describes the properties of a supported private link resource for the search service.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceProperties - Describes the properties of a supported private link resource for the search service. For
// a given API version, this represents the 'supported' groupIds when creating a shared private link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; The group ID of the private link resource.
	GroupID *string

	// READ-ONLY; The list of required members of the private link resource.
	RequiredMembers []*string

	// READ-ONLY; The list of required DNS zone names of the private link resource.
	RequiredZoneNames []*string

	// READ-ONLY; The list of resources that are onboarded to private link service and that are supported by search.
	ShareablePrivateLinkResourceTypes []*ShareablePrivateLinkResourceType
}

// PrivateLinkResourcesResult - Response containing a list of supported Private Link Resources.
type PrivateLinkResourcesResult struct {
	// READ-ONLY; The list of supported Private Link Resources.
	Value []*PrivateLinkResource
}

// QueryKey - Describes an API key for a given search service that has permissions for query operations only.
type QueryKey struct {
	// READ-ONLY; The value of the query API key.
	Key *string

	// READ-ONLY; The name of the query API key; may be empty.
	Name *string
}

// QuotaUsageResult - Describes the quota usage for a particular SKU.
type QuotaUsageResult struct {
	// The currently used up value for the particular search SKU.
	CurrentValue *int32

	// The resource ID of the quota usage SKU endpoint for Microsoft.Search provider.
	ID *string

	// The quota limit for the particular search SKU.
	Limit *int32

	// The unit of measurement for the search SKU.
	Unit *string

	// READ-ONLY; The name of the SKU supported by Azure AI Search.
	Name *QuotaUsageResultName
}

// QuotaUsageResultName - The name of the SKU supported by Azure AI Search.
type QuotaUsageResultName struct {
	// The localized string value for the SKU name.
	LocalizedValue *string

	// The SKU name supported by Azure AI Search.
	Value *string
}

// QuotaUsagesListResult - Response containing the quota usage information for all the supported SKUs of Azure AI Search.
type QuotaUsagesListResult struct {
	// READ-ONLY; Request URL that can be used to query next page of quota usages. Returned when the total number of requested
	// quota usages exceed maximum page size.
	NextLink *string

	// READ-ONLY; The quota usages for the SKUs supported by Azure AI Search.
	Value []*QuotaUsageResult
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

// SKU - Defines the SKU of a search service, which determines billing rate and capacity limits.
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

// Service - Describes a search service and its current state.
type Service struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The identity of the resource.
	Identity *Identity

	// Properties of the search service.
	Properties *ServiceProperties

	// The SKU of the search service, which determines billing rate and capacity limits. This property is required when creating
	// a new search service.
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

// ServiceListResult - Response containing a list of search services.
type ServiceListResult struct {
	// READ-ONLY; Request URL that can be used to query next page of search services. Returned when the total number of requested
	// search services exceed maximum page size.
	NextLink *string

	// READ-ONLY; The list of search services.
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

	// Network-specific rules that determine how the search service may be reached.
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

	// Sets options that control the availability of semantic search. This configuration is only possible for certain search SKUs
	// in certain locations.
	SemanticSearch *SearchSemanticSearch

	// READ-ONLY; The list of private endpoint connections to the search service.
	PrivateEndpointConnections []*PrivateEndpointConnection

	// READ-ONLY; The state of the last provisioning operation performed on the search service. Provisioning is an intermediate
	// state that occurs while service capacity is being established. After capacity is set up,
	// provisioningState changes to either 'succeeded' or 'failed'. Client applications can poll provisioning status (the recommended
	// polling interval is from 30 seconds to one minute) by using the Get
	// Search Service operation to see when an operation is completed. If you are using the free service, this value tends to
	// come back as 'succeeded' directly in the call to Create search service. This is
	// because the free service uses capacity that is already set up.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The list of shared private link resources managed by the search service.
	SharedPrivateLinkResources []*SharedPrivateLinkResource

	// READ-ONLY; The status of the search service. Possible values include: 'running': The search service is running and no provisioning
	// operations are underway. 'provisioning': The search service is being provisioned
	// or scaled up or down. 'deleting': The search service is being deleted. 'degraded': The search service is degraded. This
	// can occur when the underlying search units are not healthy. The search service
	// is most likely operational, but performance might be slow and some requests might be dropped. 'disabled': The search service
	// is disabled. In this state, the service will reject all API requests.
	// 'error': The search service is in an error state. If your service is in the degraded, disabled, or error states, Microsoft
	// is actively investigating the underlying issue. Dedicated services in these
	// states are still chargeable based on the number of search units provisioned.
	Status *SearchServiceStatus

	// READ-ONLY; The details of the search service status.
	StatusDetails *string
}

// ServiceUpdate - The parameters used to update a search service.
type ServiceUpdate struct {
	// The identity of the resource.
	Identity *Identity

	// The geographic location of the resource. This must be one of the supported and registered Azure geo regions (for example,
	// West US, East US, Southeast Asia, and so forth). This property is required
	// when creating a new resource.
	Location *string

	// Properties of the search service.
	Properties *ServiceProperties

	// The SKU of the search service, which determines the billing rate and capacity limits. This property is required when creating
	// a new search service.
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

// ShareablePrivateLinkResourceProperties - Describes the properties of a resource type that has been onboarded to private
// link service and that's supported by search.
type ShareablePrivateLinkResourceProperties struct {
	// READ-ONLY; The description of the resource type that has been onboarded to private link service and that's supported by
	// search.
	Description *string

	// READ-ONLY; The resource provider group id for the resource that has been onboarded to private link service and that's supported
	// by search.
	GroupID *string

	// READ-ONLY; The resource provider type for the resource that has been onboarded to private link service and that's supported
	// by search.
	Type *string
}

// ShareablePrivateLinkResourceType - Describes a resource type that has been onboarded to private link service and that's
// supported by search.
type ShareablePrivateLinkResourceType struct {
	// READ-ONLY; The name of the resource type that has been onboarded to private link service and that's supported by search.
	Name *string

	// READ-ONLY; Describes the properties of a resource type that has been onboarded to private link service and that's supported
	// by search.
	Properties *ShareablePrivateLinkResourceProperties
}

// SharedPrivateLinkResource - Describes a Shared Private Link Resource managed by the search service.
type SharedPrivateLinkResource struct {
	// Describes the properties of a Shared Private Link Resource managed by the search service.
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
// search service.
type SharedPrivateLinkResourceProperties struct {
	// The group id from the provider of resource the shared private link resource is for.
	GroupID *string

	// The resource id of the resource the shared private link resource is for.
	PrivateLinkResourceID *string

	// The provisioning state of the shared private link resource. Valid values are Updating, Deleting, Failed, Succeeded or Incomplete.
	ProvisioningState *SharedPrivateLinkResourceProvisioningState

	// The request message for requesting approval of the shared private link resource.
	RequestMessage *string

	// Optional. Can be used to specify the Azure Resource Manager location of the resource to which a shared private link is
	// to be created. This is only required for those resources whose DNS configuration
	// are regional (such as Azure Kubernetes Service).
	ResourceRegion *string

	// Status of the shared private link resource. Valid values are Pending, Approved, Rejected or Disconnected.
	Status *SharedPrivateLinkResourceStatus
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
