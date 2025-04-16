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
	"strings"
)

// DatabaseMigrationsSQLMiClient contains the methods for the DatabaseMigrationsSQLMi group.
// Don't use this type directly, use NewDatabaseMigrationsSQLMiClient() instead.
type DatabaseMigrationsSQLMiClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseMigrationsSQLMiClient creates a new instance of DatabaseMigrationsSQLMiClient with the specified values.
//   - subscriptionID - Subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseMigrationsSQLMiClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseMigrationsSQLMiClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseMigrationsSQLMiClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCancel - Stop in-progress database migration to SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - parameters - Required migration operation ID for which cancel will be initiated.
//   - options - DatabaseMigrationsSQLMiClientBeginCancelOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCancel
//     method.
func (client *DatabaseMigrationsSQLMiClient) BeginCancel(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCancelOptions) (*runtime.Poller[DatabaseMigrationsSQLMiClientCancelResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cancel(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseMigrationsSQLMiClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseMigrationsSQLMiClientCancelResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Cancel - Stop in-progress database migration to SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
func (client *DatabaseMigrationsSQLMiClient) cancel(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCancelOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseMigrationsSQLMiClient.BeginCancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
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
func (client *DatabaseMigrationsSQLMiClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, _ *DatabaseMigrationsSQLMiClientBeginCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cancel"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreateOrUpdate - Create a new database migration to a given SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - parameters - Details of SqlMigrationService resource.
//   - options - DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate
//     method.
func (client *DatabaseMigrationsSQLMiClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi, options *DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (*runtime.Poller[DatabaseMigrationsSQLMiClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseMigrationsSQLMiClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseMigrationsSQLMiClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create a new database migration to a given SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
func (client *DatabaseMigrationsSQLMiClient) createOrUpdate(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi, options *DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseMigrationsSQLMiClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
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
func (client *DatabaseMigrationsSQLMiClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters DatabaseMigrationSQLMi, _ *DatabaseMigrationsSQLMiClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCutover - Initiate cutover for in-progress online database migration to SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - parameters - Required migration operation ID for which cutover will be initiated.
//   - options - DatabaseMigrationsSQLMiClientBeginCutoverOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.BeginCutover
//     method.
func (client *DatabaseMigrationsSQLMiClient) BeginCutover(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCutoverOptions) (*runtime.Poller[DatabaseMigrationsSQLMiClientCutoverResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.cutover(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabaseMigrationsSQLMiClientCutoverResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DatabaseMigrationsSQLMiClientCutoverResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Cutover - Initiate cutover for in-progress online database migration to SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
func (client *DatabaseMigrationsSQLMiClient) cutover(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, options *DatabaseMigrationsSQLMiClientBeginCutoverOptions) (*http.Response, error) {
	var err error
	const operationName = "DatabaseMigrationsSQLMiClient.BeginCutover"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cutoverCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, parameters, options)
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

// cutoverCreateRequest creates the Cutover request.
func (client *DatabaseMigrationsSQLMiClient) cutoverCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, parameters MigrationOperationInput, _ *DatabaseMigrationsSQLMiClientBeginCutoverOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}/cutover"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	reqQP.Set("api-version", "2025-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Retrieve the specified database migration for a given SQL Managed Instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-15-preview
//   - resourceGroupName - Name of the resource group that contains the resource. You can obtain this value from the Azure Resource
//     Manager API or the portal.
//   - targetDbName - The name of the target database.
//   - options - DatabaseMigrationsSQLMiClientGetOptions contains the optional parameters for the DatabaseMigrationsSQLMiClient.Get
//     method.
func (client *DatabaseMigrationsSQLMiClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, options *DatabaseMigrationsSQLMiClientGetOptions) (DatabaseMigrationsSQLMiClientGetResponse, error) {
	var err error
	const operationName = "DatabaseMigrationsSQLMiClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, targetDbName, options)
	if err != nil {
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseMigrationsSQLMiClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, targetDbName string, options *DatabaseMigrationsSQLMiClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/providers/Microsoft.DataMigration/databaseMigrations/{targetDbName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if managedInstanceName == "" {
		return nil, errors.New("parameter managedInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{managedInstanceName}", url.PathEscape(managedInstanceName))
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
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2025-03-15-preview")
	if options != nil && options.MigrationOperationID != nil {
		reqQP.Set("migrationOperationId", *options.MigrationOperationID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseMigrationsSQLMiClient) getHandleResponse(resp *http.Response) (DatabaseMigrationsSQLMiClientGetResponse, error) {
	result := DatabaseMigrationsSQLMiClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseMigrationSQLMi); err != nil {
		return DatabaseMigrationsSQLMiClientGetResponse{}, err
	}
	return result, nil
}
