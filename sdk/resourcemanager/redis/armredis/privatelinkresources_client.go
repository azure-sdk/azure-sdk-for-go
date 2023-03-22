//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredis

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

// PrivateLinkResourcesClient contains the methods for the PrivateLinkResources group.
// Don't use this type directly, use NewPrivateLinkResourcesClient() instead.
type PrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateLinkResourcesClient creates a new instance of PrivateLinkResourcesClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateLinkResourcesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByRedisCachePager - Gets the private link resources that need to be created for a redis cache.
//
// Generated from API version 2022-06-01-privatepreview
//   - resourceGroupName - The name of the resource group.
//   - cacheName - The name of the Redis cache.
//   - options - PrivateLinkResourcesClientListByRedisCacheOptions contains the optional parameters for the PrivateLinkResourcesClient.NewListByRedisCachePager
//     method.
func (client *PrivateLinkResourcesClient) NewListByRedisCachePager(resourceGroupName string, cacheName string, options *PrivateLinkResourcesClientListByRedisCacheOptions) *runtime.Pager[PrivateLinkResourcesClientListByRedisCacheResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateLinkResourcesClientListByRedisCacheResponse]{
		More: func(page PrivateLinkResourcesClientListByRedisCacheResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *PrivateLinkResourcesClientListByRedisCacheResponse) (PrivateLinkResourcesClientListByRedisCacheResponse, error) {
			req, err := client.listByRedisCacheCreateRequest(ctx, resourceGroupName, cacheName, options)
			if err != nil {
				return PrivateLinkResourcesClientListByRedisCacheResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateLinkResourcesClientListByRedisCacheResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateLinkResourcesClientListByRedisCacheResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByRedisCacheHandleResponse(resp)
		},
	})
}

// listByRedisCacheCreateRequest creates the ListByRedisCache request.
func (client *PrivateLinkResourcesClient) listByRedisCacheCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *PrivateLinkResourcesClientListByRedisCacheOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/privateLinkResources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-privatepreview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByRedisCacheHandleResponse handles the ListByRedisCache response.
func (client *PrivateLinkResourcesClient) listByRedisCacheHandleResponse(resp *http.Response) (PrivateLinkResourcesClientListByRedisCacheResponse, error) {
	result := PrivateLinkResourcesClientListByRedisCacheResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return PrivateLinkResourcesClientListByRedisCacheResponse{}, err
	}
	return result, nil
}
