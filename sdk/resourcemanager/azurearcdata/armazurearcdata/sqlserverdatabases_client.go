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

// SQLServerDatabasesClient contains the methods for the SQLServerDatabases group.
// Don't use this type directly, use NewSQLServerDatabasesClient() instead.
type SQLServerDatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSQLServerDatabasesClient creates a new instance of SQLServerDatabasesClient with the specified values.
//   - subscriptionID - The ID of the Azure subscription
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSQLServerDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SQLServerDatabasesClient, error) {
	cl, err := arm.NewClient(moduleName+".SQLServerDatabasesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SQLServerDatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Creates or replaces an Arc Sql Server Database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlServerInstanceName - Name of SQL Server Instance
//   - databaseName - Name of the database
//   - sqlServerDatabaseResource - The request body for database resource.
//   - options - SQLServerDatabasesClientCreateOptions contains the optional parameters for the SQLServerDatabasesClient.Create
//     method.
func (client *SQLServerDatabasesClient) Create(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, sqlServerDatabaseResource SQLServerDatabaseResource, options *SQLServerDatabasesClientCreateOptions) (SQLServerDatabasesClientCreateResponse, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, databaseName, sqlServerDatabaseResource, options)
	if err != nil {
		return SQLServerDatabasesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerDatabasesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerDatabasesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SQLServerDatabasesClient) createCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, sqlServerDatabaseResource SQLServerDatabaseResource, options *SQLServerDatabasesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sqlServerDatabaseResource); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SQLServerDatabasesClient) createHandleResponse(resp *http.Response) (SQLServerDatabasesClientCreateResponse, error) {
	result := SQLServerDatabasesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerDatabaseResource); err != nil {
		return SQLServerDatabasesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Arc Sql Server database resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlServerInstanceName - Name of SQL Server Instance
//   - databaseName - Name of the database
//   - options - SQLServerDatabasesClientDeleteOptions contains the optional parameters for the SQLServerDatabasesClient.Delete
//     method.
func (client *SQLServerDatabasesClient) Delete(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, options *SQLServerDatabasesClientDeleteOptions) (SQLServerDatabasesClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, databaseName, options)
	if err != nil {
		return SQLServerDatabasesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerDatabasesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerDatabasesClientDeleteResponse{}, err
	}
	return SQLServerDatabasesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SQLServerDatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, options *SQLServerDatabasesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves an Arc Sql Server database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlServerInstanceName - Name of SQL Server Instance
//   - databaseName - Name of the database
//   - options - SQLServerDatabasesClientGetOptions contains the optional parameters for the SQLServerDatabasesClient.Get method.
func (client *SQLServerDatabasesClient) Get(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, options *SQLServerDatabasesClientGetOptions) (SQLServerDatabasesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, databaseName, options)
	if err != nil {
		return SQLServerDatabasesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerDatabasesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SQLServerDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, options *SQLServerDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SQLServerDatabasesClient) getHandleResponse(resp *http.Response) (SQLServerDatabasesClientGetResponse, error) {
	result := SQLServerDatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerDatabaseResource); err != nil {
		return SQLServerDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the databases associated with the given Arc Sql Server.
//
// Generated from API version 2023-01-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlServerInstanceName - Name of SQL Server Instance
//   - options - SQLServerDatabasesClientListOptions contains the optional parameters for the SQLServerDatabasesClient.NewListPager
//     method.
func (client *SQLServerDatabasesClient) NewListPager(resourceGroupName string, sqlServerInstanceName string, options *SQLServerDatabasesClientListOptions) *runtime.Pager[SQLServerDatabasesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SQLServerDatabasesClientListResponse]{
		More: func(page SQLServerDatabasesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SQLServerDatabasesClientListResponse) (SQLServerDatabasesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SQLServerDatabasesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SQLServerDatabasesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SQLServerDatabasesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SQLServerDatabasesClient) listCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, options *SQLServerDatabasesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}/databases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SQLServerDatabasesClient) listHandleResponse(resp *http.Response) (SQLServerDatabasesClientListResponse, error) {
	result := SQLServerDatabasesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ArcSQLServerDatabaseListResult); err != nil {
		return SQLServerDatabasesClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-15-preview
//   - resourceGroupName - The name of the Azure resource group
//   - sqlServerInstanceName - Name of SQL Server Instance
//   - databaseName - Name of the database
//   - sqlServerDatabaseUpdate - The requested database resource state.
//   - options - SQLServerDatabasesClientUpdateOptions contains the optional parameters for the SQLServerDatabasesClient.Update
//     method.
func (client *SQLServerDatabasesClient) Update(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, sqlServerDatabaseUpdate SQLServerDatabaseUpdate, options *SQLServerDatabasesClientUpdateOptions) (SQLServerDatabasesClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, sqlServerInstanceName, databaseName, sqlServerDatabaseUpdate, options)
	if err != nil {
		return SQLServerDatabasesClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SQLServerDatabasesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SQLServerDatabasesClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *SQLServerDatabasesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, sqlServerInstanceName string, databaseName string, sqlServerDatabaseUpdate SQLServerDatabaseUpdate, options *SQLServerDatabasesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureArcData/sqlServerInstances/{sqlServerInstanceName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlServerInstanceName == "" {
		return nil, errors.New("parameter sqlServerInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlServerInstanceName}", url.PathEscape(sqlServerInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, sqlServerDatabaseUpdate); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *SQLServerDatabasesClient) updateHandleResponse(resp *http.Response) (SQLServerDatabasesClientUpdateResponse, error) {
	result := SQLServerDatabasesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SQLServerDatabaseResource); err != nil {
		return SQLServerDatabasesClientUpdateResponse{}, err
	}
	return result, nil
}
