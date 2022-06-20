//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// TenantAccessClient contains the methods for the TenantAccess group.
// Don't use this type directly, use NewTenantAccessClient() instead.
type TenantAccessClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewTenantAccessClient creates a new instance of TenantAccessClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTenantAccessClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TenantAccessClient, error) {
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
	client := &TenantAccessClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Get tenant access information details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientGetOptions contains the optional parameters for the TenantAccessClient.Get method.
func (client *TenantAccessClient) Get(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientGetOptions) (TenantAccessClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return TenantAccessClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TenantAccessClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TenantAccessClient) getHandleResponse(resp *http.Response) (TenantAccessClientGetResponse, error) {
	result := TenantAccessClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AccessInformationContract); err != nil {
		return TenantAccessClientGetResponse{}, err
	}
	return result, nil
}

// RegeneratePrimaryKey - Regenerate primary access key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientRegeneratePrimaryKeyOptions contains the optional parameters for the TenantAccessClient.RegeneratePrimaryKey
// method.
func (client *TenantAccessClient) RegeneratePrimaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegeneratePrimaryKeyOptions) (TenantAccessClientRegeneratePrimaryKeyResponse, error) {
	req, err := client.regeneratePrimaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientRegeneratePrimaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientRegeneratePrimaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessClientRegeneratePrimaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessClientRegeneratePrimaryKeyResponse{}, nil
}

// regeneratePrimaryKeyCreateRequest creates the RegeneratePrimaryKey request.
func (client *TenantAccessClient) regeneratePrimaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegeneratePrimaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regeneratePrimaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// RegenerateSecondaryKey - Regenerate secondary access key.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// options - TenantAccessClientRegenerateSecondaryKeyOptions contains the optional parameters for the TenantAccessClient.RegenerateSecondaryKey
// method.
func (client *TenantAccessClient) RegenerateSecondaryKey(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegenerateSecondaryKeyOptions) (TenantAccessClientRegenerateSecondaryKeyResponse, error) {
	req, err := client.regenerateSecondaryKeyCreateRequest(ctx, resourceGroupName, serviceName, accessName, options)
	if err != nil {
		return TenantAccessClientRegenerateSecondaryKeyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientRegenerateSecondaryKeyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessClientRegenerateSecondaryKeyResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessClientRegenerateSecondaryKeyResponse{}, nil
}

// regenerateSecondaryKeyCreateRequest creates the RegenerateSecondaryKey request.
func (client *TenantAccessClient) regenerateSecondaryKeyCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, options *TenantAccessClientRegenerateSecondaryKeyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}/regenerateSecondaryKey"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Update - Update tenant access information details.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-09
// resourceGroupName - The name of the resource group.
// serviceName - The name of the API Management service.
// accessName - The identifier of the Access configuration.
// ifMatch - The entity state (Etag) version of the property to update. A value of "*" can be used for If-Match to unconditionally
// apply the operation.
// parameters - Parameters supplied to retrieve the Tenant Access Information.
// options - TenantAccessClientUpdateOptions contains the optional parameters for the TenantAccessClient.Update method.
func (client *TenantAccessClient) Update(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationUpdateParameters, options *TenantAccessClientUpdateOptions) (TenantAccessClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, accessName, ifMatch, parameters, options)
	if err != nil {
		return TenantAccessClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TenantAccessClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusNoContent) {
		return TenantAccessClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return TenantAccessClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *TenantAccessClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, accessName AccessIDName, ifMatch string, parameters AccessInformationUpdateParameters, options *TenantAccessClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tenant/{accessName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if accessName == "" {
		return nil, errors.New("parameter accessName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accessName}", url.PathEscape(string(accessName)))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-09")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["If-Match"] = []string{ifMatch}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}
