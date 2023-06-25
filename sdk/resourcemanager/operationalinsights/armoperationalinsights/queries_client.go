//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armoperationalinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// QueriesClient contains the methods for the Queries group.
// Don't use this type directly, use NewQueriesClient() instead.
type QueriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewQueriesClient creates a new instance of QueriesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewQueriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*QueriesClient, error) {
	cl, err := arm.NewClient(moduleName+".QueriesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &QueriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Deletes a specific Query defined within an Log Analytics QueryPack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - queryPackName - The name of the Log Analytics QueryPack resource.
//   - id - The id of a specific query defined in the Log Analytics QueryPack
//   - options - QueriesClientDeleteOptions contains the optional parameters for the QueriesClient.Delete method.
func (client *QueriesClient) Delete(ctx context.Context, resourceGroupName string, queryPackName string, id string, options *QueriesClientDeleteOptions) (QueriesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, queryPackName, id, options)
	if err != nil {
		return QueriesClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueriesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return QueriesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return QueriesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *QueriesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, queryPackName string, id string, options *QueriesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if queryPackName == "" {
		return nil, errors.New("parameter queryPackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryPackName}", url.PathEscape(queryPackName))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a specific Log Analytics Query defined within a Log Analytics QueryPack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - queryPackName - The name of the Log Analytics QueryPack resource.
//   - id - The id of a specific query defined in the Log Analytics QueryPack
//   - options - QueriesClientGetOptions contains the optional parameters for the QueriesClient.Get method.
func (client *QueriesClient) Get(ctx context.Context, resourceGroupName string, queryPackName string, id string, options *QueriesClientGetOptions) (QueriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, queryPackName, id, options)
	if err != nil {
		return QueriesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *QueriesClient) getCreateRequest(ctx context.Context, resourceGroupName string, queryPackName string, id string, options *QueriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if queryPackName == "" {
		return nil, errors.New("parameter queryPackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryPackName}", url.PathEscape(queryPackName))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *QueriesClient) getHandleResponse(resp *http.Response) (QueriesClientGetResponse, error) {
	result := QueriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogAnalyticsQueryPackQuery); err != nil {
		return QueriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of Queries defined within a Log Analytics QueryPack.
//
// Generated from API version 2019-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - queryPackName - The name of the Log Analytics QueryPack resource.
//   - options - QueriesClientListOptions contains the optional parameters for the QueriesClient.NewListPager method.
func (client *QueriesClient) NewListPager(resourceGroupName string, queryPackName string, options *QueriesClientListOptions) *runtime.Pager[QueriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueriesClientListResponse]{
		More: func(page QueriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueriesClientListResponse) (QueriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, queryPackName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return QueriesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return QueriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return QueriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *QueriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, queryPackName string, options *QueriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if queryPackName == "" {
		return nil, errors.New("parameter queryPackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryPackName}", url.PathEscape(queryPackName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(*options.Top, 10))
	}
	if options != nil && options.IncludeBody != nil {
		reqQP.Set("includeBody", strconv.FormatBool(*options.IncludeBody))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *QueriesClient) listHandleResponse(resp *http.Response) (QueriesClientListResponse, error) {
	result := QueriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogAnalyticsQueryPackQueryListResult); err != nil {
		return QueriesClientListResponse{}, err
	}
	return result, nil
}

// Put - Adds or Updates a specific Query within a Log Analytics QueryPack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - queryPackName - The name of the Log Analytics QueryPack resource.
//   - id - The id of a specific query defined in the Log Analytics QueryPack
//   - queryPayload - Properties that need to be specified to create a new query and add it to a Log Analytics QueryPack.
//   - options - QueriesClientPutOptions contains the optional parameters for the QueriesClient.Put method.
func (client *QueriesClient) Put(ctx context.Context, resourceGroupName string, queryPackName string, id string, queryPayload LogAnalyticsQueryPackQuery, options *QueriesClientPutOptions) (QueriesClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, resourceGroupName, queryPackName, id, queryPayload, options)
	if err != nil {
		return QueriesClientPutResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueriesClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *QueriesClient) putCreateRequest(ctx context.Context, resourceGroupName string, queryPackName string, id string, queryPayload LogAnalyticsQueryPackQuery, options *QueriesClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if queryPackName == "" {
		return nil, errors.New("parameter queryPackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryPackName}", url.PathEscape(queryPackName))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, queryPayload)
}

// putHandleResponse handles the Put response.
func (client *QueriesClient) putHandleResponse(resp *http.Response) (QueriesClientPutResponse, error) {
	result := QueriesClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogAnalyticsQueryPackQuery); err != nil {
		return QueriesClientPutResponse{}, err
	}
	return result, nil
}

// NewSearchPager - Search a list of Queries defined within a Log Analytics QueryPack according to given search properties.
//
// Generated from API version 2019-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - queryPackName - The name of the Log Analytics QueryPack resource.
//   - querySearchProperties - Properties by which to search queries in the given Log Analytics QueryPack.
//   - options - QueriesClientSearchOptions contains the optional parameters for the QueriesClient.NewSearchPager method.
func (client *QueriesClient) NewSearchPager(resourceGroupName string, queryPackName string, querySearchProperties LogAnalyticsQueryPackQuerySearchProperties, options *QueriesClientSearchOptions) *runtime.Pager[QueriesClientSearchResponse] {
	return runtime.NewPager(runtime.PagingHandler[QueriesClientSearchResponse]{
		More: func(page QueriesClientSearchResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *QueriesClientSearchResponse) (QueriesClientSearchResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.searchCreateRequest(ctx, resourceGroupName, queryPackName, querySearchProperties, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return QueriesClientSearchResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return QueriesClientSearchResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return QueriesClientSearchResponse{}, runtime.NewResponseError(resp)
			}
			return client.searchHandleResponse(resp)
		},
	})
}

// searchCreateRequest creates the Search request.
func (client *QueriesClient) searchCreateRequest(ctx context.Context, resourceGroupName string, queryPackName string, querySearchProperties LogAnalyticsQueryPackQuerySearchProperties, options *QueriesClientSearchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/search"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if queryPackName == "" {
		return nil, errors.New("parameter queryPackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryPackName}", url.PathEscape(queryPackName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(*options.Top, 10))
	}
	if options != nil && options.IncludeBody != nil {
		reqQP.Set("includeBody", strconv.FormatBool(*options.IncludeBody))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, querySearchProperties)
}

// searchHandleResponse handles the Search response.
func (client *QueriesClient) searchHandleResponse(resp *http.Response) (QueriesClientSearchResponse, error) {
	result := QueriesClientSearchResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogAnalyticsQueryPackQueryListResult); err != nil {
		return QueriesClientSearchResponse{}, err
	}
	return result, nil
}

// Update - Adds or Updates a specific Query within a Log Analytics QueryPack.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - queryPackName - The name of the Log Analytics QueryPack resource.
//   - id - The id of a specific query defined in the Log Analytics QueryPack
//   - queryPayload - Properties that need to be specified to create a new query and add it to a Log Analytics QueryPack.
//   - options - QueriesClientUpdateOptions contains the optional parameters for the QueriesClient.Update method.
func (client *QueriesClient) Update(ctx context.Context, resourceGroupName string, queryPackName string, id string, queryPayload LogAnalyticsQueryPackQuery, options *QueriesClientUpdateOptions) (QueriesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, queryPackName, id, queryPayload, options)
	if err != nil {
		return QueriesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return QueriesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return QueriesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *QueriesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, queryPackName string, id string, queryPayload LogAnalyticsQueryPackQuery, options *QueriesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/queryPacks/{queryPackName}/queries/{id}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if queryPackName == "" {
		return nil, errors.New("parameter queryPackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{queryPackName}", url.PathEscape(queryPackName))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, queryPayload)
}

// updateHandleResponse handles the Update response.
func (client *QueriesClient) updateHandleResponse(resp *http.Response) (QueriesClientUpdateResponse, error) {
	result := QueriesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogAnalyticsQueryPackQuery); err != nil {
		return QueriesClientUpdateResponse{}, err
	}
	return result, nil
}
