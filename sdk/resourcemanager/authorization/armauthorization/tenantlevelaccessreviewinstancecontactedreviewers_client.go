//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armauthorization

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

// TenantLevelAccessReviewInstanceContactedReviewersClient contains the methods for the TenantLevelAccessReviewInstanceContactedReviewers group.
// Don't use this type directly, use NewTenantLevelAccessReviewInstanceContactedReviewersClient() instead.
type TenantLevelAccessReviewInstanceContactedReviewersClient struct {
	internal *arm.Client
}

// NewTenantLevelAccessReviewInstanceContactedReviewersClient creates a new instance of TenantLevelAccessReviewInstanceContactedReviewersClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTenantLevelAccessReviewInstanceContactedReviewersClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantLevelAccessReviewInstanceContactedReviewersClient, error) {
	cl, err := arm.NewClient(moduleName+".TenantLevelAccessReviewInstanceContactedReviewersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TenantLevelAccessReviewInstanceContactedReviewersClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Get access review instance contacted reviewers
//
// Generated from API version 2021-12-01-preview
//   - scheduleDefinitionID - The id of the access review schedule definition.
//   - id - The id of the access review instance.
//   - options - TenantLevelAccessReviewInstanceContactedReviewersClientListOptions contains the optional parameters for the TenantLevelAccessReviewInstanceContactedReviewersClient.NewListPager
//     method.
func (client *TenantLevelAccessReviewInstanceContactedReviewersClient) NewListPager(scheduleDefinitionID string, id string, options *TenantLevelAccessReviewInstanceContactedReviewersClientListOptions) *runtime.Pager[TenantLevelAccessReviewInstanceContactedReviewersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TenantLevelAccessReviewInstanceContactedReviewersClientListResponse]{
		More: func(page TenantLevelAccessReviewInstanceContactedReviewersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TenantLevelAccessReviewInstanceContactedReviewersClientListResponse) (TenantLevelAccessReviewInstanceContactedReviewersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scheduleDefinitionID, id, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return TenantLevelAccessReviewInstanceContactedReviewersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return TenantLevelAccessReviewInstanceContactedReviewersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return TenantLevelAccessReviewInstanceContactedReviewersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *TenantLevelAccessReviewInstanceContactedReviewersClient) listCreateRequest(ctx context.Context, scheduleDefinitionID string, id string, options *TenantLevelAccessReviewInstanceContactedReviewersClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Authorization/accessReviewScheduleDefinitions/{scheduleDefinitionId}/instances/{id}/contactedReviewers"
	if scheduleDefinitionID == "" {
		return nil, errors.New("parameter scheduleDefinitionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scheduleDefinitionId}", url.PathEscape(scheduleDefinitionID))
	if id == "" {
		return nil, errors.New("parameter id cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{id}", url.PathEscape(id))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TenantLevelAccessReviewInstanceContactedReviewersClient) listHandleResponse(resp *http.Response) (TenantLevelAccessReviewInstanceContactedReviewersClientListResponse, error) {
	result := TenantLevelAccessReviewInstanceContactedReviewersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessReviewContactedReviewerListResult); err != nil {
		return TenantLevelAccessReviewInstanceContactedReviewersClientListResponse{}, err
	}
	return result, nil
}
