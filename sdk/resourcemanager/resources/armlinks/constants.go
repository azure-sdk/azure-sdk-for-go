//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlinks

const (
	moduleName    = "github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armlinks"
	moduleVersion = "v2.0.0"
)

// AssignmentType - The type of policy assignment. Possible values are NotSpecified, System, SystemHidden, and Custom. Immutable.
type AssignmentType string

const (
	AssignmentTypeCustom       AssignmentType = "Custom"
	AssignmentTypeNotSpecified AssignmentType = "NotSpecified"
	AssignmentTypeSystem       AssignmentType = "System"
	AssignmentTypeSystemHidden AssignmentType = "SystemHidden"
)

// PossibleAssignmentTypeValues returns the possible values for the AssignmentType const type.
func PossibleAssignmentTypeValues() []AssignmentType {
	return []AssignmentType{
		AssignmentTypeCustom,
		AssignmentTypeNotSpecified,
		AssignmentTypeSystem,
		AssignmentTypeSystemHidden,
	}
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// EnforcementMode - The policy assignment enforcement mode. Possible values are Default and DoNotEnforce.
type EnforcementMode string

const (
	// EnforcementModeDefault - The policy effect is enforced during resource creation or update.
	EnforcementModeDefault EnforcementMode = "Default"
	// EnforcementModeDoNotEnforce - The policy effect is not enforced during resource creation or update.
	EnforcementModeDoNotEnforce EnforcementMode = "DoNotEnforce"
)

// PossibleEnforcementModeValues returns the possible values for the EnforcementMode const type.
func PossibleEnforcementModeValues() []EnforcementMode {
	return []EnforcementMode{
		EnforcementModeDefault,
		EnforcementModeDoNotEnforce,
	}
}

// OverrideKind - The override kind.
type OverrideKind string

const (
	// OverrideKindDefinitionVersion - It will override the definition version property value of the policy assignment.
	OverrideKindDefinitionVersion OverrideKind = "definitionVersion"
	// OverrideKindPolicyEffect - It will override the policy effect type.
	OverrideKindPolicyEffect OverrideKind = "policyEffect"
)

// PossibleOverrideKindValues returns the possible values for the OverrideKind const type.
func PossibleOverrideKindValues() []OverrideKind {
	return []OverrideKind{
		OverrideKindDefinitionVersion,
		OverrideKindPolicyEffect,
	}
}

// ParameterType - The data type of the parameter.
type ParameterType string

const (
	ParameterTypeArray    ParameterType = "Array"
	ParameterTypeBoolean  ParameterType = "Boolean"
	ParameterTypeDateTime ParameterType = "DateTime"
	ParameterTypeFloat    ParameterType = "Float"
	ParameterTypeInteger  ParameterType = "Integer"
	ParameterTypeObject   ParameterType = "Object"
	ParameterTypeString   ParameterType = "String"
)

// PossibleParameterTypeValues returns the possible values for the ParameterType const type.
func PossibleParameterTypeValues() []ParameterType {
	return []ParameterType{
		ParameterTypeArray,
		ParameterTypeBoolean,
		ParameterTypeDateTime,
		ParameterTypeFloat,
		ParameterTypeInteger,
		ParameterTypeObject,
		ParameterTypeString,
	}
}

// PolicyType - The type of policy definition. Possible values are NotSpecified, BuiltIn, Custom, and Static.
type PolicyType string

const (
	PolicyTypeBuiltIn      PolicyType = "BuiltIn"
	PolicyTypeCustom       PolicyType = "Custom"
	PolicyTypeNotSpecified PolicyType = "NotSpecified"
	PolicyTypeStatic       PolicyType = "Static"
)

// PossiblePolicyTypeValues returns the possible values for the PolicyType const type.
func PossiblePolicyTypeValues() []PolicyType {
	return []PolicyType{
		PolicyTypeBuiltIn,
		PolicyTypeCustom,
		PolicyTypeNotSpecified,
		PolicyTypeStatic,
	}
}

// ResourceIdentityType - The identity type. This is the only required field when adding a system or user assigned identity
// to a resource.
type ResourceIdentityType string

const (
	// ResourceIdentityTypeNone - Indicates that no identity is associated with the resource or that the existing identity should
	// be removed.
	ResourceIdentityTypeNone ResourceIdentityType = "None"
	// ResourceIdentityTypeSystemAssigned - Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeUserAssigned - Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// SelectorKind - The selector kind.
type SelectorKind string

const (
	// SelectorKindPolicyDefinitionReferenceID - The selector kind to filter policies by the policy definition reference ID.
	SelectorKindPolicyDefinitionReferenceID SelectorKind = "policyDefinitionReferenceId"
	// SelectorKindResourceLocation - The selector kind to filter policies by the resource location.
	SelectorKindResourceLocation SelectorKind = "resourceLocation"
	// SelectorKindResourceType - The selector kind to filter policies by the resource type.
	SelectorKindResourceType SelectorKind = "resourceType"
	// SelectorKindResourceWithoutLocation - The selector kind to filter policies by the resource without location.
	SelectorKindResourceWithoutLocation SelectorKind = "resourceWithoutLocation"
)

// PossibleSelectorKindValues returns the possible values for the SelectorKind const type.
func PossibleSelectorKindValues() []SelectorKind {
	return []SelectorKind{
		SelectorKindPolicyDefinitionReferenceID,
		SelectorKindResourceLocation,
		SelectorKindResourceType,
		SelectorKindResourceWithoutLocation,
	}
}
