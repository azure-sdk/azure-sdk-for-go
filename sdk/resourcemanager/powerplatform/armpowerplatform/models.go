//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpowerplatform

import "time"

// Account - Definition of the account.
type Account struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The properties that define configuration for the account.
	Properties *AccountProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// AccountList - The response of the list accounts operation.
type AccountList struct {
	// Next page link if any.
	NextLink *string

	// Result of the list accounts operation.
	Value []*Account
}

// AccountProperties - The properties that define configuration for the account.
type AccountProperties struct {
	// The description of the account.
	Description *string

	// READ-ONLY; The internally assigned unique identifier of the resource.
	SystemID *string
}

// AccountsClientCreateOrUpdateOptions contains the optional parameters for the AccountsClient.CreateOrUpdate method.
type AccountsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientDeleteOptions contains the optional parameters for the AccountsClient.Delete method.
type AccountsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientGetOptions contains the optional parameters for the AccountsClient.Get method.
type AccountsClientGetOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListByResourceGroupOptions contains the optional parameters for the AccountsClient.NewListByResourceGroupPager
// method.
type AccountsClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientListBySubscriptionOptions contains the optional parameters for the AccountsClient.NewListBySubscriptionPager
// method.
type AccountsClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// AccountsClientUpdateOptions contains the optional parameters for the AccountsClient.Update method.
type AccountsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePoliciesClientCreateOrUpdateOptions contains the optional parameters for the EnterprisePoliciesClient.CreateOrUpdate
// method.
type EnterprisePoliciesClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePoliciesClientDeleteOptions contains the optional parameters for the EnterprisePoliciesClient.Delete method.
type EnterprisePoliciesClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePoliciesClientGetOptions contains the optional parameters for the EnterprisePoliciesClient.Get method.
type EnterprisePoliciesClientGetOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePoliciesClientListByResourceGroupOptions contains the optional parameters for the EnterprisePoliciesClient.NewListByResourceGroupPager
// method.
type EnterprisePoliciesClientListByResourceGroupOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePoliciesClientListBySubscriptionOptions contains the optional parameters for the EnterprisePoliciesClient.NewListBySubscriptionPager
// method.
type EnterprisePoliciesClientListBySubscriptionOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePoliciesClientUpdateOptions contains the optional parameters for the EnterprisePoliciesClient.Update method.
type EnterprisePoliciesClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// EnterprisePolicy - Definition of the EnterprisePolicy.
type EnterprisePolicy struct {
	// REQUIRED; The kind (type) of Enterprise Policy.
	Kind *EnterprisePolicyKind

	// REQUIRED; The geo-location where the resource lives
	Location *string

	// The identity of the EnterprisePolicy.
	Identity *EnterprisePolicyIdentity

	// The properties that define configuration for the enterprise policy
	Properties *Properties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EnterprisePolicyIdentity - The identity of the EnterprisePolicy.
type EnterprisePolicyIdentity struct {
	// The type of identity used for the EnterprisePolicy. Currently, the only supported type is 'SystemAssigned', which implicitly
	// creates an identity.
	Type *ResourceIdentityType

	// READ-ONLY; The principal id of EnterprisePolicy identity.
	SystemAssignedIdentityPrincipalID *string

	// READ-ONLY; The tenant id associated with the EnterprisePolicy.
	TenantID *string
}

// EnterprisePolicyList - The response of the list EnterprisePolicy operation.
type EnterprisePolicyList struct {
	// Next page link if any.
	NextLink *string

	// Result of the list EnterprisePolicy operation.
	Value []*EnterprisePolicy
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

// KeyProperties - Url and version of the KeyVault Secret
type KeyProperties struct {
	// The identifier of the key vault key used to encrypt data.
	Name *string

	// The version of the identity which will be used to access key vault.
	Version *string
}

// KeyVaultProperties - Settings concerning key vault encryption for a configuration store.
type KeyVaultProperties struct {
	// Uri of KeyVault
	ID *string

	// Identity of the secret that includes name and version.
	Key *KeyProperties
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

// OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
type OperationsClientListOptions struct {
	// placeholder for future optional parameters
}

// PatchAccount - Definition of the account.
type PatchAccount struct {
	// The geo-location where the resource lives
	Location *string

	// The properties that define configuration for the account.
	Properties *AccountProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PatchEnterprisePolicy - Definition of the EnterprisePolicy.
type PatchEnterprisePolicy struct {
	// The identity of the EnterprisePolicy.
	Identity *EnterprisePolicyIdentity

	// The kind (type) of Enterprise Policy.
	Kind *EnterprisePolicyKind

	// The geo-location where the resource lives
	Location *string

	// The properties that define configuration for the enterprise policy
	Properties *Properties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PatchTrackedResource - The resource model definition for an Azure Resource Manager tracked top level resource which has
// 'tags' and a 'location'
type PatchTrackedResource struct {
	// The geo-location where the resource lives
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

// PrivateEndpoint - The Private Endpoint resource.
type PrivateEndpoint struct {
	// READ-ONLY; The ARM identifier for Private Endpoint
	ID *string
}

// PrivateEndpointConnection - A private endpoint connection
type PrivateEndpointConnection struct {
	// Resource properties.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateEndpointConnectionListResult - A list of private endpoint connections
type PrivateEndpointConnectionListResult struct {
	// Array of private endpoint connections
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of the PrivateEndpointConnectProperties.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The resource of private end point.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning state of the private endpoint connection resource.
	ProvisioningState *PrivateEndpointConnectionProvisioningState
}

// PrivateEndpointConnectionsClientCreateOrUpdateOptions contains the optional parameters for the PrivateEndpointConnectionsClient.CreateOrUpdate
// method.
type PrivateEndpointConnectionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
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

// PrivateEndpointConnectionsClientListByEnterprisePolicyOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByEnterprisePolicyPager
// method.
type PrivateEndpointConnectionsClientListByEnterprisePolicyOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResource - A private link resource
type PrivateLinkResource struct {
	// Resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// PrivateLinkResourceListResult - A list of private link resources
type PrivateLinkResourceListResult struct {
	// Array of private link resources
	Value []*PrivateLinkResource
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

// PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get method.
type PrivateLinkResourcesClientGetOptions struct {
	// placeholder for future optional parameters
}

// PrivateLinkResourcesClientListByEnterprisePolicyOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByEnterprisePolicyPager
// method.
type PrivateLinkResourcesClientListByEnterprisePolicyOptions struct {
	// placeholder for future optional parameters
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

// Properties - The properties that define configuration for the enterprise policy.
type Properties struct {
	// The encryption settings for a configuration store.
	Encryption *PropertiesEncryption

	// The health status of the resource.
	HealthStatus *HealthStatus

	// Settings concerning lockbox.
	Lockbox *PropertiesLockbox

	// Settings concerning network injection.
	NetworkInjection *PropertiesNetworkInjection

	// READ-ONLY; The internally assigned unique identifier of the resource.
	SystemID *string
}

// PropertiesEncryption - The encryption settings for a configuration store.
type PropertiesEncryption struct {
	// Key vault properties.
	KeyVault *KeyVaultProperties

	// The state of onboarding, which only appears in the response.
	State *State
}

// PropertiesLockbox - Settings concerning lockbox.
type PropertiesLockbox struct {
	// lockbox configuration
	State *State
}

// PropertiesNetworkInjection - Settings concerning network injection.
type PropertiesNetworkInjection struct {
	// Network injection configuration
	VirtualNetworks *VirtualNetworkPropertiesList
}

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
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

// SubnetProperties - Properties of a subnet.
type SubnetProperties struct {
	// Subnet name.
	Name *string
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

// VirtualNetworkProperties - Settings concerning the virtual network.
type VirtualNetworkProperties struct {
	// Uri of the virtual network.
	ID *string

	// Properties of a subnet.
	Subnet *SubnetProperties
}

// VirtualNetworkPropertiesList - A list of private link resources
type VirtualNetworkPropertiesList struct {
	// Next page link if any.
	NextLink *string

	// Array of virtual networks.
	Value []*VirtualNetworkProperties
}
