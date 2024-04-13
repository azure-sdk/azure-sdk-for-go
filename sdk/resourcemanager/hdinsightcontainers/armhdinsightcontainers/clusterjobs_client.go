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

// ClusterJobsClient contains the methods for the ClusterJobs group.
// Don't use this type directly, use NewClusterJobsClient() instead.
type ClusterJobsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterJobsClient creates a new instance of ClusterJobsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterJobsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterJobsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterJobsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Get jobs of HDInsight on AKS cluster.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - options - ClusterJobsClientListOptions contains the optional parameters for the ClusterJobsClient.NewListPager method.
func (client *ClusterJobsClient) NewListPager(resourceGroupName string, clusterPoolName string, clusterName string, options *ClusterJobsClientListOptions) *runtime.Pager[ClusterJobsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterJobsClientListResponse]{
		More: func(page ClusterJobsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterJobsClientListResponse) (ClusterJobsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ClusterJobsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, options)
			}, nil)
			if err != nil {
				return ClusterJobsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ClusterJobsClient) listCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, options *ClusterJobsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/jobs"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ClusterJobsClient) listHandleResponse(resp *http.Response) (ClusterJobsClientListResponse, error) {
	result := ClusterJobsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterJobList); err != nil {
		return ClusterJobsClientListResponse{}, err
	}
	return result, nil
}

// BeginRunJob - Operations on jobs of HDInsight on AKS cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterName - The name of the HDInsight cluster.
//   - clusterJob - The Cluster job.
//   - options - ClusterJobsClientBeginRunJobOptions contains the optional parameters for the ClusterJobsClient.BeginRunJob method.
func (client *ClusterJobsClient) BeginRunJob(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterJob ClusterJob, options *ClusterJobsClientBeginRunJobOptions) (*runtime.Poller[ClusterJobsClientRunJobResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.runJob(ctx, resourceGroupName, clusterPoolName, clusterName, clusterJob, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClusterJobsClientRunJobResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClusterJobsClientRunJobResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// RunJob - Operations on jobs of HDInsight on AKS cluster.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
func (client *ClusterJobsClient) runJob(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterJob ClusterJob, options *ClusterJobsClientBeginRunJobOptions) (*http.Response, error) {
	var err error
	const operationName = "ClusterJobsClient.BeginRunJob"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.runJobCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterName, clusterJob, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// runJobCreateRequest creates the RunJob request.
func (client *ClusterJobsClient) runJobCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterName string, clusterJob ClusterJob, options *ClusterJobsClientBeginRunJobOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}/clusters/{clusterName}/runJob"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterJob); err != nil {
		return nil, err
	}
	return req, nil
}
