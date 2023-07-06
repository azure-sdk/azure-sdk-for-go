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

// L2NetworksClient contains the methods for the L2Networks group.
// Don't use this type directly, use NewL2NetworksClient() instead.
type L2NetworksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewL2NetworksClient creates a new instance of L2NetworksClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewL2NetworksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*L2NetworksClient, error) {
	cl, err := arm.NewClient(moduleName+".L2NetworksClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &L2NetworksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create a new layer 2 (L2) network or update the properties of the existing network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2NetworkName - The name of the L2 network.
//   - l2NetworkParameters - The request body.
//   - options - L2NetworksClientBeginCreateOrUpdateOptions contains the optional parameters for the L2NetworksClient.BeginCreateOrUpdate
//     method.
func (client *L2NetworksClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, l2NetworkName string, l2NetworkParameters L2Network, options *L2NetworksClientBeginCreateOrUpdateOptions) (*runtime.Poller[L2NetworksClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, l2NetworkName, l2NetworkParameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2NetworksClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[L2NetworksClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create a new layer 2 (L2) network or update the properties of the existing network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *L2NetworksClient) createOrUpdate(ctx context.Context, resourceGroupName string, l2NetworkName string, l2NetworkParameters L2Network, options *L2NetworksClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, l2NetworkName, l2NetworkParameters, options)
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
func (client *L2NetworksClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, l2NetworkName string, l2NetworkParameters L2Network, options *L2NetworksClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l2Networks/{l2NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2NetworkName == "" {
		return nil, errors.New("parameter l2NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2NetworkName}", url.PathEscape(l2NetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, l2NetworkParameters)
}

// BeginDelete - Delete the provided layer 2 (L2) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2NetworkName - The name of the L2 network.
//   - options - L2NetworksClientBeginDeleteOptions contains the optional parameters for the L2NetworksClient.BeginDelete method.
func (client *L2NetworksClient) BeginDelete(ctx context.Context, resourceGroupName string, l2NetworkName string, options *L2NetworksClientBeginDeleteOptions) (*runtime.Poller[L2NetworksClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, l2NetworkName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[L2NetworksClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[L2NetworksClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete the provided layer 2 (L2) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
func (client *L2NetworksClient) deleteOperation(ctx context.Context, resourceGroupName string, l2NetworkName string, options *L2NetworksClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, l2NetworkName, options)
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
func (client *L2NetworksClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, l2NetworkName string, options *L2NetworksClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l2Networks/{l2NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2NetworkName == "" {
		return nil, errors.New("parameter l2NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2NetworkName}", url.PathEscape(l2NetworkName))
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

// Get - Get properties of the provided layer 2 (L2) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2NetworkName - The name of the L2 network.
//   - options - L2NetworksClientGetOptions contains the optional parameters for the L2NetworksClient.Get method.
func (client *L2NetworksClient) Get(ctx context.Context, resourceGroupName string, l2NetworkName string, options *L2NetworksClientGetOptions) (L2NetworksClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, l2NetworkName, options)
	if err != nil {
		return L2NetworksClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return L2NetworksClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return L2NetworksClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *L2NetworksClient) getCreateRequest(ctx context.Context, resourceGroupName string, l2NetworkName string, options *L2NetworksClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l2Networks/{l2NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2NetworkName == "" {
		return nil, errors.New("parameter l2NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2NetworkName}", url.PathEscape(l2NetworkName))
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
func (client *L2NetworksClient) getHandleResponse(resp *http.Response) (L2NetworksClientGetResponse, error) {
	result := L2NetworksClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2Network); err != nil {
		return L2NetworksClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Get a list of layer 2 (L2) networks in the provided resource group.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - L2NetworksClientListByResourceGroupOptions contains the optional parameters for the L2NetworksClient.NewListByResourceGroupPager
//     method.
func (client *L2NetworksClient) NewListByResourceGroupPager(resourceGroupName string, options *L2NetworksClientListByResourceGroupOptions) *runtime.Pager[L2NetworksClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[L2NetworksClientListByResourceGroupResponse]{
		More: func(page L2NetworksClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L2NetworksClientListByResourceGroupResponse) (L2NetworksClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return L2NetworksClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return L2NetworksClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return L2NetworksClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *L2NetworksClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *L2NetworksClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l2Networks"
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
func (client *L2NetworksClient) listByResourceGroupHandleResponse(resp *http.Response) (L2NetworksClientListByResourceGroupResponse, error) {
	result := L2NetworksClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2NetworkList); err != nil {
		return L2NetworksClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Get a list of layer 2 (L2) networks in the provided subscription.
//
// Generated from API version 2023-07-01
//   - options - L2NetworksClientListBySubscriptionOptions contains the optional parameters for the L2NetworksClient.NewListBySubscriptionPager
//     method.
func (client *L2NetworksClient) NewListBySubscriptionPager(options *L2NetworksClientListBySubscriptionOptions) *runtime.Pager[L2NetworksClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[L2NetworksClientListBySubscriptionResponse]{
		More: func(page L2NetworksClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *L2NetworksClientListBySubscriptionResponse) (L2NetworksClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return L2NetworksClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return L2NetworksClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return L2NetworksClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *L2NetworksClient) listBySubscriptionCreateRequest(ctx context.Context, options *L2NetworksClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.NetworkCloud/l2Networks"
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
func (client *L2NetworksClient) listBySubscriptionHandleResponse(resp *http.Response) (L2NetworksClientListBySubscriptionResponse, error) {
	result := L2NetworksClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2NetworkList); err != nil {
		return L2NetworksClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// Update - Update tags associated with the provided layer 2 (L2) network.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - l2NetworkName - The name of the L2 network.
//   - l2NetworkUpdateParameters - The request body.
//   - options - L2NetworksClientUpdateOptions contains the optional parameters for the L2NetworksClient.Update method.
func (client *L2NetworksClient) Update(ctx context.Context, resourceGroupName string, l2NetworkName string, l2NetworkUpdateParameters L2NetworkPatchParameters, options *L2NetworksClientUpdateOptions) (L2NetworksClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, l2NetworkName, l2NetworkUpdateParameters, options)
	if err != nil {
		return L2NetworksClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return L2NetworksClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return L2NetworksClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *L2NetworksClient) updateCreateRequest(ctx context.Context, resourceGroupName string, l2NetworkName string, l2NetworkUpdateParameters L2NetworkPatchParameters, options *L2NetworksClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/l2Networks/{l2NetworkName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if l2NetworkName == "" {
		return nil, errors.New("parameter l2NetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{l2NetworkName}", url.PathEscape(l2NetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, l2NetworkUpdateParameters)
}

// updateHandleResponse handles the Update response.
func (client *L2NetworksClient) updateHandleResponse(resp *http.Response) (L2NetworksClientUpdateResponse, error) {
	result := L2NetworksClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.L2Network); err != nil {
		return L2NetworksClientUpdateResponse{}, err
	}
	return result, nil
}
