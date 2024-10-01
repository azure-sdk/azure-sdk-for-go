//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmigrate

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

// VmwareHostControllerClient contains the methods for the VmwareHostController group.
// Don't use this type directly, use NewVmwareHostControllerClient() instead.
type VmwareHostControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVmwareHostControllerClient creates a new instance of VmwareHostControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVmwareHostControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VmwareHostControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VmwareHostControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a VmwareHost
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - hostName - Hosts name
//   - options - VmwareHostControllerClientGetOptions contains the optional parameters for the VmwareHostControllerClient.Get
//     method.
func (client *VmwareHostControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *VmwareHostControllerClientGetOptions) (VmwareHostControllerClientGetResponse, error) {
	var err error
	const operationName = "VmwareHostControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, hostName, options)
	if err != nil {
		return VmwareHostControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VmwareHostControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VmwareHostControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VmwareHostControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, hostName string, options *VmwareHostControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/hosts/{hostName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	if hostName == "" {
		return nil, errors.New("parameter hostName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hostName}", url.PathEscape(hostName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VmwareHostControllerClient) getHandleResponse(resp *http.Response) (VmwareHostControllerClientGetResponse, error) {
	result := VmwareHostControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VmwareHost); err != nil {
		return VmwareHostControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVmwareSitePager - List VmwareHost resources by VmwareSite
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - VmwareHostControllerClientListByVmwareSiteOptions contains the optional parameters for the VmwareHostControllerClient.NewListByVmwareSitePager
//     method.
func (client *VmwareHostControllerClient) NewListByVmwareSitePager(resourceGroupName string, siteName string, options *VmwareHostControllerClientListByVmwareSiteOptions) *runtime.Pager[VmwareHostControllerClientListByVmwareSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[VmwareHostControllerClientListByVmwareSiteResponse]{
		More: func(page VmwareHostControllerClientListByVmwareSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VmwareHostControllerClientListByVmwareSiteResponse) (VmwareHostControllerClientListByVmwareSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VmwareHostControllerClient.NewListByVmwareSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVmwareSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return VmwareHostControllerClientListByVmwareSiteResponse{}, err
			}
			return client.listByVmwareSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVmwareSiteCreateRequest creates the ListByVmwareSite request.
func (client *VmwareHostControllerClient) listByVmwareSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *VmwareHostControllerClientListByVmwareSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/hosts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if siteName == "" {
		return nil, errors.New("parameter siteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{siteName}", url.PathEscape(siteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByVmwareSiteHandleResponse handles the ListByVmwareSite response.
func (client *VmwareHostControllerClient) listByVmwareSiteHandleResponse(resp *http.Response) (VmwareHostControllerClientListByVmwareSiteResponse, error) {
	result := VmwareHostControllerClientListByVmwareSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VmwareHostListResult); err != nil {
		return VmwareHostControllerClientListByVmwareSiteResponse{}, err
	}
	return result, nil
}
