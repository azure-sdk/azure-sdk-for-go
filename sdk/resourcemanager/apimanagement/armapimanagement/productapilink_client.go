//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// ProductAPILinkClient contains the methods for the ProductAPILink group.
// Don't use this type directly, use NewProductAPILinkClient() instead.
type ProductAPILinkClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewProductAPILinkClient creates a new instance of ProductAPILinkClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewProductAPILinkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ProductAPILinkClient, error) {
	cl, err := arm.NewClient(moduleName+".ProductAPILinkClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ProductAPILinkClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Adds an API to the specified product via link.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Product-API link identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - ProductAPILinkClientCreateOrUpdateOptions contains the optional parameters for the ProductAPILinkClient.CreateOrUpdate
//     method.
func (client *ProductAPILinkClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiLinkID string, parameters ProductAPILinkContract, options *ProductAPILinkClientCreateOrUpdateOptions) (ProductAPILinkClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, productID, apiLinkID, parameters, options)
	if err != nil {
		return ProductAPILinkClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductAPILinkClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ProductAPILinkClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ProductAPILinkClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiLinkID string, parameters ProductAPILinkContract, options *ProductAPILinkClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ProductAPILinkClient) createOrUpdateHandleResponse(resp *http.Response) (ProductAPILinkClientCreateOrUpdateResponse, error) {
	result := ProductAPILinkClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductAPILinkContract); err != nil {
		return ProductAPILinkClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified API from the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Product-API link identifier. Must be unique in the current API Management service instance.
//   - options - ProductAPILinkClientDeleteOptions contains the optional parameters for the ProductAPILinkClient.Delete method.
func (client *ProductAPILinkClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiLinkID string, options *ProductAPILinkClientDeleteOptions) (ProductAPILinkClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, productID, apiLinkID, options)
	if err != nil {
		return ProductAPILinkClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductAPILinkClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ProductAPILinkClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ProductAPILinkClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ProductAPILinkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiLinkID string, options *ProductAPILinkClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the API link for the product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - apiLinkID - Product-API link identifier. Must be unique in the current API Management service instance.
//   - options - ProductAPILinkClientGetOptions contains the optional parameters for the ProductAPILinkClient.Get method.
func (client *ProductAPILinkClient) Get(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiLinkID string, options *ProductAPILinkClientGetOptions) (ProductAPILinkClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, productID, apiLinkID, options)
	if err != nil {
		return ProductAPILinkClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ProductAPILinkClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ProductAPILinkClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ProductAPILinkClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, apiLinkID string, options *ProductAPILinkClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apiLinks/{apiLinkId}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ProductAPILinkClient) getHandleResponse(resp *http.Response) (ProductAPILinkClientGetResponse, error) {
	result := ProductAPILinkClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductAPILinkContract); err != nil {
		return ProductAPILinkClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - Lists a collection of the API links associated with a product.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - options - ProductAPILinkClientListByProductOptions contains the optional parameters for the ProductAPILinkClient.NewListByProductPager
//     method.
func (client *ProductAPILinkClient) NewListByProductPager(resourceGroupName string, serviceName string, productID string, options *ProductAPILinkClientListByProductOptions) *runtime.Pager[ProductAPILinkClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[ProductAPILinkClientListByProductResponse]{
		More: func(page ProductAPILinkClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ProductAPILinkClientListByProductResponse) (ProductAPILinkClientListByProductResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, productID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ProductAPILinkClientListByProductResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ProductAPILinkClientListByProductResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ProductAPILinkClientListByProductResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProductHandleResponse(resp)
		},
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *ProductAPILinkClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, productID string, options *ProductAPILinkClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/apiLinks"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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
	reqQP.Set("api-version", "2023-03-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByProductHandleResponse handles the ListByProduct response.
func (client *ProductAPILinkClient) listByProductHandleResponse(resp *http.Response) (ProductAPILinkClientListByProductResponse, error) {
	result := ProductAPILinkClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductAPILinkCollection); err != nil {
		return ProductAPILinkClientListByProductResponse{}, err
	}
	return result, nil
}
