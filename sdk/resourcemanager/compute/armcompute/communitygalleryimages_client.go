//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

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

// CommunityGalleryImagesClient contains the methods for the CommunityGalleryImages group.
// Don't use this type directly, use NewCommunityGalleryImagesClient() instead.
type CommunityGalleryImagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewCommunityGalleryImagesClient creates a new instance of CommunityGalleryImagesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewCommunityGalleryImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*CommunityGalleryImagesClient, error) {
	cl, err := arm.NewClient(moduleName+".CommunityGalleryImagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &CommunityGalleryImagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a community gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-03
//   - location - Resource location.
//   - publicGalleryName - The public name of the community gallery.
//   - galleryImageName - The name of the community gallery image definition.
//   - options - CommunityGalleryImagesClientGetOptions contains the optional parameters for the CommunityGalleryImagesClient.Get
//     method.
func (client *CommunityGalleryImagesClient) Get(ctx context.Context, location string, publicGalleryName string, galleryImageName string, options *CommunityGalleryImagesClientGetOptions) (CommunityGalleryImagesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, location, publicGalleryName, galleryImageName, options)
	if err != nil {
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *CommunityGalleryImagesClient) getCreateRequest(ctx context.Context, location string, publicGalleryName string, galleryImageName string, options *CommunityGalleryImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images/{galleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publicGalleryName == "" {
		return nil, errors.New("parameter publicGalleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicGalleryName}", url.PathEscape(publicGalleryName))
	if galleryImageName == "" {
		return nil, errors.New("parameter galleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryImageName}", url.PathEscape(galleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *CommunityGalleryImagesClient) getHandleResponse(resp *http.Response) (CommunityGalleryImagesClientGetResponse, error) {
	result := CommunityGalleryImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityGalleryImage); err != nil {
		return CommunityGalleryImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List community gallery images inside a gallery.
//
// Generated from API version 2022-08-03
//   - location - Resource location.
//   - publicGalleryName - The public name of the community gallery.
//   - options - CommunityGalleryImagesClientListOptions contains the optional parameters for the CommunityGalleryImagesClient.NewListPager
//     method.
func (client *CommunityGalleryImagesClient) NewListPager(location string, publicGalleryName string, options *CommunityGalleryImagesClientListOptions) *runtime.Pager[CommunityGalleryImagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[CommunityGalleryImagesClientListResponse]{
		More: func(page CommunityGalleryImagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *CommunityGalleryImagesClientListResponse) (CommunityGalleryImagesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, publicGalleryName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return CommunityGalleryImagesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return CommunityGalleryImagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return CommunityGalleryImagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *CommunityGalleryImagesClient) listCreateRequest(ctx context.Context, location string, publicGalleryName string, options *CommunityGalleryImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/communityGalleries/{publicGalleryName}/images"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if publicGalleryName == "" {
		return nil, errors.New("parameter publicGalleryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{publicGalleryName}", url.PathEscape(publicGalleryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CommunityGalleryImagesClient) listHandleResponse(resp *http.Response) (CommunityGalleryImagesClientListResponse, error) {
	result := CommunityGalleryImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CommunityGalleryImageList); err != nil {
		return CommunityGalleryImagesClientListResponse{}, err
	}
	return result, nil
}
