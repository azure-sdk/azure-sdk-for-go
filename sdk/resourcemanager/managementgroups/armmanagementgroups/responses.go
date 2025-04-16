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
	// The results of Azure-AsyncOperation.
	AzureAsyncOperationResults
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

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// Describes the result of the request to list management groups.
	ManagementGroupListResult
}

// ClientUpdateResponse contains the response from method Client.Update.
type ClientUpdateResponse struct {
	// The management group details.
	ManagementGroup
}

// EntitiesClientListResponse contains the response from method EntitiesClient.NewListPager.
type EntitiesClientListResponse struct {
	// Describes the result of the request to view entities.
	EntityListResult
}

// HierarchySettingsClientCreateOrUpdateResponse contains the response from method HierarchySettingsClient.CreateOrUpdate.
type HierarchySettingsClientCreateOrUpdateResponse struct {
	// Settings defined at the Management Group scope.
	HierarchySettings
}

// HierarchySettingsClientDeleteResponse contains the response from method HierarchySettingsClient.Delete.
type HierarchySettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// HierarchySettingsClientGetResponse contains the response from method HierarchySettingsClient.Get.
type HierarchySettingsClientGetResponse struct {
	// Settings defined at the Management Group scope.
	HierarchySettings
}

// HierarchySettingsClientListResponse contains the response from method HierarchySettingsClient.List.
type HierarchySettingsClientListResponse struct {
	// Lists all hierarchy settings.
	HierarchySettingsList
}

// HierarchySettingsClientUpdateResponse contains the response from method HierarchySettingsClient.Update.
type HierarchySettingsClientUpdateResponse struct {
	// Settings defined at the Management Group scope.
	HierarchySettings
}

// ManagementGroupSubscriptionsClientCreateResponse contains the response from method ManagementGroupSubscriptionsClient.Create.
type ManagementGroupSubscriptionsClientCreateResponse struct {
	// The details of subscription under management group.
	SubscriptionUnderManagementGroup
}

// ManagementGroupSubscriptionsClientDeleteResponse contains the response from method ManagementGroupSubscriptionsClient.Delete.
type ManagementGroupSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ManagementGroupSubscriptionsClientGetSubscriptionResponse contains the response from method ManagementGroupSubscriptionsClient.GetSubscription.
type ManagementGroupSubscriptionsClientGetSubscriptionResponse struct {
	// The details of subscription under management group.
	SubscriptionUnderManagementGroup
}

// ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse contains the response from method ManagementGroupSubscriptionsClient.NewGetSubscriptionsUnderManagementGroupPager.
type ManagementGroupSubscriptionsClientGetSubscriptionsUnderManagementGroupResponse struct {
	// The details of all subscriptions under management group.
	ListSubscriptionUnderManagementGroup
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Describes the result of the request to list Microsoft.Management operations.
	OperationListResult
}
