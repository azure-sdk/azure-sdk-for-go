//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// ActiveSessionHostConfigurationsClient contains the methods for the ActiveSessionHostConfigurations group.
// Don't use this type directly, use NewActiveSessionHostConfigurationsClient() instead.
type ActiveSessionHostConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewActiveSessionHostConfigurationsClient creates a new instance of ActiveSessionHostConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewActiveSessionHostConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ActiveSessionHostConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ActiveSessionHostConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the ActiveSessionHostConfiguration for the hostPool that is currently being used for update operations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-06-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - ActiveSessionHostConfigurationsClientGetOptions contains the optional parameters for the ActiveSessionHostConfigurationsClient.Get
//     method.
func (client *ActiveSessionHostConfigurationsClient) Get(ctx context.Context, resourceGroupName string, hostPoolName string, options *ActiveSessionHostConfigurationsClientGetOptions) (ActiveSessionHostConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "ActiveSessionHostConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, hostPoolName, options)
	if err != nil {
		return ActiveSessionHostConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ActiveSessionHostConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ActiveSessionHostConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ActiveSessionHostConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *ActiveSessionHostConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/activeSessionHostConfigurations/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-06-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ActiveSessionHostConfigurationsClient) getHandleResponse(resp *http.Response) (ActiveSessionHostConfigurationsClientGetResponse, error) {
	result := ActiveSessionHostConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActiveSessionHostConfiguration); err != nil {
		return ActiveSessionHostConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHostPoolPager - List activeSessionHostConfigurations.
//
// Generated from API version 2024-03-06-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - options - ActiveSessionHostConfigurationsClientListByHostPoolOptions contains the optional parameters for the ActiveSessionHostConfigurationsClient.NewListByHostPoolPager
//     method.
func (client *ActiveSessionHostConfigurationsClient) NewListByHostPoolPager(resourceGroupName string, hostPoolName string, options *ActiveSessionHostConfigurationsClientListByHostPoolOptions) *runtime.Pager[ActiveSessionHostConfigurationsClientListByHostPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[ActiveSessionHostConfigurationsClientListByHostPoolResponse]{
		More: func(page ActiveSessionHostConfigurationsClientListByHostPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ActiveSessionHostConfigurationsClientListByHostPoolResponse) (ActiveSessionHostConfigurationsClientListByHostPoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ActiveSessionHostConfigurationsClient.NewListByHostPoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHostPoolCreateRequest(ctx, resourceGroupName, hostPoolName, options)
			}, nil)
			if err != nil {
				return ActiveSessionHostConfigurationsClientListByHostPoolResponse{}, err
			}
			return client.listByHostPoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHostPoolCreateRequest creates the ListByHostPool request.
func (client *ActiveSessionHostConfigurationsClient) listByHostPoolCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, options *ActiveSessionHostConfigurationsClientListByHostPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/activeSessionHostConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-06-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHostPoolHandleResponse handles the ListByHostPool response.
func (client *ActiveSessionHostConfigurationsClient) listByHostPoolHandleResponse(resp *http.Response) (ActiveSessionHostConfigurationsClientListByHostPoolResponse, error) {
	result := ActiveSessionHostConfigurationsClientListByHostPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ActiveSessionHostConfigurationList); err != nil {
		return ActiveSessionHostConfigurationsClientListByHostPoolResponse{}, err
	}
	return result, nil
}
