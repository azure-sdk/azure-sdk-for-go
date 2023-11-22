//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// GalleryAppSKUsClient contains the methods for the GalleryAppSKUs group.
// Don't use this type directly, use NewGalleryAppSKUsClient() instead.
type GalleryAppSKUsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGalleryAppSKUsClient creates a new instance of GalleryAppSKUsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGalleryAppSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GalleryAppSKUsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GalleryAppSKUsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a gallery application SKU for test runs under a Test Base account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - galleryAppName - The resource name of a gallery application.
//   - galleryAppSKUName - The resource name of a gallery application SKU.
//   - options - GalleryAppSKUsClientGetOptions contains the optional parameters for the GalleryAppSKUsClient.Get method.
func (client *GalleryAppSKUsClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, galleryAppName string, galleryAppSKUName string, options *GalleryAppSKUsClientGetOptions) (GalleryAppSKUsClientGetResponse, error) {
	var err error
	const operationName = "GalleryAppSKUsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, galleryAppName, galleryAppSKUName, options)
	if err != nil {
		return GalleryAppSKUsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GalleryAppSKUsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GalleryAppSKUsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GalleryAppSKUsClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, galleryAppName string, galleryAppSKUName string, options *GalleryAppSKUsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/galleryApps/{galleryAppName}/galleryAppSkus/{galleryAppSkuName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if galleryAppName == "" {
		return nil, errors.New("parameter galleryAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryAppName}", url.PathEscape(galleryAppName))
	if galleryAppSKUName == "" {
		return nil, errors.New("parameter galleryAppSKUName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryAppSkuName}", url.PathEscape(galleryAppSKUName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GalleryAppSKUsClient) getHandleResponse(resp *http.Response) (GalleryAppSKUsClientGetResponse, error) {
	result := GalleryAppSKUsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryAppSKUResource); err != nil {
		return GalleryAppSKUsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all SKUs of a gallery application currently available for test runs under a Test Base account.
//
// Generated from API version 2023-11-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - galleryAppName - The resource name of a gallery application.
//   - options - GalleryAppSKUsClientListOptions contains the optional parameters for the GalleryAppSKUsClient.NewListPager method.
func (client *GalleryAppSKUsClient) NewListPager(resourceGroupName string, testBaseAccountName string, galleryAppName string, options *GalleryAppSKUsClientListOptions) *runtime.Pager[GalleryAppSKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GalleryAppSKUsClientListResponse]{
		More: func(page GalleryAppSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GalleryAppSKUsClientListResponse) (GalleryAppSKUsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GalleryAppSKUsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, galleryAppName, options)
			}, nil)
			if err != nil {
				return GalleryAppSKUsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GalleryAppSKUsClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, galleryAppName string, options *GalleryAppSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/galleryApps/{galleryAppName}/galleryAppSkus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if galleryAppName == "" {
		return nil, errors.New("parameter galleryAppName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{galleryAppName}", url.PathEscape(galleryAppName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GalleryAppSKUsClient) listHandleResponse(resp *http.Response) (GalleryAppSKUsClientListResponse, error) {
	result := GalleryAppSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GalleryAppSKUListResult); err != nil {
		return GalleryAppSKUsClientListResponse{}, err
	}
	return result, nil
}
