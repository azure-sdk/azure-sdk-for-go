//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

const (
	moduleName    = "armresourcegraph"
	moduleVersion = "v0.9.0"
)

// AuthorizationScopeFilter - Defines what level of authorization resources should be returned based on the which subscriptions
// and management groups are passed as scopes.
type AuthorizationScopeFilter string

const (
	AuthorizationScopeFilterAtScopeAboveAndBelow AuthorizationScopeFilter = "AtScopeAboveAndBelow"
	AuthorizationScopeFilterAtScopeAndAbove      AuthorizationScopeFilter = "AtScopeAndAbove"
	AuthorizationScopeFilterAtScopeAndBelow      AuthorizationScopeFilter = "AtScopeAndBelow"
	AuthorizationScopeFilterAtScopeExact         AuthorizationScopeFilter = "AtScopeExact"
)

// PossibleAuthorizationScopeFilterValues returns the possible values for the AuthorizationScopeFilter const type.
func PossibleAuthorizationScopeFilterValues() []AuthorizationScopeFilter {
	return []AuthorizationScopeFilter{
		AuthorizationScopeFilterAtScopeAboveAndBelow,
		AuthorizationScopeFilterAtScopeAndAbove,
		AuthorizationScopeFilterAtScopeAndBelow,
		AuthorizationScopeFilterAtScopeExact,
	}
}

// ChangeCategory - The change category.
type ChangeCategory string

const (
	ChangeCategorySystem ChangeCategory = "System"
	ChangeCategoryUser   ChangeCategory = "User"
)

// PossibleChangeCategoryValues returns the possible values for the ChangeCategory const type.
func PossibleChangeCategoryValues() []ChangeCategory {
	return []ChangeCategory{
		ChangeCategorySystem,
		ChangeCategoryUser,
	}
}

// ChangeType - The change type for snapshot. PropertyChanges will be provided in case of Update change type
type ChangeType string

const (
	ChangeTypeCreate ChangeType = "Create"
	ChangeTypeDelete ChangeType = "Delete"
	ChangeTypeUpdate ChangeType = "Update"
)

// PossibleChangeTypeValues returns the possible values for the ChangeType const type.
func PossibleChangeTypeValues() []ChangeType {
	return []ChangeType{
		ChangeTypeCreate,
		ChangeTypeDelete,
		ChangeTypeUpdate,
	}
}

// ColumnDataType - Data type of a column in a table.
type ColumnDataType string

const (
	ColumnDataTypeBoolean  ColumnDataType = "boolean"
	ColumnDataTypeDatetime ColumnDataType = "datetime"
	ColumnDataTypeInteger  ColumnDataType = "integer"
	ColumnDataTypeNumber   ColumnDataType = "number"
	ColumnDataTypeObject   ColumnDataType = "object"
	ColumnDataTypeString   ColumnDataType = "string"
)

// PossibleColumnDataTypeValues returns the possible values for the ColumnDataType const type.
func PossibleColumnDataTypeValues() []ColumnDataType {
	return []ColumnDataType{
		ColumnDataTypeBoolean,
		ColumnDataTypeDatetime,
		ColumnDataTypeInteger,
		ColumnDataTypeNumber,
		ColumnDataTypeObject,
		ColumnDataTypeString,
	}
}

// FacetSortOrder - The sorting order by the selected column (count by default).
type FacetSortOrder string

const (
	FacetSortOrderAsc  FacetSortOrder = "asc"
	FacetSortOrderDesc FacetSortOrder = "desc"
)

// PossibleFacetSortOrderValues returns the possible values for the FacetSortOrder const type.
func PossibleFacetSortOrderValues() []FacetSortOrder {
	return []FacetSortOrder{
		FacetSortOrderAsc,
		FacetSortOrderDesc,
	}
}

// PropertyChangeType - The property change Type
type PropertyChangeType string

const (
	PropertyChangeTypeInsert PropertyChangeType = "Insert"
	PropertyChangeTypeRemove PropertyChangeType = "Remove"
	PropertyChangeTypeUpdate PropertyChangeType = "Update"
)

// PossiblePropertyChangeTypeValues returns the possible values for the PropertyChangeType const type.
func PossiblePropertyChangeTypeValues() []PropertyChangeType {
	return []PropertyChangeType{
		PropertyChangeTypeInsert,
		PropertyChangeTypeRemove,
		PropertyChangeTypeUpdate,
	}
}

// ResultFormat - Defines in which format query result returned.
type ResultFormat string

const (
	ResultFormatObjectArray ResultFormat = "objectArray"
	ResultFormatTable       ResultFormat = "table"
)

// PossibleResultFormatValues returns the possible values for the ResultFormat const type.
func PossibleResultFormatValues() []ResultFormat {
	return []ResultFormat{
		ResultFormatObjectArray,
		ResultFormatTable,
	}
}

// ResultTruncated - Indicates whether the query results are truncated.
type ResultTruncated string

const (
	ResultTruncatedFalse ResultTruncated = "false"
	ResultTruncatedTrue  ResultTruncated = "true"
)

// PossibleResultTruncatedValues returns the possible values for the ResultTruncated const type.
func PossibleResultTruncatedValues() []ResultTruncated {
	return []ResultTruncated{
		ResultTruncatedFalse,
		ResultTruncatedTrue,
	}
}

// Role - The role which generates a specific message. Restricted to the user or assistant role.
type Role string

const (
	// RoleAssistant - If a message was generated by ARG Query Generation service, then the role should be assistant
	RoleAssistant Role = "assistant"
	// RoleSystem - If a message provides high level instructions for the query generation, then the role should be system
	RoleSystem Role = "system"
	// RoleUser - If a message was generated by a human, the role should be assigned user
	RoleUser Role = "user"
)

// PossibleRoleValues returns the possible values for the Role const type.
func PossibleRoleValues() []Role {
	return []Role{
		RoleAssistant,
		RoleSystem,
		RoleUser,
	}
}

// StatusCategory - Status Category
type StatusCategory string

const (
	// StatusCategoryFailed - Failed when query was not generated
	StatusCategoryFailed StatusCategory = "Failed"
	// StatusCategorySucceeded - Succeeded when query was generated successfully
	StatusCategorySucceeded StatusCategory = "Succeeded"
)

// PossibleStatusCategoryValues returns the possible values for the StatusCategory const type.
func PossibleStatusCategoryValues() []StatusCategory {
	return []StatusCategory{
		StatusCategoryFailed,
		StatusCategorySucceeded,
	}
}

// Versions - Versions Info.
type Versions string

const (
	// VersionsV20230901Preview - The 2023-09-01-preview version.
	VersionsV20230901Preview Versions = "2023-09-01-preview"
)

// PossibleVersionsValues returns the possible values for the Versions const type.
func PossibleVersionsValues() []Versions {
	return []Versions{
		VersionsV20230901Preview,
	}
}
