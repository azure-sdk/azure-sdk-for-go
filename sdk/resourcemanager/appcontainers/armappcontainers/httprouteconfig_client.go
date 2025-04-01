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

// HTTPRouteConfigClient contains the methods for the HTTPRouteConfig group.
// Don't use this type directly, use NewHTTPRouteConfigClient() instead.
type HTTPRouteConfigClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewHTTPRouteConfigClient creates a new instance of HTTPRouteConfigClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewHTTPRouteConfigClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HTTPRouteConfigClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &HTTPRouteConfigClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Create or Update a Http Route Config.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - httpRouteName - Name of the Http Route Config Resource.
//   - options - HTTPRouteConfigClientCreateOrUpdateOptions contains the optional parameters for the HTTPRouteConfigClient.CreateOrUpdate
//     method.
func (client *HTTPRouteConfigClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *HTTPRouteConfigClientCreateOrUpdateOptions) (HTTPRouteConfigClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "HTTPRouteConfigClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, environmentName, httpRouteName, options)
	if err != nil {
		return HTTPRouteConfigClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRouteConfigClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRouteConfigClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *HTTPRouteConfigClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *HTTPRouteConfigClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/httpRouteConfigs/{httpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.HTTPRouteConfigEnvelope != nil {
		if err := runtime.MarshalAsJSON(req, *options.HTTPRouteConfigEnvelope); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *HTTPRouteConfigClient) createOrUpdateHandleResponse(resp *http.Response) (HTTPRouteConfigClientCreateOrUpdateResponse, error) {
	result := HTTPRouteConfigClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteConfig); err != nil {
		return HTTPRouteConfigClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified Managed Http Route.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - httpRouteName - Name of the Http Route Config Resource.
//   - options - HTTPRouteConfigClientDeleteOptions contains the optional parameters for the HTTPRouteConfigClient.Delete method.
func (client *HTTPRouteConfigClient) Delete(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *HTTPRouteConfigClientDeleteOptions) (HTTPRouteConfigClientDeleteResponse, error) {
	var err error
	const operationName = "HTTPRouteConfigClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, environmentName, httpRouteName, options)
	if err != nil {
		return HTTPRouteConfigClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRouteConfigClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRouteConfigClientDeleteResponse{}, err
	}
	return HTTPRouteConfigClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *HTTPRouteConfigClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, _ *HTTPRouteConfigClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/httpRouteConfigs/{httpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the specified Managed Http Route Config.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - httpRouteName - Name of the Http Route Config Resource.
//   - options - HTTPRouteConfigClientGetOptions contains the optional parameters for the HTTPRouteConfigClient.Get method.
func (client *HTTPRouteConfigClient) Get(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, options *HTTPRouteConfigClientGetOptions) (HTTPRouteConfigClientGetResponse, error) {
	var err error
	const operationName = "HTTPRouteConfigClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, environmentName, httpRouteName, options)
	if err != nil {
		return HTTPRouteConfigClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRouteConfigClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRouteConfigClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *HTTPRouteConfigClient) getCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, _ *HTTPRouteConfigClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/httpRouteConfigs/{httpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HTTPRouteConfigClient) getHandleResponse(resp *http.Response) (HTTPRouteConfigClientGetResponse, error) {
	result := HTTPRouteConfigClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteConfig); err != nil {
		return HTTPRouteConfigClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Get the Managed Http Routes in a given managed environment.
//
// Generated from API version 2025-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - options - HTTPRouteConfigClientListOptions contains the optional parameters for the HTTPRouteConfigClient.NewListPager
//     method.
func (client *HTTPRouteConfigClient) NewListPager(resourceGroupName string, environmentName string, options *HTTPRouteConfigClientListOptions) *runtime.Pager[HTTPRouteConfigClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[HTTPRouteConfigClientListResponse]{
		More: func(page HTTPRouteConfigClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *HTTPRouteConfigClientListResponse) (HTTPRouteConfigClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "HTTPRouteConfigClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, environmentName, options)
			}, nil)
			if err != nil {
				return HTTPRouteConfigClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *HTTPRouteConfigClient) listCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, _ *HTTPRouteConfigClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/httpRouteConfigs"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *HTTPRouteConfigClient) listHandleResponse(resp *http.Response) (HTTPRouteConfigClientListResponse, error) {
	result := HTTPRouteConfigClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteConfigCollection); err != nil {
		return HTTPRouteConfigClientListResponse{}, err
	}
	return result, nil
}

// Update - Patches an http route config resource. Only patching of tags is supported
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-02-02-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - environmentName - Name of the Managed Environment.
//   - httpRouteName - Name of the Http Route Config Resource.
//   - httpRouteConfigEnvelope - Properties of http route config that need to be updated
//   - options - HTTPRouteConfigClientUpdateOptions contains the optional parameters for the HTTPRouteConfigClient.Update method.
func (client *HTTPRouteConfigClient) Update(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, httpRouteConfigEnvelope HTTPRouteConfig, options *HTTPRouteConfigClientUpdateOptions) (HTTPRouteConfigClientUpdateResponse, error) {
	var err error
	const operationName = "HTTPRouteConfigClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, environmentName, httpRouteName, httpRouteConfigEnvelope, options)
	if err != nil {
		return HTTPRouteConfigClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPRouteConfigClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return HTTPRouteConfigClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *HTTPRouteConfigClient) updateCreateRequest(ctx context.Context, resourceGroupName string, environmentName string, httpRouteName string, httpRouteConfigEnvelope HTTPRouteConfig, _ *HTTPRouteConfigClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.App/managedEnvironments/{environmentName}/httpRouteConfigs/{httpRouteName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	if httpRouteName == "" {
		return nil, errors.New("parameter httpRouteName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{httpRouteName}", url.PathEscape(httpRouteName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-02-02-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, httpRouteConfigEnvelope); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *HTTPRouteConfigClient) updateHandleResponse(resp *http.Response) (HTTPRouteConfigClientUpdateResponse, error) {
	result := HTTPRouteConfigClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HTTPRouteConfig); err != nil {
		return HTTPRouteConfigClientUpdateResponse{}, err
	}
	return result, nil
}
