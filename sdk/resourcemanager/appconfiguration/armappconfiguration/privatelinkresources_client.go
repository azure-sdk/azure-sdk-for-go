//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The Microsoft Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a private link resource that need to be created for a configuration store.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - configStoreName - The name of the configuration store.
//   - groupName - The name of the private link resource group.
//   - options - PrivateLinkResourcesClientGetOptions contains the optional parameters for the PrivateLinkResourcesClient.Get
//     method.
func (client *PrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, configStoreName string, groupName string, options *PrivateLinkResourcesClientGetOptions) (PrivateLinkResourcesClientGetResponse, error) {
	var err error
	const operationName = "PrivateLinkResourcesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, configStoreName, groupName, options)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesClient) getCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, groupName string, options *PrivateLinkResourcesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/privateLinkResources/{groupName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
	if groupName == "" {
		return nil, errors.New("parameter groupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupName}", url.PathEscape(groupName))
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
func (client *PrivateLinkResourcesClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesClientGetResponse, error) {
	result := PrivateLinkResourcesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByConfigurationStorePager - Gets the private link resources that need to be created for a configuration store.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group to which the container registry belongs.
//   - configStoreName - The name of the configuration store.
//   - options - PrivateLinkResourcesClientListByConfigurationStoreOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByConfigurationStorePager
//     method.
func (client *PrivateLinkResourcesClient) NewListByConfigurationStorePager(resourceGroupName string, configStoreName string, options *PrivateLinkResourcesClientListByConfigurationStoreOptions) *runtime.Pager[PrivateLinkResourcesClientListByConfigurationStoreResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListByConfigurationStoreResponse]{
		More: func(page PrivateLinkResourcesClientListByConfigurationStoreResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListByConfigurationStoreResponse) (PrivateLinkResourcesClientListByConfigurationStoreResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkResourcesClient.NewListByConfigurationStorePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByConfigurationStoreCreateRequest(ctx, resourceGroupName, configStoreName, options)
			}, nil)
			if err != nil {
				return PrivateLinkResourcesClientListByConfigurationStoreResponse{}, err
			}
			return client.listByConfigurationStoreHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByConfigurationStoreCreateRequest creates the ListByConfigurationStore request.
func (client *PrivateLinkResourcesClient) listByConfigurationStoreCreateRequest(ctx context.Context, resourceGroupName string, configStoreName string, options *PrivateLinkResourcesClientListByConfigurationStoreOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppConfiguration/configurationStores/{configStoreName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if configStoreName == "" {
		return nil, errors.New("parameter configStoreName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configStoreName}", url.PathEscape(configStoreName))
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

// listByConfigurationStoreHandleResponse handles the ListByConfigurationStore response.
func (client *PrivateLinkResourcesClient) listByConfigurationStoreHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByConfigurationStoreResponse, error) {
	result := PrivateLinkResourcesClientListByConfigurationStoreResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByConfigurationStoreResponse{}, err
	}
	return result, nil
}
