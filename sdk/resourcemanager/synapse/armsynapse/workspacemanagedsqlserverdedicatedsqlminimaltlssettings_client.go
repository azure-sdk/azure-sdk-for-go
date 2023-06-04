//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient contains the methods for the WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettings group.
// Don't use this type directly, use NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient() instead.
type WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient creates a new instance of WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get workspace managed sql server's minimal tls settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dedicatedSQLminimalTLSSettingsName - The name of the dedicated sql minimal tls settings.
//   - options - WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetOptions contains the optional parameters for
//     the WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.Get method.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName string, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetOptions) (WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, dedicatedSQLminimalTLSSettingsName, options)
	if err != nil {
		return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName string, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings/{dedicatedSQLminimalTlsSettingsName}"
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
	if dedicatedSQLminimalTLSSettingsName == "" {
		return nil, errors.New("parameter dedicatedSQLminimalTLSSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedSQLminimalTlsSettingsName}", url.PathEscape(dedicatedSQLminimalTLSSettingsName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) getHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse, error) {
	result := WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedSQLminimalTLSSettings); err != nil {
		return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List workspace managed sql server's minimal tls settings.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListOptions contains the optional parameters for
//     the WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.NewListPager method.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) NewListPager(resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListOptions) *runtime.Pager[WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse]{
		More: func(page WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse) (WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings"
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
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) listHandleResponse(resp *http.Response) (WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse, error) {
	result := WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedSQLminimalTLSSettingsListResult); err != nil {
		return WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update workspace managed sql server's minimal tls settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - dedicatedSQLminimalTLSSettingsName - The name of the dedicated sql minimal tls settings.
//   - parameters - minimal tls settings
//   - options - WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientBeginUpdateOptions contains the optional parameters
//     for the WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient.BeginUpdate method.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName DedicatedSQLMinimalTLSSettingsName, parameters DedicatedSQLminimalTLSSettings, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientBeginUpdateOptions) (*runtime.Poller[WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, workspaceName, dedicatedSQLminimalTLSSettingsName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update workspace managed sql server's minimal tls settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) update(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName DedicatedSQLMinimalTLSSettingsName, parameters DedicatedSQLminimalTLSSettings, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, dedicatedSQLminimalTLSSettingsName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, dedicatedSQLminimalTLSSettingsName DedicatedSQLMinimalTLSSettingsName, parameters DedicatedSQLminimalTLSSettings, options *WorkspaceManagedSQLServerDedicatedSQLMinimalTLSSettingsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/dedicatedSQLminimalTlsSettings/{dedicatedSQLminimalTlsSettingsName}"
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
	if dedicatedSQLminimalTLSSettingsName == "" {
		return nil, errors.New("parameter dedicatedSQLminimalTLSSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedSQLminimalTlsSettingsName}", url.PathEscape(string(dedicatedSQLminimalTLSSettingsName)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
