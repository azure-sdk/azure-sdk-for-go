//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

import "time"

// APIKeyOwnerEntity - API Key Owner details which can be a user or service account
type APIKeyOwnerEntity struct {
	// API Key owner id
	ID *string

	// Type of the owner service or user account
	Kind *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// APIKeyProperties - API Key Properties
type APIKeyProperties struct {
	// Metadata of the record
	Metadata *SCMetadataEntity

	// Specification of the API Key
	Spec *APIKeySpecEntity
}

// APIKeyRecord - Details API key
type APIKeyRecord struct {
	// Id of the api key
	ID *string

	// Type of api key
	Kind *string

	// API Key Properties
	Properties *APIKeyProperties
}

// APIKeyResourceEntity - API Key Resource details which can be kafka cluster or schema registry cluster
type APIKeyResourceEntity struct {
	// The environment of the api key
	Environment *string

	// Id of the resource
	ID *string

	// Type of the owner which can be service or user account
	Kind *string

	// API URL for accessing or modifying the api key resource object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// APIKeySpecEntity - Spec of the API Key record
type APIKeySpecEntity struct {
	// The description of the API Key
	Description *string

	// The name of the API Key
	Name *string

	// Specification of the cluster
	Owner *APIKeyOwnerEntity

	// Specification of the cluster
	Resource *APIKeyResourceEntity

	// API Key Secret
	Secret *string
}

// AccessCreateRoleBindingRequestModel - Create role binding request model
type AccessCreateRoleBindingRequestModel struct {
	// A CRN that specifies the scope and resource patterns necessary for the role to bind
	CrnPattern *string

	// The principal User or Group to bind the role to
	Principal *string

	// The name of the role to bind to the principal
	RoleName *string
}

// AccessInviteUserAccountModel - Invite User Account model
type AccessInviteUserAccountModel struct {
	// Email of the logged in user
	Email *string

	// Details of the user who is being invited
	InvitedUserDetails *AccessInvitedUserDetails

	// Id of the organization
	OrganizationID *string

	// Upn of the logged in user
	Upn *string
}

// AccessInvitedUserDetails - Details of the user being invited
type AccessInvitedUserDetails struct {
	// Auth type of the user
	AuthType *string

	// UPN/Email of the user who is being invited
	InvitedEmail *string
}

// AccessListClusterSuccessResponse - Details of the clusters returned on successful response
type AccessListClusterSuccessResponse struct {
	// List of clusters
	Data []*ClusterRecord

	// Type of response
	Kind *string

	// Metadata of the list
	Metadata *ListMetadata
}

// AccessListEnvironmentsSuccessResponse - Details of the environments returned on successful response
type AccessListEnvironmentsSuccessResponse struct {
	// Environment list data
	Data []*EnvironmentRecord

	// Type of response
	Kind *string

	// Metadata of the environment list
	Metadata *ListMetadata
}

// AccessListInvitationsSuccessResponse - List invitations success response
type AccessListInvitationsSuccessResponse struct {
	// Data of the invitations list
	Data []*InvitationRecord

	// Type of response
	Kind *string

	// Metadata of the list
	Metadata *ListMetadata
}

// AccessListRoleBindingsSuccessResponse - Details of the role bindings returned on successful response
type AccessListRoleBindingsSuccessResponse struct {
	// List of role binding
	Data []*RoleBindingRecord

	// Type of response
	Kind *string

	// Metadata of the list
	Metadata *ListMetadata
}

// AccessListServiceAccountsSuccessResponse - List service accounts success response
type AccessListServiceAccountsSuccessResponse struct {
	// Data of the service accounts list
	Data []*ServiceAccountRecord

	// Type of response
	Kind *string

	// Metadata of the list
	Metadata *ListMetadata
}

// AccessListUsersSuccessResponse - List users success response
type AccessListUsersSuccessResponse struct {
	// Data of the users list
	Data []*UserRecord

	// Type of response
	Kind *string

	// Metadata of the list
	Metadata *ListMetadata
}

// AccessRoleBindingNameListSuccessResponse - Details of the role binding names returned on successful response
type AccessRoleBindingNameListSuccessResponse struct {
	// List of role binding names
	Data []*string

	// Type of response
	Kind *string

	// Metadata of the list
	Metadata *ListMetadata
}

// AgreementProperties - Terms properties for Marketplace and Confluent.
type AgreementProperties struct {
	// If any version of the terms have been accepted, otherwise false.
	Accepted *bool

	// Link to HTML with Microsoft and Publisher terms.
	LicenseTextLink *string

	// Plan identifier string.
	Plan *string

	// Link to the privacy policy of the publisher.
	PrivacyPolicyLink *string

	// Product identifier string.
	Product *string

	// Publisher identifier string.
	Publisher *string

	// Date and time in UTC of when the terms were accepted. This is empty if Accepted is false.
	RetrieveDatetime *time.Time

	// Terms signature.
	Signature *string
}

// AgreementResource - Agreement Terms definition
type AgreementResource struct {
	// Represents the properties of the resource.
	Properties *AgreementProperties

	// READ-ONLY; The ARM id of the resource.
	ID *string

	// READ-ONLY; The name of the agreement.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData

	// READ-ONLY; The type of the agreement.
	Type *string
}

// AgreementResourceListResponse - Response of a list operation.
type AgreementResourceListResponse struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Results of a list operation.
	Value []*AgreementResource
}

// AzureBlobStorageSinkConnectorServiceInfo - The authentication info when auth_type is azureBlobStorageSinkConnector
type AzureBlobStorageSinkConnectorServiceInfo struct {
	// REQUIRED; The connector service type.
	ConnectorServiceType *ConnectorServiceType

	// Azure Blob Storage Account Key
	StorageAccountKey *string

	// Azure Blob Storage Account Name
	StorageAccountName *string

	// Azure Blob Storage Account Container Name
	StorageContainerName *string
}

// GetConnectorServiceTypeInfoBase implements the ConnectorServiceTypeInfoBaseClassification interface for type AzureBlobStorageSinkConnectorServiceInfo.
func (a *AzureBlobStorageSinkConnectorServiceInfo) GetConnectorServiceTypeInfoBase() *ConnectorServiceTypeInfoBase {
	return &ConnectorServiceTypeInfoBase{
		ConnectorServiceType: a.ConnectorServiceType,
	}
}

// AzureBlobStorageSourceConnectorServiceInfo - The connector service type is AzureBlobStorageSourceConnector
type AzureBlobStorageSourceConnectorServiceInfo struct {
	// REQUIRED; The connector service type.
	ConnectorServiceType *ConnectorServiceType

	// Azure Blob Storage Account Key
	StorageAccountKey *string

	// Azure Blob Storage Account Name
	StorageAccountName *string

	// Azure Blob Storage Account Container Name
	StorageContainerName *string
}

// GetConnectorServiceTypeInfoBase implements the ConnectorServiceTypeInfoBaseClassification interface for type AzureBlobStorageSourceConnectorServiceInfo.
func (a *AzureBlobStorageSourceConnectorServiceInfo) GetConnectorServiceTypeInfoBase() *ConnectorServiceTypeInfoBase {
	return &ConnectorServiceTypeInfoBase{
		ConnectorServiceType: a.ConnectorServiceType,
	}
}

// AzureCosmosDBSinkConnectorServiceInfo - The authentication info when auth_type is AzureCosmosDBSinkConnector
type AzureCosmosDBSinkConnectorServiceInfo struct {
	// REQUIRED; The connector service type.
	ConnectorServiceType *ConnectorServiceType

	// Azure Cosmos Database Connection Endpoint
	CosmosConnectionEndpoint *string

	// Azure Cosmos Database Containers Topic Mapping
	CosmosContainersTopicMapping *string

	// Azure Cosmos Database Name
	CosmosDatabaseName *string

	// Azure Cosmos Database Id Strategy
	CosmosIDStrategy *string

	// Azure Cosmos Database Master Key
	CosmosMasterKey *string
}

// GetConnectorServiceTypeInfoBase implements the ConnectorServiceTypeInfoBaseClassification interface for type AzureCosmosDBSinkConnectorServiceInfo.
func (a *AzureCosmosDBSinkConnectorServiceInfo) GetConnectorServiceTypeInfoBase() *ConnectorServiceTypeInfoBase {
	return &ConnectorServiceTypeInfoBase{
		ConnectorServiceType: a.ConnectorServiceType,
	}
}

// AzureCosmosDBSourceConnectorServiceInfo - The authentication info when auth_type is AzureCosmosDBSourceConnector
type AzureCosmosDBSourceConnectorServiceInfo struct {
	// REQUIRED; The connector service type.
	ConnectorServiceType *ConnectorServiceType

	// Azure Cosmos Database Connection Endpoint
	CosmosConnectionEndpoint *string

	// Azure Cosmos Database Containers Topic Mapping
	CosmosContainersTopicMapping *string

	// Azure Cosmos Database Name
	CosmosDatabaseName *string

	// Azure Cosmos Database Master Key
	CosmosMasterKey *string

	// Azure Cosmos Database Message Key Enabled
	CosmosMessageKeyEnabled *bool

	// Azure Cosmos Database Message Key Field
	CosmosMessageKeyField *string
}

// GetConnectorServiceTypeInfoBase implements the ConnectorServiceTypeInfoBaseClassification interface for type AzureCosmosDBSourceConnectorServiceInfo.
func (a *AzureCosmosDBSourceConnectorServiceInfo) GetConnectorServiceTypeInfoBase() *ConnectorServiceTypeInfoBase {
	return &ConnectorServiceTypeInfoBase{
		ConnectorServiceType: a.ConnectorServiceType,
	}
}

// AzureSynapseAnalyticsSinkConnectorServiceInfo - The authentication info when auth_type is AzureSynapseAnalyticsSinkConnector
type AzureSynapseAnalyticsSinkConnectorServiceInfo struct {
	// REQUIRED; The connector service type.
	ConnectorServiceType *ConnectorServiceType

	// Azure Synapse Dedicated SQL Pool Database Name
	SynapseSQLDatabaseName *string

	// Azure Synapse SQL login details
	SynapseSQLPassword *string

	// Azure Synapse Analytics SQL Server Name
	SynapseSQLServerName *string

	// Azure Synapse SQL login details
	SynapseSQLUser *string
}

// GetConnectorServiceTypeInfoBase implements the ConnectorServiceTypeInfoBaseClassification interface for type AzureSynapseAnalyticsSinkConnectorServiceInfo.
func (a *AzureSynapseAnalyticsSinkConnectorServiceInfo) GetConnectorServiceTypeInfoBase() *ConnectorServiceTypeInfoBase {
	return &ConnectorServiceTypeInfoBase{
		ConnectorServiceType: a.ConnectorServiceType,
	}
}

// ClusterByokEntity - The network associated with this object
type ClusterByokEntity struct {
	// ID of the referred resource
	ID *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// ClusterConfigEntity - The configuration of the Kafka cluster
type ClusterConfigEntity struct {
	// The lifecycle phase of the cluster
	Kind *string
}

// ClusterEnvironmentEntity - The environment to which cluster belongs
type ClusterEnvironmentEntity struct {
	// Environment of the referred resource
	Environment *string

	// ID of the referred resource
	ID *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// ClusterNetworkEntity - The network associated with this object
type ClusterNetworkEntity struct {
	// Environment of the referred resource
	Environment *string

	// ID of the referred resource
	ID *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// ClusterProperties - Service Connector Cluster Properties
type ClusterProperties struct {
	// Metadata of the record
	Metadata *SCMetadataEntity

	// Specification of the cluster
	Spec *SCClusterSpecEntity

	// Specification of the cluster status
	Status *ClusterStatusEntity
}

// ClusterRecord - Details of cluster record
type ClusterRecord struct {
	// Display name of the cluster
	DisplayName *string

	// Id of the cluster
	ID *string

	// Type of cluster
	Kind *string

	// Metadata of the record
	Metadata *MetadataEntity

	// Specification of the cluster
	Spec *ClusterSpecEntity

	// Specification of the cluster
	Status *ClusterStatusEntity
}

// ClusterSpecEntity - Spec of the cluster record
type ClusterSpecEntity struct {
	// The Kafka API cluster endpoint
	APIEndpoint *string

	// The availability zone configuration of the cluster
	Availability *string

	// Specification of the cluster
	Byok *ClusterByokEntity

	// The cloud service provider
	Cloud *string

	// Specification of the cluster
	Config *ClusterConfigEntity

	// The name of the cluster
	DisplayName *string

	// Specification of the cluster
	Environment *ClusterEnvironmentEntity

	// The cluster HTTP request URL.
	HTTPEndpoint *string

	// The bootstrap endpoint used by Kafka clients to connect to the cluster
	KafkaBootstrapEndpoint *string

	// Specification of the cluster
	Network *ClusterNetworkEntity

	// The cloud service provider region
	Region *string

	// type of zone availability
	Zone *string
}

// ClusterStatusEntity - Status of the cluster record
type ClusterStatusEntity struct {
	// The number of Confluent Kafka Units
	Cku *int32

	// The lifecycle phase of the cluster
	Phase *string
}

// ConnectorInfoBase - Connector Info Base properties
type ConnectorInfoBase struct {
	// Connector Class
	ConnectorClass *ConnectorClass

	// Connector Id
	ConnectorID *string

	// Connector Name
	ConnectorName *string

	// Connector Status
	ConnectorState *ConnectorStatus

	// Connector Type
	ConnectorType *ConnectorType
}

// ConnectorResource - Details of connector record
type ConnectorResource struct {
	// REQUIRED; The properties of the Connector
	Properties *ConnectorResourceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ConnectorResourceProperties - The resource properties of the Connector
type ConnectorResourceProperties struct {
	// Connector Info Base
	ConnectorBasicInfo *ConnectorInfoBase

	// Connector Service type info base properties.
	ConnectorServiceTypeInfo ConnectorServiceTypeInfoBaseClassification

	// The connection information consumed by applications.
	PartnerConnectorInfo PartnerInfoBaseClassification
}

// ConnectorServiceTypeInfoBase - The connector service type info
type ConnectorServiceTypeInfoBase struct {
	// REQUIRED; The connector service type.
	ConnectorServiceType *ConnectorServiceType
}

// GetConnectorServiceTypeInfoBase implements the ConnectorServiceTypeInfoBaseClassification interface for type ConnectorServiceTypeInfoBase.
func (c *ConnectorServiceTypeInfoBase) GetConnectorServiceTypeInfoBase() *ConnectorServiceTypeInfoBase {
	return c
}

// CreateAPIKeyModel - Create API Key model
type CreateAPIKeyModel struct {
	// Description of the API Key
	Description *string

	// Name of the API Key
	Name *string
}

// EnvironmentProperties - Environment resource property
type EnvironmentProperties struct {
	// Metadata of the record
	Metadata *SCMetadataEntity

	// Stream governance configuration
	StreamGovernanceConfig *StreamGovernanceConfig
}

// EnvironmentRecord - Details about environment name, metadata and environment id of an environment
type EnvironmentRecord struct {
	// Display name of the user
	DisplayName *string

	// Id of the environment
	ID *string

	// Type of environment
	Kind *string

	// Metadata of the record
	Metadata *MetadataEntity
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

// ErrorResponseBody - Response body of Error
type ErrorResponseBody struct {
	// READ-ONLY; Error code
	Code *string

	// READ-ONLY; Error detail
	Details []*ErrorResponseBody

	// READ-ONLY; Error message
	Message *string

	// READ-ONLY; Error target
	Target *string
}

// GetEnvironmentsResponse - Result of GET request to list Confluent operations.
type GetEnvironmentsResponse struct {
	// URL to get the next set of environment records if there are any.
	NextLink *string

	// List of environments in a confluent organization
	Value []*SCEnvironmentRecord
}

// InvitationRecord - Record of the invitation
type InvitationRecord struct {
	// Accepted date time of the invitation
	AcceptedAt *string

	// Auth type of the user
	AuthType *string

	// Email of the user
	Email *string

	// Expiration date time of the invitation
	ExpiresAt *string

	// Id of the invitation
	ID *string

	// Type of account
	Kind *string

	// Metadata of the record
	Metadata *MetadataEntity

	// Status of the invitation
	Status *string
}

// KafkaAzureBlobStorageSinkConnectorInfo - The partner connector type is KafkaAzureBlobStorageSink
type KafkaAzureBlobStorageSinkConnectorInfo struct {
	// REQUIRED; The partner connector type.
	PartnerConnectorType *PartnerConnectorType

	// Kafka API Key
	APIKey *string

	// Kafka API Key Secret
	APISecret *string

	// Kafka Auth Type
	AuthType *AuthType

	// Flush size
	FlushSize *string

	// Kafka Input Data Format Type
	InputFormat *DataFormatType

	// Maximum Tasks
	MaxTasks *string

	// Kafka Output Data Format Type
	OutputFormat *DataFormatType

	// Kafka Service Account Id
	ServiceAccountID *string

	// Time Interval
	TimeInterval *string

	// Kafka topics list
	Topics []*string

	// Kafka topics directory
	TopicsDir *string
}

// GetPartnerInfoBase implements the PartnerInfoBaseClassification interface for type KafkaAzureBlobStorageSinkConnectorInfo.
func (k *KafkaAzureBlobStorageSinkConnectorInfo) GetPartnerInfoBase() *PartnerInfoBase {
	return &PartnerInfoBase{
		PartnerConnectorType: k.PartnerConnectorType,
	}
}

// KafkaAzureBlobStorageSourceConnectorInfo - The partner connector type is KafkaAzureBlobStorageSource
type KafkaAzureBlobStorageSourceConnectorInfo struct {
	// REQUIRED; The partner connector type.
	PartnerConnectorType *PartnerConnectorType

	// Kafka API Key
	APIKey *string

	// Kafka API Secret
	APISecret *string

	// Kafka Auth Type
	AuthType *AuthType

	// Kafka Input Data Format Type
	InputFormat *DataFormatType

	// Maximum Tasks
	MaxTasks *string

	// Kafka Output Data Format Type
	OutputFormat *DataFormatType

	// Kafka Service Account Id
	ServiceAccountID *string

	// Kafka topics Regex pattern
	TopicRegex *string

	// Kafka topics directory
	TopicsDir *string
}

// GetPartnerInfoBase implements the PartnerInfoBaseClassification interface for type KafkaAzureBlobStorageSourceConnectorInfo.
func (k *KafkaAzureBlobStorageSourceConnectorInfo) GetPartnerInfoBase() *PartnerInfoBase {
	return &PartnerInfoBase{
		PartnerConnectorType: k.PartnerConnectorType,
	}
}

// KafkaAzureCosmosDBSinkConnectorInfo - The partner connector type is KafkaAzureCosmosDBSink
type KafkaAzureCosmosDBSinkConnectorInfo struct {
	// REQUIRED; The partner connector type.
	PartnerConnectorType *PartnerConnectorType

	// Kafka API Key
	APIKey *string

	// Kafka API Key Secret
	APISecret *string

	// Kafka Auth Type
	AuthType *AuthType

	// Flush size
	FlushSize *string

	// Kafka Input Data Format Type
	InputFormat *DataFormatType

	// Maximum Tasks
	MaxTasks *string

	// Kafka Output Data Format Type
	OutputFormat *DataFormatType

	// Kafka Service Account Id
	ServiceAccountID *string

	// Time Interval
	TimeInterval *string

	// Kafka topics list
	Topics []*string

	// Kafka topics directory
	TopicsDir *string
}

// GetPartnerInfoBase implements the PartnerInfoBaseClassification interface for type KafkaAzureCosmosDBSinkConnectorInfo.
func (k *KafkaAzureCosmosDBSinkConnectorInfo) GetPartnerInfoBase() *PartnerInfoBase {
	return &PartnerInfoBase{
		PartnerConnectorType: k.PartnerConnectorType,
	}
}

// KafkaAzureCosmosDBSourceConnectorInfo - The partner connector type is KafkaAzureCosmosDBSource
type KafkaAzureCosmosDBSourceConnectorInfo struct {
	// REQUIRED; The partner connector type.
	PartnerConnectorType *PartnerConnectorType

	// Kafka API Key
	APIKey *string

	// Kafka API Secret
	APISecret *string

	// Kafka Auth Type
	AuthType *AuthType

	// Kafka Input Data Format Type
	InputFormat *DataFormatType

	// Maximum Tasks
	MaxTasks *string

	// Kafka Output Data Format Type
	OutputFormat *DataFormatType

	// Kafka Service Account Id
	ServiceAccountID *string

	// Kafka topics Regex pattern
	TopicRegex *string

	// Kafka topics directory
	TopicsDir *string
}

// GetPartnerInfoBase implements the PartnerInfoBaseClassification interface for type KafkaAzureCosmosDBSourceConnectorInfo.
func (k *KafkaAzureCosmosDBSourceConnectorInfo) GetPartnerInfoBase() *PartnerInfoBase {
	return &PartnerInfoBase{
		PartnerConnectorType: k.PartnerConnectorType,
	}
}

// KafkaAzureSynapseAnalyticsSinkConnectorInfo - The partner connector type is KafkaAzureSynapseAnalyticsSink
type KafkaAzureSynapseAnalyticsSinkConnectorInfo struct {
	// REQUIRED; The partner connector type.
	PartnerConnectorType *PartnerConnectorType

	// Kafka API Key
	APIKey *string

	// Kafka API Key Secret
	APISecret *string

	// Kafka Auth Type
	AuthType *AuthType

	// Flush size
	FlushSize *string

	// Kafka Input Data Format Type
	InputFormat *DataFormatType

	// Maximum Tasks
	MaxTasks *string

	// Kafka Output Data Format Type
	OutputFormat *DataFormatType

	// Kafka Service Account Id
	ServiceAccountID *string

	// Time Interval
	TimeInterval *string

	// Kafka topics list
	Topics []*string

	// Kafka topics directory
	TopicsDir *string
}

// GetPartnerInfoBase implements the PartnerInfoBaseClassification interface for type KafkaAzureSynapseAnalyticsSinkConnectorInfo.
func (k *KafkaAzureSynapseAnalyticsSinkConnectorInfo) GetPartnerInfoBase() *PartnerInfoBase {
	return &PartnerInfoBase{
		PartnerConnectorType: k.PartnerConnectorType,
	}
}

// LinkOrganization - Link an existing Confluent organization
type LinkOrganization struct {
	// REQUIRED; User auth token
	Token *string
}

// ListAccessRequestModel - List Access Request Model
type ListAccessRequestModel struct {
	// Search filters for the request
	SearchFilters map[string]*string
}

// ListClustersSuccessResponse - Result of GET request to list clusters in the environment of a confluent organization
type ListClustersSuccessResponse struct {
	// URL to get the next set of cluster records if there are any.
	NextLink *string

	// List of clusters in an environment of a confluent organization
	Value []*SCClusterRecord
}

// ListConnectorsSuccessResponse - Result of GET request to list connectors in the cluster of a confluent organization
type ListConnectorsSuccessResponse struct {
	// URL to get the next set of connectors records if there are any.
	NextLink *string

	// List of connectors in a cluster of a confluent organization
	Value []*ConnectorResource
}

// ListMetadata - Metadata of the list
type ListMetadata struct {
	// First page of the list
	First *string

	// Last page of the list
	Last *string

	// Next page of the list
	Next *string

	// Previous page of the list
	Prev *string

	// Total size of the list
	TotalSize *int32
}

// ListRegionsSuccessResponse - Result of POST request to list regions supported by confluent
type ListRegionsSuccessResponse struct {
	// List of regions supported by confluent
	Data []*RegionRecord
}

// ListSchemaRegistryClustersResponse - Result of GET request to list schema registry clusters in the environment of a confluent
// organization
type ListSchemaRegistryClustersResponse struct {
	// URL to get the next set of schema registry cluster records if there are any.
	NextLink *string

	// List of schema registry clusters in an environment of a confluent organization
	Value []*SchemaRegistryClusterRecord
}

// ListTopicsSuccessResponse - Result of GET request to list topics in the cluster of a confluent organization
type ListTopicsSuccessResponse struct {
	// URL to get the next set of topics records if there are any.
	NextLink *string

	// List of topics in a cluster of a confluent organization
	Value []*TopicRecord
}

// MetadataEntity - Metadata of the data record
type MetadataEntity struct {
	// Created Date Time
	CreatedAt *string

	// Deleted Date time
	DeletedAt *string

	// Resource name of the record
	ResourceName *string

	// Self lookup url
	Self *string

	// Updated Date time
	UpdatedAt *string
}

// OfferDetail - Confluent Offer detail
type OfferDetail struct {
	// REQUIRED; Offer Id
	ID *string

	// REQUIRED; Offer Plan Id
	PlanID *string

	// REQUIRED; Offer Plan Name
	PlanName *string

	// REQUIRED; Publisher Id
	PublisherID *string

	// REQUIRED; Offer Plan Term unit
	TermUnit *string

	// Private Offer Id
	PrivateOfferID *string

	// Array of Private Offer Ids
	PrivateOfferIDs []*string

	// SaaS Offer Status
	Status *SaaSOfferStatus

	// Offer Plan Term Id
	TermID *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation, e.g., 'Write confluent'.
	Description *string

	// Operation type, e.g., read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.Confluent
	Provider *string

	// Type on which the operation is performed, e.g., 'clusters'.
	Resource *string
}

// OperationListResult - Result of GET request to list Confluent operations.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Confluent operations supported by the Microsoft.Confluent provider.
	Value []*OperationResult
}

// OperationResult - An Confluent REST API operation.
type OperationResult struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Indicates whether the operation is a data action
	IsDataAction *bool

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OrganizationResource - Organization resource.
type OrganizationResource struct {
	// REQUIRED; Organization resource properties
	Properties *OrganizationResourceProperties

	// Location of Organization resource
	Location *string

	// Organization resource tags
	Tags map[string]*string

	// READ-ONLY; The ARM id of the resource.
	ID *string

	// READ-ONLY; The name of the resource.
	Name *string

	// READ-ONLY; Metadata pertaining to creation and last modification of the resource
	SystemData *SystemData

	// READ-ONLY; The type of the resource.
	Type *string
}

// OrganizationResourceListResult - The response of a list operation.
type OrganizationResourceListResult struct {
	// Link to the next set of results, if any.
	NextLink *string

	// Result of a list operation.
	Value []*OrganizationResource
}

// OrganizationResourceProperties - Organization resource property
type OrganizationResourceProperties struct {
	// REQUIRED; Confluent offer detail
	OfferDetail *OfferDetail

	// REQUIRED; Subscriber detail
	UserDetail *UserDetail

	// Link an existing Confluent organization
	LinkOrganization *LinkOrganization

	// READ-ONLY; The creation time of the resource.
	CreatedTime *time.Time

	// READ-ONLY; Id of the Confluent organization.
	OrganizationID *string

	// READ-ONLY; Provision states for confluent RP
	ProvisioningState *ProvisionState

	// READ-ONLY; SSO url for the Confluent organization.
	SsoURL *string
}

// OrganizationResourceUpdate - Organization Resource update
type OrganizationResourceUpdate struct {
	// ARM resource tags
	Tags map[string]*string
}

// PartnerInfoBase - The partner info base
type PartnerInfoBase struct {
	// REQUIRED; The partner connector type.
	PartnerConnectorType *PartnerConnectorType
}

// GetPartnerInfoBase implements the PartnerInfoBaseClassification interface for type PartnerInfoBase.
func (p *PartnerInfoBase) GetPartnerInfoBase() *PartnerInfoBase { return p }

// ProxyResource - The resource model definition for a Azure Resource Manager proxy resource. It will not have tags and a
// location
type ProxyResource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// RegionProperties - Region Properties
type RegionProperties struct {
	// Metadata of the record
	Metadata *SCMetadataEntity

	// Specification of the region
	Spec *RegionSpecEntity
}

// RegionRecord - Details of region record
type RegionRecord struct {
	// Id of the cluster
	ID *string

	// Kind of the cluster
	Kind *string

	// Region Properties
	Properties *RegionProperties
}

// RegionSpecEntity - Region spec details
type RegionSpecEntity struct {
	// Cloud provider name
	Cloud *string

	// Display Name of the region
	Name     *string
	Packages []*string

	// Region name
	RegionName *string
}

// Resource - Common fields that are returned in the response for all Azure Resource Manager resources
type Resource struct {
	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ResourceProviderDefaultErrorResponse - Default error response for resource provider
type ResourceProviderDefaultErrorResponse struct {
	// READ-ONLY; Response body of Error
	Error *ErrorResponseBody
}

// RoleBindingRecord - Details on principal, role name and crn pattern of a role binding
type RoleBindingRecord struct {
	// A CRN that specifies the scope and resource patterns necessary for the role to bind
	CrnPattern *string

	// Id of the role binding
	ID *string

	// The type of the resource.
	Kind *string

	// Metadata of the record
	Metadata *MetadataEntity

	// The principal User or Group to bind the role to
	Principal *string

	// The name of the role to bind to the principal
	RoleName *string
}

// SCClusterByokEntity - The network associated with this object
type SCClusterByokEntity struct {
	// ID of the referred resource
	ID *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// SCClusterNetworkEnvironmentEntity - The environment or the network to which cluster belongs
type SCClusterNetworkEnvironmentEntity struct {
	// Environment of the referred resource
	Environment *string

	// ID of the referred resource
	ID *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// SCClusterRecord - Details of cluster record
type SCClusterRecord struct {
	// Id of the cluster
	ID *string

	// Type of cluster
	Kind *string

	// Display name of the cluster
	Name *string

	// Cluster Properties
	Properties *ClusterProperties

	// Type of the resource
	Type *string
}

// SCClusterSpecEntity - Spec of the cluster record
type SCClusterSpecEntity struct {
	// The Kafka API cluster endpoint
	APIEndpoint *string

	// The availability zone configuration of the cluster
	Availability *string

	// Specification of the cluster byok
	Byok *SCClusterByokEntity

	// The cloud service provider
	Cloud *string

	// Specification of the cluster configuration
	Config *ClusterConfigEntity

	// Specification of the cluster environment
	Environment *SCClusterNetworkEnvironmentEntity

	// The cluster HTTP request URL.
	HTTPEndpoint *string

	// The bootstrap endpoint used by Kafka clients to connect to the cluster
	KafkaBootstrapEndpoint *string

	// The name of the cluster
	Name *string

	// Specification of the cluster network
	Network *SCClusterNetworkEnvironmentEntity

	// Stream governance configuration
	Package *Package

	// The cloud service provider region
	Region *string

	// type of zone availability
	Zone *string
}

// SCConfluentListMetadata - Metadata of the list
type SCConfluentListMetadata struct {
	// First page of the list
	First *string

	// Last page of the list
	Last *string

	// Next page of the list
	Next *string

	// Previous page of the list
	Prev *string

	// Total size of the list
	TotalSize *int32
}

// SCEnvironmentRecord - Details about environment name, metadata and environment id of an environment
type SCEnvironmentRecord struct {
	// Id of the environment
	ID *string

	// Type of environment
	Kind *string

	// Display name of the environment
	Name *string

	// Environment properties
	Properties *EnvironmentProperties

	// Type of the resource
	Type *string
}

// SCMetadataEntity - Metadata of the data record
type SCMetadataEntity struct {
	// Created Date Time
	CreatedTimestamp *string

	// Deleted Date time
	DeletedTimestamp *string

	// Resource name of the record
	ResourceName *string

	// Self lookup url
	Self *string

	// Updated Date time
	UpdatedTimestamp *string
}

// SchemaRegistryClusterEnvironmentRegionEntity - The environment associated with this object
type SchemaRegistryClusterEnvironmentRegionEntity struct {
	// ID of the referred resource
	ID *string

	// API URL for accessing or modifying the referred object
	Related *string

	// CRN reference to the referred resource
	ResourceName *string
}

// SchemaRegistryClusterProperties - Schema Registry Cluster Properties
type SchemaRegistryClusterProperties struct {
	// Metadata of the record
	Metadata *SCMetadataEntity

	// Specification of the schema registry cluster
	Spec *SchemaRegistryClusterSpecEntity

	// Specification of the cluster status
	Status *SchemaRegistryClusterStatusEntity
}

// SchemaRegistryClusterRecord - Details of schema registry cluster record
type SchemaRegistryClusterRecord struct {
	// Id of the cluster
	ID *string

	// Kind of the cluster
	Kind *string

	// Schema Registry Cluster Properties
	Properties *SchemaRegistryClusterProperties
}

// SchemaRegistryClusterSpecEntity - Details of schema registry cluster spec
type SchemaRegistryClusterSpecEntity struct {
	// The cloud service provider
	Cloud *string

	// Environment details of the schema registry cluster
	Environment *SchemaRegistryClusterEnvironmentRegionEntity

	// Http endpoint of the cluster
	HTTPEndpoint *string

	// Name of the schema registry cluster
	Name *string

	// Type of the cluster package Advanced, essentials
	Package *string

	// Region details of the schema registry cluster
	Region *SchemaRegistryClusterEnvironmentRegionEntity
}

// SchemaRegistryClusterStatusEntity - Status of the schema registry cluster record
type SchemaRegistryClusterStatusEntity struct {
	// The lifecycle phase of the cluster
	Phase *string
}

// ServiceAccountRecord - Record of the service account
type ServiceAccountRecord struct {
	// Description of the service account
	Description *string

	// Name of the service account
	DisplayName *string

	// Id of the service account
	ID *string

	// Type of account
	Kind *string

	// Metadata of the record
	Metadata *MetadataEntity
}

// StreamGovernanceConfig - Stream governance configuration
type StreamGovernanceConfig struct {
	// Stream governance configuration
	Package *Package
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

// TopicMetadataEntity - Metadata of the data record
type TopicMetadataEntity struct {
	// Resource name of the record
	ResourceName *string

	// Self lookup url
	Self *string
}

// TopicProperties - Topic Properties
type TopicProperties struct {
	// Config Specification of the topic
	Configs *TopicsRelatedLink

	// Input Config Specification of the topic
	InputConfigs []*TopicsInputConfig

	// Type of topic
	Kind *string

	// Metadata of the record
	Metadata *TopicMetadataEntity

	// Partition Specification of the topic
	Partitions *TopicsRelatedLink

	// Partition count of the topic
	PartitionsCount *string

	// Partition Reassignment Specification of the topic
	PartitionsReassignments *TopicsRelatedLink

	// Replication factor of the topic
	ReplicationFactor *string

	// Topic Id returned by Confluent
	TopicID *string
}

// TopicRecord - Details of topic record
type TopicRecord struct {
	// Topic Properties
	Properties *TopicProperties

	// READ-ONLY; The ARM Resource Id of the Topic
	ID *string

	// READ-ONLY; Display name of the topic
	Name *string

	// READ-ONLY; The type of the resource.
	Type *string
}

// TopicsInputConfig - Topics input config
type TopicsInputConfig struct {
	// Name of the topic input config
	Name *string

	// Value of the topic input config
	Value *string
}

// TopicsRelatedLink - Partition Config spec of the topic record
type TopicsRelatedLink struct {
	// Relationship of the topic
	Related *string
}

// UserDetail - Subscriber detail
type UserDetail struct {
	// REQUIRED; Email address
	EmailAddress *string

	// AAD email address
	AADEmail *string

	// First name
	FirstName *string

	// Last name
	LastName *string

	// User principal name
	UserPrincipalName *string
}

// UserRecord - Record of the user
type UserRecord struct {
	// Auth type of the user
	AuthType *string

	// Email of the user
	Email *string

	// Name of the user
	FullName *string

	// Id of the user
	ID *string

	// Type of account
	Kind *string

	// Metadata of the record
	Metadata *MetadataEntity
}

// ValidationResponse - Validation response from the provider
type ValidationResponse struct {
	// Info from the response
	Info map[string]*string
}
