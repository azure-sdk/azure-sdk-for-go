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

// MaintenanceWindowOptionsClient contains the methods for the MaintenanceWindowOptions group.
// Don't use this type directly, use NewMaintenanceWindowOptionsClient() instead.
type MaintenanceWindowOptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMaintenanceWindowOptionsClient creates a new instance of MaintenanceWindowOptionsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMaintenanceWindowOptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MaintenanceWindowOptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MaintenanceWindowOptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a list of available maintenance windows.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database to get maintenance windows options for.
//   - maintenanceWindowOptionsName - Maintenance window options name.
//   - options - MaintenanceWindowOptionsClientGetOptions contains the optional parameters for the MaintenanceWindowOptionsClient.Get
//     method.
func (client *MaintenanceWindowOptionsClient) Get(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowOptionsName string, options *MaintenanceWindowOptionsClientGetOptions) (MaintenanceWindowOptionsClientGetResponse, error) {
	var err error
	const operationName = "MaintenanceWindowOptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, databaseName, maintenanceWindowOptionsName, options)
	if err != nil {
		return MaintenanceWindowOptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MaintenanceWindowOptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MaintenanceWindowOptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MaintenanceWindowOptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, maintenanceWindowOptionsName string, options *MaintenanceWindowOptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/maintenanceWindowOptions/current"
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
	reqQP.Set("api-version", "2024-02-01-preview")
	reqQP.Set("maintenanceWindowOptionsName", maintenanceWindowOptionsName)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MaintenanceWindowOptionsClient) getHandleResponse(resp *http.Response) (MaintenanceWindowOptionsClientGetResponse, error) {
	result := MaintenanceWindowOptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MaintenanceWindowOptions); err != nil {
		return MaintenanceWindowOptionsClientGetResponse{}, err
	}
	return result, nil
}
