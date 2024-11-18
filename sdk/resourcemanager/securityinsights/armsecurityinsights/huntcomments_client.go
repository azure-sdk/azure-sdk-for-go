//go:build go1.18
// +build go1.18

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

// HuntCommentsClient contains the methods for the HuntComments group.
// Don't use this type directly, use NewHuntCommentsClient() instead.
type HuntCommentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHuntCommentsClient creates a new instance of HuntCommentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHuntCommentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HuntCommentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HuntCommentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a hunt relation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - huntID - The hunt id (GUID)
//   - huntCommentID - The hunt comment id (GUID)
//   - huntComment - The hunt comment
//   - options - HuntCommentsClientCreateOrUpdateOptions contains the optional parameters for the HuntCommentsClient.CreateOrUpdate
//     method.
func (client *HuntCommentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntCommentID string, huntComment HuntComment, options *HuntCommentsClientCreateOrUpdateOptions) (HuntCommentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "HuntCommentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, huntID, huntCommentID, huntComment, options)
	if err != nil {
		return HuntCommentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HuntCommentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HuntCommentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HuntCommentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntCommentID string, huntComment HuntComment, options *HuntCommentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/hunts/{huntId}/comments/{huntCommentId}"
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
	if huntID == "" {
		return nil, errors.New("parameter huntID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntId}", url.PathEscape(huntID))
	if huntCommentID == "" {
		return nil, errors.New("parameter huntCommentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntCommentId}", url.PathEscape(huntCommentID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, huntComment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HuntCommentsClient) createOrUpdateHandleResponse(resp *http.Response) (HuntCommentsClientCreateOrUpdateResponse, error) {
	result := HuntCommentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HuntComment); err != nil {
		return HuntCommentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a hunt comment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - huntID - The hunt id (GUID)
//   - huntCommentID - The hunt comment id (GUID)
//   - options - HuntCommentsClientDeleteOptions contains the optional parameters for the HuntCommentsClient.Delete method.
func (client *HuntCommentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntCommentID string, options *HuntCommentsClientDeleteOptions) (HuntCommentsClientDeleteResponse, error) {
	var err error
	const operationName = "HuntCommentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, huntID, huntCommentID, options)
	if err != nil {
		return HuntCommentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HuntCommentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HuntCommentsClientDeleteResponse{}, err
	}
	return HuntCommentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HuntCommentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntCommentID string, options *HuntCommentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/hunts/{huntId}/comments/{huntCommentId}"
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
	if huntID == "" {
		return nil, errors.New("parameter huntID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntId}", url.PathEscape(huntID))
	if huntCommentID == "" {
		return nil, errors.New("parameter huntCommentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntCommentId}", url.PathEscape(huntCommentID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a hunt comment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - huntID - The hunt id (GUID)
//   - huntCommentID - The hunt comment id (GUID)
//   - options - HuntCommentsClientGetOptions contains the optional parameters for the HuntCommentsClient.Get method.
func (client *HuntCommentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntCommentID string, options *HuntCommentsClientGetOptions) (HuntCommentsClientGetResponse, error) {
	var err error
	const operationName = "HuntCommentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, huntID, huntCommentID, options)
	if err != nil {
		return HuntCommentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HuntCommentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HuntCommentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HuntCommentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, huntCommentID string, options *HuntCommentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/hunts/{huntId}/comments/{huntCommentId}"
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
	if huntID == "" {
		return nil, errors.New("parameter huntID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntId}", url.PathEscape(huntID))
	if huntCommentID == "" {
		return nil, errors.New("parameter huntCommentID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntCommentId}", url.PathEscape(huntCommentID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HuntCommentsClient) getHandleResponse(resp *http.Response) (HuntCommentsClientGetResponse, error) {
	result := HuntCommentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HuntComment); err != nil {
		return HuntCommentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all hunt comments
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - huntID - The hunt id (GUID)
//   - options - HuntCommentsClientListOptions contains the optional parameters for the HuntCommentsClient.NewListPager method.
func (client *HuntCommentsClient) NewListPager(resourceGroupName string, workspaceName string, huntID string, options *HuntCommentsClientListOptions) *runtime.Pager[HuntCommentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HuntCommentsClientListResponse]{
		More: func(page HuntCommentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HuntCommentsClientListResponse) (HuntCommentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HuntCommentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, huntID, options)
			}, nil)
			if err != nil {
				return HuntCommentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HuntCommentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, huntID string, options *HuntCommentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/hunts/{huntId}/comments"
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
	if huntID == "" {
		return nil, errors.New("parameter huntID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{huntId}", url.PathEscape(huntID))
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
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HuntCommentsClient) listHandleResponse(resp *http.Response) (HuntCommentsClientListResponse, error) {
	result := HuntCommentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HuntCommentList); err != nil {
		return HuntCommentsClientListResponse{}, err
	}
	return result, nil
}
