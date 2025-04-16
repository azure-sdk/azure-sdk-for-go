// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armbilling

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

// AvailableBalancesClient contains the methods for the AvailableBalances group.
// Don't use this type directly, use NewAvailableBalancesClient() instead.
type AvailableBalancesClient struct {
	internal *arm.Client
}

// NewAvailableBalancesClient creates a new instance of AvailableBalancesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAvailableBalancesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AvailableBalancesClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AvailableBalancesClient{
		internal: cl,
	}
	return client, nil
}

// GetByBillingAccount - The Available Credit or Payment on Account Balance for a billing account. The credit balance can
// be used to settle due or past due invoices and is supported for billing accounts with agreement type
// Microsoft Customer Agreement. The payment on account balance is supported for billing accounts with agreement type Microsoft
// Customer Agreement or Microsoft Online Services Program.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - options - AvailableBalancesClientGetByBillingAccountOptions contains the optional parameters for the AvailableBalancesClient.GetByBillingAccount
//     method.
func (client *AvailableBalancesClient) GetByBillingAccount(ctx context.Context, billingAccountName string, options *AvailableBalancesClientGetByBillingAccountOptions) (AvailableBalancesClientGetByBillingAccountResponse, error) {
	var err error
	const operationName = "AvailableBalancesClient.GetByBillingAccount"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByBillingAccountCreateRequest(ctx, billingAccountName, options)
	if err != nil {
		return AvailableBalancesClientGetByBillingAccountResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailableBalancesClientGetByBillingAccountResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AvailableBalancesClientGetByBillingAccountResponse{}, err
	}
	resp, err := client.getByBillingAccountHandleResponse(httpResp)
	return resp, err
}

// getByBillingAccountCreateRequest creates the GetByBillingAccount request.
func (client *AvailableBalancesClient) getByBillingAccountCreateRequest(ctx context.Context, billingAccountName string, _ *AvailableBalancesClientGetByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/availableBalance/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingAccountHandleResponse handles the GetByBillingAccount response.
func (client *AvailableBalancesClient) getByBillingAccountHandleResponse(resp *http.Response) (AvailableBalancesClientGetByBillingAccountResponse, error) {
	result := AvailableBalancesClientGetByBillingAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableBalance); err != nil {
		return AvailableBalancesClientGetByBillingAccountResponse{}, err
	}
	return result, nil
}

// GetByBillingProfile - The Available Credit or Payment on Account Balance for a billing profile. The credit balance can
// be used to settle due or past due invoices and is supported for billing accounts with agreement type
// Microsoft Customer Agreement. The payment on account balance is supported for billing accounts with agreement type Microsoft
// Customer Agreement.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - billingAccountName - The ID that uniquely identifies a billing account.
//   - billingProfileName - The ID that uniquely identifies a billing profile.
//   - options - AvailableBalancesClientGetByBillingProfileOptions contains the optional parameters for the AvailableBalancesClient.GetByBillingProfile
//     method.
func (client *AvailableBalancesClient) GetByBillingProfile(ctx context.Context, billingAccountName string, billingProfileName string, options *AvailableBalancesClientGetByBillingProfileOptions) (AvailableBalancesClientGetByBillingProfileResponse, error) {
	var err error
	const operationName = "AvailableBalancesClient.GetByBillingProfile"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getByBillingProfileCreateRequest(ctx, billingAccountName, billingProfileName, options)
	if err != nil {
		return AvailableBalancesClientGetByBillingProfileResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AvailableBalancesClientGetByBillingProfileResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AvailableBalancesClientGetByBillingProfileResponse{}, err
	}
	resp, err := client.getByBillingProfileHandleResponse(httpResp)
	return resp, err
}

// getByBillingProfileCreateRequest creates the GetByBillingProfile request.
func (client *AvailableBalancesClient) getByBillingProfileCreateRequest(ctx context.Context, billingAccountName string, billingProfileName string, _ *AvailableBalancesClientGetByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/availableBalance/default"
	if billingAccountName == "" {
		return nil, errors.New("parameter billingAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountName}", url.PathEscape(billingAccountName))
	if billingProfileName == "" {
		return nil, errors.New("parameter billingProfileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileName}", url.PathEscape(billingProfileName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getByBillingProfileHandleResponse handles the GetByBillingProfile response.
func (client *AvailableBalancesClient) getByBillingProfileHandleResponse(resp *http.Response) (AvailableBalancesClientGetByBillingProfileResponse, error) {
	result := AvailableBalancesClientGetByBillingProfileResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AvailableBalance); err != nil {
		return AvailableBalancesClientGetByBillingProfileResponse{}, err
	}
	return result, nil
}
