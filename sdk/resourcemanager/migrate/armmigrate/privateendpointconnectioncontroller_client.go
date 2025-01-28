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
	"strconv"
	"strings"
)

// PrivateEndpointConnectionControllerClient contains the methods for the PrivateEndpointConnectionController group.
// Don't use this type directly, use NewPrivateEndpointConnectionControllerClient() instead.
type PrivateEndpointConnectionControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionControllerClient creates a new instance of PrivateEndpointConnectionControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Gets the private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - peConnectionName - Private link resource name.
//   - body - Resource create parameters.
//   - options - PrivateEndpointConnectionControllerClientCreateOptions contains the optional parameters for the PrivateEndpointConnectionControllerClient.Create
//     method.
func (client *PrivateEndpointConnectionControllerClient) Create(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, body PrivateEndpointConnection, options *PrivateEndpointConnectionControllerClientCreateOptions) (PrivateEndpointConnectionControllerClientCreateResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionControllerClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, siteName, peConnectionName, body, options)
	if err != nil {
		return PrivateEndpointConnectionControllerClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionControllerClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionControllerClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *PrivateEndpointConnectionControllerClient) createCreateRequest(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, body PrivateEndpointConnection, options *PrivateEndpointConnectionControllerClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/privateEndpointConnections/{peConnectionName}"
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
	if peConnectionName == "" {
		return nil, errors.New("parameter peConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peConnectionName}", url.PathEscape(peConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *PrivateEndpointConnectionControllerClient) createHandleResponse(resp *http.Response) (PrivateEndpointConnectionControllerClientCreateResponse, error) {
	result := PrivateEndpointConnectionControllerClientCreateResponse{}
	if val := resp.Header.Get("Retry-After"); val != "" {
		retryAfter32, err := strconv.ParseInt(val, 10, 32)
		retryAfter := int32(retryAfter32)
		if err != nil {
			return PrivateEndpointConnectionControllerClientCreateResponse{}, err
		}
		result.RetryAfter = &retryAfter
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionControllerClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - peConnectionName - Private link resource name.
//   - options - PrivateEndpointConnectionControllerClientDeleteOptions contains the optional parameters for the PrivateEndpointConnectionControllerClient.Delete
//     method.
func (client *PrivateEndpointConnectionControllerClient) Delete(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, options *PrivateEndpointConnectionControllerClientDeleteOptions) (PrivateEndpointConnectionControllerClientDeleteResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionControllerClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, siteName, peConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionControllerClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionControllerClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionControllerClientDeleteResponse{}, err
	}
	return PrivateEndpointConnectionControllerClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PrivateEndpointConnectionControllerClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, options *PrivateEndpointConnectionControllerClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/privateEndpointConnections/{peConnectionName}"
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
	if peConnectionName == "" {
		return nil, errors.New("parameter peConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peConnectionName}", url.PathEscape(peConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - peConnectionName - Private link resource name.
//   - options - PrivateEndpointConnectionControllerClientGetOptions contains the optional parameters for the PrivateEndpointConnectionControllerClient.Get
//     method.
func (client *PrivateEndpointConnectionControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, options *PrivateEndpointConnectionControllerClientGetOptions) (PrivateEndpointConnectionControllerClientGetResponse, error) {
	var err error
	const operationName = "PrivateEndpointConnectionControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, peConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateEndpointConnectionControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, peConnectionName string, options *PrivateEndpointConnectionControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/privateEndpointConnections/{peConnectionName}"
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
	if peConnectionName == "" {
		return nil, errors.New("parameter peConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peConnectionName}", url.PathEscape(peConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionControllerClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionControllerClientGetResponse, error) {
	result := PrivateEndpointConnectionControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnection); err != nil {
		return PrivateEndpointConnectionControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMasterSitePager - Gets the private link resource.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - PrivateEndpointConnectionControllerClientListByMasterSiteOptions contains the optional parameters for the PrivateEndpointConnectionControllerClient.NewListByMasterSitePager
//     method.
func (client *PrivateEndpointConnectionControllerClient) NewListByMasterSitePager(resourceGroupName string, siteName string, options *PrivateEndpointConnectionControllerClientListByMasterSiteOptions) *runtime.Pager[PrivateEndpointConnectionControllerClientListByMasterSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionControllerClientListByMasterSiteResponse]{
		More: func(page PrivateEndpointConnectionControllerClientListByMasterSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionControllerClientListByMasterSiteResponse) (PrivateEndpointConnectionControllerClientListByMasterSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateEndpointConnectionControllerClient.NewListByMasterSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMasterSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return PrivateEndpointConnectionControllerClientListByMasterSiteResponse{}, err
			}
			return client.listByMasterSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMasterSiteCreateRequest creates the ListByMasterSite request.
func (client *PrivateEndpointConnectionControllerClient) listByMasterSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *PrivateEndpointConnectionControllerClientListByMasterSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/privateEndpointConnections"
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
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMasterSiteHandleResponse handles the ListByMasterSite response.
func (client *PrivateEndpointConnectionControllerClient) listByMasterSiteHandleResponse(resp *http.Response) (PrivateEndpointConnectionControllerClientListByMasterSiteResponse, error) {
	result := PrivateEndpointConnectionControllerClientListByMasterSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionListResult); err != nil {
		return PrivateEndpointConnectionControllerClientListByMasterSiteResponse{}, err
	}
	return result, nil
}
