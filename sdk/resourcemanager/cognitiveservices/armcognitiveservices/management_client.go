//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcognitiveservices

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

// ManagementClient contains the methods for the CognitiveServicesManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewManagementClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ManagementClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ManagementClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CheckDomainAvailability - Check whether a domain is available.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - parameters - Check Domain Availability parameter.
//   - options - ManagementClientCheckDomainAvailabilityOptions contains the optional parameters for the ManagementClient.CheckDomainAvailability
//     method.
func (client *ManagementClient) CheckDomainAvailability(ctx context.Context, parameters CheckDomainAvailabilityParameter, options *ManagementClientCheckDomainAvailabilityOptions) (ManagementClientCheckDomainAvailabilityResponse, error) {
	var err error
	const operationName = "ManagementClient.CheckDomainAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkDomainAvailabilityCreateRequest(ctx, parameters, options)
	if err != nil {
		return ManagementClientCheckDomainAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientCheckDomainAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementClientCheckDomainAvailabilityResponse{}, err
	}
	resp, err := client.checkDomainAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkDomainAvailabilityCreateRequest creates the CheckDomainAvailability request.
func (client *ManagementClient) checkDomainAvailabilityCreateRequest(ctx context.Context, parameters CheckDomainAvailabilityParameter, options *ManagementClientCheckDomainAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/checkDomainAvailability"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkDomainAvailabilityHandleResponse handles the CheckDomainAvailability response.
func (client *ManagementClient) checkDomainAvailabilityHandleResponse(resp *http.Response) (ManagementClientCheckDomainAvailabilityResponse, error) {
	result := ManagementClientCheckDomainAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.DomainAvailability); err != nil {
		return ManagementClientCheckDomainAvailabilityResponse{}, err
	}
	return result, nil
}

// CheckSKUAvailability - Check available SKUs.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - location - Resource location.
//   - parameters - Check SKU Availability POST body.
//   - options - ManagementClientCheckSKUAvailabilityOptions contains the optional parameters for the ManagementClient.CheckSKUAvailability
//     method.
func (client *ManagementClient) CheckSKUAvailability(ctx context.Context, location string, parameters CheckSKUAvailabilityParameter, options *ManagementClientCheckSKUAvailabilityOptions) (ManagementClientCheckSKUAvailabilityResponse, error) {
	var err error
	const operationName = "ManagementClient.CheckSKUAvailability"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.checkSKUAvailabilityCreateRequest(ctx, location, parameters, options)
	if err != nil {
		return ManagementClientCheckSKUAvailabilityResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientCheckSKUAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementClientCheckSKUAvailabilityResponse{}, err
	}
	resp, err := client.checkSKUAvailabilityHandleResponse(httpResp)
	return resp, err
}

// checkSKUAvailabilityCreateRequest creates the CheckSKUAvailability request.
func (client *ManagementClient) checkSKUAvailabilityCreateRequest(ctx context.Context, location string, parameters CheckSKUAvailabilityParameter, options *ManagementClientCheckSKUAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.CognitiveServices/locations/{location}/checkSkuAvailability"
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
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// checkSKUAvailabilityHandleResponse handles the CheckSKUAvailability response.
func (client *ManagementClient) checkSKUAvailabilityHandleResponse(resp *http.Response) (ManagementClientCheckSKUAvailabilityResponse, error) {
	result := ManagementClientCheckSKUAvailabilityResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUAvailabilityListResult); err != nil {
		return ManagementClientCheckSKUAvailabilityResponse{}, err
	}
	return result, nil
}
