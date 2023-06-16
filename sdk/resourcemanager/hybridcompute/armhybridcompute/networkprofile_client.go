//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armhybridcompute

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

// NetworkProfileClient contains the methods for the NetworkProfile group.
// Don't use this type directly, use NewNetworkProfileClient() instead.
type NetworkProfileClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkProfileClient creates a new instance of NetworkProfileClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkProfileClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkProfileClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkProfileClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkProfileClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - The operation to get network information of hybrid machine
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-25-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - machineName - The name of the hybrid machine.
//   - options - NetworkProfileClientGetOptions contains the optional parameters for the NetworkProfileClient.Get method.
func (client *NetworkProfileClient) Get(ctx context.Context, resourceGroupName string, machineName string, options *NetworkProfileClientGetOptions) (NetworkProfileClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, machineName, options)
	if err != nil {
		return NetworkProfileClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkProfileClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NetworkProfileClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NetworkProfileClient) getCreateRequest(ctx context.Context, resourceGroupName string, machineName string, options *NetworkProfileClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HybridCompute/machines/{machineName}/networkProfile"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if machineName == "" {
		return nil, errors.New("parameter machineName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{machineName}", url.PathEscape(machineName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-04-25-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *NetworkProfileClient) getHandleResponse(resp *http.Response) (NetworkProfileClientGetResponse, error) {
	result := NetworkProfileClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkProfile); err != nil {
		return NetworkProfileClientGetResponse{}, err
	}
	return result, nil
}
