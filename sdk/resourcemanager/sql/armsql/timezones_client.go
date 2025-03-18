// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// TimeZonesClient contains the methods for the TimeZones group.
// Don't use this type directly, use NewTimeZonesClient() instead.
type TimeZonesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewTimeZonesClient creates a new instance of TimeZonesClient with the specified values.
//   - subscriptionID - The subscription ID that identifies an Azure subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewTimeZonesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*TimeZonesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &TimeZonesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a managed instance time zone.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-01-preview
//   - options - TimeZonesClientGetOptions contains the optional parameters for the TimeZonesClient.Get method.
func (client *TimeZonesClient) Get(ctx context.Context, locationName string, timeZoneID string, options *TimeZonesClientGetOptions) (TimeZonesClientGetResponse, error) {
	var err error
	const operationName = "TimeZonesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, locationName, timeZoneID, options)
	if err != nil {
		return TimeZonesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return TimeZonesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return TimeZonesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *TimeZonesClient) getCreateRequest(ctx context.Context, locationName string, timeZoneID string, _ *TimeZonesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/timeZones/{timeZoneId}"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if timeZoneID == "" {
		return nil, errors.New("parameter timeZoneID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{timeZoneId}", url.PathEscape(timeZoneID))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TimeZonesClient) getHandleResponse(resp *http.Response) (TimeZonesClientGetResponse, error) {
	result := TimeZonesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeZone); err != nil {
		return TimeZonesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - Gets a list of managed instance time zones by location.
//
// Generated from API version 2020-11-01-preview
//   - options - TimeZonesClientListByLocationOptions contains the optional parameters for the TimeZonesClient.NewListByLocationPager
//     method.
func (client *TimeZonesClient) NewListByLocationPager(locationName string, options *TimeZonesClientListByLocationOptions) *runtime.Pager[TimeZonesClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[TimeZonesClientListByLocationResponse]{
		More: func(page TimeZonesClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *TimeZonesClientListByLocationResponse) (TimeZonesClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "TimeZonesClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, locationName, options)
			}, nil)
			if err != nil {
				return TimeZonesClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *TimeZonesClient) listByLocationCreateRequest(ctx context.Context, locationName string, _ *TimeZonesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Sql/locations/{locationName}/timeZones"
	if locationName == "" {
		return nil, errors.New("parameter locationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{locationName}", url.PathEscape(locationName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByLocationHandleResponse handles the ListByLocation response.
func (client *TimeZonesClient) listByLocationHandleResponse(resp *http.Response) (TimeZonesClientListByLocationResponse, error) {
	result := TimeZonesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TimeZoneListResult); err != nil {
		return TimeZonesClientListByLocationResponse{}, err
	}
	return result, nil
}
