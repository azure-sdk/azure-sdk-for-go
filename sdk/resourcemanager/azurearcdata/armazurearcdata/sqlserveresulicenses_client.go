//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurearcdata

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

// SQLServerEsuLicensesClient contains the methods for the SQLServerEsuLicenses group.
// Don't use this type directly, use NewSQLServerEsuLicensesClient() instead.
type SQLServerEsuLicensesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLServerEsuLicensesClient creates a new instance of SQLServerEsuLicensesClient with the specified values.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLServerEsuLicensesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLServerEsuLicensesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLServerEsuLicensesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or replaces a SQL Server ESU license resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sqlServerEsuLicenseName - Name of SQL Server ESU License
//   - sqlServerEsuLicense - The SQL Server ESU license to be created or updated.
//   - options - SQLServerEsuLicensesClientCreateOptions contains the optional parameters for the SQLServerEsuLicensesClient.Create
//     method.
func (client *SQLServerEsuLicensesClient) Create(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, sqlServerEsuLicense SQLServerEsuLicense, options *SQLServerEsuLicensesClientCreateOptions) (SQLServerEsuLicensesClientCreateResponse, error) {
	var err error
	const operationName = "SQLServerEsuLicensesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, sqlServerEsuLicenseName, sqlServerEsuLicense, options)
	if err != nil {
		return SQLServerEsuLicensesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerEsuLicensesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerEsuLicensesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SQLServerEsuLicensesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, sqlServerEsuLicense SQLServerEsuLicense, options *SQLServerEsuLicensesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerEsuLicenses/{sqlServerEsuLicenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerEsuLicenseName == "" {
		return nil, errors.New("parameter sqlServerEsuLicenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerEsuLicenseName}", url.PathEscape(sqlServerEsuLicenseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sqlServerEsuLicense); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SQLServerEsuLicensesClient) createHandleResponse(resp *http.Response) (SQLServerEsuLicensesClientCreateResponse, error) {
	result := SQLServerEsuLicensesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerEsuLicense); err != nil {
		return SQLServerEsuLicensesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a SQL Server ESU license resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sqlServerEsuLicenseName - Name of SQL Server ESU License
//   - options - SQLServerEsuLicensesClientDeleteOptions contains the optional parameters for the SQLServerEsuLicensesClient.Delete
//     method.
func (client *SQLServerEsuLicensesClient) Delete(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, options *SQLServerEsuLicensesClientDeleteOptions) (SQLServerEsuLicensesClientDeleteResponse, error) {
	var err error
	const operationName = "SQLServerEsuLicensesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlServerEsuLicenseName, options)
	if err != nil {
		return SQLServerEsuLicensesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerEsuLicensesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerEsuLicensesClientDeleteResponse{}, err
	}
	return SQLServerEsuLicensesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLServerEsuLicensesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, options *SQLServerEsuLicensesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerEsuLicenses/{sqlServerEsuLicenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerEsuLicenseName == "" {
		return nil, errors.New("parameter sqlServerEsuLicenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerEsuLicenseName}", url.PathEscape(sqlServerEsuLicenseName))
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

// Get - Retrieves a SQL Server ESU license resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sqlServerEsuLicenseName - Name of SQL Server ESU License
//   - options - SQLServerEsuLicensesClientGetOptions contains the optional parameters for the SQLServerEsuLicensesClient.Get
//     method.
func (client *SQLServerEsuLicensesClient) Get(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, options *SQLServerEsuLicensesClientGetOptions) (SQLServerEsuLicensesClientGetResponse, error) {
	var err error
	const operationName = "SQLServerEsuLicensesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlServerEsuLicenseName, options)
	if err != nil {
		return SQLServerEsuLicensesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerEsuLicensesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerEsuLicensesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLServerEsuLicensesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, options *SQLServerEsuLicensesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerEsuLicenses/{sqlServerEsuLicenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerEsuLicenseName == "" {
		return nil, errors.New("parameter sqlServerEsuLicenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerEsuLicenseName}", url.PathEscape(sqlServerEsuLicenseName))
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
func (client *SQLServerEsuLicensesClient) getHandleResponse(resp *http.Response) (SQLServerEsuLicensesClientGetResponse, error) {
	result := SQLServerEsuLicensesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerEsuLicense); err != nil {
		return SQLServerEsuLicensesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List sqlServerEsuLicense resources in the subscription
//
// Generated from API version 2024-05-01-preview
//   - options - SQLServerEsuLicensesClientListOptions contains the optional parameters for the SQLServerEsuLicensesClient.NewListPager
//     method.
func (client *SQLServerEsuLicensesClient) NewListPager(options *SQLServerEsuLicensesClientListOptions) *runtime.Pager[SQLServerEsuLicensesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLServerEsuLicensesClientListResponse]{
		More: func(page SQLServerEsuLicensesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLServerEsuLicensesClientListResponse) (SQLServerEsuLicensesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLServerEsuLicensesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return SQLServerEsuLicensesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SQLServerEsuLicensesClient) listCreateRequest(ctx context.Context, options *SQLServerEsuLicensesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureArcData/sqlServerEsuLicenses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listHandleResponse handles the List response.
func (client *SQLServerEsuLicensesClient) listHandleResponse(resp *http.Response) (SQLServerEsuLicensesClientListResponse, error) {
	result := SQLServerEsuLicensesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerEsuLicenseListResult); err != nil {
		return SQLServerEsuLicensesClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all sqlServerEsuLicenses in a resource group.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - SQLServerEsuLicensesClientListByResourceGroupOptions contains the optional parameters for the SQLServerEsuLicensesClient.NewListByResourceGroupPager
//     method.
func (client *SQLServerEsuLicensesClient) NewListByResourceGroupPager(resourceGroupName string, options *SQLServerEsuLicensesClientListByResourceGroupOptions) *runtime.Pager[SQLServerEsuLicensesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLServerEsuLicensesClientListByResourceGroupResponse]{
		More: func(page SQLServerEsuLicensesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLServerEsuLicensesClientListByResourceGroupResponse) (SQLServerEsuLicensesClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SQLServerEsuLicensesClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return SQLServerEsuLicensesClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *SQLServerEsuLicensesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *SQLServerEsuLicensesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerEsuLicenses"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
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

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *SQLServerEsuLicensesClient) listByResourceGroupHandleResponse(resp *http.Response) (SQLServerEsuLicensesClientListByResourceGroupResponse, error) {
	result := SQLServerEsuLicensesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerEsuLicenseListResult); err != nil {
		return SQLServerEsuLicensesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// Update - Updates a SQL Server ESU license resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - sqlServerEsuLicenseName - Name of SQL Server ESU License
//   - parameters - The SQL Server ESU license.
//   - options - SQLServerEsuLicensesClientUpdateOptions contains the optional parameters for the SQLServerEsuLicensesClient.Update
//     method.
func (client *SQLServerEsuLicensesClient) Update(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, parameters SQLServerEsuLicenseUpdate, options *SQLServerEsuLicensesClientUpdateOptions) (SQLServerEsuLicensesClientUpdateResponse, error) {
	var err error
	const operationName = "SQLServerEsuLicensesClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlServerEsuLicenseName, parameters, options)
	if err != nil {
		return SQLServerEsuLicensesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerEsuLicensesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerEsuLicensesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SQLServerEsuLicensesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlServerEsuLicenseName string, parameters SQLServerEsuLicenseUpdate, options *SQLServerEsuLicensesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerEsuLicenses/{sqlServerEsuLicenseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerEsuLicenseName == "" {
		return nil, errors.New("parameter sqlServerEsuLicenseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerEsuLicenseName}", url.PathEscape(sqlServerEsuLicenseName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SQLServerEsuLicensesClient) updateHandleResponse(resp *http.Response) (SQLServerEsuLicensesClientUpdateResponse, error) {
	result := SQLServerEsuLicensesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerEsuLicense); err != nil {
		return SQLServerEsuLicensesClientUpdateResponse{}, err
	}
	return result, nil
}
