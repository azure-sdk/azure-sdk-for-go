//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armpolicyinsights

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// PolicyMetadataClient contains the methods for the PolicyMetadata group.
// Don't use this type directly, use NewPolicyMetadataClient() instead.
type PolicyMetadataClient struct {
	host string
	pl   runtime.Pipeline
}

// NewPolicyMetadataClient creates a new instance of PolicyMetadataClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPolicyMetadataClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*PolicyMetadataClient, error) {
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
	client := &PolicyMetadataClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// GetResource - Get policy metadata resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2019-10-01
//   - resourceName - The name of the policy metadata resource.
//   - options - PolicyMetadataClientGetResourceOptions contains the optional parameters for the PolicyMetadataClient.GetResource
//     method.
func (client *PolicyMetadataClient) GetResource(ctx context.Context, resourceName string, options *PolicyMetadataClientGetResourceOptions) (PolicyMetadataClientGetResourceResponse, error) {
	req, err := client.getResourceCreateRequest(ctx, resourceName, options)
	if err != nil {
		return PolicyMetadataClientGetResourceResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PolicyMetadataClientGetResourceResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PolicyMetadataClientGetResourceResponse{}, runtime.NewResponseError(resp)
	}
	return client.getResourceHandleResponse(resp)
}

// getResourceCreateRequest creates the GetResource request.
func (client *PolicyMetadataClient) getResourceCreateRequest(ctx context.Context, resourceName string, options *PolicyMetadataClientGetResourceOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PolicyInsights/policyMetadata/{resourceName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", resourceName)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getResourceHandleResponse handles the GetResource response.
func (client *PolicyMetadataClient) getResourceHandleResponse(resp *http.Response) (PolicyMetadataClientGetResourceResponse, error) {
	result := PolicyMetadataClientGetResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyMetadata); err != nil {
		return PolicyMetadataClientGetResourceResponse{}, err
	}
	return result, nil
}

// NewListPager - Get a list of the policy metadata resources.
//
// Generated from API version 2019-10-01
//   - QueryOptions - QueryOptions contains a group of parameters for the PolicyTrackedResourcesClient.ListQueryResultsForManagementGroup
//     method.
//   - options - PolicyMetadataClientListOptions contains the optional parameters for the PolicyMetadataClient.NewListPager method.
func (client *PolicyMetadataClient) NewListPager(queryOptions *QueryOptions, options *PolicyMetadataClientListOptions) *runtime.Pager[PolicyMetadataClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[PolicyMetadataClientListResponse]{
		More: func(page PolicyMetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *PolicyMetadataClientListResponse) (PolicyMetadataClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, queryOptions, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return PolicyMetadataClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return PolicyMetadataClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return PolicyMetadataClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *PolicyMetadataClient) listCreateRequest(ctx context.Context, queryOptions *QueryOptions, options *PolicyMetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.PolicyInsights/policyMetadata"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	if queryOptions != nil && queryOptions.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*queryOptions.Top), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *PolicyMetadataClient) listHandleResponse(resp *http.Response) (PolicyMetadataClientListResponse, error) {
	result := PolicyMetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyMetadataCollection); err != nil {
		return PolicyMetadataClientListResponse{}, err
	}
	return result, nil
}
