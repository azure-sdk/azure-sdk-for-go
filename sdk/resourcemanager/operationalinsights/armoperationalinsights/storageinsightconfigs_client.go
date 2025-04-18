// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// StorageInsightConfigsClient contains the methods for the StorageInsightConfigs group.
// Don't use this type directly, use NewStorageInsightConfigsClient() instead.
type StorageInsightConfigsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageInsightConfigsClient creates a new instance of StorageInsightConfigsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageInsightConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageInsightConfigsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageInsightConfigsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a storage insight.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - storageInsightName - Name of the storageInsightsConfigs resource
//   - parameters - The parameters required to create or update a storage insight.
//   - options - StorageInsightConfigsClientCreateOrUpdateOptions contains the optional parameters for the StorageInsightConfigsClient.CreateOrUpdate
//     method.
func (client *StorageInsightConfigsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters StorageInsight, options *StorageInsightConfigsClientCreateOrUpdateOptions) (StorageInsightConfigsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "StorageInsightConfigsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, storageInsightName, parameters, options)
	if err != nil {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageInsightConfigsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, parameters StorageInsight, _ *StorageInsightConfigsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if storageInsightName == "" {
		return nil, errors.New("parameter storageInsightName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageInsightName}", url.PathEscape(storageInsightName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *StorageInsightConfigsClient) createOrUpdateHandleResponse(resp *http.Response) (StorageInsightConfigsClientCreateOrUpdateResponse, error) {
	result := StorageInsightConfigsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageInsight); err != nil {
		return StorageInsightConfigsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a storageInsightsConfigs resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - storageInsightName - Name of the storageInsightsConfigs resource
//   - options - StorageInsightConfigsClientDeleteOptions contains the optional parameters for the StorageInsightConfigsClient.Delete
//     method.
func (client *StorageInsightConfigsClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *StorageInsightConfigsClientDeleteOptions) (StorageInsightConfigsClientDeleteResponse, error) {
	var err error
	const operationName = "StorageInsightConfigsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, storageInsightName, options)
	if err != nil {
		return StorageInsightConfigsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageInsightConfigsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return StorageInsightConfigsClientDeleteResponse{}, err
	}
	return StorageInsightConfigsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageInsightConfigsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, _ *StorageInsightConfigsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if storageInsightName == "" {
		return nil, errors.New("parameter storageInsightName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageInsightName}", url.PathEscape(storageInsightName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a storage insight instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - storageInsightName - Name of the storageInsightsConfigs resource
//   - options - StorageInsightConfigsClientGetOptions contains the optional parameters for the StorageInsightConfigsClient.Get
//     method.
func (client *StorageInsightConfigsClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, options *StorageInsightConfigsClientGetOptions) (StorageInsightConfigsClientGetResponse, error) {
	var err error
	const operationName = "StorageInsightConfigsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, storageInsightName, options)
	if err != nil {
		return StorageInsightConfigsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageInsightConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageInsightConfigsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StorageInsightConfigsClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, storageInsightName string, _ *StorageInsightConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs/{storageInsightName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if storageInsightName == "" {
		return nil, errors.New("parameter storageInsightName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageInsightName}", url.PathEscape(storageInsightName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageInsightConfigsClient) getHandleResponse(resp *http.Response) (StorageInsightConfigsClientGetResponse, error) {
	result := StorageInsightConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageInsight); err != nil {
		return StorageInsightConfigsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWorkspacePager - Lists the storage insight instances within a workspace
//
// Generated from API version 2025-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - StorageInsightConfigsClientListByWorkspaceOptions contains the optional parameters for the StorageInsightConfigsClient.NewListByWorkspacePager
//     method.
func (client *StorageInsightConfigsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *StorageInsightConfigsClientListByWorkspaceOptions) *runtime.Pager[StorageInsightConfigsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageInsightConfigsClientListByWorkspaceResponse]{
		More: func(page StorageInsightConfigsClientListByWorkspaceResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageInsightConfigsClientListByWorkspaceResponse) (StorageInsightConfigsClientListByWorkspaceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StorageInsightConfigsClient.NewListByWorkspacePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.ODataNextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return StorageInsightConfigsClientListByWorkspaceResponse{}, err
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *StorageInsightConfigsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, _ *StorageInsightConfigsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/storageInsightConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *StorageInsightConfigsClient) listByWorkspaceHandleResponse(resp *http.Response) (StorageInsightConfigsClientListByWorkspaceResponse, error) {
	result := StorageInsightConfigsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageInsightListResult); err != nil {
		return StorageInsightConfigsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
