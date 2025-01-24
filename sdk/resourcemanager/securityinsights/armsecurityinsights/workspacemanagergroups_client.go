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

// WorkspaceManagerGroupsClient contains the methods for the WorkspaceManagerGroups group.
// Don't use this type directly, use NewWorkspaceManagerGroupsClient() instead.
type WorkspaceManagerGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagerGroupsClient creates a new instance of WorkspaceManagerGroupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagerGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagerGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagerGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a workspace manager group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerGroupName - The name of the workspace manager group
//   - workspaceManagerGroup - The workspace manager group object
//   - options - WorkspaceManagerGroupsClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceManagerGroupsClient.CreateOrUpdate
//     method.
func (client *WorkspaceManagerGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerGroupName string, workspaceManagerGroup WorkspaceManagerGroup, options *WorkspaceManagerGroupsClientCreateOrUpdateOptions) (WorkspaceManagerGroupsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceManagerGroupsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerGroupName, workspaceManagerGroup, options)
	if err != nil {
		return WorkspaceManagerGroupsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagerGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerGroupName string, workspaceManagerGroup WorkspaceManagerGroup, options *WorkspaceManagerGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerGroups/{workspaceManagerGroupName}"
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
	if workspaceManagerGroupName == "" {
		return nil, errors.New("parameter workspaceManagerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerGroupName}", url.PathEscape(workspaceManagerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, workspaceManagerGroup); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceManagerGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceManagerGroupsClientCreateOrUpdateResponse, error) {
	result := WorkspaceManagerGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerGroup); err != nil {
		return WorkspaceManagerGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a workspace manager group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerGroupName - The name of the workspace manager group
//   - options - WorkspaceManagerGroupsClientDeleteOptions contains the optional parameters for the WorkspaceManagerGroupsClient.Delete
//     method.
func (client *WorkspaceManagerGroupsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerGroupName string, options *WorkspaceManagerGroupsClientDeleteOptions) (WorkspaceManagerGroupsClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceManagerGroupsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerGroupName, options)
	if err != nil {
		return WorkspaceManagerGroupsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerGroupsClientDeleteResponse{}, err
	}
	return WorkspaceManagerGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceManagerGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerGroupName string, options *WorkspaceManagerGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerGroups/{workspaceManagerGroupName}"
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
	if workspaceManagerGroupName == "" {
		return nil, errors.New("parameter workspaceManagerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerGroupName}", url.PathEscape(workspaceManagerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a workspace manager group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerGroupName - The name of the workspace manager group
//   - options - WorkspaceManagerGroupsClientGetOptions contains the optional parameters for the WorkspaceManagerGroupsClient.Get
//     method.
func (client *WorkspaceManagerGroupsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerGroupName string, options *WorkspaceManagerGroupsClientGetOptions) (WorkspaceManagerGroupsClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceManagerGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerGroupName, options)
	if err != nil {
		return WorkspaceManagerGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagerGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerGroupName string, options *WorkspaceManagerGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerGroups/{workspaceManagerGroupName}"
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
	if workspaceManagerGroupName == "" {
		return nil, errors.New("parameter workspaceManagerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerGroupName}", url.PathEscape(workspaceManagerGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagerGroupsClient) getHandleResponse(resp *http.Response) (WorkspaceManagerGroupsClientGetResponse, error) {
	result := WorkspaceManagerGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerGroup); err != nil {
		return WorkspaceManagerGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all workspace manager groups in the Sentinel workspace manager
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagerGroupsClientListOptions contains the optional parameters for the WorkspaceManagerGroupsClient.NewListPager
//     method.
func (client *WorkspaceManagerGroupsClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagerGroupsClientListOptions) *runtime.Pager[WorkspaceManagerGroupsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagerGroupsClientListResponse]{
		More: func(page WorkspaceManagerGroupsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagerGroupsClientListResponse) (WorkspaceManagerGroupsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceManagerGroupsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return WorkspaceManagerGroupsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagerGroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagerGroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerGroups"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagerGroupsClient) listHandleResponse(resp *http.Response) (WorkspaceManagerGroupsClientListResponse, error) {
	result := WorkspaceManagerGroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerGroupList); err != nil {
		return WorkspaceManagerGroupsClientListResponse{}, err
	}
	return result, nil
}
