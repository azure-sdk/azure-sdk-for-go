//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// OperationStatusResultClient contains the methods for the OperationStatusResult group.
// Don't use this type directly, use NewOperationStatusResultClient() instead.
type OperationStatusResultClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationStatusResultClient creates a new instance of OperationStatusResultClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationStatusResultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationStatusResultClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationStatusResultClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the status of a specific operation in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - operationID - The ID of an ongoing async operation.
//   - options - OperationStatusResultClientGetOptions contains the optional parameters for the OperationStatusResultClient.Get
//     method.
func (client *OperationStatusResultClient) Get(ctx context.Context, resourceGroupName string, resourceName string, operationID string, options *OperationStatusResultClientGetOptions) (OperationStatusResultClientGetResponse, error) {
	var err error
	const operationName = "OperationStatusResultClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, operationID, options)
	if err != nil {
		return OperationStatusResultClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationStatusResultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationStatusResultClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OperationStatusResultClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, operationID string, options *OperationStatusResultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/operations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationStatusResultClient) getHandleResponse(resp *http.Response) (OperationStatusResultClientGetResponse, error) {
	result := OperationStatusResultClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatusResult); err != nil {
		return OperationStatusResultClientGetResponse{}, err
	}
	return result, nil
}

// GetByAgentPool - Get the status of a specific operation in the specified agent pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - agentPoolName - The name of the agent pool.
//   - operationID - The ID of an ongoing async operation.
//   - options - OperationStatusResultClientGetByAgentPoolOptions contains the optional parameters for the OperationStatusResultClient.GetByAgentPool
//     method.
func (client *OperationStatusResultClient) GetByAgentPool(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, operationID string, options *OperationStatusResultClientGetByAgentPoolOptions) (OperationStatusResultClientGetByAgentPoolResponse, error) {
	var err error
	const operationName = "OperationStatusResultClient.GetByAgentPool"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByAgentPoolCreateRequest(ctx, resourceGroupName, resourceName, agentPoolName, operationID, options)
	if err != nil {
		return OperationStatusResultClientGetByAgentPoolResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationStatusResultClientGetByAgentPoolResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OperationStatusResultClientGetByAgentPoolResponse{}, err
	}
	resp, err := client.getByAgentPoolHandleResponse(httpResp)
	return resp, err
}

// getByAgentPoolCreateRequest creates the GetByAgentPool request.
func (client *OperationStatusResultClient) getByAgentPoolCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, agentPoolName string, operationID string, options *OperationStatusResultClientGetByAgentPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/agentPools/{agentPoolName}/operations/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if agentPoolName == "" {
		return nil, errors.New("parameter agentPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentPoolName}", url.PathEscape(agentPoolName))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByAgentPoolHandleResponse handles the GetByAgentPool response.
func (client *OperationStatusResultClient) getByAgentPoolHandleResponse(resp *http.Response) (OperationStatusResultClientGetByAgentPoolResponse, error) {
	result := OperationStatusResultClientGetByAgentPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatusResult); err != nil {
		return OperationStatusResultClientGetByAgentPoolResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets a list of operations in the specified managedCluster
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - OperationStatusResultClientListOptions contains the optional parameters for the OperationStatusResultClient.NewListPager
//     method.
func (client *OperationStatusResultClient) NewListPager(resourceGroupName string, resourceName string, options *OperationStatusResultClientListOptions) *runtime.Pager[OperationStatusResultClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationStatusResultClientListResponse]{
		More: func(page OperationStatusResultClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationStatusResultClientListResponse) (OperationStatusResultClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OperationStatusResultClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return OperationStatusResultClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OperationStatusResultClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *OperationStatusResultClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/operations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationStatusResultClient) listHandleResponse(resp *http.Response) (OperationStatusResultClientListResponse, error) {
	result := OperationStatusResultClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationStatusResultList); err != nil {
		return OperationStatusResultClientListResponse{}, err
	}
	return result, nil
}
