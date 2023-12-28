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

// DeletedBackupInstancesClient contains the methods for the DeletedBackupInstances group.
// Don't use this type directly, use NewDeletedBackupInstancesClient() instead.
type DeletedBackupInstancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeletedBackupInstancesClient creates a new instance of DeletedBackupInstancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeletedBackupInstancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeletedBackupInstancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeletedBackupInstancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a deleted backup instance with name in a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - backupInstanceName - The name of the deleted backup instance
//   - options - DeletedBackupInstancesClientGetOptions contains the optional parameters for the DeletedBackupInstancesClient.Get
//     method.
func (client *DeletedBackupInstancesClient) Get(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *DeletedBackupInstancesClientGetOptions) (DeletedBackupInstancesClientGetResponse, error) {
	var err error
	const operationName = "DeletedBackupInstancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, backupInstanceName, options)
	if err != nil {
		return DeletedBackupInstancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeletedBackupInstancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeletedBackupInstancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DeletedBackupInstancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *DeletedBackupInstancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/deletedBackupInstances/{backupInstanceName}"
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
	if backupInstanceName == "" {
		return nil, errors.New("parameter backupInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupInstanceName}", url.PathEscape(backupInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DeletedBackupInstancesClient) getHandleResponse(resp *http.Response) (DeletedBackupInstancesClientGetResponse, error) {
	result := DeletedBackupInstancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedBackupInstanceResource); err != nil {
		return DeletedBackupInstancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets deleted backup instances belonging to a backup vault
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - DeletedBackupInstancesClientListOptions contains the optional parameters for the DeletedBackupInstancesClient.NewListPager
//     method.
func (client *DeletedBackupInstancesClient) NewListPager(resourceGroupName string, vaultName string, options *DeletedBackupInstancesClientListOptions) *runtime.Pager[DeletedBackupInstancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedBackupInstancesClientListResponse]{
		More: func(page DeletedBackupInstancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedBackupInstancesClientListResponse) (DeletedBackupInstancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeletedBackupInstancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
			}, nil)
			if err != nil {
				return DeletedBackupInstancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DeletedBackupInstancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *DeletedBackupInstancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/deletedBackupInstances"
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
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedBackupInstancesClient) listHandleResponse(resp *http.Response) (DeletedBackupInstancesClientListResponse, error) {
	result := DeletedBackupInstancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedBackupInstanceResourceList); err != nil {
		return DeletedBackupInstancesClientListResponse{}, err
	}
	return result, nil
}

// BeginUndelete -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - backupInstanceName - The name of the deleted backup instance
//   - options - DeletedBackupInstancesClientBeginUndeleteOptions contains the optional parameters for the DeletedBackupInstancesClient.BeginUndelete
//     method.
func (client *DeletedBackupInstancesClient) BeginUndelete(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *DeletedBackupInstancesClientBeginUndeleteOptions) (*runtime.Poller[DeletedBackupInstancesClientUndeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.undelete(ctx, resourceGroupName, vaultName, backupInstanceName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[DeletedBackupInstancesClientUndeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[DeletedBackupInstancesClientUndeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Undelete -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01
func (client *DeletedBackupInstancesClient) undelete(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *DeletedBackupInstancesClientBeginUndeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "DeletedBackupInstancesClient.BeginUndelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.undeleteCreateRequest(ctx, resourceGroupName, vaultName, backupInstanceName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// undeleteCreateRequest creates the Undelete request.
func (client *DeletedBackupInstancesClient) undeleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *DeletedBackupInstancesClientBeginUndeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/deletedBackupInstances/{backupInstanceName}/undelete"
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
	if backupInstanceName == "" {
		return nil, errors.New("parameter backupInstanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupInstanceName}", url.PathEscape(backupInstanceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
