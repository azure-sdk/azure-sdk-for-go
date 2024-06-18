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

// DNSPrivateZonesClient contains the methods for the DNSPrivateZones group.
// Don't use this type directly, use NewDNSPrivateZonesClient() instead.
type DNSPrivateZonesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewDNSPrivateZonesClient creates a new instance of DNSPrivateZonesClient with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewDNSPrivateZonesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*DNSPrivateZonesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &DNSPrivateZonesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get a DnsPrivateZone
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-09-01
//   - location - The name of the Azure region.
//   - dnsprivatezonename - DnsPrivateZone name
//   - options - DNSPrivateZonesClientGetOptions contains the optional parameters for the DNSPrivateZonesClient.Get method.
func (client *DNSPrivateZonesClient) Get(ctx context.Context, location string, dnsprivatezonename string, options *DNSPrivateZonesClientGetOptions) (DNSPrivateZonesClientGetResponse, error) {
	var err error
	const operationName = "DNSPrivateZonesClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, dnsprivatezonename, options)
	if err != nil {
		return DNSPrivateZonesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DNSPrivateZonesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return DNSPrivateZonesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *DNSPrivateZonesClient) getCreateRequest(ctx context.Context, location string, dnsprivatezonename string, options *DNSPrivateZonesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/dnsPrivateZones/{dnsprivatezonename}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if dnsprivatezonename == "" {
		return nil, errors.New("parameter dnsprivatezonename cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dnsprivatezonename}", url.PathEscape(dnsprivatezonename))
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
func (client *DNSPrivateZonesClient) getHandleResponse(resp *http.Response) (DNSPrivateZonesClientGetResponse, error) {
	result := DNSPrivateZonesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSPrivateZone); err != nil {
		return DNSPrivateZonesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByLocationPager - List DnsPrivateZone resources by Location
//
// Generated from API version 2023-09-01
//   - location - The name of the Azure region.
//   - options - DNSPrivateZonesClientListByLocationOptions contains the optional parameters for the DNSPrivateZonesClient.NewListByLocationPager
//     method.
func (client *DNSPrivateZonesClient) NewListByLocationPager(location string, options *DNSPrivateZonesClientListByLocationOptions) *runtime.Pager[DNSPrivateZonesClientListByLocationResponse] {
	return runtime.NewPager(runtime.PagingHandler[DNSPrivateZonesClientListByLocationResponse]{
		More: func(page DNSPrivateZonesClientListByLocationResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *DNSPrivateZonesClientListByLocationResponse) (DNSPrivateZonesClientListByLocationResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "DNSPrivateZonesClient.NewListByLocationPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByLocationCreateRequest(ctx, location, options)
			}, nil)
			if err != nil {
				return DNSPrivateZonesClientListByLocationResponse{}, err
			}
			return client.listByLocationHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByLocationCreateRequest creates the ListByLocation request.
func (client *DNSPrivateZonesClient) listByLocationCreateRequest(ctx context.Context, location string, options *DNSPrivateZonesClientListByLocationOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Oracle.Database/locations/{location}/dnsPrivateZones"
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
func (client *DNSPrivateZonesClient) listByLocationHandleResponse(resp *http.Response) (DNSPrivateZonesClientListByLocationResponse, error) {
	result := DNSPrivateZonesClientListByLocationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DNSPrivateZoneListResult); err != nil {
		return DNSPrivateZonesClientListByLocationResponse{}, err
	}
	return result, nil
}
