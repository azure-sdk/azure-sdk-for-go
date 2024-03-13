//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// ContentPackagesClient contains the methods for the ContentPackages group.
// Don't use this type directly, use NewContentPackagesClient() instead.
type ContentPackagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContentPackagesClient creates a new instance of ContentPackagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContentPackagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContentPackagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContentPackagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets an installed packages by its id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - packageID - package Id
//   - options - ContentPackagesClientGetOptions contains the optional parameters for the ContentPackagesClient.Get method.
func (client *ContentPackagesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, packageID string, options *ContentPackagesClientGetOptions) (ContentPackagesClientGetResponse, error) {
	var err error
	const operationName = "ContentPackagesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, packageID, options)
	if err != nil {
		return ContentPackagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentPackagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContentPackagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContentPackagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, packageID string, options *ContentPackagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentPackages/{packageId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if packageID == "" {
		return nil, errors.New("parameter packageID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packageId}", url.PathEscape(packageID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContentPackagesClient) getHandleResponse(resp *http.Response) (ContentPackagesClientGetResponse, error) {
	result := ContentPackagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageModel); err != nil {
		return ContentPackagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all installed packages.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - ContentPackagesClientListOptions contains the optional parameters for the ContentPackagesClient.NewListPager
//     method.
func (client *ContentPackagesClient) NewListPager(resourceGroupName string, workspaceName string, options *ContentPackagesClientListOptions) *runtime.Pager[ContentPackagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContentPackagesClientListResponse]{
		More: func(page ContentPackagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContentPackagesClientListResponse) (ContentPackagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContentPackagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return ContentPackagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ContentPackagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *ContentPackagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/contentPackages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Count != nil {
		reqQP.Set("$count", strconv.FormatBool(*options.Count))
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Search != nil {
		reqQP.Set("$search", *options.Search)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ContentPackagesClient) listHandleResponse(resp *http.Response) (ContentPackagesClientListResponse, error) {
	result := ContentPackagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PackageList); err != nil {
		return ContentPackagesClientListResponse{}, err
	}
	return result, nil
}
