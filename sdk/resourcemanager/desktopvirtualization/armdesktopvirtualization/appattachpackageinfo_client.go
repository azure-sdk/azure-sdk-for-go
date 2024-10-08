//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdesktopvirtualization

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

// AppAttachPackageInfoClient contains the methods for the AppAttachPackageInfo group.
// Don't use this type directly, use NewAppAttachPackageInfoClient() instead.
type AppAttachPackageInfoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAppAttachPackageInfoClient creates a new instance of AppAttachPackageInfoClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAppAttachPackageInfoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AppAttachPackageInfoClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AppAttachPackageInfoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewImportPager - Gets information from a package given the path to the package.
//
// Generated from API version 2024-04-03
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - hostPoolName - The name of the host pool within the specified resource group
//   - importPackageInfoRequest - Object containing URI to package image and other optional properties
//   - options - AppAttachPackageInfoClientImportOptions contains the optional parameters for the AppAttachPackageInfoClient.NewImportPager
//     method.
func (client *AppAttachPackageInfoClient) NewImportPager(resourceGroupName string, hostPoolName string, importPackageInfoRequest ImportPackageInfoRequest, options *AppAttachPackageInfoClientImportOptions) *runtime.Pager[AppAttachPackageInfoClientImportResponse] {
	return runtime.NewPager(runtime.PagingHandler[AppAttachPackageInfoClientImportResponse]{
		More: func(page AppAttachPackageInfoClientImportResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppAttachPackageInfoClientImportResponse) (AppAttachPackageInfoClientImportResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AppAttachPackageInfoClient.NewImportPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.importCreateRequest(ctx, resourceGroupName, hostPoolName, importPackageInfoRequest, options)
			}, nil)
			if err != nil {
				return AppAttachPackageInfoClientImportResponse{}, err
			}
			return client.importHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// importCreateRequest creates the Import request.
func (client *AppAttachPackageInfoClient) importCreateRequest(ctx context.Context, resourceGroupName string, hostPoolName string, importPackageInfoRequest ImportPackageInfoRequest, options *AppAttachPackageInfoClientImportOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/hostPools/{hostPoolName}/importAppAttachPackageInfo"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hostPoolName == "" {
		return nil, errors.New("parameter hostPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostPoolName}", url.PathEscape(hostPoolName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-03")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, importPackageInfoRequest); err != nil {
		return nil, err
	}
	return req, nil
}

// importHandleResponse handles the Import response.
func (client *AppAttachPackageInfoClient) importHandleResponse(resp *http.Response) (AppAttachPackageInfoClientImportResponse, error) {
	result := AppAttachPackageInfoClientImportResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppAttachPackageList); err != nil {
		return AppAttachPackageInfoClientImportResponse{}, err
	}
	return result, nil
}
