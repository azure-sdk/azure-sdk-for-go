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

// AutonomousDatabaseCharacterSetsClient contains the methods for the AutonomousDatabaseCharacterSets group.
// Don't use this type directly, use NewAutonomousDatabaseCharacterSetsClient() instead.
type AutonomousDatabaseCharacterSetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAutonomousDatabaseCharacterSetsClient creates a new instance of AutonomousDatabaseCharacterSetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAutonomousDatabaseCharacterSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AutonomousDatabaseCharacterSetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AutonomousDatabaseCharacterSetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AutonomousDatabaseCharacterSet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01-preview
//   - location - The name of the Azure region.
//   - adbscharsetname - AutonomousDatabaseCharacterSet name
//   - options - AutonomousDatabaseCharacterSetsClientGetOptions contains the optional parameters for the AutonomousDatabaseCharacterSetsClient.Get
//     method.
func (client *AutonomousDatabaseCharacterSetsClient) Get(ctx context.Context, location string, adbscharsetname string, options *AutonomousDatabaseCharacterSetsClientGetOptions) (AutonomousDatabaseCharacterSetsClientGetResponse, error) {
	var err error
	const operationName = "AutonomousDatabaseCharacterSetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, adbscharsetname, options)
	if err != nil {
		return AutonomousDatabaseCharacterSetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutonomousDatabaseCharacterSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AutonomousDatabaseCharacterSetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AutonomousDatabaseCharacterSetsClient) getCreateRequest(ctx context.Context, location string, adbscharsetname string, options *AutonomousDatabaseCharacterSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/autonomousDatabaseCharacterSets/{adbscharsetname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if adbscharsetname == "" {
		return nil, errors.New("parameter adbscharsetname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{adbscharsetname}", url.PathEscape(adbscharsetname))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AutonomousDatabaseCharacterSetsClient) getHandleResponse(resp *http.Response) (AutonomousDatabaseCharacterSetsClientGetResponse, error) {
	result := AutonomousDatabaseCharacterSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutonomousDatabaseCharacterSet); err != nil {
		return AutonomousDatabaseCharacterSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - List AutonomousDatabaseCharacterSet resources by Location
//
// Generated from API version 2023-09-01-preview
//   - location - The name of the Azure region.
//   - options - AutonomousDatabaseCharacterSetsClientListByLocationOptions contains the optional parameters for the AutonomousDatabaseCharacterSetsClient.NewListByLocationPager
//     method.
func (client *AutonomousDatabaseCharacterSetsClient) NewListByLocationPager(location string, options *AutonomousDatabaseCharacterSetsClientListByLocationOptions) *runtime.Pager[AutonomousDatabaseCharacterSetsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[AutonomousDatabaseCharacterSetsClientListByLocationResponse]{
		More: func(page AutonomousDatabaseCharacterSetsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AutonomousDatabaseCharacterSetsClientListByLocationResponse) (AutonomousDatabaseCharacterSetsClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AutonomousDatabaseCharacterSetsClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return AutonomousDatabaseCharacterSetsClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *AutonomousDatabaseCharacterSetsClient) listByLocationCreateRequest(ctx context.Context, location string, options *AutonomousDatabaseCharacterSetsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/autonomousDatabaseCharacterSets"
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
	reqQP.Set("api-version", "2023-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *AutonomousDatabaseCharacterSetsClient) listByLocationHandleResponse(resp *http.Response) (AutonomousDatabaseCharacterSetsClientListByLocationResponse, error) {
	result := AutonomousDatabaseCharacterSetsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutonomousDatabaseCharacterSetListResult); err != nil {
		return AutonomousDatabaseCharacterSetsClientListByLocationResponse{}, err
	}
	return result, nil
}
