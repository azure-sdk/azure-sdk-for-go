//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armrecoveryservices

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

// Client contains the methods for the RecoveryServices group.
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
	cl, err := arm.NewClient(moduleName+".Client", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &Client{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Capabilities - API to get details about capabilities provided by Microsoft.RecoveryServices RP
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - location - Location of the resource
//   - input - Contains information about Resource type and properties to get capabilities
//   - options - ClientCapabilitiesOptions contains the optional parameters for the Client.Capabilities method.
func (client *Client) Capabilities(ctx context.Context, location string, input ResourceCapabilities, options *ClientCapabilitiesOptions) (ClientCapabilitiesResponse, error) {
	var err error
	req, err := client.capabilitiesCreateRequest(ctx, location, input, options)
	if err != nil {
		return ClientCapabilitiesResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCapabilitiesResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientCapabilitiesResponse{}, err
	}
	resp, err := client.capabilitiesHandleResponse(httpResp)
	return resp, err
}

// capabilitiesCreateRequest creates the Capabilities request.
func (client *Client) capabilitiesCreateRequest(ctx context.Context, location string, input ResourceCapabilities, options *ClientCapabilitiesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.RecoveryServices/locations/{location}/capabilities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// capabilitiesHandleResponse handles the Capabilities response.
func (client *Client) capabilitiesHandleResponse(resp *http.Response) (ClientCapabilitiesResponse, error) {
	result := ClientCapabilitiesResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CapabilitiesResponse); err != nil {
		return ClientCapabilitiesResponse{}, err
	}
	return result, nil
}

// CheckNameAvailability - API to check for resource name availability. A name is available if no other resource exists that
// has the same SubscriptionId, Resource Name and Type or if one or more such resources exist, each of
// these must be GC'd and their time of deletion be more than 24 Hours Ago
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-06-01
//   - resourceGroupName - The name of the resource group. The name is case insensitive.
//   - location - Location of the resource
//   - input - Contains information about Resource type and Resource name
//   - options - ClientCheckNameAvailabilityOptions contains the optional parameters for the Client.CheckNameAvailability method.
func (client *Client) CheckNameAvailability(ctx context.Context, resourceGroupName string, location string, input CheckNameAvailabilityParameters, options *ClientCheckNameAvailabilityOptions) (ClientCheckNameAvailabilityResponse, error) {
	var err error
	req, err := client.checkNameAvailabilityCreateRequest(ctx, resourceGroupName, location, input, options)
	if err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.checkNameAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *Client) checkNameAvailabilityCreateRequest(ctx context.Context, resourceGroupName string, location string, input CheckNameAvailabilityParameters, options *ClientCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.RecoveryServices/locations/{location}/checkNameAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-06-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *Client) checkNameAvailabilityHandleResponse(resp *http.Response) (ClientCheckNameAvailabilityResponse, error) {
	result := ClientCheckNameAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckNameAvailabilityResult); err != nil {
		return ClientCheckNameAvailabilityResponse{}, err
	}
	return result, nil
}
