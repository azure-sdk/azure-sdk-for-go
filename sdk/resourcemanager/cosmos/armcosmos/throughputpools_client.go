//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcosmos

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

// ThroughputPoolsClient contains the methods for the ThroughputPools group.
// Don't use this type directly, use NewThroughputPoolsClient() instead.
type ThroughputPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewThroughputPoolsClient creates a new instance of ThroughputPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewThroughputPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ThroughputPoolsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ThroughputPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all the Azure Cosmos DB Throughput Pools available under the subscription.
//
// Generated from API version 2024-05-15-preview
//   - options - ThroughputPoolsClientListOptions contains the optional parameters for the ThroughputPoolsClient.NewListPager
//     method.
func (client *ThroughputPoolsClient) NewListPager(options *ThroughputPoolsClientListOptions) *runtime.Pager[ThroughputPoolsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ThroughputPoolsClientListResponse]{
		More: func(page ThroughputPoolsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ThroughputPoolsClientListResponse) (ThroughputPoolsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ThroughputPoolsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ThroughputPoolsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ThroughputPoolsClient) listCreateRequest(ctx context.Context, options *ThroughputPoolsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DocumentDB/throughputPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ThroughputPoolsClient) listHandleResponse(resp *http.Response) (ThroughputPoolsClientListResponse, error) {
	result := ThroughputPoolsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThroughputPoolsListResult); err != nil {
		return ThroughputPoolsClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - List all the ThroughputPools in a given resource group.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ThroughputPoolsClientListByResourceGroupOptions contains the optional parameters for the ThroughputPoolsClient.NewListByResourceGroupPager
//     method.
func (client *ThroughputPoolsClient) NewListByResourceGroupPager(resourceGroupName string, options *ThroughputPoolsClientListByResourceGroupOptions) *runtime.Pager[ThroughputPoolsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ThroughputPoolsClientListByResourceGroupResponse]{
		More: func(page ThroughputPoolsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ThroughputPoolsClientListByResourceGroupResponse) (ThroughputPoolsClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ThroughputPoolsClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return ThroughputPoolsClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ThroughputPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ThroughputPoolsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/throughputPools"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ThroughputPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (ThroughputPoolsClientListByResourceGroupResponse, error) {
	result := ThroughputPoolsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ThroughputPoolsListResult); err != nil {
		return ThroughputPoolsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}
