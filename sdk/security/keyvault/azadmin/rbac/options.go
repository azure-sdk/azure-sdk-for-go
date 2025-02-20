// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package rbac

// CreateOrUpdateRoleDefinitionOptions contains the optional parameters for the Client.CreateOrUpdateRoleDefinition method.
type CreateOrUpdateRoleDefinitionOptions struct {
	// placeholder for future optional parameters
}

// CreateRoleAssignmentOptions contains the optional parameters for the Client.CreateRoleAssignment method.
type CreateRoleAssignmentOptions struct {
	// placeholder for future optional parameters
}

// DeleteRoleAssignmentOptions contains the optional parameters for the Client.DeleteRoleAssignment method.
type DeleteRoleAssignmentOptions struct {
	// placeholder for future optional parameters
}

// DeleteRoleDefinitionOptions contains the optional parameters for the Client.DeleteRoleDefinition method.
type DeleteRoleDefinitionOptions struct {
	// placeholder for future optional parameters
}

// GetRoleAssignmentOptions contains the optional parameters for the Client.GetRoleAssignment method.
type GetRoleAssignmentOptions struct {
	// placeholder for future optional parameters
}

// GetRoleDefinitionOptions contains the optional parameters for the Client.GetRoleDefinition method.
type GetRoleDefinitionOptions struct {
	// placeholder for future optional parameters
}

// ListRoleAssignmentsOptions contains the optional parameters for the Client.NewListRoleAssignmentsPager method.
type ListRoleAssignmentsOptions struct {
	// The filter to apply on the operation. Use $filter=atScope() to return all role assignments at or above the scope. Use $filter=principalId
	// eq {id} to return all role assignments at, above or below the scope for the specified principal.
	Filter *string
}

// ListRoleDefinitionsOptions contains the optional parameters for the Client.NewListRoleDefinitionsPager method.
type ListRoleDefinitionsOptions struct {
	// The filter to apply on the operation. Use atScopeAndBelow filter to search below the given scope as well.
	Filter *string
}
