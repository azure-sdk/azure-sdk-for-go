//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ServerAdvisorsClient contains the methods for the ServerAdvisors group.
// Don't use this type directly, use NewServerAdvisorsClient() instead.
type ServerAdvisorsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerAdvisorsClient creates a new instance of ServerAdvisorsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerAdvisorsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerAdvisorsClient, error) {
	cl, err := arm.NewClient(moduleName+".ServerAdvisorsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerAdvisorsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a server advisor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - advisorName - The name of the Server Advisor.
//   - options - ServerAdvisorsClientGetOptions contains the optional parameters for the ServerAdvisorsClient.Get method.
func (client *ServerAdvisorsClient) Get(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *ServerAdvisorsClientGetOptions) (ServerAdvisorsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, advisorName, options)
	if err != nil {
		return ServerAdvisorsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAdvisorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdvisorsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerAdvisorsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, options *ServerAdvisorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
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
func (client *ServerAdvisorsClient) getHandleResponse(resp *http.Response) (ServerAdvisorsClientGetResponse, error) {
	result := ServerAdvisorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Advisor); err != nil {
		return ServerAdvisorsClientGetResponse{}, err
	}
	return result, nil
}

// ListByServer - Gets a list of server advisors.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - ServerAdvisorsClientListByServerOptions contains the optional parameters for the ServerAdvisorsClient.ListByServer
//     method.
func (client *ServerAdvisorsClient) ListByServer(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdvisorsClientListByServerOptions) (ServerAdvisorsClientListByServerResponse, error) {
	req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return ServerAdvisorsClientListByServerResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAdvisorsClientListByServerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdvisorsClientListByServerResponse{}, runtime.NewResponseError(resp)
	}
	return client.listByServerHandleResponse(resp)
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerAdvisorsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerAdvisorsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors"
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerAdvisorsClient) listByServerHandleResponse(resp *http.Response) (ServerAdvisorsClientListByServerResponse, error) {
	result := ServerAdvisorsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AdvisorArray); err != nil {
		return ServerAdvisorsClientListByServerResponse{}, err
	}
	return result, nil
}

// Update - Updates a server advisor.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - advisorName - The name of the Server Advisor.
//   - parameters - The requested advisor resource state.
//   - options - ServerAdvisorsClientUpdateOptions contains the optional parameters for the ServerAdvisorsClient.Update method.
func (client *ServerAdvisorsClient) Update(ctx context.Context, resourceGroupName string, serverName string, advisorName string, parameters Advisor, options *ServerAdvisorsClientUpdateOptions) (ServerAdvisorsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, advisorName, parameters, options)
	if err != nil {
		return ServerAdvisorsClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerAdvisorsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerAdvisorsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *ServerAdvisorsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, advisorName string, parameters Advisor, options *ServerAdvisorsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/advisors/{advisorName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if advisorName == "" {
		return nil, errors.New("parameter advisorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advisorName}", url.PathEscape(advisorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *ServerAdvisorsClient) updateHandleResponse(resp *http.Response) (ServerAdvisorsClientUpdateResponse, error) {
	result := ServerAdvisorsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Advisor); err != nil {
		return ServerAdvisorsClientUpdateResponse{}, err
	}
	return result, nil
}
