//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ManagedDatabaseQueriesClient contains the methods for the ManagedDatabaseQueries group.
// Don't use this type directly, use NewManagedDatabaseQueriesClient() instead.
type ManagedDatabaseQueriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedDatabaseQueriesClient creates a new instance of ManagedDatabaseQueriesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedDatabaseQueriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseQueriesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedDatabaseQueriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get query by query id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedDatabaseQueriesClientGetOptions contains the optional parameters for the ManagedDatabaseQueriesClient.Get
//     method.
func (client *ManagedDatabaseQueriesClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesClientGetOptions) (ManagedDatabaseQueriesClientGetResponse, error) {
	var err error
	const operationName = "ManagedDatabaseQueriesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, queryID, options)
	if err != nil {
		return ManagedDatabaseQueriesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseQueriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedDatabaseQueriesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseQueriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/queries/{queryId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if queryID == "" {
		return nil, errors.New("parameter queryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryId}", url.PathEscape(queryID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseQueriesClient) getHandleResponse(resp *http.Response) (ManagedDatabaseQueriesClientGetResponse, error) {
	result := ManagedDatabaseQueriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceQuery); err != nil {
		return ManagedDatabaseQueriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByQueryPager - Get query execution statistics by query id.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedDatabaseQueriesClientListByQueryOptions contains the optional parameters for the ManagedDatabaseQueriesClient.NewListByQueryPager
//     method.
func (client *ManagedDatabaseQueriesClient) NewListByQueryPager(resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesClientListByQueryOptions) *runtime.Pager[ManagedDatabaseQueriesClientListByQueryResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedDatabaseQueriesClientListByQueryResponse]{
		More: func(page ManagedDatabaseQueriesClientListByQueryResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseQueriesClientListByQueryResponse) (ManagedDatabaseQueriesClientListByQueryResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedDatabaseQueriesClient.NewListByQueryPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByQueryCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, queryID, options)
			}, nil)
			if err != nil {
				return ManagedDatabaseQueriesClientListByQueryResponse{}, err
			}
			return client.listByQueryHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByQueryCreateRequest creates the ListByQuery request.
func (client *ManagedDatabaseQueriesClient) listByQueryCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, queryID string, options *ManagedDatabaseQueriesClientListByQueryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/queries/{queryId}/statistics"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if queryID == "" {
		return nil, errors.New("parameter queryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryId}", url.PathEscape(queryID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	if options != nil && options.EndTime != nil {
		reqQP.Set("endTime", *options.EndTime)
	}
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", string(*options.Interval))
	}
	if options != nil && options.StartTime != nil {
		reqQP.Set("startTime", *options.StartTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByQueryHandleResponse handles the ListByQuery response.
func (client *ManagedDatabaseQueriesClient) listByQueryHandleResponse(resp *http.Response) (ManagedDatabaseQueriesClientListByQueryResponse, error) {
	result := ManagedDatabaseQueriesClientListByQueryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedInstanceQueryStatistics); err != nil {
		return ManagedDatabaseQueriesClientListByQueryResponse{}, err
	}
	return result, nil
}
