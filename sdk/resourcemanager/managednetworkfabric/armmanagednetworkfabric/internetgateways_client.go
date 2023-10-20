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

// InternetGatewaysClient contains the methods for the InternetGateways group.
// Don't use this type directly, use NewInternetGatewaysClient() instead.
type InternetGatewaysClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewInternetGatewaysClient creates a new instance of InternetGatewaysClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewInternetGatewaysClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*InternetGatewaysClient, error) {
	cl, err := arm.NewClient(moduleName+".InternetGatewaysClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &InternetGatewaysClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Creates a Network Fabric Service Internet Gateway resource instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayName - Name of the Internet Gateway.
//   - body - Request payload.
//   - options - InternetGatewaysClientBeginCreateOptions contains the optional parameters for the InternetGatewaysClient.BeginCreate
//     method.
func (client *InternetGatewaysClient) BeginCreate(ctx context.Context, resourceGroupName string, internetGatewayName string, body InternetGateway, options *InternetGatewaysClientBeginCreateOptions) (*runtime.Poller[InternetGatewaysClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, internetGatewayName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InternetGatewaysClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InternetGatewaysClientCreateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Create - Creates a Network Fabric Service Internet Gateway resource instance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *InternetGatewaysClient) create(ctx context.Context, resourceGroupName string, internetGatewayName string, body InternetGateway, options *InternetGatewaysClientBeginCreateOptions) (*http.Response, error) {
	var err error
	req, err := client.createCreateRequest(ctx, resourceGroupName, internetGatewayName, body, options)
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
func (client *InternetGatewaysClient) createCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayName string, body InternetGateway, options *InternetGatewaysClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGateways/{internetGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayName == "" {
		return nil, errors.New("parameter internetGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayName}", url.PathEscape(internetGatewayName))
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

// BeginDelete - Execute a delete on Network Fabric Service Internet Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayName - Name of the Internet Gateway.
//   - options - InternetGatewaysClientBeginDeleteOptions contains the optional parameters for the InternetGatewaysClient.BeginDelete
//     method.
func (client *InternetGatewaysClient) BeginDelete(ctx context.Context, resourceGroupName string, internetGatewayName string, options *InternetGatewaysClientBeginDeleteOptions) (*runtime.Poller[InternetGatewaysClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, internetGatewayName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InternetGatewaysClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InternetGatewaysClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Execute a delete on Network Fabric Service Internet Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *InternetGatewaysClient) deleteOperation(ctx context.Context, resourceGroupName string, internetGatewayName string, options *InternetGatewaysClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, internetGatewayName, options)
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
func (client *InternetGatewaysClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayName string, options *InternetGatewaysClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGateways/{internetGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayName == "" {
		return nil, errors.New("parameter internetGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayName}", url.PathEscape(internetGatewayName))
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

// Get - Implements Gateway GET method.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayName - Name of the Internet Gateway.
//   - options - InternetGatewaysClientGetOptions contains the optional parameters for the InternetGatewaysClient.Get method.
func (client *InternetGatewaysClient) Get(ctx context.Context, resourceGroupName string, internetGatewayName string, options *InternetGatewaysClientGetOptions) (InternetGatewaysClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, internetGatewayName, options)
	if err != nil {
		return InternetGatewaysClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return InternetGatewaysClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return InternetGatewaysClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *InternetGatewaysClient) getCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayName string, options *InternetGatewaysClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGateways/{internetGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayName == "" {
		return nil, errors.New("parameter internetGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayName}", url.PathEscape(internetGatewayName))
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
func (client *InternetGatewaysClient) getHandleResponse(resp *http.Response) (InternetGatewaysClientGetResponse, error) {
	result := InternetGatewaysClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternetGateway); err != nil {
		return InternetGatewaysClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Displays Internet Gateways list by resource group GET method.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - InternetGatewaysClientListByResourceGroupOptions contains the optional parameters for the InternetGatewaysClient.NewListByResourceGroupPager
//     method.
func (client *InternetGatewaysClient) NewListByResourceGroupPager(resourceGroupName string, options *InternetGatewaysClientListByResourceGroupOptions) *runtime.Pager[InternetGatewaysClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[InternetGatewaysClientListByResourceGroupResponse]{
		More: func(page InternetGatewaysClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InternetGatewaysClientListByResourceGroupResponse) (InternetGatewaysClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InternetGatewaysClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InternetGatewaysClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InternetGatewaysClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *InternetGatewaysClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *InternetGatewaysClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGateways"
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
func (client *InternetGatewaysClient) listByResourceGroupHandleResponse(resp *http.Response) (InternetGatewaysClientListByResourceGroupResponse, error) {
	result := InternetGatewaysClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternetGatewaysListResult); err != nil {
		return InternetGatewaysClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Displays Internet Gateways list by subscription GET method.
//
// Generated from API version 2023-06-15
//   - options - InternetGatewaysClientListBySubscriptionOptions contains the optional parameters for the InternetGatewaysClient.NewListBySubscriptionPager
//     method.
func (client *InternetGatewaysClient) NewListBySubscriptionPager(options *InternetGatewaysClientListBySubscriptionOptions) *runtime.Pager[InternetGatewaysClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[InternetGatewaysClientListBySubscriptionResponse]{
		More: func(page InternetGatewaysClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *InternetGatewaysClientListBySubscriptionResponse) (InternetGatewaysClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return InternetGatewaysClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return InternetGatewaysClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return InternetGatewaysClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *InternetGatewaysClient) listBySubscriptionCreateRequest(ctx context.Context, options *InternetGatewaysClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ManagedNetworkFabric/internetGateways"
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
func (client *InternetGatewaysClient) listBySubscriptionHandleResponse(resp *http.Response) (InternetGatewaysClientListBySubscriptionResponse, error) {
	result := InternetGatewaysClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InternetGatewaysListResult); err != nil {
		return InternetGatewaysClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Execute patch on Network Fabric Service Internet Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - internetGatewayName - Name of the Internet Gateway.
//   - body - API to update certain properties of the L2 Isolation Domain resource..
//   - options - InternetGatewaysClientBeginUpdateOptions contains the optional parameters for the InternetGatewaysClient.BeginUpdate
//     method.
func (client *InternetGatewaysClient) BeginUpdate(ctx context.Context, resourceGroupName string, internetGatewayName string, body InternetGatewayPatch, options *InternetGatewaysClientBeginUpdateOptions) (*runtime.Poller[InternetGatewaysClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, internetGatewayName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[InternetGatewaysClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[InternetGatewaysClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Execute patch on Network Fabric Service Internet Gateway.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-15
func (client *InternetGatewaysClient) update(ctx context.Context, resourceGroupName string, internetGatewayName string, body InternetGatewayPatch, options *InternetGatewaysClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, internetGatewayName, body, options)
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
func (client *InternetGatewaysClient) updateCreateRequest(ctx context.Context, resourceGroupName string, internetGatewayName string, body InternetGatewayPatch, options *InternetGatewaysClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedNetworkFabric/internetGateways/{internetGatewayName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if internetGatewayName == "" {
		return nil, errors.New("parameter internetGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{internetGatewayName}", url.PathEscape(internetGatewayName))
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
