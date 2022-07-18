//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armresourcehealth

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
	"strings"
)

// EventsClient contains the methods for the Events group.
// Don't use this type directly, use NewEventsClient() instead.
type EventsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewEventsClient creates a new instance of EventsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEventsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EventsClient, error) {
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
	client := &EventsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListBySingleResourcePager - Lists current service health events for given resource.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// resourceURI - The fully qualified ID of the resource, including the resource name and resource type. Currently the API
// support not nested and one nesting level resource types :
// /subscriptions/{subscriptionId}/resourceGroups/{resource-group-name}/providers/{resource-provider-name}/{resource-type}/{resource-name}
// and
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
// options - EventsClientListBySingleResourceOptions contains the optional parameters for the EventsClient.ListBySingleResource
// method.
func (client *EventsClient) NewListBySingleResourcePager(resourceURI string, options *EventsClientListBySingleResourceOptions) *runtime.Pager[EventsClientListBySingleResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventsClientListBySingleResourceResponse]{
		More: func(page EventsClientListBySingleResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EventsClientListBySingleResourceResponse) (EventsClientListBySingleResourceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySingleResourceCreateRequest(ctx, resourceURI, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EventsClientListBySingleResourceResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EventsClientListBySingleResourceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EventsClientListBySingleResourceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySingleResourceHandleResponse(resp)
		},
	})
}

// listBySingleResourceCreateRequest creates the ListBySingleResource request.
func (client *EventsClient) listBySingleResourceCreateRequest(ctx context.Context, resourceURI string, options *EventsClientListBySingleResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/events"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.View != nil {
		reqQP.Set("view", *options.View)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySingleResourceHandleResponse handles the ListBySingleResource response.
func (client *EventsClient) listBySingleResourceHandleResponse(resp *http.Response) (EventsClientListBySingleResourceResponse, error) {
	result := EventsClientListBySingleResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListBySingleResourceResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionIDPager - Lists current service health events in the subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-06-01
// options - EventsClientListBySubscriptionIDOptions contains the optional parameters for the EventsClient.ListBySubscriptionID
// method.
func (client *EventsClient) NewListBySubscriptionIDPager(options *EventsClientListBySubscriptionIDOptions) *runtime.Pager[EventsClientListBySubscriptionIDResponse] {
	return runtime.NewPager(runtime.PagingHandler[EventsClientListBySubscriptionIDResponse]{
		More: func(page EventsClientListBySubscriptionIDResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EventsClientListBySubscriptionIDResponse) (EventsClientListBySubscriptionIDResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionIDCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EventsClientListBySubscriptionIDResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return EventsClientListBySubscriptionIDResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EventsClientListBySubscriptionIDResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionIDHandleResponse(resp)
		},
	})
}

// listBySubscriptionIDCreateRequest creates the ListBySubscriptionID request.
func (client *EventsClient) listBySubscriptionIDCreateRequest(ctx context.Context, options *EventsClientListBySubscriptionIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.QueryStartTime != nil {
		reqQP.Set("queryStartTime", *options.QueryStartTime)
	}
	if options != nil && options.View != nil {
		reqQP.Set("view", *options.View)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionIDHandleResponse handles the ListBySubscriptionID response.
func (client *EventsClient) listBySubscriptionIDHandleResponse(resp *http.Response) (EventsClientListBySubscriptionIDResponse, error) {
	result := EventsClientListBySubscriptionIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListBySubscriptionIDResponse{}, err
	}
	return result, nil
}
