//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise

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

// DatabasesClient contains the methods for the Databases group.
// Don't use this type directly, use NewDatabasesClient() instead.
type DatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabasesClient creates a new instance of DatabasesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabasesClient, error) {
	cl, err := arm.NewClient(moduleName+".DatabasesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - resource - Resource create parameters.
//   - options - DatabasesClientBeginCreateOptions contains the optional parameters for the DatabasesClient.BeginCreate method.
func (client *DatabasesClient) BeginCreate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, resource DatabaseCreateOrUpdate, options *DatabasesClientBeginCreateOptions) (*runtime.Poller[DatabasesClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, clusterName, databaseName, resource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) create(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, resource DatabaseCreateOrUpdate, options *DatabasesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, clusterName, databaseName, resource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *DatabasesClient) createCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, resource DatabaseCreateOrUpdate, options *DatabasesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, resource)
}

// BeginDelete - Deletes a database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - options - DatabasesClientBeginDeleteOptions contains the optional parameters for the DatabasesClient.BeginDelete method.
func (client *DatabasesClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*runtime.Poller[DatabasesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, databaseName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DatabasesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// BeginExportRdb - Exports RDB file(s)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - body - The content of the action request
//   - options - DatabasesClientBeginExportRdbOptions contains the optional parameters for the DatabasesClient.BeginExportRdb
//     method.
func (client *DatabasesClient) BeginExportRdb(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ExportParameters, options *DatabasesClientBeginExportRdbOptions) (*runtime.Poller[DatabasesClientExportRdbResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.exportRdb(ctx, resourceGroupName, clusterName, databaseName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatabasesClientExportRdbResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientExportRdbResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ExportRdb - Exports RDB file(s)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) exportRdb(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ExportParameters, options *DatabasesClientBeginExportRdbOptions) (*http.Response, error) {
	req, err := client.exportRdbCreateRequest(ctx, resourceGroupName, clusterName, databaseName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// exportRdbCreateRequest creates the ExportRdb request.
func (client *DatabasesClient) exportRdbCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ExportParameters, options *DatabasesClientBeginExportRdbOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/export"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginForceUnlink - Forcibly unlinks one or more databases from a replication group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - body - The content of the action request
//   - options - DatabasesClientBeginForceUnlinkOptions contains the optional parameters for the DatabasesClient.BeginForceUnlink
//     method.
func (client *DatabasesClient) BeginForceUnlink(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ForceUnlinkParameters, options *DatabasesClientBeginForceUnlinkOptions) (*runtime.Poller[DatabasesClientForceUnlinkResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.forceUnlink(ctx, resourceGroupName, clusterName, databaseName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatabasesClientForceUnlinkResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientForceUnlinkResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ForceUnlink - Forcibly unlinks one or more databases from a replication group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) forceUnlink(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ForceUnlinkParameters, options *DatabasesClientBeginForceUnlinkOptions) (*http.Response, error) {
	req, err := client.forceUnlinkCreateRequest(ctx, resourceGroupName, clusterName, databaseName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// forceUnlinkCreateRequest creates the ForceUnlink request.
func (client *DatabasesClient) forceUnlinkCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ForceUnlinkParameters, options *DatabasesClientBeginForceUnlinkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/forceUnlink"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginFush - Forcibly flushes data from all databases in a replication group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - body - The content of the action request
//   - options - DatabasesClientBeginFushOptions contains the optional parameters for the DatabasesClient.BeginFush method.
func (client *DatabasesClient) BeginFush(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body FlushParameters, options *DatabasesClientBeginFushOptions) (*runtime.Poller[DatabasesClientFushResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.fush(ctx, resourceGroupName, clusterName, databaseName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatabasesClientFushResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientFushResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Fush - Forcibly flushes data from all databases in a replication group
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) fush(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body FlushParameters, options *DatabasesClientBeginFushOptions) (*http.Response, error) {
	req, err := client.fushCreateRequest(ctx, resourceGroupName, clusterName, databaseName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// fushCreateRequest creates the Fush request.
func (client *DatabasesClient) fushCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body FlushParameters, options *DatabasesClientBeginFushOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/fush"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// Get - Gets information about a database in a RedisEnterprise cluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - options - DatabasesClientGetOptions contains the optional parameters for the DatabasesClient.Get method.
func (client *DatabasesClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientGetOptions) (DatabasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *DatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabasesClient) getHandleResponse(resp *http.Response) (DatabasesClientGetResponse, error) {
	result := DatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Database); err != nil {
		return DatabasesClientGetResponse{}, err
	}
	return result, nil
}

// BeginImportRdb - Imports RDB file(s)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - body - The content of the action request
//   - options - DatabasesClientBeginImportRdbOptions contains the optional parameters for the DatabasesClient.BeginImportRdb
//     method.
func (client *DatabasesClient) BeginImportRdb(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ImportParameters, options *DatabasesClientBeginImportRdbOptions) (*runtime.Poller[DatabasesClientImportRdbResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.importRdb(ctx, resourceGroupName, clusterName, databaseName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[DatabasesClientImportRdbResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientImportRdbResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// ImportRdb - Imports RDB file(s)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) importRdb(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ImportParameters, options *DatabasesClientBeginImportRdbOptions) (*http.Response, error) {
	req, err := client.importRdbCreateRequest(ctx, resourceGroupName, clusterName, databaseName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// importRdbCreateRequest creates the ImportRdb request.
func (client *DatabasesClient) importRdbCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body ImportParameters, options *DatabasesClientBeginImportRdbOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/import"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// NewListByClusterPager - Lists all databases in a RedisEnterprise cluster.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - options - DatabasesClientListByClusterOptions contains the optional parameters for the DatabasesClient.NewListByClusterPager
//     method.
func (client *DatabasesClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *DatabasesClientListByClusterOptions) *runtime.Pager[DatabasesClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabasesClientListByClusterResponse]{
		More: func(page DatabasesClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabasesClientListByClusterResponse) (DatabasesClientListByClusterResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DatabasesClientListByClusterResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DatabasesClientListByClusterResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DatabasesClientListByClusterResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByClusterHandleResponse(resp)
		},
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *DatabasesClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *DatabasesClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *DatabasesClient) listByClusterHandleResponse(resp *http.Response) (DatabasesClientListByClusterResponse, error) {
	result := DatabasesClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseListResult); err != nil {
		return DatabasesClientListByClusterResponse{}, err
	}
	return result, nil
}

// ListKeys - Retrieves the access keys for the RedisEnterprise database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - options - DatabasesClientListKeysOptions contains the optional parameters for the DatabasesClient.ListKeys method.
func (client *DatabasesClient) ListKeys(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientListKeysOptions) (DatabasesClientListKeysResponse, error) {
	req, err := client.listKeysCreateRequest(ctx, resourceGroupName, clusterName, databaseName, options)
	if err != nil {
		return DatabasesClientListKeysResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabasesClientListKeysResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DatabasesClientListKeysResponse{}, runtime.NewResponseError(resp)
	}
	return client.listKeysHandleResponse(resp)
}

// listKeysCreateRequest creates the ListKeys request.
func (client *DatabasesClient) listKeysCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, options *DatabasesClientListKeysOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/listKeys"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listKeysHandleResponse handles the ListKeys response.
func (client *DatabasesClient) listKeysHandleResponse(resp *http.Response) (DatabasesClientListKeysResponse, error) {
	result := DatabasesClientListKeysResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessKeys); err != nil {
		return DatabasesClientListKeysResponse{}, err
	}
	return result, nil
}

// BeginRegenerateKey - Regenerates an access key for the RedisEnterprise database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - body - The content of the action request
//   - options - DatabasesClientBeginRegenerateKeyOptions contains the optional parameters for the DatabasesClient.BeginRegenerateKey
//     method.
func (client *DatabasesClient) BeginRegenerateKey(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body RegenerateKeyParameters, options *DatabasesClientBeginRegenerateKeyOptions) (*runtime.Poller[DatabasesClientRegenerateKeyResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.regenerateKey(ctx, resourceGroupName, clusterName, databaseName, body, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientRegenerateKeyResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientRegenerateKeyResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RegenerateKey - Regenerates an access key for the RedisEnterprise database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) regenerateKey(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body RegenerateKeyParameters, options *DatabasesClientBeginRegenerateKeyOptions) (*http.Response, error) {
	req, err := client.regenerateKeyCreateRequest(ctx, resourceGroupName, clusterName, databaseName, body, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// regenerateKeyCreateRequest creates the RegenerateKey request.
func (client *DatabasesClient) regenerateKeyCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, body RegenerateKeyParameters, options *DatabasesClientBeginRegenerateKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}/regenerateKey"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, body)
}

// BeginUpdate - Updates a database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - Name of cluster
//   - databaseName - Name of database
//   - properties - The resource properties to be updated.
//   - options - DatabasesClientBeginUpdateOptions contains the optional parameters for the DatabasesClient.BeginUpdate method.
func (client *DatabasesClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, properties any, options *DatabasesClientBeginUpdateOptions) (*runtime.Poller[DatabasesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, databaseName, properties, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DatabasesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[DatabasesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates a database
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *DatabasesClient) update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, properties any, options *DatabasesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, databaseName, properties, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *DatabasesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, properties any, options *DatabasesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redisEnterprise/{clusterName}/databases/{databaseName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, properties)
}
