//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaps

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// Client contains the methods for the Maps group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal *arm.Client
}

// NewClient creates a new instance of Client with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		internal: cl,
	}
	return client, nil
}

// NewListOperationsPager - List operations available for the Maps Resource Provider
//
// Generated from API version 2023-12-01-preview
//   - options - ClientListOperationsOptions contains the optional parameters for the Client.NewListOperationsPager method.
func (client *Client) NewListOperationsPager(options *ClientListOperationsOptions) *runtime.Pager[ClientListOperationsResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListOperationsResponse]{
		More: func(page ClientListOperationsResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListOperationsResponse) (ClientListOperationsResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.NewListOperationsPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listOperationsCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return ClientListOperationsResponse{}, err
			}
			return client.listOperationsHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listOperationsCreateRequest creates the ListOperations request.
func (client *Client) listOperationsCreateRequest(ctx context.Context, options *ClientListOperationsOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Maps/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listOperationsHandleResponse handles the ListOperations response.
func (client *Client) listOperationsHandleResponse(resp *http.Response) (ClientListOperationsResponse, error) {
	result := ClientListOperationsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return ClientListOperationsResponse{}, err
	}
	return result, nil
}
