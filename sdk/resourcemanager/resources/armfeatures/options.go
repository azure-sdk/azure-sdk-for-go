//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armfeatures

// PolicyAssignmentsClientCreateByIDOptions contains the optional parameters for the PolicyAssignmentsClient.CreateByID method.
type PolicyAssignmentsClientCreateByIDOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientCreateOptions contains the optional parameters for the PolicyAssignmentsClient.Create method.
type PolicyAssignmentsClientCreateOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientDeleteByIDOptions contains the optional parameters for the PolicyAssignmentsClient.DeleteByID method.
type PolicyAssignmentsClientDeleteByIDOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientDeleteOptions contains the optional parameters for the PolicyAssignmentsClient.Delete method.
type PolicyAssignmentsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientGetByIDOptions contains the optional parameters for the PolicyAssignmentsClient.GetByID method.
type PolicyAssignmentsClientGetByIDOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientGetOptions contains the optional parameters for the PolicyAssignmentsClient.Get method.
type PolicyAssignmentsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientListForManagementGroupOptions contains the optional parameters for the PolicyAssignmentsClient.NewListForManagementGroupPager
// method.
type PolicyAssignmentsClientListForManagementGroupOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atScope()', 'atExactScope()' or 'policyDefinitionId
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atScope() is provided, the returned list only includes all policy assignments that apply to the scope, which is
	// everything in the unfiltered list except those applied to sub scopes contained
	// within the given scope. If $filter=atExactScope() is provided, the returned list only includes all policy assignments that
	// at the given scope. If $filter=policyDefinitionId eq '{value}' is provided,
	// the returned list includes all policy assignments of the policy definition whose id is {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyAssignmentsClientListForResourceGroupOptions contains the optional parameters for the PolicyAssignmentsClient.NewListForResourceGroupPager
// method.
type PolicyAssignmentsClientListForResourceGroupOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atScope()', 'atExactScope()' or 'policyDefinitionId
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atScope() is provided, the returned list only includes all policy assignments that apply to the scope, which is
	// everything in the unfiltered list except those applied to sub scopes contained
	// within the given scope. If $filter=atExactScope() is provided, the returned list only includes all policy assignments that
	// at the given scope. If $filter=policyDefinitionId eq '{value}' is provided,
	// the returned list includes all policy assignments of the policy definition whose id is {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyAssignmentsClientListForResourceOptions contains the optional parameters for the PolicyAssignmentsClient.NewListForResourcePager
// method.
type PolicyAssignmentsClientListForResourceOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atScope()', 'atExactScope()' or 'policyDefinitionId
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atScope() is provided, the returned list only includes all policy assignments that apply to the scope, which is
	// everything in the unfiltered list except those applied to sub scopes contained
	// within the given scope. If $filter=atExactScope() is provided, the returned list only includes all policy assignments that
	// at the given scope. If $filter=policyDefinitionId eq '{value}' is provided,
	// the returned list includes all policy assignments of the policy definition whose id is {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyAssignmentsClientListOptions contains the optional parameters for the PolicyAssignmentsClient.NewListPager method.
type PolicyAssignmentsClientListOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atScope()', 'atExactScope()' or 'policyDefinitionId
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atScope() is provided, the returned list only includes all policy assignments that apply to the scope, which is
	// everything in the unfiltered list except those applied to sub scopes contained
	// within the given scope. If $filter=atExactScope() is provided, the returned list only includes all policy assignments that
	// at the given scope. If $filter=policyDefinitionId eq '{value}' is provided,
	// the returned list includes all policy assignments of the policy definition whose id is {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyAssignmentsClientUpdateByIDOptions contains the optional parameters for the PolicyAssignmentsClient.UpdateByID method.
type PolicyAssignmentsClientUpdateByIDOptions struct {
	// placeholder for future optional parameters
}

// PolicyAssignmentsClientUpdateOptions contains the optional parameters for the PolicyAssignmentsClient.Update method.
type PolicyAssignmentsClientUpdateOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionVersionsClient.CreateOrUpdateAtManagementGroup
// method.
type PolicyDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientCreateOrUpdateOptions contains the optional parameters for the PolicyDefinitionVersionsClient.CreateOrUpdate
// method.
type PolicyDefinitionVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientDeleteAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionVersionsClient.DeleteAtManagementGroup
// method.
type PolicyDefinitionVersionsClientDeleteAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientDeleteOptions contains the optional parameters for the PolicyDefinitionVersionsClient.Delete
// method.
type PolicyDefinitionVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientGetAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionVersionsClient.GetAtManagementGroup
// method.
type PolicyDefinitionVersionsClientGetAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientGetBuiltInOptions contains the optional parameters for the PolicyDefinitionVersionsClient.GetBuiltIn
// method.
type PolicyDefinitionVersionsClientGetBuiltInOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientGetOptions contains the optional parameters for the PolicyDefinitionVersionsClient.Get method.
type PolicyDefinitionVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientListAllAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionVersionsClient.ListAllAtManagementGroup
// method.
type PolicyDefinitionVersionsClientListAllAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientListAllBuiltinsOptions contains the optional parameters for the PolicyDefinitionVersionsClient.ListAllBuiltins
// method.
type PolicyDefinitionVersionsClientListAllBuiltinsOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientListAllOptions contains the optional parameters for the PolicyDefinitionVersionsClient.ListAll
// method.
type PolicyDefinitionVersionsClientListAllOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionVersionsClientListBuiltInOptions contains the optional parameters for the PolicyDefinitionVersionsClient.NewListBuiltInPager
// method.
type PolicyDefinitionVersionsClientListBuiltInOptions struct {
	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyDefinitionVersionsClientListByManagementGroupOptions contains the optional parameters for the PolicyDefinitionVersionsClient.NewListByManagementGroupPager
// method.
type PolicyDefinitionVersionsClientListByManagementGroupOptions struct {
	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyDefinitionVersionsClientListOptions contains the optional parameters for the PolicyDefinitionVersionsClient.NewListPager
// method.
type PolicyDefinitionVersionsClientListOptions struct {
	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyDefinitionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionsClient.CreateOrUpdateAtManagementGroup
// method.
type PolicyDefinitionsClientCreateOrUpdateAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the PolicyDefinitionsClient.CreateOrUpdate
// method.
type PolicyDefinitionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientDeleteAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionsClient.DeleteAtManagementGroup
// method.
type PolicyDefinitionsClientDeleteAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientDeleteOptions contains the optional parameters for the PolicyDefinitionsClient.Delete method.
type PolicyDefinitionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientGetAtManagementGroupOptions contains the optional parameters for the PolicyDefinitionsClient.GetAtManagementGroup
// method.
type PolicyDefinitionsClientGetAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientGetBuiltInOptions contains the optional parameters for the PolicyDefinitionsClient.GetBuiltIn method.
type PolicyDefinitionsClientGetBuiltInOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientGetOptions contains the optional parameters for the PolicyDefinitionsClient.Get method.
type PolicyDefinitionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PolicyDefinitionsClientListBuiltInOptions contains the optional parameters for the PolicyDefinitionsClient.NewListBuiltInPager
// method.
type PolicyDefinitionsClientListBuiltInOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or 'category
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atExactScope() is provided, the returned list only includes all policy definitions that at the given scope. If
	// $filter='policyType -eq {value}' is provided, the returned list only includes all
	// policy definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom, and Static.
	// If $filter='category -eq {value}' is provided, the returned list only
	// includes all policy definitions whose category match the {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyDefinitionsClientListByManagementGroupOptions contains the optional parameters for the PolicyDefinitionsClient.NewListByManagementGroupPager
// method.
type PolicyDefinitionsClientListByManagementGroupOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or 'category
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atExactScope() is provided, the returned list only includes all policy definitions that at the given scope. If
	// $filter='policyType -eq {value}' is provided, the returned list only includes all
	// policy definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom, and Static.
	// If $filter='category -eq {value}' is provided, the returned list only
	// includes all policy definitions whose category match the {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicyDefinitionsClientListOptions contains the optional parameters for the PolicyDefinitionsClient.NewListPager method.
type PolicyDefinitionsClientListOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or 'category
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atExactScope() is provided, the returned list only includes all policy definitions that at the given scope. If
	// $filter='policyType -eq {value}' is provided, the returned list only includes all
	// policy definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom, and Static.
	// If $filter='category -eq {value}' is provided, the returned list only
	// includes all policy definitions whose category match the {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicySetDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.CreateOrUpdateAtManagementGroup
// method.
type PolicySetDefinitionVersionsClientCreateOrUpdateAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientCreateOrUpdateOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.CreateOrUpdate
// method.
type PolicySetDefinitionVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientDeleteAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.DeleteAtManagementGroup
// method.
type PolicySetDefinitionVersionsClientDeleteAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientDeleteOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.Delete
// method.
type PolicySetDefinitionVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientGetAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.GetAtManagementGroup
// method.
type PolicySetDefinitionVersionsClientGetAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientGetBuiltInOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.GetBuiltIn
// method.
type PolicySetDefinitionVersionsClientGetBuiltInOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientGetOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.Get
// method.
type PolicySetDefinitionVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientListAllAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.ListAllAtManagementGroup
// method.
type PolicySetDefinitionVersionsClientListAllAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientListAllBuiltinsOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.ListAllBuiltins
// method.
type PolicySetDefinitionVersionsClientListAllBuiltinsOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientListAllOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.ListAll
// method.
type PolicySetDefinitionVersionsClientListAllOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionVersionsClientListBuiltInOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.NewListBuiltInPager
// method.
type PolicySetDefinitionVersionsClientListBuiltInOptions struct {
	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicySetDefinitionVersionsClientListByManagementGroupOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.NewListByManagementGroupPager
// method.
type PolicySetDefinitionVersionsClientListByManagementGroupOptions struct {
	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicySetDefinitionVersionsClientListOptions contains the optional parameters for the PolicySetDefinitionVersionsClient.NewListPager
// method.
type PolicySetDefinitionVersionsClientListOptions struct {
	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.CreateOrUpdateAtManagementGroup
// method.
type PolicySetDefinitionsClientCreateOrUpdateAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientCreateOrUpdateOptions contains the optional parameters for the PolicySetDefinitionsClient.CreateOrUpdate
// method.
type PolicySetDefinitionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientDeleteAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.DeleteAtManagementGroup
// method.
type PolicySetDefinitionsClientDeleteAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientDeleteOptions contains the optional parameters for the PolicySetDefinitionsClient.Delete method.
type PolicySetDefinitionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientGetAtManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.GetAtManagementGroup
// method.
type PolicySetDefinitionsClientGetAtManagementGroupOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientGetBuiltInOptions contains the optional parameters for the PolicySetDefinitionsClient.GetBuiltIn
// method.
type PolicySetDefinitionsClientGetBuiltInOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientGetOptions contains the optional parameters for the PolicySetDefinitionsClient.Get method.
type PolicySetDefinitionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// PolicySetDefinitionsClientListBuiltInOptions contains the optional parameters for the PolicySetDefinitionsClient.NewListBuiltInPager
// method.
type PolicySetDefinitionsClientListBuiltInOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or 'category
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atExactScope() is provided, the returned list only includes all policy set definitions that at the given scope.
	// If $filter='policyType -eq {value}' is provided, the returned list only includes
	// all policy set definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom,
	// and Static. If $filter='category -eq {value}' is provided, the returned list only
	// includes all policy set definitions whose category match the {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicySetDefinitionsClientListByManagementGroupOptions contains the optional parameters for the PolicySetDefinitionsClient.NewListByManagementGroupPager
// method.
type PolicySetDefinitionsClientListByManagementGroupOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or 'category
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atExactScope() is provided, the returned list only includes all policy set definitions that at the given scope.
	// If $filter='policyType -eq {value}' is provided, the returned list only includes
	// all policy set definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom,
	// and Static. If $filter='category -eq {value}' is provided, the returned list only
	// includes all policy set definitions whose category match the {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}

// PolicySetDefinitionsClientListOptions contains the optional parameters for the PolicySetDefinitionsClient.NewListPager
// method.
type PolicySetDefinitionsClientListOptions struct {
	// The filter to apply on the operation. Valid values for $filter are: 'atExactScope()', 'policyType -eq {value}' or 'category
	// eq '{value}''. If $filter is not provided, no filtering is performed. If
	// $filter=atExactScope() is provided, the returned list only includes all policy set definitions that at the given scope.
	// If $filter='policyType -eq {value}' is provided, the returned list only includes
	// all policy set definitions whose type match the {value}. Possible policyType values are NotSpecified, BuiltIn, Custom,
	// and Static. If $filter='category -eq {value}' is provided, the returned list only
	// includes all policy set definitions whose category match the {value}.
	Filter *string

	// Maximum number of records to return. When the $top filter is not provided, it will return 500 records.
	Top *int32
}
