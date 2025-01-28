//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// APIProductClient contains the methods for the APIProduct group.
// Don't use this type directly, use NewAPIProductClient() instead.
type APIProductClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAPIProductClient creates a new instance of APIProductClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAPIProductClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*APIProductClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &APIProductClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByApisPager - Lists all Products, which the API is part of.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - apiID - API identifier. Must be unique in the current API Management service instance.
//   - options - APIProductClientListByApisOptions contains the optional parameters for the APIProductClient.NewListByApisPager
//     method.
func (client *APIProductClient) NewListByApisPager(resourceGroupName string, serviceName string, apiID string, options *APIProductClientListByApisOptions) *runtime.Pager[APIProductClientListByApisResponse] {
	return runtime.NewPager(runtime.PagingHandler[APIProductClientListByApisResponse]{
		More: func(page APIProductClientListByApisResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *APIProductClientListByApisResponse) (APIProductClientListByApisResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "APIProductClient.NewListByApisPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByApisCreateRequest(ctx, resourceGroupName, serviceName, apiID, options)
			}, nil)
			if err != nil {
				return APIProductClientListByApisResponse{}, err
			}
			return client.listByApisHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByApisCreateRequest creates the ListByApis request.
func (client *APIProductClient) listByApisCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, apiID string, options *APIProductClientListByApisOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/apis/{apiId}/products"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if apiID == "" {
		return nil, errors.New("parameter apiID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiId}", url.PathEscape(apiID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByApisHandleResponse handles the ListByApis response.
func (client *APIProductClient) listByApisHandleResponse(resp *http.Response) (APIProductClientListByApisResponse, error) {
	result := APIProductClientListByApisResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductCollection); err != nil {
		return APIProductClientListByApisResponse{}, err
	}
	return result, nil
}
