// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoveryPointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoveryPointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecoveryPointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the details of specified recovery point.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - The fabric name.
//   - protectionContainerName - The protection container name.
//   - replicatedProtectedItemName - The replication protected item name.
//   - recoveryPointName - The recovery point name.
//   - options - RecoveryPointsClientGetOptions contains the optional parameters for the RecoveryPointsClient.Get method.
func (client *RecoveryPointsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string, options *RecoveryPointsClientGetOptions) (RecoveryPointsClientGetResponse, error) {
	var err error
	const operationName = "RecoveryPointsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName, recoveryPointName, options)
	if err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecoveryPointsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecoveryPointsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, recoveryPointName string, _ *RecoveryPointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints/{recoveryPointName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if replicatedProtectedItemName == "" {
		return nil, errors.New("parameter replicatedProtectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicatedProtectedItemName}", url.PathEscape(replicatedProtectedItemName))
	if recoveryPointName == "" {
		return nil, errors.New("parameter recoveryPointName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointName}", url.PathEscape(recoveryPointName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoveryPointsClient) getHandleResponse(resp *http.Response) (RecoveryPointsClientGetResponse, error) {
	result := RecoveryPointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPoint); err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByReplicationProtectedItemsPager - Lists the available recovery points for a replication protected item.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - The fabric name.
//   - protectionContainerName - The protection container name.
//   - replicatedProtectedItemName - The replication protected item name.
//   - options - RecoveryPointsClientListByReplicationProtectedItemsOptions contains the optional parameters for the RecoveryPointsClient.NewListByReplicationProtectedItemsPager
//     method.
func (client *RecoveryPointsClient) NewListByReplicationProtectedItemsPager(resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, options *RecoveryPointsClientListByReplicationProtectedItemsOptions) *runtime.Pager[RecoveryPointsClientListByReplicationProtectedItemsResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecoveryPointsClientListByReplicationProtectedItemsResponse]{
		More: func(page RecoveryPointsClientListByReplicationProtectedItemsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecoveryPointsClientListByReplicationProtectedItemsResponse) (RecoveryPointsClientListByReplicationProtectedItemsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecoveryPointsClient.NewListByReplicationProtectedItemsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByReplicationProtectedItemsCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, replicatedProtectedItemName, options)
			}, nil)
			if err != nil {
				return RecoveryPointsClientListByReplicationProtectedItemsResponse{}, err
			}
			return client.listByReplicationProtectedItemsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByReplicationProtectedItemsCreateRequest creates the ListByReplicationProtectedItems request.
func (client *RecoveryPointsClient) listByReplicationProtectedItemsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicatedProtectedItemName string, _ *RecoveryPointsClientListByReplicationProtectedItemsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectedItems/{replicatedProtectedItemName}/recoveryPoints"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if replicatedProtectedItemName == "" {
		return nil, errors.New("parameter replicatedProtectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicatedProtectedItemName}", url.PathEscape(replicatedProtectedItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationProtectedItemsHandleResponse handles the ListByReplicationProtectedItems response.
func (client *RecoveryPointsClient) listByReplicationProtectedItemsHandleResponse(resp *http.Response) (RecoveryPointsClientListByReplicationProtectedItemsResponse, error) {
	result := RecoveryPointsClientListByReplicationProtectedItemsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointCollection); err != nil {
		return RecoveryPointsClientListByReplicationProtectedItemsResponse{}, err
	}
	return result, nil
}
