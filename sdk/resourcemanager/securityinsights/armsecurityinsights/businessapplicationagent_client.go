// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// BusinessApplicationAgentClient contains the methods for the BusinessApplicationAgent group.
// Don't use this type directly, use NewBusinessApplicationAgentClient() instead.
type BusinessApplicationAgentClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBusinessApplicationAgentClient creates a new instance of BusinessApplicationAgentClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBusinessApplicationAgentClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BusinessApplicationAgentClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BusinessApplicationAgentClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets Business Application Agent.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-04-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - agentResourceName - Business Application Agent Name
//   - options - BusinessApplicationAgentClientGetOptions contains the optional parameters for the BusinessApplicationAgentClient.Get
//     method.
func (client *BusinessApplicationAgentClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, options *BusinessApplicationAgentClientGetOptions) (BusinessApplicationAgentClientGetResponse, error) {
	var err error
	const operationName = "BusinessApplicationAgentClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, agentResourceName, options)
	if err != nil {
		return BusinessApplicationAgentClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BusinessApplicationAgentClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BusinessApplicationAgentClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BusinessApplicationAgentClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, agentResourceName string, _ *BusinessApplicationAgentClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/businessApplicationAgents/{agentResourceName}"
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
	if agentResourceName == "" {
		return nil, errors.New("parameter agentResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{agentResourceName}", url.PathEscape(agentResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BusinessApplicationAgentClient) getHandleResponse(resp *http.Response) (BusinessApplicationAgentClientGetResponse, error) {
	result := BusinessApplicationAgentClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BusinessApplicationAgentResource); err != nil {
		return BusinessApplicationAgentClientGetResponse{}, err
	}
	return result, nil
}
