//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient contains the methods for the WorkspaceManagedSQLServerExtendedBlobAuditingPolicies group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient() instead.
type WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient creates a new instance of WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or Update a workspace managed sql server's extended blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - blobAuditingPolicyName - The name of the blob auditing policy.
//   - parameters - Properties of extended blob auditing policy.
//   - options - WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters
//     for the WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.BeginCreateOrUpdate method.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ExtendedServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or Update a workspace managed sql server's extended blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ExtendedServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, parameters ExtendedServerBlobAuditingPolicy, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings/{blobAuditingPolicyName}"
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
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a workspace SQL server's extended blob auditing policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - blobAuditingPolicyName - The name of the blob auditing policy.
//   - options - WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetOptions contains the optional parameters for the
//     WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.Get method.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetOptions) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, blobAuditingPolicyName, options)
	if err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, blobAuditingPolicyName BlobAuditingPolicyName, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings/{blobAuditingPolicyName}"
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
	if blobAuditingPolicyName == "" {
		return nil, errors.New("parameter blobAuditingPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{blobAuditingPolicyName}", url.PathEscape(string(blobAuditingPolicyName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse, error) {
	result := WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedServerBlobAuditingPolicy); err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - List workspace managed sql server's extended blob auditing policies.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceOptions contains the optional parameters
//     for the WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient.NewListByWorkspacePager method.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceOptions) *runtime.Pager[WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse]{
		More: func(page WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/extendedAuditingSettings"
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
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClient) listByWorkspaceHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse, error) {
	result := WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExtendedServerBlobAuditingPolicyListResult); err != nil {
		return WorkspaceManagedSQLServerExtendedBlobAuditingPoliciesClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
