//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataprotection

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

// DppResourceGuardProxyClient contains the methods for the DppResourceGuardProxy group.
// Don't use this type directly, use NewDppResourceGuardProxyClient() instead.
type DppResourceGuardProxyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDppResourceGuardProxyClient creates a new instance of DppResourceGuardProxyClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDppResourceGuardProxyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DppResourceGuardProxyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DppResourceGuardProxyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates a ResourceGuardProxy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - parameters - Request body for operation
//   - options - DppResourceGuardProxyClientCreateOrUpdateOptions contains the optional parameters for the DppResourceGuardProxyClient.CreateOrUpdate
//     method.
func (client *DppResourceGuardProxyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *DppResourceGuardProxyClientCreateOrUpdateOptions) (DppResourceGuardProxyClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DppResourceGuardProxyClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DppResourceGuardProxyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters ResourceGuardProxyBaseResource, options *DppResourceGuardProxyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DppResourceGuardProxyClient) createOrUpdateHandleResponse(resp *http.Response) (DppResourceGuardProxyClientCreateOrUpdateResponse, error) {
	result := DppResourceGuardProxyClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return DppResourceGuardProxyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the ResourceGuardProxy
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - options - DppResourceGuardProxyClientDeleteOptions contains the optional parameters for the DppResourceGuardProxyClient.Delete
//     method.
func (client *DppResourceGuardProxyClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientDeleteOptions) (DppResourceGuardProxyClientDeleteResponse, error) {
	var err error
	const operationName = "DppResourceGuardProxyClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, options)
	if err != nil {
		return DppResourceGuardProxyClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DppResourceGuardProxyClientDeleteResponse{}, err
	}
	return DppResourceGuardProxyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DppResourceGuardProxyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns the ResourceGuardProxy object associated with the vault, and that matches the name in the request
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - options - DppResourceGuardProxyClientGetOptions contains the optional parameters for the DppResourceGuardProxyClient.Get
//     method.
func (client *DppResourceGuardProxyClient) Get(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientGetOptions) (DppResourceGuardProxyClientGetResponse, error) {
	var err error
	const operationName = "DppResourceGuardProxyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, options)
	if err != nil {
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DppResourceGuardProxyClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, options *DppResourceGuardProxyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DppResourceGuardProxyClient) getHandleResponse(resp *http.Response) (DppResourceGuardProxyClientGetResponse, error) {
	result := DppResourceGuardProxyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResource); err != nil {
		return DppResourceGuardProxyClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns the list of ResourceGuardProxies associated with the vault
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - DppResourceGuardProxyClientListOptions contains the optional parameters for the DppResourceGuardProxyClient.NewListPager
//     method.
func (client *DppResourceGuardProxyClient) NewListPager(resourceGroupName string, vaultName string, options *DppResourceGuardProxyClientListOptions) *runtime.Pager[DppResourceGuardProxyClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DppResourceGuardProxyClientListResponse]{
		More: func(page DppResourceGuardProxyClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DppResourceGuardProxyClientListResponse) (DppResourceGuardProxyClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DppResourceGuardProxyClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
			}, nil)
			if err != nil {
				return DppResourceGuardProxyClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DppResourceGuardProxyClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *DppResourceGuardProxyClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DppResourceGuardProxyClient) listHandleResponse(resp *http.Response) (DppResourceGuardProxyClientListResponse, error) {
	result := DppResourceGuardProxyClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceGuardProxyBaseResourceList); err != nil {
		return DppResourceGuardProxyClientListResponse{}, err
	}
	return result, nil
}

// UnlockDelete - UnlockDelete call for ResourceGuardProxy, executed before one can delete it
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - resourceGuardProxyName - name of the resource guard proxy
//   - parameters - Request body for operation
//   - options - DppResourceGuardProxyClientUnlockDeleteOptions contains the optional parameters for the DppResourceGuardProxyClient.UnlockDelete
//     method.
func (client *DppResourceGuardProxyClient) UnlockDelete(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *DppResourceGuardProxyClientUnlockDeleteOptions) (DppResourceGuardProxyClientUnlockDeleteResponse, error) {
	var err error
	const operationName = "DppResourceGuardProxyClient.UnlockDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.unlockDeleteCreateRequest(ctx, resourceGroupName, vaultName, resourceGuardProxyName, parameters, options)
	if err != nil {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	resp, err := client.unlockDeleteHandleResponse(httpResp)
	return resp, err
}

// unlockDeleteCreateRequest creates the UnlockDelete request.
func (client *DppResourceGuardProxyClient) unlockDeleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, resourceGuardProxyName string, parameters UnlockDeleteRequest, options *DppResourceGuardProxyClientUnlockDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupResourceGuardProxies/{resourceGuardProxyName}/unlockDelete"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if resourceGuardProxyName == "" {
		return nil, errors.New("parameter resourceGuardProxyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGuardProxyName}", url.PathEscape(resourceGuardProxyName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
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

// unlockDeleteHandleResponse handles the UnlockDelete response.
func (client *DppResourceGuardProxyClient) unlockDeleteHandleResponse(resp *http.Response) (DppResourceGuardProxyClientUnlockDeleteResponse, error) {
	result := DppResourceGuardProxyClientUnlockDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UnlockDeleteResponse); err != nil {
		return DppResourceGuardProxyClientUnlockDeleteResponse{}, err
	}
	return result, nil
}
