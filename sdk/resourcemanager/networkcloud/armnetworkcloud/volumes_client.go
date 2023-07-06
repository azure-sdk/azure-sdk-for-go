//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetworkcloud

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
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
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

// BeginCreateOrUpdate - Create a new volume or update the properties of the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - volumeName - The name of the volume.
//   - volumeParameters - The request body.
//   - options - VolumesClientBeginCreateOrUpdateOptions contains the optional parameters for the VolumesClient.BeginCreateOrUpdate
//     method.
func (client *VolumesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, volumeName string, volumeParameters Volume, options *VolumesClientBeginCreateOrUpdateOptions) (*runtime.Poller[VolumesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, volumeName, volumeParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VolumesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new volume or update the properties of the existing one.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *VolumesClient) createOrUpdate(ctx context.Context, resourceGroupName string, volumeName string, volumeParameters Volume, options *VolumesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, volumeName, volumeParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VolumesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, volumeName string, volumeParameters Volume, options *VolumesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/volumes/{volumeName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, volumeParameters)
}

// BeginDelete - Delete the provided volume.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - volumeName - The name of the volume.
//   - options - VolumesClientBeginDeleteOptions contains the optional parameters for the VolumesClient.BeginDelete method.
func (client *VolumesClient) BeginDelete(ctx context.Context, resourceGroupName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*runtime.Poller[VolumesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, volumeName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VolumesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VolumesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided volume.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *VolumesClient) deleteOperation(ctx context.Context, resourceGroupName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, volumeName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *VolumesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, volumeName string, options *VolumesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/volumes/{volumeName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get properties of the provided volume.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - volumeName - The name of the volume.
//   - options - VolumesClientGetOptions contains the optional parameters for the VolumesClient.Get method.
func (client *VolumesClient) Get(ctx context.Context, resourceGroupName string, volumeName string, options *VolumesClientGetOptions) (VolumesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, volumeName, options)
	if err != nil {
		return VolumesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VolumesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VolumesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VolumesClient) getCreateRequest(ctx context.Context, resourceGroupName string, volumeName string, options *VolumesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/volumes/{volumeName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
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

// NewListByResourceGroupPager - Get a list of volumes in the provided resource group.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - VolumesClientListByResourceGroupOptions contains the optional parameters for the VolumesClient.NewListByResourceGroupPager
//     method.
func (client *VolumesClient) NewListByResourceGroupPager(resourceGroupName string, options *VolumesClientListByResourceGroupOptions) *runtime.Pager[VolumesClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumesClientListByResourceGroupResponse]{
		More: func(page VolumesClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VolumesClientListByResourceGroupResponse) (VolumesClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VolumesClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VolumesClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumesClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VolumesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VolumesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/volumes"
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
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VolumesClient) listByResourceGroupHandleResponse(resp *http.Response) (VolumesClientListByResourceGroupResponse, error) {
	result := VolumesClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeList); err != nil {
		return VolumesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of volumes in the provided subscription.
//
// Generated from API version 2023-07-01
//   - options - VolumesClientListBySubscriptionOptions contains the optional parameters for the VolumesClient.NewListBySubscriptionPager
//     method.
func (client *VolumesClient) NewListBySubscriptionPager(options *VolumesClientListBySubscriptionOptions) *runtime.Pager[VolumesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VolumesClientListBySubscriptionResponse]{
		More: func(page VolumesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VolumesClientListBySubscriptionResponse) (VolumesClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VolumesClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VolumesClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VolumesClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VolumesClient) listBySubscriptionCreateRequest(ctx context.Context, options *VolumesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/volumes"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VolumesClient) listBySubscriptionHandleResponse(resp *http.Response) (VolumesClientListBySubscriptionResponse, error) {
	result := VolumesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VolumeList); err != nil {
		return VolumesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update tags associated with the provided volume.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - volumeName - The name of the volume.
//   - volumeUpdateParameters - The request body.
//   - options - VolumesClientUpdateOptions contains the optional parameters for the VolumesClient.Update method.
func (client *VolumesClient) Update(ctx context.Context, resourceGroupName string, volumeName string, volumeUpdateParameters VolumePatchParameters, options *VolumesClientUpdateOptions) (VolumesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, volumeName, volumeUpdateParameters, options)
	if err != nil {
		return VolumesClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VolumesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VolumesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *VolumesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, volumeName string, volumeUpdateParameters VolumePatchParameters, options *VolumesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/volumes/{volumeName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if volumeName == "" {
		return nil, errors.New("parameter volumeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{volumeName}", url.PathEscape(volumeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, volumeUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *VolumesClient) updateHandleResponse(resp *http.Response) (VolumesClientUpdateResponse, error) {
	result := VolumesClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Volume); err != nil {
		return VolumesClientUpdateResponse{}, err
	}
	return result, nil
}
