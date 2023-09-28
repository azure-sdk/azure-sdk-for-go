//go:build go1.18
// +build go1.18

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

// InterfaceIPConfigurationsClient contains the methods for the NetworkInterfaceIPConfigurations group.
// Don't use this type directly, use NewInterfaceIPConfigurationsClient() instead.
type InterfaceIPConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInterfaceIPConfigurationsClient creates a new instance of InterfaceIPConfigurationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInterfaceIPConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InterfaceIPConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".InterfaceIPConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InterfaceIPConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the specified network interface ip configuration.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - networkInterfaceName - The name of the network interface.
//   - ipConfigurationName - The name of the ip configuration name.
//   - options - InterfaceIPConfigurationsClientGetOptions contains the optional parameters for the InterfaceIPConfigurationsClient.Get
//     method.
func (client *InterfaceIPConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *InterfaceIPConfigurationsClientGetOptions) (InterfaceIPConfigurationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkInterfaceName, ipConfigurationName, options)
	if err != nil {
		return InterfaceIPConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InterfaceIPConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InterfaceIPConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InterfaceIPConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, ipConfigurationName string, options *InterfaceIPConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations/{ipConfigurationName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if ipConfigurationName == "" {
		return nil, errors.New("parameter ipConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ipConfigurationName}", url.PathEscape(ipConfigurationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *InterfaceIPConfigurationsClient) getHandleResponse(resp *http.Response) (InterfaceIPConfigurationsClientGetResponse, error) {
	result := InterfaceIPConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InterfaceIPConfiguration); err != nil {
		return InterfaceIPConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all ip configurations in a network interface.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group.
//   - networkInterfaceName - The name of the network interface.
//   - options - InterfaceIPConfigurationsClientListOptions contains the optional parameters for the InterfaceIPConfigurationsClient.NewListPager
//     method.
func (client *InterfaceIPConfigurationsClient) NewListPager(resourceGroupName string, networkInterfaceName string, options *InterfaceIPConfigurationsClientListOptions) *runtime.Pager[InterfaceIPConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[InterfaceIPConfigurationsClientListResponse]{
		More: func(page InterfaceIPConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InterfaceIPConfigurationsClientListResponse) (InterfaceIPConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkInterfaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InterfaceIPConfigurationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InterfaceIPConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InterfaceIPConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *InterfaceIPConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkInterfaceName string, options *InterfaceIPConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkInterfaces/{networkInterfaceName}/ipConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkInterfaceName == "" {
		return nil, errors.New("parameter networkInterfaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkInterfaceName}", url.PathEscape(networkInterfaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *InterfaceIPConfigurationsClient) listHandleResponse(resp *http.Response) (InterfaceIPConfigurationsClientListResponse, error) {
	result := InterfaceIPConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InterfaceIPConfigurationListResult); err != nil {
		return InterfaceIPConfigurationsClientListResponse{}, err
	}
	return result, nil
}
