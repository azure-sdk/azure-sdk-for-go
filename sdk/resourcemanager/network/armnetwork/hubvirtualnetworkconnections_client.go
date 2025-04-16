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

// HubVirtualNetworkConnectionsClient contains the methods for the HubVirtualNetworkConnections group.
// Don't use this type directly, use NewHubVirtualNetworkConnectionsClient() instead.
type HubVirtualNetworkConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHubVirtualNetworkConnectionsClient creates a new instance of HubVirtualNetworkConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHubVirtualNetworkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HubVirtualNetworkConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HubVirtualNetworkConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a hub virtual network connection if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The resource group name of the HubVirtualNetworkConnection.
//   - virtualHubName - The name of the VirtualHub.
//   - connectionName - The name of the HubVirtualNetworkConnection.
//   - hubVirtualNetworkConnectionParameters - Parameters supplied to create or update a hub virtual network connection.
//   - options - HubVirtualNetworkConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the HubVirtualNetworkConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *HubVirtualNetworkConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection, options *HubVirtualNetworkConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[HubVirtualNetworkConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualHubName, connectionName, hubVirtualNetworkConnectionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HubVirtualNetworkConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HubVirtualNetworkConnectionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a hub virtual network connection if it doesn't exist else updates the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
func (client *HubVirtualNetworkConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection, options *HubVirtualNetworkConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "HubVirtualNetworkConnectionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, hubVirtualNetworkConnectionParameters, options)
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
func (client *HubVirtualNetworkConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, hubVirtualNetworkConnectionParameters HubVirtualNetworkConnection, _ *HubVirtualNetworkConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
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
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, hubVirtualNetworkConnectionParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a HubVirtualNetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - connectionName - The name of the HubVirtualNetworkConnection.
//   - options - HubVirtualNetworkConnectionsClientBeginDeleteOptions contains the optional parameters for the HubVirtualNetworkConnectionsClient.BeginDelete
//     method.
func (client *HubVirtualNetworkConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsClientBeginDeleteOptions) (*runtime.Poller[HubVirtualNetworkConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualHubName, connectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HubVirtualNetworkConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HubVirtualNetworkConnectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a HubVirtualNetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
func (client *HubVirtualNetworkConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "HubVirtualNetworkConnectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
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
func (client *HubVirtualNetworkConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, _ *HubVirtualNetworkConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
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
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves the details of a HubVirtualNetworkConnection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - connectionName - The name of the vpn connection.
//   - options - HubVirtualNetworkConnectionsClientGetOptions contains the optional parameters for the HubVirtualNetworkConnectionsClient.Get
//     method.
func (client *HubVirtualNetworkConnectionsClient) Get(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, options *HubVirtualNetworkConnectionsClientGetOptions) (HubVirtualNetworkConnectionsClientGetResponse, error) {
	var err error
	const operationName = "HubVirtualNetworkConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualHubName, connectionName, options)
	if err != nil {
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HubVirtualNetworkConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, connectionName string, _ *HubVirtualNetworkConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections/{connectionName}"
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
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HubVirtualNetworkConnectionsClient) getHandleResponse(resp *http.Response) (HubVirtualNetworkConnectionsClientGetResponse, error) {
	result := HubVirtualNetworkConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HubVirtualNetworkConnection); err != nil {
		return HubVirtualNetworkConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves the details of all HubVirtualNetworkConnections.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The resource group name of the VirtualHub.
//   - virtualHubName - The name of the VirtualHub.
//   - options - HubVirtualNetworkConnectionsClientListOptions contains the optional parameters for the HubVirtualNetworkConnectionsClient.NewListPager
//     method.
func (client *HubVirtualNetworkConnectionsClient) NewListPager(resourceGroupName string, virtualHubName string, options *HubVirtualNetworkConnectionsClientListOptions) *runtime.Pager[HubVirtualNetworkConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HubVirtualNetworkConnectionsClientListResponse]{
		More: func(page HubVirtualNetworkConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HubVirtualNetworkConnectionsClientListResponse) (HubVirtualNetworkConnectionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HubVirtualNetworkConnectionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			}, nil)
			if err != nil {
				return HubVirtualNetworkConnectionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HubVirtualNetworkConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, _ *HubVirtualNetworkConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/hubVirtualNetworkConnections"
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
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HubVirtualNetworkConnectionsClient) listHandleResponse(resp *http.Response) (HubVirtualNetworkConnectionsClientListResponse, error) {
	result := HubVirtualNetworkConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListHubVirtualNetworkConnectionsResult); err != nil {
		return HubVirtualNetworkConnectionsClientListResponse{}, err
	}
	return result, nil
}
