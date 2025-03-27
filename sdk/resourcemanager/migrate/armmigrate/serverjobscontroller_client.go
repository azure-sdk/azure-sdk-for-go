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

// ServerJobsControllerClient contains the methods for the ServerJobsController group.
// Don't use this type directly, use NewServerJobsControllerClient() instead.
type ServerJobsControllerClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewServerJobsControllerClient creates a new instance of ServerJobsControllerClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewServerJobsControllerClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerJobsControllerClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ServerJobsControllerClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a ServerJob
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - jobName - Jobs name
//   - options - ServerJobsControllerClientGetOptions contains the optional parameters for the ServerJobsControllerClient.Get
//     method.
func (client *ServerJobsControllerClient) Get(ctx context.Context, resourceGroupName string, siteName string, jobName string, options *ServerJobsControllerClientGetOptions) (ServerJobsControllerClientGetResponse, error) {
	var err error
	const operationName = "ServerJobsControllerClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, jobName, options)
	if err != nil {
		return ServerJobsControllerClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ServerJobsControllerClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ServerJobsControllerClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ServerJobsControllerClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, jobName string, _ *ServerJobsControllerClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/jobs/{jobName}"
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
	if jobName == "" {
		return nil, errors.New("parameter jobName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{jobName}", url.PathEscape(jobName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerJobsControllerClient) getHandleResponse(resp *http.Response) (ServerJobsControllerClientGetResponse, error) {
	result := ServerJobsControllerClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerJob); err != nil {
		return ServerJobsControllerClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerSiteResourcePager - List ServerJob resources by ServerSiteResource
//
// Generated from API version 2024-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - Site name
//   - options - ServerJobsControllerClientListByServerSiteResourceOptions contains the optional parameters for the ServerJobsControllerClient.NewListByServerSiteResourcePager
//     method.
func (client *ServerJobsControllerClient) NewListByServerSiteResourcePager(resourceGroupName string, siteName string, options *ServerJobsControllerClientListByServerSiteResourceOptions) *runtime.Pager[ServerJobsControllerClientListByServerSiteResourceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ServerJobsControllerClientListByServerSiteResourceResponse]{
		More: func(page ServerJobsControllerClientListByServerSiteResourceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerJobsControllerClientListByServerSiteResourceResponse) (ServerJobsControllerClientListByServerSiteResourceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ServerJobsControllerClient.NewListByServerSiteResourcePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServerSiteResourceCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return ServerJobsControllerClientListByServerSiteResourceResponse{}, err
			}
			return client.listByServerSiteResourceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServerSiteResourceCreateRequest creates the ListByServerSiteResource request.
func (client *ServerJobsControllerClient) listByServerSiteResourceCreateRequest(ctx context.Context, resourceGroupName string, siteName string, _ *ServerJobsControllerClientListByServerSiteResourceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/serverSites/{siteName}/jobs"
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
	reqQP.Set("api-version", "2024-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerSiteResourceHandleResponse handles the ListByServerSiteResource response.
func (client *ServerJobsControllerClient) listByServerSiteResourceHandleResponse(resp *http.Response) (ServerJobsControllerClientListByServerSiteResourceResponse, error) {
	result := ServerJobsControllerClientListByServerSiteResourceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerJobListResult); err != nil {
		return ServerJobsControllerClientListByServerSiteResourceResponse{}, err
	}
	return result, nil
}
