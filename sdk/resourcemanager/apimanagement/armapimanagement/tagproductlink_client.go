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

// TagProductLinkClient contains the methods for the TagProductLink group.
// Don't use this type directly, use NewTagProductLinkClient() instead.
type TagProductLinkClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTagProductLinkClient creates a new instance of TagProductLinkClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTagProductLinkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TagProductLinkClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TagProductLinkClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Adds a product to the specified tag via link.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - productLinkID - Tag-product link identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - TagProductLinkClientCreateOrUpdateOptions contains the optional parameters for the TagProductLinkClient.CreateOrUpdate
//     method.
func (client *TagProductLinkClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, tagID string, productLinkID string, parameters TagProductLinkContract, options *TagProductLinkClientCreateOrUpdateOptions) (TagProductLinkClientCreateOrUpdateResponse, error) {
	var err error
	const operationName = "TagProductLinkClient.CreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, tagID, productLinkID, parameters, options)
	if err != nil {
		return TagProductLinkClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagProductLinkClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return TagProductLinkClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *TagProductLinkClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, productLinkID string, parameters TagProductLinkContract, _ *TagProductLinkClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/productLinks/{productLinkId}"
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
	if productLinkID == "" {
		return nil, errors.New("parameter productLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productLinkId}", url.PathEscape(productLinkID))
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
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *TagProductLinkClient) createOrUpdateHandleResponse(resp *http.Response) (TagProductLinkClientCreateOrUpdateResponse, error) {
	result := TagProductLinkClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagProductLinkContract); err != nil {
		return TagProductLinkClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified product from the specified tag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - productLinkID - Tag-product link identifier. Must be unique in the current API Management service instance.
//   - options - TagProductLinkClientDeleteOptions contains the optional parameters for the TagProductLinkClient.Delete method.
func (client *TagProductLinkClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, tagID string, productLinkID string, options *TagProductLinkClientDeleteOptions) (TagProductLinkClientDeleteResponse, error) {
	var err error
	const operationName = "TagProductLinkClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, tagID, productLinkID, options)
	if err != nil {
		return TagProductLinkClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagProductLinkClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return TagProductLinkClientDeleteResponse{}, err
	}
	return TagProductLinkClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *TagProductLinkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, productLinkID string, _ *TagProductLinkClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/productLinks/{productLinkId}"
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
	if productLinkID == "" {
		return nil, errors.New("parameter productLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productLinkId}", url.PathEscape(productLinkID))
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
	return req, nil
}

// Get - Gets the product link for the tag.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - productLinkID - Tag-product link identifier. Must be unique in the current API Management service instance.
//   - options - TagProductLinkClientGetOptions contains the optional parameters for the TagProductLinkClient.Get method.
func (client *TagProductLinkClient) Get(ctx context.Context, resourceGroupName string, serviceName string, tagID string, productLinkID string, options *TagProductLinkClientGetOptions) (TagProductLinkClientGetResponse, error) {
	var err error
	const operationName = "TagProductLinkClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, tagID, productLinkID, options)
	if err != nil {
		return TagProductLinkClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TagProductLinkClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TagProductLinkClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TagProductLinkClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, productLinkID string, _ *TagProductLinkClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/productLinks/{productLinkId}"
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
	if productLinkID == "" {
		return nil, errors.New("parameter productLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productLinkId}", url.PathEscape(productLinkID))
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
func (client *TagProductLinkClient) getHandleResponse(resp *http.Response) (TagProductLinkClientGetResponse, error) {
	result := TagProductLinkClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagProductLinkContract); err != nil {
		return TagProductLinkClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - Lists a collection of the product links associated with a tag.
//
// Generated from API version 2024-06-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - tagID - Tag identifier. Must be unique in the current API Management service instance.
//   - options - TagProductLinkClientListByProductOptions contains the optional parameters for the TagProductLinkClient.NewListByProductPager
//     method.
func (client *TagProductLinkClient) NewListByProductPager(resourceGroupName string, serviceName string, tagID string, options *TagProductLinkClientListByProductOptions) *runtime.Pager[TagProductLinkClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[TagProductLinkClientListByProductResponse]{
		More: func(page TagProductLinkClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TagProductLinkClientListByProductResponse) (TagProductLinkClientListByProductResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TagProductLinkClient.NewListByProductPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, tagID, options)
			}, nil)
			if err != nil {
				return TagProductLinkClientListByProductResponse{}, err
			}
			return client.listByProductHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *TagProductLinkClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, tagID string, options *TagProductLinkClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/tags/{tagId}/productLinks"
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
	reqQP.Set("api-version", "2024-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *TagProductLinkClient) listByProductHandleResponse(resp *http.Response) (TagProductLinkClientListByProductResponse, error) {
	result := TagProductLinkClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagProductLinkCollection); err != nil {
		return TagProductLinkClientListByProductResponse{}, err
	}
	return result, nil
}
