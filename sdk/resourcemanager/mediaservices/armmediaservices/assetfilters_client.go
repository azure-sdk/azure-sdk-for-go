//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmediaservices

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

// AssetFiltersClient contains the methods for the AssetFilters group.
// Don't use this type directly, use NewAssetFiltersClient() instead.
type AssetFiltersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAssetFiltersClient creates a new instance of AssetFiltersClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAssetFiltersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AssetFiltersClient, error) {
	cl, err := arm.NewClient(moduleName+".AssetFiltersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AssetFiltersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates or updates an Asset Filter associated with the specified Asset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - assetName - The Asset name.
//   - filterName - The Asset Filter name
//   - parameters - The request parameters
//   - options - AssetFiltersClientCreateOrUpdateOptions contains the optional parameters for the AssetFiltersClient.CreateOrUpdate
//     method.
func (client *AssetFiltersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, parameters AssetFilter, options *AssetFiltersClientCreateOrUpdateOptions) (AssetFiltersClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, accountName, assetName, filterName, parameters, options)
	if err != nil {
		return AssetFiltersClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssetFiltersClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return AssetFiltersClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AssetFiltersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, parameters AssetFilter, options *AssetFiltersClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *AssetFiltersClient) createOrUpdateHandleResponse(resp *http.Response) (AssetFiltersClientCreateOrUpdateResponse, error) {
	result := AssetFiltersClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetFilter); err != nil {
		return AssetFiltersClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes an Asset Filter associated with the specified Asset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - assetName - The Asset name.
//   - filterName - The Asset Filter name
//   - options - AssetFiltersClientDeleteOptions contains the optional parameters for the AssetFiltersClient.Delete method.
func (client *AssetFiltersClient) Delete(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, options *AssetFiltersClientDeleteOptions) (AssetFiltersClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, accountName, assetName, filterName, options)
	if err != nil {
		return AssetFiltersClientDeleteResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssetFiltersClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return AssetFiltersClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return AssetFiltersClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AssetFiltersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, options *AssetFiltersClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get the details of an Asset Filter associated with the specified Asset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - assetName - The Asset name.
//   - filterName - The Asset Filter name
//   - options - AssetFiltersClientGetOptions contains the optional parameters for the AssetFiltersClient.Get method.
func (client *AssetFiltersClient) Get(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, options *AssetFiltersClientGetOptions) (AssetFiltersClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, accountName, assetName, filterName, options)
	if err != nil {
		return AssetFiltersClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssetFiltersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetFiltersClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssetFiltersClient) getCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, options *AssetFiltersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssetFiltersClient) getHandleResponse(resp *http.Response) (AssetFiltersClientGetResponse, error) {
	result := AssetFiltersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetFilter); err != nil {
		return AssetFiltersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List Asset Filters associated with the specified Asset.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - assetName - The Asset name.
//   - options - AssetFiltersClientListOptions contains the optional parameters for the AssetFiltersClient.NewListPager method.
func (client *AssetFiltersClient) NewListPager(resourceGroupName string, accountName string, assetName string, options *AssetFiltersClientListOptions) *runtime.Pager[AssetFiltersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssetFiltersClientListResponse]{
		More: func(page AssetFiltersClientListResponse) bool {
			return page.ODataNextLink != nil && len(*page.ODataNextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssetFiltersClientListResponse) (AssetFiltersClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, accountName, assetName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.ODataNextLink)
			}
			if err != nil {
				return AssetFiltersClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AssetFiltersClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssetFiltersClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AssetFiltersClient) listCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, options *AssetFiltersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AssetFiltersClient) listHandleResponse(resp *http.Response) (AssetFiltersClientListResponse, error) {
	result := AssetFiltersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetFilterCollection); err != nil {
		return AssetFiltersClientListResponse{}, err
	}
	return result, nil
}

// Update - Updates an existing Asset Filter associated with the specified Asset.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-01-01
//   - resourceGroupName - The name of the resource group within the Azure subscription.
//   - accountName - The Media Services account name.
//   - assetName - The Asset name.
//   - filterName - The Asset Filter name
//   - parameters - The request parameters
//   - options - AssetFiltersClientUpdateOptions contains the optional parameters for the AssetFiltersClient.Update method.
func (client *AssetFiltersClient) Update(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, parameters AssetFilter, options *AssetFiltersClientUpdateOptions) (AssetFiltersClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, accountName, assetName, filterName, parameters, options)
	if err != nil {
		return AssetFiltersClientUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AssetFiltersClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssetFiltersClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *AssetFiltersClient) updateCreateRequest(ctx context.Context, resourceGroupName string, accountName string, assetName string, filterName string, parameters AssetFilter, options *AssetFiltersClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Media/mediaServices/{accountName}/assets/{assetName}/assetFilters/{filterName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if accountName == "" {
		return nil, errors.New("parameter accountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{accountName}", url.PathEscape(accountName))
	if assetName == "" {
		return nil, errors.New("parameter assetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{assetName}", url.PathEscape(assetName))
	if filterName == "" {
		return nil, errors.New("parameter filterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{filterName}", url.PathEscape(filterName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateHandleResponse handles the Update response.
func (client *AssetFiltersClient) updateHandleResponse(resp *http.Response) (AssetFiltersClientUpdateResponse, error) {
	result := AssetFiltersClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssetFilter); err != nil {
		return AssetFiltersClientUpdateResponse{}, err
	}
	return result, nil
}
