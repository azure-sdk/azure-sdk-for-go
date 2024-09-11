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

// RunAsAccountsControllerClient contains the methods for the RunAsAccountsController group.
// Don't use this type directly, use NewRunAsAccountsControllerClient() instead.
type RunAsAccountsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRunAsAccountsControllerClient creates a new instance of RunAsAccountsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRunAsAccountsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RunAsAccountsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RunAsAccountsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a VmwareRunAsAccountResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - accountName - RunAsAccounts name
//   - options - RunAsAccountsControllerClientGetOptions contains the optional parameters for the RunAsAccountsControllerClient.Get
//     method.
func (client *RunAsAccountsControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *RunAsAccountsControllerClientGetOptions) (RunAsAccountsControllerClientGetResponse, error) {
	var err error
	const operationName = "RunAsAccountsControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, accountName, options)
	if err != nil {
		return RunAsAccountsControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RunAsAccountsControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RunAsAccountsControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RunAsAccountsControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, accountName string, options *RunAsAccountsControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/runAsAccounts/{accountName}"
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
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
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
func (client *RunAsAccountsControllerClient) getHandleResponse(resp *http.Response) (RunAsAccountsControllerClientGetResponse, error) {
	result := RunAsAccountsControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VmwareRunAsAccountResource); err != nil {
		return RunAsAccountsControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVmwareSitePager - List VmwareRunAsAccountResource resources by VmwareSite
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - RunAsAccountsControllerClientListByVmwareSiteOptions contains the optional parameters for the RunAsAccountsControllerClient.NewListByVmwareSitePager
//     method.
func (client *RunAsAccountsControllerClient) NewListByVmwareSitePager(resourceGroupName string, siteName string, options *RunAsAccountsControllerClientListByVmwareSiteOptions) *runtime.Pager[RunAsAccountsControllerClientListByVmwareSiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[RunAsAccountsControllerClientListByVmwareSiteResponse]{
		More: func(page RunAsAccountsControllerClientListByVmwareSiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RunAsAccountsControllerClientListByVmwareSiteResponse) (RunAsAccountsControllerClientListByVmwareSiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RunAsAccountsControllerClient.NewListByVmwareSitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVmwareSiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return RunAsAccountsControllerClientListByVmwareSiteResponse{}, err
			}
			return client.listByVmwareSiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVmwareSiteCreateRequest creates the ListByVmwareSite request.
func (client *RunAsAccountsControllerClient) listByVmwareSiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *RunAsAccountsControllerClientListByVmwareSiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/vmwareSites/{siteName}/runAsAccounts"
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
func (client *RunAsAccountsControllerClient) listByVmwareSiteHandleResponse(resp *http.Response) (RunAsAccountsControllerClientListByVmwareSiteResponse, error) {
	result := RunAsAccountsControllerClientListByVmwareSiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VmwareRunAsAccountResourceListResult); err != nil {
		return RunAsAccountsControllerClientListByVmwareSiteResponse{}, err
	}
	return result, nil
}