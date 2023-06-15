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

// WorkspaceProductGroupLinkClient contains the methods for the WorkspaceProductGroupLink group.
// Don't use this type directly, use NewWorkspaceProductGroupLinkClient() instead.
type WorkspaceProductGroupLinkClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewWorkspaceProductGroupLinkClient creates a new instance of WorkspaceProductGroupLinkClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewWorkspaceProductGroupLinkClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*WorkspaceProductGroupLinkClient, error) {
	cl, err := arm.NewClient(moduleName+".WorkspaceProductGroupLinkClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &WorkspaceProductGroupLinkClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Adds a group to the specified product via link.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - groupLinkID - Product-Group link identifier. Must be unique in the current API Management service instance.
//   - parameters - Create or update parameters.
//   - options - WorkspaceProductGroupLinkClientCreateOrUpdateOptions contains the optional parameters for the WorkspaceProductGroupLinkClient.CreateOrUpdate
//     method.
func (client *WorkspaceProductGroupLinkClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, groupLinkID string, parameters ProductGroupLinkContract, options *WorkspaceProductGroupLinkClientCreateOrUpdateOptions) (WorkspaceProductGroupLinkClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, productID, groupLinkID, parameters, options)
	if err != nil {
		return WorkspaceProductGroupLinkClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceProductGroupLinkClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return WorkspaceProductGroupLinkClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspaceProductGroupLinkClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, groupLinkID string, parameters ProductGroupLinkContract, options *WorkspaceProductGroupLinkClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/products/{productId}/groupLinks/{groupLinkId}"
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
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupLinkID == "" {
		return nil, errors.New("parameter groupLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupLinkId}", url.PathEscape(groupLinkID))
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
func (client *WorkspaceProductGroupLinkClient) createOrUpdateHandleResponse(resp *http.Response) (WorkspaceProductGroupLinkClientCreateOrUpdateResponse, error) {
	result := WorkspaceProductGroupLinkClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductGroupLinkContract); err != nil {
		return WorkspaceProductGroupLinkClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the specified group from the specified product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - groupLinkID - Product-Group link identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceProductGroupLinkClientDeleteOptions contains the optional parameters for the WorkspaceProductGroupLinkClient.Delete
//     method.
func (client *WorkspaceProductGroupLinkClient) Delete(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, groupLinkID string, options *WorkspaceProductGroupLinkClientDeleteOptions) (WorkspaceProductGroupLinkClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, productID, groupLinkID, options)
	if err != nil {
		return WorkspaceProductGroupLinkClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceProductGroupLinkClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return WorkspaceProductGroupLinkClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return WorkspaceProductGroupLinkClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspaceProductGroupLinkClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, groupLinkID string, options *WorkspaceProductGroupLinkClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/products/{productId}/groupLinks/{groupLinkId}"
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
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupLinkID == "" {
		return nil, errors.New("parameter groupLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupLinkId}", url.PathEscape(groupLinkID))
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

// Get - Gets the group link for the product.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - groupLinkID - Product-Group link identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceProductGroupLinkClientGetOptions contains the optional parameters for the WorkspaceProductGroupLinkClient.Get
//     method.
func (client *WorkspaceProductGroupLinkClient) Get(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, groupLinkID string, options *WorkspaceProductGroupLinkClientGetOptions) (WorkspaceProductGroupLinkClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, productID, groupLinkID, options)
	if err != nil {
		return WorkspaceProductGroupLinkClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return WorkspaceProductGroupLinkClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspaceProductGroupLinkClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspaceProductGroupLinkClient) getCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, groupLinkID string, options *WorkspaceProductGroupLinkClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/products/{productId}/groupLinks/{groupLinkId}"
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
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
	if groupLinkID == "" {
		return nil, errors.New("parameter groupLinkID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{groupLinkId}", url.PathEscape(groupLinkID))
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
func (client *WorkspaceProductGroupLinkClient) getHandleResponse(resp *http.Response) (WorkspaceProductGroupLinkClientGetResponse, error) {
	result := WorkspaceProductGroupLinkClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductGroupLinkContract); err != nil {
		return WorkspaceProductGroupLinkClientGetResponse{}, err
	}
	return result, nil
}

// NewListByProductPager - Lists a collection of the group links associated with a product.
//
// Generated from API version 2023-03-01-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - workspaceID - Workspace identifier. Must be unique in the current API Management service instance.
//   - productID - Product identifier. Must be unique in the current API Management service instance.
//   - options - WorkspaceProductGroupLinkClientListByProductOptions contains the optional parameters for the WorkspaceProductGroupLinkClient.NewListByProductPager
//     method.
func (client *WorkspaceProductGroupLinkClient) NewListByProductPager(resourceGroupName string, serviceName string, workspaceID string, productID string, options *WorkspaceProductGroupLinkClientListByProductOptions) *runtime.Pager[WorkspaceProductGroupLinkClientListByProductResponse] {
	return runtime.NewPager(runtime.PagingHandler[WorkspaceProductGroupLinkClientListByProductResponse]{
		More: func(page WorkspaceProductGroupLinkClientListByProductResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *WorkspaceProductGroupLinkClientListByProductResponse) (WorkspaceProductGroupLinkClientListByProductResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByProductCreateRequest(ctx, resourceGroupName, serviceName, workspaceID, productID, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return WorkspaceProductGroupLinkClientListByProductResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return WorkspaceProductGroupLinkClientListByProductResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return WorkspaceProductGroupLinkClientListByProductResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByProductHandleResponse(resp)
		},
	})
}

// listByProductCreateRequest creates the ListByProduct request.
func (client *WorkspaceProductGroupLinkClient) listByProductCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, workspaceID string, productID string, options *WorkspaceProductGroupLinkClientListByProductOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/workspaces/{workspaceId}/products/{productId}/groupLinks"
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
	if productID == "" {
		return nil, errors.New("parameter productID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{productId}", url.PathEscape(productID))
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
func (client *WorkspaceProductGroupLinkClient) listByProductHandleResponse(resp *http.Response) (WorkspaceProductGroupLinkClientListByProductResponse, error) {
	result := WorkspaceProductGroupLinkClientListByProductResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ProductGroupLinkCollection); err != nil {
		return WorkspaceProductGroupLinkClientListByProductResponse{}, err
	}
	return result, nil
}
