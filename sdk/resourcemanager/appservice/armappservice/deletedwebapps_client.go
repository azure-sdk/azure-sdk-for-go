//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappservice

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

// DeletedWebAppsClient contains the methods for the DeletedWebApps group.
// Don't use this type directly, use NewDeletedWebAppsClient() instead.
type DeletedWebAppsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDeletedWebAppsClient creates a new instance of DeletedWebAppsClient with the specified values.
//   - subscriptionID - Your Azure subscription ID. This is a GUID-formatted string (e.g. 00000000-0000-0000-0000-000000000000).
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDeletedWebAppsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DeletedWebAppsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DeletedWebAppsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// GetDeletedWebAppByLocation - Description for Get deleted app for a subscription at location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-01-01
//   - deletedSiteID - The numeric ID of the deleted app, e.g. 12345
//   - options - DeletedWebAppsClientGetDeletedWebAppByLocationOptions contains the optional parameters for the DeletedWebAppsClient.GetDeletedWebAppByLocation
//     method.
func (client *DeletedWebAppsClient) GetDeletedWebAppByLocation(ctx context.Context, location string, deletedSiteID string, options *DeletedWebAppsClientGetDeletedWebAppByLocationOptions) (DeletedWebAppsClientGetDeletedWebAppByLocationResponse, error) {
	var err error
	const operationName = "DeletedWebAppsClient.GetDeletedWebAppByLocation"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getDeletedWebAppByLocationCreateRequest(ctx, location, deletedSiteID, options)
	if err != nil {
		return DeletedWebAppsClientGetDeletedWebAppByLocationResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DeletedWebAppsClientGetDeletedWebAppByLocationResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DeletedWebAppsClientGetDeletedWebAppByLocationResponse{}, err
	}
	resp, err := client.getDeletedWebAppByLocationHandleResponse(httpResp)
	return resp, err
}

// getDeletedWebAppByLocationCreateRequest creates the GetDeletedWebAppByLocation request.
func (client *DeletedWebAppsClient) getDeletedWebAppByLocationCreateRequest(ctx context.Context, location string, deletedSiteID string, options *DeletedWebAppsClientGetDeletedWebAppByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites/{deletedSiteId}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if deletedSiteID == "" {
		return nil, errors.New("parameter deletedSiteID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{deletedSiteId}", url.PathEscape(deletedSiteID))
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

// getDeletedWebAppByLocationHandleResponse handles the GetDeletedWebAppByLocation response.
func (client *DeletedWebAppsClient) getDeletedWebAppByLocationHandleResponse(resp *http.Response) (DeletedWebAppsClientGetDeletedWebAppByLocationResponse, error) {
	result := DeletedWebAppsClientGetDeletedWebAppByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedSite); err != nil {
		return DeletedWebAppsClientGetDeletedWebAppByLocationResponse{}, err
	}
	return result, nil
}

// NewListPager - Description for Get all deleted apps for a subscription.
//
// Generated from API version 2024-01-01
//   - options - DeletedWebAppsClientListOptions contains the optional parameters for the DeletedWebAppsClient.NewListPager method.
func (client *DeletedWebAppsClient) NewListPager(options *DeletedWebAppsClientListOptions) *runtime.Pager[DeletedWebAppsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedWebAppsClientListResponse]{
		More: func(page DeletedWebAppsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedWebAppsClientListResponse) (DeletedWebAppsClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeletedWebAppsClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return DeletedWebAppsClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *DeletedWebAppsClient) listCreateRequest(ctx context.Context, options *DeletedWebAppsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/deletedSites"
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

// listHandleResponse handles the List response.
func (client *DeletedWebAppsClient) listHandleResponse(resp *http.Response) (DeletedWebAppsClientListResponse, error) {
	result := DeletedWebAppsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedWebAppCollection); err != nil {
		return DeletedWebAppsClientListResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Description for Get all deleted apps for a subscription at location
//
// Generated from API version 2024-01-01
//   - options - DeletedWebAppsClientListByLocationOptions contains the optional parameters for the DeletedWebAppsClient.NewListByLocationPager
//     method.
func (client *DeletedWebAppsClient) NewListByLocationPager(location string, options *DeletedWebAppsClientListByLocationOptions) *runtime.Pager[DeletedWebAppsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[DeletedWebAppsClientListByLocationResponse]{
		More: func(page DeletedWebAppsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DeletedWebAppsClientListByLocationResponse) (DeletedWebAppsClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DeletedWebAppsClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return DeletedWebAppsClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *DeletedWebAppsClient) listByLocationCreateRequest(ctx context.Context, location string, options *DeletedWebAppsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Web/locations/{location}/deletedSites"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
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

// listByLocationHandleResponse handles the ListByLocation response.
func (client *DeletedWebAppsClient) listByLocationHandleResponse(resp *http.Response) (DeletedWebAppsClientListByLocationResponse, error) {
	result := DeletedWebAppsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DeletedWebAppCollection); err != nil {
		return DeletedWebAppsClientListByLocationResponse{}, err
	}
	return result, nil
}
