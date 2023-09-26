//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armresourcehealth

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// ChildAvailabilityStatusesClient contains the methods for the ChildAvailabilityStatuses group.
// Don't use this type directly, use NewChildAvailabilityStatusesClient() instead.
type ChildAvailabilityStatusesClient struct {
	internal *arm.Client
}

// NewChildAvailabilityStatusesClient creates a new instance of ChildAvailabilityStatusesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewChildAvailabilityStatusesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*ChildAvailabilityStatusesClient, error) {
	cl, err := arm.NewClient(moduleName+".ChildAvailabilityStatusesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ChildAvailabilityStatusesClient{
		internal: cl,
	}
	return client, nil
}

// GetByResource - Gets current availability status for a single resource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-01-preview
//   - resourceURI - The fully qualified ID of the resource, including the resource name and resource type. Currently the API
//     only support one nesting level resource types :
//     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
//   - options - ChildAvailabilityStatusesClientGetByResourceOptions contains the optional parameters for the ChildAvailabilityStatusesClient.GetByResource
//     method.
func (client *ChildAvailabilityStatusesClient) GetByResource(ctx context.Context, resourceURI string, options *ChildAvailabilityStatusesClientGetByResourceOptions) (ChildAvailabilityStatusesClientGetByResourceResponse, error) {
	var err error
	req, err := client.getByResourceCreateRequest(ctx, resourceURI, options)
	if err != nil {
		return ChildAvailabilityStatusesClientGetByResourceResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ChildAvailabilityStatusesClientGetByResourceResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ChildAvailabilityStatusesClientGetByResourceResponse{}, err
	}
	resp, err := client.getByResourceHandleResponse(httpResp)
	return resp, err
}

// getByResourceCreateRequest creates the GetByResource request.
func (client *ChildAvailabilityStatusesClient) getByResourceCreateRequest(ctx context.Context, resourceURI string, options *ChildAvailabilityStatusesClientGetByResourceOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/childAvailabilityStatuses/current"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByResourceHandleResponse handles the GetByResource response.
func (client *ChildAvailabilityStatusesClient) getByResourceHandleResponse(resp *http.Response) (ChildAvailabilityStatusesClientGetByResourceResponse, error) {
	result := ChildAvailabilityStatusesClientGetByResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatus); err != nil {
		return ChildAvailabilityStatusesClientGetByResourceResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the historical availability statuses for a single child resource. Use the nextLink property in the
// response to get the next page of availability status
//
// Generated from API version 2023-07-01-preview
//   - resourceURI - The fully qualified ID of the resource, including the resource name and resource type. Currently the API
//     only support one nesting level resource types :
//     /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resource-provider-name}/{parentResourceType}/{parentResourceName}/{resourceType}/{resourceName}
//   - options - ChildAvailabilityStatusesClientListOptions contains the optional parameters for the ChildAvailabilityStatusesClient.NewListPager
//     method.
func (client *ChildAvailabilityStatusesClient) NewListPager(resourceURI string, options *ChildAvailabilityStatusesClientListOptions) *runtime.Pager[ChildAvailabilityStatusesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ChildAvailabilityStatusesClientListResponse]{
		More: func(page ChildAvailabilityStatusesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ChildAvailabilityStatusesClientListResponse) (ChildAvailabilityStatusesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceURI, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ChildAvailabilityStatusesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ChildAvailabilityStatusesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ChildAvailabilityStatusesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ChildAvailabilityStatusesClient) listCreateRequest(ctx context.Context, resourceURI string, options *ChildAvailabilityStatusesClientListOptions) (*policy.Request, error) {
	urlPath := "/{resourceUri}/providers/Microsoft.ResourceHealth/childAvailabilityStatuses"
	urlPath = strings.ReplaceAll(urlPath, "{resourceUri}", resourceURI)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-01-preview")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ChildAvailabilityStatusesClient) listHandleResponse(resp *http.Response) (ChildAvailabilityStatusesClientListResponse, error) {
	result := ChildAvailabilityStatusesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailabilityStatusListResult); err != nil {
		return ChildAvailabilityStatusesClientListResponse{}, err
	}
	return result, nil
}
