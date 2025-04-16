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
	"strings"
)

// ContentTypeClient contains the methods for the ContentType group.
// Don't use this type directly, use NewContentTypeClient() instead.
type ContentTypeClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewContentTypeClient creates a new instance of ContentTypeClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewContentTypeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ContentTypeClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ContentTypeClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates the developer portal's content type. Content types describe content items' properties,
// validation rules, and constraints. Custom content types' identifiers need to start with the c-
// prefix. Built-in content types can't be modified.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - contentTypeID - Content type identifier.
//   - parameters - Create or update parameters.
//   - options - ContentTypeClientCreateOrUpdateOptions contains the optional parameters for the ContentTypeClient.CreateOrUpdate
//     method.
func (client *ContentTypeClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, parameters ContentTypeContract, options *ContentTypeClientCreateOrUpdateOptions) (ContentTypeClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "ContentTypeClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, parameters, options)
	if err != nil {
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ContentTypeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, parameters ContentTypeContract, options *ContentTypeClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.IfMatch != nil {
		req.Raw().Header["If-Match"] = []string{*options.IfMatch}
	}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ContentTypeClient) createOrUpdateHandleResponse(resp *http.Response) (ContentTypeClientCreateOrUpdateResponse, error) {
	result := ContentTypeClientCreateOrUpdateResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentTypeContract); err != nil {
		return ContentTypeClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Removes the specified developer portal's content type. Content types describe content items' properties, validation
// rules, and constraints. Built-in content types (with identifiers starting with the
// c- prefix) can't be removed.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - contentTypeID - Content type identifier.
//   - ifMatch - ETag of the Entity. ETag should match the current entity state from the header response of the GET request or
//     it should be * for unconditional update.
//   - options - ContentTypeClientDeleteOptions contains the optional parameters for the ContentTypeClient.Delete method.
func (client *ContentTypeClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, ifMatch string, options *ContentTypeClientDeleteOptions) (ContentTypeClientDeleteResponse, error) {
	var err error
	const operationName = "ContentTypeClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, ifMatch, options)
	if err != nil {
		return ContentTypeClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentTypeClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return ContentTypeClientDeleteResponse{}, err
	}
	return ContentTypeClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ContentTypeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, ifMatch string, _ *ContentTypeClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	req.Raw().Header["If-Match"] = []string{ifMatch}
	return req, nil
}

// Get - Gets the details of the developer portal's content type. Content types describe content items' properties, validation
// rules, and constraints.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - contentTypeID - Content type identifier.
//   - options - ContentTypeClientGetOptions contains the optional parameters for the ContentTypeClient.Get method.
func (client *ContentTypeClient) Get(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, options *ContentTypeClientGetOptions) (ContentTypeClientGetResponse, error) {
	var err error
	const operationName = "ContentTypeClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, contentTypeID, options)
	if err != nil {
		return ContentTypeClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ContentTypeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ContentTypeClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *ContentTypeClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, contentTypeID string, _ *ContentTypeClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes/{contentTypeId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if contentTypeID == "" {
		return nil, errors.New("parameter contentTypeID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{contentTypeId}", url.PathEscape(contentTypeID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ContentTypeClient) getHandleResponse(resp *http.Response) (ContentTypeClientGetResponse, error) {
	result := ContentTypeClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentTypeContract); err != nil {
		return ContentTypeClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServicePager - Lists the developer portal's content types. Content types describe content items' properties, validation
// rules, and constraints.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - options - ContentTypeClientListByServiceOptions contains the optional parameters for the ContentTypeClient.NewListByServicePager
//     method.
func (client *ContentTypeClient) NewListByServicePager(resourceGroupName string, serviceName string, options *ContentTypeClientListByServiceOptions) *runtime.Pager[ContentTypeClientListByServiceResponse] {
	return runtime.NewPager(runtime.PagingHandler[ContentTypeClientListByServiceResponse]{
		More: func(page ContentTypeClientListByServiceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ContentTypeClientListByServiceResponse) (ContentTypeClientListByServiceResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "ContentTypeClient.NewListByServicePager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByServiceCreateRequest(ctx, resourceGroupName, serviceName, options)
			}, nil)
			if err != nil {
				return ContentTypeClientListByServiceResponse{}, err
			}
			return client.listByServiceHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByServiceCreateRequest creates the ListByService request.
func (client *ContentTypeClient) listByServiceCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, _ *ContentTypeClientListByServiceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/contentTypes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByServiceHandleResponse handles the ListByService response.
func (client *ContentTypeClient) listByServiceHandleResponse(resp *http.Response) (ContentTypeClientListByServiceResponse, error) {
	result := ContentTypeClientListByServiceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ContentTypeCollection); err != nil {
		return ContentTypeClientListByServiceResponse{}, err
	}
	return result, nil
}
