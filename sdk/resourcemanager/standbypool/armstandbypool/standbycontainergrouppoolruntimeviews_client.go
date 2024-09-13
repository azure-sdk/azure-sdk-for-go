//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstandbypool

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

// StandbyContainerGroupPoolRuntimeViewsClient contains the methods for the StandbyContainerGroupPoolRuntimeViews group.
// Don't use this type directly, use NewStandbyContainerGroupPoolRuntimeViewsClient() instead.
type StandbyContainerGroupPoolRuntimeViewsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStandbyContainerGroupPoolRuntimeViewsClient creates a new instance of StandbyContainerGroupPoolRuntimeViewsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStandbyContainerGroupPoolRuntimeViewsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StandbyContainerGroupPoolRuntimeViewsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StandbyContainerGroupPoolRuntimeViewsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a StandbyContainerGroupPoolRuntimeViewResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - standbyContainerGroupPoolName - Name of the standby container group pool
//   - runtimeView - The unique identifier for the runtime view. The input string should be the word 'latest', which will get
//     the latest runtime view of the pool, otherwise the request will fail with NotFound exception.
//   - options - StandbyContainerGroupPoolRuntimeViewsClientGetOptions contains the optional parameters for the StandbyContainerGroupPoolRuntimeViewsClient.Get
//     method.
func (client *StandbyContainerGroupPoolRuntimeViewsClient) Get(ctx context.Context, resourceGroupName string, standbyContainerGroupPoolName string, runtimeView string, options *StandbyContainerGroupPoolRuntimeViewsClientGetOptions) (StandbyContainerGroupPoolRuntimeViewsClientGetResponse, error) {
	var err error
	const operationName = "StandbyContainerGroupPoolRuntimeViewsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, standbyContainerGroupPoolName, runtimeView, options)
	if err != nil {
		return StandbyContainerGroupPoolRuntimeViewsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StandbyContainerGroupPoolRuntimeViewsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StandbyContainerGroupPoolRuntimeViewsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StandbyContainerGroupPoolRuntimeViewsClient) getCreateRequest(ctx context.Context, resourceGroupName string, standbyContainerGroupPoolName string, runtimeView string, options *StandbyContainerGroupPoolRuntimeViewsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StandbyPool/standbyContainerGroupPools/{standbyContainerGroupPoolName}/runtimeViews/{runtimeView}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standbyContainerGroupPoolName == "" {
		return nil, errors.New("parameter standbyContainerGroupPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standbyContainerGroupPoolName}", url.PathEscape(standbyContainerGroupPoolName))
	if runtimeView == "" {
		return nil, errors.New("parameter runtimeView cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{runtimeView}", url.PathEscape(runtimeView))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StandbyContainerGroupPoolRuntimeViewsClient) getHandleResponse(resp *http.Response) (StandbyContainerGroupPoolRuntimeViewsClientGetResponse, error) {
	result := StandbyContainerGroupPoolRuntimeViewsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StandbyContainerGroupPoolRuntimeViewResource); err != nil {
		return StandbyContainerGroupPoolRuntimeViewsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByStandbyPoolPager - List StandbyContainerGroupPoolRuntimeViewResource resources by StandbyContainerGroupPoolResource
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - standbyContainerGroupPoolName - Name of the standby container group pool
//   - options - StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolOptions contains the optional parameters for the
//     StandbyContainerGroupPoolRuntimeViewsClient.NewListByStandbyPoolPager method.
func (client *StandbyContainerGroupPoolRuntimeViewsClient) NewListByStandbyPoolPager(resourceGroupName string, standbyContainerGroupPoolName string, options *StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolOptions) *runtime.Pager[StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse]{
		More: func(page StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse) (StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "StandbyContainerGroupPoolRuntimeViewsClient.NewListByStandbyPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByStandbyPoolCreateRequest(ctx, resourceGroupName, standbyContainerGroupPoolName, options)
			}, nil)
			if err != nil {
				return StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse{}, err
			}
			return client.listByStandbyPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByStandbyPoolCreateRequest creates the ListByStandbyPool request.
func (client *StandbyContainerGroupPoolRuntimeViewsClient) listByStandbyPoolCreateRequest(ctx context.Context, resourceGroupName string, standbyContainerGroupPoolName string, options *StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.StandbyPool/standbyContainerGroupPools/{standbyContainerGroupPoolName}/runtimeViews"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if standbyContainerGroupPoolName == "" {
		return nil, errors.New("parameter standbyContainerGroupPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{standbyContainerGroupPoolName}", url.PathEscape(standbyContainerGroupPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByStandbyPoolHandleResponse handles the ListByStandbyPool response.
func (client *StandbyContainerGroupPoolRuntimeViewsClient) listByStandbyPoolHandleResponse(resp *http.Response) (StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse, error) {
	result := StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StandbyContainerGroupPoolRuntimeViewResourceListResult); err != nil {
		return StandbyContainerGroupPoolRuntimeViewsClientListByStandbyPoolResponse{}, err
	}
	return result, nil
}
