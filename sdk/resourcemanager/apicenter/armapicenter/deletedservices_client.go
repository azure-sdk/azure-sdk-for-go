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

// DeletedServicesClient contains the methods for the DeletedServices group.
// Don't use this type directly, use NewDeletedServicesClient() instead.
type DeletedServicesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeletedServicesClient creates a new instance of DeletedServicesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeletedServicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeletedServicesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeletedServicesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Delete - Permanently deletes specified service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deletedServiceName - The name of the deleted service.
//   - options - DeletedServicesClientDeleteOptions contains the optional parameters for the DeletedServicesClient.Delete method.
func (client *DeletedServicesClient) Delete(ctx context.Context, resourceGroupName string, deletedServiceName string, options *DeletedServicesClientDeleteOptions) (DeletedServicesClientDeleteResponse, error) {
	var err error
	const operationName = "DeletedServicesClient.Delete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, deletedServiceName, options)
	if err != nil {
		return DeletedServicesClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeletedServicesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return DeletedServicesClientDeleteResponse{}, err
	}
	return DeletedServicesClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *DeletedServicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, deletedServiceName string, options *DeletedServicesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/deletedServices/{deletedServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deletedServiceName == "" {
		return nil, errors.New("parameter deletedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedServiceName}", url.PathEscape(deletedServiceName))
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

// Get - Returns details of the soft-deleted service.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - deletedServiceName - The name of the deleted service.
//   - options - DeletedServicesClientGetOptions contains the optional parameters for the DeletedServicesClient.Get method.
func (client *DeletedServicesClient) Get(ctx context.Context, resourceGroupName string, deletedServiceName string, options *DeletedServicesClientGetOptions) (DeletedServicesClientGetResponse, error) {
	var err error
	const operationName = "DeletedServicesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, deletedServiceName, options)
	if err != nil {
		return DeletedServicesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeletedServicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeletedServicesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DeletedServicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, deletedServiceName string, options *DeletedServicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/deletedServices/{deletedServiceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if deletedServiceName == "" {
		return nil, errors.New("parameter deletedServiceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedServiceName}", url.PathEscape(deletedServiceName))
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
func (client *DeletedServicesClient) getHandleResponse(resp *http.Response) (DeletedServicesClientGetResponse, error) {
	result := DeletedServicesClientGetResponse{}
	if val := resp.Header.Get("ETag"); val != "" {
		result.ETag = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedService); err != nil {
		return DeletedServicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists soft-deleted services.
//
// Generated from API version 2024-03-15-preview
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - options - DeletedServicesClientListOptions contains the optional parameters for the DeletedServicesClient.NewListPager
//     method.
func (client *DeletedServicesClient) NewListPager(resourceGroupName string, options *DeletedServicesClientListOptions) *runtime.Pager[DeletedServicesClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedServicesClientListResponse]{
		More: func(page DeletedServicesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedServicesClientListResponse) (DeletedServicesClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeletedServicesClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return DeletedServicesClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DeletedServicesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *DeletedServicesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiCenter/deletedServices"
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
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	reqQP.Set("api-version", "2024-03-15-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *DeletedServicesClient) listHandleResponse(resp *http.Response) (DeletedServicesClientListResponse, error) {
	result := DeletedServicesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServiceListResult); err != nil {
		return DeletedServicesClientListResponse{}, err
	}
	return result, nil
}

// NewListBySubscriptionPager - Lists services within an Azure subscription.
//
// Generated from API version 2024-03-15-preview
//   - options - DeletedServicesClientListBySubscriptionOptions contains the optional parameters for the DeletedServicesClient.NewListBySubscriptionPager
//     method.
func (client *DeletedServicesClient) NewListBySubscriptionPager(options *DeletedServicesClientListBySubscriptionOptions) *runtime.Pager[DeletedServicesClientListBySubscriptionResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedServicesClientListBySubscriptionResponse]{
		More: func(page DeletedServicesClientListBySubscriptionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedServicesClientListBySubscriptionResponse) (DeletedServicesClientListBySubscriptionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeletedServicesClient.NewListBySubscriptionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listBySubscriptionCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DeletedServicesClientListBySubscriptionResponse{}, err
			}
			return client.listBySubscriptionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *DeletedServicesClient) listBySubscriptionCreateRequest(ctx context.Context, options *DeletedServicesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ApiCenter/deletedServices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *DeletedServicesClient) listBySubscriptionHandleResponse(resp *http.Response) (DeletedServicesClientListBySubscriptionResponse, error) {
	result := DeletedServicesClientListBySubscriptionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedServiceListResult); err != nil {
		return DeletedServicesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}
