//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmanagednetworkfabric

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

// NetworkPacketBrokersClient contains the methods for the NetworkPacketBrokers group.
// Don't use this type directly, use NewNetworkPacketBrokersClient() instead.
type NetworkPacketBrokersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkPacketBrokersClient creates a new instance of NetworkPacketBrokersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkPacketBrokersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkPacketBrokersClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkPacketBrokersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkPacketBrokersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a Network Packet Broker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkPacketBrokerName - Name of the Network Packet Broker.
//   - body - Request payload.
//   - options - NetworkPacketBrokersClientBeginCreateOptions contains the optional parameters for the NetworkPacketBrokersClient.BeginCreate
//     method.
func (client *NetworkPacketBrokersClient) BeginCreate(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, body NetworkPacketBroker, options *NetworkPacketBrokersClientBeginCreateOptions) (*runtime.Poller[NetworkPacketBrokersClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, networkPacketBrokerName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkPacketBrokersClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[NetworkPacketBrokersClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a Network Packet Broker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkPacketBrokersClient) create(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, body NetworkPacketBroker, options *NetworkPacketBrokersClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkPacketBrokerName, body, options)
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

// createCreateRequest creates the Create request.
func (client *NetworkPacketBrokersClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, body NetworkPacketBroker, options *NetworkPacketBrokersClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/{networkPacketBrokerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkPacketBrokerName == "" {
		return nil, errors.New("parameter networkPacketBrokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkPacketBrokerName}", url.PathEscape(networkPacketBrokerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes Network Packet Broker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkPacketBrokerName - Name of the Network Packet Broker.
//   - options - NetworkPacketBrokersClientBeginDeleteOptions contains the optional parameters for the NetworkPacketBrokersClient.BeginDelete
//     method.
func (client *NetworkPacketBrokersClient) BeginDelete(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, options *NetworkPacketBrokersClientBeginDeleteOptions) (*runtime.Poller[NetworkPacketBrokersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkPacketBrokerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkPacketBrokersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[NetworkPacketBrokersClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes Network Packet Broker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkPacketBrokersClient) deleteOperation(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, options *NetworkPacketBrokersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkPacketBrokerName, options)
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
func (client *NetworkPacketBrokersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, options *NetworkPacketBrokersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/{networkPacketBrokerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkPacketBrokerName == "" {
		return nil, errors.New("parameter networkPacketBrokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkPacketBrokerName}", url.PathEscape(networkPacketBrokerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieves details of this Network Packet Broker.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkPacketBrokerName - Name of the Network Packet Broker.
//   - options - NetworkPacketBrokersClientGetOptions contains the optional parameters for the NetworkPacketBrokersClient.Get
//     method.
func (client *NetworkPacketBrokersClient) Get(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, options *NetworkPacketBrokersClientGetOptions) (NetworkPacketBrokersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkPacketBrokerName, options)
	if err != nil {
		return NetworkPacketBrokersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkPacketBrokersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkPacketBrokersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *NetworkPacketBrokersClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, options *NetworkPacketBrokersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/{networkPacketBrokerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkPacketBrokerName == "" {
		return nil, errors.New("parameter networkPacketBrokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkPacketBrokerName}", url.PathEscape(networkPacketBrokerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkPacketBrokersClient) getHandleResponse(resp *http.Response) (NetworkPacketBrokersClientGetResponse, error) {
	result := NetworkPacketBrokersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkPacketBroker); err != nil {
		return NetworkPacketBrokersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Displays NetworkPacketBrokers list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - NetworkPacketBrokersClientListByResourceGroupOptions contains the optional parameters for the NetworkPacketBrokersClient.NewListByResourceGroupPager
//     method.
func (client *NetworkPacketBrokersClient) NewListByResourceGroupPager(resourceGroupName string, options *NetworkPacketBrokersClientListByResourceGroupOptions) *runtime.Pager[NetworkPacketBrokersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkPacketBrokersClientListByResourceGroupResponse]{
		More: func(page NetworkPacketBrokersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkPacketBrokersClientListByResourceGroupResponse) (NetworkPacketBrokersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkPacketBrokersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkPacketBrokersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkPacketBrokersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *NetworkPacketBrokersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *NetworkPacketBrokersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers"
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
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *NetworkPacketBrokersClient) listByResourceGroupHandleResponse(resp *http.Response) (NetworkPacketBrokersClientListByResourceGroupResponse, error) {
	result := NetworkPacketBrokersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkPacketBrokersListResult); err != nil {
		return NetworkPacketBrokersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Displays Network Packet Brokers list by subscription GET method.
//
// Generated from API version 2023-06-15
//   - options - NetworkPacketBrokersClientListBySubscriptionOptions contains the optional parameters for the NetworkPacketBrokersClient.NewListBySubscriptionPager
//     method.
func (client *NetworkPacketBrokersClient) NewListBySubscriptionPager(options *NetworkPacketBrokersClientListBySubscriptionOptions) *runtime.Pager[NetworkPacketBrokersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[NetworkPacketBrokersClientListBySubscriptionResponse]{
		More: func(page NetworkPacketBrokersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *NetworkPacketBrokersClientListBySubscriptionResponse) (NetworkPacketBrokersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return NetworkPacketBrokersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return NetworkPacketBrokersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return NetworkPacketBrokersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *NetworkPacketBrokersClient) listBySubscriptionCreateRequest(ctx context.Context, options *NetworkPacketBrokersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *NetworkPacketBrokersClient) listBySubscriptionHandleResponse(resp *http.Response) (NetworkPacketBrokersClientListBySubscriptionResponse, error) {
	result := NetworkPacketBrokersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkPacketBrokersListResult); err != nil {
		return NetworkPacketBrokersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - API to update certain properties of the Network Packet Broker resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - networkPacketBrokerName - Name of the Network Packet Broker.
//   - body - Network Packet Broker properties to update.
//   - options - NetworkPacketBrokersClientBeginUpdateOptions contains the optional parameters for the NetworkPacketBrokersClient.BeginUpdate
//     method.
func (client *NetworkPacketBrokersClient) BeginUpdate(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, body NetworkPacketBrokerPatch, options *NetworkPacketBrokersClientBeginUpdateOptions) (*runtime.Poller[NetworkPacketBrokersClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, networkPacketBrokerName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[NetworkPacketBrokersClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[NetworkPacketBrokersClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - API to update certain properties of the Network Packet Broker resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *NetworkPacketBrokersClient) update(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, body NetworkPacketBrokerPatch, options *NetworkPacketBrokersClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, networkPacketBrokerName, body, options)
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
func (client *NetworkPacketBrokersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, networkPacketBrokerName string, body NetworkPacketBrokerPatch, options *NetworkPacketBrokersClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/networkPacketBrokers/{networkPacketBrokerName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkPacketBrokerName == "" {
		return nil, errors.New("parameter networkPacketBrokerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkPacketBrokerName}", url.PathEscape(networkPacketBrokerName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}
