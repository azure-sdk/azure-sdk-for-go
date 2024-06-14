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
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConnectedPartnerResourcesClient contains the methods for the ConnectedPartnerResources group.
// Don't use this type directly, use NewConnectedPartnerResourcesClient() instead.
type ConnectedPartnerResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectedPartnerResourcesClient creates a new instance of ConnectedPartnerResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectedPartnerResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectedPartnerResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectedPartnerResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - List of all active deployments that are associated with the marketplace subscription linked to the given
// monitor.
//
// Generated from API version 2024-06-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - monitorName - Monitor resource name
//   - options - ConnectedPartnerResourcesClientListOptions contains the optional parameters for the ConnectedPartnerResourcesClient.NewListPager
//     method.
func (client *ConnectedPartnerResourcesClient) NewListPager(resourceGroupName string, monitorName string, options *ConnectedPartnerResourcesClientListOptions) *runtime.Pager[ConnectedPartnerResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectedPartnerResourcesClientListResponse]{
		More: func(page ConnectedPartnerResourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectedPartnerResourcesClientListResponse) (ConnectedPartnerResourcesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectedPartnerResourcesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, monitorName, options)
			}, nil)
			if err != nil {
				return ConnectedPartnerResourcesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConnectedPartnerResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, monitorName string, options *ConnectedPartnerResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Elastic/monitors/{monitorName}/listConnectedPartnerResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if monitorName == "" {
		return nil, errors.New("parameter monitorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{monitorName}", url.PathEscape(monitorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectedPartnerResourcesClient) listHandleResponse(resp *http.Response) (ConnectedPartnerResourcesClientListResponse, error) {
	result := ConnectedPartnerResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectedPartnerResourcesListResponse); err != nil {
		return ConnectedPartnerResourcesClientListResponse{}, err
	}
	return result, nil
}
