// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

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
	internal       *arm.Client
	subscriptionID string
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnvironmentsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnvironmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnvironmentsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - environmentName - The name of the environment.
//   - payload - Resource create parameters.
//   - options - EnvironmentsClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentsClient.CreateOrUpdate
//     method.
func (client *EnvironmentsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, payload Environment, options *EnvironmentsClientCreateOrUpdateOptions) (EnvironmentsClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "EnvironmentsClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceName, environmentName, payload, options)
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
func (client *EnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, payload Environment, _ *EnvironmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
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
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["Content-Type"] = []string{"application/json"}
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
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - environmentName - The name of the environment.
//   - options - EnvironmentsClientDeleteOptions contains the optional parameters for the EnvironmentsClient.Delete method.
func (client *EnvironmentsClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, options *EnvironmentsClientDeleteOptions) (EnvironmentsClientDeleteResponse, error) {
	var err error
	const operationName = "EnvironmentsClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceName, environmentName, options)
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
func (client *EnvironmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, _ *EnvironmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
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
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Returns details of the environment.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - environmentName - The name of the environment.
//   - options - EnvironmentsClientGetOptions contains the optional parameters for the EnvironmentsClient.Get method.
func (client *EnvironmentsClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, options *EnvironmentsClientGetOptions) (EnvironmentsClientGetResponse, error) {
	var err error
	const operationName = "EnvironmentsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceName, environmentName, options)
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
func (client *EnvironmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, _ *EnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
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
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-15-preview")
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
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - environmentName - The name of the environment.
//   - options - EnvironmentsClientHeadOptions contains the optional parameters for the EnvironmentsClient.Head method.
func (client *EnvironmentsClient) Head(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, options *EnvironmentsClientHeadOptions) (EnvironmentsClientHeadResponse, error) {
	var err error
	const operationName = "EnvironmentsClient.Head"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.headCreateRequest(ctx, resourceGroupName, serviceName, workspaceName, environmentName, options)
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
func (client *EnvironmentsClient) headCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, environmentName string, _ *EnvironmentsClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments/{environmentName}"
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
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListPager - Returns a collection of environments.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - workspaceName - The name of the workspace.
//   - options - EnvironmentsClientListOptions contains the optional parameters for the EnvironmentsClient.NewListPager method.
func (client *EnvironmentsClient) NewListPager(resourceGroupName string, serviceName string, workspaceName string, options *EnvironmentsClientListOptions) *runtime.Pager[EnvironmentsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[EnvironmentsClientListResponse]{
		More: func(page EnvironmentsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnvironmentsClientListResponse) (EnvironmentsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "EnvironmentsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, serviceName, workspaceName, options)
			}, nil)
			if err != nil {
				return EnvironmentsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *EnvironmentsClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceName string, options *EnvironmentsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/workspaces/{workspaceName}/environments"
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
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *EnvironmentsClient) listHandleResponse(resp *http.Response) (EnvironmentsClientListResponse, error) {
	result := EnvironmentsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentListResult); err != nil {
		return EnvironmentsClientListResponse{}, err
	}
	return result, nil
}
