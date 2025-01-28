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

// SQLAvailabilityGroupsControllerClient contains the methods for the SQLAvailabilityGroupsController group.
// Don't use this type directly, use NewSQLAvailabilityGroupsControllerClient() instead.
type SQLAvailabilityGroupsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLAvailabilityGroupsControllerClient creates a new instance of SQLAvailabilityGroupsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLAvailabilityGroupsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLAvailabilityGroupsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLAvailabilityGroupsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the sql availability group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - sqlAvailabilityGroupName - SQL availability group name.
//   - options - SQLAvailabilityGroupsControllerClientGetOptions contains the optional parameters for the SQLAvailabilityGroupsControllerClient.Get
//     method.
func (client *SQLAvailabilityGroupsControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, sqlAvailabilityGroupName string, options *SQLAvailabilityGroupsControllerClientGetOptions) (SQLAvailabilityGroupsControllerClientGetResponse, error) {
	var err error
	const operationName = "SQLAvailabilityGroupsControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, sqlAvailabilityGroupName, options)
	if err != nil {
		return SQLAvailabilityGroupsControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLAvailabilityGroupsControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLAvailabilityGroupsControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLAvailabilityGroupsControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, sqlAvailabilityGroupName string, options *SQLAvailabilityGroupsControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/sqlAvailabilityGroups/{sqlAvailabilityGroupName}"
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
	if sqlAvailabilityGroupName == "" {
		return nil, errors.New("parameter sqlAvailabilityGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlAvailabilityGroupName}", url.PathEscape(sqlAvailabilityGroupName))
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
func (client *SQLAvailabilityGroupsControllerClient) getHandleResponse(resp *http.Response) (SQLAvailabilityGroupsControllerClientGetResponse, error) {
	result := SQLAvailabilityGroupsControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLAvailabilityGroup); err != nil {
		return SQLAvailabilityGroupsControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySQLSitePager - Gets the sql availability groups.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - sqlSiteName - SQL site name.
//   - options - SQLAvailabilityGroupsControllerClientListBySQLSiteOptions contains the optional parameters for the SQLAvailabilityGroupsControllerClient.NewListBySQLSitePager
//     method.
func (client *SQLAvailabilityGroupsControllerClient) NewListBySQLSitePager(resourceGroupName string, siteName string, sqlSiteName string, options *SQLAvailabilityGroupsControllerClientListBySQLSiteOptions) *runtime.Pager[SQLAvailabilityGroupsControllerClientListBySQLSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLAvailabilityGroupsControllerClientListBySQLSiteResponse]{
		More: func(page SQLAvailabilityGroupsControllerClientListBySQLSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLAvailabilityGroupsControllerClientListBySQLSiteResponse) (SQLAvailabilityGroupsControllerClientListBySQLSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLAvailabilityGroupsControllerClient.NewListBySQLSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySQLSiteCreateRequest(ctx, resourceGroupName, siteName, sqlSiteName, options)
			}, nil)
			if err != nil {
				return SQLAvailabilityGroupsControllerClientListBySQLSiteResponse{}, err
			}
			return client.listBySQLSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySQLSiteCreateRequest creates the ListBySQLSite request.
func (client *SQLAvailabilityGroupsControllerClient) listBySQLSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, sqlSiteName string, options *SQLAvailabilityGroupsControllerClientListBySQLSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/sqlSites/{sqlSiteName}/sqlAvailabilityGroups"
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

// listBySQLSiteHandleResponse handles the ListBySQLSite response.
func (client *SQLAvailabilityGroupsControllerClient) listBySQLSiteHandleResponse(resp *http.Response) (SQLAvailabilityGroupsControllerClientListBySQLSiteResponse, error) {
	result := SQLAvailabilityGroupsControllerClientListBySQLSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLAvailabilityGroupListResult); err != nil {
		return SQLAvailabilityGroupsControllerClientListBySQLSiteResponse{}, err
	}
	return result, nil
}
