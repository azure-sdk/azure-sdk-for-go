//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// ServerRunAsAccountsControllerClient contains the methods for the ServerRunAsAccountsController group.
// Don't use this type directly, use NewServerRunAsAccountsControllerClient() instead.
type ServerRunAsAccountsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerRunAsAccountsControllerClient creates a new instance of ServerRunAsAccountsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerRunAsAccountsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerRunAsAccountsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerRunAsAccountsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a ServerRunAsAccount
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - accountName - RunAsAccounts name
//   - options - ServerRunAsAccountsControllerClientGetOptions contains the optional parameters for the ServerRunAsAccountsControllerClient.Get
//     method.
func (client *ServerRunAsAccountsControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *ServerRunAsAccountsControllerClientGetOptions) (ServerRunAsAccountsControllerClientGetResponse, error) {
	var err error
	const operationName = "ServerRunAsAccountsControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, accountName, options)
	if err != nil {
		return ServerRunAsAccountsControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerRunAsAccountsControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerRunAsAccountsControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerRunAsAccountsControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *ServerRunAsAccountsControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/runAsAccounts/{accountName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerRunAsAccountsControllerClient) getHandleResponse(resp *http.Response) (ServerRunAsAccountsControllerClientGetResponse, error) {
	result := ServerRunAsAccountsControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerRunAsAccount); err != nil {
		return ServerRunAsAccountsControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerSiteResourcePager - List ServerRunAsAccount resources by ServerSiteResource
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - ServerRunAsAccountsControllerClientListByServerSiteResourceOptions contains the optional parameters for the ServerRunAsAccountsControllerClient.NewListByServerSiteResourcePager
//     method.
func (client *ServerRunAsAccountsControllerClient) NewListByServerSiteResourcePager(resourceGroupName string, siteName string, options *ServerRunAsAccountsControllerClientListByServerSiteResourceOptions) *runtime.Pager[ServerRunAsAccountsControllerClientListByServerSiteResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerRunAsAccountsControllerClientListByServerSiteResourceResponse]{
		More: func(page ServerRunAsAccountsControllerClientListByServerSiteResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerRunAsAccountsControllerClientListByServerSiteResourceResponse) (ServerRunAsAccountsControllerClientListByServerSiteResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerRunAsAccountsControllerClient.NewListByServerSiteResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerSiteResourceCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return ServerRunAsAccountsControllerClientListByServerSiteResourceResponse{}, err
			}
			return client.listByServerSiteResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerSiteResourceCreateRequest creates the ListByServerSiteResource request.
func (client *ServerRunAsAccountsControllerClient) listByServerSiteResourceCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *ServerRunAsAccountsControllerClientListByServerSiteResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/runAsAccounts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerSiteResourceHandleResponse handles the ListByServerSiteResource response.
func (client *ServerRunAsAccountsControllerClient) listByServerSiteResourceHandleResponse(resp *http.Response) (ServerRunAsAccountsControllerClientListByServerSiteResourceResponse, error) {
	result := ServerRunAsAccountsControllerClientListByServerSiteResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerRunAsAccountListResult); err != nil {
		return ServerRunAsAccountsControllerClientListByServerSiteResourceResponse{}, err
	}
	return result, nil
}
