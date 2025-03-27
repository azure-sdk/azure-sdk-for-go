// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrationdiscovery

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

// WebAppRunAsAccountsControllerClient contains the methods for the WebAppRunAsAccountsController group.
// Don't use this type directly, use NewWebAppRunAsAccountsControllerClient() instead.
type WebAppRunAsAccountsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebAppRunAsAccountsControllerClient creates a new instance of WebAppRunAsAccountsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebAppRunAsAccountsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebAppRunAsAccountsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebAppRunAsAccountsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Method to get run as account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - accountName - Run as account ARM name.
//   - options - WebAppRunAsAccountsControllerClientGetOptions contains the optional parameters for the WebAppRunAsAccountsControllerClient.Get
//     method.
func (client *WebAppRunAsAccountsControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, accountName string, options *WebAppRunAsAccountsControllerClientGetOptions) (WebAppRunAsAccountsControllerClientGetResponse, error) {
	var err error
	const operationName = "WebAppRunAsAccountsControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, accountName, options)
	if err != nil {
		return WebAppRunAsAccountsControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppRunAsAccountsControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppRunAsAccountsControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebAppRunAsAccountsControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, accountName string, _ *WebAppRunAsAccountsControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/runasaccounts/{accountName}"
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
	if webAppSiteName == "" {
		return nil, errors.New("parameter webAppSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webAppSiteName}", url.PathEscape(webAppSiteName))
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
func (client *WebAppRunAsAccountsControllerClient) getHandleResponse(resp *http.Response) (WebAppRunAsAccountsControllerClientGetResponse, error) {
	result := WebAppRunAsAccountsControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppRunAsAccount); err != nil {
		return WebAppRunAsAccountsControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWebAppSitePager - Method to get all run as accounts.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebAppRunAsAccountsControllerClientListByWebAppSiteOptions contains the optional parameters for the WebAppRunAsAccountsControllerClient.NewListByWebAppSitePager
//     method.
func (client *WebAppRunAsAccountsControllerClient) NewListByWebAppSitePager(resourceGroupName string, siteName string, webAppSiteName string, options *WebAppRunAsAccountsControllerClientListByWebAppSiteOptions) *runtime.Pager[WebAppRunAsAccountsControllerClientListByWebAppSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebAppRunAsAccountsControllerClientListByWebAppSiteResponse]{
		More: func(page WebAppRunAsAccountsControllerClientListByWebAppSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebAppRunAsAccountsControllerClientListByWebAppSiteResponse) (WebAppRunAsAccountsControllerClientListByWebAppSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebAppRunAsAccountsControllerClient.NewListByWebAppSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWebAppSiteCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
			}, nil)
			if err != nil {
				return WebAppRunAsAccountsControllerClientListByWebAppSiteResponse{}, err
			}
			return client.listByWebAppSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWebAppSiteCreateRequest creates the ListByWebAppSite request.
func (client *WebAppRunAsAccountsControllerClient) listByWebAppSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, _ *WebAppRunAsAccountsControllerClientListByWebAppSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/runasaccounts"
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
	if webAppSiteName == "" {
		return nil, errors.New("parameter webAppSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{webAppSiteName}", url.PathEscape(webAppSiteName))
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

// listByWebAppSiteHandleResponse handles the ListByWebAppSite response.
func (client *WebAppRunAsAccountsControllerClient) listByWebAppSiteHandleResponse(resp *http.Response) (WebAppRunAsAccountsControllerClientListByWebAppSiteResponse, error) {
	result := WebAppRunAsAccountsControllerClientListByWebAppSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppRunAsAccountListResult); err != nil {
		return WebAppRunAsAccountsControllerClientListByWebAppSiteResponse{}, err
	}
	return result, nil
}
