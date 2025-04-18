// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armproviderhub

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

// Client contains the methods for the ProviderHub group.
// Don't use this type directly, use NewClient() instead.
type Client struct {
	internal       *arm.Client
	subscriptionID string
}

// NewClient creates a new instance of Client with the specified values.
//   - subscriptionID - The ID of the target subscription.
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

// CheckinManifest - Checkin the manifest.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - checkinManifestParams - The required body parameters supplied to the checkin manifest operation.
//   - options - ClientCheckinManifestOptions contains the optional parameters for the Client.CheckinManifest method.
func (client *Client) CheckinManifest(ctx context.Context, providerNamespace string, checkinManifestParams CheckinManifestParams, options *ClientCheckinManifestOptions) (ClientCheckinManifestResponse, error) {
	var err error
	const operationName = "Client.CheckinManifest"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkinManifestCreateRequest(ctx, providerNamespace, checkinManifestParams, options)
	if err != nil {
		return ClientCheckinManifestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCheckinManifestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientCheckinManifestResponse{}, err
	}
	resp, err := client.checkinManifestHandleResponse(httpResp)
	return resp, err
}

// checkinManifestCreateRequest creates the CheckinManifest request.
func (client *Client) checkinManifestCreateRequest(ctx context.Context, providerNamespace string, checkinManifestParams CheckinManifestParams, _ *ClientCheckinManifestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/checkinManifest"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, checkinManifestParams); err != nil {
		return nil, err
	}
	return req, nil
}

// checkinManifestHandleResponse handles the CheckinManifest response.
func (client *Client) checkinManifestHandleResponse(resp *http.Response) (ClientCheckinManifestResponse, error) {
	result := ClientCheckinManifestResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckinManifestInfo); err != nil {
		return ClientCheckinManifestResponse{}, err
	}
	return result, nil
}

// GenerateManifest - Generates the manifest for the given provider.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2020-11-20
//   - providerNamespace - The name of the resource provider hosted within ProviderHub.
//   - options - ClientGenerateManifestOptions contains the optional parameters for the Client.GenerateManifest method.
func (client *Client) GenerateManifest(ctx context.Context, providerNamespace string, options *ClientGenerateManifestOptions) (ClientGenerateManifestResponse, error) {
	var err error
	const operationName = "Client.GenerateManifest"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.generateManifestCreateRequest(ctx, providerNamespace, options)
	if err != nil {
		return ClientGenerateManifestResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGenerateManifestResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGenerateManifestResponse{}, err
	}
	resp, err := client.generateManifestHandleResponse(httpResp)
	return resp, err
}

// generateManifestCreateRequest creates the GenerateManifest request.
func (client *Client) generateManifestCreateRequest(ctx context.Context, providerNamespace string, _ *ClientGenerateManifestOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ProviderHub/providerRegistrations/{providerNamespace}/generateManifest"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if providerNamespace == "" {
		return nil, errors.New("parameter providerNamespace cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{providerNamespace}", url.PathEscape(providerNamespace))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-20")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// generateManifestHandleResponse handles the GenerateManifest response.
func (client *Client) generateManifestHandleResponse(resp *http.Response) (ClientGenerateManifestResponse, error) {
	result := ClientGenerateManifestResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ResourceProviderManifest); err != nil {
		return ClientGenerateManifestResponse{}, err
	}
	return result, nil
}
