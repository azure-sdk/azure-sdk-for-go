//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// UsagesClient contains the methods for the Usages group.
// Don't use this type directly, use NewUsagesClient() instead.
type UsagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewUsagesClient creates a new instance of UsagesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewUsagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*UsagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &UsagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByInstancePoolPager - Gets all instance pool usage metrics
//
// Generated from API version 2021-02-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - instancePoolName - The name of the instance pool to be retrieved.
//   - options - UsagesClientListByInstancePoolOptions contains the optional parameters for the UsagesClient.NewListByInstancePoolPager
//     method.
func (client *UsagesClient) NewListByInstancePoolPager(resourceGroupName string, instancePoolName string, options *UsagesClientListByInstancePoolOptions) *runtime.Pager[UsagesClientListByInstancePoolResponse] {
	return runtime.NewPager(runtime.PagingHandler[UsagesClientListByInstancePoolResponse]{
		More: func(page UsagesClientListByInstancePoolResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *UsagesClientListByInstancePoolResponse) (UsagesClientListByInstancePoolResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "UsagesClient.NewListByInstancePoolPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByInstancePoolCreateRequest(ctx, resourceGroupName, instancePoolName, options)
			}, nil)
			if err != nil {
				return UsagesClientListByInstancePoolResponse{}, err
			}
			return client.listByInstancePoolHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByInstancePoolCreateRequest creates the ListByInstancePool request.
func (client *UsagesClient) listByInstancePoolCreateRequest(ctx context.Context, resourceGroupName string, instancePoolName string, options *UsagesClientListByInstancePoolOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/instancePools/{instancePoolName}/usages"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if instancePoolName == "" {
		return nil, errors.New("parameter instancePoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{instancePoolName}", url.PathEscape(instancePoolName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	if options != nil && options.ExpandChildren != nil {
		reqQP.Set("expandChildren", strconv.FormatBool(*options.ExpandChildren))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByInstancePoolHandleResponse handles the ListByInstancePool response.
func (client *UsagesClient) listByInstancePoolHandleResponse(resp *http.Response) (UsagesClientListByInstancePoolResponse, error) {
	result := UsagesClientListByInstancePoolResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.UsageListResult); err != nil {
		return UsagesClientListByInstancePoolResponse{}, err
	}
	return result, nil
}
