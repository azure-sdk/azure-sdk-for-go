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

// BackupWorkloadItemsClient contains the methods for the BackupWorkloadItems group.
// Don't use this type directly, use NewBackupWorkloadItemsClient() instead.
type BackupWorkloadItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBackupWorkloadItemsClient creates a new instance of BackupWorkloadItemsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBackupWorkloadItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BackupWorkloadItemsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BackupWorkloadItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Provides a pageable list of workload item of a specific container according to the query filter and the
// pagination parameters.
//
// Generated from API version 2024-10-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the container.
//   - containerName - Name of the container.
//   - options - BackupWorkloadItemsClientListOptions contains the optional parameters for the BackupWorkloadItemsClient.NewListPager
//     method.
func (client *BackupWorkloadItemsClient) NewListPager(vaultName string, resourceGroupName string, fabricName string, containerName string, options *BackupWorkloadItemsClientListOptions) *runtime.Pager[BackupWorkloadItemsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[BackupWorkloadItemsClientListResponse]{
		More: func(page BackupWorkloadItemsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BackupWorkloadItemsClientListResponse) (BackupWorkloadItemsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "BackupWorkloadItemsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, options)
			}, nil)
			if err != nil {
				return BackupWorkloadItemsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *BackupWorkloadItemsClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, options *BackupWorkloadItemsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/items"
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
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if containerName == "" {
		return nil, errors.New("parameter containerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{containerName}", url.PathEscape(containerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *BackupWorkloadItemsClient) listHandleResponse(resp *http.Response) (BackupWorkloadItemsClientListResponse, error) {
	result := BackupWorkloadItemsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadItemResourceList); err != nil {
		return BackupWorkloadItemsClientListResponse{}, err
	}
	return result, nil
}
