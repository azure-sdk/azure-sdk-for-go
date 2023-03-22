//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// HealthReportsClient contains the methods for the HealthReports group.
// Don't use this type directly, use NewHealthReportsClient() instead.
type HealthReportsClient struct {
	internal *arm.Client
}

// NewHealthReportsClient creates a new instance of HealthReportsClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHealthReportsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*HealthReportsClient, error) {
	cl, err := arm.NewClient(moduleName+".HealthReportsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HealthReportsClient{
		internal: cl,
	}
	return client, nil
}

// NewListPager - Get a list of all health reports inside a scope. Valid scopes are: subscription (format: 'subscriptions/{subscriptionId}'),
// or security connector (format:
// 'subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName})'
//
// Generated from API version 2023-02-01-preview
//   - scope - The scope at which the operation is performed.
//   - options - HealthReportsClientListOptions contains the optional parameters for the HealthReportsClient.NewListPager method.
func (client *HealthReportsClient) NewListPager(scope string, options *HealthReportsClientListOptions) *runtime.Pager[HealthReportsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HealthReportsClientListResponse]{
		More: func(page HealthReportsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HealthReportsClientListResponse) (HealthReportsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return HealthReportsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return HealthReportsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return HealthReportsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *HealthReportsClient) listCreateRequest(ctx context.Context, scope string, options *HealthReportsClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/healthReports"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HealthReportsClient) listHandleResponse(resp *http.Response) (HealthReportsClientListResponse, error) {
	result := HealthReportsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HealthReportsList); err != nil {
		return HealthReportsClientListResponse{}, err
	}
	return result, nil
}
