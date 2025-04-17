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

// DatabaseAdvancedThreatProtectionSettingsClient contains the methods for the DatabaseAdvancedThreatProtectionSettings group.
// Don't use this type directly, use NewDatabaseAdvancedThreatProtectionSettingsClient() instead.
type DatabaseAdvancedThreatProtectionSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseAdvancedThreatProtectionSettingsClient creates a new instance of DatabaseAdvancedThreatProtectionSettingsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseAdvancedThreatProtectionSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseAdvancedThreatProtectionSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseAdvancedThreatProtectionSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a database's Advanced Threat Protection state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - advancedThreatProtectionName - The name of the Advanced Threat Protection state.
//   - parameters - The database Advanced Threat Protection state.
//   - options - DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateOptions contains the optional parameters for the
//     DatabaseAdvancedThreatProtectionSettingsClient.CreateOrUpdate method.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters DatabaseAdvancedThreatProtection, options *DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateOptions) (DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DatabaseAdvancedThreatProtectionSettingsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, advancedThreatProtectionName, parameters, options)
	if err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advancedThreatProtectionName AdvancedThreatProtectionName, parameters DatabaseAdvancedThreatProtection, _ *DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}"
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
	if advancedThreatProtectionName == "" {
		return nil, errors.New("parameter advancedThreatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advancedThreatProtectionName}", url.PathEscape(string(advancedThreatProtectionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse, error) {
	result := DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseAdvancedThreatProtection); err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Gets a database's Advanced Threat Protection state.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - advancedThreatProtectionName - The name of the Advanced Threat Protection state.
//   - options - DatabaseAdvancedThreatProtectionSettingsClientGetOptions contains the optional parameters for the DatabaseAdvancedThreatProtectionSettingsClient.Get
//     method.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advancedThreatProtectionName AdvancedThreatProtectionName, options *DatabaseAdvancedThreatProtectionSettingsClientGetOptions) (DatabaseAdvancedThreatProtectionSettingsClientGetResponse, error) {
	var err error
	const operationName = "DatabaseAdvancedThreatProtectionSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, advancedThreatProtectionName, options)
	if err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, advancedThreatProtectionName AdvancedThreatProtectionName, _ *DatabaseAdvancedThreatProtectionSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings/{advancedThreatProtectionName}"
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
	if advancedThreatProtectionName == "" {
		return nil, errors.New("parameter advancedThreatProtectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{advancedThreatProtectionName}", url.PathEscape(string(advancedThreatProtectionName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) getHandleResponse(resp *http.Response) (DatabaseAdvancedThreatProtectionSettingsClientGetResponse, error) {
	result := DatabaseAdvancedThreatProtectionSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseAdvancedThreatProtection); err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets a list of database's Advanced Threat Protection states.
//
// Generated from API version 2021-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseOptions contains the optional parameters for the
//     DatabaseAdvancedThreatProtectionSettingsClient.NewListByDatabasePager method.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseOptions) *runtime.Pager[DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse]{
		More: func(page DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse) (DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DatabaseAdvancedThreatProtectionSettingsClient.NewListByDatabasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			}, nil)
			if err != nil {
				return DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse{}, err
			}
			return client.listByDatabaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, _ *DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/advancedThreatProtectionSettings"
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
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *DatabaseAdvancedThreatProtectionSettingsClient) listByDatabaseHandleResponse(resp *http.Response) (DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse, error) {
	result := DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseAdvancedThreatProtectionListResult); err != nil {
		return DatabaseAdvancedThreatProtectionSettingsClientListByDatabaseResponse{}, err
	}
	return result, nil
}
