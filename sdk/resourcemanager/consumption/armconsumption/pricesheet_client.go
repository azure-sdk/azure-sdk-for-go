//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PriceSheetClient contains the methods for the PriceSheet group.
// Don't use this type directly, use NewPriceSheetClient() instead.
type PriceSheetClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewPriceSheetClient creates a new instance of PriceSheetClient with the specified values.
//   - subscriptionID - Azure Subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewPriceSheetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*PriceSheetClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &PriceSheetClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginDownloadByBillingAccountPeriod - Generates the pricesheet for the provided billing period asynchronously based on
// the enrollment id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - billingAccountID - BillingAccount ID
//   - billingPeriodName - Billing Period Name.
//   - options - PriceSheetClientBeginDownloadByBillingAccountPeriodOptions contains the optional parameters for the PriceSheetClient.BeginDownloadByBillingAccountPeriod
//     method.
func (client *PriceSheetClient) BeginDownloadByBillingAccountPeriod(ctx context.Context, billingAccountID string, billingPeriodName string, options *PriceSheetClientBeginDownloadByBillingAccountPeriodOptions) (*runtime.Poller[PriceSheetClientDownloadByBillingAccountPeriodResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.downloadByBillingAccountPeriod(ctx, billingAccountID, billingPeriodName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[PriceSheetClientDownloadByBillingAccountPeriodResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[PriceSheetClientDownloadByBillingAccountPeriodResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// DownloadByBillingAccountPeriod - Generates the pricesheet for the provided billing period asynchronously based on the enrollment
// id
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
func (client *PriceSheetClient) downloadByBillingAccountPeriod(ctx context.Context, billingAccountID string, billingPeriodName string, options *PriceSheetClientBeginDownloadByBillingAccountPeriodOptions) (*http.Response, error) {
	var err error
	const operationName = "PriceSheetClient.BeginDownloadByBillingAccountPeriod"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.downloadByBillingAccountPeriodCreateRequest(ctx, billingAccountID, billingPeriodName, options)
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

// downloadByBillingAccountPeriodCreateRequest creates the DownloadByBillingAccountPeriod request.
func (client *PriceSheetClient) downloadByBillingAccountPeriodCreateRequest(ctx context.Context, billingAccountID string, billingPeriodName string, options *PriceSheetClientBeginDownloadByBillingAccountPeriodOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/download"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingPeriodName == "" {
		return nil, errors.New("parameter billingPeriodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingPeriodName}", url.PathEscape(billingPeriodName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the price sheet for a subscription. Price sheet is available via this API only for May 1, 2014 or later.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - options - PriceSheetClientGetOptions contains the optional parameters for the PriceSheetClient.Get method.
func (client *PriceSheetClient) Get(ctx context.Context, options *PriceSheetClientGetOptions) (PriceSheetClientGetResponse, error) {
	var err error
	const operationName = "PriceSheetClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return PriceSheetClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PriceSheetClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PriceSheetClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *PriceSheetClient) getCreateRequest(ctx context.Context, options *PriceSheetClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/pricesheets/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PriceSheetClient) getHandleResponse(resp *http.Response) (PriceSheetClientGetResponse, error) {
	result := PriceSheetClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PriceSheetResult); err != nil {
		return PriceSheetClientGetResponse{}, err
	}
	return result, nil
}

// GetByBillingPeriod - Get the price sheet for a scope by subscriptionId and billing period. Price sheet is available via
// this API only for May 1, 2014 or later.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-08-01
//   - billingPeriodName - Billing Period Name.
//   - options - PriceSheetClientGetByBillingPeriodOptions contains the optional parameters for the PriceSheetClient.GetByBillingPeriod
//     method.
func (client *PriceSheetClient) GetByBillingPeriod(ctx context.Context, billingPeriodName string, options *PriceSheetClientGetByBillingPeriodOptions) (PriceSheetClientGetByBillingPeriodResponse, error) {
	var err error
	const operationName = "PriceSheetClient.GetByBillingPeriod"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByBillingPeriodCreateRequest(ctx, billingPeriodName, options)
	if err != nil {
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	resp, err := client.getByBillingPeriodHandleResponse(httpResp)
	return resp, err
}

// getByBillingPeriodCreateRequest creates the GetByBillingPeriod request.
func (client *PriceSheetClient) getByBillingPeriodCreateRequest(ctx context.Context, billingPeriodName string, options *PriceSheetClientGetByBillingPeriodOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if billingPeriodName == "" {
		return nil, errors.New("parameter billingPeriodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingPeriodName}", url.PathEscape(billingPeriodName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2024-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingPeriodHandleResponse handles the GetByBillingPeriod response.
func (client *PriceSheetClient) getByBillingPeriodHandleResponse(resp *http.Response) (PriceSheetClientGetByBillingPeriodResponse, error) {
	result := PriceSheetClientGetByBillingPeriodResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PriceSheetResult); err != nil {
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	return result, nil
}
