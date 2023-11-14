//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armspringappdiscovery

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

// SummariesClient contains the methods for the Summaries group.
// Don't use this type directly, use NewSummariesClient() instead.
type SummariesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSummariesClient creates a new instance of SummariesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSummariesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SummariesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SummariesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets the Summaries resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - summaryName - The name of summary
//   - options - SummariesClientGetOptions contains the optional parameters for the SummariesClient.Get method.
func (client *SummariesClient) Get(ctx context.Context, resourceGroupName string, siteName string, summaryName string, options *SummariesClientGetOptions) (SummariesClientGetResponse, error) {
	var err error
	const operationName = "SummariesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, siteName, summaryName, options)
	if err != nil {
		return SummariesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SummariesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SummariesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SummariesClient) getCreateRequest(ctx context.Context, resourceGroupName string, siteName string, summaryName string, options *SummariesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/summaries/{summaryName}"
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
	if summaryName == "" {
		return nil, errors.New("parameter summaryName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{summaryName}", url.PathEscape(summaryName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SummariesClient) getHandleResponse(resp *http.Response) (SummariesClientGetResponse, error) {
	result := SummariesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Summary); err != nil {
		return SummariesClientGetResponse{}, err
	}
	return result, nil
}

// NewListBySitePager - Lists the Summaries resource in springbootsites.
//
// Generated from API version 2023-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - siteName - The springbootsites name.
//   - options - SummariesClientListBySiteOptions contains the optional parameters for the SummariesClient.NewListBySitePager
//     method.
func (client *SummariesClient) NewListBySitePager(resourceGroupName string, siteName string, options *SummariesClientListBySiteOptions) *runtime.Pager[SummariesClientListBySiteResponse] {
	return runtime.NewPager(runtime.PagingHandler[SummariesClientListBySiteResponse]{
		More: func(page SummariesClientListBySiteResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SummariesClientListBySiteResponse) (SummariesClientListBySiteResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SummariesClient.NewListBySitePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySiteCreateRequest(ctx, resourceGroupName, siteName, options)
			}, nil)
			if err != nil {
				return SummariesClientListBySiteResponse{}, err
			}
			return client.listBySiteHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySiteCreateRequest creates the ListBySite request.
func (client *SummariesClient) listBySiteCreateRequest(ctx context.Context, resourceGroupName string, siteName string, options *SummariesClientListBySiteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzureSpringBoot/springbootsites/{siteName}/summaries"
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
	reqQP.Set("api-version", "2023-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listBySiteHandleResponse handles the ListBySite response.
func (client *SummariesClient) listBySiteHandleResponse(resp *http.Response) (SummariesClientListBySiteResponse, error) {
	result := SummariesClientListBySiteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SummaryList); err != nil {
		return SummariesClientListBySiteResponse{}, err
	}
	return result, nil
}
