// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

package armdevopsinfrastructure

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

// ImageVersionsClient contains the methods for the ImageVersions group.
// Don't use this type directly, use NewImageVersionsClient() instead.
type ImageVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewImageVersionsClient creates a new instance of ImageVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewImageVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ImageVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ImageVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByImagePager - List ImageVersion resources by Image
//
// Generated from API version 2025-01-21
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - imageName - Name of the image.
//   - options - ImageVersionsClientListByImageOptions contains the optional parameters for the ImageVersionsClient.NewListByImagePager
//     method.
func (client *ImageVersionsClient) NewListByImagePager(resourceGroupName string, imageName string, options *ImageVersionsClientListByImageOptions) *runtime.Pager[ImageVersionsClientListByImageResponse] {
	return runtime.NewPager(runtime.PagingHandler[ImageVersionsClientListByImageResponse]{
		More: func(page ImageVersionsClientListByImageResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ImageVersionsClientListByImageResponse) (ImageVersionsClientListByImageResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ImageVersionsClient.NewListByImagePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByImageCreateRequest(ctx, resourceGroupName, imageName, options)
			}, nil)
			if err != nil {
				return ImageVersionsClientListByImageResponse{}, err
			}
			return client.listByImageHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByImageCreateRequest creates the ListByImage request.
func (client *ImageVersionsClient) listByImageCreateRequest(ctx context.Context, resourceGroupName string, imageName string, _ *ImageVersionsClientListByImageOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevOpsInfrastructure/images/{imageName}/versions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if imageName == "" {
		return nil, errors.New("parameter imageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{imageName}", url.PathEscape(imageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-21")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByImageHandleResponse handles the ListByImage response.
func (client *ImageVersionsClient) listByImageHandleResponse(resp *http.Response) (ImageVersionsClientListByImageResponse, error) {
	result := ImageVersionsClientListByImageResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ImageVersionListResult); err != nil {
		return ImageVersionsClientListByImageResponse{}, err
	}
	return result, nil
}
