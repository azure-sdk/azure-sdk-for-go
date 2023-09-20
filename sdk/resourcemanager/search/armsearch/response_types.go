//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

// AdminKeysClientGetResponse contains the response from method AdminKeysClient.Get.
type AdminKeysClientGetResponse struct {
	// Response containing the primary and secondary admin API keys for a given Azure Cognitive Search service.
	AdminKeyResult
}

// AdminKeysClientRegenerateResponse contains the response from method AdminKeysClient.Regenerate.
type AdminKeysClientRegenerateResponse struct {
	// Response containing the primary and secondary admin API keys for a given Azure Cognitive Search service.
	AdminKeyResult
}

// ManagementClientUsageBySubscriptionSKUResponse contains the response from method ManagementClient.UsageBySubscriptionSKU.
type ManagementClientUsageBySubscriptionSKUResponse struct {
	// Describes the quota usage for a particular sku supported by Azure Cognitive Search.
	QuotaUsageResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// The result of the request to list REST API operations. It contains a list of operations and a URL to get the next set of
	// results.
	OperationListResult
}

// PrivateEndpointConnectionsClientDeleteResponse contains the response from method PrivateEndpointConnectionsClient.Delete.
type PrivateEndpointConnectionsClientDeleteResponse struct {
	// Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientGetResponse contains the response from method PrivateEndpointConnectionsClient.Get.
type PrivateEndpointConnectionsClientGetResponse struct {
	// Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
	PrivateEndpointConnection
}

// PrivateEndpointConnectionsClientListByServiceResponse contains the response from method PrivateEndpointConnectionsClient.NewListByServicePager.
type PrivateEndpointConnectionsClientListByServiceResponse struct {
	// Response containing a list of Private Endpoint connections.
	PrivateEndpointConnectionListResult
}

// PrivateEndpointConnectionsClientUpdateResponse contains the response from method PrivateEndpointConnectionsClient.Update.
type PrivateEndpointConnectionsClientUpdateResponse struct {
	// Describes an existing Private Endpoint connection to the Azure Cognitive Search service.
	PrivateEndpointConnection
}

// PrivateLinkResourcesClientListSupportedResponse contains the response from method PrivateLinkResourcesClient.NewListSupportedPager.
type PrivateLinkResourcesClientListSupportedResponse struct {
	// Response containing a list of supported Private Link Resources.
	PrivateLinkResourcesResult
}

// QueryKeysClientCreateResponse contains the response from method QueryKeysClient.Create.
type QueryKeysClientCreateResponse struct {
	// Describes an API key for a given Azure Cognitive Search service that has permissions for query operations only.
	QueryKey
}

// QueryKeysClientDeleteResponse contains the response from method QueryKeysClient.Delete.
type QueryKeysClientDeleteResponse struct {
	// placeholder for future response values
}

// QueryKeysClientListBySearchServiceResponse contains the response from method QueryKeysClient.NewListBySearchServicePager.
type QueryKeysClientListBySearchServiceResponse struct {
	// Response containing the query API keys for a given Azure Cognitive Search service.
	ListQueryKeysResult
}

// ServicesClientCheckNameAvailabilityResponse contains the response from method ServicesClient.CheckNameAvailability.
type ServicesClientCheckNameAvailabilityResponse struct {
	// Output of check name availability API.
	CheckNameAvailabilityOutput
}

// ServicesClientCreateOrUpdateResponse contains the response from method ServicesClient.BeginCreateOrUpdate.
type ServicesClientCreateOrUpdateResponse struct {
	// Describes an Azure Cognitive Search service and its current state.
	Service
}

// ServicesClientDeleteResponse contains the response from method ServicesClient.Delete.
type ServicesClientDeleteResponse struct {
	// placeholder for future response values
}

// ServicesClientGetResponse contains the response from method ServicesClient.Get.
type ServicesClientGetResponse struct {
	// Describes an Azure Cognitive Search service and its current state.
	Service
}

// ServicesClientListByResourceGroupResponse contains the response from method ServicesClient.NewListByResourceGroupPager.
type ServicesClientListByResourceGroupResponse struct {
	// Response containing a list of Azure Cognitive Search services.
	ServiceListResult
}

// ServicesClientListBySubscriptionResponse contains the response from method ServicesClient.NewListBySubscriptionPager.
type ServicesClientListBySubscriptionResponse struct {
	// Response containing a list of Azure Cognitive Search services.
	ServiceListResult
}

// ServicesClientUpdateResponse contains the response from method ServicesClient.Update.
type ServicesClientUpdateResponse struct {
	// Describes an Azure Cognitive Search service and its current state.
	Service
}

// SharedPrivateLinkResourcesClientCreateOrUpdateResponse contains the response from method SharedPrivateLinkResourcesClient.BeginCreateOrUpdate.
type SharedPrivateLinkResourcesClientCreateOrUpdateResponse struct {
	// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
	SharedPrivateLinkResource
}

// SharedPrivateLinkResourcesClientDeleteResponse contains the response from method SharedPrivateLinkResourcesClient.BeginDelete.
type SharedPrivateLinkResourcesClientDeleteResponse struct {
	// placeholder for future response values
}

// SharedPrivateLinkResourcesClientGetResponse contains the response from method SharedPrivateLinkResourcesClient.Get.
type SharedPrivateLinkResourcesClientGetResponse struct {
	// Describes a Shared Private Link Resource managed by the Azure Cognitive Search service.
	SharedPrivateLinkResource
}

// SharedPrivateLinkResourcesClientListByServiceResponse contains the response from method SharedPrivateLinkResourcesClient.NewListByServicePager.
type SharedPrivateLinkResourcesClientListByServiceResponse struct {
	// Response containing a list of Shared Private Link Resources.
	SharedPrivateLinkResourceListResult
}

// UsagesClientListBySubscriptionResponse contains the response from method UsagesClient.NewListBySubscriptionPager.
type UsagesClientListBySubscriptionResponse struct {
	// Response containing the quota usage information for all the supported skus of Azure Cognitive Search service.
	QuotaUsagesListResult
}
