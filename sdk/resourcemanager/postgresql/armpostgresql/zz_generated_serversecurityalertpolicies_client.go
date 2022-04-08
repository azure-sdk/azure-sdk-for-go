//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresql

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ServerSecurityAlertPoliciesClient contains the methods for the ServerSecurityAlertPolicies group.
// Don't use this type directly, use NewServerSecurityAlertPoliciesClient() instead.
type ServerSecurityAlertPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerSecurityAlertPoliciesClient creates a new instance of ServerSecurityAlertPoliciesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerSecurityAlertPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *ServerSecurityAlertPoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &ServerSecurityAlertPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a threat detection policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// securityAlertPolicyName - The name of the threat detection policy.
// parameters - The server security alert policy.
// options - ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerSecurityAlertPoliciesClient.BeginCreateOrUpdate
// method.
func (client *ServerSecurityAlertPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyName, parameters ServerSecurityAlertPolicy, options *ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (ServerSecurityAlertPoliciesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return ServerSecurityAlertPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result := ServerSecurityAlertPoliciesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("ServerSecurityAlertPoliciesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return ServerSecurityAlertPoliciesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &ServerSecurityAlertPoliciesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a threat detection policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerSecurityAlertPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyName, parameters ServerSecurityAlertPolicy, options *ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, securityAlertPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerSecurityAlertPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyName, parameters ServerSecurityAlertPolicy, options *ServerSecurityAlertPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Get a server's security alert policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// securityAlertPolicyName - The name of the security alert policy.
// options - ServerSecurityAlertPoliciesClientGetOptions contains the optional parameters for the ServerSecurityAlertPoliciesClient.Get
// method.
func (client *ServerSecurityAlertPoliciesClient) Get(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyName, options *ServerSecurityAlertPoliciesClientGetOptions) (ServerSecurityAlertPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, securityAlertPolicyName, options)
	if err != nil {
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerSecurityAlertPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerSecurityAlertPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, securityAlertPolicyName SecurityAlertPolicyName, options *ServerSecurityAlertPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies/{securityAlertPolicyName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerSecurityAlertPoliciesClient) getHandleResponse(resp *http.Response) (ServerSecurityAlertPoliciesClientGetResponse, error) {
	result := ServerSecurityAlertPoliciesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerSecurityAlertPolicy); err != nil {
		return ServerSecurityAlertPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Get the server's threat detection policies.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group. The name is case insensitive.
// serverName - The name of the server.
// options - ServerSecurityAlertPoliciesClientListByServerOptions contains the optional parameters for the ServerSecurityAlertPoliciesClient.ListByServer
// method.
func (client *ServerSecurityAlertPoliciesClient) ListByServer(resourceGroupName string, serverName string, options *ServerSecurityAlertPoliciesClientListByServerOptions) *ServerSecurityAlertPoliciesClientListByServerPager {
	return &ServerSecurityAlertPoliciesClientListByServerPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
		},
		advancer: func(ctx context.Context, resp ServerSecurityAlertPoliciesClientListByServerResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ServerSecurityAlertPolicyListResult.NextLink)
		},
	}
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerSecurityAlertPoliciesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerSecurityAlertPoliciesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/servers/{serverName}/securityAlertPolicies"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerSecurityAlertPoliciesClient) listByServerHandleResponse(resp *http.Response) (ServerSecurityAlertPoliciesClientListByServerResponse, error) {
	result := ServerSecurityAlertPoliciesClientListByServerResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerSecurityAlertPolicyListResult); err != nil {
		return ServerSecurityAlertPoliciesClientListByServerResponse{}, err
	}
	return result, nil
}
