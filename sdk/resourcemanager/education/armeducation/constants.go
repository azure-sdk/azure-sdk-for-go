//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armeducation

const (
	moduleName    = "armeducation"
	moduleVersion = "v0.2.1"
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

// GrantStatus - Grant status
type GrantStatus string

const (
	GrantStatusActive   GrantStatus = "Active"
	GrantStatusInactive GrantStatus = "Inactive"
)

// PossibleGrantStatusValues returns the possible values for the GrantStatus const type.
func PossibleGrantStatusValues() []GrantStatus {
	return []GrantStatus{
		GrantStatusActive,
		GrantStatusInactive,
	}
}

// GrantType - Grant Offer Type
type GrantType string

const (
	GrantTypeAcademic GrantType = "Academic"
	GrantTypeStudent  GrantType = "Student"
)

// PossibleGrantTypeValues returns the possible values for the GrantType const type.
func PossibleGrantTypeValues() []GrantType {
	return []GrantType{
		GrantTypeAcademic,
		GrantTypeStudent,
	}
}

// JoinRequestStatus - Join request status
type JoinRequestStatus string

const (
	JoinRequestStatusDenied  JoinRequestStatus = "Denied"
	JoinRequestStatusPending JoinRequestStatus = "Pending"
)

// PossibleJoinRequestStatusValues returns the possible values for the JoinRequestStatus const type.
func PossibleJoinRequestStatusValues() []JoinRequestStatus {
	return []JoinRequestStatus{
		JoinRequestStatusDenied,
		JoinRequestStatusPending,
	}
}

// LabStatus - The status of this lab
type LabStatus string

const (
	LabStatusActive  LabStatus = "Active"
	LabStatusDeleted LabStatus = "Deleted"
	LabStatusPending LabStatus = "Pending"
)

// PossibleLabStatusValues returns the possible values for the LabStatus const type.
func PossibleLabStatusValues() []LabStatus {
	return []LabStatus{
		LabStatusActive,
		LabStatusDeleted,
		LabStatusPending,
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

// StudentLabStatus - Student Lab Status
type StudentLabStatus string

const (
	StudentLabStatusActive   StudentLabStatus = "Active"
	StudentLabStatusDeleted  StudentLabStatus = "Deleted"
	StudentLabStatusDisabled StudentLabStatus = "Disabled"
	StudentLabStatusExpired  StudentLabStatus = "Expired"
	StudentLabStatusPending  StudentLabStatus = "Pending"
)

// PossibleStudentLabStatusValues returns the possible values for the StudentLabStatus const type.
func PossibleStudentLabStatusValues() []StudentLabStatus {
	return []StudentLabStatus{
		StudentLabStatusActive,
		StudentLabStatusDeleted,
		StudentLabStatusDisabled,
		StudentLabStatusExpired,
		StudentLabStatusPending,
	}
}

// StudentRole - Student Role
type StudentRole string

const (
	StudentRoleAdmin   StudentRole = "Admin"
	StudentRoleStudent StudentRole = "Student"
)

// PossibleStudentRoleValues returns the possible values for the StudentRole const type.
func PossibleStudentRoleValues() []StudentRole {
	return []StudentRole{
		StudentRoleAdmin,
		StudentRoleStudent,
	}
}
