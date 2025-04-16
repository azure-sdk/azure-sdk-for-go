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

// WebAppDiscoverySiteDataSourcesControllerClient contains the methods for the WebAppDiscoverySiteDataSourcesController group.
// Don't use this type directly, use NewWebAppDiscoverySiteDataSourcesControllerClient() instead.
type WebAppDiscoverySiteDataSourcesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebAppDiscoverySiteDataSourcesControllerClient creates a new instance of WebAppDiscoverySiteDataSourcesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebAppDiscoverySiteDataSourcesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebAppDiscoverySiteDataSourcesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebAppDiscoverySiteDataSourcesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Method to create or update a Web app data source in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - discoverySiteDataSourceName - Data Source ARM name.
//   - body - Resource create parameters.
//   - options - WebAppDiscoverySiteDataSourcesControllerClientBeginCreateOptions contains the optional parameters for the WebAppDiscoverySiteDataSourcesControllerClient.BeginCreate
//     method.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) BeginCreate(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, body DiscoverySiteDataSource, options *WebAppDiscoverySiteDataSourcesControllerClientBeginCreateOptions) (*runtime.Poller[WebAppDiscoverySiteDataSourcesControllerClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, siteName, webAppSiteName, discoverySiteDataSourceName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppDiscoverySiteDataSourcesControllerClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppDiscoverySiteDataSourcesControllerClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Method to create or update a Web app data source in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *WebAppDiscoverySiteDataSourcesControllerClient) create(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, body DiscoverySiteDataSource, options *WebAppDiscoverySiteDataSourcesControllerClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppDiscoverySiteDataSourcesControllerClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, discoverySiteDataSourceName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, body DiscoverySiteDataSource, _ *WebAppDiscoverySiteDataSourcesControllerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/discoverySiteDataSources/{discoverySiteDataSourceName}"
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
	if discoverySiteDataSourceName == "" {
		return nil, errors.New("parameter discoverySiteDataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{discoverySiteDataSourceName}", url.PathEscape(discoverySiteDataSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Method to delete a Web app data source in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - discoverySiteDataSourceName - Data Source ARM name.
//   - options - WebAppDiscoverySiteDataSourcesControllerClientBeginDeleteOptions contains the optional parameters for the WebAppDiscoverySiteDataSourcesControllerClient.BeginDelete
//     method.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) BeginDelete(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, options *WebAppDiscoverySiteDataSourcesControllerClientBeginDeleteOptions) (*runtime.Poller[WebAppDiscoverySiteDataSourcesControllerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, siteName, webAppSiteName, discoverySiteDataSourceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppDiscoverySiteDataSourcesControllerClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppDiscoverySiteDataSourcesControllerClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Method to delete a Web app data source in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *WebAppDiscoverySiteDataSourcesControllerClient) deleteOperation(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, options *WebAppDiscoverySiteDataSourcesControllerClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppDiscoverySiteDataSourcesControllerClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, discoverySiteDataSourceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, _ *WebAppDiscoverySiteDataSourcesControllerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/discoverySiteDataSources/{discoverySiteDataSourceName}"
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
	if discoverySiteDataSourceName == "" {
		return nil, errors.New("parameter discoverySiteDataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{discoverySiteDataSourceName}", url.PathEscape(discoverySiteDataSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Method to get a Web app data source in site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - discoverySiteDataSourceName - Data Source ARM name.
//   - options - WebAppDiscoverySiteDataSourcesControllerClientGetOptions contains the optional parameters for the WebAppDiscoverySiteDataSourcesControllerClient.Get
//     method.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, options *WebAppDiscoverySiteDataSourcesControllerClientGetOptions) (WebAppDiscoverySiteDataSourcesControllerClientGetResponse, error) {
	var err error
	const operationName = "WebAppDiscoverySiteDataSourcesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, discoverySiteDataSourceName, options)
	if err != nil {
		return WebAppDiscoverySiteDataSourcesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppDiscoverySiteDataSourcesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppDiscoverySiteDataSourcesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, discoverySiteDataSourceName string, _ *WebAppDiscoverySiteDataSourcesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/discoverySiteDataSources/{discoverySiteDataSourceName}"
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
	if discoverySiteDataSourceName == "" {
		return nil, errors.New("parameter discoverySiteDataSourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{discoverySiteDataSourceName}", url.PathEscape(discoverySiteDataSourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) getHandleResponse(resp *http.Response) (WebAppDiscoverySiteDataSourcesControllerClientGetResponse, error) {
	result := WebAppDiscoverySiteDataSourcesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoverySiteDataSource); err != nil {
		return WebAppDiscoverySiteDataSourcesControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByWebAppSitePager - Method to get all Web app data sources in site.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteOptions contains the optional parameters for the
//     WebAppDiscoverySiteDataSourcesControllerClient.NewListByWebAppSitePager method.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) NewListByWebAppSitePager(resourceGroupName string, siteName string, webAppSiteName string, options *WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteOptions) *runtime.Pager[WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse]{
		More: func(page WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse) (WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebAppDiscoverySiteDataSourcesControllerClient.NewListByWebAppSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByWebAppSiteCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
			}, nil)
			if err != nil {
				return WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse{}, err
			}
			return client.listByWebAppSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByWebAppSiteCreateRequest creates the ListByWebAppSite request.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) listByWebAppSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, _ *WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/discoverySiteDataSources"
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
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWebAppSiteHandleResponse handles the ListByWebAppSite response.
func (client *WebAppDiscoverySiteDataSourcesControllerClient) listByWebAppSiteHandleResponse(resp *http.Response) (WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse, error) {
	result := WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoverySiteDataSourceListResult); err != nil {
		return WebAppDiscoverySiteDataSourcesControllerClientListByWebAppSiteResponse{}, err
	}
	return result, nil
}
