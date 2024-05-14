//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// RouteMapsClient contains the methods for the RouteMaps group.
// Don't use this type directly, use NewRouteMapsClient() instead.
type RouteMapsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRouteMapsClient creates a new instance of RouteMapsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRouteMapsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RouteMapsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RouteMapsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a RouteMap if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The resource group name of the RouteMap's resource group.
//   - virtualHubName - The name of the VirtualHub containing the RouteMap.
//   - routeMapName - The name of the RouteMap.
//   - routeMapParameters - Parameters supplied to create or update a RouteMap.
//   - options - RouteMapsClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteMapsClient.BeginCreateOrUpdate
//     method.
func (client *RouteMapsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters RouteMap, options *RouteMapsClientBeginCreateOrUpdateOptions) (*runtime.Poller[RouteMapsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, routeMapName, routeMapParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RouteMapsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RouteMapsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a RouteMap if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *RouteMapsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters RouteMap, options *RouteMapsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RouteMapsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeMapName, routeMapParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteMapsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, routeMapParameters RouteMap, options *RouteMapsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeMapName == "" {
		return nil, errors.New("parameter routeMapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeMapName}", url.PathEscape(routeMapName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, routeMapParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a RouteMap.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The resource group name of the RouteMap's resource group.
//   - virtualHubName - The name of the VirtualHub containing the RouteMap.
//   - routeMapName - The name of the RouteMap.
//   - options - RouteMapsClientBeginDeleteOptions contains the optional parameters for the RouteMapsClient.BeginDelete method.
func (client *RouteMapsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientBeginDeleteOptions) (*runtime.Poller[RouteMapsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, routeMapName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RouteMapsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RouteMapsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a RouteMap.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
func (client *RouteMapsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RouteMapsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeMapName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RouteMapsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeMapName == "" {
		return nil, errors.New("parameter routeMapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeMapName}", url.PathEscape(routeMapName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a RouteMap.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The resource group name of the RouteMap's resource group.
//   - virtualHubName - The name of the VirtualHub containing the RouteMap.
//   - routeMapName - The name of the RouteMap.
//   - options - RouteMapsClientGetOptions contains the optional parameters for the RouteMapsClient.Get method.
func (client *RouteMapsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientGetOptions) (RouteMapsClientGetResponse, error) {
	var err error
	const operationName = "RouteMapsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, routeMapName, options)
	if err != nil {
		return RouteMapsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RouteMapsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RouteMapsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RouteMapsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeMapName string, options *RouteMapsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps/{routeMapName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	if routeMapName == "" {
		return nil, errors.New("parameter routeMapName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeMapName}", url.PathEscape(routeMapName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteMapsClient) getHandleResponse(resp *http.Response) (RouteMapsClientGetResponse, error) {
	result := RouteMapsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteMap); err != nil {
		return RouteMapsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all RouteMaps.
//
// Generated from API version 2023-11-01
//   - resourceGroupName - The resource group name of the RouteMap's resource group'.
//   - virtualHubName - The name of the VirtualHub containing the RouteMap.
//   - options - RouteMapsClientListOptions contains the optional parameters for the RouteMapsClient.NewListPager method.
func (client *RouteMapsClient) NewListPager(resourceGroupName string, virtualHubName string, options *RouteMapsClientListOptions) *runtime.Pager[RouteMapsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteMapsClientListResponse]{
		More: func(page RouteMapsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteMapsClientListResponse) (RouteMapsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RouteMapsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			}, nil)
			if err != nil {
				return RouteMapsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RouteMapsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *RouteMapsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/routeMaps"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RouteMapsClient) listHandleResponse(resp *http.Response) (RouteMapsClientListResponse, error) {
	result := RouteMapsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListRouteMapsResult); err != nil {
		return RouteMapsClientListResponse{}, err
	}
	return result, nil
}
