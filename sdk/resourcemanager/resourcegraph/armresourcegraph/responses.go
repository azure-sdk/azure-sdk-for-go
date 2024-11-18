//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcegraph

// ClientResourceChangeDetailsResponse contains the response from method Client.ResourceChangeDetails.
type ClientResourceChangeDetailsResponse struct {
	// A list of change details associated with a resource over a specific time interval.
	ResourceChangeDataArray []*ResourceChangeData
}

// ClientResourceChangesResponse contains the response from method Client.ResourceChanges.
type ClientResourceChangesResponse struct {
	// A list of changes associated with a resource over a specific time interval.
	ResourceChangeList
}

// ClientResourcesHistoryResponse contains the response from method Client.ResourcesHistory.
type ClientResourcesHistoryResponse struct {
	// Dictionary of <any>
	Value map[string]any
}

// ClientResourcesResponse contains the response from method Client.Resources.
type ClientResourcesResponse struct {
	// Query result.
	QueryResponse
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	// Result of the request to list Resource Graph operations. It contains a list of operations and a URL link to get the next
	// set of results.
	OperationListResult
}

// QueryClientGenerateQueryResponse contains the response from method QueryClient.GenerateQuery.
type QueryClientGenerateQueryResponse struct {
	// Query Generation Response
	QueryGenerationResponse
}