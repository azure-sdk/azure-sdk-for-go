//go:build go1.18
// +build go1.18

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

// ConnectionRaiBlocklistItemClient contains the methods for the ConnectionRaiBlocklistItem group.
// Don't use this type directly, use NewConnectionRaiBlocklistItemClient() instead.
type ConnectionRaiBlocklistItemClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectionRaiBlocklistItemClient creates a new instance of ConnectionRaiBlocklistItemClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectionRaiBlocklistItemClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectionRaiBlocklistItemClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectionRaiBlocklistItemClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreate - Update the state of specified blocklist item associated with the Azure OpenAI connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - connectionName - Friendly name of the workspace connection
//   - raiBlocklistName - The name of the RaiBlocklist.
//   - raiBlocklistItemName - Name of the RaiBlocklist Item
//   - options - ConnectionRaiBlocklistItemClientBeginCreateOptions contains the optional parameters for the ConnectionRaiBlocklistItemClient.BeginCreate
//     method.
func (client *ConnectionRaiBlocklistItemClient) BeginCreate(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, raiBlocklistItemName string, body RaiBlocklistItemPropertiesBasicResource, options *ConnectionRaiBlocklistItemClientBeginCreateOptions) (*runtime.Poller[ConnectionRaiBlocklistItemClientCreateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.create(ctx, resourceGroupName, workspaceName, connectionName, raiBlocklistName, raiBlocklistItemName, body, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectionRaiBlocklistItemClientCreateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConnectionRaiBlocklistItemClientCreateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Create - Update the state of specified blocklist item associated with the Azure OpenAI connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ConnectionRaiBlocklistItemClient) create(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, raiBlocklistItemName string, body RaiBlocklistItemPropertiesBasicResource, options *ConnectionRaiBlocklistItemClientBeginCreateOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectionRaiBlocklistItemClient.BeginCreate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, raiBlocklistName, raiBlocklistItemName, body, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createCreateRequest creates the Create request.
func (client *ConnectionRaiBlocklistItemClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, raiBlocklistItemName string, body RaiBlocklistItemPropertiesBasicResource, options *ConnectionRaiBlocklistItemClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}/raiBlocklists/{raiBlocklistName}/raiBlocklistItems/{raiBlocklistItemName}"
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
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	if raiBlocklistItemName == "" {
		return nil, errors.New("parameter raiBlocklistItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistItemName}", url.PathEscape(raiBlocklistItemName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified custom blocklist item associated with the Azure OpenAI connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - connectionName - Friendly name of the workspace connection
//   - raiBlocklistName - The name of the RaiBlocklist.
//   - raiBlocklistItemName - Name of the RaiBlocklist Item
//   - options - ConnectionRaiBlocklistItemClientBeginDeleteOptions contains the optional parameters for the ConnectionRaiBlocklistItemClient.BeginDelete
//     method.
func (client *ConnectionRaiBlocklistItemClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, raiBlocklistItemName string, options *ConnectionRaiBlocklistItemClientBeginDeleteOptions) (*runtime.Poller[ConnectionRaiBlocklistItemClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, connectionName, raiBlocklistName, raiBlocklistItemName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectionRaiBlocklistItemClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ConnectionRaiBlocklistItemClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified custom blocklist item associated with the Azure OpenAI connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
func (client *ConnectionRaiBlocklistItemClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, raiBlocklistItemName string, options *ConnectionRaiBlocklistItemClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "ConnectionRaiBlocklistItemClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, raiBlocklistName, raiBlocklistItemName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionRaiBlocklistItemClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, raiBlocklistItemName string, options *ConnectionRaiBlocklistItemClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}/raiBlocklists/{raiBlocklistName}/raiBlocklistItems/{raiBlocklistItemName}"
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
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	if raiBlocklistItemName == "" {
		return nil, errors.New("parameter raiBlocklistItemName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistItemName}", url.PathEscape(raiBlocklistItemName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified custom blocklist associated with the Azure OpenAI connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - Name of Azure Machine Learning workspace.
//   - connectionName - Friendly name of the workspace connection
//   - raiBlocklistName - The name of the RaiBlocklist.
//   - options - ConnectionRaiBlocklistItemClientGetOptions contains the optional parameters for the ConnectionRaiBlocklistItemClient.Get
//     method.
func (client *ConnectionRaiBlocklistItemClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, options *ConnectionRaiBlocklistItemClientGetOptions) (ConnectionRaiBlocklistItemClientGetResponse, error) {
	var err error
	const operationName = "ConnectionRaiBlocklistItemClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, connectionName, raiBlocklistName, options)
	if err != nil {
		return ConnectionRaiBlocklistItemClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectionRaiBlocklistItemClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConnectionRaiBlocklistItemClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConnectionRaiBlocklistItemClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, connectionName string, raiBlocklistName string, options *ConnectionRaiBlocklistItemClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearningServices/workspaces/{workspaceName}/connections/{connectionName}/raiBlocklists/{raiBlocklistName}"
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
	if raiBlocklistName == "" {
		return nil, errors.New("parameter raiBlocklistName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{raiBlocklistName}", url.PathEscape(raiBlocklistName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionRaiBlocklistItemClient) getHandleResponse(resp *http.Response) (ConnectionRaiBlocklistItemClientGetResponse, error) {
	result := ConnectionRaiBlocklistItemClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RaiBlocklistPropertiesBasicResource); err != nil {
		return ConnectionRaiBlocklistItemClientGetResponse{}, err
	}
	return result, nil
}
