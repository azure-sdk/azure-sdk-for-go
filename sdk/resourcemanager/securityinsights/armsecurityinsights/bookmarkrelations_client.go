// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// BookmarkRelationsClient contains the methods for the BookmarkRelations group.
// Don't use this type directly, use NewBookmarkRelationsClient() instead.
type BookmarkRelationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBookmarkRelationsClient creates a new instance of BookmarkRelationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBookmarkRelationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BookmarkRelationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BookmarkRelationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates the bookmark relation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - bookmarkID - Bookmark ID
//   - relationName - Relation Name
//   - relation - The relation model
//   - options - BookmarkRelationsClientCreateOrUpdateOptions contains the optional parameters for the BookmarkRelationsClient.CreateOrUpdate
//     method.
func (client *BookmarkRelationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, relation Relation, options *BookmarkRelationsClientCreateOrUpdateOptions) (BookmarkRelationsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "BookmarkRelationsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, bookmarkID, relationName, relation, options)
	if err != nil {
		return BookmarkRelationsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BookmarkRelationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return BookmarkRelationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BookmarkRelationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, relation Relation, _ *BookmarkRelationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if bookmarkID == "" {
		return nil, errors.New("parameter bookmarkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bookmarkId}", url.PathEscape(bookmarkID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, relation); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BookmarkRelationsClient) createOrUpdateHandleResponse(resp *http.Response) (BookmarkRelationsClientCreateOrUpdateResponse, error) {
	result := BookmarkRelationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Relation); err != nil {
		return BookmarkRelationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the bookmark relation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - bookmarkID - Bookmark ID
//   - relationName - Relation Name
//   - options - BookmarkRelationsClientDeleteOptions contains the optional parameters for the BookmarkRelationsClient.Delete
//     method.
func (client *BookmarkRelationsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, options *BookmarkRelationsClientDeleteOptions) (BookmarkRelationsClientDeleteResponse, error) {
	var err error
	const operationName = "BookmarkRelationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, bookmarkID, relationName, options)
	if err != nil {
		return BookmarkRelationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BookmarkRelationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BookmarkRelationsClientDeleteResponse{}, err
	}
	return BookmarkRelationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BookmarkRelationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, _ *BookmarkRelationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if bookmarkID == "" {
		return nil, errors.New("parameter bookmarkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bookmarkId}", url.PathEscape(bookmarkID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a bookmark relation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - bookmarkID - Bookmark ID
//   - relationName - Relation Name
//   - options - BookmarkRelationsClientGetOptions contains the optional parameters for the BookmarkRelationsClient.Get method.
func (client *BookmarkRelationsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, options *BookmarkRelationsClientGetOptions) (BookmarkRelationsClientGetResponse, error) {
	var err error
	const operationName = "BookmarkRelationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, bookmarkID, relationName, options)
	if err != nil {
		return BookmarkRelationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BookmarkRelationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BookmarkRelationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BookmarkRelationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, relationName string, _ *BookmarkRelationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations/{relationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if bookmarkID == "" {
		return nil, errors.New("parameter bookmarkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bookmarkId}", url.PathEscape(bookmarkID))
	if relationName == "" {
		return nil, errors.New("parameter relationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{relationName}", url.PathEscape(relationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BookmarkRelationsClient) getHandleResponse(resp *http.Response) (BookmarkRelationsClientGetResponse, error) {
	result := BookmarkRelationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Relation); err != nil {
		return BookmarkRelationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all bookmark relations.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - bookmarkID - Bookmark ID
//   - options - BookmarkRelationsClientListOptions contains the optional parameters for the BookmarkRelationsClient.NewListPager
//     method.
func (client *BookmarkRelationsClient) NewListPager(resourceGroupName string, workspaceName string, bookmarkID string, options *BookmarkRelationsClientListOptions) *runtime.Pager[BookmarkRelationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BookmarkRelationsClientListResponse]{
		More: func(page BookmarkRelationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BookmarkRelationsClientListResponse) (BookmarkRelationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BookmarkRelationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, bookmarkID, options)
			}, nil)
			if err != nil {
				return BookmarkRelationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BookmarkRelationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, bookmarkID string, options *BookmarkRelationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/bookmarks/{bookmarkId}/relations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if bookmarkID == "" {
		return nil, errors.New("parameter bookmarkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{bookmarkId}", url.PathEscape(bookmarkID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BookmarkRelationsClient) listHandleResponse(resp *http.Response) (BookmarkRelationsClientListResponse, error) {
	result := BookmarkRelationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationList); err != nil {
		return BookmarkRelationsClientListResponse{}, err
	}
	return result, nil
}
