//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetapp

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

// BackupsClient contains the methods for the Backups group.
// Don't use this type directly, use NewBackupsClient() instead.
type BackupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupsClient creates a new instance of BackupsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetVolumeRestoreStatus - Get the status of the restore for a volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - accountName - The name of the NetApp account
//   - poolName - The name of the capacity pool
//   - volumeName - The name of the volume
//   - options - BackupsClientGetVolumeRestoreStatusOptions contains the optional parameters for the BackupsClient.GetVolumeRestoreStatus
//     method.
func (client *BackupsClient) GetVolumeRestoreStatus(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetVolumeRestoreStatusOptions) (BackupsClientGetVolumeRestoreStatusResponse, error) {
	var err error
	const operationName = "BackupsClient.GetVolumeRestoreStatus"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getVolumeRestoreStatusCreateRequest(ctx, resourceGroupName, accountName, poolName, volumeName, options)
	if err != nil {
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	resp, err := client.getVolumeRestoreStatusHandleResponse(httpResp)
	return resp, err
}

// getVolumeRestoreStatusCreateRequest creates the GetVolumeRestoreStatus request.
func (client *BackupsClient) getVolumeRestoreStatusCreateRequest(ctx context.Context, resourceGroupName string, accountName string, poolName string, volumeName string, options *BackupsClientGetVolumeRestoreStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetApp/netAppAccounts/{accountName}/capacityPools/{poolName}/volumes/{volumeName}/restoreStatus"
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
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getVolumeRestoreStatusHandleResponse handles the GetVolumeRestoreStatus response.
func (client *BackupsClient) getVolumeRestoreStatusHandleResponse(resp *http.Response) (BackupsClientGetVolumeRestoreStatusResponse, error) {
	result := BackupsClientGetVolumeRestoreStatusResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestoreStatus); err != nil {
		return BackupsClientGetVolumeRestoreStatusResponse{}, err
	}
	return result, nil
}
