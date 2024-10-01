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
	"strings"
)

// SourceControlClient contains the methods for the SourceControl group.
// Don't use this type directly, use NewSourceControlClient() instead.
type SourceControlClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSourceControlClient creates a new instance of SourceControlClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSourceControlClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SourceControlClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SourceControlClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListRepositoriesPager - Gets a list of repositories metadata.
//
// Generated from API version 2024-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - repositoryAccess - The repository access credentials.
//   - options - SourceControlClientListRepositoriesOptions contains the optional parameters for the SourceControlClient.NewListRepositoriesPager
//     method.
func (client *SourceControlClient) NewListRepositoriesPager(resourceGroupName string, workspaceName string, repositoryAccess RepositoryAccessProperties, options *SourceControlClientListRepositoriesOptions) *runtime.Pager[SourceControlClientListRepositoriesResponse] {
	return runtime.NewPager(runtime.PagingHandler[SourceControlClientListRepositoriesResponse]{
		More: func(page SourceControlClientListRepositoriesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SourceControlClientListRepositoriesResponse) (SourceControlClientListRepositoriesResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SourceControlClient.NewListRepositoriesPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listRepositoriesCreateRequest(ctx, resourceGroupName, workspaceName, repositoryAccess, options)
			}, nil)
			if err != nil {
				return SourceControlClientListRepositoriesResponse{}, err
			}
			return client.listRepositoriesHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listRepositoriesCreateRequest creates the ListRepositories request.
func (client *SourceControlClient) listRepositoriesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, repositoryAccess RepositoryAccessProperties, options *SourceControlClientListRepositoriesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/listRepositories"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, repositoryAccess); err != nil {
		return nil, err
	}
	return req, nil
}

// listRepositoriesHandleResponse handles the ListRepositories response.
func (client *SourceControlClient) listRepositoriesHandleResponse(resp *http.Response) (SourceControlClientListRepositoriesResponse, error) {
	result := SourceControlClientListRepositoriesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RepoList); err != nil {
		return SourceControlClientListRepositoriesResponse{}, err
	}
	return result, nil
}
