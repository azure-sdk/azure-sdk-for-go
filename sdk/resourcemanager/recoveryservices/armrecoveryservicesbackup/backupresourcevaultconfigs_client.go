// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesbackup

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

// BackupResourceVaultConfigsClient contains the methods for the BackupResourceVaultConfigs group.
// Don't use this type directly, use NewBackupResourceVaultConfigsClient() instead.
type BackupResourceVaultConfigsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupResourceVaultConfigsClient creates a new instance of BackupResourceVaultConfigsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupResourceVaultConfigsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupResourceVaultConfigsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupResourceVaultConfigsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Fetches resource vault config.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - BackupResourceVaultConfigsClientGetOptions contains the optional parameters for the BackupResourceVaultConfigsClient.Get
//     method.
func (client *BackupResourceVaultConfigsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, options *BackupResourceVaultConfigsClientGetOptions) (BackupResourceVaultConfigsClientGetResponse, error) {
	var err error
	const operationName = "BackupResourceVaultConfigsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, options)
	if err != nil {
		return BackupResourceVaultConfigsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupResourceVaultConfigsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupResourceVaultConfigsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BackupResourceVaultConfigsClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, _ *BackupResourceVaultConfigsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupResourceVaultConfigsClient) getHandleResponse(resp *http.Response) (BackupResourceVaultConfigsClientGetResponse, error) {
	result := BackupResourceVaultConfigsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResourceVaultConfigResource); err != nil {
		return BackupResourceVaultConfigsClientGetResponse{}, err
	}
	return result, nil
}

// Put - Updates vault security config.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - resource config request
//   - options - BackupResourceVaultConfigsClientPutOptions contains the optional parameters for the BackupResourceVaultConfigsClient.Put
//     method.
func (client *BackupResourceVaultConfigsClient) Put(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceVaultConfigResource, options *BackupResourceVaultConfigsClientPutOptions) (BackupResourceVaultConfigsClientPutResponse, error) {
	var err error
	const operationName = "BackupResourceVaultConfigsClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
	if err != nil {
		return BackupResourceVaultConfigsClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupResourceVaultConfigsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupResourceVaultConfigsClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *BackupResourceVaultConfigsClient) putCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceVaultConfigResource, options *BackupResourceVaultConfigsClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.XMSAuthorizationAuxiliary != nil {
		req.Raw().Header["x-ms-authorization-auxiliary"] = []string{*options.XMSAuthorizationAuxiliary}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *BackupResourceVaultConfigsClient) putHandleResponse(resp *http.Response) (BackupResourceVaultConfigsClientPutResponse, error) {
	result := BackupResourceVaultConfigsClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResourceVaultConfigResource); err != nil {
		return BackupResourceVaultConfigsClientPutResponse{}, err
	}
	return result, nil
}

// Update - Updates vault security config.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - resource config request
//   - options - BackupResourceVaultConfigsClientUpdateOptions contains the optional parameters for the BackupResourceVaultConfigsClient.Update
//     method.
func (client *BackupResourceVaultConfigsClient) Update(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceVaultConfigResource, options *BackupResourceVaultConfigsClientUpdateOptions) (BackupResourceVaultConfigsClientUpdateResponse, error) {
	var err error
	const operationName = "BackupResourceVaultConfigsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, vaultName, resourceGroupName, parameters, options)
	if err != nil {
		return BackupResourceVaultConfigsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupResourceVaultConfigsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupResourceVaultConfigsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *BackupResourceVaultConfigsClient) updateCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, parameters BackupResourceVaultConfigResource, options *BackupResourceVaultConfigsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupconfig/vaultconfig"
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.XMSAuthorizationAuxiliary != nil {
		req.Raw().Header["x-ms-authorization-auxiliary"] = []string{*options.XMSAuthorizationAuxiliary}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *BackupResourceVaultConfigsClient) updateHandleResponse(resp *http.Response) (BackupResourceVaultConfigsClientUpdateResponse, error) {
	result := BackupResourceVaultConfigsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupResourceVaultConfigResource); err != nil {
		return BackupResourceVaultConfigsClientUpdateResponse{}, err
	}
	return result, nil
}
