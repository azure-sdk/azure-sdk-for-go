// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhardwaresecuritymodules

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

// CloudHsmClusterPrivateLinkResourcesClient contains the methods for the CloudHsmClusterPrivateLinkResources group.
// Don't use this type directly, use NewCloudHsmClusterPrivateLinkResourcesClient() instead.
type CloudHsmClusterPrivateLinkResourcesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCloudHsmClusterPrivateLinkResourcesClient creates a new instance of CloudHsmClusterPrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCloudHsmClusterPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CloudHsmClusterPrivateLinkResourcesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CloudHsmClusterPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByCloudHsmClusterPager - Gets the private link resources supported for the Cloud Hsm Cluster.
//
// Generated from API version 2024-06-30-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - cloudHsmClusterName - The name of the Cloud HSM Cluster within the specified resource group. Cloud HSM Cluster names must
//     be between 3 and 23 characters in length.
//   - options - CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterOptions contains the optional parameters for the
//     CloudHsmClusterPrivateLinkResourcesClient.NewListByCloudHsmClusterPager method.
func (client *CloudHsmClusterPrivateLinkResourcesClient) NewListByCloudHsmClusterPager(resourceGroupName string, cloudHsmClusterName string, options *CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterOptions) *runtime.Pager[CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse] {
	return runtime.NewPager(runtime.PagingHandler[CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse]{
		More: func(page CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse) (CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "CloudHsmClusterPrivateLinkResourcesClient.NewListByCloudHsmClusterPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByCloudHsmClusterCreateRequest(ctx, resourceGroupName, cloudHsmClusterName, options)
			}, nil)
			if err != nil {
				return CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse{}, err
			}
			return client.listByCloudHsmClusterHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByCloudHsmClusterCreateRequest creates the ListByCloudHsmCluster request.
func (client *CloudHsmClusterPrivateLinkResourcesClient) listByCloudHsmClusterCreateRequest(ctx context.Context, resourceGroupName string, cloudHsmClusterName string, _ *CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HardwareSecurityModules/cloudHsmClusters/{cloudHsmClusterName}/privateLinkResources"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if cloudHsmClusterName == "" {
		return nil, errors.New("parameter cloudHsmClusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cloudHsmClusterName}", url.PathEscape(cloudHsmClusterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-30-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByCloudHsmClusterHandleResponse handles the ListByCloudHsmCluster response.
func (client *CloudHsmClusterPrivateLinkResourcesClient) listByCloudHsmClusterHandleResponse(resp *http.Response) (CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse, error) {
	result := CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateLinkResourceListResult); err != nil {
		return CloudHsmClusterPrivateLinkResourcesClientListByCloudHsmClusterResponse{}, err
	}
	return result, nil
}
