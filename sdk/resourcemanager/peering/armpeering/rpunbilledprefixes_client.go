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
	"strconv"
	"strings"
)

// RpUnbilledPrefixesClient contains the methods for the RpUnbilledPrefixes group.
// Don't use this type directly, use NewRpUnbilledPrefixesClient() instead.
type RpUnbilledPrefixesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRpUnbilledPrefixesClient creates a new instance of RpUnbilledPrefixesClient with the specified values.
//   - subscriptionID - The Azure subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRpUnbilledPrefixesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RpUnbilledPrefixesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RpUnbilledPrefixesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all of the RP unbilled prefixes for the specified peering
//
// Generated from API version 2022-10-01
//   - resourceGroupName - The Azure resource group name.
//   - peeringName - The peering name.
//   - options - RpUnbilledPrefixesClientListOptions contains the optional parameters for the RpUnbilledPrefixesClient.NewListPager
//     method.
func (client *RpUnbilledPrefixesClient) NewListPager(resourceGroupName string, peeringName string, options *RpUnbilledPrefixesClientListOptions) *runtime.Pager[RpUnbilledPrefixesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RpUnbilledPrefixesClientListResponse]{
		More: func(page RpUnbilledPrefixesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RpUnbilledPrefixesClientListResponse) (RpUnbilledPrefixesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RpUnbilledPrefixesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, peeringName, options)
			}, nil)
			if err != nil {
				return RpUnbilledPrefixesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RpUnbilledPrefixesClient) listCreateRequest(ctx context.Context, resourceGroupName string, peeringName string, options *RpUnbilledPrefixesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Peering/peerings/{peeringName}/rpUnbilledPrefixes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if peeringName == "" {
		return nil, errors.New("parameter peeringName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{peeringName}", url.PathEscape(peeringName))
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
	if options != nil && options.Consolidate != nil {
		reqQP.Set("consolidate", strconv.FormatBool(*options.Consolidate))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RpUnbilledPrefixesClient) listHandleResponse(resp *http.Response) (RpUnbilledPrefixesClientListResponse, error) {
	result := RpUnbilledPrefixesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RpUnbilledPrefixListResult); err != nil {
		return RpUnbilledPrefixesClientListResponse{}, err
	}
	return result, nil
}
