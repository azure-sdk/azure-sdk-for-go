//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

// AvailabilityStatusesClientGetByResourceResponse contains the response from method AvailabilityStatusesClient.GetByResource.
type AvailabilityStatusesClientGetByResourceResponse struct {
	AvailabilityStatus
}

// AvailabilityStatusesClientListByResourceGroupResponse contains the response from method AvailabilityStatusesClient.ListByResourceGroup.
type AvailabilityStatusesClientListByResourceGroupResponse struct {
	AvailabilityStatusListResult
}

// AvailabilityStatusesClientListBySubscriptionIDResponse contains the response from method AvailabilityStatusesClient.ListBySubscriptionID.
type AvailabilityStatusesClientListBySubscriptionIDResponse struct {
	AvailabilityStatusListResult
}

// AvailabilityStatusesClientListResponse contains the response from method AvailabilityStatusesClient.List.
type AvailabilityStatusesClientListResponse struct {
	AvailabilityStatusListResult
}

// ChildAvailabilityStatusesClientGetByResourceResponse contains the response from method ChildAvailabilityStatusesClient.GetByResource.
type ChildAvailabilityStatusesClientGetByResourceResponse struct {
	AvailabilityStatus
}

// ChildAvailabilityStatusesClientListResponse contains the response from method ChildAvailabilityStatusesClient.List.
type ChildAvailabilityStatusesClientListResponse struct {
	AvailabilityStatusListResult
}

// ChildResourcesClientListResponse contains the response from method ChildResourcesClient.List.
type ChildResourcesClientListResponse struct {
	AvailabilityStatusListResult
}

// EmergingIssuesClientGetResponse contains the response from method EmergingIssuesClient.Get.
type EmergingIssuesClientGetResponse struct {
	EmergingIssuesGetResult
}

// EmergingIssuesClientListResponse contains the response from method EmergingIssuesClient.List.
type EmergingIssuesClientListResponse struct {
	EmergingIssueListResult
}

// EventsClientListBySingleResourceResponse contains the response from method EventsClient.ListBySingleResource.
type EventsClientListBySingleResourceResponse struct {
	Events
}

// EventsClientListBySubscriptionIDResponse contains the response from method EventsClient.ListBySubscriptionID.
type EventsClientListBySubscriptionIDResponse struct {
	Events
}

// ImpactedResourcesClientListBySubscriptionIDResponse contains the response from method ImpactedResourcesClient.ListBySubscriptionID.
type ImpactedResourcesClientListBySubscriptionIDResponse struct {
	ImpactedResourceListResult
}

// MetadataClientGetEntityResponse contains the response from method MetadataClient.GetEntity.
type MetadataClientGetEntityResponse struct {
	MetadataEntity
}

// MetadataClientListResponse contains the response from method MetadataClient.List.
type MetadataClientListResponse struct {
	MetadataEntityListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}
