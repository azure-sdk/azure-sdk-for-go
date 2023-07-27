//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerDNSAliasesClient contains the methods for the ServerDNSAliases group.
// Don't use this type directly, use NewServerDNSAliasesClient() instead.
type ServerDNSAliasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerDNSAliasesClient creates a new instance of ServerDNSAliasesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerDNSAliasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerDNSAliasesClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerDNSAliasesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerDNSAliasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginAcquire - Acquires server DNS alias from another server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server that the alias is pointing to.
//   - dnsAliasName - The name of the server dns alias.
//   - options - ServerDNSAliasesClientBeginAcquireOptions contains the optional parameters for the ServerDNSAliasesClient.BeginAcquire
//     method.
func (client *ServerDNSAliasesClient) BeginAcquire(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, parameters ServerDNSAliasAcquisition, options *ServerDNSAliasesClientBeginAcquireOptions) (*runtime.Poller[ServerDNSAliasesClientAcquireResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.acquire(ctx, resourceGroupName, serverName, dnsAliasName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerDNSAliasesClientAcquireResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerDNSAliasesClientAcquireResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Acquire - Acquires server DNS alias from another server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ServerDNSAliasesClient) acquire(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, parameters ServerDNSAliasAcquisition, options *ServerDNSAliasesClientBeginAcquireOptions) (*http.Response, error) {
	var err error
	req, err := client.acquireCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, parameters, options)
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

// acquireCreateRequest creates the Acquire request.
func (client *ServerDNSAliasesClient) acquireCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, parameters ServerDNSAliasAcquisition, options *ServerDNSAliasesClientBeginAcquireOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}/acquire"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Creates a server DNS alias.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server that the alias is pointing to.
//   - dnsAliasName - The name of the server dns alias.
//   - options - ServerDNSAliasesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerDNSAliasesClient.BeginCreateOrUpdate
//     method.
func (client *ServerDNSAliasesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerDNSAliasesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, dnsAliasName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerDNSAliasesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerDNSAliasesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a server DNS alias.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ServerDNSAliasesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerDNSAliasesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginDelete - Deletes the server DNS alias with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server that the alias is pointing to.
//   - dnsAliasName - The name of the server dns alias.
//   - options - ServerDNSAliasesClientBeginDeleteOptions contains the optional parameters for the ServerDNSAliasesClient.BeginDelete
//     method.
func (client *ServerDNSAliasesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginDeleteOptions) (*runtime.Poller[ServerDNSAliasesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, dnsAliasName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ServerDNSAliasesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ServerDNSAliasesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the server DNS alias with the given name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *ServerDNSAliasesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, options)
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
func (client *ServerDNSAliasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a server DNS alias.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server that the alias is pointing to.
//   - dnsAliasName - The name of the server dns alias.
//   - options - ServerDNSAliasesClientGetOptions contains the optional parameters for the ServerDNSAliasesClient.Get method.
func (client *ServerDNSAliasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientGetOptions) (ServerDNSAliasesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, dnsAliasName, options)
	if err != nil {
		return ServerDNSAliasesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerDNSAliasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerDNSAliasesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerDNSAliasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, dnsAliasName string, options *ServerDNSAliasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases/{dnsAliasName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if dnsAliasName == "" {
		return nil, errors.New("parameter dnsAliasName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsAliasName}", url.PathEscape(dnsAliasName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerDNSAliasesClient) getHandleResponse(resp *http.Response) (ServerDNSAliasesClientGetResponse, error) {
	result := ServerDNSAliasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDNSAlias); err != nil {
		return ServerDNSAliasesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of server DNS aliases for a server.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server that the alias is pointing to.
//   - options - ServerDNSAliasesClientListByServerOptions contains the optional parameters for the ServerDNSAliasesClient.NewListByServerPager
//     method.
func (client *ServerDNSAliasesClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerDNSAliasesClientListByServerOptions) *runtime.Pager[ServerDNSAliasesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerDNSAliasesClientListByServerResponse]{
		More: func(page ServerDNSAliasesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerDNSAliasesClientListByServerResponse) (ServerDNSAliasesClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerDNSAliasesClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ServerDNSAliasesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerDNSAliasesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerDNSAliasesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerDNSAliasesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/dnsAliases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerDNSAliasesClient) listByServerHandleResponse(resp *http.Response) (ServerDNSAliasesClientListByServerResponse, error) {
	result := ServerDNSAliasesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDNSAliasListResult); err != nil {
		return ServerDNSAliasesClientListByServerResponse{}, err
	}
	return result, nil
}
