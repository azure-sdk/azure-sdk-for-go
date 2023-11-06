//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazuredatatransfer

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

// ListPendingConnectionsClient contains the methods for the ListPendingConnections group.
// Don't use this type directly, use NewListPendingConnectionsClient() instead.
type ListPendingConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewListPendingConnectionsClient creates a new instance of ListPendingConnectionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewListPendingConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ListPendingConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".ListPendingConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ListPendingConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListPager - Lists all pending connections for a connection.
//
// Generated from API version 2023-10-11-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - connectionName - The name for the connection that is to be requested.
//   - options - ListPendingConnectionsClientListOptions contains the optional parameters for the ListPendingConnectionsClient.NewListPager
//     method.
func (client *ListPendingConnectionsClient) NewListPager(resourceGroupName string, connectionName string, options *ListPendingConnectionsClientListOptions) *runtime.Pager[ListPendingConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ListPendingConnectionsClientListResponse]{
		More: func(page ListPendingConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ListPendingConnectionsClientListResponse) (ListPendingConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, connectionName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ListPendingConnectionsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ListPendingConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ListPendingConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ListPendingConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, connectionName string, options *ListPendingConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureDataTransfer/connections/{connectionName}/listPendingConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-11-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ListPendingConnectionsClient) listHandleResponse(resp *http.Response) (ListPendingConnectionsClientListResponse, error) {
	result := ListPendingConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PendingConnectionsListResult); err != nil {
		return ListPendingConnectionsClientListResponse{}, err
	}
	return result, nil
}
