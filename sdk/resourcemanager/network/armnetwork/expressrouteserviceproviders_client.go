//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ExpressRouteServiceProvidersClient contains the methods for the ExpressRouteServiceProviders group.
// Don't use this type directly, use NewExpressRouteServiceProvidersClient() instead.
type ExpressRouteServiceProvidersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExpressRouteServiceProvidersClient creates a new instance of ExpressRouteServiceProvidersClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExpressRouteServiceProvidersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteServiceProvidersClient, error) {
	cl, err := arm.NewClient(moduleName+".ExpressRouteServiceProvidersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteServiceProvidersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets all the available express route service providers.
//
// Generated from API version 2023-02-01
//   - options - ExpressRouteServiceProvidersClientListOptions contains the optional parameters for the ExpressRouteServiceProvidersClient.NewListPager
//     method.
func (client *ExpressRouteServiceProvidersClient) NewListPager(options *ExpressRouteServiceProvidersClientListOptions) *runtime.Pager[ExpressRouteServiceProvidersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRouteServiceProvidersClientListResponse]{
		More: func(page ExpressRouteServiceProvidersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRouteServiceProvidersClientListResponse) (ExpressRouteServiceProvidersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ExpressRouteServiceProvidersClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRouteServiceProvidersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ExpressRouteServiceProvidersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRouteServiceProvidersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRouteServiceProvidersClient) listCreateRequest(ctx context.Context, options *ExpressRouteServiceProvidersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteServiceProviders"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteServiceProvidersClient) listHandleResponse(resp *http.Response) (ExpressRouteServiceProvidersClientListResponse, error) {
	result := ExpressRouteServiceProvidersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteServiceProviderListResult); err != nil {
		return ExpressRouteServiceProvidersClientListResponse{}, err
	}
	return result, nil
}
