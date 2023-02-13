//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armedgeorder

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
	"strconv"
	"strings"
)

// OrdersClient contains the methods for the Orders group.
// Don't use this type directly, use NewOrdersClient() instead.
type OrdersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewOrdersClient creates a new instance of OrdersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrdersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrdersClient, error) {
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
	client := &OrdersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get an order.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - location - The name of Azure region.
//   - orderName - The name of the order.
//   - options - OrdersClientGetOptions contains the optional parameters for the OrdersClient.Get method.
func (client *OrdersClient) Get(ctx context.Context, resourceGroupName string, location string, orderName string, options *OrdersClientGetOptions) (OrdersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, location, orderName, options)
	if err != nil {
		return OrdersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return OrdersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return OrdersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *OrdersClient) getCreateRequest(ctx context.Context, resourceGroupName string, location string, orderName string, options *OrdersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/locations/{location}/orders/{orderName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if orderName == "" {
		return nil, errors.New("parameter orderName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderName}", url.PathEscape(orderName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrdersClient) getHandleResponse(resp *http.Response) (OrdersClientGetResponse, error) {
	result := OrdersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderResource); err != nil {
		return OrdersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List orders at resource group level.
//
// Generated from API version 2022-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - OrdersClientListByResourceGroupOptions contains the optional parameters for the OrdersClient.NewListByResourceGroupPager
//     method.
func (client *OrdersClient) NewListByResourceGroupPager(resourceGroupName string, options *OrdersClientListByResourceGroupOptions) *runtime.Pager[OrdersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrdersClientListByResourceGroupResponse]{
		More: func(page OrdersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrdersClientListByResourceGroupResponse) (OrdersClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrdersClientListByResourceGroupResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrdersClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrdersClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OrdersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OrdersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *OrdersClient) listByResourceGroupHandleResponse(resp *http.Response) (OrdersClientListByResourceGroupResponse, error) {
	result := OrdersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderResourceList); err != nil {
		return OrdersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List orders at subscription level.
//
// Generated from API version 2022-05-01-preview
//   - options - OrdersClientListBySubscriptionOptions contains the optional parameters for the OrdersClient.NewListBySubscriptionPager
//     method.
func (client *OrdersClient) NewListBySubscriptionPager(options *OrdersClientListBySubscriptionOptions) *runtime.Pager[OrdersClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrdersClientListBySubscriptionResponse]{
		More: func(page OrdersClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrdersClientListBySubscriptionResponse) (OrdersClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OrdersClientListBySubscriptionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return OrdersClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OrdersClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OrdersClient) listBySubscriptionCreateRequest(ctx context.Context, options *OrdersClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/orders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *OrdersClient) listBySubscriptionHandleResponse(resp *http.Response) (OrdersClientListBySubscriptionResponse, error) {
	result := OrdersClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderResourceList); err != nil {
		return OrdersClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
