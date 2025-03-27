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
	"strconv"
	"strings"
)

// WebAppExtendedMachinesControllerClient contains the methods for the WebAppExtendedMachinesController group.
// Don't use this type directly, use NewWebAppExtendedMachinesControllerClient() instead.
type WebAppExtendedMachinesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebAppExtendedMachinesControllerClient creates a new instance of WebAppExtendedMachinesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebAppExtendedMachinesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebAppExtendedMachinesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebAppExtendedMachinesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Method to get a extended machine.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - extendedMachineName - Extended machine name.
//   - options - WebAppExtendedMachinesControllerClientGetOptions contains the optional parameters for the WebAppExtendedMachinesControllerClient.Get
//     method.
func (client *WebAppExtendedMachinesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, extendedMachineName string, options *WebAppExtendedMachinesControllerClientGetOptions) (WebAppExtendedMachinesControllerClientGetResponse, error) {
	var err error
	const operationName = "WebAppExtendedMachinesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, extendedMachineName, options)
	if err != nil {
		return WebAppExtendedMachinesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppExtendedMachinesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppExtendedMachinesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebAppExtendedMachinesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, extendedMachineName string, _ *WebAppExtendedMachinesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/extendedMachines/{extendedMachineName}"
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
	if extendedMachineName == "" {
		return nil, errors.New("parameter extendedMachineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{extendedMachineName}", url.PathEscape(extendedMachineName))
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
func (client *WebAppExtendedMachinesControllerClient) getHandleResponse(resp *http.Response) (WebAppExtendedMachinesControllerClientGetResponse, error) {
	result := WebAppExtendedMachinesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppExtendedMachine); err != nil {
		return WebAppExtendedMachinesControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWebAppSitePager - Method to get all extended machines.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebAppExtendedMachinesControllerClientListByWebAppSiteOptions contains the optional parameters for the WebAppExtendedMachinesControllerClient.NewListByWebAppSitePager
//     method.
func (client *WebAppExtendedMachinesControllerClient) NewListByWebAppSitePager(resourceGroupName string, siteName string, webAppSiteName string, options *WebAppExtendedMachinesControllerClientListByWebAppSiteOptions) *runtime.Pager[WebAppExtendedMachinesControllerClientListByWebAppSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebAppExtendedMachinesControllerClientListByWebAppSiteResponse]{
		More: func(page WebAppExtendedMachinesControllerClientListByWebAppSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebAppExtendedMachinesControllerClientListByWebAppSiteResponse) (WebAppExtendedMachinesControllerClientListByWebAppSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebAppExtendedMachinesControllerClient.NewListByWebAppSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWebAppSiteCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
			}, nil)
			if err != nil {
				return WebAppExtendedMachinesControllerClientListByWebAppSiteResponse{}, err
			}
			return client.listByWebAppSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWebAppSiteCreateRequest creates the ListByWebAppSite request.
func (client *WebAppExtendedMachinesControllerClient) listByWebAppSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *WebAppExtendedMachinesControllerClientListByWebAppSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/extendedMachines"
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
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", *options.Top)
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWebAppSiteHandleResponse handles the ListByWebAppSite response.
func (client *WebAppExtendedMachinesControllerClient) listByWebAppSiteHandleResponse(resp *http.Response) (WebAppExtendedMachinesControllerClientListByWebAppSiteResponse, error) {
	result := WebAppExtendedMachinesControllerClientListByWebAppSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppExtendedMachineListResult); err != nil {
		return WebAppExtendedMachinesControllerClientListByWebAppSiteResponse{}, err
	}
	return result, nil
}
