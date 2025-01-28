//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoperationalinsights

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

// IntelligencePacksClient contains the methods for the IntelligencePacks group.
// Don't use this type directly, use NewIntelligencePacksClient() instead.
type IntelligencePacksClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewIntelligencePacksClient creates a new instance of IntelligencePacksClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewIntelligencePacksClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*IntelligencePacksClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &IntelligencePacksClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Disable - Disables an intelligence pack for a given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - intelligencePackName - The name of the intelligence pack to be disabled.
//   - options - IntelligencePacksClientDisableOptions contains the optional parameters for the IntelligencePacksClient.Disable
//     method.
func (client *IntelligencePacksClient) Disable(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string, options *IntelligencePacksClientDisableOptions) (IntelligencePacksClientDisableResponse, error) {
	var err error
	const operationName = "IntelligencePacksClient.Disable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.disableCreateRequest(ctx, resourceGroupName, workspaceName, intelligencePackName, options)
	if err != nil {
		return IntelligencePacksClientDisableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntelligencePacksClientDisableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntelligencePacksClientDisableResponse{}, err
	}
	return IntelligencePacksClientDisableResponse{}, nil
}

// disableCreateRequest creates the Disable request.
func (client *IntelligencePacksClient) disableCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string, options *IntelligencePacksClientDisableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/intelligencePacks/{intelligencePackName}/Disable"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if intelligencePackName == "" {
		return nil, errors.New("parameter intelligencePackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intelligencePackName}", url.PathEscape(intelligencePackName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Enable - Enables an intelligence pack for a given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - intelligencePackName - The name of the intelligence pack to be enabled.
//   - options - IntelligencePacksClientEnableOptions contains the optional parameters for the IntelligencePacksClient.Enable
//     method.
func (client *IntelligencePacksClient) Enable(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string, options *IntelligencePacksClientEnableOptions) (IntelligencePacksClientEnableResponse, error) {
	var err error
	const operationName = "IntelligencePacksClient.Enable"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.enableCreateRequest(ctx, resourceGroupName, workspaceName, intelligencePackName, options)
	if err != nil {
		return IntelligencePacksClientEnableResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntelligencePacksClientEnableResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntelligencePacksClientEnableResponse{}, err
	}
	return IntelligencePacksClientEnableResponse{}, nil
}

// enableCreateRequest creates the Enable request.
func (client *IntelligencePacksClient) enableCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, intelligencePackName string, options *IntelligencePacksClientEnableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/intelligencePacks/{intelligencePackName}/Enable"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if intelligencePackName == "" {
		return nil, errors.New("parameter intelligencePackName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{intelligencePackName}", url.PathEscape(intelligencePackName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// List - Lists all the intelligence packs possible and whether they are enabled or disabled for a given workspace.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - IntelligencePacksClientListOptions contains the optional parameters for the IntelligencePacksClient.List method.
func (client *IntelligencePacksClient) List(ctx context.Context, resourceGroupName string, workspaceName string, options *IntelligencePacksClientListOptions) (IntelligencePacksClientListResponse, error) {
	var err error
	const operationName = "IntelligencePacksClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return IntelligencePacksClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return IntelligencePacksClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return IntelligencePacksClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *IntelligencePacksClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *IntelligencePacksClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/intelligencePacks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntelligencePacksClient) listHandleResponse(resp *http.Response) (IntelligencePacksClientListResponse, error) {
	result := IntelligencePacksClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntelligencePackArray); err != nil {
		return IntelligencePacksClientListResponse{}, err
	}
	return result, nil
}
