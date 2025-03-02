// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// CompliancesClient contains the methods for the Compliances group.
// Don't use this type directly, use NewCompliancesClient() instead.
type CompliancesClient struct {
	internal *arm.Client
}

// NewCompliancesClient creates a new instance of CompliancesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCompliancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*CompliancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CompliancesClient{
		internal: cl,
	}
	return client, nil
}

// Get - Details of a specific Compliance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2017-08-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - complianceName - name of the Compliance
//   - options - CompliancesClientGetOptions contains the optional parameters for the CompliancesClient.Get method.
func (client *CompliancesClient) Get(ctx context.Context, scope string, complianceName string, options *CompliancesClientGetOptions) (CompliancesClientGetResponse, error) {
	var err error
	const operationName = "CompliancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, scope, complianceName, options)
	if err != nil {
		return CompliancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CompliancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CompliancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CompliancesClient) getCreateRequest(ctx context.Context, scope string, complianceName string, _ *CompliancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/compliances/{complianceName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if complianceName == "" {
		return nil, errors.New("parameter complianceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{complianceName}", url.PathEscape(complianceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CompliancesClient) getHandleResponse(resp *http.Response) (CompliancesClientGetResponse, error) {
	result := CompliancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Compliance); err != nil {
		return CompliancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - The Compliance scores of the specific management group.
//
// Generated from API version 2017-08-01-preview
//   - scope - Scope of the query, can be subscription (/subscriptions/0b06d9ea-afe6-4779-bd59-30e5c2d9d13f) or management group
//     (/providers/Microsoft.Management/managementGroups/mgName).
//   - options - CompliancesClientListOptions contains the optional parameters for the CompliancesClient.NewListPager method.
func (client *CompliancesClient) NewListPager(scope string, options *CompliancesClientListOptions) *runtime.Pager[CompliancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CompliancesClientListResponse]{
		More: func(page CompliancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CompliancesClientListResponse) (CompliancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CompliancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, scope, options)
			}, nil)
			if err != nil {
				return CompliancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CompliancesClient) listCreateRequest(ctx context.Context, scope string, _ *CompliancesClientListOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Security/compliances"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2017-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CompliancesClient) listHandleResponse(resp *http.Response) (CompliancesClientListResponse, error) {
	result := CompliancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComplianceList); err != nil {
		return CompliancesClientListResponse{}, err
	}
	return result, nil
}
