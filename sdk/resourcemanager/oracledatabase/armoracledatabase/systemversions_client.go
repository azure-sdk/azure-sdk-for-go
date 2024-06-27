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

// SystemVersionsClient contains the methods for the SystemVersions group.
// Don't use this type directly, use NewSystemVersionsClient() instead.
type SystemVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewSystemVersionsClient creates a new instance of SystemVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewSystemVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SystemVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &SystemVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a SystemVersion
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - location - The name of the Azure region.
//   - systemversionname - SystemVersion name
//   - options - SystemVersionsClientGetOptions contains the optional parameters for the SystemVersionsClient.Get method.
func (client *SystemVersionsClient) Get(ctx context.Context, location string, systemversionname string, options *SystemVersionsClientGetOptions) (SystemVersionsClientGetResponse, error) {
	var err error
	const operationName = "SystemVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, systemversionname, options)
	if err != nil {
		return SystemVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return SystemVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return SystemVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *SystemVersionsClient) getCreateRequest(ctx context.Context, location string, systemversionname string, options *SystemVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/systemVersions/{systemversionname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if systemversionname == "" {
		return nil, errors.New("parameter systemversionname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{systemversionname}", url.PathEscape(systemversionname))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SystemVersionsClient) getHandleResponse(resp *http.Response) (SystemVersionsClientGetResponse, error) {
	result := SystemVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemVersion); err != nil {
		return SystemVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - List SystemVersion resources by Location
//
// Generated from API version 2023-09-01
//   - location - The name of the Azure region.
//   - options - SystemVersionsClientListByLocationOptions contains the optional parameters for the SystemVersionsClient.NewListByLocationPager
//     method.
func (client *SystemVersionsClient) NewListByLocationPager(location string, options *SystemVersionsClientListByLocationOptions) *runtime.Pager[SystemVersionsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[SystemVersionsClientListByLocationResponse]{
		More: func(page SystemVersionsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SystemVersionsClientListByLocationResponse) (SystemVersionsClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "SystemVersionsClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return SystemVersionsClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *SystemVersionsClient) listByLocationCreateRequest(ctx context.Context, location string, options *SystemVersionsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/systemVersions"
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
	reqQP.Set("api-version", "2023-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *SystemVersionsClient) listByLocationHandleResponse(resp *http.Response) (SystemVersionsClientListByLocationResponse, error) {
	result := SystemVersionsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SystemVersionListResult); err != nil {
		return SystemVersionsClientListByLocationResponse{}, err
	}
	return result, nil
}
