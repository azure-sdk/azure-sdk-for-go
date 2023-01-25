//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armnetwork

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

// ExpressRoutePortsLocationsClient contains the methods for the ExpressRoutePortsLocations group.
// Don't use this type directly, use NewExpressRoutePortsLocationsClient() instead.
type ExpressRoutePortsLocationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewExpressRoutePortsLocationsClient creates a new instance of ExpressRoutePortsLocationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExpressRoutePortsLocationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRoutePortsLocationsClient, error) {
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
	client := &ExpressRoutePortsLocationsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Retrieves a single ExpressRoutePort peering location, including the list of available bandwidths available at said
// peering location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-07-01
//   - locationName - Name of the requested ExpressRoutePort peering location.
//   - options - ExpressRoutePortsLocationsClientGetOptions contains the optional parameters for the ExpressRoutePortsLocationsClient.Get
//     method.
func (client *ExpressRoutePortsLocationsClient) Get(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsClientGetOptions) (ExpressRoutePortsLocationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, locationName, options)
	if err != nil {
		return ExpressRoutePortsLocationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ExpressRoutePortsLocationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ExpressRoutePortsLocationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ExpressRoutePortsLocationsClient) getCreateRequest(ctx context.Context, locationName string, options *ExpressRoutePortsLocationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations/{locationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ExpressRoutePortsLocationsClient) getHandleResponse(resp *http.Response) (ExpressRoutePortsLocationsClientGetResponse, error) {
	result := ExpressRoutePortsLocationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortsLocation); err != nil {
		return ExpressRoutePortsLocationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Retrieves all ExpressRoutePort peering locations. Does not return available bandwidths for each location.
// Available bandwidths can only be obtained when retrieving a specific peering location.
//
// Generated from API version 2022-07-01
//   - options - ExpressRoutePortsLocationsClientListOptions contains the optional parameters for the ExpressRoutePortsLocationsClient.NewListPager
//     method.
func (client *ExpressRoutePortsLocationsClient) NewListPager(options *ExpressRoutePortsLocationsClientListOptions) *runtime.Pager[ExpressRoutePortsLocationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ExpressRoutePortsLocationsClientListResponse]{
		More: func(page ExpressRoutePortsLocationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ExpressRoutePortsLocationsClientListResponse) (ExpressRoutePortsLocationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ExpressRoutePortsLocationsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ExpressRoutePortsLocationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ExpressRoutePortsLocationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ExpressRoutePortsLocationsClient) listCreateRequest(ctx context.Context, options *ExpressRoutePortsLocationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ExpressRoutePortsLocations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRoutePortsLocationsClient) listHandleResponse(resp *http.Response) (ExpressRoutePortsLocationsClientListResponse, error) {
	result := ExpressRoutePortsLocationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRoutePortsLocationListResult); err != nil {
		return ExpressRoutePortsLocationsClientListResponse{}, err
	}
	return result, nil
}
