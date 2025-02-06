//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquantum

// OfferingsClientListResponse contains the response from method OfferingsClient.NewListPager.
type OfferingsClientListResponse struct {
	// The response of a list Providers operation.
	OfferingsListResult
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// WorkspacesClientCheckNameAvailabilityResponse contains the response from method WorkspacesClient.CheckNameAvailability.
type WorkspacesClientCheckNameAvailabilityResponse struct {
	// The check availability result.
	CheckNameAvailabilityResponse
}

// WorkspacesClientCreateOrUpdateResponse contains the response from method WorkspacesClient.BeginCreateOrUpdate.
type WorkspacesClientCreateOrUpdateResponse struct {
	// The resource proxy definition object for Quantum Workspace.
	Workspace
}

// WorkspacesClientDeleteResponse contains the response from method WorkspacesClient.BeginDelete.
type WorkspacesClientDeleteResponse struct {
	// placeholder for future response values
}

// WorkspacesClientGetResponse contains the response from method WorkspacesClient.Get.
type WorkspacesClientGetResponse struct {
	// The resource proxy definition object for Quantum Workspace.
	Workspace
}

// WorkspacesClientListByResourceGroupResponse contains the response from method WorkspacesClient.NewListByResourceGroupPager.
type WorkspacesClientListByResourceGroupResponse struct {
	// The response of a QuantumWorkspace list operation.
	WorkspaceListResult
}

// WorkspacesClientListBySubscriptionResponse contains the response from method WorkspacesClient.NewListBySubscriptionPager.
type WorkspacesClientListBySubscriptionResponse struct {
	// The response of a QuantumWorkspace list operation.
	WorkspaceListResult
}

// WorkspacesClientListKeysResponse contains the response from method WorkspacesClient.ListKeys.
type WorkspacesClientListKeysResponse struct {
	// Result of list Api keys and connection strings.
	ListKeysResult
}

// WorkspacesClientRegenerateKeysResponse contains the response from method WorkspacesClient.RegenerateKeys.
type WorkspacesClientRegenerateKeysResponse struct {
	// placeholder for future response values
}

// WorkspacesClientUpdateTagsResponse contains the response from method WorkspacesClient.UpdateTags.
type WorkspacesClientUpdateTagsResponse struct {
	// The resource proxy definition object for Quantum Workspace.
	Workspace
}
