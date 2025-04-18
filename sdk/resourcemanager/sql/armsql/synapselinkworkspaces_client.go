// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// SynapseLinkWorkspacesClient contains the methods for the SynapseLinkWorkspaces group.
// Don't use this type directly, use NewSynapseLinkWorkspacesClient() instead.
type SynapseLinkWorkspacesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSynapseLinkWorkspacesClient creates a new instance of SynapseLinkWorkspacesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSynapseLinkWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SynapseLinkWorkspacesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SynapseLinkWorkspacesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByDatabasePager - Gets all synapselink workspaces for a database.
//
// Generated from API version 2022-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - databaseName - The name of the database.
//   - options - SynapseLinkWorkspacesClientListByDatabaseOptions contains the optional parameters for the SynapseLinkWorkspacesClient.NewListByDatabasePager
//     method.
func (client *SynapseLinkWorkspacesClient) NewListByDatabasePager(resourceGroupName string, serverName string, databaseName string, options *SynapseLinkWorkspacesClientListByDatabaseOptions) *runtime.Pager[SynapseLinkWorkspacesClientListByDatabaseResponse] {
	return runtime.NewPager(runtime.PagingHandler[SynapseLinkWorkspacesClientListByDatabaseResponse]{
		More: func(page SynapseLinkWorkspacesClientListByDatabaseResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SynapseLinkWorkspacesClientListByDatabaseResponse) (SynapseLinkWorkspacesClientListByDatabaseResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SynapseLinkWorkspacesClient.NewListByDatabasePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByDatabaseCreateRequest(ctx, resourceGroupName, serverName, databaseName, options)
			}, nil)
			if err != nil {
				return SynapseLinkWorkspacesClientListByDatabaseResponse{}, err
			}
			return client.listByDatabaseHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByDatabaseCreateRequest creates the ListByDatabase request.
func (client *SynapseLinkWorkspacesClient) listByDatabaseCreateRequest(ctx context.Context, resourceGroupName string, serverName string, databaseName string, _ *SynapseLinkWorkspacesClientListByDatabaseOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/databases/{databaseName}/linkWorkspaces"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByDatabaseHandleResponse handles the ListByDatabase response.
func (client *SynapseLinkWorkspacesClient) listByDatabaseHandleResponse(resp *http.Response) (SynapseLinkWorkspacesClientListByDatabaseResponse, error) {
	result := SynapseLinkWorkspacesClientListByDatabaseResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SynapseLinkWorkspaceListResult); err != nil {
		return SynapseLinkWorkspacesClientListByDatabaseResponse{}, err
	}
	return result, nil
}
