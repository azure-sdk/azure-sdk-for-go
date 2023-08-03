//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhdinsightonaks

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

// ClusterPoolsClient contains the methods for the ClusterPools group.
// Don't use this type directly, use NewClusterPoolsClient() instead.
type ClusterPoolsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClusterPoolsClient creates a new instance of ClusterPoolsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClusterPoolsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ClusterPoolsClient, error) {
	cl, err := arm.NewClient(moduleName+".ClusterPoolsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ClusterPoolsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a cluster pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterPool - The Cluster Pool to create.
//   - options - ClusterPoolsClientBeginCreateOrUpdateOptions contains the optional parameters for the ClusterPoolsClient.BeginCreateOrUpdate
//     method.
func (client *ClusterPoolsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterPool ClusterPool, options *ClusterPoolsClientBeginCreateOrUpdateOptions) (*runtime.Poller[ClusterPoolsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, clusterPoolName, clusterPool, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClusterPoolsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClusterPoolsClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a cluster pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClusterPoolsClient) createOrUpdate(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterPool ClusterPool, options *ClusterPoolsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterPool, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ClusterPoolsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterPool ClusterPool, options *ClusterPoolsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterPool); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a Cluster Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - options - ClusterPoolsClientBeginDeleteOptions contains the optional parameters for the ClusterPoolsClient.BeginDelete
//     method.
func (client *ClusterPoolsClient) BeginDelete(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClusterPoolsClientBeginDeleteOptions) (*runtime.Poller[ClusterPoolsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, clusterPoolName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[ClusterPoolsClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClusterPoolsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a Cluster Pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClusterPoolsClient) deleteOperation(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClusterPoolsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, clusterPoolName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ClusterPoolsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClusterPoolsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a cluster pool.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - options - ClusterPoolsClientGetOptions contains the optional parameters for the ClusterPoolsClient.Get method.
func (client *ClusterPoolsClient) Get(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClusterPoolsClientGetOptions) (ClusterPoolsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, clusterPoolName, options)
	if err != nil {
		return ClusterPoolsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClusterPoolsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClusterPoolsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ClusterPoolsClient) getCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, options *ClusterPoolsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ClusterPoolsClient) getHandleResponse(resp *http.Response) (ClusterPoolsClientGetResponse, error) {
	result := ClusterPoolsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPool); err != nil {
		return ClusterPoolsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Lists the HDInsight cluster pools under a resource group.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - ClusterPoolsClientListByResourceGroupOptions contains the optional parameters for the ClusterPoolsClient.NewListByResourceGroupPager
//     method.
func (client *ClusterPoolsClient) NewListByResourceGroupPager(resourceGroupName string, options *ClusterPoolsClientListByResourceGroupOptions) *runtime.Pager[ClusterPoolsClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterPoolsClientListByResourceGroupResponse]{
		More: func(page ClusterPoolsClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterPoolsClientListByResourceGroupResponse) (ClusterPoolsClientListByResourceGroupResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClusterPoolsClientListByResourceGroupResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClusterPoolsClientListByResourceGroupResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClusterPoolsClientListByResourceGroupResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *ClusterPoolsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *ClusterPoolsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *ClusterPoolsClient) listByResourceGroupHandleResponse(resp *http.Response) (ClusterPoolsClientListByResourceGroupResponse, error) {
	result := ClusterPoolsClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPoolListResult); err != nil {
		return ClusterPoolsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Gets the list of Cluster Pools within a Subscription.
//
// Generated from API version 2023-06-01-preview
//   - options - ClusterPoolsClientListBySubscriptionOptions contains the optional parameters for the ClusterPoolsClient.NewListBySubscriptionPager
//     method.
func (client *ClusterPoolsClient) NewListBySubscriptionPager(options *ClusterPoolsClientListBySubscriptionOptions) *runtime.Pager[ClusterPoolsClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClusterPoolsClientListBySubscriptionResponse]{
		More: func(page ClusterPoolsClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClusterPoolsClientListBySubscriptionResponse) (ClusterPoolsClientListBySubscriptionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listBySubscriptionCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClusterPoolsClientListBySubscriptionResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClusterPoolsClientListBySubscriptionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClusterPoolsClientListBySubscriptionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *ClusterPoolsClient) listBySubscriptionCreateRequest(ctx context.Context, options *ClusterPoolsClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HDInsight/clusterpools"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *ClusterPoolsClient) listBySubscriptionHandleResponse(resp *http.Response) (ClusterPoolsClientListBySubscriptionResponse, error) {
	result := ClusterPoolsClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ClusterPoolListResult); err != nil {
		return ClusterPoolsClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdateTags - Updates an existing Cluster Pool Tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterPoolName - The name of the cluster pool.
//   - clusterPoolTags - Parameters supplied to update tags.
//   - options - ClusterPoolsClientBeginUpdateTagsOptions contains the optional parameters for the ClusterPoolsClient.BeginUpdateTags
//     method.
func (client *ClusterPoolsClient) BeginUpdateTags(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterPoolTags TagsObject, options *ClusterPoolsClientBeginUpdateTagsOptions) (*runtime.Poller[ClusterPoolsClientUpdateTagsResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updateTags(ctx, resourceGroupName, clusterPoolName, clusterPoolTags, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClusterPoolsClientUpdateTagsResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[ClusterPoolsClientUpdateTagsResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdateTags - Updates an existing Cluster Pool Tags.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01-preview
func (client *ClusterPoolsClient) updateTags(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterPoolTags TagsObject, options *ClusterPoolsClientBeginUpdateTagsOptions) (*http.Response, error) {
	var err error
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, clusterPoolName, clusterPoolTags, options)
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

// updateTagsCreateRequest creates the UpdateTags request.
func (client *ClusterPoolsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, clusterPoolName string, clusterPoolTags TagsObject, options *ClusterPoolsClientBeginUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HDInsight/clusterpools/{clusterPoolName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if clusterPoolName == "" {
		return nil, errors.New("parameter clusterPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterPoolName}", url.PathEscape(clusterPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, clusterPoolTags); err != nil {
		return nil, err
	}
	return req, nil
}
