// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpeering

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

// CdnPeeringPrefixesClient contains the methods for the CdnPeeringPrefixes group.
// Don't use this type directly, use NewCdnPeeringPrefixesClient() instead.
type CdnPeeringPrefixesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCdnPeeringPrefixesClient creates a new instance of CdnPeeringPrefixesClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCdnPeeringPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CdnPeeringPrefixesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CdnPeeringPrefixesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all of the advertised prefixes for the specified peering location
//
// Generated from API version 2022-10-01
//   - peeringLocation - The peering location.
//   - options - CdnPeeringPrefixesClientListOptions contains the optional parameters for the CdnPeeringPrefixesClient.NewListPager
//     method.
func (client *CdnPeeringPrefixesClient) NewListPager(peeringLocation string, options *CdnPeeringPrefixesClientListOptions) *runtime.Pager[CdnPeeringPrefixesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CdnPeeringPrefixesClientListResponse]{
		More: func(page CdnPeeringPrefixesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CdnPeeringPrefixesClientListResponse) (CdnPeeringPrefixesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CdnPeeringPrefixesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, peeringLocation, options)
			}, nil)
			if err != nil {
				return CdnPeeringPrefixesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *CdnPeeringPrefixesClient) listCreateRequest(ctx context.Context, peeringLocation string, _ *CdnPeeringPrefixesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Peering/cdnPeeringPrefixes"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-10-01")
	reqQP.Set("peeringLocation", peeringLocation)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CdnPeeringPrefixesClient) listHandleResponse(resp *http.Response) (CdnPeeringPrefixesClientListResponse, error) {
	result := CdnPeeringPrefixesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CdnPeeringPrefixListResult); err != nil {
		return CdnPeeringPrefixesClientListResponse{}, err
	}
	return result, nil
}
