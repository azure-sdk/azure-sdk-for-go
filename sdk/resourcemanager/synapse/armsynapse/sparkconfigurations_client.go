//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsynapse

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

// SparkConfigurationsClient contains the methods for the SparkConfigurations group.
// Don't use this type directly, use NewSparkConfigurationsClient() instead.
type SparkConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSparkConfigurationsClient creates a new instance of SparkConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSparkConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SparkConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".SparkConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SparkConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByWorkspacePager - List sparkConfigurations in a workspace.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - SparkConfigurationsClientListByWorkspaceOptions contains the optional parameters for the SparkConfigurationsClient.NewListByWorkspacePager
//     method.
func (client *SparkConfigurationsClient) NewListByWorkspacePager(resourceGroupName string, workspaceName string, options *SparkConfigurationsClientListByWorkspaceOptions) *runtime.Pager[SparkConfigurationsClientListByWorkspaceResponse] {
	return runtime.NewPager(runtime.PagingHandler[SparkConfigurationsClientListByWorkspaceResponse]{
		More: func(page SparkConfigurationsClientListByWorkspaceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SparkConfigurationsClientListByWorkspaceResponse) (SparkConfigurationsClientListByWorkspaceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByWorkspaceCreateRequest(ctx, resourceGroupName, workspaceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SparkConfigurationsClientListByWorkspaceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return SparkConfigurationsClientListByWorkspaceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SparkConfigurationsClientListByWorkspaceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByWorkspaceHandleResponse(resp)
		},
	})
}

// listByWorkspaceCreateRequest creates the ListByWorkspace request.
func (client *SparkConfigurationsClient) listByWorkspaceCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SparkConfigurationsClientListByWorkspaceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/sparkconfigurations"
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
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByWorkspaceHandleResponse handles the ListByWorkspace response.
func (client *SparkConfigurationsClient) listByWorkspaceHandleResponse(resp *http.Response) (SparkConfigurationsClientListByWorkspaceResponse, error) {
	result := SparkConfigurationsClientListByWorkspaceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SparkConfigurationListResponse); err != nil {
		return SparkConfigurationsClientListByWorkspaceResponse{}, err
	}
	return result, nil
}
