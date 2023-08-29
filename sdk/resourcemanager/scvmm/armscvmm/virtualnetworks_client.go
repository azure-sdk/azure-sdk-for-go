//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armscvmm

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// VirtualNetworksClient contains the methods for the VirtualNetworks group.
// Don't use this type directly, use NewVirtualNetworksClient() instead.
type VirtualNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVirtualNetworksClient creates a new instance of VirtualNetworksClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVirtualNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualNetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".VirtualNetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Onboards the ScVmm virtual network as an Azure virtual network resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - Name of the VirtualNetwork.
//   - body - Request payload.
//   - options - VirtualNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the VirtualNetworksClient.BeginCreateOrUpdate
//     method.
func (client *VirtualNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, body VirtualNetwork, options *VirtualNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[VirtualNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, virtualNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Onboards the ScVmm virtual network as an Azure virtual network resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *VirtualNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, body VirtualNetwork, options *VirtualNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, virtualNetworkName, body, options)
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
func (client *VirtualNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, body VirtualNetwork, options *VirtualNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deregisters the ScVmm virtual network from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - Name of the VirtualNetwork.
//   - options - VirtualNetworksClientBeginDeleteOptions contains the optional parameters for the VirtualNetworksClient.BeginDelete
//     method.
func (client *VirtualNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*runtime.Poller[VirtualNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, virtualNetworkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deregisters the ScVmm virtual network from Azure.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *VirtualNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
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
func (client *VirtualNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Implements VirtualNetwork GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - Name of the VirtualNetwork.
//   - options - VirtualNetworksClientGetOptions contains the optional parameters for the VirtualNetworksClient.Get method.
func (client *VirtualNetworksClient) Get(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (VirtualNetworksClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, virtualNetworkName, options)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VirtualNetworksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VirtualNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, options *VirtualNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualNetworks/{virtualNetworkName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VirtualNetworksClient) getHandleResponse(resp *http.Response) (VirtualNetworksClientGetResponse, error) {
	result := VirtualNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetwork); err != nil {
		return VirtualNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List of VirtualNetworks in a resource group.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group.
//   - options - VirtualNetworksClientListByResourceGroupOptions contains the optional parameters for the VirtualNetworksClient.NewListByResourceGroupPager
//     method.
func (client *VirtualNetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *VirtualNetworksClientListByResourceGroupOptions) *runtime.Pager[VirtualNetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListByResourceGroupResponse]{
		More: func(page VirtualNetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListByResourceGroupResponse) (VirtualNetworksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VirtualNetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VirtualNetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualNetworks"
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
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VirtualNetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (VirtualNetworksClientListByResourceGroupResponse, error) {
	result := VirtualNetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResult); err != nil {
		return VirtualNetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List of VirtualNetworks in a subscription.
//
// Generated from API version 2023-04-01-preview
//   - options - VirtualNetworksClientListBySubscriptionOptions contains the optional parameters for the VirtualNetworksClient.NewListBySubscriptionPager
//     method.
func (client *VirtualNetworksClient) NewListBySubscriptionPager(options *VirtualNetworksClientListBySubscriptionOptions) *runtime.Pager[VirtualNetworksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VirtualNetworksClientListBySubscriptionResponse]{
		More: func(page VirtualNetworksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualNetworksClientListBySubscriptionResponse) (VirtualNetworksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualNetworksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VirtualNetworksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualNetworksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *VirtualNetworksClient) listBySubscriptionCreateRequest(ctx context.Context, options *VirtualNetworksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ScVmm/virtualNetworks"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *VirtualNetworksClient) listBySubscriptionHandleResponse(resp *http.Response) (VirtualNetworksClientListBySubscriptionResponse, error) {
	result := VirtualNetworksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VirtualNetworkListResult); err != nil {
		return VirtualNetworksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the VirtualNetworks resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - resourceGroupName - The name of the resource group.
//   - virtualNetworkName - Name of the VirtualNetwork.
//   - body - VirtualNetworks patch payload.
//   - options - VirtualNetworksClientBeginUpdateOptions contains the optional parameters for the VirtualNetworksClient.BeginUpdate
//     method.
func (client *VirtualNetworksClient) BeginUpdate(ctx context.Context, resourceGroupName string, virtualNetworkName string, body ResourcePatch, options *VirtualNetworksClientBeginUpdateOptions) (*runtime.Poller[VirtualNetworksClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, virtualNetworkName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VirtualNetworksClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[VirtualNetworksClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Updates the VirtualNetworks resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
func (client *VirtualNetworksClient) update(ctx context.Context, resourceGroupName string, virtualNetworkName string, body ResourcePatch, options *VirtualNetworksClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, virtualNetworkName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *VirtualNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, virtualNetworkName string, body ResourcePatch, options *VirtualNetworksClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ScVmm/virtualNetworks/{virtualNetworkName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if virtualNetworkName == "" {
		return nil, errors.New("parameter virtualNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualNetworkName}", url.PathEscape(virtualNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
