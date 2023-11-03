//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpostgresqlflexibleservers

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

// ReplicasClient contains the methods for the Replicas group.
// Don't use this type directly, use NewReplicasClient() instead.
type ReplicasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewReplicasClient creates a new instance of ReplicasClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewReplicasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ReplicasClient, error) {
	cl, err := arm.NewClient(moduleName+".ReplicasClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ReplicasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// NewListByServerPager - List all the replicas for a given server.
//
// Generated from API version 2023-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - ReplicasClientListByServerOptions contains the optional parameters for the ReplicasClient.NewListByServerPager
//     method.
func (client *ReplicasClient) NewListByServerPager(resourceGroupName string, serverName string, options *ReplicasClientListByServerOptions) *runtime.Pager[ReplicasClientListByServerResponse] {
	return runtime.NewPager(runtime.PagingHandler[ReplicasClientListByServerResponse]{
		More: func(page ReplicasClientListByServerResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ReplicasClientListByServerResponse) (ReplicasClientListByServerResponse, error) {
			req, err := client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			if err != nil {
				return ReplicasClientListByServerResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ReplicasClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ReplicasClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ReplicasClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ReplicasClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/replicas"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ReplicasClient) listByServerHandleResponse(resp *http.Response) (ReplicasClientListByServerResponse, error) {
	result := ReplicasClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerListResult); err != nil {
		return ReplicasClientListByServerResponse{}, err
	}
	return result, nil
}
