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
	"strconv"
	"strings"
)

// WorkspaceTagAPILinkClient contains the methods for the WorkspaceTagAPILink group.
// Don't use this type directly, use NewWorkspaceTagAPILinkClient() instead.
type WorkspaceTagAPILinkClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceTagAPILinkClient creates a new instance of WorkspaceTagAPILinkClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceTagAPILinkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceTagAPILinkClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceTagAPILinkClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Adds an API to the specified tag via link.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Tag-API link identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - WorkspaceTagAPILinkClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceTagAPILinkClient.CreateOrUpdate
//     method.
func (client *WorkspaceTagAPILinkClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, apiLinkID string, parameters TagAPILinkContract, options *WorkspaceTagAPILinkClientCreateOrUpdateOptions) (WorkspaceTagAPILinkClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "WorkspaceTagAPILinkClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, tagID, apiLinkID, parameters, options)
	if err != nil {
		return WorkspaceTagAPILinkClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceTagAPILinkClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceTagAPILinkClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceTagAPILinkClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, apiLinkID string, parameters TagAPILinkContract, options *WorkspaceTagAPILinkClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/tags/{tagId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if tagID == "" {
		return nil, errors.New("parameter tagID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagId}", url.PathEscape(tagID))
	if apiLinkID == "" {
		return nil, errors.New("parameter apiLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiLinkId}", url.PathEscape(apiLinkID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WorkspaceTagAPILinkClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceTagAPILinkClientCreateOrUpdateResponse, error) {
	result := WorkspaceTagAPILinkClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagAPILinkContract); err != nil {
		return WorkspaceTagAPILinkClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified API from the specified tag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Tag-API link identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceTagAPILinkClientDeleteOptions contains the optional parameters for the WorkspaceTagAPILinkClient.Delete
//     method.
func (client *WorkspaceTagAPILinkClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, apiLinkID string, options *WorkspaceTagAPILinkClientDeleteOptions) (WorkspaceTagAPILinkClientDeleteResponse, error) {
	var err error
	const operationName = "WorkspaceTagAPILinkClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, tagID, apiLinkID, options)
	if err != nil {
		return WorkspaceTagAPILinkClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceTagAPILinkClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceTagAPILinkClientDeleteResponse{}, err
	}
	return WorkspaceTagAPILinkClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceTagAPILinkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, apiLinkID string, options *WorkspaceTagAPILinkClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/tags/{tagId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if tagID == "" {
		return nil, errors.New("parameter tagID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagId}", url.PathEscape(tagID))
	if apiLinkID == "" {
		return nil, errors.New("parameter apiLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiLinkId}", url.PathEscape(apiLinkID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the API link for the tag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Tag-API link identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceTagAPILinkClientGetOptions contains the optional parameters for the WorkspaceTagAPILinkClient.Get method.
func (client *WorkspaceTagAPILinkClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, apiLinkID string, options *WorkspaceTagAPILinkClientGetOptions) (WorkspaceTagAPILinkClientGetResponse, error) {
	var err error
	const operationName = "WorkspaceTagAPILinkClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, tagID, apiLinkID, options)
	if err != nil {
		return WorkspaceTagAPILinkClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceTagAPILinkClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return WorkspaceTagAPILinkClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *WorkspaceTagAPILinkClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, apiLinkID string, options *WorkspaceTagAPILinkClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/tags/{tagId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if tagID == "" {
		return nil, errors.New("parameter tagID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagId}", url.PathEscape(tagID))
	if apiLinkID == "" {
		return nil, errors.New("parameter apiLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{apiLinkId}", url.PathEscape(apiLinkID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspaceTagAPILinkClient) getHandleResponse(resp *http.Response) (WorkspaceTagAPILinkClientGetResponse, error) {
	result := WorkspaceTagAPILinkClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagAPILinkContract); err != nil {
		return WorkspaceTagAPILinkClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - Lists a collection of the API links associated with a tag.
//
// Generated from API version 2023-05-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceTagAPILinkClientListByProductOptions contains the optional parameters for the WorkspaceTagAPILinkClient.NewListByProductPager
//     method.
func (client *WorkspaceTagAPILinkClient) NewListByProductPager(resourceGroupName string, serviceName string, workspaceID string, tagID string, options *WorkspaceTagAPILinkClientListByProductOptions) *runtime.Pager[WorkspaceTagAPILinkClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceTagAPILinkClientListByProductResponse]{
		More: func(page WorkspaceTagAPILinkClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceTagAPILinkClientListByProductResponse) (WorkspaceTagAPILinkClientListByProductResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "WorkspaceTagAPILinkClient.NewListByProductPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, tagID, options)
			}, nil)
			if err != nil {
				return WorkspaceTagAPILinkClientListByProductResponse{}, err
			}
			return client.listByProductHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *WorkspaceTagAPILinkClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, tagID string, options *WorkspaceTagAPILinkClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/tags/{tagId}/apiLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if workspaceID == "" {
		return nil, errors.New("parameter workspaceID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceId}", url.PathEscape(workspaceID))
	if tagID == "" {
		return nil, errors.New("parameter tagID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{tagId}", url.PathEscape(tagID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	reqQP.Set("api-version", "2023-05-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *WorkspaceTagAPILinkClient) listByProductHandleResponse(resp *http.Response) (WorkspaceTagAPILinkClientListByProductResponse, error) {
	result := WorkspaceTagAPILinkClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagAPILinkCollection); err != nil {
		return WorkspaceTagAPILinkClientListByProductResponse{}, err
	}
	return result, nil
}
