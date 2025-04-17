// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

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
	internal       *arm.Client
	subscriptionID string
}

// NewEventClient creates a new instance of EventClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEventClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EventClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EventClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// FetchDetailsBySubscriptionIDAndTrackingID - Service health event details in the subscription by event tracking id. This
// can be used to fetch sensitive properties for Security Advisory events
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - options - EventClientFetchDetailsBySubscriptionIDAndTrackingIDOptions contains the optional parameters for the EventClient.FetchDetailsBySubscriptionIDAndTrackingID
//     method.
func (client *EventClient) FetchDetailsBySubscriptionIDAndTrackingID(ctx context.Context, eventTrackingID string, options *EventClientFetchDetailsBySubscriptionIDAndTrackingIDOptions) (EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse, error) {
	var err error
	const operationName = "EventClient.FetchDetailsBySubscriptionIDAndTrackingID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fetchDetailsBySubscriptionIDAndTrackingIDCreateRequest(ctx, eventTrackingID, options)
	if err != nil {
		return EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse{}, err
	}
	resp, err := client.fetchDetailsBySubscriptionIDAndTrackingIDHandleResponse(httpResp)
	return resp, err
}

// fetchDetailsBySubscriptionIDAndTrackingIDCreateRequest creates the FetchDetailsBySubscriptionIDAndTrackingID request.
func (client *EventClient) fetchDetailsBySubscriptionIDAndTrackingIDCreateRequest(ctx context.Context, eventTrackingID string, _ *EventClientFetchDetailsBySubscriptionIDAndTrackingIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/fetchEventDetails"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// fetchDetailsBySubscriptionIDAndTrackingIDHandleResponse handles the FetchDetailsBySubscriptionIDAndTrackingID response.
func (client *EventClient) fetchDetailsBySubscriptionIDAndTrackingIDHandleResponse(resp *http.Response) (EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse, error) {
	result := EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Event); err != nil {
		return EventClientFetchDetailsBySubscriptionIDAndTrackingIDResponse{}, err
	}
	return result, nil
}

// FetchDetailsByTenantIDAndTrackingID - Service health event details in the tenant by event tracking id. This can be used
// to fetch sensitive properties for Security Advisory events
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - options - EventClientFetchDetailsByTenantIDAndTrackingIDOptions contains the optional parameters for the EventClient.FetchDetailsByTenantIDAndTrackingID
//     method.
func (client *EventClient) FetchDetailsByTenantIDAndTrackingID(ctx context.Context, eventTrackingID string, options *EventClientFetchDetailsByTenantIDAndTrackingIDOptions) (EventClientFetchDetailsByTenantIDAndTrackingIDResponse, error) {
	var err error
	const operationName = "EventClient.FetchDetailsByTenantIDAndTrackingID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.fetchDetailsByTenantIDAndTrackingIDCreateRequest(ctx, eventTrackingID, options)
	if err != nil {
		return EventClientFetchDetailsByTenantIDAndTrackingIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventClientFetchDetailsByTenantIDAndTrackingIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventClientFetchDetailsByTenantIDAndTrackingIDResponse{}, err
	}
	resp, err := client.fetchDetailsByTenantIDAndTrackingIDHandleResponse(httpResp)
	return resp, err
}

// fetchDetailsByTenantIDAndTrackingIDCreateRequest creates the FetchDetailsByTenantIDAndTrackingID request.
func (client *EventClient) fetchDetailsByTenantIDAndTrackingIDCreateRequest(ctx context.Context, eventTrackingID string, _ *EventClientFetchDetailsByTenantIDAndTrackingIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/events/{eventTrackingId}/fetchEventDetails"
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// fetchDetailsByTenantIDAndTrackingIDHandleResponse handles the FetchDetailsByTenantIDAndTrackingID response.
func (client *EventClient) fetchDetailsByTenantIDAndTrackingIDHandleResponse(resp *http.Response) (EventClientFetchDetailsByTenantIDAndTrackingIDResponse, error) {
	result := EventClientFetchDetailsByTenantIDAndTrackingIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Event); err != nil {
		return EventClientFetchDetailsByTenantIDAndTrackingIDResponse{}, err
	}
	return result, nil
}

// GetBySubscriptionIDAndTrackingID - Service health event in the subscription by event tracking id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - options - EventClientGetBySubscriptionIDAndTrackingIDOptions contains the optional parameters for the EventClient.GetBySubscriptionIDAndTrackingID
//     method.
func (client *EventClient) GetBySubscriptionIDAndTrackingID(ctx context.Context, eventTrackingID string, options *EventClientGetBySubscriptionIDAndTrackingIDOptions) (EventClientGetBySubscriptionIDAndTrackingIDResponse, error) {
	var err error
	const operationName = "EventClient.GetBySubscriptionIDAndTrackingID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getBySubscriptionIDAndTrackingIDCreateRequest(ctx, eventTrackingID, options)
	if err != nil {
		return EventClientGetBySubscriptionIDAndTrackingIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventClientGetBySubscriptionIDAndTrackingIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventClientGetBySubscriptionIDAndTrackingIDResponse{}, err
	}
	resp, err := client.getBySubscriptionIDAndTrackingIDHandleResponse(httpResp)
	return resp, err
}

// getBySubscriptionIDAndTrackingIDCreateRequest creates the GetBySubscriptionIDAndTrackingID request.
func (client *EventClient) getBySubscriptionIDAndTrackingIDCreateRequest(ctx context.Context, eventTrackingID string, options *EventClientGetBySubscriptionIDAndTrackingIDOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ResourceHealth/events/{eventTrackingId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.QueryStartTime != nil {
		reqQP.Set("queryStartTime", *options.QueryStartTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getBySubscriptionIDAndTrackingIDHandleResponse handles the GetBySubscriptionIDAndTrackingID response.
func (client *EventClient) getBySubscriptionIDAndTrackingIDHandleResponse(resp *http.Response) (EventClientGetBySubscriptionIDAndTrackingIDResponse, error) {
	result := EventClientGetBySubscriptionIDAndTrackingIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Event); err != nil {
		return EventClientGetBySubscriptionIDAndTrackingIDResponse{}, err
	}
	return result, nil
}

// GetByTenantIDAndTrackingID - Service health event in the tenant by event tracking id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - eventTrackingID - Event Id which uniquely identifies ServiceHealth event.
//   - options - EventClientGetByTenantIDAndTrackingIDOptions contains the optional parameters for the EventClient.GetByTenantIDAndTrackingID
//     method.
func (client *EventClient) GetByTenantIDAndTrackingID(ctx context.Context, eventTrackingID string, options *EventClientGetByTenantIDAndTrackingIDOptions) (EventClientGetByTenantIDAndTrackingIDResponse, error) {
	var err error
	const operationName = "EventClient.GetByTenantIDAndTrackingID"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByTenantIDAndTrackingIDCreateRequest(ctx, eventTrackingID, options)
	if err != nil {
		return EventClientGetByTenantIDAndTrackingIDResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EventClientGetByTenantIDAndTrackingIDResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EventClientGetByTenantIDAndTrackingIDResponse{}, err
	}
	resp, err := client.getByTenantIDAndTrackingIDHandleResponse(httpResp)
	return resp, err
}

// getByTenantIDAndTrackingIDCreateRequest creates the GetByTenantIDAndTrackingID request.
func (client *EventClient) getByTenantIDAndTrackingIDCreateRequest(ctx context.Context, eventTrackingID string, options *EventClientGetByTenantIDAndTrackingIDOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.ResourceHealth/events/{eventTrackingId}"
	if eventTrackingID == "" {
		return nil, errors.New("parameter eventTrackingID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{eventTrackingId}", url.PathEscape(eventTrackingID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.QueryStartTime != nil {
		reqQP.Set("queryStartTime", *options.QueryStartTime)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByTenantIDAndTrackingIDHandleResponse handles the GetByTenantIDAndTrackingID response.
func (client *EventClient) getByTenantIDAndTrackingIDHandleResponse(resp *http.Response) (EventClientGetByTenantIDAndTrackingIDResponse, error) {
	result := EventClientGetByTenantIDAndTrackingIDResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Event); err != nil {
		return EventClientGetByTenantIDAndTrackingIDResponse{}, err
	}
	return result, nil
}
