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

// TagAPILinkClient contains the methods for the TagAPILink group.
// Don't use this type directly, use NewTagAPILinkClient() instead.
type TagAPILinkClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTagAPILinkClient creates a new instance of TagAPILinkClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTagAPILinkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TagAPILinkClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TagAPILinkClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Adds an API to the specified tag via link.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Tag-API link identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - TagAPILinkClientCreateOrUpdateOptions contains the optional parameters for the TagAPILinkClient.CreateOrUpdate
//     method.
func (client *TagAPILinkClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, tagID string, apiLinkID string, parameters TagAPILinkContract, options *TagAPILinkClientCreateOrUpdateOptions) (TagAPILinkClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TagAPILinkClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, tagID, apiLinkID, parameters, options)
	if err != nil {
		return TagAPILinkClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagAPILinkClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TagAPILinkClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagAPILinkClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, apiLinkID string, parameters TagAPILinkContract, options *TagAPILinkClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagAPILinkClient) createOrUpdateHandleResponse(resp *http.Response) (TagAPILinkClientCreateOrUpdateResponse, error) {
	result := TagAPILinkClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagAPILinkContract); err != nil {
		return TagAPILinkClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified API from the specified tag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Tag-API link identifier. Must be unique in the current API Management service instance.
//   - options - TagAPILinkClientDeleteOptions contains the optional parameters for the TagAPILinkClient.Delete method.
func (client *TagAPILinkClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, tagID string, apiLinkID string, options *TagAPILinkClientDeleteOptions) (TagAPILinkClientDeleteResponse, error) {
	var err error
	const operationName = "TagAPILinkClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, tagID, apiLinkID, options)
	if err != nil {
		return TagAPILinkClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagAPILinkClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TagAPILinkClientDeleteResponse{}, err
	}
	return TagAPILinkClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TagAPILinkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, apiLinkID string, options *TagAPILinkClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the API link for the tag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Tag-API link identifier. Must be unique in the current API Management service instance.
//   - options - TagAPILinkClientGetOptions contains the optional parameters for the TagAPILinkClient.Get method.
func (client *TagAPILinkClient) Get(ctx context.Context, resourceGroupName string, serviceName string, tagID string, apiLinkID string, options *TagAPILinkClientGetOptions) (TagAPILinkClientGetResponse, error) {
	var err error
	const operationName = "TagAPILinkClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, tagID, apiLinkID, options)
	if err != nil {
		return TagAPILinkClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagAPILinkClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TagAPILinkClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TagAPILinkClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, apiLinkID string, options *TagAPILinkClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagAPILinkClient) getHandleResponse(resp *http.Response) (TagAPILinkClientGetResponse, error) {
	result := TagAPILinkClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagAPILinkContract); err != nil {
		return TagAPILinkClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - Lists a collection of the API links associated with a tag.
//
// Generated from API version 2024-05-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - options - TagAPILinkClientListByProductOptions contains the optional parameters for the TagAPILinkClient.NewListByProductPager
//     method.
func (client *TagAPILinkClient) NewListByProductPager(resourceGroupName string, serviceName string, tagID string, options *TagAPILinkClientListByProductOptions) *runtime.Pager[TagAPILinkClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagAPILinkClientListByProductResponse]{
		More: func(page TagAPILinkClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TagAPILinkClientListByProductResponse) (TagAPILinkClientListByProductResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TagAPILinkClient.NewListByProductPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, tagID, options)
			}, nil)
			if err != nil {
				return TagAPILinkClientListByProductResponse{}, err
			}
			return client.listByProductHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *TagAPILinkClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, options *TagAPILinkClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/apiLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
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
	if options != nil && options.Skip != nil {
		reqQP.Set("$skip", strconv.FormatInt(int64(*options.Skip), 10))
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *TagAPILinkClient) listByProductHandleResponse(resp *http.Response) (TagAPILinkClientListByProductResponse, error) {
	result := TagAPILinkClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagAPILinkCollection); err != nil {
		return TagAPILinkClientListByProductResponse{}, err
	}
	return result, nil
}
