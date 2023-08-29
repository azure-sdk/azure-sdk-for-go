//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdomainservices

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DomainServiceOperationsClient contains the methods for the DomainServiceOperations group.
// Don't use this type directly, use NewDomainServiceOperationsClient() instead.
type DomainServiceOperationsClient struct {
	internal *arm.Client
}

// NewDomainServiceOperationsClient creates a new instance of DomainServiceOperationsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDomainServiceOperationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*DomainServiceOperationsClient, error) {
	cl, err := arm.NewClient(moduleName+".DomainServiceOperationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DomainServiceOperationsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Lists all the available Domain Services operations.
//
// Generated from API version 2022-12-01
//   - options - DomainServiceOperationsClientListOptions contains the optional parameters for the DomainServiceOperationsClient.NewListPager
//     method.
func (client *DomainServiceOperationsClient) NewListPager(options *DomainServiceOperationsClientListOptions) *runtime.Pager[DomainServiceOperationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DomainServiceOperationsClientListResponse]{
		More: func(page DomainServiceOperationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DomainServiceOperationsClientListResponse) (DomainServiceOperationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DomainServiceOperationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DomainServiceOperationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DomainServiceOperationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DomainServiceOperationsClient) listCreateRequest(ctx context.Context, options *DomainServiceOperationsClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.AAD/operations"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DomainServiceOperationsClient) listHandleResponse(resp *http.Response) (DomainServiceOperationsClientListResponse, error) {
	result := DomainServiceOperationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OperationEntityListResult); err != nil {
		return DomainServiceOperationsClientListResponse{}, err
	}
	return result, nil
}
