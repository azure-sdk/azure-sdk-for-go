//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armadvisor

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

// WorkloadsClient contains the methods for the Workloads group.
// Don't use this type directly, use NewWorkloadsClient() instead.
type WorkloadsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkloadsClient creates a new instance of WorkloadsClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkloadsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkloadsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkloadsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get list of Workloads.
//
// Generated from API version 2023-09-01-preview
//   - options - WorkloadsClientListOptions contains the optional parameters for the WorkloadsClient.NewListPager method.
func (client *WorkloadsClient) NewListPager(options *WorkloadsClientListOptions) *runtime.Pager[WorkloadsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkloadsClientListResponse]{
		More: func(page WorkloadsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkloadsClientListResponse) (WorkloadsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkloadsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return WorkloadsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *WorkloadsClient) listCreateRequest(ctx context.Context, options *WorkloadsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Advisor/workloads"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *WorkloadsClient) listHandleResponse(resp *http.Response) (WorkloadsClientListResponse, error) {
	result := WorkloadsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkloadListResult); err != nil {
		return WorkloadsClientListResponse{}, err
	}
	return result, nil
}
