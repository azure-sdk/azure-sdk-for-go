//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// TenantSettingsClient contains the methods for the TenantSettings group.
// Don't use this type directly, use NewTenantSettingsClient() instead.
type TenantSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTenantSettingsClient creates a new instance of TenantSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTenantSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".TenantSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TenantSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get tenant settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - settingsType - The identifier of the settings.
//   - options - TenantSettingsClientGetOptions contains the optional parameters for the TenantSettingsClient.Get method.
func (client *TenantSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, settingsType SettingsTypeName, options *TenantSettingsClientGetOptions) (TenantSettingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, settingsType, options)
	if err != nil {
		return TenantSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TenantSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TenantSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TenantSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, settingsType SettingsTypeName, options *TenantSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/settings/{settingsType}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if settingsType == "" {
		return nil, errors.New("parameter settingsType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{settingsType}", url.PathEscape(string(settingsType)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TenantSettingsClient) getHandleResponse(resp *http.Response) (TenantSettingsClientGetResponse, error) {
	result := TenantSettingsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantSettingsContract); err != nil {
		return TenantSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Public settings.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - TenantSettingsClientListByServiceOptions contains the optional parameters for the TenantSettingsClient.NewListByServicePager
//     method.
func (client *TenantSettingsClient) NewListByServicePager(resourceGroupName string, serviceName string, options *TenantSettingsClientListByServiceOptions) *runtime.Pager[TenantSettingsClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[TenantSettingsClientListByServiceResponse]{
		More: func(page TenantSettingsClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantSettingsClientListByServiceResponse) (TenantSettingsClientListByServiceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TenantSettingsClientListByServiceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TenantSettingsClientListByServiceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TenantSettingsClientListByServiceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServiceHandleResponse(resp)
		},
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *TenantSettingsClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *TenantSettingsClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/settings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *TenantSettingsClient) listByServiceHandleResponse(resp *http.Response) (TenantSettingsClientListByServiceResponse, error) {
	result := TenantSettingsClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TenantSettingsCollection); err != nil {
		return TenantSettingsClientListByServiceResponse{}, err
	}
	return result, nil
}
