//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// AppResiliencyClient contains the methods for the AppResiliency group.
// Don't use this type directly, use NewAppResiliencyClient() instead.
type AppResiliencyClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAppResiliencyClient creates a new instance of AppResiliencyClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAppResiliencyClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AppResiliencyClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AppResiliencyClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update container app resiliency policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appName - Name of the Container App.
//   - name - Name of the resiliency policy.
//   - resiliencyEnvelope - The resiliency policy to create or update.
//   - options - AppResiliencyClientCreateOrUpdateOptions contains the optional parameters for the AppResiliencyClient.CreateOrUpdate
//     method.
func (client *AppResiliencyClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, appName string, name string, resiliencyEnvelope AppResiliency, options *AppResiliencyClientCreateOrUpdateOptions) (AppResiliencyClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "AppResiliencyClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, appName, name, resiliencyEnvelope, options)
	if err != nil {
		return AppResiliencyClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppResiliencyClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return AppResiliencyClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AppResiliencyClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, appName string, name string, resiliencyEnvelope AppResiliency, options *AppResiliencyClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{appName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resiliencyEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AppResiliencyClient) createOrUpdateHandleResponse(resp *http.Response) (AppResiliencyClientCreateOrUpdateResponse, error) {
	result := AppResiliencyClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppResiliency); err != nil {
		return AppResiliencyClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete container app resiliency policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appName - Name of the Container App.
//   - name - Name of the resiliency policy.
//   - options - AppResiliencyClientDeleteOptions contains the optional parameters for the AppResiliencyClient.Delete method.
func (client *AppResiliencyClient) Delete(ctx context.Context, resourceGroupName string, appName string, name string, options *AppResiliencyClientDeleteOptions) (AppResiliencyClientDeleteResponse, error) {
	var err error
	const operationName = "AppResiliencyClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, appName, name, options)
	if err != nil {
		return AppResiliencyClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppResiliencyClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return AppResiliencyClientDeleteResponse{}, err
	}
	return AppResiliencyClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AppResiliencyClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, appName string, name string, options *AppResiliencyClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{appName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get container app resiliency policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appName - Name of the Container App.
//   - name - Name of the resiliency policy.
//   - options - AppResiliencyClientGetOptions contains the optional parameters for the AppResiliencyClient.Get method.
func (client *AppResiliencyClient) Get(ctx context.Context, resourceGroupName string, appName string, name string, options *AppResiliencyClientGetOptions) (AppResiliencyClientGetResponse, error) {
	var err error
	const operationName = "AppResiliencyClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, appName, name, options)
	if err != nil {
		return AppResiliencyClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppResiliencyClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppResiliencyClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AppResiliencyClient) getCreateRequest(ctx context.Context, resourceGroupName string, appName string, name string, options *AppResiliencyClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{appName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AppResiliencyClient) getHandleResponse(resp *http.Response) (AppResiliencyClientGetResponse, error) {
	result := AppResiliencyClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppResiliency); err != nil {
		return AppResiliencyClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List container app resiliency policies.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appName - Name of the Container App.
//   - options - AppResiliencyClientListOptions contains the optional parameters for the AppResiliencyClient.NewListPager method.
func (client *AppResiliencyClient) NewListPager(resourceGroupName string, appName string, options *AppResiliencyClientListOptions) *runtime.Pager[AppResiliencyClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AppResiliencyClientListResponse]{
		More: func(page AppResiliencyClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AppResiliencyClientListResponse) (AppResiliencyClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AppResiliencyClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, appName, options)
			}, nil)
			if err != nil {
				return AppResiliencyClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *AppResiliencyClient) listCreateRequest(ctx context.Context, resourceGroupName string, appName string, options *AppResiliencyClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{appName}/resiliencyPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AppResiliencyClient) listHandleResponse(resp *http.Response) (AppResiliencyClientListResponse, error) {
	result := AppResiliencyClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppResiliencyCollection); err != nil {
		return AppResiliencyClientListResponse{}, err
	}
	return result, nil
}

// Update - Update container app resiliency policy.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - appName - Name of the Container App.
//   - name - Name of the resiliency policy.
//   - resiliencyEnvelope - The resiliency policy to update.
//   - options - AppResiliencyClientUpdateOptions contains the optional parameters for the AppResiliencyClient.Update method.
func (client *AppResiliencyClient) Update(ctx context.Context, resourceGroupName string, appName string, name string, resiliencyEnvelope AppResiliency, options *AppResiliencyClientUpdateOptions) (AppResiliencyClientUpdateResponse, error) {
	var err error
	const operationName = "AppResiliencyClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, appName, name, resiliencyEnvelope, options)
	if err != nil {
		return AppResiliencyClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AppResiliencyClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AppResiliencyClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *AppResiliencyClient) updateCreateRequest(ctx context.Context, resourceGroupName string, appName string, name string, resiliencyEnvelope AppResiliency, options *AppResiliencyClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/containerApps/{appName}/resiliencyPolicies/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if appName == "" {
		return nil, errors.New("parameter appName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{appName}", url.PathEscape(appName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resiliencyEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *AppResiliencyClient) updateHandleResponse(resp *http.Response) (AppResiliencyClientUpdateResponse, error) {
	result := AppResiliencyClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AppResiliency); err != nil {
		return AppResiliencyClientUpdateResponse{}, err
	}
	return result, nil
}
