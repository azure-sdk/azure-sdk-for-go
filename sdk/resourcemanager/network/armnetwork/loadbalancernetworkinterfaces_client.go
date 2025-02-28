// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// LoadBalancerNetworkInterfacesClient contains the methods for the LoadBalancerNetworkInterfaces group.
// Don't use this type directly, use NewLoadBalancerNetworkInterfacesClient() instead.
type LoadBalancerNetworkInterfacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLoadBalancerNetworkInterfacesClient creates a new instance of LoadBalancerNetworkInterfacesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLoadBalancerNetworkInterfacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LoadBalancerNetworkInterfacesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LoadBalancerNetworkInterfacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets associated load balancer network interfaces.
//
// Generated from API version 2024-07-01
//   - resourceGroupName - The name of the resource group.
//   - loadBalancerName - The name of the load balancer.
//   - options - LoadBalancerNetworkInterfacesClientListOptions contains the optional parameters for the LoadBalancerNetworkInterfacesClient.NewListPager
//     method.
func (client *LoadBalancerNetworkInterfacesClient) NewListPager(resourceGroupName string, loadBalancerName string, options *LoadBalancerNetworkInterfacesClientListOptions) *runtime.Pager[LoadBalancerNetworkInterfacesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[LoadBalancerNetworkInterfacesClientListResponse]{
		More: func(page LoadBalancerNetworkInterfacesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *LoadBalancerNetworkInterfacesClientListResponse) (LoadBalancerNetworkInterfacesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "LoadBalancerNetworkInterfacesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, loadBalancerName, options)
			}, nil)
			if err != nil {
				return LoadBalancerNetworkInterfacesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *LoadBalancerNetworkInterfacesClient) listCreateRequest(ctx context.Context, resourceGroupName string, loadBalancerName string, _ *LoadBalancerNetworkInterfacesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/loadBalancers/{loadBalancerName}/networkInterfaces"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if loadBalancerName == "" {
		return nil, errors.New("parameter loadBalancerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{loadBalancerName}", url.PathEscape(loadBalancerName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *LoadBalancerNetworkInterfacesClient) listHandleResponse(resp *http.Response) (LoadBalancerNetworkInterfacesClientListResponse, error) {
	result := LoadBalancerNetworkInterfacesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.InterfaceListResult); err != nil {
		return LoadBalancerNetworkInterfacesClientListResponse{}, err
	}
	return result, nil
}
