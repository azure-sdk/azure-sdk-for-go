//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armauthorization

// ClassicAdministratorsClientListResponse contains the response from method ClassicAdministratorsClient.List.
type ClassicAdministratorsClientListResponse struct {
	ClassicAdministratorListResult
}

// DenyAssignmentsClientGetByIDResponse contains the response from method DenyAssignmentsClient.GetByID.
type DenyAssignmentsClientGetByIDResponse struct {
	DenyAssignment
}

// DenyAssignmentsClientGetResponse contains the response from method DenyAssignmentsClient.Get.
type DenyAssignmentsClientGetResponse struct {
	DenyAssignment
}

// DenyAssignmentsClientListForResourceGroupResponse contains the response from method DenyAssignmentsClient.ListForResourceGroup.
type DenyAssignmentsClientListForResourceGroupResponse struct {
	DenyAssignmentListResult
}

// DenyAssignmentsClientListForResourceResponse contains the response from method DenyAssignmentsClient.ListForResource.
type DenyAssignmentsClientListForResourceResponse struct {
	DenyAssignmentListResult
}

// DenyAssignmentsClientListForScopeResponse contains the response from method DenyAssignmentsClient.ListForScope.
type DenyAssignmentsClientListForScopeResponse struct {
	DenyAssignmentListResult
}

// DenyAssignmentsClientListResponse contains the response from method DenyAssignmentsClient.List.
type DenyAssignmentsClientListResponse struct {
	DenyAssignmentListResult
}

// EligibleChildResourcesClientGetResponse contains the response from method EligibleChildResourcesClient.Get.
type EligibleChildResourcesClientGetResponse struct {
	EligibleChildResourcesListResult
}

// GlobalAdministratorClientElevateAccessResponse contains the response from method GlobalAdministratorClient.ElevateAccess.
type GlobalAdministratorClientElevateAccessResponse struct {
	// placeholder for future response values
}

// PermissionsClientListForResourceGroupResponse contains the response from method PermissionsClient.ListForResourceGroup.
type PermissionsClientListForResourceGroupResponse struct {
	PermissionGetResult
}

// PermissionsClientListForResourceResponse contains the response from method PermissionsClient.ListForResource.
type PermissionsClientListForResourceResponse struct {
	PermissionGetResult
}

// ProviderOperationsMetadataClientGetResponse contains the response from method ProviderOperationsMetadataClient.Get.
type ProviderOperationsMetadataClientGetResponse struct {
	ProviderOperationsMetadata
}

// ProviderOperationsMetadataClientListResponse contains the response from method ProviderOperationsMetadataClient.List.
type ProviderOperationsMetadataClientListResponse struct {
	ProviderOperationsMetadataListResult
}

// RoleAssignmentScheduleInstancesClientGetResponse contains the response from method RoleAssignmentScheduleInstancesClient.Get.
type RoleAssignmentScheduleInstancesClientGetResponse struct {
	RoleAssignmentScheduleInstance
}

// RoleAssignmentScheduleInstancesClientListForScopeResponse contains the response from method RoleAssignmentScheduleInstancesClient.ListForScope.
type RoleAssignmentScheduleInstancesClientListForScopeResponse struct {
	RoleAssignmentScheduleInstanceListResult
}

// RoleAssignmentScheduleRequestsClientCancelResponse contains the response from method RoleAssignmentScheduleRequestsClient.Cancel.
type RoleAssignmentScheduleRequestsClientCancelResponse struct {
	// placeholder for future response values
}

// RoleAssignmentScheduleRequestsClientCreateResponse contains the response from method RoleAssignmentScheduleRequestsClient.Create.
type RoleAssignmentScheduleRequestsClientCreateResponse struct {
	RoleAssignmentScheduleRequest
}

// RoleAssignmentScheduleRequestsClientGetResponse contains the response from method RoleAssignmentScheduleRequestsClient.Get.
type RoleAssignmentScheduleRequestsClientGetResponse struct {
	RoleAssignmentScheduleRequest
}

// RoleAssignmentScheduleRequestsClientListForScopeResponse contains the response from method RoleAssignmentScheduleRequestsClient.ListForScope.
type RoleAssignmentScheduleRequestsClientListForScopeResponse struct {
	RoleAssignmentScheduleRequestListResult
}

// RoleAssignmentScheduleRequestsClientValidateResponse contains the response from method RoleAssignmentScheduleRequestsClient.Validate.
type RoleAssignmentScheduleRequestsClientValidateResponse struct {
	RoleAssignmentScheduleRequest
}

// RoleAssignmentSchedulesClientGetResponse contains the response from method RoleAssignmentSchedulesClient.Get.
type RoleAssignmentSchedulesClientGetResponse struct {
	RoleAssignmentSchedule
}

// RoleAssignmentSchedulesClientListForScopeResponse contains the response from method RoleAssignmentSchedulesClient.ListForScope.
type RoleAssignmentSchedulesClientListForScopeResponse struct {
	RoleAssignmentScheduleListResult
}

// RoleAssignmentsClientCreateByIDResponse contains the response from method RoleAssignmentsClient.CreateByID.
type RoleAssignmentsClientCreateByIDResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientCreateResponse contains the response from method RoleAssignmentsClient.Create.
type RoleAssignmentsClientCreateResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientDeleteByIDResponse contains the response from method RoleAssignmentsClient.DeleteByID.
type RoleAssignmentsClientDeleteByIDResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientDeleteResponse contains the response from method RoleAssignmentsClient.Delete.
type RoleAssignmentsClientDeleteResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientGetByIDResponse contains the response from method RoleAssignmentsClient.GetByID.
type RoleAssignmentsClientGetByIDResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientGetResponse contains the response from method RoleAssignmentsClient.Get.
type RoleAssignmentsClientGetResponse struct {
	RoleAssignment
}

// RoleAssignmentsClientListForResourceGroupResponse contains the response from method RoleAssignmentsClient.ListForResourceGroup.
type RoleAssignmentsClientListForResourceGroupResponse struct {
	RoleAssignmentListResult
}

// RoleAssignmentsClientListForResourceResponse contains the response from method RoleAssignmentsClient.ListForResource.
type RoleAssignmentsClientListForResourceResponse struct {
	RoleAssignmentListResult
}

// RoleAssignmentsClientListForScopeResponse contains the response from method RoleAssignmentsClient.ListForScope.
type RoleAssignmentsClientListForScopeResponse struct {
	RoleAssignmentListResult
}

// RoleAssignmentsClientListResponse contains the response from method RoleAssignmentsClient.List.
type RoleAssignmentsClientListResponse struct {
	RoleAssignmentListResult
}

// RoleDefinitionsClientCreateOrUpdateResponse contains the response from method RoleDefinitionsClient.CreateOrUpdate.
type RoleDefinitionsClientCreateOrUpdateResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientDeleteResponse contains the response from method RoleDefinitionsClient.Delete.
type RoleDefinitionsClientDeleteResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientGetByIDResponse contains the response from method RoleDefinitionsClient.GetByID.
type RoleDefinitionsClientGetByIDResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientGetResponse contains the response from method RoleDefinitionsClient.Get.
type RoleDefinitionsClientGetResponse struct {
	RoleDefinition
}

// RoleDefinitionsClientListResponse contains the response from method RoleDefinitionsClient.List.
type RoleDefinitionsClientListResponse struct {
	RoleDefinitionListResult
}

// RoleEligibilityScheduleInstancesClientGetResponse contains the response from method RoleEligibilityScheduleInstancesClient.Get.
type RoleEligibilityScheduleInstancesClientGetResponse struct {
	RoleEligibilityScheduleInstance
}

// RoleEligibilityScheduleInstancesClientListForScopeResponse contains the response from method RoleEligibilityScheduleInstancesClient.ListForScope.
type RoleEligibilityScheduleInstancesClientListForScopeResponse struct {
	RoleEligibilityScheduleInstanceListResult
}

// RoleEligibilityScheduleRequestsClientCancelResponse contains the response from method RoleEligibilityScheduleRequestsClient.Cancel.
type RoleEligibilityScheduleRequestsClientCancelResponse struct {
	// placeholder for future response values
}

// RoleEligibilityScheduleRequestsClientCreateResponse contains the response from method RoleEligibilityScheduleRequestsClient.Create.
type RoleEligibilityScheduleRequestsClientCreateResponse struct {
	RoleEligibilityScheduleRequest
}

// RoleEligibilityScheduleRequestsClientGetResponse contains the response from method RoleEligibilityScheduleRequestsClient.Get.
type RoleEligibilityScheduleRequestsClientGetResponse struct {
	RoleEligibilityScheduleRequest
}

// RoleEligibilityScheduleRequestsClientListForScopeResponse contains the response from method RoleEligibilityScheduleRequestsClient.ListForScope.
type RoleEligibilityScheduleRequestsClientListForScopeResponse struct {
	RoleEligibilityScheduleRequestListResult
}

// RoleEligibilityScheduleRequestsClientValidateResponse contains the response from method RoleEligibilityScheduleRequestsClient.Validate.
type RoleEligibilityScheduleRequestsClientValidateResponse struct {
	RoleEligibilityScheduleRequest
}

// RoleEligibilitySchedulesClientGetResponse contains the response from method RoleEligibilitySchedulesClient.Get.
type RoleEligibilitySchedulesClientGetResponse struct {
	RoleEligibilitySchedule
}

// RoleEligibilitySchedulesClientListForScopeResponse contains the response from method RoleEligibilitySchedulesClient.ListForScope.
type RoleEligibilitySchedulesClientListForScopeResponse struct {
	RoleEligibilityScheduleListResult
}

// RoleManagementPoliciesClientDeleteResponse contains the response from method RoleManagementPoliciesClient.Delete.
type RoleManagementPoliciesClientDeleteResponse struct {
	// placeholder for future response values
}

// RoleManagementPoliciesClientGetResponse contains the response from method RoleManagementPoliciesClient.Get.
type RoleManagementPoliciesClientGetResponse struct {
	RoleManagementPolicy
}

// RoleManagementPoliciesClientListForScopeResponse contains the response from method RoleManagementPoliciesClient.ListForScope.
type RoleManagementPoliciesClientListForScopeResponse struct {
	RoleManagementPolicyListResult
}

// RoleManagementPoliciesClientUpdateResponse contains the response from method RoleManagementPoliciesClient.Update.
type RoleManagementPoliciesClientUpdateResponse struct {
	RoleManagementPolicy
}

// RoleManagementPolicyAssignmentsClientCreateResponse contains the response from method RoleManagementPolicyAssignmentsClient.Create.
type RoleManagementPolicyAssignmentsClientCreateResponse struct {
	RoleManagementPolicyAssignment
}

// RoleManagementPolicyAssignmentsClientDeleteResponse contains the response from method RoleManagementPolicyAssignmentsClient.Delete.
type RoleManagementPolicyAssignmentsClientDeleteResponse struct {
	// placeholder for future response values
}

// RoleManagementPolicyAssignmentsClientGetResponse contains the response from method RoleManagementPolicyAssignmentsClient.Get.
type RoleManagementPolicyAssignmentsClientGetResponse struct {
	RoleManagementPolicyAssignment
}

// RoleManagementPolicyAssignmentsClientListForScopeResponse contains the response from method RoleManagementPolicyAssignmentsClient.ListForScope.
type RoleManagementPolicyAssignmentsClientListForScopeResponse struct {
	RoleManagementPolicyAssignmentListResult
}
