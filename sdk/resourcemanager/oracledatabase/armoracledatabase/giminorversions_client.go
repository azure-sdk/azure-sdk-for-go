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

// GiMinorVersionsClient contains the methods for the GiMinorVersions group.
// Don't use this type directly, use NewGiMinorVersionsClient() instead.
type GiMinorVersionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewGiMinorVersionsClient creates a new instance of GiMinorVersionsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewGiMinorVersionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*GiMinorVersionsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &GiMinorVersionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a GiMinorVersion
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - location - The name of the Azure region.
//   - giversionname - GiVersion name
//   - giMinorVersionName - The name of the GiMinorVersion
//   - options - GiMinorVersionsClientGetOptions contains the optional parameters for the GiMinorVersionsClient.Get method.
func (client *GiMinorVersionsClient) Get(ctx context.Context, location string, giversionname string, giMinorVersionName string, options *GiMinorVersionsClientGetOptions) (GiMinorVersionsClientGetResponse, error) {
	var err error
	const operationName = "GiMinorVersionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, giversionname, giMinorVersionName, options)
	if err != nil {
		return GiMinorVersionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return GiMinorVersionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return GiMinorVersionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *GiMinorVersionsClient) getCreateRequest(ctx context.Context, location string, giversionname string, giMinorVersionName string, _ *GiMinorVersionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/giVersions/{giversionname}/giMinorVersions/{giMinorVersionName}"
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
	if giMinorVersionName == "" {
		return nil, errors.New("parameter giMinorVersionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{giMinorVersionName}", url.PathEscape(giMinorVersionName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *GiMinorVersionsClient) getHandleResponse(resp *http.Response) (GiMinorVersionsClientGetResponse, error) {
	result := GiMinorVersionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GiMinorVersion); err != nil {
		return GiMinorVersionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByParentPager - List GiMinorVersion resources by GiVersion
//
// Generated from API version 2025-03-01
//   - location - The name of the Azure region.
//   - giversionname - GiVersion name
//   - options - GiMinorVersionsClientListByParentOptions contains the optional parameters for the GiMinorVersionsClient.NewListByParentPager
//     method.
func (client *GiMinorVersionsClient) NewListByParentPager(location string, giversionname string, options *GiMinorVersionsClientListByParentOptions) *runtime.Pager[GiMinorVersionsClientListByParentResponse] {
	return runtime.NewPager(runtime.PagingHandler[GiMinorVersionsClientListByParentResponse]{
		More: func(page GiMinorVersionsClientListByParentResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *GiMinorVersionsClientListByParentResponse) (GiMinorVersionsClientListByParentResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "GiMinorVersionsClient.NewListByParentPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByParentCreateRequest(ctx, location, giversionname, options)
			}, nil)
			if err != nil {
				return GiMinorVersionsClientListByParentResponse{}, err
			}
			return client.listByParentHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByParentCreateRequest creates the ListByParent request.
func (client *GiMinorVersionsClient) listByParentCreateRequest(ctx context.Context, location string, giversionname string, options *GiMinorVersionsClientListByParentOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/giVersions/{giversionname}/giMinorVersions"
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
	reqQP.Set("api-version", "2025-03-01")
	if options != nil && options.ShapeFamily != nil {
		reqQP.Set("shapeFamily", string(*options.ShapeFamily))
	}
	if options != nil && options.Zone != nil {
		reqQP.Set("zone", *options.Zone)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByParentHandleResponse handles the ListByParent response.
func (client *GiMinorVersionsClient) listByParentHandleResponse(resp *http.Response) (GiMinorVersionsClientListByParentResponse, error) {
	result := GiMinorVersionsClientListByParentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.GiMinorVersionListResult); err != nil {
		return GiMinorVersionsClientListByParentResponse{}, err
	}
	return result, nil
}
