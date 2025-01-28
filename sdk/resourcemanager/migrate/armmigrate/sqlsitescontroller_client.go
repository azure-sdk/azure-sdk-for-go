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
	"strconv"
	"strings"
)

// SQLSitesControllerClient contains the methods for the SQLSitesController group.
// Don't use this type directly, use NewSQLSitesControllerClient() instead.
type SQLSitesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLSitesControllerClient creates a new instance of SQLSitesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLSitesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLSitesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLSitesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Method to create a SQL site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - body - Resource create parameters.
//   - options - SQLSitesControllerClientCreateOptions contains the optional parameters for the SQLSitesControllerClient.Create
//     method.
func (client *SQLSitesControllerClient) Create(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body SQLSite, options *SQLSitesControllerClientCreateOptions) (SQLSitesControllerClientCreateResponse, error) {
	var err error
	const operationName = "SQLSitesControllerClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
	if err != nil {
		return SQLSitesControllerClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLSitesControllerClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SQLSitesControllerClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SQLSitesControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body SQLSite, options *SQLSitesControllerClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
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

// createHandleResponse handles the Create response.
func (client *SQLSitesControllerClient) createHandleResponse(resp *http.Response) (SQLSitesControllerClientCreateResponse, error) {
	result := SQLSitesControllerClientCreateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return SQLSitesControllerClientCreateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLSite); err != nil {
		return SQLSitesControllerClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the SQL site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - options - SQLSitesControllerClientDeleteOptions contains the optional parameters for the SQLSitesControllerClient.Delete
//     method.
func (client *SQLSitesControllerClient) Delete(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLSitesControllerClientDeleteOptions) (SQLSitesControllerClientDeleteResponse, error) {
	var err error
	const operationName = "SQLSitesControllerClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, options)
	if err != nil {
		return SQLSitesControllerClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLSitesControllerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SQLSitesControllerClientDeleteResponse{}, err
	}
	return SQLSitesControllerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLSitesControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLSitesControllerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
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

// ErrorSummary - Method to get error summary from SQL site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - body - The content of the action request
//   - options - SQLSitesControllerClientErrorSummaryOptions contains the optional parameters for the SQLSitesControllerClient.ErrorSummary
//     method.
func (client *SQLSitesControllerClient) ErrorSummary(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body ErrorSummaryRequest, options *SQLSitesControllerClientErrorSummaryOptions) (SQLSitesControllerClientErrorSummaryResponse, error) {
	var err error
	const operationName = "SQLSitesControllerClient.ErrorSummary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.errorSummaryCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
	if err != nil {
		return SQLSitesControllerClientErrorSummaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLSitesControllerClientErrorSummaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLSitesControllerClientErrorSummaryResponse{}, err
	}
	resp, err := client.errorSummaryHandleResponse(httpResp)
	return resp, err
}

// errorSummaryCreateRequest creates the ErrorSummary request.
func (client *SQLSitesControllerClient) errorSummaryCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body ErrorSummaryRequest, options *SQLSitesControllerClientErrorSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/errorSummary"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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

// errorSummaryHandleResponse handles the ErrorSummary response.
func (client *SQLSitesControllerClient) errorSummaryHandleResponse(resp *http.Response) (SQLSitesControllerClientErrorSummaryResponse, error) {
	result := SQLSitesControllerClientErrorSummaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SiteErrorSummary); err != nil {
		return SQLSitesControllerClientErrorSummaryResponse{}, err
	}
	return result, nil
}

// BeginExportSQLServerErrors - Method to generate report containing SQL servers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - body - The content of the action request
//   - options - SQLSitesControllerClientBeginExportSQLServerErrorsOptions contains the optional parameters for the SQLSitesControllerClient.BeginExportSQLServerErrors
//     method.
func (client *SQLSitesControllerClient) BeginExportSQLServerErrors(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body any, options *SQLSitesControllerClientBeginExportSQLServerErrorsOptions) (*runtime.Poller[SQLSitesControllerClientExportSQLServerErrorsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportSQLServerErrors(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLSitesControllerClientExportSQLServerErrorsResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLSitesControllerClientExportSQLServerErrorsResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportSQLServerErrors - Method to generate report containing SQL servers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *SQLSitesControllerClient) exportSQLServerErrors(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body any, options *SQLSitesControllerClientBeginExportSQLServerErrorsOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLSitesControllerClient.BeginExportSQLServerErrors"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportSQLServerErrorsCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
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

// exportSQLServerErrorsCreateRequest creates the ExportSQLServerErrors request.
func (client *SQLSitesControllerClient) exportSQLServerErrorsCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body any, options *SQLSitesControllerClientBeginExportSQLServerErrorsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/exportSqlServerErrors"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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

// BeginExportSQLServers - Method to generate report containing SQL servers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - body - The content of the action request
//   - options - SQLSitesControllerClientBeginExportSQLServersOptions contains the optional parameters for the SQLSitesControllerClient.BeginExportSQLServers
//     method.
func (client *SQLSitesControllerClient) BeginExportSQLServers(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body ExportSQLServersRequest, options *SQLSitesControllerClientBeginExportSQLServersOptions) (*runtime.Poller[SQLSitesControllerClientExportSQLServersResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportSQLServers(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLSitesControllerClientExportSQLServersResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLSitesControllerClientExportSQLServersResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ExportSQLServers - Method to generate report containing SQL servers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *SQLSitesControllerClient) exportSQLServers(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body ExportSQLServersRequest, options *SQLSitesControllerClientBeginExportSQLServersOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLSitesControllerClient.BeginExportSQLServers"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.exportSQLServersCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
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

// exportSQLServersCreateRequest creates the ExportSQLServers request.
func (client *SQLSitesControllerClient) exportSQLServersCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body ExportSQLServersRequest, options *SQLSitesControllerClientBeginExportSQLServersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/exportSqlServers"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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

// Get - Method to get a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - options - SQLSitesControllerClientGetOptions contains the optional parameters for the SQLSitesControllerClient.Get method.
func (client *SQLSitesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLSitesControllerClientGetOptions) (SQLSitesControllerClientGetResponse, error) {
	var err error
	const operationName = "SQLSitesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, options)
	if err != nil {
		return SQLSitesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLSitesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLSitesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLSitesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLSitesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
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
func (client *SQLSitesControllerClient) getHandleResponse(resp *http.Response) (SQLSitesControllerClientGetResponse, error) {
	result := SQLSitesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLSite); err != nil {
		return SQLSitesControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMasterSitePager - Method to get all sites.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - SQLSitesControllerClientListByMasterSiteOptions contains the optional parameters for the SQLSitesControllerClient.NewListByMasterSitePager
//     method.
func (client *SQLSitesControllerClient) NewListByMasterSitePager(resourceGroupName string, siteName string, options *SQLSitesControllerClientListByMasterSiteOptions) *runtime.Pager[SQLSitesControllerClientListByMasterSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLSitesControllerClientListByMasterSiteResponse]{
		More: func(page SQLSitesControllerClientListByMasterSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLSitesControllerClientListByMasterSiteResponse) (SQLSitesControllerClientListByMasterSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLSitesControllerClient.NewListByMasterSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMasterSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return SQLSitesControllerClientListByMasterSiteResponse{}, err
			}
			return client.listByMasterSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMasterSiteCreateRequest creates the ListByMasterSite request.
func (client *SQLSitesControllerClient) listByMasterSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *SQLSitesControllerClientListByMasterSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites"
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
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMasterSiteHandleResponse handles the ListByMasterSite response.
func (client *SQLSitesControllerClient) listByMasterSiteHandleResponse(resp *http.Response) (SQLSitesControllerClientListByMasterSiteResponse, error) {
	result := SQLSitesControllerClientListByMasterSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLSiteListResult); err != nil {
		return SQLSitesControllerClientListByMasterSiteResponse{}, err
	}
	return result, nil
}

// BeginRefresh - Method to refresh a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - body - The content of the action request
//   - options - SQLSitesControllerClientBeginRefreshOptions contains the optional parameters for the SQLSitesControllerClient.BeginRefresh
//     method.
func (client *SQLSitesControllerClient) BeginRefresh(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body SQLSiteRefreshBody, options *SQLSitesControllerClientBeginRefreshOptions) (*runtime.Poller[SQLSitesControllerClientRefreshResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refresh(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLSitesControllerClientRefreshResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLSitesControllerClientRefreshResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Refresh - Method to refresh a site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *SQLSitesControllerClient) refresh(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body SQLSiteRefreshBody, options *SQLSitesControllerClientBeginRefreshOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLSitesControllerClient.BeginRefresh"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.refreshCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, body, options)
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
func (client *SQLSitesControllerClient) refreshCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, body SQLSiteRefreshBody, options *SQLSitesControllerClientBeginRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/refresh"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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

// Summary - Method to get site usage/summary.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - options - SQLSitesControllerClientSummaryOptions contains the optional parameters for the SQLSitesControllerClient.Summary
//     method.
func (client *SQLSitesControllerClient) Summary(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLSitesControllerClientSummaryOptions) (SQLSitesControllerClientSummaryResponse, error) {
	var err error
	const operationName = "SQLSitesControllerClient.Summary"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.summaryCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, options)
	if err != nil {
		return SQLSitesControllerClientSummaryResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLSitesControllerClientSummaryResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLSitesControllerClientSummaryResponse{}, err
	}
	resp, err := client.summaryHandleResponse(httpResp)
	return resp, err
}

// summaryCreateRequest creates the Summary request.
func (client *SQLSitesControllerClient) summaryCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLSitesControllerClientSummaryOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/summary"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// summaryHandleResponse handles the Summary response.
func (client *SQLSitesControllerClient) summaryHandleResponse(resp *http.Response) (SQLSitesControllerClientSummaryResponse, error) {
	result := SQLSitesControllerClientSummaryResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLSiteUsage); err != nil {
		return SQLSitesControllerClientSummaryResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Method to update an existing site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - properties - The resource properties to be updated.
//   - options - SQLSitesControllerClientBeginUpdateOptions contains the optional parameters for the SQLSitesControllerClient.BeginUpdate
//     method.
func (client *SQLSitesControllerClient) BeginUpdate(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, properties SQLSiteUpdate, options *SQLSitesControllerClientBeginUpdateOptions) (*runtime.Poller[SQLSitesControllerClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, siteName, sqlSiteName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SQLSitesControllerClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SQLSitesControllerClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Method to update an existing site.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
func (client *SQLSitesControllerClient) update(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, properties SQLSiteUpdate, options *SQLSitesControllerClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SQLSitesControllerClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, properties, options)
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
func (client *SQLSitesControllerClient) updateCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, properties SQLSiteUpdate, options *SQLSitesControllerClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}"
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
	if sqlSiteName == "" {
		return nil, errors.New("parameter sqlSiteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlSiteName}", url.PathEscape(sqlSiteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
