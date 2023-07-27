//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappcontainers

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

// BillingMetersClient contains the methods for the BillingMeters group.
// Don't use this type directly, use NewBillingMetersClient() instead.
type BillingMetersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBillingMetersClient creates a new instance of BillingMetersClient with the specified values.
//   - subscriptionID - The ID of the target subscription.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBillingMetersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BillingMetersClient, error) {
	cl, err := arm.NewClient(moduleName+".BillingMetersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BillingMetersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get all billingMeters for a location.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-04-01-preview
//   - location - The name of Azure region.
//   - options - BillingMetersClientGetOptions contains the optional parameters for the BillingMetersClient.Get method.
func (client *BillingMetersClient) Get(ctx context.Context, location string, options *BillingMetersClientGetOptions) (BillingMetersClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, location, options)
	if err != nil {
		return BillingMetersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BillingMetersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return BillingMetersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *BillingMetersClient) getCreateRequest(ctx context.Context, location string, options *BillingMetersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.App/locations/{location}/billingMeters"
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
	reqQP.Set("api-version", "2023-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *BillingMetersClient) getHandleResponse(resp *http.Response) (BillingMetersClientGetResponse, error) {
	result := BillingMetersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BillingMeterCollection); err != nil {
		return BillingMetersClientGetResponse{}, err
	}
	return result, nil
}
