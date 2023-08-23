//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armiotsecurity

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

// DevicesClient contains the methods for the Devices group.
// Don't use this type directly, use NewDevicesClient() instead.
type DevicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDevicesClient creates a new instance of DevicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDevicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DevicesClient, error) {
	cl, err := arm.NewClient(moduleName+".DevicesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DevicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get device
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01-preview
//   - iotDefenderLocation - Defender for IoT location
//   - deviceGroupName - Device group name
//   - deviceID - Device Id
//   - options - DevicesClientGetOptions contains the optional parameters for the DevicesClient.Get method.
func (client *DevicesClient) Get(ctx context.Context, iotDefenderLocation string, deviceGroupName string, deviceID string, options *DevicesClientGetOptions) (DevicesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, iotDefenderLocation, deviceGroupName, deviceID, options)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DevicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DevicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DevicesClient) getCreateRequest(ctx context.Context, iotDefenderLocation string, deviceGroupName string, deviceID string, options *DevicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/devices/{deviceId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if iotDefenderLocation == "" {
		return nil, errors.New("parameter iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	if deviceID == "" {
		return nil, errors.New("parameter deviceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceId}", url.PathEscape(deviceID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DevicesClient) getHandleResponse(resp *http.Response) (DevicesClientGetResponse, error) {
	result := DevicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceModel); err != nil {
		return DevicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List devices
//
// Generated from API version 2021-02-01-preview
//   - iotDefenderLocation - Defender for IoT location
//   - deviceGroupName - Device group name
//   - options - DevicesClientListOptions contains the optional parameters for the DevicesClient.NewListPager method.
func (client *DevicesClient) NewListPager(iotDefenderLocation string, deviceGroupName string, options *DevicesClientListOptions) *runtime.Pager[DevicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DevicesClientListResponse]{
		More: func(page DevicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DevicesClientListResponse) (DevicesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, iotDefenderLocation, deviceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return DevicesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return DevicesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return DevicesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *DevicesClient) listCreateRequest(ctx context.Context, iotDefenderLocation string, deviceGroupName string, options *DevicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.IoTSecurity/locations/{iotDefenderLocation}/deviceGroups/{deviceGroupName}/devices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if iotDefenderLocation == "" {
		return nil, errors.New("parameter iotDefenderLocation cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{iotDefenderLocation}", url.PathEscape(iotDefenderLocation))
	if deviceGroupName == "" {
		return nil, errors.New("parameter deviceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deviceGroupName}", url.PathEscape(deviceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01-preview")
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DevicesClient) listHandleResponse(resp *http.Response) (DevicesClientListResponse, error) {
	result := DevicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeviceList); err != nil {
		return DevicesClientListResponse{}, err
	}
	return result, nil
}
