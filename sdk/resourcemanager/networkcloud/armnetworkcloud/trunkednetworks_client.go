//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// TrunkedNetworksClient contains the methods for the TrunkedNetworks group.
// Don't use this type directly, use NewTrunkedNetworksClient() instead.
type TrunkedNetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTrunkedNetworksClient creates a new instance of TrunkedNetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTrunkedNetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TrunkedNetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".TrunkedNetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TrunkedNetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new trunked network or update the properties of the existing trunked network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trunkedNetworkName - The name of the trunked network.
//   - trunkedNetworkParameters - The request body.
//   - options - TrunkedNetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the TrunkedNetworksClient.BeginCreateOrUpdate
//     method.
func (client *TrunkedNetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, trunkedNetworkName string, trunkedNetworkParameters TrunkedNetwork, options *TrunkedNetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[TrunkedNetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, trunkedNetworkName, trunkedNetworkParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TrunkedNetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TrunkedNetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new trunked network or update the properties of the existing trunked network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *TrunkedNetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, trunkedNetworkName string, trunkedNetworkParameters TrunkedNetwork, options *TrunkedNetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, trunkedNetworkName, trunkedNetworkParameters, options)
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
func (client *TrunkedNetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, trunkedNetworkName string, trunkedNetworkParameters TrunkedNetwork, options *TrunkedNetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/trunkedNetworks/{trunkedNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trunkedNetworkName == "" {
		return nil, errors.New("parameter trunkedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trunkedNetworkName}", url.PathEscape(trunkedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, trunkedNetworkParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the provided trunked network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trunkedNetworkName - The name of the trunked network.
//   - options - TrunkedNetworksClientBeginDeleteOptions contains the optional parameters for the TrunkedNetworksClient.BeginDelete
//     method.
func (client *TrunkedNetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, trunkedNetworkName string, options *TrunkedNetworksClientBeginDeleteOptions) (*runtime.Poller[TrunkedNetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, trunkedNetworkName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[TrunkedNetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[TrunkedNetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided trunked network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *TrunkedNetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, trunkedNetworkName string, options *TrunkedNetworksClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, trunkedNetworkName, options)
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
func (client *TrunkedNetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, trunkedNetworkName string, options *TrunkedNetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/trunkedNetworks/{trunkedNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trunkedNetworkName == "" {
		return nil, errors.New("parameter trunkedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trunkedNetworkName}", url.PathEscape(trunkedNetworkName))
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

// Get - Get properties of the provided trunked network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trunkedNetworkName - The name of the trunked network.
//   - options - TrunkedNetworksClientGetOptions contains the optional parameters for the TrunkedNetworksClient.Get method.
func (client *TrunkedNetworksClient) Get(ctx context.Context, resourceGroupName string, trunkedNetworkName string, options *TrunkedNetworksClientGetOptions) (TrunkedNetworksClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, trunkedNetworkName, options)
	if err != nil {
		return TrunkedNetworksClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrunkedNetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TrunkedNetworksClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TrunkedNetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, trunkedNetworkName string, options *TrunkedNetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/trunkedNetworks/{trunkedNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trunkedNetworkName == "" {
		return nil, errors.New("parameter trunkedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trunkedNetworkName}", url.PathEscape(trunkedNetworkName))
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
func (client *TrunkedNetworksClient) getHandleResponse(resp *http.Response) (TrunkedNetworksClientGetResponse, error) {
	result := TrunkedNetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrunkedNetwork); err != nil {
		return TrunkedNetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of trunked networks in the provided resource group.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - TrunkedNetworksClientListByResourceGroupOptions contains the optional parameters for the TrunkedNetworksClient.NewListByResourceGroupPager
//     method.
func (client *TrunkedNetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *TrunkedNetworksClientListByResourceGroupOptions) *runtime.Pager[TrunkedNetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrunkedNetworksClientListByResourceGroupResponse]{
		More: func(page TrunkedNetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrunkedNetworksClientListByResourceGroupResponse) (TrunkedNetworksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TrunkedNetworksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TrunkedNetworksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TrunkedNetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *TrunkedNetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *TrunkedNetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/trunkedNetworks"
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
func (client *TrunkedNetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (TrunkedNetworksClientListByResourceGroupResponse, error) {
	result := TrunkedNetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrunkedNetworkList); err != nil {
		return TrunkedNetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of trunked networks in the provided subscription.
//
// Generated from API version 2023-07-01
//   - options - TrunkedNetworksClientListBySubscriptionOptions contains the optional parameters for the TrunkedNetworksClient.NewListBySubscriptionPager
//     method.
func (client *TrunkedNetworksClient) NewListBySubscriptionPager(options *TrunkedNetworksClientListBySubscriptionOptions) *runtime.Pager[TrunkedNetworksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[TrunkedNetworksClientListBySubscriptionResponse]{
		More: func(page TrunkedNetworksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TrunkedNetworksClientListBySubscriptionResponse) (TrunkedNetworksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TrunkedNetworksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TrunkedNetworksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TrunkedNetworksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *TrunkedNetworksClient) listBySubscriptionCreateRequest(ctx context.Context, options *TrunkedNetworksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/trunkedNetworks"
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
func (client *TrunkedNetworksClient) listBySubscriptionHandleResponse(resp *http.Response) (TrunkedNetworksClientListBySubscriptionResponse, error) {
	result := TrunkedNetworksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrunkedNetworkList); err != nil {
		return TrunkedNetworksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update tags associated with the provided trunked network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - trunkedNetworkName - The name of the trunked network.
//   - trunkedNetworkUpdateParameters - The request body.
//   - options - TrunkedNetworksClientUpdateOptions contains the optional parameters for the TrunkedNetworksClient.Update method.
func (client *TrunkedNetworksClient) Update(ctx context.Context, resourceGroupName string, trunkedNetworkName string, trunkedNetworkUpdateParameters TrunkedNetworkPatchParameters, options *TrunkedNetworksClientUpdateOptions) (TrunkedNetworksClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, trunkedNetworkName, trunkedNetworkUpdateParameters, options)
	if err != nil {
		return TrunkedNetworksClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TrunkedNetworksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TrunkedNetworksClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *TrunkedNetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, trunkedNetworkName string, trunkedNetworkUpdateParameters TrunkedNetworkPatchParameters, options *TrunkedNetworksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/trunkedNetworks/{trunkedNetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if trunkedNetworkName == "" {
		return nil, errors.New("parameter trunkedNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{trunkedNetworkName}", url.PathEscape(trunkedNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, trunkedNetworkUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *TrunkedNetworksClient) updateHandleResponse(resp *http.Response) (TrunkedNetworksClientUpdateResponse, error) {
	result := TrunkedNetworksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TrunkedNetwork); err != nil {
		return TrunkedNetworksClientUpdateResponse{}, err
	}
	return result, nil
}
