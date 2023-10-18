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

// WorkspaceManagerMembersClient contains the methods for the WorkspaceManagerMembers group.
// Don't use this type directly, use NewWorkspaceManagerMembersClient() instead.
type WorkspaceManagerMembersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagerMembersClient creates a new instance of WorkspaceManagerMembersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagerMembersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagerMembersClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceManagerMembersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagerMembersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a workspace manager member
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerMemberName - The name of the workspace manager member
//   - workspaceManagerMember - The workspace manager member object
//   - options - WorkspaceManagerMembersClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceManagerMembersClient.CreateOrUpdate
//     method.
func (client *WorkspaceManagerMembersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerMemberName string, workspaceManagerMember WorkspaceManagerMember, options *WorkspaceManagerMembersClientCreateOrUpdateOptions) (WorkspaceManagerMembersClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerMemberName, workspaceManagerMember, options)
	if err != nil {
		return WorkspaceManagerMembersClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerMembersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerMembersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagerMembersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerMemberName string, workspaceManagerMember WorkspaceManagerMember, options *WorkspaceManagerMembersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerMembers/{workspaceManagerMemberName}"
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
	if workspaceManagerMemberName == "" {
		return nil, errors.New("parameter workspaceManagerMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerMemberName}", url.PathEscape(workspaceManagerMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, workspaceManagerMember); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceManagerMembersClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceManagerMembersClientCreateOrUpdateResponse, error) {
	result := WorkspaceManagerMembersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerMember); err != nil {
		return WorkspaceManagerMembersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a workspace manager member
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerMemberName - The name of the workspace manager member
//   - options - WorkspaceManagerMembersClientDeleteOptions contains the optional parameters for the WorkspaceManagerMembersClient.Delete
//     method.
func (client *WorkspaceManagerMembersClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerMemberName string, options *WorkspaceManagerMembersClientDeleteOptions) (WorkspaceManagerMembersClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerMemberName, options)
	if err != nil {
		return WorkspaceManagerMembersClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerMembersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerMembersClientDeleteResponse{}, err
	}
	return WorkspaceManagerMembersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceManagerMembersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerMemberName string, options *WorkspaceManagerMembersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerMembers/{workspaceManagerMemberName}"
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
	if workspaceManagerMemberName == "" {
		return nil, errors.New("parameter workspaceManagerMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerMemberName}", url.PathEscape(workspaceManagerMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a workspace manager member
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - workspaceManagerMemberName - The name of the workspace manager member
//   - options - WorkspaceManagerMembersClientGetOptions contains the optional parameters for the WorkspaceManagerMembersClient.Get
//     method.
func (client *WorkspaceManagerMembersClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerMemberName string, options *WorkspaceManagerMembersClientGetOptions) (WorkspaceManagerMembersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, workspaceManagerMemberName, options)
	if err != nil {
		return WorkspaceManagerMembersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagerMembersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagerMembersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagerMembersClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspaceManagerMemberName string, options *WorkspaceManagerMembersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerMembers/{workspaceManagerMemberName}"
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
	if workspaceManagerMemberName == "" {
		return nil, errors.New("parameter workspaceManagerMemberName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceManagerMemberName}", url.PathEscape(workspaceManagerMemberName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagerMembersClient) getHandleResponse(resp *http.Response) (WorkspaceManagerMembersClientGetResponse, error) {
	result := WorkspaceManagerMembersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerMember); err != nil {
		return WorkspaceManagerMembersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all workspace manager members that exist for the given Sentinel workspace manager
//
// Generated from API version 2023-09-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagerMembersClientListOptions contains the optional parameters for the WorkspaceManagerMembersClient.NewListPager
//     method.
func (client *WorkspaceManagerMembersClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagerMembersClientListOptions) *runtime.Pager[WorkspaceManagerMembersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagerMembersClientListResponse]{
		More: func(page WorkspaceManagerMembersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagerMembersClientListResponse) (WorkspaceManagerMembersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceManagerMembersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceManagerMembersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceManagerMembersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagerMembersClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagerMembersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/workspaceManagerMembers/"
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
	reqQP.Set("api-version", "2023-09-01-preview")
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
func (client *WorkspaceManagerMembersClient) listHandleResponse(resp *http.Response) (WorkspaceManagerMembersClientListResponse, error) {
	result := WorkspaceManagerMembersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceManagerMembersList); err != nil {
		return WorkspaceManagerMembersClientListResponse{}, err
	}
	return result, nil
}
