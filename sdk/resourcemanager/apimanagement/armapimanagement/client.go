//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapimanagement

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

// Client contains the methods for the APIManagementClient group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - The ID of the target subscription. The value must be an UUID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*Client, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginPerformConnectivityCheckAsync - Performs a connectivity check between the API Management service and a given destination,
// and returns metrics for the connection, as well as errors encountered while trying to establish it.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - serviceName - The name of the API Management service.
//   - connectivityCheckRequestParams - Connectivity Check request parameters.
//   - options - ClientBeginPerformConnectivityCheckAsyncOptions contains the optional parameters for the Client.BeginPerformConnectivityCheckAsync
//     method.
func (client *Client) BeginPerformConnectivityCheckAsync(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest, options *ClientBeginPerformConnectivityCheckAsyncOptions) (*runtime.Poller[ClientPerformConnectivityCheckAsyncResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.performConnectivityCheckAsync(ctx, resourceGroupName, serviceName, connectivityCheckRequestParams, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ClientPerformConnectivityCheckAsyncResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[ClientPerformConnectivityCheckAsyncResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// PerformConnectivityCheckAsync - Performs a connectivity check between the API Management service and a given destination,
// and returns metrics for the connection, as well as errors encountered while trying to establish it.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-06-01
func (client *Client) performConnectivityCheckAsync(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest, options *ClientBeginPerformConnectivityCheckAsyncOptions) (*http.Response, error) {
	var err error
	const operationName = "Client.BeginPerformConnectivityCheckAsync"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.performConnectivityCheckAsyncCreateRequest(ctx, resourceGroupName, serviceName, connectivityCheckRequestParams, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// performConnectivityCheckAsyncCreateRequest creates the PerformConnectivityCheckAsync request.
func (client *Client) performConnectivityCheckAsyncCreateRequest(ctx context.Context, resourceGroupName string, serviceName string, connectivityCheckRequestParams ConnectivityCheckRequest, options *ClientBeginPerformConnectivityCheckAsyncOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/connectivityCheck"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serviceName == "" {
		return nil, errors.New("parameter serviceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serviceName}", url.PathEscape(serviceName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectivityCheckRequestParams); err != nil {
		return nil, err
	}
	return req, nil
}
