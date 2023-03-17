//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armcompute

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

// SharedGalleriesClient contains the methods for the SharedGalleries group.
// Don't use this type directly, use NewSharedGalleriesClient() instead.
type SharedGalleriesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSharedGalleriesClient creates a new instance of SharedGalleriesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharedGalleriesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharedGalleriesClient, error) {
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
	client := &SharedGalleriesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get a shared gallery by subscription id or tenant id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-03
//   - location - Resource location.
//   - galleryUniqueName - The unique name of the Shared Gallery.
//   - options - SharedGalleriesClientGetOptions contains the optional parameters for the SharedGalleriesClient.Get method.
func (client *SharedGalleriesClient) Get(ctx context.Context, location string, galleryUniqueName string, options *SharedGalleriesClientGetOptions) (SharedGalleriesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, location, galleryUniqueName, options)
	if err != nil {
		return SharedGalleriesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SharedGalleriesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SharedGalleriesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SharedGalleriesClient) getCreateRequest(ctx context.Context, location string, galleryUniqueName string, options *SharedGalleriesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if galleryUniqueName == "" {
		return nil, errors.New("parameter galleryUniqueName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryUniqueName}", url.PathEscape(galleryUniqueName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SharedGalleriesClient) getHandleResponse(resp *http.Response) (SharedGalleriesClientGetResponse, error) {
	result := SharedGalleriesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGallery); err != nil {
		return SharedGalleriesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List shared galleries by subscription id or tenant id.
//
// Generated from API version 2022-08-03
//   - location - Resource location.
//   - options - SharedGalleriesClientListOptions contains the optional parameters for the SharedGalleriesClient.NewListPager
//     method.
func (client *SharedGalleriesClient) NewListPager(location string, options *SharedGalleriesClientListOptions) *runtime.Pager[SharedGalleriesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharedGalleriesClientListResponse]{
		More: func(page SharedGalleriesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharedGalleriesClientListResponse) (SharedGalleriesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SharedGalleriesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SharedGalleriesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SharedGalleriesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *SharedGalleriesClient) listCreateRequest(ctx context.Context, location string, options *SharedGalleriesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-03")
	if options != nil && options.SharedTo != nil {
		reqQP.Set("sharedTo", string(*options.SharedTo))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SharedGalleriesClient) listHandleResponse(resp *http.Response) (SharedGalleriesClientListResponse, error) {
	result := SharedGalleriesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryList); err != nil {
		return SharedGalleriesClientListResponse{}, err
	}
	return result, nil
}
