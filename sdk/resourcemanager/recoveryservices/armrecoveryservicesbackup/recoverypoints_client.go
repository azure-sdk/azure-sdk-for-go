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

// Get - Provides the information of the backed up data identified using RecoveryPointID. This is an asynchronous operation.
// To know the status of the operation, call the GetProtectedItemOperationResult API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with backed up item.
//   - containerName - Container name associated with backed up item.
//   - protectedItemName - Backed up item name whose backup data needs to be fetched.
//   - recoveryPointID - RecoveryPointID represents the backed up data to be fetched.
//   - options - RecoveryPointsClientGetOptions contains the optional parameters for the RecoveryPointsClient.Get method.
func (client *RecoveryPointsClient) Get(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, options *RecoveryPointsClientGetOptions) (RecoveryPointsClientGetResponse, error) {
	var err error
	const operationName = "RecoveryPointsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, recoveryPointID, options)
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
func (client *RecoveryPointsClient) getCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, recoveryPointID string, options *RecoveryPointsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints/{recoveryPointId}"
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
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	if recoveryPointID == "" {
		return nil, errors.New("parameter recoveryPointID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{recoveryPointId}", url.PathEscape(recoveryPointID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoveryPointsClient) getHandleResponse(resp *http.Response) (RecoveryPointsClientGetResponse, error) {
	result := RecoveryPointsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointResource); err != nil {
		return RecoveryPointsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the backup copies for the backed up item.
//
// Generated from API version 2024-01-01
//   - vaultName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name associated with the backed up item.
//   - containerName - Container name associated with the backed up item.
//   - protectedItemName - Backed up item whose backup copies are to be fetched.
//   - options - RecoveryPointsClientListOptions contains the optional parameters for the RecoveryPointsClient.NewListPager method.
func (client *RecoveryPointsClient) NewListPager(vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, options *RecoveryPointsClientListOptions) *runtime.Pager[RecoveryPointsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RecoveryPointsClientListResponse]{
		More: func(page RecoveryPointsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RecoveryPointsClientListResponse) (RecoveryPointsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RecoveryPointsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, vaultName, resourceGroupName, fabricName, containerName, protectedItemName, options)
			}, nil)
			if err != nil {
				return RecoveryPointsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RecoveryPointsClient) listCreateRequest(ctx context.Context, vaultName string, resourceGroupName string, fabricName string, containerName string, protectedItemName string, options *RecoveryPointsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/backupFabrics/{fabricName}/protectionContainers/{containerName}/protectedItems/{protectedItemName}/recoveryPoints"
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
	if protectedItemName == "" {
		return nil, errors.New("parameter protectedItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectedItemName}", url.PathEscape(protectedItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RecoveryPointsClient) listHandleResponse(resp *http.Response) (RecoveryPointsClientListResponse, error) {
	result := RecoveryPointsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryPointResourceList); err != nil {
		return RecoveryPointsClientListResponse{}, err
	}
	return result, nil
}
