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

// FreeHourBalancesClient contains the methods for the FreeHourBalances group.
// Don't use this type directly, use NewFreeHourBalancesClient() instead.
type FreeHourBalancesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFreeHourBalancesClient creates a new instance of FreeHourBalancesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFreeHourBalancesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FreeHourBalancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FreeHourBalancesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Return the Test Base free hour balance.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - freeHourBalanceName - The name of the free hour balance of a Test Base Account.
//   - options - FreeHourBalancesClientGetOptions contains the optional parameters for the FreeHourBalancesClient.Get method.
func (client *FreeHourBalancesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, freeHourBalanceName FreeHourBalanceName, options *FreeHourBalancesClientGetOptions) (FreeHourBalancesClientGetResponse, error) {
	var err error
	const operationName = "FreeHourBalancesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, freeHourBalanceName, options)
	if err != nil {
		return FreeHourBalancesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FreeHourBalancesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FreeHourBalancesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FreeHourBalancesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, freeHourBalanceName FreeHourBalanceName, _ *FreeHourBalancesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/freeHourBalances/{freeHourBalanceName}"
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
	if freeHourBalanceName == "" {
		return nil, errors.New("parameter freeHourBalanceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{freeHourBalanceName}", url.PathEscape(string(freeHourBalanceName)))
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
func (client *FreeHourBalancesClient) getHandleResponse(resp *http.Response) (FreeHourBalancesClientGetResponse, error) {
	result := FreeHourBalancesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FreeHourBalanceResource); err != nil {
		return FreeHourBalancesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Return the Test Base free hour balances list.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - FreeHourBalancesClientListOptions contains the optional parameters for the FreeHourBalancesClient.NewListPager
//     method.
func (client *FreeHourBalancesClient) NewListPager(resourceGroupName string, testBaseAccountName string, options *FreeHourBalancesClientListOptions) *runtime.Pager[FreeHourBalancesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FreeHourBalancesClientListResponse]{
		More: func(page FreeHourBalancesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FreeHourBalancesClientListResponse) (FreeHourBalancesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "FreeHourBalancesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			}, nil)
			if err != nil {
				return FreeHourBalancesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *FreeHourBalancesClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, _ *FreeHourBalancesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/freeHourBalances"
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
func (client *FreeHourBalancesClient) listHandleResponse(resp *http.Response) (FreeHourBalancesClientListResponse, error) {
	result := FreeHourBalancesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FreeHourBalancesListResult); err != nil {
		return FreeHourBalancesClientListResponse{}, err
	}
	return result, nil
}
