//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapicenter

import "time"

// API entity.
type API struct {
	// API properties.
	Properties *APIProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// APICollection - Paginated collection of APIs.
type APICollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*API
}

// APIDefinition - API definition entity.
type APIDefinition struct {
	// API version properties entity.
	Properties *APIDefinitionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// APIDefinitionCollection - Paginated collection of API definitions.
type APIDefinitionCollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*APIDefinition
}

// APIDefinitionProperties - API version properties entity.
type APIDefinitionProperties struct {
	// REQUIRED; API version.
	Title *string

	// API definition description.
	Description *string

	// READ-ONLY
	Specification *APIDefinitionPropertiesSpecification
}

type APIDefinitionPropertiesSpecification struct {
	Name    *string
	Version *string
}

// APIProperties - API properties.
type APIProperties struct {
	// REQUIRED; API title.
	Title *string

	// REQUIRED; Kind of API. For example, REST or GraphQL.
	Type     *APIType
	Contacts []*Contact

	// The custom metadata defined for API catalog entities.
	CustomProperties any

	// Description of the API.
	Description           *string
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
	Format *APISpecExportResultFormat

	// The result of the export operation.
	Value *string
}

// APISpecImportSource - The API specification source entity properties.
type APISpecImportSource struct {
	// Format of the API specification source.
	Format        *APISpecImportSourceFormat
	Specification *APISpecImportSourceSpecification

	// Value of the API specification source.
	Value *string
}

type APISpecImportSourceSpecification struct {
	Name  *string
	Value *string
}

// APIVersion - API version entity.
type APIVersion struct {
	// API version properties entity.
	Properties *APIVersionProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// APIVersionCollection - Paginated collection of API versions.
type APIVersionCollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*APIVersion
}

// APIVersionProperties - API version properties entity.
type APIVersionProperties struct {
	// REQUIRED; API version.
	Title *string

	// Current lifecycle stage of the API.
	LifecycleStage *LifecycleStage
}

type Contact struct {
	// Email address of the contact.
	Email *string

	// Name of the contact.
	Name *string

	// URL for the contact.
	URL *string
}

// Deployment - API deployment entity.
type Deployment struct {
	// API deployment entity properties.
	Properties *DeploymentProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// DeploymentCollection - Paginated collection of API deployments.
type DeploymentCollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*Deployment
}

// DeploymentProperties - API deployment entity properties.
type DeploymentProperties struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties any

	// API center-scoped definition resource ID.
	DefinitionID *string

	// Description of the deployment.
	Description *string

	// API center-scoped environment resource ID.
	EnvironmentID *string

	// Server
	Server *DeploymentServer

	// State of API deployment.
	State *DeploymentState

	// Title
	Title *string
}

// DeploymentServer - Server
type DeploymentServer struct {
	// Base runtime URLs for this deployment.
	RuntimeURI []*string
}

// Environment entity.
type Environment struct {
	// Environment properties entity.
	Properties *EnvironmentProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// EnvironmentCollection - Paginated collection of environments.
type EnvironmentCollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*Environment
}

// EnvironmentProperties - Environment properties entity.
type EnvironmentProperties struct {
	// The custom metadata defined for API catalog entities.
	CustomProperties any

	// Description.
	Description *string
	Onboarding  *Onboarding

	// Server information of the environment.
	Server *EnvironmentServer

	// Title
	Title *string

	// Environment type.
	Type *EnvironmentType
}

// EnvironmentServer - Server information of the environment.
type EnvironmentServer struct {
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
	// SPDX license information for the API. The identifier field is mutually exclusive of the URL field.
	Identifier *string

	// Name of the license.
	Name *string

	// URL pointing to the license details. The URL field is mutually exclusive of the identifier field.
	URL *string
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

type MetadataAssignment struct {
	Deprecated *bool

	// The entities this metadata schema component gets applied to.
	Entity   *MetadataAssignmentEntity
	Required *bool
}

// MetadataSchema - Metadata schema entity. Used to define metadata for the entities in API catalog.
type MetadataSchema struct {
	// Metadata schema properties.
	Properties *MetadataSchemaProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// MetadataSchemaCollection - Paginated collection of metadata schemas.
type MetadataSchemaCollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*MetadataSchema
}

// MetadataSchemaExportResult - The metadata schema export result.
type MetadataSchemaExportResult struct {
	Format *MetadataSchemaExportFormat

	// The result of the export operation.
	Value *string
}

// MetadataSchemaProperties - Metadata schema properties.
type MetadataSchemaProperties struct {
	// REQUIRED; YAML schema defining the type.
	Schema     *string
	AssignedTo []*MetadataAssignment
}

type Onboarding struct {
	DeveloperPortalURI []*string

	// Onboarding guide.
	Instructions *string
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

// Service - The service entity.
type Service struct {
	// REQUIRED; The geo-location where the resource lives
	Location *string

	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentity

	// The properties of the service.
	Properties *ServiceProperties

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

// ServiceCollection - The response of a Service list operation.
type ServiceCollection struct {
	// REQUIRED; The Service items on this page
	Value []*Service

	// The link to the next page of items
	NextLink *string
}

// ServiceProperties - The properties of the service.
type ServiceProperties struct {
	// READ-ONLY; Provisioning state of the service.
	ProvisioningState *ProvisioningState
}

// ServiceUpdate - The service properties to be updated.
type ServiceUpdate struct {
	// The properties of the service.
	Properties *ServiceProperties
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
	// Workspace properties.
	Properties *WorkspaceProperties

	// READ-ONLY; Fully qualified resource ID for the resource. E.g. "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	ID *string

	// READ-ONLY; The name of the resource
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string
}

// WorkspaceCollection - Paginated collection of workspaces.
type WorkspaceCollection struct {
	// Total record count number across all pages.
	Count *int64

	// READ-ONLY; Next page link if any.
	NextLink *string

	// READ-ONLY; Page items.
	Value []*Workspace
}

// WorkspaceProperties - Workspace properties.
type WorkspaceProperties struct {
	// REQUIRED; Workspace display name.
	Title *string

	// Workspace description.
	Description *string
}
