//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// VaultsClient contains the methods for the Vaults group.
// Don't use this type directly, use NewVaultsClient() instead.
type VaultsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVaultsClient creates a new instance of VaultsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVaultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VaultsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VaultsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a Recovery Services vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - vault - Recovery Services Vault to be created.
//   - options - VaultsClientBeginCreateOrUpdateOptions contains the optional parameters for the VaultsClient.BeginCreateOrUpdate
//     method.
func (client *VaultsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, vault Vault, options *VaultsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VaultsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, vaultName, vault, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VaultsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VaultsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a Recovery Services vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *VaultsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, vault Vault, options *VaultsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VaultsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, vault, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VaultsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, vault Vault, options *VaultsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, vault); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Deletes a vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - VaultsClientDeleteOptions contains the optional parameters for the VaultsClient.Delete method.
func (client *VaultsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientDeleteOptions) (VaultsClientDeleteResponse, error) {
	var err error
	const operationName = "VaultsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VaultsClientDeleteResponse{}, err
	}
	return VaultsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VaultsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the Vault details.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - options - VaultsClientGetOptions contains the optional parameters for the VaultsClient.Get method.
func (client *VaultsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientGetOptions) (VaultsClientGetResponse, error) {
	var err error
	const operationName = "VaultsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VaultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VaultsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VaultsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VaultsClient) getHandleResponse(resp *http.Response) (VaultsClientGetResponse, error) {
	result := VaultsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Vault); err != nil {
		return VaultsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Retrieve a list of Vaults.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - VaultsClientListByResourceGroupOptions contains the optional parameters for the VaultsClient.NewListByResourceGroupPager
//     method.
func (client *VaultsClient) NewListByResourceGroupPager(resourceGroupName string, options *VaultsClientListByResourceGroupOptions) *runtime.Pager[VaultsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultsClientListByResourceGroupResponse]{
		More: func(page VaultsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultsClientListByResourceGroupResponse) (VaultsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VaultsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return VaultsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VaultsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VaultsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VaultsClient) listByResourceGroupHandleResponse(resp *http.Response) (VaultsClientListByResourceGroupResponse, error) {
	result := VaultsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultList); err != nil {
		return VaultsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionIDPager - Fetches all the resources of the specified type in the subscription.
//
// Generated from API version 2024-03-01
//   - options - VaultsClientListBySubscriptionIDOptions contains the optional parameters for the VaultsClient.NewListBySubscriptionIDPager
//     method.
func (client *VaultsClient) NewListBySubscriptionIDPager(options *VaultsClientListBySubscriptionIDOptions) *runtime.Pager[VaultsClientListBySubscriptionIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[VaultsClientListBySubscriptionIDResponse]{
		More: func(page VaultsClientListBySubscriptionIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VaultsClientListBySubscriptionIDResponse) (VaultsClientListBySubscriptionIDResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VaultsClient.NewListBySubscriptionIDPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionIDCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return VaultsClientListBySubscriptionIDResponse{}, err
			}
			return client.listBySubscriptionIDHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *VaultsClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *VaultsClientListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/vaults"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *VaultsClient) listBySubscriptionIDHandleResponse(resp *http.Response) (VaultsClientListBySubscriptionIDResponse, error) {
	result := VaultsClientListBySubscriptionIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VaultList); err != nil {
		return VaultsClientListBySubscriptionIDResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - vaultName - The name of the recovery services vault.
//   - vault - Recovery Services Vault to be created.
//   - options - VaultsClientBeginUpdateOptions contains the optional parameters for the VaultsClient.BeginUpdate method.
func (client *VaultsClient) BeginUpdate(ctx context.Context, resourceGroupName string, vaultName string, vault PatchVault, options *VaultsClientBeginUpdateOptions) (*runtime.Poller[VaultsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, vaultName, vault, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VaultsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VaultsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
func (client *VaultsClient) update(ctx context.Context, resourceGroupName string, vaultName string, vault PatchVault, options *VaultsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VaultsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, vault, options)
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
func (client *VaultsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, vault PatchVault, options *VaultsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{vaultName}"
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
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, vault); err != nil {
		return nil, err
	}
	return req, nil
}
