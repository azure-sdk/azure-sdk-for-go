//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// EntitiesRelationsClient contains the methods for the EntitiesRelations group.
// Don't use this type directly, use NewEntitiesRelationsClient() instead.
type EntitiesRelationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewEntitiesRelationsClient creates a new instance of EntitiesRelationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEntitiesRelationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EntitiesRelationsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EntitiesRelationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets all relations of an entity.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - entityID - entity ID
//   - options - EntitiesRelationsClientListOptions contains the optional parameters for the EntitiesRelationsClient.NewListPager
//     method.
func (client *EntitiesRelationsClient) NewListPager(resourceGroupName string, workspaceName string, entityID string, options *EntitiesRelationsClientListOptions) *runtime.Pager[EntitiesRelationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EntitiesRelationsClientListResponse]{
		More: func(page EntitiesRelationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EntitiesRelationsClientListResponse) (EntitiesRelationsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EntitiesRelationsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, entityID, options)
			}, nil)
			if err != nil {
				return EntitiesRelationsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EntitiesRelationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, entityID string, options *EntitiesRelationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/entities/{entityId}/relations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if entityID == "" {
		return nil, errors.New("parameter entityID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{entityId}", url.PathEscape(entityID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EntitiesRelationsClient) listHandleResponse(resp *http.Response) (EntitiesRelationsClientListResponse, error) {
	result := EntitiesRelationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RelationList); err != nil {
		return EntitiesRelationsClientListResponse{}, err
	}
	return result, nil
}
