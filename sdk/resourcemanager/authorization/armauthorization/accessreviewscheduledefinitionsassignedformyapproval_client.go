// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// AccessReviewScheduleDefinitionsAssignedForMyApprovalClient contains the methods for the AccessReviewScheduleDefinitionsAssignedForMyApproval
// group.
// Don't use this type directly, use NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient() instead.
type AccessReviewScheduleDefinitionsAssignedForMyApprovalClient struct {
	internal *arm.Client
}

// NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient creates a new instance of AccessReviewScheduleDefinitionsAssignedForMyApprovalClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAccessReviewScheduleDefinitionsAssignedForMyApprovalClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AccessReviewScheduleDefinitionsAssignedForMyApprovalClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AccessReviewScheduleDefinitionsAssignedForMyApprovalClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Get access review instances assigned for my approval.
//
// Generated from API version 2021-12-01-preview
//   - options - AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListOptions contains the optional parameters for the
//     AccessReviewScheduleDefinitionsAssignedForMyApprovalClient.NewListPager method.
func (client *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient) NewListPager(options *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListOptions) *runtime.Pager[AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse]{
		More: func(page AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse) (AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AccessReviewScheduleDefinitionsAssignedForMyApprovalClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient) listCreateRequest(ctx context.Context, options *AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	unencodedParams := []string{req.Raw().URL.RawQuery}
	if options != nil && options.Filter != nil {
		unencodedParams = append(unencodedParams, "$filter="+*options.Filter)
	}
	req.Raw().URL.RawQuery = strings.Join(unencodedParams, "&")
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AccessReviewScheduleDefinitionsAssignedForMyApprovalClient) listHandleResponse(resp *http.Response) (AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse, error) {
	result := AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewScheduleDefinitionListResult); err != nil {
		return AccessReviewScheduleDefinitionsAssignedForMyApprovalClientListResponse{}, err
	}
	return result, nil
}
