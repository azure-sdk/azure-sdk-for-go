//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armloadtesting

import "time"

// AzureResourceManagerCommonTypesManagedServiceIdentity - Managed service identity (system assigned and/or user assigned
// identities)
type AzureResourceManagerCommonTypesManagedServiceIdentity struct {
	// The type of managed identity assigned to this resource.
	Type *ManagedServiceIdentityType

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity
}

// CheckQuotaAvailabilityResponse - Check quota availability response object.
type CheckQuotaAvailabilityResponse struct {
	// Check quota availability response properties.
	Properties *CheckQuotaAvailabilityResponseProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData
}

// CheckQuotaAvailabilityResponseProperties - Check quota availability response properties.
type CheckQuotaAvailabilityResponseProperties struct {
	// Message indicating additional details to add to quota support request.
	AvailabilityStatus *string

	// True/False indicating whether the quota request be granted based on availability.
	IsAvailable *bool
}

// EncryptionProperties - Key and identity details for Customer Managed Key encryption of load test resource.
type EncryptionProperties struct {
	// All identity configuration for Customer-managed key settings defining which identity should be used to auth to Key Vault.
	Identity *EncryptionPropertiesIdentity

	// key encryption key Url, versioned. Ex: https://contosovault.vault.azure.net/keys/contosokek/562a4bb76b524a1493a6afe8e536ee78
	// or https://contosovault.vault.azure.net/keys/contosokek.
	KeyURL *string
}

// EncryptionPropertiesIdentity - All identity configuration for Customer-managed key settings defining which identity should
// be used to auth to Key Vault.
type EncryptionPropertiesIdentity struct {
	// User assigned identity to use for accessing key encryption key Url. Ex: /subscriptions/fa5fc227-a624-475e-b696-cdd604c735bc/resourceGroups/
	// /providers/Microsoft.ManagedIdentity/userAssignedIdentities/myId.
	ResourceID *string

	// Managed identity type to use for accessing encryption key Url.
	Type *Type
}

// EndpointDependency - A domain name and connection details used to access a dependency.
type EndpointDependency struct {
	// READ-ONLY; The domain name of the dependency. Domain names may be fully qualified or may contain a * wildcard.
	DomainName *string

	// READ-ONLY; Human-readable supplemental information about the dependency and when it is applicable.
	Description *string

	// READ-ONLY; The list of connection details for this endpoint.
	EndpointDetails []*EndpointDetail
}

// EndpointDetail - Details about the connection between the Batch service and the endpoint.
type EndpointDetail struct {
	// READ-ONLY; The port an endpoint is connected to.
	Port *int32
}

// LoadTestProperties - LoadTest resource properties.
type LoadTestProperties struct {
	// Description of the resource.
	Description *string

	// CMK Encryption property.
	Encryption *EncryptionProperties

	// READ-ONLY; Resource data plane URI.
	DataPlaneURI *string

	// READ-ONLY; Resource provisioning state.
	ProvisioningState *ResourceState
}

// LoadTestResource - LoadTest details.
type LoadTestResource struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *LoadTestProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// LoadTestResourceListResult - The response of a LoadTestResource list operation.
type LoadTestResourceListResult struct {
	// REQUIRED; The LoadTestResource items on this page
	Value []*LoadTestResource

	// The link to the next page of items
	NextLink *string
}

// LoadTestResourceUpdate - The type used for update operations of the LoadTestResource.
type LoadTestResourceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *AzureResourceManagerCommonTypesManagedServiceIdentity

	// The updatable properties of the LoadTestResource.
	Properties *LoadTestResourceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// LoadTestResourceUpdateProperties - The updatable properties of the LoadTestResource.
type LoadTestResourceUpdateProperties struct {
	// Description of the resource.
	Description *string

	// CMK Encryption property.
	Encryption *EncryptionProperties
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

// OutboundEnvironmentEndpoint - A collection of related endpoints from the same service for which the Batch service requires
// outbound access.
type OutboundEnvironmentEndpoint struct {
	// READ-ONLY; The type of service that Azure Load Testing connects to.
	Category *string

	// READ-ONLY; The endpoints for this service to which the Batch service makes outbound calls.
	Endpoints []*EndpointDependency
}

// PagedOutboundEnvironmentEndpoint - Values returned by the List operation.
type PagedOutboundEnvironmentEndpoint struct {
	// REQUIRED; The OutboundEnvironmentEndpoint items on this page
	Value []*OutboundEnvironmentEndpoint

	// The link to the next page of items
	NextLink *string
}

// QuotaBucketRequest - Request object of new quota for a quota bucket.
type QuotaBucketRequest struct {
	// Request object of new quota for a quota bucket.
	Properties *QuotaBucketRequestProperties
}

// QuotaBucketRequestProperties - New quota request request properties.
type QuotaBucketRequestProperties struct {
	// Current quota limit of the quota bucket.
	CurrentQuota *int32

	// Current quota usage of the quota bucket.
	CurrentUsage *int32

	// Dimensions for new quota request.
	Dimensions *QuotaBucketRequestPropertiesDimensions

	// New quota limit of the quota bucket.
	NewQuota *int32
}

// QuotaBucketRequestPropertiesDimensions - Dimensions for new quota request.
type QuotaBucketRequestPropertiesDimensions struct {
	// Location dimension for new quota request of the quota bucket.
	Location *string

	// Subscription Id dimension for new quota request of the quota bucket.
	SubscriptionID *string
}

// QuotaResource - Quota bucket details object.
type QuotaResource struct {
	// The resource-specific properties for this resource.
	Properties *QuotaResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// QuotaResourceListResult - The response of a QuotaResource list operation.
type QuotaResourceListResult struct {
	// REQUIRED; The QuotaResource items on this page
	Value []*QuotaResource

	// The link to the next page of items
	NextLink *string
}

// QuotaResourceProperties - Quota bucket resource properties.
type QuotaResourceProperties struct {
	// Current quota limit of the quota bucket.
	Limit *int32

	// Current quota usage of the quota bucket.
	Usage *int32

	// READ-ONLY; Resource provisioning state.
	ProvisioningState *ResourceState
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

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}
