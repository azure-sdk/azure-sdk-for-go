//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// Code generated by Microsoft (R) AutoRest Code Generator.Changes may cause incorrect behavior and will be lost if the code
// is regenerated.
// Code generated by @autorest/go. DO NOT EDIT.

package armelastic

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VersionsClient contains the methods for the ElasticVersions group.
// Don't use this type directly, use NewVersionsClient() instead.
type VersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVersionsClient creates a new instance of VersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VersionsClient, error) {
	cl, err := arm.NewClient(moduleName+".VersionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get a list of available versions for a region.
//
// Generated from API version 2023-10-01-preview
//   - region - Region where elastic deployment will take place.
//   - options - VersionsClientListOptions contains the optional parameters for the VersionsClient.NewListPager method.
func (client *VersionsClient) NewListPager(region string, options *VersionsClientListOptions) *runtime.Pager[VersionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VersionsClientListResponse]{
		More: func(page VersionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VersionsClientListResponse) (VersionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, region, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VersionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return VersionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VersionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VersionsClient) listCreateRequest(ctx context.Context, region string, options *VersionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Elastic/elasticVersions"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	reqQP.Set("region", region)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VersionsClient) listHandleResponse(resp *http.Response) (VersionsClientListResponse, error) {
	result := VersionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VersionsListResponse); err != nil {
		return VersionsClientListResponse{}, err
	}
	return result, nil
}
