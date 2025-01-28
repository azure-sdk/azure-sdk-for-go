//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// LinkedWorkspaceClient contains the methods for the LinkedWorkspace group.
// Don't use this type directly, use NewLinkedWorkspaceClient() instead.
type LinkedWorkspaceClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewLinkedWorkspaceClient creates a new instance of LinkedWorkspaceClient with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewLinkedWorkspaceClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*LinkedWorkspaceClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &LinkedWorkspaceClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Retrieve the linked workspace for the account id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-10-23
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - options - LinkedWorkspaceClientGetOptions contains the optional parameters for the LinkedWorkspaceClient.Get method.
func (client *LinkedWorkspaceClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, options *LinkedWorkspaceClientGetOptions) (LinkedWorkspaceClientGetResponse, error) {
	var err error
	const operationName = "LinkedWorkspaceClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, options)
	if err != nil {
		return LinkedWorkspaceClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return LinkedWorkspaceClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return LinkedWorkspaceClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *LinkedWorkspaceClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *LinkedWorkspaceClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/linkedWorkspace"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-10-23")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *LinkedWorkspaceClient) getHandleResponse(resp *http.Response) (LinkedWorkspaceClientGetResponse, error) {
	result := LinkedWorkspaceClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.LinkedWorkspace); err != nil {
		return LinkedWorkspaceClientGetResponse{}, err
	}
	return result, nil
}
