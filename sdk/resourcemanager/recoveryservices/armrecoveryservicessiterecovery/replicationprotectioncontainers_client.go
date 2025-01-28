//go:build go1.18
// +build go1.18

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

// ReplicationProtectionContainersClient contains the methods for the ReplicationProtectionContainers group.
// Don't use this type directly, use NewReplicationProtectionContainersClient() instead.
type ReplicationProtectionContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicationProtectionContainersClient creates a new instance of ReplicationProtectionContainersClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationProtectionContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationProtectionContainersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationProtectionContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Operation to create a protection container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Unique fabric ARM name.
//   - protectionContainerName - Unique protection container ARM name.
//   - creationInput - Creation input.
//   - options - ReplicationProtectionContainersClientBeginCreateOptions contains the optional parameters for the ReplicationProtectionContainersClient.BeginCreate
//     method.
func (client *ReplicationProtectionContainersClient) BeginCreate(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput, options *ReplicationProtectionContainersClientBeginCreateOptions) (*runtime.Poller[ReplicationProtectionContainersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, creationInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainersClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainersClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Operation to create a protection container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ReplicationProtectionContainersClient) create(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput, options *ReplicationProtectionContainersClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainersClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, creationInput, options)
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

// createCreateRequest creates the Create request.
func (client *ReplicationProtectionContainersClient) createCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, creationInput CreateProtectionContainerInput, options *ReplicationProtectionContainersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, creationInput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Operation to remove a protection container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Unique fabric ARM name.
//   - protectionContainerName - Unique protection container ARM name.
//   - options - ReplicationProtectionContainersClientBeginDeleteOptions contains the optional parameters for the ReplicationProtectionContainersClient.BeginDelete
//     method.
func (client *ReplicationProtectionContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainersClientBeginDeleteOptions) (*runtime.Poller[ReplicationProtectionContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Operation to remove a protection container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ReplicationProtectionContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, options)
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
func (client *ReplicationProtectionContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/remove"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginDiscoverProtectableItem - The operation to a add a protectable item to a protection container(Add physical server).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - The name of the fabric.
//   - protectionContainerName - The name of the protection container.
//   - discoverProtectableItemRequest - The request object to add a protectable item.
//   - options - ReplicationProtectionContainersClientBeginDiscoverProtectableItemOptions contains the optional parameters for
//     the ReplicationProtectionContainersClient.BeginDiscoverProtectableItem method.
func (client *ReplicationProtectionContainersClient) BeginDiscoverProtectableItem(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest, options *ReplicationProtectionContainersClientBeginDiscoverProtectableItemOptions) (*runtime.Poller[ReplicationProtectionContainersClientDiscoverProtectableItemResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.discoverProtectableItem(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, discoverProtectableItemRequest, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainersClientDiscoverProtectableItemResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainersClientDiscoverProtectableItemResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DiscoverProtectableItem - The operation to a add a protectable item to a protection container(Add physical server).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ReplicationProtectionContainersClient) discoverProtectableItem(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest, options *ReplicationProtectionContainersClientBeginDiscoverProtectableItemOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainersClient.BeginDiscoverProtectableItem"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.discoverProtectableItemCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, discoverProtectableItemRequest, options)
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

// discoverProtectableItemCreateRequest creates the DiscoverProtectableItem request.
func (client *ReplicationProtectionContainersClient) discoverProtectableItemCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, discoverProtectableItemRequest DiscoverProtectableItemRequest, options *ReplicationProtectionContainersClientBeginDiscoverProtectableItemOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/discoverProtectableItem"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, discoverProtectableItemRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// Get - Gets the details of a protection container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - options - ReplicationProtectionContainersClientGetOptions contains the optional parameters for the ReplicationProtectionContainersClient.Get
//     method.
func (client *ReplicationProtectionContainersClient) Get(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainersClientGetOptions) (ReplicationProtectionContainersClientGetResponse, error) {
	var err error
	const operationName = "ReplicationProtectionContainersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, options)
	if err != nil {
		return ReplicationProtectionContainersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ReplicationProtectionContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ReplicationProtectionContainersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ReplicationProtectionContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, options *ReplicationProtectionContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationProtectionContainersClient) getHandleResponse(resp *http.Response) (ReplicationProtectionContainersClientGetResponse, error) {
	result := ReplicationProtectionContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainer); err != nil {
		return ReplicationProtectionContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the protection containers in a vault.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - options - ReplicationProtectionContainersClientListOptions contains the optional parameters for the ReplicationProtectionContainersClient.NewListPager
//     method.
func (client *ReplicationProtectionContainersClient) NewListPager(resourceGroupName string, resourceName string, options *ReplicationProtectionContainersClientListOptions) *runtime.Pager[ReplicationProtectionContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionContainersClientListResponse]{
		More: func(page ReplicationProtectionContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionContainersClientListResponse) (ReplicationProtectionContainersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationProtectionContainersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return ReplicationProtectionContainersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationProtectionContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ReplicationProtectionContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainers"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationProtectionContainersClient) listHandleResponse(resp *http.Response) (ReplicationProtectionContainersClientListResponse, error) {
	result := ReplicationProtectionContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerCollection); err != nil {
		return ReplicationProtectionContainersClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationFabricsPager - Lists the protection containers in the specified fabric.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Fabric name.
//   - options - ReplicationProtectionContainersClientListByReplicationFabricsOptions contains the optional parameters for the
//     ReplicationProtectionContainersClient.NewListByReplicationFabricsPager method.
func (client *ReplicationProtectionContainersClient) NewListByReplicationFabricsPager(resourceGroupName string, resourceName string, fabricName string, options *ReplicationProtectionContainersClientListByReplicationFabricsOptions) *runtime.Pager[ReplicationProtectionContainersClientListByReplicationFabricsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionContainersClientListByReplicationFabricsResponse]{
		More: func(page ReplicationProtectionContainersClientListByReplicationFabricsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionContainersClientListByReplicationFabricsResponse) (ReplicationProtectionContainersClientListByReplicationFabricsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ReplicationProtectionContainersClient.NewListByReplicationFabricsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByReplicationFabricsCreateRequest(ctx, resourceGroupName, resourceName, fabricName, options)
			}, nil)
			if err != nil {
				return ReplicationProtectionContainersClientListByReplicationFabricsResponse{}, err
			}
			return client.listByReplicationFabricsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByReplicationFabricsCreateRequest creates the ListByReplicationFabrics request.
func (client *ReplicationProtectionContainersClient) listByReplicationFabricsCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, options *ReplicationProtectionContainersClientListByReplicationFabricsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationFabricsHandleResponse handles the ListByReplicationFabrics response.
func (client *ReplicationProtectionContainersClient) listByReplicationFabricsHandleResponse(resp *http.Response) (ReplicationProtectionContainersClientListByReplicationFabricsResponse, error) {
	result := ReplicationProtectionContainersClientListByReplicationFabricsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerCollection); err != nil {
		return ReplicationProtectionContainersClientListByReplicationFabricsResponse{}, err
	}
	return result, nil
}

// BeginSwitchClusterProtection - Operation to switch protection from one container to another.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - switchInput - Switch protection input.
//   - options - ReplicationProtectionContainersClientBeginSwitchClusterProtectionOptions contains the optional parameters for
//     the ReplicationProtectionContainersClient.BeginSwitchClusterProtection method.
func (client *ReplicationProtectionContainersClient) BeginSwitchClusterProtection(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, switchInput SwitchClusterProtectionInput, options *ReplicationProtectionContainersClientBeginSwitchClusterProtectionOptions) (*runtime.Poller[ReplicationProtectionContainersClientSwitchClusterProtectionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.switchClusterProtection(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, switchInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainersClientSwitchClusterProtectionResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainersClientSwitchClusterProtectionResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// SwitchClusterProtection - Operation to switch protection from one container to another.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ReplicationProtectionContainersClient) switchClusterProtection(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, switchInput SwitchClusterProtectionInput, options *ReplicationProtectionContainersClientBeginSwitchClusterProtectionOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainersClient.BeginSwitchClusterProtection"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.switchClusterProtectionCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, switchInput, options)
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

// switchClusterProtectionCreateRequest creates the SwitchClusterProtection request.
func (client *ReplicationProtectionContainersClient) switchClusterProtectionCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, switchInput SwitchClusterProtectionInput, options *ReplicationProtectionContainersClientBeginSwitchClusterProtectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/switchClusterProtection"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, switchInput); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginSwitchProtection - Operation to switch protection from one container to another or one replication provider to another.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Unique fabric name.
//   - protectionContainerName - Protection container name.
//   - switchInput - Switch protection input.
//   - options - ReplicationProtectionContainersClientBeginSwitchProtectionOptions contains the optional parameters for the ReplicationProtectionContainersClient.BeginSwitchProtection
//     method.
func (client *ReplicationProtectionContainersClient) BeginSwitchProtection(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, switchInput SwitchProtectionInput, options *ReplicationProtectionContainersClientBeginSwitchProtectionOptions) (*runtime.Poller[ReplicationProtectionContainersClientSwitchProtectionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.switchProtection(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, switchInput, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ReplicationProtectionContainersClientSwitchProtectionResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ReplicationProtectionContainersClientSwitchProtectionResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// SwitchProtection - Operation to switch protection from one container to another or one replication provider to another.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01
func (client *ReplicationProtectionContainersClient) switchProtection(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, switchInput SwitchProtectionInput, options *ReplicationProtectionContainersClientBeginSwitchProtectionOptions) (*http.Response, error) {
	var err error
	const operationName = "ReplicationProtectionContainersClient.BeginSwitchProtection"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.switchProtectionCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, switchInput, options)
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

// switchProtectionCreateRequest creates the SwitchProtection request.
func (client *ReplicationProtectionContainersClient) switchProtectionCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, switchInput SwitchProtectionInput, options *ReplicationProtectionContainersClientBeginSwitchProtectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/switchprotection"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, switchInput); err != nil {
		return nil, err
	}
	return req, nil
}
