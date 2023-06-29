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

// IntegrationRuntimeNodeIPAddressClient contains the methods for the IntegrationRuntimeNodeIPAddress group.
// Don't use this type directly, use NewIntegrationRuntimeNodeIPAddressClient() instead.
type IntegrationRuntimeNodeIPAddressClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntegrationRuntimeNodeIPAddressClient creates a new instance of IntegrationRuntimeNodeIPAddressClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntegrationRuntimeNodeIPAddressClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntegrationRuntimeNodeIPAddressClient, error) {
	cl, err := arm.NewClient(moduleName+".IntegrationRuntimeNodeIPAddressClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntegrationRuntimeNodeIPAddressClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the IP address of an integration runtime node
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - integrationRuntimeName - Integration runtime name
//   - nodeName - Integration runtime node name
//   - options - IntegrationRuntimeNodeIPAddressClientGetOptions contains the optional parameters for the IntegrationRuntimeNodeIPAddressClient.Get
//     method.
func (client *IntegrationRuntimeNodeIPAddressClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodeIPAddressClientGetOptions) (IntegrationRuntimeNodeIPAddressClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, integrationRuntimeName, nodeName, options)
	if err != nil {
		return IntegrationRuntimeNodeIPAddressClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntegrationRuntimeNodeIPAddressClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return IntegrationRuntimeNodeIPAddressClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *IntegrationRuntimeNodeIPAddressClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, integrationRuntimeName string, nodeName string, options *IntegrationRuntimeNodeIPAddressClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Synapse/workspaces/{workspaceName}/integrationRuntimes/{integrationRuntimeName}/nodes/{nodeName}/ipAddress"
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
	if nodeName == "" {
		return nil, errors.New("parameter nodeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{nodeName}", url.PathEscape(nodeName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *IntegrationRuntimeNodeIPAddressClient) getHandleResponse(resp *http.Response) (IntegrationRuntimeNodeIPAddressClientGetResponse, error) {
	result := IntegrationRuntimeNodeIPAddressClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationRuntimeNodeIPAddress); err != nil {
		return IntegrationRuntimeNodeIPAddressClientGetResponse{}, err
	}
	return result, nil
}
