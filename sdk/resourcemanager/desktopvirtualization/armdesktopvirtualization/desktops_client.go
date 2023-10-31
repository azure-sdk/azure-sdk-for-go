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
	"strconv"
	"strings"
)

// DesktopsClient contains the methods for the Desktops group.
// Don't use this type directly, use NewDesktopsClient() instead.
type DesktopsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDesktopsClient creates a new instance of DesktopsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDesktopsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DesktopsClient, error) {
	cl, err := arm.NewClient(moduleName+".DesktopsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DesktopsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a desktop.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - desktopName - The name of the desktop within the specified desktop group
//   - options - DesktopsClientGetOptions contains the optional parameters for the DesktopsClient.Get method.
func (client *DesktopsClient) Get(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, options *DesktopsClientGetOptions) (DesktopsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, applicationGroupName, desktopName, options)
	if err != nil {
		return DesktopsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DesktopsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DesktopsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DesktopsClient) getCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, options *DesktopsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/desktops/{desktopName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	if desktopName == "" {
		return nil, errors.New("parameter desktopName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{desktopName}", url.PathEscape(desktopName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DesktopsClient) getHandleResponse(resp *http.Response) (DesktopsClientGetResponse, error) {
	result := DesktopsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Desktop); err != nil {
		return DesktopsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List desktops.
//
// Generated from API version 2023-10-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - options - DesktopsClientListOptions contains the optional parameters for the DesktopsClient.NewListPager method.
func (client *DesktopsClient) NewListPager(resourceGroupName string, applicationGroupName string, options *DesktopsClientListOptions) *runtime.Pager[DesktopsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DesktopsClientListResponse]{
		More: func(page DesktopsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DesktopsClientListResponse) (DesktopsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, applicationGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DesktopsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DesktopsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DesktopsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DesktopsClient) listCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, options *DesktopsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/desktops"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-04-preview")
	if options != nil && options.PageSize != nil {
		reqQP.Set("pageSize", strconv.FormatInt(int64(*options.PageSize), 10))
	}
	if options != nil && options.IsDescending != nil {
		reqQP.Set("isDescending", strconv.FormatBool(*options.IsDescending))
	}
	if options != nil && options.InitialSkip != nil {
		reqQP.Set("initialSkip", strconv.FormatInt(int64(*options.InitialSkip), 10))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DesktopsClient) listHandleResponse(resp *http.Response) (DesktopsClientListResponse, error) {
	result := DesktopsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DesktopList); err != nil {
		return DesktopsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update a desktop.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-04-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - applicationGroupName - The name of the application group
//   - desktopName - The name of the desktop within the specified desktop group
//   - options - DesktopsClientUpdateOptions contains the optional parameters for the DesktopsClient.Update method.
func (client *DesktopsClient) Update(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, options *DesktopsClientUpdateOptions) (DesktopsClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, applicationGroupName, desktopName, options)
	if err != nil {
		return DesktopsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DesktopsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DesktopsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *DesktopsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, applicationGroupName string, desktopName string, options *DesktopsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DesktopVirtualization/applicationGroups/{applicationGroupName}/desktops/{desktopName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if applicationGroupName == "" {
		return nil, errors.New("parameter applicationGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{applicationGroupName}", url.PathEscape(applicationGroupName))
	if desktopName == "" {
		return nil, errors.New("parameter desktopName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{desktopName}", url.PathEscape(desktopName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-04-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Desktop != nil {
		if err := runtime.MarshalAsJSON(req, *options.Desktop); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *DesktopsClient) updateHandleResponse(resp *http.Response) (DesktopsClientUpdateResponse, error) {
	result := DesktopsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Desktop); err != nil {
		return DesktopsClientUpdateResponse{}, err
	}
	return result, nil
}
