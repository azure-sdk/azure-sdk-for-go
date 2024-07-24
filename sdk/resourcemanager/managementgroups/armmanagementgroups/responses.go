//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagementgroups

// APIClientCheckNameAvailabilityResponse contains the response from method APIClient.CheckNameAvailability.
type APIClientCheckNameAvailabilityResponse struct {
	// Describes the result of the request to check management group name availability.
	CheckNameAvailabilityResult
}

// APIClientStartTenantBackfillResponse contains the response from method APIClient.StartTenantBackfill.
type APIClientStartTenantBackfillResponse struct {
	// The tenant backfill status
	TenantBackfillStatusResult
}

// APIClientTenantBackfillStatusResponse contains the response from method APIClient.TenantBackfillStatus.
type APIClientTenantBackfillStatusResponse struct {
	// The tenant backfill status
	TenantBackfillStatusResult
}

// ClientCreateOrUpdateResponse contains the response from method Client.BeginCreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	// The management group details.
	ManagementGroup
}

// ClientDeleteResponse contains the response from method Client.BeginDelete.
type ClientDeleteResponse struct {
	// The management group operation acceptance details.
	ManagementGroupOperationAcceptance
}

// ClientGetDescendantsResponse contains the response from method Client.NewGetDescendantsPager.
type ClientGetDescendantsResponse struct {
	// Describes the result of the request to view descendants.
	DescendantListResult
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// The management group details.
	ManagementGroup
}

// ClientListSettingsResponse contains the response from method Client.NewListSettingsPager.
type ClientListSettingsResponse struct {
	// The response of a HierarchySettings list operation.
	HierarchySettingsListResult
}

// ClientUpdateResponse contains the response from method Client.Update.
type ClientUpdateResponse struct {
	// The management group details.
	ManagementGroup
}

// EntitiesOperationsClientListResponse contains the response from method EntitiesOperationsClient.NewListPager.
type EntitiesOperationsClientListResponse struct {
	// Describes the result of the request to view entities.
	EntityListResult
}

// HierarchySettingsOperationGroupClientCreateOrUpdateResponse contains the response from method HierarchySettingsOperationGroupClient.CreateOrUpdate.
type HierarchySettingsOperationGroupClientCreateOrUpdateResponse struct {
	// Settings defined at the Management Group scope.
	HierarchySettings
}

// HierarchySettingsOperationGroupClientDeleteResponse contains the response from method HierarchySettingsOperationGroupClient.Delete.
type HierarchySettingsOperationGroupClientDeleteResponse struct {
	// placeholder for future response values
}

// HierarchySettingsOperationGroupClientGetResponse contains the response from method HierarchySettingsOperationGroupClient.Get.
type HierarchySettingsOperationGroupClientGetResponse struct {
	// Settings defined at the Management Group scope.
	HierarchySettings
}

// HierarchySettingsOperationGroupClientListResponse contains the response from method HierarchySettingsOperationGroupClient.NewListPager.
type HierarchySettingsOperationGroupClientListResponse struct {
	// The response of a HierarchySettings list operation.
	HierarchySettingsListResult
}

// HierarchySettingsOperationGroupClientUpdateResponse contains the response from method HierarchySettingsOperationGroupClient.Update.
type HierarchySettingsOperationGroupClientUpdateResponse struct {
	// Settings defined at the Management Group scope.
	HierarchySettings
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// SubscriptionUnderManagementGroupsClientCreateResponse contains the response from method SubscriptionUnderManagementGroupsClient.Create.
type SubscriptionUnderManagementGroupsClientCreateResponse struct {
	// The details of subscription under management group.
	SubscriptionUnderManagementGroup
}

// SubscriptionUnderManagementGroupsClientDeleteResponse contains the response from method SubscriptionUnderManagementGroupsClient.Delete.
type SubscriptionUnderManagementGroupsClientDeleteResponse struct {
	// placeholder for future response values
}

// SubscriptionUnderManagementGroupsClientGetSubscriptionResponse contains the response from method SubscriptionUnderManagementGroupsClient.GetSubscription.
type SubscriptionUnderManagementGroupsClientGetSubscriptionResponse struct {
	// The details of subscription under management group.
	SubscriptionUnderManagementGroup
}

// SubscriptionUnderManagementGroupsClientGetSubscriptionsUnderManagementGroupResponse contains the response from method SubscriptionUnderManagementGroupsClient.NewGetSubscriptionsUnderManagementGroupPager.
type SubscriptionUnderManagementGroupsClientGetSubscriptionsUnderManagementGroupResponse struct {
	// The response of a SubscriptionUnderManagementGroup list operation.
	SubscriptionUnderManagementGroupListResult
}
