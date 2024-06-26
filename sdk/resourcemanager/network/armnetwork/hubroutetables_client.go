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

// HubRouteTablesClient contains the methods for the HubRouteTables group.
// Don't use this type directly, use NewHubRouteTablesClient() instead.
type HubRouteTablesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHubRouteTablesClient creates a new instance of HubRouteTablesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHubRouteTablesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HubRouteTablesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HubRouteTablesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - routeTableName - The name of the RouteTable.
//   - routeTableParameters - Parameters supplied to create or update RouteTable.
//   - options - HubRouteTablesClientBeginCreateOrUpdateOptions contains the optional parameters for the HubRouteTablesClient.BeginCreateOrUpdate
//     method.
func (client *HubRouteTablesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesClientBeginCreateOrUpdateOptions) (*runtime.Poller[HubRouteTablesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, routeTableName, routeTableParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HubRouteTablesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HubRouteTablesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a RouteTable resource if it doesn't exist else updates the existing RouteTable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *HubRouteTablesClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "HubRouteTablesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, routeTableParameters, options)
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
func (client *HubRouteTablesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, routeTableParameters HubRouteTable, options *HubRouteTablesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}"
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
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, routeTableParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a RouteTable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The resource group name of the RouteTable.
//   - virtualHubName - The name of the VirtualHub.
//   - routeTableName - The name of the RouteTable.
//   - options - HubRouteTablesClientBeginDeleteOptions contains the optional parameters for the HubRouteTablesClient.BeginDelete
//     method.
func (client *HubRouteTablesClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesClientBeginDeleteOptions) (*runtime.Poller[HubRouteTablesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, routeTableName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HubRouteTablesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HubRouteTablesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a RouteTable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *HubRouteTablesClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "HubRouteTablesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
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
func (client *HubRouteTablesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}"
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
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a RouteTable.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - routeTableName - The name of the RouteTable.
//   - options - HubRouteTablesClientGetOptions contains the optional parameters for the HubRouteTablesClient.Get method.
func (client *HubRouteTablesClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesClientGetOptions) (HubRouteTablesClientGetResponse, error) {
	var err error
	const operationName = "HubRouteTablesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, routeTableName, options)
	if err != nil {
		return HubRouteTablesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HubRouteTablesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HubRouteTablesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HubRouteTablesClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, routeTableName string, options *HubRouteTablesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables/{routeTableName}"
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
	if routeTableName == "" {
		return nil, errors.New("parameter routeTableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeTableName}", url.PathEscape(routeTableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HubRouteTablesClient) getHandleResponse(resp *http.Response) (HubRouteTablesClientGetResponse, error) {
	result := HubRouteTablesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubRouteTable); err != nil {
		return HubRouteTablesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all RouteTables.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - options - HubRouteTablesClientListOptions contains the optional parameters for the HubRouteTablesClient.NewListPager method.
func (client *HubRouteTablesClient) NewListPager(resourceGroupName string, virtualHubName string, options *HubRouteTablesClientListOptions) *runtime.Pager[HubRouteTablesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HubRouteTablesClientListResponse]{
		More: func(page HubRouteTablesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HubRouteTablesClientListResponse) (HubRouteTablesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HubRouteTablesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			}, nil)
			if err != nil {
				return HubRouteTablesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HubRouteTablesClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *HubRouteTablesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubRouteTables"
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
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HubRouteTablesClient) listHandleResponse(resp *http.Response) (HubRouteTablesClientListResponse, error) {
	result := HubRouteTablesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListHubRouteTablesResult); err != nil {
		return HubRouteTablesClientListResponse{}, err
	}
	return result, nil
}
