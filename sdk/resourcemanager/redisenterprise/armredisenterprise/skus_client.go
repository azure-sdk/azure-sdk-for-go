//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armredisenterprise

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

// SKUsClient contains the methods for the SKUs group.
// Don't use this type directly, use NewSKUsClient() instead.
type SKUsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSKUsClient creates a new instance of SKUsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SKUsClient, error) {
	cl, err := arm.NewClient(moduleName+".SKUsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SKUsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets information about skus in specified location for the given subscription id
//
// Generated from API version 2023-07-01-preview
//   - location - The name of Azure region.
//   - options - SKUsClientListOptions contains the optional parameters for the SKUsClient.NewListPager method.
func (client *SKUsClient) NewListPager(location string, options *SKUsClientListOptions) *runtime.Pager[SKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SKUsClientListResponse]{
		More: func(page SKUsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *SKUsClientListResponse) (SKUsClientListResponse, error) {
			req, err := client.listCreateRequest(ctx, location, options)
			if err != nil {
				return SKUsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SKUsClient) listCreateRequest(ctx context.Context, location string, options *SKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Cache/locations/{location}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SKUsClient) listHandleResponse(resp *http.Response) (SKUsClientListResponse, error) {
	result := SKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RegionSKUDetails); err != nil {
		return SKUsClientListResponse{}, err
	}
	return result, nil
}
