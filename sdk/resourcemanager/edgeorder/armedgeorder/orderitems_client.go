//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armedgeorder

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

// OrderItemsClient contains the methods for the OrderItems group.
// Don't use this type directly, use NewOrderItemsClient() instead.
type OrderItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOrderItemsClient creates a new instance of OrderItemsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrderItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OrderItemsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OrderItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Cancel - Cancel order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - orderItemName - The name of the order item.
//   - cancellationReason - Reason for cancellation.
//   - options - OrderItemsClientCancelOptions contains the optional parameters for the OrderItemsClient.Cancel method.
func (client *OrderItemsClient) Cancel(ctx context.Context, resourceGroupName string, orderItemName string, cancellationReason CancellationReason, options *OrderItemsClientCancelOptions) (OrderItemsClientCancelResponse, error) {
	var err error
	const operationName = "OrderItemsClient.Cancel"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.cancelCreateRequest(ctx, resourceGroupName, orderItemName, cancellationReason, options)
	if err != nil {
		return OrderItemsClientCancelResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrderItemsClientCancelResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OrderItemsClientCancelResponse{}, err
	}
	return OrderItemsClientCancelResponse{}, nil
}

// cancelCreateRequest creates the Cancel request.
func (client *OrderItemsClient) cancelCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, cancellationReason CancellationReason, options *OrderItemsClientCancelOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}/cancel"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, cancellationReason); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginCreate - Create an order item. Existing order item cannot be updated with this api and should instead be updated with
// the Update order item API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - orderItemName - The name of the order item.
//   - orderItemResource - Order item details from request body.
//   - options - OrderItemsClientBeginCreateOptions contains the optional parameters for the OrderItemsClient.BeginCreate method.
func (client *OrderItemsClient) BeginCreate(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource OrderItemResource, options *OrderItemsClientBeginCreateOptions) (*runtime.Poller[OrderItemsClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, orderItemName, orderItemResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrderItemsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrderItemsClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create an order item. Existing order item cannot be updated with this api and should instead be updated with the
// Update order item API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OrderItemsClient) create(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource OrderItemResource, options *OrderItemsClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "OrderItemsClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, orderItemName, orderItemResource, options)
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

// createCreateRequest creates the Create request.
func (client *OrderItemsClient) createCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, orderItemResource OrderItemResource, options *OrderItemsClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, orderItemResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - orderItemName - The name of the order item.
//   - options - OrderItemsClientBeginDeleteOptions contains the optional parameters for the OrderItemsClient.BeginDelete method.
func (client *OrderItemsClient) BeginDelete(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientBeginDeleteOptions) (*runtime.Poller[OrderItemsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, orderItemName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrderItemsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrderItemsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete an order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OrderItemsClient) deleteOperation(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "OrderItemsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, orderItemName, options)
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
func (client *OrderItemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - orderItemName - The name of the order item.
//   - options - OrderItemsClientGetOptions contains the optional parameters for the OrderItemsClient.Get method.
func (client *OrderItemsClient) Get(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientGetOptions) (OrderItemsClientGetResponse, error) {
	var err error
	const operationName = "OrderItemsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, orderItemName, options)
	if err != nil {
		return OrderItemsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OrderItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OrderItemsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OrderItemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, options *OrderItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OrderItemsClient) getHandleResponse(resp *http.Response) (OrderItemsClientGetResponse, error) {
	result := OrderItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderItemResource); err != nil {
		return OrderItemsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List order items at resource group level.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - OrderItemsClientListByResourceGroupOptions contains the optional parameters for the OrderItemsClient.NewListByResourceGroupPager
//     method.
func (client *OrderItemsClient) NewListByResourceGroupPager(resourceGroupName string, options *OrderItemsClientListByResourceGroupOptions) *runtime.Pager[OrderItemsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrderItemsClientListByResourceGroupResponse]{
		More: func(page OrderItemsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrderItemsClientListByResourceGroupResponse) (OrderItemsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrderItemsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return OrderItemsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *OrderItemsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *OrderItemsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems"
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
	reqQP.Set("api-version", "2024-02-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
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
func (client *OrderItemsClient) listByResourceGroupHandleResponse(resp *http.Response) (OrderItemsClientListByResourceGroupResponse, error) {
	result := OrderItemsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderItemResourceList); err != nil {
		return OrderItemsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - List order items at subscription level.
//
// Generated from API version 2024-02-01
//   - options - OrderItemsClientListBySubscriptionOptions contains the optional parameters for the OrderItemsClient.NewListBySubscriptionPager
//     method.
func (client *OrderItemsClient) NewListBySubscriptionPager(options *OrderItemsClientListBySubscriptionOptions) *runtime.Pager[OrderItemsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrderItemsClientListBySubscriptionResponse]{
		More: func(page OrderItemsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrderItemsClientListBySubscriptionResponse) (OrderItemsClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrderItemsClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OrderItemsClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *OrderItemsClient) listBySubscriptionCreateRequest(ctx context.Context, options *OrderItemsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.EdgeOrder/orderItems"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
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
func (client *OrderItemsClient) listBySubscriptionHandleResponse(resp *http.Response) (OrderItemsClientListBySubscriptionResponse, error) {
	result := OrderItemsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OrderItemResourceList); err != nil {
		return OrderItemsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginReturn - Return order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - orderItemName - The name of the order item.
//   - returnOrderItemDetails - Return order item details.
//   - options - OrderItemsClientBeginReturnOptions contains the optional parameters for the OrderItemsClient.BeginReturn method.
func (client *OrderItemsClient) BeginReturn(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails ReturnOrderItemDetails, options *OrderItemsClientBeginReturnOptions) (*runtime.Poller[OrderItemsClientReturnResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.returnOperation(ctx, resourceGroupName, orderItemName, returnOrderItemDetails, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrderItemsClientReturnResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrderItemsClientReturnResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Return - Return order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OrderItemsClient) returnOperation(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails ReturnOrderItemDetails, options *OrderItemsClientBeginReturnOptions) (*http.Response, error) {
	var err error
	const operationName = "OrderItemsClient.BeginReturn"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.returnCreateRequest(ctx, resourceGroupName, orderItemName, returnOrderItemDetails, options)
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

// returnCreateRequest creates the Return request.
func (client *OrderItemsClient) returnCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, returnOrderItemDetails ReturnOrderItemDetails, options *OrderItemsClientBeginReturnOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}/return"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, returnOrderItemDetails); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdate - Update the properties of an existing order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - orderItemName - The name of the order item.
//   - orderItemUpdateParameter - Order item update parameters from request body.
//   - options - OrderItemsClientBeginUpdateOptions contains the optional parameters for the OrderItemsClient.BeginUpdate method.
func (client *OrderItemsClient) BeginUpdate(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter OrderItemUpdateParameter, options *OrderItemsClientBeginUpdateOptions) (*runtime.Poller[OrderItemsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, orderItemName, orderItemUpdateParameter, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[OrderItemsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[OrderItemsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update the properties of an existing order item.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-02-01
func (client *OrderItemsClient) update(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter OrderItemUpdateParameter, options *OrderItemsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "OrderItemsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, orderItemName, orderItemUpdateParameter, options)
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
func (client *OrderItemsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, orderItemName string, orderItemUpdateParameter OrderItemUpdateParameter, options *OrderItemsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EdgeOrder/orderItems/{orderItemName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if orderItemName == "" {
		return nil, errors.New("parameter orderItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{orderItemName}", url.PathEscape(orderItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, orderItemUpdateParameter); err != nil {
		return nil, err
	}
	return req, nil
}
