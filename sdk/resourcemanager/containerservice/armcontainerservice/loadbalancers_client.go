//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcontainerservice

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

// LoadBalancersClient contains the methods for the LoadBalancers group.
// Don't use this type directly, use NewLoadBalancersClient() instead.
type LoadBalancersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLoadBalancersClient creates a new instance of LoadBalancersClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadBalancersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates a load balancer in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - loadBalancerName - The name of the load balancer.
//   - parameters - The load balancer to create or update.
//   - options - LoadBalancersClientCreateOrUpdateOptions contains the optional parameters for the LoadBalancersClient.CreateOrUpdate
//     method.
func (client *LoadBalancersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientCreateOrUpdateOptions) (LoadBalancersClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "LoadBalancersClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, loadBalancerName, parameters, options)
	if err != nil {
		return LoadBalancersClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *LoadBalancersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, parameters LoadBalancer, options *LoadBalancersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/loadBalancers/{loadBalancerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *LoadBalancersClient) createOrUpdateHandleResponse(resp *http.Response) (LoadBalancersClientCreateOrUpdateResponse, error) {
	result := LoadBalancersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a load balancer in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancersClientBeginDeleteOptions contains the optional parameters for the LoadBalancersClient.BeginDelete
//     method.
func (client *LoadBalancersClient) BeginDelete(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*runtime.Poller[LoadBalancersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, resourceName, loadBalancerName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LoadBalancersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a load balancer in the specified managed cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
func (client *LoadBalancersClient) deleteOperation(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, loadBalancerName, options)
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
func (client *LoadBalancersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, options *LoadBalancersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/loadBalancers/{loadBalancerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified load balancer.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancersClientGetOptions contains the optional parameters for the LoadBalancersClient.Get method.
func (client *LoadBalancersClient) Get(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (LoadBalancersClientGetResponse, error) {
	var err error
	const operationName = "LoadBalancersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, loadBalancerName, options)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LoadBalancersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LoadBalancersClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, loadBalancerName string, options *LoadBalancersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/loadBalancers/{loadBalancerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LoadBalancersClient) getHandleResponse(resp *http.Response) (LoadBalancersClientGetResponse, error) {
	result := LoadBalancersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancer); err != nil {
		return LoadBalancersClientGetResponse{}, err
	}
	return result, nil
}

// NewListByManagedClusterPager - Gets a list of load balancers in the specified managed cluster.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - options - LoadBalancersClientListByManagedClusterOptions contains the optional parameters for the LoadBalancersClient.NewListByManagedClusterPager
//     method.
func (client *LoadBalancersClient) NewListByManagedClusterPager(resourceGroupName string, resourceName string, options *LoadBalancersClientListByManagedClusterOptions) *runtime.Pager[LoadBalancersClientListByManagedClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancersClientListByManagedClusterResponse]{
		More: func(page LoadBalancersClientListByManagedClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancersClientListByManagedClusterResponse) (LoadBalancersClientListByManagedClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadBalancersClient.NewListByManagedClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByManagedClusterCreateRequest(ctx, resourceGroupName, resourceName, options)
			}, nil)
			if err != nil {
				return LoadBalancersClientListByManagedClusterResponse{}, err
			}
			return client.listByManagedClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByManagedClusterCreateRequest creates the ListByManagedCluster request.
func (client *LoadBalancersClient) listByManagedClusterCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *LoadBalancersClientListByManagedClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/loadBalancers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByManagedClusterHandleResponse handles the ListByManagedCluster response.
func (client *LoadBalancersClient) listByManagedClusterHandleResponse(resp *http.Response) (LoadBalancersClientListByManagedClusterResponse, error) {
	result := LoadBalancersClientListByManagedClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LoadBalancerListResult); err != nil {
		return LoadBalancersClientListByManagedClusterResponse{}, err
	}
	return result, nil
}

// BeginRebalance - Rebalance nodes across specific load balancers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the managed cluster resource.
//   - parameters - The names of the load balancers to be rebalanced. If set to empty, all load balancers will be rebalanced.
//   - options - LoadBalancersClientBeginRebalanceOptions contains the optional parameters for the LoadBalancersClient.BeginRebalance
//     method.
func (client *LoadBalancersClient) BeginRebalance(ctx context.Context, resourceGroupName string, resourceName string, parameters RebalanceLoadBalancersRequestBody, options *LoadBalancersClientBeginRebalanceOptions) (*runtime.Poller[LoadBalancersClientRebalanceResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.rebalance(ctx, resourceGroupName, resourceName, parameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[LoadBalancersClientRebalanceResponse]{
			Tracer: client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[LoadBalancersClientRebalanceResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Rebalance - Rebalance nodes across specific load balancers.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-02-preview
func (client *LoadBalancersClient) rebalance(ctx context.Context, resourceGroupName string, resourceName string, parameters RebalanceLoadBalancersRequestBody, options *LoadBalancersClientBeginRebalanceOptions) (*http.Response, error) {
	var err error
	const operationName = "LoadBalancersClient.BeginRebalance"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.rebalanceCreateRequest(ctx, resourceGroupName, resourceName, parameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// rebalanceCreateRequest creates the Rebalance request.
func (client *LoadBalancersClient) rebalanceCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, parameters RebalanceLoadBalancersRequestBody, options *LoadBalancersClientBeginRebalanceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/managedClusters/{resourceName}/rebalanceLoadBalancers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
