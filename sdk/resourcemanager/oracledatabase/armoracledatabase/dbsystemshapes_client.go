// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) Go Code Generator. DO NOT EDIT.

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

// DbSystemShapesClient contains the methods for the DbSystemShapes group.
// Don't use this type directly, use NewDbSystemShapesClient() instead.
type DbSystemShapesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDbSystemShapesClient creates a new instance of DbSystemShapesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDbSystemShapesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DbSystemShapesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DbSystemShapesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a DbSystemShape
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-03-01
//   - location - The name of the Azure region.
//   - dbsystemshapename - DbSystemShape name
//   - options - DbSystemShapesClientGetOptions contains the optional parameters for the DbSystemShapesClient.Get method.
func (client *DbSystemShapesClient) Get(ctx context.Context, location string, dbsystemshapename string, options *DbSystemShapesClientGetOptions) (DbSystemShapesClientGetResponse, error) {
	var err error
	const operationName = "DbSystemShapesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, dbsystemshapename, options)
	if err != nil {
		return DbSystemShapesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DbSystemShapesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DbSystemShapesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DbSystemShapesClient) getCreateRequest(ctx context.Context, location string, dbsystemshapename string, _ *DbSystemShapesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/dbSystemShapes/{dbsystemshapename}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if dbsystemshapename == "" {
		return nil, errors.New("parameter dbsystemshapename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dbsystemshapename}", url.PathEscape(dbsystemshapename))
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
func (client *DbSystemShapesClient) getHandleResponse(resp *http.Response) (DbSystemShapesClientGetResponse, error) {
	result := DbSystemShapesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DbSystemShape); err != nil {
		return DbSystemShapesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - List DbSystemShape resources by SubscriptionLocationResource
//
// Generated from API version 2025-03-01
//   - location - The name of the Azure region.
//   - options - DbSystemShapesClientListByLocationOptions contains the optional parameters for the DbSystemShapesClient.NewListByLocationPager
//     method.
func (client *DbSystemShapesClient) NewListByLocationPager(location string, options *DbSystemShapesClientListByLocationOptions) *runtime.Pager[DbSystemShapesClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[DbSystemShapesClientListByLocationResponse]{
		More: func(page DbSystemShapesClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DbSystemShapesClientListByLocationResponse) (DbSystemShapesClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DbSystemShapesClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return DbSystemShapesClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *DbSystemShapesClient) listByLocationCreateRequest(ctx context.Context, location string, options *DbSystemShapesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/dbSystemShapes"
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
	reqQP.Set("api-version", "2025-03-01")
	if options != nil && options.Zone != nil {
		reqQP.Set("zone", *options.Zone)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *DbSystemShapesClient) listByLocationHandleResponse(resp *http.Response) (DbSystemShapesClientListByLocationResponse, error) {
	result := DbSystemShapesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DbSystemShapeListResult); err != nil {
		return DbSystemShapesClientListByLocationResponse{}, err
	}
	return result, nil
}
