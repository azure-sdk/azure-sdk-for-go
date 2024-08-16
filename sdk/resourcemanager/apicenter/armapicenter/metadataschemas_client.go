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

// MetadataSchemasClient contains the methods for the MetadataSchemas group.
// Don't use this type directly, use NewMetadataSchemasClient() instead.
type MetadataSchemasClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMetadataSchemasClient creates a new instance of MetadataSchemasClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMetadataSchemasClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MetadataSchemasClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MetadataSchemasClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates new or updates existing metadata schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - metadataSchemaName - The name of the metadata schema.
//   - resource - Resource create parameters.
//   - options - MetadataSchemasClientCreateOrUpdateOptions contains the optional parameters for the MetadataSchemasClient.CreateOrUpdate
//     method.
func (client *MetadataSchemasClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, resource MetadataSchema, options *MetadataSchemasClientCreateOrUpdateOptions) (MetadataSchemasClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "MetadataSchemasClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, metadataSchemaName, resource, options)
	if err != nil {
		return MetadataSchemasClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataSchemasClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return MetadataSchemasClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MetadataSchemasClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, resource MetadataSchema, options *MetadataSchemasClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/metadataSchemas/{metadataSchemaName}"
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
	if metadataSchemaName == "" {
		return nil, errors.New("parameter metadataSchemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataSchemaName}", url.PathEscape(metadataSchemaName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *MetadataSchemasClient) createOrUpdateHandleResponse(resp *http.Response) (MetadataSchemasClientCreateOrUpdateResponse, error) {
	result := MetadataSchemasClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataSchema); err != nil {
		return MetadataSchemasClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes specified metadata schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - metadataSchemaName - The name of the metadata schema.
//   - options - MetadataSchemasClientDeleteOptions contains the optional parameters for the MetadataSchemasClient.Delete method.
func (client *MetadataSchemasClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, options *MetadataSchemasClientDeleteOptions) (MetadataSchemasClientDeleteResponse, error) {
	var err error
	const operationName = "MetadataSchemasClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, metadataSchemaName, options)
	if err != nil {
		return MetadataSchemasClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataSchemasClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return MetadataSchemasClientDeleteResponse{}, err
	}
	return MetadataSchemasClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MetadataSchemasClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, options *MetadataSchemasClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/metadataSchemas/{metadataSchemaName}"
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
	if metadataSchemaName == "" {
		return nil, errors.New("parameter metadataSchemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataSchemaName}", url.PathEscape(metadataSchemaName))
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

// Get - Returns details of the metadata schema.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - metadataSchemaName - The name of the metadata schema.
//   - options - MetadataSchemasClientGetOptions contains the optional parameters for the MetadataSchemasClient.Get method.
func (client *MetadataSchemasClient) Get(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, options *MetadataSchemasClientGetOptions) (MetadataSchemasClientGetResponse, error) {
	var err error
	const operationName = "MetadataSchemasClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, metadataSchemaName, options)
	if err != nil {
		return MetadataSchemasClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataSchemasClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetadataSchemasClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MetadataSchemasClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, options *MetadataSchemasClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/metadataSchemas/{metadataSchemaName}"
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
	if metadataSchemaName == "" {
		return nil, errors.New("parameter metadataSchemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataSchemaName}", url.PathEscape(metadataSchemaName))
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
func (client *MetadataSchemasClient) getHandleResponse(resp *http.Response) (MetadataSchemasClientGetResponse, error) {
	result := MetadataSchemasClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataSchema); err != nil {
		return MetadataSchemasClientGetResponse{}, err
	}
	return result, nil
}

// Head - Checks if specified metadata schema exists.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - metadataSchemaName - The name of the metadata schema.
//   - options - MetadataSchemasClientHeadOptions contains the optional parameters for the MetadataSchemasClient.Head method.
func (client *MetadataSchemasClient) Head(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, options *MetadataSchemasClientHeadOptions) (MetadataSchemasClientHeadResponse, error) {
	var err error
	const operationName = "MetadataSchemasClient.Head"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.headCreateRequest(ctx, resourceGroupName, serviceName, metadataSchemaName, options)
	if err != nil {
		return MetadataSchemasClientHeadResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MetadataSchemasClientHeadResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MetadataSchemasClientHeadResponse{}, err
	}
	return MetadataSchemasClientHeadResponse{Success: httpResp.StatusCode >= 200 && httpResp.StatusCode < 300}, nil
}

// headCreateRequest creates the Head request.
func (client *MetadataSchemasClient) headCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, metadataSchemaName string, options *MetadataSchemasClientHeadOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/metadataSchemas/{metadataSchemaName}"
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
	if metadataSchemaName == "" {
		return nil, errors.New("parameter metadataSchemaName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{metadataSchemaName}", url.PathEscape(metadataSchemaName))
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

// NewListPager - Returns a collection of metadata schemas.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of Azure API Center service.
//   - options - MetadataSchemasClientListOptions contains the optional parameters for the MetadataSchemasClient.NewListPager
//     method.
func (client *MetadataSchemasClient) NewListPager(resourceGroupName string, serviceName string, options *MetadataSchemasClientListOptions) *runtime.Pager[MetadataSchemasClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MetadataSchemasClientListResponse]{
		More: func(page MetadataSchemasClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MetadataSchemasClientListResponse) (MetadataSchemasClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MetadataSchemasClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, serviceName, options)
			}, nil)
			if err != nil {
				return MetadataSchemasClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MetadataSchemasClient) listCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, options *MetadataSchemasClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/services/{serviceName}/metadataSchemas"
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
func (client *MetadataSchemasClient) listHandleResponse(resp *http.Response) (MetadataSchemasClientListResponse, error) {
	result := MetadataSchemasClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MetadataSchemaListResult); err != nil {
		return MetadataSchemasClientListResponse{}, err
	}
	return result, nil
}
