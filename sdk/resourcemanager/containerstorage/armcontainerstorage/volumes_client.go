//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerstorage

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

// VolumesClient contains the methods for the Volumes group.
// Don't use this type directly, use NewVolumesClient() instead.
type VolumesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVolumesClient creates a new instance of VolumesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVolumesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VolumesClient, error) {
	cl, err := arm.NewClient(moduleName+".VolumesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VolumesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - poolName - Pool Object
//   - volumeName - Volume Resource
//   - resource - Resource create parameters.
//   - options - VolumesClientBeginCreateOrUpdateOptions contains the optional parameters for the VolumesClient.BeginCreateOrUpdate
//     method.
func (client *VolumesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, poolName string, volumeName string, resource Volume, options *VolumesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VolumesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, poolName, volumeName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VolumesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *VolumesClient) createOrUpdate(ctx context.Context, resourceGroupName string, poolName string, volumeName string, resource Volume, options *VolumesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, poolName, volumeName, resource, options)
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
func (client *VolumesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, poolName string, volumeName string, resource Volume, options *VolumesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerStorage/pools/{poolName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - poolName - Pool Object
//   - volumeName - Volume Resource
//   - options - VolumesClientBeginDeleteOptions contains the optional parameters for the VolumesClient.BeginDelete method.
func (client *VolumesClient) BeginDelete(ctx context.Context, resourceGroupName string, poolName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*runtime.Poller[VolumesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, poolName, volumeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VolumesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *VolumesClient) deleteOperation(ctx context.Context, resourceGroupName string, poolName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, poolName, volumeName, options)
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
func (client *VolumesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, poolName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerStorage/pools/{poolName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - poolName - Pool Object
//   - volumeName - Volume Resource
//   - options - VolumesClientGetOptions contains the optional parameters for the VolumesClient.Get method.
func (client *VolumesClient) Get(ctx context.Context, resourceGroupName string, poolName string, volumeName string, options *VolumesClientGetOptions) (VolumesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, poolName, volumeName, options)
	if err != nil {
		return VolumesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VolumesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VolumesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VolumesClient) getCreateRequest(ctx context.Context, resourceGroupName string, poolName string, volumeName string, options *VolumesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerStorage/pools/{poolName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VolumesClient) getHandleResponse(resp *http.Response) (VolumesClientGetResponse, error) {
	result := VolumesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Volume); err != nil {
		return VolumesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByPoolPager - List Volume resources by Pool
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - poolName - Pool Object
//   - options - VolumesClientListByPoolOptions contains the optional parameters for the VolumesClient.NewListByPoolPager method.
func (client *VolumesClient) NewListByPoolPager(resourceGroupName string, poolName string, options *VolumesClientListByPoolOptions) *runtime.Pager[VolumesClientListByPoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumesClientListByPoolResponse]{
		More: func(page VolumesClientListByPoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VolumesClientListByPoolResponse) (VolumesClientListByPoolResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPoolCreateRequest(ctx, resourceGroupName, poolName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VolumesClientListByPoolResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VolumesClientListByPoolResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumesClientListByPoolResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPoolHandleResponse(resp)
		},
	})
}

// listByPoolCreateRequest creates the ListByPool request.
func (client *VolumesClient) listByPoolCreateRequest(ctx context.Context, resourceGroupName string, poolName string, options *VolumesClientListByPoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerStorage/pools/{poolName}/volumes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPoolHandleResponse handles the ListByPool response.
func (client *VolumesClient) listByPoolHandleResponse(resp *http.Response) (VolumesClientListByPoolResponse, error) {
	result := VolumesClientListByPoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeListResult); err != nil {
		return VolumesClientListByPoolResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - poolName - Pool Object
//   - volumeName - Volume Resource
//   - properties - The resource properties to be updated.
//   - options - VolumesClientBeginUpdateOptions contains the optional parameters for the VolumesClient.BeginUpdate method.
func (client *VolumesClient) BeginUpdate(ctx context.Context, resourceGroupName string, poolName string, volumeName string, properties VolumeUpdate, options *VolumesClientBeginUpdateOptions) (*runtime.Poller[VolumesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, poolName, volumeName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VolumesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a Volume
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
func (client *VolumesClient) update(ctx context.Context, resourceGroupName string, poolName string, volumeName string, properties VolumeUpdate, options *VolumesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, poolName, volumeName, properties, options)
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
func (client *VolumesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, poolName string, volumeName string, properties VolumeUpdate, options *VolumesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerStorage/pools/{poolName}/volumes/{volumeName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if poolName == "" {
		return nil, errors.New("parameter poolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{poolName}", url.PathEscape(poolName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
		return nil, err
	}
	return req, nil
}
