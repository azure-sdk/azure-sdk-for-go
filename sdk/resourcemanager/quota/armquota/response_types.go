//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

// ClientCreateOrUpdateResponse contains the response from method Client.BeginCreateOrUpdate.
type ClientCreateOrUpdateResponse struct {
	// Quota limit.
	CurrentQuotaLimitBase
}

// ClientGetResponse contains the response from method Client.Get.
type ClientGetResponse struct {
	// Quota limit.
	CurrentQuotaLimitBase

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// ClientListResponse contains the response from method Client.NewListPager.
type ClientListResponse struct {
	// Quota limits.
	Limits

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// ClientUpdateResponse contains the response from method Client.BeginUpdate.
type ClientUpdateResponse struct {
	// Quota limit.
	CurrentQuotaLimitBase
}

// GroupQuotaEnforcementClientCreateOrUpdateResponse contains the response from method GroupQuotaEnforcementClient.BeginCreateOrUpdate.
type GroupQuotaEnforcementClientCreateOrUpdateResponse struct {
	// The GroupQuota Enforcement status for a Azure Location/Region.
	GroupQuotasEnforcementResponse
}

// GroupQuotaEnforcementClientGetResponse contains the response from method GroupQuotaEnforcementClient.Get.
type GroupQuotaEnforcementClientGetResponse struct {
	// The GroupQuota Enforcement status for a Azure Location/Region.
	GroupQuotasEnforcementResponse
}

// GroupQuotaEnforcementClientListResponse contains the response from method GroupQuotaEnforcementClient.NewListPager.
type GroupQuotaEnforcementClientListResponse struct {
	// List of Azure regions, where the group quotas is enabled for enforcement.
	GroupQuotasEnforcementListResponse
}

// GroupQuotaEnforcementClientUpdateResponse contains the response from method GroupQuotaEnforcementClient.BeginUpdate.
type GroupQuotaEnforcementClientUpdateResponse struct {
	// The GroupQuota Enforcement status for a Azure Location/Region.
	GroupQuotasEnforcementResponse
}

// GroupQuotaLimitsClientCreateOrUpdateResponse contains the response from method GroupQuotaLimitsClient.BeginCreateOrUpdate.
type GroupQuotaLimitsClientCreateOrUpdateResponse struct {
	// Group Quota limit.
	GroupQuotaLimit
}

// GroupQuotaLimitsClientGetResponse contains the response from method GroupQuotaLimitsClient.Get.
type GroupQuotaLimitsClientGetResponse struct {
	// Group Quota limit.
	GroupQuotaLimit
}

// GroupQuotaLimitsClientListResponse contains the response from method GroupQuotaLimitsClient.NewListPager.
type GroupQuotaLimitsClientListResponse struct {
	// List of Group Quota Limit details.
	GroupQuotaLimitList
}

// GroupQuotaLimitsClientUpdateResponse contains the response from method GroupQuotaLimitsClient.BeginUpdate.
type GroupQuotaLimitsClientUpdateResponse struct {
	// Group Quota limit.
	GroupQuotaLimit
}

// GroupQuotaLimitsRequestsClientGetResponse contains the response from method GroupQuotaLimitsRequestsClient.Get.
type GroupQuotaLimitsRequestsClientGetResponse struct {
	// Status of a single GroupQuota request.
	SubmittedResourceRequestStatus
}

// GroupQuotaLimitsRequestsClientListResponse contains the response from method GroupQuotaLimitsRequestsClient.NewListPager.
type GroupQuotaLimitsRequestsClientListResponse struct {
	// Share Quota Entity list.
	SubmittedResourceRequestStatusList
}

// GroupQuotaSubscriptionQuotaAllocationClientCreateOrUpdateResponse contains the response from method GroupQuotaSubscriptionQuotaAllocationClient.BeginCreateOrUpdate.
type GroupQuotaSubscriptionQuotaAllocationClientCreateOrUpdateResponse struct {
	// Quota allocated to a subscription for the specific Resource Provider, Location, ResourceName. This will include the GroupQuota
	// and total quota allocated to the subscription. Only the Group quota allocated to the subscription can be allocated back
	// to the MG Group Quota.
	SubscriptionQuotaAllocations
}

// GroupQuotaSubscriptionQuotaAllocationClientGetResponse contains the response from method GroupQuotaSubscriptionQuotaAllocationClient.Get.
type GroupQuotaSubscriptionQuotaAllocationClientGetResponse struct {
	// Quota allocated to a subscription for the specific Resource Provider, Location, ResourceName. This will include the GroupQuota
	// and total quota allocated to the subscription. Only the Group quota allocated to the subscription can be allocated back
	// to the MG Group Quota.
	SubscriptionQuotaAllocations
}

// GroupQuotaSubscriptionQuotaAllocationClientListResponse contains the response from method GroupQuotaSubscriptionQuotaAllocationClient.NewListPager.
type GroupQuotaSubscriptionQuotaAllocationClientListResponse struct {
	// Subscription quota list.
	SubscriptionQuotaAllocationsList
}

// GroupQuotaSubscriptionQuotaAllocationClientUpdateResponse contains the response from method GroupQuotaSubscriptionQuotaAllocationClient.BeginUpdate.
type GroupQuotaSubscriptionQuotaAllocationClientUpdateResponse struct {
	// Quota allocated to a subscription for the specific Resource Provider, Location, ResourceName. This will include the GroupQuota
	// and total quota allocated to the subscription. Only the Group quota allocated to the subscription can be allocated back
	// to the MG Group Quota.
	SubscriptionQuotaAllocations
}

// GroupQuotaSubscriptionQuotaAllocationRequestsClientGetResponse contains the response from method GroupQuotaSubscriptionQuotaAllocationRequestsClient.Get.
type GroupQuotaSubscriptionQuotaAllocationRequestsClientGetResponse struct {
	// The new quota limit request status.
	AllocationRequestStatus
}

// GroupQuotaSubscriptionQuotaAllocationRequestsClientListResponse contains the response from method GroupQuotaSubscriptionQuotaAllocationRequestsClient.NewListPager.
type GroupQuotaSubscriptionQuotaAllocationRequestsClientListResponse struct {
	// List of QuotaAllocation Request Status
	AllocationRequestStatusList
}

// GroupQuotaSubscriptionsClientCreateOrUpdateResponse contains the response from method GroupQuotaSubscriptionsClient.BeginCreateOrUpdate.
type GroupQuotaSubscriptionsClientCreateOrUpdateResponse struct {
	// This represents a Azure subscriptionId that is associated with a GroupQuotasEntity.
	GroupQuotaSubscriptionID
}

// GroupQuotaSubscriptionsClientDeleteResponse contains the response from method GroupQuotaSubscriptionsClient.Delete.
type GroupQuotaSubscriptionsClientDeleteResponse struct {
	// placeholder for future response values
}

// GroupQuotaSubscriptionsClientGetResponse contains the response from method GroupQuotaSubscriptionsClient.Get.
type GroupQuotaSubscriptionsClientGetResponse struct {
	// This represents a Azure subscriptionId that is associated with a GroupQuotasEntity.
	GroupQuotaSubscriptionID
}

// GroupQuotaSubscriptionsClientListResponse contains the response from method GroupQuotaSubscriptionsClient.NewListPager.
type GroupQuotaSubscriptionsClientListResponse struct {
	// List of GroupQuotaSubscriptionIds
	GroupQuotaSubscriptionIDList
}

// GroupQuotaSubscriptionsClientUpdateResponse contains the response from method GroupQuotaSubscriptionsClient.BeginUpdate.
type GroupQuotaSubscriptionsClientUpdateResponse struct {
	// This represents a Azure subscriptionId that is associated with a GroupQuotasEntity.
	GroupQuotaSubscriptionID
}

// GroupQuotasClientCreateOrUpdateResponse contains the response from method GroupQuotasClient.BeginCreateOrUpdate.
type GroupQuotasClientCreateOrUpdateResponse struct {
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	GroupQuotasEntity
}

// GroupQuotasClientDeleteResponse contains the response from method GroupQuotasClient.Delete.
type GroupQuotasClientDeleteResponse struct {
	// placeholder for future response values
}

// GroupQuotasClientGetResponse contains the response from method GroupQuotasClient.Get.
type GroupQuotasClientGetResponse struct {
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	GroupQuotasEntity
}

// GroupQuotasClientListResponse contains the response from method GroupQuotasClient.NewListPager.
type GroupQuotasClientListResponse struct {
	// List of Group Quotas at MG level.
	GroupQuotaList
}

// GroupQuotasClientUpdateResponse contains the response from method GroupQuotasClient.BeginUpdate.
type GroupQuotasClientUpdateResponse struct {
	// Properties and filters for ShareQuota. The request parameter is optional, if there are no filters specified.
	GroupQuotasEntity
}

// OperationClientListResponse contains the response from method OperationClient.NewListPager.
type OperationClientListResponse struct {
	OperationList
}

// RequestStatusClientGetResponse contains the response from method RequestStatusClient.Get.
type RequestStatusClientGetResponse struct {
	// List of quota requests with details.
	RequestDetails
}

// RequestStatusClientListResponse contains the response from method RequestStatusClient.NewListPager.
type RequestStatusClientListResponse struct {
	// Quota request information.
	RequestDetailsList
}

// SubscriptionRequestsClientGetResponse contains the response from method SubscriptionRequestsClient.Get.
type SubscriptionRequestsClientGetResponse struct {
	// The new quota limit request status.
	GroupQuotaSubscriptionRequestStatus
}

// SubscriptionRequestsClientListResponse contains the response from method SubscriptionRequestsClient.NewListPager.
type SubscriptionRequestsClientListResponse struct {
	// List of GroupQuotaSubscriptionRequests Status
	GroupQuotaSubscriptionRequestStatusList
}

// UsagesClientGetResponse contains the response from method UsagesClient.Get.
type UsagesClientGetResponse struct {
	// Resource usage.
	CurrentUsagesBase

	// ETag contains the information returned from the ETag header response.
	ETag *string
}

// UsagesClientListResponse contains the response from method UsagesClient.NewListPager.
type UsagesClientListResponse struct {
	// Quota limits.
	UsagesLimits

	// ETag contains the information returned from the ETag header response.
	ETag *string
}
