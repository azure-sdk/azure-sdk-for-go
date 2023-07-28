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

// NetworkSecurityPerimeterConfigurationsClient contains the methods for the NetworkSecurityPerimeterConfigurations group.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationsClient() instead.
type NetworkSecurityPerimeterConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterConfigurationsClient creates a new instance of NetworkSecurityPerimeterConfigurationsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkSecurityPerimeterConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a network security perimeter configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - NetworkSecurityPerimeterConfigurationsClientGetOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.Get
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) Get(ctx context.Context, resourceGroupName string, serverName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientGetOptions) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, nspConfigName, options)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkSecurityPerimeterConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/networkSecurityPerimeterConfigurations/{nspConfigName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if nspConfigName == "" {
		return nil, errors.New("parameter nspConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nspConfigName}", url.PathEscape(nspConfigName))
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
func (client *NetworkSecurityPerimeterConfigurationsClient) getHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientGetResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfiguration); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of NSP configurations for a server.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - NetworkSecurityPerimeterConfigurationsClientListByServerOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.NewListByServerPager
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) NewListByServerPager(resourceGroupName string, serverName string, options *NetworkSecurityPerimeterConfigurationsClientListByServerOptions) *runtime.Pager[NetworkSecurityPerimeterConfigurationsClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkSecurityPerimeterConfigurationsClientListByServerResponse]{
		More: func(page NetworkSecurityPerimeterConfigurationsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkSecurityPerimeterConfigurationsClientListByServerResponse) (NetworkSecurityPerimeterConfigurationsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkSecurityPerimeterConfigurationsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkSecurityPerimeterConfigurationsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *NetworkSecurityPerimeterConfigurationsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *NetworkSecurityPerimeterConfigurationsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/networkSecurityPerimeterConfigurations"
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
func (client *NetworkSecurityPerimeterConfigurationsClient) listByServerHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationsClientListByServerResponse, error) {
	result := NetworkSecurityPerimeterConfigurationsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationListResult); err != nil {
		return NetworkSecurityPerimeterConfigurationsClientListByServerResponse{}, err
	}
	return result, nil
}

// BeginReconcile - Reconcile network security perimeter configuration for SQL Resource Provider
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationsClient.BeginReconcile
//     method.
func (client *NetworkSecurityPerimeterConfigurationsClient) BeginReconcile(ctx context.Context, resourceGroupName string, serverName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*runtime.Poller[NetworkSecurityPerimeterConfigurationsClientReconcileResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.reconcile(ctx, resourceGroupName, serverName, nspConfigName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkSecurityPerimeterConfigurationsClientReconcileResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[NetworkSecurityPerimeterConfigurationsClientReconcileResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Reconcile - Reconcile network security perimeter configuration for SQL Resource Provider
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcile(ctx context.Context, resourceGroupName string, serverName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*http.Response, error) {
	var err error
	req, err := client.reconcileCreateRequest(ctx, resourceGroupName, serverName, nspConfigName, options)
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

// reconcileCreateRequest creates the Reconcile request.
func (client *NetworkSecurityPerimeterConfigurationsClient) reconcileCreateRequest(ctx context.Context, resourceGroupName string, serverName string, nspConfigName string, options *NetworkSecurityPerimeterConfigurationsClientBeginReconcileOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/networkSecurityPerimeterConfigurations/{nspConfigName}/reconcile"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if nspConfigName == "" {
		return nil, errors.New("parameter nspConfigName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nspConfigName}", url.PathEscape(nspConfigName))
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
	return req, nil
}
