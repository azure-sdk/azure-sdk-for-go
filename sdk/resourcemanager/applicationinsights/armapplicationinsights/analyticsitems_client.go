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
	"strconv"
	"strings"
)

// AnalyticsItemsClient contains the methods for the AnalyticsItems group.
// Don't use this type directly, use NewAnalyticsItemsClient() instead.
type AnalyticsItemsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAnalyticsItemsClient creates a new instance of AnalyticsItemsClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAnalyticsItemsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AnalyticsItemsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AnalyticsItemsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Deletes a specific Analytics Items defined within an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - scopePath - Enum indicating if this item definition is owned by a specific user or is shared between all users with access
//     to the Application Insights component.
//   - options - AnalyticsItemsClientDeleteOptions contains the optional parameters for the AnalyticsItemsClient.Delete method.
func (client *AnalyticsItemsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, options *AnalyticsItemsClientDeleteOptions) (AnalyticsItemsClientDeleteResponse, error) {
	var err error
	const operationName = "AnalyticsItemsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, scopePath, options)
	if err != nil {
		return AnalyticsItemsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AnalyticsItemsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AnalyticsItemsClientDeleteResponse{}, err
	}
	return AnalyticsItemsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AnalyticsItemsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, options *AnalyticsItemsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if scopePath == "" {
		return nil, errors.New("parameter scopePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopePath}", url.PathEscape(string(scopePath)))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	if options != nil && options.ID != nil {
		reqQP.Set("id", *options.ID)
	}
	if options != nil && options.Name != nil {
		reqQP.Set("name", *options.Name)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	return req, nil
}

// Get - Gets a specific Analytics Items defined within an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - scopePath - Enum indicating if this item definition is owned by a specific user or is shared between all users with access
//     to the Application Insights component.
//   - options - AnalyticsItemsClientGetOptions contains the optional parameters for the AnalyticsItemsClient.Get method.
func (client *AnalyticsItemsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, options *AnalyticsItemsClientGetOptions) (AnalyticsItemsClientGetResponse, error) {
	var err error
	const operationName = "AnalyticsItemsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, scopePath, options)
	if err != nil {
		return AnalyticsItemsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AnalyticsItemsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AnalyticsItemsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AnalyticsItemsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, options *AnalyticsItemsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if scopePath == "" {
		return nil, errors.New("parameter scopePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopePath}", url.PathEscape(string(scopePath)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	if options != nil && options.ID != nil {
		reqQP.Set("id", *options.ID)
	}
	if options != nil && options.Name != nil {
		reqQP.Set("name", *options.Name)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AnalyticsItemsClient) getHandleResponse(resp *http.Response) (AnalyticsItemsClientGetResponse, error) {
	result := AnalyticsItemsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAnalyticsItem); err != nil {
		return AnalyticsItemsClientGetResponse{}, err
	}
	return result, nil
}

// List - Gets a list of Analytics Items defined within an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - scopePath - Enum indicating if this item definition is owned by a specific user or is shared between all users with access
//     to the Application Insights component.
//   - options - AnalyticsItemsClientListOptions contains the optional parameters for the AnalyticsItemsClient.List method.
func (client *AnalyticsItemsClient) List(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, options *AnalyticsItemsClientListOptions) (AnalyticsItemsClientListResponse, error) {
	var err error
	const operationName = "AnalyticsItemsClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, resourceGroupName, resourceName, scopePath, options)
	if err != nil {
		return AnalyticsItemsClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AnalyticsItemsClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AnalyticsItemsClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *AnalyticsItemsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, options *AnalyticsItemsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if scopePath == "" {
		return nil, errors.New("parameter scopePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopePath}", url.PathEscape(string(scopePath)))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	if options != nil && options.IncludeContent != nil {
		reqQP.Set("includeContent", strconv.FormatBool(*options.IncludeContent))
	}
	if options != nil && options.Scope != nil {
		reqQP.Set("scope", string(*options.Scope))
	}
	if options != nil && options.Type != nil {
		reqQP.Set("type", string(*options.Type))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AnalyticsItemsClient) listHandleResponse(resp *http.Response) (AnalyticsItemsClientListResponse, error) {
	result := AnalyticsItemsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAnalyticsItemArray); err != nil {
		return AnalyticsItemsClientListResponse{}, err
	}
	return result, nil
}

// Put - Adds or Updates a specific Analytics Item within an Application Insights component.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2015-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - resourceName - The name of the Application Insights component resource.
//   - scopePath - Enum indicating if this item definition is owned by a specific user or is shared between all users with access
//     to the Application Insights component.
//   - itemProperties - Properties that need to be specified to create a new item and add it to an Application Insights component.
//   - options - AnalyticsItemsClientPutOptions contains the optional parameters for the AnalyticsItemsClient.Put method.
func (client *AnalyticsItemsClient) Put(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, itemProperties ComponentAnalyticsItem, options *AnalyticsItemsClientPutOptions) (AnalyticsItemsClientPutResponse, error) {
	var err error
	const operationName = "AnalyticsItemsClient.Put"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.putCreateRequest(ctx, resourceGroupName, resourceName, scopePath, itemProperties, options)
	if err != nil {
		return AnalyticsItemsClientPutResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AnalyticsItemsClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AnalyticsItemsClientPutResponse{}, err
	}
	resp, err := client.putHandleResponse(httpResp)
	return resp, err
}

// putCreateRequest creates the Put request.
func (client *AnalyticsItemsClient) putCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, scopePath ItemScopePath, itemProperties ComponentAnalyticsItem, options *AnalyticsItemsClientPutOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/microsoft.insights/components/{resourceName}/{scopePath}/item"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if scopePath == "" {
		return nil, errors.New("parameter scopePath cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{scopePath}", url.PathEscape(string(scopePath)))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-05-01")
	if options != nil && options.OverrideItem != nil {
		reqQP.Set("overrideItem", strconv.FormatBool(*options.OverrideItem))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, itemProperties); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *AnalyticsItemsClient) putHandleResponse(resp *http.Response) (AnalyticsItemsClientPutResponse, error) {
	result := AnalyticsItemsClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ComponentAnalyticsItem); err != nil {
		return AnalyticsItemsClientPutResponse{}, err
	}
	return result, nil
}
