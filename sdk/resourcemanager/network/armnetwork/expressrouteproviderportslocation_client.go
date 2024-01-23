//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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

// ExpressRouteProviderPortsLocationClient contains the methods for the ExpressRouteProviderPortsLocation group.
// Don't use this type directly, use NewExpressRouteProviderPortsLocationClient() instead.
type ExpressRouteProviderPortsLocationClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewExpressRouteProviderPortsLocationClient creates a new instance of ExpressRouteProviderPortsLocationClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewExpressRouteProviderPortsLocationClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ExpressRouteProviderPortsLocationClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ExpressRouteProviderPortsLocationClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// List - Retrieves all the ExpressRouteProviderPorts in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-11-01
//   - options - ExpressRouteProviderPortsLocationClientListOptions contains the optional parameters for the ExpressRouteProviderPortsLocationClient.List
//     method.
func (client *ExpressRouteProviderPortsLocationClient) List(ctx context.Context, options *ExpressRouteProviderPortsLocationClientListOptions) (ExpressRouteProviderPortsLocationClientListResponse, error) {
	var err error
	const operationName = "ExpressRouteProviderPortsLocationClient.List"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	resp, err := client.listHandleResponse(httpResp)
	return resp, err
}

// listCreateRequest creates the List request.
func (client *ExpressRouteProviderPortsLocationClient) listCreateRequest(ctx context.Context, options *ExpressRouteProviderPortsLocationClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/expressRouteProviderPorts"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-11-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ExpressRouteProviderPortsLocationClient) listHandleResponse(resp *http.Response) (ExpressRouteProviderPortsLocationClientListResponse, error) {
	result := ExpressRouteProviderPortsLocationClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ExpressRouteProviderPortListResult); err != nil {
		return ExpressRouteProviderPortsLocationClientListResponse{}, err
	}
	return result, nil
}
