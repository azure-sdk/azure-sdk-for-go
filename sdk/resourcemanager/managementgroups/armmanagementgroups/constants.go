//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmanagementgroups

const (
	moduleName    = "armmanagementgroups"
	moduleVersion = "v1.1.1"
)

type EntitySearchType string

const (
	EntitySearchTypeAllowedChildren             EntitySearchType = "AllowedChildren"
	EntitySearchTypeAllowedParents              EntitySearchType = "AllowedParents"
	EntitySearchTypeChildrenOnly                EntitySearchType = "ChildrenOnly"
	EntitySearchTypeParentAndFirstLevelChildren EntitySearchType = "ParentAndFirstLevelChildren"
	EntitySearchTypeParentOnly                  EntitySearchType = "ParentOnly"
)

// PossibleEntitySearchTypeValues returns the possible values for the EntitySearchType const type.
func PossibleEntitySearchTypeValues() []EntitySearchType {
	return []EntitySearchType{
		EntitySearchTypeAllowedChildren,
		EntitySearchTypeAllowedParents,
		EntitySearchTypeChildrenOnly,
		EntitySearchTypeParentAndFirstLevelChildren,
		EntitySearchTypeParentOnly,
	}
}

type EntityViewParameterType string

const (
	EntityViewParameterTypeAudit             EntityViewParameterType = "Audit"
	EntityViewParameterTypeFullHierarchy     EntityViewParameterType = "FullHierarchy"
	EntityViewParameterTypeGroupsOnly        EntityViewParameterType = "GroupsOnly"
	EntityViewParameterTypeSubscriptionsOnly EntityViewParameterType = "SubscriptionsOnly"
)

// PossibleEntityViewParameterTypeValues returns the possible values for the EntityViewParameterType const type.
func PossibleEntityViewParameterTypeValues() []EntityViewParameterType {
	return []EntityViewParameterType{
		EntityViewParameterTypeAudit,
		EntityViewParameterTypeFullHierarchy,
		EntityViewParameterTypeGroupsOnly,
		EntityViewParameterTypeSubscriptionsOnly,
	}
}

// ManagementGroupChildType - The type of child resource.
type ManagementGroupChildType string

const (
	ManagementGroupChildTypeMicrosoftManagementManagementGroups ManagementGroupChildType = "Microsoft.Management/managementGroups"
	ManagementGroupChildTypeSubscriptions                       ManagementGroupChildType = "/subscriptions"
)

// PossibleManagementGroupChildTypeValues returns the possible values for the ManagementGroupChildType const type.
func PossibleManagementGroupChildTypeValues() []ManagementGroupChildType {
	return []ManagementGroupChildType{
		ManagementGroupChildTypeMicrosoftManagementManagementGroups,
		ManagementGroupChildTypeSubscriptions,
	}
}

type ManagementGroupExpandType string

const (
	ManagementGroupExpandTypeAncestors ManagementGroupExpandType = "ancestors"
	ManagementGroupExpandTypeChildren  ManagementGroupExpandType = "children"
	ManagementGroupExpandTypePath      ManagementGroupExpandType = "path"
)

// PossibleManagementGroupExpandTypeValues returns the possible values for the ManagementGroupExpandType const type.
func PossibleManagementGroupExpandTypeValues() []ManagementGroupExpandType {
	return []ManagementGroupExpandType{
		ManagementGroupExpandTypeAncestors,
		ManagementGroupExpandTypeChildren,
		ManagementGroupExpandTypePath,
	}
}

// Permissions - The users specific permissions to this item.
type Permissions string

const (
	PermissionsNoaccess Permissions = "noaccess"
	PermissionsView     Permissions = "view"
	PermissionsEdit     Permissions = "edit"
	PermissionsDelete   Permissions = "delete"
)

// PossiblePermissionsValues returns the possible values for the Permissions const type.
func PossiblePermissionsValues() []Permissions {
	return []Permissions{
		PermissionsNoaccess,
		PermissionsView,
		PermissionsEdit,
		PermissionsDelete,
	}
}

// Reason - Required if nameAvailable == false. Invalid indicates the name provided does not match the resource provider's
// naming requirements (incorrect length, unsupported characters, etc.) AlreadyExists
// indicates that the name is already in use and is therefore unavailable.
type Reason string

const (
	ReasonInvalid       Reason = "Invalid"
	ReasonAlreadyExists Reason = "AlreadyExists"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonInvalid,
		ReasonAlreadyExists,
	}
}

// Status - The status of the Tenant Backfill
type Status string

const (
	StatusNotStarted               Status = "NotStarted"
	StatusNotStartedButGroupsExist Status = "NotStartedButGroupsExist"
	StatusStarted                  Status = "Started"
	StatusFailed                   Status = "Failed"
	StatusCancelled                Status = "Cancelled"
	StatusCompleted                Status = "Completed"
)

// PossibleStatusValues returns the possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{
		StatusNotStarted,
		StatusNotStartedButGroupsExist,
		StatusStarted,
		StatusFailed,
		StatusCancelled,
		StatusCompleted,
	}
}
