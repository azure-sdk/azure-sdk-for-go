//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetworkcloud

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

// MetricsConfigurationsClient contains the methods for the MetricsConfigurations group.
// Don't use this type directly, use NewMetricsConfigurationsClient() instead.
type MetricsConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMetricsConfigurationsClient creates a new instance of MetricsConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetricsConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetricsConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetricsConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create new or update the existing metrics configuration of the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - metricsConfigurationName - The name of the metrics configuration for the cluster.
//   - metricsConfigurationParameters - The request body.
//   - options - MetricsConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the MetricsConfigurationsClient.BeginCreateOrUpdate
//     method.
func (client *MetricsConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, metricsConfigurationParameters ClusterMetricsConfiguration, options *MetricsConfigurationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[MetricsConfigurationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterName, metricsConfigurationName, metricsConfigurationParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MetricsConfigurationsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MetricsConfigurationsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Create new or update the existing metrics configuration of the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
func (client *MetricsConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, metricsConfigurationParameters ClusterMetricsConfiguration, options *MetricsConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MetricsConfigurationsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterName, metricsConfigurationName, metricsConfigurationParameters, options)
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
func (client *MetricsConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, metricsConfigurationParameters ClusterMetricsConfiguration, options *MetricsConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/metricsConfigurations/{metricsConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if metricsConfigurationName == "" {
		return nil, errors.New("parameter metricsConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metricsConfigurationName}", url.PathEscape(metricsConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, metricsConfigurationParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete the metrics configuration of the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - metricsConfigurationName - The name of the metrics configuration for the cluster.
//   - options - MetricsConfigurationsClientBeginDeleteOptions contains the optional parameters for the MetricsConfigurationsClient.BeginDelete
//     method.
func (client *MetricsConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, options *MetricsConfigurationsClientBeginDeleteOptions) (*runtime.Poller[MetricsConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterName, metricsConfigurationName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MetricsConfigurationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MetricsConfigurationsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete the metrics configuration of the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
func (client *MetricsConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, options *MetricsConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MetricsConfigurationsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterName, metricsConfigurationName, options)
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
func (client *MetricsConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, options *MetricsConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/metricsConfigurations/{metricsConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if metricsConfigurationName == "" {
		return nil, errors.New("parameter metricsConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metricsConfigurationName}", url.PathEscape(metricsConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get metrics configuration of the provided cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - metricsConfigurationName - The name of the metrics configuration for the cluster.
//   - options - MetricsConfigurationsClientGetOptions contains the optional parameters for the MetricsConfigurationsClient.Get
//     method.
func (client *MetricsConfigurationsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, options *MetricsConfigurationsClientGetOptions) (MetricsConfigurationsClientGetResponse, error) {
	var err error
	const operationName = "MetricsConfigurationsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterName, metricsConfigurationName, options)
	if err != nil {
		return MetricsConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetricsConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetricsConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MetricsConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, options *MetricsConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/metricsConfigurations/{metricsConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if metricsConfigurationName == "" {
		return nil, errors.New("parameter metricsConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metricsConfigurationName}", url.PathEscape(metricsConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MetricsConfigurationsClient) getHandleResponse(resp *http.Response) (MetricsConfigurationsClientGetResponse, error) {
	result := MetricsConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterMetricsConfiguration); err != nil {
		return MetricsConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByClusterPager - Get a list of metrics configurations for the provided cluster.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - options - MetricsConfigurationsClientListByClusterOptions contains the optional parameters for the MetricsConfigurationsClient.NewListByClusterPager
//     method.
func (client *MetricsConfigurationsClient) NewListByClusterPager(resourceGroupName string, clusterName string, options *MetricsConfigurationsClientListByClusterOptions) *runtime.Pager[MetricsConfigurationsClientListByClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetricsConfigurationsClientListByClusterResponse]{
		More: func(page MetricsConfigurationsClientListByClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MetricsConfigurationsClientListByClusterResponse) (MetricsConfigurationsClientListByClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MetricsConfigurationsClient.NewListByClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByClusterCreateRequest(ctx, resourceGroupName, clusterName, options)
			}, nil)
			if err != nil {
				return MetricsConfigurationsClientListByClusterResponse{}, err
			}
			return client.listByClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByClusterCreateRequest creates the ListByCluster request.
func (client *MetricsConfigurationsClient) listByClusterCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, options *MetricsConfigurationsClientListByClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/metricsConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByClusterHandleResponse handles the ListByCluster response.
func (client *MetricsConfigurationsClient) listByClusterHandleResponse(resp *http.Response) (MetricsConfigurationsClientListByClusterResponse, error) {
	result := MetricsConfigurationsClientListByClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterMetricsConfigurationList); err != nil {
		return MetricsConfigurationsClientListByClusterResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch properties of metrics configuration for the provided cluster, or update the tags associated with it.
// Properties and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the cluster.
//   - metricsConfigurationName - The name of the metrics configuration for the cluster.
//   - metricsConfigurationUpdateParameters - The request body.
//   - options - MetricsConfigurationsClientBeginUpdateOptions contains the optional parameters for the MetricsConfigurationsClient.BeginUpdate
//     method.
func (client *MetricsConfigurationsClient) BeginUpdate(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, metricsConfigurationUpdateParameters ClusterMetricsConfigurationPatchParameters, options *MetricsConfigurationsClientBeginUpdateOptions) (*runtime.Poller[MetricsConfigurationsClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, clusterName, metricsConfigurationName, metricsConfigurationUpdateParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MetricsConfigurationsClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MetricsConfigurationsClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - Patch properties of metrics configuration for the provided cluster, or update the tags associated with it. Properties
// and tag updates can be done independently.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-01-preview
func (client *MetricsConfigurationsClient) update(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, metricsConfigurationUpdateParameters ClusterMetricsConfigurationPatchParameters, options *MetricsConfigurationsClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MetricsConfigurationsClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, clusterName, metricsConfigurationName, metricsConfigurationUpdateParameters, options)
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
func (client *MetricsConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, metricsConfigurationName string, metricsConfigurationUpdateParameters ClusterMetricsConfigurationPatchParameters, options *MetricsConfigurationsClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.NetworkCloud/clusters/{clusterName}/metricsConfigurations/{metricsConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if metricsConfigurationName == "" {
		return nil, errors.New("parameter metricsConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metricsConfigurationName}", url.PathEscape(metricsConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, metricsConfigurationUpdateParameters); err != nil {
		return nil, err
	}
	return req, nil
}
