//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armchaos

const (
	moduleName    = "armchaos"
	moduleVersion = "v0.8.0"
)

// ActionType - Enum. Indicates the action type. "Internal" refers to actions that are for internal only APIs.
type ActionType string

const (
	ActionTypeInternal ActionType = "Internal"
)

// PossibleActionTypeValues returns the possible values for the ActionType const type.
func PossibleActionTypeValues() []ActionType {
	return []ActionType{
		ActionTypeInternal,
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

// FilterType - Enum that discriminates between filter types. Currently only Simple type is supported.
type FilterType string

const (
	FilterTypeSimple FilterType = "Simple"
)

// PossibleFilterTypeValues returns the possible values for the FilterType const type.
func PossibleFilterTypeValues() []FilterType {
	return []FilterType{
		FilterTypeSimple,
	}
}

// Origin - The intended executor of the operation; as in Resource Based Access Control (RBAC) and audit logs UX. Default
// value is "user,system"
type Origin string

const (
	OriginSystem     Origin = "system"
	OriginUser       Origin = "user"
	OriginUserSystem Origin = "user,system"
)

// PossibleOriginValues returns the possible values for the Origin const type.
func PossibleOriginValues() []Origin {
	return []Origin{
		OriginSystem,
		OriginUser,
		OriginUserSystem,
	}
}

// ResourceIdentityType - String of the resource identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeNone           ResourceIdentityType = "None"
	ResourceIdentityTypeSystemAssigned ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned   ResourceIdentityType = "UserAssigned"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeNone,
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
	}
}

// SelectorType - Enum of the selector type.
type SelectorType string

const (
	SelectorTypeList  SelectorType = "List"
	SelectorTypeQuery SelectorType = "Query"
)

// PossibleSelectorTypeValues returns the possible values for the SelectorType const type.
func PossibleSelectorTypeValues() []SelectorType {
	return []SelectorType{
		SelectorTypeList,
		SelectorTypeQuery,
	}
}

// TargetReferenceType - Enum of the Target reference type.
type TargetReferenceType string

const (
	TargetReferenceTypeChaosTarget TargetReferenceType = "ChaosTarget"
)

// PossibleTargetReferenceTypeValues returns the possible values for the TargetReferenceType const type.
func PossibleTargetReferenceTypeValues() []TargetReferenceType {
	return []TargetReferenceType{
		TargetReferenceTypeChaosTarget,
	}
}
