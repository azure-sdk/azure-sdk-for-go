//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// StorageContainersClient contains the methods for the StorageContainers group.
// Don't use this type directly, use NewStorageContainersClient() instead.
type StorageContainersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewStorageContainersClient creates a new instance of StorageContainersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewStorageContainersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*StorageContainersClient, error) {
	cl, err := arm.NewClient(moduleName+".StorageContainersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &StorageContainersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a storage container. Please note some properties can be set only
// during storage container creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageContainerName - Name of the storage container
//   - options - StorageContainersClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageContainersClient.BeginCreateOrUpdate
//     method.
func (client *StorageContainersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, storageContainerName string, storageContainers StorageContainers, options *StorageContainersClientBeginCreateOrUpdateOptions) (*runtime.Poller[StorageContainersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, storageContainerName, storageContainers, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageContainersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageContainersClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - The operation to create or update a storage container. Please note some properties can be set only during
// storage container creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *StorageContainersClient) createOrUpdate(ctx context.Context, resourceGroupName string, storageContainerName string, storageContainers StorageContainers, options *StorageContainersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, storageContainerName, storageContainers, options)
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
func (client *StorageContainersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, storageContainerName string, storageContainers StorageContainers, options *StorageContainersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/storageContainers/{storageContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageContainerName == "" {
		return nil, errors.New("parameter storageContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageContainerName}", url.PathEscape(storageContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, storageContainers); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a storage container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageContainerName - Name of the storage container
//   - options - StorageContainersClientBeginDeleteOptions contains the optional parameters for the StorageContainersClient.BeginDelete
//     method.
func (client *StorageContainersClient) BeginDelete(ctx context.Context, resourceGroupName string, storageContainerName string, options *StorageContainersClientBeginDeleteOptions) (*runtime.Poller[StorageContainersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, storageContainerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageContainersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageContainersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - The operation to delete a storage container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *StorageContainersClient) deleteOperation(ctx context.Context, resourceGroupName string, storageContainerName string, options *StorageContainersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, storageContainerName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageContainersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, storageContainerName string, options *StorageContainersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/storageContainers/{storageContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageContainerName == "" {
		return nil, errors.New("parameter storageContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageContainerName}", url.PathEscape(storageContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a storage container
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageContainerName - Name of the storage container
//   - options - StorageContainersClientGetOptions contains the optional parameters for the StorageContainersClient.Get method.
func (client *StorageContainersClient) Get(ctx context.Context, resourceGroupName string, storageContainerName string, options *StorageContainersClientGetOptions) (StorageContainersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, storageContainerName, options)
	if err != nil {
		return StorageContainersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return StorageContainersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return StorageContainersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *StorageContainersClient) getCreateRequest(ctx context.Context, resourceGroupName string, storageContainerName string, options *StorageContainersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/storageContainers/{storageContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageContainerName == "" {
		return nil, errors.New("parameter storageContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageContainerName}", url.PathEscape(storageContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageContainersClient) getHandleResponse(resp *http.Response) (StorageContainersClientGetResponse, error) {
	result := StorageContainersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainers); err != nil {
		return StorageContainersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the storage containers in the specified resource group. Use the nextLink property in the response
// to get the next page of storage containers.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - StorageContainersClientListOptions contains the optional parameters for the StorageContainersClient.NewListPager
//     method.
func (client *StorageContainersClient) NewListPager(resourceGroupName string, options *StorageContainersClientListOptions) *runtime.Pager[StorageContainersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageContainersClientListResponse]{
		More: func(page StorageContainersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageContainersClientListResponse) (StorageContainersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageContainersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return StorageContainersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageContainersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *StorageContainersClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *StorageContainersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/storageContainers"
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
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *StorageContainersClient) listHandleResponse(resp *http.Response) (StorageContainersClientListResponse, error) {
	result := StorageContainersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainersListResult); err != nil {
		return StorageContainersClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Lists all of the storage containers in the specified subscription. Use the nextLink property in the response
// to get the next page of storage containers.
//
// Generated from API version 2023-07-01-preview
//   - options - StorageContainersClientListAllOptions contains the optional parameters for the StorageContainersClient.NewListAllPager
//     method.
func (client *StorageContainersClient) NewListAllPager(options *StorageContainersClientListAllOptions) *runtime.Pager[StorageContainersClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[StorageContainersClientListAllResponse]{
		More: func(page StorageContainersClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *StorageContainersClientListAllResponse) (StorageContainersClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return StorageContainersClientListAllResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return StorageContainersClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return StorageContainersClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *StorageContainersClient) listAllCreateRequest(ctx context.Context, options *StorageContainersClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureStackHCI/storageContainers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *StorageContainersClient) listAllHandleResponse(resp *http.Response) (StorageContainersClientListAllResponse, error) {
	result := StorageContainersClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageContainersListResult); err != nil {
		return StorageContainersClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update a storage container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - storageContainerName - Name of the storage container
//   - options - StorageContainersClientBeginUpdateOptions contains the optional parameters for the StorageContainersClient.BeginUpdate
//     method.
func (client *StorageContainersClient) BeginUpdate(ctx context.Context, resourceGroupName string, storageContainerName string, storageContainers StorageContainersUpdateRequest, options *StorageContainersClientBeginUpdateOptions) (*runtime.Poller[StorageContainersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, storageContainerName, storageContainers, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[StorageContainersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[StorageContainersClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - The operation to update a storage container.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
func (client *StorageContainersClient) update(ctx context.Context, resourceGroupName string, storageContainerName string, storageContainers StorageContainersUpdateRequest, options *StorageContainersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, storageContainerName, storageContainers, options)
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
func (client *StorageContainersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, storageContainerName string, storageContainers StorageContainersUpdateRequest, options *StorageContainersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/storageContainers/{storageContainerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if storageContainerName == "" {
		return nil, errors.New("parameter storageContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageContainerName}", url.PathEscape(storageContainerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, storageContainers); err != nil {
		return nil, err
	}
	return req, nil
}
