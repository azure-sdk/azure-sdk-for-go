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
	"strconv"
	"strings"
)

// MetadataClient contains the methods for the Metadata group.
// Don't use this type directly, use NewMetadataClient() instead.
type MetadataClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMetadataClient creates a new instance of MetadataClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetadataClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetadataClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetadataClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Create - Create a Metadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - metadataName - The Metadata name.
//   - metadata - Metadata resource.
//   - options - MetadataClientCreateOptions contains the optional parameters for the MetadataClient.Create method.
func (client *MetadataClient) Create(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, metadata MetadataModel, options *MetadataClientCreateOptions) (MetadataClientCreateResponse, error) {
	var err error
	const operationName = "MetadataClient.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, resourceGroupName, workspaceName, metadataName, metadata, options)
	if err != nil {
		return MetadataClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MetadataClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *MetadataClient) createCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, metadata MetadataModel, _ *MetadataClientCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/metadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, metadata); err != nil {
		return nil, err
	}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *MetadataClient) createHandleResponse(resp *http.Response) (MetadataClientCreateResponse, error) {
	result := MetadataClientCreateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataModel); err != nil {
		return MetadataClientCreateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a Metadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - metadataName - The Metadata name.
//   - options - MetadataClientDeleteOptions contains the optional parameters for the MetadataClient.Delete method.
func (client *MetadataClient) Delete(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, options *MetadataClientDeleteOptions) (MetadataClientDeleteResponse, error) {
	var err error
	const operationName = "MetadataClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, metadataName, options)
	if err != nil {
		return MetadataClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MetadataClientDeleteResponse{}, err
	}
	return MetadataClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MetadataClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, _ *MetadataClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/metadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a Metadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - metadataName - The Metadata name.
//   - options - MetadataClientGetOptions contains the optional parameters for the MetadataClient.Get method.
func (client *MetadataClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, options *MetadataClientGetOptions) (MetadataClientGetResponse, error) {
	var err error
	const operationName = "MetadataClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, metadataName, options)
	if err != nil {
		return MetadataClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetadataClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MetadataClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, _ *MetadataClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/metadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MetadataClient) getHandleResponse(resp *http.Response) (MetadataClientGetResponse, error) {
	result := MetadataClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataModel); err != nil {
		return MetadataClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List of all metadata
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - options - MetadataClientListOptions contains the optional parameters for the MetadataClient.NewListPager method.
func (client *MetadataClient) NewListPager(resourceGroupName string, workspaceName string, options *MetadataClientListOptions) *runtime.Pager[MetadataClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetadataClientListResponse]{
		More: func(page MetadataClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MetadataClientListResponse) (MetadataClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MetadataClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, workspaceName, options)
			}, nil)
			if err != nil {
				return MetadataClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MetadataClient) listCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *MetadataClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/metadata"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MetadataClient) listHandleResponse(resp *http.Response) (MetadataClientListResponse, error) {
	result := MetadataClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataList); err != nil {
		return MetadataClientListResponse{}, err
	}
	return result, nil
}

// Update - Update an existing Metadata.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - workspaceName - The name of the workspace.
//   - metadataName - The Metadata name.
//   - metadataPatch - Partial metadata request.
//   - options - MetadataClientUpdateOptions contains the optional parameters for the MetadataClient.Update method.
func (client *MetadataClient) Update(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, metadataPatch MetadataPatch, options *MetadataClientUpdateOptions) (MetadataClientUpdateResponse, error) {
	var err error
	const operationName = "MetadataClient.Update"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, metadataName, metadataPatch, options)
	if err != nil {
		return MetadataClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetadataClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *MetadataClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, metadataName string, metadataPatch MetadataPatch, _ *MetadataClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/providers/Microsoft.SecurityInsights/metadata/{metadataName}"
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
	if metadataName == "" {
		return nil, errors.New("parameter metadataName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataName}", url.PathEscape(metadataName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, metadataPatch); err != nil {
		return nil, err
	}
	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *MetadataClient) updateHandleResponse(resp *http.Response) (MetadataClientUpdateResponse, error) {
	result := MetadataClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataModel); err != nil {
		return MetadataClientUpdateResponse{}, err
	}
	return result, nil
}
