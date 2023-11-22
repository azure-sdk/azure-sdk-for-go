//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// TestTypesClient contains the methods for the TestTypes group.
// Don't use this type directly, use NewTestTypesClient() instead.
type TestTypesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTestTypesClient creates a new instance of TestTypesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTestTypesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TestTypesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TestTypesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a test type of a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - testTypeResourceName - The resource name of a test type.
//   - options - TestTypesClientGetOptions contains the optional parameters for the TestTypesClient.Get method.
func (client *TestTypesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, testTypeResourceName string, options *TestTypesClientGetOptions) (TestTypesClientGetResponse, error) {
	var err error
	const operationName = "TestTypesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, testTypeResourceName, options)
	if err != nil {
		return TestTypesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TestTypesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TestTypesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TestTypesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, testTypeResourceName string, options *TestTypesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/testTypes/{testTypeResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if testTypeResourceName == "" {
		return nil, errors.New("parameter testTypeResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testTypeResourceName}", url.PathEscape(testTypeResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TestTypesClient) getHandleResponse(resp *http.Response) (TestTypesClientGetResponse, error) {
	result := TestTypesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestTypeResource); err != nil {
		return TestTypesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the test types of a Test Base Account.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - TestTypesClientListOptions contains the optional parameters for the TestTypesClient.NewListPager method.
func (client *TestTypesClient) NewListPager(resourceGroupName string, testBaseAccountName string, options *TestTypesClientListOptions) *runtime.Pager[TestTypesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[TestTypesClientListResponse]{
		More: func(page TestTypesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TestTypesClientListResponse) (TestTypesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TestTypesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			}, nil)
			if err != nil {
				return TestTypesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *TestTypesClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *TestTypesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/testTypes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *TestTypesClient) listHandleResponse(resp *http.Response) (TestTypesClientListResponse, error) {
	result := TestTypesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TestTypeListResult); err != nil {
		return TestTypesClientListResponse{}, err
	}
	return result, nil
}
