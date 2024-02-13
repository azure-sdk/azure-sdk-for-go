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

// WorkspaceManagerConfigurationsClient contains the methods for the WorkspaceManagerConfigurations group.
// Don't use this type directly, use NewWorkspaceManagerConfigurationsClient() instead.
type WorkspaceManagerConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagerConfigurationsClient creates a new instance of WorkspaceManagerConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagerConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagerConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagerConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a workspace manager configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerConfigurationName - The name of the workspace manager configuration
//   - workspaceManagerConfiguration - The workspace manager configuration
//   - options - WorkspaceManagerConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceManagerConfigurationsClient.CreateOrUpdate
//     method.
func (client *WorkspaceManagerConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerConfigurationName string, workspaceManagerConfiguration WorkspaceManagerConfiguration, options *WorkspaceManagerConfigurationsClientCreateOrUpdateOptions) (WorkspaceManagerConfigurationsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceManagerConfigurationsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerConfigurationName, workspaceManagerConfiguration, options)
	if err != nil {
		return WorkspaceManagerConfigurationsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerConfigurationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerConfigurationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagerConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerConfigurationName string, workspaceManagerConfiguration WorkspaceManagerConfiguration, options *WorkspaceManagerConfigurationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerConfigurations/{workspaceManagerConfigurationName}"
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
	if workspaceManagerConfigurationName == "" {
		return nil, errors.New("parameter workspaceManagerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerConfigurationName}", url.PathEscape(workspaceManagerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, workspaceManagerConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceManagerConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceManagerConfigurationsClientCreateOrUpdateResponse, error) {
	result := WorkspaceManagerConfigurationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerConfiguration); err != nil {
		return WorkspaceManagerConfigurationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a workspace manager configuration
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerConfigurationName - The name of the workspace manager configuration
//   - options - WorkspaceManagerConfigurationsClientDeleteOptions contains the optional parameters for the WorkspaceManagerConfigurationsClient.Delete
//     method.
func (client *WorkspaceManagerConfigurationsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerConfigurationName string, options *WorkspaceManagerConfigurationsClientDeleteOptions) (WorkspaceManagerConfigurationsClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceManagerConfigurationsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerConfigurationName, options)
	if err != nil {
		return WorkspaceManagerConfigurationsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerConfigurationsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerConfigurationsClientDeleteResponse{}, err
	}
	return WorkspaceManagerConfigurationsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceManagerConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerConfigurationName string, options *WorkspaceManagerConfigurationsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerConfigurations/{workspaceManagerConfigurationName}"
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
	if workspaceManagerConfigurationName == "" {
		return nil, errors.New("parameter workspaceManagerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerConfigurationName}", url.PathEscape(workspaceManagerConfigurationName))
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

// Get - Gets a workspace manager configuration
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerConfigurationName - The name of the workspace manager configuration
//   - options - WorkspaceManagerConfigurationsClientGetOptions contains the optional parameters for the WorkspaceManagerConfigurationsClient.Get
//     method.
func (client *WorkspaceManagerConfigurationsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerConfigurationName string, options *WorkspaceManagerConfigurationsClientGetOptions) (WorkspaceManagerConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceManagerConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerConfigurationName, options)
	if err != nil {
		return WorkspaceManagerConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagerConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerConfigurationName string, options *WorkspaceManagerConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerConfigurations/{workspaceManagerConfigurationName}"
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
	if workspaceManagerConfigurationName == "" {
		return nil, errors.New("parameter workspaceManagerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerConfigurationName}", url.PathEscape(workspaceManagerConfigurationName))
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
func (client *WorkspaceManagerConfigurationsClient) getHandleResponse(resp *http.Response) (WorkspaceManagerConfigurationsClientGetResponse, error) {
	result := WorkspaceManagerConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerConfiguration); err != nil {
		return WorkspaceManagerConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all workspace manager configurations for a Sentinel workspace.
//
// Generated from API version 2024-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagerConfigurationsClientListOptions contains the optional parameters for the WorkspaceManagerConfigurationsClient.NewListPager
//     method.
func (client *WorkspaceManagerConfigurationsClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagerConfigurationsClientListOptions) *runtime.Pager[WorkspaceManagerConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagerConfigurationsClientListResponse]{
		More: func(page WorkspaceManagerConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagerConfigurationsClientListResponse) (WorkspaceManagerConfigurationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceManagerConfigurationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return WorkspaceManagerConfigurationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagerConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagerConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerConfigurations"
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
	reqQP.Set("api-version", "2024-01-01-preview")
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagerConfigurationsClient) listHandleResponse(resp *http.Response) (WorkspaceManagerConfigurationsClientListResponse, error) {
	result := WorkspaceManagerConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerConfigurationList); err != nil {
		return WorkspaceManagerConfigurationsClientListResponse{}, err
	}
	return result, nil
}
