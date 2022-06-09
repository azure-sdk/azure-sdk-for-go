//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SourceControlClient contains the methods for the SourceControl group.
// Don't use this type directly, use NewSourceControlClient() instead.
type SourceControlClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSourceControlClient creates a new instance of SourceControlClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSourceControlClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SourceControlClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &SourceControlClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListRepositoriesPager - Gets a list of repositories metadata.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-07-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// workspaceName - The name of the workspace.
// repoType - The repo type.
// options - SourceControlClientListRepositoriesOptions contains the optional parameters for the SourceControlClient.ListRepositories
// method.
func (client *SourceControlClient) NewListRepositoriesPager(resourceGroupName string, workspaceName string, repoType RepoType, options *SourceControlClientListRepositoriesOptions) *runtime.Pager[SourceControlClientListRepositoriesResponse] {
	return runtime.NewPager(runtime.PagingHandler[SourceControlClientListRepositoriesResponse]{
		More: func(page SourceControlClientListRepositoriesResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SourceControlClientListRepositoriesResponse) (SourceControlClientListRepositoriesResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listRepositoriesCreateRequest(ctx, resourceGroupName, workspaceName, repoType, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SourceControlClientListRepositoriesResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SourceControlClientListRepositoriesResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SourceControlClientListRepositoriesResponse{}, runtime.NewResponseError(resp)
			}
			return client.listRepositoriesHandleResponse(resp)
		},
	})
}

// listRepositoriesCreateRequest creates the ListRepositories request.
func (client *SourceControlClient) listRepositoriesCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, repoType RepoType, options *SourceControlClientListRepositoriesOptions) (*policy.Request, error) {
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, repoType)
}

// listRepositoriesHandleResponse handles the ListRepositories response.
func (client *SourceControlClient) listRepositoriesHandleResponse(resp *http.Response) (SourceControlClientListRepositoriesResponse, error) {
	result := SourceControlClientListRepositoriesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RepoList); err != nil {
		return SourceControlClientListRepositoriesResponse{}, err
	}
	return result, nil
}
