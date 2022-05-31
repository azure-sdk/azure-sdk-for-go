//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpolicy

const (
	moduleName    = "armpolicy"
	moduleVersion = "v0.6.1"
)

// AliasPathAttributes - The attributes of the token that the alias path is referring to.
type AliasPathAttributes string

const (
	// AliasPathAttributesModifiable - The token that the alias path is referring to is modifiable by policies with 'modify' effect.
	AliasPathAttributesModifiable AliasPathAttributes = "Modifiable"
	// AliasPathAttributesNone - The token that the alias path is referring to has no attributes.
	AliasPathAttributesNone AliasPathAttributes = "None"
)

// PossibleAliasPathAttributesValues returns the possible values for the AliasPathAttributes const type.
func PossibleAliasPathAttributesValues() []AliasPathAttributes {
	return []AliasPathAttributes{
		AliasPathAttributesModifiable,
		AliasPathAttributesNone,
	}
}

// AliasPathTokenType - The type of the token that the alias path is referring to.
type AliasPathTokenType string

const (
	// AliasPathTokenTypeAny - The token type can be anything.
	AliasPathTokenTypeAny AliasPathTokenType = "Any"
	// AliasPathTokenTypeArray - The token type is array.
	AliasPathTokenTypeArray AliasPathTokenType = "Array"
	// AliasPathTokenTypeBoolean - The token type is boolean.
	AliasPathTokenTypeBoolean AliasPathTokenType = "Boolean"
	// AliasPathTokenTypeInteger - The token type is integer.
	AliasPathTokenTypeInteger AliasPathTokenType = "Integer"
	// AliasPathTokenTypeNotSpecified - The token type is not specified.
	AliasPathTokenTypeNotSpecified AliasPathTokenType = "NotSpecified"
	// AliasPathTokenTypeNumber - The token type is number.
	AliasPathTokenTypeNumber AliasPathTokenType = "Number"
	// AliasPathTokenTypeObject - The token type is object.
	AliasPathTokenTypeObject AliasPathTokenType = "Object"
	// AliasPathTokenTypeString - The token type is string.
	AliasPathTokenTypeString AliasPathTokenType = "String"
)

// PossibleAliasPathTokenTypeValues returns the possible values for the AliasPathTokenType const type.
func PossibleAliasPathTokenTypeValues() []AliasPathTokenType {
	return []AliasPathTokenType{
		AliasPathTokenTypeAny,
		AliasPathTokenTypeArray,
		AliasPathTokenTypeBoolean,
		AliasPathTokenTypeInteger,
		AliasPathTokenTypeNotSpecified,
		AliasPathTokenTypeNumber,
		AliasPathTokenTypeObject,
		AliasPathTokenTypeString,
	}
}

// AliasPatternType - The type of alias pattern
type AliasPatternType string

const (
	// AliasPatternTypeNotSpecified - NotSpecified is not allowed.
	AliasPatternTypeNotSpecified AliasPatternType = "NotSpecified"
	// AliasPatternTypeExtract - Extract is the only allowed value.
	AliasPatternTypeExtract AliasPatternType = "Extract"
)

// PossibleAliasPatternTypeValues returns the possible values for the AliasPatternType const type.
func PossibleAliasPatternTypeValues() []AliasPatternType {
	return []AliasPatternType{
		AliasPatternTypeNotSpecified,
		AliasPatternTypeExtract,
	}
}

// AliasType - The type of the alias.
type AliasType string

const (
	// AliasTypeNotSpecified - Alias type is unknown (same as not providing alias type).
	AliasTypeNotSpecified AliasType = "NotSpecified"
	// AliasTypePlainText - Alias value is not secret.
	AliasTypePlainText AliasType = "PlainText"
	// AliasTypeMask - Alias value is secret.
	AliasTypeMask AliasType = "Mask"
)

// PossibleAliasTypeValues returns the possible values for the AliasType const type.
func PossibleAliasTypeValues() []AliasType {
	return []AliasType{
		AliasTypeNotSpecified,
		AliasTypePlainText,
		AliasTypeMask,
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

// ExemptionCategory - The policy exemption category. Possible values are Waiver and Mitigated.
type ExemptionCategory string

const (
	// ExemptionCategoryMitigated - This category of exemptions usually means the mitigation actions have been applied to the
	// scope.
	ExemptionCategoryMitigated ExemptionCategory = "Mitigated"
	// ExemptionCategoryWaiver - This category of exemptions usually means the scope is not applicable for the policy.
	ExemptionCategoryWaiver ExemptionCategory = "Waiver"
)

// PossibleExemptionCategoryValues returns the possible values for the ExemptionCategory const type.
func PossibleExemptionCategoryValues() []ExemptionCategory {
	return []ExemptionCategory{
		ExemptionCategoryMitigated,
		ExemptionCategoryWaiver,
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
	// ResourceIdentityTypeSystemAssigned - Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	// ResourceIdentityTypeUserAssigned - Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeUserAssigned ResourceIdentityType = "UserAssigned"
	// ResourceIdentityTypeNone - Indicates that no identity is associated with the resource or that the existing identity should
	// be removed.
	ResourceIdentityTypeNone ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeNone,
	}
}
