// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// PrivateEndpointConnectionsClient contains the methods for the PrivateEndpointConnections group.
// Don't use this type directly, use NewPrivateEndpointConnectionsClient() instead.
type PrivateEndpointConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsClient creates a new instance of PrivateEndpointConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// DeleteByHostPool - Remove a connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource.
//   - options - PrivateEndpointConnectionsClientDeleteByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.DeleteByHostPool
//     method.
func (client *PrivateEndpointConnectionsClient) DeleteByHostPool(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientDeleteByHostPoolOptions) (PrivateEndpointConnectionsClientDeleteByHostPoolResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.DeleteByHostPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteByHostPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteByHostPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientDeleteByHostPoolResponse{}, err
	}
	return PrivateEndpointConnectionsClientDeleteByHostPoolResponse{}, nil
}

// deleteByHostPoolCreateRequest creates the DeleteByHostPool request.
func (client *PrivateEndpointConnectionsClient) deleteByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, _ *PrivateEndpointConnectionsClientDeleteByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// DeleteByWorkspace - Remove a connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource.
//   - options - PrivateEndpointConnectionsClientDeleteByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.DeleteByWorkspace
//     method.
func (client *PrivateEndpointConnectionsClient) DeleteByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientDeleteByWorkspaceOptions) (PrivateEndpointConnectionsClientDeleteByWorkspaceResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.DeleteByWorkspace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteByWorkspaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientDeleteByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientDeleteByWorkspaceResponse{}, err
	}
	return PrivateEndpointConnectionsClientDeleteByWorkspaceResponse{}, nil
}

// deleteByWorkspaceCreateRequest creates the DeleteByWorkspace request.
func (client *PrivateEndpointConnectionsClient) deleteByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, _ *PrivateEndpointConnectionsClientDeleteByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetByHostPool - Get a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource.
//   - options - PrivateEndpointConnectionsClientGetByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.GetByHostPool
//     method.
func (client *PrivateEndpointConnectionsClient) GetByHostPool(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetByHostPoolOptions) (PrivateEndpointConnectionsClientGetByHostPoolResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.GetByHostPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetByHostPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetByHostPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientGetByHostPoolResponse{}, err
	}
	resp, err := client.getByHostPoolHandleResponse(httpResp)
	return resp, err
}

// getByHostPoolCreateRequest creates the GetByHostPool request.
func (client *PrivateEndpointConnectionsClient) getByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, _ *PrivateEndpointConnectionsClientGetByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByHostPoolHandleResponse handles the GetByHostPool response.
func (client *PrivateEndpointConnectionsClient) getByHostPoolHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetByHostPoolResponse, error) {
	result := PrivateEndpointConnectionsClientGetByHostPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionWithSystemData); err != nil {
		return PrivateEndpointConnectionsClientGetByHostPoolResponse{}, err
	}
	return result, nil
}

// GetByWorkspace - Get a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource.
//   - options - PrivateEndpointConnectionsClientGetByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.GetByWorkspace
//     method.
func (client *PrivateEndpointConnectionsClient) GetByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsClientGetByWorkspaceOptions) (PrivateEndpointConnectionsClientGetByWorkspaceResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.GetByWorkspace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsClientGetByWorkspaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientGetByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientGetByWorkspaceResponse{}, err
	}
	resp, err := client.getByWorkspaceHandleResponse(httpResp)
	return resp, err
}

// getByWorkspaceCreateRequest creates the GetByWorkspace request.
func (client *PrivateEndpointConnectionsClient) getByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, _ *PrivateEndpointConnectionsClientGetByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByWorkspaceHandleResponse handles the GetByWorkspace response.
func (client *PrivateEndpointConnectionsClient) getByWorkspaceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientGetByWorkspaceResponse, error) {
	result := PrivateEndpointConnectionsClientGetByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionWithSystemData); err != nil {
		return PrivateEndpointConnectionsClientGetByWorkspaceResponse{}, err
	}
	return result, nil
}

// NewListByHostPoolPager - List private endpoint connections associated with hostpool.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - PrivateEndpointConnectionsClientListByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByHostPoolPager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByHostPoolPager(resourceGroupName string, hostPoolName string, options *PrivateEndpointConnectionsClientListByHostPoolOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByHostPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByHostPoolResponse]{
		More: func(page PrivateEndpointConnectionsClientListByHostPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByHostPoolResponse) (PrivateEndpointConnectionsClientListByHostPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateEndpointConnectionsClient.NewListByHostPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			}, nil)
			if err != nil {
				return PrivateEndpointConnectionsClientListByHostPoolResponse{}, err
			}
			return client.listByHostPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHostPoolCreateRequest creates the ListByHostPool request.
func (client *PrivateEndpointConnectionsClient) listByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *PrivateEndpointConnectionsClientListByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHostPoolHandleResponse handles the ListByHostPool response.
func (client *PrivateEndpointConnectionsClient) listByHostPoolHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByHostPoolResponse, error) {
	result := PrivateEndpointConnectionsClientListByHostPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResultWithSystemData); err != nil {
		return PrivateEndpointConnectionsClientListByHostPoolResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - List private endpoint connections.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace
//   - options - PrivateEndpointConnectionsClientListByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.NewListByWorkspacePager
//     method.
func (client *PrivateEndpointConnectionsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *PrivateEndpointConnectionsClientListByWorkspaceOptions) *runtime.Pager[PrivateEndpointConnectionsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsClientListByWorkspaceResponse]{
		More: func(page PrivateEndpointConnectionsClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsClientListByWorkspaceResponse) (PrivateEndpointConnectionsClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateEndpointConnectionsClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return PrivateEndpointConnectionsClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *PrivateEndpointConnectionsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, _ *PrivateEndpointConnectionsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/workspaces/{workspaceName}/privateEndpointConnections"
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
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *PrivateEndpointConnectionsClient) listByWorkspaceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientListByWorkspaceResponse, error) {
	result := PrivateEndpointConnectionsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResultWithSystemData); err != nil {
		return PrivateEndpointConnectionsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}

// UpdateByHostPool - Approve or reject a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource.
//   - connection - Object containing the updated connection.
//   - options - PrivateEndpointConnectionsClientUpdateByHostPoolOptions contains the optional parameters for the PrivateEndpointConnectionsClient.UpdateByHostPool
//     method.
func (client *PrivateEndpointConnectionsClient) UpdateByHostPool(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, connection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientUpdateByHostPoolOptions) (PrivateEndpointConnectionsClientUpdateByHostPoolResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.UpdateByHostPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, privateEndpointConnectionName, connection, options)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateByHostPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateByHostPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientUpdateByHostPoolResponse{}, err
	}
	resp, err := client.updateByHostPoolHandleResponse(httpResp)
	return resp, err
}

// updateByHostPoolCreateRequest creates the UpdateByHostPool request.
func (client *PrivateEndpointConnectionsClient) updateByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, privateEndpointConnectionName string, connection PrivateEndpointConnection, _ *PrivateEndpointConnectionsClientUpdateByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connection); err != nil {
		return nil, err
	}
	return req, nil
}

// updateByHostPoolHandleResponse handles the UpdateByHostPool response.
func (client *PrivateEndpointConnectionsClient) updateByHostPoolHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientUpdateByHostPoolResponse, error) {
	result := PrivateEndpointConnectionsClientUpdateByHostPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionWithSystemData); err != nil {
		return PrivateEndpointConnectionsClientUpdateByHostPoolResponse{}, err
	}
	return result, nil
}

// UpdateByWorkspace - Approve or reject a private endpoint connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace
//   - privateEndpointConnectionName - The name of the private endpoint connection associated with the Azure resource.
//   - connection - Object containing the updated connection.
//   - options - PrivateEndpointConnectionsClientUpdateByWorkspaceOptions contains the optional parameters for the PrivateEndpointConnectionsClient.UpdateByWorkspace
//     method.
func (client *PrivateEndpointConnectionsClient) UpdateByWorkspace(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, connection PrivateEndpointConnection, options *PrivateEndpointConnectionsClientUpdateByWorkspaceOptions) (PrivateEndpointConnectionsClientUpdateByWorkspaceResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionsClient.UpdateByWorkspace"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, privateEndpointConnectionName, connection, options)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateByWorkspaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsClientUpdateByWorkspaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionsClientUpdateByWorkspaceResponse{}, err
	}
	resp, err := client.updateByWorkspaceHandleResponse(httpResp)
	return resp, err
}

// updateByWorkspaceCreateRequest creates the UpdateByWorkspace request.
func (client *PrivateEndpointConnectionsClient) updateByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, privateEndpointConnectionName string, connection PrivateEndpointConnection, _ *PrivateEndpointConnectionsClientUpdateByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/workspaces/{workspaceName}/privateEndpointConnections/{privateEndpointConnectionName}"
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
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connection); err != nil {
		return nil, err
	}
	return req, nil
}

// updateByWorkspaceHandleResponse handles the UpdateByWorkspace response.
func (client *PrivateEndpointConnectionsClient) updateByWorkspaceHandleResponse(resp *http.Response) (PrivateEndpointConnectionsClientUpdateByWorkspaceResponse, error) {
	result := PrivateEndpointConnectionsClientUpdateByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionWithSystemData); err != nil {
		return PrivateEndpointConnectionsClientUpdateByWorkspaceResponse{}, err
	}
	return result, nil
}
