//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcdn

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

// ResourceUsageClient contains the methods for the ResourceUsage group.
// Don't use this type directly, use NewResourceUsageClient() instead.
type ResourceUsageClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewResourceUsageClient creates a new instance of ResourceUsageClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewResourceUsageClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ResourceUsageClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ResourceUsageClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Check the quota and actual usage of the CDN profiles under the given subscription.
//
// Generated from API version 2024-09-01
//   - options - ResourceUsageClientListOptions contains the optional parameters for the ResourceUsageClient.NewListPager method.
func (client *ResourceUsageClient) NewListPager(options *ResourceUsageClientListOptions) *runtime.Pager[ResourceUsageClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ResourceUsageClientListResponse]{
		More: func(page ResourceUsageClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ResourceUsageClientListResponse) (ResourceUsageClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ResourceUsageClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ResourceUsageClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ResourceUsageClient) listCreateRequest(ctx context.Context, options *ResourceUsageClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cdn/checkResourceUsage"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ResourceUsageClient) listHandleResponse(resp *http.Response) (ResourceUsageClientListResponse, error) {
	result := ResourceUsageClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceUsageListResult); err != nil {
		return ResourceUsageClientListResponse{}, err
	}
	return result, nil
}
