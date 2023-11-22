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

// ServerSecurityAlertPoliciesClient contains the methods for the ServerSecurityAlertPolicies group.
// Don't use this type directly, use NewServerSecurityAlertPoliciesClient() instead.
type ServerSecurityAlertPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerSecurityAlertPoliciesClient creates a new instance of ServerSecurityAlertPoliciesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerSecurityAlertPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerSecurityAlertPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerSecurityAlertPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a threat detection policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - securityAlertPolicyName - The name of the threat detection policy.
//   - parameters - The server security alert policy.
//   - options - ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate
//     method.
func (client *ServerSecurityAlertPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, parameters ServerSecurityAlertPolicy, options *ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (*runtime.Poller[ServerSecurityAlertPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, securityAlertPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ServerSecurityAlertPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ServerSecurityAlertPoliciesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a threat detection policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
func (client *ServerSecurityAlertPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, parameters ServerSecurityAlertPolicy, options *ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, securityAlertPolicyName, parameters, options)
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
func (client *ServerSecurityAlertPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, parameters ServerSecurityAlertPolicy, options *ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Get a server's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - securityAlertPolicyName - The name of the security alert policy.
//   - options - ServerSecurityAlertPoliciesClientGetOptions contains the optional parameters for the ServerSecurityAlertPoliciesClient.Get
//     method.
func (client *ServerSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, options *ServerSecurityAlertPoliciesClientGetOptions) (ServerSecurityAlertPoliciesClientGetResponse, error) {
	var err error
	const operationName = "ServerSecurityAlertPoliciesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, securityAlertPolicyName, options)
	if err != nil {
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerSecurityAlertPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyNameAutoGenerated, options *ServerSecurityAlertPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if securityAlertPolicyName == "" {
		return nil, errors.New("parameter securityAlertPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityAlertPolicyName}", url.PathEscape(string(securityAlertPolicyName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerSecurityAlertPoliciesClient) getHandleResponse(resp *http.Response) (ServerSecurityAlertPoliciesClientGetResponse, error) {
	result := ServerSecurityAlertPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerSecurityAlertPolicy); err != nil {
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Get the server's threat detection policies.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerSecurityAlertPoliciesClientListByServerOptions contains the optional parameters for the ServerSecurityAlertPoliciesClient.NewListByServerPager
//     method.
func (client *ServerSecurityAlertPoliciesClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerSecurityAlertPoliciesClientListByServerOptions) *runtime.Pager[ServerSecurityAlertPoliciesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerSecurityAlertPoliciesClientListByServerResponse]{
		More: func(page ServerSecurityAlertPoliciesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerSecurityAlertPoliciesClientListByServerResponse) (ServerSecurityAlertPoliciesClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerSecurityAlertPoliciesClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return ServerSecurityAlertPoliciesClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerSecurityAlertPoliciesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerSecurityAlertPoliciesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/securityAlertPolicies"
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
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerSecurityAlertPoliciesClient) listByServerHandleResponse(resp *http.Response) (ServerSecurityAlertPoliciesClientListByServerResponse, error) {
	result := ServerSecurityAlertPoliciesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LogicalServerSecurityAlertPolicyListResult); err != nil {
		return ServerSecurityAlertPoliciesClientListByServerResponse{}, err
	}
	return result, nil
}
