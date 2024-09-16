//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

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

// GitHubOwnersClient contains the methods for the GitHubOwners group.
// Don't use this type directly, use NewGitHubOwnersClient() instead.
type GitHubOwnersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGitHubOwnersClient creates a new instance of GitHubOwnersClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGitHubOwnersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GitHubOwnersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GitHubOwnersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns a monitored GitHub owner.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - ownerName - The GitHub owner name.
//   - options - GitHubOwnersClientGetOptions contains the optional parameters for the GitHubOwnersClient.Get method.
func (client *GitHubOwnersClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, ownerName string, options *GitHubOwnersClientGetOptions) (GitHubOwnersClientGetResponse, error) {
	var err error
	const operationName = "GitHubOwnersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, ownerName, options)
	if err != nil {
		return GitHubOwnersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GitHubOwnersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GitHubOwnersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GitHubOwnersClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ownerName string, options *GitHubOwnersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitHubOwners/{ownerName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	if ownerName == "" {
		return nil, errors.New("parameter ownerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ownerName}", url.PathEscape(ownerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GitHubOwnersClient) getHandleResponse(resp *http.Response) (GitHubOwnersClientGetResponse, error) {
	result := GitHubOwnersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubOwner); err != nil {
		return GitHubOwnersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of GitHub owners onboarded to the connector.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - GitHubOwnersClientListOptions contains the optional parameters for the GitHubOwnersClient.NewListPager method.
func (client *GitHubOwnersClient) NewListPager(resourceGroupName string, securityConnectorName string, options *GitHubOwnersClientListOptions) *runtime.Pager[GitHubOwnersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GitHubOwnersClientListResponse]{
		More: func(page GitHubOwnersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GitHubOwnersClientListResponse) (GitHubOwnersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GitHubOwnersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
			}, nil)
			if err != nil {
				return GitHubOwnersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GitHubOwnersClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *GitHubOwnersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitHubOwners"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GitHubOwnersClient) listHandleResponse(resp *http.Response) (GitHubOwnersClientListResponse, error) {
	result := GitHubOwnersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubOwnerListResponse); err != nil {
		return GitHubOwnersClientListResponse{}, err
	}
	return result, nil
}

// ListAvailable - Returns a list of all GitHub owners accessible by the user token consumed by the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - options - GitHubOwnersClientListAvailableOptions contains the optional parameters for the GitHubOwnersClient.ListAvailable
//     method.
func (client *GitHubOwnersClient) ListAvailable(ctx context.Context, resourceGroupName string, securityConnectorName string, options *GitHubOwnersClientListAvailableOptions) (GitHubOwnersClientListAvailableResponse, error) {
	var err error
	const operationName = "GitHubOwnersClient.ListAvailable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listAvailableCreateRequest(ctx, resourceGroupName, securityConnectorName, options)
	if err != nil {
		return GitHubOwnersClientListAvailableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GitHubOwnersClientListAvailableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GitHubOwnersClientListAvailableResponse{}, err
	}
	resp, err := client.listAvailableHandleResponse(httpResp)
	return resp, err
}

// listAvailableCreateRequest creates the ListAvailable request.
func (client *GitHubOwnersClient) listAvailableCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, options *GitHubOwnersClientListAvailableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/listAvailableGitHubOwners"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if securityConnectorName == "" {
		return nil, errors.New("parameter securityConnectorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{securityConnectorName}", url.PathEscape(securityConnectorName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAvailableHandleResponse handles the ListAvailable response.
func (client *GitHubOwnersClient) listAvailableHandleResponse(resp *http.Response) (GitHubOwnersClientListAvailableResponse, error) {
	result := GitHubOwnersClientListAvailableResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubOwnerListResponse); err != nil {
		return GitHubOwnersClientListAvailableResponse{}, err
	}
	return result, nil
}
