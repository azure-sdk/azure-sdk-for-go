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

// WorkspaceManagerAssignmentJobsClient contains the methods for the WorkspaceManagerAssignmentJobs group.
// Don't use this type directly, use NewWorkspaceManagerAssignmentJobsClient() instead.
type WorkspaceManagerAssignmentJobsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagerAssignmentJobsClient creates a new instance of WorkspaceManagerAssignmentJobsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagerAssignmentJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagerAssignmentJobsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagerAssignmentJobsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a job for the specified workspace manager assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - options - WorkspaceManagerAssignmentJobsClientCreateOptions contains the optional parameters for the WorkspaceManagerAssignmentJobsClient.Create
//     method.
func (client *WorkspaceManagerAssignmentJobsClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentJobsClientCreateOptions) (WorkspaceManagerAssignmentJobsClientCreateResponse, error) {
	var err error
	const operationName = "WorkspaceManagerAssignmentJobsClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, options)
	if err != nil {
		return WorkspaceManagerAssignmentJobsClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerAssignmentJobsClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerAssignmentJobsClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *WorkspaceManagerAssignmentJobsClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentJobsClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}/jobs"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *WorkspaceManagerAssignmentJobsClient) createHandleResponse(resp *http.Response) (WorkspaceManagerAssignmentJobsClientCreateResponse, error) {
	result := WorkspaceManagerAssignmentJobsClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return WorkspaceManagerAssignmentJobsClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified job from the specified workspace manager assignment
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - jobName - The job name
//   - options - WorkspaceManagerAssignmentJobsClientDeleteOptions contains the optional parameters for the WorkspaceManagerAssignmentJobsClient.Delete
//     method.
func (client *WorkspaceManagerAssignmentJobsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, jobName string, options *WorkspaceManagerAssignmentJobsClientDeleteOptions) (WorkspaceManagerAssignmentJobsClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceManagerAssignmentJobsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, jobName, options)
	if err != nil {
		return WorkspaceManagerAssignmentJobsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerAssignmentJobsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerAssignmentJobsClientDeleteResponse{}, err
	}
	return WorkspaceManagerAssignmentJobsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceManagerAssignmentJobsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, jobName string, options *WorkspaceManagerAssignmentJobsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}/jobs/{jobName}"
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
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a job
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - jobName - The job name
//   - options - WorkspaceManagerAssignmentJobsClientGetOptions contains the optional parameters for the WorkspaceManagerAssignmentJobsClient.Get
//     method.
func (client *WorkspaceManagerAssignmentJobsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, jobName string, options *WorkspaceManagerAssignmentJobsClientGetOptions) (WorkspaceManagerAssignmentJobsClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceManagerAssignmentJobsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, jobName, options)
	if err != nil {
		return WorkspaceManagerAssignmentJobsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerAssignmentJobsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerAssignmentJobsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagerAssignmentJobsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, jobName string, options *WorkspaceManagerAssignmentJobsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}/jobs/{jobName}"
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
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagerAssignmentJobsClient) getHandleResponse(resp *http.Response) (WorkspaceManagerAssignmentJobsClientGetResponse, error) {
	result := WorkspaceManagerAssignmentJobsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Job); err != nil {
		return WorkspaceManagerAssignmentJobsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all jobs for the specified workspace manager assignment
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerAssignmentName - The name of the workspace manager assignment
//   - options - WorkspaceManagerAssignmentJobsClientListOptions contains the optional parameters for the WorkspaceManagerAssignmentJobsClient.NewListPager
//     method.
func (client *WorkspaceManagerAssignmentJobsClient) NewListPager(resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentJobsClientListOptions) *runtime.Pager[WorkspaceManagerAssignmentJobsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagerAssignmentJobsClientListResponse]{
		More: func(page WorkspaceManagerAssignmentJobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagerAssignmentJobsClientListResponse) (WorkspaceManagerAssignmentJobsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceManagerAssignmentJobsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerAssignmentName, options)
			}, nil)
			if err != nil {
				return WorkspaceManagerAssignmentJobsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagerAssignmentJobsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerAssignmentName string, options *WorkspaceManagerAssignmentJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerAssignments/{workspaceManagerAssignmentName}/jobs"
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
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagerAssignmentJobsClient) listHandleResponse(resp *http.Response) (WorkspaceManagerAssignmentJobsClientListResponse, error) {
	result := WorkspaceManagerAssignmentJobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.JobList); err != nil {
		return WorkspaceManagerAssignmentJobsClientListResponse{}, err
	}
	return result, nil
}
