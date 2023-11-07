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

// OSUpdatesClient contains the methods for the OSUpdates group.
// Don't use this type directly, use NewOSUpdatesClient() instead.
type OSUpdatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOSUpdatesClient creates a new instance of OSUpdatesClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOSUpdatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OSUpdatesClient, error) {
	cl, err := arm.NewClient(moduleName+".OSUpdatesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OSUpdatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an OS Update by name in which the package was tested before.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-04-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - osUpdateResourceName - The resource name of an OS Update.
//   - options - OSUpdatesClientGetOptions contains the optional parameters for the OSUpdatesClient.Get method.
func (client *OSUpdatesClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateResourceName string, options *OSUpdatesClientGetOptions) (OSUpdatesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, osUpdateResourceName, options)
	if err != nil {
		return OSUpdatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OSUpdatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return OSUpdatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OSUpdatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateResourceName string, options *OSUpdatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/osUpdates/{osUpdateResourceName}"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	if osUpdateResourceName == "" {
		return nil, errors.New("parameter osUpdateResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{osUpdateResourceName}", url.PathEscape(osUpdateResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OSUpdatesClient) getHandleResponse(resp *http.Response) (OSUpdatesClientGetResponse, error) {
	result := OSUpdatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSUpdateResource); err != nil {
		return OSUpdatesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists the OS Updates in which the package were tested before.
//
// Generated from API version 2022-04-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - packageName - The resource name of the Test Base Package.
//   - osUpdateType - The type of the OS Update.
//   - options - OSUpdatesClientListOptions contains the optional parameters for the OSUpdatesClient.NewListPager method.
func (client *OSUpdatesClient) NewListPager(resourceGroupName string, testBaseAccountName string, packageName string, osUpdateType OsUpdateType, options *OSUpdatesClientListOptions) *runtime.Pager[OSUpdatesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[OSUpdatesClientListResponse]{
		More: func(page OSUpdatesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *OSUpdatesClientListResponse) (OSUpdatesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, packageName, osUpdateType, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return OSUpdatesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return OSUpdatesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return OSUpdatesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *OSUpdatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, packageName string, osUpdateType OsUpdateType, options *OSUpdatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/packages/{packageName}/osUpdates"
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
	if packageName == "" {
		return nil, errors.New("parameter packageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageName}", url.PathEscape(packageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("osUpdateType", string(osUpdateType))
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *OSUpdatesClient) listHandleResponse(resp *http.Response) (OSUpdatesClientListResponse, error) {
	result := OSUpdatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.OSUpdateListResult); err != nil {
		return OSUpdatesClientListResponse{}, err
	}
	return result, nil
}
