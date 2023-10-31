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

// ManagedDatabaseRestoreDetailsClient contains the methods for the ManagedDatabaseRestoreDetails group.
// Don't use this type directly, use NewManagedDatabaseRestoreDetailsClient() instead.
type ManagedDatabaseRestoreDetailsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagedDatabaseRestoreDetailsClient creates a new instance of ManagedDatabaseRestoreDetailsClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagedDatabaseRestoreDetailsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagedDatabaseRestoreDetailsClient, error) {
	cl, err := arm.NewClient(moduleName+".ManagedDatabaseRestoreDetailsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagedDatabaseRestoreDetailsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets managed database restore details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - managedInstanceName - The name of the managed instance.
//   - databaseName - The name of the database.
//   - restoreDetailsName - The name of the restore details to retrieve.
//   - options - ManagedDatabaseRestoreDetailsClientGetOptions contains the optional parameters for the ManagedDatabaseRestoreDetailsClient.Get
//     method.
func (client *ManagedDatabaseRestoreDetailsClient) Get(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, restoreDetailsName RestoreDetailsName, options *ManagedDatabaseRestoreDetailsClientGetOptions) (ManagedDatabaseRestoreDetailsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, managedInstanceName, databaseName, restoreDetailsName, options)
	if err != nil {
		return ManagedDatabaseRestoreDetailsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagedDatabaseRestoreDetailsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagedDatabaseRestoreDetailsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ManagedDatabaseRestoreDetailsClient) getCreateRequest(ctx context.Context, resourceGroupName string, managedInstanceName string, databaseName string, restoreDetailsName RestoreDetailsName, options *ManagedDatabaseRestoreDetailsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/managedInstances/{managedInstanceName}/databases/{databaseName}/restoreDetails/{restoreDetailsName}"
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
	if restoreDetailsName == "" {
		return nil, errors.New("parameter restoreDetailsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restoreDetailsName}", url.PathEscape(string(restoreDetailsName)))
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
func (client *ManagedDatabaseRestoreDetailsClient) getHandleResponse(resp *http.Response) (ManagedDatabaseRestoreDetailsClientGetResponse, error) {
	result := ManagedDatabaseRestoreDetailsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ManagedDatabaseRestoreDetailsResult); err != nil {
		return ManagedDatabaseRestoreDetailsClientGetResponse{}, err
	}
	return result, nil
}
