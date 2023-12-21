//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurityinsights

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

// SentinelOnboardingStatesClient contains the methods for the SentinelOnboardingStates group.
// Don't use this type directly, use NewSentinelOnboardingStatesClient() instead.
type SentinelOnboardingStatesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSentinelOnboardingStatesClient creates a new instance of SentinelOnboardingStatesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSentinelOnboardingStatesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SentinelOnboardingStatesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SentinelOnboardingStatesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create Sentinel onboarding state
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sentinelOnboardingStateName - The Sentinel onboarding state name. Supports - default
//   - options - SentinelOnboardingStatesClientCreateOptions contains the optional parameters for the SentinelOnboardingStatesClient.Create
//     method.
func (client *SentinelOnboardingStatesClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, options *SentinelOnboardingStatesClientCreateOptions) (SentinelOnboardingStatesClientCreateResponse, error) {
	var err error
	const operationName = "SentinelOnboardingStatesClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, sentinelOnboardingStateName, options)
	if err != nil {
		return SentinelOnboardingStatesClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SentinelOnboardingStatesClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return SentinelOnboardingStatesClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *SentinelOnboardingStatesClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, options *SentinelOnboardingStatesClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/onboardingStates/{sentinelOnboardingStateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sentinelOnboardingStateName == "" {
		return nil, errors.New("parameter sentinelOnboardingStateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sentinelOnboardingStateName}", url.PathEscape(sentinelOnboardingStateName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.SentinelOnboardingStateParameter != nil {
		if err := runtime.MarshalAsJSON(req, *options.SentinelOnboardingStateParameter); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *SentinelOnboardingStatesClient) createHandleResponse(resp *http.Response) (SentinelOnboardingStatesClientCreateResponse, error) {
	result := SentinelOnboardingStatesClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SentinelOnboardingState); err != nil {
		return SentinelOnboardingStatesClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete Sentinel onboarding state
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sentinelOnboardingStateName - The Sentinel onboarding state name. Supports - default
//   - options - SentinelOnboardingStatesClientDeleteOptions contains the optional parameters for the SentinelOnboardingStatesClient.Delete
//     method.
func (client *SentinelOnboardingStatesClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, options *SentinelOnboardingStatesClientDeleteOptions) (SentinelOnboardingStatesClientDeleteResponse, error) {
	var err error
	const operationName = "SentinelOnboardingStatesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, sentinelOnboardingStateName, options)
	if err != nil {
		return SentinelOnboardingStatesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SentinelOnboardingStatesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return SentinelOnboardingStatesClientDeleteResponse{}, err
	}
	return SentinelOnboardingStatesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SentinelOnboardingStatesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, options *SentinelOnboardingStatesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/onboardingStates/{sentinelOnboardingStateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sentinelOnboardingStateName == "" {
		return nil, errors.New("parameter sentinelOnboardingStateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sentinelOnboardingStateName}", url.PathEscape(sentinelOnboardingStateName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get Sentinel onboarding state
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - sentinelOnboardingStateName - The Sentinel onboarding state name. Supports - default
//   - options - SentinelOnboardingStatesClientGetOptions contains the optional parameters for the SentinelOnboardingStatesClient.Get
//     method.
func (client *SentinelOnboardingStatesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, options *SentinelOnboardingStatesClientGetOptions) (SentinelOnboardingStatesClientGetResponse, error) {
	var err error
	const operationName = "SentinelOnboardingStatesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, sentinelOnboardingStateName, options)
	if err != nil {
		return SentinelOnboardingStatesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SentinelOnboardingStatesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SentinelOnboardingStatesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SentinelOnboardingStatesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, sentinelOnboardingStateName string, options *SentinelOnboardingStatesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/onboardingStates/{sentinelOnboardingStateName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if sentinelOnboardingStateName == "" {
		return nil, errors.New("parameter sentinelOnboardingStateName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sentinelOnboardingStateName}", url.PathEscape(sentinelOnboardingStateName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SentinelOnboardingStatesClient) getHandleResponse(resp *http.Response) (SentinelOnboardingStatesClientGetResponse, error) {
	result := SentinelOnboardingStatesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SentinelOnboardingState); err != nil {
		return SentinelOnboardingStatesClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets all Sentinel onboarding states
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-12-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - SentinelOnboardingStatesClientListOptions contains the optional parameters for the SentinelOnboardingStatesClient.List
//     method.
func (client *SentinelOnboardingStatesClient) List(ctx context.Context, resourceGroupName string, workspaceName string, options *SentinelOnboardingStatesClientListOptions) (SentinelOnboardingStatesClientListResponse, error) {
	var err error
	const operationName = "SentinelOnboardingStatesClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return SentinelOnboardingStatesClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SentinelOnboardingStatesClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SentinelOnboardingStatesClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *SentinelOnboardingStatesClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *SentinelOnboardingStatesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/onboardingStates"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-12-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SentinelOnboardingStatesClient) listHandleResponse(resp *http.Response) (SentinelOnboardingStatesClientListResponse, error) {
	result := SentinelOnboardingStatesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SentinelOnboardingStatesList); err != nil {
		return SentinelOnboardingStatesClientListResponse{}, err
	}
	return result, nil
}
