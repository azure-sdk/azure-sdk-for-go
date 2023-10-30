//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armquota

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// OperationClient contains the methods for the QuotaOperation group.
// Don't use this type directly, use NewOperationClient() instead.
type OperationClient struct {
	internal *arm.Client
}

// NewOperationClient creates a new instance of OperationClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationClient, error) {
	cl, err := arm.NewClient(moduleName+".OperationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - List all the operations supported by the Microsoft.Quota resource provider.
//
// Generated from API version 2023-02-01
//   - options - OperationClientListOptions contains the optional parameters for the OperationClient.NewListPager method.
func (client *OperationClient) NewListPager(options *OperationClientListOptions) *runtime.Pager[OperationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OperationClientListResponse]{
		More: func(page OperationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OperationClientListResponse) (OperationClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OperationClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OperationClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OperationClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OperationClient) listCreateRequest(ctx context.Context, options *OperationClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Quota/operations"
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
func (client *OperationClient) listHandleResponse(resp *http.Response) (OperationClientListResponse, error) {
	result := OperationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationList); err != nil {
		return OperationClientListResponse{}, err
	}
	return result, nil
}
