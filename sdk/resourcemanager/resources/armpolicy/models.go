//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

import "time"

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

// Assignment - The policy assignment.
type Assignment struct {
	// The managed identity associated with the policy assignment.
	Identity *Identity

	// The location of the policy assignment. Only required when utilizing managed identity.
	Location *string

	// Properties for the policy assignment.
	Properties *AssignmentProperties

	// READ-ONLY; The ID of the policy assignment.
	ID *string

	// READ-ONLY; The name of the policy assignment.
	Name *string

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the policy assignment.
	Type *string
}

// AssignmentListResult - List of policy assignments.
type AssignmentListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// An array of policy assignments.
	Value []*Assignment
}

// AssignmentProperties - The policy assignment properties.
type AssignmentProperties struct {
	// This message will be part of response in case of policy violation.
	Description *string

	// The display name of the policy assignment.
	DisplayName *string

	// The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
	EnforcementMode *EnforcementMode

	// The policy assignment metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata any

	// The messages that describe why a resource is non-compliant with the policy.
	NonComplianceMessages []*NonComplianceMessage

	// The policy's excluded scopes.
	NotScopes []*string

	// The policy property value override.
	Overrides []*Override

	// The parameter values for the assigned policy rule. The keys are the parameter names.
	Parameters map[string]*ParameterValuesValue

	// The ID of the policy definition or policy set definition being assigned.
	PolicyDefinitionID *string

	// The resource selector list to filter policies by resource properties.
	ResourceSelectors []*ResourceSelector

	// READ-ONLY; The scope for the policy assignment.
	Scope *string
}

// AssignmentUpdate - The policy assignment for Patch request.
type AssignmentUpdate struct {
	// The managed identity associated with the policy assignment.
	Identity *Identity

	// The location of the policy assignment. Only required when utilizing managed identity.
	Location *string

	// The policy assignment properties for Patch request.
	Properties *AssignmentUpdateProperties
}

// AssignmentUpdateProperties - The policy assignment properties for Patch request.
type AssignmentUpdateProperties struct {
	// The policy property value override.
	Overrides []*Override

	// The resource selector list to filter policies by resource properties.
	ResourceSelectors []*ResourceSelector
}

// DataEffect - The data effect definition.
type DataEffect struct {
	// The data effect details schema.
	DetailsSchema any

	// The data effect name.
	Name *string
}

// DataManifestCustomResourceFunctionDefinition - The custom resource function definition.
type DataManifestCustomResourceFunctionDefinition struct {
	// A value indicating whether the custom properties within the property bag are allowed. Needs api-version to be specified
	// in the policy rule eg - vault('2019-06-01').
	AllowCustomProperties *bool

	// The top-level properties that can be selected on the function's output. eg - [ "name", "location" ] if vault().name and
	// vault().location are supported
	DefaultProperties []*string

	// The fully qualified control plane resource type that this function represents. eg - 'Microsoft.KeyVault/vaults'.
	FullyQualifiedResourceType *string

	// The function name as it will appear in the policy rule. eg - 'vault'.
	Name *string
}

// DataManifestResourceFunctionsDefinition - The resource functions supported by a manifest
type DataManifestResourceFunctionsDefinition struct {
	// An array of data manifest custom resource definition.
	Custom []*DataManifestCustomResourceFunctionDefinition

	// The standard resource functions (subscription and/or resourceGroup).
	Standard []*string
}

// DataPolicyManifest - The data policy manifest.
type DataPolicyManifest struct {
	// The data policy manifest properties.
	Properties *DataPolicyManifestProperties

	// READ-ONLY; The ID of the data policy manifest.
	ID *string

	// READ-ONLY; The name of the data policy manifest (it's the same as the Policy Mode).
	Name *string

	// READ-ONLY; The type of the resource (Microsoft.Authorization/dataPolicyManifests).
	Type *string
}

// DataPolicyManifestListResult - List of data policy manifests.
type DataPolicyManifestListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// An array of data policy manifests.
	Value []*DataPolicyManifest
}

// DataPolicyManifestProperties - The properties of the data policy manifest.
type DataPolicyManifestProperties struct {
	// The effect definition.
	Effects []*DataEffect

	// The non-alias field accessor values that can be used in the policy rule.
	FieldValues []*string

	// A value indicating whether policy mode is allowed only in built-in definitions.
	IsBuiltInOnly *bool

	// The list of namespaces for the data policy manifest.
	Namespaces []*string

	// The policy mode of the data policy manifest.
	PolicyMode *string

	// The resource functions definition specified in the data manifest.
	ResourceFunctions *DataManifestResourceFunctionsDefinition

	// An array of resource type aliases.
	ResourceTypeAliases []*ResourceTypeAliases
}

// Definition - The policy definition.
type Definition struct {
	// The policy definition properties.
	Properties *DefinitionProperties

	// READ-ONLY; The ID of the policy definition.
	ID *string

	// READ-ONLY; The name of the policy definition.
	Name *string

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource (Microsoft.Authorization/policyDefinitions).
	Type *string
}

// DefinitionGroup - The policy definition group.
type DefinitionGroup struct {
	// REQUIRED; The name of the group.
	Name *string

	// A resource ID of a resource that contains additional metadata about the group.
	AdditionalMetadataID *string

	// The group's category.
	Category *string

	// The group's description.
	Description *string

	// The group's display name.
	DisplayName *string
}

// DefinitionListResult - List of policy definitions.
type DefinitionListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// An array of policy definitions.
	Value []*Definition
}

// DefinitionProperties - The policy definition properties.
type DefinitionProperties struct {
	// The policy definition description.
	Description *string

	// The display name of the policy definition.
	DisplayName *string

	// The policy definition metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata any

	// The policy definition mode. Some examples are All, Indexed, Microsoft.KeyVault.Data.
	Mode *string

	// The parameter definitions for parameters used in the policy rule. The keys are the parameter names.
	Parameters map[string]*ParameterDefinitionsValue

	// The policy rule.
	PolicyRule any

	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *PolicyType
}

// DefinitionReference - The policy definition reference.
type DefinitionReference struct {
	// REQUIRED; The ID of the policy definition or policy set definition.
	PolicyDefinitionID *string

	// The name of the groups that this policy definition reference belongs to.
	GroupNames []*string

	// The parameter values for the referenced policy rule. The keys are the parameter names.
	Parameters map[string]*ParameterValuesValue

	// A unique id (within the policy set definition) for this policy definition reference.
	PolicyDefinitionReferenceID *string
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

// Exemption - The policy exemption.
type Exemption struct {
	// REQUIRED; Properties for the policy exemption.
	Properties *ExemptionProperties

	// READ-ONLY; The ID of the policy exemption.
	ID *string

	// READ-ONLY; The name of the policy exemption.
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource (Microsoft.Authorization/policyExemptions).
	Type *string
}

// ExemptionListResult - List of policy exemptions.
type ExemptionListResult struct {
	// An array of policy exemptions.
	Value []*Exemption

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// ExemptionProperties - The policy exemption properties.
type ExemptionProperties struct {
	// REQUIRED; The policy exemption category. Possible values are Waiver and Mitigated.
	ExemptionCategory *ExemptionCategory

	// REQUIRED; The ID of the policy assignment that is being exempted.
	PolicyAssignmentID *string

	// The option whether validate the exemption is at or under the assignment scope.
	AssignmentScopeValidation *AssignmentScopeValidation

	// The description of the policy exemption.
	Description *string

	// The display name of the policy exemption.
	DisplayName *string

	// The expiration date and time (in UTC ISO 8601 format yyyy-MM-ddTHH:mm:ssZ) of the policy exemption.
	ExpiresOn *time.Time

	// The policy exemption metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata any

	// The policy definition reference ID list when the associated policy assignment is an assignment of a policy set definition.
	PolicyDefinitionReferenceIDs []*string

	// The resource selector list to filter policies by resource properties.
	ResourceSelectors []*ResourceSelector
}

// ExemptionUpdate - The policy exemption for Patch request.
type ExemptionUpdate struct {
	// The policy exemption properties for Patch request.
	Properties *ExemptionUpdateProperties
}

// ExemptionUpdateProperties - The policy exemption properties for Patch request.
type ExemptionUpdateProperties struct {
	// The option whether validate the exemption is at or under the assignment scope.
	AssignmentScopeValidation *AssignmentScopeValidation

	// The resource selector list to filter policies by resource properties.
	ResourceSelectors []*ResourceSelector
}

// Identity for the resource. Policy assignments support a maximum of one identity. That is either a system assigned identity
// or a single user assigned identity.
type Identity struct {
	// The identity type. This is the only required field when adding a system or user assigned identity to a resource.
	Type *ResourceIdentityType

	// The user identity associated with the policy. The user identity dictionary key references will be ARM resource ids in the
	// form:
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{identityName}'.
	UserAssignedIdentities map[string]*UserAssignedIdentitiesValue

	// READ-ONLY; The principal ID of the resource identity. This property will only be provided for a system assigned identity
	PrincipalID *string

	// READ-ONLY; The tenant ID of the resource identity. This property will only be provided for a system assigned identity
	TenantID *string
}

// NonComplianceMessage - A message that describes why a resource is non-compliant with the policy. This is shown in 'deny'
// error messages and on resource's non-compliant compliance results.
type NonComplianceMessage struct {
	// REQUIRED; A message that describes why a resource is non-compliant with the policy. This is shown in 'deny' error messages
	// and on resource's non-compliant compliance results.
	Message *string

	// The policy definition reference ID within a policy set definition the message is intended for. This is only applicable
	// if the policy assignment assigns a policy set definition. If this is not provided
	// the message applies to all policies assigned by this policy assignment.
	PolicyDefinitionReferenceID *string
}

// Override - The policy property value override.
type Override struct {
	// The override kind.
	Kind *OverrideKind

	// The list of the selector expressions.
	Selectors []*Selector

	// The value to override the policy property.
	Value *string
}

// ParameterDefinitionsValue - The definition of a parameter that can be provided to the policy.
type ParameterDefinitionsValue struct {
	// The allowed values for the parameter.
	AllowedValues []any

	// The default value for the parameter if no value is provided.
	DefaultValue any

	// General metadata for the parameter.
	Metadata *ParameterDefinitionsValueMetadata

	// The data type of the parameter.
	Type *ParameterType
}

// ParameterDefinitionsValueMetadata - General metadata for the parameter.
type ParameterDefinitionsValueMetadata struct {
	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]any

	// Set to true to have Azure portal create role assignments on the resource ID or resource scope value of this parameter during
	// policy assignment. This property is useful in case you wish to assign
	// permissions outside the assignment scope.
	AssignPermissions *bool

	// The description of the parameter.
	Description *string

	// The display name for the parameter.
	DisplayName *string

	// Used when assigning the policy definition through the portal. Provides a context aware list of values for the user to choose
	// from.
	StrongType *string
}

// ParameterValuesValue - The value of a parameter.
type ParameterValuesValue struct {
	// The value of the parameter.
	Value any
}

// ResourceSelector - The resource selector to filter policies by resource properties.
type ResourceSelector struct {
	// The name of the resource selector.
	Name *string

	// The list of the selector expressions.
	Selectors []*Selector
}

// ResourceTypeAliases - The resource type aliases definition.
type ResourceTypeAliases struct {
	// The aliases for property names.
	Aliases []*Alias

	// The resource type name.
	ResourceType *string
}

// Selector - The selector expression.
type Selector struct {
	// The list of values to filter in.
	In []*string

	// The selector kind.
	Kind *SelectorKind

	// The list of values to filter out.
	NotIn []*string
}

// SetDefinition - The policy set definition.
type SetDefinition struct {
	// The policy definition properties.
	Properties *SetDefinitionProperties

	// READ-ONLY; The ID of the policy set definition.
	ID *string

	// READ-ONLY; The name of the policy set definition.
	Name *string

	// READ-ONLY; The system metadata relating to this resource.
	SystemData *SystemData

	// READ-ONLY; The type of the resource (Microsoft.Authorization/policySetDefinitions).
	Type *string
}

// SetDefinitionListResult - List of policy set definitions.
type SetDefinitionListResult struct {
	// The URL to use for getting the next set of results.
	NextLink *string

	// An array of policy set definitions.
	Value []*SetDefinition
}

// SetDefinitionProperties - The policy set definition properties.
type SetDefinitionProperties struct {
	// REQUIRED; An array of policy definition references.
	PolicyDefinitions []*DefinitionReference

	// The policy set definition description.
	Description *string

	// The display name of the policy set definition.
	DisplayName *string

	// The policy set definition metadata. Metadata is an open ended object and is typically a collection of key value pairs.
	Metadata any

	// The policy set definition parameters that can be used in policy definition references.
	Parameters map[string]*ParameterDefinitionsValue

	// The metadata describing groups of policy definition references within the policy set definition.
	PolicyDefinitionGroups []*DefinitionGroup

	// The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
	PolicyType *PolicyType
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

type UserAssignedIdentitiesValue struct {
	// READ-ONLY; The client id of user assigned identity.
	ClientID *string

	// READ-ONLY; The principal id of user assigned identity.
	PrincipalID *string
}

// Variable - The variable.
type Variable struct {
	// REQUIRED; Properties for the variable.
	Properties *VariableProperties

	// READ-ONLY; The ID of the variable.
	ID *string

	// READ-ONLY; The name of the variable.
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource (Microsoft.Authorization/variables).
	Type *string
}

// VariableColumn - The variable column.
type VariableColumn struct {
	// REQUIRED; The name of this policy variable column.
	ColumnName *string
}

// VariableListResult - List of variables.
type VariableListResult struct {
	// An array of variables.
	Value []*Variable

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// VariableProperties - The variable properties.
type VariableProperties struct {
	// REQUIRED; Variable column definitions.
	Columns []*VariableColumn
}

// VariableValue - The variable value.
type VariableValue struct {
	// REQUIRED; Properties for the variable value.
	Properties *VariableValueProperties

	// READ-ONLY; The ID of the variable.
	ID *string

	// READ-ONLY; The name of the variable.
	Name *string

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData

	// READ-ONLY; The type of the resource (Microsoft.Authorization/variables/values).
	Type *string
}

// VariableValueColumnValue - The name value tuple for this variable value column.
type VariableValueColumnValue struct {
	// REQUIRED; Column name for the variable value
	ColumnName *string

	// REQUIRED; Column value for the variable value; this can be an integer, double, boolean, null or a string.
	ColumnValue any
}

// VariableValueListResult - List of variable values.
type VariableValueListResult struct {
	// An array of variable values.
	Value []*VariableValue

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string
}

// VariableValueProperties - The variable value properties.
type VariableValueProperties struct {
	// REQUIRED; Variable value column value array.
	Values []*VariableValueColumnValue
}
