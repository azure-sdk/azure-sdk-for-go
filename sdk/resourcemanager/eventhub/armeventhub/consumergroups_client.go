//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

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

// ConsumerGroupsClient contains the methods for the ConsumerGroups group.
// Don't use this type directly, use NewConsumerGroupsClient() instead.
type ConsumerGroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConsumerGroupsClient creates a new instance of ConsumerGroupsClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConsumerGroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConsumerGroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConsumerGroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an Event Hubs consumer group as a nested resource within a Namespace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - eventHubName - The Event Hub name
//   - consumerGroupName - The consumer group name
//   - parameters - Parameters supplied to create or update a consumer group resource.
//   - options - ConsumerGroupsClientCreateOrUpdateOptions contains the optional parameters for the ConsumerGroupsClient.CreateOrUpdate
//     method.
func (client *ConsumerGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters ConsumerGroup, options *ConsumerGroupsClientCreateOrUpdateOptions) (ConsumerGroupsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConsumerGroupsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName, parameters, options)
	if err != nil {
		return ConsumerGroupsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConsumerGroupsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConsumerGroupsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConsumerGroupsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, parameters ConsumerGroup, options *ConsumerGroupsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if consumerGroupName == "" {
		return nil, errors.New("parameter consumerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consumerGroupName}", url.PathEscape(consumerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConsumerGroupsClient) createOrUpdateHandleResponse(resp *http.Response) (ConsumerGroupsClientCreateOrUpdateResponse, error) {
	result := ConsumerGroupsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerGroup); err != nil {
		return ConsumerGroupsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes a consumer group from the specified Event Hub and resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - eventHubName - The Event Hub name
//   - consumerGroupName - The consumer group name
//   - options - ConsumerGroupsClientDeleteOptions contains the optional parameters for the ConsumerGroupsClient.Delete method.
func (client *ConsumerGroupsClient) Delete(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, options *ConsumerGroupsClientDeleteOptions) (ConsumerGroupsClientDeleteResponse, error) {
	var err error
	const operationName = "ConsumerGroupsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName, options)
	if err != nil {
		return ConsumerGroupsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConsumerGroupsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConsumerGroupsClientDeleteResponse{}, err
	}
	return ConsumerGroupsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConsumerGroupsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, options *ConsumerGroupsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if consumerGroupName == "" {
		return nil, errors.New("parameter consumerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consumerGroupName}", url.PathEscape(consumerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a description for the specified consumer group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - eventHubName - The Event Hub name
//   - consumerGroupName - The consumer group name
//   - options - ConsumerGroupsClientGetOptions contains the optional parameters for the ConsumerGroupsClient.Get method.
func (client *ConsumerGroupsClient) Get(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, options *ConsumerGroupsClientGetOptions) (ConsumerGroupsClientGetResponse, error) {
	var err error
	const operationName = "ConsumerGroupsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, consumerGroupName, options)
	if err != nil {
		return ConsumerGroupsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConsumerGroupsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConsumerGroupsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConsumerGroupsClient) getCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, consumerGroupName string, options *ConsumerGroupsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if consumerGroupName == "" {
		return nil, errors.New("parameter consumerGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{consumerGroupName}", url.PathEscape(consumerGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConsumerGroupsClient) getHandleResponse(resp *http.Response) (ConsumerGroupsClientGetResponse, error) {
	result := ConsumerGroupsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerGroup); err != nil {
		return ConsumerGroupsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByEventHubPager - Gets all the consumer groups in a Namespace. An empty feed is returned if no consumer group exists
// in the Namespace.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - eventHubName - The Event Hub name
//   - options - ConsumerGroupsClientListByEventHubOptions contains the optional parameters for the ConsumerGroupsClient.NewListByEventHubPager
//     method.
func (client *ConsumerGroupsClient) NewListByEventHubPager(resourceGroupName string, namespaceName string, eventHubName string, options *ConsumerGroupsClientListByEventHubOptions) *runtime.Pager[ConsumerGroupsClientListByEventHubResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConsumerGroupsClientListByEventHubResponse]{
		More: func(page ConsumerGroupsClientListByEventHubResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConsumerGroupsClientListByEventHubResponse) (ConsumerGroupsClientListByEventHubResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConsumerGroupsClient.NewListByEventHubPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByEventHubCreateRequest(ctx, resourceGroupName, namespaceName, eventHubName, options)
			}, nil)
			if err != nil {
				return ConsumerGroupsClientListByEventHubResponse{}, err
			}
			return client.listByEventHubHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByEventHubCreateRequest creates the ListByEventHub request.
func (client *ConsumerGroupsClient) listByEventHubCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, eventHubName string, options *ConsumerGroupsClientListByEventHubOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if eventHubName == "" {
		return nil, errors.New("parameter eventHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventHubName}", url.PathEscape(eventHubName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByEventHubHandleResponse handles the ListByEventHub response.
func (client *ConsumerGroupsClient) listByEventHubHandleResponse(resp *http.Response) (ConsumerGroupsClientListByEventHubResponse, error) {
	result := ConsumerGroupsClientListByEventHubResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConsumerGroupListResult); err != nil {
		return ConsumerGroupsClientListByEventHubResponse{}, err
	}
	return result, nil
}
