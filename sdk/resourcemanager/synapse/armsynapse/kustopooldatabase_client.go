//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// KustoPoolDatabaseClient contains the methods for the KustoPoolDatabase group.
// Don't use this type directly, use NewKustoPoolDatabaseClient() instead.
type KustoPoolDatabaseClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewKustoPoolDatabaseClient creates a new instance of KustoPoolDatabaseClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewKustoPoolDatabaseClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*KustoPoolDatabaseClient, error) {
	cl, err := arm.NewClient(moduleName+".KustoPoolDatabaseClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &KustoPoolDatabaseClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// InviteFollower - Generates an invitation token that allows attaching a follower database to this database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - kustoPoolName - The name of the Kusto pool.
//   - databaseName - The name of the database in the Kusto pool.
//   - parameters - The follower invitation request parameters.
//   - options - KustoPoolDatabaseClientInviteFollowerOptions contains the optional parameters for the KustoPoolDatabaseClient.InviteFollower
//     method.
func (client *KustoPoolDatabaseClient) InviteFollower(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DatabaseInviteFollowerRequest, options *KustoPoolDatabaseClientInviteFollowerOptions) (KustoPoolDatabaseClientInviteFollowerResponse, error) {
	req, err := client.inviteFollowerCreateRequest(ctx, resourceGroupName, workspaceName, kustoPoolName, databaseName, parameters, options)
	if err != nil {
		return KustoPoolDatabaseClientInviteFollowerResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return KustoPoolDatabaseClientInviteFollowerResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return KustoPoolDatabaseClientInviteFollowerResponse{}, runtime.NewResponseError(resp)
	}
	return client.inviteFollowerHandleResponse(resp)
}

// inviteFollowerCreateRequest creates the InviteFollower request.
func (client *KustoPoolDatabaseClient) inviteFollowerCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, kustoPoolName string, databaseName string, parameters DatabaseInviteFollowerRequest, options *KustoPoolDatabaseClientInviteFollowerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/kustoPools/{kustoPoolName}/databases/{databaseName}/inviteFollower"
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
	if kustoPoolName == "" {
		return nil, errors.New("parameter kustoPoolName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{kustoPoolName}", url.PathEscape(kustoPoolName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// inviteFollowerHandleResponse handles the InviteFollower response.
func (client *KustoPoolDatabaseClient) inviteFollowerHandleResponse(resp *http.Response) (KustoPoolDatabaseClientInviteFollowerResponse, error) {
	result := KustoPoolDatabaseClientInviteFollowerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseInviteFollowerResult); err != nil {
		return KustoPoolDatabaseClientInviteFollowerResponse{}, err
	}
	return result, nil
}
