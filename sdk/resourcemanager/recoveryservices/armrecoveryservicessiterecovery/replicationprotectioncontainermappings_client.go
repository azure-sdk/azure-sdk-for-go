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
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ReplicationProtectionContainerMappingsClient contains the methods for the ReplicationProtectionContainerMappings group.
// Don't use this type directly, use NewReplicationProtectionContainerMappingsClient() instead.
type ReplicationProtectionContainerMappingsClient struct {
	host              string
	resourceName      string
	resourceGroupName string
	subscriptionID    string
	pl                runtime.Pipeline
}

// NewReplicationProtectionContainerMappingsClient creates a new instance of ReplicationProtectionContainerMappingsClient with the specified values.
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicationProtectionContainerMappingsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicationProtectionContainerMappingsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicationProtectionContainerMappingsClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		host:              ep,
		pl:                pl,
	}
	return client, nil
}

// BeginCreate - The operation to create a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - creationInput - Mapping creation input.
//   - options - ReplicationProtectionContainerMappingsClientBeginCreateOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginCreate
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginCreate(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginCreateOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, fabricName, protectionContainerName, mappingName, creationInput, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationProtectionContainerMappingsClientCreateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationProtectionContainerMappingsClientCreateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Create - The operation to create a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationProtectionContainerMappingsClient) create(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, fabricName, protectionContainerName, mappingName, creationInput, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *ReplicationProtectionContainerMappingsClient) createCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, creationInput CreateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
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
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, creationInput)
}

// BeginDelete - The operation to delete or remove a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - removalInput - Removal input.
//   - options - ReplicationProtectionContainerMappingsClientBeginDeleteOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginDelete
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginDelete(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginDeleteOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, fabricName, protectionContainerName, mappingName, removalInput, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationProtectionContainerMappingsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationProtectionContainerMappingsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - The operation to delete or remove a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationProtectionContainerMappingsClient) deleteOperation(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, fabricName, protectionContainerName, mappingName, removalInput, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ReplicationProtectionContainerMappingsClient) deleteCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, removalInput RemoveProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}/remove"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
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
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, runtime.MarshalAsJSON(req, removalInput)
}

// Get - Gets the details of a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection Container mapping name.
//   - options - ReplicationProtectionContainerMappingsClientGetOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.Get
//     method.
func (client *ReplicationProtectionContainerMappingsClient) Get(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientGetOptions) (ReplicationProtectionContainerMappingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, fabricName, protectionContainerName, mappingName, options)
	if err != nil {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ReplicationProtectionContainerMappingsClient) getCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
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
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ReplicationProtectionContainerMappingsClient) getHandleResponse(resp *http.Response) (ReplicationProtectionContainerMappingsClientGetResponse, error) {
	result := ReplicationProtectionContainerMappingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerMapping); err != nil {
		return ReplicationProtectionContainerMappingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the protection container mappings in the vault.
//
// Generated from API version 2023-02-01
//   - options - ReplicationProtectionContainerMappingsClientListOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.NewListPager
//     method.
func (client *ReplicationProtectionContainerMappingsClient) NewListPager(options *ReplicationProtectionContainerMappingsClientListOptions) *runtime.Pager[ReplicationProtectionContainerMappingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionContainerMappingsClientListResponse]{
		More: func(page ReplicationProtectionContainerMappingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionContainerMappingsClientListResponse) (ReplicationProtectionContainerMappingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationProtectionContainerMappingsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationProtectionContainerMappingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationProtectionContainerMappingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ReplicationProtectionContainerMappingsClient) listCreateRequest(ctx context.Context, options *ReplicationProtectionContainerMappingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationProtectionContainerMappings"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ReplicationProtectionContainerMappingsClient) listHandleResponse(resp *http.Response) (ReplicationProtectionContainerMappingsClientListResponse, error) {
	result := ReplicationProtectionContainerMappingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerMappingCollection); err != nil {
		return ReplicationProtectionContainerMappingsClientListResponse{}, err
	}
	return result, nil
}

// NewListByReplicationProtectionContainersPager - Lists the protection container mappings for a protection container.
//
// Generated from API version 2023-02-01
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - options - ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersOptions contains the optional
//     parameters for the ReplicationProtectionContainerMappingsClient.NewListByReplicationProtectionContainersPager method.
func (client *ReplicationProtectionContainerMappingsClient) NewListByReplicationProtectionContainersPager(fabricName string, protectionContainerName string, options *ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersOptions) *runtime.Pager[ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse]{
		More: func(page ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse) (ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByReplicationProtectionContainersCreateRequest(ctx, fabricName, protectionContainerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByReplicationProtectionContainersHandleResponse(resp)
		},
	})
}

// listByReplicationProtectionContainersCreateRequest creates the ListByReplicationProtectionContainers request.
func (client *ReplicationProtectionContainerMappingsClient) listByReplicationProtectionContainersCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, options *ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationProtectionContainersHandleResponse handles the ListByReplicationProtectionContainers response.
func (client *ReplicationProtectionContainerMappingsClient) listByReplicationProtectionContainersHandleResponse(resp *http.Response) (ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse, error) {
	result := ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProtectionContainerMappingCollection); err != nil {
		return ReplicationProtectionContainerMappingsClientListByReplicationProtectionContainersResponse{}, err
	}
	return result, nil
}

// BeginPurge - The operation to purge(force delete) a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - options - ReplicationProtectionContainerMappingsClientBeginPurgeOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginPurge
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginPurge(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientBeginPurgeOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientPurgeResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.purge(ctx, fabricName, protectionContainerName, mappingName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationProtectionContainerMappingsClientPurgeResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationProtectionContainerMappingsClientPurgeResponse](options.ResumeToken, client.pl, nil)
	}
}

// Purge - The operation to purge(force delete) a protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationProtectionContainerMappingsClient) purge(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientBeginPurgeOptions) (*http.Response, error) {
	req, err := client.purgeCreateRequest(ctx, fabricName, protectionContainerName, mappingName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// purgeCreateRequest creates the Purge request.
func (client *ReplicationProtectionContainerMappingsClient) purgeCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, options *ReplicationProtectionContainerMappingsClientBeginPurgeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
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
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// BeginUpdate - The operation to update protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - mappingName - Protection container mapping name.
//   - updateInput - Mapping update input.
//   - options - ReplicationProtectionContainerMappingsClientBeginUpdateOptions contains the optional parameters for the ReplicationProtectionContainerMappingsClient.BeginUpdate
//     method.
func (client *ReplicationProtectionContainerMappingsClient) BeginUpdate(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginUpdateOptions) (*runtime.Poller[ReplicationProtectionContainerMappingsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, fabricName, protectionContainerName, mappingName, updateInput, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ReplicationProtectionContainerMappingsClientUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ReplicationProtectionContainerMappingsClientUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Update - The operation to update protection container mapping.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
func (client *ReplicationProtectionContainerMappingsClient) update(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, fabricName, protectionContainerName, mappingName, updateInput, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *ReplicationProtectionContainerMappingsClient) updateCreateRequest(ctx context.Context, fabricName string, protectionContainerName string, mappingName string, updateInput UpdateProtectionContainerMappingInput, options *ReplicationProtectionContainerMappingsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionContainerMappings/{mappingName}"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
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
	if mappingName == "" {
		return nil, errors.New("parameter mappingName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mappingName}", url.PathEscape(mappingName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, updateInput)
}
