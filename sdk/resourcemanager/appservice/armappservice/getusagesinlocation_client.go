// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// GetUsagesInLocationClient contains the methods for the GetUsagesInLocation group.
// Don't use this type directly, use NewGetUsagesInLocationClient() instead.
type GetUsagesInLocationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGetUsagesInLocationClient creates a new instance of GetUsagesInLocationClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGetUsagesInLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GetUsagesInLocationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GetUsagesInLocationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List usages in cores for all skus used by a subscription in a given location, for a specific quota type.
//
// Generated from API version 2024-11-01
//   - location - The name of the Azure region.
//   - options - GetUsagesInLocationClientListOptions contains the optional parameters for the GetUsagesInLocationClient.NewListPager
//     method.
func (client *GetUsagesInLocationClient) NewListPager(location string, options *GetUsagesInLocationClientListOptions) *runtime.Pager[GetUsagesInLocationClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GetUsagesInLocationClientListResponse]{
		More: func(page GetUsagesInLocationClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GetUsagesInLocationClientListResponse) (GetUsagesInLocationClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GetUsagesInLocationClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return GetUsagesInLocationClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GetUsagesInLocationClient) listCreateRequest(ctx context.Context, location string, _ *GetUsagesInLocationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/usages"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-11-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GetUsagesInLocationClient) listHandleResponse(resp *http.Response) (GetUsagesInLocationClientListResponse, error) {
	result := GetUsagesInLocationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CsmUsageQuotaCollection); err != nil {
		return GetUsagesInLocationClientListResponse{}, err
	}
	return result, nil
}
