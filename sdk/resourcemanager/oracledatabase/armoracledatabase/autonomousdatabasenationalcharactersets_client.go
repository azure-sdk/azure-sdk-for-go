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

// AutonomousDatabaseNationalCharacterSetsClient contains the methods for the AutonomousDatabaseNationalCharacterSets group.
// Don't use this type directly, use NewAutonomousDatabaseNationalCharacterSetsClient() instead.
type AutonomousDatabaseNationalCharacterSetsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewAutonomousDatabaseNationalCharacterSetsClient creates a new instance of AutonomousDatabaseNationalCharacterSetsClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAutonomousDatabaseNationalCharacterSetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*AutonomousDatabaseNationalCharacterSetsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AutonomousDatabaseNationalCharacterSetsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a AutonomousDatabaseNationalCharacterSet
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - location - The name of the Azure region.
//   - adbsncharsetname - AutonomousDatabaseNationalCharacterSets name
//   - options - AutonomousDatabaseNationalCharacterSetsClientGetOptions contains the optional parameters for the AutonomousDatabaseNationalCharacterSetsClient.Get
//     method.
func (client *AutonomousDatabaseNationalCharacterSetsClient) Get(ctx context.Context, location string, adbsncharsetname string, options *AutonomousDatabaseNationalCharacterSetsClientGetOptions) (AutonomousDatabaseNationalCharacterSetsClientGetResponse, error) {
	var err error
	const operationName = "AutonomousDatabaseNationalCharacterSetsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, adbsncharsetname, options)
	if err != nil {
		return AutonomousDatabaseNationalCharacterSetsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AutonomousDatabaseNationalCharacterSetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AutonomousDatabaseNationalCharacterSetsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AutonomousDatabaseNationalCharacterSetsClient) getCreateRequest(ctx context.Context, location string, adbsncharsetname string, options *AutonomousDatabaseNationalCharacterSetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/autonomousDatabaseNationalCharacterSets/{adbsncharsetname}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if adbsncharsetname == "" {
		return nil, errors.New("parameter adbsncharsetname cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{adbsncharsetname}", url.PathEscape(adbsncharsetname))
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
func (client *AutonomousDatabaseNationalCharacterSetsClient) getHandleResponse(resp *http.Response) (AutonomousDatabaseNationalCharacterSetsClientGetResponse, error) {
	result := AutonomousDatabaseNationalCharacterSetsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutonomousDatabaseNationalCharacterSet); err != nil {
		return AutonomousDatabaseNationalCharacterSetsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - List AutonomousDatabaseNationalCharacterSet resources by Location
//
// Generated from API version 2024-06-01
//   - location - The name of the Azure region.
//   - options - AutonomousDatabaseNationalCharacterSetsClientListByLocationOptions contains the optional parameters for the AutonomousDatabaseNationalCharacterSetsClient.NewListByLocationPager
//     method.
func (client *AutonomousDatabaseNationalCharacterSetsClient) NewListByLocationPager(location string, options *AutonomousDatabaseNationalCharacterSetsClientListByLocationOptions) *runtime.Pager[AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse]{
		More: func(page AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse) (AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "AutonomousDatabaseNationalCharacterSetsClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *AutonomousDatabaseNationalCharacterSetsClient) listByLocationCreateRequest(ctx context.Context, location string, options *AutonomousDatabaseNationalCharacterSetsClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/autonomousDatabaseNationalCharacterSets"
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
func (client *AutonomousDatabaseNationalCharacterSetsClient) listByLocationHandleResponse(resp *http.Response) (AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse, error) {
	result := AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AutonomousDatabaseNationalCharacterSetListResult); err != nil {
		return AutonomousDatabaseNationalCharacterSetsClientListByLocationResponse{}, err
	}
	return result, nil
}
