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
	"strconv"
	"strings"
)

// SensitivityLabelsClient contains the methods for the SensitivityLabels group.
// Don't use this type directly, use NewSensitivityLabelsClient() instead.
type SensitivityLabelsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSensitivityLabelsClient creates a new instance of SensitivityLabelsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSensitivityLabelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SensitivityLabelsClient, error) {
	cl, err := arm.NewClient(moduleName+".SensitivityLabelsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SensitivityLabelsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the sensitivity label of a given column
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - columnName - The name of the column.
//   - parameters - The column sensitivity label resource.
//   - options - SensitivityLabelsClientCreateOrUpdateOptions contains the optional parameters for the SensitivityLabelsClient.CreateOrUpdate
//     method.
func (client *SensitivityLabelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel, options *SensitivityLabelsClientCreateOrUpdateOptions) (SensitivityLabelsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, parameters, options)
	if err != nil {
		return SensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SensitivityLabelsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, parameters SensitivityLabel, options *SensitivityLabelsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
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
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("current"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SensitivityLabelsClient) createOrUpdateHandleResponse(resp *http.Response) (SensitivityLabelsClientCreateOrUpdateResponse, error) {
	result := SensitivityLabelsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabel); err != nil {
		return SensitivityLabelsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the sensitivity label of a given column
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - columnName - The name of the column.
//   - options - SensitivityLabelsClientDeleteOptions contains the optional parameters for the SensitivityLabelsClient.Delete
//     method.
func (client *SensitivityLabelsClient) Delete(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *SensitivityLabelsClientDeleteOptions) (SensitivityLabelsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return SensitivityLabelsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensitivityLabelsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SensitivityLabelsClientDeleteResponse{}, err
	}
	return SensitivityLabelsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SensitivityLabelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *SensitivityLabelsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
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
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("current"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// DisableRecommendation - Disables sensitivity recommendations on a given column
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - columnName - The name of the column.
//   - options - SensitivityLabelsClientDisableRecommendationOptions contains the optional parameters for the SensitivityLabelsClient.DisableRecommendation
//     method.
func (client *SensitivityLabelsClient) DisableRecommendation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *SensitivityLabelsClientDisableRecommendationOptions) (SensitivityLabelsClientDisableRecommendationResponse, error) {
	var err error
	req, err := client.disableRecommendationCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return SensitivityLabelsClientDisableRecommendationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensitivityLabelsClientDisableRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SensitivityLabelsClientDisableRecommendationResponse{}, err
	}
	return SensitivityLabelsClientDisableRecommendationResponse{}, nil
}

// disableRecommendationCreateRequest creates the DisableRecommendation request.
func (client *SensitivityLabelsClient) disableRecommendationCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *SensitivityLabelsClientDisableRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/disable"
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
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("recommended"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// EnableRecommendation - Enables sensitivity recommendations on a given column (recommendations are enabled by default on
// all columns)
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - columnName - The name of the column.
//   - options - SensitivityLabelsClientEnableRecommendationOptions contains the optional parameters for the SensitivityLabelsClient.EnableRecommendation
//     method.
func (client *SensitivityLabelsClient) EnableRecommendation(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *SensitivityLabelsClientEnableRecommendationOptions) (SensitivityLabelsClientEnableRecommendationResponse, error) {
	var err error
	req, err := client.enableRecommendationCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, options)
	if err != nil {
		return SensitivityLabelsClientEnableRecommendationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensitivityLabelsClientEnableRecommendationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SensitivityLabelsClientEnableRecommendationResponse{}, err
	}
	return SensitivityLabelsClientEnableRecommendationResponse{}, nil
}

// enableRecommendationCreateRequest creates the EnableRecommendation request.
func (client *SensitivityLabelsClient) enableRecommendationCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, options *SensitivityLabelsClientEnableRecommendationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}/enable"
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
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape("recommended"))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the sensitivity label of a given column
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - schemaName - The name of the schema.
//   - tableName - The name of the table.
//   - columnName - The name of the column.
//   - sensitivityLabelSource - The source of the sensitivity label.
//   - options - SensitivityLabelsClientGetOptions contains the optional parameters for the SensitivityLabelsClient.Get method.
func (client *SensitivityLabelsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource, options *SensitivityLabelsClientGetOptions) (SensitivityLabelsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, schemaName, tableName, columnName, sensitivityLabelSource, options)
	if err != nil {
		return SensitivityLabelsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensitivityLabelsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SensitivityLabelsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SensitivityLabelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, schemaName string, tableName string, columnName string, sensitivityLabelSource SensitivityLabelSource, options *SensitivityLabelsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/schemas/{schemaName}/tables/{tableName}/columns/{columnName}/sensitivityLabels/{sensitivityLabelSource}"
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
	if columnName == "" {
		return nil, errors.New("parameter columnName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{columnName}", url.PathEscape(columnName))
	if sensitivityLabelSource == "" {
		return nil, errors.New("parameter sensitivityLabelSource cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sensitivityLabelSource}", url.PathEscape(string(sensitivityLabelSource)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SensitivityLabelsClient) getHandleResponse(resp *http.Response) (SensitivityLabelsClientGetResponse, error) {
	result := SensitivityLabelsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabel); err != nil {
		return SensitivityLabelsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDatabasePager - Gets the sensitivity labels of a given database
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - SensitivityLabelsClientListByDatabaseOptions contains the optional parameters for the SensitivityLabelsClient.NewListByDatabasePager
//     method.
func (client *SensitivityLabelsClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *SensitivityLabelsClientListByDatabaseOptions) *runtime.Pager[SensitivityLabelsClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[SensitivityLabelsClientListByDatabaseResponse]{
		More: func(page SensitivityLabelsClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SensitivityLabelsClientListByDatabaseResponse) (SensitivityLabelsClientListByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SensitivityLabelsClientListByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SensitivityLabelsClientListByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SensitivityLabelsClientListByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDatabaseHandleResponse(resp)
		},
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *SensitivityLabelsClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *SensitivityLabelsClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/sensitivityLabels"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *SensitivityLabelsClient) listByDatabaseHandleResponse(resp *http.Response) (SensitivityLabelsClientListByDatabaseResponse, error) {
	result := SensitivityLabelsClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return SensitivityLabelsClientListByDatabaseResponse{}, err
	}
	return result, nil
}

// NewListCurrentByDatabasePager - Gets the sensitivity labels of a given database
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - SensitivityLabelsClientListCurrentByDatabaseOptions contains the optional parameters for the SensitivityLabelsClient.NewListCurrentByDatabasePager
//     method.
func (client *SensitivityLabelsClient) NewListCurrentByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *SensitivityLabelsClientListCurrentByDatabaseOptions) *runtime.Pager[SensitivityLabelsClientListCurrentByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[SensitivityLabelsClientListCurrentByDatabaseResponse]{
		More: func(page SensitivityLabelsClientListCurrentByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SensitivityLabelsClientListCurrentByDatabaseResponse) (SensitivityLabelsClientListCurrentByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCurrentByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SensitivityLabelsClientListCurrentByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SensitivityLabelsClientListCurrentByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SensitivityLabelsClientListCurrentByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listCurrentByDatabaseHandleResponse(resp)
		},
	})
}

// listCurrentByDatabaseCreateRequest creates the ListCurrentByDatabase request.
func (client *SensitivityLabelsClient) listCurrentByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *SensitivityLabelsClientListCurrentByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/currentSensitivityLabels"
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
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCurrentByDatabaseHandleResponse handles the ListCurrentByDatabase response.
func (client *SensitivityLabelsClient) listCurrentByDatabaseHandleResponse(resp *http.Response) (SensitivityLabelsClientListCurrentByDatabaseResponse, error) {
	result := SensitivityLabelsClientListCurrentByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return SensitivityLabelsClientListCurrentByDatabaseResponse{}, err
	}
	return result, nil
}

// NewListRecommendedByDatabasePager - Gets the sensitivity labels of a given database
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - SensitivityLabelsClientListRecommendedByDatabaseOptions contains the optional parameters for the SensitivityLabelsClient.NewListRecommendedByDatabasePager
//     method.
func (client *SensitivityLabelsClient) NewListRecommendedByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *SensitivityLabelsClientListRecommendedByDatabaseOptions) *runtime.Pager[SensitivityLabelsClientListRecommendedByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[SensitivityLabelsClientListRecommendedByDatabaseResponse]{
		More: func(page SensitivityLabelsClientListRecommendedByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SensitivityLabelsClientListRecommendedByDatabaseResponse) (SensitivityLabelsClientListRecommendedByDatabaseResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listRecommendedByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SensitivityLabelsClientListRecommendedByDatabaseResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SensitivityLabelsClientListRecommendedByDatabaseResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SensitivityLabelsClientListRecommendedByDatabaseResponse{}, runtime.NewResponseError(resp)
			}
			return client.listRecommendedByDatabaseHandleResponse(resp)
		},
	})
}

// listRecommendedByDatabaseCreateRequest creates the ListRecommendedByDatabase request.
func (client *SensitivityLabelsClient) listRecommendedByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *SensitivityLabelsClientListRecommendedByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/recommendedSensitivityLabels"
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
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.IncludeDisabledRecommendations != nil {
		reqQP.Set("includeDisabledRecommendations", strconv.FormatBool(*options.IncludeDisabledRecommendations))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listRecommendedByDatabaseHandleResponse handles the ListRecommendedByDatabase response.
func (client *SensitivityLabelsClient) listRecommendedByDatabaseHandleResponse(resp *http.Response) (SensitivityLabelsClientListRecommendedByDatabaseResponse, error) {
	result := SensitivityLabelsClientListRecommendedByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SensitivityLabelListResult); err != nil {
		return SensitivityLabelsClientListRecommendedByDatabaseResponse{}, err
	}
	return result, nil
}

// Update - Update sensitivity labels of a given database using an operations batch.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-11-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - SensitivityLabelsClientUpdateOptions contains the optional parameters for the SensitivityLabelsClient.Update
//     method.
func (client *SensitivityLabelsClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters SensitivityLabelUpdateList, options *SensitivityLabelsClientUpdateOptions) (SensitivityLabelsClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, parameters, options)
	if err != nil {
		return SensitivityLabelsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SensitivityLabelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SensitivityLabelsClientUpdateResponse{}, err
	}
	return SensitivityLabelsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *SensitivityLabelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters SensitivityLabelUpdateList, options *SensitivityLabelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/currentSensitivityLabels"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
