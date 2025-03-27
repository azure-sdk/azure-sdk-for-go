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
	"strconv"
	"strings"
)

// WebAppSitesControllerClient contains the methods for the WebAppSitesController group.
// Don't use this type directly, use NewWebAppSitesControllerClient() instead.
type WebAppSitesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWebAppSitesControllerClient creates a new instance of WebAppSitesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWebAppSitesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WebAppSitesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WebAppSitesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Method to create a WebApp site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - body - Resource create parameters.
//   - options - WebAppSitesControllerClientCreateOptions contains the optional parameters for the WebAppSitesControllerClient.Create
//     method.
func (client *WebAppSitesControllerClient) Create(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body WebAppSite, options *WebAppSitesControllerClientCreateOptions) (WebAppSitesControllerClientCreateResponse, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, body, options)
	if err != nil {
		return WebAppSitesControllerClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppSitesControllerClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WebAppSitesControllerClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *WebAppSitesControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body WebAppSite, _ *WebAppSitesControllerClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *WebAppSitesControllerClient) createHandleResponse(resp *http.Response) (WebAppSitesControllerClientCreateResponse, error) {
	result := WebAppSitesControllerClientCreateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return WebAppSitesControllerClientCreateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppSite); err != nil {
		return WebAppSitesControllerClientCreateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes the WebApp site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebAppSitesControllerClientBeginDeleteOptions contains the optional parameters for the WebAppSitesControllerClient.BeginDelete
//     method.
func (client *WebAppSitesControllerClient) BeginDelete(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *WebAppSitesControllerClientBeginDeleteOptions) (*runtime.Poller[WebAppSitesControllerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, siteName, webAppSiteName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppSitesControllerClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppSitesControllerClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the WebApp site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
func (client *WebAppSitesControllerClient) deleteOperation(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *WebAppSitesControllerClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
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
func (client *WebAppSitesControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, _ *WebAppSitesControllerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// ErrorSummary - MMethod to get error summary from web app site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - body - The content of the action request
//   - options - WebAppSitesControllerClientErrorSummaryOptions contains the optional parameters for the WebAppSitesControllerClient.ErrorSummary
//     method.
func (client *WebAppSitesControllerClient) ErrorSummary(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ErrorSummaryRequest, options *WebAppSitesControllerClientErrorSummaryOptions) (WebAppSitesControllerClientErrorSummaryResponse, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.ErrorSummary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.errorSummaryCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, body, options)
	if err != nil {
		return WebAppSitesControllerClientErrorSummaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppSitesControllerClientErrorSummaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppSitesControllerClientErrorSummaryResponse{}, err
	}
	resp, err := client.errorSummaryHandleResponse(httpResp)
	return resp, err
}

// errorSummaryCreateRequest creates the ErrorSummary request.
func (client *WebAppSitesControllerClient) errorSummaryCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ErrorSummaryRequest, _ *WebAppSitesControllerClientErrorSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/errorSummary"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// errorSummaryHandleResponse handles the ErrorSummary response.
func (client *WebAppSitesControllerClient) errorSummaryHandleResponse(resp *http.Response) (WebAppSitesControllerClientErrorSummaryResponse, error) {
	result := WebAppSitesControllerClientErrorSummaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteErrorSummary); err != nil {
		return WebAppSitesControllerClientErrorSummaryResponse{}, err
	}
	return result, nil
}

// BeginExportInventory - Method to generate report containing web app inventory.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - body - The content of the action request
//   - options - WebAppSitesControllerClientBeginExportInventoryOptions contains the optional parameters for the WebAppSitesControllerClient.BeginExportInventory
//     method.
func (client *WebAppSitesControllerClient) BeginExportInventory(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ExportWebAppsRequest, options *WebAppSitesControllerClientBeginExportInventoryOptions) (*runtime.Poller[WebAppSitesControllerClientExportInventoryResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportInventory(ctx, resourceGroupName, siteName, webAppSiteName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppSitesControllerClientExportInventoryResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppSitesControllerClientExportInventoryResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportInventory - Method to generate report containing web app inventory.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
func (client *WebAppSitesControllerClient) exportInventory(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ExportWebAppsRequest, options *WebAppSitesControllerClientBeginExportInventoryOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.BeginExportInventory"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportInventoryCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, body, options)
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

// exportInventoryCreateRequest creates the ExportInventory request.
func (client *WebAppSitesControllerClient) exportInventoryCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ExportWebAppsRequest, _ *WebAppSitesControllerClientBeginExportInventoryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/exportInventory"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Method to get a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebAppSitesControllerClientGetOptions contains the optional parameters for the WebAppSitesControllerClient.Get
//     method.
func (client *WebAppSitesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *WebAppSitesControllerClientGetOptions) (WebAppSitesControllerClientGetResponse, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
	if err != nil {
		return WebAppSitesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppSitesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppSitesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WebAppSitesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, _ *WebAppSitesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}"
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
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebAppSitesControllerClient) getHandleResponse(resp *http.Response) (WebAppSitesControllerClientGetResponse, error) {
	result := WebAppSitesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppSite); err != nil {
		return WebAppSitesControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMasterSitePager - Method to get all sites.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - WebAppSitesControllerClientListByMasterSiteOptions contains the optional parameters for the WebAppSitesControllerClient.NewListByMasterSitePager
//     method.
func (client *WebAppSitesControllerClient) NewListByMasterSitePager(resourceGroupName string, siteName string, options *WebAppSitesControllerClientListByMasterSiteOptions) *runtime.Pager[WebAppSitesControllerClientListByMasterSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[WebAppSitesControllerClientListByMasterSiteResponse]{
		More: func(page WebAppSitesControllerClientListByMasterSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WebAppSitesControllerClientListByMasterSiteResponse) (WebAppSitesControllerClientListByMasterSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WebAppSitesControllerClient.NewListByMasterSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMasterSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return WebAppSitesControllerClientListByMasterSiteResponse{}, err
			}
			return client.listByMasterSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMasterSiteCreateRequest creates the ListByMasterSite request.
func (client *WebAppSitesControllerClient) listByMasterSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, _ *WebAppSitesControllerClientListByMasterSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites"
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
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMasterSiteHandleResponse handles the ListByMasterSite response.
func (client *WebAppSitesControllerClient) listByMasterSiteHandleResponse(resp *http.Response) (WebAppSitesControllerClientListByMasterSiteResponse, error) {
	result := WebAppSitesControllerClientListByMasterSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppSiteListResult); err != nil {
		return WebAppSitesControllerClientListByMasterSiteResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Method to refresh a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - body - The content of the action request
//   - options - WebAppSitesControllerClientBeginRefreshOptions contains the optional parameters for the WebAppSitesControllerClient.BeginRefresh
//     method.
func (client *WebAppSitesControllerClient) BeginRefresh(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ProxySiteRefreshBody, options *WebAppSitesControllerClientBeginRefreshOptions) (*runtime.Poller[WebAppSitesControllerClientRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refresh(ctx, resourceGroupName, siteName, webAppSiteName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppSitesControllerClientRefreshResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppSitesControllerClientRefreshResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Refresh - Method to refresh a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
func (client *WebAppSitesControllerClient) refresh(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ProxySiteRefreshBody, options *WebAppSitesControllerClientBeginRefreshOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.BeginRefresh"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, body, options)
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

// refreshCreateRequest creates the Refresh request.
func (client *WebAppSitesControllerClient) refreshCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, body ProxySiteRefreshBody, _ *WebAppSitesControllerClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/refresh"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Summary - Method to get site usage/summary.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - options - WebAppSitesControllerClientSummaryOptions contains the optional parameters for the WebAppSitesControllerClient.Summary
//     method.
func (client *WebAppSitesControllerClient) Summary(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, options *WebAppSitesControllerClientSummaryOptions) (WebAppSitesControllerClientSummaryResponse, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.Summary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.summaryCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, options)
	if err != nil {
		return WebAppSitesControllerClientSummaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WebAppSitesControllerClientSummaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WebAppSitesControllerClientSummaryResponse{}, err
	}
	resp, err := client.summaryHandleResponse(httpResp)
	return resp, err
}

// summaryCreateRequest creates the Summary request.
func (client *WebAppSitesControllerClient) summaryCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, _ *WebAppSitesControllerClientSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}/summary"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// summaryHandleResponse handles the Summary response.
func (client *WebAppSitesControllerClient) summaryHandleResponse(resp *http.Response) (WebAppSitesControllerClientSummaryResponse, error) {
	result := WebAppSitesControllerClientSummaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WebAppSiteUsage); err != nil {
		return WebAppSitesControllerClientSummaryResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Method to update an existing site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - webAppSiteName - Web app site name.
//   - properties - The resource properties to be updated.
//   - options - WebAppSitesControllerClientBeginUpdateOptions contains the optional parameters for the WebAppSitesControllerClient.BeginUpdate
//     method.
func (client *WebAppSitesControllerClient) BeginUpdate(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, properties WebAppSiteUpdate, options *WebAppSitesControllerClientBeginUpdateOptions) (*runtime.Poller[WebAppSitesControllerClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, siteName, webAppSiteName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[WebAppSitesControllerClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[WebAppSitesControllerClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Method to update an existing site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
func (client *WebAppSitesControllerClient) update(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, properties WebAppSiteUpdate, options *WebAppSitesControllerClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "WebAppSitesControllerClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, siteName, webAppSiteName, properties, options)
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

// updateCreateRequest creates the Update request.
func (client *WebAppSitesControllerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, siteName string, webAppSiteName string, properties WebAppSiteUpdate, _ *WebAppSitesControllerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/webAppSites/{webAppSiteName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
