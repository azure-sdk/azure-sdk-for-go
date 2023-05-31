//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservices

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

// OperationsClient contains the methods for the Operations group.
// Don't use this type directly, use NewOperationsClient() instead.
type OperationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationsClient creates a new instance of OperationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetOperationResult - Gets the operation result for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - OperationsClientGetOperationResultOptions contains the optional parameters for the OperationsClient.GetOperationResult
//     method.
func (client *OperationsClient) GetOperationResult(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *OperationsClientGetOperationResultOptions) (OperationsClientGetOperationResultResponse, error) {
	req, err := client.getOperationResultCreateRequest(ctx, resourceGroupName, vaultName, operationID, options)
	if err != nil {
		return OperationsClientGetOperationResultResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientGetOperationResultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return OperationsClientGetOperationResultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getOperationResultHandleResponse(resp)
}

// getOperationResultCreateRequest creates the GetOperationResult request.
func (client *OperationsClient) getOperationResultCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *OperationsClientGetOperationResultOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/operationResults/{operationId}"
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
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getOperationResultHandleResponse handles the GetOperationResult response.
func (client *OperationsClient) getOperationResultHandleResponse(resp *http.Response) (OperationsClientGetOperationResultResponse, error) {
	result := OperationsClientGetOperationResultResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Vault); err != nil {
		return OperationsClientGetOperationResultResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns the list of available operations.
//
// Generated from API version 2023-04-01
//   - options - OperationsClientListOptions contains the optional parameters for the OperationsClient.NewListPager method.
func (client *OperationsClient) NewListPager(options *OperationsClientListOptions) *runtime.Pager[OperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationsClientListResponse]{
		More: func(page OperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationsClientListResponse) (OperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OperationsClient) listCreateRequest(ctx context.Context, options *OperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.RecoveryServices/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OperationsClient) listHandleResponse(resp *http.Response) (OperationsClientListResponse, error) {
	result := OperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClientDiscoveryResponse); err != nil {
		return OperationsClientListResponse{}, err
	}
	return result, nil
}

// OperationStatusGet - Gets the operation status for a resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - OperationsClientOperationStatusGetOptions contains the optional parameters for the OperationsClient.OperationStatusGet
//     method.
func (client *OperationsClient) OperationStatusGet(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *OperationsClientOperationStatusGetOptions) (OperationsClientOperationStatusGetResponse, error) {
	req, err := client.operationStatusGetCreateRequest(ctx, resourceGroupName, vaultName, operationID, options)
	if err != nil {
		return OperationsClientOperationStatusGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationsClientOperationStatusGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OperationsClientOperationStatusGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.operationStatusGetHandleResponse(resp)
}

// operationStatusGetCreateRequest creates the OperationStatusGet request.
func (client *OperationsClient) operationStatusGetCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationID string, options *OperationsClientOperationStatusGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}/operationStatus/{operationId}"
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
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// operationStatusGetHandleResponse handles the OperationStatusGet response.
func (client *OperationsClient) operationStatusGetHandleResponse(resp *http.Response) (OperationsClientOperationStatusGetResponse, error) {
	result := OperationsClientOperationStatusGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationResource); err != nil {
		return OperationsClientOperationStatusGetResponse{}, err
	}
	return result, nil
}
