//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhelp

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// DiscoverySolutionClient contains the methods for the DiscoverySolution group.
// Don't use this type directly, use NewDiscoverySolutionClient() instead.
type DiscoverySolutionClient struct {
	internal  *arm.Client
	scope     string
	filter    *string
	skiptoken *string
}

// NewDiscoverySolutionClient creates a new instance of DiscoverySolutionClient with the specified values.
//   - scope - This is an extension resource provider and only resource level extension is supported at the moment.
//   - filter - Can be used to filter solutionIds by 'ProblemClassificationId'. The filter supports only 'and' and 'eq' operators.
//     Example: $filter=ProblemClassificationId eq '1ddda5b4-cf6c-4d4f-91ad-bc38ab0e811e'
//     and ProblemClassificationId eq '0a9673c2-7af6-4e19-90d3-4ee2461076d9'.
//   - skiptoken - Skiptoken is only used if a previous operation returned a partial result. If a previous response contains a
//     nextLink element, the value of the nextLink element will include a skiptoken parameter that
//     specifies a starting point to use for subsequent calls.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDiscoverySolutionClient(scope string, filter *string, skiptoken *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DiscoverySolutionClient, error) {
	cl, err := arm.NewClient(moduleName+".DiscoverySolutionClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DiscoverySolutionClient{
		scope:     scope,
		filter:    filter,
		skiptoken: skiptoken,
		internal:  cl,
	}
	return client, nil
}

// NewListPager - Solutions Discovery is the initial point of entry within Help API, which helps you identify the relevant
// solutions for your Azure issue.
// You can discover solutions using resourceUri OR resourceUri + problemClassificationId.
// We will do our best in returning relevant diagnostics for your Azure issue.
// Get the problemClassificationId(s) using this reference [https://learn.microsoft.com/en-us/rest/api/support/problem-classifications/list?tabs=HTTP].
// Note: ‘requiredParameterSets’ from Solutions Discovery API response must be passed via ‘additionalParameters’ as an input
// to Diagnostics API.
//
// Generated from API version 2023-01-01-preview
//   - options - DiscoverySolutionClientListOptions contains the optional parameters for the DiscoverySolutionClient.NewListPager
//     method.
func (client *DiscoverySolutionClient) NewListPager(options *DiscoverySolutionClientListOptions) *runtime.Pager[DiscoverySolutionClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DiscoverySolutionClientListResponse]{
		More: func(page DiscoverySolutionClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DiscoverySolutionClientListResponse) (DiscoverySolutionClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DiscoverySolutionClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DiscoverySolutionClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DiscoverySolutionClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DiscoverySolutionClient) listCreateRequest(ctx context.Context, options *DiscoverySolutionClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Help/discoverySolutions"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", client.scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	if client.skiptoken != nil {
		reqQP.Set("$skiptoken", *client.skiptoken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if client.filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*client.filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DiscoverySolutionClient) listHandleResponse(resp *http.Response) (DiscoverySolutionClientListResponse, error) {
	result := DiscoverySolutionClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DiscoveryResponse); err != nil {
		return DiscoverySolutionClientListResponse{}, err
	}
	return result, nil
}
