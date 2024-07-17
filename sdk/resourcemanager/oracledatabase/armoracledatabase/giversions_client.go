//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armoracledatabase

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

// GiVersionsClient contains the methods for the GiVersions group.
// Don't use this type directly, use NewGiVersionsClient() instead.
type GiVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGiVersionsClient creates a new instance of GiVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGiVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GiVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GiVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a GiVersion
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - location - The name of the Azure region.
//   - giversionname - GiVersion name
//   - options - GiVersionsClientGetOptions contains the optional parameters for the GiVersionsClient.Get method.
func (client *GiVersionsClient) Get(ctx context.Context, location string, giversionname string, options *GiVersionsClientGetOptions) (GiVersionsClientGetResponse, error) {
	var err error
	const operationName = "GiVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, giversionname, options)
	if err != nil {
		return GiVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GiVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GiVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GiVersionsClient) getCreateRequest(ctx context.Context, location string, giversionname string, options *GiVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/giVersions/{giversionname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if giversionname == "" {
		return nil, errors.New("parameter giversionname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{giversionname}", url.PathEscape(giversionname))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GiVersionsClient) getHandleResponse(resp *http.Response) (GiVersionsClientGetResponse, error) {
	result := GiVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GiVersion); err != nil {
		return GiVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - List GiVersion resources by Location
//
// Generated from API version 2024-06-01
//   - location - The name of the Azure region.
//   - options - GiVersionsClientListByLocationOptions contains the optional parameters for the GiVersionsClient.NewListByLocationPager
//     method.
func (client *GiVersionsClient) NewListByLocationPager(location string, options *GiVersionsClientListByLocationOptions) *runtime.Pager[GiVersionsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[GiVersionsClientListByLocationResponse]{
		More: func(page GiVersionsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GiVersionsClientListByLocationResponse) (GiVersionsClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GiVersionsClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return GiVersionsClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *GiVersionsClient) listByLocationCreateRequest(ctx context.Context, location string, options *GiVersionsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/giVersions"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *GiVersionsClient) listByLocationHandleResponse(resp *http.Response) (GiVersionsClientListByLocationResponse, error) {
	result := GiVersionsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GiVersionListResult); err != nil {
		return GiVersionsClientListByLocationResponse{}, err
	}
	return result, nil
}
