// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmariadb

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

// RecoverableServersClient contains the methods for the RecoverableServers group.
// Don't use this type directly, use NewRecoverableServersClient() instead.
type RecoverableServersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRecoverableServersClient creates a new instance of RecoverableServersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRecoverableServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RecoverableServersClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RecoverableServersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a recoverable MariaDB Server.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2018-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serverName - The name of the server.
//   - options - RecoverableServersClientGetOptions contains the optional parameters for the RecoverableServersClient.Get method.
func (client *RecoverableServersClient) Get(ctx context.Context, resourceGroupName string, serverName string, options *RecoverableServersClientGetOptions) (RecoverableServersClientGetResponse, error) {
	var err error
	const operationName = "RecoverableServersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, options)
	if err != nil {
		return RecoverableServersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RecoverableServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RecoverableServersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RecoverableServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, _ *RecoverableServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMariaDB/servers/{serverName}/recoverableServers"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
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
	reqQP.Set("api-version", "2018-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RecoverableServersClient) getHandleResponse(resp *http.Response) (RecoverableServersClientGetResponse, error) {
	result := RecoverableServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecoverableServerResource); err != nil {
		return RecoverableServersClientGetResponse{}, err
	}
	return result, nil
}
