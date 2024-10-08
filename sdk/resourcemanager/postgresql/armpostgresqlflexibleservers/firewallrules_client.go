//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// FirewallRulesClient contains the methods for the FirewallRules group.
// Don't use this type directly, use NewFirewallRulesClient() instead.
type FirewallRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFirewallRulesClient creates a new instance of FirewallRulesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFirewallRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallRulesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new firewall rule or updates an existing firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the server firewall rule.
//   - parameters - The required parameters for creating or updating a firewall rule.
//   - options - FirewallRulesClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.BeginCreateOrUpdate
//     method.
func (client *FirewallRulesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule, options *FirewallRulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[FirewallRulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, firewallRuleName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallRulesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FirewallRulesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a new firewall rule or updates an existing firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
func (client *FirewallRulesClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule, options *FirewallRulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallRulesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, parameters, options)
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
func (client *FirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule, options *FirewallRulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}"
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
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a PostgreSQL server firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the server firewall rule.
//   - options - FirewallRulesClientBeginDeleteOptions contains the optional parameters for the FirewallRulesClient.BeginDelete
//     method.
func (client *FirewallRulesClient) BeginDelete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientBeginDeleteOptions) (*runtime.Poller[FirewallRulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, serverName, firewallRuleName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[FirewallRulesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[FirewallRulesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a PostgreSQL server firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
func (client *FirewallRulesClient) deleteOperation(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "FirewallRulesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
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
func (client *FirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}"
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
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - List all the firewall rules in a given server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the server firewall rule.
//   - options - FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
func (client *FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientGetOptions) (FirewallRulesClientGetResponse, error) {
	var err error
	const operationName = "FirewallRulesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
	if err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FirewallRulesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FirewallRulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}"
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
	if firewallRuleName == "" {
		return nil, errors.New("parameter firewallRuleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallRuleName}", url.PathEscape(firewallRuleName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallRulesClient) getHandleResponse(resp *http.Response) (FirewallRulesClientGetResponse, error) {
	result := FirewallRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRule); err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - List all the firewall rules in a given PostgreSQL server.
//
// Generated from API version 2024-08-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - FirewallRulesClientListByServerOptions contains the optional parameters for the FirewallRulesClient.NewListByServerPager
//     method.
func (client *FirewallRulesClient) NewListByServerPager(resourceGroupName string, serverName string, options *FirewallRulesClientListByServerOptions) *runtime.Pager[FirewallRulesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[FirewallRulesClientListByServerResponse]{
		More: func(page FirewallRulesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallRulesClientListByServerResponse) (FirewallRulesClientListByServerResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FirewallRulesClient.NewListByServerPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			}, nil)
			if err != nil {
				return FirewallRulesClientListByServerResponse{}, err
			}
			return client.listByServerHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *FirewallRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *FirewallRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules"
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
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *FirewallRulesClient) listByServerHandleResponse(resp *http.Response) (FirewallRulesClientListByServerResponse, error) {
	result := FirewallRulesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRuleListResult); err != nil {
		return FirewallRulesClientListByServerResponse{}, err
	}
	return result, nil
}
