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

// BackupPoliciesClient contains the methods for the BackupPolicies group.
// Don't use this type directly, use NewBackupPoliciesClient() instead.
type BackupPoliciesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupPoliciesClient creates a new instance of BackupPoliciesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupPoliciesClient, error) {
	cl, err := arm.NewClient(moduleName+".BackupPoliciesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupPoliciesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or Updates a backup policy belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - backupPolicyName - Name of the policy
//   - parameters - Request body for operation
//   - options - BackupPoliciesClientCreateOrUpdateOptions contains the optional parameters for the BackupPoliciesClient.CreateOrUpdate
//     method.
func (client *BackupPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, parameters BaseBackupPolicyResource, options *BackupPoliciesClientCreateOrUpdateOptions) (BackupPoliciesClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, backupPolicyName, parameters, options)
	if err != nil {
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BackupPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, parameters BaseBackupPolicyResource, options *BackupPoliciesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies/{backupPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", url.PathEscape(backupPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *BackupPoliciesClient) createOrUpdateHandleResponse(resp *http.Response) (BackupPoliciesClientCreateOrUpdateResponse, error) {
	result := BackupPoliciesClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BaseBackupPolicyResource); err != nil {
		return BackupPoliciesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a backup policy belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - BackupPoliciesClientDeleteOptions contains the optional parameters for the BackupPoliciesClient.Delete method.
func (client *BackupPoliciesClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, options *BackupPoliciesClientDeleteOptions) (BackupPoliciesClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, backupPolicyName, options)
	if err != nil {
		return BackupPoliciesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupPoliciesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return BackupPoliciesClientDeleteResponse{}, err
	}
	return BackupPoliciesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *BackupPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, options *BackupPoliciesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies/{backupPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", url.PathEscape(backupPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a backup policy belonging to a backup vault
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - BackupPoliciesClientGetOptions contains the optional parameters for the BackupPoliciesClient.Get method.
func (client *BackupPoliciesClient) Get(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, options *BackupPoliciesClientGetOptions) (BackupPoliciesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, backupPolicyName, options)
	if err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BackupPoliciesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BackupPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupPolicyName string, options *BackupPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies/{backupPolicyName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if backupPolicyName == "" {
		return nil, errors.New("parameter backupPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{backupPolicyName}", url.PathEscape(backupPolicyName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BackupPoliciesClient) getHandleResponse(resp *http.Response) (BackupPoliciesClientGetResponse, error) {
	result := BackupPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BaseBackupPolicyResource); err != nil {
		return BackupPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of backup policies belonging to a backup vault
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - options - BackupPoliciesClientListOptions contains the optional parameters for the BackupPoliciesClient.NewListPager method.
func (client *BackupPoliciesClient) NewListPager(resourceGroupName string, vaultName string, options *BackupPoliciesClientListOptions) *runtime.Pager[BackupPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupPoliciesClientListResponse]{
		More: func(page BackupPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupPoliciesClientListResponse) (BackupPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, vaultName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BackupPoliciesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BackupPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BackupPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *BackupPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *BackupPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupPolicies"
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
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupPoliciesClient) listHandleResponse(resp *http.Response) (BackupPoliciesClientListResponse, error) {
	result := BackupPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BaseBackupPolicyResourceList); err != nil {
		return BackupPoliciesClientListResponse{}, err
	}
	return result, nil
}
