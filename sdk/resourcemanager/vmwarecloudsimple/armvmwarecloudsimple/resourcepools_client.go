// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armvmwarecloudsimple

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

// ResourcePoolsClient contains the methods for the ResourcePools group.
// Don't use this type directly, use NewResourcePoolsClient() instead.
type ResourcePoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourcePoolsClient creates a new instance of ResourcePoolsClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourcePoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourcePoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourcePoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns resource pool templates by its name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-04-01
//   - regionID - The region Id (westus, eastus)
//   - pcName - The private cloud name
//   - resourcePoolName - resource pool id (vsphereId)
//   - options - ResourcePoolsClientGetOptions contains the optional parameters for the ResourcePoolsClient.Get method.
func (client *ResourcePoolsClient) Get(ctx context.Context, regionID string, pcName string, resourcePoolName string, options *ResourcePoolsClientGetOptions) (ResourcePoolsClientGetResponse, error) {
	var err error
	const operationName = "ResourcePoolsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, regionID, pcName, resourcePoolName, options)
	if err != nil {
		return ResourcePoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ResourcePoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ResourcePoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ResourcePoolsClient) getCreateRequest(ctx context.Context, regionID string, pcName string, resourcePoolName string, _ *ResourcePoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools/{resourcePoolName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	if resourcePoolName == "" {
		return nil, errors.New("parameter resourcePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourcePoolName}", url.PathEscape(resourcePoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ResourcePoolsClient) getHandleResponse(resp *http.Response) (ResourcePoolsClientGetResponse, error) {
	result := ResourcePoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcePool); err != nil {
		return ResourcePoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns list of resource pools in region for private cloud
//
// Generated from API version 2019-04-01
//   - regionID - The region Id (westus, eastus)
//   - pcName - The private cloud name
//   - options - ResourcePoolsClientListOptions contains the optional parameters for the ResourcePoolsClient.NewListPager method.
func (client *ResourcePoolsClient) NewListPager(regionID string, pcName string, options *ResourcePoolsClientListOptions) *runtime.Pager[ResourcePoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourcePoolsClientListResponse]{
		More: func(page ResourcePoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourcePoolsClientListResponse) (ResourcePoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourcePoolsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, regionID, pcName, options)
			}, nil)
			if err != nil {
				return ResourcePoolsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ResourcePoolsClient) listCreateRequest(ctx context.Context, regionID string, pcName string, _ *ResourcePoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.VMwareCloudSimple/locations/{regionId}/privateClouds/{pcName}/resourcePools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if regionID == "" {
		return nil, errors.New("parameter regionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{regionId}", url.PathEscape(regionID))
	if pcName == "" {
		return nil, errors.New("parameter pcName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{pcName}", url.PathEscape(pcName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourcePoolsClient) listHandleResponse(resp *http.Response) (ResourcePoolsClientListResponse, error) {
	result := ResourcePoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourcePoolsListResponse); err != nil {
		return ResourcePoolsClientListResponse{}, err
	}
	return result, nil
}
