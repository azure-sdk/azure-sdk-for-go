// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdatadog

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

// MonitoredSubscriptionsClient contains the methods for the MonitoredSubscriptions group.
// Don't use this type directly, use NewMonitoredSubscriptionsClient() instead.
type MonitoredSubscriptionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMonitoredSubscriptionsClient creates a new instance of MonitoredSubscriptionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMonitoredSubscriptionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MonitoredSubscriptionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MonitoredSubscriptionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateorUpdate - Add the subscriptions that should be monitored by the Datadog monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - configurationName - The configuration name. Only 'default' value is supported.
//   - options - MonitoredSubscriptionsClientBeginCreateorUpdateOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginCreateorUpdate
//     method.
func (client *MonitoredSubscriptionsClient) BeginCreateorUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginCreateorUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientCreateorUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createorUpdate(ctx, resourceGroupName, monitorName, configurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MonitoredSubscriptionsClientCreateorUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MonitoredSubscriptionsClientCreateorUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateorUpdate - Add the subscriptions that should be monitored by the Datadog monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
func (client *MonitoredSubscriptionsClient) createorUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginCreateorUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MonitoredSubscriptionsClient.BeginCreateorUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createorUpdateCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createorUpdateCreateRequest creates the CreateorUpdate request.
func (client *MonitoredSubscriptionsClient) createorUpdateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginCreateorUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/monitoredSubscriptions/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginDelete - Updates the subscriptions that are being monitored by the Datadog monitor resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - configurationName - Configuration name
//   - options - MonitoredSubscriptionsClientBeginDeleteOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginDelete
//     method.
func (client *MonitoredSubscriptionsClient) BeginDelete(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginDeleteOptions) (*runtime.Poller[MonitoredSubscriptionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, monitorName, configurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MonitoredSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MonitoredSubscriptionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Updates the subscriptions that are being monitored by the Datadog monitor resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
func (client *MonitoredSubscriptionsClient) deleteOperation(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MonitoredSubscriptionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
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
func (client *MonitoredSubscriptionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, _ *MonitoredSubscriptionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/monitoredSubscriptions/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - List the subscriptions currently being monitored by the Datadog monitor resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - configurationName - The configuration name. Only 'default' value is supported.
//   - options - MonitoredSubscriptionsClientGetOptions contains the optional parameters for the MonitoredSubscriptionsClient.Get
//     method.
func (client *MonitoredSubscriptionsClient) Get(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientGetOptions) (MonitoredSubscriptionsClientGetResponse, error) {
	var err error
	const operationName = "MonitoredSubscriptionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
	if err != nil {
		return MonitoredSubscriptionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MonitoredSubscriptionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MonitoredSubscriptionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MonitoredSubscriptionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, _ *MonitoredSubscriptionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/monitoredSubscriptions/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitoredSubscriptionsClient) getHandleResponse(resp *http.Response) (MonitoredSubscriptionsClientGetResponse, error) {
	result := MonitoredSubscriptionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoredSubscriptionProperties); err != nil {
		return MonitoredSubscriptionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List the subscriptions currently being monitored by the Datadog monitor resource.
//
// Generated from API version 2023-10-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - options - MonitoredSubscriptionsClientListOptions contains the optional parameters for the MonitoredSubscriptionsClient.NewListPager
//     method.
func (client *MonitoredSubscriptionsClient) NewListPager(resourceGroupName string, monitorName string, options *MonitoredSubscriptionsClientListOptions) *runtime.Pager[MonitoredSubscriptionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MonitoredSubscriptionsClientListResponse]{
		More: func(page MonitoredSubscriptionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MonitoredSubscriptionsClientListResponse) (MonitoredSubscriptionsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MonitoredSubscriptionsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			}, nil)
			if err != nil {
				return MonitoredSubscriptionsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MonitoredSubscriptionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, _ *MonitoredSubscriptionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/monitoredSubscriptions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MonitoredSubscriptionsClient) listHandleResponse(resp *http.Response) (MonitoredSubscriptionsClientListResponse, error) {
	result := MonitoredSubscriptionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoredSubscriptionPropertiesList); err != nil {
		return MonitoredSubscriptionsClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Updates the subscriptions that are being monitored by the Datadog monitor resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - configurationName - The configuration name. Only 'default' value is supported.
//   - options - MonitoredSubscriptionsClientBeginUpdateOptions contains the optional parameters for the MonitoredSubscriptionsClient.BeginUpdate
//     method.
func (client *MonitoredSubscriptionsClient) BeginUpdate(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginUpdateOptions) (*runtime.Poller[MonitoredSubscriptionsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, monitorName, configurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MonitoredSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MonitoredSubscriptionsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Updates the subscriptions that are being monitored by the Datadog monitor resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-20
func (client *MonitoredSubscriptionsClient) update(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MonitoredSubscriptionsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, monitorName, configurationName, options)
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
func (client *MonitoredSubscriptionsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, configurationName string, options *MonitoredSubscriptionsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Datadog/monitors/{monitorName}/monitoredSubscriptions/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Body != nil {
		if err := runtime.MarshalAsJSON(req, *options.Body); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
