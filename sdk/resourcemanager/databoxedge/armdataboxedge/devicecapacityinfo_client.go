//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// DeviceCapacityInfoClient contains the methods for the DeviceCapacityInfo group.
// Don't use this type directly, use NewDeviceCapacityInfoClient() instead.
type DeviceCapacityInfoClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeviceCapacityInfoClient creates a new instance of DeviceCapacityInfoClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeviceCapacityInfoClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeviceCapacityInfoClient, error) {
	cl, err := arm.NewClient(moduleName+".DeviceCapacityInfoClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeviceCapacityInfoClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetDeviceCapacityInfo - Gets the properties of the specified device capacity info.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-07-17
//   - resourceGroupName - The resource group name.
//   - deviceName - The device name.
//   - options - DeviceCapacityInfoClientGetDeviceCapacityInfoOptions contains the optional parameters for the DeviceCapacityInfoClient.GetDeviceCapacityInfo
//     method.
func (client *DeviceCapacityInfoClient) GetDeviceCapacityInfo(ctx context.Context, resourceGroupName string, deviceName string, options *DeviceCapacityInfoClientGetDeviceCapacityInfoOptions) (DeviceCapacityInfoClientGetDeviceCapacityInfoResponse, error) {
	var err error
	req, err := client.getDeviceCapacityInfoCreateRequest(ctx, resourceGroupName, deviceName, options)
	if err != nil {
		return DeviceCapacityInfoClientGetDeviceCapacityInfoResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeviceCapacityInfoClientGetDeviceCapacityInfoResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeviceCapacityInfoClientGetDeviceCapacityInfoResponse{}, err
	}
	resp, err := client.getDeviceCapacityInfoHandleResponse(httpResp)
	return resp, err
}

// getDeviceCapacityInfoCreateRequest creates the GetDeviceCapacityInfo request.
func (client *DeviceCapacityInfoClient) getDeviceCapacityInfoCreateRequest(ctx context.Context, resourceGroupName string, deviceName string, options *DeviceCapacityInfoClientGetDeviceCapacityInfoOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/deviceCapacityInfo/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deviceName == "" {
		return nil, errors.New("parameter deviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-07-17")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getDeviceCapacityInfoHandleResponse handles the GetDeviceCapacityInfo response.
func (client *DeviceCapacityInfoClient) getDeviceCapacityInfoHandleResponse(resp *http.Response) (DeviceCapacityInfoClientGetDeviceCapacityInfoResponse, error) {
	result := DeviceCapacityInfoClientGetDeviceCapacityInfoResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceCapacityInfo); err != nil {
		return DeviceCapacityInfoClientGetDeviceCapacityInfoResponse{}, err
	}
	return result, nil
}
