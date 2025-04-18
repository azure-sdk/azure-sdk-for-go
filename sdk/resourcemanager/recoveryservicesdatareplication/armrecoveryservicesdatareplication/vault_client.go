// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicesdatareplication

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

// VaultClient contains the methods for the Vault group.
// Don't use this type directly, use NewVaultClient() instead.
type VaultClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVaultClient creates a new instance of VaultClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVaultClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VaultClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VaultClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - body - Vault properties.
//   - options - VaultClientBeginCreateOptions contains the optional parameters for the VaultClient.BeginCreate method.
func (client *VaultClient) BeginCreate(ctx context.Context, resourceGroupName string, vaultName string, body VaultModel, options *VaultClientBeginCreateOptions) (*runtime.Poller[VaultClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, vaultName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VaultClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VaultClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Creates the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *VaultClient) create(ctx context.Context, resourceGroupName string, vaultName string, body VaultModel, options *VaultClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "VaultClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, vaultName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *VaultClient) createCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, body VaultModel, _ *VaultClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Removes the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - options - VaultClientBeginDeleteOptions contains the optional parameters for the VaultClient.BeginDelete method.
func (client *VaultClient) BeginDelete(ctx context.Context, resourceGroupName string, vaultName string, options *VaultClientBeginDeleteOptions) (*runtime.Poller[VaultClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, vaultName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VaultClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VaultClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Removes the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *VaultClient) deleteOperation(ctx context.Context, resourceGroupName string, vaultName string, options *VaultClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VaultClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VaultClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, _ *VaultClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the details of the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - options - VaultClientGetOptions contains the optional parameters for the VaultClient.Get method.
func (client *VaultClient) Get(ctx context.Context, resourceGroupName string, vaultName string, options *VaultClientGetOptions) (VaultClientGetResponse, error) {
	var err error
	const operationName = "VaultClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VaultClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VaultClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, _ *VaultClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}"
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
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VaultClient) getHandleResponse(resp *http.Response) (VaultClientGetResponse, error) {
	result := VaultClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultModel); err != nil {
		return VaultClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of vaults in the given subscription and resource group.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - VaultClientListOptions contains the optional parameters for the VaultClient.NewListPager method.
func (client *VaultClient) NewListPager(resourceGroupName string, options *VaultClientListOptions) *runtime.Pager[VaultClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultClientListResponse]{
		More: func(page VaultClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultClientListResponse) (VaultClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VaultClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return VaultClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *VaultClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *VaultClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VaultClient) listHandleResponse(resp *http.Response) (VaultClientListResponse, error) {
	result := VaultClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultModelListResult); err != nil {
		return VaultClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets the list of vaults in the given subscription.
//
// Generated from API version 2024-09-01
//   - options - VaultClientListBySubscriptionOptions contains the optional parameters for the VaultClient.NewListBySubscriptionPager
//     method.
func (client *VaultClient) NewListBySubscriptionPager(options *VaultClientListBySubscriptionOptions) *runtime.Pager[VaultClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultClientListBySubscriptionResponse]{
		More: func(page VaultClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultClientListBySubscriptionResponse) (VaultClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VaultClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return VaultClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VaultClient) listBySubscriptionCreateRequest(ctx context.Context, _ *VaultClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DataReplication/replicationVaults"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VaultClient) listBySubscriptionHandleResponse(resp *http.Response) (VaultClientListBySubscriptionResponse, error) {
	result := VaultClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultModelListResult); err != nil {
		return VaultClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Performs update on the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The vault name.
//   - body - Vault properties.
//   - options - VaultClientBeginUpdateOptions contains the optional parameters for the VaultClient.BeginUpdate method.
func (client *VaultClient) BeginUpdate(ctx context.Context, resourceGroupName string, vaultName string, body VaultModelUpdate, options *VaultClientBeginUpdateOptions) (*runtime.Poller[VaultClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vaultName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VaultClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VaultClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Performs update on the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-09-01
func (client *VaultClient) update(ctx context.Context, resourceGroupName string, vaultName string, body VaultModelUpdate, options *VaultClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VaultClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, body, options)
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

// updateCreateRequest creates the Update request.
func (client *VaultClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, body VaultModelUpdate, _ *VaultClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
