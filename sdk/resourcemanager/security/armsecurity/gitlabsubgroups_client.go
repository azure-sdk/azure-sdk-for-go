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

// GitLabSubgroupsClient contains the methods for the GitLabSubgroups group.
// Don't use this type directly, use NewGitLabSubgroupsClient() instead.
type GitLabSubgroupsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGitLabSubgroupsClient creates a new instance of GitLabSubgroupsClient with the specified values.
//   - subscriptionID - Azure subscription ID
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGitLabSubgroupsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GitLabSubgroupsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GitLabSubgroupsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Gets nested subgroups of given GitLab Group which are onboarded to the connector.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - securityConnectorName - The security connector name.
//   - groupFQName - The GitLab group fully-qualified name.
//   - options - GitLabSubgroupsClientListOptions contains the optional parameters for the GitLabSubgroupsClient.List method.
func (client *GitLabSubgroupsClient) List(ctx context.Context, resourceGroupName string, securityConnectorName string, groupFQName string, options *GitLabSubgroupsClientListOptions) (GitLabSubgroupsClientListResponse, error) {
	var err error
	const operationName = "GitLabSubgroupsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, securityConnectorName, groupFQName, options)
	if err != nil {
		return GitLabSubgroupsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GitLabSubgroupsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GitLabSubgroupsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *GitLabSubgroupsClient) listCreateRequest(ctx context.Context, resourceGroupName string, securityConnectorName string, groupFQName string, _ *GitLabSubgroupsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Security/securityConnectors/{securityConnectorName}/devops/default/gitLabGroups/{groupFQName}/listSubgroups"
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
	if groupFQName == "" {
		return nil, errors.New("parameter groupFQName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupFQName}", url.PathEscape(groupFQName))
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

// listHandleResponse handles the List response.
func (client *GitLabSubgroupsClient) listHandleResponse(resp *http.Response) (GitLabSubgroupsClientListResponse, error) {
	result := GitLabSubgroupsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GitLabGroupListResponse); err != nil {
		return GitLabSubgroupsClientListResponse{}, err
	}
	return result, nil
}
