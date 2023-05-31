//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ReplicationRecoveryServicesProvidersClient contains the methods for the ReplicationRecoveryServicesProviders group.
// Don't use this type directly, use NewReplicationRecoveryServicesProvidersClient() instead.
type ReplicationRecoveryServicesProvidersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationRecoveryServicesProvidersClient creates a new instance of ReplicationRecoveryServicesProvidersClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationRecoveryServicesProvidersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationRecoveryServicesProvidersClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicationRecoveryServicesProvidersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationRecoveryServicesProvidersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - The operation to add a recovery services provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - providerName - Recovery services provider name.
//   - addProviderInput - Add provider input.
//   - options - ReplicationRecoveryServicesProvidersClientBeginCreateOptions contains the optional parameters for the ReplicationRecoveryServicesProvidersClient.BeginCreate
//     method.
func (client *ReplicationRecoveryServicesProvidersClient) BeginCreate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, addProviderInput AddRecoveryServicesProviderInput, options *ReplicationRecoveryServicesProvidersClientBeginCreateOptions) (*runtime.Poller[ReplicationRecoveryServicesProvidersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceName, resourceGroupName, fabricName, providerName, addProviderInput, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationRecoveryServicesProvidersClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationRecoveryServicesProvidersClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - The operation to add a recovery services provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *ReplicationRecoveryServicesProvidersClient) create(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, addProviderInput AddRecoveryServicesProviderInput, options *ReplicationRecoveryServicesProvidersClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceName, resourceGroupName, fabricName, providerName, addProviderInput, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationRecoveryServicesProvidersClient) createCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, addProviderInput AddRecoveryServicesProviderInput, options *ReplicationRecoveryServicesProvidersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, addProviderInput)
}

// BeginDelete - The operation to removes/delete(unregister) a recovery services provider from the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - providerName - Recovery services provider name.
//   - options - ReplicationRecoveryServicesProvidersClientBeginDeleteOptions contains the optional parameters for the ReplicationRecoveryServicesProvidersClient.BeginDelete
//     method.
func (client *ReplicationRecoveryServicesProvidersClient) BeginDelete(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginDeleteOptions) (*runtime.Poller[ReplicationRecoveryServicesProvidersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationRecoveryServicesProvidersClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationRecoveryServicesProvidersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to removes/delete(unregister) a recovery services provider from the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *ReplicationRecoveryServicesProvidersClient) deleteOperation(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationRecoveryServicesProvidersClient) deleteCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}/remove"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the details of registered recovery services provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - providerName - Recovery services provider name.
//   - options - ReplicationRecoveryServicesProvidersClientGetOptions contains the optional parameters for the ReplicationRecoveryServicesProvidersClient.Get
//     method.
func (client *ReplicationRecoveryServicesProvidersClient) Get(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientGetOptions) (ReplicationRecoveryServicesProvidersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
	if err != nil {
		return ReplicationRecoveryServicesProvidersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationRecoveryServicesProvidersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationRecoveryServicesProvidersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationRecoveryServicesProvidersClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
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

// getHandleResponse handles the Get response.
func (client *ReplicationRecoveryServicesProvidersClient) getHandleResponse(resp *http.Response) (ReplicationRecoveryServicesProvidersClientGetResponse, error) {
	result := ReplicationRecoveryServicesProvidersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryServicesProvider); err != nil {
		return ReplicationRecoveryServicesProvidersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the registered recovery services providers in the vault.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationRecoveryServicesProvidersClientListOptions contains the optional parameters for the ReplicationRecoveryServicesProvidersClient.NewListPager
//     method.
func (client *ReplicationRecoveryServicesProvidersClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationRecoveryServicesProvidersClientListOptions) *runtime.Pager[ReplicationRecoveryServicesProvidersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationRecoveryServicesProvidersClientListResponse]{
		More: func(page ReplicationRecoveryServicesProvidersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationRecoveryServicesProvidersClientListResponse) (ReplicationRecoveryServicesProvidersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationRecoveryServicesProvidersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationRecoveryServicesProvidersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationRecoveryServicesProvidersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationRecoveryServicesProvidersClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationRecoveryServicesProvidersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationRecoveryServicesProviders"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationRecoveryServicesProvidersClient) listHandleResponse(resp *http.Response) (ReplicationRecoveryServicesProvidersClientListResponse, error) {
	result := ReplicationRecoveryServicesProvidersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryServicesProviderCollection); err != nil {
		return ReplicationRecoveryServicesProvidersClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationFabricsPager - Lists the registered recovery services providers for the specified fabric.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - options - ReplicationRecoveryServicesProvidersClientListByReplicationFabricsOptions contains the optional parameters for
//     the ReplicationRecoveryServicesProvidersClient.NewListByReplicationFabricsPager method.
func (client *ReplicationRecoveryServicesProvidersClient) NewListByReplicationFabricsPager(resourceName string, resourceGroupName string, fabricName string, options *ReplicationRecoveryServicesProvidersClientListByReplicationFabricsOptions) *runtime.Pager[ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse]{
		More: func(page ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse) (ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationFabricsCreateRequest(ctx, resourceName, resourceGroupName, fabricName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationFabricsHandleResponse(resp)
		},
	})
}

// listByReplicationFabricsCreateRequest creates the ListByReplicationFabrics request.
func (client *ReplicationRecoveryServicesProvidersClient) listByReplicationFabricsCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, options *ReplicationRecoveryServicesProvidersClientListByReplicationFabricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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

// listByReplicationFabricsHandleResponse handles the ListByReplicationFabrics response.
func (client *ReplicationRecoveryServicesProvidersClient) listByReplicationFabricsHandleResponse(resp *http.Response) (ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse, error) {
	result := ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoveryServicesProviderCollection); err != nil {
		return ReplicationRecoveryServicesProvidersClientListByReplicationFabricsResponse{}, err
	}
	return result, nil
}

// BeginPurge - The operation to purge(force delete) a recovery services provider from the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - providerName - Recovery services provider name.
//   - options - ReplicationRecoveryServicesProvidersClientBeginPurgeOptions contains the optional parameters for the ReplicationRecoveryServicesProvidersClient.BeginPurge
//     method.
func (client *ReplicationRecoveryServicesProvidersClient) BeginPurge(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginPurgeOptions) (*runtime.Poller[ReplicationRecoveryServicesProvidersClientPurgeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purge(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationRecoveryServicesProvidersClientPurgeResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationRecoveryServicesProvidersClientPurgeResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Purge - The operation to purge(force delete) a recovery services provider from the vault.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *ReplicationRecoveryServicesProvidersClient) purge(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginPurgeOptions) (*http.Response, error) {
	req, err := client.purgeCreateRequest(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// purgeCreateRequest creates the Purge request.
func (client *ReplicationRecoveryServicesProvidersClient) purgeCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginPurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginRefreshProvider - The operation to refresh the information from the recovery services provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Fabric name.
//   - providerName - Recovery services provider name.
//   - options - ReplicationRecoveryServicesProvidersClientBeginRefreshProviderOptions contains the optional parameters for the
//     ReplicationRecoveryServicesProvidersClient.BeginRefreshProvider method.
func (client *ReplicationRecoveryServicesProvidersClient) BeginRefreshProvider(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginRefreshProviderOptions) (*runtime.Poller[ReplicationRecoveryServicesProvidersClientRefreshProviderResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.refreshProvider(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationRecoveryServicesProvidersClientRefreshProviderResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationRecoveryServicesProvidersClientRefreshProviderResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// RefreshProvider - The operation to refresh the information from the recovery services provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01
func (client *ReplicationRecoveryServicesProvidersClient) refreshProvider(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginRefreshProviderOptions) (*http.Response, error) {
	req, err := client.refreshProviderCreateRequest(ctx, resourceName, resourceGroupName, fabricName, providerName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// refreshProviderCreateRequest creates the RefreshProvider request.
func (client *ReplicationRecoveryServicesProvidersClient) refreshProviderCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, providerName string, options *ReplicationRecoveryServicesProvidersClientBeginRefreshProviderOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationRecoveryServicesProviders/{providerName}/refreshProvider"
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
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
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
