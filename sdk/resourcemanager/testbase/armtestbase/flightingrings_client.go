//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtestbase

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

// FlightingRingsClient contains the methods for the FlightingRings group.
// Don't use this type directly, use NewFlightingRingsClient() instead.
type FlightingRingsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewFlightingRingsClient creates a new instance of FlightingRingsClient with the specified values.
//   - subscriptionID - The Azure subscription ID. This is a GUID-formatted string.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewFlightingRingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FlightingRingsClient, error) {
	cl, err := arm.NewClient(moduleName+".FlightingRingsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &FlightingRingsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Gets a flighting ring of a Test Base Account.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-04-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - flightingRingResourceName - The resource name of a flighting ring.
//   - options - FlightingRingsClientGetOptions contains the optional parameters for the FlightingRingsClient.Get method.
func (client *FlightingRingsClient) Get(ctx context.Context, resourceGroupName string, testBaseAccountName string, flightingRingResourceName string, options *FlightingRingsClientGetOptions) (FlightingRingsClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, resourceGroupName, testBaseAccountName, flightingRingResourceName, options)
	if err != nil {
		return FlightingRingsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FlightingRingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return FlightingRingsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *FlightingRingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, flightingRingResourceName string, options *FlightingRingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/flightingRings/{flightingRingResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	if flightingRingResourceName == "" {
		return nil, errors.New("parameter flightingRingResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{flightingRingResourceName}", url.PathEscape(flightingRingResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FlightingRingsClient) getHandleResponse(resp *http.Response) (FlightingRingsClientGetResponse, error) {
	result := FlightingRingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlightingRingResource); err != nil {
		return FlightingRingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the flighting rings of a Test Base Account.
//
// Generated from API version 2022-04-01-preview
//   - resourceGroupName - The name of the resource group that contains the resource.
//   - testBaseAccountName - The resource name of the Test Base Account.
//   - options - FlightingRingsClientListOptions contains the optional parameters for the FlightingRingsClient.NewListPager method.
func (client *FlightingRingsClient) NewListPager(resourceGroupName string, testBaseAccountName string, options *FlightingRingsClientListOptions) *runtime.Pager[FlightingRingsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FlightingRingsClientListResponse]{
		More: func(page FlightingRingsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FlightingRingsClientListResponse) (FlightingRingsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, testBaseAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FlightingRingsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return FlightingRingsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FlightingRingsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FlightingRingsClient) listCreateRequest(ctx context.Context, resourceGroupName string, testBaseAccountName string, options *FlightingRingsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.TestBase/testBaseAccounts/{testBaseAccountName}/flightingRings"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if testBaseAccountName == "" {
		return nil, errors.New("parameter testBaseAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{testBaseAccountName}", url.PathEscape(testBaseAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FlightingRingsClient) listHandleResponse(resp *http.Response) (FlightingRingsClientListResponse, error) {
	result := FlightingRingsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FlightingRingListResult); err != nil {
		return FlightingRingsClientListResponse{}, err
	}
	return result, nil
}
