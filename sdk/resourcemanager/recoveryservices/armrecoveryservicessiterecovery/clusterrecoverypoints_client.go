// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservicessiterecovery

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

// ClusterRecoveryPointsClient contains the methods for the ClusterRecoveryPoints group.
// Don't use this type directly, use NewClusterRecoveryPointsClient() instead.
type ClusterRecoveryPointsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterRecoveryPointsClient creates a new instance of ClusterRecoveryPointsClient with the specified values.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterRecoveryPointsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterRecoveryPointsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterRecoveryPointsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByReplicationProtectionClusterPager - The list of cluster recovery points.
//
// Generated from API version 2025-01-01
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - resourceName - The name of the recovery services vault.
//   - fabricName - Fabric name.
//   - protectionContainerName - Protection container name.
//   - replicationProtectionClusterName - Replication protection cluster name.
//   - options - ClusterRecoveryPointsClientListByReplicationProtectionClusterOptions contains the optional parameters for the
//     ClusterRecoveryPointsClient.NewListByReplicationProtectionClusterPager method.
func (client *ClusterRecoveryPointsClient) NewListByReplicationProtectionClusterPager(resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicationProtectionClusterName string, options *ClusterRecoveryPointsClientListByReplicationProtectionClusterOptions) *runtime.Pager[ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse]{
		More: func(page ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse) (ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ClusterRecoveryPointsClient.NewListByReplicationProtectionClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByReplicationProtectionClusterCreateRequest(ctx, resourceGroupName, resourceName, fabricName, protectionContainerName, replicationProtectionClusterName, options)
			}, nil)
			if err != nil {
				return ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse{}, err
			}
			return client.listByReplicationProtectionClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByReplicationProtectionClusterCreateRequest creates the ListByReplicationProtectionCluster request.
func (client *ClusterRecoveryPointsClient) listByReplicationProtectionClusterCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, fabricName string, protectionContainerName string, replicationProtectionClusterName string, _ *ClusterRecoveryPointsClientListByReplicationProtectionClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationFabrics/{fabricName}/replicationProtectionContainers/{protectionContainerName}/replicationProtectionClusters/{replicationProtectionClusterName}/recoveryPoints"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if fabricName == "" {
		return nil, errors.New("parameter fabricName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{fabricName}", url.PathEscape(fabricName))
	if protectionContainerName == "" {
		return nil, errors.New("parameter protectionContainerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{protectionContainerName}", url.PathEscape(protectionContainerName))
	if replicationProtectionClusterName == "" {
		return nil, errors.New("parameter replicationProtectionClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{replicationProtectionClusterName}", url.PathEscape(replicationProtectionClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByReplicationProtectionClusterHandleResponse handles the ListByReplicationProtectionCluster response.
func (client *ClusterRecoveryPointsClient) listByReplicationProtectionClusterHandleResponse(resp *http.Response) (ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse, error) {
	result := ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterRecoveryPointCollection); err != nil {
		return ClusterRecoveryPointsClientListByReplicationProtectionClusterResponse{}, err
	}
	return result, nil
}
