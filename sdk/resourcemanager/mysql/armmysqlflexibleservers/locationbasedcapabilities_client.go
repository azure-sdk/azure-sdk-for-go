//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmysqlflexibleservers

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

// LocationBasedCapabilitiesClient contains the methods for the LocationBasedCapabilities group.
// Don't use this type directly, use NewLocationBasedCapabilitiesClient() instead.
type LocationBasedCapabilitiesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLocationBasedCapabilitiesClient creates a new instance of LocationBasedCapabilitiesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLocationBasedCapabilitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LocationBasedCapabilitiesClient, error) {
	cl, err := arm.NewClient(moduleName+".LocationBasedCapabilitiesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LocationBasedCapabilitiesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get capabilities at specified location in a given subscription.
//
// Generated from API version 2023-06-30
//   - locationName - The name of the location.
//   - options - LocationBasedCapabilitiesClientListOptions contains the optional parameters for the LocationBasedCapabilitiesClient.NewListPager
//     method.
func (client *LocationBasedCapabilitiesClient) NewListPager(locationName string, options *LocationBasedCapabilitiesClientListOptions) *runtime.Pager[LocationBasedCapabilitiesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LocationBasedCapabilitiesClientListResponse]{
		More: func(page LocationBasedCapabilitiesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LocationBasedCapabilitiesClientListResponse) (LocationBasedCapabilitiesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, locationName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return LocationBasedCapabilitiesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return LocationBasedCapabilitiesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return LocationBasedCapabilitiesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *LocationBasedCapabilitiesClient) listCreateRequest(ctx context.Context, locationName string, options *LocationBasedCapabilitiesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.DBforMySQL/locations/{locationName}/capabilities"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-30")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LocationBasedCapabilitiesClient) listHandleResponse(resp *http.Response) (LocationBasedCapabilitiesClientListResponse, error) {
	result := LocationBasedCapabilitiesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilitiesListResult); err != nil {
		return LocationBasedCapabilitiesClientListResponse{}, err
	}
	return result, nil
}
