//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// ServiceTagsClient contains the methods for the ServiceTags group.
// Don't use this type directly, use NewServiceTagsClient() instead.
type ServiceTagsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServiceTagsClient creates a new instance of ServiceTagsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServiceTagsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServiceTagsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ServiceTagsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// List - Gets a list of service tag information resources.
// If the operation fails it returns an *azcore.ResponseError type.
// location - The location that will be used as a reference for version (not as a filter based on location, you will get the
// list of service tags with prefix details across all regions but limited to the cloud that
// your subscription belongs to).
// options - ServiceTagsClientListOptions contains the optional parameters for the ServiceTagsClient.List method.
func (client *ServiceTagsClient) List(ctx context.Context, location string, options *ServiceTagsClientListOptions) (ServiceTagsClientListResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return ServiceTagsClientListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServiceTagsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServiceTagsClientListResponse{}, runtime.NewResponseError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *ServiceTagsClient) listCreateRequest(ctx context.Context, location string, options *ServiceTagsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/locations/{location}/serviceTags"
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
	reqQP.Set("api-version", "2021-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ServiceTagsClient) listHandleResponse(resp *http.Response) (ServiceTagsClientListResponse, error) {
	result := ServiceTagsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServiceTagsListResult); err != nil {
		return ServiceTagsClientListResponse{}, err
	}
	return result, nil
}
