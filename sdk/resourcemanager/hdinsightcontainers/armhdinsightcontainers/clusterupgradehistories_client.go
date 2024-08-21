//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightcontainers

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

// ClusterUpgradeHistoriesClient contains the methods for the ClusterUpgradeHistories group.
// Don't use this type directly, use NewClusterUpgradeHistoriesClient() instead.
type ClusterUpgradeHistoriesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterUpgradeHistoriesClient creates a new instance of ClusterUpgradeHistoriesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterUpgradeHistoriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterUpgradeHistoriesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterUpgradeHistoriesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Returns a list of upgrade history.
//
// Generated from API version 2024-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClusterUpgradeHistoriesClientListOptions contains the optional parameters for the ClusterUpgradeHistoriesClient.NewListPager
//     method.
func (client *ClusterUpgradeHistoriesClient) NewListPager(resourceGroupName string, clusterPoolName string, clusterName string, options *ClusterUpgradeHistoriesClientListOptions) *runtime.Pager[ClusterUpgradeHistoriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterUpgradeHistoriesClientListResponse]{
		More: func(page ClusterUpgradeHistoriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterUpgradeHistoriesClientListResponse) (ClusterUpgradeHistoriesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ClusterUpgradeHistoriesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
			}, nil)
			if err != nil {
				return ClusterUpgradeHistoriesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ClusterUpgradeHistoriesClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClusterUpgradeHistoriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/upgradeHistories"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ClusterUpgradeHistoriesClient) listHandleResponse(resp *http.Response) (ClusterUpgradeHistoriesClientListResponse, error) {
	result := ClusterUpgradeHistoriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterUpgradeHistoryListResult); err != nil {
		return ClusterUpgradeHistoriesClientListResponse{}, err
	}
	return result, nil
}