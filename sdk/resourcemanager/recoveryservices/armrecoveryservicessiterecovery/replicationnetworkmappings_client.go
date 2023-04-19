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

// ReplicationNetworkMappingsClient contains the methods for the ReplicationNetworkMappings group.
// Don't use this type directly, use NewReplicationNetworkMappingsClient() instead.
type ReplicationNetworkMappingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationNetworkMappingsClient creates a new instance of ReplicationNetworkMappingsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationNetworkMappingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationNetworkMappingsClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicationNetworkMappingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationNetworkMappingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - The operation to create an ASR network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Primary fabric name.
//   - networkName - Primary network name.
//   - networkMappingName - Network mapping name.
//   - input - Create network mapping input.
//   - options - ReplicationNetworkMappingsClientBeginCreateOptions contains the optional parameters for the ReplicationNetworkMappingsClient.BeginCreate
//     method.
func (client *ReplicationNetworkMappingsClient) BeginCreate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput, options *ReplicationNetworkMappingsClientBeginCreateOptions) (*runtime.Poller[ReplicationNetworkMappingsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, input, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationNetworkMappingsClientCreateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationNetworkMappingsClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - The operation to create an ASR network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ReplicationNetworkMappingsClient) create(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput, options *ReplicationNetworkMappingsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, input, options)
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
func (client *ReplicationNetworkMappingsClient) createCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, input CreateNetworkMappingInput, options *ReplicationNetworkMappingsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
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
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, input)
}

// BeginDelete - The operation to delete a network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Primary fabric name.
//   - networkName - Primary network name.
//   - networkMappingName - ARM Resource Name for network mapping.
//   - options - ReplicationNetworkMappingsClientBeginDeleteOptions contains the optional parameters for the ReplicationNetworkMappingsClient.BeginDelete
//     method.
func (client *ReplicationNetworkMappingsClient) BeginDelete(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsClientBeginDeleteOptions) (*runtime.Poller[ReplicationNetworkMappingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationNetworkMappingsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationNetworkMappingsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete a network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ReplicationNetworkMappingsClient) deleteOperation(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, options)
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
func (client *ReplicationNetworkMappingsClient) deleteCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
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
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets the details of an ASR network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Primary fabric name.
//   - networkName - Primary network name.
//   - networkMappingName - Network mapping name.
//   - options - ReplicationNetworkMappingsClientGetOptions contains the optional parameters for the ReplicationNetworkMappingsClient.Get
//     method.
func (client *ReplicationNetworkMappingsClient) Get(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsClientGetOptions) (ReplicationNetworkMappingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, options)
	if err != nil {
		return ReplicationNetworkMappingsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationNetworkMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationNetworkMappingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationNetworkMappingsClient) getCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, options *ReplicationNetworkMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
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
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationNetworkMappingsClient) getHandleResponse(resp *http.Response) (ReplicationNetworkMappingsClientGetResponse, error) {
	result := ReplicationNetworkMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkMapping); err != nil {
		return ReplicationNetworkMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all ASR network mappings in the vault.
//
// Generated from API version 2023-03-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - options - ReplicationNetworkMappingsClientListOptions contains the optional parameters for the ReplicationNetworkMappingsClient.NewListPager
//     method.
func (client *ReplicationNetworkMappingsClient) NewListPager(resourceName string, resourceGroupName string, options *ReplicationNetworkMappingsClientListOptions) *runtime.Pager[ReplicationNetworkMappingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationNetworkMappingsClientListResponse]{
		More: func(page ReplicationNetworkMappingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationNetworkMappingsClientListResponse) (ReplicationNetworkMappingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationNetworkMappingsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationNetworkMappingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationNetworkMappingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationNetworkMappingsClient) listCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, options *ReplicationNetworkMappingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationNetworkMappings"
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
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationNetworkMappingsClient) listHandleResponse(resp *http.Response) (ReplicationNetworkMappingsClientListResponse, error) {
	result := ReplicationNetworkMappingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkMappingCollection); err != nil {
		return ReplicationNetworkMappingsClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationNetworksPager - Lists all ASR network mappings for the specified network.
//
// Generated from API version 2023-03-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Primary fabric name.
//   - networkName - Primary network name.
//   - options - ReplicationNetworkMappingsClientListByReplicationNetworksOptions contains the optional parameters for the ReplicationNetworkMappingsClient.NewListByReplicationNetworksPager
//     method.
func (client *ReplicationNetworkMappingsClient) NewListByReplicationNetworksPager(resourceName string, resourceGroupName string, fabricName string, networkName string, options *ReplicationNetworkMappingsClientListByReplicationNetworksOptions) *runtime.Pager[ReplicationNetworkMappingsClientListByReplicationNetworksResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationNetworkMappingsClientListByReplicationNetworksResponse]{
		More: func(page ReplicationNetworkMappingsClientListByReplicationNetworksResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationNetworkMappingsClientListByReplicationNetworksResponse) (ReplicationNetworkMappingsClientListByReplicationNetworksResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationNetworksCreateRequest(ctx, resourceName, resourceGroupName, fabricName, networkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationNetworkMappingsClientListByReplicationNetworksResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicationNetworkMappingsClientListByReplicationNetworksResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationNetworkMappingsClientListByReplicationNetworksResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationNetworksHandleResponse(resp)
		},
	})
}

// listByReplicationNetworksCreateRequest creates the ListByReplicationNetworks request.
func (client *ReplicationNetworkMappingsClient) listByReplicationNetworksCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, options *ReplicationNetworkMappingsClientListByReplicationNetworksOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings"
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
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationNetworksHandleResponse handles the ListByReplicationNetworks response.
func (client *ReplicationNetworkMappingsClient) listByReplicationNetworksHandleResponse(resp *http.Response) (ReplicationNetworkMappingsClientListByReplicationNetworksResponse, error) {
	result := ReplicationNetworkMappingsClientListByReplicationNetworksResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkMappingCollection); err != nil {
		return ReplicationNetworkMappingsClientListByReplicationNetworksResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update an ASR network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - fabricName - Primary fabric name.
//   - networkName - Primary network name.
//   - networkMappingName - Network mapping name.
//   - input - Update network mapping input.
//   - options - ReplicationNetworkMappingsClientBeginUpdateOptions contains the optional parameters for the ReplicationNetworkMappingsClient.BeginUpdate
//     method.
func (client *ReplicationNetworkMappingsClient) BeginUpdate(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput, options *ReplicationNetworkMappingsClientBeginUpdateOptions) (*runtime.Poller[ReplicationNetworkMappingsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, input, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationNetworkMappingsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationNetworkMappingsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update an ASR network mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01
func (client *ReplicationNetworkMappingsClient) update(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput, options *ReplicationNetworkMappingsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceName, resourceGroupName, fabricName, networkName, networkMappingName, input, options)
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

// updateCreateRequest creates the Update request.
func (client *ReplicationNetworkMappingsClient) updateCreateRequest(ctx context.Context, resourceName string, resourceGroupName string, fabricName string, networkName string, networkMappingName string, input UpdateNetworkMappingInput, options *ReplicationNetworkMappingsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationNetworks/{networkName}/replicationNetworkMappings/{networkMappingName}"
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
	if networkName == "" {
		return nil, errors.New("parameter networkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkName}", url.PathEscape(networkName))
	if networkMappingName == "" {
		return nil, errors.New("parameter networkMappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkMappingName}", url.PathEscape(networkMappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, input)
}
