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

// ApplicationGatewayPrivateLinkResourcesClient contains the methods for the ApplicationGatewayPrivateLinkResources group.
// Don't use this type directly, use NewApplicationGatewayPrivateLinkResourcesClient() instead.
type ApplicationGatewayPrivateLinkResourcesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewApplicationGatewayPrivateLinkResourcesClient creates a new instance of ApplicationGatewayPrivateLinkResourcesClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewApplicationGatewayPrivateLinkResourcesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ApplicationGatewayPrivateLinkResourcesClient, error) {
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
	client := &ApplicationGatewayPrivateLinkResourcesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Lists all private link resources on an application gateway.
//
// Generated from API version 2022-07-01
//   - resourceGroupName - The name of the resource group.
//   - applicationGatewayName - The name of the application gateway.
//   - options - ApplicationGatewayPrivateLinkResourcesClientListOptions contains the optional parameters for the ApplicationGatewayPrivateLinkResourcesClient.NewListPager
//     method.
func (client *ApplicationGatewayPrivateLinkResourcesClient) NewListPager(resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateLinkResourcesClientListOptions) *runtime.Pager[ApplicationGatewayPrivateLinkResourcesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ApplicationGatewayPrivateLinkResourcesClientListResponse]{
		More: func(page ApplicationGatewayPrivateLinkResourcesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ApplicationGatewayPrivateLinkResourcesClientListResponse) (ApplicationGatewayPrivateLinkResourcesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, applicationGatewayName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ApplicationGatewayPrivateLinkResourcesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ApplicationGatewayPrivateLinkResourcesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ApplicationGatewayPrivateLinkResourcesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ApplicationGatewayPrivateLinkResourcesClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationGatewayName string, options *ApplicationGatewayPrivateLinkResourcesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/applicationGateways/{applicationGatewayName}/privateLinkResources"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGatewayName == "" {
		return nil, errors.New("parameter applicationGatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGatewayName}", url.PathEscape(applicationGatewayName))
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
func (client *ApplicationGatewayPrivateLinkResourcesClient) listHandleResponse(resp *http.Response) (ApplicationGatewayPrivateLinkResourcesClientListResponse, error) {
	result := ApplicationGatewayPrivateLinkResourcesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ApplicationGatewayPrivateLinkResourceListResult); err != nil {
		return ApplicationGatewayPrivateLinkResourcesClientListResponse{}, err
	}
	return result, nil
}
