//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armavs

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

// ScriptPackagesClient contains the methods for the ScriptPackages group.
// Don't use this type directly, use NewScriptPackagesClient() instead.
type ScriptPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewScriptPackagesClient creates a new instance of ScriptPackagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewScriptPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ScriptPackagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ScriptPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a ScriptPackage
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - scriptPackageName - Name of the script package.
//   - options - ScriptPackagesClientGetOptions contains the optional parameters for the ScriptPackagesClient.Get method.
func (client *ScriptPackagesClient) Get(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string, options *ScriptPackagesClientGetOptions) (ScriptPackagesClientGetResponse, error) {
	var err error
	const operationName = "ScriptPackagesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, privateCloudName, scriptPackageName, options)
	if err != nil {
		return ScriptPackagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ScriptPackagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ScriptPackagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ScriptPackagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, scriptPackageName string, options *ScriptPackagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages/{scriptPackageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	if scriptPackageName == "" {
		return nil, errors.New("parameter scriptPackageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scriptPackageName}", url.PathEscape(scriptPackageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ScriptPackagesClient) getHandleResponse(resp *http.Response) (ScriptPackagesClientGetResponse, error) {
	result := ScriptPackagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptPackage); err != nil {
		return ScriptPackagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List ScriptPackage resources by PrivateCloud
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - privateCloudName - Name of the private cloud
//   - options - ScriptPackagesClientListOptions contains the optional parameters for the ScriptPackagesClient.NewListPager method.
func (client *ScriptPackagesClient) NewListPager(resourceGroupName string, privateCloudName string, options *ScriptPackagesClientListOptions) *runtime.Pager[ScriptPackagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ScriptPackagesClientListResponse]{
		More: func(page ScriptPackagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ScriptPackagesClientListResponse) (ScriptPackagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ScriptPackagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, privateCloudName, options)
			}, nil)
			if err != nil {
				return ScriptPackagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ScriptPackagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, privateCloudName string, options *ScriptPackagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AVS/privateClouds/{privateCloudName}/scriptPackages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if privateCloudName == "" {
		return nil, errors.New("parameter privateCloudName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{privateCloudName}", url.PathEscape(privateCloudName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ScriptPackagesClient) listHandleResponse(resp *http.Response) (ScriptPackagesClientListResponse, error) {
	result := ScriptPackagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScriptPackageListResult); err != nil {
		return ScriptPackagesClientListResponse{}, err
	}
	return result, nil
}
