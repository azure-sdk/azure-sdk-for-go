//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armappplatform

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// ConfigServersClient contains the methods for the ConfigServers group.
// Don't use this type directly, use NewConfigServersClient() instead.
type ConfigServersClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConfigServersClient creates a new instance of ConfigServersClient with the specified values.
// subscriptionID - Gets subscription ID which uniquely identify the Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConfigServersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigServersClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigServersClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get the config server and its properties.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// options - ConfigServersClientGetOptions contains the optional parameters for the ConfigServersClient.Get method.
func (client *ConfigServersClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *ConfigServersClientGetOptions) (ConfigServersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return ConfigServersClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConfigServersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConfigServersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConfigServersClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *ConfigServersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigServersClient) getHandleResponse(resp *http.Response) (ConfigServersClientGetResponse, error) {
	result := ConfigServersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigServerResource); err != nil {
		return ConfigServersClientGetResponse{}, err
	}
	return result, nil
}

// BeginUpdatePatch - Update the config server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// configServerResource - Parameters for the update operation
// options - ConfigServersClientBeginUpdatePatchOptions contains the optional parameters for the ConfigServersClient.BeginUpdatePatch
// method.
func (client *ConfigServersClient) BeginUpdatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersClientBeginUpdatePatchOptions) (*runtime.Poller[ConfigServersClientUpdatePatchResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePatch(ctx, resourceGroupName, serviceName, configServerResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ConfigServersClientUpdatePatchResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConfigServersClientUpdatePatchResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdatePatch - Update the config server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
func (client *ConfigServersClient) updatePatch(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersClientBeginUpdatePatchOptions) (*http.Response, error) {
	req, err := client.updatePatchCreateRequest(ctx, resourceGroupName, serviceName, configServerResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updatePatchCreateRequest creates the UpdatePatch request.
func (client *ConfigServersClient) updatePatchCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersClientBeginUpdatePatchOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, configServerResource)
}

// BeginUpdatePut - Update the config server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// configServerResource - Parameters for the update operation
// options - ConfigServersClientBeginUpdatePutOptions contains the optional parameters for the ConfigServersClient.BeginUpdatePut
// method.
func (client *ConfigServersClient) BeginUpdatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersClientBeginUpdatePutOptions) (*runtime.Poller[ConfigServersClientUpdatePutResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.updatePut(ctx, resourceGroupName, serviceName, configServerResource, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[ConfigServersClientUpdatePutResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[ConfigServersClientUpdatePutResponse](options.ResumeToken, client.pl, nil)
	}
}

// UpdatePut - Update the config server.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
func (client *ConfigServersClient) updatePut(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersClientBeginUpdatePutOptions) (*http.Response, error) {
	req, err := client.updatePutCreateRequest(ctx, resourceGroupName, serviceName, configServerResource, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// updatePutCreateRequest creates the UpdatePut request.
func (client *ConfigServersClient) updatePutCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, configServerResource ConfigServerResource, options *ConfigServersClientBeginUpdatePutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/default"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, configServerResource)
}

// BeginValidate - Check if the config server settings are valid.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serviceName - The name of the Service resource.
// configServerSettings - Config server settings to be validated
// options - ConfigServersClientBeginValidateOptions contains the optional parameters for the ConfigServersClient.BeginValidate
// method.
func (client *ConfigServersClient) BeginValidate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings, options *ConfigServersClientBeginValidateOptions) (*runtime.Poller[ConfigServersClientValidateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.validate(ctx, resourceGroupName, serviceName, configServerSettings, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[ConfigServersClientValidateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ConfigServersClientValidateResponse](options.ResumeToken, client.pl, nil)
	}
}

// Validate - Check if the config server settings are valid.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-12-01
func (client *ConfigServersClient) validate(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings, options *ConfigServersClientBeginValidateOptions) (*http.Response, error) {
	req, err := client.validateCreateRequest(ctx, resourceGroupName, serviceName, configServerSettings, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// validateCreateRequest creates the Validate request.
func (client *ConfigServersClient) validateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, configServerSettings ConfigServerSettings, options *ConfigServersClientBeginValidateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AppPlatform/Spring/{serviceName}/configServers/validate"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-12-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, configServerSettings)
}
