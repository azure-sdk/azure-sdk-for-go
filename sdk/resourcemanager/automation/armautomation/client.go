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

// Client contains the methods for the AutomationClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
//     forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName+".Client", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// ConvertGraphRunbookContent - Post operation to serialize or deserialize GraphRunbookContent
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-08-01
//   - resourceGroupName - Name of an Azure Resource group.
//   - automationAccountName - The name of the automation account.
//   - parameters - Input data describing the graphical runbook.
//   - options - ClientConvertGraphRunbookContentOptions contains the optional parameters for the Client.ConvertGraphRunbookContent
//     method.
func (client *Client) ConvertGraphRunbookContent(ctx context.Context, resourceGroupName string, automationAccountName string, parameters GraphicalRunbookContent, options *ClientConvertGraphRunbookContentOptions) (ClientConvertGraphRunbookContentResponse, error) {
	var err error
	req, err := client.convertGraphRunbookContentCreateRequest(ctx, resourceGroupName, automationAccountName, parameters, options)
	if err != nil {
		return ClientConvertGraphRunbookContentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientConvertGraphRunbookContentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientConvertGraphRunbookContentResponse{}, err
	}
	resp, err := client.convertGraphRunbookContentHandleResponse(httpResp)
	return resp, err
}

// convertGraphRunbookContentCreateRequest creates the ConvertGraphRunbookContent request.
func (client *Client) convertGraphRunbookContentCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, parameters GraphicalRunbookContent, options *ClientConvertGraphRunbookContentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/convertGraphRunbookContent"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// convertGraphRunbookContentHandleResponse handles the ConvertGraphRunbookContent response.
func (client *Client) convertGraphRunbookContentHandleResponse(resp *http.Response) (ClientConvertGraphRunbookContentResponse, error) {
	result := ClientConvertGraphRunbookContentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GraphicalRunbookContent); err != nil {
		return ClientConvertGraphRunbookContentResponse{}, err
	}
	return result, nil
}
