//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkusto

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

// DatabaseClient contains the methods for the Database group.
// Don't use this type directly, use NewDatabaseClient() instead.
type DatabaseClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDatabaseClient creates a new instance of DatabaseClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDatabaseClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DatabaseClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DatabaseClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// InviteFollower - Generates an invitation token that allows attaching a follower database to this database.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-13
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - clusterName - The name of the Kusto cluster.
//   - databaseName - The name of the database in the Kusto cluster.
//   - parameters - The follower invitation request parameters.
//   - options - DatabaseClientInviteFollowerOptions contains the optional parameters for the DatabaseClient.InviteFollower method.
func (client *DatabaseClient) InviteFollower(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseInviteFollowerRequest, options *DatabaseClientInviteFollowerOptions) (DatabaseClientInviteFollowerResponse, error) {
	var err error
	const operationName = "DatabaseClient.InviteFollower"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.inviteFollowerCreateRequest(ctx, resourceGroupName, clusterName, databaseName, parameters, options)
	if err != nil {
		return DatabaseClientInviteFollowerResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DatabaseClientInviteFollowerResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DatabaseClientInviteFollowerResponse{}, err
	}
	resp, err := client.inviteFollowerHandleResponse(httpResp)
	return resp, err
}

// inviteFollowerCreateRequest creates the InviteFollower request.
func (client *DatabaseClient) inviteFollowerCreateRequest(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DatabaseInviteFollowerRequest, options *DatabaseClientInviteFollowerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/inviteFollower"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if clusterName == "" {
		return nil, errors.New("parameter clusterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{clusterName}", url.PathEscape(clusterName))
	if databaseName == "" {
		return nil, errors.New("parameter databaseName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{databaseName}", url.PathEscape(databaseName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-13")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// inviteFollowerHandleResponse handles the InviteFollower response.
func (client *DatabaseClient) inviteFollowerHandleResponse(resp *http.Response) (DatabaseClientInviteFollowerResponse, error) {
	result := DatabaseClientInviteFollowerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DatabaseInviteFollowerResult); err != nil {
		return DatabaseClientInviteFollowerResponse{}, err
	}
	return result, nil
}
