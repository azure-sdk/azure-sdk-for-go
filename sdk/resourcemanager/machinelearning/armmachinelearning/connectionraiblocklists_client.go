// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmachinelearning

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

// ConnectionRaiBlocklistsClient contains the methods for the ConnectionRaiBlocklists group.
// Don't use this type directly, use NewConnectionRaiBlocklistsClient() instead.
type ConnectionRaiBlocklistsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectionRaiBlocklistsClient creates a new instance of ConnectionRaiBlocklistsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectionRaiBlocklistsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectionRaiBlocklistsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectionRaiBlocklistsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Gets the custom blocklists associated with the Azure OpenAI connection.
//
// Generated from API version 2025-01-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Azure Machine Learning Workspace Name
//   - connectionName - Friendly name of the workspace connection
//   - options - ConnectionRaiBlocklistsClientListOptions contains the optional parameters for the ConnectionRaiBlocklistsClient.NewListPager
//     method.
func (client *ConnectionRaiBlocklistsClient) NewListPager(resourceGroupName string, workspaceName string, connectionName string, options *ConnectionRaiBlocklistsClientListOptions) *runtime.Pager[ConnectionRaiBlocklistsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectionRaiBlocklistsClientListResponse]{
		More: func(page ConnectionRaiBlocklistsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectionRaiBlocklistsClientListResponse) (ConnectionRaiBlocklistsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConnectionRaiBlocklistsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, options)
			}, nil)
			if err != nil {
				return ConnectionRaiBlocklistsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConnectionRaiBlocklistsClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, options *ConnectionRaiBlocklistsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}/raiBlocklists"
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
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-01-01-preview")
	if options != nil && options.ProxyAPIVersion != nil {
		reqQP.Set("proxy-api-version", *options.ProxyAPIVersion)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectionRaiBlocklistsClient) listHandleResponse(resp *http.Response) (ConnectionRaiBlocklistsClientListResponse, error) {
	result := ConnectionRaiBlocklistsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlocklistPropertiesBasicResourceArmPaginatedResult); err != nil {
		return ConnectionRaiBlocklistsClientListResponse{}, err
	}
	return result, nil
}
