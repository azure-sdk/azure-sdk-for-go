//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// RestorableDroppedDatabasesClient contains the methods for the RestorableDroppedDatabases group.
// Don't use this type directly, use NewRestorableDroppedDatabasesClient() instead.
type RestorableDroppedDatabasesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRestorableDroppedDatabasesClient creates a new instance of RestorableDroppedDatabasesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRestorableDroppedDatabasesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RestorableDroppedDatabasesClient, error) {
	cl, err := arm.NewClient(moduleName+".RestorableDroppedDatabasesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RestorableDroppedDatabasesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a restorable dropped database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - RestorableDroppedDatabasesClientGetOptions contains the optional parameters for the RestorableDroppedDatabasesClient.Get
//     method.
func (client *RestorableDroppedDatabasesClient) Get(ctx context.Context, resourceGroupName string, serverName string, restorableDroppedDatabaseID string, options *RestorableDroppedDatabasesClientGetOptions) (RestorableDroppedDatabasesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, restorableDroppedDatabaseID, options)
	if err != nil {
		return RestorableDroppedDatabasesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RestorableDroppedDatabasesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return RestorableDroppedDatabasesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *RestorableDroppedDatabasesClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, restorableDroppedDatabaseID string, options *RestorableDroppedDatabasesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases/{restorableDroppedDatabaseId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if restorableDroppedDatabaseID == "" {
		return nil, errors.New("parameter restorableDroppedDatabaseID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{restorableDroppedDatabaseId}", url.PathEscape(restorableDroppedDatabaseID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RestorableDroppedDatabasesClient) getHandleResponse(resp *http.Response) (RestorableDroppedDatabasesClientGetResponse, error) {
	result := RestorableDroppedDatabasesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedDatabase); err != nil {
		return RestorableDroppedDatabasesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Gets a list of restorable dropped databases.
//
// Generated from API version 2022-08-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serverName - The name of the server.
//   - options - RestorableDroppedDatabasesClientListByServerOptions contains the optional parameters for the RestorableDroppedDatabasesClient.NewListByServerPager
//     method.
func (client *RestorableDroppedDatabasesClient) NewListByServerPager(resourceGroupName string, serverName string, options *RestorableDroppedDatabasesClientListByServerOptions) *runtime.Pager[RestorableDroppedDatabasesClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[RestorableDroppedDatabasesClientListByServerResponse]{
		More: func(page RestorableDroppedDatabasesClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RestorableDroppedDatabasesClientListByServerResponse) (RestorableDroppedDatabasesClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return RestorableDroppedDatabasesClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return RestorableDroppedDatabasesClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return RestorableDroppedDatabasesClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *RestorableDroppedDatabasesClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *RestorableDroppedDatabasesClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/restorableDroppedDatabases"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *RestorableDroppedDatabasesClient) listByServerHandleResponse(resp *http.Response) (RestorableDroppedDatabasesClientListByServerResponse, error) {
	result := RestorableDroppedDatabasesClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RestorableDroppedDatabaseListResult); err != nil {
		return RestorableDroppedDatabasesClientListByServerResponse{}, err
	}
	return result, nil
}
