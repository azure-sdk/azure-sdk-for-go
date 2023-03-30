//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsynapse

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

// PrivateEndpointConnectionsPrivateLinkHubClient contains the methods for the PrivateEndpointConnectionsPrivateLinkHub group.
// Don't use this type directly, use NewPrivateEndpointConnectionsPrivateLinkHubClient() instead.
type PrivateEndpointConnectionsPrivateLinkHubClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPrivateEndpointConnectionsPrivateLinkHubClient creates a new instance of PrivateEndpointConnectionsPrivateLinkHubClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPrivateEndpointConnectionsPrivateLinkHubClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PrivateEndpointConnectionsPrivateLinkHubClient, error) {
	cl, err := arm.NewClient(moduleName+".PrivateEndpointConnectionsPrivateLinkHubClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PrivateEndpointConnectionsPrivateLinkHubClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get all PrivateEndpointConnection in the PrivateLinkHub by name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateLinkHubName - Name of the privateLinkHub
//   - privateEndpointConnectionName - Name of the privateEndpointConnection
//   - options - PrivateEndpointConnectionsPrivateLinkHubClientGetOptions contains the optional parameters for the PrivateEndpointConnectionsPrivateLinkHubClient.Get
//     method.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) Get(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsPrivateLinkHubClientGetOptions) (PrivateEndpointConnectionsPrivateLinkHubClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateLinkHubName, privateEndpointConnectionName, options)
	if err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PrivateEndpointConnectionsPrivateLinkHubClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, privateEndpointConnectionName string, options *PrivateEndpointConnectionsPrivateLinkHubClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections/{privateEndpointConnectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateLinkHubName == "" {
		return nil, errors.New("parameter privateLinkHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkHubName}", url.PathEscape(privateLinkHubName))
	if privateEndpointConnectionName == "" {
		return nil, errors.New("parameter privateEndpointConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateEndpointConnectionName}", url.PathEscape(privateEndpointConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) getHandleResponse(resp *http.Response) (PrivateEndpointConnectionsPrivateLinkHubClientGetResponse, error) {
	result := PrivateEndpointConnectionsPrivateLinkHubClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionForPrivateLinkHub); err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get all PrivateEndpointConnections in the PrivateLinkHub
//
// Generated from API version 2021-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateLinkHubName - Name of the privateLinkHub
//   - options - PrivateEndpointConnectionsPrivateLinkHubClientListOptions contains the optional parameters for the PrivateEndpointConnectionsPrivateLinkHubClient.NewListPager
//     method.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) NewListPager(resourceGroupName string, privateLinkHubName string, options *PrivateEndpointConnectionsPrivateLinkHubClientListOptions) *runtime.Pager[PrivateEndpointConnectionsPrivateLinkHubClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PrivateEndpointConnectionsPrivateLinkHubClientListResponse]{
		More: func(page PrivateEndpointConnectionsPrivateLinkHubClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PrivateEndpointConnectionsPrivateLinkHubClientListResponse) (PrivateEndpointConnectionsPrivateLinkHubClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, privateLinkHubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PrivateEndpointConnectionsPrivateLinkHubClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return PrivateEndpointConnectionsPrivateLinkHubClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PrivateEndpointConnectionsPrivateLinkHubClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateLinkHubName string, options *PrivateEndpointConnectionsPrivateLinkHubClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/privateLinkHubs/{privateLinkHubName}/privateEndpointConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateLinkHubName == "" {
		return nil, errors.New("parameter privateLinkHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateLinkHubName}", url.PathEscape(privateLinkHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PrivateEndpointConnectionsPrivateLinkHubClient) listHandleResponse(resp *http.Response) (PrivateEndpointConnectionsPrivateLinkHubClientListResponse, error) {
	result := PrivateEndpointConnectionsPrivateLinkHubClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PrivateEndpointConnectionForPrivateLinkHubResourceCollectionResponse); err != nil {
		return PrivateEndpointConnectionsPrivateLinkHubClientListResponse{}, err
	}
	return result, nil
}
