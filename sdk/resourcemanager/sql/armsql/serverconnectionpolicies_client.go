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

// ServerConnectionPoliciesClient contains the methods for the ServerConnectionPolicies group.
// Don't use this type directly, use NewServerConnectionPoliciesClient() instead.
type ServerConnectionPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerConnectionPoliciesClient creates a new instance of ServerConnectionPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerConnectionPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerConnectionPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerConnectionPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Updates a server connection policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - connectionPolicyName - The name of the connection policy.
//   - parameters - The required parameters for updating a server connection policy.
//   - options - ServerConnectionPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerConnectionPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *ServerConnectionPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerConnectionPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, connectionPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerConnectionPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerConnectionPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Updates a server connection policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
func (client *ServerConnectionPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerConnectionPoliciesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, connectionPolicyName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerConnectionPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, parameters ServerConnectionPolicy, options *ServerConnectionPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if connectionPolicyName == "" {
		return nil, errors.New("parameter connectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionPolicyName}", url.PathEscape(string(connectionPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets a server connection policy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - connectionPolicyName - The name of the connection policy.
//   - options - ServerConnectionPoliciesClientGetOptions contains the optional parameters for the ServerConnectionPoliciesClient.Get
//     method.
func (client *ServerConnectionPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, options *ServerConnectionPoliciesClientGetOptions) (ServerConnectionPoliciesClientGetResponse, error) {
	var err error
	const operationName = "ServerConnectionPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, connectionPolicyName, options)
	if err != nil {
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerConnectionPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, connectionPolicyName ConnectionPolicyName, options *ServerConnectionPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies/{connectionPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if connectionPolicyName == "" {
		return nil, errors.New("parameter connectionPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionPolicyName}", url.PathEscape(string(connectionPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerConnectionPoliciesClient) getHandleResponse(resp *http.Response) (ServerConnectionPoliciesClientGetResponse, error) {
	result := ServerConnectionPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConnectionPolicy); err != nil {
		return ServerConnectionPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Lists connection policy
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerConnectionPoliciesClientListByServerOptions contains the optional parameters for the ServerConnectionPoliciesClient.NewListByServerPager
//     method.
func (client *ServerConnectionPoliciesClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerConnectionPoliciesClientListByServerOptions) *runtime.Pager[ServerConnectionPoliciesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerConnectionPoliciesClientListByServerResponse]{
		More: func(page ServerConnectionPoliciesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerConnectionPoliciesClientListByServerResponse) (ServerConnectionPoliciesClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerConnectionPoliciesClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerConnectionPoliciesClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerConnectionPoliciesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerConnectionPoliciesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/connectionPolicies"
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
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerConnectionPoliciesClient) listByServerHandleResponse(resp *http.Response) (ServerConnectionPoliciesClientListByServerResponse, error) {
	result := ServerConnectionPoliciesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerConnectionPolicyListResult); err != nil {
		return ServerConnectionPoliciesClientListByServerResponse{}, err
	}
	return result, nil
}
