//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armeventhub

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

// NetworkSecurityPerimeterConfigurationClient contains the methods for the NetworkSecurityPerimeterConfiguration group.
// Don't use this type directly, use NewNetworkSecurityPerimeterConfigurationClient() instead.
type NetworkSecurityPerimeterConfigurationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewNetworkSecurityPerimeterConfigurationClient creates a new instance of NetworkSecurityPerimeterConfigurationClient with the specified values.
//   - subscriptionID - Subscription credentials that uniquely identify a Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewNetworkSecurityPerimeterConfigurationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*NetworkSecurityPerimeterConfigurationClient, error) {
	cl, err := arm.NewClient(moduleName+".NetworkSecurityPerimeterConfigurationClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &NetworkSecurityPerimeterConfigurationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Gets list of current NetworkSecurityPerimeterConfiguration for Namespace
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - Name of the resource group within the azure subscription.
//   - namespaceName - The Namespace name
//   - options - NetworkSecurityPerimeterConfigurationClientListOptions contains the optional parameters for the NetworkSecurityPerimeterConfigurationClient.List
//     method.
func (client *NetworkSecurityPerimeterConfigurationClient) List(ctx context.Context, resourceGroupName string, namespaceName string, options *NetworkSecurityPerimeterConfigurationClientListOptions) (NetworkSecurityPerimeterConfigurationClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, resourceGroupName, namespaceName, options)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return NetworkSecurityPerimeterConfigurationClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return NetworkSecurityPerimeterConfigurationClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *NetworkSecurityPerimeterConfigurationClient) listCreateRequest(ctx context.Context, resourceGroupName string, namespaceName string, options *NetworkSecurityPerimeterConfigurationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/networkSecurityPerimeterConfigurations"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if namespaceName == "" {
		return nil, errors.New("parameter namespaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{namespaceName}", url.PathEscape(namespaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *NetworkSecurityPerimeterConfigurationClient) listHandleResponse(resp *http.Response) (NetworkSecurityPerimeterConfigurationClientListResponse, error) {
	result := NetworkSecurityPerimeterConfigurationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.NetworkSecurityPerimeterConfigurationList); err != nil {
		return NetworkSecurityPerimeterConfigurationClientListResponse{}, err
	}
	return result, nil
}
