//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// DelegationSettingsClient contains the methods for the DelegationSettings group.
// Don't use this type directly, use NewDelegationSettingsClient() instead.
type DelegationSettingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDelegationSettingsClient creates a new instance of DelegationSettingsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDelegationSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DelegationSettingsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DelegationSettingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update Delegation settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - parameters - Create or update parameters.
//   - options - DelegationSettingsClientCreateOrUpdateOptions contains the optional parameters for the DelegationSettingsClient.CreateOrUpdate
//     method.
func (client *DelegationSettingsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalDelegationSettings, options *DelegationSettingsClientCreateOrUpdateOptions) (DelegationSettingsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "DelegationSettingsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, parameters, options)
	if err != nil {
		return DelegationSettingsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DelegationSettingsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DelegationSettingsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *DelegationSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, parameters PortalDelegationSettings, options *DelegationSettingsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
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
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *DelegationSettingsClient) createOrUpdateHandleResponse(resp *http.Response) (DelegationSettingsClientCreateOrUpdateResponse, error) {
	result := DelegationSettingsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalDelegationSettings); err != nil {
		return DelegationSettingsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Get - Get Delegation Settings for the Portal.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - DelegationSettingsClientGetOptions contains the optional parameters for the DelegationSettingsClient.Get method.
func (client *DelegationSettingsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsClientGetOptions) (DelegationSettingsClientGetResponse, error) {
	var err error
	const operationName = "DelegationSettingsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DelegationSettingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DelegationSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DelegationSettingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DelegationSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
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
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *DelegationSettingsClient) getHandleResponse(resp *http.Response) (DelegationSettingsClientGetResponse, error) {
	result := DelegationSettingsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalDelegationSettings); err != nil {
		return DelegationSettingsClientGetResponse{}, err
	}
	return result, nil
}

// GetEntityTag - Gets the entity state (Etag) version of the DelegationSettings.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - DelegationSettingsClientGetEntityTagOptions contains the optional parameters for the DelegationSettingsClient.GetEntityTag
//     method.
func (client *DelegationSettingsClient) GetEntityTag(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsClientGetEntityTagOptions) (DelegationSettingsClientGetEntityTagResponse, error) {
	var err error
	const operationName = "DelegationSettingsClient.GetEntityTag"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getEntityTagCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DelegationSettingsClientGetEntityTagResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DelegationSettingsClientGetEntityTagResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DelegationSettingsClientGetEntityTagResponse{}, err
	}
	resp, err := client.getEntityTagHandleResponse(httpResp)
	return resp, err
}

// getEntityTagCreateRequest creates the GetEntityTag request.
func (client *DelegationSettingsClient) getEntityTagCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsClientGetEntityTagOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
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
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getEntityTagHandleResponse handles the GetEntityTag response.
func (client *DelegationSettingsClient) getEntityTagHandleResponse(resp *http.Response) (DelegationSettingsClientGetEntityTagResponse, error) {
	result := DelegationSettingsClientGetEntityTagResponse{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	return result, nil
}

// ListSecrets - Gets the secret validation key of the DelegationSettings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - DelegationSettingsClientListSecretsOptions contains the optional parameters for the DelegationSettingsClient.ListSecrets
//     method.
func (client *DelegationSettingsClient) ListSecrets(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsClientListSecretsOptions) (DelegationSettingsClientListSecretsResponse, error) {
	var err error
	const operationName = "DelegationSettingsClient.ListSecrets"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listSecretsCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return DelegationSettingsClientListSecretsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DelegationSettingsClientListSecretsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DelegationSettingsClientListSecretsResponse{}, err
	}
	resp, err := client.listSecretsHandleResponse(httpResp)
	return resp, err
}

// listSecretsCreateRequest creates the ListSecrets request.
func (client *DelegationSettingsClient) listSecretsCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *DelegationSettingsClientListSecretsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation/listSecrets"
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
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listSecretsHandleResponse handles the ListSecrets response.
func (client *DelegationSettingsClient) listSecretsHandleResponse(resp *http.Response) (DelegationSettingsClientListSecretsResponse, error) {
	result := DelegationSettingsClientListSecretsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PortalSettingValidationKeyContract); err != nil {
		return DelegationSettingsClientListSecretsResponse{}, err
	}
	return result, nil
}

// Update - Update Delegation settings.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - parameters - Update Delegation settings.
//   - options - DelegationSettingsClientUpdateOptions contains the optional parameters for the DelegationSettingsClient.Update
//     method.
func (client *DelegationSettingsClient) Update(ctx context.Context, resourceGroupName string, serviceName string, ifMatch string, parameters PortalDelegationSettings, options *DelegationSettingsClientUpdateOptions) (DelegationSettingsClientUpdateResponse, error) {
	var err error
	const operationName = "DelegationSettingsClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, serviceName, ifMatch, parameters, options)
	if err != nil {
		return DelegationSettingsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DelegationSettingsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DelegationSettingsClientUpdateResponse{}, err
	}
	return DelegationSettingsClientUpdateResponse{}, nil
}

// updateCreateRequest creates the Update request.
func (client *DelegationSettingsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, ifMatch string, parameters PortalDelegationSettings, options *DelegationSettingsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/portalsettings/delegation"
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
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}
