//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armrecoveryservicessiterecovery

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

// SupportedOperatingSystemsClient contains the methods for the SupportedOperatingSystems group.
// Don't use this type directly, use NewSupportedOperatingSystemsClient() instead.
type SupportedOperatingSystemsClient struct {
	internal          *arm.Client
	resourceName      string
	resourceGroupName string
	subscriptionID    string
}

// NewSupportedOperatingSystemsClient creates a new instance of SupportedOperatingSystemsClient with the specified values.
//   - resourceName - The name of the recovery services vault.
//   - resourceGroupName - The name of the resource group where the recovery services vault is present.
//   - subscriptionID - The subscription Id.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSupportedOperatingSystemsClient(resourceName string, resourceGroupName string, subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SupportedOperatingSystemsClient, error) {
	cl, err := arm.NewClient(moduleName+".SupportedOperatingSystemsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SupportedOperatingSystemsClient{
		resourceName:      resourceName,
		resourceGroupName: resourceGroupName,
		subscriptionID:    subscriptionID,
		internal:          cl,
	}
	return client, nil
}

// Get - Gets the data of supported operating systems by SRS.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-02-01
//   - options - SupportedOperatingSystemsClientGetOptions contains the optional parameters for the SupportedOperatingSystemsClient.Get
//     method.
func (client *SupportedOperatingSystemsClient) Get(ctx context.Context, options *SupportedOperatingSystemsClientGetOptions) (SupportedOperatingSystemsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return SupportedOperatingSystemsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SupportedOperatingSystemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SupportedOperatingSystemsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SupportedOperatingSystemsClient) getCreateRequest(ctx context.Context, options *SupportedOperatingSystemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/vaults/{resourceName}/replicationSupportedOperatingSystems"
	if client.resourceName == "" {
		return nil, errors.New("parameter client.resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(client.resourceName))
	if client.resourceGroupName == "" {
		return nil, errors.New("parameter client.resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(client.resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-02-01")
	if options != nil && options.InstanceType != nil {
		reqQP.Set("instanceType", *options.InstanceType)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SupportedOperatingSystemsClient) getHandleResponse(resp *http.Response) (SupportedOperatingSystemsClientGetResponse, error) {
	result := SupportedOperatingSystemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SupportedOperatingSystems); err != nil {
		return SupportedOperatingSystemsClientGetResponse{}, err
	}
	return result, nil
}
