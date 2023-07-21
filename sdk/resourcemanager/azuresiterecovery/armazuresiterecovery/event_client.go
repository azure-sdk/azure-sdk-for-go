//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuresiterecovery

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

// EventClient contains the methods for the Event group.
// Don't use this type directly, use NewEventClient() instead.
type EventClient struct {
	internal *arm.Client
}

// NewEventClient creates a new instance of EventClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEventClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*EventClient, error) {
	cl, err := arm.NewClient(moduleName+".EventClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EventClient{
		internal: cl,
	}
	return client, nil
}

// Get - Gets the details of the event.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-16-preview
//   - subscriptionID - The subscription Id.
//   - resourceGroupName - Resource group name.
//   - vaultName - Vault name.
//   - eventName - Event name.
//   - options - EventClientGetOptions contains the optional parameters for the EventClient.Get method.
func (client *EventClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, eventName string, options *EventClientGetOptions) (EventClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, vaultName, eventName, options)
	if err != nil {
		return EventClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EventClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, eventName string, options *EventClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/events/{eventName}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	if eventName == "" {
		return nil, errors.New("parameter eventName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventName}", url.PathEscape(eventName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EventClient) getHandleResponse(resp *http.Response) (EventClientGetResponse, error) {
	result := EventClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventModel); err != nil {
		return EventClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of events in the given vault.
//
// Generated from API version 2021-02-16-preview
//   - subscriptionID - The subscription Id.
//   - resourceGroupName - Resource group name.
//   - vaultName - Vault name.
//   - options - EventClientListOptions contains the optional parameters for the EventClient.NewListPager method.
func (client *EventClient) NewListPager(subscriptionID string, resourceGroupName string, vaultName string, options *EventClientListOptions) *runtime.Pager[EventClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventClientListResponse]{
		More: func(page EventClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EventClientListResponse) (EventClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, subscriptionID, resourceGroupName, vaultName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EventClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EventClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EventClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EventClient) listCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, vaultName string, options *EventClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataReplication/replicationVaults/{vaultName}/events"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vaultName == "" {
		return nil, errors.New("parameter vaultName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.ContinuationToken != nil {
		reqQP.Set("continuationToken", *options.ContinuationToken)
	}
	reqQP.Set("api-version", "2021-02-16-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EventClient) listHandleResponse(resp *http.Response) (EventClientListResponse, error) {
	result := EventClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EventModelCollection); err != nil {
		return EventClientListResponse{}, err
	}
	return result, nil
}
