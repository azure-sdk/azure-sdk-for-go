//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapicenter

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

// EnvironmentsClient contains the methods for the Environments group.
// Don't use this type directly, use NewEnvironmentsClient() instead.
type EnvironmentsClient struct {
	internal        *arm.Client
	subscriptionID  string
	workspaceName   string
	filter          *string
	environmentName string
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - workspaceName - The name of the workspace.
//   - environmentName - The name of the environment.
//   - filter - OData filter parameter.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnvironmentsClient(subscriptionID string, workspaceName string, environmentName string, filter *string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnvironmentsClient, error) {
	cl, err := arm.NewClient(moduleName+".EnvironmentsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnvironmentsClient{
		subscriptionID:  subscriptionID,
		workspaceName:   workspaceName,
		environmentName: environmentName,
		filter:          filter,
		internal:        cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - payload - Environment entity.
//   - options - EnvironmentsClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentsClient.CreateOrUpdate
//     method.
func (client *EnvironmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, payload Environment, options *EnvironmentsClientCreateOrUpdateOptions) (EnvironmentsClientCreateOrUpdateResponse, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, payload, options)
	if err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, payload Environment, options *EnvironmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.environmentName == "" {
		return nil, errors.New("parameter client.environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(client.environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, payload); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EnvironmentsClient) createOrUpdateHandleResponse(resp *http.Response) (EnvironmentsClientCreateOrUpdateResponse, error) {
	result := EnvironmentsClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Environment); err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - EnvironmentsClientDeleteOptions contains the optional parameters for the EnvironmentsClient.Delete method.
func (client *EnvironmentsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientDeleteOptions) (EnvironmentsClientDeleteResponse, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return EnvironmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientDeleteResponse{}, err
	}
	return EnvironmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.environmentName == "" {
		return nil, errors.New("parameter client.environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(client.environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns details of the environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - EnvironmentsClientGetOptions contains the optional parameters for the EnvironmentsClient.Get method.
func (client *EnvironmentsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientGetOptions) (EnvironmentsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.environmentName == "" {
		return nil, errors.New("parameter client.environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(client.environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *EnvironmentsClient) getHandleResponse(resp *http.Response) (EnvironmentsClientGetResponse, error) {
	result := EnvironmentsClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.Environment); err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks if specified environment exists.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - EnvironmentsClientHeadOptions contains the optional parameters for the EnvironmentsClient.Head method.
func (client *EnvironmentsClient) Head(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientHeadOptions) (EnvironmentsClientHeadResponse, error) {
	var err error
	req, err := client.headCreateRequest(ctx, resourceGroupName, serviceName, options)
	if err != nil {
		return EnvironmentsClientHeadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientHeadResponse{}, err
	}
	return EnvironmentsClientHeadResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// headCreateRequest creates the Head request.
func (client *EnvironmentsClient) headCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	if client.environmentName == "" {
		return nil, errors.New("parameter client.environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(client.environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListPager - Returns a collection of environments.
//
// Generated from API version 2024-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - Service name
//   - options - EnvironmentsClientListOptions contains the optional parameters for the EnvironmentsClient.NewListPager method.
func (client *EnvironmentsClient) NewListPager(resourceGroupName string, serviceName string, options *EnvironmentsClientListOptions) *runtime.Pager[EnvironmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnvironmentsClientListResponse]{
		More: func(page EnvironmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnvironmentsClientListResponse) (EnvironmentsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return EnvironmentsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return EnvironmentsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return EnvironmentsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *EnvironmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *EnvironmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.workspaceName == "" {
		return nil, errors.New("parameter client.workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(client.workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-01")
	if client.filter != nil {
		reqQP.Set("$filter", *client.filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EnvironmentsClient) listHandleResponse(resp *http.Response) (EnvironmentsClientListResponse, error) {
	result := EnvironmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentCollection); err != nil {
		return EnvironmentsClientListResponse{}, err
	}
	return result, nil
}
