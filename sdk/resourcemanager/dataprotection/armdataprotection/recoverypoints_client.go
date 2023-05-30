//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// RecoveryPointsClient contains the methods for the RecoveryPoints group.
// Don't use this type directly, use NewRecoveryPointsClient() instead.
type RecoveryPointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRecoveryPointsClient creates a new instance of RecoveryPointsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoveryPointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoveryPointsClient, error) {
	cl, err := arm.NewClient(moduleName+".RecoveryPointsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecoveryPointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a Recovery Point using recoveryPointId for a Datasource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - backupInstanceName - The name of the backup instance.
//   - options - RecoveryPointsClientGetOptions contains the optional parameters for the RecoveryPointsClient.Get method.
func (client *RecoveryPointsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, recoveryPointID string, options *RecoveryPointsClientGetOptions) (RecoveryPointsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, backupInstanceName, recoveryPointID, options)
	if err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RecoveryPointsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RecoveryPointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, recoveryPointID string, options *RecoveryPointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/recoveryPoints/{recoveryPointId}"
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
	if recoveryPointID == "" {
		return nil, errors.New("parameter recoveryPointID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointId}", url.PathEscape(recoveryPointID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoveryPointsClient) getHandleResponse(resp *http.Response) (RecoveryPointsClientGetResponse, error) {
	result := RecoveryPointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBackupRecoveryPointResource); err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of Recovery Points for a DataSource in a vault.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the backup vault.
//   - backupInstanceName - The name of the backup instance.
//   - options - RecoveryPointsClientListOptions contains the optional parameters for the RecoveryPointsClient.NewListPager method.
func (client *RecoveryPointsClient) NewListPager(resourceGroupName string, vaultName string, backupInstanceName string, options *RecoveryPointsClientListOptions) *runtime.Pager[RecoveryPointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecoveryPointsClientListResponse]{
		More: func(page RecoveryPointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecoveryPointsClientListResponse) (RecoveryPointsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, vaultName, backupInstanceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RecoveryPointsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RecoveryPointsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RecoveryPointsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *RecoveryPointsClient) listCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, backupInstanceName string, options *RecoveryPointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataProtection/backupVaults/{vaultName}/backupInstances/{backupInstanceName}/recoveryPoints"
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
	reqQP.Set("api-version", "2023-04-01-preview")
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
func (client *RecoveryPointsClient) listHandleResponse(resp *http.Response) (RecoveryPointsClientListResponse, error) {
	result := RecoveryPointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AzureBackupRecoveryPointResourceList); err != nil {
		return RecoveryPointsClientListResponse{}, err
	}
	return result, nil
}
