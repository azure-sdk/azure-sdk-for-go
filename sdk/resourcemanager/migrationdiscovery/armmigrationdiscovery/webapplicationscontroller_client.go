//go:build go1.18
// +build go1.18

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

// WebApplicationsControllerClient contains the methods for the WebApplicationsController group.
// Don't use this type directly, use NewWebApplicationsControllerClient() instead.
type WebApplicationsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebApplicationsControllerClient creates a new instance of WebApplicationsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebApplicationsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebApplicationsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebApplicationsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByWebAppSitePager - Method to get all IIS web applications.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebApplicationsControllerClientListByWebAppSiteOptions contains the optional parameters for the WebApplicationsControllerClient.NewListByWebAppSitePager
//     method.
func (client *WebApplicationsControllerClient) NewListByWebAppSitePager(resourceGroupName string, siteName string, webAppSiteName string, options *WebApplicationsControllerClientListByWebAppSiteOptions) *runtime.Pager[WebApplicationsControllerClientListByWebAppSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebApplicationsControllerClientListByWebAppSiteResponse]{
		More: func(page WebApplicationsControllerClientListByWebAppSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebApplicationsControllerClientListByWebAppSiteResponse) (WebApplicationsControllerClientListByWebAppSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebApplicationsControllerClient.NewListByWebAppSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWebAppSiteCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
			}, nil)
			if err != nil {
				return WebApplicationsControllerClientListByWebAppSiteResponse{}, err
			}
			return client.listByWebAppSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWebAppSiteCreateRequest creates the ListByWebAppSite request.
func (client *WebApplicationsControllerClient) listByWebAppSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *WebApplicationsControllerClientListByWebAppSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/webApplications"
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
func (client *WebApplicationsControllerClient) listByWebAppSiteHandleResponse(resp *http.Response) (WebApplicationsControllerClientListByWebAppSiteResponse, error) {
	result := WebApplicationsControllerClientListByWebAppSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebApplicationListResult); err != nil {
		return WebApplicationsControllerClientListByWebAppSiteResponse{}, err
	}
	return result, nil
}
