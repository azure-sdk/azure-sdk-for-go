//go:build go1.18
// +build go1.18

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

// BackupEnginesClient contains the methods for the BackupEngines group.
// Don't use this type directly, use NewBackupEnginesClient() instead.
type BackupEnginesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupEnginesClient creates a new instance of BackupEnginesClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupEnginesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupEnginesClient, error) {
	cl, err := arm.NewClient(moduleName+".BackupEnginesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupEnginesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns backup management server registered to Recovery Services Vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - backupEngineName - Name of the backup management server.
//   - options - BackupEnginesClientGetOptions contains the optional parameters for the BackupEnginesClient.Get method.
func (client *BackupEnginesClient) Get(ctx context.Context, vaultName string, resourceGroupName string, backupEngineName string, options *BackupEnginesClientGetOptions) (BackupEnginesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, backupEngineName, options)
	if err != nil {
		return BackupEnginesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupEnginesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupEnginesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BackupEnginesClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, backupEngineName string, options *BackupEnginesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupEngines/{backupEngineName}"
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
	if backupEngineName == "" {
		return nil, errors.New("parameter backupEngineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupEngineName}", url.PathEscape(backupEngineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupEnginesClient) getHandleResponse(resp *http.Response) (BackupEnginesClientGetResponse, error) {
	result := BackupEnginesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupEngineBaseResource); err != nil {
		return BackupEnginesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Backup management servers registered to Recovery Services Vault. Returns a pageable list of servers.
//
// Generated from API version 2023-06-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - BackupEnginesClientListOptions contains the optional parameters for the BackupEnginesClient.NewListPager method.
func (client *BackupEnginesClient) NewListPager(vaultName string, resourceGroupName string, options *BackupEnginesClientListOptions) *runtime.Pager[BackupEnginesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupEnginesClientListResponse]{
		More: func(page BackupEnginesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupEnginesClientListResponse) (BackupEnginesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vaultName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupEnginesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BackupEnginesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupEnginesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupEnginesClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, options *BackupEnginesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupEngines"
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
	reqQP.Set("api-version", "2023-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupEnginesClient) listHandleResponse(resp *http.Response) (BackupEnginesClientListResponse, error) {
	result := BackupEnginesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BackupEngineBaseResourceList); err != nil {
		return BackupEnginesClientListResponse{}, err
	}
	return result, nil
}
