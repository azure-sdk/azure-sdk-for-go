//go:build go1.18
// +build go1.18

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

// HypervHostControllerClient contains the methods for the HypervHostController group.
// Don't use this type directly, use NewHypervHostControllerClient() instead.
type HypervHostControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHypervHostControllerClient creates a new instance of HypervHostControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHypervHostControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HypervHostControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HypervHostControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Create a HypervHost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - hostName - Host name
//   - body - Resource create parameters.
//   - options - HypervHostControllerClientBeginCreateOptions contains the optional parameters for the HypervHostControllerClient.BeginCreate
//     method.
func (client *HypervHostControllerClient) BeginCreate(ctx context.Context, resourceGroupName string, siteName string, hostName string, body HypervHost, options *HypervHostControllerClientBeginCreateOptions) (*runtime.Poller[HypervHostControllerClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, siteName, hostName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HypervHostControllerClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HypervHostControllerClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Create a HypervHost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *HypervHostControllerClient) create(ctx context.Context, resourceGroupName string, siteName string, hostName string, body HypervHost, options *HypervHostControllerClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "HypervHostControllerClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, siteName, hostName, body, options)
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

// createCreateRequest creates the Create request.
func (client *HypervHostControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, siteName string, hostName string, body HypervHost, options *HypervHostControllerClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/hosts/{hostName}"
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
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Delete a HypervHost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - hostName - Host name
//   - options - HypervHostControllerClientBeginDeleteOptions contains the optional parameters for the HypervHostControllerClient.BeginDelete
//     method.
func (client *HypervHostControllerClient) BeginDelete(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *HypervHostControllerClientBeginDeleteOptions) (*runtime.Poller[HypervHostControllerClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, siteName, hostName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[HypervHostControllerClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[HypervHostControllerClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Delete a HypervHost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *HypervHostControllerClient) deleteOperation(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *HypervHostControllerClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "HypervHostControllerClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, hostName, options)
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
func (client *HypervHostControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *HypervHostControllerClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/hosts/{hostName}"
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
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a HypervHost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - hostName - Host name
//   - options - HypervHostControllerClientGetOptions contains the optional parameters for the HypervHostControllerClient.Get
//     method.
func (client *HypervHostControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *HypervHostControllerClientGetOptions) (HypervHostControllerClientGetResponse, error) {
	var err error
	const operationName = "HypervHostControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, hostName, options)
	if err != nil {
		return HypervHostControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HypervHostControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HypervHostControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HypervHostControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *HypervHostControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/hosts/{hostName}"
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
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HypervHostControllerClient) getHandleResponse(resp *http.Response) (HypervHostControllerClientGetResponse, error) {
	result := HypervHostControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HypervHost); err != nil {
		return HypervHostControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByHypervSitePager - List HypervHost resources by HypervSite
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - HypervHostControllerClientListByHypervSiteOptions contains the optional parameters for the HypervHostControllerClient.NewListByHypervSitePager
//     method.
func (client *HypervHostControllerClient) NewListByHypervSitePager(resourceGroupName string, siteName string, options *HypervHostControllerClientListByHypervSiteOptions) *runtime.Pager[HypervHostControllerClientListByHypervSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[HypervHostControllerClientListByHypervSiteResponse]{
		More: func(page HypervHostControllerClientListByHypervSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HypervHostControllerClientListByHypervSiteResponse) (HypervHostControllerClientListByHypervSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HypervHostControllerClient.NewListByHypervSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByHypervSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return HypervHostControllerClientListByHypervSiteResponse{}, err
			}
			return client.listByHypervSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByHypervSiteCreateRequest creates the ListByHypervSite request.
func (client *HypervHostControllerClient) listByHypervSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *HypervHostControllerClientListByHypervSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/hypervSites/{siteName}/hosts"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHypervSiteHandleResponse handles the ListByHypervSite response.
func (client *HypervHostControllerClient) listByHypervSiteHandleResponse(resp *http.Response) (HypervHostControllerClientListByHypervSiteResponse, error) {
	result := HypervHostControllerClientListByHypervSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HypervHostListResult); err != nil {
		return HypervHostControllerClientListByHypervSiteResponse{}, err
	}
	return result, nil
}
