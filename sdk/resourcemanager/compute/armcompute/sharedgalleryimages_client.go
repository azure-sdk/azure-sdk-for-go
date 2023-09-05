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

// SharedGalleryImagesClient contains the methods for the SharedGalleryImages group.
// Don't use this type directly, use NewSharedGalleryImagesClient() instead.
type SharedGalleryImagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSharedGalleryImagesClient creates a new instance of SharedGalleryImagesClient with the specified values.
//   - subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSharedGalleryImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SharedGalleryImagesClient, error) {
	cl, err := arm.NewClient(moduleName+".SharedGalleryImagesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SharedGalleryImagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a shared gallery image by subscription id or tenant id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-03
//   - location - Resource location.
//   - galleryUniqueName - The unique name of the Shared Gallery.
//   - galleryImageName - The name of the Shared Gallery Image Definition from which the Image Versions are to be listed.
//   - options - SharedGalleryImagesClientGetOptions contains the optional parameters for the SharedGalleryImagesClient.Get method.
func (client *SharedGalleryImagesClient) Get(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, options *SharedGalleryImagesClientGetOptions) (SharedGalleryImagesClientGetResponse, error) {
	var err error
	const operationName = "SharedGalleryImagesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, galleryUniqueName, galleryImageName, options)
	if err != nil {
		return SharedGalleryImagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SharedGalleryImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SharedGalleryImagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SharedGalleryImagesClient) getCreateRequest(ctx context.Context, location string, galleryUniqueName string, galleryImageName string, options *SharedGalleryImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images/{galleryImageName}"
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
func (client *SharedGalleryImagesClient) getHandleResponse(resp *http.Response) (SharedGalleryImagesClientGetResponse, error) {
	result := SharedGalleryImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryImage); err != nil {
		return SharedGalleryImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List shared gallery images by subscription id or tenant id.
//
// Generated from API version 2022-08-03
//   - location - Resource location.
//   - galleryUniqueName - The unique name of the Shared Gallery.
//   - options - SharedGalleryImagesClientListOptions contains the optional parameters for the SharedGalleryImagesClient.NewListPager
//     method.
func (client *SharedGalleryImagesClient) NewListPager(location string, galleryUniqueName string, options *SharedGalleryImagesClientListOptions) *runtime.Pager[SharedGalleryImagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[SharedGalleryImagesClientListResponse]{
		More: func(page SharedGalleryImagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SharedGalleryImagesClientListResponse) (SharedGalleryImagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SharedGalleryImagesClient.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, location, galleryUniqueName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SharedGalleryImagesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SharedGalleryImagesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SharedGalleryImagesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *SharedGalleryImagesClient) listCreateRequest(ctx context.Context, location string, galleryUniqueName string, options *SharedGalleryImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/sharedGalleries/{galleryUniqueName}/images"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
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
func (client *SharedGalleryImagesClient) listHandleResponse(resp *http.Response) (SharedGalleryImagesClientListResponse, error) {
	result := SharedGalleryImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SharedGalleryImageList); err != nil {
		return SharedGalleryImagesClientListResponse{}, err
	}
	return result, nil
}
