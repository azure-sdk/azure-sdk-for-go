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

// AvailableEndpointServicesClient contains the methods for the AvailableEndpointServices group.
// Don't use this type directly, use NewAvailableEndpointServicesClient() instead.
type AvailableEndpointServicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewAvailableEndpointServicesClient creates a new instance of AvailableEndpointServicesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAvailableEndpointServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableEndpointServicesClient, error) {
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
	client := &AvailableEndpointServicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - List what values of endpoint services are available for use.
// Generated from API version 2022-06-01-preview
// location - The location to check available endpoint services.
// options - AvailableEndpointServicesClientListOptions contains the optional parameters for the AvailableEndpointServicesClient.List
// method.
func (client *AvailableEndpointServicesClient) NewListPager(location string, options *AvailableEndpointServicesClientListOptions) *runtime.Pager[AvailableEndpointServicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AvailableEndpointServicesClientListResponse]{
		More: func(page AvailableEndpointServicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AvailableEndpointServicesClientListResponse) (AvailableEndpointServicesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AvailableEndpointServicesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AvailableEndpointServicesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AvailableEndpointServicesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AvailableEndpointServicesClient) listCreateRequest(ctx context.Context, location string, options *AvailableEndpointServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/virtualNetworkAvailableEndpointServices"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AvailableEndpointServicesClient) listHandleResponse(resp *http.Response) (AvailableEndpointServicesClientListResponse, error) {
	result := AvailableEndpointServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EndpointServicesListResult); err != nil {
		return AvailableEndpointServicesClientListResponse{}, err
	}
	return result, nil
}
