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

// FirewallRulesClient contains the methods for the FirewallRules group.
// Don't use this type directly, use NewFirewallRulesClient() instead.
type FirewallRulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFirewallRulesClient creates a new instance of FirewallRulesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFirewallRulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FirewallRulesClient, error) {
	cl, err := arm.NewClient(moduleName+".FirewallRulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FirewallRulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the firewall rule.
//   - parameters - The required parameters for creating or updating a firewall rule.
//   - options - FirewallRulesClientCreateOrUpdateOptions contains the optional parameters for the FirewallRulesClient.CreateOrUpdate
//     method.
func (client *FirewallRulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule, options *FirewallRulesClientCreateOrUpdateOptions) (FirewallRulesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, parameters, options)
	if err != nil {
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallRulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, parameters FirewallRule, options *FirewallRulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}"
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
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FirewallRulesClient) createOrUpdateHandleResponse(resp *http.Response) (FirewallRulesClientCreateOrUpdateResponse, error) {
	result := FirewallRulesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRule); err != nil {
		return FirewallRulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the firewall rule.
//   - options - FirewallRulesClientDeleteOptions contains the optional parameters for the FirewallRulesClient.Delete method.
func (client *FirewallRulesClient) Delete(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientDeleteOptions) (FirewallRulesClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, firewallRuleName, options)
	if err != nil {
		return FirewallRulesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallRulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return FirewallRulesClientDeleteResponse{}, err
	}
	return FirewallRulesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FirewallRulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}"
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

// Get - Gets a firewall rule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - firewallRuleName - The name of the firewall rule.
//   - options - FirewallRulesClientGetOptions contains the optional parameters for the FirewallRulesClient.Get method.
func (client *FirewallRulesClient) Get(ctx context.Context, resourceGroupName string, serverName string, firewallRuleName string, options *FirewallRulesClientGetOptions) (FirewallRulesClientGetResponse, error) {
	var err error
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
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}"
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
func (client *FirewallRulesClient) getHandleResponse(resp *http.Response) (FirewallRulesClientGetResponse, error) {
	result := FirewallRulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRule); err != nil {
		return FirewallRulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of firewall rules.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - FirewallRulesClientListByServerOptions contains the optional parameters for the FirewallRulesClient.NewListByServerPager
//     method.
func (client *FirewallRulesClient) NewListByServerPager(resourceGroupName string, serverName string, options *FirewallRulesClientListByServerOptions) *runtime.Pager[FirewallRulesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[FirewallRulesClientListByServerResponse]{
		More: func(page FirewallRulesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallRulesClientListByServerResponse) (FirewallRulesClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FirewallRulesClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FirewallRulesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FirewallRulesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *FirewallRulesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *FirewallRulesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules"
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
func (client *FirewallRulesClient) listByServerHandleResponse(resp *http.Response) (FirewallRulesClientListByServerResponse, error) {
	result := FirewallRulesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRuleListResult); err != nil {
		return FirewallRulesClientListByServerResponse{}, err
	}
	return result, nil
}

// Replace - Replaces all firewall rules on the server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - FirewallRulesClientReplaceOptions contains the optional parameters for the FirewallRulesClient.Replace method.
func (client *FirewallRulesClient) Replace(ctx context.Context, resourceGroupName string, serverName string, parameters FirewallRuleList, options *FirewallRulesClientReplaceOptions) (FirewallRulesClientReplaceResponse, error) {
	var err error
	req, err := client.replaceCreateRequest(ctx, resourceGroupName, serverName, parameters, options)
	if err != nil {
		return FirewallRulesClientReplaceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FirewallRulesClientReplaceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return FirewallRulesClientReplaceResponse{}, err
	}
	resp, err := client.replaceHandleResponse(httpResp)
	return resp, err
}

// replaceCreateRequest creates the Replace request.
func (client *FirewallRulesClient) replaceCreateRequest(ctx context.Context, resourceGroupName string, serverName string, parameters FirewallRuleList, options *FirewallRulesClientReplaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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

// replaceHandleResponse handles the Replace response.
func (client *FirewallRulesClient) replaceHandleResponse(resp *http.Response) (FirewallRulesClientReplaceResponse, error) {
	result := FirewallRulesClientReplaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallRule); err != nil {
		return FirewallRulesClientReplaceResponse{}, err
	}
	return result, nil
}
