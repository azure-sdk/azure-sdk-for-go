//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// TableResourcesClient contains the methods for the TableResources group.
// Don't use this type directly, use NewTableResourcesClient() instead.
type TableResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTableResourcesClient creates a new instance of TableResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTableResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TableResourcesClient, error) {
	cl, err := arm.NewClient(moduleName+".TableResourcesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TableResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateUpdateTable - Create or update an Azure Cosmos DB Table
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - createUpdateTableParameters - The parameters to provide for the current Table.
//   - options - TableResourcesClientBeginCreateUpdateTableOptions contains the optional parameters for the TableResourcesClient.BeginCreateUpdateTable
//     method.
func (client *TableResourcesClient) BeginCreateUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesClientBeginCreateUpdateTableOptions) (*runtime.Poller[TableResourcesClientCreateUpdateTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createUpdateTable(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TableResourcesClientCreateUpdateTableResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TableResourcesClientCreateUpdateTableResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateUpdateTable - Create or update an Azure Cosmos DB Table
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
func (client *TableResourcesClient) createUpdateTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesClientBeginCreateUpdateTableOptions) (*http.Response, error) {
	var err error
	req, err := client.createUpdateTableCreateRequest(ctx, resourceGroupName, accountName, tableName, createUpdateTableParameters, options)
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

// createUpdateTableCreateRequest creates the CreateUpdateTable request.
func (client *TableResourcesClient) createUpdateTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, createUpdateTableParameters TableCreateUpdateParameters, options *TableResourcesClientBeginCreateUpdateTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, createUpdateTableParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDeleteTable - Deletes an existing Azure Cosmos DB Table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - options - TableResourcesClientBeginDeleteTableOptions contains the optional parameters for the TableResourcesClient.BeginDeleteTable
//     method.
func (client *TableResourcesClient) BeginDeleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginDeleteTableOptions) (*runtime.Poller[TableResourcesClientDeleteTableResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteTable(ctx, resourceGroupName, accountName, tableName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TableResourcesClientDeleteTableResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TableResourcesClientDeleteTableResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// DeleteTable - Deletes an existing Azure Cosmos DB Table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
func (client *TableResourcesClient) deleteTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginDeleteTableOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteTableCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteTableCreateRequest creates the DeleteTable request.
func (client *TableResourcesClient) deleteTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginDeleteTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// GetTable - Gets the Tables under an existing Azure Cosmos DB database account with the provided name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - options - TableResourcesClientGetTableOptions contains the optional parameters for the TableResourcesClient.GetTable method.
func (client *TableResourcesClient) GetTable(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientGetTableOptions) (TableResourcesClientGetTableResponse, error) {
	var err error
	req, err := client.getTableCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesClientGetTableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TableResourcesClientGetTableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TableResourcesClientGetTableResponse{}, err
	}
	resp, err := client.getTableHandleResponse(httpResp)
	return resp, err
}

// getTableCreateRequest creates the GetTable request.
func (client *TableResourcesClient) getTableCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientGetTableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTableHandleResponse handles the GetTable response.
func (client *TableResourcesClient) getTableHandleResponse(resp *http.Response) (TableResourcesClientGetTableResponse, error) {
	result := TableResourcesClientGetTableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TableGetResults); err != nil {
		return TableResourcesClientGetTableResponse{}, err
	}
	return result, nil
}

// GetTableThroughput - Gets the RUs per second of the Table under an existing Azure Cosmos DB database account with the provided
// name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - options - TableResourcesClientGetTableThroughputOptions contains the optional parameters for the TableResourcesClient.GetTableThroughput
//     method.
func (client *TableResourcesClient) GetTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientGetTableThroughputOptions) (TableResourcesClientGetTableThroughputResponse, error) {
	var err error
	req, err := client.getTableThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
	if err != nil {
		return TableResourcesClientGetTableThroughputResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TableResourcesClientGetTableThroughputResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TableResourcesClientGetTableThroughputResponse{}, err
	}
	resp, err := client.getTableThroughputHandleResponse(httpResp)
	return resp, err
}

// getTableThroughputCreateRequest creates the GetTableThroughput request.
func (client *TableResourcesClient) getTableThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientGetTableThroughputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getTableThroughputHandleResponse handles the GetTableThroughput response.
func (client *TableResourcesClient) getTableThroughputHandleResponse(resp *http.Response) (TableResourcesClientGetTableThroughputResponse, error) {
	result := TableResourcesClientGetTableThroughputResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThroughputSettingsGetResults); err != nil {
		return TableResourcesClientGetTableThroughputResponse{}, err
	}
	return result, nil
}

// NewListTablesPager - Lists the Tables under an existing Azure Cosmos DB database account.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - options - TableResourcesClientListTablesOptions contains the optional parameters for the TableResourcesClient.NewListTablesPager
//     method.
func (client *TableResourcesClient) NewListTablesPager(resourceGroupName string, accountName string, options *TableResourcesClientListTablesOptions) *runtime.Pager[TableResourcesClientListTablesResponse] {
	return runtime.NewPager(runtime.PagingHandler[TableResourcesClientListTablesResponse]{
		More: func(page TableResourcesClientListTablesResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *TableResourcesClientListTablesResponse) (TableResourcesClientListTablesResponse, error) {
			req, err := client.listTablesCreateRequest(ctx, resourceGroupName, accountName, options)
			if err != nil {
				return TableResourcesClientListTablesResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TableResourcesClientListTablesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TableResourcesClientListTablesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listTablesHandleResponse(resp)
		},
	})
}

// listTablesCreateRequest creates the ListTables request.
func (client *TableResourcesClient) listTablesCreateRequest(ctx context.Context, resourceGroupName string, accountName string, options *TableResourcesClientListTablesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listTablesHandleResponse handles the ListTables response.
func (client *TableResourcesClient) listTablesHandleResponse(resp *http.Response) (TableResourcesClientListTablesResponse, error) {
	result := TableResourcesClientListTablesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TableListResult); err != nil {
		return TableResourcesClientListTablesResponse{}, err
	}
	return result, nil
}

// BeginMigrateTableToAutoscale - Migrate an Azure Cosmos DB Table from manual throughput to autoscale
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - options - TableResourcesClientBeginMigrateTableToAutoscaleOptions contains the optional parameters for the TableResourcesClient.BeginMigrateTableToAutoscale
//     method.
func (client *TableResourcesClient) BeginMigrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginMigrateTableToAutoscaleOptions) (*runtime.Poller[TableResourcesClientMigrateTableToAutoscaleResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.migrateTableToAutoscale(ctx, resourceGroupName, accountName, tableName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TableResourcesClientMigrateTableToAutoscaleResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TableResourcesClientMigrateTableToAutoscaleResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// MigrateTableToAutoscale - Migrate an Azure Cosmos DB Table from manual throughput to autoscale
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
func (client *TableResourcesClient) migrateTableToAutoscale(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginMigrateTableToAutoscaleOptions) (*http.Response, error) {
	var err error
	req, err := client.migrateTableToAutoscaleCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
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

// migrateTableToAutoscaleCreateRequest creates the MigrateTableToAutoscale request.
func (client *TableResourcesClient) migrateTableToAutoscaleCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginMigrateTableToAutoscaleOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToAutoscale"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginMigrateTableToManualThroughput - Migrate an Azure Cosmos DB Table from autoscale to manual throughput
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - options - TableResourcesClientBeginMigrateTableToManualThroughputOptions contains the optional parameters for the TableResourcesClient.BeginMigrateTableToManualThroughput
//     method.
func (client *TableResourcesClient) BeginMigrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginMigrateTableToManualThroughputOptions) (*runtime.Poller[TableResourcesClientMigrateTableToManualThroughputResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.migrateTableToManualThroughput(ctx, resourceGroupName, accountName, tableName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TableResourcesClientMigrateTableToManualThroughputResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TableResourcesClientMigrateTableToManualThroughputResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// MigrateTableToManualThroughput - Migrate an Azure Cosmos DB Table from autoscale to manual throughput
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
func (client *TableResourcesClient) migrateTableToManualThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginMigrateTableToManualThroughputOptions) (*http.Response, error) {
	var err error
	req, err := client.migrateTableToManualThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, options)
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

// migrateTableToManualThroughputCreateRequest creates the MigrateTableToManualThroughput request.
func (client *TableResourcesClient) migrateTableToManualThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, options *TableResourcesClientBeginMigrateTableToManualThroughputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default/migrateToManualThroughput"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginRetrieveContinuousBackupInformation - Retrieves continuous backup information for a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - location - The name of the continuous backup restore location.
//   - options - TableResourcesClientBeginRetrieveContinuousBackupInformationOptions contains the optional parameters for the
//     TableResourcesClient.BeginRetrieveContinuousBackupInformation method.
func (client *TableResourcesClient) BeginRetrieveContinuousBackupInformation(ctx context.Context, resourceGroupName string, accountName string, tableName string, location ContinuousBackupRestoreLocation, options *TableResourcesClientBeginRetrieveContinuousBackupInformationOptions) (*runtime.Poller[TableResourcesClientRetrieveContinuousBackupInformationResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.retrieveContinuousBackupInformation(ctx, resourceGroupName, accountName, tableName, location, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TableResourcesClientRetrieveContinuousBackupInformationResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TableResourcesClientRetrieveContinuousBackupInformationResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RetrieveContinuousBackupInformation - Retrieves continuous backup information for a table.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
func (client *TableResourcesClient) retrieveContinuousBackupInformation(ctx context.Context, resourceGroupName string, accountName string, tableName string, location ContinuousBackupRestoreLocation, options *TableResourcesClientBeginRetrieveContinuousBackupInformationOptions) (*http.Response, error) {
	var err error
	req, err := client.retrieveContinuousBackupInformationCreateRequest(ctx, resourceGroupName, accountName, tableName, location, options)
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

// retrieveContinuousBackupInformationCreateRequest creates the RetrieveContinuousBackupInformation request.
func (client *TableResourcesClient) retrieveContinuousBackupInformationCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, location ContinuousBackupRestoreLocation, options *TableResourcesClientBeginRetrieveContinuousBackupInformationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/retrieveContinuousBackupInformation"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, location); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdateTableThroughput - Update RUs per second of an Azure Cosmos DB Table
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - Cosmos DB database account name.
//   - tableName - Cosmos DB table name.
//   - updateThroughputParameters - The parameters to provide for the RUs per second of the current Table.
//   - options - TableResourcesClientBeginUpdateTableThroughputOptions contains the optional parameters for the TableResourcesClient.BeginUpdateTableThroughput
//     method.
func (client *TableResourcesClient) BeginUpdateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesClientBeginUpdateTableThroughputOptions) (*runtime.Poller[TableResourcesClientUpdateTableThroughputResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTableThroughput(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[TableResourcesClientUpdateTableThroughputResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TableResourcesClientUpdateTableThroughputResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTableThroughput - Update RUs per second of an Azure Cosmos DB Table
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-15-preview
func (client *TableResourcesClient) updateTableThroughput(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesClientBeginUpdateTableThroughputOptions) (*http.Response, error) {
	var err error
	req, err := client.updateTableThroughputCreateRequest(ctx, resourceGroupName, accountName, tableName, updateThroughputParameters, options)
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

// updateTableThroughputCreateRequest creates the UpdateTableThroughput request.
func (client *TableResourcesClient) updateTableThroughputCreateRequest(ctx context.Context, resourceGroupName string, accountName string, tableName string, updateThroughputParameters ThroughputSettingsUpdateParameters, options *TableResourcesClientBeginUpdateTableThroughputOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/tables/{tableName}/throughputSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if tableName == "" {
		return nil, errors.New("parameter tableName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tableName}", url.PathEscape(tableName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, updateThroughputParameters); err != nil {
		return nil, err
	}
	return req, nil
}
