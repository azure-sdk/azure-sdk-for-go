//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ManagedDatabaseColumnsClient contains the methods for the ManagedDatabaseColumns group.
// Don't use this type directly, use NewManagedDatabaseColumnsClient() instead.
type ManagedDatabaseColumnsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedDatabaseColumnsClient creates a new instance of ManagedDatabaseColumnsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedDatabaseColumnsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseColumnsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedDatabaseColumnsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get managed database column
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - columnName - The name of the column.
//   - options - ManagedDatabaseColumnsClientGetOptions contains the optional parameters for the ManagedDatabaseColumnsClient.Get
//     method.
func (client *ManagedDatabaseColumnsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseColumnsClientGetOptions) (ManagedDatabaseColumnsClientGetResponse, error) {
	var err error
	const operationName = "ManagedDatabaseColumnsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return ManagedDatabaseColumnsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseColumnsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedDatabaseColumnsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseColumnsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, columnName string, options *ManagedDatabaseColumnsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ManagedDatabaseColumnsClient) getHandleResponse(resp *http.Response) (ManagedDatabaseColumnsClientGetResponse, error) {
	result := ManagedDatabaseColumnsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumn); err != nil {
		return ManagedDatabaseColumnsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - List managed database columns
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - options - ManagedDatabaseColumnsClientListByDatabaseOptions contains the optional parameters for the ManagedDatabaseColumnsClient.NewListByDatabasePager
//     method.
func (client *ManagedDatabaseColumnsClient) NewListByDatabasePager(resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseColumnsClientListByDatabaseOptions) *runtime.Pager[ManagedDatabaseColumnsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedDatabaseColumnsClientListByDatabaseResponse]{
		More: func(page ManagedDatabaseColumnsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseColumnsClientListByDatabaseResponse) (ManagedDatabaseColumnsClientListByDatabaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedDatabaseColumnsClient.NewListByDatabasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDatabaseCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, options)
			}, nil)
			if err != nil {
				return ManagedDatabaseColumnsClientListByDatabaseResponse{}, err
			}
			return client.listByDatabaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *ManagedDatabaseColumnsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, options *ManagedDatabaseColumnsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/columns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2024-02-01-preview")
	if options != nil && options.Column != nil {
		for _, qv := range options.Column {
			reqQP.Add("column", qv)
		}
	}
	if options != nil && options.OrderBy != nil {
		for _, qv := range options.OrderBy {
			reqQP.Add("orderBy", qv)
		}
	}
	if options != nil && options.Schema != nil {
		for _, qv := range options.Schema {
			reqQP.Add("schema", qv)
		}
	}
	if options != nil && options.Table != nil {
		for _, qv := range options.Table {
			reqQP.Add("table", qv)
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *ManagedDatabaseColumnsClient) listByDatabaseHandleResponse(resp *http.Response) (ManagedDatabaseColumnsClientListByDatabaseResponse, error) {
	result := ManagedDatabaseColumnsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumnListResult); err != nil {
		return ManagedDatabaseColumnsClientListByDatabaseResponse{}, err
	}
	return result, nil
}

// NewListByTablePager - List managed database columns
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - options - ManagedDatabaseColumnsClientListByTableOptions contains the optional parameters for the ManagedDatabaseColumnsClient.NewListByTablePager
//     method.
func (client *ManagedDatabaseColumnsClient) NewListByTablePager(resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, options *ManagedDatabaseColumnsClientListByTableOptions) *runtime.Pager[ManagedDatabaseColumnsClientListByTableResponse] {
	return runtime.NewPager(runtime.PagingHandler[ManagedDatabaseColumnsClientListByTableResponse]{
		More: func(page ManagedDatabaseColumnsClientListByTableResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ManagedDatabaseColumnsClientListByTableResponse) (ManagedDatabaseColumnsClientListByTableResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ManagedDatabaseColumnsClient.NewListByTablePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByTableCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, schemaName, tableName, options)
			}, nil)
			if err != nil {
				return ManagedDatabaseColumnsClientListByTableResponse{}, err
			}
			return client.listByTableHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByTableCreateRequest creates the ListByTable request.
func (client *ManagedDatabaseColumnsClient) listByTableCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, schemaName string, tableName string, options *ManagedDatabaseColumnsClientListByTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByTableHandleResponse handles the ListByTable response.
func (client *ManagedDatabaseColumnsClient) listByTableHandleResponse(resp *http.Response) (ManagedDatabaseColumnsClientListByTableResponse, error) {
	result := ManagedDatabaseColumnsClientListByTableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseColumnListResult); err != nil {
		return ManagedDatabaseColumnsClientListByTableResponse{}, err
	}
	return result, nil
}
