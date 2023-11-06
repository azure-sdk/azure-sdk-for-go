//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights

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

// ProactiveDetectionConfigurationsClient contains the methods for the ProactiveDetectionConfigurations group.
// Don't use this type directly, use NewProactiveDetectionConfigurationsClient() instead.
type ProactiveDetectionConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProactiveDetectionConfigurationsClient creates a new instance of ProactiveDetectionConfigurationsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProactiveDetectionConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProactiveDetectionConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ProactiveDetectionConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProactiveDetectionConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get the ProactiveDetection configuration for this configuration id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - configurationID - The ProactiveDetection configuration ID. This is unique within a Application Insights component.
//   - options - ProactiveDetectionConfigurationsClientGetOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.Get
//     method.
func (client *ProactiveDetectionConfigurationsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, options *ProactiveDetectionConfigurationsClientGetOptions) (ProactiveDetectionConfigurationsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, configurationID, options)
	if err != nil {
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ProactiveDetectionConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, options *ProactiveDetectionConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationID == "" {
		return nil, errors.New("parameter configurationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ConfigurationId}", url.PathEscape(configurationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProactiveDetectionConfigurationsClient) getHandleResponse(resp *http.Response) (ProactiveDetectionConfigurationsClientGetResponse, error) {
	result := ProactiveDetectionConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentProactiveDetectionConfiguration); err != nil {
		return ProactiveDetectionConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of ProactiveDetection configurations of an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - options - ProactiveDetectionConfigurationsClientListOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.List
//     method.
func (client *ProactiveDetectionConfigurationsClient) List(ctx context.Context, resourceGroupName string, resourceName string, options *ProactiveDetectionConfigurationsClientListOptions) (ProactiveDetectionConfigurationsClientListResponse, error) {
	var err error
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
	if err != nil {
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *ProactiveDetectionConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *ProactiveDetectionConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ProactiveDetectionConfigurationsClient) listHandleResponse(resp *http.Response) (ProactiveDetectionConfigurationsClientListResponse, error) {
	result := ProactiveDetectionConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentProactiveDetectionConfigurationArray); err != nil {
		return ProactiveDetectionConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// Update - Update the ProactiveDetection configuration for this configuration id.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - configurationID - The ProactiveDetection configuration ID. This is unique within a Application Insights component.
//   - proactiveDetectionProperties - Properties that need to be specified to update the ProactiveDetection configuration.
//   - options - ProactiveDetectionConfigurationsClientUpdateOptions contains the optional parameters for the ProactiveDetectionConfigurationsClient.Update
//     method.
func (client *ProactiveDetectionConfigurationsClient) Update(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, proactiveDetectionProperties ComponentProactiveDetectionConfiguration, options *ProactiveDetectionConfigurationsClientUpdateOptions) (ProactiveDetectionConfigurationsClientUpdateResponse, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, resourceGroupName, resourceName, configurationID, proactiveDetectionProperties, options)
	if err != nil {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *ProactiveDetectionConfigurationsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, configurationID string, proactiveDetectionProperties ComponentProactiveDetectionConfiguration, options *ProactiveDetectionConfigurationsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/components/{resourceName}/ProactiveDetectionConfigs/{ConfigurationId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationID == "" {
		return nil, errors.New("parameter configurationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{ConfigurationId}", url.PathEscape(configurationID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, proactiveDetectionProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *ProactiveDetectionConfigurationsClient) updateHandleResponse(resp *http.Response) (ProactiveDetectionConfigurationsClientUpdateResponse, error) {
	result := ProactiveDetectionConfigurationsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentProactiveDetectionConfiguration); err != nil {
		return ProactiveDetectionConfigurationsClientUpdateResponse{}, err
	}
	return result, nil
}
