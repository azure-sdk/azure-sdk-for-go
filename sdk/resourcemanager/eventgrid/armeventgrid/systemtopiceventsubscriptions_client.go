//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventgrid

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

// SystemTopicEventSubscriptionsClient contains the methods for the SystemTopicEventSubscriptions group.
// Don't use this type directly, use NewSystemTopicEventSubscriptionsClient() instead.
type SystemTopicEventSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSystemTopicEventSubscriptionsClient creates a new instance of SystemTopicEventSubscriptionsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSystemTopicEventSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SystemTopicEventSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SystemTopicEventSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates or updates an event subscription with the specified parameters. Existing event
// subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 64
//     characters in length and use alphanumeric letters only.
//   - eventSubscriptionInfo - Event subscription properties containing the destination and filter information.
//   - options - SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.BeginCreateOrUpdate
//     method.
func (client *SystemTopicEventSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[SystemTopicEventSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SystemTopicEventSubscriptionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SystemTopicEventSubscriptionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Asynchronously creates or updates an event subscription with the specified parameters. Existing event
// subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
func (client *SystemTopicEventSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SystemTopicEventSubscriptionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SystemTopicEventSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *SystemTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSubscriptionInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - eventSubscriptionName - Name of the event subscription to be deleted.
//   - options - SystemTopicEventSubscriptionsClientBeginDeleteOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.BeginDelete
//     method.
func (client *SystemTopicEventSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[SystemTopicEventSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SystemTopicEventSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SystemTopicEventSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
func (client *SystemTopicEventSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "SystemTopicEventSubscriptionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
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
func (client *SystemTopicEventSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get an event subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - eventSubscriptionName - Name of the event subscription to be found.
//   - options - SystemTopicEventSubscriptionsClientGetOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.Get
//     method.
func (client *SystemTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetOptions) (SystemTopicEventSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "SystemTopicEventSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SystemTopicEventSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemTopicEventSubscriptionsClient) getHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientGetResponse, error) {
	result := SystemTopicEventSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscription); err != nil {
		return SystemTopicEventSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeliveryAttributes - Get all delivery attributes for an event subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - eventSubscriptionName - Name of the event subscription.
//   - options - SystemTopicEventSubscriptionsClientGetDeliveryAttributesOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.GetDeliveryAttributes
//     method.
func (client *SystemTopicEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	var err error
	const operationName = "SystemTopicEventSubscriptionsClient.GetDeliveryAttributes"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDeliveryAttributesCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	resp, err := client.getDeliveryAttributesHandleResponse(httpResp)
	return resp, err
}

// getDeliveryAttributesCreateRequest creates the GetDeliveryAttributes request.
func (client *SystemTopicEventSubscriptionsClient) getDeliveryAttributesCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeliveryAttributesHandleResponse handles the GetDeliveryAttributes response.
func (client *SystemTopicEventSubscriptionsClient) getDeliveryAttributesHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	result := SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeliveryAttributeListResult); err != nil {
		return SystemTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	return result, nil
}

// GetFullURL - Get the full endpoint URL for an event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - eventSubscriptionName - Name of the event subscription.
//   - options - SystemTopicEventSubscriptionsClientGetFullURLOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.GetFullURL
//     method.
func (client *SystemTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetFullURLOptions) (SystemTopicEventSubscriptionsClientGetFullURLResponse, error) {
	var err error
	const operationName = "SystemTopicEventSubscriptionsClient.GetFullURL"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, options)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	resp, err := client.getFullURLHandleResponse(httpResp)
	return resp, err
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *SystemTopicEventSubscriptionsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, options *SystemTopicEventSubscriptionsClientGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *SystemTopicEventSubscriptionsClient) getFullURLHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientGetFullURLResponse, error) {
	result := SystemTopicEventSubscriptionsClientGetFullURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionFullURL); err != nil {
		return SystemTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	return result, nil
}

// NewListBySystemTopicPager - List event subscriptions that belong to a specific system topic.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - options - SystemTopicEventSubscriptionsClientListBySystemTopicOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.NewListBySystemTopicPager
//     method.
func (client *SystemTopicEventSubscriptionsClient) NewListBySystemTopicPager(resourceGroupName string, systemTopicName string, options *SystemTopicEventSubscriptionsClientListBySystemTopicOptions) *runtime.Pager[SystemTopicEventSubscriptionsClientListBySystemTopicResponse] {
	return runtime.NewPager(runtime.PagingHandler[SystemTopicEventSubscriptionsClientListBySystemTopicResponse]{
		More: func(page SystemTopicEventSubscriptionsClientListBySystemTopicResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemTopicEventSubscriptionsClientListBySystemTopicResponse) (SystemTopicEventSubscriptionsClientListBySystemTopicResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SystemTopicEventSubscriptionsClient.NewListBySystemTopicPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySystemTopicCreateRequest(ctx, resourceGroupName, systemTopicName, options)
			}, nil)
			if err != nil {
				return SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}, err
			}
			return client.listBySystemTopicHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySystemTopicCreateRequest creates the ListBySystemTopic request.
func (client *SystemTopicEventSubscriptionsClient) listBySystemTopicCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, options *SystemTopicEventSubscriptionsClientListBySystemTopicOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySystemTopicHandleResponse handles the ListBySystemTopic response.
func (client *SystemTopicEventSubscriptionsClient) listBySystemTopicHandleResponse(resp *http.Response) (SystemTopicEventSubscriptionsClientListBySystemTopicResponse, error) {
	result := SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionsListResult); err != nil {
		return SystemTopicEventSubscriptionsClientListBySystemTopicResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - systemTopicName - Name of the system topic.
//   - eventSubscriptionName - Name of the event subscription to be updated.
//   - eventSubscriptionUpdateParameters - Updated event subscription information.
//   - options - SystemTopicEventSubscriptionsClientBeginUpdateOptions contains the optional parameters for the SystemTopicEventSubscriptionsClient.BeginUpdate
//     method.
func (client *SystemTopicEventSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *SystemTopicEventSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[SystemTopicEventSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[SystemTopicEventSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[SystemTopicEventSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Update an existing event subscription of a system topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-15-preview
func (client *SystemTopicEventSubscriptionsClient) update(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *SystemTopicEventSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "SystemTopicEventSubscriptionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, systemTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *SystemTopicEventSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, systemTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *SystemTopicEventSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/systemTopics/{systemTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if systemTopicName == "" {
		return nil, errors.New("parameter systemTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemTopicName}", url.PathEscape(systemTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSubscriptionUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
