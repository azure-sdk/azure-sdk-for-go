// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// HypervClusterControllerClient contains the methods for the HypervClusterController group.
// Don't use this type directly, use NewHypervClusterControllerClient() instead.
type HypervClusterControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHypervClusterControllerClient creates a new instance of HypervClusterControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHypervClusterControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HypervClusterControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HypervClusterControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateCluster - Method to create or update a Hyper-V cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - clusterName - Cluster ARM name
//   - body - Resource create parameters.
//   - options - HypervClusterControllerClientBeginCreateClusterOptions contains the optional parameters for the HypervClusterControllerClient.BeginCreateCluster
//     method.
func (client *HypervClusterControllerClient) BeginCreateCluster(ctx context.Context, resourceGroupName string, siteName string, clusterName string, body HypervCluster, options *HypervClusterControllerClientBeginCreateClusterOptions) (*runtime.Poller[HypervClusterControllerClientCreateClusterResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createCluster(ctx, resourceGroupName, siteName, clusterName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HypervClusterControllerClientCreateClusterResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HypervClusterControllerClientCreateClusterResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateCluster - Method to create or update a Hyper-V cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
func (client *HypervClusterControllerClient) createCluster(ctx context.Context, resourceGroupName string, siteName string, clusterName string, body HypervCluster, options *HypervClusterControllerClientBeginCreateClusterOptions) (*http.Response, error) {
	var err error
	const operationName = "HypervClusterControllerClient.BeginCreateCluster"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createClusterCreateRequest(ctx, resourceGroupName, siteName, clusterName, body, options)
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

// createClusterCreateRequest creates the CreateCluster request.
func (client *HypervClusterControllerClient) createClusterCreateRequest(ctx context.Context, resourceGroupName string, siteName string, clusterName string, body HypervCluster, _ *HypervClusterControllerClientBeginCreateClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete - Delete a HypervCluster
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - clusterName - Cluster ARM name
//   - options - HypervClusterControllerClientDeleteOptions contains the optional parameters for the HypervClusterControllerClient.Delete
//     method.
func (client *HypervClusterControllerClient) Delete(ctx context.Context, resourceGroupName string, siteName string, clusterName string, options *HypervClusterControllerClientDeleteOptions) (HypervClusterControllerClientDeleteResponse, error) {
	var err error
	const operationName = "HypervClusterControllerClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, clusterName, options)
	if err != nil {
		return HypervClusterControllerClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HypervClusterControllerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HypervClusterControllerClientDeleteResponse{}, err
	}
	return HypervClusterControllerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HypervClusterControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, clusterName string, _ *HypervClusterControllerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// GetCluster - Method to get a Hyper-V cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - clusterName - Cluster ARM name
//   - options - HypervClusterControllerClientGetClusterOptions contains the optional parameters for the HypervClusterControllerClient.GetCluster
//     method.
func (client *HypervClusterControllerClient) GetCluster(ctx context.Context, resourceGroupName string, siteName string, clusterName string, options *HypervClusterControllerClientGetClusterOptions) (HypervClusterControllerClientGetClusterResponse, error) {
	var err error
	const operationName = "HypervClusterControllerClient.GetCluster"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getClusterCreateRequest(ctx, resourceGroupName, siteName, clusterName, options)
	if err != nil {
		return HypervClusterControllerClientGetClusterResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HypervClusterControllerClientGetClusterResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HypervClusterControllerClientGetClusterResponse{}, err
	}
	resp, err := client.getClusterHandleResponse(httpResp)
	return resp, err
}

// getClusterCreateRequest creates the GetCluster request.
func (client *HypervClusterControllerClient) getClusterCreateRequest(ctx context.Context, resourceGroupName string, siteName string, clusterName string, _ *HypervClusterControllerClientGetClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/clusters/{clusterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getClusterHandleResponse handles the GetCluster response.
func (client *HypervClusterControllerClient) getClusterHandleResponse(resp *http.Response) (HypervClusterControllerClientGetClusterResponse, error) {
	result := HypervClusterControllerClientGetClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HypervCluster); err != nil {
		return HypervClusterControllerClientGetClusterResponse{}, err
	}
	return result, nil
}

// NewListByHypervSitePager - List HypervCluster resources by HypervSite
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - HypervClusterControllerClientListByHypervSiteOptions contains the optional parameters for the HypervClusterControllerClient.NewListByHypervSitePager
//     method.
func (client *HypervClusterControllerClient) NewListByHypervSitePager(resourceGroupName string, siteName string, options *HypervClusterControllerClientListByHypervSiteOptions) *runtime.Pager[HypervClusterControllerClientListByHypervSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[HypervClusterControllerClientListByHypervSiteResponse]{
		More: func(page HypervClusterControllerClientListByHypervSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HypervClusterControllerClientListByHypervSiteResponse) (HypervClusterControllerClientListByHypervSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HypervClusterControllerClient.NewListByHypervSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHypervSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return HypervClusterControllerClientListByHypervSiteResponse{}, err
			}
			return client.listByHypervSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHypervSiteCreateRequest creates the ListByHypervSite request.
func (client *HypervClusterControllerClient) listByHypervSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *HypervClusterControllerClientListByHypervSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/clusters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHypervSiteHandleResponse handles the ListByHypervSite response.
func (client *HypervClusterControllerClient) listByHypervSiteHandleResponse(resp *http.Response) (HypervClusterControllerClientListByHypervSiteResponse, error) {
	result := HypervClusterControllerClientListByHypervSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HypervClusterListResult); err != nil {
		return HypervClusterControllerClientListByHypervSiteResponse{}, err
	}
	return result, nil
}
