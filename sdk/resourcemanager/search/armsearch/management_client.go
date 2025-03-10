// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsearch

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

// ManagementClient contains the methods for the SearchManagementClient group.
// Don't use this type directly, use NewManagementClient() instead.
type ManagementClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewManagementClient creates a new instance of ManagementClient with the specified values.
//   - subscriptionID - The unique identifier for a Microsoft Azure subscription. You can obtain this value from the Azure Resource
//     Manager API, command line tools, or the portal.
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

// UsageBySubscriptionSKU - Gets the quota usage for a search SKU in the given subscription.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2025-05-01
//   - location - The unique location name for a Microsoft Azure geographic region.
//   - skuName - The unique SKU name that identifies a billable tier.
//   - SearchManagementRequestOptions - SearchManagementRequestOptions contains a group of parameters for the AdminKeysClient.Get
//     method.
//   - options - ManagementClientUsageBySubscriptionSKUOptions contains the optional parameters for the ManagementClient.UsageBySubscriptionSKU
//     method.
func (client *ManagementClient) UsageBySubscriptionSKU(ctx context.Context, location string, skuName string, searchManagementRequestOptions *SearchManagementRequestOptions, options *ManagementClientUsageBySubscriptionSKUOptions) (ManagementClientUsageBySubscriptionSKUResponse, error) {
	var err error
	const operationName = "ManagementClient.UsageBySubscriptionSKU"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.usageBySubscriptionSKUCreateRequest(ctx, location, skuName, searchManagementRequestOptions, options)
	if err != nil {
		return ManagementClientUsageBySubscriptionSKUResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ManagementClientUsageBySubscriptionSKUResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ManagementClientUsageBySubscriptionSKUResponse{}, err
	}
	resp, err := client.usageBySubscriptionSKUHandleResponse(httpResp)
	return resp, err
}

// usageBySubscriptionSKUCreateRequest creates the UsageBySubscriptionSKU request.
func (client *ManagementClient) usageBySubscriptionSKUCreateRequest(ctx context.Context, location string, skuName string, searchManagementRequestOptions *SearchManagementRequestOptions, _ *ManagementClientUsageBySubscriptionSKUOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Search/locations/{location}/usages/{skuName}"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2025-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if searchManagementRequestOptions != nil && searchManagementRequestOptions.ClientRequestID != nil {
		req.Raw().Header["x-ms-client-request-id"] = []string{*searchManagementRequestOptions.ClientRequestID}
	}
	return req, nil
}

// usageBySubscriptionSKUHandleResponse handles the UsageBySubscriptionSKU response.
func (client *ManagementClient) usageBySubscriptionSKUHandleResponse(resp *http.Response) (ManagementClientUsageBySubscriptionSKUResponse, error) {
	result := ManagementClientUsageBySubscriptionSKUResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.QuotaUsageResult); err != nil {
		return ManagementClientUsageBySubscriptionSKUResponse{}, err
	}
	return result, nil
}
