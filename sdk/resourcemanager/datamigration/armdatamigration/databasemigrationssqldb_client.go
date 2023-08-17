//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatamigration

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

// DatabaseMigrationsSQLDbClient contains the methods for the DatabaseMigrationsSQLDb group.
// Don't use this type directly, use NewDatabaseMigrationsSQLDbClient() instead.
type DatabaseMigrationsSQLDbClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseMigrationsSQLDbClient creates a new instance of DatabaseMigrationsSQLDbClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseMigrationsSQLDbClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseMigrationsSQLDbClient, error) {
	cl, err := arm.NewClient(moduleName+".DatabaseMigrationsSQLDbClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseMigrationsSQLDbClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - Stop on going migration for the database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - parameters - Required migration operation ID for which cancel will be initiated.
//   - options - DatabaseMigrationsSQLDbClientBeginCancelOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.BeginCancel
//     method.
func (client *DatabaseMigrationsSQLDbClient) BeginCancel(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLDbClientBeginCancelOptions) (*runtime.Poller[DatabaseMigrationsSQLDbClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DatabaseMigrationsSQLDbClientCancelResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DatabaseMigrationsSQLDbClientCancelResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Cancel - Stop on going migration for the database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *DatabaseMigrationsSQLDbClient) cancel(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLDbClientBeginCancelOptions) (*http.Response, error) {
	var err error
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, parameters, options)
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

// cancelCreateRequest creates the Cancel request.
func (client *DatabaseMigrationsSQLDbClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLDbClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{sqlDbInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDbInstanceName == "" {
		return nil, errors.New("parameter sqlDbInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDbInstanceName}", url.PathEscape(sqlDbInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Create or Update Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - parameters - Details of Sql Db migration resource.
//   - options - DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.BeginCreateOrUpdate
//     method.
func (client *DatabaseMigrationsSQLDbClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters DatabaseMigrationSQLDb, options *DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions) (*runtime.Poller[DatabaseMigrationsSQLDbClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DatabaseMigrationsSQLDbClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DatabaseMigrationsSQLDbClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or Update Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *DatabaseMigrationsSQLDbClient) createOrUpdate(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters DatabaseMigrationSQLDb, options *DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabaseMigrationsSQLDbClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, parameters DatabaseMigrationSQLDb, options *DatabaseMigrationsSQLDbClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{sqlDbInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDbInstanceName == "" {
		return nil, errors.New("parameter sqlDbInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDbInstanceName}", url.PathEscape(sqlDbInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - options - DatabaseMigrationsSQLDbClientBeginDeleteOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.BeginDelete
//     method.
func (client *DatabaseMigrationsSQLDbClient) BeginDelete(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *DatabaseMigrationsSQLDbClientBeginDeleteOptions) (*runtime.Poller[DatabaseMigrationsSQLDbClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[DatabaseMigrationsSQLDbClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[DatabaseMigrationsSQLDbClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
func (client *DatabaseMigrationsSQLDbClient) deleteOperation(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *DatabaseMigrationsSQLDbClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, options)
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
func (client *DatabaseMigrationsSQLDbClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *DatabaseMigrationsSQLDbClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{sqlDbInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDbInstanceName == "" {
		return nil, errors.New("parameter sqlDbInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDbInstanceName}", url.PathEscape(sqlDbInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Retrieve the Database Migration resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - options - DatabaseMigrationsSQLDbClientGetOptions contains the optional parameters for the DatabaseMigrationsSQLDbClient.Get
//     method.
func (client *DatabaseMigrationsSQLDbClient) Get(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *DatabaseMigrationsSQLDbClientGetOptions) (DatabaseMigrationsSQLDbClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, sqlDbInstanceName, targetDbName, options)
	if err != nil {
		return DatabaseMigrationsSQLDbClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseMigrationsSQLDbClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseMigrationsSQLDbClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseMigrationsSQLDbClient) getCreateRequest(ctx context.Context, resourceGroupName string, sqlDbInstanceName string, targetDbName string, options *DatabaseMigrationsSQLDbClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{sqlDbInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if sqlDbInstanceName == "" {
		return nil, errors.New("parameter sqlDbInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sqlDbInstanceName}", url.PathEscape(sqlDbInstanceName))
	if targetDbName == "" {
		return nil, errors.New("parameter targetDbName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{targetDbName}", url.PathEscape(targetDbName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.MigrationOperationID != nil {
		reqQP.Set("migrationOperationId", *options.MigrationOperationID)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2023-07-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseMigrationsSQLDbClient) getHandleResponse(resp *http.Response) (DatabaseMigrationsSQLDbClientGetResponse, error) {
	result := DatabaseMigrationsSQLDbClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationSQLDb); err != nil {
		return DatabaseMigrationsSQLDbClientGetResponse{}, err
	}
	return result, nil
}
