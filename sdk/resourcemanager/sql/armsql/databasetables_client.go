//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// DatabaseTablesClient contains the methods for the DatabaseTables group.
// Don't use this type directly, use NewDatabaseTablesClient() instead.
type DatabaseTablesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseTablesClient creates a new instance of DatabaseTablesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseTablesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseTablesClient, error) {
	cl, err := arm.NewClient(moduleName+".DatabaseTablesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseTablesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get database table
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - options - DatabaseTablesClientGetOptions contains the optional parameters for the DatabaseTablesClient.Get method.
func (client *DatabaseTablesClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *DatabaseTablesClientGetOptions) (DatabaseTablesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, options)
	if err != nil {
		return DatabaseTablesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseTablesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabaseTablesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabaseTablesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, options *DatabaseTablesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
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
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseTablesClient) getHandleResponse(resp *http.Response) (DatabaseTablesClientGetResponse, error) {
	result := DatabaseTablesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseTable); err != nil {
		return DatabaseTablesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySchemaPager - List database tables
//
// Generated from API version 2023-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - options - DatabaseTablesClientListBySchemaOptions contains the optional parameters for the DatabaseTablesClient.NewListBySchemaPager
//     method.
func (client *DatabaseTablesClient) NewListBySchemaPager(resourceGroupName string, serverName string, databaseName string, schemaName string, options *DatabaseTablesClientListBySchemaOptions) *runtime.Pager[DatabaseTablesClientListBySchemaResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseTablesClientListBySchemaResponse]{
		More: func(page DatabaseTablesClientListBySchemaResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabaseTablesClientListBySchemaResponse) (DatabaseTablesClientListBySchemaResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySchemaCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DatabaseTablesClientListBySchemaResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DatabaseTablesClientListBySchemaResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabaseTablesClientListBySchemaResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySchemaHandleResponse(resp)
		},
	})
}

// listBySchemaCreateRequest creates the ListBySchema request.
func (client *DatabaseTablesClient) listBySchemaCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, options *DatabaseTablesClientListBySchemaOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if schemaName == "" {
		return nil, errors.New("parameter schemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{schemaName}", url.PathEscape(schemaName))
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
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySchemaHandleResponse handles the ListBySchema response.
func (client *DatabaseTablesClient) listBySchemaHandleResponse(resp *http.Response) (DatabaseTablesClientListBySchemaResponse, error) {
	result := DatabaseTablesClientListBySchemaResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseTableListResult); err != nil {
		return DatabaseTablesClientListBySchemaResponse{}, err
	}
	return result, nil
}
