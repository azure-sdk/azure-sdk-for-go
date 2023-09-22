//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

import "time"

// APIKey - An API key used for authenticating with a configuration store endpoint.
type APIKey struct {
	// READ-ONLY; A connection string that can be used by supporting clients for authentication.
	ConnectionString *string

	// READ-ONLY; The key ID.
	ID *string

	// READ-ONLY; The last time any of the key's properties were modified.
	LastModified *time.Time

	// READ-ONLY; A name for the key describing its usage.
	Name *string

	// READ-ONLY; Whether this key can only be used for read operations.
	ReadOnly *bool

	// READ-ONLY; The value of the key that is used for authentication purposes.
	Value *string
}

// APIKeyListResult - The result of a request to list API keys.
type APIKeyListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*APIKey
}

// CheckNameAvailabilityParameters - Parameters used for checking whether a resource name is available.
type CheckNameAvailabilityParameters struct {
	// REQUIRED; The name to check for availability.
	Name *string

	// REQUIRED; The resource type to check for name availability.
	Type *ConfigurationResourceType
}

// ConfigurationStore - The configuration store along with all resource properties. The Configuration Store will have all
// information to begin utilizing it.
type ConfigurationStore struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// REQUIRED; The sku of the configuration store.
	SKU *SKU

	// The managed identity information, if configured.
	Identity *ResourceIdentity

	// The properties of a configuration store.
	Properties *ConfigurationStoreProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Resource system metadata.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ConfigurationStoreListResult - The result of a request to list configuration stores.
type ConfigurationStoreListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*ConfigurationStore
}

// ConfigurationStoreProperties - The properties of a configuration store.
type ConfigurationStoreProperties struct {
	// Indicates whether the configuration store need to be recovered.
	CreateMode *CreateMode

	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool

	// Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection *bool

	// The encryption settings of the configuration store.
	Encryption *EncryptionProperties

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess

	// The amount of time in days that the configuration store will be retained when it is soft deleted.
	SoftDeleteRetentionInDays *int32

	// READ-ONLY; The creation date of configuration store.
	CreationDate *time.Time

	// READ-ONLY; The DNS endpoint where the configuration store API will be available.
	Endpoint *string

	// READ-ONLY; The list of private endpoint connections that are set up for this resource.
	PrivateEndpointConnections []*PrivateEndpointConnectionReference

	// READ-ONLY; The provisioning state of the configuration store.
	ProvisioningState *ProvisioningState
}

// ConfigurationStorePropertiesUpdateParameters - The properties for updating a configuration store.
type ConfigurationStorePropertiesUpdateParameters struct {
	// Disables all authentication methods other than AAD authentication.
	DisableLocalAuth *bool

	// Property specifying whether protection against purge is enabled for this configuration store.
	EnablePurgeProtection *bool

	// The encryption settings of the configuration store.
	Encryption *EncryptionProperties

	// Control permission for data plane traffic coming from public networks while private endpoint is enabled.
	PublicNetworkAccess *PublicNetworkAccess
}

// ConfigurationStoreUpdateParameters - The parameters for updating a configuration store.
type ConfigurationStoreUpdateParameters struct {
	// The managed identity information for the configuration store.
	Identity *ResourceIdentity

	// The properties for updating a configuration store.
	Properties *ConfigurationStorePropertiesUpdateParameters

	// The SKU of the configuration store.
	SKU *SKU

	// The ARM resource tags.
	Tags map[string]*string
}

// DeletedConfigurationStore - Deleted configuration store information with extended details.
type DeletedConfigurationStore struct {
	// Properties of the deleted configuration store
	Properties *DeletedConfigurationStoreProperties

	// READ-ONLY; The resource ID for the deleted configuration store.
	ID *string

	// READ-ONLY; The name of the configuration store.
	Name *string

	// READ-ONLY; The resource type of the configuration store.
	Type *string
}

// DeletedConfigurationStoreListResult - List of deleted configuration stores
type DeletedConfigurationStoreListResult struct {
	// The URL to get the next set of deleted configuration stores.
	NextLink *string

	// The list of deleted configuration store.
	Value []*DeletedConfigurationStore
}

// DeletedConfigurationStoreProperties - Properties of the deleted configuration store.
type DeletedConfigurationStoreProperties struct {
	// READ-ONLY; The resource id of the original configuration store.
	ConfigurationStoreID *string

	// READ-ONLY; The deleted date.
	DeletionDate *time.Time

	// READ-ONLY; The location of the original configuration store.
	Location *string

	// READ-ONLY; Purge protection status of the original configuration store.
	PurgeProtectionEnabled *bool

	// READ-ONLY; The scheduled purged date.
	ScheduledPurgeDate *time.Time

	// READ-ONLY; Tags of the original configuration store.
	Tags map[string]*string
}

// EncryptionProperties - The encryption settings for a configuration store.
type EncryptionProperties struct {
	// Key vault properties.
	KeyVaultProperties *KeyVaultProperties
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorDetails - The details of the error.
type ErrorDetails struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; Error code.
	Code *string

	// READ-ONLY; Error message indicating why the operation failed.
	Message *string
}

// ErrorResponse - Error response indicates that the service is not able to process the incoming request. The reason is provided
// in the error message.
type ErrorResponse struct {
	// The details of the error.
	Error *ErrorDetails
}

// KeyValue - The key-value resource along with all resource properties.
type KeyValue struct {
	// All key-value properties.
	Properties *KeyValueProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// KeyValueListResult - The result of a request to list key-values.
type KeyValueListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*KeyValue
}

// KeyValueProperties - All key-value properties.
type KeyValueProperties struct {
	// The content type of the key-value's value. Providing a proper content-type can enable transformations of values when they
	// are retrieved by applications.
	ContentType *string

	// A dictionary of tags that can help identify what a key-value may be applicable for.
	Tags map[string]*string

	// The value of the key-value.
	Value *string

	// READ-ONLY; An ETag indicating the state of a key-value within a configuration store.
	ETag *string

	// READ-ONLY; The primary identifier of a key-value. The key is used in unison with the label to uniquely identify a key-value.
	Key *string

	// READ-ONLY; A value used to group key-values. The label is used in unison with the key to uniquely identify a key-value.
	Label *string

	// READ-ONLY; The last time a modifying operation was performed on the given key-value.
	LastModified *time.Time

	// READ-ONLY; A value indicating whether the key-value is locked. A locked key-value may not be modified until it is unlocked.
	Locked *bool
}

// KeyVaultProperties - Settings concerning key vault encryption for a configuration store.
type KeyVaultProperties struct {
	// The client id of the identity which will be used to access key vault.
	IdentityClientID *string

	// The URI of the key vault key used to encrypt data.
	KeyIdentifier *string
}

// LogSpecification - Specifications of the Log for Azure Monitoring
type LogSpecification struct {
	// Blob duration of the log
	BlobDuration *string

	// Localized friendly display name of the log
	DisplayName *string

	// Name of the log
	Name *string
}

// MetricDimension - Specifications of the Dimension of metrics
type MetricDimension struct {
	// Localized friendly display name of the dimension
	DisplayName *string

	// Internal name of the dimension.
	InternalName *string

	// Name of the dimension
	Name *string
}

// MetricSpecification - Specifications of the Metrics for Azure Monitoring
type MetricSpecification struct {
	// Only provide one value for this field. Valid values: Average, Minimum, Maximum, Total, Count.
	AggregationType *string

	// Dimensions of the metric
	Dimensions []*MetricDimension

	// Localized friendly description of the metric
	DisplayDescription *string

	// Localized friendly display name of the metric
	DisplayName *string

	// Optional. If set to true, then zero will be returned for time duration where no metric is emitted/published.
	FillGapWithZero *bool

	// Internal metric name.
	InternalMetricName *string

	// Name of the metric
	Name *string

	// Unit that makes sense for the metric
	Unit *string
}

// NameAvailabilityStatus - The result of a request to check the availability of a resource name.
type NameAvailabilityStatus struct {
	// READ-ONLY; If any, the error message that provides more detail for the reason that the name is not available.
	Message *string

	// READ-ONLY; The value indicating whether the resource name is available.
	NameAvailable *bool

	// READ-ONLY; If any, the reason that the name is not available.
	Reason *string
}

type NetworkSecurityPerimeter struct {
	ID            *string
	Location      *string
	PerimeterGUID *string
}

// NetworkSecurityPerimeterConfiguration - Network security perimeter configuration for a configuration store.
type NetworkSecurityPerimeterConfiguration struct {
	// The resource ID.
	ID *string

	// The name of the network security perimeter configuration.
	Name *string

	// Resource properties.
	Properties *NetworkSecurityPerimeterConfigurationProperties

	// The type of the resource.
	Type *string
}

// NetworkSecurityPerimeterConfigurationListResult - The result of a request to list network security perimeter configurations.
type NetworkSecurityPerimeterConfigurationListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*NetworkSecurityPerimeterConfiguration
}

// NetworkSecurityPerimeterConfigurationProperties - The properties of network security perimeter configuration.
type NetworkSecurityPerimeterConfigurationProperties struct {
	NetworkSecurityPerimeter *NetworkSecurityPerimeter
	Profile                  *NetworkSecurityPerimeterProfile

	// Provisioning issues of network security perimeter configuration.
	ProvisioningIssues []*NetworkSecurityPerimeterProvisioningIssue

	// Provisioning state of network security perimeter configuration.
	ProvisioningState   *string
	ResourceAssociation *NetworkSecurityPerimeterResourceAssociation
}

type NetworkSecurityPerimeterProfile struct {
	AccessRules               []*NetworkSecurityPerimeterProfileAccessRule
	AccessRulesVersion        *string
	DiagnosticSettingsVersion *string
	EnabledLogCategories      []*string
	Name                      *string
}

type NetworkSecurityPerimeterProfileAccessRule struct {
	Name       *string
	Properties *NetworkSecurityPerimeterProfileAccessRuleProperties
}

type NetworkSecurityPerimeterProfileAccessRuleProperties struct {
	AddressPrefixes           []*string
	Direction                 *string
	EmailAddresses            []*string
	FullyQualifiedDomainNames []*string
	NetworkSecurityPerimeters []*NetworkSecurityPerimeter
	PhoneNumbers              []*string
	Subscriptions             []*NspAccessRuleSubscription
}

type NetworkSecurityPerimeterProvisioningIssue struct {
	Name       *string
	Properties *NetworkSecurityPerimeterProvisioningIssueProperties
}

type NetworkSecurityPerimeterProvisioningIssueProperties struct {
	Description          *string
	IssueType            *string
	Severity             *string
	SuggestedAccessRules []*string
	SuggestedResourceIDs []*string
}

type NetworkSecurityPerimeterResourceAssociation struct {
	AccessMode *string
	Name       *string
}

type NspAccessRuleSubscription struct {
	ID *string
}

// OperationDefinition - The definition of a configuration store operation.
type OperationDefinition struct {
	// The display information for the configuration store operation.
	Display *OperationDefinitionDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name: {provider}/{resource}/{operation}.
	Name *string

	// Origin of the operation
	Origin *string

	// Properties of the operation
	Properties *OperationProperties
}

// OperationDefinitionDisplay - The display information for a configuration store operation.
type OperationDefinitionDisplay struct {
	// The description for the operation.
	Description *string

	// The operation that users can perform.
	Operation *string

	// The resource on which the operation is performed.
	Resource *string

	// READ-ONLY; The resource provider name: Microsoft App Configuration."
	Provider *string
}

// OperationDefinitionListResult - The result of a request to list configuration store operations.
type OperationDefinitionListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*OperationDefinition
}

// OperationProperties - Extra Operation properties
type OperationProperties struct {
	// Service specifications of the operation
	ServiceSpecification *ServiceSpecification
}

// PrivateEndpoint - Private endpoint which a connection belongs to.
type PrivateEndpoint struct {
	// The resource Id for private endpoint
	ID *string
}

// PrivateEndpointConnection - A private endpoint connection
type PrivateEndpointConnection struct {
	// The properties of a private endpoint.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// PrivateEndpointConnectionListResult - A list of private endpoint connections
type PrivateEndpointConnectionListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*PrivateEndpointConnection
}

// PrivateEndpointConnectionProperties - Properties of a private endpoint connection.
type PrivateEndpointConnectionProperties struct {
	// REQUIRED; A collection of information about the state of the connection between service consumer and provider.
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState

	// The resource of private endpoint.
	PrivateEndpoint *PrivateEndpoint

	// READ-ONLY; The provisioning status of the private endpoint connection.
	ProvisioningState *ProvisioningState
}

// PrivateEndpointConnectionReference - A reference to a related private endpoint connection.
type PrivateEndpointConnectionReference struct {
	// The properties of a private endpoint connection.
	Properties *PrivateEndpointConnectionProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// PrivateLinkResource - A resource that supports private link capabilities.
type PrivateLinkResource struct {
	// Private link resource properties.
	Properties *PrivateLinkResourceProperties

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// PrivateLinkResourceListResult - A list of private link resources.
type PrivateLinkResourceListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*PrivateLinkResource
}

// PrivateLinkResourceProperties - Properties of a private link resource.
type PrivateLinkResourceProperties struct {
	// READ-ONLY; The private link resource group id.
	GroupID *string

	// READ-ONLY; The private link resource required member names.
	RequiredMembers []*string

	// READ-ONLY; The list of required DNS zone names of the private link resource.
	RequiredZoneNames []*string
}

// PrivateLinkServiceConnectionState - The state of a private link service connection.
type PrivateLinkServiceConnectionState struct {
	// The private link service connection description.
	Description *string

	// The private link service connection status.
	Status *ConnectionStatus

	// READ-ONLY; Any action that is required beyond basic workflow (approve/ reject/ disconnect)
	ActionsRequired *ActionsRequired
}

// RegenerateKeyParameters - The parameters used to regenerate an API key.
type RegenerateKeyParameters struct {
	// The id of the key to regenerate.
	ID *string
}

// Replica - The replica resource.
type Replica struct {
	// The location of the replica.
	Location *string

	// READ-ONLY; The resource ID.
	ID *string

	// READ-ONLY; The name of the replica.
	Name *string

	// READ-ONLY; All replica properties.
	Properties *ReplicaProperties

	// READ-ONLY; Resource system metadata.
	SystemData *SystemData

	// READ-ONLY; The type of the resource.
	Type *string
}

// ReplicaListResult - The result of a request to list replicas.
type ReplicaListResult struct {
	// The URI that can be used to request the next set of paged results.
	NextLink *string

	// The collection value.
	Value []*Replica
}

// ReplicaProperties - All replica properties.
type ReplicaProperties struct {
	// READ-ONLY; The URI of the replica where the replica API will be available.
	Endpoint *string

	// READ-ONLY; The provisioning state of the replica.
	ProvisioningState *ReplicaProvisioningState
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

// ResourceIdentity - An identity that can be associated with a resource.
type ResourceIdentity struct {
	// The type of managed identity used. The type 'SystemAssigned, UserAssigned' includes both an implicitly created identity
	// and a set of user-assigned identities. The type 'None' will remove any
	// identities.
	Type *IdentityType

	// The list of user-assigned identities associated with the resource. The user-assigned identity dictionary keys will be ARM
	// resource ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]*UserIdentity

	// READ-ONLY; The principal id of the identity. This property will only be provided for a system-assigned identity.
	PrincipalID *string

	// READ-ONLY; The tenant id associated with the resource's identity. This property will only be provided for a system-assigned
	// identity.
	TenantID *string
}

// SKU - Describes a configuration store SKU.
type SKU struct {
	// REQUIRED; The SKU name of the configuration store.
	Name *string
}

// ServiceSpecification - Service specification payload
type ServiceSpecification struct {
	// Specifications of the Log for Azure Monitoring
	LogSpecifications []*LogSpecification

	// Specifications of the Metrics for Azure Monitoring
	MetricSpecifications []*MetricSpecification
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

// UserIdentity - A resource identity that is managed by the user of the service.
type UserIdentity struct {
	// READ-ONLY; The client ID of the user-assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the user-assigned identity.
	PrincipalID *string
}
