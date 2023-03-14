//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armsecurity

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AllowedConnectionsClient contains the methods for the AllowedConnections group.
// Don't use this type directly, use NewAllowedConnectionsClient() instead.
type AllowedConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAllowedConnectionsClient creates a new instance of AllowedConnectionsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAllowedConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AllowedConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &AllowedConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets the list of all possible traffic between resources for the subscription and location, based on connection type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-01-01
//   - resourceGroupName - The name of the resource group within the user's subscription. The name is case insensitive.
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - connectionType - The type of allowed connections (Internal, External)
//   - options - AllowedConnectionsClientGetOptions contains the optional parameters for the AllowedConnectionsClient.Get method.
func (client *AllowedConnectionsClient) Get(ctx context.Context, resourceGroupName string, ascLocation string, connectionType ConnectionType, options *AllowedConnectionsClientGetOptions) (AllowedConnectionsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, ascLocation, connectionType, options)
	if err != nil {
		return AllowedConnectionsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AllowedConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AllowedConnectionsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AllowedConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, ascLocation string, connectionType ConnectionType, options *AllowedConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections/{connectionType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	if connectionType == "" {
		return nil, errors.New("parameter connectionType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionType}", url.PathEscape(string(connectionType)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AllowedConnectionsClient) getHandleResponse(resp *http.Response) (AllowedConnectionsClientGetResponse, error) {
	result := AllowedConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedConnectionsResource); err != nil {
		return AllowedConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets the list of all possible traffic between resources for the subscription
//
// Generated from API version 2020-01-01
//   - options - AllowedConnectionsClientListOptions contains the optional parameters for the AllowedConnectionsClient.NewListPager
//     method.
func (client *AllowedConnectionsClient) NewListPager(options *AllowedConnectionsClientListOptions) *runtime.Pager[AllowedConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AllowedConnectionsClientListResponse]{
		More: func(page AllowedConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AllowedConnectionsClientListResponse) (AllowedConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AllowedConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AllowedConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AllowedConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AllowedConnectionsClient) listCreateRequest(ctx context.Context, options *AllowedConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/allowedConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AllowedConnectionsClient) listHandleResponse(resp *http.Response) (AllowedConnectionsClientListResponse, error) {
	result := AllowedConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedConnectionsList); err != nil {
		return AllowedConnectionsClientListResponse{}, err
	}
	return result, nil
}

// NewListByHomeRegionPager - Gets the list of all possible traffic between resources for the subscription and location.
//
// Generated from API version 2020-01-01
//   - ascLocation - The location where ASC stores the data of the subscription. can be retrieved from Get locations
//   - options - AllowedConnectionsClientListByHomeRegionOptions contains the optional parameters for the AllowedConnectionsClient.NewListByHomeRegionPager
//     method.
func (client *AllowedConnectionsClient) NewListByHomeRegionPager(ascLocation string, options *AllowedConnectionsClientListByHomeRegionOptions) *runtime.Pager[AllowedConnectionsClientListByHomeRegionResponse] {
	return runtime.NewPager(runtime.PagingHandler[AllowedConnectionsClientListByHomeRegionResponse]{
		More: func(page AllowedConnectionsClientListByHomeRegionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AllowedConnectionsClientListByHomeRegionResponse) (AllowedConnectionsClientListByHomeRegionResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByHomeRegionCreateRequest(ctx, ascLocation, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AllowedConnectionsClientListByHomeRegionResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AllowedConnectionsClientListByHomeRegionResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AllowedConnectionsClientListByHomeRegionResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByHomeRegionHandleResponse(resp)
		},
	})
}

// listByHomeRegionCreateRequest creates the ListByHomeRegion request.
func (client *AllowedConnectionsClient) listByHomeRegionCreateRequest(ctx context.Context, ascLocation string, options *AllowedConnectionsClientListByHomeRegionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/locations/{ascLocation}/allowedConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if ascLocation == "" {
		return nil, errors.New("parameter ascLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ascLocation}", url.PathEscape(ascLocation))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByHomeRegionHandleResponse handles the ListByHomeRegion response.
func (client *AllowedConnectionsClient) listByHomeRegionHandleResponse(resp *http.Response) (AllowedConnectionsClientListByHomeRegionResponse, error) {
	result := AllowedConnectionsClientListByHomeRegionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AllowedConnectionsList); err != nil {
		return AllowedConnectionsClientListByHomeRegionResponse{}, err
	}
	return result, nil
}
