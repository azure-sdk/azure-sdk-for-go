//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnewrelicobservability

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

// PlansClient contains the methods for the Plans group.
// Don't use this type directly, use NewPlansClient() instead.
type PlansClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPlansClient creates a new instance of PlansClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPlansClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PlansClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PlansClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List plans data
//
// Generated from API version 2023-11-01-preview
//   - options - PlansClientListOptions contains the optional parameters for the PlansClient.NewListPager method.
func (client *PlansClient) NewListPager(options *PlansClientListOptions) *runtime.Pager[PlansClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PlansClientListResponse]{
		More: func(page PlansClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PlansClientListResponse) (PlansClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "PlansClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return PlansClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *PlansClient) listCreateRequest(ctx context.Context, options *PlansClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/NewRelic.Observability/plans"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	if options != nil && options.AccountID != nil {
		reqQP.Set("accountId", *options.AccountID)
	}
	if options != nil && options.OrganizationID != nil {
		reqQP.Set("organizationId", *options.OrganizationID)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PlansClient) listHandleResponse(resp *http.Response) (PlansClientListResponse, error) {
	result := PlansClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PlanDataListResponse); err != nil {
		return PlansClientListResponse{}, err
	}
	return result, nil
}
