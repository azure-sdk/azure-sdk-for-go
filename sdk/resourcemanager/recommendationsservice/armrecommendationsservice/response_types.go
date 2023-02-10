//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecommendationsservice

// AccountsClientCheckNameAvailabilityResponse contains the response from method AccountsClient.CheckNameAvailability.
type AccountsClientCheckNameAvailabilityResponse struct {
	CheckNameAvailabilityResponse
}

// AccountsClientCreateOrUpdateResponse contains the response from method AccountsClient.BeginCreateOrUpdate.
type AccountsClientCreateOrUpdateResponse struct {
	AccountResource
}

// AccountsClientDeleteResponse contains the response from method AccountsClient.BeginDelete.
type AccountsClientDeleteResponse struct {
	// placeholder for future response values
}

// AccountsClientGetResponse contains the response from method AccountsClient.Get.
type AccountsClientGetResponse struct {
	AccountResource
}

// AccountsClientGetStatusResponse contains the response from method AccountsClient.GetStatus.
type AccountsClientGetStatusResponse struct {
	AccountStatus
}

// AccountsClientListByResourceGroupResponse contains the response from method AccountsClient.NewListByResourceGroupPager.
type AccountsClientListByResourceGroupResponse struct {
	AccountResourceList
}

// AccountsClientListBySubscriptionResponse contains the response from method AccountsClient.NewListBySubscriptionPager.
type AccountsClientListBySubscriptionResponse struct {
	AccountResourceList
}

// AccountsClientUpdateResponse contains the response from method AccountsClient.BeginUpdate.
type AccountsClientUpdateResponse struct {
	AccountResource
}

// ModelingClientCreateOrUpdateResponse contains the response from method ModelingClient.BeginCreateOrUpdate.
type ModelingClientCreateOrUpdateResponse struct {
	ModelingResource
}

// ModelingClientDeleteResponse contains the response from method ModelingClient.BeginDelete.
type ModelingClientDeleteResponse struct {
	// placeholder for future response values
}

// ModelingClientGetResponse contains the response from method ModelingClient.Get.
type ModelingClientGetResponse struct {
	ModelingResource
}

// ModelingClientListByAccountResourceResponse contains the response from method ModelingClient.NewListByAccountResourcePager.
type ModelingClientListByAccountResourceResponse struct {
	ModelingResourceList
}

// ModelingClientUpdateResponse contains the response from method ModelingClient.BeginUpdate.
type ModelingClientUpdateResponse struct {
	ModelingResource
}

// OperationStatusesClientGetResponse contains the response from method OperationStatusesClient.Get.
type OperationStatusesClientGetResponse struct {
	OperationStatusResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// ServiceEndpointsClientCreateOrUpdateResponse contains the response from method ServiceEndpointsClient.BeginCreateOrUpdate.
type ServiceEndpointsClientCreateOrUpdateResponse struct {
	ServiceEndpointResource
}

// ServiceEndpointsClientDeleteResponse contains the response from method ServiceEndpointsClient.BeginDelete.
type ServiceEndpointsClientDeleteResponse struct {
	// placeholder for future response values
}

// ServiceEndpointsClientGetResponse contains the response from method ServiceEndpointsClient.Get.
type ServiceEndpointsClientGetResponse struct {
	ServiceEndpointResource
}

// ServiceEndpointsClientListByAccountResourceResponse contains the response from method ServiceEndpointsClient.NewListByAccountResourcePager.
type ServiceEndpointsClientListByAccountResourceResponse struct {
	ServiceEndpointResourceList
}

// ServiceEndpointsClientUpdateResponse contains the response from method ServiceEndpointsClient.BeginUpdate.
type ServiceEndpointsClientUpdateResponse struct {
	ServiceEndpointResource
}
