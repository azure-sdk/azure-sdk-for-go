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

// GitHubReposClient contains the methods for the GitHubRepos group.
// Don't use this type directly, use NewGitHubReposClient() instead.
type GitHubReposClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGitHubReposClient creates a new instance of GitHubReposClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGitHubReposClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GitHubReposClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GitHubReposClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Returns a monitored GitHub repository.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - ownerName - The GitHub owner name.
//   - repoName - The repository name.
//   - options - GitHubReposClientGetOptions contains the optional parameters for the GitHubReposClient.Get method.
func (client *GitHubReposClient) Get(ctx context.Context, resourceGroupName string, securityConnectorName string, ownerName string, repoName string, options *GitHubReposClientGetOptions) (GitHubReposClientGetResponse, error) {
	var err error
	const operationName = "GitHubReposClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, securityConnectorName, ownerName, repoName, options)
	if err != nil {
		return GitHubReposClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GitHubReposClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GitHubReposClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GitHubReposClient) getCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ownerName string, repoName string, _ *GitHubReposClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitHubOwners/{ownerName}/repos/{repoName}"
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
	if repoName == "" {
		return nil, errors.New("parameter repoName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{repoName}", url.PathEscape(repoName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GitHubReposClient) getHandleResponse(resp *http.Response) (GitHubReposClientGetResponse, error) {
	result := GitHubReposClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubRepository); err != nil {
		return GitHubReposClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Returns a list of GitHub repositories onboarded to the connector.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - ownerName - The GitHub owner name.
//   - options - GitHubReposClientListOptions contains the optional parameters for the GitHubReposClient.NewListPager method.
func (client *GitHubReposClient) NewListPager(resourceGroupName string, securityConnectorName string, ownerName string, options *GitHubReposClientListOptions) *runtime.Pager[GitHubReposClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[GitHubReposClientListResponse]{
		More: func(page GitHubReposClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GitHubReposClientListResponse) (GitHubReposClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GitHubReposClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, ownerName, options)
			}, nil)
			if err != nil {
				return GitHubReposClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *GitHubReposClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, ownerName string, _ *GitHubReposClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitHubOwners/{ownerName}/repos"
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
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *GitHubReposClient) listHandleResponse(resp *http.Response) (GitHubReposClientListResponse, error) {
	result := GitHubReposClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitHubRepositoryListResponse); err != nil {
		return GitHubReposClientListResponse{}, err
	}
	return result, nil
}
