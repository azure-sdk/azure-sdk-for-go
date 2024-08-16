//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

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

// MarketplaceGalleryImagesClient contains the methods for the MarketplaceGalleryImages group.
// Don't use this type directly, use NewMarketplaceGalleryImagesClient() instead.
type MarketplaceGalleryImagesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMarketplaceGalleryImagesClient creates a new instance of MarketplaceGalleryImagesClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMarketplaceGalleryImagesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MarketplaceGalleryImagesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MarketplaceGalleryImagesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - The operation to create or update a marketplace gallery image. Please note some properties can be
// set only during marketplace gallery image creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - marketplaceGalleryImageName - Name of the marketplace gallery image
//   - options - MarketplaceGalleryImagesClientBeginCreateOrUpdateOptions contains the optional parameters for the MarketplaceGalleryImagesClient.BeginCreateOrUpdate
//     method.
func (client *MarketplaceGalleryImagesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, marketplaceGalleryImages MarketplaceGalleryImages, options *MarketplaceGalleryImagesClientBeginCreateOrUpdateOptions) (*runtime.Poller[MarketplaceGalleryImagesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, marketplaceGalleryImageName, marketplaceGalleryImages, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MarketplaceGalleryImagesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MarketplaceGalleryImagesClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - The operation to create or update a marketplace gallery image. Please note some properties can be set
// only during marketplace gallery image creation.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *MarketplaceGalleryImagesClient) createOrUpdate(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, marketplaceGalleryImages MarketplaceGalleryImages, options *MarketplaceGalleryImagesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MarketplaceGalleryImagesClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, marketplaceGalleryImageName, marketplaceGalleryImages, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MarketplaceGalleryImagesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, marketplaceGalleryImages MarketplaceGalleryImages, options *MarketplaceGalleryImagesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/{marketplaceGalleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if marketplaceGalleryImageName == "" {
		return nil, errors.New("parameter marketplaceGalleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{marketplaceGalleryImageName}", url.PathEscape(marketplaceGalleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, marketplaceGalleryImages); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - The operation to delete a marketplace gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - marketplaceGalleryImageName - Name of the marketplace gallery image
//   - options - MarketplaceGalleryImagesClientBeginDeleteOptions contains the optional parameters for the MarketplaceGalleryImagesClient.BeginDelete
//     method.
func (client *MarketplaceGalleryImagesClient) BeginDelete(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, options *MarketplaceGalleryImagesClientBeginDeleteOptions) (*runtime.Poller[MarketplaceGalleryImagesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, marketplaceGalleryImageName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MarketplaceGalleryImagesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MarketplaceGalleryImagesClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - The operation to delete a marketplace gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *MarketplaceGalleryImagesClient) deleteOperation(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, options *MarketplaceGalleryImagesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "MarketplaceGalleryImagesClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, marketplaceGalleryImageName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *MarketplaceGalleryImagesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, options *MarketplaceGalleryImagesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/{marketplaceGalleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if marketplaceGalleryImageName == "" {
		return nil, errors.New("parameter marketplaceGalleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{marketplaceGalleryImageName}", url.PathEscape(marketplaceGalleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a marketplace gallery image
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - marketplaceGalleryImageName - Name of the marketplace gallery image
//   - options - MarketplaceGalleryImagesClientGetOptions contains the optional parameters for the MarketplaceGalleryImagesClient.Get
//     method.
func (client *MarketplaceGalleryImagesClient) Get(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, options *MarketplaceGalleryImagesClientGetOptions) (MarketplaceGalleryImagesClientGetResponse, error) {
	var err error
	const operationName = "MarketplaceGalleryImagesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, marketplaceGalleryImageName, options)
	if err != nil {
		return MarketplaceGalleryImagesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MarketplaceGalleryImagesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MarketplaceGalleryImagesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MarketplaceGalleryImagesClient) getCreateRequest(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, options *MarketplaceGalleryImagesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/{marketplaceGalleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if marketplaceGalleryImageName == "" {
		return nil, errors.New("parameter marketplaceGalleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{marketplaceGalleryImageName}", url.PathEscape(marketplaceGalleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MarketplaceGalleryImagesClient) getHandleResponse(resp *http.Response) (MarketplaceGalleryImagesClientGetResponse, error) {
	result := MarketplaceGalleryImagesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceGalleryImages); err != nil {
		return MarketplaceGalleryImagesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all of the marketplace gallery images in the specified resource group. Use the nextLink property in
// the response to get the next page of marketplace gallery images.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - MarketplaceGalleryImagesClientListOptions contains the optional parameters for the MarketplaceGalleryImagesClient.NewListPager
//     method.
func (client *MarketplaceGalleryImagesClient) NewListPager(resourceGroupName string, options *MarketplaceGalleryImagesClientListOptions) *runtime.Pager[MarketplaceGalleryImagesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MarketplaceGalleryImagesClientListResponse]{
		More: func(page MarketplaceGalleryImagesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MarketplaceGalleryImagesClientListResponse) (MarketplaceGalleryImagesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MarketplaceGalleryImagesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return MarketplaceGalleryImagesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *MarketplaceGalleryImagesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *MarketplaceGalleryImagesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MarketplaceGalleryImagesClient) listHandleResponse(resp *http.Response) (MarketplaceGalleryImagesClientListResponse, error) {
	result := MarketplaceGalleryImagesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceGalleryImagesListResult); err != nil {
		return MarketplaceGalleryImagesClientListResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Lists all of the marketplace gallery images in the specified subscription. Use the nextLink property
// in the response to get the next page of marketplace gallery images.
//
// Generated from API version 2024-01-01
//   - options - MarketplaceGalleryImagesClientListAllOptions contains the optional parameters for the MarketplaceGalleryImagesClient.NewListAllPager
//     method.
func (client *MarketplaceGalleryImagesClient) NewListAllPager(options *MarketplaceGalleryImagesClientListAllOptions) *runtime.Pager[MarketplaceGalleryImagesClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[MarketplaceGalleryImagesClientListAllResponse]{
		More: func(page MarketplaceGalleryImagesClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MarketplaceGalleryImagesClientListAllResponse) (MarketplaceGalleryImagesClientListAllResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MarketplaceGalleryImagesClient.NewListAllPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listAllCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return MarketplaceGalleryImagesClientListAllResponse{}, err
			}
			return client.listAllHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *MarketplaceGalleryImagesClient) listAllCreateRequest(ctx context.Context, options *MarketplaceGalleryImagesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *MarketplaceGalleryImagesClient) listAllHandleResponse(resp *http.Response) (MarketplaceGalleryImagesClientListAllResponse, error) {
	result := MarketplaceGalleryImagesClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MarketplaceGalleryImagesListResult); err != nil {
		return MarketplaceGalleryImagesClientListAllResponse{}, err
	}
	return result, nil
}

// BeginUpdate - The operation to update a marketplace gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - marketplaceGalleryImageName - Name of the marketplace gallery image
//   - options - MarketplaceGalleryImagesClientBeginUpdateOptions contains the optional parameters for the MarketplaceGalleryImagesClient.BeginUpdate
//     method.
func (client *MarketplaceGalleryImagesClient) BeginUpdate(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, marketplaceGalleryImages MarketplaceGalleryImagesUpdateRequest, options *MarketplaceGalleryImagesClientBeginUpdateOptions) (*runtime.Poller[MarketplaceGalleryImagesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, resourceGroupName, marketplaceGalleryImageName, marketplaceGalleryImages, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[MarketplaceGalleryImagesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[MarketplaceGalleryImagesClientUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Update - The operation to update a marketplace gallery image.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
func (client *MarketplaceGalleryImagesClient) update(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, marketplaceGalleryImages MarketplaceGalleryImagesUpdateRequest, options *MarketplaceGalleryImagesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "MarketplaceGalleryImagesClient.BeginUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, resourceGroupName, marketplaceGalleryImageName, marketplaceGalleryImages, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *MarketplaceGalleryImagesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, marketplaceGalleryImageName string, marketplaceGalleryImages MarketplaceGalleryImagesUpdateRequest, options *MarketplaceGalleryImagesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.AzureStackHCI/marketplaceGalleryImages/{marketplaceGalleryImageName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if marketplaceGalleryImageName == "" {
		return nil, errors.New("parameter marketplaceGalleryImageName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{marketplaceGalleryImageName}", url.PathEscape(marketplaceGalleryImageName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, marketplaceGalleryImages); err != nil {
		return nil, err
	}
	return req, nil
}
