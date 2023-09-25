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

// DatabaseAutomaticTuningClient contains the methods for the DatabaseAutomaticTuning group.
// Don't use this type directly, use NewDatabaseAutomaticTuningClient() instead.
type DatabaseAutomaticTuningClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseAutomaticTuningClient creates a new instance of DatabaseAutomaticTuningClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseAutomaticTuningClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseAutomaticTuningClient, error) {
	cl, err := arm.NewClient(moduleName+".DatabaseAutomaticTuningClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseAutomaticTuningClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a database's automatic tuning.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - DatabaseAutomaticTuningClientGetOptions contains the optional parameters for the DatabaseAutomaticTuningClient.Get
//     method.
func (client *DatabaseAutomaticTuningClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DatabaseAutomaticTuningClientGetOptions) (DatabaseAutomaticTuningClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
	if err != nil {
		return DatabaseAutomaticTuningClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseAutomaticTuningClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseAutomaticTuningClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DatabaseAutomaticTuningClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, options *DatabaseAutomaticTuningClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/automaticTuning/current"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DatabaseAutomaticTuningClient) getHandleResponse(resp *http.Response) (DatabaseAutomaticTuningClientGetResponse, error) {
	result := DatabaseAutomaticTuningClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseAutomaticTuning); err != nil {
		return DatabaseAutomaticTuningClientGetResponse{}, err
	}
	return result, nil
}

// Update - Update automatic tuning properties for target database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - parameters - The requested automatic tuning resource state.
//   - options - DatabaseAutomaticTuningClientUpdateOptions contains the optional parameters for the DatabaseAutomaticTuningClient.Update
//     method.
func (client *DatabaseAutomaticTuningClient) Update(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DatabaseAutomaticTuning, options *DatabaseAutomaticTuningClientUpdateOptions) (DatabaseAutomaticTuningClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serverName, databaseName, parameters, options)
	if err != nil {
		return DatabaseAutomaticTuningClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseAutomaticTuningClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseAutomaticTuningClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DatabaseAutomaticTuningClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, parameters DatabaseAutomaticTuning, options *DatabaseAutomaticTuningClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/automaticTuning/current"
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
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DatabaseAutomaticTuningClient) updateHandleResponse(resp *http.Response) (DatabaseAutomaticTuningClientUpdateResponse, error) {
	result := DatabaseAutomaticTuningClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseAutomaticTuning); err != nil {
		return DatabaseAutomaticTuningClientUpdateResponse{}, err
	}
	return result, nil
}
