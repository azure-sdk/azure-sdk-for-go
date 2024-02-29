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

// WorkspaceManagerAssignmentsClient contains the methods for the WorkspaceManagerAssignments group.
// Don't use this type directly, use NewWorkspaceManagerAssignmentsClient() instead.
type WorkspaceManagerAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagerAssignmentsClient creates a new instance of WorkspaceManagerAssignmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagerAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagerAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagerAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a workspace manager assignment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - workspaceManagerAssignment - The workspace manager assignment
//   - options - WorkspaceManagerAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceManagerAssignmentsClient.CreateOrUpdate
//     method.
func (client *WorkspaceManagerAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, workspaceManagerAssignment WorkspaceManagerAssignment, options *WorkspaceManagerAssignmentsClientCreateOrUpdateOptions) (WorkspaceManagerAssignmentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceManagerAssignmentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, workspaceManagerAssignment, options)
	if err != nil {
		return WorkspaceManagerAssignmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagerAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, workspaceManagerAssignment WorkspaceManagerAssignment, options *WorkspaceManagerAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}"
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
	if workspaceManagerAssignmentName == "" {
		return nil, errors.New("parameter workspaceManagerAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerAssignmentName}", url.PathEscape(workspaceManagerAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, workspaceManagerAssignment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceManagerAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceManagerAssignmentsClientCreateOrUpdateResponse, error) {
	result := WorkspaceManagerAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerAssignment); err != nil {
		return WorkspaceManagerAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a workspace manager assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - options - WorkspaceManagerAssignmentsClientDeleteOptions contains the optional parameters for the WorkspaceManagerAssignmentsClient.Delete
//     method.
func (client *WorkspaceManagerAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentsClientDeleteOptions) (WorkspaceManagerAssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceManagerAssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, options)
	if err != nil {
		return WorkspaceManagerAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerAssignmentsClientDeleteResponse{}, err
	}
	return WorkspaceManagerAssignmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceManagerAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}"
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
	if workspaceManagerAssignmentName == "" {
		return nil, errors.New("parameter workspaceManagerAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerAssignmentName}", url.PathEscape(workspaceManagerAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a workspace manager assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - options - WorkspaceManagerAssignmentsClientGetOptions contains the optional parameters for the WorkspaceManagerAssignmentsClient.Get
//     method.
func (client *WorkspaceManagerAssignmentsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentsClientGetOptions) (WorkspaceManagerAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceManagerAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, options)
	if err != nil {
		return WorkspaceManagerAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagerAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}"
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
	if workspaceManagerAssignmentName == "" {
		return nil, errors.New("parameter workspaceManagerAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerAssignmentName}", url.PathEscape(workspaceManagerAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagerAssignmentsClient) getHandleResponse(resp *http.Response) (WorkspaceManagerAssignmentsClientGetResponse, error) {
	result := WorkspaceManagerAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerAssignment); err != nil {
		return WorkspaceManagerAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all workspace manager assignments for the Sentinel workspace manager.
//
// Generated from API version 2024-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagerAssignmentsClientListOptions contains the optional parameters for the WorkspaceManagerAssignmentsClient.NewListPager
//     method.
func (client *WorkspaceManagerAssignmentsClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagerAssignmentsClientListOptions) *runtime.Pager[WorkspaceManagerAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagerAssignmentsClientListResponse]{
		More: func(page WorkspaceManagerAssignmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagerAssignmentsClientListResponse) (WorkspaceManagerAssignmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceManagerAssignmentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return WorkspaceManagerAssignmentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagerAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagerAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments"
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
	reqQP.Set("api-version", "2024-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagerAssignmentsClient) listHandleResponse(resp *http.Response) (WorkspaceManagerAssignmentsClientListResponse, error) {
	result := WorkspaceManagerAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerAssignmentList); err != nil {
		return WorkspaceManagerAssignmentsClientListResponse{}, err
	}
	return result, nil
}
