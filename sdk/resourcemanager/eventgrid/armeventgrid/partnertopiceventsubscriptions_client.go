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

// PartnerTopicEventSubscriptionsClient contains the methods for the PartnerTopicEventSubscriptions group.
// Don't use this type directly, use NewPartnerTopicEventSubscriptionsClient() instead.
type PartnerTopicEventSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPartnerTopicEventSubscriptionsClient creates a new instance of PartnerTopicEventSubscriptionsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPartnerTopicEventSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PartnerTopicEventSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName+".PartnerTopicEventSubscriptionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PartnerTopicEventSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Asynchronously creates or updates an event subscription of a partner topic with the specified parameters.
// Existing event subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
//     characters in length and use alphanumeric letters only.
//   - eventSubscriptionInfo - Event subscription properties containing the destination and filter information.
//   - options - PartnerTopicEventSubscriptionsClientBeginCreateOrUpdateOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.BeginCreateOrUpdate
//     method.
func (client *PartnerTopicEventSubscriptionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *PartnerTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PartnerTopicEventSubscriptionsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Asynchronously creates or updates an event subscription of a partner topic with the specified parameters.
// Existing event subscriptions will be updated with this API.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *PartnerTopicEventSubscriptionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *PartnerTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionInfo, options)
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
func (client *PartnerTopicEventSubscriptionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionInfo EventSubscription, options *PartnerTopicEventSubscriptionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSubscriptionInfo); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete an existing event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
//     characters in length and use alphanumeric letters only.
//   - options - PartnerTopicEventSubscriptionsClientBeginDeleteOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.BeginDelete
//     method.
func (client *PartnerTopicEventSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[PartnerTopicEventSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PartnerTopicEventSubscriptionsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PartnerTopicEventSubscriptionsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete an existing event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *PartnerTopicEventSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
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
func (client *PartnerTopicEventSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Get properties of an event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - eventSubscriptionName - Name of the event subscription to be found. Event subscription names must be between 3 and 100
//     characters in length and use alphanumeric letters only.
//   - options - PartnerTopicEventSubscriptionsClientGetOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.Get
//     method.
func (client *PartnerTopicEventSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientGetOptions) (PartnerTopicEventSubscriptionsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerTopicEventSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerTopicEventSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PartnerTopicEventSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PartnerTopicEventSubscriptionsClient) getHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsClientGetResponse, error) {
	result := PartnerTopicEventSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscription); err != nil {
		return PartnerTopicEventSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// GetDeliveryAttributes - Get all delivery attributes for an event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
//     characters in length and use alphanumeric letters only.
//   - options - PartnerTopicEventSubscriptionsClientGetDeliveryAttributesOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.GetDeliveryAttributes
//     method.
func (client *PartnerTopicEventSubscriptionsClient) GetDeliveryAttributes(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	var err error
	req, err := client.getDeliveryAttributesCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	resp, err := client.getDeliveryAttributesHandleResponse(httpResp)
	return resp, err
}

// getDeliveryAttributesCreateRequest creates the GetDeliveryAttributes request.
func (client *PartnerTopicEventSubscriptionsClient) getDeliveryAttributesCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientGetDeliveryAttributesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}/getDeliveryAttributes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeliveryAttributesHandleResponse handles the GetDeliveryAttributes response.
func (client *PartnerTopicEventSubscriptionsClient) getDeliveryAttributesHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse, error) {
	result := PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeliveryAttributeListResult); err != nil {
		return PartnerTopicEventSubscriptionsClientGetDeliveryAttributesResponse{}, err
	}
	return result, nil
}

// GetFullURL - Get the full endpoint URL for an event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
//     characters in length and use alphanumeric letters only.
//   - options - PartnerTopicEventSubscriptionsClientGetFullURLOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.GetFullURL
//     method.
func (client *PartnerTopicEventSubscriptionsClient) GetFullURL(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientGetFullURLOptions) (PartnerTopicEventSubscriptionsClientGetFullURLResponse, error) {
	var err error
	req, err := client.getFullURLCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, options)
	if err != nil {
		return PartnerTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PartnerTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PartnerTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	resp, err := client.getFullURLHandleResponse(httpResp)
	return resp, err
}

// getFullURLCreateRequest creates the GetFullURL request.
func (client *PartnerTopicEventSubscriptionsClient) getFullURLCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, options *PartnerTopicEventSubscriptionsClientGetFullURLOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}/getFullUrl"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getFullURLHandleResponse handles the GetFullURL response.
func (client *PartnerTopicEventSubscriptionsClient) getFullURLHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsClientGetFullURLResponse, error) {
	result := PartnerTopicEventSubscriptionsClientGetFullURLResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionFullURL); err != nil {
		return PartnerTopicEventSubscriptionsClientGetFullURLResponse{}, err
	}
	return result, nil
}

// NewListByPartnerTopicPager - List event subscriptions that belong to a specific partner topic.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - options - PartnerTopicEventSubscriptionsClientListByPartnerTopicOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.NewListByPartnerTopicPager
//     method.
func (client *PartnerTopicEventSubscriptionsClient) NewListByPartnerTopicPager(resourceGroupName string, partnerTopicName string, options *PartnerTopicEventSubscriptionsClientListByPartnerTopicOptions) *runtime.Pager[PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse] {
	return runtime.NewPager(runtime.PagingHandler[PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse]{
		More: func(page PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse) (PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByPartnerTopicCreateRequest(ctx, resourceGroupName, partnerTopicName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByPartnerTopicHandleResponse(resp)
		},
	})
}

// listByPartnerTopicCreateRequest creates the ListByPartnerTopic request.
func (client *PartnerTopicEventSubscriptionsClient) listByPartnerTopicCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, options *PartnerTopicEventSubscriptionsClientListByPartnerTopicOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByPartnerTopicHandleResponse handles the ListByPartnerTopic response.
func (client *PartnerTopicEventSubscriptionsClient) listByPartnerTopicHandleResponse(resp *http.Response) (PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse, error) {
	result := PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventSubscriptionsListResult); err != nil {
		return PartnerTopicEventSubscriptionsClientListByPartnerTopicResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update an existing event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group within the user's subscription.
//   - partnerTopicName - Name of the partner topic.
//   - eventSubscriptionName - Name of the event subscription to be created. Event subscription names must be between 3 and 100
//     characters in length and use alphanumeric letters only.
//   - eventSubscriptionUpdateParameters - Updated event subscription information.
//   - options - PartnerTopicEventSubscriptionsClientBeginUpdateOptions contains the optional parameters for the PartnerTopicEventSubscriptionsClient.BeginUpdate
//     method.
func (client *PartnerTopicEventSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *PartnerTopicEventSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[PartnerTopicEventSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[PartnerTopicEventSubscriptionsClientUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[PartnerTopicEventSubscriptionsClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update an existing event subscription of a partner topic.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *PartnerTopicEventSubscriptionsClient) update(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *PartnerTopicEventSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, partnerTopicName, eventSubscriptionName, eventSubscriptionUpdateParameters, options)
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
func (client *PartnerTopicEventSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, partnerTopicName string, eventSubscriptionName string, eventSubscriptionUpdateParameters EventSubscriptionUpdateParameters, options *PartnerTopicEventSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventGrid/partnerTopics/{partnerTopicName}/eventSubscriptions/{eventSubscriptionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if partnerTopicName == "" {
		return nil, errors.New("parameter partnerTopicName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{partnerTopicName}", url.PathEscape(partnerTopicName))
	if eventSubscriptionName == "" {
		return nil, errors.New("parameter eventSubscriptionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventSubscriptionName}", url.PathEscape(eventSubscriptionName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, eventSubscriptionUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
