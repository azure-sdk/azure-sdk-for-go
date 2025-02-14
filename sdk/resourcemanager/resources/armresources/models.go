//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresources

import "time"

type APIProfile struct {
	// READ-ONLY; The API version.
	APIVersion *string

	// READ-ONLY; The profile version.
	ProfileVersion *string
}

// Alias - The alias type.
type Alias struct {
	// The default path for an alias.
	DefaultPath *string

	// The default pattern for an alias.
	DefaultPattern *AliasPattern

	// The alias name.
	Name *string

	// The paths for an alias.
	Paths []*AliasPath

	// The type of the alias.
	Type *AliasType

	// READ-ONLY; The default alias path metadata. Applies to the default path and to any alias path that doesn't have metadata
	DefaultMetadata *AliasPathMetadata
}

// AliasPath - The type of the paths for alias.
type AliasPath struct {
	// The API versions.
	APIVersions []*string

	// The path of an alias.
	Path *string

	// The pattern for an alias path.
	Pattern *AliasPattern

	// READ-ONLY; The metadata of the alias path. If missing, fall back to the default metadata of the alias.
	Metadata *AliasPathMetadata
}

type AliasPathMetadata struct {
	// READ-ONLY; The attributes of the token that the alias path is referring to.
	Attributes *AliasPathAttributes

	// READ-ONLY; The type of the token that the alias path is referring to.
	Type *AliasPathTokenType
}

// AliasPattern - The type of the pattern for an alias path.
type AliasPattern struct {
	// The alias pattern phrase.
	Phrase *string

	// The type of alias pattern
	Type *AliasPatternType

	// The alias pattern variable.
	Variable *string
}

// BasicDependency - Deployment dependency information.
type BasicDependency struct {
	// The ID of the dependency.
	ID *string

	// The dependency resource name.
	ResourceName *string

	// The dependency resource type.
	ResourceType *string
}

// DebugSetting - The debug setting.
type DebugSetting struct {
	// Specifies the type of information to log for debugging. The permitted values are none, requestContent, responseContent,
	// or both requestContent and responseContent separated by a comma. The default is
	// none. When setting this value, carefully consider the type of information you are passing in during deployment. By logging
	// information about the request or response, you could potentially expose
	// sensitive data that is retrieved through the deployment operations.
	DetailLevel *string
}

// Dependency - Deployment dependency information.
type Dependency struct {
	// The list of dependencies.
	DependsOn []*BasicDependency

	// The ID of the dependency.
	ID *string

	// The dependency resource name.
	ResourceName *string

	// The dependency resource type.
	ResourceType *string
}

// Deployment operation parameters.
type Deployment struct {
	// REQUIRED; The deployment properties.
	Properties *DeploymentProperties

	// The location to store the deployment data.
	Location *string

	// Deployment tags
	Tags map[string]*string
}

type DeploymentDiagnosticsDefinition struct {
	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; Denotes the additional response level.
	Level *Level

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error target.
	Target *string
}

// DeploymentExportResult - The deployment export result.
type DeploymentExportResult struct {
	// The template content.
	Template any
}

// DeploymentExtended - Deployment information.
type DeploymentExtended struct {
	// the location of the deployment.
	Location *string

	// Deployment properties.
	Properties *DeploymentPropertiesExtended

	// Deployment tags
	Tags map[string]*string

	// READ-ONLY; The ID of the deployment.
	ID *string

	// READ-ONLY; The name of the deployment.
	Name *string

	// READ-ONLY; The type of the deployment.
	Type *string
}

// DeploymentExtendedFilter - Deployment filter.
type DeploymentExtendedFilter struct {
	// The provisioning state.
	ProvisioningState *string
}

// DeploymentListResult - List of deployments.
type DeploymentListResult struct {
	// An array of deployments.
	Value []*DeploymentExtended

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// DeploymentOperation - Deployment operation information.
type DeploymentOperation struct {
	// Deployment properties.
	Properties *DeploymentOperationProperties

	// READ-ONLY; Full deployment operation ID.
	ID *string

	// READ-ONLY; Deployment operation ID.
	OperationID *string
}

// DeploymentOperationProperties - Deployment operation properties.
type DeploymentOperationProperties struct {
	// READ-ONLY; The duration of the operation.
	Duration *string

	// READ-ONLY; The name of the current provisioning operation.
	ProvisioningOperation *ProvisioningOperation

	// READ-ONLY; The state of the provisioning.
	ProvisioningState *string

	// READ-ONLY; The HTTP request message.
	Request *HTTPMessage

	// READ-ONLY; The HTTP response message.
	Response *HTTPMessage

	// READ-ONLY; Deployment operation service request id.
	ServiceRequestID *string

	// READ-ONLY; Operation status code from the resource provider. This property may not be set if a response has not yet been
	// received.
	StatusCode *string

	// READ-ONLY; Operation status message from the resource provider. This property is optional. It will only be provided if
	// an error was received from the resource provider.
	StatusMessage *StatusMessage

	// READ-ONLY; The target resource.
	TargetResource *TargetResource

	// READ-ONLY; The date and time of the operation.
	Timestamp *time.Time
}

// DeploymentOperationsListResult - List of deployment operations.
type DeploymentOperationsListResult struct {
	// An array of deployment operations.
	Value []*DeploymentOperation

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// DeploymentParameter - Deployment parameter for the template.
type DeploymentParameter struct {
	// Azure Key Vault parameter reference.
	Reference *KeyVaultParameterReference

	// Input value to the parameter .
	Value any
}

// DeploymentProperties - Deployment properties.
type DeploymentProperties struct {
	// REQUIRED; The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode,
	// resources are deployed without deleting existing resources that are not included in
	// the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included
	// in the template are deleted. Be careful when using Complete mode as you may
	// unintentionally delete resources.
	Mode *DeploymentMode

	// The debug setting of the deployment.
	DebugSetting *DebugSetting

	// Specifies whether template expressions are evaluated within the scope of the parent template or nested template. Only applicable
	// to nested templates. If not specified, default value is outer.
	ExpressionEvaluationOptions *ExpressionEvaluationOptions

	// The deployment on error behavior.
	OnErrorDeployment *OnErrorDeployment

	// Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide
	// the parameter values directly in the request rather than link to an existing
	// parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or
	// a well formed JSON string.
	Parameters map[string]*DeploymentParameter

	// The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink
	// property or the parameters property, but not both.
	ParametersLink *ParametersLink

	// The template content. You use this element when you want to pass the template syntax directly in the request rather than
	// link to an existing template. It can be a JObject or well-formed JSON string.
	// Use either the templateLink property or the template property, but not both.
	Template any

	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *TemplateLink

	// The validation level of the deployment
	ValidationLevel *ValidationLevel
}

// DeploymentPropertiesExtended - Deployment properties with additional details.
type DeploymentPropertiesExtended struct {
	// The validation level of the deployment
	ValidationLevel *ValidationLevel

	// READ-ONLY; The correlation ID of the deployment.
	CorrelationID *string

	// READ-ONLY; The debug setting of the deployment.
	DebugSetting *DebugSetting

	// READ-ONLY; The list of deployment dependencies.
	Dependencies []*Dependency

	// READ-ONLY; Contains diagnostic information collected during validation process.
	Diagnostics []*DeploymentDiagnosticsDefinition

	// READ-ONLY; The duration of the template deployment.
	Duration *string

	// READ-ONLY; The deployment error.
	Error *ErrorResponse

	// READ-ONLY; The deployment mode. Possible values are Incremental and Complete.
	Mode *DeploymentMode

	// READ-ONLY; The deployment on error behavior.
	OnErrorDeployment *OnErrorDeploymentExtended

	// READ-ONLY; Array of provisioned resources.
	OutputResources []*ResourceReference

	// READ-ONLY; Key/value pairs that represent deployment output.
	Outputs any

	// READ-ONLY; Deployment parameters.
	Parameters any

	// READ-ONLY; The URI referencing the parameters.
	ParametersLink *ParametersLink

	// READ-ONLY; The list of resource providers needed for the deployment.
	Providers []*Provider

	// READ-ONLY; Denotes the state of provisioning.
	ProvisioningState *ProvisioningState

	// READ-ONLY; The hash produced for the template.
	TemplateHash *string

	// READ-ONLY; The URI referencing the template.
	TemplateLink *TemplateLink

	// READ-ONLY; The timestamp of the template deployment.
	Timestamp *time.Time

	// READ-ONLY; Array of validated resources.
	ValidatedResources []*ResourceReference
}

// DeploymentValidateResult - Information from validate template deployment response.
type DeploymentValidateResult struct {
	// The template deployment properties.
	Properties *DeploymentPropertiesExtended

	// READ-ONLY; The deployment validation error.
	Error *ErrorResponse

	// READ-ONLY; The ID of the deployment.
	ID *string

	// READ-ONLY; The name of the deployment.
	Name *string

	// READ-ONLY; The type of the deployment.
	Type *string
}

// DeploymentWhatIf - Deployment What-if operation parameters.
type DeploymentWhatIf struct {
	// REQUIRED; The deployment properties.
	Properties *DeploymentWhatIfProperties

	// The location to store the deployment data.
	Location *string
}

// DeploymentWhatIfProperties - Deployment What-if properties.
type DeploymentWhatIfProperties struct {
	// REQUIRED; The mode that is used to deploy resources. This value can be either Incremental or Complete. In Incremental mode,
	// resources are deployed without deleting existing resources that are not included in
	// the template. In Complete mode, resources are deployed and existing resources in the resource group that are not included
	// in the template are deleted. Be careful when using Complete mode as you may
	// unintentionally delete resources.
	Mode *DeploymentMode

	// The debug setting of the deployment.
	DebugSetting *DebugSetting

	// Specifies whether template expressions are evaluated within the scope of the parent template or nested template. Only applicable
	// to nested templates. If not specified, default value is outer.
	ExpressionEvaluationOptions *ExpressionEvaluationOptions

	// The deployment on error behavior.
	OnErrorDeployment *OnErrorDeployment

	// Name and value pairs that define the deployment parameters for the template. You use this element when you want to provide
	// the parameter values directly in the request rather than link to an existing
	// parameter file. Use either the parametersLink property or the parameters property, but not both. It can be a JObject or
	// a well formed JSON string.
	Parameters map[string]*DeploymentParameter

	// The URI of parameters file. You use this element to link to an existing parameters file. Use either the parametersLink
	// property or the parameters property, but not both.
	ParametersLink *ParametersLink

	// The template content. You use this element when you want to pass the template syntax directly in the request rather than
	// link to an existing template. It can be a JObject or well-formed JSON string.
	// Use either the templateLink property or the template property, but not both.
	Template any

	// The URI of the template. Use either the templateLink property or the template property, but not both.
	TemplateLink *TemplateLink

	// The validation level of the deployment
	ValidationLevel *ValidationLevel

	// Optional What-If operation settings.
	WhatIfSettings *DeploymentWhatIfSettings
}

// DeploymentWhatIfSettings - Deployment What-If operation settings.
type DeploymentWhatIfSettings struct {
	// The format of the What-If results
	ResultFormat *WhatIfResultFormat
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info any

	// READ-ONLY; The additional info type.
	Type *string
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.)
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo

	// READ-ONLY; The error code.
	Code *string

	// READ-ONLY; The error details.
	Details []*ErrorResponse

	// READ-ONLY; The error message.
	Message *string

	// READ-ONLY; The error target.
	Target *string
}

// ExportTemplateRequest - Export resource group template request parameters.
type ExportTemplateRequest struct {
	// The export template options. A CSV-formatted list containing zero or more of the following: 'IncludeParameterDefaultValue',
	// 'IncludeComments', 'SkipResourceNameParameterization',
	// 'SkipAllParameterization'
	Options *string

	// The output format for the exported resources.
	OutputFormat *ExportTemplateOutputFormat

	// The IDs of the resources to filter the export by. To export all resources, supply an array with single entry '*'.
	Resources []*string
}

// ExpressionEvaluationOptions - Specifies whether template expressions are evaluated within the scope of the parent template
// or nested template.
type ExpressionEvaluationOptions struct {
	// The scope to be used for evaluation of parameters, variables and functions in a nested template.
	Scope *ExpressionEvaluationOptionsScopeType
}

// ExtendedLocation - Resource extended location.
type ExtendedLocation struct {
	// The extended location name.
	Name *string

	// The extended location type.
	Type *ExtendedLocationType
}

// GenericResource - Resource information.
type GenericResource struct {
	// Resource extended location.
	ExtendedLocation *ExtendedLocation

	// The identity of the resource.
	Identity *Identity

	// The kind of the resource.
	Kind *string

	// Resource location
	Location *string

	// ID of the resource that manages this resource.
	ManagedBy *string

	// The plan of the resource.
	Plan *Plan

	// The resource properties.
	Properties any

	// The SKU of the resource.
	SKU *SKU

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource ID
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; Resource type
	Type *string
}

// GenericResourceExpanded - Resource information.
type GenericResourceExpanded struct {
	// Resource extended location.
	ExtendedLocation *ExtendedLocation

	// The identity of the resource.
	Identity *Identity

	// The kind of the resource.
	Kind *string

	// Resource location
	Location *string

	// ID of the resource that manages this resource.
	ManagedBy *string

	// The plan of the resource.
	Plan *Plan

	// The resource properties.
	Properties any

	// The SKU of the resource.
	SKU *SKU

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; The changed time of the resource. This is only present if requested via the $expand query parameter.
	ChangedTime *time.Time

	// READ-ONLY; The created time of the resource. This is only present if requested via the $expand query parameter.
	CreatedTime *time.Time

	// READ-ONLY; Resource ID
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; The provisioning state of the resource. This is only present if requested via the $expand query parameter.
	ProvisioningState *string

	// READ-ONLY; Resource type
	Type *string
}

// GenericResourceFilter - Resource filter.
type GenericResourceFilter struct {
	// The resource type.
	ResourceType *string

	// The tag name.
	Tagname *string

	// The tag value.
	Tagvalue *string
}

// HTTPMessage - HTTP message.
type HTTPMessage struct {
	// HTTP message content.
	Content any
}

// Identity for the resource.
type Identity struct {
	// The identity type.
	Type *ResourceIdentityType

	// The list of user identities associated with the resource. The user identity dictionary key references will be ARM resource
	// ids in the form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]*IdentityUserAssignedIdentitiesValue

	// READ-ONLY; The principal ID of resource identity.
	PrincipalID *string

	// READ-ONLY; The tenant ID of resource.
	TenantID *string
}

type IdentityUserAssignedIdentitiesValue struct {
	// READ-ONLY; The client id of user assigned identity.
	ClientID *string

	// READ-ONLY; The principal id of user assigned identity.
	PrincipalID *string
}

// KeyVaultParameterReference - Azure Key Vault parameter reference.
type KeyVaultParameterReference struct {
	// REQUIRED; Azure Key Vault reference.
	KeyVault *KeyVaultReference

	// REQUIRED; Azure Key Vault secret name.
	SecretName *string

	// Azure Key Vault secret version.
	SecretVersion *string
}

// KeyVaultReference - Azure Key Vault reference.
type KeyVaultReference struct {
	// REQUIRED; Azure Key Vault resource id.
	ID *string
}

// MoveInfo - Parameters of move resources.
type MoveInfo struct {
	// The IDs of the resources.
	Resources []*string

	// The target resource group.
	TargetResourceGroup *string
}

// OnErrorDeployment - Deployment on error behavior.
type OnErrorDeployment struct {
	// The deployment to be used on error case.
	DeploymentName *string

	// The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
	Type *OnErrorDeploymentType
}

// OnErrorDeploymentExtended - Deployment on error behavior with additional details.
type OnErrorDeploymentExtended struct {
	// The deployment to be used on error case.
	DeploymentName *string

	// The deployment on error behavior type. Possible values are LastSuccessful and SpecificDeployment.
	Type *OnErrorDeploymentType

	// READ-ONLY; The state of the provisioning for the on error deployment.
	ProvisioningState *string
}

// Operation - Microsoft.Resources operation
type Operation struct {
	// The object that represents the operation.
	Display *OperationDisplay

	// Operation name: {provider}/{resource}/{operation}
	Name *string
}

// OperationDisplay - The object that represents the operation.
type OperationDisplay struct {
	// Description of the operation.
	Description *string

	// Operation type: Read, write, delete, etc.
	Operation *string

	// Service provider: Microsoft.Resources
	Provider *string

	// Resource on which the operation is performed: Profile, endpoint, etc.
	Resource *string
}

// OperationListResult - Result of the request to list Microsoft.Resources operations. It contains a list of operations and
// a URL link to get the next set of results.
type OperationListResult struct {
	// URL to get the next set of operation list results if there are any.
	NextLink *string

	// List of Microsoft.Resources operations.
	Value []*Operation
}

// ParametersLink - Entity representing the reference to the deployment parameters.
type ParametersLink struct {
	// REQUIRED; The URI of the parameters file.
	URI *string

	// If included, must match the ContentVersion in the template.
	ContentVersion *string
}

// Permission - Role definition permissions.
type Permission struct {
	// Allowed actions.
	Actions []*string

	// Allowed Data actions.
	DataActions []*string

	// Denied actions.
	NotActions []*string

	// Denied Data actions.
	NotDataActions []*string
}

// Plan for the resource.
type Plan struct {
	// The plan ID.
	Name *string

	// The offer ID.
	Product *string

	// The promotion code.
	PromotionCode *string

	// The publisher ID.
	Publisher *string

	// The plan's version.
	Version *string
}

// Provider - Resource provider information.
type Provider struct {
	// The namespace of the resource provider.
	Namespace *string

	// The provider authorization consent state.
	ProviderAuthorizationConsentState *ProviderAuthorizationConsentState

	// READ-ONLY; The provider ID.
	ID *string

	// READ-ONLY; The registration policy of the resource provider.
	RegistrationPolicy *string

	// READ-ONLY; The registration state of the resource provider.
	RegistrationState *string

	// READ-ONLY; The collection of provider resource types.
	ResourceTypes []*ProviderResourceType
}

// ProviderConsentDefinition - The provider consent.
type ProviderConsentDefinition struct {
	// A value indicating whether authorization is consented or not.
	ConsentToAuthorization *bool
}

// ProviderExtendedLocation - The provider extended location.
type ProviderExtendedLocation struct {
	// The extended locations for the azure location.
	ExtendedLocations []*string

	// The azure location.
	Location *string

	// The extended location type.
	Type *string
}

// ProviderListResult - List of resource providers.
type ProviderListResult struct {
	// An array of resource providers.
	Value []*Provider

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// ProviderPermission - The provider permission
type ProviderPermission struct {
	// The application id.
	ApplicationID *string

	// Role definition properties.
	ManagedByRoleDefinition *RoleDefinition

	// The provider authorization consent state.
	ProviderAuthorizationConsentState *ProviderAuthorizationConsentState

	// Role definition properties.
	RoleDefinition *RoleDefinition
}

// ProviderPermissionListResult - List of provider permissions.
type ProviderPermissionListResult struct {
	// An array of provider permissions.
	Value []*ProviderPermission

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// ProviderRegistrationRequest - The provider registration definition.
type ProviderRegistrationRequest struct {
	// The provider consent.
	ThirdPartyProviderConsent *ProviderConsentDefinition
}

// ProviderResourceType - Resource type managed by the resource provider.
type ProviderResourceType struct {
	// The API version.
	APIVersions []*string

	// The aliases that are supported by this resource type.
	Aliases []*Alias

	// The additional capabilities offered by this resource type.
	Capabilities *string

	// The location mappings that are supported by this resource type.
	LocationMappings []*ProviderExtendedLocation

	// The collection of locations where this resource type can be created.
	Locations []*string

	// The properties.
	Properties map[string]*string

	// The resource type.
	ResourceType *string
	ZoneMappings []*ZoneMapping

	// READ-ONLY; The API profiles for the resource provider.
	APIProfiles []*APIProfile

	// READ-ONLY; The default API version.
	DefaultAPIVersion *string
}

// ProviderResourceTypeListResult - List of resource types of a resource provider.
type ProviderResourceTypeListResult struct {
	// An array of resource types.
	Value []*ProviderResourceType

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// Resource - Specified resource.
type Resource struct {
	// Resource extended location.
	ExtendedLocation *ExtendedLocation

	// Resource location
	Location *string

	// Resource tags
	Tags map[string]*string

	// READ-ONLY; Resource ID
	ID *string

	// READ-ONLY; Resource name
	Name *string

	// READ-ONLY; Resource type
	Type *string
}

// ResourceGroup - Resource group information.
type ResourceGroup struct {
	// REQUIRED; The location of the resource group. It cannot be changed after the resource group has been created. It must be
	// one of the supported Azure locations.
	Location *string

	// The ID of the resource that manages this resource group.
	ManagedBy *string

	// The resource group properties.
	Properties *ResourceGroupProperties

	// The tags attached to the resource group.
	Tags map[string]*string

	// READ-ONLY; The ID of the resource group.
	ID *string

	// READ-ONLY; The name of the resource group.
	Name *string

	// READ-ONLY; The type of the resource group.
	Type *string
}

// ResourceGroupExportResult - Resource group export result.
type ResourceGroupExportResult struct {
	// The template export error.
	Error *ErrorResponse

	// The formatted export content. Used if outputFormat is set to 'Bicep'.
	Output *string

	// The template content. Used if outputFormat is empty or set to 'Json'.
	Template any
}

// ResourceGroupFilter - Resource group filter.
type ResourceGroupFilter struct {
	// The tag name.
	TagName *string

	// The tag value.
	TagValue *string
}

// ResourceGroupListResult - List of resource groups.
type ResourceGroupListResult struct {
	// An array of resource groups.
	Value []*ResourceGroup

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// ResourceGroupPatchable - Resource group information.
type ResourceGroupPatchable struct {
	// The ID of the resource that manages this resource group.
	ManagedBy *string

	// The name of the resource group.
	Name *string

	// The resource group properties.
	Properties *ResourceGroupProperties

	// The tags attached to the resource group.
	Tags map[string]*string
}

// ResourceGroupProperties - The resource group properties.
type ResourceGroupProperties struct {
	// READ-ONLY; The provisioning state.
	ProvisioningState *string
}

// ResourceListResult - List of resource groups.
type ResourceListResult struct {
	// An array of resources.
	Value []*GenericResourceExpanded

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// ResourceProviderOperationDisplayProperties - Resource provider operation's display properties.
type ResourceProviderOperationDisplayProperties struct {
	// Operation description.
	Description *string

	// Resource provider operation.
	Operation *string

	// Operation provider.
	Provider *string

	// Operation description.
	Publisher *string

	// Operation resource.
	Resource *string
}

// ResourceReference - The resource Id model.
type ResourceReference struct {
	// READ-ONLY; The fully qualified resource Id.
	ID *string
}

// RoleDefinition - Role definition properties.
type RoleDefinition struct {
	// The role definition ID.
	ID *string

	// If this is a service role.
	IsServiceRole *bool

	// The role definition name.
	Name *string

	// Role definition permissions.
	Permissions []*Permission

	// Role definition assignable scopes.
	Scopes []*string
}

// SKU for the resource.
type SKU struct {
	// The SKU capacity.
	Capacity *int32

	// The SKU family.
	Family *string

	// The SKU model.
	Model *string

	// The SKU name.
	Name *string

	// The SKU size.
	Size *string

	// The SKU tier.
	Tier *string
}

// ScopedDeployment - Deployment operation parameters.
type ScopedDeployment struct {
	// REQUIRED; The location to store the deployment data.
	Location *string

	// REQUIRED; The deployment properties.
	Properties *DeploymentProperties

	// Deployment tags
	Tags map[string]*string
}

// ScopedDeploymentWhatIf - Deployment What-if operation parameters.
type ScopedDeploymentWhatIf struct {
	// REQUIRED; The location to store the deployment data.
	Location *string

	// REQUIRED; The deployment properties.
	Properties *DeploymentWhatIfProperties
}

// StatusMessage - Operation status message object.
type StatusMessage struct {
	// The error reported by the operation.
	Error *ErrorResponse

	// Status of the deployment operation.
	Status *string
}

// SubResource - Sub-resource.
type SubResource struct {
	// Resource ID
	ID *string
}

// TagCount - Tag count.
type TagCount struct {
	// Type of count.
	Type *string

	// Value of count.
	Value *int32
}

// TagDetails - Tag details.
type TagDetails struct {
	// The total number of resources that use the resource tag. When a tag is initially created and has no associated resources,
	// the value is 0.
	Count *TagCount

	// The tag name.
	TagName *string

	// The list of tag values.
	Values []*TagValue

	// READ-ONLY; The tag name ID.
	ID *string
}

// TagValue - Tag information.
type TagValue struct {
	// The tag value count.
	Count *TagCount

	// The tag value.
	TagValue *string

	// READ-ONLY; The tag value ID.
	ID *string
}

// Tags - A dictionary of name and value pairs.
type Tags struct {
	// Dictionary of
	Tags map[string]*string
}

// TagsListResult - List of subscription tags.
type TagsListResult struct {
	// An array of tags.
	Value []*TagDetails

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// TagsPatchResource - Wrapper resource for tags patch API request only.
type TagsPatchResource struct {
	// The operation type for the patch API.
	Operation *TagsPatchOperation

	// The set of tags.
	Properties *Tags
}

// TagsResource - Wrapper resource for tags API requests and responses.
type TagsResource struct {
	// REQUIRED; The set of tags.
	Properties *Tags

	// READ-ONLY; The ID of the tags wrapper resource.
	ID *string

	// READ-ONLY; The name of the tags wrapper resource.
	Name *string

	// READ-ONLY; The type of the tags wrapper resource.
	Type *string
}

// TargetResource - Target resource.
type TargetResource struct {
	// The ID of the resource.
	ID *string

	// The name of the resource.
	ResourceName *string

	// The type of the resource.
	ResourceType *string
}

// TemplateHashResult - Result of the request to calculate template hash. It contains a string of minified template and its
// hash.
type TemplateHashResult struct {
	// The minified template string.
	MinifiedTemplate *string

	// The template hash.
	TemplateHash *string
}

// TemplateLink - Entity representing the reference to the template.
type TemplateLink struct {
	// If included, must match the ContentVersion in the template.
	ContentVersion *string

	// The resource id of a Template Spec. Use either the id or uri property, but not both.
	ID *string

	// The query string (for example, a SAS token) to be used with the templateLink URI.
	QueryString *string

	// The relativePath property can be used to deploy a linked template at a location relative to the parent. If the parent template
	// was linked with a TemplateSpec, this will reference an artifact in the
	// TemplateSpec. If the parent was linked with a URI, the child deployment will be a combination of the parent and relativePath
	// URIs
	RelativePath *string

	// The URI of the template to deploy. Use either the uri or id property, but not both.
	URI *string
}

// WhatIfChange - Information about a single resource change predicted by What-If operation.
type WhatIfChange struct {
	// REQUIRED; Type of change that will be made to the resource when the deployment is executed.
	ChangeType *ChangeType

	// The predicted snapshot of the resource after the deployment is executed.
	After any

	// The snapshot of the resource before the deployment is executed.
	Before any

	// The predicted changes to resource properties.
	Delta []*WhatIfPropertyChange

	// The resource id of the Deployment responsible for this change.
	DeploymentID *string

	// A subset of properties that uniquely identify a Bicep extensible resource because it lacks a resource id like an Azure
	// resource has.
	Identifiers any

	// Resource ID
	ResourceID *string

	// The symbolic name of the resource responsible for this change.
	SymbolicName *string

	// The explanation about why the resource is unsupported by What-If.
	UnsupportedReason *string
}

// WhatIfOperationProperties - Deployment operation properties.
type WhatIfOperationProperties struct {
	// List of resource changes predicted by What-If operation.
	Changes []*WhatIfChange

	// List of resource changes predicted by What-If operation.
	PotentialChanges []*WhatIfChange

	// READ-ONLY; List of resource diagnostics detected by What-If operation.
	Diagnostics []*DeploymentDiagnosticsDefinition
}

// WhatIfOperationResult - Result of the What-If operation. Contains a list of predicted changes and a URL link to get to
// the next set of results.
type WhatIfOperationResult struct {
	// Error when What-If operation fails.
	Error *ErrorResponse

	// What-If operation properties.
	Properties *WhatIfOperationProperties

	// Status of the What-If operation.
	Status *string
}

// WhatIfPropertyChange - The predicted change to the resource property.
type WhatIfPropertyChange struct {
	// REQUIRED; The path of the property.
	Path *string

	// REQUIRED; The type of property change.
	PropertyChangeType *PropertyChangeType

	// The value of the property after the deployment is executed.
	After any

	// The value of the property before the deployment is executed.
	Before any

	// Nested property changes.
	Children []*WhatIfPropertyChange
}

type ZoneMapping struct {
	// The location of the zone mapping.
	Location *string
	Zones    []*string
}
