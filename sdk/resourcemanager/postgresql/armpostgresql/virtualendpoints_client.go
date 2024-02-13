//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

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

// VirtualEndpointsClient contains the methods for the VirtualEndpoints group.
// Don't use this type directly, use NewVirtualEndpointsClient() instead.
type VirtualEndpointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualEndpointsClient creates a new instance of VirtualEndpointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualEndpointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualEndpointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualEndpointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a new virtual endpoint for PostgreSQL flexible server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualEndpointName - The name of the virtual endpoint.
//   - parameters - The required parameters for creating or updating virtual endpoints.
//   - options - VirtualEndpointsClientBeginCreateOptions contains the optional parameters for the VirtualEndpointsClient.BeginCreate
//     method.
func (client *VirtualEndpointsClient) BeginCreate(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters VirtualEndpointResource, options *VirtualEndpointsClientBeginCreateOptions) (*runtime.Poller[VirtualEndpointsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, serverName, virtualEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualEndpointsClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualEndpointsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates a new virtual endpoint for PostgreSQL flexible server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *VirtualEndpointsClient) create(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters VirtualEndpointResource, options *VirtualEndpointsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualEndpointsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, serverName, virtualEndpointName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *VirtualEndpointsClient) createCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters VirtualEndpointResource, options *VirtualEndpointsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/virtualendpoints/{virtualEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualEndpointName == "" {
		return nil, errors.New("parameter virtualEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualEndpointName}", url.PathEscape(virtualEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a virtual endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualEndpointName - The name of the virtual endpoint.
//   - options - VirtualEndpointsClientBeginDeleteOptions contains the optional parameters for the VirtualEndpointsClient.BeginDelete
//     method.
func (client *VirtualEndpointsClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *VirtualEndpointsClientBeginDeleteOptions) (*runtime.Poller[VirtualEndpointsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, virtualEndpointName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualEndpointsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualEndpointsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a virtual endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *VirtualEndpointsClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *VirtualEndpointsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualEndpointsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, virtualEndpointName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VirtualEndpointsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *VirtualEndpointsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/virtualendpoints/{virtualEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualEndpointName == "" {
		return nil, errors.New("parameter virtualEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualEndpointName}", url.PathEscape(virtualEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about a virtual endpoint.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualEndpointName - The name of the virtual endpoint.
//   - options - VirtualEndpointsClientGetOptions contains the optional parameters for the VirtualEndpointsClient.Get method.
func (client *VirtualEndpointsClient) Get(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *VirtualEndpointsClientGetOptions) (VirtualEndpointsClientGetResponse, error) {
	var err error
	const operationName = "VirtualEndpointsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, virtualEndpointName, options)
	if err != nil {
		return VirtualEndpointsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualEndpointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualEndpointsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualEndpointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, options *VirtualEndpointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/virtualendpoints/{virtualEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualEndpointName == "" {
		return nil, errors.New("parameter virtualEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualEndpointName}", url.PathEscape(virtualEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualEndpointsClient) getHandleResponse(resp *http.Response) (VirtualEndpointsClientGetResponse, error) {
	result := VirtualEndpointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualEndpointResource); err != nil {
		return VirtualEndpointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - List all the servers in a given resource group.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - VirtualEndpointsClientListByServerOptions contains the optional parameters for the VirtualEndpointsClient.NewListByServerPager
//     method.
func (client *VirtualEndpointsClient) NewListByServerPager(resourceGroupName string, serverName string, options *VirtualEndpointsClientListByServerOptions) *runtime.Pager[VirtualEndpointsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualEndpointsClientListByServerResponse]{
		More: func(page VirtualEndpointsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualEndpointsClientListByServerResponse) (VirtualEndpointsClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VirtualEndpointsClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return VirtualEndpointsClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *VirtualEndpointsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *VirtualEndpointsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/virtualendpoints"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *VirtualEndpointsClient) listByServerHandleResponse(resp *http.Response) (VirtualEndpointsClientListByServerResponse, error) {
	result := VirtualEndpointsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualEndpointsListResult); err != nil {
		return VirtualEndpointsClientListByServerResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates an existing virtual endpoint. The request body can contain one to many of the properties present
// in the normal virtual endpoint definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - virtualEndpointName - The name of the virtual endpoint.
//   - parameters - The required parameters for updating a server.
//   - options - VirtualEndpointsClientBeginUpdateOptions contains the optional parameters for the VirtualEndpointsClient.BeginUpdate
//     method.
func (client *VirtualEndpointsClient) BeginUpdate(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters VirtualEndpointResourceForPatch, options *VirtualEndpointsClientBeginUpdateOptions) (*runtime.Poller[VirtualEndpointsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, serverName, virtualEndpointName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualEndpointsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VirtualEndpointsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates an existing virtual endpoint. The request body can contain one to many of the properties present in the
// normal virtual endpoint definition.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
func (client *VirtualEndpointsClient) update(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters VirtualEndpointResourceForPatch, options *VirtualEndpointsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VirtualEndpointsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, virtualEndpointName, parameters, options)
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

// updateCreateRequest creates the Update request.
func (client *VirtualEndpointsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, virtualEndpointName string, parameters VirtualEndpointResourceForPatch, options *VirtualEndpointsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/virtualendpoints/{virtualEndpointName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if virtualEndpointName == "" {
		return nil, errors.New("parameter virtualEndpointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualEndpointName}", url.PathEscape(virtualEndpointName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
