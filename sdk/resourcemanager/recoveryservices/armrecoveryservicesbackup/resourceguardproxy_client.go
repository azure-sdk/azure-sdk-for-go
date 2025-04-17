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

// ResourceGuardProxyClient contains the methods for the ResourceGuardProxy group.
// Don't use this type directly, use NewResourceGuardProxyClient() instead.
type ResourceGuardProxyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceGuardProxyClient creates a new instance of ResourceGuardProxyClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceGuardProxyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceGuardProxyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceGuardProxyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Delete ResourceGuardProxy under vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ResourceGuardProxyClientDeleteOptions contains the optional parameters for the ResourceGuardProxyClient.Delete
//     method.
func (client *ResourceGuardProxyClient) Delete(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *ResourceGuardProxyClientDeleteOptions) (ResourceGuardProxyClientDeleteResponse, error) {
	var err error
	const operationName = "ResourceGuardProxyClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, options)
	if err != nil {
		return ResourceGuardProxyClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGuardProxyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGuardProxyClientDeleteResponse{}, err
	}
	return ResourceGuardProxyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ResourceGuardProxyClient) deleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, _ *ResourceGuardProxyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
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
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns ResourceGuardProxy under vault and with the name referenced in request
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ResourceGuardProxyClientGetOptions contains the optional parameters for the ResourceGuardProxyClient.Get method.
func (client *ResourceGuardProxyClient) Get(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, options *ResourceGuardProxyClientGetOptions) (ResourceGuardProxyClientGetResponse, error) {
	var err error
	const operationName = "ResourceGuardProxyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, options)
	if err != nil {
		return ResourceGuardProxyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGuardProxyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGuardProxyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ResourceGuardProxyClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, _ *ResourceGuardProxyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
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
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
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
func (client *ResourceGuardProxyClient) getHandleResponse(resp *http.Response) (ResourceGuardProxyClientGetResponse, error) {
	result := ResourceGuardProxyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return ResourceGuardProxyClientGetResponse{}, err
	}
	return result, nil
}

// Put - Add or Update ResourceGuardProxy under vault Secures vault critical operations
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - Request body for operation
//   - options - ResourceGuardProxyClientPutOptions contains the optional parameters for the ResourceGuardProxyClient.Put method.
func (client *ResourceGuardProxyClient) Put(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *ResourceGuardProxyClientPutOptions) (ResourceGuardProxyClientPutResponse, error) {
	var err error
	const operationName = "ResourceGuardProxyClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return ResourceGuardProxyClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGuardProxyClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGuardProxyClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *ResourceGuardProxyClient) putCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, _ *ResourceGuardProxyClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
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
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *ResourceGuardProxyClient) putHandleResponse(resp *http.Response) (ResourceGuardProxyClientPutResponse, error) {
	result := ResourceGuardProxyClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return ResourceGuardProxyClientPutResponse{}, err
	}
	return result, nil
}

// UnlockDelete - Secures delete ResourceGuardProxy operations.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - parameters - Request body for operation
//   - options - ResourceGuardProxyClientUnlockDeleteOptions contains the optional parameters for the ResourceGuardProxyClient.UnlockDelete
//     method.
func (client *ResourceGuardProxyClient) UnlockDelete(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *ResourceGuardProxyClientUnlockDeleteOptions) (ResourceGuardProxyClientUnlockDeleteResponse, error) {
	var err error
	const operationName = "ResourceGuardProxyClient.UnlockDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unlockDeleteCreateRequest(ctx, vaultName, resourceGroupName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	resp, err := client.unlockDeleteHandleResponse(httpResp)
	return resp, err
}

// unlockDeleteCreateRequest creates the UnlockDelete request.
func (client *ResourceGuardProxyClient) unlockDeleteCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, _ *ResourceGuardProxyClientUnlockDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}/unlockDelete"
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
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// unlockDeleteHandleResponse handles the UnlockDelete response.
func (client *ResourceGuardProxyClient) unlockDeleteHandleResponse(resp *http.Response) (ResourceGuardProxyClientUnlockDeleteResponse, error) {
	result := ResourceGuardProxyClientUnlockDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnlockDeleteResponse); err != nil {
		return ResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	return result, nil
}
