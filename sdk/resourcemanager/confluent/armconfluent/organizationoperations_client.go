//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconfluent

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OrganizationOperationsClient contains the methods for the OrganizationOperations group.
// Don't use this type directly, use NewOrganizationOperationsClient() instead.
type OrganizationOperationsClient struct {
	internal *arm.Client
}

// NewOrganizationOperationsClient creates a new instance of OrganizationOperationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOrganizationOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OrganizationOperationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OrganizationOperationsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - List all operations provided by Microsoft.Confluent.
//
// Generated from API version 2024-01-19
//   - options - OrganizationOperationsClientListOptions contains the optional parameters for the OrganizationOperationsClient.NewListPager
//     method.
func (client *OrganizationOperationsClient) NewListPager(options *OrganizationOperationsClientListOptions) *runtime.Pager[OrganizationOperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OrganizationOperationsClientListResponse]{
		More: func(page OrganizationOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OrganizationOperationsClientListResponse) (OrganizationOperationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "OrganizationOperationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return OrganizationOperationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *OrganizationOperationsClient) listCreateRequest(ctx context.Context, options *OrganizationOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Confluent/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-19")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OrganizationOperationsClient) listHandleResponse(resp *http.Response) (OrganizationOperationsClientListResponse, error) {
	result := OrganizationOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationListResult); err != nil {
		return OrganizationOperationsClientListResponse{}, err
	}
	return result, nil
}
