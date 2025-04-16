// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmaintenance

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

// ConfigurationAssignmentsClient contains the methods for the ConfigurationAssignments group.
// Don't use this type directly, use NewConfigurationAssignmentsClient() instead.
type ConfigurationAssignmentsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConfigurationAssignmentsClient creates a new instance of ConfigurationAssignmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConfigurationAssignmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConfigurationAssignmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConfigurationAssignmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Register configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Configuration assignment name
//   - configurationAssignment - The configurationAssignment
//   - options - ConfigurationAssignmentsClientCreateOrUpdateOptions contains the optional parameters for the ConfigurationAssignmentsClient.CreateOrUpdate
//     method.
func (client *ConfigurationAssignmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsClientCreateOrUpdateOptions) (ConfigurationAssignmentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, configurationAssignment, options)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConfigurationAssignmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, _ *ConfigurationAssignmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configurationAssignment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConfigurationAssignmentsClient) createOrUpdateHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientCreateOrUpdateResponse, error) {
	result := ConfigurationAssignmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// CreateOrUpdateParent - Register configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Configuration assignment name
//   - configurationAssignment - The configurationAssignment
//   - options - ConfigurationAssignmentsClientCreateOrUpdateParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.CreateOrUpdateParent
//     method.
func (client *ConfigurationAssignmentsClient) CreateOrUpdateParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, options *ConfigurationAssignmentsClientCreateOrUpdateParentOptions) (ConfigurationAssignmentsClientCreateOrUpdateParentResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsClient.CreateOrUpdateParent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, configurationAssignment, options)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	resp, err := client.createOrUpdateParentHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateParentCreateRequest creates the CreateOrUpdateParent request.
func (client *ConfigurationAssignmentsClient) createOrUpdateParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, configurationAssignment ConfigurationAssignment, _ *ConfigurationAssignmentsClientCreateOrUpdateParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, configurationAssignment); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateParentHandleResponse handles the CreateOrUpdateParent response.
func (client *ConfigurationAssignmentsClient) createOrUpdateParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientCreateOrUpdateParentResponse, error) {
	result := ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientCreateOrUpdateParentResponse{}, err
	}
	return result, nil
}

// Delete - Unregister configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Unique configuration assignment name
//   - options - ConfigurationAssignmentsClientDeleteOptions contains the optional parameters for the ConfigurationAssignmentsClient.Delete
//     method.
func (client *ConfigurationAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientDeleteOptions) (ConfigurationAssignmentsClientDeleteResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	resp, err := client.deleteHandleResponse(httpResp)
	return resp, err
}

// deleteCreateRequest creates the Delete request.
func (client *ConfigurationAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, _ *ConfigurationAssignmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteHandleResponse handles the Delete response.
func (client *ConfigurationAssignmentsClient) deleteHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientDeleteResponse, error) {
	result := ConfigurationAssignmentsClientDeleteResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientDeleteResponse{}, err
	}
	return result, nil
}

// DeleteParent - Unregister configuration for resource.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Unique configuration assignment name
//   - options - ConfigurationAssignmentsClientDeleteParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.DeleteParent
//     method.
func (client *ConfigurationAssignmentsClient) DeleteParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientDeleteParentOptions) (ConfigurationAssignmentsClientDeleteParentResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsClient.DeleteParent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	resp, err := client.deleteParentHandleResponse(httpResp)
	return resp, err
}

// deleteParentCreateRequest creates the DeleteParent request.
func (client *ConfigurationAssignmentsClient) deleteParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, _ *ConfigurationAssignmentsClientDeleteParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// deleteParentHandleResponse handles the DeleteParent response.
func (client *ConfigurationAssignmentsClient) deleteParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientDeleteParentResponse, error) {
	result := ConfigurationAssignmentsClientDeleteParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientDeleteParentResponse{}, err
	}
	return result, nil
}

// Get - Get configuration assignment for resource..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Configuration assignment name
//   - options - ConfigurationAssignmentsClientGetOptions contains the optional parameters for the ConfigurationAssignmentsClient.Get
//     method.
func (client *ConfigurationAssignmentsClient) Get(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientGetOptions) (ConfigurationAssignmentsClientGetResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ConfigurationAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, configurationAssignmentName string, _ *ConfigurationAssignmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConfigurationAssignmentsClient) getHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientGetResponse, error) {
	result := ConfigurationAssignmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientGetResponse{}, err
	}
	return result, nil
}

// GetParent - Get configuration assignment for resource..
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - configurationAssignmentName - Configuration assignment name
//   - options - ConfigurationAssignmentsClientGetParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.GetParent
//     method.
func (client *ConfigurationAssignmentsClient) GetParent(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, options *ConfigurationAssignmentsClientGetParentOptions) (ConfigurationAssignmentsClientGetParentResponse, error) {
	var err error
	const operationName = "ConfigurationAssignmentsClient.GetParent"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, configurationAssignmentName, options)
	if err != nil {
		return ConfigurationAssignmentsClientGetParentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConfigurationAssignmentsClientGetParentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ConfigurationAssignmentsClientGetParentResponse{}, err
	}
	resp, err := client.getParentHandleResponse(httpResp)
	return resp, err
}

// getParentCreateRequest creates the GetParent request.
func (client *ConfigurationAssignmentsClient) getParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, configurationAssignmentName string, _ *ConfigurationAssignmentsClientGetParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments/{configurationAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if configurationAssignmentName == "" {
		return nil, errors.New("parameter configurationAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationAssignmentName}", url.PathEscape(configurationAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getParentHandleResponse handles the GetParent response.
func (client *ConfigurationAssignmentsClient) getParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientGetParentResponse, error) {
	result := ConfigurationAssignmentsClientGetParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConfigurationAssignment); err != nil {
		return ConfigurationAssignmentsClientGetParentResponse{}, err
	}
	return result, nil
}

// NewListPager - List configurationAssignments for resource.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - ConfigurationAssignmentsClientListOptions contains the optional parameters for the ConfigurationAssignmentsClient.NewListPager
//     method.
func (client *ConfigurationAssignmentsClient) NewListPager(resourceGroupName string, providerName string, resourceType string, resourceName string, options *ConfigurationAssignmentsClientListOptions) *runtime.Pager[ConfigurationAssignmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationAssignmentsClientListResponse]{
		More: func(page ConfigurationAssignmentsClientListResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationAssignmentsClientListResponse) (ConfigurationAssignmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConfigurationAssignmentsClient.NewListPager")
			req, err := client.listCreateRequest(ctx, resourceGroupName, providerName, resourceType, resourceName, options)
			if err != nil {
				return ConfigurationAssignmentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationAssignmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationAssignmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *ConfigurationAssignmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceType string, resourceName string, _ *ConfigurationAssignmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConfigurationAssignmentsClient) listHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientListResponse, error) {
	result := ConfigurationAssignmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListConfigurationAssignmentsResult); err != nil {
		return ConfigurationAssignmentsClientListResponse{}, err
	}
	return result, nil
}

// NewListParentPager - List configurationAssignments for resource.
//
// Generated from API version 2023-10-01-preview
//   - resourceGroupName - Resource group name
//   - providerName - Resource provider name
//   - resourceParentType - Resource parent type
//   - resourceParentName - Resource parent identifier
//   - resourceType - Resource type
//   - resourceName - Resource identifier
//   - options - ConfigurationAssignmentsClientListParentOptions contains the optional parameters for the ConfigurationAssignmentsClient.NewListParentPager
//     method.
func (client *ConfigurationAssignmentsClient) NewListParentPager(resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, options *ConfigurationAssignmentsClientListParentOptions) *runtime.Pager[ConfigurationAssignmentsClientListParentResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConfigurationAssignmentsClientListParentResponse]{
		More: func(page ConfigurationAssignmentsClientListParentResponse) bool {
			return false
		},
		Fetcher: func(ctx context.Context, page *ConfigurationAssignmentsClientListParentResponse) (ConfigurationAssignmentsClientListParentResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ConfigurationAssignmentsClient.NewListParentPager")
			req, err := client.listParentCreateRequest(ctx, resourceGroupName, providerName, resourceParentType, resourceParentName, resourceType, resourceName, options)
			if err != nil {
				return ConfigurationAssignmentsClientListParentResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConfigurationAssignmentsClientListParentResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConfigurationAssignmentsClientListParentResponse{}, runtime.NewResponseError(resp)
			}
			return client.listParentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listParentCreateRequest creates the ListParent request.
func (client *ConfigurationAssignmentsClient) listParentCreateRequest(ctx context.Context, resourceGroupName string, providerName string, resourceParentType string, resourceParentName string, resourceType string, resourceName string, _ *ConfigurationAssignmentsClientListParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{providerName}/{resourceParentType}/{resourceParentName}/{resourceType}/{resourceName}/providers/Microsoft.Maintenance/configurationAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if providerName == "" {
		return nil, errors.New("parameter providerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerName}", url.PathEscape(providerName))
	if resourceParentType == "" {
		return nil, errors.New("parameter resourceParentType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentType}", url.PathEscape(resourceParentType))
	if resourceParentName == "" {
		return nil, errors.New("parameter resourceParentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceParentName}", url.PathEscape(resourceParentName))
	if resourceType == "" {
		return nil, errors.New("parameter resourceType cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceType}", url.PathEscape(resourceType))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listParentHandleResponse handles the ListParent response.
func (client *ConfigurationAssignmentsClient) listParentHandleResponse(resp *http.Response) (ConfigurationAssignmentsClientListParentResponse, error) {
	result := ConfigurationAssignmentsClientListParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListConfigurationAssignmentsResult); err != nil {
		return ConfigurationAssignmentsClientListParentResponse{}, err
	}
	return result, nil
}
