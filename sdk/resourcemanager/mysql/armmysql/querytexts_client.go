// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysql

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

// QueryTextsClient contains the methods for the QueryTexts group.
// Don't use this type directly, use NewQueryTextsClient() instead.
type QueryTextsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQueryTextsClient creates a new instance of QueryTextsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQueryTextsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QueryTextsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QueryTextsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve the Query-Store query texts for the queryId.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - queryID - The Query-Store query identifier.
//   - options - QueryTextsClientGetOptions contains the optional parameters for the QueryTextsClient.Get method.
func (client *QueryTextsClient) Get(ctx context.Context, resourceGroupName string, serverName string, queryID string, options *QueryTextsClientGetOptions) (QueryTextsClientGetResponse, error) {
	var err error
	const operationName = "QueryTextsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, queryID, options)
	if err != nil {
		return QueryTextsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueryTextsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return QueryTextsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *QueryTextsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, queryID string, _ *QueryTextsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/queryTexts/{queryId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if queryID == "" {
		return nil, errors.New("parameter queryID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryId}", url.PathEscape(queryID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QueryTextsClient) getHandleResponse(resp *http.Response) (QueryTextsClientGetResponse, error) {
	result := QueryTextsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryText); err != nil {
		return QueryTextsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Retrieve the Query-Store query texts for specified queryIds.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - queryIDs - The query identifiers
//   - options - QueryTextsClientListByServerOptions contains the optional parameters for the QueryTextsClient.NewListByServerPager
//     method.
func (client *QueryTextsClient) NewListByServerPager(resourceGroupName string, serverName string, queryIDs []string, options *QueryTextsClientListByServerOptions) *runtime.Pager[QueryTextsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueryTextsClientListByServerResponse]{
		More: func(page QueryTextsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueryTextsClientListByServerResponse) (QueryTextsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "QueryTextsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, queryIDs, options)
			}, nil)
			if err != nil {
				return QueryTextsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *QueryTextsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, queryIDs []string, _ *QueryTextsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/servers/{serverName}/queryTexts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-06-01")
	for _, qv := range queryIDs {
		reqQP.Add("queryIds", qv)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *QueryTextsClient) listByServerHandleResponse(resp *http.Response) (QueryTextsClientListByServerResponse, error) {
	result := QueryTextsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QueryTextsResultList); err != nil {
		return QueryTextsClientListByServerResponse{}, err
	}
	return result, nil
}
