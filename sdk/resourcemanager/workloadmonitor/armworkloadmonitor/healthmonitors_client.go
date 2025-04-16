// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armworkloadmonitor

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
	"time"
)

// HealthMonitorsClient contains the methods for the HealthMonitors group.
// Don't use this type directly, use NewHealthMonitorsClient() instead.
type HealthMonitorsClient struct {
	internal *arm.Client
}

// NewHealthMonitorsClient creates a new instance of HealthMonitorsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHealthMonitorsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*HealthMonitorsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HealthMonitorsClient{
		internal: cl,
	}
	return client, nil
}

// Get - Get the current health status of a monitor of a virtual machine. Optional parameter: $expand (retrieve the monitor's
// evidence and configuration).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - subscriptionID - The subscription Id of the virtual machine.
//   - resourceGroupName - The resource group of the virtual machine.
//   - providerName - The provider name (ex: Microsoft.Compute for virtual machines).
//   - resourceCollectionName - The resource collection name (ex: virtualMachines for virtual machines).
//   - resourceName - The name of the virtual machine.
//   - monitorID - The monitor Id of the virtual machine.
//   - options - HealthMonitorsClientGetOptions contains the optional parameters for the HealthMonitorsClient.Get method.
func (client *HealthMonitorsClient) Get(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsClientGetOptions) (HealthMonitorsClientGetResponse, error) {
	var err error
	const operationName = "HealthMonitorsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, options)
	if err != nil {
		return HealthMonitorsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HealthMonitorsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HealthMonitorsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HealthMonitorsClient) getCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if monitorID == "" {
		return nil, errors.New("parameter monitorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorId}", url.PathEscape(monitorID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HealthMonitorsClient) getHandleResponse(resp *http.Response) (HealthMonitorsClientGetResponse, error) {
	result := HealthMonitorsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitor); err != nil {
		return HealthMonitorsClientGetResponse{}, err
	}
	return result, nil
}

// GetStateChange - Get the health state change of a monitor of a virtual machine at the provided timestamp. Optional parameter:
// $expand (retrieve the monitor's evidence and configuration).
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-13-preview
//   - subscriptionID - The subscription Id of the virtual machine.
//   - resourceGroupName - The resource group of the virtual machine.
//   - providerName - The provider name (ex: Microsoft.Compute for virtual machines).
//   - resourceCollectionName - The resource collection name (ex: virtualMachines for virtual machines).
//   - resourceName - The name of the virtual machine.
//   - monitorID - The monitor Id of the virtual machine.
//   - timestampUnix - The timestamp of the state change (unix format).
//   - options - HealthMonitorsClientGetStateChangeOptions contains the optional parameters for the HealthMonitorsClient.GetStateChange
//     method.
func (client *HealthMonitorsClient) GetStateChange(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, timestampUnix string, options *HealthMonitorsClientGetStateChangeOptions) (HealthMonitorsClientGetStateChangeResponse, error) {
	var err error
	const operationName = "HealthMonitorsClient.GetStateChange"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getStateChangeCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, timestampUnix, options)
	if err != nil {
		return HealthMonitorsClientGetStateChangeResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HealthMonitorsClientGetStateChangeResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HealthMonitorsClientGetStateChangeResponse{}, err
	}
	resp, err := client.getStateChangeHandleResponse(httpResp)
	return resp, err
}

// getStateChangeCreateRequest creates the GetStateChange request.
func (client *HealthMonitorsClient) getStateChangeCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, timestampUnix string, options *HealthMonitorsClientGetStateChangeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history/{timestampUnix}"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if monitorID == "" {
		return nil, errors.New("parameter monitorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorId}", url.PathEscape(monitorID))
	if timestampUnix == "" {
		return nil, errors.New("parameter timestampUnix cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timestampUnix}", url.PathEscape(timestampUnix))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getStateChangeHandleResponse handles the GetStateChange response.
func (client *HealthMonitorsClient) getStateChangeHandleResponse(resp *http.Response) (HealthMonitorsClientGetStateChangeResponse, error) {
	result := HealthMonitorsClientGetStateChangeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitorStateChange); err != nil {
		return HealthMonitorsClientGetStateChangeResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the current health status of all monitors of a virtual machine. Optional parameters: $expand (retrieve
// the monitor's evidence and configuration) and $filter (filter by monitor name).
//
// Generated from API version 2020-01-13-preview
//   - subscriptionID - The subscription Id of the virtual machine.
//   - resourceGroupName - The resource group of the virtual machine.
//   - providerName - The provider name (ex: Microsoft.Compute for virtual machines).
//   - resourceCollectionName - The resource collection name (ex: virtualMachines for virtual machines).
//   - resourceName - The name of the virtual machine.
//   - options - HealthMonitorsClientListOptions contains the optional parameters for the HealthMonitorsClient.NewListPager method.
func (client *HealthMonitorsClient) NewListPager(subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, options *HealthMonitorsClientListOptions) *runtime.Pager[HealthMonitorsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HealthMonitorsClientListResponse]{
		More: func(page HealthMonitorsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HealthMonitorsClientListResponse) (HealthMonitorsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HealthMonitorsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, options)
			}, nil)
			if err != nil {
				return HealthMonitorsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HealthMonitorsClient) listCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, options *HealthMonitorsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HealthMonitorsClient) listHandleResponse(resp *http.Response) (HealthMonitorsClientListResponse, error) {
	result := HealthMonitorsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitorList); err != nil {
		return HealthMonitorsClientListResponse{}, err
	}
	return result, nil
}

// NewListStateChangesPager - Get the health state changes of a monitor of a virtual machine within the provided time window
// (default is the last 24 hours). Optional parameters: $expand (retrieve the monitor's evidence and
// configuration) and $filter (filter by heartbeat condition).
//
// Generated from API version 2020-01-13-preview
//   - subscriptionID - The subscription Id of the virtual machine.
//   - resourceGroupName - The resource group of the virtual machine.
//   - providerName - The provider name (ex: Microsoft.Compute for virtual machines).
//   - resourceCollectionName - The resource collection name (ex: virtualMachines for virtual machines).
//   - resourceName - The name of the virtual machine.
//   - monitorID - The monitor Id of the virtual machine.
//   - options - HealthMonitorsClientListStateChangesOptions contains the optional parameters for the HealthMonitorsClient.NewListStateChangesPager
//     method.
func (client *HealthMonitorsClient) NewListStateChangesPager(subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsClientListStateChangesOptions) *runtime.Pager[HealthMonitorsClientListStateChangesResponse] {
	return runtime.NewPager(runtime.PagingHandler[HealthMonitorsClientListStateChangesResponse]{
		More: func(page HealthMonitorsClientListStateChangesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HealthMonitorsClientListStateChangesResponse) (HealthMonitorsClientListStateChangesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HealthMonitorsClient.NewListStateChangesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listStateChangesCreateRequest(ctx, subscriptionID, resourceGroupName, providerName, resourceCollectionName, resourceName, monitorID, options)
			}, nil)
			if err != nil {
				return HealthMonitorsClientListStateChangesResponse{}, err
			}
			return client.listStateChangesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listStateChangesCreateRequest creates the ListStateChanges request.
func (client *HealthMonitorsClient) listStateChangesCreateRequest(ctx context.Context, subscriptionID string, resourceGroupName string, providerName string, resourceCollectionName string, resourceName string, monitorID string, options *HealthMonitorsClientListStateChangesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{providerName}/{resourceCollectionName}/{resourceName}/providers/Microsoft.WorkloadMonitor/monitors/{monitorId}/history"
	if subscriptionID == "" {
		return nil, errors.New("parameter subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceCollectionName == "" {
		return nil, errors.New("parameter resourceCollectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceCollectionName}", url.PathEscape(resourceCollectionName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if monitorID == "" {
		return nil, errors.New("parameter monitorID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorId}", url.PathEscape(monitorID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2020-01-13-preview")
	if options != nil && options.EndTimestampUTC != nil {
		reqQP.Set("endTimestampUtc", options.EndTimestampUTC.Format(time.RFC3339Nano))
	}
	if options != nil && options.StartTimestampUTC != nil {
		reqQP.Set("startTimestampUtc", options.StartTimestampUTC.Format(time.RFC3339Nano))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listStateChangesHandleResponse handles the ListStateChanges response.
func (client *HealthMonitorsClient) listStateChangesHandleResponse(resp *http.Response) (HealthMonitorsClientListStateChangesResponse, error) {
	result := HealthMonitorsClientListStateChangesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthMonitorStateChangeList); err != nil {
		return HealthMonitorsClientListStateChangesResponse{}, err
	}
	return result, nil
}
