// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

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

// PatchSchedulesClient contains the methods for the PatchSchedules group.
// Don't use this type directly, use NewPatchSchedulesClient() instead.
type PatchSchedulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPatchSchedulesClient creates a new instance of PatchSchedulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPatchSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PatchSchedulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PatchSchedulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or replace the patching schedule for Redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the Redis cache.
//   - defaultParam - Default string modeled as parameter for auto generation to work correctly.
//   - parameters - Parameters to set the patching schedule for Redis cache.
//   - options - PatchSchedulesClientCreateOrUpdateOptions contains the optional parameters for the PatchSchedulesClient.CreateOrUpdate
//     method.
func (client *PatchSchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, parameters PatchSchedule, options *PatchSchedulesClientCreateOrUpdateOptions) (PatchSchedulesClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "PatchSchedulesClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, name, defaultParam, parameters, options)
	if err != nil {
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *PatchSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, parameters PatchSchedule, _ *PatchSchedulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *PatchSchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (PatchSchedulesClientCreateOrUpdateResponse, error) {
	result := PatchSchedulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchSchedule); err != nil {
		return PatchSchedulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the patching schedule of a redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the redis cache.
//   - defaultParam - Default string modeled as parameter for auto generation to work correctly.
//   - options - PatchSchedulesClientDeleteOptions contains the optional parameters for the PatchSchedulesClient.Delete method.
func (client *PatchSchedulesClient) Delete(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesClientDeleteOptions) (PatchSchedulesClientDeleteResponse, error) {
	var err error
	const operationName = "PatchSchedulesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, name, defaultParam, options)
	if err != nil {
		return PatchSchedulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PatchSchedulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PatchSchedulesClientDeleteResponse{}, err
	}
	return PatchSchedulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PatchSchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, _ *PatchSchedulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the patching schedule of a redis cache.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - name - The name of the redis cache.
//   - defaultParam - Default string modeled as parameter for auto generation to work correctly.
//   - options - PatchSchedulesClientGetOptions contains the optional parameters for the PatchSchedulesClient.Get method.
func (client *PatchSchedulesClient) Get(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, options *PatchSchedulesClientGetOptions) (PatchSchedulesClientGetResponse, error) {
	var err error
	const operationName = "PatchSchedulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, name, defaultParam, options)
	if err != nil {
		return PatchSchedulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PatchSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PatchSchedulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PatchSchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, name string, defaultParam DefaultName, _ *PatchSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/{default}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	if defaultParam == "" {
		return nil, errors.New("parameter defaultParam cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{default}", url.PathEscape(string(defaultParam)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PatchSchedulesClient) getHandleResponse(resp *http.Response) (PatchSchedulesClientGetResponse, error) {
	result := PatchSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchSchedule); err != nil {
		return PatchSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByRedisResourcePager - Gets all patch schedules in the specified redis cache (there is only one).
//
// Generated from API version 2024-11-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cacheName - The name of the Redis cache.
//   - options - PatchSchedulesClientListByRedisResourceOptions contains the optional parameters for the PatchSchedulesClient.NewListByRedisResourcePager
//     method.
func (client *PatchSchedulesClient) NewListByRedisResourcePager(resourceGroupName string, cacheName string, options *PatchSchedulesClientListByRedisResourceOptions) *runtime.Pager[PatchSchedulesClientListByRedisResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[PatchSchedulesClientListByRedisResourceResponse]{
		More: func(page PatchSchedulesClientListByRedisResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PatchSchedulesClientListByRedisResourceResponse) (PatchSchedulesClientListByRedisResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PatchSchedulesClient.NewListByRedisResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByRedisResourceCreateRequest(ctx, resourceGroupName, cacheName, options)
			}, nil)
			if err != nil {
				return PatchSchedulesClientListByRedisResourceResponse{}, err
			}
			return client.listByRedisResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByRedisResourceCreateRequest creates the ListByRedisResource request.
func (client *PatchSchedulesClient) listByRedisResourceCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, _ *PatchSchedulesClientListByRedisResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/patchSchedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRedisResourceHandleResponse handles the ListByRedisResource response.
func (client *PatchSchedulesClient) listByRedisResourceHandleResponse(resp *http.Response) (PatchSchedulesClientListByRedisResourceResponse, error) {
	result := PatchSchedulesClientListByRedisResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PatchScheduleListResult); err != nil {
		return PatchSchedulesClientListByRedisResourceResponse{}, err
	}
	return result, nil
}
