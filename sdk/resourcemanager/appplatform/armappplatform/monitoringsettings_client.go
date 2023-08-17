//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappplatform

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

// MonitoringSettingsClient contains the methods for the MonitoringSettings group.
// Don't use this type directly, use NewMonitoringSettingsClient() instead.
type MonitoringSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMonitoringSettingsClient creates a new instance of MonitoringSettingsClient with the specified values.
//   - subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
//     part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMonitoringSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MonitoringSettingsClient, error) {
	cl, err := arm.NewClient(moduleName+".MonitoringSettingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MonitoringSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the Monitoring Setting and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - options - MonitoringSettingsClientGetOptions contains the optional parameters for the MonitoringSettingsClient.Get method.
func (client *MonitoringSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *MonitoringSettingsClientGetOptions) (MonitoringSettingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return MonitoringSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MonitoringSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MonitoringSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MonitoringSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *MonitoringSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitoringSettingsClient) getHandleResponse(resp *http.Response) (MonitoringSettingsClientGetResponse, error) {
	result := MonitoringSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringSettingResource); err != nil {
		return MonitoringSettingsClientGetResponse{}, err
	}
	return result, nil
}

// BeginUpdatePatch - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - monitoringSettingResource - Parameters for the update operation
//   - options - MonitoringSettingsClientBeginUpdatePatchOptions contains the optional parameters for the MonitoringSettingsClient.BeginUpdatePatch
//     method.
func (client *MonitoringSettingsClient) BeginUpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePatchOptions) (*runtime.Poller[MonitoringSettingsClientUpdatePatchResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePatch(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[MonitoringSettingsClientUpdatePatchResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[MonitoringSettingsClientUpdatePatchResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdatePatch - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *MonitoringSettingsClient) updatePatch(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePatchOptions) (*http.Response, error) {
	var err error
	req, err := client.updatePatchCreateRequest(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updatePatchCreateRequest creates the UpdatePatch request.
func (client *MonitoringSettingsClient) updatePatchCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, monitoringSettingResource); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginUpdatePut - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
//     Resource Manager API or the portal.
//   - serviceName - The name of the Service resource.
//   - monitoringSettingResource - Parameters for the update operation
//   - options - MonitoringSettingsClientBeginUpdatePutOptions contains the optional parameters for the MonitoringSettingsClient.BeginUpdatePut
//     method.
func (client *MonitoringSettingsClient) BeginUpdatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePutOptions) (*runtime.Poller[MonitoringSettingsClientUpdatePutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePut(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[MonitoringSettingsClientUpdatePutResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[MonitoringSettingsClientUpdatePutResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// UpdatePut - Update the Monitoring Setting.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
func (client *MonitoringSettingsClient) updatePut(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePutOptions) (*http.Response, error) {
	var err error
	req, err := client.updatePutCreateRequest(ctx, resourceGroupName, serviceName, monitoringSettingResource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updatePutCreateRequest creates the UpdatePut request.
func (client *MonitoringSettingsClient) updatePutCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, monitoringSettingResource MonitoringSettingResource, options *MonitoringSettingsClientBeginUpdatePutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/monitoringSettings/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, monitoringSettingResource); err != nil {
		return nil, err
	}
	return req, nil
}
