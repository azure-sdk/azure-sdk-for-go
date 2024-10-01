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

// PrivateLinkResourcesControllerClient contains the methods for the PrivateLinkResourcesController group.
// Don't use this type directly, use NewPrivateLinkResourcesControllerClient() instead.
type PrivateLinkResourcesControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesControllerClient creates a new instance of PrivateLinkResourcesControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the private link resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - privateLinkResourceName - Private link resource name.
//   - options - PrivateLinkResourcesControllerClientGetOptions contains the optional parameters for the PrivateLinkResourcesControllerClient.Get
//     method.
func (client *PrivateLinkResourcesControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, privateLinkResourceName string, options *PrivateLinkResourcesControllerClientGetOptions) (PrivateLinkResourcesControllerClientGetResponse, error) {
	var err error
	const operationName = "PrivateLinkResourcesControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, privateLinkResourceName, options)
	if err != nil {
		return PrivateLinkResourcesControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateLinkResourcesControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PrivateLinkResourcesControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PrivateLinkResourcesControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, privateLinkResourceName string, options *PrivateLinkResourcesControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/privateLinkResources/{privateLinkResourceName}"
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
	if privateLinkResourceName == "" {
		return nil, errors.New("parameter privateLinkResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkResourceName}", url.PathEscape(privateLinkResourceName))
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
func (client *PrivateLinkResourcesControllerClient) getHandleResponse(resp *http.Response) (PrivateLinkResourcesControllerClientGetResponse, error) {
	result := PrivateLinkResourcesControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResource); err != nil {
		return PrivateLinkResourcesControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMasterSitePager - Gets the private link resource.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - PrivateLinkResourcesControllerClientListByMasterSiteOptions contains the optional parameters for the PrivateLinkResourcesControllerClient.NewListByMasterSitePager
//     method.
func (client *PrivateLinkResourcesControllerClient) NewListByMasterSitePager(resourceGroupName string, siteName string, options *PrivateLinkResourcesControllerClientListByMasterSiteOptions) *runtime.Pager[PrivateLinkResourcesControllerClientListByMasterSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesControllerClientListByMasterSiteResponse]{
		More: func(page PrivateLinkResourcesControllerClientListByMasterSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesControllerClientListByMasterSiteResponse) (PrivateLinkResourcesControllerClientListByMasterSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PrivateLinkResourcesControllerClient.NewListByMasterSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByMasterSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return PrivateLinkResourcesControllerClientListByMasterSiteResponse{}, err
			}
			return client.listByMasterSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByMasterSiteCreateRequest creates the ListByMasterSite request.
func (client *PrivateLinkResourcesControllerClient) listByMasterSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *PrivateLinkResourcesControllerClientListByMasterSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/masterSites/{siteName}/privateLinkResources"
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
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMasterSiteHandleResponse handles the ListByMasterSite response.
func (client *PrivateLinkResourcesControllerClient) listByMasterSiteHandleResponse(resp *http.Response) (PrivateLinkResourcesControllerClientListByMasterSiteResponse, error) {
	result := PrivateLinkResourcesControllerClientListByMasterSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesControllerClientListByMasterSiteResponse{}, err
	}
	return result, nil
}
