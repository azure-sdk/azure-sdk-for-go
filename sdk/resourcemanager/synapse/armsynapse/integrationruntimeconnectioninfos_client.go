// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

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

// IntegrationRuntimeConnectionInfosClient contains the methods for the IntegrationRuntimeConnectionInfos group.
// Don't use this type directly, use NewIntegrationRuntimeConnectionInfosClient() instead.
type IntegrationRuntimeConnectionInfosClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationRuntimeConnectionInfosClient creates a new instance of IntegrationRuntimeConnectionInfosClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationRuntimeConnectionInfosClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeConnectionInfosClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationRuntimeConnectionInfosClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get connection info for an integration runtime
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - integrationRuntimeName - Integration runtime name
//   - options - IntegrationRuntimeConnectionInfosClientGetOptions contains the optional parameters for the IntegrationRuntimeConnectionInfosClient.Get
//     method.
func (client *IntegrationRuntimeConnectionInfosClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, options *IntegrationRuntimeConnectionInfosClientGetOptions) (IntegrationRuntimeConnectionInfosClientGetResponse, error) {
	var err error
	const operationName = "IntegrationRuntimeConnectionInfosClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, options)
	if err != nil {
		return IntegrationRuntimeConnectionInfosClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeConnectionInfosClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntegrationRuntimeConnectionInfosClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeConnectionInfosClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, _ *IntegrationRuntimeConnectionInfosClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/getConnectionInfo"
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
	if integrationRuntimeName == "" {
		return nil, errors.New("parameter integrationRuntimeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationRuntimeName}", url.PathEscape(integrationRuntimeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimeConnectionInfosClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeConnectionInfosClientGetResponse, error) {
	result := IntegrationRuntimeConnectionInfosClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeConnectionInfo); err != nil {
		return IntegrationRuntimeConnectionInfosClientGetResponse{}, err
	}
	return result, nil
}
