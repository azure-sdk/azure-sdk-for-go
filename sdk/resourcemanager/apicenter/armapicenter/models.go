// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armapicenter

import "time"

// API entity.
type API struct {
	// The resource-specific properties for this resource.
	Properties *APIProperties

	// READ-ONLY; The name of the API.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// APIDefinition - API definition entity.
type APIDefinition struct {
	// The resource-specific properties for this resource.
	Properties *APIDefinitionProperties

	// READ-ONLY; The name of the API definition.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// APIDefinitionListResult - The response of a ApiDefinition list operation.
type APIDefinitionListResult struct {
	// READ-ONLY; The ApiDefinition items on this page
	Value []*APIDefinition

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// APIDefinitionProperties - API definition properties entity.
type APIDefinitionProperties struct {
	// REQUIRED; API definition title.
	Title *string

	// API definition description.
	Description *string

	// READ-ONLY; API specification details.
	Specification *APIDefinitionPropertiesSpecification
}

// APIDefinitionPropertiesSpecification - API specification details.
type APIDefinitionPropertiesSpecification struct {
	// Specification name.
	Name *string

	// Specification version.
	Version *string
}

// APIImportSuccess - The API specification was successfully imported.
type APIImportSuccess struct {
}

// APIListResult - The response of a Api list operation.
type APIListResult struct {
	// READ-ONLY; The Api items on this page
	Value []*API

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// APIProperties - API properties.
type APIProperties struct {
	// REQUIRED; Kind of API. For example, REST or GraphQL.
	Kind *APIKind

	// REQUIRED; API title.
	Title *string

	// The set of contacts
	Contacts []*Contact

	// The custom metadata defined for API catalog entities.
	CustomProperties *CustomProperties

	// Description of the API.
	Description *string

	// The set of external documentation
	ExternalDocumentation []*ExternalDocumentation

	// The license information for the API.
	License *License

	// Short description of the API.
	Summary *string

	// Terms of service for the API.
	TermsOfService *TermsOfService

	// READ-ONLY; Current lifecycle stage of the API.
	LifecycleStage *LifecycleStage
}

// APISpecExportResult - The API specification export result.
type APISpecExportResult struct {
	// The format of exported result
	Format *APISpecExportResultFormat

	// The result of the export operation.
	Value *string
}

// APISpecImportRequest - The API specification source entity properties.
type APISpecImportRequest struct {
	// Format of the API specification source.
	Format *APISpecImportSourceFormat

	// API specification details.
	Specification *APISpecImportRequestSpecification

	// Value of the API specification source.
	Value *string
}

// APISpecImportRequestSpecification - API specification details.
type APISpecImportRequestSpecification struct {
	// Specification name.
	Name *string

	// Specification version.
	Version *string
}

// APIVersion - API version entity.
type APIVersion struct {
	// The resource-specific properties for this resource.
	Properties *APIVersionProperties

	// READ-ONLY; The name of the API version.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// APIVersionListResult - The response of a ApiVersion list operation.
type APIVersionListResult struct {
	// READ-ONLY; The ApiVersion items on this page
	Value []*APIVersion

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// APIVersionProperties - API version properties entity.
type APIVersionProperties struct {
	// REQUIRED; Current lifecycle stage of the API.
	LifecycleStage *LifecycleStage

	// REQUIRED; API version title.
	Title *string
}

// Contact information
type Contact struct {
	// Email address of the contact.
	Email *string

	// Name of the contact.
	Name *string

	// URL for the contact.
	URL *string
}

// CustomProperties - Custom Properties
type CustomProperties struct {
}

// DeletedService - Soft-deleted service entity.
type DeletedService struct {
	// The resource-specific properties for this resource.
	Properties *DeletedServiceProperties

	// READ-ONLY; The name of the deleted service.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DeletedServiceListResult - The response of a DeletedService list operation.
type DeletedServiceListResult struct {
	// READ-ONLY; The DeletedService items on this page
	Value []*DeletedService

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// DeletedServiceProperties - Deleted service properties.
type DeletedServiceProperties struct {
	// UTC date and time when the service will be automatically purged. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ
	// as specified by the ISO 8601 standard.
	ScheduledPurgeDate *time.Time

	// UTC date and time when the service was soft-deleted. The date conforms to the following format: yyyy-MM-ddTHH:mm:ssZ as
	// specified by the ISO 8601 standard.
	SoftDeletionDate *time.Time
}

// Deployment - API deployment entity.
type Deployment struct {
	// The resource-specific properties for this resource.
	Properties *DeploymentProperties

	// READ-ONLY; The name of the API deployment.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DeploymentListResult - The response of a Deployment list operation.
type DeploymentListResult struct {
	// READ-ONLY; The Deployment items on this page
	Value []*Deployment

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// DeploymentProperties - API deployment entity properties.
type DeploymentProperties struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties *CustomProperties

	// API center-scoped definition resource ID.
	DefinitionID *string

	// Description of the deployment.
	Description *string

	// API center-scoped environment resource ID.
	EnvironmentID *string

	// The deployment server
	Server *DeploymentServer

	// State of API deployment.
	State *DeploymentState

	// API deployment title
	Title *string
}

// DeploymentServer - Server
type DeploymentServer struct {
	// Base runtime URLs for this deployment.
	RuntimeURI []*string
}

// Environment entity.
type Environment struct {
	// The resource-specific properties for this resource.
	Properties *EnvironmentProperties

	// READ-ONLY; The name of the environment.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EnvironmentListResult - The response of a Environment list operation.
type EnvironmentListResult struct {
	// READ-ONLY; The Environment items on this page
	Value []*Environment

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// EnvironmentProperties - Environment properties entity.
type EnvironmentProperties struct {
	// REQUIRED; Environment kind.
	Kind *EnvironmentKind

	// REQUIRED; Environment title.
	Title *string

	// The custom metadata defined for API catalog entities.
	CustomProperties *CustomProperties

	// The environment description.
	Description *string

	// Environment onboarding information
	Onboarding *Onboarding

	// Server information of the environment.
	Server *EnvironmentServer
}

// EnvironmentServer - Server information of the environment.
type EnvironmentServer struct {
	// The location of the management portal
	ManagementPortalURI []*string

	// Type of the server that represents the environment.
	Type *EnvironmentServerType
}

// ExternalDocumentation - Additional, external documentation for the API.
type ExternalDocumentation struct {
	// REQUIRED; URL pointing to the documentation.
	URL *string

	// Description of the documentation.
	Description *string

	// Title of the documentation.
	Title *string
}

// License - The license information for the API.
type License struct {
	// SPDX license information for the API. The identifier field is mutually
	// exclusive of the URL field.
	Identifier *string

	// Name of the license.
	Name *string

	// URL pointing to the license details. The URL field is mutually exclusive of the
	// identifier field.
	URL *string
}

// ManagedServiceIdentity - Managed service identity (system assigned and/or user assigned identities)
type ManagedServiceIdentity struct {
	// REQUIRED; The type of managed identity assigned to this resource.
	Type *ManagedServiceIdentityType

	// The identities assigned to this resource by the user.
	UserAssignedIdentities map[string]*UserAssignedIdentity

	// READ-ONLY; The service principal ID of the system assigned identity. This property will only be provided for a system assigned
	// identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of the system assigned identity. This property will only be provided for a system assigned identity.
	TenantID *string
}

// MetadataAssignment - Assignment metadata
type MetadataAssignment struct {
	// Deprecated assignment
	Deprecated *bool

	// The entities this metadata schema component gets applied to.
	Entity *MetadataAssignmentEntity

	// Required assignment
	Required *bool
}

// MetadataSchema - Metadata schema entity. Used to define metadata for the entities in API catalog.
type MetadataSchema struct {
	// The resource-specific properties for this resource.
	Properties *MetadataSchemaProperties

	// READ-ONLY; The name of the metadata schema.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// MetadataSchemaExportRequest - The metadata schema export request.
type MetadataSchemaExportRequest struct {
	// An entity the metadata schema is requested for.
	AssignedTo *MetadataAssignmentEntity
}

// MetadataSchemaExportResult - The metadata schema export result.
type MetadataSchemaExportResult struct {
	// The export format for the schema
	Format *MetadataSchemaExportFormat

	// The result of the export operation.
	Value *string
}

// MetadataSchemaListResult - The response of a MetadataSchema list operation.
type MetadataSchemaListResult struct {
	// READ-ONLY; The MetadataSchema items on this page
	Value []*MetadataSchema

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// MetadataSchemaProperties - Metadata schema properties.
type MetadataSchemaProperties struct {
	// REQUIRED; The schema defining the type.
	Schema *string

	// The assignees
	AssignedTo []*MetadataAssignment
}

// Onboarding information
type Onboarding struct {
	// The location of the development portal
	DeveloperPortalURI []*string

	// Onboarding guide.
	Instructions *string
}

// Operation - Details of a REST API operation, returned from the Resource Provider Operations API
type Operation struct {
	// Extensible enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
	ActionType *ActionType

	// READ-ONLY; Localized display information for this particular operation.
	Display *OperationDisplay

	// READ-ONLY; Whether the operation applies to data-plane. This is "true" for data-plane operations and "false" for Azure
	// Resource Manager/control-plane operations.
	IsDataAction *bool

	// READ-ONLY; The name of the operation, as per Resource-Based Access Control (RBAC). Examples: "Microsoft.Compute/virtualMachines/write",
	// "Microsoft.Compute/virtualMachines/capture/action"
	Name *string

	// READ-ONLY; The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
	// value is "user,system"
	Origin *Origin
}

// OperationDisplay - Localized display information for and operation.
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
	// READ-ONLY; The Operation items on this page
	Value []*Operation

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// Service - The service entity.
type Service struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// READ-ONLY; The name of Azure API Center service.
	Name *string

	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *ServiceProperties

	// Resource tags.
	Tags map[string]*string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// ServiceListResult - The response of a Service list operation.
type ServiceListResult struct {
	// READ-ONLY; The Service items on this page
	Value []*Service

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// ServiceProperties - The properties of the service.
type ServiceProperties struct {
	// Flag used to restore soft-deleted API Center service. If specified and set to 'true' all other properties will be ignored.
	Restore *bool

	// READ-ONLY; Provisioning state of the service.
	ProvisioningState *ProvisioningState
}

// ServiceUpdate - The type used for update operations of the Service.
type ServiceUpdate struct {
	// The managed service identities assigned to this resource.
	Identity *ManagedServiceIdentity

	// The resource-specific properties for this resource.
	Properties *ServiceUpdateProperties

	// Resource tags.
	Tags map[string]*string
}

// ServiceUpdateProperties - The updatable properties of the Service.
type ServiceUpdateProperties struct {
	// Flag used to restore soft-deleted API Center service. If specified and set to 'true' all other properties will be ignored.
	Restore *bool
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

// TermsOfService - Terms of service for the API.
type TermsOfService struct {
	// REQUIRED; URL pointing to the terms of service.
	URL *string
}

// UserAssignedIdentity - User assigned identity properties
type UserAssignedIdentity struct {
	// READ-ONLY; The client ID of the assigned identity.
	ClientID *string

	// READ-ONLY; The principal ID of the assigned identity.
	PrincipalID *string
}

// Workspace entity.
type Workspace struct {
	// The resource-specific properties for this resource.
	Properties *WorkspaceProperties

	// READ-ONLY; The name of the workspace.
	Name *string

	// READ-ONLY; Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	ID *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WorkspaceListResult - The response of a Workspace list operation.
type WorkspaceListResult struct {
	// READ-ONLY; The Workspace items on this page
	Value []*Workspace

	// READ-ONLY; The link to the next page of items
	NextLink *string
}

// WorkspaceProperties - Workspace properties.
type WorkspaceProperties struct {
	// REQUIRED; Workspace title.
	Title *string

	// Workspace description.
	Description *string
}
