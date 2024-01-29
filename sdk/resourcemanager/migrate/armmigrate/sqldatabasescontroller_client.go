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

// SQLDatabasesControllerClient contains the methods for the SQLDatabasesController group.
// Don't use this type directly, use NewSQLDatabasesControllerClient() instead.
type SQLDatabasesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLDatabasesControllerClient creates a new instance of SQLDatabasesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLDatabasesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLDatabasesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLDatabasesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the sql Database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - sqlDatabaseName - SQL Database name.
//   - options - SQLDatabasesControllerClientGetOptions contains the optional parameters for the SQLDatabasesControllerClient.Get
//     method.
func (client *SQLDatabasesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, sqlDatabaseName string, options *SQLDatabasesControllerClientGetOptions) (SQLDatabasesControllerClientGetResponse, error) {
	var err error
	const operationName = "SQLDatabasesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, sqlDatabaseName, options)
	if err != nil {
		return SQLDatabasesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLDatabasesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLDatabasesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLDatabasesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, sqlDatabaseName string, options *SQLDatabasesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/sqlDatabases/{sqlDatabaseName}"
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
	if sqlDatabaseName == "" {
		return nil, errors.New("parameter sqlDatabaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDatabaseName}", url.PathEscape(sqlDatabaseName))
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
func (client *SQLDatabasesControllerClient) getHandleResponse(resp *http.Response) (SQLDatabasesControllerClientGetResponse, error) {
	result := SQLDatabasesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLDatabaseV2); err != nil {
		return SQLDatabasesControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySQLSitePager - Gets the sql Databases.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - options - SQLDatabasesControllerClientListBySQLSiteOptions contains the optional parameters for the SQLDatabasesControllerClient.NewListBySQLSitePager
//     method.
func (client *SQLDatabasesControllerClient) NewListBySQLSitePager(resourceGroupName string, siteName string, sqlSiteName string, options *SQLDatabasesControllerClientListBySQLSiteOptions) *runtime.Pager[SQLDatabasesControllerClientListBySQLSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLDatabasesControllerClientListBySQLSiteResponse]{
		More: func(page SQLDatabasesControllerClientListBySQLSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLDatabasesControllerClientListBySQLSiteResponse) (SQLDatabasesControllerClientListBySQLSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLDatabasesControllerClient.NewListBySQLSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySQLSiteCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, options)
			}, nil)
			if err != nil {
				return SQLDatabasesControllerClientListBySQLSiteResponse{}, err
			}
			return client.listBySQLSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySQLSiteCreateRequest creates the ListBySQLSite request.
func (client *SQLDatabasesControllerClient) listBySQLSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLDatabasesControllerClientListBySQLSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/sqlDatabases"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("top", *options.Top)
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	if options != nil && options.TotalRecordCount != nil {
		reqQP.Set("totalRecordCount", strconv.FormatInt(int64(*options.TotalRecordCount), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySQLSiteHandleResponse handles the ListBySQLSite response.
func (client *SQLDatabasesControllerClient) listBySQLSiteHandleResponse(resp *http.Response) (SQLDatabasesControllerClientListBySQLSiteResponse, error) {
	result := SQLDatabasesControllerClientListBySQLSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLDatabaseV2ListResult); err != nil {
		return SQLDatabasesControllerClientListBySQLSiteResponse{}, err
	}
	return result, nil
}
